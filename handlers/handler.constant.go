package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nutwreck/admin-loker-service/constants"
	"github.com/nutwreck/admin-loker-service/helpers"
)

/**
* =============================================
* Handler Results All Jenis Kelamin Teritory
*==============================================
 */
// GetConstants godoc
// @Summary		Get List Jenis Kelamin
// @Description	Get List Jenis Kelamin
// @Tags		Constant
// @Accept		json
// @Produce		json
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/constant/jenis-kelamin [get]
func HandlerJenisKelamin(ctx *gin.Context) {
	if len(constants.JenisKelamins) == 0 {
		helpers.APIResponse(ctx, "Jenis Kelamin data not found", http.StatusInternalServerError, nil)
		return
	}
	helpers.APIResponse(ctx, "Jenis Kelamin data already to use", http.StatusOK, constants.JenisKelamins)
}

/**
* =============================================
* Handler Results All Status Pernikahan Teritory
*==============================================
 */
// GetConstants godoc
// @Summary		Get List Status Pernikahan
// @Description	Get List Status Pernikahan
// @Tags		Constant
// @Accept		json
// @Produce		json
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/constant/status-pernikahan [get]
func HandlerStatusPernikahan(ctx *gin.Context) {
	if len(constants.StatusPernikahans) == 0 {
		helpers.APIResponse(ctx, "Status Pernikahan data not found", http.StatusInternalServerError, nil)
		return
	}
	helpers.APIResponse(ctx, "Status Pernikahan data already to use", http.StatusOK, constants.StatusPernikahans)
}
