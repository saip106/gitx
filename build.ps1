# Build script for gitx
# Usage: ./build.ps1 [version]
# If version is not provided, defaults to 'dev'.

param(
    [string]$Version = "dev"
)

$ErrorActionPreference = "Stop"

Write-Host "Building gitx with version: $Version"

go build -ldflags="-X 'main.Version=$Version'" -o gitx.exe ./cmd/gitx

if ($LASTEXITCODE -ne 0) {
    Write-Error "Build failed."
    exit 1
}

Write-Host "Build succeeded. Binary: gitx.exe"

# Show version
./gitx.exe --version
