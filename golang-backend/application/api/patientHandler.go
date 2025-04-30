package api

import (
	entity "demo/application/internal"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Server) addPatient(c *gin.Context) {
	var patient entity.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}
	ctx := c.Request.Context()
	id, err := s.patientService.CreatePatient(ctx, patient)
	if err != nil {
		s.logger.Error("Failed to create patient", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create patient"})
		return
	}
	s.logger.Info("Patient Inserted", zap.String("Patient id", id))
}

func (s *Server) updatePatientStatus(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		NewStatus string `json:"new_status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	ctx := c.Request.Context()
	patient, err := s.patientService.UpdatePatientStatus(ctx, id, req.NewStatus)
	if err != nil {
		s.logger.Error("Failed to update patient status", zap.String("PatientID", id), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update patient status"})
		return
	}
	s.logger.Info("Patient Status Updated", zap.String("PatientID", id))
	c.JSON(http.StatusOK, patient)
}

func (s *Server) getPatient(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	patient, err := s.patientService.GetPatient(ctx, id)
	if err != nil {
		s.logger.Error("Failed to get patient", zap.String("PatientID", id), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch patient"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func (s *Server) listPatients(c *gin.Context) {
	ctx := c.Request.Context()

	patients, err := s.patientService.ListPatients(ctx)
	if err != nil {
		s.logger.Error("Failed to list patients", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch patients"})
		return
	}
	c.JSON(http.StatusOK, patients)
}
