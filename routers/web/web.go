// hunk: @@ -129,13 +129,13 @@ func webAuth(authMethod auth_service.Method) func(*context.Context) {
			// ensure the session uid is deleted
			_ = ctx.Session.Delete("uid")
		}
	}
}

// verifyAuthWithOptions checks authentication according to options
func verifyAuthWithOptions(options *common.VerifyOptions) func(ctx *context.Context) {
	crossOriginProtection := http.NewCrossOriginProtection()

	return func(ctx *context.Context) {
		// Check prohibit login users.
		if ctx.IsSigned {
// hunk: @@ -178,9 +178,9 @@ func verifyAuthWithOptions(options *common.VerifyOptions) func(ctx *context.Cont
			return
		}

		if !options.SignOutRequired && !options.DisableCrossOriginProtection {
			if err := crossOriginProtection.Check(ctx.Req); err != nil {
				http.Error(ctx.Resp, err.Error(), http.StatusForbidden)
				return
			}
		}
// hunk: @@ -292,7 +292,12 @@ func Routes() *web.Router {
	return routes
}

// optSignInFromAnyOrigin means that the user can (optionally) be signed in from any origin (no cross-origin protection)
//   - With CORS middleware: CORS middleware does the preflight request handling, the requests has Sec-Fetch-Site header.
//     The CORS mechanism already protects cross-origin requests, and the CrossOriginProtection has no "allowed origin" list, so disable CrossOriginProtection.
//   - For non-browser client requests: git clone via http, no Sec-Fetch-Site header.
//     Such requests are not cross-origin requests, so disable CrossOriginProtection.
var optSignInFromAnyOrigin = verifyAuthWithOptions(&common.VerifyOptions{DisableCrossOriginProtection: true})

// registerWebRoutes register routes
func registerWebRoutes(m *web.Router) {
// hunk: @@ -489,7 +494,7 @@ func registerWebRoutes(m *web.Router) {
	m.Post("/-/markup", reqSignIn, web.Bind(structs.MarkupOption{}), misc.Markup)

	m.Get("/-/web-theme/list", misc.WebThemeList)
	m.Post("/-/web-theme/apply", optSignIn, misc.WebThemeApply)

	m.Group("/explore", func() {
		m.Get("", func(ctx *context.Context) {
// hunk: @@ -565,12 +570,14 @@ func registerWebRoutes(m *web.Router) {
			m.Post("/grant", web.Bind(forms.GrantApplicationForm{}), auth.GrantApplicationOAuth)
			// TODO manage redirection
			m.Post("/authorize", web.Bind(forms.AuthorizationForm{}), auth.AuthorizeOAuth)
		}, reqSignIn)

		m.Group("", func() {
			m.Methods("GET, POST, OPTIONS", "/userinfo", auth.InfoOAuth)
			m.Methods("POST, OPTIONS", "/access_token", web.Bind(forms.AccessTokenForm{}), auth.AccessTokenOAuth)
			m.Methods("GET, OPTIONS", "/keys", auth.OIDCKeys)
			m.Methods("POST, OPTIONS", "/introspect", web.Bind(forms.IntrospectTokenForm{}), auth.IntrospectOAuth)
		}, optionsCorsHandler(), optSignInFromAnyOrigin)
	}, oauth2Enabled)

	m.Group("/user/settings", func() {
// hunk: @@ -1653,7 +1660,7 @@ func registerWebRoutes(m *web.Router) {
		m.Post("/action/{action:accept_transfer|reject_transfer}", reqSignIn, repo.ActionTransfer)
	}, optSignIn, context.RepoAssignment)

	common.AddOwnerRepoGitLFSRoutes(m, lfsServerEnabled, repo.CorsHandler(), optSignInFromAnyOrigin) // "/{username}/{reponame}/{lfs-paths}": git-lfs support, see also addOwnerRepoGitHTTPRouters

	addOwnerRepoGitHTTPRouters(m) // "/{username}/{reponame}/{git-paths}": git http support

