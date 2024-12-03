INSERT INTO library (song_id, group_name, song_name, song_text, link, release_date)
SELECT COALESCE(MAX(song_id), 0) + 1, $1, $2, $3, $4, $5
FROM library;