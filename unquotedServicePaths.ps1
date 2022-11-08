#Retrieve all services
$services = Get-CimInstance win32_service | select State, Name, DisplayName, PathName, StartName

#Find unquoted service paths
#Paths that do not contain " quotes and do contain space chars before the .exe
$services | where{($_.PathName -notlike '"*"*') -and ($_.PathName -like '* *.exe*')} | FT