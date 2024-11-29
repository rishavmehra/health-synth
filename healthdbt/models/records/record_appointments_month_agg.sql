SELECT
    month_year,
    COUNT(DISTINCT AppointmentID) AS total_appointments,
    COUNT(DISTINCT PatientID) AS total_patients,
    COUNT(DISTINCT DoctorID) AS total_doctors
FROM {{ ref("newenv_appointments") }}
GROUP BY month_year
