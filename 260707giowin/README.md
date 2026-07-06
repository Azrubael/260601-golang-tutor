# Prepare the environment

```powershell
go mod init "github.com/Azrubael/260601-golang-tutor/260707giowin"
go mod tidy
```

The right set of the environment variables on Windows 11
```powershell
PS C:\> go env
set AR=ar
set CC=gcc
set CGO_CFLAGS=-O2 -g
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-O2 -g
set CGO_ENABLED=0
set CGO_FFLAGS=-O2 -g
set CGO_LDFLAGS=-O2 -g
set CXX=g++
set GCCGO=gccgo
set GO111MODULE=
set GOAMD64=v1
set GOARCH=amd64
set GOAUTH=netrc
set GOBIN=
set GOCACHE=C:\Users\User\AppData\Local\go-build
set GOCACHEPROG=
set GODEBUG=
set GOENV=C:\Users\User\AppData\Roaming\go\env
set GOEXE=.exe
set GOEXPERIMENT=
set GOFIPS140=off
set GOFLAGS=
set GOGCCFLAGS=-m64 -fno-caret-diagnostics -Qunused-arguments -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=C:\Users\User\AppData\Local\Temp\go-build3685056939=/tmp/go-build -gno-record-gcc-switches
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOINSECURE=
set GOMOD=NUL
set GOMODCACHE=C:\Users\User\go\pkg\mod
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
set GOPATH=C:\Users\User\go
set GOPRIVATE=
set GOPROXY=https://proxy.golang.org,direct
set GOROOT=C:\Program Files\Go
set GOSUMDB=sum.golang.org
set GOTELEMETRY=local
set GOTELEMETRYDIR=C:\Users\User\AppData\Roaming\go\telemetry
set GOTMPDIR=
set GOTOOLCHAIN=auto
set GOTOOLDIR=C:\Program Files\Go\pkg\tool\windows_amd64
set GOVCS=
set GOVERSION=go1.26.3
set GOWORK=
set PKG_CONFIG=pkg-config
```


## To run the tests for all the module`s code:
```pwsh
go test -v ./...
```

## To test all the code of the package `gio_win`
```pwsh
go test -v ./gio_win
```

## To run tests and generate coverage profile of the package `gio_win`
```pwsh
go test -coverprofile=coverage.out -v ./gio_win
```

## To run tests from the `gio_win_test` package
```pwsh
go test -v ./gio_win/gio_win_test
```

## To run an exact test function `TestReadShpkFile` in the package `gio_win`
```pwsh
go clean -testcache
go test -run -v ^TestReadShpkFile$ -v ./gio_win
```
    OR
```pwsh
go clean -testcache
go test -v -run ^TestReadShpkFile$ ./gio_win/gio_win_test
go test -v -run ^TestPrepareReportPPD$ ./gio_win/gio_win_test
```

#### To check all the list of the test files:
```pwsh
Get-ChildItem -Recurse -Filter "*_test.go"
go list -f "{{.TestGoFiles}}" ./gio_win/gio_win_test
```