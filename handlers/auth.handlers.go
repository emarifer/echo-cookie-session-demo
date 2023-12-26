package handlers

import (
	// "fmt"
	"net/http"
	"os/exec"

	"github.com/emarifer/echo-cookie-session-demo/views/user"

	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	auth_sessions_key string = "authenticate-sessions"
	auth_key          string = "authenticated"
	user_id_key       string = "user_id"
)

func loginHandler(c echo.Context) error {
	if c.Request().Method == "POST" {
		sess, _ := session.Get(auth_sessions_key, c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   3600, // in seconds
			HttpOnly: true,
		}
		// Authentication goes here
		// ...

		// Set user as authenticated and their ID example
		uuid, _ := exec.Command("uuidgen").Output() // generate UUID from OS
		sess.Values = map[interface{}]interface{}{
			auth_key:    true,
			user_id_key: string(uuid), // e.g.
		}
		sess.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusSeeOther, "/protected/secret")
	}

	return renderView(c, user.LoginIndex("| Login", user.Login()))
}

func secretHandler(c echo.Context) error {
	data := echo.Map{
		"content": "You're able to read a secret",
		"user_id": c.Get(user_id_key).(string), // get user_id from context
	}

	return renderView(c, user.SecretIndex(
		"| Protected Page",
		user.Secret(data),
	))
}

func logoutHandler(c echo.Context) error {
	sess, _ := session.Get(auth_sessions_key, c)
	// Revoke users authentication
	sess.Values = map[interface{}]interface{}{
		auth_key:    false,
		user_id_key: "",
	}
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusSeeOther, "/login")
}

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get(auth_sessions_key, c)
		if auth, ok := sess.Values[auth_key].(bool); !ok || !auth {
			// fmt.Println(ok, auth)
			return echo.NewHTTPError(echo.ErrUnauthorized.Code, "Please provide valid credentials")
		}

		if userId, ok := sess.Values[user_id_key].(string); ok && len(userId) != 0 {
			// fmt.Println(ok, userId)
			c.Set(user_id_key, userId) // set the user_id in the context
		}

		return next(c)
	}
}
