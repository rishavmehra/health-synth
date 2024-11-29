SELECT 
    doctorid,
    doctorname,
    specialty,
    SUM(total_billing) AS total_billing,
    SUM(total_patients) AS total_patients,
    CASE 
        WHEN SUM(total_patients) > 0 THEN SUM(total_billing) / SUM(total_patients)
        ELSE 0 
    END AS avg_billing_per_patient
FROM 
    record_doctor_performance
GROUP BY 
    doctorid, doctorname, specialty
ORDER BY 
    avg_billing_per_patient DESC;