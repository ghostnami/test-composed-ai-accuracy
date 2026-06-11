// hunk: @@ -255,7 +255,7 @@ func newMarkupRenderer(name string, sec ConfigSection) {
	}

	// ATTENTION! at the moment, only a safe set like "allow-scripts" are allowed for sandbox mode.
	// "allow-same-origin" should NEVER be used, it leads to XSS attack: makes the JS in iframe can access parent window's config and send requests with user's credentials.
	renderContentSandbox := sec.Key("RENDER_CONTENT_SANDBOX").MustString("allow-scripts allow-popups")
	if renderContentSandbox == "disabled" {
		renderContentSandbox = ""
