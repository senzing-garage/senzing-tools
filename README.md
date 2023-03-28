# senzing-tools

## :warning: WARNING: senzing-tools is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

`senzing-tools` is a suite of tools to help use the Senzing API.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/senzing-tools.svg)](https://pkg.go.dev/github.com/senzing/senzing-tools)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/senzing-tools)](https://goreportcard.com/report/github.com/senzing/senzing-tools)
[![go-test.yaml](https://github.com/Senzing/senzing-tools/actions/workflows/go-test.yaml/badge.svg)](https://github.com/Senzing/senzing-tools/actions/workflows/go-test.yaml)

## Overview

Senzing's `senzing-tools` has the following tools/commands:

1. [init-database](https://github.com/Senzing/init-database) - Used to create a Senzing schema and configuration in PostgreSQL, MySQL, MsSQL and SQLite databases.
1. [servegrpc](https://github.com/Senzing/servegrpc) - A gRPC server of the Senzing API

### Install

1. Visit [Releases](https://github.com/Senzing/senzing-tools/releases) page.
1. For the desired versioned release, in the "Assets" section,
   download the appropriate installation package.
    1. Use `.deb` file for Debian, Ubuntu and
       [others](https://en.wikipedia.org/wiki/List_of_Linux_distributions#Debian-based)
    1. Use `.rpm` file for Red Hat, CentOS, openSuse and
       [others](https://en.wikipedia.org/wiki/List_of_Linux_distributions#RPM-based).

1. :pencil2: Example installation for `.deb` file:

    ```console
    sudo apt install ./senzing-tools-0.0.0.deb
    ```

1. :pencil2: Example installation for `.rpm` file:

    ```console
    sudo yum install ./senzing-tools-0.0.0.rpm
    ```

## Use

1. **Important:** Prior to using the `senzing-tools` command,
   the `LD_LIBRARY_PATH` environment variable must be set
   to the location of the Senzing binaries.
   Example:

    ```console
    export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
    ```

### Using command line options

Simple examples.

1. :pencil2: A `senzing-tools init-database` example:

    ```console
    export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
    senzing-tools init-database \
        --database-url postgresql://username:password@postgres.example.com:5432/G2
    ```

Each command has it's own list of supported command line options.
Documentation for the command line options:

1. Online documentation, see
   [hub.senzing.com/senzing-tools](https://hub.senzing.com/senzing-tools)

1. Runtime documentation, run:

    ```console
    export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
    senzing-tools [command] --help
    ```

1. Detailed documentation, visit:
    1. [init-database](https://github.com/Senzing/init-database#using-command-line-options)
    1. [servegrpc](https://github.com/Senzing/servegrpc#using-command-line-options)

### Using environment variables

Environment variables may be used instead of command-line options.
Each command has it's own list of supported environment variables.
So see a specific list, visit the appropriate command.

1. Commands:
    1. [init-database](https://github.com/Senzing/init-database#using-environment-variables)
    1. [servegrpc](https://github.com/Senzing/servegrpc#using-environment-variables)

### Using Docker

The `senzing-tools` can be run from the `senzing/senzing-tools` Docker container.
Each command has it's own list of supported environment variables and command line options.
So see a specific list, visit the appropriate command.

1. Commands:
    1. [init-database](https://github.com/Senzing/init-database#using-docker)
    1. [servegrpc](https://github.com/Senzing/servegrpc#using-docker)

This usage shows how to initialze a database with a Docker container.

1. :pencil2: A `senzing/senzing-tools init-database`example:

    ```console
    docker run \
        --env SENZING_TOOLS_DATABASE_URL=postgresql://username:password@postgres.example.com:5432/G2 \
        senzing/senzing-tools init-database
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
            senzing/senzing-tools init-database
        ```

### Parameters

See individual commands for parameters:

1. [init-database](https://github.com/Senzing/init-database#parameters)
1. [servegrpc](https://github.com/Senzing/servegrpc#parameters)

## References

- [Command reference](docs/senzing-tools.md)
- [Development](docs/development.md)
- [Errors](docs/errors.md)
- [Examples](docs/examples.md)
