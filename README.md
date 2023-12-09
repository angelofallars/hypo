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

## Status

Currently implemented commands:

Literals
  - [x] `<s>`
  - [x] `<data>`
  - [ ] `<ol>`
  - [ ] `<table>`

Math Commands

  - [x] `<dd>`
  - [x] `<sub>`
  - [x] `<ul>`
  - [x] `<div>`

Stack Manipulation Commands
  - [x] `<dt>`
  - [x] `<del>`

Comparison Commands
  - [ ] `<big>`
  - [ ] `<small>`
  - [ ] `<em>`

Logical Operators
  - [ ] `<b>`
  - [ ] `<bdi>`
  - [ ] `<bdo>`

Control Flow
  - [ ] `<i>`
  - [ ] `<rt>`
  - [ ] `<a>`

Variables
  - [x] `<var>`
  - [x] `<cite>`

I/O
  - [ ] `<input>`
  - [x] `<output>`
  - [ ] `<wbr>`

Properties
  - [ ] `<rp>`
  - [ ] `<samp>`

Arrays/Dynamic Properties
  - [ ] `<address>`
  - [ ] `<ins>`

Functions
  - [ ] `<dfn>`

Programs
  - [ ] `<main>`
  - [ ] `<body>`
