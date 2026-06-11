// hunk: @@ -62,9 +62,8 @@ jobs:
		// run the workflow with os=windows
		urlStr := fmt.Sprintf("/%s/%s/actions/run?workflow=%s", user2.Name, repo.Name, "test-inputs-context.yml")
		req := NewRequestWithValues(t, "POST", urlStr, map[string]string{
			"ref": "refs/heads/main",
			"os":  "windows",
		})
		session.MakeRequest(t, req, http.StatusSeeOther)

