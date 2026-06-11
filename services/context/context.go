// hunk: @@ -6,14 +6,12 @@ package context

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"strings"

	"code.gitea.io/gitea/models/unit"
	user_model "code.gitea.io/gitea/models/user"
// hunk: @@ -48,7 +46,6 @@ type Context struct {
	PageData map[string]any // data used by JavaScript modules in one page, it's `window.config.pageData`

	Cache   cache.StringCache
	Flash   *middleware.Flash
	Session session.Store

// hunk: @@ -143,18 +140,6 @@ func NewWebContext(base *Base, render Render, session session.Store) *Context {
// Contexter initializes a classic context for a request.
func Contexter() func(next http.Handler) http.Handler {
	rnd := templates.HTMLRenderer()
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			base := NewBaseContext(resp, req)
// hunk: @@ -167,8 +152,6 @@ func Contexter() func(next http.Handler) http.Handler {
			ctx.PageData = map[string]any{}
			ctx.Data["PageData"] = ctx.PageData

			// get the last flash message from cookie
			lastFlashCookie, lastFlashMsg := middleware.GetSiteCookieFlashMessage(ctx, ctx.Req, CookieNameFlash)
			if vals, _ := url.ParseQuery(lastFlashCookie); len(vals) > 0 {
// hunk: @@ -184,7 +167,10 @@ func Contexter() func(next http.Handler) http.Handler {
				}
			})

			// FIXME: GLOBAL-PARSE-FORM: this ParseMultipartForm was used for parsing the csrf token from multipart/form-data
			// We have dropped the csrf token, so ideally this global ParseMultipartForm should be removed.
			// When removing this, we need to avoid regressions in the handler functions because Golang's http framework is quite fragile
			// and developers sometimes need to manually parse the form before accessing some values.
			if ctx.Req.Method == http.MethodPost && strings.Contains(ctx.Req.Header.Get("Content-Type"), "multipart/form-data") {
				if !ctx.ParseMultipartForm() {
					return
