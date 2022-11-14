$regpathServices = "HKLM:\SYSTEM\CurrentControlSet\Services" #Location of all services in the registry
foreach ($path in Get-ChildItem $regpathServices){ #Loop through all services
    #Create a test value to check if the path is writable
    try{
        $null = New-ItemProperty -Path $path.PSPath -Name "test" -Value "teststring" -PropertyType "String" -ErrorAction Stop #Try to write
        write-host $path #If i get here i can write
        $null = Remove-ItemProperty -Path $path.PSPath -Name "test" #Cleanup
        
    }
    catch{
        #Do nothing next itteration
    }
}