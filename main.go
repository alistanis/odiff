package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

var (
	f1 string
	f2 string
)

func init() {
	flag.StringVar(&f1, "f1", "", "The first file")
	flag.StringVar(&f2, "f2", "", "The second file")

}

func main() {
	flag.Parse()
	if f1 == "" || f2 == "" {
		exitError(errors.New("Must provide -f1 and -f2 flags"))
	}

	d1, err := ioutil.ReadFile(f1)
	if err != nil {
		exitError(err)
	}

	d2, err := ioutil.ReadFile(f2)
	if err != nil {
		exitError(err)
	}



	slc1 := strings.Split(string(d1), "")
	slc2 := strings.Split(string(d2), "")

	sort.Strings(slc1)
	sort.Strings(slc2)

	s1 := strings.Join(slc1, "")
	s2 := strings.Join(slc2, "")

	if s1 == s2 {
		fmt.Println("The files contain the same data")
	} else {
		fmt.Println("The files do not contain the same data")
	}

}

func exitError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(-1)
}
