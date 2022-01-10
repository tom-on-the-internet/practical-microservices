CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO videos (owner_id, name, description, transcoding_status, view_count) VALUES
    (uuid_generate_v4(), 'Video A', 'A video', 'transcoded', 21),
    (uuid_generate_v4(), 'Video B', 'Another video', 'transcoded', 38),
    (uuid_generate_v4(), 'Video C', 'A bad video', 'transcoded', 0),
    (uuid_generate_v4(), 'Video C', 'A risqué video', 'transcoded', 98);

