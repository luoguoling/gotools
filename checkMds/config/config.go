package config
import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)
type Config struct {
	WebHook	[]string	`yaml:"webhook,flow"`
	ProjectName	string	`yaml:"projectname"`
	DbUrl	string	`yaml:"dburl"`
	DbPort	int	`yaml:"dbport"`
	DbUser	string	`yaml:"dbuser"`
	DbPwd	string	`yaml:"dbpwd"`
	PlatformAlias	[]string	`yaml:"platformAlias,flow"`
	DbName	string	`yaml:"dbname"`
	Interval     int      `yaml:"interval"`
}
var conf Config
func InitConfig(configPath string) error {
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return errors.Wrap(err, "Read config file failed")
	}

	if err = yaml.Unmarshal(configFile, &conf); err != nil {
		return errors.Wrap(err, "Unmarshal config file failed.")
	}
	return nil
}

func GetConfig() Config {
	return conf
}