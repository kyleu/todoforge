// Package cutil - Content managed by Project Forge, see [projectforge.md] for details.
package cutil

import (
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/todoforge/app/util"
)

type Arg struct {
	Key         string   `json:"key"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Type        string   `json:"type,omitempty"`
	Default     string   `json:"default,omitempty"`
	Choices     []string `json:"choices,omitempty"`
}

type Args []*Arg

type ArgResults struct {
	Args    Args          `json:"args"`
	Values  util.ValueMap `json:"values"`
	Missing []string      `json:"missing,omitempty"`
}

func (a *ArgResults) HasMissing() bool {
	return len(a.Missing) > 0
}

func CollectArgs(rc *fasthttp.RequestCtx, args Args) *ArgResults {
	ret := make(util.ValueMap, len(args))
	var missing []string
	lo.ForEach(args, func(arg *Arg, _ int) {
		qa := rc.URI().QueryArgs()
		if !qa.Has(arg.Key) {
			missing = append(missing, arg.Key)
		}
		ret[arg.Key] = string(qa.Peek(arg.Key))
	})
	return &ArgResults{Args: args, Values: ret, Missing: missing}
}
