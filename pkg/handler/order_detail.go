package handler

import (
	"jatis/pkg/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

func GetOrderDetail(c echo.Context) error {

	orderID, err := strconv.ParseInt(c.Param("orderID"), 10, 64)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Order ID",
		})
	}

	res, err := service.GetOrderDetails(orderID)
	if err != nil {
		log.Println(err)
		if err == service.ErrOrderNotFound {
			return c.JSON(http.StatusNotFound, ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "Order ID not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, res)
}
