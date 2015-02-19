package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"io"
	"testing"
)

func TestMashNext(t *testing.T) {
	email := "test@mojn.com"
	buf := bytes.NewBufferString(email)
	mash := Mash{
		bufio.NewScanner(buf),
		md5.New(),
	}

	hash, err := mash.Next()
	if err != nil {
		t.Error(err)
	}

	t.Logf("Hashed %s to %s", email, hash)

	if hash != "b15bd93d1cd88787c7b65502378b9600" {
		t.Error(errors.New("Computed hash did not match"))
	}
}

func BenchmarkMash(b *testing.B) {
	b.StopTimer()

	bufs := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		bufs[i] = make([]byte, 32)
		rand.Read(bufs[i])
	}

	b.Logf("Seeded %d lines", b.N)

	buf := bytes.NewBuffer(bytes.Join(bufs, []byte("\n")))
	mash := Mash{
		bufio.NewScanner(buf),
		md5.New(),
	}

	b.StartTimer()
	for {
		_, err := mash.Next()
		if err == io.EOF {
			return
		}
		if err != nil {
			b.Fatal(err)
		}
	}
}
