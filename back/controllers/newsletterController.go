package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"newsletter-back/entities"
	"newsletter-back/services"
)

type NewsletterController struct {
	newsletterService *services.NewsletterService
	emailService      services.EmailService
}

func (c *NewsletterController) RouteSetup(rg *gin.RouterGroup) {
	rg.GET("", c.index)
	rg.POST("", c.create)
	rg.GET("/:newsletterId", c.show)
	rg.PUT("/:newsletterId", c.update)
	rg.DELETE("/:newsletterId", c.delete)
}

func NewNewsletterController(newsletterService *services.NewsletterService, emailService services.EmailService) *NewsletterController {
	return &NewsletterController{
		newsletterService,
		emailService,
	}
}

func (c *NewsletterController) index(ctx *gin.Context) {
	newsletters, err := c.newsletterService.GetAllNewsletters()
	if err != nil {
		log.Println(err)

		abortWithInternalServerError(ctx)

		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusOK,
		Data: newsletters,
	})
}

func (c *NewsletterController) show(ctx *gin.Context) {
	newsletterId := ctx.Param("newsletterId")
	newsletter, err := c.newsletterService.GetNewsletter(newsletterId)
	if err != nil {
		log.Println("newsletter does not exists: ", err)

		abortWithNotFound(ctx)

		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusOK,
		Data: newsletter,
	})
}

type CreateNewsletter struct {
	Subject     string    `json:"subject" binding:"required"`
	Content     string    `json:"content" binding:"required"`
	SendAt      time.Time `json:"send_at" binding:"required"`
	Recipients  []string  `json:"recipients" binding:"required"`
	Attachments []string  `json:"attachments"`
}

func (c *NewsletterController) create(ctx *gin.Context) {
	var requestBody CreateNewsletter
	if errs := validateRequest(ctx, &requestBody); errs != nil {
		abortWithValidationError(ctx, errs)

		return
	}

	recipients := make([]entities.Recipient, len(requestBody.Recipients))
	for index, recipient := range requestBody.Recipients {
		recipients[index] = entities.Recipient{
			ID: recipient,
		}
	}

	attachments := make([]entities.Media, len(requestBody.Attachments))
	for index, attachment := range requestBody.Attachments {
		attachments[index] = entities.Media{
			ID: attachment,
		}
	}

	newsletter := entities.Newsletter{
		Subject:     requestBody.Subject,
		Content:     requestBody.Content,
		SendAt:      requestBody.SendAt,
		Recipients:  recipients,
		Attachments: attachments,
	}

	err := c.newsletterService.CreateNewsletter(&newsletter)
	if err != nil {
		log.Println("error creating newsletter: ", err)
		abortWithInternalServerError(ctx)
		return
	}

	nw, err := c.newsletterService.GetNewsletter(newsletter.ID)

	emailRecipients := make([]string, len(nw.Recipients))
	for index, recipient := range nw.Recipients {
		emailRecipients[index] = recipient.Email
	}

	emailAttachments := make([]services.Attachment, len(nw.Attachments))
	for index, attachment := range nw.Attachments {
		emailAttachments[index] = services.Attachment{
			Filename: attachment.FileName,
			Location: attachment.Location,
		}
	}

	log.Println(emailRecipients, newsletter.Subject, newsletter.Content, emailAttachments)

	err = c.emailService.SendEmail(emailRecipients, newsletter.Subject, newsletter.Content, emailAttachments)
	if err != nil {
		log.Println("error sending email: ", err)
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusCreated,
		Data: newsletter,
	})
}

func (c *NewsletterController) update(ctx *gin.Context) {
	newsletterId := ctx.Param("newsletterId")

	var updatedNewsletter entities.Newsletter
	if errs := validateRequest(ctx, &updatedNewsletter); errs != nil {
		abortWithValidationError(ctx, errs)

		return
	}

	updatedNewsletter.ID = newsletterId

	if err := c.newsletterService.UpdateNewsletter(&updatedNewsletter); err != nil {
		log.Println("error updating newsletter: ", err)
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusOK,
		Data: updatedNewsletter,
	})
}

func (c *NewsletterController) delete(ctx *gin.Context) {
	newsletterId := ctx.Param("newsletterId")
	if err := c.newsletterService.DeleteNewsletter(newsletterId); err != nil {
		log.Println("error deleting newsletter: ", err)
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusNoContent,
	})
}
