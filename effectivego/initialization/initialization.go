package initialization

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

const (
	none = iota
	a
	b
	c
)

func Print() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(MB)

	fmt.Println(a, b, c)

	size := ByteSize(1 << 10)

	fmt.Println(size)

	fmt.Println(YB, ByteSize(1e13))
	//1.00YB 9.09TB

	fmt.Println(Home, gopath)
}

var (
	Home   = os.Getenv("HOME")
	User   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

func init() {
	if User == "" {
		log.Fatal("$USER not set")
	}
	if Home == "" {
		Home = "/home/" + User
	}
	if gopath == "" {
		gopath = Home + "/go"
	}
	// gopath may be overridden by --gopath flag on command line.
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}
