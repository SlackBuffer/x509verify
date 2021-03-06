package main

import (
	"crypto/x509"
	"encoding/pem"
)

func main() {
	// Verifying with a custom list of root certificates.
	const rootPEM = `
-----BEGIN CERTIFICATE-----
MIIFgzCCA2ugAwIBAgIJAJ9ZfMolM8QcMA0GCSqGSIb3DQEBCwUAMFgxCzAJBgNV
BAYTAmNuMQswCQYDVQQIDAJ6ajELMAkGA1UEBwwCaHoxITAfBgNVBAoMGEludGVy
bmV0IFdpZGdpdHMgUHR5IEx0ZDEMMAoGA1UEAwwDaGhoMB4XDTIyMDEwMjExNTg1
N1oXDTI0MTAyMjExNTg1N1owWDELMAkGA1UEBhMCY24xCzAJBgNVBAgMAnpqMQsw
CQYDVQQHDAJoejEhMB8GA1UECgwYSW50ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMQww
CgYDVQQDDANoaGgwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQCxR+aA
6Z3BSezrXK8iYlFKoYmaomb14ydTilk4qHRc5mCnL+PEIgdFcxsGIpGsy1u86/2L
uUzIeVnLJ7KlkNhGoNr2fEVdtctrO97hJdp6YZZ/77tKAs6RZAhpHlm767UlqCXQ
IQ9uF5XEIl126w54JEjhjh86Q9Hsm6rf/e1WP9Wc5KK6yubzIu31Wct9jQYcVqPR
bCsE5unzoik4j8jY1jYr0fKVmC+sHb7MqbLhe6X2zYvnykHCU8+fXBt5JEH99J/6
CtgdBf96TclpqKeOfXXTnn949uiFs2imwvC0cBXZxA1AWwFfLXcr+jT1zRJfu8Sp
ys2JbgeUz4/ffD2dAkUp4ADa/W+Gyuu7sbo/p0cuIHyO/JbwrO6ZuJYiveK2Z5BV
RWsk1JNb4tXfAgcuoXVie0nCErLz1ILek3OflxyjQKjpCdKJRPdCY1rLjnRnEoq3
9yyzRlfbmoeljJdZAOYTiVBin7AtB6tExbmQgheWzLEVWhwJZ30cR3Vo3yzrJIG8
bKPIIX/B8UmWl1Pz2BqP2nLc/f5myuDZtcePNSIYUx6FhDAXFM1nU3KlKZkaq7nO
LAL0mexjcTZ02ZCOF3miwXYQqNX4VRYiJr4OJBsxm02G+h7hogD82UuWWhmugutt
vwwSiQ2p+6KRLVUyDwWwWTHDhmFpwBTZj2xPKwIDAQABo1AwTjAdBgNVHQ4EFgQU
uSvjeH8igrb70GSVw4LkgWv/EJgwHwYDVR0jBBgwFoAUuSvjeH8igrb70GSVw4Lk
gWv/EJgwDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAgEAOk+9iYqBkARO
u1CdPQvST1FDMQaUYkZH151pmA4hCnBGHf01yrhCvVzezE1bLkzn6fL/JjliqTsE
Lixlp5JPsIey9h6IXeg4CoV6CfsnG9i4lpo5biCug8r88Wqv+G0aYnJcCX+0FFKH
k539CdTCqnziOxqOJu7rMemzbBPWr4WXcpKxEDiF6pUWE8a/+mEcf7z6KXwWcGJu
WRvRELItLpqonez46q3Kd/MjjXjrc3y1/cpEa/dtYG2ozl15FHh8lPHfxr/d3r91
tWh9hNNJ+uZg3uY4XO+vs1wPEd41MfWZEPwHdo1HTfOooqchBPwK0lmD5fUSviHx
DEMNdfTlpuH8mEDMaS17xfBGP0MCPeThh80iyCnRRTBZtP9dT11uAQUtIsXc/YzC
JV503SaBlJWXp8fAKmzgl97sOp0/zJ6rgeoA/7htilEALBUs27KrNDCqzTFUk/Yo
GbD+1Tb4+oXZrNaII0S9/h3QKxbDRb01bMzDM9oo0v5eJXqldtHcvQE1WrFJT+hG
k/hlZ1agluBtR/Yz5pX4Eqo9DAvrMThYPoK3dFQojWVxsqLykyj3uBJ0PBqCgGsP
rz8sjRzMCyJGqclURyjk7NY9o/YcPJM9EHi2U7qxASErjiaSQbe1G0vvPiDzPPCO
LSubQBWhJDoLyk9b5AEPRZTTtH87yA8=
-----END CERTIFICATE-----`

	const certPEM = `
-----BEGIN CERTIFICATE-----
MIIEGzCCAgMCCQCgG9TauwmWtTANBgkqhkiG9w0BAQsFADBYMQswCQYDVQQGEwJj
bjELMAkGA1UECAwCemoxCzAJBgNVBAcMAmh6MSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQxDDAKBgNVBAMMA2hoaDAeFw0yMjAxMDIxMjAyNDZaFw0y
NDEwMjIxMjAyNDZaMEcxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJDQTEUMBIGA1UE
CgwLTXlPcmcsIEluYy4xFTATBgNVBAMMDG15ZG9tYWluLmNvbTCCASIwDQYJKoZI
hvcNAQEBBQADggEPADCCAQoCggEBALfo+7EzFYDeAIvgUqYxu1YVpYlCxYN+8KqQ
ubl2aZxaKaYf6qKU3EgxGQPh5Xhzk5n6DGatSb9cHmbrQVguq9BmG/G86ZcIW4NH
AIu5JDy00r5QF87t4l5hq2N4f1syBZfl4e4iIRL2wOL17g5PA0QogS+MvkLboaZG
BcgtXZpeXAe1NhCKKLn74kYpr1bN7eRxk9scU3tFEIiZpYyTP5FT4DfvYqWv7335
1RntUh7pgcFpw5hOqSgTqYVL4qgOBGvm03pFil4vfzWu1rDfw7T7xzeKZqjb83o7
E8szRIG1Ev3Cl5ZGdGxO/jclMveq0J6oYZ/oNUHcwnhlhrBLI5ECAwEAATANBgkq
hkiG9w0BAQsFAAOCAgEAjSu5baD9eDwc1MJowmEU86YPyr+On9wXphVOklnM/Trs
nCEbJ669n3qpIIPDDsmkU1Uc/RF/ELqXO5XIErpwjEVTCh1vou5yjkSNVDS7H20b
d3CvtzSwT8Oxg3h1RxQW1eipVus7j1LsBySqUlRygRFxXrrkvTc/ex4qCGPuw5Bj
DHJw0Wj2hS1inm6/K8fxdnnlzbrXIvtVxe7KqjlJA/5KMQb8bTFvvCixMGxZel6F
6R8Y62D/j5mxqJVZi7cvfYuxmZHcepdaxruJ7DucOYywUwBNPXhsVr1OHMVPieWu
KFp5re7wdQhsPkyZevXzEG6pB0Wir0xbtnKLPih/Dp6ot7saowZ6GNX2Ol887D8A
muFXtV7pocRR4ULbLvGa0H1LXwzNpZSJq3tZ4YAnhDejeMvUjL7G5b1IdC33b4aO
nxkcqubXjNrHJfGnTjLUwUbYaV3fNpJHW9IlCrAaSTRKmvvV8OYWWHDP316hfWLq
WAAcBVN/PweFg04zRmrDN6wVBXhFsMKi/IOzcCAHUAwYglQZTANlX5o2mQF/+Uen
HRym0RR7Ax9n5vxXlsyFDNjwGgjIiBh3sAbP3uOg+s6SxdKxddqrbOurCeckTnbK
+oKIFH4QtbkA/Er1/nrR4H5e+LwQpHBxHy5QEGDFwsIPH2R3+4JZWrdOBegU30I=
-----END CERTIFICATE-----`

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}

	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		panic("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic("failed to parse certificate: " + err.Error())
	}

	opts := x509.VerifyOptions{
		Roots: roots,
	}

	if _, err := cert.Verify(opts); err != nil {
		panic("failed to verify certificate: " + err.Error())
	}
}
