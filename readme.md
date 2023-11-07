# Docker with go!

This repo shows how To build a docker image for apps written in Go.

every folder has a different example for you To try and learn about!

Hope you enjoy it!

Useful commands:

To build:
```bash
docker build . -t name_of_the_image
```
To run:
```bash
docker run -p 8080:8080 name_of_the_image
```
To run in the background:
```bash
docker run -d -p 8080:8080 name_of_the_image
```
To see the logs:
```bash
docker logs container_id
```
To stop the container:
```bash
docker stop container_id
```
To see the running containers:
```bash
docker ps
```