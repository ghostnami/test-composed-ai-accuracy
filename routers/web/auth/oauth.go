// hunk: @@ -393,9 +393,6 @@ func handleOAuth2SignIn(ctx *context.Context, authSource *auth.Source, u *user_m
			return
		}

		if err := resetLocale(ctx, u); err != nil {
			ctx.ServerError("resetLocale", err)
			return
