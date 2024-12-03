CREATE TABLE IF NOT EXISTS library (
    song_id INT NOT NULL UNIQUE,
    group_name VARCHAR(50) NOT NULL,
    song_name VARCHAR(100) NOT NULL,
    song_text text,
    link text,
    release_date VARCHAR(30),
    UNIQUE (group_name, song_name)
);
