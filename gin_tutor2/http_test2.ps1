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