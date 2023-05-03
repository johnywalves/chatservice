-- name: FIndChatByID :one
SELECT * FROM chats WHERE id = ?;