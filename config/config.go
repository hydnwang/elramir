package config

var (
	Mode string
	Port string
)

func SetDefault() {
	Mode = "release"
	Port = "80"
}
