SELECT song_text
FROM library
WHERE group_name = $1 AND song_name = $2
ORDER BY song_id;