package configuration

type RouterConfig struct {
	Router string `mapstructure:"router"`
}
type Authentication struct {
	Param Param `mapstructure:"param"`
}
type Param struct {
	TokenUrl     string `yaml:"token-url" mapstructure:"token-url" json:"token-url"`
	ClientID     string `yaml:"client-id" mapstructure:"client-id" json:"client-id"`
	ClientSecret string `yaml:"client-secret" mapstructure:"client-secret" json:"client-secret"`
}
