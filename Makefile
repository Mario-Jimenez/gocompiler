.PHONY: parser

parser:
	java -jar /usr/local/lib/antlr-4.8-complete.jar -Dlanguage=Go -no-listener -no-visitor ./parser/Compiler.g4

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o compiler main.go

run:
	./compiler