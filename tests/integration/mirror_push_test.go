// hunk: @@ -79,7 +79,6 @@ func testMirrorPush(t *testing.T, u *url.URL) {

func testCreatePushMirror(t *testing.T, session *TestSession, owner, repo, address, username, password, interval string) {
	req := NewRequestWithValues(t, "POST", fmt.Sprintf("/%s/%s/settings", url.PathEscape(owner), url.PathEscape(repo)), map[string]string{
		"action":               "push-mirror-add",
		"push_mirror_address":  address,
		"push_mirror_username": username,
// hunk: @@ -94,7 +93,6 @@ func testCreatePushMirror(t *testing.T, session *TestSession, owner, repo, addre

func doRemovePushMirror(t *testing.T, session *TestSession, owner, repo string, pushMirrorID int64) bool {
	req := NewRequestWithValues(t, "POST", fmt.Sprintf("/%s/%s/settings", url.PathEscape(owner), url.PathEscape(repo)), map[string]string{
		"action":         "push-mirror-remove",
		"push_mirror_id": strconv.FormatInt(pushMirrorID, 10),
	})
// hunk: @@ -105,7 +103,6 @@ func doRemovePushMirror(t *testing.T, session *TestSession, owner, repo string,

func doUpdatePushMirror(t *testing.T, session *TestSession, owner, repo string, pushMirrorID int64, interval string) bool {
	req := NewRequestWithValues(t, "POST", fmt.Sprintf("/%s/%s/settings", owner, repo), map[string]string{
		"action":                 "push-mirror-update",
		"push_mirror_id":         strconv.FormatInt(pushMirrorID, 10),
		"push_mirror_interval":   interval,
