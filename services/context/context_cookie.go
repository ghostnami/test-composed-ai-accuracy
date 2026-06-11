// hunk: @@ -25,13 +25,11 @@ func removeSessionCookieHeader(w http.ResponseWriter) {
}

// SetSiteCookie convenience function to set most cookies consistently
func (ctx *Context) SetSiteCookie(name, value string, maxAge int) {
	middleware.SetSiteCookie(ctx.Resp, name, value, maxAge)
}

// DeleteSiteCookie convenience function to delete most cookies consistently
func (ctx *Context) DeleteSiteCookie(name string) {
	middleware.SetSiteCookie(ctx.Resp, name, "", -1)
}
