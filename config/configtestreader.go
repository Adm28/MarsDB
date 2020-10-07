package config

import (
  "gopkg.in/yaml.v2"
  "os"
  "io/ioutil"
)
const configFilePath string =  "config.yml"

type Config struct {
  ClusterConfig struct { Nodes []string `yaml:"Nodes"`
  } `yaml:ClusterConfig`
}


func InitializeFromConfig()(Config,error) {

  var config Config
  file,err := os.Open(configFilePath)
  defer file.Close()
  if(err!=nil) {
    return Config{},err
  }
  bytes,_:=ioutil.ReadAll(file)
  err = yaml.Unmarshal(bytes,&config)
  if err!= nil {
    return Config{},err
  }
  return config,nil;
}

func InitializeConfiguration() Config {
  config,err := InitializeFromConfig()
  if err!= nil {
    config.SetDefaultParameters()
  }
  return config
}

func (config *Config) SetDefaultParameters() {
}

/*
func main() {
  config := InitializeConfiguration()
  fmt.Println(config)
}
*/
