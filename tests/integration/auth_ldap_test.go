// hunk: @@ -125,7 +125,7 @@ type ldapAuthOptions struct {
	groupTeamMapRemoval   string
}

func (te *ldapTestEnv) buildAuthSourcePayload(opts ...ldapAuthOptions) map[string]string {
	opt := util.OptionalArg(opts)
	// Modify user filter to test group filter explicitly
	userFilter := "(&(objectClass=inetOrgPerson)(memberOf=cn=git,ou=people,dc=planetexpress,dc=com)(uid=%s))"
// hunk: @@ -134,7 +134,6 @@ func (te *ldapTestEnv) buildAuthSourcePayload(csrf string, opts ...ldapAuthOptio
	}

	return map[string]string{
		"type":                     "2",
		"name":                     "ldap",
		"host":                     te.serverHost,
// hunk: @@ -164,8 +163,7 @@ func (te *ldapTestEnv) buildAuthSourcePayload(csrf string, opts ...ldapAuthOptio

func (te *ldapTestEnv) addAuthSource(t *testing.T, opts ...ldapAuthOptions) {
	session := loginUser(t, "user1")
	req := NewRequestWithValues(t, "POST", "/-/admin/auths/new", te.buildAuthSourcePayload(opts...))
	session.MakeRequest(t, req, http.StatusSeeOther)
}

// hunk: @@ -212,13 +210,12 @@ func TestLDAPAuthChange(t *testing.T) {
	req = NewRequest(t, "GET", href)
	resp = session.MakeRequest(t, req, http.StatusOK)
	doc = NewHTMLParser(t, resp.Body)
	host, _ := doc.Find(`input[name="host"]`).Attr("value")
	assert.Equal(t, te.serverHost, host)
	binddn, _ := doc.Find(`input[name="bind_dn"]`).Attr("value")
	assert.Equal(t, "uid=gitea,ou=service,dc=planetexpress,dc=com", binddn)

	req = NewRequestWithValues(t, "POST", href, te.buildAuthSourcePayload(ldapAuthOptions{groupTeamMapRemoval: "off"}))
	session.MakeRequest(t, req, http.StatusSeeOther)

	req = NewRequest(t, "GET", href)
// hunk: @@ -267,8 +264,7 @@ func TestLDAPUserSyncWithEmptyUsernameAttribute(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	session := loginUser(t, "user1")
	payload := te.buildAuthSourcePayload()
	payload["attribute_username"] = ""
	req := NewRequestWithValues(t, "POST", "/-/admin/auths/new", payload)
	session.MakeRequest(t, req, http.StatusSeeOther)
// hunk: @@ -285,7 +281,6 @@ func TestLDAPUserSyncWithEmptyUsernameAttribute(t *testing.T) {

	for _, u := range te.gitLDAPUsers {
		req := NewRequestWithValues(t, "POST", "/user/login", map[string]string{
			"user_name": u.UserName,
			"password":  u.Password,
		})
// hunk: @@ -512,8 +507,7 @@ func TestLDAPPreventInvalidGroupTeamMap(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	session := loginUser(t, "user1")
	payload := te.buildAuthSourcePayload(ldapAuthOptions{groupTeamMap: `{"NOT_A_VALID_JSON"["MISSING_DOUBLE_POINT"]}`, groupTeamMapRemoval: "off"})
	req := NewRequestWithValues(t, "POST", "/-/admin/auths/new", payload)
	session.MakeRequest(t, req, http.StatusOK) // StatusOK = failed, StatusSeeOther = ok
}
