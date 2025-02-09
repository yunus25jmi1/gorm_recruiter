package handlers

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *EmployerHandler) GetApplicantData(c *gin.Context) {
	//! Get Applicant ID
	applicantID := c.Query("applicant_id")
	if applicantID == "" {
		response := models.Response{Message: "Applicant ID is required"}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//! Fetch data from DB
	applicant, err := h.repo.FetchApplicantByID(applicantID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response := models.Response{Message: "Record Not found"}
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := models.Response{Message: "Error fetching applicant"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := models.Response{
		Message: "Applicant fetched successfully!",
		Data:    applicant,
	}
	c.JSON(http.StatusOK, response)
}
