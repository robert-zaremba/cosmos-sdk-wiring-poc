package main

type ProdKind int

const (
	ProdKindProd ProdKind = iota
	ProdKindTest
)

type ModuleA struct {
	Key ModuleKey
	Cdc Codec
	// store

	kind ProdKind
}

func NewModuleATest(k ModuleKey, cdc Codec) ModuleA {
	return ModuleA{k, cdc, ProdKindTest}
}

func NewModuleAProd(k ModuleKey, cdc Codec) ModuleA {
	return ModuleA{k, cdc, ProdKindProd}
}

func (m ModuleA) RegisterKeeper() KeeperA {
	if m.kind == ProdKindProd {
		return keeperA{m}
	}
	return mockKeeperA{m}
}

// KEEPER A - production

type KeeperA interface {
	GetAccount(id string) string
}

type keeperA struct {
	M ModuleA
}

func (k keeperA) GetAccount(id string) string {
	return "Account " + id
}

// KEEPER A - mock implementation

type mockKeeperA struct {
	M ModuleA
}

func (k mockKeeperA) GetAccount(id string) string {
	return "Mock Account " + id
}
