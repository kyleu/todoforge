// Package routes - Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"github.com/fasthttp/router"

	"github.com/kyleu/todoforge/app/controller"
	"github.com/kyleu/todoforge/app/controller/ccollection"
)

func generatedRoutes(r *router.Router) {
	r.GET("/collection", controller.CollectionList)
	r.GET("/collection/_new", controller.CollectionCreateForm)
	r.POST("/collection/_new", controller.CollectionCreate)
	r.GET("/collection/_random", controller.CollectionRandom)
	r.GET("/collection/{id}", controller.CollectionDetail)
	r.GET("/collection/{id}/edit", controller.CollectionEditForm)
	r.POST("/collection/{id}/edit", controller.CollectionEdit)
	r.GET("/collection/{id}/delete", controller.CollectionDelete)
	r.GET("/collection/item", ccollection.ItemList)
	r.GET("/collection/item/_new", ccollection.ItemCreateForm)
	r.POST("/collection/item/_new", ccollection.ItemCreate)
	r.GET("/collection/item/_random", ccollection.ItemRandom)
	r.GET("/collection/item/{id}", ccollection.ItemDetail)
	r.GET("/collection/item/{id}/edit", ccollection.ItemEditForm)
	r.POST("/collection/item/{id}/edit", ccollection.ItemEdit)
	r.GET("/collection/item/{id}/delete", ccollection.ItemDelete)
}
