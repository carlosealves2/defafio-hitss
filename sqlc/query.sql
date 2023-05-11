-- name: CreateUser :one
INSERT INTO users (name, surname, contact, address, birth, cpf)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: UpdateUser :one
UPDATE users set
                 name = $2,
                 surname = $3,
                 contact = $4,
                 address = $5,
                 birth = $6,
                 cpf = $7
WHERE id = $1
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING id;