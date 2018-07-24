/*

usage:


	level := astro.DebugLevel
	formatter := astro.ConsoleFormatter

	if os.Gentenv("GO_ENV") == "production" {
		level = astro.InfoLevel
		formatter = astro.JSONFormatter
	}

	log.Init( // or log.Config
		astro.Token(...),
		astro.With({"app" "api"}, {"host": host}),
		astro.DisplayLogLevel(level),
		astro.Formatter(formatter),
		astro.Sample(sampler),
		// DisableTimeSTamp
		// default astro.SendLogLevel(astro.InfoLevel),
	)

	....

	log.With("port", Port).Info("server stared")
```
*/
package astro
