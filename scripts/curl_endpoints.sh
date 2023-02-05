#!/usr/bin/env bash

# set -xeu

PORT=3000

echo "\n\n==========================="
echo "Healthcheck, endpoint 'GET /'"
curl http://localhost:$PORT/


echo "\n\n==========================="
echo "Get all todos, endpoint 'GET /todos'"
curl http://localhost:$PORT/todos


echo "\n\n==========================="
echo "Get a todo by id, endpoint 'GET /todo/:id'"
curl http://localhost:$PORT/todo/2


echo "\n\n==========================="
echo "Create new todo, endpoint 'POST /todo'"
curl http://localhost:$PORT/todo \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": 4,"description": "Get a research for best practices with Go","priority": 4}'


echo "\n\n==========================="
echo "Delete a todo, endpoint 'DELETE /todo/:id'"
curl http://localhost:$PORT/todo/1 \
    --include \
    --request "DELETE"


echo "\n\n==========================="
echo "Update a todo, endpoint 'GET /todo/:id'"
curl http://localhost:$PORT/todo/2 \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"id": 2,"description": "Participate in Open Source with Go or a volunteer position","priority": 4}'








