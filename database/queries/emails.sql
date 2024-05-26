-- name: GetEmail :one
SELECT
    *
FROM
    emails
WHERE
    id = $1
LIMIT
    1;

-- name: ListEmails :many
SELECT
    *
FROM
    emails
ORDER BY
    id;

-- name: CreateEmail :one
INSERT INTO
    emails (owner_email, email_html, url_hash)
VALUES
    ($1, $2, $3)
RETURNING
    *;

-- name: UpdateEmail :one
UPDATE emails
SET
    owner_email = $2,
    email_html = $3,
    url_hash = $4
WHERE
    id = $1
RETURNING
    *;

-- name: DeleteEmail :exec
DELETE FROM emails
WHERE
    id = $1;