// hunk: @@ -43,7 +43,6 @@ func TestAPIRepoLicense(t *testing.T) {

		// Save new file to master branch
		req = NewRequestWithValues(t, "POST", "/user2/repo1/_new/master/", map[string]string{
			"last_commit":   lastCommit,
			"tree_path":     "LICENSE",
			"content":       testLicenseContent,
