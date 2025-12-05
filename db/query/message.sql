-- name: CreateMessage :one
INSERT INTO message (thread, sender, content)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetMessageByID :one
SELECT * FROM message
WHERE id = $1;

-- name: GetAllMessages :many
SELECT * FROM message
ORDER BY created_at DESC;

-- name: GetMessagesByThread :many
SELECT * FROM message
WHERE thread = $1
ORDER BY created_at DESC;

-- name: DeleteMessageByID :exec
DELETE FROM message
WHERE id = $1;  

-- name: UpdateMessageContent :exec
UPDATE message
SET content = $2
WHERE id = $1;

-- name: CountMessagesInThread :one
SELECT COUNT(*) AS count
FROM message
WHERE thread = $1;