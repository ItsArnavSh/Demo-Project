# 🏥 Hospital Management API Documentation

This API enables role-based access for `receptionist` and `doctor` functionalities using middleware-based authentication and authorization. All endpoints require a valid token in the `Authorization` header after signup/login.



---

## 🚀 Boot Instructions

Follow these steps to launch the full stack (Docker + Go backend + Database):

### 1. Start Docker Containers

Ensure Docker is installed and running. From the project root directory:

```bash
docker compose up
```

> This starts the PostgreSQL database using the configuration in the `database/` folder.

---

### 2. Set Up and Run the Backend

Navigate to the Go backend directory:

```bash
cd golang-backend
```

Run the setup script (requires `make`):

```bash
make setup
```

> This step:
>
> - Runs SQL migrations
>
> - Generates Go code using `sqlc` based on the SQL queries
>

Now, start the Go server:

```bash
make run
```

---

## 🔐 Authentication

### `POST /signup`

Registers a new user with a specific role (`receptionist` or `doctor`) and returns a JWT token.

#### Request Body

```json
{
  "name": "abcd",
  "password": "12312",
  "type": "receptionist"
}
```

#### Response

```json
{
  "token": "eyJ..."
}
```

---

## 🧾 Receptionist Endpoints

### `POST /receptionist/addPatient`

Adds a new patient to the system.

#### Request Body

```json
{
  "name": "abcd",
  "email": "cool@gmail.com",
  "phno": "9817177189",
  "DOB": "01-02-2000"
}
```

---

### `GET /receptionist/listPatients`

Retrieves a list of all registered patients.

#### Response

```json
[
  {
    "Pid": "2048a3ec-3425-49e7-9a29-9642266b699a",
    "Name": "abcd",
    "Email": "hello",
    "PhNo": "9817177889",
    "DOB": "19July",
    "Status": "",
    "Appointments": null
  }
]
```

---

### `GET /receptionist/getPatient/{patient_id}`

Retrieves detailed information of a specific patient using their ID.

#### Response

```json
{
  "Pid": "2048a3ec-3425-49e7-9a29-9642266b699a",
  "Name": "abcd",
  "Email": "hello",
  "PhNo": "9817177889",
  "DOB": "19July",
  "Status": "",
  "Appointments": null
}
```

---

## 🩺 Doctor Endpoints

### `POST /doctor/createAppointment/{patient_id}`

Creates a new appointment for a specific patient.

#### Request Body

```json
{
  "Title": "Regular Checkup",
  "Diagnosis": "Hypertension",
  "Treatment": "Lifestyle modification and medication"
}
```

#### Response

```json
{
  "appointment_id": "a2a4d430-69cb-4a68-8e73-4ea28e80a734"
}
```

---

### `GET /doctor/listAppointmentsByPatient/{patient_id}`

Fetches all appointments associated with a specific patient.

#### Response

```json
[
  {
    "AppID": "a2a4d430-69cb-4a68-8e73-4ea28e80a734",
    "Title": "Regular Checkup",
    "Diagnosis": "Hypertension",
    "Treatment": "Lifestyle modification and medication",
    "Status": "",
    "Date": "0001-01-01T00:00:00Z"
  }
]
```

---

### `PUT /doctor/updatePatientStatus/{patient_id}`

Updates the medical status of a patient.

#### Request Body

```json
{
  "new_status": "He is alright"
}
```

#### Response

```json
{
  "Pid": "2048a3ec-3425-49e7-9a29-9642266b699a",
  "Name": "abcd",
  "Email": "hello",
  "PhNo": "9817177889",
  "DOB": "19July",
  "Status": "He is alright",
  "Appointments": null
}
```

---

## 🔐 Authorization

All endpoints (except `/signup`) require a JWT token in the `Authorization` header:

```
Authorization: Bearer <token>
```

Middleware ensures:

- `receptionist` routes are only accessible to users with the role `receptionist`.

- `doctor` routes are only accessible to users with the role `doctor`.


---
# Demo-Project
