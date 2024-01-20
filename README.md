# Braza

**Braza** is a fork of [Horusec](https://github.com/ZupIT/horusec) Open Source cli SAST.

The original source code is copyrighted and owned by **ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA** currently released under Apache 2.0. 

This fork aims on updating and adding new functionalities not currently present in the original.

## About

The main objective of the project is, not only, keeping this great SAST alive with newer updates and maintenance but also, to implement some DevSecOps focused features, to allow easy implementation of this software in:
- CI/CD pipelines.
- Kubernetes Clusters.
- Integration with **Observability** tools.

Overall making it easy for **Application Security** and **DevSecOps** teams to work with it while maintaining its Command Line Interface for fast mitigation by the Development teams.

The project is still in early development and there is no ROADMAP. 

The first steps will be as described in the Todos section below.

## Running

You can simple run ```make build-dev``` and then ```./braza-dev version``` (PS: Docker may require ```sudo``` for running in your system)

## Todos

### Short term
- ✅ **Make it a monorepo (for now)**
    - ~~Join horusec-engine and horusec-devkit into horusec cli project, for easier development (I'm a lonely dev).~~
- **Update and Upgrade every package used**
    - ✅ Go mod provides easy management of dependencies, I will be focusing on updating all the dependencies and making it work with the latest version of everything.
    - Remove deprecated and archived dependencies. (This one may take some time)
    - Check all the security tools used for licensing, depecration, security related issues.
- **Clear the project**
    - Remove legacy pkgs and files that were required by the another horusec projects (Web App, Operator, etc...), and are not required now.
- **Fix all the tests**
    - After all the cleaning, tests should be working again
- **Fix Pipelines**
    - I've removed all the workflows available for github but I still want the release/quality/security ones in the future. For now, manually building the project will be enough.
- **Improve Output**
    - Add better output for humans to easily identify and correct the vulnerabilities when running from cli.
    - Make it easier to send the report file to some external API.
    - Improve integration with Github, making it easier to configure Issue generation and handling of vulnerabilities. Further increase these integrations with another platforms like Gitlab and Bitbucket.

### Long term
- **Overall Performance**:
    - I want to remove the docker dependency, allowing to use alternatives like [containerd](https://containerd.io/) or [podman](https://podman.io).
    - Study new forms of improving the analysis speed and computer resource usage.
    - Find solutions for the **False positive** problem, very common in mostly SAST tools.
    - Study new ways of handling the Dockerfiles requirements, maybe find some form of templating useful for adding extra tooling to the already huge inventory of possible scans.
    - Allow users to better configure the underlying tools, allowing them to use their own pro versions of available tools.
    - Improve integration with open source Security tools like OWASP [DefectDojo](https://www.defectdojo.org/).


## Contributing

Not accepting any contribution for now.

