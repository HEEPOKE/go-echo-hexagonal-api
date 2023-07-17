package utils

import (
	"fmt"
	"io"

	"github.com/mattn/go-colorable"
)

type writerFuncWrapper struct {
	writeFunc func(p []byte) (n int, err error)
}

func (w *writerFuncWrapper) Write(p []byte) (n int, err error) {
	return w.writeFunc(p)
}

func ColorLoggerOutput() io.Writer {
	colorOutput := colorable.NewColorableStdout()

	return &writerFuncWrapper{
		writeFunc: func(p []byte) (n int, err error) {
			coloredOutput := fmt.Sprintf("\x1b[32m%s\x1b[0m", p)

			return colorOutput.Write([]byte(coloredOutput))
		},
	}
}
