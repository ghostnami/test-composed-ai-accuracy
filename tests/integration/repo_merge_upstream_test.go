// hunk: @@ -49,7 +49,6 @@ func TestRepoMergeUpstream(t *testing.T) {

		// create fork-branch
		req = NewRequestWithValues(t, "POST", fmt.Sprintf("/%s/test-repo-fork/branches/_new/branch/master", forkUser.Name), map[string]string{
			"new_branch_name": "fork-branch",
		})
		session.MakeRequest(t, req, http.StatusSeeOther)
// hunk: @@ -81,7 +80,6 @@ func TestRepoMergeUpstream(t *testing.T) {
			t.Run("DetectSameBranch", func(t *testing.T) {
				// if the fork-branch name also exists in the base repo, then use that branch instead
				req = NewRequestWithValues(t, "POST", "/user2/repo1/branches/_new/branch/master", map[string]string{
					"new_branch_name": "fork-branch",
				})
				sessionBaseUser.MakeRequest(t, req, http.StatusSeeOther)
// hunk: @@ -99,14 +97,12 @@ func TestRepoMergeUpstream(t *testing.T) {
			})

			// click the "sync fork" button
			req = NewRequest(t, "POST", mergeUpstreamLink)
			session.MakeRequest(t, req, http.StatusOK)
			checkFileContent("fork-branch", "test-content-1")

			// delete the "fork-branch" from the base repo
			req = NewRequest(t, "POST", "/user2/repo1/branches/delete?name=fork-branch")
			sessionBaseUser.MakeRequest(t, req, http.StatusOK)
		})

// hunk: @@ -151,7 +147,6 @@ func TestRepoMergeUpstream(t *testing.T) {
		t.Run("FastForwardOnly", func(t *testing.T) {
			// Create a clean branch for fast-forward testing
			req = NewRequestWithValues(t, "POST", fmt.Sprintf("/%s/test-repo-fork/branches/_new/branch/master", forkUser.Name), map[string]string{
				"new_branch_name": "ff-test-branch",
			})
			session.MakeRequest(t, req, http.StatusSeeOther)
