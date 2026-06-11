// hunk: @@ -129,13 +129,7 @@ func TestUserSettingsUpdatePassword(t *testing.T) {
		defer tests.PrintCurrentTest(t)()

		session := loginUser(t, "user2")
		req := NewRequestWithValues(t, "POST", "/user/settings/account", map[string]string{
			"old_password": "password",
			"password":     "password",
			"retype":       "password",
// hunk: @@ -147,16 +141,8 @@ func TestUserSettingsUpdatePassword(t *testing.T) {
		defer tests.PrintCurrentTest(t)()

		WithDisabledFeatures(t, setting.UserFeatureManageCredentials)
		session := loginUser(t, "user2")
		req := NewRequest(t, "POST", "/user/settings/account")
		session.MakeRequest(t, req, http.StatusNotFound)
	})
}
// hunk: @@ -168,16 +154,8 @@ func TestUserSettingsUpdateEmail(t *testing.T) {
		defer tests.PrintCurrentTest(t)()

		WithDisabledFeatures(t, setting.UserFeatureManageCredentials)
		session := loginUser(t, "user2")
		req := NewRequest(t, "POST", "/user/settings/account/email")
		session.MakeRequest(t, req, http.StatusNotFound)
	})
}
// hunk: @@ -189,16 +167,8 @@ func TestUserSettingsDeleteEmail(t *testing.T) {
		defer tests.PrintCurrentTest(t)()

		WithDisabledFeatures(t, setting.UserFeatureManageCredentials)
		session := loginUser(t, "user2")
		req := NewRequest(t, "POST", "/user/settings/account/email/delete")
		session.MakeRequest(t, req, http.StatusNotFound)
	})
}
// hunk: @@ -212,14 +182,7 @@ func TestUserSettingsDelete(t *testing.T) {
		WithDisabledFeatures(t, setting.UserFeatureDeletion)

		session := loginUser(t, "user2")
		req := NewRequest(t, "POST", "/user/settings/account/delete")
		session.MakeRequest(t, req, http.StatusNotFound)
	})
}
// hunk: @@ -308,15 +271,10 @@ func TestUserSettingsApplications(t *testing.T) {
		t.Run("OAuthApplicationsEdit", func(t *testing.T) {
			defer tests.PrintCurrentTest(t)()

			t.Run("Invalid URL", func(t *testing.T) {
				defer tests.PrintCurrentTest(t)()

				req := NewRequestWithValues(t, "POST", "/user/settings/applications/oauth2/2", map[string]string{
					"application_name":    "Test native app",
					"redirect_uris":       "ftp://127.0.0.1",
					"confidential_client": "false",
// hunk: @@ -332,7 +290,6 @@ func TestUserSettingsApplications(t *testing.T) {
				defer tests.PrintCurrentTest(t)()

				req := NewRequestWithValues(t, "POST", "/user/settings/applications/oauth2/2", map[string]string{
					"application_name":    "Test native app",
					"redirect_uris":       "http://127.0.0.1",
					"confidential_client": "false",
