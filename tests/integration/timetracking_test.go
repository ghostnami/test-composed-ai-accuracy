// hunk: @@ -46,9 +46,7 @@ func testViewTimetrackingControls(t *testing.T, session *TestSession, user, repo
	AssertHTMLElement(t, htmlDoc, ".issue-add-time", canTrackTime)

	issueLink := path.Join(user, repo, "issues", issue)
	reqStart := NewRequest(t, "POST", path.Join(issueLink, "times", "stopwatch", "start"))
	if canTrackTime {
		session.MakeRequest(t, reqStart, http.StatusOK)

// hunk: @@ -65,9 +63,7 @@ func testViewTimetrackingControls(t *testing.T, session *TestSession, user, repo
		// Sleep for 1 second to not get wrong order for stopping timer
		time.Sleep(time.Second)

		reqStop := NewRequest(t, "POST", path.Join(issueLink, "times", "stopwatch", "stop"))
		session.MakeRequest(t, reqStop, http.StatusOK)

		req = NewRequest(t, "GET", issueLink)
