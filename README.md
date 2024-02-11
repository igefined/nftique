# Backend service for NFT marketplace

[![Coverage Status](https://coveralls.io/repos/github/igefined/nftique/badge.svg?branch=main)](https://coveralls.io/github/igefined/nftique?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/igefined/nftique)](https://goreportcard.com/report/github.com/igefined/nftique)
![Go version](https://img.shields.io/github/go-mod/go-version/igefined/nftique)
![Last commit](https://img.shields.io/github/last-commit/igefined/nftique)

This is simple project for understanding how to work with abigen, smartcontracts in golang. It list, add and buy a NFT via smartcontracts
## Usage
```shell
#run linter
make lint

#build an application
make build

#build an application image
make build-image

#run the application
make run

#test application into container
make test

#upload new abi generated smartcontracts from https://github.com/igefined/nftique-smartcontracts
make generate-abigen
```