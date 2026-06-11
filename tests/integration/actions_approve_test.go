// hunk: @@ -127,10 +127,7 @@ jobs:
		)

		// user2 approves all runs
		req = NewRequest(t, "POST", dataURL)
		user2Session.MakeRequest(t, req, http.StatusOK)

		// check runs
