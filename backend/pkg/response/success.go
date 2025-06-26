package response

import "github.com/kataras/iris/v12"

// SuccessResponse represents a successful API response.
type SuccessResponse[T any] struct {
	Code    int    `json:"-"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func (s SuccessResponse[T]) Dispatch(ctx iris.Context) {
	ctx.StatusCode(s.Code)
	ctx.JSON(iris.Map{
		"status":  "success",
		"message": s.Message,
		"data":    s.Data,
	})
}
