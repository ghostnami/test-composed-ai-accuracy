// hunk: @@ -133,7 +133,7 @@ func loadOAuth2From(rootCfg ConfigProvider) {

	// FIXME: at the moment, no matter oauth2 is enabled or not, it must generate a "oauth2 JWT_SECRET"
	// Because this secret is also used as GeneralTokenSigningSecret (as a quick not-that-breaking fix for some legacy problems).
	// Including: account validation token, etc ...
	// In main branch, the signing token should be refactored (eg: one unique for LFS/OAuth2/etc ...)
	jwtSecretBase64 := loadSecret(sec, "JWT_SECRET_URI", "JWT_SECRET")
	if InstallLock {
