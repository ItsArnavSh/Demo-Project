package repository

import (
	"context"
	entity "demo/application/internal"
)

// db interface defines methods for interacting with patients and appointments in the database.
type Repository interface {
	//Login Methods
	RegisterUser(ctx context.Context, staff entity.Staff) error
	GetPassword(ctx context.Context, id string) (string, error)
	// Patient-related methods
	CreatePatient(ctx context.Context, patient entity.Patient) error
	GetPatientByID(ctx context.Context, id string) (entity.Patient, error)
	ListPatients(ctx context.Context) ([]entity.Patient, error)
	UpdatePatientStatus(ctx context.Context, id string, newStatus string) (entity.Patient, error)
	// Appointment-related methods
	CreateAppointment(ctx context.Context, patient_id string, appointment entity.Appointment) error
	GetAppointmentByID(ctx context.Context, appID string) (entity.Appointment, error)
	ListAppointmentsByPatientID(ctx context.Context, patientID string) ([]entity.Appointment, error)
	UpdateAppointment(ctx context.Context, appID string, appointment entity.Appointment) error
	DeleteAppointment(ctx context.Context, appID string) error
}
