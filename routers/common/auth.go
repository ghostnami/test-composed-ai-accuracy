// hunk: @@ -38,8 +38,8 @@ func AuthShared(ctx *context.Base, sessionStore auth_service.SessionStore, authM

// VerifyOptions contains required or check options
type VerifyOptions struct {
	SignInRequired               bool
	SignOutRequired              bool
	AdminRequired                bool
	DisableCrossOriginProtection bool
}
