package base

import "fmt"

type ProxyConfig struct {
	Port            string `xml:"port,attr"`
	PrintConsoleLog bool   `xml:"printConsoleLog,attr"`
	PrintSummary    bool   `xml:"printSummary,attr"`

	ReverseProxys *ReverseProxys `xml:"ReverseProxys"`

	AllowIpStr string

	BuffSize int

	AllowMaxConn int32
	AllowMaxWait int32

	Timeout int
}

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