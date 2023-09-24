package gql

import (
	_ "embed"

	"github.com/kyleu/todoforge/app/lib/graphql"
	"github.com/kyleu/todoforge/app/util"
)

//go:embed schema.graphql
var schemaString string

type Schema struct {
	svc *graphql.Service
	sch string
}

func NewSchema(svc *graphql.Service) *Schema {
	ret := &Schema{svc: svc, sch: schemaString}
	err := ret.svc.RegisterStringSchema(util.AppKey, util.AppName, ret.sch, ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func (s *Schema) Hello() string {
	return "Howdy!"
}

func (s *Schema) Kill() string {
	return "OK!"
}
