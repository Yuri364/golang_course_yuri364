package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main()  {
	// cat text.txt | go run main.go
	scanner := bufio.NewScanner(os.Stdin)

	bufSize := 1024 * 1024 * 5 // 5 Mb
	buf := make([]byte, 0, bufSize)
	scanner.Buffer(buf, bufSize)

	i:=0
	for scanner.Scan() {
		text := scanner.Text()

		for _, line := range strings.Split(text, "\n") {

			if line == "" {
				continue
			}

			fmt.Println(i, line)

			i++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("%v", err)
	}
}