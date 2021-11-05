package actions

import (
	"dndapi/models"
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/x/responder"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Character)
// DB Table: Plural (characters)
// Resource: Plural (Characters)
// Path: Plural (/characters)
// View Template Folder: Plural (/templates/characters/)

// CharactersResource is the resource for the Character model
type CharactersResource struct {
	buffalo.Resource
}

// List gets all Characters. This function is mapped to the path
// GET /characters// ListAccounts godoc
// @Summary List characters
// @Description get characters
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} models.Characters
// @Router /characters [get]
func (v CharactersResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	characters := &models.Characters{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Characters from the DB
	if err := q.All(characters); err != nil {
		return err
	}

	//return responder.Wants("json", func (c buffalo.Context) error {
	//  // Add the paginator to the context so it can be used in the template.
	//  c.Set("pagination", q.Paginator)

	//  c.Set("characters", characters)
	//  return c.Render(http.StatusOK, r.HTML("/characters/index.plush.html"))
	//}).Wants("json", func (c buffalo.Context) error {
	//  return c.Render(200, r.JSON(characters))
	//}).Wants("xml", func (c buffalo.Context) error {
	//  return c.Render(200, r.XML(characters))
	//}).Respond(c)
	return c.Render(http.StatusOK, r.JSON(characters))
}

// Show gets the data for one Character. This function is mapped to
// the path GET /characters/{character_id}
func (v CharactersResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Character
	character := &models.Character{}

	// To find the Character the parameter character_id is used.
	if err := tx.Find(character, c.Param("character_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return c.Render(200, r.JSON(character))
}

// Create adds a Character to the DB. This function is mapped to the
// path POST /characters
func (v CharactersResource) Create(c buffalo.Context) error {
	// Allocate an empty Character
	character := &models.Character{}

	// Bind character to the html form elements
	if err := c.Bind(character); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(character)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		//Render errors as json
		return c.Render(400, r.JSON(verrs))
	}

	return c.Render(201, r.JSON(character))
}

// Update changes a Character in the DB. This function is mapped to
// the path PUT /characters/{character_id}
func (v CharactersResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Character
	character := &models.Character{}

	if err := tx.Find(character, c.Param("character_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Character to the html form elements
	if err := c.Bind(character); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(character)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("character", character)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("/characters/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "character.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/characters/%v", character.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(character))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(character))
	}).Respond(c)
}

// Destroy deletes a Character from the DB. This function is mapped
// to the path DELETE /characters/{character_id}
func (v CharactersResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Character
	character := &models.Character{}

	// To find the Character the parameter character_id is used.
	if err := tx.Find(character, c.Param("character_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(character); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "character.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/characters")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(character))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(character))
	}).Respond(c)
}
