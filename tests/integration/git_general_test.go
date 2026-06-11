// hunk: @@ -507,10 +507,7 @@ type doProtectBranchOptions struct {
func doProtectBranchExt(ctx APITestContext, ruleName string, opts doProtectBranchOptions) func(t *testing.T) {
	// We are going to just use the owner to set the protection.
	return func(t *testing.T) {
		formData := map[string]string{
			"rule_name":                 ruleName,
			"unprotected_file_patterns": opts.UnprotectedFilePatterns,
			"protected_file_patterns":   opts.ProtectedFilePatterns,
// hunk: @@ -694,11 +691,7 @@ func doPushCreate(ctx APITestContext, u *url.URL) func(t *testing.T) {

func doBranchDelete(ctx APITestContext, owner, repo, branch string) func(*testing.T) {
	return func(t *testing.T) {
		req := NewRequest(t, "POST", fmt.Sprintf("/%s/%s/branches/delete?name=%s", url.PathEscape(owner), url.PathEscape(repo), url.QueryEscape(branch)))
		ctx.Session.MakeRequest(t, req, http.StatusOK)
	}
}
