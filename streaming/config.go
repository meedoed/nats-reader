package streaming

type Config struct {
	Address  string `toml:"address"`
	Cluster  string `toml:"cluster"`
	ClientID string `toml:"client_id"`
}

func NewConfig() *Config {
	return &Config{
		Address:  ":4222",
		Cluster:  "whschool_cluster",
		ClientID: "reader",
	}
}
