package rest

import (
	"includemy/model"
	"includemy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) UserJoinPremiumCourse(ctx *gin.Context) {
	param := model.CreateUserJoinCourse{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	result, err := r.service.PaymentService.GetPaymentCourse(&param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get url payment", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Success to get url payment", result)

}

func (r *Rest) Callback(ctx *gin.Context) {
	var notificationPayload map[string]interface{}

	err := ctx.ShouldBind(&notificationPayload)
	if err != nil {
		return
	}

	_, exists := notificationPayload["order_id"].(string)
	if !exists {
		return
	}

	r.service.PaymentService.CallbackCourse(notificationPayload)
}

func (r *Rest) UserJoinPremiumSertif(ctx *gin.Context) {
	param := model.CreateSertificationUser{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	result, err := r.service.PaymentService.GetPaymentSertif(&param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get url payment", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Success to get url payment", result)

}

func (r *Rest) CallbackSertif(ctx *gin.Context) {
	var notificationPayload map[string]interface{}

	err := ctx.ShouldBind(&notificationPayload)
	if err != nil {
		return
	}

	_, exists := notificationPayload["order_id"].(string)
	if !exists {
		return
	}

	r.service.PaymentService.CallbackSertif(notificationPayload)
}
