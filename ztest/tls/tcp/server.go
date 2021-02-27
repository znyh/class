package main

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
)

func main() {

	log.SetFlags(log.Lshortfile)

	cer, err := tls.LoadX509KeyPair("cert.pem", "key.pem")

	if err != nil {

		log.Println(err)

		return

	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	ln, err := tls.Listen("tcp", ":8000", config)

	if err != nil {

		log.Println(err)

		return

	}

	defer ln.Close()

	for {

		conn, err := ln.Accept()

		if err != nil {

			log.Println(err)

			continue

		}

		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	r := bufio.NewReader(conn)

	for {

		msg, err := r.ReadString('\n')

		if err != nil {

			log.Println(err)

			return

		}

		println(msg)

		n, err := conn.Write([]byte("world\n"))

		if err != nil {

			log.Println(n, err)

			return

		}

	}

}

var (
	_cert = `-----BEGIN CERTIFICATE-----
MIID5TCCAs2gAwIBAgIJAPX3G0uznEvZMA0GCSqGSIb3DQEBCwUAMIGIMQswCQYD
VQQGEwJDTjERMA8GA1UECAwIY2hpbmEtc3oxEjAQBgNVBAcMCWd1YW5nZG9uZzES
MBAGA1UECgwJc2hlbmd6aGVuMQ0wCwYDVQQLDAR6bnloMQ0wCwYDVQQDDAR6bnlo
MSAwHgYJKoZIhvcNAQkBFhExNTE4Njg0Nzk0QHFxLmNvbTAeFw0xOTEwMjcwNTA4
MzNaFw0yOTEwMjQwNTA4MzNaMIGIMQswCQYDVQQGEwJDTjERMA8GA1UECAwIY2hp
bmEtc3oxEjAQBgNVBAcMCWd1YW5nZG9uZzESMBAGA1UECgwJc2hlbmd6aGVuMQ0w
CwYDVQQLDAR6bnloMQ0wCwYDVQQDDAR6bnloMSAwHgYJKoZIhvcNAQkBFhExNTE4
Njg0Nzk0QHFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAO41
RiLFUVeBv0eBgxCQ0WoUULpCqOJyIUbDYRLScY8FFg4aTtOWG+lUqBnaOOMm9Jmu
tPScPIBJw6SC7pD48QweRSTCEeJ6XcGNEnB4ETQt0p2i2M+a1Ibng4iHxpsnhNZC
Bd3A4wF9cPNXtRGM3MSn11oC60cu945MwMOHllOj5gDAtr/ehFNI5AHC/UL2WqnV
RLhu73tF9KT99rbXNJzco6nVWNxhnWTug1j4Cdf0oL99gRpkQmuFjkbkZJPgkNlF
2o/8ouxggPw5kuKzhVh+OkeDeOmJd/Y4SNA2fVhKl7pXrWeDUHRTDrzX/5FVolNR
/SsyH+26eebAoQKJsUMCAwEAAaNQME4wHQYDVR0OBBYEFNvN7+z8XMLbEb2F7MD6
LTGCtZV8MB8GA1UdIwQYMBaAFNvN7+z8XMLbEb2F7MD6LTGCtZV8MAwGA1UdEwQF
MAMBAf8wDQYJKoZIhvcNAQELBQADggEBAMe2OOZXx9EumtnXOQwt+yuluxhjTNGW
raiKw/M1pHVy44so0wqyBpzfQddTwewkxjMnvjmJbhSEGxXiMkyCIijzhtH2I8Xf
e3LMDvz8xNDxjIcJgmvNuKH6KLCdWZHRkxyfX7AVabeHROlRZ5Dcv5scy98NGh2z
ZyqUHWAlc5FgIg4O//asr2YK11QA2E6JfsneSRILBbVacy3vxCKBmhAXUjMSwp2a
Pqi71w4KO6mazLaaiWFev8b4AY+UrBejlcRlYD/K8XKoIP0VI44KNQwpI1p9WurO
LrcJ0OqNlSMUnzcNbpnoDAtCKqzk7HWB/lcDpvCW9Mdb7Sn/etiiHtA=
-----END CERTIFICATE-----
`

	_key = "-----BEGIN RSA PRIVATE KEY-----MIIEpAIBAAKCAQEA7jVGIsVRV4G/R4GDEJDRahRQukKo4nIhRsNhEtJxjwUWDhpO05Yb6VSoGdo44yb0ma609Jw8gEnDpILukPjxDB5FJMIR4npdwY0ScHgRNC3SnaLYz5rUhueDiIfGmyeE1kIF3cDjAX1w81e1EYzcxKfXWgLrRy73jkzAw4eWU6PmAMC2v96EU0jkAcL9QvZaqdVEuG7ve0X0pP32ttc0nNyjqdVY3GGdZO6DWPgJ1/Sgv32BGmRCa4WORuRkk+CQ2UXaj/yi7GCA/DmS4rOFWH46R4N46Yl39jhI0DZ9WEqXuletZ4NQdFMOvNf/kVWiU1H9KzIf7bp55sChAomxQwIDAQABAoIBAQCS3TmuumUaRJ6AIbODBSZ39qqHDA4//wnRLSiuiB0HhqVAcKvk/AmdZFp4BflI432vOu3KWoFavx/mT+tyammDhS4wKY8JSVSAvs99lClXOGFAW2S9KsKZoieQh9XFVyopR2+Cdyf+hS6ceZjjhyud+7Vkg6Q3sRdkHqJCa1gdnCDuz/tZkXVer+K7ZwyxIj/IL4ck0X6UpPjOn0Ki1RsLCaz1o8DT8xLYJqOl+opHiPyJQQ2vv5fT3SrCL8SuOpln+eM9Nx8xZXa+/HLHx7otwsknBPd42A1CQHpzXDtDQpNZDOdwDbTEdgeMVXEWszQO6+hi3sIhH0A1dLXLw7PRAoGBAPmguuCfWeyJmCGqZvWvSQSl2DvEdf3ar4VZS+yvsg/OdFftrn+d2DsPeeCoFQHGjrJ6LDG99zpP0Hl/QCN9pmUnUuDMSXalHVb1K6YcfhDBOXumcRtEkNbCSexGrUj0/NRfQMI+XuqXyRLnJf7TofTAY/at+wVwGCPDdPpaS3Y3AoGBAPRJ6wkRn7qlWoi3XzjhS8ku3C6sbP4BoMCzAT1yo1n/9eyGybx+qsskzP6N+dneG+qu6S7qzmA9vIPstvogiJvfCdO+SXP4Q8i/GURruB9+PIcd+QuV8iqo/bO65jvnvwDgtSseqxopXE2mzg+dx+JS+j61BoPapWHAyEdgx5dVAoGAeBbC0FzYvdpTvZ5keb0hu4KtS9Nvj/gsiFI1HGBJoSEJB2Polqww4fjIATPbJ0eGToZHzIY+8WaEPG7jRC+mZPNOA6dDkDhOrQ61OOxsGVNbfAjOUPfjVe2R8oDtzDNPjjtkxA2NS/5JVTauKLFXMI99h6abJDqQ4mOdWvMlanECgYEAxwSxhuHkz5qG93rtCNoDdBPJLf0rjoSlAljTWcxbaKcvI1NsZzSK0JVarMfeFbDt7XQyZRbGq4GgY5omOADpSpOa3fbYwLIyBwoo+6UGjCG8DhhNHTM756Vli1do2sUixwRKlKaXnuizWk/a8L4KakhPt2ajuCgZNFpq0Dv4X0ECgYA4IgudmJak/JyZDFydg3X+KGGuJR+w3JK9SqHliX4vgK6ZOijxYjmZOtmbAxBdA04TJzgNv0yXKQO7EoD5nvK5XERgksw3N2a2U5YGNk41ALH3kkWOMQJia5j/Umgat8Wd8gGIdXbiYcjFdiaem+XeqqAv4ikRHbHwUGs1k2mDiQ==-----END RSA PRIVATE KEY-----"
)