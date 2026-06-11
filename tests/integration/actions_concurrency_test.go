// hunk: @@ -51,7 +51,7 @@ func TestWorkflowConcurrency(t *testing.T) {

		wf1TreePath := ".gitea/workflows/concurrent-workflow-1.yml"
		wf1FileContent := `name: concurrent-workflow-1
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-1.yml'
// hunk: @@ -65,7 +65,7 @@ jobs:
`
		wf2TreePath := ".gitea/workflows/concurrent-workflow-2.yml"
		wf2FileContent := `name: concurrent-workflow-2
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-2.yml'
// hunk: @@ -79,7 +79,7 @@ jobs:
`
		wf3TreePath := ".gitea/workflows/concurrent-workflow-3.yml"
		wf3FileContent := `name: concurrent-workflow-3
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-3.yml'
// hunk: @@ -159,7 +159,7 @@ func TestWorkflowConcurrencyShort(t *testing.T) {

		wf1TreePath := ".gitea/workflows/concurrent-workflow-1.yml"
		wf1FileContent := `name: concurrent-workflow-1
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-1.yml'
// hunk: @@ -172,7 +172,7 @@ jobs:
`
		wf2TreePath := ".gitea/workflows/concurrent-workflow-2.yml"
		wf2FileContent := `name: concurrent-workflow-2
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-2.yml'
// hunk: @@ -185,7 +185,7 @@ jobs:
`
		wf3TreePath := ".gitea/workflows/concurrent-workflow-3.yml"
		wf3FileContent := `name: concurrent-workflow-3
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-3.yml'
// hunk: @@ -264,7 +264,7 @@ func TestWorkflowConcurrencyShortJson(t *testing.T) {

		wf1TreePath := ".gitea/workflows/concurrent-workflow-1.yml"
		wf1FileContent := `name: concurrent-workflow-1
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-1.yml'
// hunk: @@ -281,7 +281,7 @@ jobs:
`
		wf2TreePath := ".gitea/workflows/concurrent-workflow-2.yml"
		wf2FileContent := `name: concurrent-workflow-2
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-2.yml'
// hunk: @@ -298,7 +298,7 @@ jobs:
`
		wf3TreePath := ".gitea/workflows/concurrent-workflow-3.yml"
		wf3FileContent := `name: concurrent-workflow-3
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-3.yml'
// hunk: @@ -454,11 +454,7 @@ jobs:
		runner.fetchNoTask(t)
		// user2 approves the run
		pr2Run1 := unittest.AssertExistsAndLoadBean(t, &actions_model.ActionRun{RepoID: baseRepo.ID, TriggerUserID: user4.ID})
		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/approve", baseRepo.OwnerName, baseRepo.Name, pr2Run1.Index))
		user2Session.MakeRequest(t, req, http.StatusOK)
		// fetch the task and the previous task has been cancelled
		pr2Task1 := runner.fetchTask(t)
// hunk: @@ -532,7 +528,7 @@ func TestJobConcurrency(t *testing.T) {

		wf1TreePath := ".gitea/workflows/concurrent-workflow-1.yml"
		wf1FileContent := `name: concurrent-workflow-1
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-1.yml'
// hunk: @@ -546,7 +542,7 @@ jobs:
`
		wf2TreePath := ".gitea/workflows/concurrent-workflow-2.yml"
		wf2FileContent := `name: concurrent-workflow-2
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-2.yml'
// hunk: @@ -554,7 +550,7 @@ jobs:
  wf2-job1:
    runs-on: runner2
    outputs:
      version: ${{ steps.version_step.outputs.app_version }}
    steps:
      - id: version_step
        run: echo "app_version=v1.23.0" >> "$GITHUB_OUTPUT"
// hunk: @@ -568,7 +564,7 @@ jobs:
`
		wf3TreePath := ".gitea/workflows/concurrent-workflow-3.yml"
		wf3FileContent := `name: concurrent-workflow-3
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-3.yml'
// hunk: @@ -624,9 +620,7 @@ jobs:
		assert.Equal(t, actions_model.StatusCancelled, wf2Job2ActionJob.Status)

		// rerun wf2
		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/rerun", user2.Name, repo.Name, wf2Run.Index))
		_ = session.MakeRequest(t, req, http.StatusOK)

		// (rerun1) cannot fetch wf2-job2
// hunk: @@ -650,9 +644,7 @@ jobs:
		assert.Equal(t, "job-main-v1.24.0", wf2Job2Rerun1Job.ConcurrencyGroup)

		// rerun wf2-job2
		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/jobs/%d/rerun", user2.Name, repo.Name, wf2Run.Index, 1))
		_ = session.MakeRequest(t, req, http.StatusOK)
		// (rerun2) fetch and exec wf2-job2
		wf2Job2Rerun2Task := runner1.fetchTask(t)
