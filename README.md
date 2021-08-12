
# Cord Tools Base AEM Image
This project contains the Cord Tools base AEM docker images and the code to generate them. The current dockerfiles live in `versions/`, and extend a linux-based image and include dependencies that AEM environments typically require.

The base images contain a JDK version compatible with some or all of the 6.x AEM versions. They do ***not*** contain AEM, and require a volume that contains the following:

1. *license.properties*
2. An unpacked AEM crx-quickstart directory **or** An Adobe AEM Quickstart JAR

## Generating and Building Dockerfiles

To generate base dockerfiles from the configs in `config/`, run the following:

```bash
go run generate.go
```

You can build a Docker image with your own tag by running the following command in the root directory of this project:

```bash
docker build -t public.ecr.aws/cord-tools/base:<version-tag> -f ./versions/<jdk-version>/Dockerfile .
```

### Example

```bash
docker build -t public.ecr.aws/cord-tools/base:aemcloud-jdk-11-my-tag -f ./versions/jdk-11/Dockerfile .
```

## Running a Base Image

Run the following command to create a Docker container from a built image that was created with your own tag:

```bash
docker run [-p 4502:<port> | -p 4503:<port>] [-v <path/to/aem>:/opt/aem] public.ecr.aws/cord-tools/base:<version-tag>
```

## Persisting AEM 

Every new container will be a new install of AEM with no existing content. To persist the content, a volume can be mounted from the host into the container that has the following:

1. *license.properties*
2. **crx-quickstart**/

### Example

```bash
docker run --rm \
        --volume ${HOME}/backups/:/opt/backups \
        --volume ${HOME}/aem/:/opt/aem  \
        -p 4502:4502 \
        public.ecr.aws/cord-tools/base:aemcloud-jdk-8
```

The initial start of the container will put the correct files within the volume. Subsequent restarts will use the contents of the volume.
