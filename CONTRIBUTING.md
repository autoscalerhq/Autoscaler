# Contributing


## Project Setup


### Prerequisites
* Install pnpm globally `npm install -g pnpm`
* Install docker. If using windows, [use docker desktop](https://docs.docker.com/desktop/install/windows-install/). 

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
pnpm install
pnpm run dev
```

You will need to sign up for an account once the project is running.


### Naming Conventions


* folders and files use lower case one word if possible, if not, use snake_case.
* All url routes should be kebab-case.
* function names must be PascalCase when exported, and camelCase when internal.