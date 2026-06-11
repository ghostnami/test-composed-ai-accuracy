// hunk: @@ -826,7 +826,6 @@ func TestPackageContainer(t *testing.T) {
		newOwnerName := "newUsername"

		req := NewRequestWithValues(t, "POST", "/user/settings", map[string]string{
			"name":     newOwnerName,
			"email":    "user2@example.com",
			"language": "en-US",
// hunk: @@ -836,7 +835,6 @@ func TestPackageContainer(t *testing.T) {
		t.Run(fmt.Sprintf("Catalog[%s]", newOwnerName), checkCatalog(newOwnerName))

		req = NewRequestWithValues(t, "POST", "/user/settings", map[string]string{
			"name":     user.Name,
			"email":    "user2@example.com",
			"language": "en-US",
