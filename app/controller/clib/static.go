// Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/assets"
)

func Favicon(rc *fasthttp.RequestCtx) {
	data, contentType, err := assets.EmbedAsset("favicon.ico")
	assetResponse(rc, data, contentType, err)
}

func RobotsTxt(rc *fasthttp.RequestCtx) {
	data, contentType, err := assets.EmbedAsset("robots.txt")
	assetResponse(rc, data, contentType, err)
}

func Static(rc *fasthttp.RequestCtx) {
	p := strings.TrimPrefix(string(rc.Request.URI().Path()), "/assets")
	p = strings.TrimPrefix(p, "/")
	if strings.Contains(p, "../") {
		rc.Error("invalid path", fasthttp.StatusNotFound)
	} else {
		data, contentType, e := assets.EmbedAsset(p)
		assetResponse(rc, data, contentType, e)
	}
}

func assetResponse(rc *fasthttp.RequestCtx, data []byte, contentType string, err error) {
	if err == nil {
		rc.Response.Header.SetContentType(contentType)
		rc.SetStatusCode(fasthttp.StatusOK)
		cutil.WriteCORS(rc)
		_, _ = rc.Write(data)
	} else {
		rc.Error(err.Error(), fasthttp.StatusNotFound)
	}
}
