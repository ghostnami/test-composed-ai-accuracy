// hunk: @@ -27,7 +27,6 @@ func testRepoMigrate(t testing.TB, session *TestSession, cloneAddr, repoName str
	assert.True(t, exists, "The template has changed")

	req = NewRequestWithValues(t, "POST", link, map[string]string{
		"clone_addr": cloneAddr,
		"uid":        uid,
		"repo_name":  repoName,
