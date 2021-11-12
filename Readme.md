# CatMQ

CatMQ is a queue server that allows offline data processing.
The main purpose of this server is to analyze text snippets and reject those that have offensive terms.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purpose.

There are two ways to get started:
<details>
<summary>Get started with local Golang;</summary>

### Prerequisites

- [Golang](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. You need the version 1.17.

### Installation

1. Clone this repository;

2. Run the following commands:

```Makefile
cd new-server
make setup
make run
```

### Running tests

1. Run the server
```Makefile
make run
```

2. In another terminal, run the tests:

```Makefile
make test
```

</details>

<details>
<summary>Get started with Docker;</summary>
### Prerequisites

- [Docker](https://www.docker.com/) - is an open platform for developing, shipping, and running applications. Docker enables you to separate your applications from your infrastructure so you can deliver software quickly.

### Installation

1. Clone this repository;

2. Run the following commands:

```Makefile
cd new-server
make docker-build
make docker-run
```

4. For kill container's Docker, run the following command:

```Makefile
make docker-kill
```

</details>

## How it works?

### Server

The server receives TCP connections on port `23023` and follows a message sending and replying protocol.
Client connections remain open until the client process is closed.
Messages sent to and from the server are messages ending in `'\n'` with the format specified in Protocol.

### Protocol

#### Publish

To send a message to the queue (publish), the client must send to the server the following message:

`{"operation":"PUT","data":"<some data>"}\n`

The server responds with:

`{"operation":"OK","data":""}\n`

In case of error, the server should respond:

`{"operation":"ERR","data":"Operation failed!"}\n`

#### Get

To receive a message from the queue, the client must send the following message to the server:

`{"operation":"GET","data":""}\n`

If there is any message in the queue, the server responds with:

`{"operation":"MSG","data":"<some data>"}\n`

If the queue is empty, the server responds with:

`{"operation":"EMP","data":""}\n`

In case of error, the server should respond:

`{"operation":"ERR","data":"Operation failed!"}\n`

### Client

The client library must have a Client class that receives the host and server port and have the following methods:

#### Connect

Connects to the server and leaves a socket open with the server.

#### Disconnect

Closes the socket opened with the server. Does nothing if the connection already
it's closed.

#### Publish

Sends a message to the queue. This method takes the message to be sent as an argument.

#### Get

Read a message from the queue. If the server responds with EMP, it returns null (no error).

## Versions

We use [Semantic version](http://semver.org) for versioning. For versions available, see [changelog](Changelog.md).

## Contribute Us

Contributions are what make the open source community such an amazing place to learn, inspire, and create.
Any contributions you make are greatly appreciated. See [contribute policy](Contribute.md).

## License

This project is licensed under the [MIT License](LICENSE).
