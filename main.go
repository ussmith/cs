package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/ussmith/cs/data"
	"github.com/ussmith/cs/processor"
)

func main() {

	var scannerList []processor.Scanner

	scannerList = append(scannerList, processor.FastScanner{
		Name: "fs1",
	})
	scannerList = append(scannerList, processor.SlowScanner{
		Name: "sl1",
	})
	scannerList = append(scannerList, processor.MediumScanner{
		Name: "ms1",
	})
	scannerList = append(scannerList, processor.FastScanner{
		Name: "fs2",
	})
	scannerList = append(scannerList, processor.FastScanner{
		Name: "fs3",
	})
	scannerList = append(scannerList, processor.FastScanner{
		Name: "fs4",
	})
	scannerList = append(scannerList, processor.SlowScanner{
		Name: "sl2",
	})
	scannerList = append(scannerList, processor.MediumScanner{
		Name: "ms2",
	})
	scannerList = append(scannerList, processor.SlowScanner{
		Name: "sl3",
	})
	scannerList = append(scannerList, processor.MediumScanner{
		Name: "ms3",
	})
	scannerList = append(scannerList, processor.SlowScanner{
		Name: "sl4",
	})
	scannerList = append(scannerList, processor.MediumScanner{
		Name: "ms4",
	})

	fmt.Printf("There are %d scanners\n", len(scannerList))
	chans := make(chan data.ScanStatus, len(scannerList))

	//var wg sync.WaitGroup
	fmt.Println("Creating the context")
	ct := context.Background()
	ctWithTimeout, f := context.WithTimeout(ct, time.Millisecond*2000)
	defer f()

	//wg.Add(len(scannerList))

	for _, v := range scannerList {
		go v.Scan(ctWithTimeout, "somewhere", chans)
	}

	md := string(md5.New().Sum([]byte("How now brown cow")))

	c := cache.New(5*time.Minute, 10*time.Minute)

	fmt.Println("Reading from the chans")

	scanPackage := data.ScanPackage{
		Viruses: make(map[string][]string),
	}

	received := 0
	for r := range chans {
		received++
		if r.Err != nil {
			scanPackage.FailedJobs++
			//fmt.Printf("Failed Scan Results: %v -- total = %d\n", r.Err, scanPackage.FailedJobs)
		} else {
			if r.Found == true {
				scanPackage.Found = true
				//fmt.Printf("Set found to %v\n", scanPackage.Found)
				scanPackage.Viruses[r.ScannerName] = r.Viruses
			}
			//	fmt.Printf("Received a result %v\n", r)
		}

		//process(scanPackage)
		//fmt.Printf("true or false %v\n", scanPackage.Found)
		c.Add(md, scanPackage, cache.DefaultExpiration)
		// raw, ok := c.Get(md)

		// if !ok {
		// 	fmt.Println("Where'd it go?")
		// 	os.Exit(-1)
		// }

		// sp := raw.(data.ScanPackage)
		// fmt.Println("After cache")
		// process(sp)
		// fmt.Printf("\n\n")

		if received == len(scannerList) {
			// Anti-Pattern?
			//fmt.Println("Received all, closing channel")
			close(chans)
			//raw, ok := c.Get(md)

			//if !ok {
			//fmt.Println("Where'd it go?")
			//os.Exit(-1)
			//}

			//sp := raw.(data.ScanPackage)
			//fmt.Printf("From cache, true or false: %v\n", sp.Found)

			process(scanPackage)
		}
	}
}

func process(sp data.ScanPackage) {
	fmt.Printf("Viruses found? %v\n", sp.Found)
	fmt.Printf("Failed Jobs? %d\n", sp.FailedJobs)

	for k, v := range sp.Viruses {
		fmt.Printf("%s\n", k)
		fmt.Println("-------------------------------------------------------------")

		for _, vv := range v {
			fmt.Printf("%s\n", vv)
		}
	}
}
