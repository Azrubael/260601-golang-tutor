Write-Output "Привіт з контейнера!"

# Показати загальну доступну пам'ять системи
$os = Get-CimInstance Win32_OperatingSystem
$freeMemMB = [math]::Round($os.FreePhysicalMemory / 1024, 2)
Write-Output "Доступна фізична пам'ять системи: $freeMemMB MB"

$counter = Get-Counter '\Memory\Available MBytes'
Write-Output ("Доступна системна пам'ять: {0} MB" -f $counter.CounterSamples.CookedValue)

# Показати використання пам'яті самим процесом PowerShell
$proc = Get-Process -Id $PID
$usedMemMB = [math]::Round($proc.WorkingSet64 / 1MB, 2)
Write-Output "Поточне використання пам'яті контейнером: $usedMemMB MB"

Write-Output "Введіть 'exit' щоб завершити контейнер..."

while ($true) {
    $input = Read-Host "Введіть команду"
    if ($input -eq "exit") {
        break
    } else {
        Write-Output "Ви ввели: $input"
    }
}