// hunk: @@ -33,7 +33,7 @@ func testGeneratePngBytes() []byte {
	return buff.Bytes()
}

func testCreateIssueAttachment(t *testing.T, session *TestSession, repoURL, filename string, content []byte, expectedStatus int) string {
	body := &bytes.Buffer{}

	// Setup multi-part
// hunk: @@ -46,7 +46,6 @@ func testCreateIssueAttachment(t *testing.T, session *TestSession, csrf, repoURL
	assert.NoError(t, err)

	req := NewRequestWithBody(t, "POST", repoURL+"/issues/attachments", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	resp := session.MakeRequest(t, req, expectedStatus)

// hunk: @@ -79,20 +78,20 @@ func testUploadAttachmentDeleteTemp(t *testing.T) {
	defer web.RouteMock(route_web.RouterMockPointBeforeWebRoutes, func(resp http.ResponseWriter, req *http.Request) {
		tmpFileCountDuringUpload = countTmpFile()
	})()
	_ = testCreateIssueAttachment(t, session, "user2/repo1", "image.png", testGeneratePngBytes(), http.StatusOK)
	assert.Equal(t, 1, tmpFileCountDuringUpload, "the temp file should exist when uploaded size exceeds the parse form's max memory")
	assert.Equal(t, 0, countTmpFile(), "the temp file should be deleted after upload")
}

func testCreateAnonymousAttachment(t *testing.T) {
	session := emptyTestSession(t)
	testCreateIssueAttachment(t, session, "user2/repo1", "image.png", testGeneratePngBytes(), http.StatusSeeOther)
}

func testCreateUser2IssueAttachment(t *testing.T) {
	const repoURL = "user2/repo1"
	session := loginUser(t, "user2")
	uuid := testCreateIssueAttachment(t, session, repoURL, "image.png", testGeneratePngBytes(), http.StatusOK)

	req := NewRequest(t, "GET", repoURL+"/issues/new")
	resp := session.MakeRequest(t, req, http.StatusOK)
// hunk: @@ -102,7 +101,6 @@ func testCreateUser2IssueAttachment(t *testing.T) {
	assert.True(t, exists, "The template has changed")

	postData := map[string]string{
		"title":   "New Issue With Attachment",
		"content": "some content",
		"files":   uuid,
