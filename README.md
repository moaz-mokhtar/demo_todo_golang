# Demo for a todo app with GoLang, MySQL and Svelte.

This repo shows a sample for how to create a todo application with GoLang, MySQL and Svelte.


## Guidance:
- Install [golang](https://go.dev/doc/install)
- Install gcc build-essential (gcc)
- Run `mkdir demo_todo_golang` to create a directory for the project.
- Run `cd demo_todo_golang` to enter the project directory.
- Run `go mod init training/demo_todo` to create the dependencies file `go.mod`.
- Run `go get .` to download dependencies.
- Install docker, also can run script `sh scripts/get_docker.sh`.
- Run `docker comopse up` to start `mysql` container. Also you can run `docker compose down -v` to stop `mysql` container.
- Run `go run .` to run the program. 
- 