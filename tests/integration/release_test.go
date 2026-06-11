// hunk: @@ -28,7 +28,6 @@ func createNewRelease(t *testing.T, session *TestSession, repoURL, tag, title st
	assert.True(t, exists, "The template has changed")

	postData := map[string]string{
		"tag_name":   tag,
		"tag_target": "master",
		"title":      title,
