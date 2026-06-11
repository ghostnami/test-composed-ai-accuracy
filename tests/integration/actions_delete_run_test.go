// hunk: @@ -133,9 +133,7 @@ jobs:
		}

		for i := 0; i < len(testCase.outcomes); i++ {
			req := NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%s/jobs/%d", user2.Name, apiRepo.Name, runIndex, i))
			resp := session.MakeRequest(t, req, http.StatusOK)
			var listResp actions.ViewResponse
			err := json.Unmarshal(resp.Body.Bytes(), &listResp)
// hunk: @@ -147,30 +145,20 @@ jobs:
			MakeRequest(t, req, http.StatusOK)
		}

		req := NewRequest(t, "GET", fmt.Sprintf("/%s/%s/actions/runs/%s", user2.Name, apiRepo.Name, runIndex))
		session.MakeRequest(t, req, http.StatusOK)

		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%s/delete", user2.Name, apiRepo.Name, runIndex))
		session.MakeRequest(t, req, http.StatusOK)

		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%s/delete", user2.Name, apiRepo.Name, runIndex))
		session.MakeRequest(t, req, http.StatusNotFound)

		req = NewRequest(t, "GET", fmt.Sprintf("/%s/%s/actions/runs/%s", user2.Name, apiRepo.Name, runIndex))
		session.MakeRequest(t, req, http.StatusNotFound)

		for i := 0; i < len(testCase.outcomes); i++ {
			req := NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%s/jobs/%d", user2.Name, apiRepo.Name, runIndex, i))
			session.MakeRequest(t, req, http.StatusNotFound)

			req = NewRequest(t, "GET", fmt.Sprintf("/%s/%s/actions/runs/%s/jobs/%d/logs", user2.Name, apiRepo.Name, runIndex, i)).
