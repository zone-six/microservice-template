# Zone Six Microservice Template

## Usage
- Clone and remove git history
- Find and replace "microservice-template" with name of your service
- Delete whatever example code you don't need
- Init git

## Architecture

- All microservices should follow the architectural patterns laid out in the book [Righting Software](https://www.amazon.com/Righting-Software-Juval-L%C3%B6wy/dp/0136524036).

## Dependency Injection

- No DI frameworks are used. Instead, DI is achieved using a simple pattern where a component can call any component in they layer below it or any utility. Each layer's component's are organized into a container. Each component has access to the container of the layer below it. Each component has a public interface and private implementation and publishes a "New" method that is used to instantiate the component in a given layer's container.

## Logging

- Using [zerolog](https://github.com/rs/zerolog) should be acceptable. If apps are containerized, we should be able to stream StdOut and StdErr to some central logging service.
- Debugging from logs should be possible.
- Each log message should contain a CorrelationID if present in the context.

## Data

- Every service has own database and is responsible for its own migrations and managment

## Data Sharing/PubSub

- Data shared between services using pub-sub pattern (via NATS streaming)
  - Each service should keep a record of the last time it got a record from pub sub
  - When the service starts, it should ask NATS for all the records since the last record it recieved and write those to its database and/or process those records accordingly.
- Each services subjects should be namespaced (e.g. {zonesix}.{service-name}.{record-type}.{optional})
- Each service should publish a package with the subject names, exported as constants and an example of how to subscribe to these.
- Each service should publish a package with the types (i.e. structs) of messages that it posts along with what subject they're posted to.
- Messages are published via a/the PubSub utility (exported in a package)

## Clients

- Each service can clients (GraphQL, Rest, etc) that is best for the problem it is trying to solve.
- Clients could be as follows:
  - PubSub client that recieves messages from NATS streaming and processes these messages.
  - Jobs client that runs background processes.
    - Note: If performance and uptime are key, it may be best to have a separate main.go and build pipeline for the jobs portion of the service, so that running jobs won't impact the performance of a service
  - REST/GraphQL client for accepting incoming web traffic.
- In general, clients should be kept lightweight, calling into a manager to perform any sort of business logic.
- Public facing clients (REST, GraphQL, etc.) should be well documented, using some sort of lightweight client or a published document (e.g. Swagger, GraphQL playground)

## Authentication/Authorization

- Each service is responsible for managing what a user can access (i.e. access control).
- Identity of a user can be found through use of a JWT
- API Keys can be used if there is a need

## Libraries

- A service may also publish a package or library that will help other services to interact with it. This may be a client library for calling its REST or GraphQL endpoints or middleware that runs pre or post http call.
- Thought: For PubSub, instead of each service building its own client, a service could publish a package that has a method that takes a call back for each subject that it publishes a message to. This method could take care of the subscribing and unmarshalling of the encoded message, which could/should result in better type safety.

## Configuration
- Configuration should be handled using environment vars (see the internal/config folder)

## File Storage
- If file storage is needed, it is acceptable to read and write from disk. When running in containers, these paths should map to a mounted volume so that storage will be persisted.

## Middleware
- CorrelationID middleware
- UserContext middleware
- API key validation

## TODO: 
- If more than one server (e.g. jobs, REST, etc.) need to look at multiple docker files and builds.
