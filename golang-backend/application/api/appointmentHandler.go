package api

import (
	entity "demo/application/internal"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Server) createAppointment(c *gin.Context) {
	patientID := c.Param("patient_id")
	var appointment entity.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	ctx := c.Request.Context()
	id, err := s.patientService.CreateAppointment(ctx, patientID, appointment)
	if err != nil {
		s.logger.Error("Failed to create appointment", zap.String("PatientID", patientID), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create appointment"})
		return
	}
	s.logger.Info("Appointment Created", zap.String("AppointmentID", id))
	c.JSON(http.StatusOK, gin.H{"appointment_id": id})
}

func (s *Server) getAppointment(c *gin.Context) {
	appID := c.Param("appointment_id")

	ctx := c.Request.Context()
	appointment, err := s.patientService.GetAppointment(ctx, appID)
	if err != nil {
		s.logger.Error("Failed to get appointment", zap.String("AppointmentID", appID), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch appointment"})
		return
	}
	c.JSON(http.StatusOK, appointment)
}
func (s *Server) listAppointmentsByPatient(c *gin.Context) {
	patientID := c.Param("patient_id")

	ctx := c.Request.Context()
	appointments, err := s.patientService.ListAppointmentsByPatientID(ctx, patientID)
	if err != nil {
		s.logger.Error("Failed to list appointments", zap.String("PatientID", patientID), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch appointments"})
		return
	}
	c.JSON(http.StatusOK, appointments)
}
func (s *Server) updateAppointment(c *gin.Context) {
	appID := c.Param("appointment_id")
	var appointment entity.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	ctx := c.Request.Context()
	err := s.patientService.UpdateAppointment(ctx, appID, appointment)
	if err != nil {
		s.logger.Error("Failed to update appointment", zap.String("AppointmentID", appID), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update appointment"})
		return
	}
	s.logger.Info("Appointment Updated", zap.String("AppointmentID", appID))
	c.JSON(http.StatusOK, gin.H{"message": "Appointment Updated"})
}

func (s *Server) deleteAppointment(c *gin.Context) {
	appID := c.Param("appointment_id")

	ctx := c.Request.Context()
	err := s.patientService.DeleteAppointment(ctx, appID)
	if err != nil {
		s.logger.Error("Failed to delete appointment", zap.String("AppointmentID", appID), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete appointment"})
		return
	}
	s.logger.Info("Appointment Deleted", zap.String("AppointmentID", appID))
	c.JSON(http.StatusOK, gin.H{"message": "appointment deleted"})
}
