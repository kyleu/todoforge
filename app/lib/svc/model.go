// Package svc - Content managed by Project Forge, see [projectforge.md] for details.
package svc

import "fmt"

type Model interface {
	fmt.Stringer
	TitleString() string
	Strings() []string
	ToCSV() ([]string, [][]string)
	WebPath() string
	ToData() []any
}

type ModelSeq []Model
