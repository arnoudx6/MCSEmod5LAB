Set-Location $PSScriptRoot
$currentUser = [System.Security.Principal.WindowsIdentity]::GetCurrent().Name
$currentUser | Out-File .\username.txt