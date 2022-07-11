package command

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/mikfreedman/harvey/operation"
)

type Strip struct {
	Input   *os.File `arg:"" help:"input file, specify stdin with -"`
	InPlace bool     `default:"false" short:"i"`
}

func (s *Strip) Run() error {
	var multi io.Writer
	var a *bytes.Buffer

	multi = os.Stdout

	if s.InPlace {
		a = new(bytes.Buffer)
		multi = a
	}

	operation.StripFrontMatter(s.Input, multi)

	if s.InPlace {
		s.Input.Close()
		f, err := s.Input.Stat()
		if err != nil {
			log.Fatal(err)
		}
		os.WriteFile(s.Input.Name(), a.Bytes(), f.Mode())
	}
	return nil
}
