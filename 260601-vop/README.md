### 2026-06-01
-------------------

[1]
To initialize the plain system you should add the nessessary dependency on "excelize/v2"
```powershell
go mod init myproject   # initialize module
go get github.com/xuri/excelize/v2   # add dependency
go mod tidy             # clean up go.mod and go.sum
```

[2]
