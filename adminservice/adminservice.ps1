$user = New-LocalUser 'usrPentestMegacorp' -FullName 'Pentest Megacorp' -Password (ConvertTo-SecureString -AsPlainText '123456789' -Force)
Add-LocalGroupMember -Member $user -Group 'administrators'