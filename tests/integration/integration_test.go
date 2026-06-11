// hunk: @@ -225,16 +225,11 @@ func loginUser(t testing.TB, userName string) *TestSession {

func loginUserWithPassword(t testing.TB, userName, password string) *TestSession {
	t.Helper()
	req := NewRequestWithValues(t, "POST", "/user/login", map[string]string{
		"user_name": userName,
		"password":  password,
	})
	resp := MakeRequest(t, req, http.StatusSeeOther)

	ch := http.Header{}
	ch.Add("Cookie", strings.Join(resp.Header()["Set-Cookie"], ";"))
// hunk: @@ -256,7 +251,6 @@ var tokenCounter int64
func getTokenForLoggedInUser(t testing.TB, session *TestSession, scopes ...auth.AccessTokenScope) string {
	t.Helper()
	urlValues := url.Values{}
	urlValues.Add("name", fmt.Sprintf("api-testing-token-%d", atomic.AddInt64(&tokenCounter, 1)))
	for _, scope := range scopes {
		urlValues.Add("scope-dummy", string(scope)) // it only needs to start with "scope-" to be accepted
// hunk: @@ -436,20 +430,3 @@ func VerifyJSONSchema(t testing.TB, resp *httptest.ResponseRecorder, schemaFile
	assert.Empty(t, result.Errors())
	assert.True(t, result.Valid())
}
