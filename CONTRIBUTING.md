# Contributing


## Project Setup


### Prerequisites
* [Install bun](https://bun.sh/docs/installation)
* Install Docker Desktop. [Windows Install](https://docs.docker.com/desktop/install/windows-install/), [Mac Install](https://docs.docker.com/desktop/install/mac-install/)

Install docker and run the following docker compose script
```bash
docker compose -f docker/docker-compose.yml -p autoscaler up 
```

Run the backend go project
```bash
cd services/api
go run main.go
```

Run the frontend/ssr nextjs project
```bash
cd website
bun install
bun run dev
```

You will need to sign up for an account once the project is running.


### Naming Conventions

* folders and files use lower case one word if possible, if not, use snake_case.
* All url routes should be kebab-case.
* function names must be PascalCase when exported, and camelCase when internal. For more information read [Effective Go](https://go.dev/doc/effective_go)

### Validation

All code must pass the following validation checks:
* Typescript via the "typ-check" script in the package.json
* Linting via the "lint" script in the package.json