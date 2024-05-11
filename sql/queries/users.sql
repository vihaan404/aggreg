
-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name , API_KEY)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex')
       )
RETURNING *;


-- name: GetUserApiKey :one
SELECT * from users
where API_KEY = $1 ;
