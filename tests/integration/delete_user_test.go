// hunk: @@ -32,11 +32,8 @@ func TestUserDeleteAccount(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	session := loginUser(t, "user8")
	urlStr := "/user/settings/account/delete?password=" + userPassword
	req := NewRequest(t, "POST", urlStr)
	session.MakeRequest(t, req, http.StatusSeeOther)

	assertUserDeleted(t, 8)
// hunk: @@ -47,11 +44,8 @@ func TestUserDeleteAccountStillOwnRepos(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	session := loginUser(t, "user2")
	urlStr := "/user/settings/account/delete?password=" + userPassword
	req := NewRequest(t, "POST", urlStr)
	session.MakeRequest(t, req, http.StatusSeeOther)

	// user should not have been deleted, because the user still owns repos
