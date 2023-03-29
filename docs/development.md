# senzing-tools development

## Install Go

1. See Go's [Download and install](https://go.dev/doc/install)

## Install Senzing C library

Since the Senzing library is a prerequisite, it must be installed first.

1. Verify Senzing C shared objects, configuration, and SDK header files are installed.
    1. `/opt/senzing/g2/lib`
    1. `/opt/senzing/g2/sdk/c`
    1. `/etc/opt/senzing`

1. If not installed, see
   [How to Install Senzing for Go Development](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/install-senzing-for-go-development.md).

## Install Git repository

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=senzing
    export GIT_REPOSITORY=senzing-tools
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow steps in
   [clone-repository](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/clone-repository.md) to install the Git repository.

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

1. Clean up.
   Example:

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
   The `senzing-tools-...` filename will need modification.
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
    export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
    senzing-tools init-database

    ```

1. Remove `senzing-tools` from system.
   Example:

    ```console
    sudo apt-get remove senzing-tools

    ```

### Test RPM package on Centos

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
   Example:

    ```console
    export SENZING_TOOLS_DATABASE_URL=sqlite3://na:na@/tmp/sqlite/G2C.db

    ```

1. Run command.
   Example:

    ```console
    export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
    senzing-tools init-database

    ```

1. Remove `senzing-tools` from system.
   Example:

    ```console
    sudo yum remove senzing-tools

    ```

## Make documents

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
