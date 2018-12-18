package astro

type Formatter interface {
	Format(entry Event) []byte
}
