// hunk: @@ -92,7 +92,6 @@ func TestAuthorizeShow(t *testing.T) {

	htmlDoc := NewHTMLParser(t, resp.Body)
	AssertHTMLElement(t, htmlDoc, "#authorize-app", true)
}

func TestAuthorizeRedirectWithExistingGrant(t *testing.T) {
