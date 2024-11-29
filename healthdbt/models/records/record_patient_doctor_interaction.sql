SELECT
    r.month_year,
    r.PatientID,
    r.DoctorID,
    COUNT(DISTINCT a.AppointmentID) AS total_appointments,
    MAX(r.LastReportDate) AS last_visit_date,
    MAX(r.HealthDisease) AS last_diagnosis
FROM {{ ref("newenv_records") }} r
LEFT JOIN {{ ref("newenv_appointments") }} a
    ON r.PatientID = a.PatientID AND r.DoctorID = a.DoctorID
GROUP BY r.month_year, r.PatientID, r.DoctorID