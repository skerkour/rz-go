package main

import (
	"github.com/astroflow/astro-go"
	"github.com/astroflow/astro-go/log"
)

func main() {
	log.Init(
		astro.SetFormatter(astro.NewConsoleFormatter()),
		astro.SetLevel(astro.InfoLevel),
	)
	sublogger := log.With("4343", "4343")

	sublogger.Info("lol")
	sublogger.Debug("lol")
	log.Warn("lol2")

	/*
		for i := 0; i < 10000; i++ {
			go func() {
				log.With("user", struct{ Name string }{}, "n", " \\ \"ee").Info("lol\"\\dede")
			}()
		}
	*/
}
