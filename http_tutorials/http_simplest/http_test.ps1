# Константа для базової адреси
$BaseUrl = "http://127.0.0.1:8080"

function Test-Endpoint {
    param (
        [string]$Url,
        [string]$Name
    )
    try {
        $r = Invoke-WebRequest -Uri $Url -Method Get -TimeoutSec 5 -ErrorAction Stop
        Write-Output "$Name OK"
        $r.StatusCode
        $r.Content
		Read-Host "Press Enter again:"
    } catch {
        $_.Exception.Message
        $_.Exception.Response.StatusCode
    }
}

# Хеш-таблиця з трьома парами ключ-значення
$endpoints = @{
    root  = "$BaseUrl/root"
    users = "$BaseUrl/users"
    posts = "$BaseUrl/posts"
}

# Перевірка порту
netstat -ano | findstr LISTENING | findstr :8080

Read-Host "Press Enter to exit"

# Виклик функції для кожної пари
foreach ($key in $endpoints.Keys) {
    Test-Endpoint -Url $endpoints[$key] -Name $key
}
Write-Output "The server $Name tested successfully!"