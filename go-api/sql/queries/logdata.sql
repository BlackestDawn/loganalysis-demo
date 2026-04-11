-- name: StoreLogData :exec
INSERT INTO logdata (timestamp, level, service, message) VALUES ($1, $2, $3, $4);

-- name: GetLogData :many
SELECT timestamp, level, service, message FROM logdata;
