// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener-extension-provider-openstack/pkg/openstack/client (interfaces: Factory,FactoryFactory,Compute,DNS,Networking)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	openstack "github.com/gardener/gardener-extension-provider-openstack/pkg/openstack"
	client "github.com/gardener/gardener-extension-provider-openstack/pkg/openstack/client"
	gomock "github.com/golang/mock/gomock"
	floatingips "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/floatingips"
	servergroups "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/servergroups"
	servers "github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	floatingips0 "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/floatingips"
	groups "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
	rules "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
	networks "github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// Compute mocks base method.
func (m *MockFactory) Compute(arg0 ...client.Option) (client.Compute, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Compute", varargs...)
	ret0, _ := ret[0].(client.Compute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compute indicates an expected call of Compute.
func (mr *MockFactoryMockRecorder) Compute(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compute", reflect.TypeOf((*MockFactory)(nil).Compute), arg0...)
}

// DNS mocks base method.
func (m *MockFactory) DNS(arg0 ...client.Option) (client.DNS, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DNS", varargs...)
	ret0, _ := ret[0].(client.DNS)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DNS indicates an expected call of DNS.
func (mr *MockFactoryMockRecorder) DNS(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DNS", reflect.TypeOf((*MockFactory)(nil).DNS), arg0...)
}

// Networking mocks base method.
func (m *MockFactory) Networking(arg0 ...client.Option) (client.Networking, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Networking", varargs...)
	ret0, _ := ret[0].(client.Networking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Networking indicates an expected call of Networking.
func (mr *MockFactoryMockRecorder) Networking(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Networking", reflect.TypeOf((*MockFactory)(nil).Networking), arg0...)
}

// Storage mocks base method.
func (m *MockFactory) Storage(arg0 ...client.Option) (client.Storage, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Storage", varargs...)
	ret0, _ := ret[0].(client.Storage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Storage indicates an expected call of Storage.
func (mr *MockFactoryMockRecorder) Storage(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Storage", reflect.TypeOf((*MockFactory)(nil).Storage), arg0...)
}

// MockFactoryFactory is a mock of FactoryFactory interface.
type MockFactoryFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryFactoryMockRecorder
}

// MockFactoryFactoryMockRecorder is the mock recorder for MockFactoryFactory.
type MockFactoryFactoryMockRecorder struct {
	mock *MockFactoryFactory
}

// NewMockFactoryFactory creates a new mock instance.
func NewMockFactoryFactory(ctrl *gomock.Controller) *MockFactoryFactory {
	mock := &MockFactoryFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactoryFactory) EXPECT() *MockFactoryFactoryMockRecorder {
	return m.recorder
}

// NewFactory mocks base method.
func (m *MockFactoryFactory) NewFactory(arg0 *openstack.Credentials) (client.Factory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewFactory", arg0)
	ret0, _ := ret[0].(client.Factory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewFactory indicates an expected call of NewFactory.
func (mr *MockFactoryFactoryMockRecorder) NewFactory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFactory", reflect.TypeOf((*MockFactoryFactory)(nil).NewFactory), arg0)
}

// MockCompute is a mock of Compute interface.
type MockCompute struct {
	ctrl     *gomock.Controller
	recorder *MockComputeMockRecorder
}

// MockComputeMockRecorder is the mock recorder for MockCompute.
type MockComputeMockRecorder struct {
	mock *MockCompute
}

// NewMockCompute creates a new mock instance.
func NewMockCompute(ctrl *gomock.Controller) *MockCompute {
	mock := &MockCompute{ctrl: ctrl}
	mock.recorder = &MockComputeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompute) EXPECT() *MockComputeMockRecorder {
	return m.recorder
}

// AssociateFIPWithInstance mocks base method.
func (m *MockCompute) AssociateFIPWithInstance(arg0 string, arg1 floatingips.AssociateOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssociateFIPWithInstance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssociateFIPWithInstance indicates an expected call of AssociateFIPWithInstance.
func (mr *MockComputeMockRecorder) AssociateFIPWithInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssociateFIPWithInstance", reflect.TypeOf((*MockCompute)(nil).AssociateFIPWithInstance), arg0, arg1)
}

// CreateServer mocks base method.
func (m *MockCompute) CreateServer(arg0 servers.CreateOpts) (*servers.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServer", arg0)
	ret0, _ := ret[0].(*servers.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServer indicates an expected call of CreateServer.
func (mr *MockComputeMockRecorder) CreateServer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServer", reflect.TypeOf((*MockCompute)(nil).CreateServer), arg0)
}

// CreateServerGroup mocks base method.
func (m *MockCompute) CreateServerGroup(arg0, arg1 string) (*servergroups.ServerGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServerGroup", arg0, arg1)
	ret0, _ := ret[0].(*servergroups.ServerGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServerGroup indicates an expected call of CreateServerGroup.
func (mr *MockComputeMockRecorder) CreateServerGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServerGroup", reflect.TypeOf((*MockCompute)(nil).CreateServerGroup), arg0, arg1)
}

// DeleteServer mocks base method.
func (m *MockCompute) DeleteServer(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServer indicates an expected call of DeleteServer.
func (mr *MockComputeMockRecorder) DeleteServer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServer", reflect.TypeOf((*MockCompute)(nil).DeleteServer), arg0)
}

// DeleteServerGroup mocks base method.
func (m *MockCompute) DeleteServerGroup(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServerGroup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServerGroup indicates an expected call of DeleteServerGroup.
func (mr *MockComputeMockRecorder) DeleteServerGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServerGroup", reflect.TypeOf((*MockCompute)(nil).DeleteServerGroup), arg0)
}

// FindFloatingIDbyInstnaceID mocks base method.
func (m *MockCompute) FindFloatingIDbyInstnaceID(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFloatingIDbyInstnaceID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFloatingIDbyInstnaceID indicates an expected call of FindFloatingIDbyInstnaceID.
func (mr *MockComputeMockRecorder) FindFloatingIDbyInstnaceID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFloatingIDbyInstnaceID", reflect.TypeOf((*MockCompute)(nil).FindFloatingIDbyInstnaceID), arg0)
}

// FindServersByName mocks base method.
func (m *MockCompute) FindServersByName(arg0 string) ([]servers.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindServersByName", arg0)
	ret0, _ := ret[0].([]servers.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServersByName indicates an expected call of FindServersByName.
func (mr *MockComputeMockRecorder) FindServersByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServersByName", reflect.TypeOf((*MockCompute)(nil).FindServersByName), arg0)
}

// GetServerGroup mocks base method.
func (m *MockCompute) GetServerGroup(arg0 string) (*servergroups.ServerGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerGroup", arg0)
	ret0, _ := ret[0].(*servergroups.ServerGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerGroup indicates an expected call of GetServerGroup.
func (mr *MockComputeMockRecorder) GetServerGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerGroup", reflect.TypeOf((*MockCompute)(nil).GetServerGroup), arg0)
}

// ListServerGroups mocks base method.
func (m *MockCompute) ListServerGroups() ([]servergroups.ServerGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServerGroups")
	ret0, _ := ret[0].([]servergroups.ServerGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServerGroups indicates an expected call of ListServerGroups.
func (mr *MockComputeMockRecorder) ListServerGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServerGroups", reflect.TypeOf((*MockCompute)(nil).ListServerGroups))
}

// MockDNS is a mock of DNS interface.
type MockDNS struct {
	ctrl     *gomock.Controller
	recorder *MockDNSMockRecorder
}

// MockDNSMockRecorder is the mock recorder for MockDNS.
type MockDNSMockRecorder struct {
	mock *MockDNS
}

// NewMockDNS creates a new mock instance.
func NewMockDNS(ctrl *gomock.Controller) *MockDNS {
	mock := &MockDNS{ctrl: ctrl}
	mock.recorder = &MockDNSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDNS) EXPECT() *MockDNSMockRecorder {
	return m.recorder
}

// CreateOrUpdateRecordSet mocks base method.
func (m *MockDNS) CreateOrUpdateRecordSet(arg0 context.Context, arg1, arg2, arg3 string, arg4 []string, arg5 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateRecordSet", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateRecordSet indicates an expected call of CreateOrUpdateRecordSet.
func (mr *MockDNSMockRecorder) CreateOrUpdateRecordSet(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateRecordSet", reflect.TypeOf((*MockDNS)(nil).CreateOrUpdateRecordSet), arg0, arg1, arg2, arg3, arg4, arg5)
}

// DeleteRecordSet mocks base method.
func (m *MockDNS) DeleteRecordSet(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecordSet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRecordSet indicates an expected call of DeleteRecordSet.
func (mr *MockDNSMockRecorder) DeleteRecordSet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecordSet", reflect.TypeOf((*MockDNS)(nil).DeleteRecordSet), arg0, arg1, arg2, arg3)
}

// GetZones mocks base method.
func (m *MockDNS) GetZones(arg0 context.Context) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetZones", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetZones indicates an expected call of GetZones.
func (mr *MockDNSMockRecorder) GetZones(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetZones", reflect.TypeOf((*MockDNS)(nil).GetZones), arg0)
}

// MockNetworking is a mock of Networking interface.
type MockNetworking struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkingMockRecorder
}

// MockNetworkingMockRecorder is the mock recorder for MockNetworking.
type MockNetworkingMockRecorder struct {
	mock *MockNetworking
}

// NewMockNetworking creates a new mock instance.
func NewMockNetworking(ctrl *gomock.Controller) *MockNetworking {
	mock := &MockNetworking{ctrl: ctrl}
	mock.recorder = &MockNetworkingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworking) EXPECT() *MockNetworkingMockRecorder {
	return m.recorder
}

// CreateFloatingIP mocks base method.
func (m *MockNetworking) CreateFloatingIP(arg0 floatingips0.CreateOpts) (*floatingips0.FloatingIP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFloatingIP", arg0)
	ret0, _ := ret[0].(*floatingips0.FloatingIP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFloatingIP indicates an expected call of CreateFloatingIP.
func (mr *MockNetworkingMockRecorder) CreateFloatingIP(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFloatingIP", reflect.TypeOf((*MockNetworking)(nil).CreateFloatingIP), arg0)
}

// CreateRule mocks base method.
func (m *MockNetworking) CreateRule(arg0 rules.CreateOpts) (*rules.SecGroupRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRule", arg0)
	ret0, _ := ret[0].(*rules.SecGroupRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRule indicates an expected call of CreateRule.
func (mr *MockNetworkingMockRecorder) CreateRule(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRule", reflect.TypeOf((*MockNetworking)(nil).CreateRule), arg0)
}

// CreateSecurityGroup mocks base method.
func (m *MockNetworking) CreateSecurityGroup(arg0 groups.CreateOpts) (*groups.SecGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecurityGroup", arg0)
	ret0, _ := ret[0].(*groups.SecGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecurityGroup indicates an expected call of CreateSecurityGroup.
func (mr *MockNetworkingMockRecorder) CreateSecurityGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecurityGroup", reflect.TypeOf((*MockNetworking)(nil).CreateSecurityGroup), arg0)
}

// DeleteFloatingIP mocks base method.
func (m *MockNetworking) DeleteFloatingIP(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFloatingIP", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFloatingIP indicates an expected call of DeleteFloatingIP.
func (mr *MockNetworkingMockRecorder) DeleteFloatingIP(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFloatingIP", reflect.TypeOf((*MockNetworking)(nil).DeleteFloatingIP), arg0)
}

// DeleteRule mocks base method.
func (m *MockNetworking) DeleteRule(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRule", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRule indicates an expected call of DeleteRule.
func (mr *MockNetworkingMockRecorder) DeleteRule(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRule", reflect.TypeOf((*MockNetworking)(nil).DeleteRule), arg0)
}

// DeleteSecurityGroup mocks base method.
func (m *MockNetworking) DeleteSecurityGroup(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecurityGroup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecurityGroup indicates an expected call of DeleteSecurityGroup.
func (mr *MockNetworkingMockRecorder) DeleteSecurityGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecurityGroup", reflect.TypeOf((*MockNetworking)(nil).DeleteSecurityGroup), arg0)
}

// GetExternalNetworkInfoByName mocks base method.
func (m *MockNetworking) GetExternalNetworkInfoByName(arg0 string) ([]networks.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalNetworkInfoByName", arg0)
	ret0, _ := ret[0].([]networks.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExternalNetworkInfoByName indicates an expected call of GetExternalNetworkInfoByName.
func (mr *MockNetworkingMockRecorder) GetExternalNetworkInfoByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalNetworkInfoByName", reflect.TypeOf((*MockNetworking)(nil).GetExternalNetworkInfoByName), arg0)
}

// GetExternalNetworkNames mocks base method.
func (m *MockNetworking) GetExternalNetworkNames(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalNetworkNames", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExternalNetworkNames indicates an expected call of GetExternalNetworkNames.
func (mr *MockNetworkingMockRecorder) GetExternalNetworkNames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalNetworkNames", reflect.TypeOf((*MockNetworking)(nil).GetExternalNetworkNames), arg0)
}

// GetFipbyName mocks base method.
func (m *MockNetworking) GetFipbyName(arg0 string) ([]floatingips0.FloatingIP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFipbyName", arg0)
	ret0, _ := ret[0].([]floatingips0.FloatingIP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFipbyName indicates an expected call of GetFipbyName.
func (mr *MockNetworkingMockRecorder) GetFipbyName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFipbyName", reflect.TypeOf((*MockNetworking)(nil).GetFipbyName), arg0)
}

// GetNetworkByName mocks base method.
func (m *MockNetworking) GetNetworkByName(arg0 string) ([]networks.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkByName", arg0)
	ret0, _ := ret[0].([]networks.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkByName indicates an expected call of GetNetworkByName.
func (mr *MockNetworkingMockRecorder) GetNetworkByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkByName", reflect.TypeOf((*MockNetworking)(nil).GetNetworkByName), arg0)
}

// GetRulebyName mocks base method.
func (m *MockNetworking) GetRulebyName(arg0 string) ([]rules.SecGroupRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRulebyName", arg0)
	ret0, _ := ret[0].([]rules.SecGroupRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRulebyName indicates an expected call of GetRulebyName.
func (mr *MockNetworkingMockRecorder) GetRulebyName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRulebyName", reflect.TypeOf((*MockNetworking)(nil).GetRulebyName), arg0)
}

// GetSecurityGroupbyName mocks base method.
func (m *MockNetworking) GetSecurityGroupbyName(arg0 string) ([]groups.SecGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecurityGroupbyName", arg0)
	ret0, _ := ret[0].([]groups.SecGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecurityGroupbyName indicates an expected call of GetSecurityGroupbyName.
func (mr *MockNetworkingMockRecorder) GetSecurityGroupbyName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecurityGroupbyName", reflect.TypeOf((*MockNetworking)(nil).GetSecurityGroupbyName), arg0)
}

// ListFip mocks base method.
func (m *MockNetworking) ListFip(arg0 floatingips0.ListOpts) ([]floatingips0.FloatingIP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFip", arg0)
	ret0, _ := ret[0].([]floatingips0.FloatingIP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFip indicates an expected call of ListFip.
func (mr *MockNetworkingMockRecorder) ListFip(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFip", reflect.TypeOf((*MockNetworking)(nil).ListFip), arg0)
}

// ListNetwork mocks base method.
func (m *MockNetworking) ListNetwork(arg0 networks.ListOpts) ([]networks.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNetwork", arg0)
	ret0, _ := ret[0].([]networks.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNetwork indicates an expected call of ListNetwork.
func (mr *MockNetworkingMockRecorder) ListNetwork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNetwork", reflect.TypeOf((*MockNetworking)(nil).ListNetwork), arg0)
}

// ListRules mocks base method.
func (m *MockNetworking) ListRules(arg0 rules.ListOpts) ([]rules.SecGroupRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRules", arg0)
	ret0, _ := ret[0].([]rules.SecGroupRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRules indicates an expected call of ListRules.
func (mr *MockNetworkingMockRecorder) ListRules(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRules", reflect.TypeOf((*MockNetworking)(nil).ListRules), arg0)
}

// ListSecurityGroup mocks base method.
func (m *MockNetworking) ListSecurityGroup(arg0 groups.ListOpts) ([]groups.SecGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecurityGroup", arg0)
	ret0, _ := ret[0].([]groups.SecGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecurityGroup indicates an expected call of ListSecurityGroup.
func (mr *MockNetworkingMockRecorder) ListSecurityGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecurityGroup", reflect.TypeOf((*MockNetworking)(nil).ListSecurityGroup), arg0)
}
