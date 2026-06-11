// hunk: @@ -52,9 +52,7 @@ func testSuccessfullEdit(t *testing.T, formData user_model.User) {

func makeRequest(t *testing.T, formData user_model.User, headerCode int) {
	session := loginUser(t, "user1")
	req := NewRequestWithValues(t, "POST", "/-/admin/users/"+strconv.Itoa(int(formData.ID))+"/edit", map[string]string{
		"user_name":  formData.Name,
		"login_name": formData.LoginName,
		"login_type": "0-0",
// hunk: @@ -96,10 +94,7 @@ func TestAdminDeleteUser(t *testing.T) {
				query = "?purge=true"
			}

			req := NewRequest(t, "POST", fmt.Sprintf("/-/admin/users/%d/delete%s", entry.userID, query))
			session.MakeRequest(t, req, http.StatusSeeOther)

			assertUserDeleted(t, entry.userID)
