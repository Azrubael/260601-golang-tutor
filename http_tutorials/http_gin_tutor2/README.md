```powershell
go mod init gin_tutor2
go mod tidy
go run .
```

### In another powershell
```powershell
PS D:\Project\code\Go\az\gin_tutor2> Get-Content .\http_test.ps1
netstat -ano | findstr LISTENING | findstr :8080

Write-Output "Hello World!"
Read-Host "Press Enter to exit"

try {
  $r = Invoke-WebRequest -Uri "http://127.0.0.1:8080/ping" -Method Get -TimeoutSec 5 -ErrorAction Stop
  $r.StatusCode
  $r.Content
} catch {
  $_.Exception.Message
  $_.Exception.Response.StatusCode
}
```

### Invoke the powershell script `http_test.ps1`:
```powershell
PS D:\Project\code\Go\az\gin_tutor2> ./http_test2.ps1

  TCP    0.0.0.0:8080           0.0.0.0:0              LISTENING       16060
  TCP    [::]:8080              [::]:0                 LISTENING       16060
Hello World!
Press Enter to exit:


Security Warning: Script Execution Risk
Invoke-WebRequest parses the content of the web page. Script code in the web page might be run when
 the page is parsed.
      RECOMMENDED ACTION:
      Use the -UseBasicParsing switch to avoid script code execution.

      Do you want to continue?

[Y] Yes  [A] Yes to All  [N] No  [L] No to All  [S] Suspend  [?] Help (default is "N"): a
200
{"message":"pong"}
```