#!/bin/bash
set -o errexit

# ubuntu@myserver:~$ docker pull --platform linux/arm64 gcr.io/google-samples/hello-app:1.0
# 1.0: Pulling from google-samples/hello-app
# Digest: sha256:b1455e1c4fcc5ea1023c9e3b584cd84b64eb920e332feff690a2829696e379e7
# Status: Image is up to date for gcr.io/google-samples/hello-app:1.0
# image with reference gcr.io/google-samples/hello-app:1.0 was found but does not match the specified platform: wanted linux/arm64, actual: linux/amd64

# https://github.com/traefik/whoami
# https://hub.docker.com/r/stefanscherer/whoami/tags?page=&page_size=&ordering=&name=arm
# stefanscherer's docker image is badly tagged for linux/amd64 even for his arm releases, but it will still do for the test.
export TEST_IMG="stefanscherer/whoami:linux-arm64-2.0.1"


kubectl create ns test

kubectl -n test create deployment whoami-server --image=${TEST_IMG}

kubectl -n test get all



exit 0

# (kindest/node:v1.30.0)