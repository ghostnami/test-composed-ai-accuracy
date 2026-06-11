// hunk: @@ -34,7 +34,6 @@ func TestOrgProjectAccess(t *testing.T) {
	// change the org's visibility to private
	session := loginUser(t, "user2")
	req = NewRequestWithValues(t, "POST", "/org/org3/settings", map[string]string{
		"name":       "org3",
		"visibility": "2",
	})
// hunk: @@ -48,7 +47,6 @@ func TestOrgProjectAccess(t *testing.T) {
	// disable team1's project unit
	session = loginUser(t, "user2")
	req = NewRequestWithValues(t, "POST", "/org/org3/teams/team1/edit", map[string]string{
		"team_name":   "team1",
		"repo_access": "specific",
		"permission":  "read",
