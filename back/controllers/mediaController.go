package controllers

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"newsletter-back/entities"
	"newsletter-back/services"
)

type MediaController struct {
	mediaService *services.MediaService
}

func (c *MediaController) RouteSetup(r *gin.RouterGroup) {
	r.GET("", c.index)
	r.POST("", c.create)
	r.GET("/:mediaId", c.show)
	r.PUT("/:mediaId", c.update)
	r.DELETE("/:mediaId", c.delete)
}

func NewMediaController(mediaService *services.MediaService) *MediaController {
	return &MediaController{
		mediaService,
	}
}

func (c *MediaController) index(ctx *gin.Context) {
	medias, err := c.mediaService.GetAllMedia()
	if err != nil {
		log.Println(err)
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusOK,
		Data: medias,
	})
}

func (c *MediaController) show(ctx *gin.Context) {
	mediaId := ctx.Param("mediaId")
	media, err := c.mediaService.GetMediaByID(mediaId)
	if errors.Is(services.ErrRecordNotFound, err) {
		abortWithNotFound(ctx)
		return
	}
	if err != nil {
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusOK,
		Data: media,
	})
}

func (c *MediaController) create(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Println("could not parse the files in request: ", err)
		abortWithInternalServerError(ctx)
		return
	}

	files := form.File["media[]"]
	if len(files) == 0 {
		log.Println("no file in request")
		abortWithValidationError(ctx, []InputError{{
			Field:   "media",
			Message: "the file is missing",
		}})
		return
	}

	file := files[0]
	fileReader, err := file.Open()
	if err != nil {
		log.Println("error opening file: ", err)
		abortWithInternalServerError(ctx)
		return
	}
	defer fileReader.Close()

	ext := filepath.Ext(file.Filename)
	if ext != ".pdf" && ext != ".png" {
		log.Println("invalid file format")
		abortWithValidationError(ctx, []InputError{{
			Field:   "media",
			Message: "only PDF and PNG files are allowed",
		}})
		return
	}

	fileBytes, err := io.ReadAll(fileReader)
	if err != nil {
		log.Println("error reading file content: ", err)
		abortWithInternalServerError(ctx)
		return
	}

	media := &entities.Media{
		FileName:    file.Filename,
		ContentType: file.Header.Get("Content-Type"),
		Content:     bytes.NewReader(fileBytes),
	}

	err = c.mediaService.CreateMedia(media)
	if err != nil {
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusCreated,
		Data: media,
	})
}

func (c *MediaController) update(ctx *gin.Context) {
	mediaId := ctx.Param("mediaId")

	form, err := ctx.MultipartForm()
	if err != nil {
		log.Println("could not parse the files in request: ", err)
		abortWithInternalServerError(ctx)
		return
	}

	files := form.File["media[]"]
	if len(files) == 0 {
		log.Println("no file in request")
		abortWithValidationError(ctx, []InputError{{
			Field:   "media",
			Message: "the file is missing",
		}})
		return
	}

	file := files[0]
	fileReader, err := file.Open()
	if err != nil {
		log.Println("error opening file: ", err)
		abortWithInternalServerError(ctx)
		return
	}
	defer fileReader.Close()

	fileBytes, err := io.ReadAll(fileReader)
	if err != nil {
		log.Println("error reading file content: ", err)
		abortWithInternalServerError(ctx)
		return
	}

	media := entities.Media{
		ID:          mediaId,
		FileName:    file.Filename,
		ContentType: file.Header.Get("Content-Type"),
		Content:     bytes.NewReader(fileBytes),
	}

	err = c.mediaService.UpdateMedia(&media)
	if errors.Is(services.ErrRecordNotFound, err) {
		abortWithNotFound(ctx)
		return
	}
	if err != nil {
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusOK,
		Data: media,
	})
}

func (c *MediaController) delete(ctx *gin.Context) {
	mediaId := ctx.Param("mediaId")

	err := c.mediaService.DeleteMediaByID(mediaId)
	if errors.Is(services.ErrRecordNotFound, err) {
		abortWithNotFound(ctx)
		return
	}
	if err != nil {
		abortWithInternalServerError(ctx)
		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Code: http.StatusNoContent,
	})
}
