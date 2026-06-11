// hunk: @@ -21,7 +21,6 @@ func TestXSSUserFullName(t *testing.T) {

	session := loginUser(t, user.Name)
	req := NewRequestWithValues(t, "POST", "/user/settings", map[string]string{
		"name":      user.Name,
		"full_name": fullName,
		"email":     user.Email,
