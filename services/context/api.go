// hunk: @@ -227,7 +227,7 @@ func APIContexter() func(http.Handler) http.Handler {

			ctx.SetContextValue(apiContextKey, ctx)

			// FIXME: GLOBAL-PARSE-FORM: see more details in another FIXME comment
			if ctx.Req.Method == http.MethodPost && strings.Contains(ctx.Req.Header.Get("Content-Type"), "multipart/form-data") {
				if !ctx.ParseMultipartForm() {
					return
