# Recomended PowerShell Core Version: 7.5.0 or later

Write-Host "Building Quickly Browse for macOS (Apple Silicon, arm64)..."
Write-Host "Removing old executables..."
Remove-Item ./dist/mac/q-brow -Force

$env:GOOS = "darwin"
$env:GOARCH = "arm64"
go build -o ./dist/mac/q-brow main.go
Remove-Item Env:GOOS
Remove-Item Env:GOARCH
Write-Host "Building completed."
Write-Host "Executable created at: ./dist/mac/q-brow"
Write-Host "Server run command: ./dist/mac/q-brow"