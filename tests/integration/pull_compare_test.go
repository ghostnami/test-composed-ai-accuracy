// hunk: @@ -133,7 +133,6 @@ func TestPullCompare_EnableAllowEditsFromMaintainer(t *testing.T) {
		dataURL, exists := htmlDoc.doc.Find("#allow-edits-from-maintainers").Attr("data-url")
		assert.True(t, exists)
		req := NewRequestWithValues(t, "POST", dataURL+"/set_allow_maintainer_edit", map[string]string{
			"allow_maintainer_edit": "true",
		})
		user4Session.MakeRequest(t, req, http.StatusOK)
// hunk: @@ -159,7 +158,6 @@ func TestPullCompare_EnableAllowEditsFromMaintainer(t *testing.T) {
				lastCommit := htmlDoc.GetInputValueByName("last_commit")
				assert.NotEmpty(t, lastCommit)
				req := NewRequestWithValues(t, "POST", editFileLink, map[string]string{
					"last_commit":    lastCommit,
					"tree_path":      "README.md",
					"content":        "File is edited by the maintainer user2",
