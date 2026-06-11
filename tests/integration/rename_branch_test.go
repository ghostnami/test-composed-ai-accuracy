// hunk: @@ -27,14 +27,9 @@ func testRenameBranch(t *testing.T, u *url.URL) {

	// get branch setting page
	session := loginUser(t, "user2")
	req := NewRequestWithValues(t, "POST", "/user2/repo1/branches/rename", map[string]string{
		"from": "master",
		"to":   "main",
	})
	session.MakeRequest(t, req, http.StatusSeeOther)

// hunk: @@ -44,7 +39,7 @@ func testRenameBranch(t *testing.T, u *url.URL) {

	// check old branch link
	req = NewRequestWithValues(t, "GET", "/user2/repo1/src/branch/master/README.md", nil)
	resp := session.MakeRequest(t, req, http.StatusSeeOther)
	location := resp.Header().Get("Location")
	assert.Equal(t, "/user2/repo1/src/branch/main/README.md", location)

// hunk: @@ -53,10 +48,7 @@ func testRenameBranch(t *testing.T, u *url.URL) {
	assert.Equal(t, "main", repo1.DefaultBranch)

	// create branch1
	req = NewRequestWithValues(t, "POST", "/user2/repo1/branches/_new/branch/main", map[string]string{
		"new_branch_name": "branch1",
	})
	session.MakeRequest(t, req, http.StatusSeeOther)
// hunk: @@ -66,7 +58,6 @@ func testRenameBranch(t *testing.T, u *url.URL) {

	// create branch2
	req = NewRequestWithValues(t, "POST", "/user2/repo1/branches/_new/branch/main", map[string]string{
		"new_branch_name": "branch2",
	})
	session.MakeRequest(t, req, http.StatusSeeOther)
// hunk: @@ -76,9 +67,8 @@ func testRenameBranch(t *testing.T, u *url.URL) {

	// rename branch2 to branch1
	req = NewRequestWithValues(t, "POST", "/user2/repo1/branches/rename", map[string]string{
		"from": "branch2",
		"to":   "branch1",
	})
	session.MakeRequest(t, req, http.StatusSeeOther)
	flashMsg := session.GetCookieFlashMessage()
// hunk: @@ -91,8 +81,7 @@ func testRenameBranch(t *testing.T, u *url.URL) {

	// delete branch1
	req = NewRequestWithValues(t, "POST", "/user2/repo1/branches/delete", map[string]string{
		"name": "branch1",
	})
	session.MakeRequest(t, req, http.StatusOK)
	branch2 = unittest.AssertExistsAndLoadBean(t, &git_model.Branch{RepoID: repo1.ID, Name: "branch2"})
// hunk: @@ -102,9 +91,8 @@ func testRenameBranch(t *testing.T, u *url.URL) {

	// rename branch2 to branch1 again
	req = NewRequestWithValues(t, "POST", "/user2/repo1/branches/rename", map[string]string{
		"from": "branch2",
		"to":   "branch1",
	})
	session.MakeRequest(t, req, http.StatusSeeOther)

