// hunk: @@ -130,7 +130,6 @@ func testNewIssue(t *testing.T, session *TestSession, user, repo, title, content
	link, exists := htmlDoc.doc.Find("form.ui.form").Attr("action")
	assert.True(t, exists, "The template has changed")
	req = NewRequestWithValues(t, "POST", link, map[string]string{
		"title":   title,
		"content": content,
	})
// hunk: @@ -150,15 +149,12 @@ func testNewIssue(t *testing.T, session *TestSession, user, repo, title, content
}

func testIssueDelete(t *testing.T, session *TestSession, issueURL string) {
	req := NewRequest(t, "POST", path.Join(issueURL, "delete"))
	session.MakeRequest(t, req, http.StatusSeeOther)
}

func testIssueAssign(t *testing.T, session *TestSession, repoLink string, issueID, assigneeID int64) {
	req := NewRequestWithValues(t, "POST", fmt.Sprintf(repoLink+"/issues/assignee?issue_ids=%d", issueID), map[string]string{
		"id":     strconv.FormatInt(assigneeID, 10),
		"action": "", // empty action means assign
	})
// hunk: @@ -176,7 +172,6 @@ func testIssueAddComment(t *testing.T, session *TestSession, issueURL, content,
	commentCount := htmlDoc.doc.Find(".comment-list .comment .render-content").Length()

	req = NewRequestWithValues(t, "POST", link, map[string]string{
		"content": content,
		"status":  status,
	})
// hunk: @@ -200,8 +195,7 @@ func testIssueAddComment(t *testing.T, session *TestSession, issueURL, content,

func testIssueChangeMilestone(t *testing.T, session *TestSession, repoLink string, issueID, milestoneID int64) {
	req := NewRequestWithValues(t, "POST", fmt.Sprintf(repoLink+"/issues/milestone?issue_ids=%d", issueID), map[string]string{
		"id": strconv.FormatInt(milestoneID, 10),
	})
	resp := session.MakeRequest(t, req, http.StatusOK)
	assert.Equal(t, `{"ok":true}`, strings.TrimSpace(resp.Body.String()))
// hunk: @@ -219,21 +213,18 @@ func TestEditIssue(t *testing.T) {
	issueURL := testNewIssue(t, session, "user2", "repo1", "Title", "Description")

	req := NewRequestWithValues(t, "POST", issueURL+"/content", map[string]string{
		"content": "modified content",
		"context": fmt.Sprintf("/%s/%s", "user2", "repo1"),
	})
	session.MakeRequest(t, req, http.StatusOK)

	req = NewRequestWithValues(t, "POST", issueURL+"/content", map[string]string{
		"content": "modified content",
		"context": fmt.Sprintf("/%s/%s", "user2", "repo1"),
	})
	session.MakeRequest(t, req, http.StatusBadRequest)

	req = NewRequestWithValues(t, "POST", issueURL+"/content", map[string]string{
		"content":         "modified content",
		"content_version": "1",
		"context":         fmt.Sprintf("/%s/%s", "user2", "repo1"),
// hunk: @@ -267,13 +258,9 @@ func TestIssueCommentDelete(t *testing.T) {
	assert.Equal(t, comment1, comment.Content)

	// Using the ID of a comment that does not belong to the repository must fail
	req := NewRequest(t, "POST", fmt.Sprintf("/%s/%s/comments/%d/delete", "user5", "repo4", commentID))
	session.MakeRequest(t, req, http.StatusNotFound)
	req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/comments/%d/delete", "user2", "repo1", commentID))
	session.MakeRequest(t, req, http.StatusOK)
	unittest.AssertNotExistsBean(t, &issues_model.Comment{ID: commentID})
}
// hunk: @@ -292,13 +279,11 @@ func TestIssueCommentUpdate(t *testing.T) {

	// Using the ID of a comment that does not belong to the repository must fail
	req := NewRequestWithValues(t, "POST", fmt.Sprintf("/%s/%s/comments/%d", "user5", "repo4", commentID), map[string]string{
		"content": modifiedContent,
	})
	session.MakeRequest(t, req, http.StatusNotFound)

	req = NewRequestWithValues(t, "POST", fmt.Sprintf("/%s/%s/comments/%d", "user2", "repo1", commentID), map[string]string{
		"content": modifiedContent,
	})
	session.MakeRequest(t, req, http.StatusOK)
// hunk: @@ -320,7 +305,6 @@ func TestIssueCommentUpdateSimultaneously(t *testing.T) {
	modifiedContent := comment.Content + "MODIFIED"

	req := NewRequestWithValues(t, "POST", fmt.Sprintf("/%s/%s/comments/%d", "user2", "repo1", commentID), map[string]string{
		"content": modifiedContent,
	})
	session.MakeRequest(t, req, http.StatusOK)
// hunk: @@ -328,13 +312,11 @@ func TestIssueCommentUpdateSimultaneously(t *testing.T) {
	modifiedContent = comment.Content + "2"

	req = NewRequestWithValues(t, "POST", fmt.Sprintf("/%s/%s/comments/%d", "user2", "repo1", commentID), map[string]string{
		"content": modifiedContent,
	})
	session.MakeRequest(t, req, http.StatusBadRequest)

	req = NewRequestWithValues(t, "POST", fmt.Sprintf("/%s/%s/comments/%d", "user2", "repo1", commentID), map[string]string{
		"content":         modifiedContent,
		"content_version": "1",
	})
// hunk: @@ -350,22 +332,15 @@ func TestIssueReaction(t *testing.T) {
	session := loginUser(t, "user2")
	issueURL := testNewIssue(t, session, "user2", "repo1", "Title", "Description")

	req := NewRequestWithValues(t, "POST", path.Join(issueURL, "/reactions/react"), map[string]string{
		"content": "8ball",
	})
	session.MakeRequest(t, req, http.StatusInternalServerError)
	req = NewRequestWithValues(t, "POST", path.Join(issueURL, "/reactions/react"), map[string]string{
		"content": "eyes",
	})
	session.MakeRequest(t, req, http.StatusOK)
	req = NewRequestWithValues(t, "POST", path.Join(issueURL, "/reactions/unreact"), map[string]string{
		"content": "eyes",
	})
	session.MakeRequest(t, req, http.StatusOK)
// hunk: @@ -459,14 +434,8 @@ func testIssueWithBean(t *testing.T, user string, repoID int64, title, content s

func testIssueChangeInfo(t *testing.T, user, issueURL, info, value string) {
	session := loginUser(t, user)
	req := NewRequestWithValues(t, "POST", path.Join(issueURL, info), map[string]string{
		info: value,
	})
	_ = session.MakeRequest(t, req, http.StatusOK)
}
// hunk: @@ -700,7 +669,7 @@ func TestUpdateIssueDeadline(t *testing.T) {
	assert.Equal(t, api.StateOpen, issueBefore.State())

	session := loginUser(t, owner.Name)
	urlStr := fmt.Sprintf("%s/%s/issues/%d/deadline", owner.Name, repoBefore.Name, issueBefore.Index)

	req := NewRequestWithValues(t, "POST", urlStr, map[string]string{"deadline": "2022-04-06"})
	session.MakeRequest(t, req, http.StatusOK)
