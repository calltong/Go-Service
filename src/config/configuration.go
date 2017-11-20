package config

import (
    "encoding/json"
    "os"
)

type FolderMain struct {
  Product string `json:"product" bson:"product"`
  Page string `json:"page" bson:"page"`
  Order string `json:"order" bson:"order"`
}

type LazadaMain struct {
  Url string `json:"url" bson:"url"`
}

type GeneralData struct {
  DbName string `json:"db_name" bson:"db_name"`
  Root string `json:"root_folder" bson:"root_folder"`
  Address string `json:"address" bson:"address"`
}

type Mapping struct {
  Name string `json:"name" bson:"name"`
  General GeneralData `json:"general" bson:"general"`
}

type ConfigData struct {
  Folder FolderMain `json:"folder" bson:"folder"`
  Lazada LazadaMain `json:"lazada" bson:"lazada"`
  Mappings []Mapping `json:"mapping" bson:"mapping"`
}

type Login struct {
	Username string `json:"username"`
  Password string `json:"password"`
}

type DBConfig struct {
	IpAddress string `json:"ip"`
  Port int `json:"port"`
  DbName string `json:"db_name"`
	Admin Login `json:"admin"`
	User Login `json:"user"`
}

type Configuration struct {
  IpAddress string `json:"ip"`
  Port int `json:"port"`
  AcceptList []string `json:"accept_list"`
	Database DBConfig `json:"database"`
  Data ConfigData `json:"data" bson:"data"`
}

var configuration Configuration

func LoadConfiguration(path string) (Configuration, error) {
	file, err := os.Open(path)
  if err == nil {
    decoder := json.NewDecoder(file)
  	err = decoder.Decode(&configuration)
  }

  return configuration, err
}

func GetConfiguration() Configuration {
	return configuration
}

func SetConfiguration(data ConfigData) {
  configuration.Data = data
}

func GetDatabaseName() string {
	return configuration.Database.DbName
}