// hunk: @@ -684,7 +676,7 @@ func TestMatrixConcurrency(t *testing.T) {

		wf1TreePath := ".gitea/workflows/concurrent-workflow-1.yml"
		wf1FileContent := `name: concurrent-workflow-1
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-1.yml'
// hunk: @@ -702,7 +694,7 @@ jobs:

		wf2TreePath := ".gitea/workflows/concurrent-workflow-2.yml"
		wf2FileContent := `name: concurrent-workflow-2
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-2.yml'
// hunk: @@ -806,7 +798,6 @@ jobs:
		// run the workflow with appVersion=v1.21 and cancel=false
		urlStr := fmt.Sprintf("/%s/%s/actions/run?workflow=%s", user2.Name, repo.Name, "workflow-dispatch-concurrency.yml")
		req := NewRequestWithValues(t, "POST", urlStr, map[string]string{
			"ref":        "refs/heads/main",
			"appVersion": "v1.21",
		})
// hunk: @@ -817,7 +808,6 @@ jobs:

		// run the workflow with appVersion=v1.22 and cancel=false
		req = NewRequestWithValues(t, "POST", urlStr, map[string]string{
			"ref":        "refs/heads/main",
			"appVersion": "v1.22",
		})
// hunk: @@ -828,7 +818,6 @@ jobs:

		// run the workflow with appVersion=v1.22 and cancel=false again
		req = NewRequestWithValues(t, "POST", urlStr, map[string]string{
			"ref":        "refs/heads/main",
			"appVersion": "v1.22",
		})
