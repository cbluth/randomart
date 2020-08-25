package main

import (
	"fmt"
	"log"
	"encoding/binary"
	"math/rand"
)

var (
	box string = buildFrame()
	charmap = []string{
		" ", ".", "o",
		"+", "=", "*",
		"B", "O", "X",
		"@", "%", "&",
		"#", "/", "^",
		"S", "E", "~",
		"$", ":", "-",
		"8",
	}
)

const (
	rows int = 11
	cols int = 19
	hoz string = "-"
	vert string = "|"
	corner string = "+"
	undStart string = "\033[4m"
	undStop  string = "\033[0m"
)

func randomArt(b []byte, banner string) string {
	return printBox(b[:])
}

func buildFrame() string {
	s := ""
	for i := 1 ; i < rows ; i++ {
		if i == 1 {
			s += corner+"%s"+corner+"\n"
		}
		if i == rows-1 {
			s += corner+"%s"+corner
		} else {
			s += vert+"%s"+vert+"\n"
		}		
	}
	return s
}

func printBox(b []byte) string {
	rs := makeRows(b)
	// fillRows(b)
	return fmt.Sprintf(box,
		rs[0],
		rs[1],
		rs[2],
		rs[3],
		rs[4],
		rs[5],
		rs[6],
		rs[7],
		rs[8],
		rs[9],
		rs[10],
	)
}

func showByte(b byte) bool {
	c := false
	rand.Seed(int64(int(b)))
	r := rand.Intn(4)
	if r == 0 {
		c = true
	}
	return c
}

func getChar(b byte) string {
	rand.Seed(int64(int(b)))
	return charmap[rand.Intn(len(charmap))]	
}

func makeRows(b []byte) []string {
	rs := []string{}
	fr := fillRows(b) 
	for i := 1 ; i < rows ; i++ {
		if i == 1 {
			rs = append(rs, "------[TOP]------")
		}
		if i == rows-1 {
			rs = append(rs, "----[BOTTOM]-----")
		} else {
			rs = append(rs, fr[i-1])
		}		
	}
	return rs
}

func bytesToInt64(bytes []byte) int64 {
	i := binary.BigEndian.Uint64(bytes)
	return int64(i)
}

func getBytes(b []byte) [153]byte {
	src := rand.New(rand.NewSource(bytesToInt64(b)))
	rb := [153]byte{}
	_, err := src.Read(rb[:])
	if err != nil {
		log.Fatalln(err)
	}
	return rb
}

func fillRows(bytes []byte) []string {
	rb := getBytes(bytes)
	s := ""
	for _, b := range rb {
		if showByte(b) {
			s += getChar(b)
		} else {
			s += " "
		}
	}
	ss := []string{}
	for i := 0 ; i < rows-2 ; i++ {
		ss = append(ss, s[(i) * (cols-2):(i+1) * (cols-2)])
	}
	return ss
}















// +---[RSA 3072]----+
// |           =*o.+ |
// |       .   .*+= .|
// |      = + ...=.+ |
// |     + = +.*. =  |
// |      + S.oo=+ o |
// |     . o.+o+* o o|
// |      . ..+  B  .|
// |       .    E    |
// |             .   |
// +----[SHA256]-----+


// +---[RSA 2048]----+
// |       o=.       |
// |    o  o++E      |
// |   + . Ooo.      |
// |    + O B..      |
// |     = *S.       |
// |      o          |
// |                 |
// |                 |
// |                 |
// +--------|--------+
