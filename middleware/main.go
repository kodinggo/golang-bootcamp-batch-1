package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ReqParam struct {
	Id string `param:"id"`
}

type ReqQuery struct {
	Filter string `query:"filter"`
	Sort   string `query:"sort"`
}

type ReqBody struct {
	Message string `json:"message"`
}

type ReqHeader struct {
	Authorization string `header:"Authorization"`
}

type HTTPResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/users/:id", func(c echo.Context) error {
		reqParam := new(ReqParam)
		if err := c.Bind(reqParam); err != nil {
			return err
		}

		reqQuery := new(ReqQuery)
		if err := c.Bind(reqQuery); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"param":  reqParam,
			"filter": reqQuery.Filter,
			"sort":   reqQuery.Sort,
		})
	})

	protected := e.Group("/admin")
	protected.Use(middlewareWithContext)

	protected.POST("/users", postUser)
	protected.PATCH("/users", postUser)
	protected.DELETE("/users", postUser)

	e.Logger.Fatal(e.Start(":8000"))
}

func postUser(c echo.Context) error {
	fmt.Println("Handler start")
	var reqBody ReqBody
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, HTTPResponse{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
			Success: false,
			Data:    nil,
		})
	}

	var reqHeader ReqHeader

	err := (&echo.DefaultBinder{}).BindHeaders(c, &reqHeader)
	if err != nil {
		return c.JSON(http.StatusBadRequest, HTTPResponse{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
			Success: false,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, HTTPResponse{
		Message: "Success",
		Status:  http.StatusOK,
		Success: true,
		Data: map[string]interface{}{
			"message":       reqBody.Message,
			"Authorization": reqHeader.Authorization,
			"User-Session":  c.Request().Context().Value("User-Session"),
		},
	})
}

type UserSession struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func middlewareWithContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get value of Header Authorization
		auth := c.Request().Header.Get("Authorization")

		// check if value of Header Authorization is NOT equal to "ValidToken"
		// "ValidToken" is a dummy token, you can change it to your own token
		// but should send equally in request header
		// "Authorization": "ValidToken/YourTokenHere"
		if auth != "ValidToken" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		session := UserSession{
			Id:       "1",
			Username: "joe",
			Role:     "admin",
		}

		// if value "ValidToken", then set request to context
		// WithValue accept parent context and key-value pair
		parentCtx := c.Request().Context()
		ctx := context.WithValue(parentCtx, "User-Session", session)
		req := c.Request().WithContext(ctx)

		c.SetRequest(req)

		// continue to request
		return next(c)
	}
}
