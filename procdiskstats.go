//    1: 'reads_completed',
//    2: 'reads_merged',
//    3: 'sectors_read',
//    4: 'reading_milliseconds',
//    5: 'writes_completed',
//    6: 'writes_merged',
//    7: 'sectors_written',
//    8: 'writing_milliseconds',
//    9: 'io_inprogress',
//    10: 'io_milliseconds',
//    11: 'io_milliseconds_weighted'

package procdiskstats

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

// ReadCompleted will print ReadCompletion field of disk name diskd
func ReadCompleted(diskd string) {
	// Search whether given diskd is available on device or not

}
