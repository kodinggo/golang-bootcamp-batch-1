package main

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func createUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Ini API Get User, id nya: "+id)
}

func getUsers(c echo.Context) error {
	qname := c.QueryParam("name")
	return c.String(http.StatusOK, "Ini API Get Users, query name: "+qname)
}

func updateUser(c echo.Context) error {
	byteBody, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Ini API Update User, body: "+string(byteBody))
}

func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "Ini API Delete User")
}

func main() {
	e := echo.New()

	g := e.Group("/users")
	g.POST("", createUser)
	g.GET("", getUsers)
	g.GET("/:id", getUser)
	g.PUT("/:id", updateUser)
	g.DELETE(":id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
