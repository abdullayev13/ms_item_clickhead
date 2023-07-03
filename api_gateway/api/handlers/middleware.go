package handlers

import (
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/api/http"
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/genproto/auth"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckUrl(ctx *gin.Context) {
	uri := ctx.Request.RequestURI

	if uri == "/api/auth/sign-up" || uri == "/api/auth/log-in" {
		ctx.Next()
		return
	}

	token := ctx.GetHeader("Authorization")

	res, err := h.services.AuthService().CheckUri(ctx.Request.Context(),
		&auth.CheckUriRequest{Uri: uri, Token: token})
	if err != nil {
		h.handleResponse(ctx, http.BadRequest, err.Error())
		return
	}

	if !res.Ok {
		h.handleResponse(ctx, http.BadRequest, res.Message)
		return
	}

	ctx.Set("user_id", res.UserId)
}

func (h *Handler) mustGetUserId(ctx *gin.Context) int32 {
	userId, exists := ctx.Get("user_id")
	id, ok := userId.(int32)

	if !exists || !ok {
		h.handleResponse(ctx, http.InternalServerError, "user_id not found!")
		panic("user_id not found!")
	}

	return id
}
