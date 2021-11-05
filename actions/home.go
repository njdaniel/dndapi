package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
// @Summary Home
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}
