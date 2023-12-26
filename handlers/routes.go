package handlers

import "github.com/labstack/echo/v4"

func SetupRoutes(e *echo.Echo) {
	e.GET("/login", loginHandler)
	e.POST("/login", loginHandler)

	/* ↓ Protected Routes ↓ */
	protectedGroup := e.Group("/protected", authMiddleware)
	protectedGroup.GET("/secret", secretHandler)
	protectedGroup.POST("/logout", logoutHandler)
}
