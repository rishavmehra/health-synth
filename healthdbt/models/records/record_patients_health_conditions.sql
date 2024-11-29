SELECT
    month_year,
    SUM(CASE WHEN Hipertension THEN 1 ELSE 0 END) AS total_hypertension,
    SUM(CASE WHEN Diabetes THEN 1 ELSE 0 END) AS total_diabetes,
    SUM(CASE WHEN Alcoholism THEN 1 ELSE 0 END) AS total_alcoholism,
    SUM(CASE WHEN Handicap THEN 1 ELSE 0 END) AS total_handicap
FROM {{ ref("newenv_patients") }}
GROUP BY month_year