// File path: api/handler.go

package api

import (
	"github.com/achimgrolimund/csp-report-api/pkg/domain/csp_report"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *csp_report.ReportService
}

func NewHandler(service *csp_report.ReportService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SaveReportHandler(c *gin.Context) {
	var report csp_report.Report

	if err := c.BindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.SaveReport(c.Request.Context(), &report); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
