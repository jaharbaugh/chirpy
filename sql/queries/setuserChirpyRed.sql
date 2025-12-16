-- name: SetUserChirpyRed :exec
update users
set is_chirpy_red =$2
where id = $1;