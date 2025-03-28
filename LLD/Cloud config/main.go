package main

import "fmt"

type IprotoType interface {
	clone() IprotoType
}

type env_config struct {
	name       string
	logging    string
	databaeUrl string
}

func (ec *env_config) clone() IprotoType {
	return &env_config{
		name:       ec.name,
		logging:    ec.logging,
		databaeUrl: ec.databaeUrl,
	}
}

func main() {
	ec := env_config{
		name:       "dev",
		logging:    "info only",
		databaeUrl: "local instance",
	}
	ec_qa := ec.clone().(*env_config)
	ec_qa.name = "qa"
	fmt.Println(" dev env config ", ec)
	fmt.Println(" qa env config ", ec_qa)
}
