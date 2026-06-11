// hunk: @@ -44,7 +44,6 @@ func testRepoGenerate(t *testing.T, session *TestSession, templateID, templateOw
	_, exists = htmlDoc.doc.Find(fmt.Sprintf(`#repo_owner_dropdown .item[data-value="%d"]`, generateOwner.ID)).Attr("data-value")
	assert.True(t, exists, "Generate owner '%s' is not present in select box", generateOwnerName)
	req = NewRequestWithValues(t, "POST", link, map[string]string{
		"uid":           strconv.FormatInt(generateOwner.ID, 10),
		"repo_name":     generateRepoName,
		"repo_template": templateID,
