SELECT *
FROM AppCropRotationView
WHERE OrganizationId = 1
  AND Date >= '2022-01-01'
ORDER BY Date DESC