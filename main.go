package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	EMPTYSTRING = ""
	CLITOOLNAME = "mmtrace"
	MAXHOPS     = 64
	PACKETSIZE  = 32
)

func main() {
	var (
		command       string
		websiteDomain string
	)

	// take input
	_, err := fmt.Scanln(&command, &websiteDomain)
	if err != nil {
		log.Printf("err: %v; unable to read input\n", err.Error())
		os.Exit(1)
	}

	// basic validation of provided input
	if commandErr := validateCommand(command); commandErr != nil {
		log.Printf("err: %v\n", commandErr.Error())
		os.Exit(1)
	}
	if domainErr := validateDomain(websiteDomain); domainErr != nil {
		log.Printf("err: %v\n", domainErr.Error())
		os.Exit(1)
	}

	websiteDomainIPAddr, err := fetchIPViaDNSLookup(websiteDomain)
	if err != nil {
		log.Printf("err: %v, unable to resolve website domain\n", err.Error())
		os.Exit(1)
	}

	if websiteDomainIPAddr == EMPTYSTRING {
		log.Printf("err: unable to find IP Addr for: %v\n", websiteDomain)
		os.Exit(1)
	}

	log.Printf("traceroute to %v (%v), %d hops max, %d bytes packets\n", websiteDomain, websiteDomainIPAddr, MAXHOPS, PACKETSIZE)
	os.Exit(1)
}

func validateCommand(command string) error {
	if command == EMPTYSTRING || command != CLITOOLNAME {
		return errors.New("invalid command")
	}

	return nil
}

func validateDomain(domain string) error {
	if domain == EMPTYSTRING {
		return errors.New("invalid argument")
	}
	return nil
}

func fetchIPViaDNSLookup(websiteDomain string) (string, error) {
	dnsResponse, err := net.LookupIP(websiteDomain)
	if err != nil {
		return "", err
	}

	if len(dnsResponse) > 0 {
		return dnsResponse[0].String(), nil
	}

	return "", nil
}
