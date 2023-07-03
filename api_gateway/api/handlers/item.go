package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateArticle(c *gin.Context) {
	/*
			{
			var article item.CreateArticle

			err := c.ShouldBindJSON(&article)
			if err != nil {
				h.handleResponse(c, http.BadRequest, err.Error())
				return
			}

			resp, err := h.services.ItemService().Create(
				c.Request.Context(),
				&article,
			)
			if err != nil {
				h.handleResponse(c, http.GRPCError, err.Error())
				return
			}

			h.handleResponse(c, http.Created, resp)
		}
	*/
}

func (h *Handler) GetArticleById(c *gin.Context) {
	/*
			{

			userId := c.Param("id")

			if !util.IsValidUUID(userId) {
				h.handleResponse(c, http.InvalidArgument, "article id is an invalid uuid")
				return
			}

			resp, err := h.services.ItemService().GetByID(
				context.Background(),
				&item.ArticlePrimaryKey{
					Id: userId,
				},
			)
			if err != nil {
				h.handleResponse(c, http.GRPCError, err.Error())
				return
			}

			h.handleResponse(c, http.OK, resp)
		}
	*/
}

func (h *Handler) GetAllArticles(c *gin.Context) {
	/*
			{
			resp, err := h.services.ItemService().GetList(
				c.Request.Context(),
				&item.GetListArticleRequest{},
			)

			if err != nil {
				h.handleResponse(c, http.GRPCError, err.Error())
				return
			}

			h.handleResponse(c, http.OK, resp)
		}
	*/
}

func (h *Handler) DeleteArticle(c *gin.Context) {
	/*
			{
			id := c.Param("id")

			resp, err := h.services.ItemService().Delete(
				c.Request.Context(),
				&item.ArticlePrimaryKey{
					Id: id,
				},
			)

			if err != nil {
				h.handleResponse(c, http.GRPCError, err.Error())
				return
			}

			h.handleResponse(c, http.OK, resp)
		}
	*/
}

func (h *Handler) UpdateArticle(c *gin.Context) {
	/*
			{
			id := c.Param("id")

			var article item.Article

			err := c.ShouldBindJSON(&article)
			if err != nil {
				h.handleResponse(c, http.BadRequest, err.Error())
				return
			}

			resp, err := h.services.ItemService().Update(
				c.Request.Context(),
				&item.UpdateArticle{
					Id:          id,
					Description: article.Description,
					UserId:      article.UserId,
				},
			)

			if err != nil {
				h.handleResponse(c, http.GRPCError, err.Error())
				return
			}

			h.handleResponse(c, http.OK, resp)
		}
	*/
}

func (h *Handler) PatchArticle(c *gin.Context) {
	/*
			{
			id := c.Param("id")

			var article item.UpdatePatchArticle

			err := c.ShouldBindJSON(&article)
			if err != nil {
				h.handleResponse(c, http.BadRequest, err.Error())
				return
			}

			resp, err := h.services.ItemService().UpdatePatch(
				c.Request.Context(),
				&item.UpdatePatchArticle{
					Id:     id,
					Fields: article.Fields,
				},
			)

			if err != nil {
				h.handleResponse(c, http.GRPCError, err.Error())
				return
			}

			h.handleResponse(c, http.OK, resp)
		}
	*/
}
