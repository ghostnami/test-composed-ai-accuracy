// hunk: @@ -97,9 +97,7 @@ func TestSessionFileCreation(t *testing.T) {
		// We're not logged in so there should be no session
		assert.False(t, sessionFileExist(t, tmpDir, sessionID))

		req = NewRequestWithValues(t, "POST", "/user/login", map[string]string{
			"user_name": "user2",
			"password":  userPassword,
		})
