package restapi

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

type Restapi interface {
	Send() (map[string]string, error)
}

func getData(r Restapi) (map[string]string, error) {
	return r.Send()
}

type RestData struct {
	uri    string
	method int
	data   map[string]string
	//result map[string]string
}

func (r *RestData) Send() (map[string]string, error) {
	tcpaddr, err := net.ResolveTCPAddr("tcp", "wdm.adstars.cn:80")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpaddr)
	checkError(err)

	conn.Write([]byte("GET /Test/Index/login/uid/489874 HTTP/1.1\r\n\r\n"))
	res, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(res))

	return nil, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal error -: %s", err.Error())
		os.Exit(1)
	}
}
