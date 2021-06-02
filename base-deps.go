package main

import "fmt"

type Codec interface {
	Marshall(interface{}) string
}

type ModuleKey string

func NewModuleKey(key string) ModuleKey {
	return ModuleKey(key)
}

type simpleCodec struct{}

func (c simpleCodec) Marshall(x interface{}) string {
	return fmt.Sprint(x)
}

func NewSimpleCodec() Codec {
	return simpleCodec{}
}
