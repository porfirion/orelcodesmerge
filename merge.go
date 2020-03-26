package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
)

func takeInt(reader *bufio.Reader) (value int, isEof bool) {
	bytes, _, err := reader.ReadLine()
	if err != nil {
		if err == io.EOF {
			return math.MaxInt32, true
		} else {
			panic(err)
		}

	}
	// assuming we always have correct numbers
	if n, err := strconv.Atoi(string(bytes)); err != nil {
		panic(err)
	} else {
		return n, false
	}
}

func merge(inpLocal, inpServer io.Reader, result, missing io.Writer) {
	rdLocal := bufio.NewReader(inpLocal)
	rdServer := bufio.NewReader(inpServer)

	var local, eofLocal = takeInt(rdLocal)
	var server, eofServer = takeInt(rdServer)

	for !eofLocal || !eofServer {
		for !eofLocal && local < server {
			_, _ = fmt.Fprintln(result, local)
			local, eofLocal = takeInt(rdLocal)
		}
		for !eofServer && server < local {
			_, _ = fmt.Fprintln(result, server)
			_, _ = fmt.Fprintln(missing, server)
			server, eofServer = takeInt(rdServer)
		}

		for !eofLocal && !eofServer && local == server {
			_, _ = fmt.Fprintln(result, local)
			local, eofLocal = takeInt(rdLocal)
			server, eofServer = takeInt(rdServer)
		}
	}
}
