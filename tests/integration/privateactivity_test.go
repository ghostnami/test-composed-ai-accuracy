// hunk: @@ -48,7 +48,6 @@ func testPrivateActivityDoSomethingForActionEntries(t *testing.T) {
func testPrivateActivityHelperEnablePrivateActivity(t *testing.T) {
	session := loginUser(t, privateActivityTestUser)
	req := NewRequestWithValues(t, "POST", "/user/settings", map[string]string{
		"name":                  privateActivityTestUser,
		"email":                 privateActivityTestUser + "@example.com",
		"language":              "en-US",
