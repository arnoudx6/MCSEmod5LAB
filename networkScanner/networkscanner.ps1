#Megacorp network scanner. Only works for /24 subnets.
$subnetPrefix = "192.168.56."
0..256 | foreach{
    if((ping $($subnetPrefix + $_) -n 1 -w 50) -like 'Reply*'){
        write-host $($subnetPrefix + $_)
    }
}