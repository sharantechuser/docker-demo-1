# docker-demo-1
"Simple Docker Demo"

Creating a simple Go (Golang) project and packaging it into a Docker container involves a few basic steps: writing a simple Go application, creating a Dockerfile, and building a Docker image. Below is a step-by-step guide to create a minimal Go application and Dockerize it.

## Create a Dockerfile
In the same directory, create a Dockerfile without any file extension. This file will define the steps to build your Docker image.

````
FROM golang:latest
RUN mkdir /app
ADD . /app
# EXPOSE 4000
WORKDIR /app
RUN go build -o docker1 .
CMD [ "/app/docker1" ]

ENV USER_NAME Test
````

In this example, each line represents a step in building the Docker image:

`FROM`: Specifies the base image. In this case, we use the latest Ubuntu image as the base.

`WORKDIR`: Sets the working directory inside the container.

`COPY`: Copies files from the host into the container.

`RUN`: Executes commands within the container to install dependencies, build the application, etc.

`EXPOSE`: Exposes a port that the application will use (it's a declaration and doesn't actually publish the port).

`CMD`: Specifies the default command to run when a container is started from the image.

## Run

````
docker build -t go-docker-first .

docker run -p 4000:4000 go-docker-first

//Verify the container
docker ps -a

````


### curl http://localhost:4000