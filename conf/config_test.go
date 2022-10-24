package conf_test

import (
	"os"
	"testing"

	"github.com/lxygwqf9527/demo-api/conf"
	"github.com/stretchr/testify/assert"
)

func TestCoonfigFRomToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		should.Equal("demo", conf.C().App.Name)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	should := assert.New(t)
	os.Setenv("MYSQL_DATABASE", "unit_test")

	err := conf.LoadConfigFromEnv()
	if should.NoError(err) {
		should.Equal("unit_test", conf.C().MySQL.Database)
	}
}
