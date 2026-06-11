// hunk: @@ -51,38 +51,27 @@ type MergeOptions struct {
	DeleteBranch bool
}

func testPullMerge(t *testing.T, session *TestSession, user, repo, pullNum string, mergeOptions MergeOptions) *httptest.ResponseRecorder {
	options := map[string]string{
		"do":                        string(mergeOptions.Style),
		"head_commit_id":            mergeOptions.HeadCommitID,
		"delete_branch_after_merge": util.Iif(mergeOptions.DeleteBranch, "on", ""),
	}
	var resp *httptest.ResponseRecorder
	require.Eventually(t, func() bool {
		req := NewRequestWithValues(t, "POST", fmt.Sprintf("/%s/%s/pulls/%s/merge", user, repo, pullNum), options)
		resp = session.MakeRequest(t, req, NoExpectedStatus)
		return resp.Code == http.StatusOK
	}, 5*time.Second, 50*time.Millisecond, "Timed out waiting for pull merge to succeed")

	redirect := test.RedirectURL(resp)
	assert.Equal(t, fmt.Sprintf("/%s/%s/pulls/%s", user, repo, pullNum), redirect)

	pullNumInt, err := strconv.ParseInt(pullNum, 10, 64)
	assert.NoError(t, err)
	repository, err := repo_model.GetRepositoryByOwnerAndName(t.Context(), user, repo)
	assert.NoError(t, err)
	pull, err := issues_model.GetPullRequestByIndex(t.Context(), repository.ID, pullNumInt)
	assert.NoError(t, err)
	assert.True(t, pull.HasMerged)

// hunk: @@ -97,9 +86,7 @@ func testPullCleanUp(t *testing.T, session *TestSession, user, repo, pullnum str
	htmlDoc := NewHTMLParser(t, resp.Body)
	link, exists := htmlDoc.doc.Find(".timeline-item .delete-branch-after-merge").Attr("data-url")
	assert.True(t, exists, "The template has changed, can not find delete button url")
	req = NewRequest(t, "POST", link)
	resp = session.MakeRequest(t, req, http.StatusOK)

	return resp
// hunk: @@ -844,11 +831,8 @@ func TestPullAutoMergeAfterCommitStatusSucceed(t *testing.T) {
			HeadBranch: "master",
		})

		// Change the "master" branch to "protected"
		req := NewRequestWithValues(t, "POST", "/user2/repo1/settings/branches/edit", map[string]string{
			"rule_name":             "master",
			"enable_push":           "true",
			"enable_status_check":   "true",
// hunk: @@ -937,11 +921,8 @@ func TestPullAutoMergeAfterCommitStatusSucceedAndApproval(t *testing.T) {
			HeadBranch: "master",
		})

		// Change master branch to protected
		req := NewRequestWithValues(t, "POST", "/user2/repo1/settings/branches/edit", map[string]string{
			"rule_name":             "master",
			"enable_push":           "true",
			"enable_status_check":   "true",
// hunk: @@ -993,10 +974,7 @@ func TestPullAutoMergeAfterCommitStatusSucceedAndApproval(t *testing.T) {

		// approve the PR from non-author
		approveSession := loginUser(t, "user2")
		testSubmitReview(t, approveSession, "user2", "repo1", strconv.Itoa(int(pr.Index)), sha, "approve", http.StatusOK)

		time.Sleep(2 * time.Second)

// hunk: @@ -1067,11 +1045,8 @@ func TestPullAutoMergeAfterCommitStatusSucceedAndApprovalForAgitFlow(t *testing.
		})

		session := loginUser(t, "user1")
		// Change master branch to protected
		req := NewRequestWithValues(t, "POST", "/user2/repo1/settings/branches/edit", map[string]string{
			"rule_name":             "master",
			"enable_push":           "true",
			"enable_status_check":   "true",
// hunk: @@ -1122,10 +1097,7 @@ func TestPullAutoMergeAfterCommitStatusSucceedAndApprovalForAgitFlow(t *testing.

		// approve the PR from non-author
		approveSession := loginUser(t, "user1")
		testSubmitReview(t, approveSession, "user2", "repo1", strconv.Itoa(int(pr.Index)), sha, "approve", http.StatusOK)

		// reload pr again
		pr = unittest.AssertExistsAndLoadBean(t, &issues_model.PullRequest{ID: pr.ID})
// hunk: @@ -1156,11 +1128,8 @@ func TestPullNonMergeForAdminWithBranchProtection(t *testing.T) {
			HeadBranch: "master",
		})

		// Change master branch to protected
		pbCreateReq := NewRequestWithValues(t, "POST", "/user2/repo1/settings/branches/edit", map[string]string{
			"rule_name":                  "master",
			"enable_push":                "true",
			"enable_status_check":        "true",
// hunk: @@ -1172,7 +1141,6 @@ func TestPullNonMergeForAdminWithBranchProtection(t *testing.T) {
		token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeWriteRepository)

		mergeReq := NewRequestWithValues(t, "POST", "/api/v1/repos/user2/repo1/pulls/6/merge", map[string]string{
			"head_commit_id":            "",
			"merge_when_checks_succeed": "false",
			"force_merge":               "true",
