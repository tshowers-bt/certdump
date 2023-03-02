package main

import (
	"crypto/tls"
	"fmt"
)

const addr = "dl-cdn.alpinelinux.org:443"

func main() {

	tlsVersions := map[uint16]string{
		tls.VersionSSL30: "SSL 3.0",
		tls.VersionTLS10: "TLS 1.0",
		tls.VersionTLS11: "TLS 1.1",
		tls.VersionTLS12: "TLS 1.2",
		tls.VersionTLS13: "TLS 1.3",
	}

	tlsVersionName := func(v uint16) string {
		if name, ok := tlsVersions[v]; ok {
			return name
		}
		return fmt.Sprintf("unknown (%d)", v)
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		VerifyConnection: func(cs tls.ConnectionState) error {
			fmt.Printf("Server name: %s\n", cs.ServerName)
			fmt.Printf("TLS version: %s\n", tlsVersionName(cs.Version))
			fmt.Printf("Cipher suite: %s\n", tls.CipherSuiteName(cs.CipherSuite))
			fmt.Printf("Peer certificates:\n")
			for _, c := range cs.PeerCertificates {
				fmt.Printf("- Issuer: %v\n", c.Issuer)
				fmt.Printf("  Subject: %v\n", c.Subject)
				if len(c.DNSNames) != 0 {
					fmt.Printf("  DNS names:\n")
					for _, d := range c.DNSNames {
						fmt.Printf("  - %s\n", d)
					}
				}
			}
			return nil
		},
	}

	fmt.Printf("getting TLS certificate info for %s...\n", addr)

	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	defer conn.Close()
}
