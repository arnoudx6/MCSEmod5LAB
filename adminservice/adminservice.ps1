Set-Location $PSScriptRoot #Change the working dir to the script folder
Start-Transcript ".\log.txt" #Log all console output to a relative path for debugging purpose
$user = New-LocalUser 'usrPentestMegacorp' -FullName 'Pentest Megacorp' -Password (ConvertTo-SecureString -AsPlainText '123456789' -Force) #Create a new user
Add-LocalGroupMember -Member $user -Group 'administrators' #Add the user to the admin group
Stop-Transcript