-- name: UpsertEvent :exec
INSERT INTO events (
    id,
    user_id,
    together,
    description,
    year,
    month,
    day,
    date,
    start_date,
    end_date,
    created_at,
    updated_at,
    important
)
VALUES (
    sqlc.arg(id),
    sqlc.arg(user_id),
    sqlc.arg(together),
    sqlc.arg(description),
    sqlc.arg(year),
    sqlc.arg(month),
    sqlc.arg(day),
    sqlc.arg(date),
    sqlc.arg(start_date),
    sqlc.arg(end_date),
    NOW(),
    NOW(),
    sqlc.arg(important)
)
ON CONFLICT (id) DO UPDATE
SET
    user_id     = EXCLUDED.user_id,
    together    = EXCLUDED.together,
    description = EXCLUDED.description,
    year        = EXCLUDED.year,
    month       = EXCLUDED.month,
    day         = EXCLUDED.day,
    date        = EXCLUDED.date,
    start_date  = EXCLUDED.start_date,
    end_date    = EXCLUDED.end_date,
    updated_at  = NOW(),
    important   = EXCLUDED.important;


-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = sqlc.arg(eventID);

-- name: FindEvent :one
SELECT *
FROM events
WHERE id = sqlc.arg(eventID);

-- name: FindDayOfEvent :one
SELECT *
FROM events
WHERE year  = sqlc.arg(year)
    AND month = sqlc.arg(month)
    AND day = sqlc.arg(day);

-- name: FindMonthEventIDs :many
SELECT id
FROM events
WHERE year  = sqlc.arg(year)
    AND month = sqlc.arg(month);