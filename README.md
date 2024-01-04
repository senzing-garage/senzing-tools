# senzing-tools

If you are beginning your journey with
[Senzing](https://senzing.com/),
please start with
[Senzing Quick Start guides](https://docs.senzing.com/quickstart/).

You are in the
[Senzing Garage](https://github.com/senzing-garage)
where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## :warning: WARNING: senzing-tools is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

`senzing-tools` is a suite of tools to help use the Senzing API.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing-garage/senzing-tools.svg)](https://pkg.go.dev/github.com/senzing-garage/senzing-tools)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing-garage/senzing-tools)](https://goreportcard.com/report/github.com/senzing-garage/senzing-tools)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/senzing-garage/senzing-tools/blob/main/LICENSE)

[![gosec.yaml](https://github.com/senzing-garage/senzing-tools/actions/workflows/gosec.yaml/badge.svg)](https://github.com/senzing-garage/senzing-tools/actions/workflows/gosec.yaml)
[![go-test-linux.yaml](https://github.com/senzing-garage/senzing-tools/actions/workflows/go-test-linux.yaml/badge.svg)](https://github.com/senzing-garage/senzing-tools/actions/workflows/go-test-linux.yaml)
[![go-test-darwin.yaml](https://github.com/senzing-garage/senzing-tools/actions/workflows/go-test-darwin.yaml/badge.svg)](https://github.com/senzing-garage/senzing-tools/actions/workflows/go-test-darwin.yaml)
[![go-test-windows.yaml](https://github.com/senzing-garage/senzing-tools/actions/workflows/go-test-windows.yaml/badge.svg)](https://github.com/senzing-garage/senzing-tools/actions/workflows/go-test-windows.yaml)

## Overview

Senzing's `senzing-tools` has the following tools/commands:

1. [init-database](https://github.com/senzing-garage/init-database) - Used to create a Senzing schema and configuration in PostgreSQL, MySQL, MsSQL and SQLite databases.
1. [observer](https://github.com/senzing-garage/observe) - Aggregates Observer messages
1. [serve-grpc](https://github.com/senzing-garage/serve-grpc) - A gRPC server of the Senzing API
1. [serve-http](https://github.com/senzing-garage/serve-http) - An HTTP server for Senzing Tools

### Install

1. Visit [Releases](https://github.com/senzing-garage/senzing-tools/releases) page.
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
    1. [init-database](https://github.com/senzing-garage/init-database#using-command-line-options)
    1. [serve-grpc](https://github.com/senzing-garage/serve-grpc#using-command-line-options)

### Using environment variables

Environment variables may be used instead of command-line options.
Each command has it's own list of supported environment variables.
So see a specific list, visit the appropriate command.

1. Commands:
    1. [init-database](https://github.com/senzing-garage/init-database#using-environment-variables)
    1. [serve-grpc](https://github.com/senzing-garage/serve-grpc#using-environment-variables)

### Using Docker

The `senzing-tools` can be run from the `senzing/senzing-tools` Docker container.
Each command has it's own list of supported environment variables and command line options.
So see a specific list, visit the appropriate command.

1. Commands:
    1. [init-database](https://github.com/senzing-garage/init-database#using-docker)
    1. [serve-grpc](https://github.com/senzing-garage/serve-grpc#using-docker)

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

1. [init-database](https://github.com/senzing-garage/init-database#parameters)
1. [serve-grpc](https://github.com/senzing-garage/serve-grpc#parameters)

## References

- [Command reference](docs/senzing-tools.md)
- [Development](docs/development.md)
- [Errors](docs/errors.md)
- [Examples](docs/examples.md)
