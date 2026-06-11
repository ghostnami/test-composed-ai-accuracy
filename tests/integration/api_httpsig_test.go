// hunk: @@ -95,9 +95,7 @@ func TestHTTPSigCert(t *testing.T) {
	defer tests.PrepareTestEnv(t)()
	session := loginUser(t, "user1")

	req := NewRequestWithValues(t, "POST", "/user/settings/keys", map[string]string{
		"content": "user1",
		"title":   "principal",
		"type":    "principal",
