-- name: CreatePatient :one
INSERT INTO patient (patient_id, name, email, phno, dob, status)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING patient_id, name, email, phno, dob, status;

-- name: UpdatePatient :one
UPDATE patient
SET status = $2
WHERE patient_id = $1
RETURNING *;

-- name: GetPatientByID :one
SELECT patient_id, name, email, phno, dob, status
FROM patient
WHERE patient_id = $1;

-- name: ListPatients :many
SELECT patient_id, name, email, phno, dob, status
FROM patient
ORDER BY patient_id;

-- name: CreateAppointment :one
INSERT INTO appointment (app_id, patient_id, title, diagnosis, treatment, status, date)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING app_id, patient_id, title, diagnosis, treatment, status, date;

-- name: GetAppointmentByID :one
SELECT app_id, patient_id, title, diagnosis, treatment, status, date
FROM appointment
WHERE app_id = $1;

-- name: ListAppointmentsByPatientID :many
SELECT app_id, patient_id, title, diagnosis, treatment, status, date
FROM appointment
WHERE patient_id = $1
ORDER BY date;

-- name: UpdateAppointment :one
UPDATE appointment
SET title = $2, diagnosis = $3, treatment = $4, status = $5, date = $6
WHERE app_id = $1
RETURNING app_id, patient_id, title, diagnosis, treatment, status, date;

-- name: DeleteAppointment :exec
DELETE FROM appointment
WHERE app_id = $1;

-- name: RegisterUser :exec
INSERT INTO staff (staff_id,name,password,type)
VALUES ($1, $2, $3,$4);

-- name: GetUser :one
SELECT password
FROM staff
where staff_id = $1;
