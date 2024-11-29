
SELECT
    *,
    CAST(EXTRACT(YEAR FROM DATE(visitdate)) AS VARCHAR) || '-' || 
    CAST(EXTRACT(MONTH FROM DATE(visitdate)) AS VARCHAR) AS month_year
FROM {{ source('health_log', 'patients') }}
