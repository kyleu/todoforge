package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/todoforge/app/controller"
	"github.com/kyleu/todoforge/app/controller/ccollection"
)

const routeNew, routeRandom, routeEdit, routeDelete = "/_new", "/_random", "/edit", "/delete"

func generatedRoutes(r *mux.Router) {
	generatedRoutesItem(r, "/collection/item")
	generatedRoutesCollection(r, "/collection")
}

func generatedRoutesItem(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, ccollection.ItemList)
	makeRoute(r, http.MethodGet, prefix+routeNew, ccollection.ItemCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, ccollection.ItemCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, ccollection.ItemRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, ccollection.ItemDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, ccollection.ItemEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, ccollection.ItemEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, ccollection.ItemDelete)
}

func generatedRoutesCollection(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.CollectionList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.CollectionCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.CollectionCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.CollectionRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.CollectionDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.CollectionEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.CollectionEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.CollectionDelete)
}
