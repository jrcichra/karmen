#!/bin/bash
sudo apt purge -y docker-ce
export DOCKER_CLI_EXPERIMENTAL=enabled
curl -fsSL https://get.docker.com/ -o docker-install.sh
CHANNEL=nightly sh docker-install.sh
export DOCKER_CLI_EXPERIMENTAL=enabled
docker version
sudo curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
docker-compose --version
docker buildx --help
docker run --rm --privileged aptman/qus -s -- -p
cat /proc/sys/fs/binfmt_misc/qemu-aarch64
docker buildx create --name testbuilder
docker buildx use testbuilder
docker buildx inspect --bootstrap
# Phase 2 - sign in
echo "$DOCKER_PASSWORD" | docker login --username "$DOCKER_USERNAME" --password-stdin 
# Phase 3 - build a container based on the arg passed in
cd karmen
docker build -t jrcichra/karmen:ci .
cd ..
# Phase 4 - build the CI test container
cd ci/python1
docker build --build-arg GITHUB_SHA=${GITHUB_SHA} -t jrcichra/python1:ci .
cd ../..
# Phase 5 - run the CI test docker-compose
cd ci
docker-compose up --exit-code-from python1
# Phase 6 - if that passes push this sha and latest
docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t jrcichra/karmen:${GITHUB_SHA:0:8} --push .
docker buildx imagetools inspect jrcichra/karmen:${GITHUB_SHA:0:8}
docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t jrcichra/karmen:latest --push .
docker buildx imagetools inspect jrcichra/karmen:latest
