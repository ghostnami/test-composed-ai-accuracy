// hunk: @@ -22,5 +22,5 @@ func addOwnerRepoGitHTTPRouters(m *web.Router) {
		m.Methods("GET,OPTIONS", "/objects/{head:[0-9a-f]{2}}/{hash:[0-9a-f]{38,62}}", repo.GetLooseObject)
		m.Methods("GET,OPTIONS", "/objects/pack/pack-{file:[0-9a-f]{40,64}}.pack", repo.GetPackFile)
		m.Methods("GET,OPTIONS", "/objects/pack/pack-{file:[0-9a-f]{40,64}}.idx", repo.GetIdxFile)
	}, repo.HTTPGitEnabledHandler, repo.CorsHandler(), optSignInFromAnyOrigin, context.UserAssignmentWeb())
}
