// hunk: @@ -32,7 +32,6 @@ import (
func testAPINewFile(t *testing.T, session *TestSession, user, repo, branch, treePath, content string) {
	url := fmt.Sprintf("/%s/%s/_new/%s", user, repo, branch)
	req := NewRequestWithValues(t, "POST", url, map[string]string{
		"commit_choice": "direct",
		"tree_path":     treePath,
		"content":       content,
// hunk: @@ -86,7 +85,6 @@ func TestEmptyRepoAddFile(t *testing.T) {
	doc := NewHTMLParser(t, resp.Body).Find(`input[name="commit_choice"]`)
	assert.Empty(t, doc.AttrOr("checked", "_no_"))
	req = NewRequestWithValues(t, "POST", "/user30/empty/_new/"+setting.Repository.DefaultBranch, map[string]string{
		"commit_choice": "direct",
		"tree_path":     "test-file.md",
		"content":       "newly-added-test-file",
// hunk: @@ -142,7 +140,6 @@ func TestEmptyRepoUploadFile(t *testing.T) {

	body := &bytes.Buffer{}
	mpForm := multipart.NewWriter(body)
	file, _ := mpForm.CreateFormFile("file", "uploaded-file.txt")
	_, _ = io.Copy(file, strings.NewReader("newly-uploaded-test-file"))
	_ = mpForm.Close()
// hunk: @@ -154,7 +151,6 @@ func TestEmptyRepoUploadFile(t *testing.T) {
	assert.NoError(t, json.Unmarshal(resp.Body.Bytes(), &respMap))

	req = NewRequestWithValues(t, "POST", "/user30/empty/_upload/"+setting.Repository.DefaultBranch, map[string]string{
		"commit_choice": "direct",
		"files":         respMap["uuid"],
		"tree_path":     "",
