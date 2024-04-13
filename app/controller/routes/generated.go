// Package routes - Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/todoforge/app/controller"
	"github.com/kyleu/todoforge/app/controller/ccollection"
)

func generatedRoutes(r *mux.Router) {
	makeRoute(r, http.MethodGet, "/collection/item", ccollection.ItemList)
	makeRoute(r, http.MethodGet, "/collection/item/_new", ccollection.ItemCreateForm)
	makeRoute(r, http.MethodPost, "/collection/item/_new", ccollection.ItemCreate)
	makeRoute(r, http.MethodGet, "/collection/item/_random", ccollection.ItemRandom)
	makeRoute(r, http.MethodGet, "/collection/item/{id}", ccollection.ItemDetail)
	makeRoute(r, http.MethodGet, "/collection/item/{id}/edit", ccollection.ItemEditForm)
	makeRoute(r, http.MethodPost, "/collection/item/{id}/edit", ccollection.ItemEdit)
	makeRoute(r, http.MethodGet, "/collection/item/{id}/delete", ccollection.ItemDelete)
	makeRoute(r, http.MethodGet, "/collection", controller.CollectionList)
	makeRoute(r, http.MethodGet, "/collection/_new", controller.CollectionCreateForm)
	makeRoute(r, http.MethodPost, "/collection/_new", controller.CollectionCreate)
	makeRoute(r, http.MethodGet, "/collection/_random", controller.CollectionRandom)
	makeRoute(r, http.MethodGet, "/collection/{id}", controller.CollectionDetail)
	makeRoute(r, http.MethodGet, "/collection/{id}/edit", controller.CollectionEditForm)
	makeRoute(r, http.MethodPost, "/collection/{id}/edit", controller.CollectionEdit)
	makeRoute(r, http.MethodGet, "/collection/{id}/delete", controller.CollectionDelete)
}
