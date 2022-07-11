package operation

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

func StripFrontMatter(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)
	defer writer.Flush()

	scanner.Scan()
	if scanner.Text() == "---" {
		for scanner.Scan() {
			if scanner.Text() == "---" {
				break
			}
		}

	} else {
		_, err := writer.WriteString(fmt.Sprintf("%s\n", scanner.Text()))
		if err != nil {
			log.Fatal(err)

		}
	}

	scanner.Scan()

	if scanner.Text() != "" {
		_, err := writer.WriteString(fmt.Sprintf("%s\n", scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}
	}

	for scanner.Scan() {
		_, err := writer.WriteString(fmt.Sprintf("%s\n", scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
