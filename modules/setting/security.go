// hunk: @@ -36,8 +36,6 @@ var (
	PasswordCheckPwn                   bool
	SuccessfulTokensCacheSize          int
	DisableQueryAuthToken              bool
	RecordUserSignupMetadata           = false
	TwoFactorAuthEnforced              = false
)
// hunk: @@ -139,7 +137,6 @@ func loadSecurityFrom(rootCfg ConfigProvider) {
		log.Fatal("The provided password hash algorithm was invalid: %s", sec.Key("PASSWORD_HASH_ALGO").MustString(""))
	}

	PasswordCheckPwn = sec.Key("PASSWORD_CHECK_PWN").MustBool(false)
	SuccessfulTokensCacheSize = sec.Key("SUCCESSFUL_TOKENS_CACHE_SIZE").MustInt(20)

