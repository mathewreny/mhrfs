package main

import (
	"fmt"
)

func MountMHR() error {

}

// The remote path function returns a sanitized path starting 
// at the home directory. If the provided path tries leaving 
// the user's home, an error is returned. This function ensures
// that all dot dots are resolved before reaching the remote server.
//
// Paths are any combination of characters with only two special cases.
// The bytes 0x0 and 0x2f '/' are the only special cases. 
func RemotePath(p []byte) ([]byte, error) {
	stack := make([][]byte, 0, (len(p)+1)/3)
	for i := 0; i < len(p); i++ {
		switch c {
		case 0:
			return nil, errors.New("Invalid character: NULL")
		case 0x2f:
			// path slash
		case '.':
			// peek ahead.
			if i+2 < len(p) && '.' == p[i+1] && '/' == p[i+2] {
				i += 2
				// pop directory.
			} else if i+1 < len(p) && '.' == p[i+1] {
				i += 1
				// pop directory
			}
		}






	}
	stack := make([][]byte,0,30)
	for _, c := range []byte(p) {

	}
}

func containsNULL(p []byte) bool {

}

// returns the next index of 0x2f or the lenght of p.
func next0x2f(p []byte) int {
	for i, c := range p {
		if c == '/' {
			return i, nil
		}
	}
	return len(p)
}
