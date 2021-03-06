package base

import (
	"fmt"
	"net"
)

type ProxyConfig struct {
	Port            string `xml:"port,attr"`
	Prototype       string `xml:"prototype,attr"`
	PrintConsoleLog bool   `xml:"printConsoleLog,attr"`
	PrintSummary    bool   `xml:"printSummary,attr"`

	ReverseProxys *ReverseProxys `xml:"ReverseProxys"`

	AllowIpStr string

	BuffSize int

	AllowMaxConn int32
	AllowMaxWait int32

	//Auto,Deadline
	TimeoutModel string
	Timeout      int

	LogLevel string

	//socket
	Socket SocketItem
}

type SocketItem struct {
	Socket_Auth bool   `xml:"auth,attr"`
	Socket_UID  string `xml:"uid,attr"`
	Socket_PWD  string `xml:"pwd,attr"`
}

//  <!--client Network control-->
//  <NetworkTrafficControl enable="true" maxBytes="65535">
//    <Client ip="10.21.30.159"  maxBytes="32768"/>
//	<Client ip="10.21.30.160-10.21.30.200" />
//  </NetworkTrafficControl>

type ReverseProxys struct {
	Servers []Server `xml:"Server"`
}

func (this ReverseProxys) String() string {
	return fmt.Sprintf("{Servers:%+v}", this.Servers)
}

type Server struct {
	Addr  string `xml:",chardata"`
	Score int    `xml:"score,attr"`

	//ignore field
	//0 unnormal 1 normal
	Status     int   `xml:"-"`
	HandleSum  int64 `xml:"-"`
	ErrorCount int   `xml:"-"`
}

type IpRange struct {
	Start net.IP
	End   net.IP
}
