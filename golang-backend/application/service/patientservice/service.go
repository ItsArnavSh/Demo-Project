package patientservice

import (
	"context"
	entity "demo/application/internal"
	"demo/application/internal/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type PatientService struct {
	database repository.Repository
	logger   *zap.Logger
}

func NewPatientService(ctx context.Context, logger *zap.Logger, database repository.Repository) *PatientService {
	return &PatientService{database, logger}
}
func (p *PatientService) CreatePatient(ctx context.Context, patient entity.Patient) (string, error) {
	patient.Pid = uuid.NewString()
	err := p.database.CreatePatient(ctx, patient)
	if err != nil {
		p.logger.Error("Could Not Create Patient", zap.String("error", err.Error()))
		return "", err
	}
	return patient.Pid, nil
}

func (p *PatientService) UpdatePatientStatus(ctx context.Context, id string, newStatus string) (entity.Patient, error) {
	updatedPatient, err := p.database.UpdatePatientStatus(ctx, id, newStatus)
	if err != nil {
		p.logger.Error("Could Not Update Patient Status", zap.String("error", err.Error()))
		return entity.Patient{}, err
	}
	return updatedPatient, nil
}

func (p *PatientService) GetPatient(ctx context.Context, id string) (entity.Patient, error) {
	return p.database.GetPatientByID(ctx, id)
}

func (p *PatientService) ListPatients(ctx context.Context) ([]entity.Patient, error) {
	patients, err := p.database.ListPatients(ctx)
	if err != nil {
		p.logger.Error("Could Not List Patients", zap.String("error", err.Error()))
		return nil, err
	}
	return patients, nil
}

func (p *PatientService) CreateAppointment(ctx context.Context, patientID string, appointment entity.Appointment) (string, error) {
	appointment.AppID = uuid.NewString()
	err := p.database.CreateAppointment(ctx, patientID, appointment)
	if err != nil {
		p.logger.Error("Could Not Create Appointment", zap.String("error", err.Error()))
		return "", err
	}
	return appointment.AppID, nil
}

func (p *PatientService) GetAppointment(ctx context.Context, appID string) (entity.Appointment, error) {
	appointment, err := p.database.GetAppointmentByID(ctx, appID)
	if err != nil {
		p.logger.Error("Could Not Get Appointment", zap.String("error", err.Error()))
		return entity.Appointment{}, err
	}
	return appointment, nil
}

func (p *PatientService) ListAppointmentsByPatientID(ctx context.Context, patientID string) ([]entity.Appointment, error) {
	appointments, err := p.database.ListAppointmentsByPatientID(ctx, patientID)
	if err != nil {
		p.logger.Error("Could Not List Appointments By Patient ID", zap.String("error", err.Error()))
		return nil, err
	}
	return appointments, nil
}

func (p *PatientService) UpdateAppointment(ctx context.Context, appID string, appointment entity.Appointment) error {
	err := p.database.UpdateAppointment(ctx, appID, appointment)
	if err != nil {
		p.logger.Error("Could Not Update Appointment", zap.String("error", err.Error()))
		return err
	}
	return nil
}

func (p *PatientService) DeleteAppointment(ctx context.Context, appID string) error {
	err := p.database.DeleteAppointment(ctx, appID)
	if err != nil {
		p.logger.Error("Could Not Delete Appointment", zap.String("error", err.Error()))
		return err
	}
	return nil
}
