package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

func main() {
	mash := Mash{
		bufio.NewScanner(os.Stdin),
		md5.New(),
	}
	for {
		hash, err := mash.Next()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(os.Stdout, hash)
	}
}

// Mash is a bufio.Scanner and hash.Hash that scans for lines in os.Stdin, trims
// and hashes them. Use by exhausting the Mash.Next function.
type Mash struct {
	*bufio.Scanner
	hash.Hash
}

// Next reads a line from the scanner, trims and hashes it. If an error occurs
// it is returned alongside an emptry string. If succesful, the hexadecimal
// represented hash is returned.
func (m *Mash) Next() (string, error) {
	if !m.Scanner.Scan() {
		err := m.Scanner.Err()
		if err != nil {
			return "", err
		}
		return "", io.EOF
	}
	line := bytes.ToLower(bytes.TrimSpace(m.Scanner.Bytes()))
	m.Write(line)
	hash := fmt.Sprintf("%x", m.Hash.Sum(nil))
	m.Hash.Reset()
	return hash, nil
}
