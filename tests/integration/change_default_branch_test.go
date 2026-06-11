// hunk: @@ -26,17 +26,13 @@ func TestChangeDefaultBranch(t *testing.T) {
	session := loginUser(t, owner.Name)
	branchesURL := fmt.Sprintf("/%s/%s/settings/branches", owner.Name, repo.Name)

	req := NewRequestWithValues(t, "POST", branchesURL, map[string]string{
		"action": "default_branch",
		"branch": "DefaultBranch",
	})
	session.MakeRequest(t, req, http.StatusSeeOther)

	req = NewRequestWithValues(t, "POST", branchesURL, map[string]string{
		"action": "default_branch",
		"branch": "does_not_exist",
	})
// hunk: @@ -110,9 +106,7 @@ func TestChangeDefaultBranchDivergence(t *testing.T) {

	// switch default branch
	newDefaultBranch := "good-sign-not-yet-validated"
	req := NewRequestWithValues(t, "POST", settingsBranchesURL, map[string]string{
		"action": "default_branch",
		"branch": newDefaultBranch,
	})
