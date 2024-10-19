-- name: GetCompanyByID :one
SELECT *
FROM company
WHERE id = $1 LIMIT 1;

-- name: CreateCompany :one
INSERT INTO company (name, description, employee_count, registered, type)
VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: UpdateCompany :exec
UPDATE company
SET name           = $2,
    description    = $3,
    employee_count = $4,
    registered     = $5,
    type           = $6
WHERE id = $1;

-- name: DeleteCompany :exec
DELETE
FROM company
WHERE id = $1;