// hunk: @@ -17,9 +17,7 @@ import (

func setDefaultBranch(t *testing.T, session *TestSession, user, repo, branch string) {
	location := path.Join("/", user, repo, "settings/branches")
	req := NewRequestWithValues(t, "POST", location, map[string]string{
		"action": "default_branch",
		"branch": branch,
	})
