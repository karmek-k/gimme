package config

type Config struct {
	Output OutputConfig
	Sites  map[string]SiteConfig
}

type SiteConfig struct {
	Selectors map[string]SelectorConfig
}

type OutputConfig struct {
	Format string
	Path   string
}

type SelectorConfig struct {
	CSS       *string
	Attribute *string
}
