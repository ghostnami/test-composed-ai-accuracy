// hunk: @@ -61,9 +61,7 @@ func branchAction(t *testing.T, button string) (*HTMLDoc, string) {
		t.Skip()
	}

	req = NewRequest(t, "POST", link)
	session.MakeRequest(t, req, http.StatusOK)

	url, err := url.Parse(link)
