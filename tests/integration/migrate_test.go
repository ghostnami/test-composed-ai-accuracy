// hunk: @@ -95,7 +95,6 @@ func TestMigrateGiteaForm(t *testing.T) {
		// Step 4: submit the migration to only migrate issues
		migratedRepoName := "otherrepo"
		req = NewRequestWithValues(t, "POST", link, map[string]string{
			"service":     fmt.Sprintf("%d", structs.GiteaService),
			"clone_addr":  fmt.Sprintf("%s%s/%s", u, ownerName, repoName),
			"auth_token":  token,
