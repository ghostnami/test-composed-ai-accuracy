// hunk: @@ -89,7 +89,6 @@ func testPullCommentRetarget(t *testing.T, u *url.URL, session *TestSession) {

	// do retarget
	req := NewRequestWithValues(t, "POST", fmt.Sprintf("/user2/repo1/pull/%d/target_branch", prIssue.PullRequest.Index), map[string]string{
		"target_branch": "test-branch/retarget-no-conflict",
	})
	session.MakeRequest(t, req, http.StatusOK)
