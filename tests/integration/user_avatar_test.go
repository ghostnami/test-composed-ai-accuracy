// hunk: @@ -35,7 +35,6 @@ func TestUserAvatar(t *testing.T) {
	}

	session := loginUser(t, "user2")

	imgData := &bytes.Buffer{}

// hunk: @@ -66,7 +65,6 @@ func TestUserAvatar(t *testing.T) {
	}

	req := NewRequestWithBody(t, "POST", "/user/settings/avatar", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	session.MakeRequest(t, req, http.StatusSeeOther)
