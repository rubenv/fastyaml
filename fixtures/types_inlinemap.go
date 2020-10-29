package fixtures

type InlineMap struct {
	UpdateCount int             `yaml:"UpdateCount"`
	Options     CarSetupOptions `yaml:",inline"`
}

type CarSetupOptions map[string]map[string]map[string]string
