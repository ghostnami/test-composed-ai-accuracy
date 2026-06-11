// hunk: @@ -59,7 +59,6 @@ func testPullCreate(t *testing.T, session *TestSession, user, repo string, toSel
	link, exists = htmlDoc.doc.Find("form.ui.form").Attr("action")
	assert.True(t, exists, "The template has changed")
	req = NewRequestWithValues(t, "POST", link, map[string]string{
		"title": title,
	})
	resp = session.MakeRequest(t, req, http.StatusOK)
// hunk: @@ -103,7 +102,6 @@ func testPullCreateDirectly(t *testing.T, session *TestSession, opts createPullR
	link, exists := htmlDoc.doc.Find("form.ui.form").Attr("action")
	assert.True(t, exists, "The template has changed")
	params := map[string]string{
		"title": opts.Title,
	}
	if opts.ReviewerIDs != "" {
// hunk: @@ -131,7 +129,6 @@ func testPullCreateFailure(t *testing.T, session *TestSession, baseRepoOwner, ba
	link, exists := htmlDoc.doc.Find("form.ui.form").Attr("action")
	assert.True(t, exists, "The template has changed")
	req = NewRequestWithValues(t, "POST", link, map[string]string{
		"title": title,
	})
	resp = session.MakeRequest(t, req, http.StatusBadRequest)
// hunk: @@ -159,7 +156,6 @@ func TestPullCreate(t *testing.T) {
		// test create the pull request again and it should fail now
		link := "/user2/repo1/compare/master...user1/repo1:master"
		req := NewRequestWithValues(t, "POST", link, map[string]string{
			"title": "This is a pull title",
		})
		session.MakeRequest(t, req, http.StatusBadRequest)
// hunk: @@ -200,7 +196,6 @@ func TestPullCreate_TitleEscape(t *testing.T) {
		assert.True(t, exists, "The template has changed")

		req = NewRequestWithValues(t, "POST", editTestTitleURL, map[string]string{
			"title": "<u>XSS PR</u>",
		})
		session.MakeRequest(t, req, http.StatusOK)
// hunk: @@ -219,25 +214,15 @@ func TestPullCreate_TitleEscape(t *testing.T) {

func testUIDeleteBranch(t *testing.T, session *TestSession, ownerName, repoName, branchName string) {
	relURL := "/" + path.Join(ownerName, repoName, "branches")
	req := NewRequestWithValues(t, "POST", relURL+"/delete", map[string]string{
		"name": branchName,
	})
	session.MakeRequest(t, req, http.StatusOK)
}

func testDeleteRepository(t *testing.T, session *TestSession, ownerName, repoName string) {
	relURL := "/" + path.Join(ownerName, repoName, "settings")
	req := NewRequestWithValues(t, "POST", relURL+"?action=delete", map[string]string{
		"repo_name": repoName,
	})
	session.MakeRequest(t, req, http.StatusSeeOther)
