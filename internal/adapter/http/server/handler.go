package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// POST
// paymentsReq
func (ct Controller) payments(c *gin.Context) {
	// ctx := c.Request.Context()

	var payload PaymentsRequest

	if err := c.ShouldBind(&payload); err != nil {
		zap.L().Error("bad request, could not bind json", zap.Error(err))

		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}

// optional
// GET /payments-summary?from=2020-07-10T12:34:56.000Z&to=2020-07-10T12:35:56.000Z
// from=2020-07-10T12:34:56.000Z
// to=2020-07-10T12:35:56.000Z
func (ct Controller) paymentsSummary(c* gin.Context) {
	// ctx := c.Request.Context()

	// trying to get from and to query params
	from := c.Query("from")
	to := c.Query("to")

	var response PaymentsSummaryResponse
	
	summary, err := ct.Service.GeneratePaymentsSummary(c.Request.Context(), from, to)
	if err != nil {
		response.Default = Default(summary.Default)
		response.Fallback = Fallback(summary.Fallback)

		c.JSON(200, response)
		return
	}

	zap.L().Error("could not generate payment summary", zap.String("from", from), zap.String("to", to))

	c.JSON(http.StatusInternalServerError, nil)
}