SELECT
    *,
    CAST(EXTRACT(YEAR FROM appointmentday) AS VARCHAR) || '-' || 
    CAST(EXTRACT(MONTH FROM appointmentday) AS VARCHAR) AS month_year
FROM {{ source('health_log', 'appointments') }}