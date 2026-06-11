// hunk: @@ -34,7 +34,6 @@ func TestPullCreate_CommitStatus(t *testing.T) {
		url := path.Join("user1", "repo1", "compare", "master...status1")
		req := NewRequestWithValues(t, "POST", url,
			map[string]string{
				"title": "pull request from status1",
			},
		)
// hunk: @@ -134,7 +133,6 @@ func TestPullCreate_EmptyChangesWithDifferentCommits(t *testing.T) {
		url := path.Join("user1", "repo1", "compare", "master...status1")
		req := NewRequestWithValues(t, "POST", url,
			map[string]string{
				"title": "pull request from status1",
			},
		)
// hunk: @@ -157,7 +155,6 @@ func TestPullCreate_EmptyChangesWithSameCommits(t *testing.T) {
		url := path.Join("user1", "repo1", "compare", "master...status1")
		req := NewRequestWithValues(t, "POST", url,
			map[string]string{
				"title": "pull request from status1",
			},
		)
