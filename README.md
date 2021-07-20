# Twooter

[![CI](https://github.com/tedsilb/twooter/actions/workflows/main.yml/badge.svg)](https://github.com/tedsilb/twooter/actions/workflows/main.yml)

[![CodeFactor](https://www.codefactor.io/repository/github/tedsilb/twooter/badge)](https://www.codefactor.io/repository/github/tedsilb/twooter)

Various Go projects I work on from time to time.

## Building

### Frontend

`cd frontend/`

Frontend is built using Angular.

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

Backend is written in Go and built using [Bazel](https://bazel.build).

- To build the project:
  - `bazel build ...`
- To run tests:
  - `bazel test ...`

Build rules are auto-generated using [Gazelle](https://github.com/bazelbuild/bazel-gazelle).

To update build rules, run `bazel run //:gazelle`
