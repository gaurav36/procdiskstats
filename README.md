# procdiskstats
Package procdiskstats provides an interface for /proc/diskstats


  
The /proc/diskstats file displays the I/O statistics of block devices.
Each line contains the following 14 fields:
  
Each set of stats only applies to the indicated device; if you want
system-wide stats you'll have to find all the devices and sum them all up. 
```
   1:   major number
   2:   minor number
   3:   device name
   4:   reads completed successfully
   5:   reads merged
   6:   sectors read
   7:   time spent reading (ms)
   8:   writes completed
   9:   writes merged
   10:  sectors written
   11:  time spent writing (ms)
   12:  I/Os currently in progress
   13:  time spent doing I/Os (ms)
   14:  weighted time spent doing I/Os (ms)
```
   
For more details you can refer to the linux kernel Documentation/iostats.txt
   
Once the information was updated you can access these disk states metrics
   
```
   import "github.com/gaurav36/procdiskstats"
   err := procdiskstats.Update()
```
 
for accessing diskstats metrics information you need to call respective function
with exact disk name as a argument to the function. Disk name can be found in
   ```  
   cat /proc/diskstats
   ```
or by executing
   ```  
   fdisk -l
   ```
  
for eg: getting information about I/Os currently in progress of disk sda1
   
```  
 ios, err := procdiskstats.IoInProgress ("sda1")
```
