netstat -ano | findstr LISTENING | findstr :8080

Write-Output "Hello World!"
Read-Host "Press Enter to exit"

try {
  $r = Invoke-WebRequest -Uri "http://127.0.0.1:8080/test" -Method Get -TimeoutSec 5 -ErrorAction Stop
  $r.StatusCode
  $r.Content
} catch {
  $_.Exception.Message
  $_.Exception.Response.StatusCode
}

# Test GET /videos
Write-Output "`n=== Testing GET /videos ==="
try {
  $getVideos = Invoke-WebRequest -Uri "http://127.0.0.1:8080/videos" -Method Get -TimeoutSec 5 -ErrorAction Stop
  Write-Output "Status Code: $($getVideos.StatusCode)"
  Write-Output "Content: $($getVideos.Content)"
} catch {
  Write-Output "Error: $($_.Exception.Message)"
  Write-Output "Status Code: $($_.Exception.Response.StatusCode)"
}

# Test POST /videos
Write-Output "`n=== Testing POST /videos ==="
try {
  $videoData = @{
    title = "Test Video"
    description = "A test video"
    url = "https://example.com/video"
  } | ConvertTo-Json

  $postVideo = Invoke-WebRequest -Uri "http://127.0.0.1:8080/videos" -Method Post -ContentType "application/json" -Body $videoData -TimeoutSec 5 -ErrorAction Stop
  Write-Output "Status Code: $($postVideo.StatusCode)"
  Write-Output "Content: $($postVideo.Content)"
} catch {
  Write-Output "Error: $($_.Exception.Message)"
  Write-Output "Status Code: $($_.Exception.Response.StatusCode)"
}

Read-Host "Press Enter to exit"