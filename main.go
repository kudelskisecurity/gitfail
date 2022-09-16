package main

import (
	b64 "encoding/base64"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

const (
	CONN_HOST  = "localhost"
	CONN_PORT  = "1337"
	CONN_TYPE  = "tcp"
	DNS_RECORD = "The flag is not here, you should enumerate (passively) more, try harder!"
)

func chunkSlice(slice string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func shuffle(vals []string) string {

	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, len(vals))
	n := len(vals)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(vals))
		ret[i] = vals[randIndex]
	}
	return ret[0]
}

func main() {

	sEnc := b64.StdEncoding.EncodeToString([]byte(DNS_RECORD))
	sChunked := chunkSlice(sEnc, 12)

	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn, sChunked)
	}
}

func handleRequest(conn net.Conn, sChunked []string) {
	fmt.Println(time.Now(), conn.RemoteAddr())
	conn.Write([]byte(shuffle(sChunked)))
	conn.Close()
}