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

1. Subcommands:
    1. [initdatabase](https://github.com/Senzing/initdatabase#https://github.com/Senzing/initdatabase#using-command-line-options)
    1. [servegrpc](https://github.com/Senzing/servegrpc#using-command-line-options)

Remember to start the commands with `senzing-tools [subcommand] ...`

1. :pencil2: A `senzing-tools initdatabase` example:

    ```console
    export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
    senzing-tools initdatabase \
        --database-url postgresql://username:password@postgres.example.com:5432/G2
    ```

### Using environment variables

Environment variables may be used instead of command-line options.
Each subcommand has it's own list of supported environment variables.
So see a specific list, visit the appropriate subcommand.

1. Subcommands:
    1. [initdatabase](https://github.com/Senzing/initdatabase#using-environment-variables)
    1. [servegrpc](https://github.com/Senzing/servegrpc#using-environment-variables)

Remember to start the commands with `senzing-tools [subcommand] ...`.

### Using Docker

The `senzing-tools` can be run from the `senzing/senzing-tools` Docker container.
Each subcommand has it's own list of supported environment variables and command line options.
So see a specific list, visit the appropriate subcommand.

1. Subcommands:
    1. [initdatabase](https://github.com/Senzing/initdatabase#using-docker)
    1. [servegrpc](https://github.com/Senzing/servegrpc#using-docker)

Remember to use the `senzing/senzing-tools` Docker image.

This usage shows how to initialze a database with a Docker container.

1. :pencil2: A `senzing/senzing-tools initdatabase`example:

    ```console
    docker run \
        --env SENZING_TOOLS_DATABASE_URL=postgresql://username:password@postgres.example.com:5432/G2 \
        senzing/senzing-tools initdatabase
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
            senzing/senzing-tools initdatabase
        ```

### Parameters

See individual commands for parameters:

1. [initdatabase](https://github.com/Senzing/initdatabase#parameters)
1. [servegrpc](https://github.com/Senzing/servegrpc#parameters)

## References

- [Examples](docs/examples.md)
- [Development](docs/development.md)
