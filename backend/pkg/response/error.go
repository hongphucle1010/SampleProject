package response

import "github.com/kataras/iris/v12"

// ErrorResponse represents a failed API response.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e ErrorResponse) Dispatch(ctx iris.Context) {
    ctx.StatusCode(e.Code)
    ctx.JSON(iris.Map{
        "status":  "error",
        "message": e.Message,
        "details": e.Details,
    })
}
