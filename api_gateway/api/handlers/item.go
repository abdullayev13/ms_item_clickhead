package handlers

import (
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/api/http"
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/genproto/item"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) CreateArticle(c *gin.Context) {
	req := new(item.CreateItem)

	err := c.ShouldBindJSON(req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "error on binding input: "+err.Error())
		return
	}

	resp, err := h.services.ItemService().Create(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

func (h *Handler) GetArticleById(c *gin.Context) {
	req := new(item.ItemPrimaryKey)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleResponse(c, http.BadRequest, "id must be number")
		return
	}

	req.Id = int32(id)

	resp, err := h.services.ItemService().GetByID(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) GetAllArticles(c *gin.Context) {
	req := new(item.GetListItemRequest)
	{
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil {
			h.handleResponse(c, http.BadRequest, "limit must be number")
			return
		}
		req.Limit = int64(limit)

		offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
		if err != nil {
			h.handleResponse(c, http.BadRequest, "offset must be number")
			return
		}
		req.Offset = int64(offset)

		req.Order = c.Query("order")
	}
	resp, err := h.services.ItemService().GetList(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) UpdateArticle(c *gin.Context) {
	req := new(item.UpdateItem)

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

	resp, err := h.services.ItemService().Update(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) DeleteArticle(c *gin.Context) {
	req := new(item.ItemPrimaryKey)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleResponse(c, http.BadRequest, "id must be number")
		return
	}

	req.Id = int32(id)

	resp, err := h.services.ItemService().Delete(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
