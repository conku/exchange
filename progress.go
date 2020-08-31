package exchange

import "github.com/conku/qor"

// Progress defined importing/exporting progress
type Progress struct {
	Current uint
	Total   uint
	Value   interface{}
	Errors  qor.Errors
	Cells   []Cell
}

// Cell is a data cell, which including its data and error that happened when importing/exporting data
type Cell struct {
	Header string
	Value  interface{}
	Error  error
}
