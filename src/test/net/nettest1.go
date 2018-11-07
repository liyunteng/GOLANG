package main

import (
	"fmt"
	"os"
	"net"
	"log"
	"io"
	"io/ioutil"
	"time"
)

// 将string类型的IP转换为IP
func TestNet_ParseIp() {
	name := "192.168.1.100"

	ip := net.ParseIP(name)
	if ip == nil {
		fmt.Fprintf(os.Stderr, "invalid ip")
		return
	}

	fmt.Printf("IP: %+v %T IPTo4: %#v %T\n", ip, ip,  ip.To4(), ip.To4())
	defaultMask := ip.DefaultMask()
	fmt.Printf("DefaultMask: %+v\n", defaultMask)

	ones, bits := defaultMask.Size()
	fmt.Printf("ones: %d bits: %d", ones, bits)
}

func TestNet_Mask() {
	name := "192.168.1.100"

	ip := net.ParseIP(name)
	if ip == nil {
		log.Fatal("invalid ip")
	}

	mask := ip.DefaultMask()
	network := ip.Mask(mask)

	fmt.Printf("network: %+v mask: %+v\n", network, mask)
}

func TestNet_LookupHost() {
	domain := "www.google.com"

	// ip, err := net.ResolveIPAddr("ip", domain)
	// if err != nil {
	//	log.Fatal(err)
	// }

	// fmt.Printf("%+v IP: %+v Network: %+v Zone: %+v\n",
	//	ip, ip.IP, ip.Network(), ip.Zone)

	// ns, err := net.LookupHost(domain)
	ns, err := net.LookupIP(domain)
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range ns {
		fmt.Println(n)
	}
}

func TestNet_LookupPort() {
	port, err := net.LookupPort("tcp", "telnet")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("telnet port: %d", port)
}

func TestNet_ResolveTCPAddr() {
	domain := "www.baidu.com:8080"
	addr, err := net.ResolveTCPAddr("tcp", domain)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s IP: %+v PORT: %+v\n", domain, addr.IP, addr.Port)
}

func TestNet_DialTCP() {
	url := "www.baidu.com:80"

	addr, err := net.ResolveTCPAddr("tcp4", url)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	n, err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Printf("send: %d\n", n)

	buf, err := ioutil.ReadAll(conn)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Printf("recv: %d\n%s\n", len(buf), buf)
}

func TestNet_ListenTCP() {
	addr, err := net.ResolveTCPAddr("tcp4", ":7070")
	if err != nil {
		log.Fatal(err)
	}

	listen, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		go func (conn *net.TCPConn) {
			defer conn.Close()
			now := time.Now()
			conn.Write([]byte(now.String() + "\n"))
		}(conn)
	}
}

func TestNet_DialTimeout() {
	// url := "baidu.com:80"
	url := "facebook.com:80"
	conn, err := net.DialTimeout("tcp", url, time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	n, err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("write: %d\n", n)

	buf := make([]byte, 2048)
	n, err = conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", buf)
}

func TestNet_ListenUDP() {
	addr, err := net.ResolveUDPAddr("udp", ":7070")
	if err != nil {
		log.Fatal(err)
	}

	listen, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	for {
		buf := make([]byte, 256)
		n, addr, err := listen.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v:%+v connected read %d\n", addr.IP,
			addr.Port, n)

		n, err = listen.WriteToUDP(buf, addr)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestNet_DialUDP() {
	for i := 0; i < 10; i++ {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:7070")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	n, err := conn.Write([]byte("你好阿！"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("send: %d\n", n)

	buf := make([]byte, 256)
	n, _, err = conn.ReadFromUDP(buf)
	if err != nil {
		log.Fatal(err)
	}
		fmt.Printf("recv: %s\n", buf)
	}
}

func main() {
	// TestNet_ParseIp()
	// TestNet_Mask()
	// TestNet_LookupHost()
	// TestNet_LookupPort()
	// TestNet_ResolveTCPAddr()
	// TestNet_DialTCP()
	// TestNet_ListenTCP()
	// TestNet_DialTimeout()

	go TestNet_DialUDP()
	TestNet_ListenUDP()
}
