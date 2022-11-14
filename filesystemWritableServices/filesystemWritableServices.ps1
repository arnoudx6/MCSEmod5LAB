$services = Get-CimInstance win32_service | Select-Object PathName #Get all services 
$services = $services | where {$_.PathName -ne $null} #Only services with a path
$servicePaths = $services | foreach{$_.PathName.Substring(0, $_.PathName.IndexOf('.exe'))} #Add +4 for adding .exe
$servicePaths = $servicePaths | foreach{$_ -replace '"'} #Replace the quotes from the path because powershell cant handle this
$servicePaths = $servicePaths | Select-Object -Unique #Get unique paths. Else you have svchost 20 times.

#Foreach service path
$servicePaths | foreach{
    #Checking effective permissions without a module is hard. So i cheat by testing if i can write in the folder containing the exe
    $serviceParrentDirectory = ($_ | split-path -Parent)
    try{ 
        $testFilePath = ("{0}\{1}" -f $serviceParrentDirectory, "test.txt") #Combine strings
        Out-File $testFilePath -Force #Writing an empty file
        Remove-Item $testFilePath -Force #Cleanup the file
        write-host $_ #If i get here the path is writable
    }
    catch{
        #Skip to next servicepath
    }
}
