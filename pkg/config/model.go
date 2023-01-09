package config

type Config struct {
	Output OutputConfig
	Sites  map[string]SiteConfig
}

type SiteConfig struct {
	Name      string
	Selectors []SelectorConfig
}

type OutputConfig struct {
	Format string
	Path   string
}

type SelectorConfig struct {
	FieldName string
	Selector  string
}
