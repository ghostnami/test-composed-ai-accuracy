// hunk: @@ -15,16 +15,13 @@ import (
	"code.gitea.io/gitea/models/unittest"
	"code.gitea.io/gitea/modules/test"
	"code.gitea.io/gitea/modules/translation"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func testCreateBranch(t testing.TB, session *TestSession, user, repo, oldRefSubURL, newBranchName string, expectedStatus int) string {
	req := NewRequestWithValues(t, "POST", path.Join(user, repo, "branches/_new", oldRefSubURL), map[string]string{
		"new_branch_name": newBranchName,
	})
	resp := session.MakeRequest(t, req, expectedStatus)
// hunk: @@ -126,17 +123,6 @@ func testCreateBranches(t *testing.T, giteaURL *url.URL) {
	}
}

func prepareRecentlyPushedBranchTest(t *testing.T, headSession *TestSession, baseRepo, headRepo *repo_model.Repository) {
	refSubURL := "branch/" + headRepo.DefaultBranch
	baseRepoPath := baseRepo.OwnerName + "/" + baseRepo.Name
