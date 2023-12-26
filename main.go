package main

import (
	"github.com/emarifer/echo-cookie-session-demo/handlers"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// In production, the secret key of the CookieStore
// would be obtained from a .env file
const secret_key = "secret"

func main() {

	e := echo.New()

	e.Static("/", "assets")

	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	// Helpers Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// CORS Middleware
	/* e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	})) */

	// Session Middleware
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(secret_key))))

	// Setting Routes
	handlers.SetupRoutes(e)

	// Start Server
	e.Logger.Fatal(e.Start(":8082"))
}

/*
https://gist.github.com/taforyou/544c60ffd072c9573971cf447c9fea44
https://gist.github.com/mhewedy/4e45e04186ed9d4e3c8c86e6acff0b17

https://github.com/CurtisVermeeren/gorilla-sessions-tutorial

func SkipperFn(skipURLs []string) func(echo.Context) bool {
	return func(context echo.Context) bool {
		for _, url := range skipURLs {
			if url == context.Request().URL.String() {
				return true
			}
		}
		return false
	}
}
*/
