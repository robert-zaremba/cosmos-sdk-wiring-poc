package main

import "errors"

type ModuleB struct {
	Key ModuleKey
	Cdc Codec

	KA KeeperA
}

func NewModuleB(k ModuleKey, cdc Codec, ka KeeperA) ModuleB {
	return ModuleB{k, cdc, ka}
}

func (m ModuleB) RegisterKeeper() KeeperB {
	return keeperB{m}
}

type KeeperB interface {
	Slash(id string) (string, error)
}

type keeperB struct {
	M ModuleB
}

func (k keeperB) Slash(id string) (string, error) {
	if id == "" {
		return "", errors.New("can't slash empty account")
	}
	a := k.M.KA.GetAccount(id)
	return a, nil
}
