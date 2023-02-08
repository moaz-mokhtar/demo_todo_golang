# Demo for a todo app with GoLang, MySQL and Svelte.

This repo shows a sample for how to create a todo application with GoLang, MySQL and Svelte.


## Guidance:
### Go language
https://go.dev
- Install [golang](https://go.dev/doc/install)
- Install gcc build-essential (gcc)
- Run `mkdir demo_todo_golang` to create a directory for the project.
- Run `cd demo_todo_golang` to enter the project directory.
- Run `go mod init training/demo_todo` to create the dependencies file `go.mod`.
- Run `go get .` to download dependencies.
- Run `go run .` to run the program. 

### `mysql` and `docker`
- Install docker, also can run script `sh scripts/get_docker.sh`.
- Run `docker compose up` to start `mysql` container. Also you can run `docker compose down -v` to stop `mysql` container.
Note: when facing issues with `mysql` container just follow:
```shell
$ docker compose down -v
...
$ docker compose up
...
```

### Svelte
https://github.com/sveltejs/svelte

- Install `NodeJs` and `npm` [Link](https://nodejs.org/en/download/).
Note: You can check if `NodeJs` and `npm` are installed using below commands:
```shell
$ node -v
v18.13.0
$ npm -v
8.19.3
```
- Install `SvelteKit` which supports local development with `Svelte`. Link: https://kit.svelte.dev/
- You can a good guidance https://svelte.dev/blog/svelte-for-new-developers
- To creat a Svelte project based on a template after installing SvelteKit:
```shell
npm create vite@latest app-name -- --template svelte
cd app-name
npm install 
npm run dev
```

