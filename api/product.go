package api

import (
	"errors"
	"net/http"

	db "github.com/QuyenFunc/Goaudio/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createProductRequest struct {
	CatalogID int64  `json:"catalog_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Price     int64  `json:"price" binding:"required"`
	Discount  string `json:"discount" binding:"required"`
	Content   string `json:"content" binding:"required"`
	ImageLink string `json:"imagelink" binding:"required"`
	ImageList string `json:"imagelist" binding:"required"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProducteParams{
		CatalogID: req.CatalogID,
		Name:      req.Name,
		Price:     req.Price,
		View:      0,
		Discount:  req.Discount,
		Content:   req.Content,
		ImageLink: req.ImageLink,
		ImageList: req.ImageList,
	}

	proruct, err := server.store.CreateProducte(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation || errCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, proruct)
}

type getProductRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getProduct(ctx *gin.Context) {
	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := server.store.GetProducte(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)

}

type listProductRequest struct {
	Name     string `json:"name" binding:"required"`
	PageID   int32  `form:"page_id" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listProduct(ctx *gin.Context) {
	var req listProductRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListProductesParams{
		Name:   req.Name,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	products, err := server.store.ListProductes(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, products)
}
