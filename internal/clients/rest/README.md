# Rest Server

## Swagger CLI installation
- In your shell config file, add the following
```bash
alias swagger="docker run --rm -e GOPATH=/go -v ${HOME}:${HOME} -w $(pwd) -u $(id -u):$(id -g) stratoscale/swagger:v1.0.27"
```

## Documentation:
[stratoscale/swagger](https://github.com/Stratoscale/swagger)

## Generation:

- init

```bash
swagger init spec \
  --title "A Todo list application" \
  --description "From the todo list tutorial on goswagger.io" \
  --version 1.0.0 \
  --scheme http \
  --consumes application/io.goswagger.examples.todo-list.v1+json \
  --produces application/io.goswagger.examples.todo-list.v1+json
```

- server gen

```bash
swagger generate server -A todo-list -f ./swagger.yml -t ./internal/clients/rest
go mod tidy
```

## Validation:

```bash
swagger validate ./swagger.yml
```

## SwaggerUI:
- Copy dist from [swagger-api/swagger-ui](https://github.com/swagger-api/swagger-ui)
- Update the index.html SwaggerUIBundle -> url: "/swagger.json"
