CREATE TABLE IF NOT EXISTS videos(
    id SERIAL PRIMARY KEY,
    owner_id TEXT NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    transcoding_status TEXT NOT NULL,
    view_count INTEGER NOT NULL DEFAULT 0
)
