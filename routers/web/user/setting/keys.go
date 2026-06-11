// hunk: @@ -325,7 +325,7 @@ func loadKeysData(ctx *context.Context) {
	ctx.Data["GPGKeys"] = gpgkeys
	tokenToSign := asymkey_model.VerificationToken(ctx.Doer, 1)

	// generate a new aes cipher using the token
	ctx.Data["TokenToSign"] = tokenToSign

	principals, err := db.Find[asymkey_model.PublicKey](ctx, asymkey_model.FindPublicKeyOptions{
