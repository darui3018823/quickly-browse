# Recomended PowerShell Core Version: 7.5.0 or later

Write-Host "Building Quickly Browse for Windows (amd64)..."
Write-Host "Removing old executables..."
Remove-Item ./dist/win/q-brow.exe -Force

$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -ldflags="-H=windowsgui" -o ./dist/win/q-brow.exe main.go
Remove-Item Env:GOOS
Remove-Item Env:GOARCH
Write-Host "Building completed."
Write-Host "Executable created at: ./dist/win/q-brow.exe"
Write-Host "Server run command: ./dist/win/q-brow.exe"