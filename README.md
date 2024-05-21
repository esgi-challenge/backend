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
