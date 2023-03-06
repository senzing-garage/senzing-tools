# senzing-tools

## :warning: WARNING: senzing-tools is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

The Senzing `senzing-tools` is a suite of tools to help
with common tasks when using the Senzing API.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/senzing-tools.svg)](https://pkg.go.dev/github.com/senzing/senzing-tools)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/senzing-tools)](https://goreportcard.com/report/github.com/senzing/senzing-tools)
[![go-test.yaml](https://github.com/Senzing/senzing-tools/actions/workflows/go-test.yaml/badge.svg)](https://github.com/Senzing/senzing-tools/actions/workflows/go-test.yaml)

## Overview

Senzing's `senzing-tools` has the following tools:

1. [initdatabase](https://github.com/Senzing/initdatabase) - Used to create a Senzing schema and configuration in PostgreSQL, MySQL, MsSQL and SQLite databases.
1. [servegrpc](https://github.com/Senzing/servegrpc) - A gRPC server of the Senzing API

## Use

```console
export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
senzing-tools [subcommand] [flags]
```

For subcommands and flags, see
[hub.senzing.com/senzing-tools](https://hub.senzing.com/senzing-tools/) or run:

```console
export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
senzing-tools --help
```

### Using command line options

Each subcommand has it's own list of supported command line options.
So see a specific list, run:

```console
export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
senzing-tools [subcommand] --help
```

Or visit the appropriate subcommand documentation.

1. Subcommands
    1. [initdatabase](https://github.com/Senzing/initdatabase#https://github.com/Senzing/initdatabase#using-command-line-options)
    1. [servegrpc](https://github.com/Senzing/servegrpc#using-command-line-options)

Remember to start the commands with `senzing-tools [subcommand] ...`.

1. :pencil2: A `senzing-tools initdatabase` example:

    ```console
    export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
    senzing-tools initdatabase --database-url postgresql://username:password@postgres.example.com:5432/G2
    ```

### Using environment variables

Environment variables may be used instead of command-line options.
Each subcommand has it's own list of supported environment variables.
So see a specific list, visit the appropriate subcommand.

1. Subcommands
    1. [initdatabase](https://github.com/Senzing/initdatabase#using-environment-variables)
    1. [servegrpc](https://github.com/Senzing/servegrpc#using-environment-variables)

Remember to start the commands with `senzing-tools [subcommand] ...`
or use the `SENZING_TOOLS_SUBCOMMAND` environment variable.

### Using Docker

The `senzing-tools` can be run from the `senzing/senzing-tools` Docker container.
Each subcommand has it's own list of supported environment variables and command line options.
So see a specific list, visit the appropriate subcommand.

1. Subcommands
    1. [initdatabase](https://github.com/Senzing/initdatabase#using-docker)
    1. [servegrpc](https://github.com/Senzing/servegrpc#using-docker)

Remember to use the `senzing/senzing-tools` Docker image.

This usage shows how to initialze a database with a Docker container.

1. :pencil2: A `senzing/senzing-tools initdatabase`example:

    ```console
    docker run \
        --env SENZING_TOOLS_DATABASE_URL=postgresql://username:password@postgres.example.com:5432/G2 \
        --env SENZING_TOOLS_SUBCOMMAND=initdatabase \
        senzing/senzing-tools
    ```

1. *Alternative:* Using `SENZING_TOOLS_ENGINE_CONFIGURATION_JSON` environment variable.

    1. :pencil2: Set `SENZING_TOOLS_ENGINE_CONFIGURATION_JSON`.
       Example:

        ```console
        export SENZING_TOOLS_ENGINE_CONFIGURATION_JSON='{
            "PIPELINE": {
                "CONFIGPATH": "/etc/opt/senzing",
                "RESOURCEPATH": "/opt/senzing/g2/resources",
                "SUPPORTPATH": "/opt/senzing/data"
            },
            "SQL": {
                "CONNECTION": "postgresql://username:password@postgres.example.com:5432:G2"
            }
        }'
        ```

    1. Run `senzing/senzing-tools`.
       **Note:** `SENZING_TOOLS_ENGINE_CONFIGURATION_JSON` superceeds use of `SENZING_TOOLS_DATABASE_URL`.
       Example:

        ```console
        docker run \
            --env SENZING_TOOLS_ENGINE_CONFIGURATION_JSON \
            --env SENZING_TOOLS_SUBCOMMAND=initdatabase \
            senzing/senzing-tools
        ```

### Parameters

See individual commands for parameters:

1. [initdatabase](https://github.com/Senzing/initdatabase#parameters)
1. [servegrpc](https://github.com/Senzing/servegrpc#parameters)

## Development

### Install Go

1. See Go's [Download and install](https://go.dev/doc/install)

### Install Senzing C library

Since the Senzing library is a prerequisite, it must be installed first.

1. Verify Senzing C shared objects, configuration, and SDK header files are installed.
    1. `/opt/senzing/g2/lib`
    1. `/opt/senzing/g2/sdk/c`
    1. `/etc/opt/senzing`

1. If not installed, see
   [How to Install Senzing for Go Development](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/install-senzing-for-go-development.md).

### Install Git repository

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=senzing
    export GIT_REPOSITORY=senzing-tools
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow steps in
   [clone-repository](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/clone-repository.md) to install the Git repository.

### Build

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

1. Clean up.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean

    ```

### Test using SQLite database

1. Run tests.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean test

    ```

1. **Optional:** View the SQLite database.
   Example:

    ```console
    docker run \
        --env SQLITE_DATABASE=G2C.db \
        --interactive \
        --publish 9174:8080 \
        --rm \
        --tty \
        --volume /tmp/sqlite:/data \
        coleifer/sqlite-web

    ```

   Visit [localhost:9174](http://localhost:9174).

### Test using Docker-compose stack with PostgreSql database

The following instructions show how to bring up a test stack to be used
in testing the `g2-sdk-go-base` packages.

1. Identify a directory to place docker-compose artifacts.
   The directory specified will be deleted and re-created.
   Example:

    ```console
    export SENZING_DEMO_DIR=~/my-senzing-demo

    ```

1. Bring up the docker-compose stack.
   Example:

    ```console
    export PGADMIN_DIR=${SENZING_DEMO_DIR}/pgadmin
    export POSTGRES_DIR=${SENZING_DEMO_DIR}/postgres
    export RABBITMQ_DIR=${SENZING_DEMO_DIR}/rabbitmq
    export SENZING_VAR_DIR=${SENZING_DEMO_DIR}/var
    export SENZING_UID=$(id -u)
    export SENZING_GID=$(id -g)

    rm -rf ${SENZING_DEMO_DIR:-/tmp/nowhere/for/safety}
    mkdir ${SENZING_DEMO_DIR}
    mkdir -p ${PGADMIN_DIR} ${POSTGRES_DIR} ${RABBITMQ_DIR} ${SENZING_VAR_DIR}
    chmod -R 777 ${SENZING_DEMO_DIR}

    curl -X GET \
        --output ${SENZING_DEMO_DIR}/docker-versions-stable.sh \
        https://raw.githubusercontent.com/Senzing/knowledge-base/main/lists/docker-versions-stable.sh
    source ${SENZING_DEMO_DIR}/docker-versions-stable.sh
    curl -X GET \
        --output ${SENZING_DEMO_DIR}/docker-compose.yaml \
        "https://raw.githubusercontent.com/Senzing/docker-compose-demo/main/resources/postgresql/docker-compose-postgresql-uninitialized.yaml"

    cd ${SENZING_DEMO_DIR}
    sudo --preserve-env docker-compose up

    ```

1. In a separate terminal window, set environment variables.
   Identify Database URL of database in docker-compose stack.
   Example:

    ```console
    export LOCAL_IP_ADDRESS=$(curl --silent https://raw.githubusercontent.com/Senzing/knowledge-base/main/gists/find-local-ip-address/find-local-ip-address.py | python3 -)
    export SENZING_TOOLS_DATABASE_URL=postgresql://postgres:postgres@${LOCAL_IP_ADDRESS}:5432/G2/?sslmode=disable

    ```

1. Run tests.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean test

    ```

1. **Optional:** View the PostgreSQL database.

   Visit [localhost:9171](http://localhost:9171).
   For the initial login, review the instructions at the top of the web page.
   For server password information, see the `POSTGRESQL_POSTGRES_PASSWORD` value in `${SENZING_DEMO_DIR}/docker-compose.yaml`.
   Usually, it's "postgres".

1. Cleanup.

    ```console
    cd ${SENZING_DEMO_DIR}
    sudo --preserve-env docker-compose down

    cd ${GIT_REPOSITORY_DIR}
    make clean

    ```

### Package

#### Package RPM and DEB files

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

#### Test DEB package on Ubuntu

1. Determine if `senzing-tools` is installed.
   Example:

    ```console
    apt list --installed | grep senzing-tools

    ```

1. :pencil2: Install `senzing-tools`.
   The `senzing-tools-...` filename will need modification.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}/target
    sudo apt install ./senzing-tools-0.0.0.deb

    ```

1. :pencil2: Identify database.
   One option is to bring up PostgreSql as see in
   [Test using Docker-compose stack with PostgreSql database](#test-using-docker-compose-stack-with-postgresql-database).
   Example:

    ```console
    export SENZING_TOOLS_DATABASE_URL=sqlite3://na:na@/tmp/sqlite/G2C.db

    ```

1. :pencil2: Run command.
   Example:

    ```console
    export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
    senzing-tools

    ```

1. Remove `senzing-tools` from system.
   Example:

    ```console
    sudo apt-get remove senzing-tools

    ```

#### Test RPM package on Centos

1. Determine if `senzing-tools` is installed.
   Example:

    ```console
    yum list installed | grep senzing-tools

    ```

1. :pencil2: Install `senzing-tools`.
   The `senzing-tools-...` filename will need modification.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}/target
    sudo yum install ./senzing-tools-0.0.0.rpm

    ```

1. :pencil2: Identify database.
   One option is to bring up PostgreSql as see in
   [Test using Docker-compose stack with PostgreSql database](#test-using-docker-compose-stack-with-postgresql-database).
   Example:

    ```console
    export SENZING_TOOLS_DATABASE_URL=sqlite3://na:na@/tmp/sqlite/G2C.db

    ```

1. Run command.
   Example:

    ```console
    export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
    senzing-tools

    ```

1. Remove `senzing-tools` from system.
   Example:

    ```console
    sudo yum remove senzing-tools

    ```

### Make documents

Make documents visible at
[hub.senzing.com/senzing-tools](https://hub.senzing.com/senzing-tools).

1. Identify repository.
   Example:

    ```console
    export GIT_ACCOUNT=senzing
    export GIT_REPOSITORY=senzing-tools
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Make documents.
   Example:

    ```console
    export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
    senzing-tools docs --dir ${GIT_REPOSITORY_DIR}/docs

    ```

## Error prefixes

Error identifiers are in the format `senzing-PPPPnnnn` where:

`P` is a prefix used to identify the package.
`n` is a location within the package.

Prefixes:

1. `6501` - initializer
1. `6502` - senzingconfig
1. `6503` - senzingschema
