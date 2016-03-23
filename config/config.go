// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type Config struct {
	Servicecheckbeat ServicecheckbeatConfig
}

type ServicecheckbeatConfig struct {
	Period string `yaml:"period"`
}
