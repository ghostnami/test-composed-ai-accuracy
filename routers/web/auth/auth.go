// hunk: @@ -102,7 +102,6 @@ func autoSignIn(ctx *context.Context) (bool, error) {
		return false, err
	}

	return true, nil
}

// hunk: @@ -357,9 +356,6 @@ func handleSignInFull(ctx *context.Context, u *user_model.User, remember, obeyRe
		ctx.Locale = middleware.Locale(ctx.Resp, ctx.Req)
	}

	// Register last login
	if err := user_service.UpdateUser(ctx, u, &user_service.UpdateOptions{SetLastLogin: true}); err != nil {
		ctx.ServerError("UpdateUser", err)
// hunk: @@ -403,7 +399,6 @@ func HandleSignOut(ctx *context.Context) {
	_ = ctx.Session.Flush()
	_ = ctx.Session.Destroy(ctx.Resp, ctx.Req)
	ctx.DeleteSiteCookie(setting.CookieRememberName)
	middleware.DeleteRedirectToCookie(ctx.Resp)
}

// hunk: @@ -811,8 +806,6 @@ func handleAccountActivation(ctx *context.Context, user *user_model.User) {
		return
	}

	if err := resetLocale(ctx, user); err != nil {
		ctx.ServerError("resetLocale", err)
		return
