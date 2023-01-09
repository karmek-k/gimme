package main

import (
	"fmt"

	"github.com/karmek-k/gimme/pkg/config"
)

func main() {
	headerSelectorString := "#readme-yaml-support-for-the-go-language"
	homepageSelectorString := ".js-headerLogo"
	homepageSelectorAttribute := "href"

	config := config.Config{
		Output: config.OutputConfig{
			Format: "json",
			Path:   "test.json",
		},
		Sites: map[string]config.SiteConfig{
			// TODO add host names - for example https://pkg.go.dev/gopkg.in/yaml.v3
			"go-yaml": {
				Selectors: map[string]config.SelectorConfig{
					"header": {
						CSS: &headerSelectorString,
					},
					"homepage": {
						CSS:       &homepageSelectorString,
						Attribute: &homepageSelectorAttribute,
					},
				},
			},
		},
	}

	fmt.Println(config)
}
