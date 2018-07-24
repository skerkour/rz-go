package astroflow

type Formatter interface {
	Format(entry Event) []byte
}
