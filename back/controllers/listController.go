package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"newsletter-back/entities"
	"newsletter-back/services"
)

type ListController struct {
	listService *services.ListService
}

func (c *ListController) RouteSetup(rg *gin.RouterGroup) {
	rg.GET("/", c.index)
	rg.GET("/:listId", c.show)
	rg.POST("/", c.create)
	rg.PUT("/:listId", c.update)
	rg.DELETE("/:listId", c.delete)
}

func NewListController(listService *services.ListService) *ListController {
	return &ListController{
		listService,
	}
}

func (c *ListController) index(ctx *gin.Context) {
	lists, err := c.listService.GetAllLists()
	if err != nil {
		log.Println(err)
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusOK,
		Data: lists,
	})
}

func (c *ListController) show(ctx *gin.Context) {
	listId := ctx.Param("listId")
	list, err := c.listService.GetList(listId)
	if err != nil {
		log.Println("list does not exist: ", err)
		abortWithNotFound(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusOK,
		Data: list,
	})
}

func (c *ListController) create(ctx *gin.Context) {
	var list entities.List
	if errs := validateRequest(ctx, &list); errs != nil {
		abortWithValidationError(ctx, errs)
		return
	}

	if err := c.listService.CreateList(&list); err != nil {
		log.Println("error creating list: ", err)
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusCreated,
		Data: list,
	})
}

func (c *ListController) update(ctx *gin.Context) {
	listId := ctx.Param("listId")

	var updatedList entities.List
	if errs := validateRequest(ctx, &updatedList); errs != nil {
		abortWithValidationError(ctx, errs)
		return
	}

	updatedList.ID = listId

	if err := c.listService.UpdateList(&updatedList); err != nil {
		log.Println("error updating list: ", err)
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusOK,
		Data: updatedList,
	})
}

func (c *ListController) delete(ctx *gin.Context) {
	listId := ctx.Param("listId")
	if err := c.listService.DeleteList(listId); err != nil {
		log.Println("error deleting list: ", err)
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusNoContent,
	})
}
