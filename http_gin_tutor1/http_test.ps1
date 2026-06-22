try {
  $r = Invoke-WebRequest -Uri "http://127.0.0.1:8080/users" -Method Get -TimeoutSec 5 -ErrorAction Stop
  $r.StatusCode
  $r.Content
} catch {
  $_.Exception.Message
  $_.Exception.Response.StatusCode
}