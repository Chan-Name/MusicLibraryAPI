UPDATE library
SET group_name = $2, song_name = $3, song_text = $4, link = $5, release_date = $6
WHERE song_id = $1;