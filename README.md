# Twooter

[![CI](https://github.com/tedsilb/twooter/actions/workflows/main.yml/badge.svg)](https://github.com/tedsilb/twooter/actions/workflows/main.yml)

[![CodeFactor](https://www.codefactor.io/repository/github/tedsilb/twooter/badge)](https://www.codefactor.io/repository/github/tedsilb/twooter)

Angular app talking to a gRPC backend. Prototype :)

## Building

### Frontend

`cd frontend/`

Frontend is built in Angular.

- To build the project:
  - `ng build`
- To run a devserver:
  - `ng serve --open`
- To run tests:
  - `ng test`
- To run e2e tests:
  - `ng e2e`

### Backend

`cd backend/`

Backend is written in Go.

- To build the project:
  - `go build ./...`
- To run tests:
  - `go test ./...`
