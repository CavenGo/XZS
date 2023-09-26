package config

import "github.com/panjf2000/ants/v2"

var GlobalPool *ants.Pool

func InitPool() (err error) {
	GlobalPool, err = ants.NewPool(10000)
	return err
}
