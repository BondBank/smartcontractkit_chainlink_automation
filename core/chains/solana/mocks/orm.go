// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	db "github.com/smartcontractkit/chainlink-solana/pkg/solana/db"
	mock "github.com/stretchr/testify/mock"

	pg "github.com/smartcontractkit/chainlink/core/services/pg"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// Chain provides a mock function with given fields: _a0, _a1
func (_m *ORM) Chain(_a0 string, _a1 ...pg.QOpt) (db.Chain, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 db.Chain
	if rf, ok := ret.Get(0).(func(string, ...pg.QOpt) db.Chain); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(db.Chain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...pg.QOpt) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Chains provides a mock function with given fields: offset, limit, qopts
func (_m *ORM) Chains(offset int, limit int, qopts ...pg.QOpt) ([]db.Chain, int, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, offset, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []db.Chain
	if rf, ok := ret.Get(0).(func(int, int, ...pg.QOpt) []db.Chain); ok {
		r0 = rf(offset, limit, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Chain)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int, ...pg.QOpt) int); ok {
		r1 = rf(offset, limit, qopts...)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int, ...pg.QOpt) error); ok {
		r2 = rf(offset, limit, qopts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateChain provides a mock function with given fields: id, config, qopts
func (_m *ORM) CreateChain(id string, config db.ChainCfg, qopts ...pg.QOpt) (db.Chain, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, id, config)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 db.Chain
	if rf, ok := ret.Get(0).(func(string, db.ChainCfg, ...pg.QOpt) db.Chain); ok {
		r0 = rf(id, config, qopts...)
	} else {
		r0 = ret.Get(0).(db.Chain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, db.ChainCfg, ...pg.QOpt) error); ok {
		r1 = rf(id, config, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNode provides a mock function with given fields: _a0, _a1
func (_m *ORM) CreateNode(_a0 db.NewNode, _a1 ...pg.QOpt) (db.Node, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 db.Node
	if rf, ok := ret.Get(0).(func(db.NewNode, ...pg.QOpt) db.Node); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(db.Node)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(db.NewNode, ...pg.QOpt) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteChain provides a mock function with given fields: id, qopts
func (_m *ORM) DeleteChain(id string, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...pg.QOpt) error); ok {
		r0 = rf(id, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNode provides a mock function with given fields: _a0, _a1
func (_m *ORM) DeleteNode(_a0 int32, _a1 ...pg.QOpt) error {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32, ...pg.QOpt) error); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnabledChains provides a mock function with given fields: _a0
func (_m *ORM) EnabledChains(_a0 ...pg.QOpt) ([]db.Chain, error) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []db.Chain
	if rf, ok := ret.Get(0).(func(...pg.QOpt) []db.Chain); ok {
		r0 = rf(_a0...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Chain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...pg.QOpt) error); ok {
		r1 = rf(_a0...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Node provides a mock function with given fields: _a0, _a1
func (_m *ORM) Node(_a0 int32, _a1 ...pg.QOpt) (db.Node, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 db.Node
	if rf, ok := ret.Get(0).(func(int32, ...pg.QOpt) db.Node); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(db.Node)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32, ...pg.QOpt) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeNamed provides a mock function with given fields: _a0, _a1
func (_m *ORM) NodeNamed(_a0 string, _a1 ...pg.QOpt) (db.Node, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 db.Node
	if rf, ok := ret.Get(0).(func(string, ...pg.QOpt) db.Node); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(db.Node)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...pg.QOpt) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Nodes provides a mock function with given fields: offset, limit, qopts
func (_m *ORM) Nodes(offset int, limit int, qopts ...pg.QOpt) ([]db.Node, int, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, offset, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []db.Node
	if rf, ok := ret.Get(0).(func(int, int, ...pg.QOpt) []db.Node); ok {
		r0 = rf(offset, limit, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Node)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int, ...pg.QOpt) int); ok {
		r1 = rf(offset, limit, qopts...)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int, ...pg.QOpt) error); ok {
		r2 = rf(offset, limit, qopts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NodesForChain provides a mock function with given fields: chainID, offset, limit, qopts
func (_m *ORM) NodesForChain(chainID string, offset int, limit int, qopts ...pg.QOpt) ([]db.Node, int, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, chainID, offset, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []db.Node
	if rf, ok := ret.Get(0).(func(string, int, int, ...pg.QOpt) []db.Node); ok {
		r0 = rf(chainID, offset, limit, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Node)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string, int, int, ...pg.QOpt) int); ok {
		r1 = rf(chainID, offset, limit, qopts...)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, int, int, ...pg.QOpt) error); ok {
		r2 = rf(chainID, offset, limit, qopts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateChain provides a mock function with given fields: id, enabled, config, qopts
func (_m *ORM) UpdateChain(id string, enabled bool, config db.ChainCfg, qopts ...pg.QOpt) (db.Chain, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, id, enabled, config)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 db.Chain
	if rf, ok := ret.Get(0).(func(string, bool, db.ChainCfg, ...pg.QOpt) db.Chain); ok {
		r0 = rf(id, enabled, config, qopts...)
	} else {
		r0 = ret.Get(0).(db.Chain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool, db.ChainCfg, ...pg.QOpt) error); ok {
		r1 = rf(id, enabled, config, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
