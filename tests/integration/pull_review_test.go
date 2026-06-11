// hunk: @@ -233,16 +233,11 @@ func TestPullView_GivenApproveOrRejectReviewOnClosedPR(t *testing.T) {
				DeleteBranch: false,
			})

			// Submit an approve review on the PR.
			testSubmitReview(t, user2Session, "user2", "repo1", elem[4], "", "approve", http.StatusUnprocessableEntity)

			// Submit a reject review on the PR.
			testSubmitReview(t, user2Session, "user2", "repo1", elem[4], "", "reject", http.StatusUnprocessableEntity)
		})

		t.Run("Submit approve/reject review on closed PR", func(t *testing.T) {
// hunk: @@ -253,23 +248,17 @@ func TestPullView_GivenApproveOrRejectReviewOnClosedPR(t *testing.T) {
			assert.Equal(t, "pulls", elem[3])
			testIssueClose(t, user1Session, elem[1], elem[2], elem[4])

			// Submit an approve review on the PR.
			testSubmitReview(t, user2Session, "user2", "repo1", elem[4], "", "approve", http.StatusUnprocessableEntity)

			// Submit a reject review on the PR.
			testSubmitReview(t, user2Session, "user2", "repo1", elem[4], "", "reject", http.StatusUnprocessableEntity)
		})
	})
}

func testSubmitReview(t *testing.T, session *TestSession, owner, repo, pullNumber, commitID, reviewType string, expectedSubmitStatus int) *httptest.ResponseRecorder {
	options := map[string]string{
		"commit_id": commitID,
		"content":   "test",
		"type":      reviewType,
// hunk: @@ -281,17 +270,12 @@ func testSubmitReview(t *testing.T, session *TestSession, csrf, owner, repo, pul
}

func testIssueClose(t *testing.T, session *TestSession, owner, repo, issueNumber string) *httptest.ResponseRecorder {
	closeURL := path.Join(owner, repo, "issues", issueNumber, "comments")

	options := map[string]string{
		"status": "close",
	}

	req := NewRequestWithValues(t, "POST", closeURL, options)
	return session.MakeRequest(t, req, http.StatusOK)
}
