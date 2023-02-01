package detector

import (
	"github.com/go-git/go-git/v5"
	"os"
	"testing"
)

func TestDetector(t *testing.T) {
	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})

}
