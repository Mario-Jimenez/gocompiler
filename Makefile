.PHONY: parser

parser:
	java -jar ./tools/antlr-4.8-complete.jar -Dlanguage=Go -o ./parser -Xexact-output-dir -no-listener -no-visitor ./grammar/MonkeyLexer.g4
	java -jar ./tools/antlr-4.8-complete.jar -Dlanguage=Go -o ./parser -Xexact-output-dir -no-listener -no-visitor ./grammar/MonkeyParser.g4

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o compiler main.go

run:
	./compiler
