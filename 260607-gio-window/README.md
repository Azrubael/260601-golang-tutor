# Prepare the environment

```powershell
mkdir myapp
cd myapp
go mod init win_gio
# setx GOSUMDB sum.golang.org
go get gioui.org@latest
go mod tidy
```

### https://docs.fyne.io/
c:\Users\User\go\pkg\mod\




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