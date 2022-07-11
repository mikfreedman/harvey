package command

import (
	"io"
	"os"

	"github.com/mikfreedman/harvey/operation"
)

type Strip struct {
	Input *os.File `arg:"" help:"input file, specify stdin with -"`
}

func (s *Strip) Run(w io.Writer) error {
	operation.StripFrontMatter(s.Input, w)
	return nil
}
