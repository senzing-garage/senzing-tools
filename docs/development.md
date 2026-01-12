# senzing-tools development

The following instructions are useful during development.

**Note:** This has been tested on Linux and Darwin/macOS.
It has not been tested on Windows.

## Prerequisites for development

:thinking: The following tasks need to be complete before proceeding.
These are "one-time tasks" which may already have been completed.

1. The following software programs need to be installed:
   1. [git]
   1. [make]
   1. [docker]
   1. [go]

## Install Senzing C library

Since the Senzing library is a prerequisite, it must be installed first.

1. Verify Senzing C shared objects, configuration, and SDK header files are installed.
   1. `/opt/senzing/er/lib`
   1. `/opt/senzing/er/sdk/c`
   1. `/etc/opt/senzing`

1. If not installed, see [How to Install Senzing for Go Development].

## Install Git repository

1. Identify git repository.

   ```console
   export GIT_ACCOUNT=senzing-garage
   export GIT_REPOSITORY=senzing-tools
   export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
   export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

   ```

1. Using the environment variables values just set, follow
   steps in [clone-repository] to install the Git repository.

## Dependencies

1. A one-time command to install dependencies needed for `make` targets.
   Example:

   ```console
   cd ${GIT_REPOSITORY_DIR}
   make dependencies-for-development

   ```

1. Install dependencies needed for [Go] code.
   Example:

   ```console
   cd ${GIT_REPOSITORY_DIR}
   make dependencies

   ```

## Lint

1. Run linting.
   Example:

   ```console
   cd ${GIT_REPOSITORY_DIR}
   make lint

   ```

## Build

1. Build the binaries.
   Example:

   ```console
   cd ${GIT_REPOSITORY_DIR}
   make clean build

   ```

1. The binaries will be found in the `${GIT_REPOSITORY_DIR}/target` directory.
   Example:

   ```console
   tree ${GIT_REPOSITORY_DIR}/target

   ```

## Run

1. Run program.
   Examples:
   1. Linux

      ```console
       export LD_LIBRARY_PATH=/opt/senzing/er/lib/
       ${GIT_REPOSITORY_DIR}/target/linux-amd64/senzing-tools

      ```

   1. macOS

      ```console
      export DYLD_LIBRARY_PATH=/opt/senzing/er/lib/:/opt/senzing/er/lib/macos
      ${GIT_REPOSITORY_DIR}/target/darwin-amd64/senzing-tools

      ```

   1. Windows

      ```console
      ${GIT_REPOSITORY_DIR}/target/windows-amd64/senzing-tools

      ```

1. Clean up.
   Example:

   ```console
   cd ${GIT_REPOSITORY_DIR}
   make clean

   ```

## Test

1. Run tests.
   Example:

   ```console
   cd ${GIT_REPOSITORY_DIR}
   make clean setup test

   ```

## Coverage

Create a code coverage map.

1. Run Go tests.
   Example:

   ```console
   cd ${GIT_REPOSITORY_DIR}
   make clean setup coverage

   ```

   A web-browser will show the results of the coverage.
   The goal is to have over 80% coverage.
   Anything less needs to be reflected in [testcoverage.yaml].

## Documentation

1. View documentation.
   Example:

   ```console
   cd ${GIT_REPOSITORY_DIR}
   make clean documentation

   ```

1. If a web page doesn't appear, visit [localhost:6060].
1. Senzing documentation will be in the "Third party" section.
   `github.com` > `senzing-garage` > `senzing-tools`

1. When a versioned release is published with a `v0.0.0` format tag,
   the reference can be found by clicking on the following badge at the top of the README.md page.
   Example:

   [![Go Reference Badge]][Go Reference]

1. To stop the `godoc` server, run

   ```console
   cd ${GIT_REPOSITORY_DIR}
   make clean

   ```

## Docker

1. Use make target to run a docker images that builds RPM and DEB files.
   Example:

   ```console
   cd ${GIT_REPOSITORY_DIR}
   make docker-build

   ```

1. Run docker container.
   Example:

   ```console
   docker run --rm senzing/template-go

   ```

1. **Optional:** Test using `docker-compose`.
   Example:

   ```console
   cd ${GIT_REPOSITORY_DIR}
   make clean docker-test

   ```

   To bring the `docker-compose` formation, run

   ```console
   cd ${GIT_REPOSITORY_DIR}
   make clean

   ```

## Package

### Package RPM and DEB files

1. Use make target to run a docker images that builds RPM and DEB files.
   Example:

   ```console
   cd ${GIT_REPOSITORY_DIR}
   make package

   ```

1. The results will be in the `${GIT_REPOSITORY_DIR}/target` directory.
   Example:

   ```console
   tree ${GIT_REPOSITORY_DIR}/target

   ```

### Test DEB package on Ubuntu

1. Determine if `senzing-tools` is installed.
   Example:

   ```console
   apt list --installed | grep senzing-tools

   ```

1. :pencil2: Install `senzing-tools`.
   Example:

   ```console
   cd ${GIT_REPOSITORY_DIR}/target
   sudo apt install ./senzing-tools-0.0.0.deb

   ```

1. :pencil2: Identify database.
   Example:

   ```console
   export SENZING_TOOLS_DATABASE_URL=sqlite3://na:na@/tmp/sqlite/G2C.db

   ```

1. :pencil2: Run command.
   Example:

   ```console
   export LD_LIBRARY_PATH=/opt/senzing/er/lib/
   senzing-tools init-database

   ```

1. Remove `senzing-tools` from system.
   Example:

   ```console
   sudo apt-get remove senzing-tools

   ```

## Make documents

Make documents visible at
[hub.senzing.com/senzing-tools](https://hub.senzing.com/senzing-tools).

1. Identify repository.
   Example:

   ```console
   export GIT_ACCOUNT=senzing-garage
   export GIT_REPOSITORY=senzing-tools
   export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
   export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

   ```

1. Make documents.
   Example:

   ```console
   export LD_LIBRARY_PATH=/opt/senzing/er/lib/
   senzing-tools docs --dir ${GIT_REPOSITORY_DIR}/docs

   ```

## References

[clone-repository]: https://github.com/senzing-garage/knowledge-base/blob/main/HOWTO/clone-repository.md
[docker]: https://github.com/senzing-garage/knowledge-base/blob/main/WHATIS/docker.md
[git]: https://github.com/senzing-garage/knowledge-base/blob/main/WHATIS/git.md
[Go Reference Badge]: https://pkg.go.dev/badge/github.com/senzing-garage/senzing-tools.svg
[Go Reference]: https://pkg.go.dev/github.com/senzing-garage/senzing-tools
[go]: https://github.com/senzing-garage/knowledge-base/blob/main/WHATIS/go.md
[How to Install Senzing for Go Development]: https://github.com/senzing-garage/knowledge-base/blob/main/HOWTO/install-senzing-for-go-development.md
[localhost:6060]: http://localhost:6060/pkg/github.com/senzing-garage/senzing-tools/
[make]: https://github.com/senzing-garage/knowledge-base/blob/main/WHATIS/make.md
[testcoverage.yaml]: ../.github/coverage/testcoverage.yaml
