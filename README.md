# EAI-Docker-Image-Task

This is a simple task in to build a docker image and push it to a docker hub repository. From Enterprise Application Integration course at the Telkom University.

## How to use

Pull the image from docker hub

```bash
$ docker pull faruqihafiz/eai-docker-image-faruqi
```

Run the image

```bash
$ docker run -p 3000:3000 eai-docker-image-faruqi
```

## Test Localy

Install the library

```bash
$ go mod tidy
```

Run

```bash
$ go run .\cmd\ main.go
```
