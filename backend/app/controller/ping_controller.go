package controller

import (
	"sample/app/service"
	"sample/pkg/response"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type PingController struct {
	Ctx         iris.Context
	PingService service.IPingService
}

// @Summary Health check
// @Description Ping to get pong
// @Tags ping
// @Produce  json
// @Success 200 {object} response.SuccessResponse[any]
// @Failure 500 {object} response.ErrorResponse
// @Router /ping/ [get]
func (c *PingController) Get() mvc.Result {
	message, err := c.PingService.Pong()
	if err != nil {
		return response.ErrorResponse{
			Code:    iris.StatusInternalServerError,
			Message: "Failed to ping",
			Details: err.Error(),
		}
	}
	return response.SuccessResponse[any]{
		Code:    iris.StatusOK,
		Message: "Successfully pinged",
		Data:    iris.Map{"message": message},
	}
}
