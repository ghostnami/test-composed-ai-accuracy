// hunk: @@ -34,7 +34,6 @@ func TestRenameUsername(t *testing.T) {

	session := loginUser(t, "user2")
	req := NewRequestWithValues(t, "POST", "/user/settings", map[string]string{
		"name":     "newUsername",
		"email":    "user2@example.com",
		"language": "en-US",
// hunk: @@ -57,7 +56,6 @@ func TestViewLimitedAndPrivateUserAndRename(t *testing.T) {
	oldName := org22.Name
	newName := "org22_renamed"
	req = NewRequestWithValues(t, "POST", "/org/"+oldName+"/settings/rename", map[string]string{
		"org_name":     oldName,
		"new_org_name": newName,
	})
// hunk: @@ -79,7 +77,6 @@ func TestViewLimitedAndPrivateUserAndRename(t *testing.T) {
	oldName = org23.Name
	newName = "org23_renamed"
	req = NewRequestWithValues(t, "POST", "/org/"+oldName+"/settings/rename", map[string]string{
		"org_name":     oldName,
		"new_org_name": newName,
	})
// hunk: @@ -102,7 +99,6 @@ func TestViewLimitedAndPrivateUserAndRename(t *testing.T) {
	newName = "user31_renamed"
	session2 := loginUser(t, oldName)
	req = NewRequestWithValues(t, "POST", "/user/settings", map[string]string{
		"name":       newName,
		"visibility": "2", // private
	})
// hunk: @@ -150,7 +146,6 @@ func TestRenameInvalidUsername(t *testing.T) {
		t.Logf("Testing username %s", invalidUsername)

		req := NewRequestWithValues(t, "POST", "/user/settings", map[string]string{
			"name":  invalidUsername,
			"email": "user2@example.com",
		})
// hunk: @@ -178,7 +173,6 @@ func TestRenameReservedUsername(t *testing.T) {
	locale := translation.NewLocale("en-US")
	for _, reservedUsername := range reservedUsernames {
		req := NewRequestWithValues(t, "POST", "/user/settings", map[string]string{
			"name":     reservedUsername,
			"email":    "user2@example.com",
			"language": "en-US",
// hunk: @@ -334,7 +328,6 @@ func TestUserLocationMapLink(t *testing.T) {

	session := loginUser(t, "user2")
	req := NewRequestWithValues(t, "POST", "/user/settings", map[string]string{
		"name":     "user2",
		"email":    "user@example.com",
		"language": "en-US",
