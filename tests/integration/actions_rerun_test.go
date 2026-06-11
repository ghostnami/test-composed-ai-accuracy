// hunk: @@ -33,7 +33,7 @@ func TestActionsRerun(t *testing.T) {

		wfTreePath := ".gitea/workflows/actions-rerun-workflow-1.yml"
		wfFileContent := `name: actions-rerun-workflow-1
on:
  push:
    paths:
      - '.gitea/workflows/actions-rerun-workflow-1.yml'
// hunk: @@ -59,9 +59,7 @@ jobs:
			result: runnerv1.Result_RESULT_SUCCESS,
		})
		// RERUN-FAILURE: the run is not done
		req := NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/rerun", user2.Name, repo.Name, run.Index))
		session.MakeRequest(t, req, http.StatusBadRequest)
		// fetch and exec job2
		job2Task := runner.fetchTask(t)
// hunk: @@ -70,9 +68,7 @@ jobs:
		})

		// RERUN-1: rerun the run
		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/rerun", user2.Name, repo.Name, run.Index))
		session.MakeRequest(t, req, http.StatusOK)
		// fetch and exec job1
		job1TaskR1 := runner.fetchTask(t)
// hunk: @@ -86,9 +82,7 @@ jobs:
		})

		// RERUN-2: rerun job1
		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/jobs/%d/rerun", user2.Name, repo.Name, run.Index, 0))
		session.MakeRequest(t, req, http.StatusOK)
		// job2 needs job1, so rerunning job1 will also rerun job2
		// fetch and exec job1
// hunk: @@ -103,9 +97,7 @@ jobs:
		})

		// RERUN-3: rerun job2
		req = NewRequest(t, "POST", fmt.Sprintf("/%s/%s/actions/runs/%d/jobs/%d/rerun", user2.Name, repo.Name, run.Index, 1))
		session.MakeRequest(t, req, http.StatusOK)
		// only job2 will rerun
		// fetch and exec job2
