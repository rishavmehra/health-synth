SELECT
    patientid,
    doctorid,
    lastreportdate,
    healthdisease,
    lastbilling,
    CAST(EXTRACT(YEAR FROM DATE(lastreportdate)) AS VARCHAR) || '-' || 
    CAST(EXTRACT(MONTH FROM DATE(lastreportdate)) AS VARCHAR) AS month_year
FROM {{ source('health_log', 'records') }}