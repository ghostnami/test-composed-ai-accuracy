// hunk: @@ -63,10 +63,9 @@ func TestMoveRepoProjectColumns(t *testing.T) {

	sess := loginUser(t, "user1")
	req := NewRequest(t, "GET", fmt.Sprintf("/%s/projects/%d", repo2.FullName(), project1.ID))
	sess.MakeRequest(t, req, http.StatusOK)

	req = NewRequestWithJSON(t, "POST", fmt.Sprintf("/%s/projects/%d/move", repo2.FullName(), project1.ID), map[string]any{
		"columns": []map[string]any{
			{"columnID": columns[1].ID, "sorting": 0},
			{"columnID": columns[2].ID, "sorting": 1},
