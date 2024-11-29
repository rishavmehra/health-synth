SELECT
    r.month_year,
    d.DoctorID,
    d.DoctorName,
    d.Specialty,
    COUNT(DISTINCT a.AppointmentID) AS total_appointments,
    COUNT(DISTINCT r.PatientID) AS total_patients,
    SUM(r.LastBilling) AS total_billing
FROM {{ ref("newenv_records") }} r
LEFT JOIN {{ ref("newenv_doctors") }} d ON r.DoctorID = d.DoctorID
LEFT JOIN {{ ref("newenv_appointments") }} a ON r.PatientID = a.PatientID AND r.DoctorID = a.DoctorID
GROUP BY r.month_year, d.DoctorID, d.DoctorName, d.Specialty