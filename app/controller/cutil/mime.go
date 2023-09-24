// Content managed by Project Forge, see [projectforge.md] for details.
package cutil

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"gopkg.in/yaml.v2"

	"github.com/kyleu/todoforge/app/util"
)

const (
	mimeJSON = "application/json"
	mimeXML  = "text/xml"
	mimeYAML = "application/x-yaml"
)

var (
	AllowedRequestHeaders  = "*"
	AllowedResponseHeaders = "*"
)

func WriteCORS(rc *fasthttp.RequestCtx) {
	setIfEmpty := func(k string, v string) {
		if x := rc.Response.Header.Peek(k); len(x) == 0 {
			rc.Response.Header.Set(k, v)
		}
	}
	setIfEmpty(fasthttp.HeaderAccessControlAllowHeaders, AllowedRequestHeaders)
	setIfEmpty(fasthttp.HeaderAccessControlAllowMethods, "GET,POST,DELETE,PUT,PATCH,OPTIONS,HEAD")
	if x := string(rc.Request.Header.Peek(fasthttp.HeaderReferer)); x == "" {
		setIfEmpty(fasthttp.HeaderAccessControlAllowOrigin, "*")
	} else {
		u, err := url.Parse(x)
		if err == nil {
			o := u.Hostname()
			if u.Port() != "" {
				o += ":" + u.Port()
			}
			sch := u.Scheme
			if strings.Contains(o, ".network") {
				sch = "https"
			}
			setIfEmpty(fasthttp.HeaderAccessControlAllowOrigin, sch+"://"+o)
		} else {
			setIfEmpty(fasthttp.HeaderAccessControlAllowOrigin, "*")
		}
	}
	setIfEmpty(fasthttp.HeaderAccessControlAllowCredentials, util.BoolTrue)
	setIfEmpty(fasthttp.HeaderAccessControlExposeHeaders, AllowedResponseHeaders)
}

func RespondDebug(rc *fasthttp.RequestCtx, filename string, body any) (string, error) {
	return RespondJSON(rc, filename, RequestCtxToMap(rc, body))
}

func RespondJSON(rc *fasthttp.RequestCtx, filename string, body any) (string, error) {
	b := util.ToJSONBytes(body, true)
	return RespondMIME(filename, mimeJSON, "json", b, rc)
}

type XMLResponse struct {
	Result any `xml:"result"`
}

func RespondXML(rc *fasthttp.RequestCtx, filename string, body any) (string, error) {
	body = XMLResponse{Result: body}
	b, err := xml.Marshal(body)
	if err != nil {
		return "", errors.Wrapf(err, "can't serialize response of type [%T] to XML", body)
	}
	return RespondMIME(filename, mimeXML, "xml", b, rc)
}

func RespondYAML(rc *fasthttp.RequestCtx, filename string, body any) (string, error) {
	b, err := yaml.Marshal(body)
	if err != nil {
		return "", err
	}
	return RespondMIME(filename, mimeYAML, "yaml", b, rc)
}

func RespondMIME(filename string, mime string, ext string, ba []byte, rc *fasthttp.RequestCtx) (string, error) {
	rc.Response.Header.SetContentType(mime + "; charset=UTF-8")
	if len(filename) > 0 {
		if !strings.HasSuffix(filename, "."+ext) {
			filename = filename + "." + ext
		}
		rc.Response.Header.Set("Content-Disposition", fmt.Sprintf(`attachment; filename=%q`, filename))
	}
	WriteCORS(rc)
	if len(ba) == 0 {
		return "", errors.New("no bytes available to write")
	}
	if _, err := rc.Write(ba); err != nil {
		return "", errors.Wrap(err, "cannot write to response")
	}

	return "", nil
}

func GetContentType(rc *fasthttp.RequestCtx) string {
	ret := string(rc.Request.Header.ContentType())
	if idx := strings.Index(ret, ";"); idx > -1 {
		ret = ret[0:idx]
	}
	t := string(rc.URI().QueryArgs().Peek("t"))
	switch t {
	case "debug":
		return t
	case "json":
		return mimeJSON
	case "xml":
		return mimeXML
	case "yaml", "yml":
		return mimeYAML
	default:
		return strings.TrimSpace(ret)
	}
}

func IsContentTypeJSON(c string) bool {
	return c == "application/json" || c == "text/json"
}

func IsContentTypeXML(c string) bool {
	return c == "application/xml" || c == mimeXML
}

func IsContentTypeYAML(c string) bool {
	return c == "application/x-yaml" || c == "text/yaml"
}
