package logger

import (
	"text/tabwriter"
)

type Logger interface {
	GetHeader(w *tabwriter.Writer)
	GetBody(w *tabwriter.Writer)
	Log(w *tabwriter.Writer)
}
