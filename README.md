# Braza

**Braza** is a fork of [Horusec](https://github.com/ZupIT/horusec) Open Source cli SAST.

The original source code is copyrighted and owned by **ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA** currently released under Apache 2.0. 

This fork aims on updating and adding new functionalities not currently present in the original.

## About

The main objective of the project is, not only, keeping this brilliant SAST alive with newer updates and maintenance but also, to implement some DevSecOps focused features, to allow easy implementation of this software in CI/CD pipelines, Kubernetes Clusters and Observability tools.

The project is still in early development and there is no ROADMAP. 

The first steps will be as described in the Todos section below.

## Todos

- **Make it a monorepo (for now)**.
    - Join all the necessary tools for making the CLI available at a single repo for easier development early. This will probably be undone in the future for better isolation of behaviors but for now, it will help.
- **Clear the project**.
    - Remove legacy pkgs and files that were required by the another horusec projects (Web App, Operator, etc...), and are not required for running the tool.
- **Update and Upgrade every package used**.
    - Go tooling provides easy management of dependencies, I will be focusing on updating all the dependencies and making it work with newer version of everything.
    - Some dependencies will require extra development from our side to fix/implement newer and not existing features.
    - I want to remove the docker dependency, allowing to use alternatives like [containerd](https://containerd.io/) or [podman](https://podman.io).
- **Fix Pipelines**.
    - I've removed all the workflows available for github but I still want the release/quality/security ones in the future. For now, manually building the project will be enough.

## Contributing

Not accepting any contribution for now.

## Running

You can simple run ```make build-dev``` and then ```./braza-dev version```