// hunk: @@ -50,16 +50,13 @@ func TestActionsRunnerModify(t *testing.T) {

	doUpdate := func(t *testing.T, sess *TestSession, baseURL string, id int64, description string, expectedStatus int) {
		req := NewRequestWithValues(t, "POST", fmt.Sprintf("%s/%d", baseURL, id), map[string]string{
			"description": description,
		})
		sess.MakeRequest(t, req, expectedStatus)
	}

	doDelete := func(t *testing.T, sess *TestSession, baseURL string, id int64, expectedStatus int) {
		req := NewRequest(t, "POST", fmt.Sprintf("%s/%d/delete", baseURL, id))
		sess.MakeRequest(t, req, expectedStatus)
	}

