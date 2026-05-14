# Deploy script for gitx
# Usage: ./deploy.ps1 [targetDir]
# If targetDir is not provided, defaults to user's Go bin directory.

param(
    [string]$TargetDir = $env:GOBIN
)

$ErrorActionPreference = "Stop"

if (-not $TargetDir) {
    $TargetDir = Join-Path $env:USERPROFILE "go\bin"
}

$src = "gitx.exe"
$dest = Join-Path $TargetDir $src

if (-not (Test-Path $src)) {
    Write-Error "gitx.exe not found. Run build.ps1 first."
    exit 1
}

Write-Host "Deploying gitx.exe to $dest"
Copy-Item $src $dest -Force

Write-Host "Deployed gitx.exe to $dest"
