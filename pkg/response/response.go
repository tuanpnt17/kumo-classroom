package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DataResponse struct {
	Code    int         `json:"code" validate:"required"`
	Message string      `json:"message" validate:"required"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, DataResponse{
		Code:    code,
		Message: message[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int) {
	c.JSON(http.StatusOK, DataResponse{
		Code:    code,
		Message: message[code],
	})
}
