package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type InputError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "It must be less than " + fe.Param()
	case "gte":
		return "It must be greater than " + fe.Param()
	case "eqfield":
		return "It does not match with " + fe.Param()
	case "email":
		return "The email format is not valid"
	case "number":
		return "It is not a valid number"
	case "len":
		return "The field must have at least of " + fe.Param() + " length"
	case "min":
		return "The field must have at least of " + fe.Param() + " length"
	case "max":
		return "The field must have a maximum of" + fe.Param() + " length"
	case "alphanum":
		return "Can only contain letters and numbers"
	case "alpha":
		return "Can only contain letters"
	case "e164":
		return "Phone number with invalid format"
	case "password":
		return "Minimum eight characters, maximum fifteen characters, at least one letter and one number"
	}

	return "Unknown error"
}

func validateRequest(ctx *gin.Context, input any) []InputError {
	if err := ctx.ShouldBind(&input); err != nil {
		var myError *json.UnmarshalTypeError

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errs := make([]InputError, len(ve))
			for i, fe := range ve {
				errs[i] = InputError{ToSnakeCase(fe.Field()), getErrorMsg(fe)}
			}

			return errs
		} else if errors.As(err, &myError) {
			log.Println("Error mapping request body: ", myError.Error())

			errs := make([]InputError, 1)
			errs[0] = InputError{myError.Field, "Formato errado"}

			return errs
		} else {
			log.Println("Error mapping request body: ", myError.Error())

			errs := make([]InputError, 1)
			return errs
		}
	}

	return nil
}

func abortWithValidationError(ctx *gin.Context, errs []InputError) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error":   true,
		"message": "validation errors",
		"errors":  errs,
	})
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

type SuccessResponseOptions struct {
	Message string
	Code    int
	Data    interface{}
}

func successResponse(ctx *gin.Context, opts SuccessResponseOptions) {
	response := gin.H{"error": false}

	if opts.Message != "" {
		response["message"] = opts.Message
	}

	if opts.Data != nil {
		response["data"] = opts.Data
	}

	ctx.JSON(opts.Code, response)
}

func abortWithInternalServerError(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"error":   true,
		"message": "Something went wrong. Try again.",
	})
}

func abortWithNotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"error":   true,
		"message": "element not found",
	})
}

func ToID(ID any) uint {
	u64, _ := strconv.ParseUint(fmt.Sprintf("%v", ID), 10, 32)
	return uint(u64)
}
