// Package cmenu - Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import "github.com/kyleu/todoforge/app/lib/menu"

//nolint:lll
var (
	menuItemCollection     = &menu.Item{Key: "collection", Title: "Collections", Description: "Collection of TODO items", Icon: "archive", Route: "/collection", Children: menu.Items{menuItemCollectionItem}}
	menuItemCollectionItem = &menu.Item{Key: "item", Title: "Items", Description: "TODO item", Icon: "file", Route: "/collection/item"}
)

func generatedMenu() menu.Items {
	return menu.Items{
		menuItemCollection,
	}
}
