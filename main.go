package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)



func fromFiles() {
	inpLocal, err := os.Open("input_local.txt")
	if err != nil {
		panic(err)
	}
	inpServer, err := os.Open("input_server.txt")
	if err != nil {
		panic(err)
	}
	outputResult, err := os.OpenFile("output_result.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	outputMissing, err := os.OpenFile("output_missing.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	merge(inpLocal, inpServer, outputResult, outputMissing)
}

func generate(count, serverOffset int) {
	var filenames = []string{
		"input_local.txt",
		"input_server.txt",
		"output_missing.txt",
		"output_result.txt",
		"output_missing.txt.example",
		"output_result.txt.example",
	}
	for _, fname := range filenames {
		if err := os.Remove(fname); err != nil {
			if _, ok := (err).(*os.PathError); ! ok {
				panic(err)
			}
		}
	}

	rand.Seed(time.Now().Unix())
	local, err := os.OpenFile("input_local.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	server, err := os.OpenFile("input_server.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	missing, err := os.OpenFile("output_missing.txt.example", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	result, err := os.OpenFile("output_result.txt.example", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= count; i++ {
		_, _ = fmt.Fprintln(result, i)

		if i < serverOffset {
			_, _ = fmt.Fprintln(local, i)
		} else {
			_, _ = fmt.Fprintln(server, i)
			if i > 1 && rand.Intn(2) == 0 {
				_, _ = fmt.Fprintln(missing, i)
			} else {
				_, _ = fmt.Fprintln(local, i)
			}
		}
	}

	if err = local.Close(); err != nil {
		panic(err)
	}
	if err = server.Close(); err != nil {
		panic(err)
	}
	if err = missing.Close(); err != nil {
		panic(err)
	}
	if err = result.Close(); err != nil {
		panic(err)
	}
}

var genMode = flag.NewFlagSet("gen", flag.ExitOnError)
var genCount = genMode.Int("count", 10, "how many ids to generate")
var genOffset = genMode.Int("offset", 5, "how many ids are skipped by server")

var usage = `For https://t.me/orel_codes_chat by Mikhail Vitsen

Two input files: 
	input_local.txt - sequence of ids in local storage
	input_server.txt - sequence of ids from server (with some offset)
Two output files
	output_result.txt - merge sequence of local and server ids
	output_missing.txt - ids from server that are not presented in local storage

* all files are sequences of valid ints separated by new line (with new line in the end of file)

usage: orelcodesmerge <command> [<args>]

Possible commands are:
   run   Run solution
   gen   Generate test input and example output files
       -count    how many ids to generate
       -offset    how many ids are skipped by server`

func main() {
	if len(os.Args) == 1 {
		fmt.Println(usage)
		return
	}


	switch os.Args[1] {
	case "run":
		var start = time.Now()

		fromFiles()

		fmt.Printf("Done in %d ms (md5sum output* to check)\n", time.Since(start).Milliseconds())
	case "gen":
		_ = genMode.Parse(os.Args[2:])
		generate(*genCount, *genOffset)
		fmt.Printf("Generated % d ids (server starts from % d)\n", *genCount, *genOffset)
	default:
		fmt.Println(usage)
	}
}
