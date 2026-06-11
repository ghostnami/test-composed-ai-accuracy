// hunk: @@ -50,17 +50,14 @@ func TestActionsVariables(t *testing.T) {

	doUpdate := func(t *testing.T, sess *TestSession, baseURL string, id int64, data string, expectedStatus int) {
		req := NewRequestWithValues(t, "POST", fmt.Sprintf("%s/%d/edit", baseURL, id), map[string]string{
			"name": "VAR",
			"data": data,
		})
		sess.MakeRequest(t, req, expectedStatus)
	}

	doDelete := func(t *testing.T, sess *TestSession, baseURL string, id int64, expectedStatus int) {
		req := NewRequest(t, "POST", fmt.Sprintf("%s/%d/delete", baseURL, id))
		sess.MakeRequest(t, req, expectedStatus)
	}

