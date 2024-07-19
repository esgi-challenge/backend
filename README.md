# Backend

Backend repository for the IW5 S2 challenge.

Swagger doc is reachable in `/swagger/index.html`

## Commands
All commands can be used with the Makefile and it's better and easier to use them in this way.

List of all usefull commands :

- `make run`: Run the program using *go run* command.
- `make build`: Build the program using *go build* command.
- `make test`: Launch all the tests using *go test* command.
- `make swag-init`: Create the docs for the swagger page, using the [swag](https://github.com/swaggo/swag) command.
- `make swag-fmt`: Format the comments used by the *swag-init* to generate the docs, using the [swag](https://github.com/swaggo/swag) command.
- `make swag`: Execute both *make swag-init* and *make swag-fmt* to save time.
- `make mock CRUD=<crudName>`: Generate the mocks for the Repository and UseCase of any CRUD passed as variable (change the `<crudName>` by your CRUD name in lowercase), the mocks are used by the tests. It use the [mockgen](https://github.com/uber-go/mock) command.
- `make crud`: Execute a custom python script that generate a full CRUD for any entity you want, including all basics CRUD actions, repository, usecase and tests directly included. It use the [jinja2](https://pypi.org/project/Jinja2/) pip package. It aim to work like the `php bin/console make:crud` symfony command.

## Launch

To launch the project in local, you first need to retrieve all the dependencies with this command :
```
go mod download
```

Then you need to setup the `.env` file by copying the `.env.example` file and rename it with the correct name `.env`.
You need to fill all the env variables of the `.env` file in order to make the application work, either way, the api won't start.
When all the variables are set, you can now build the app and launch it :
```
go build -o backend cmd/api/main.go
./backend
```

Or you can launch with docker compose :
```
docker compose up --build -d
```

## Tests

You can launch the test either with :
```
make test
```
or
```
go test -v --cover ./...
```

## Features :

Groupe :
- Antoine Lorin [AtoLrn](https://github.com/AtoLrn)
- Lucas Campistron [Redeltaz](https://github.com/Redeltaz)

Liste des features :
|   |   |
|---|---|
| Squelette, swagger, config, database, logger, error handler | Lucas Campistron |
| auth | Antoine Lorin |
| CRUD campus, classe, cours, document, informations, notes, filières, projets, emplois du temps, école, users | Antoine Lorin et Lucas Campistron |
| upload fichiers | Antoine Lorin |
| envois email | Lucas Campistron |
| invitation user well known | Antoine Lorin |
| chat temps réel | Lucas Campistron |
| intégration gmap suggestion | Lucas Campistron |
| tests | Lucas Campistron |
