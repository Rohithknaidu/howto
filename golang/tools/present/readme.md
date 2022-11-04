


how to use golang present tool


installation: 
```go install golang.org/x/tools/cmd/present@latest```

slide file:
simple slide file example below: basically the MD syntax
```simple.slice
Welcome to Go presentation

* Introduction

** sub topic

```

run the presentation:
```
present 
```


references:
- https://github.com/bvwells/go-present/ The repository illustrates how to use the go present tool.



step 1: 
make sure go modules are setup. 
if not
```
go mod init
```

step 2: 
make sure to run below. otherwise I was getting error `Couldn't find gopresent files: no required module provides package golang.org/x/tools/cmd/present;`
```
go get golang.org/x/tools/cmd/present
```

step 3: create the slide file. if using markdown, make sure to start the file heading with `#`
```simple.slice
# Welcome to Go presentation

* Introduction
some introduction

** sub topic
something
```

step 4: run present with defaults. make sure you change to the working directory
`present`

if you want to run with presenter notes enables, pass the flag
`present -notes`
when the presentation is running, pressing N will open the presenter mode in a new tab.
