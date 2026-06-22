The powershell command to test my localhost http server.
```powershell
Invoke-WebRequest -UseBasicParsing -Uri "http://localhost:8080/" -TimeoutSec 5
```

First, confirm what port is listening (pick your port):
```powershell
PS D:\Project\gin_tutor1> netstat -ano | findstr LISTENING | findstr :8080
  TCP    0.0.0.0:8080           0.0.0.0:0              LISTENING       30780
  TCP    [::]:8080              [::]:0                 LISTENING       30780
```




```powershell
PS D:\Project\code\Go\az\gin_tutor1> Get-Content .\http_test.ps1
try {
  $r = Invoke-WebRequest -Uri "http://127.0.0.1:8080/users" -Method Get -TimeoutSec 5 -ErrorAction Stop
  $r.StatusCode
  $r.Content
} catch {
  $_.Exception.Message
  $_.Exception.Response.StatusCode
}
```

### Let invoke the powershell script `http_test.ps1`:
```powershell
PS D:\Project\code\Go\az\gin_tutor1> ./

Security Warning: Script Execution Risk
Invoke-WebRequest parses the content of the web page. Script code in the web page might be run when the page is parsed.
      RECOMMENDED ACTION:
      Use the -UseBasicParsing switch to avoid script code execution.

      Do you want to continue?

[Y] Yes  [A] Yes to All  [N] No  [L] No to All  [S] Suspend  [?] Help (default is "N"): a
200
{"email":"john@doe.com"}
```