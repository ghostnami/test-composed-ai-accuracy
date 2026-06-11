// hunk: @@ -87,7 +87,6 @@ func testEditorProtectedBranch(t *testing.T) {
	session := loginUser(t, "user2")
	// Change the "master" branch to "protected"
	req := NewRequestWithValues(t, "POST", "/user2/repo1/settings/branches/edit", map[string]string{
		"rule_name":   "master",
		"enable_push": "true",
	})
// hunk: @@ -106,7 +105,6 @@ func testEditorActionPostRequest(t *testing.T, session *TestSession, requestPath
	resp := session.MakeRequest(t, req, http.StatusOK)
	htmlDoc := NewHTMLParser(t, resp.Body)
	form := map[string]string{
		"last_commit": htmlDoc.GetInputValueByName("last_commit"),
	}
	maps.Copy(form, params)
// hunk: @@ -150,7 +148,6 @@ func testEditFileToNewBranch(t *testing.T, session *TestSession, user, repo, bra
func testEditorDiffPreview(t *testing.T) {
	session := loginUser(t, "user2")
	req := NewRequestWithValues(t, "POST", "/user2/repo1/_preview/master/README.md", map[string]string{
		"content": "Hello, World (Edited)\n",
	})
	resp := session.MakeRequest(t, req, http.StatusOK)
// hunk: @@ -200,7 +197,6 @@ func testEditorWebGitCommitEmail(t *testing.T) {

	makeReq := func(t *testing.T, link string, params map[string]string, expectedUserName, expectedEmail string) *httptest.ResponseRecorder {
		lastCommit := getLastCommit(t)
		params["last_commit"] = lastCommit.ID.String()
		params["commit_choice"] = "direct"
		req := NewRequestWithValues(t, "POST", link, params)
// hunk: @@ -225,7 +221,6 @@ func testEditorWebGitCommitEmail(t *testing.T) {
		uploadForm := multipart.NewWriter(body)
		file, _ := uploadForm.CreateFormFile("file", name)
		_, _ = io.Copy(file, strings.NewReader(content))
		_ = uploadForm.Close()

		req := NewRequestWithBody(t, "POST", "/user2/repo1/upload-file", body)
// hunk: @@ -347,7 +342,7 @@ func testForkToEditFile(t *testing.T, session *TestSession, user, owner, repo, b
		assert.Contains(t, resp.Body.String(), "Fork Repository to Propose Changes")

		// fork the repository
		req = NewRequest(t, "POST", path.Join(owner, repo, "_fork", branch))
		resp = session.MakeRequest(t, req, http.StatusOK)
		assert.JSONEq(t, `{"redirect":""}`, resp.Body.String())
	}
// hunk: @@ -359,7 +354,6 @@ func testForkToEditFile(t *testing.T, session *TestSession, user, owner, repo, b
		// Archive the repository
		req := NewRequestWithValues(t, "POST", path.Join(user, repo, "settings"),
			map[string]string{
				"repo_name": repo,
				"action":    "archive",
			},
// hunk: @@ -374,7 +368,6 @@ func testForkToEditFile(t *testing.T, session *TestSession, user, owner, repo, b
		// Unfork the repository
		req = NewRequestWithValues(t, "POST", path.Join(user, repo, "settings"),
			map[string]string{
				"repo_name": repo,
				"action":    "convert_fork",
			},
// hunk: @@ -410,7 +403,6 @@ func testForkToEditFile(t *testing.T, session *TestSession, user, owner, repo, b
		resp := session.MakeRequest(t, req, http.StatusOK)
		htmlDoc := NewHTMLParser(t, resp.Body)
		editRequestForm := map[string]string{
			"last_commit":   htmlDoc.GetInputValueByName("last_commit"),
			"tree_path":     filePath,
			"content":       "new content in fork",
