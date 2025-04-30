package postgredb

import (
	"context"
	entity "demo/application/internal"
	"demo/application/internal/repository"
	"demo/db"
	"fmt"
)

type PostgreDB struct {
	queries *db.Queries
}

var _ repository.Repository = (*PostgreDB)(nil)

func NewRepo(queries *db.Queries) *PostgreDB {

	return &PostgreDB{queries: queries}
}

// Patient-related methods
func (p *PostgreDB) CreatePatient(ctx context.Context, patient entity.Patient) error {
	fmt.Println(patient)
	_, err := p.queries.CreatePatient(ctx, db.CreatePatientParams{
		PatientID: patient.Pid,
		Name:      patient.Name,
		Email:     patient.Email,
		Phno:      patient.PhNo,
		Dob:       patient.DOB,
	})
	return err
}

func (p *PostgreDB) UpdatePatientStatus(ctx context.Context, id string, newStatus string) (entity.Patient, error) {
	// Create the UpdatePatientParams struct with the patient ID and new status
	params := db.UpdatePatientParams{
		PatientID: id,
		Status:    newStatus,
	}

	// Call the UpdatePatient query with the params struct
	patient, err := p.queries.UpdatePatient(ctx, params)
	if err != nil {
		return entity.Patient{}, err
	}

	// Return the updated patient entity
	return entity.Patient{
		Pid:    patient.PatientID,
		Name:   patient.Name,
		Email:  patient.Email,
		PhNo:   patient.Phno,
		DOB:    patient.Dob,
		Status: patient.Status, // Updated status
	}, nil
}

func (p *PostgreDB) GetPatientByID(ctx context.Context, id string) (entity.Patient, error) {
	patient, err := p.queries.GetPatientByID(context.Background(), id)
	if err != nil {
		return entity.Patient{}, err
	}
	return entity.Patient{
		Pid:   patient.PatientID,
		Name:  patient.Name,
		Email: patient.Email,
		PhNo:  patient.Phno,
		DOB:   patient.Dob,
	}, nil
}

func (p *PostgreDB) ListPatients(ctx context.Context) ([]entity.Patient, error) {
	patients, err := p.queries.ListPatients(context.Background())
	if err != nil {
		return nil, err
	}

	var result []entity.Patient
	for _, patient := range patients {
		result = append(result, entity.Patient{
			Pid:   patient.PatientID,
			Name:  patient.Name,
			Email: patient.Email,
			PhNo:  patient.Phno,
			DOB:   patient.Dob,
		})
	}
	return result, nil
}

// Appointment-related methods
func (p *PostgreDB) CreateAppointment(ctx context.Context, patientID string, appointment entity.Appointment) error {
	_, err := p.queries.CreateAppointment(context.Background(), db.CreateAppointmentParams{
		AppID:     appointment.AppID,
		PatientID: patientID,
		Title:     appointment.Title,
		Diagnosis: appointment.Diagnosis,
		Treatment: appointment.Treatment,
		Status:    appointment.Status,
		Date:      appointment.Date,
	})
	return err
}

func (p *PostgreDB) GetAppointmentByID(ctx context.Context, appID string) (entity.Appointment, error) {
	appointment, err := p.queries.GetAppointmentByID(context.Background(), appID)
	if err != nil {
		return entity.Appointment{}, err
	}
	return entity.Appointment{
		AppID:     appointment.AppID,
		Title:     appointment.Title,
		Diagnosis: appointment.Diagnosis,
		Treatment: appointment.Treatment,
		Status:    appointment.Status,
		Date:      appointment.Date,
	}, nil
}

func (p *PostgreDB) ListAppointmentsByPatientID(ctx context.Context, patientID string) ([]entity.Appointment, error) {
	appointments, err := p.queries.ListAppointmentsByPatientID(context.Background(), patientID)
	if err != nil {
		return nil, err
	}

	var result []entity.Appointment
	for _, appointment := range appointments {
		result = append(result, entity.Appointment{
			AppID:     appointment.AppID,
			Title:     appointment.Title,
			Diagnosis: appointment.Diagnosis,
			Treatment: appointment.Treatment,
			Status:    appointment.Status,
			Date:      appointment.Date,
		})
	}

	return result, nil
}

func (p *PostgreDB) UpdateAppointment(ctx context.Context, appID string, appointment entity.Appointment) error {
	_, err := p.queries.UpdateAppointment(context.Background(), db.UpdateAppointmentParams{
		AppID:     appID,
		Title:     appointment.Title,
		Diagnosis: appointment.Diagnosis,
		Treatment: appointment.Treatment,
		Status:    appointment.Status,
		Date:      appointment.Date,
	})
	return err
}

func (p *PostgreDB) DeleteAppointment(ctx context.Context, appID string) error {
	err := p.queries.DeleteAppointment(context.Background(), appID)
	return err
}

func (p *PostgreDB) RegisterUser(ctx context.Context, staff entity.Staff) error {
	return p.queries.RegisterUser(ctx, db.RegisterUserParams{
		StaffID:  staff.StaffId,
		Name:     staff.Name,
		Password: staff.HashedPassword,
		Type:     staff.Type,
	})
}
func (p *PostgreDB) GetPassword(ctx context.Context, id string) (string, error) {
	return p.queries.GetUser(ctx, id)
}
