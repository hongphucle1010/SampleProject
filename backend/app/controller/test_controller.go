package controller

import (
	"sms/app/service"
	"sms/pkg/response"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type TestController struct {
	Ctx         iris.Context
	TestService service.ITestService
}

// @Summary Experimenting API
// @Description This API is just used for experimenting
// @Tags test
// @Produce  json
// @Success 200 {object} response.SuccessResponse[any]
// @Failure 500 {object} response.ErrorResponse
// @Router /test/ [get]
func (c *TestController) Get() mvc.Result {
	test, err := c.TestService.GetTest()
	if err != nil {
		return response.ErrorResponse{
			Code:    iris.StatusInternalServerError,
			Message: "Failed to fetch test",
			Details: err.Error(),
		}
	}
	return response.SuccessResponse[any]{
		Code:    iris.StatusOK,
		Message: "Successfully fetched test",
		Data:    test,
	}
}
