﻿(Get-ItemProperty -path "HKLM:\SOFTWARE\Policies\Microsoft\Windows\Installer" -Name "AlwaysInstallElevated")-and
(Get-ItemProperty -path "HKCU:\SOFTWARE\Policies\Microsoft\Windows\Installer" -Name "AlwaysInstallElevated")