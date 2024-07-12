package main

import (
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

	e.POST("/users", postUser)

	e.Logger.Fatal(e.Start(":8000"))
}

func postUser(c echo.Context) error {
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
