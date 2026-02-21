package cmenu

import (
	"context"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/lib/filter"
	"github.com/kyleu/todoforge/app/lib/menu"
	"github.com/kyleu/todoforge/app/lib/telemetry"
	"github.com/kyleu/todoforge/app/lib/user"
	"github.com/kyleu/todoforge/app/util"
)

func MenuFor(
	ctx context.Context, as *app.State, isAuthed bool, isAdmin bool, profile *user.Profile, params filter.ParamSet, logger util.Logger,
) (menu.Items, any, error) {
	ctx, sp, _ := telemetry.StartSpan(ctx, "menu", logger)
	defer sp.Complete()
	var ret menu.Items
	var data any
	// $PF_SECTION_START(menu)$
	ret = append(ret, generatedMenu()...)
	admin := &menu.Item{Key: "admin", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/admin"}
	ret = append(ret, menu.Separator, admin, menu.Separator, DocsMenu(logger))
	const aboutDesc = "Get assistance and advice for using " + util.AppName
	ret = append(ret, &menu.Item{Key: "about", Title: "About", Description: aboutDesc, Icon: "question", Route: "/about"})
	// $PF_SECTION_END(menu)$
	return ret, data, nil
}
