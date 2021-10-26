package logger

import (
	"text/tabwriter"
)

type Logger interface {
	GetHeader(w *tabwriter.Writer) (*tabwriter.Writer, int)
	Log(w *tabwriter.Writer) (*tabwriter.Writer, int)
}
