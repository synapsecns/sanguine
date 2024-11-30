package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/schollz/sqlite3dump"
)

func main() {
	err := func() (err error) {
		if len(os.Args) < 2 {
			err = fmt.Errorf("incorrect usage")
			return
		}
		f := bufio.NewWriter(os.Stdout)
		err = sqlite3dump.Dump(os.Args[1], f)
		f.Flush()
		return
	}()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		fmt.Fprintf(os.Stderr, "usage: sqlite3dump database.db > database.sql\n")
	} else {
		_, fname := filepath.Split(os.Args[1])
		fmt.Fprintf(os.Stderr, "dumped %s\n", fname)
	}
}
