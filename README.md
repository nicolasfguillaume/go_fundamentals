# golang fundamentals

## Advantages/drawbacks of Go vs Python

See https://getstream.io/blog/switched-python-go/

## Installation (Win10)

```powershell
choco install -y golang
```

Then restart the terminal and test the installation:
```powershell
go version
```

Display env variables:
```powershell
go env
```

## Create a Go workspace

See https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-windows-10

It is located at `$env:GOPATH`: `C:\Users\Nicolas\go` 

Best practice is:
```
└── $HOME
    └── go
        ├── bin
        └── src
```

When Go compiles and installs tools, it will put them in the $GOPATH/bin directory. 

## Run Hello World

```powershell
go run 01_hello.go
```

Built-in linter:
```powershell
gofmt -w 01_hello.go
```

Convert Go to C code (C shared library):
```
go build -buildmode=c-shared -o 01_hello.so 01_hello.go
```
This creates 2 files, 01_hello.h and 01_hello.so. The 01_hello.so file is the compiled binary and 01_hello.h is the header to describe what methods and types are defined within the binary. Then the shared library can be imported in C or Python...
