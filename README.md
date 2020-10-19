# monkeycompiler

Compiler for the Monkey programming language written in Go.

## Usage

A Makefile is included for convenience.

#### Download dependencies

Get dependencies with `$ go get ./...`.

#### Generate ANTLR files for Go

Generate files with `$ make parser`.

#### Build project

Build project with `$ make build`.

#### Run project

Run project with `$ make run` or with binary file `$ ./compiler`.

## Web service

#### Configuration

Base URL: http://locahost:4440/monkeycompiler

#### Endpoints

###### /compile

**Request:**

```
{"program": string}
```

**Response:**

```
{
    "errors": []string,
    "lines": []int,
    "tree": {...},
}
```
