# Prepare the environment

```powershell
mkdir myapp
cd myapp
go mod init main
go get gioui.org@latest
go mod tidy
```

c:\Users\User\go\pkg\mod\
https://pkg.go.dev/gioui.org/example
https://jonegil.github.io/gui-with-gio/egg_timer/05_button_low_refactored.html



Set GOSUMDB to a valid checksum database (or disable verification). Recommended: restore default GOSUMDB.

Temporary for current PowerShell session:
```powershell
$env:GOSUMDB="sum.golang.org"
go clean -modcache
go mod tidy
```

Persist across sessions:
```powershell
setx GOSUMDB "sum.golang.org"
```
Close and reopen PowerShell after setx.

If you must disable checksum verification (not recommended):
```powershell
$env:GOSUMDB="off"
go mod tidy
```

Verify current settings:
```powershell
go env GOPROXY GOSUMDB GONOPROXY
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