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

***Flags***

- *-vm*: path to Monkey Virtual Machine binary (**required**)
- *-code*: path to Monkey Instructions Code file
- *-h*: help

Run project with `$ make run` or with binary file `$ ./compiler -vm="path/to/file"`.

#### Main Function Template (not required)

```
let Main = fn(main) {
  // Your code here
}
```

## Web service

#### Configuration

Base URL: http://locahost:4440/monkeycompiler

#### Endpoints

###### /compile

**Request:**

```
{
    "program": string,
}
```

**Response:**

```
{
    "errors": []string,
    "lines": []int,
    "tree": {...},
}
```
It also creates a file with Monkey Instructions Code in the path specified with the flag `-code`.

###### /run

**Request:**

```
Empty
```

**Response:**

```
{
    "result": string,
}
```

