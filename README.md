# Go fundamentals (https://tour.golang.org)

## Installation (Windows 10)

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

## Run the Hello World Go script

```powershell
go run 01_hello.go
```
