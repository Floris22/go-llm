package groq

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func SplitAudio(audioPath string, chunkSize int) ([]string, string, error) {
	// determine the number of chunks so that each chunk is no larger than 20MB
	bytesPerSec := 16000 * 1 * 2 // sampleRate * channels * bytesPerSample (16-bit)
	segmentSecs := int(math.Max(1, math.Floor(0.95*float64(chunkSize)/float64(bytesPerSec))))

	// create a temporary directory to store the chunks
	tempDir, err := os.MkdirTemp("", "audio-chunks")
	if err != nil {
		return nil, "", fmt.Errorf("failed to create temporary directory: %w", err)
	}

	cmd := exec.Command(
		"ffmpeg", "-hide_banner", "-loglevel", "error", "-y",
		"-i", audioPath,
		"-map", "0:a:0",
		"-c:a", "pcm_s16le", // 16-bit PCM
		"-ar", "16000",
		"-ac", "1",
		"-f", "segment",
		"-segment_time", strconv.Itoa(segmentSecs),
		"-reset_timestamps", "1",
		filepath.Join(tempDir, "chunk_%03d.wav"),
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, tempDir, fmt.Errorf("failed to split audio: %w. Output: %s", err, string(output))
	}

	files, err := os.ReadDir(tempDir)
	if err != nil {
		return nil, tempDir, fmt.Errorf("failed to read temporary directory: %w", err)
	}

	var chunkPaths []string
	for _, file := range files {
		chunkPaths = append(chunkPaths, filepath.Join(tempDir, file.Name()))
	}

	return chunkPaths, tempDir, nil
}
