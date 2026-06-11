// hunk: @@ -157,7 +157,6 @@ func testViewRepoPrivate(t *testing.T) {

		// set unit code to "anonymous read"
		req = NewRequestWithValues(t, "POST", "/org3/repo3/settings/public_access", map[string]string{
			"repo-unit-access-" + strconv.Itoa(int(unit.TypeCode)): "anonymous-read",
		})
		session.MakeRequest(t, req, http.StatusSeeOther)
// hunk: @@ -168,9 +167,7 @@ func testViewRepoPrivate(t *testing.T) {
		assert.Contains(t, resp.Body.String(), `<span class="ui basic orange label">Public Access</span>`)

		// remove "anonymous read"
		req = NewRequest(t, "POST", "/org3/repo3/settings/public_access")
		session.MakeRequest(t, req, http.StatusSeeOther)

		// try to "anonymous read" (not found)
