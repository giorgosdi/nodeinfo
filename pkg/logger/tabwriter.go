package logger

import (
	"fmt"
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
	defer w.Flush()
	w, header := logger.GetHeader(w)
	fmt.Print(header)
	w, body := logger.Log(w)
	fmt.Print(body)
	fmt.Fprintf(w, "\n")
}
