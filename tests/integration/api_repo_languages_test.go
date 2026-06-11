// hunk: @@ -28,7 +28,6 @@ func TestRepoLanguages(t *testing.T) {

		// Save new file to master branch
		req = NewRequestWithValues(t, "POST", "/user2/repo1/_new/master/", map[string]string{
			"last_commit":   lastCommit,
			"tree_path":     "test.go",
			"content":       "package main",
