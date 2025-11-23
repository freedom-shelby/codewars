# Codewars
Solutions for codewars problems I am not too ashamed to put on github.

from Codewars wiki:
Codewars is a community of developers, who are called Code Warriors (or just warriors), that train on improving their development skills. Think of it like a coding dojo - where developers train with each other and help each other get better through practice

Programming languages are Go and PHP.

My current codewars rank:

<img src=https://www.codewars.com/users/arsen-hovhannisyan/badges/large>

## Setup

### Go Module Initialization

To run Go tests, initialize the module from the root directory:

```bash
go mod init github.com/freedom-shelby/codewars
```

Then you can run tests in any kata folder:

```bash
cd 8kyu/even_or_odd
go test              # Runs default solution (main.go)
go test -tags v2     # Runs alternative solution (solution2.go)
```
