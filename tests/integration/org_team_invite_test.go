// hunk: @@ -39,9 +39,7 @@ func TestOrgTeamEmailInvite(t *testing.T) {
	session := loginUser(t, "user1")

	teamURL := fmt.Sprintf("/org/%s/teams/%s", org.Name, team.Name)
	req := NewRequestWithValues(t, "POST", teamURL+"/action/add", map[string]string{
		"uid":   "1",
		"uname": user.Email,
	})
// hunk: @@ -58,10 +56,7 @@ func TestOrgTeamEmailInvite(t *testing.T) {

	// join the team
	inviteURL := "/org/invite/" + invites[0].Token
	req = NewRequest(t, "POST", inviteURL)
	resp = session.MakeRequest(t, req, http.StatusSeeOther)
	req = NewRequest(t, "GET", test.RedirectURL(resp))
	session.MakeRequest(t, req, http.StatusOK)
// hunk: @@ -93,7 +88,6 @@ func TestOrgTeamEmailInviteRedirectsExistingUser(t *testing.T) {

	teamURL := fmt.Sprintf("/org/%s/teams/%s", org.Name, team.Name)
	req := NewRequestWithValues(t, "POST", teamURL+"/action/add", map[string]string{
		"uid":   "1",
		"uname": user.Email,
	})
// hunk: @@ -111,9 +105,7 @@ func TestOrgTeamEmailInviteRedirectsExistingUser(t *testing.T) {
	req = NewRequest(t, "GET", "/user/login?redirect_to="+url.QueryEscape(inviteURL))
	resp = MakeRequest(t, req, http.StatusOK)

	req = NewRequestWithValues(t, "POST", "/user/login", map[string]string{
		"user_name": "user5",
		"password":  "password",
	})
// hunk: @@ -135,9 +127,7 @@ func TestOrgTeamEmailInviteRedirectsExistingUser(t *testing.T) {
	session.jar.SetCookies(baseURL, cr.Cookies())

	// make the request
	req = NewRequest(t, "POST", test.RedirectURL(resp))
	resp = session.MakeRequest(t, req, http.StatusSeeOther)
	req = NewRequest(t, "GET", test.RedirectURL(resp))
	session.MakeRequest(t, req, http.StatusOK)
// hunk: @@ -164,7 +154,6 @@ func TestOrgTeamEmailInviteRedirectsNewUser(t *testing.T) {

	teamURL := fmt.Sprintf("/org/%s/teams/%s", org.Name, team.Name)
	req := NewRequestWithValues(t, "POST", teamURL+"/action/add", map[string]string{
		"uid":   "1",
		"uname": "doesnotexist@example.com",
	})
// hunk: @@ -182,9 +171,7 @@ func TestOrgTeamEmailInviteRedirectsNewUser(t *testing.T) {
	req = NewRequest(t, "GET", "/user/sign_up?redirect_to="+url.QueryEscape(inviteURL))
	resp = MakeRequest(t, req, http.StatusOK)

	req = NewRequestWithValues(t, "POST", "/user/sign_up", map[string]string{
		"user_name": "doesnotexist",
		"email":     "doesnotexist@example.com",
		"password":  "examplePassword!1",
// hunk: @@ -208,9 +195,7 @@ func TestOrgTeamEmailInviteRedirectsNewUser(t *testing.T) {
	session.jar.SetCookies(baseURL, cr.Cookies())

	// make the redirected request
	req = NewRequest(t, "POST", test.RedirectURL(resp))
	resp = session.MakeRequest(t, req, http.StatusSeeOther)
	req = NewRequest(t, "GET", test.RedirectURL(resp))
	session.MakeRequest(t, req, http.StatusOK)
// hunk: @@ -243,7 +228,6 @@ func TestOrgTeamEmailInviteRedirectsNewUserWithActivation(t *testing.T) {

	teamURL := fmt.Sprintf("/org/%s/teams/%s", org.Name, team.Name)
	req := NewRequestWithValues(t, "POST", teamURL+"/action/add", map[string]string{
		"uid":   "1",
		"uname": "doesnotexist@example.com",
	})
// hunk: @@ -283,9 +267,7 @@ func TestOrgTeamEmailInviteRedirectsNewUserWithActivation(t *testing.T) {
	// should be redirected to accept the invite
	assert.Equal(t, inviteURL, test.RedirectURL(resp))

	req = NewRequest(t, "POST", test.RedirectURL(resp))
	resp = session.MakeRequest(t, req, http.StatusSeeOther)
	req = NewRequest(t, "GET", test.RedirectURL(resp))
	session.MakeRequest(t, req, http.StatusOK)
// hunk: @@ -319,7 +301,6 @@ func TestOrgTeamEmailInviteRedirectsExistingUserWithLogin(t *testing.T) {

	teamURL := fmt.Sprintf("/org/%s/teams/%s", org.Name, team.Name)
	req := NewRequestWithValues(t, "POST", teamURL+"/action/add", map[string]string{
		"uid":   "1",
		"uname": user.Email,
	})
// hunk: @@ -342,9 +323,7 @@ func TestOrgTeamEmailInviteRedirectsExistingUserWithLogin(t *testing.T) {
	assert.Equal(t, inviteURL, test.RedirectURL(resp))

	// make the request
	req = NewRequest(t, "POST", test.RedirectURL(resp))
	resp = session.MakeRequest(t, req, http.StatusSeeOther)
	req = NewRequest(t, "GET", test.RedirectURL(resp))
	session.MakeRequest(t, req, http.StatusOK)
