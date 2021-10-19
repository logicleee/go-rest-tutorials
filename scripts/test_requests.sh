#!/usr/bin/env bash
set -e

# not working
#go run . &
#go_pid=$!

printf "\n\n  => GET /albums <=\n"
echo $(curl http://localhost:8080/albums)

printf "\n\n  => POST /albums <=\n"
echo $(curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "6","title": "Duke Ellington & John Coltrane","artist": "Duke Ellington & John Coltrane","price": 19.62}')

printf "\n\n  => GET /albums <=\n"
echo $(curl http://localhost:8080/albums)

printf "\n\n  => GET /albums/2 <=\n"
echo $(curl http://localhost:8080/albums/2)

#kill $go_pid