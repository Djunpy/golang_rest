package controllers

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	db "golang_rest_app/db/sqlc"
	"golang_rest_app/schemas"
	"net/http"
	"strconv"
)

type PostController struct {
	ctx context.Context
	db  *db.Queries
}

func NewPostController(ctx context.Context, db *db.Queries) *PostController {
	return &PostController{ctx, db}
}

func (ac *PostController) CreatePost(ctx *gin.Context) {
	var payload schemas.CreatePost

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	args := &db.CreatePostParams{
		Title:    payload.Title,
		Category: payload.Category,
		Content:  payload.Content,
		UserID:   1,
	}

	post, err := ac.db.CreatePost(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "post": post})
}

func (ac *PostController) UpdatePost(ctx *gin.Context) {
	var payload *schemas.UpdatePost
	postIdStr := ctx.Param("postId")
	id, err := strconv.Atoi(postIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	args := &db.UpdatePostParams{
		ID:       int32(id),
		Title:    sql.NullString{String: payload.Title, Valid: payload.Title != ""},
		Category: sql.NullString{String: payload.Category, Valid: payload.Category != ""},
		Content:  sql.NullString{String: payload.Content, Valid: payload.Content != ""},
	}

	post, err := ac.db.UpdatePost(ctx, *args)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that ID exists"})
			return

		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "post": post})
}

func (ac *PostController) GetPostById(ctx *gin.Context) {
	postIdStr := ctx.Param("postId")
	id, _ := strconv.Atoi(postIdStr)
	post, err := ac.db.GetPostById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that ID exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "post": post})
}

func (ac *PostController) GetAllPosts(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	initPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (initPage - 1) * intLimit

	args := &db.ListPostsParams{
		Limit:  int32(intLimit),
		Offset: int32(offset),
	}

	posts, err := ac.db.ListPosts(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if posts == nil {
		posts = []db.Post{}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(posts), "data": posts})
}

func (ac *PostController) DeletePostById(ctx *gin.Context) {
	postId := ctx.Param("postId")
	id, _ := strconv.Atoi(postId)
	_, err := ac.db.GetPostById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that ID exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}
	err = ac.db.DeletePost(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"status": "success"})
}
