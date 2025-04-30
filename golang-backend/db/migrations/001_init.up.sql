-- Patient Table
CREATE TABLE IF NOT EXISTS patient (
    patient_id VARCHAR(128) PRIMARY KEY, -- Corresponds to Pid
    name VARCHAR(128) NOT NULL, -- Corresponds to Name
    email VARCHAR(128) NOT NULL UNIQUE, -- Corresponds to Email
    phno VARCHAR(128) NOT NULL UNIQUE, -- Corresponds to PhNo
    dob VARCHAR(128) NOT NULL, -- Corresponds to DOB (Date type for clarity)
    status VARCHAR(128) NOT NULL DEFAULT "HEALTHY"
);

-- Appointment Table
CREATE TABLE IF NOT EXISTS appointment (
    app_id VARCHAR(128) PRIMARY KEY, -- Corresponds to AppID
    patient_id VARCHAR(128) NOT NULL, -- Foreign Key to Patient Table
    title VARCHAR(128) NOT NULL,
    diagnosis TEXT NOT NULL,
    treatment TEXT NOT NULL,
    status VARCHAR(64) NOT NULL,
    date TIMESTAMP NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES patient (patient_id) -- Foreign Key Constraint
);

CREATE TABLE IF NOT EXISTS staff (
    staff_id TEXT PRIMARY KEY,
    type VARCHAR(32) NOT NULL,
    name TEXT NOT NULL,
    password TEXT NOT NULL
)
