// hunk: @@ -83,9 +83,7 @@ func testAPICreateWebhookForRepo(t *testing.T, session *TestSession, userName, r
}

func testCreateWebhookForRepo(t *testing.T, session *TestSession, webhookType, userName, repoName, url, eventKind string) {
	req := NewRequestWithValues(t, "POST", "/"+userName+"/"+repoName+"/settings/hooks/"+webhookType+"/new", map[string]string{
		"payload_url":  url,
		"events":       eventKind,
		"active":       "true",
// hunk: @@ -278,7 +276,6 @@ func Test_WebhookIssueComment(t *testing.T) {
			commentID := testIssueAddComment(t, session, issueURL, "issue title3 comment1", "")
			modifiedContent := "issue title2 comment1 - modified"
			req := NewRequestWithValues(t, "POST", fmt.Sprintf("/%s/%s/comments/%d", "user2", "repo1", commentID), map[string]string{
				"content": modifiedContent,
			})
			session.MakeRequest(t, req, http.StatusOK)
// hunk: @@ -306,7 +303,6 @@ func Test_WebhookIssueComment(t *testing.T) {
			payloads = make([]api.IssueCommentPayload, 0, 2)
			triggeredEvent = ""
			req := NewRequestWithValues(t, "POST", fmt.Sprintf("/%s/%s/comments/%d", "user2", "repo1", commentID), map[string]string{
				"content": commentContent,
			})
			session.MakeRequest(t, req, http.StatusOK)
// hunk: @@ -1284,9 +1280,7 @@ jobs:
	// Call cancel ui api
	// Only a web UI API exists for cancelling workflow runs, so use the UI endpoint.
	cancelURL := fmt.Sprintf("/user2/repo1/actions/runs/%d/cancel", webhookData.payloads[0].WorkflowRun.RunNumber)
	req := NewRequest(t, "POST", cancelURL)
	session.MakeRequest(t, req, http.StatusOK)

	assert.Len(t, webhookData.payloads, 2)
// hunk: @@ -1418,9 +1412,7 @@ jobs:
	// Call cancel ui api
	// Only a web UI API exists for cancelling workflow runs, so use the UI endpoint.
	cancelURL := fmt.Sprintf("/user2/repo1/actions/runs/%d/cancel", webhookData.payloads[0].WorkflowRun.RunNumber)
	req := NewRequest(t, "POST", cancelURL)
	session.MakeRequest(t, req, http.StatusOK)

	assert.Len(t, webhookData.payloads, 2)
// hunk: @@ -1438,9 +1430,7 @@ jobs:
	// Call rerun ui api
	// Only a web UI API exists for rerunning workflow runs, so use the UI endpoint.
	rerunURL := fmt.Sprintf("/user2/repo1/actions/runs/%d/rerun", webhookData.payloads[0].WorkflowRun.RunNumber)
	req = NewRequest(t, "POST", rerunURL)
	session.MakeRequest(t, req, http.StatusOK)

	assert.Len(t, webhookData.payloads, 3)
