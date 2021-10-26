package logger

import (
	"os"
	"text/tabwriter"
)

func getWriter() *tabwriter.Writer {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 5, '\t', 0)
	return w
}

func TableLogger(logger Logger) {
	w := getWriter()
	logger.GetHeader(w)
	logger.GetBody(w)
	logger.Log(w)
}