// hunk: @@ -837,7 +826,6 @@ jobs:

		// run the workflow with appVersion=v1.22 and cancel=true
		req = NewRequestWithValues(t, "POST", urlStr, map[string]string{
			"ref":        "refs/heads/main",
			"appVersion": "v1.22",
			"cancel":     "on",
// hunk: @@ -900,7 +888,6 @@ jobs:
		// run the workflow with appVersion=v1.21 and cancel=false
		urlStr := fmt.Sprintf("/%s/%s/actions/run?workflow=%s", user2.Name, repo.Name, "workflow-dispatch-concurrency.yml")
		req := NewRequestWithValues(t, "POST", urlStr, map[string]string{
			"ref":        "refs/heads/main",
			"appVersion": "v1.21",
		})
// hunk: @@ -910,7 +897,6 @@ jobs:
		assert.Equal(t, "workflow-dispatch-v1.21", run1.ConcurrencyGroup)

		req = NewRequestWithValues(t, "POST", urlStr, map[string]string{
			"ref":        "refs/heads/main",
			"appVersion": "v1.22",
		})
// hunk: @@ -921,7 +907,6 @@ jobs:

		// run the workflow with appVersion=v1.22 and cancel=false again
		req = NewRequestWithValues(t, "POST", urlStr, map[string]string{
			"ref":        "refs/heads/main",
			"appVersion": "v1.22",
		})
// hunk: @@ -931,7 +916,6 @@ jobs:

		// run the workflow with appVersion=v1.22 and cancel=true
		req = NewRequestWithValues(t, "POST", urlStr, map[string]string{
			"ref":        "refs/heads/main",
			"appVersion": "v1.22",
			"cancel":     "on",
// hunk: @@ -950,14 +934,10 @@ jobs:

		// rerun cancel true scenario

		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/rerun", user2.Name, apiRepo.Name, run2.Index))
		_ = session.MakeRequest(t, req, http.StatusOK)

		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/rerun", user2.Name, apiRepo.Name, run4.Index))
		_ = session.MakeRequest(t, req, http.StatusOK)

		task5 := runner.fetchTask(t)
// hunk: @@ -973,17 +953,13 @@ jobs:

		// rerun cancel false scenario

		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/rerun", user2.Name, apiRepo.Name, run2.Index))
		_ = session.MakeRequest(t, req, http.StatusOK)

		run2_2 := unittest.AssertExistsAndLoadBean(t, &actions_model.ActionRun{ID: run2.ID})
		assert.Equal(t, actions_model.StatusWaiting, run2_2.Status)

		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/rerun", user2.Name, apiRepo.Name, run2.Index+1))
		_ = session.MakeRequest(t, req, http.StatusOK)

		task6 := runner.fetchTask(t)
// hunk: @@ -1044,7 +1020,6 @@ jobs:
		// run the workflow with appVersion=v1.21 and cancel=false
		urlStr := fmt.Sprintf("/%s/%s/actions/run?workflow=%s", user2.Name, repo.Name, "workflow-dispatch-concurrency.yml")
		req := NewRequestWithValues(t, "POST", urlStr, map[string]string{
			"ref":        "refs/heads/main",
			"appVersion": "v1.21",
		})
// hunk: @@ -1054,7 +1029,6 @@ jobs:
		assert.Equal(t, "workflow-dispatch-v1.21", run1.ConcurrencyGroup)

		req = NewRequestWithValues(t, "POST", urlStr, map[string]string{
			"ref":        "refs/heads/main",
			"appVersion": "v1.22",
		})
// hunk: @@ -1065,7 +1039,6 @@ jobs:

		// run the workflow with appVersion=v1.22 and cancel=false again
		req = NewRequestWithValues(t, "POST", urlStr, map[string]string{
			"ref":        "refs/heads/main",
			"appVersion": "v1.22",
		})
// hunk: @@ -1075,7 +1048,6 @@ jobs:

		// run the workflow with appVersion=v1.22 and cancel=true
		req = NewRequestWithValues(t, "POST", urlStr, map[string]string{
			"ref":        "refs/heads/main",
			"appVersion": "v1.22",
			"cancel":     "on",
// hunk: @@ -1094,14 +1066,10 @@ jobs:

		// rerun cancel true scenario

		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/jobs/%d/rerun", user2.Name, apiRepo.Name, run2.Index, 1))
		_ = session.MakeRequest(t, req, http.StatusOK)

		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/jobs/%d/rerun", user2.Name, apiRepo.Name, run4.Index, 1))
		_ = session.MakeRequest(t, req, http.StatusOK)

		task5 := runner.fetchTask(t)
// hunk: @@ -1117,17 +1085,13 @@ jobs:

		// rerun cancel false scenario

		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/jobs/%d/rerun", user2.Name, apiRepo.Name, run2.Index, 1))
		_ = session.MakeRequest(t, req, http.StatusOK)

		run2_2 := unittest.AssertExistsAndLoadBean(t, &actions_model.ActionRun{ID: run2.ID})
		assert.Equal(t, actions_model.StatusWaiting, run2_2.Status)

		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/jobs/%d/rerun", user2.Name, apiRepo.Name, run2.Index+1, 1))
		_ = session.MakeRequest(t, req, http.StatusOK)

		task6 := runner.fetchTask(t)
// hunk: @@ -1259,7 +1223,7 @@ func TestWorkflowAndJobConcurrency(t *testing.T) {

		wf1TreePath := ".gitea/workflows/concurrent-workflow-1.yml"
		wf1FileContent := `name: concurrent-workflow-1
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-1.yml'
// hunk: @@ -1281,7 +1245,7 @@ jobs:
`
		wf2TreePath := ".gitea/workflows/concurrent-workflow-2.yml"
		wf2FileContent := `name: concurrent-workflow-2
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-2.yml'
// hunk: @@ -1303,7 +1267,7 @@ jobs:
`
		wf3TreePath := ".gitea/workflows/concurrent-workflow-3.yml"
		wf3FileContent := `name: concurrent-workflow-3
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-3.yml'
// hunk: @@ -1320,7 +1284,7 @@ jobs:

		wf4TreePath := ".gitea/workflows/concurrent-workflow-4.yml"
		wf4FileContent := `name: concurrent-workflow-4
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-4.yml'
// hunk: @@ -1491,9 +1455,7 @@ jobs:
		runner.fetchNoTask(t)

		// cancel the first run
		req := NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/cancel", user2.Name, repo.Name, run1.Index))
		user2Session.MakeRequest(t, req, http.StatusOK)

		// the first run has been cancelled
// hunk: @@ -1525,7 +1487,7 @@ func TestAbandonConcurrentRun(t *testing.T) {

		wf1TreePath := ".gitea/workflows/workflow-1.yml"
		wf1FileContent := `name: Workflow-1
on:
  push:
    paths:
      - '.gitea/workflows/workflow-1.yml'
// hunk: @@ -1544,7 +1506,7 @@ jobs:

		wf2TreePath := ".gitea/workflows/workflow-2.yml"
		wf2FileContent := `name: Workflow-2
on:
  push:
    paths:
      - '.gitea/workflows/workflow-2.yml'
// hunk: @@ -1624,7 +1586,7 @@ func TestRunAndJobWithSameConcurrencyGroup(t *testing.T) {

		wf1TreePath := ".gitea/workflows/concurrent-workflow-1.yml"
		wf1FileContent := `name: concurrent-workflow-1
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-1.yml'
// hunk: @@ -1638,7 +1600,7 @@ jobs:
`
		wf2TreePath := ".gitea/workflows/concurrent-workflow-2.yml"
		wf2FileContent := `name: concurrent-workflow-2
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-2.yml'
// hunk: @@ -1652,7 +1614,7 @@ jobs:
`
		wf3TreePath := ".gitea/workflows/concurrent-workflow-3.yml"
		wf3FileContent := `name: concurrent-workflow-3
on:
  push:
    paths:
      - '.gitea/workflows/concurrent-workflow-3.yml'
