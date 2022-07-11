# Detailed Installation Instructions

<!-- @import "[TOC]" {cmd="toc" depthFrom=2 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [Detailed Installation Instructions](#detailed-installation-instructions)
  - [Manual install](#manual-install)
    - [Download the latest binary](#download-the-latest-binary)
    - [wget](#wget)
  - [MacOS / Linux via Homebrew install](#macos--linux-via-homebrew-install)
  - [Docker image pull](#docker-image-pull)
    - [One-shot container use](#one-shot-container-use)
    - [Run container commands interactively](#run-container-commands-interactively)
  - [Go install](#go-install)

<!-- /code_chunk_output -->


## Manual install

### [Download the latest binary](https://github.com/mikfreedman/harvey/releases/latest)

### wget
Use wget to download the pre-compiled binaries:

```bash
wget https://github.com/mikfreedman/harvey/releases/download/${VERSION}/${BINARY}.tar.gz -O - |\
  tar xz && mv ${BINARY} /usr/bin/harvey
```

For instance, VERSION=v0.3.1 and BINARY=harvey_${VERSION}_linux_amd64

## MacOS / Linux via Homebrew install

Using [Homebrew](https://brew.sh/)  

```bash
brew tap mikfreedman/tap
brew install harvey
```

## Docker image pull

```bash
docker pull ghcr.io/mikfreedman/harvey
```

### One-shot container use

```bash
docker run --rm -v "${PWD}":/workdir ghcr.io/mikfreedman/harvey [flags]
```


### Run container commands interactively

```bash
docker run --rm -it -v "${PWD}":/workdir --entrypoint sh ghcr.io/mikfreedman/harvey
```

It can be useful to have a bash function to avoid typing the whole docker command:

```bash
harvey() {
  docker run --rm -i -v "${PWD}":/workdir ghcr.io/mikfreedman/harvey "$@"
}
```


## Go install

```bash
GO111MODULE=on go get github.com/mikfreedman/harvey/cmd/harvey
```