````powershell
# Simple integration test script for the http_gin_rest2 server.
# - Builds and runs the server
# - Waits until the server responds
# - Exercises basic CRUD endpoints under /videos (best-effort)
# - Cleans up the server process
Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition
Push-Location $scriptDir

# Build server executable
$exe = Join-Path $scriptDir 'http_gin_test_server.exe'
if (Test-Path $exe) { Remove-Item $exe -Force -ErrorAction SilentlyContinue }
Write-Host "Building server..."
& go build -o $exe server.go

# Start server
Write-Host "Starting server..."
$proc = Start-Process -FilePath $exe -PassThru -WindowStyle Hidden
try {
    $base = 'http://localhost:8080'
    # Wait for server to become ready
    $ready = $false
    for ($i=0; $i -lt 30; $i++) {
        try {
            Invoke-WebRequest -Uri $base -UseBasicParsing -TimeoutSec 2 | Out-Null
            $ready = $true
            break
        } catch { Start-Sleep -Seconds 1 }
    }
    if (-not $ready) { throw "Server did not become ready on $base within timeout." }

    Write-Host "Server ready at $base"

    # Determine basic auth credentials (from env or defaults)
    $user = $env:BASIC_AUTH_USER
    $pass = $env:BASIC_AUTH_PASS
    if ([string]::IsNullOrEmpty($user)) { $user = 'admin' }
    if ([string]::IsNullOrEmpty($pass)) { $pass = 'admin' }
    $authHeader = @{
        Authorization = 'Basic ' + [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("$user`:$pass"))
    }

    function Invoke-Api($method, $path, $body = $null, $useAuth = $false) {
        $uri = "$base$path"
        $headers = @{}
        if ($useAuth) { $headers = $authHeader }
        if ($null -ne $body) {
            $json = $body | ConvertTo-Json -Depth 10
            return Invoke-WebRequest -Uri $uri -Method $method -Body $json -Headers $headers -ContentType 'application/json' -UseBasicParsing -ErrorAction Stop
        } else {
            return Invoke-WebRequest -Uri $uri -Method $method -Headers $headers -UseBasicParsing -ErrorAction Stop
        }
    }

    # Helper to try without auth first, then with auth if 401
    function Invoke-ApiWithAuthFallback($method, $path, $body = $null) {
        try {
            return Invoke-Api $method $path $body $false
        } catch [System.Net.WebException] {
            $resp = $_.Exception.Response
            if ($resp -and $resp.StatusCode.Value__ -eq 401) {
                Write-Host "Endpoint requires auth, retrying with credentials..."
                return Invoke-Api $method $path $body $true
            }
            throw
        }
    }

    # 1) GET /videos (list)
    Write-Host "GET /videos"
    $r = Try-Api -method 'GET' -path '/videos'
    Write-Host "Status:" $r.StatusCode
    Write-Host $r.Content

    # 2) POST /videos (create)
    $new = @{
        id = 1
        title = "PowerShell Test Video"
        author = "test"
        playTime = 123
        likes = 0
    }
    Write-Host "POST /videos"
    $r = Try-Api -method 'POST' -path '/videos' -body $new
    Write-Host "Status:" $r.StatusCode
    Write-Host $r.Content

    # 3) GET /videos/1
    Write-Host "GET /videos/1"
    $r = Try-Api -method 'GET' -path '/videos/1'
    Write-Host "Status:" $r.StatusCode
    Write-Host $r.Content

    # 4) PUT /videos/1 (update)
    $upd = @{
        id = 1
        title = "Updated Title"
        author = "tester"
        playTime = 321
        likes = 10
    }
    Write-Host "PUT /videos/1"
    $r = Try-Api -method 'PUT' -path '/videos/1' -body $upd
    Write-Host "Status:" $r.StatusCode
    Write-Host $r.Content

    # 5) DELETE /videos/1
    Write-Host "DELETE /videos/1"
    $r = Try-Api -method 'DELETE' -path '/videos/1'
    Write-Host "Status:" $r.StatusCode
    Write-Host $r.Content

    # 6) Confirm deletion
    Write-Host "GET /videos/1 (should be missing)"
    try {
        $r = Try-Api -method 'GET' -path '/videos/1'
        Write-Host "Status:" $r.StatusCode
        Write-Host $r.Content
    } catch {
        Write-Host "Expected error or empty response after delete: $($_.Exception.Message)"
    }

    Write-Host "Tests completed."
} finally {
    if ($proc -and -not $proc.HasExited) {
        Write-Host "Stopping server (PID $($proc.Id))..."
        Stop-Process -Id $proc.Id -Force -ErrorAction SilentlyContinue
    }
    if (Test-Path $exe) { Remove-Item $exe -Force -ErrorAction SilentlyContinue }
    Pop-Location
}
```// filepath: d:\Project\code\Go\az\http_gin_rest2\http_gin_test2.ps1
# Simple integration test script for the http_gin_rest2 server.
# - Builds and runs the server
# - Waits until the server responds
# - Exercises basic CRUD endpoints under /videos (best-effort)
# - Cleans up the server process
Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition
Push-Location $scriptDir

# Build server executable
$exe = Join-Path $scriptDir 'http_gin_test_server.exe'
if (Test-Path $exe) { Remove-Item $exe -Force -ErrorAction SilentlyContinue }
Write-Host "Building server..."
& go build -o $exe server.go

# Start server
Write-Host "Starting server..."
$proc = Start-Process -FilePath $exe -PassThru -WindowStyle Hidden
try {
    $base = 'http://localhost:8080'
    # Wait for server to become ready
    $ready = $false
    for ($i=0; $i -lt 30; $i++) {
        try {
            Invoke-WebRequest -Uri $base -UseBasicParsing -TimeoutSec 2 | Out-Null
            $ready = $true
            break
        } catch { Start-Sleep -Seconds 1 }
    }
    if (-not $ready) { throw "Server did not become ready on $base within timeout." }

    Write-Host "Server ready at $base"

    # Determine basic auth credentials (from env or defaults)
    $user = $env:BASIC_AUTH_USER
    $pass = $env:BASIC_AUTH_PASS
    if ([string]::IsNullOrEmpty($user)) { $user = 'admin' }
    if ([string]::IsNullOrEmpty($pass)) { $pass = 'admin' }
    $authHeader = @{
        Authorization = 'Basic ' + [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("$user`:$pass"))
    }

    function Invoke-Api($method, $path, $body = $null, $useAuth = $false) {
        $uri = "$base$path"
        $headers = @{}
        if ($useAuth) { $headers = $authHeader }
        if ($null -ne $body) {
            $json = $body | ConvertTo-Json -Depth 10
            return Invoke-WebRequest -Uri $uri -Method $method -Body $json -Headers $headers -ContentType 'application/json' -UseBasicParsing -ErrorAction Stop
        } else {
            return Invoke-WebRequest -Uri $uri -Method $method -Headers $headers -UseBasicParsing -ErrorAction Stop
        }
    }

    # Helper to try without auth first, then with auth if 401
    function Invoke-ApiWithAuthFallback($method, $path, $body = $null) {
        try {
            return Invoke-Api $method $path $body $false
        } catch [System.Net.WebException] {
            $resp = $_.Exception.Response
            if ($resp -and $resp.StatusCode.Value__ -eq 401) {
                Write-Host "Endpoint requires auth, retrying with credentials..."
                return Invoke-Api $method $path $body $true
            }
            throw
        }
    }

    # 1) GET /videos (list)
    Write-Host "GET /videos"
    $r = Try-Api -method 'GET' -path '/videos'
    Write-Host "Status:" $r.StatusCode
    Write-Host $r.Content

    # 2) POST /videos (create)
    $new = @{
        id = 1
        title = "PowerShell Test Video"
        author = "test"
        playTime = 123
        likes = 0
    }
    Write-Host "POST /videos"
    $r = Try-Api -method 'POST' -path '/videos' -body $new
    Write-Host "Status:" $r.StatusCode
    Write-Host $r.Content

    # 3) GET /videos/1
    Write-Host "GET /videos/1"
    $r = Try-Api -method 'GET' -path '/videos/1'
    Write-Host "Status:" $r.StatusCode
    Write-Host $r.Content

    # 4) PUT /videos/1 (update)
    $upd = @{
        id = 1
        title = "Updated Title"
        author = "tester"
        playTime = 321
        likes = 10
    }
    Write-Host "PUT /videos/1"
    $r = Try-Api -method 'PUT' -path '/videos/1' -body $upd
    Write-Host "Status:" $r.StatusCode
    Write-Host $r.Content

    # 5) DELETE /videos/1
    Write-Host "DELETE /videos/1"
    $r = Try-Api -method 'DELETE' -path '/videos/1'
    Write-Host "Status:" $r.StatusCode
    Write-Host $r.Content

    # 6) Confirm deletion
    Write-Host "GET /videos/1 (should be missing)"
    try {
        $r = Try-Api -method 'GET' -path '/videos/1'
        Write-Host "Status:" $r.StatusCode
        Write-Host $r.Content
    } catch {
        Write-Host "Expected error or empty response after delete: $($_.Exception.Message)"
    }

    Write-Host "Tests completed."
} finally {
    if ($proc -and -not $proc.HasExited) {
        Write-Host "Stopping server (PID $($proc.Id))..."
        Stop-Process -Id $proc.Id -Force -ErrorAction SilentlyContinue
    }
    if (Test-Path $exe) { Remove-Item $exe -Force -ErrorAction SilentlyContinue }
    Pop-Location
}
