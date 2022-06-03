package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecord, err := net.LookupMX(domain) // check for mx record
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecorded, err := net.LookupTXT(domain) // check for spf record
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecorded {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecorded, err := net.LookupTXT("_dmarc." + domain) // check for marc record
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecorded {
		if strings.HasSuffix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}