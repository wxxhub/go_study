package mp

import (
	"fmt"
	"testing"
)

func PlayTest(t *testing.T) {
	fmt.Println("library mp test.")

	Play("test", "MP3")
	Play("test", "WAV")
	Play("test", "MP4")
}