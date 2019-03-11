// From https://github.com/iawia002/annie/blob/master/utils/ffmpeg.go
package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// MergeToMP4 merge video parts to MP4
func MergeToMP4(paths []string, mergedFilePath string, filename string) error {
	mergeFilePath := filename + ".txt" // merge list file should be in the current directory

	// write ffmpeg input file list
	mergeFile, _ := os.Create(mergeFilePath)
	for _, path := range paths {
		mergeFile.Write([]byte(fmt.Sprintf("file '%s'\n", path)))
	}
	mergeFile.Close()

	cmd := exec.Command(
		"ffmpeg", "-y", "-f", "concat", "-safe", "-1",
		"-i", mergeFilePath, "-c", "copy", "-bsf:a", "aac_adtstoasc", mergedFilePath,
	)
	return runMergeCmd(cmd, paths, mergeFilePath)
}

func runMergeCmd(cmd *exec.Cmd, paths []string, mergeFilePath string) error {
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%s\n%s", err, stderr.String())
	}

	if mergeFilePath != "" {
		os.Remove(mergeFilePath)
	}
	// remove parts
	for _, path := range paths {
		os.Remove(path)
	}
	return nil
}
