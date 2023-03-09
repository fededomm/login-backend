package configuration

type RouterConfig struct {
	Router string `mapstructure:"router"`
}

type Param struct {
	Realm        string `yaml:"realm" mapstructure:"realm" json:"realm"`
	KeycloakUrl  string `yaml:"keycloak-url" mapstructure:"keycloak-url" json:"keycloak-url"`
	TokenUrl     string `yaml:"token-url" mapstructure:"token-url" json:"token-url"`
	ClientID     string `yaml:"client-id" mapstructure:"client-id" json:"client-id"`
	ClientSecret string `yaml:"client-secret" mapstructure:"client-secret" json:"client-secret"`
}
