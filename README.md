# Hypo

Hypo is a hyper-fast runtime for [HTML, the programming language](https://html-lang.org).

Run HTML, the programming language code outside of the browser.

## Installation

Requirements: [Go](https://go.dev) 1.21 or later

```bash
git clone https://github.com/angelofallars/hypo
cd hypo
go install ./cmd/hypo
```

## Usage

With no arguments, Hypo will spin up a REPL for you to type and run HTML, the programming language code. You can execute an `.html` file by passing the file name as an argument to Hypo.

```bash
$ hypo example/helloworld.html
Hello world!
```
