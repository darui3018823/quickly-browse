# Recomended PowerShell Core Version: 7.5.0 or later

Write-Host "Building Quickly Browse for macOS (Intel, amd64)..."
Write-Host "Removing old executables..."
Remove-Item ./dist/mac/q-brow-intel -Force

$env:GOOS = "darwin"
$env:GOARCH = "amd64"
go build -o ./dist/mac/q-brow-intel main.go
Remove-Item Env:GOOS
Remove-Item Env:GOARCH
Write-Host "Building completed."
Write-Host "Executable created at: ./dist/mac/q-brow-intel"
Write-Host "Server run command: ./dist/mac/q-brow-intel"