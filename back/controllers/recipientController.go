package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"newsletter-back/entities"
	"newsletter-back/services"
)

type RecipientController struct {
	recipientService *services.RecipientService
}

func (c *RecipientController) RouteSetup(rg *gin.RouterGroup) {
	rg.GET("", c.index)
	rg.POST("", c.create)
	rg.GET("/:recipientId", c.show)
	rg.PUT("/:recipientId", c.update)
	rg.DELETE("/:recipientId", c.delete)
}

func NewRecipientController(recipientService *services.RecipientService) *RecipientController {
	return &RecipientController{
		recipientService,
	}
}

func (c *RecipientController) index(ctx *gin.Context) {
	subscribers, err := c.recipientService.GetAllSubscribers()
	if err != nil {
		log.Println(err)
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusOK,
		Data: subscribers,
	})
}

func (c *RecipientController) show(ctx *gin.Context) {
	recipientId := ctx.Param("recipientId")
	recipient, err := c.recipientService.GetRecipientByID(recipientId)
	if err != nil {
		log.Println("recipient does not exist: ", err)
		abortWithNotFound(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusOK,
		Data: recipient,
	})
}

type CreateRecipientRequest struct {
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"full_name" binding:"required"`
}

func (c *RecipientController) create(ctx *gin.Context) {
	var requestBody []CreateRecipientRequest

	errs := validateRequest(ctx, &requestBody)
	if errs != nil {
		abortWithValidationError(ctx, errs)
		return
	}

	log.Println("requestBody: ", requestBody)

	if len(requestBody) == 0 {
		abortWithValidationError(ctx, []InputError{{
			Field:   "recipients",
			Message: "recipients are required",
		}})
		return
	}

	recipients := make([]*entities.Recipient, len(requestBody))
	for index, request := range requestBody {
		recipients[index] = &entities.Recipient{
			FullName: request.FullName,
			Email:    request.Email,
		}
	}

	err := c.recipientService.CreateRecipients(recipients)
	if err != nil {
		log.Println("error creating recipient: ", err)
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusCreated,
		Data: recipients,
	})
}

func (c *RecipientController) update(ctx *gin.Context) {
	subscriberId := ctx.Param("subscriberId")

	var updatedSubscriber entities.Recipient
	errs := validateRequest(ctx, &updatedSubscriber)
	if errs != nil {
		abortWithValidationError(ctx, errs)
		return
	}

	updatedSubscriber.ID = subscriberId

	err := c.recipientService.UpdateSubscriber(&updatedSubscriber)
	if err != nil {
		log.Println("error updating subscriber: ", err)
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusOK,
		Data: updatedSubscriber,
	})
}

func (c *RecipientController) delete(ctx *gin.Context) {
	subscriberId := ctx.Param("subscriberId")
	err := c.recipientService.DeleteSubscriberByID(subscriberId)
	if err != nil {
		log.Println("error deleting subscriber: ", err)
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusNoContent,
	})
}
