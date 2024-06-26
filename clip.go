package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "clip.txt"
	info, err := os.Stdin.Stat()

	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	if info.Mode()&os.ModeCharDevice == 0 {
		// read from stdin and write to clip
		file, err := os.Create(fileName)

		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}

		defer file.Close()

		_, err = io.Copy(file, os.Stdin)

		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}
	} else {
		// read from clip, give to stdout
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			os.Create(fileName)
		}

		file, err := os.Open(fileName)

		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}

		defer file.Close()

		_, err = io.Copy(os.Stdout, file)

		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

	}
}
