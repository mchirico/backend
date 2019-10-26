FROM golang:1.12.6-alpine3.10 AS build
ARG DOCKER_GIT_CREDENTIALS=""
RUN apk add --no-cache git


WORKDIR /go/src/project

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
COPY . /go/src/project/


RUN git config --global credential.helper store && echo "${DOCKER_GIT_CREDENTIALS}" > ~/.git-credentials

RUN go get -v -t -d ./...
RUN go build -o /bin/project

# This results in a single layer image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /bin/project /bin/project

COPY test-fixtures/dummy.csv /dummy.csv

ENTRYPOINT ["/bin/project"]
# CMD ["--help"]

# export DOCKER_GIT_CREDENTIALS="$(cat ~/.git-for-docker)"
# docker build . --build-arg DOCKER_GIT_CREDENTIALS -t backend

# docker run -p 8888:8888

# az acr login --name acrDevelopment
# az acr repository list --name acrDevelopment --output table


# docker tag backend acrDevelopment.azurecr.io/backend:v1
# docker push acrDevelopment.azurecr.io/backend:v1