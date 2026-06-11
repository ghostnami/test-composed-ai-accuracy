// hunk: @@ -42,7 +42,6 @@ func TestActionsCollaborativeOwner(t *testing.T) {

		// add user10 to the list of collaborative owners
		req := NewRequestWithValues(t, "POST", fmt.Sprintf("/%s/%s/settings/actions/general/collaborative_owner/add", repo.Owner.UserName, repo.Name), map[string]string{
			"collaborative_owner": user10.Name,
		})
		user2Session.MakeRequest(t, req, http.StatusOK)
// hunk: @@ -51,9 +50,7 @@ func TestActionsCollaborativeOwner(t *testing.T) {
		doGitClone(dstPath, u)(t)

		// remove user10 from the list of collaborative owners
		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/settings/actions/general/collaborative_owner/delete?id=%d", repo.Owner.UserName, repo.Name, user10.ID))
		user2Session.MakeRequest(t, req, http.StatusOK)

		// the git clone will fail
