# senzing-tools examples

## Initialize database

### Initialize database using command line option

1. :pencil2: SQLite database.
   Example:

    ```console
    senzing-tools init-database --database-url sqlite3://na:na@/tmp/sqlite/G2C.db

    ```

1. :pencil2: PostgreSql database.
   Example:

    ```console
    senzing-tools init-database --database-url postgresql://postgres:postgres@postgres.example.com:5432/G2

    ```

### Initialize database using environment variable

1. :pencil2: SQLite database.
   Example:

    ```console
    export SENZING_TOOLS_DATABASE_URL=sqlite3://na:na@/tmp/sqlite/G2C.db
    senzing-tools init-database --database-url sqlite3://na:na@/tmp/sqlite/G2C.db

    ```

1. :pencil2: PostgreSql database.
   Example:

    ```console
    export SENZING_TOOLS_DATABASE_URL=postgresql://postgres:postgres@postgres.example.com:5432/G2
    senzing-tools init-database

    ```

### Initialize database using Docker

#### Initialize SQLite database

1. Make an empty directory for Docker volume.
   Example:

    ```console
    mkdir  /tmp/sqlite

    ```

1. Run `senzing/senzing-tools` Docker image with `init-database` command.
   Example:

    ```console
    docker run \
        --env SENZING_TOOLS_DATABASE_URL=sqlite3://na:na@/tmp/sqlite/G2C.db \
        --rm \
        --volume /tmp/sqlite:/tmp/sqlite \
        senzing/senzing-tools init-database

    ```

#### Initialize PostgreSql database

1. **Optional:** If an existing PostgreSql database doesn't exist,
   [bring up a new PostgreSql database in a docker-compose formation](#docker-compose-stack-with-postgresql-database).

1. :pencil2: Identify database URL, if using existing PostgreSql.
   Example:

    ```console
    export SENZING_TOOLS_DATABASE_URL=postgresql://postgres:postgres@postgres.example.com:5432/G2

    ```

1. Identify database URL, if using docker-compose stack.
   *Note:*  Assignment of `LOCAL_IP_ADDRESS` may not work in all cases.
   Examples:

    ```console
    export LOCAL_IP_ADDRESS=$(curl --silent https://raw.githubusercontent.com/Senzing/knowledge-base/main/gists/find-local-ip-address/find-local-ip-address.py | python3 -)
    export SENZING_TOOLS_DATABASE_URL=postgresql://postgres:postgres@${LOCAL_IP_ADDRESS}:5432/G2/?sslmode=disable

    ```

1. Run `senzing/senzing-tools` Docker image with `init-database` command.
   Example:

    ```console
    docker run \
        --env SENZING_TOOLS_DATABASE_URL \
        --rm \
        senzing/senzing-tools init-database

    ```

## Serve gRPC

### Serve gRPC using Docker

1. Start a gRPC server using an existing SQLite database.
   Example:

    ```console
    docker run \
        --env SENZING_TOOLS_DATABASE_URL=sqlite3://na:na@/tmp/sqlite/G2C.db \
        --rm \
        --volume /tmp/sqlite:/tmp/sqlite \
        senzing/senzing-tools servegrpc

    ```

## Miscellaneous

### Docker-compose stack with PostgreSql database

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

1. In PostgreSQL database, create Senzing schema and initial Senzing config.
   Example:

    ```console
    docker run \
        --env SENZING_TOOLS_DATABASE_URL \
        --rm \
        senzing/senzing-tools init-database

    ```

1. **Optional:** View the PostgreSQL database.

   Visit [localhost:9171](http://localhost:9171).
   For the initial login, review the instructions at the top of the web page.
   For server password information, see the `POSTGRESQL_POSTGRES_PASSWORD` value in `${SENZING_DEMO_DIR}/docker-compose.yaml`.
   Usually, it's "postgres".

1. Start a gRPC server using that database.
   Example:

    ```console
    docker run \
        --env SENZING_TOOLS_DATABASE_URL \
        --rm \
        senzing/senzing-tools servegrpc

    ```

1. Cleanup.
   Example:

    ```console
    cd ${SENZING_DEMO_DIR}
    sudo --preserve-env docker-compose down

    ```

### View SQLite database

1. **Optional:** View the database.
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

   Visit <http://localhost:9174>
