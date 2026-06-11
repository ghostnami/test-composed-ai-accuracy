// hunk: @@ -1614,7 +1614,7 @@ func TestPullRequestWithPathsRebase(t *testing.T) {
		testCreateFile(t, session, "user2", repoName, repo.DefaultBranch, "", "dir1/dir1.txt", "1")
		testCreateFile(t, session, "user2", repoName, repo.DefaultBranch, "", "dir2/dir2.txt", "2")
		wfFileContent := `name: ci
on:
  pull_request:
    paths:
      - 'dir1/**'
// hunk: @@ -1639,12 +1639,10 @@ jobs:
		apiPull, err := doAPICreatePullRequest(apiCtx, "user2", repoName, repo.DefaultBranch, "update-dir2")(t)
		runner.fetchNoTask(t)
		assert.NoError(t, err)
		// change the file in "dir1"
		testEditFile(t, session, "user2", repoName, repo.DefaultBranch, "dir1/dir1.txt", "11")
		// update by rebase
		req := NewRequest(t, "POST", fmt.Sprintf("/%s/%s/pulls/%d/update?style=rebase", "user2", repoName, apiPull.Index))
		session.MakeRequest(t, req, http.StatusSeeOther)
		runner.fetchNoTask(t)
	})
