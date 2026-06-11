// hunk: @@ -46,7 +46,6 @@ func testRepoFork(t *testing.T, session *TestSession, ownerName, repoName, forkO
	_, exists = htmlDoc.doc.Find(fmt.Sprintf(".owner.dropdown .item[data-value=\"%d\"]", forkOwner.ID)).Attr("data-value")
	assert.True(t, exists, "Fork owner '%s' is not present in select box", forkOwnerName)
	req = NewRequestWithValues(t, "POST", link, map[string]string{
		"uid":                strconv.FormatInt(forkOwner.ID, 10),
		"repo_name":          forkRepoName,
		"fork_single_branch": forkBranch,
