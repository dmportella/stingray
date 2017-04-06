# stingray
`GOLANG` Client API for stringray.

@dmportella

[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/dmportella/stingray/master/LICENSE) [![Build Status](https://travis-ci.org/dmportella/stingray.svg?branch=master)](https://travis-ci.org/dmportella/stingray) [![GoDoc](https://godoc.org/github.com/dmportella/stingray?status.svg)](https://godoc.org/github.com/dmportella/stingray) [![Go Report Card](https://goreportcard.com/badge/github.com/dmportella/stingray)](https://goreportcard.com/report/github.com/dmportella/stingray) [![Github Release](https://img.shields.io/github/release/dmportella/stingray.svg)](https://github.com/dmportella/stingray/releases)

## Brocade Stingray API

This `GOLANG` client API was tested using a Brocade Stingray installation running version of the API at 3.9 and below.

## Design of the client

`Stingray API` is built as a restful API so the entire API is browsable from the browser e.g. you can just hit the API endpoint at `/` and from there you should be able to navigate the entire API on your browse. This is great but it increases the complexity of building a client that works in a browseable/restful manner and enforces the restful `way` on the client with that in mind the this client API was built so it doesnt try to enforce the user the restful doctrine instead you have all the tools to `browse` the API if you choose too.

## Supported resources

Not all the resources are available in this client API I am working on adding more support for the other types as time goes by but for now the api supports the following types on `GET` and `PUT` methods.

Resources

- Traffic IP Group
- Virtual Server
- Pool

Lists

- Action Lists

> Note: action list is a generic resource that is used across the entire Stingray Rest API to `list` resources and `action` lists.
