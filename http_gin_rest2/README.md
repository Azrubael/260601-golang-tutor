## Go Module Init
```powershell
go mod init github.com/Azrubael/260601-golang-tutor/http_gin_rest2
```

## Gin-Gonic library: github.com/gin-gonic/gin

## Run
```powershell
go run server.go
```

#### Run the test from the project folder with PowerShell):

````powershell
# change to the test directory
Set-Location 'D:\Project\code\Go\az\http_gin_rest2'

# optionally set basic-auth creds used by the test
$env:BASIC_AUTH_USER = 'admin'
$env:BASIC_AUTH_PASS = 'admin'

# run the integration test (verbose, avoid test caching)
go test -v -run TestIntegration -count=1
````

Or run all tests in that package:

````powershell
go test -v -count=1
````
