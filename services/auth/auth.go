// hunk: @@ -19,7 +19,6 @@ import (
	"code.gitea.io/gitea/modules/session"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/web/middleware"
	user_service "code.gitea.io/gitea/services/user"
)

// hunk: @@ -162,9 +161,4 @@ func handleSignIn(resp http.ResponseWriter, req *http.Request, sess SessionStore
	}

	middleware.SetLocaleCookie(resp, user.Language, 0)
}
