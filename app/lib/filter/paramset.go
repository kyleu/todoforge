// Package filter - Content managed by Project Forge, see [projectforge.md] for details.
package filter

import (
	"strings"

	"github.com/samber/lo"

	"github.com/kyleu/todoforge/app/util"
)

type ParamSet map[string]*Params

func (s ParamSet) Get(key string, allowed []string, logger util.Logger) *Params {
	x, ok := s[key]
	if !ok {
		return &Params{Key: key}
	}
	return x.Filtered(allowed, logger).Sanitize(key)
}

func (s ParamSet) String() string {
	return strings.Join(lo.Map(lo.Values(s), func(p *Params, _ int) string {
		return p.String()
	}), ", ")
}
