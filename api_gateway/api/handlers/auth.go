package handlers

import (
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/api/http"
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/genproto/auth"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) SignUp(c *gin.Context) {
	req := new(auth.CreateUser)

	err := c.ShouldBindJSON(req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "error on binding input: "+err.Error())
		return
	}

	req.Role = "user"
	req.GenerateToken = true

	resp, err := h.services.AuthService().Create(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

func (h *Handler) Login(c *gin.Context) {
	req := new(auth.UserLoginRequest)

	err := c.ShouldBindJSON(req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "error on binding input: "+err.Error())
		return
	}

	resp, err := h.services.AuthService().Login(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) GetUserMe(c *gin.Context) {
	req := new(auth.UserPrimaryKey)

	req.Id = h.mustGetUserId(c)

	resp, err := h.services.AuthService().GetByID(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) UpdateUserMe(c *gin.Context) {
	req := new(auth.UpdateUser)

	req.Id = h.mustGetUserId(c)
	err := c.ShouldBindJSON(req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "error on binding input: "+err.Error())
		return
	}

	resp, err := h.services.AuthService().Update(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) DeleteUserMe(c *gin.Context) {
	req := new(auth.UserPrimaryKey)

	req.Id = h.mustGetUserId(c)

	resp, err := h.services.AuthService().Delete(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// apis for amin
func (h *Handler) CreateUser(c *gin.Context) {
	req := new(auth.CreateUser)

	err := c.ShouldBindJSON(req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "error on binding input: "+err.Error())
		return
	}

	req.GenerateToken = false

	resp, err := h.services.AuthService().Create(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

func (h *Handler) GetUserById(c *gin.Context) {
	req := new(auth.UserPrimaryKey)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleResponse(c, http.BadRequest, "id must be number")
		return
	}

	req.Id = int32(id)

	resp, err := h.services.AuthService().GetByID(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	req := new(auth.GetListUserRequest)
	{
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			h.handleResponse(c, http.BadRequest, "limit must be number")
			return
		}
		req.Limit = int64(limit)

		offset, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
			h.handleResponse(c, http.BadRequest, "offset must be number")
			return
		}
		req.Offset = int64(offset)

		req.Order = c.Query("order")
	}
	resp, err := h.services.AuthService().GetList(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	req := new(auth.UpdateUser)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleResponse(c, http.BadRequest, "id must be number")
		return
	}

	err = c.ShouldBindJSON(req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "error on binding input: "+err.Error())
		return
	}

	req.Id = int32(id)

	resp, err := h.services.AuthService().Update(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	req := new(auth.UserPrimaryKey)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleResponse(c, http.BadRequest, "id must be number")
		return
	}

	req.Id = int32(id)

	resp, err := h.services.AuthService().Delete(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
