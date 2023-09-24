// Content managed by Project Forge, see [projectforge.md] for details.
package graphql

import (
	"slices"

	"github.com/graph-gophers/graphql-go"
	otelgraphql "github.com/graph-gophers/graphql-go/trace/otel"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel"

	"github.com/kyleu/todoforge/app/util"
)

type reg struct {
	Title     string
	Schema    *graphql.Schema
	ExecCount int
}

type Service struct {
	schemata map[string]*reg
}

func NewService() *Service {
	return &Service{schemata: map[string]*reg{}}
}

func (s *Service) RegisterStringSchema(key string, title string, content string, target any) error {
	if _, ok := s.schemata[key]; ok {
		return errors.Errorf("duplicate registration of schema [%s]", key)
	}
	sch, err := graphql.ParseSchema(content, target, graphql.Tracer(&otelgraphql.Tracer{Tracer: otel.GetTracerProvider().Tracer(util.AppKey)}))
	if err != nil {
		return errors.Wrapf(err, "unable to parse schema for [%s]", key)
	}
	s.schemata[key] = &reg{Title: title, Schema: sch}
	return nil
}

func (s *Service) Keys() []string {
	ret := lo.Keys(s.schemata)
	slices.Sort(ret)
	return ret
}

func (s *Service) Titles() map[string]string {
	ret := make(map[string]string, len(s.schemata))
	for k, v := range s.schemata {
		ret[k] = v.Title
	}
	return ret
}

func (s *Service) Schema(key string) *graphql.Schema {
	ret := s.schemata[key]
	if ret == nil {
		return nil
	}
	return ret.Schema
}

func (s *Service) Close() error {
	return nil
}
