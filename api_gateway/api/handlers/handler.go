package handlers

import (
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/api/http"
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/config"
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/grpc/client"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg      config.Config
	services client.ServiceManagerI
}

func NewHandler(cfg config.Config, svcs client.ServiceManagerI) Handler {
	return Handler{
		cfg:      cfg,
		services: svcs,
	}
}

func (h *Handler) handleResponse(c *gin.Context, status http.Status, data interface{}) {
	c.JSON(status.Code, data)
}

func (h *Handler) handleErrorResponse(c *gin.Context, code int, message string, err interface{}) {
	c.JSON(code, ResponseModel{
		Code:    code,
		Message: message,
		Error:   err,
	})
}

func (h *Handler) handleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, ResponseModel{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
