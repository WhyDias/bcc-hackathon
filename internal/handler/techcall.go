package handler

import (
	"bcc-hackathon-go/pkg/modules"
	monitoring "bcc-hackathon-go/pkg/moniroting"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) TechCall(c *gin.Context) {
	var jsonInput modules.Request
	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		var response modules.Response
		response.Status = false
		response.Message = err.Error()
		logrus.Error(err)
		monitoring.ErrorHandler.With(prometheus.Labels{"error_message": err.Error()}).Inc()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	code, any := h.services.TechCallLogic(jsonInput)

	if code != 200 {
		logrus.Error()
		monitoring.ErrorHandler.With(prometheus.Labels{"error_message": "error"}).Inc()
	}

	c.JSON(code, any)
}
