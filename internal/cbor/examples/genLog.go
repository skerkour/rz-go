package main

import (
	"compress/zlib"
	"flag"
	"io"
	"log"
	"os"
	"time"

	"github.com/bloom42/rz-go"
)

func writeLog(fname string, count int, useCompress bool) {
	opFile := os.Stdout
	if fname != "<stdout>" {
		fil, _ := os.Create(fname)
		opFile = fil
		defer func() {
			if err := fil.Close(); err != nil {
				log.Fatal(err)
			}
		}()
	}

	var f io.WriteCloser = opFile
	if useCompress {
		f = zlib.NewWriter(f)
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
		}()

	}

	log := rz.New(rz.Writer(f), rz.TimestampFunc(func() time.Time { return time.Now().Round(time.Second) }))
	for i := 0; i < count; i++ {
		log.Error("Some Message", func(e *rz.Event) {
			e.Int("Fault", 41650+i)
		})
	}
}

func main() {
	outFile := flag.String("out", "<stdout>", "Output File to which logs will be written to (WILL overwrite if already present).")
	numLogs := flag.Int("num", 10, "Number of log messages to generate.")
	doCompress := flag.Bool("compress", false, "Enable inline compressed writer")

	flag.Parse()

	writeLog(*outFile, *numLogs, *doCompress)
}
