# bs-online-judge
A modern online judge.

## Building, running and testing locally
TODO.

## Running in large scale
TODO.

## System Architecture
This section describes the components of the system.

### Backend Applications
This section describes the backend applications of the system.

#### The API Servers
An API server is a stateless server providing an API to define and control the entities, or *resources*, of the system.

Clients who wish to create and manipulate the system resources are able to perform these operations by making requests to the API server. Each client described in this architecture has a dedicated API subset.

#### The Judges
Judges are workers consuming a queue of *submission executions*.

For each job popped from the queue, a judge is fulfilling a request of executing a *problem submission*, checking whether the outputs of this *submission execution* are correct and storing this check in the API. When talking to the API servers, a judge uses the *judge API* subset. For the execution part, the judge performs a request to an execution server.

#### The Execution Servers
An execution server is responsible for the critical mission of the system: executing untrusted code. For this, a particular execution server should be as much isolated as possible from everything else. Execution servers should not make requests to the API server directly, instead every required resource from the API should be passed in the execution request parameters by the requesting judge.

For each execution request, an execution server reads the required inputs from the file storage, compiles and executes the untrusted code, writes the associated outputs in the file storage and responds completion to the requesting judge.

#### The Timeout Engine
The timeout engine is a worker responsible for watching the problem submissions still needing automatic judgement and creating associated submission executions in the *timeout API* subset.

### Data Stores
This section describes the components of the system responsible for storing all sorts of data.

#### The Relational DBMS
A relational database should store the data behind the resources of the system exposed by the API.

#### The Cache
A cache should store the short-living auxiliar data shared by the API servers, like session tokens.

#### The Queue Message Broker
A message broker should act as a queue receiving submission execution jobs from the API servers and delivering those jobs atomically in FIFO order to the judges.

#### The File Storage
A file storage should store and serve files related to problems and submissions, like source codes, inputs and outputs.

### User Interfaces
This section describes the user interfaces of the system.

#### The Website
A website for the user to interact with the system. The website consumes the *web API* subset.

#### The Command-line Interface
A command-line interface for the user to interact with the system. The CLI consumes the *core API* subset.

## Architechture Implementation
This section describes our implementation of the architechture described above.

### Technologies
* The networking protocol chosen for communication between all clients and servers is HTTP.
* The programming language chosen for the backend applications and the CLI is Golang.
* The relational DBMS chosen is PostgreSQL.
* Redis will be used as cache and message broker.
* AWS S3 can be used as production file storage. For small scale environments we provide an implementation of an S3-API-compliant HTTP server in Golang.
* The website should be built with the fanciest technology of the moment.

### API Definitions
TODO.

### Configurations
This section describes all kinds of configurations for each system component.

TODO: Describe build time configuration files, database and cache stored settings, environment variables...

### Roadmap
TODO.

### Benchmarks
TODO.
