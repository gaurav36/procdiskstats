/* Copyright 2017 Gaurav Kumar Garg
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *  
 * Package procdiskstats provide an interface to /proc/diskstats
 * 
 * The /proc/diskstats file displays the I/O statistics of block devices.
 * Each line contains the following 14 fields:
 * 
 * Each set of stats only applies to the indicated device; if you want
 * system-wide stats you'll have to find all the devices and sum them all up.
 * 
 *  1:   major number
 *  2:   minor number
 *  3:   device name
 *  4:   reads completed successfully
 *  5:   reads merged
 *  6:   sectors read
 *  7:   time spent reading (ms)
 *  8:   writes completed
 *  9:   writes merged
 *  10:  sectors written
 *  11:  time spent writing (ms)
 *  12:  I/Os currently in progress
 *  13:  time spent doing I/Os (ms)
 *  14:  weighted time spent doing I/Os (ms)
 *  
 *  for more details you can refer to the linux kernel Documentation/iostats.txt
 *  
 *  Once the information was updated you can access these disk states metrics
 *  
 *  import "github.com/gaurav36/procdiskstats"
 *  
 *  err := procdiskstats.Update()
 *
 *  for accessing diskstats metrics information you need to call respective function
 *  with exact disk name as a argument to the function. Disk name can be found in
 *  cat /proc/diskstats or by executing fdisk -l command.
 * 
 *  for eg: getting information about I/Os currently in progress
 *  
 *  ios, err : procdiskstats.IoInProgress ("sda1")
 * 
 */

package procdiskstats

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
        "strconv"
)

// Diskstats is a array with all value found in /proc/diskstats
var DiskStats [][]string

func Update() error {

	path := filepath.Join("/proc/diskstats")
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		words := strings.Fields(text)

		// Store only if information is in correct format
		if len(words) == 14 {
			DiskStats = append(DiskStats, words)
		}
	}

	return nil
}

// PrintDiskStats will print all information contaion in /proc/diskstats
func PrintDiskStats() {
	fmt.Println(DiskStats)
}

// ReadCompleted will give you  total number of reads completed successfully.
func ReadCompleted(diskd string) (uint64, error){
	// Search whether given diskd is available on device or not
	for i:=0; i<len(DiskStats); i++  {
		if DiskStats[i][2] == diskd {
			read,_ := strconv.ParseUint(DiskStats[i][3], 0, 64);
			return read, nil;
		}
	}
	_, err := fmt.Sscanf("Disk %s not present on the device", diskd)
	return 0,err
}

// ReadMerged will give you reads_merged field of disk name diskd
func ReadMerged(diskd string) (uint64, error) {
	// Search whether given diskd is available on device or not
	for i:=0; i<len(DiskStats); i++  {
		if DiskStats[i][2] == diskd {
			readmerged,_ := strconv.ParseUint(DiskStats[i][4], 0, 64);
			return readmerged, nil;
		}
	}
	_, err := fmt.Sscanf("Disk %s not present on the device", diskd)
	return 0,err
}

// SectorsRead will give you total number of sectors read successfully.
func SectorsRead(diskd string) (uint64, error) {
	// Search whether given diskd is available on device or not
	for i:=0; i<len(DiskStats); i++  {
		if DiskStats[i][2] == diskd {
			sectorsread,_ := strconv.ParseUint(DiskStats[i][5], 0, 64);
			return sectorsread, nil;
		}
	}
	_, err := fmt.Sscanf("Disk %s not present on the device", diskd)
	return 0,err
}

// ReadingMilliSecond will give you total number of milliseconds spent by all reads by disk diskd
func ReadingMilliSecond(diskd string) (uint64, error) {
	// Search whether given diskd is available on device or not
	for i:=0; i<len(DiskStats); i++  {
		if DiskStats[i][2] == diskd {
			readtime,_ := strconv.ParseUint(DiskStats[i][6], 0, 64);
			return readtime, nil;
		}
	}
	_, err := fmt.Sscanf("Disk %s not present on the device", diskd)
	return 0,err
}

// WriteCompleted will give you total number of writes completed successfully of disk diskd
func WriteCompleted(diskd string) (uint64, error) {
	// Search whether given diskd is available on device or not
	for i:=0; i<len(DiskStats); i++  {
		if DiskStats[i][2] == diskd {
			sectorsread,_ := strconv.ParseUint(DiskStats[i][7], 0, 64);
			return sectorsread, nil;
		}
	}
	_, err := fmt.Sscanf("Disk %s not present on the device", diskd)
	return 0,err
}

// WriteMerged will give you write merged information of disk diskd
func WriteMerged(diskd string) (uint64, error) {
	// Search whether given diskd is available on device or not
	for i:=0; i<len(DiskStats); i++  {
		if DiskStats[i][2] == diskd {
			sectorsread,_ := strconv.ParseUint(DiskStats[i][8], 0, 64);
			return sectorsread, nil;
		}
	}
	_, err := fmt.Sscanf("Disk %s not present on the device", diskd)
	return 0,err
}

// SectorWritten will give you total number of sectors written successfully.
func SectorWritten(diskd string) (uint64, error) {
	// Search whether given diskd is available on device or not
	for i:=0; i<len(DiskStats); i++  {
		if DiskStats[i][2] == diskd {
			sectorsread,_ := strconv.ParseUint(DiskStats[i][9], 0, 64);
			return sectorsread, nil;
		}
	}
	_, err := fmt.Sscanf("Disk %s not present on the device", diskd)
	return 0,err
}

// WriteMilliSecond will give you milliseconds spent during by all write of disk diskd
func WriteMilliSecond(diskd string) (uint64, error) {
	// Search whether given diskd is available on device or not
	for i:=0; i<len(DiskStats); i++  {
		if DiskStats[i][2] == diskd {
			sectorsread,_ := strconv.ParseUint(DiskStats[i][10], 0, 64);
			return sectorsread, nil;
		}
	}
	_, err := fmt.Sscanf("Disk %s not present on the device", diskd)
	return 0,err
}

// IoInProgress will give you I/Os currently in progress of disk diskd
func IoInProgress(diskd string) (uint64, error) {
	// Search whether given diskd is available on device or not
	for i:=0; i<len(DiskStats); i++  {
		if DiskStats[i][2] == diskd {
			sectorsread,_ := strconv.ParseUint(DiskStats[i][11], 0, 64);
			return sectorsread, nil;
		}
	}
	_, err := fmt.Sscanf("Disk %s not present on the device", diskd)
	return 0,err
}

// IoMilliSecond will give you milliseconds spent doing I/Os
func IoMilliSecond(diskd string) (uint64, error) {
	// Search whether given diskd is available on device or not
	for i:=0; i<len(DiskStats); i++  {
		if DiskStats[i][2] == diskd {
			sectorsread,_ := strconv.ParseUint(DiskStats[i][12], 0, 64);
			return sectorsread, nil;
		}
	}
	_, err := fmt.Sscanf("Disk %s not present on the device", diskd)
	return 0,err
}

// IoMilliSecondWeighted will give you weighted of milliseconds spent doing I/Os
func IoMilliSecondWeighted(diskd string) (uint64, error) {
	// Search whether given diskd is available on device or not
	for i:=0; i<len(DiskStats); i++  {
		if DiskStats[i][2] == diskd {
			sectorsread,_ := strconv.ParseUint(DiskStats[i][13], 0, 64);
			return sectorsread, nil;
		}
	}
	_, err := fmt.Sscanf("Disk %s not present on the device", diskd)
	return 0,err
}
