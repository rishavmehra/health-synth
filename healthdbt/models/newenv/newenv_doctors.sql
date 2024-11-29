SELECT
    *,
    CAST(EXTRACT(YEAR FROM CURRENT_DATE) AS VARCHAR) || '-' || 
    CAST(EXTRACT(MONTH FROM CURRENT_DATE) AS VARCHAR) AS month_year
FROM {{ source('health_log', 'doctors') }}