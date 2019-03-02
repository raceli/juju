// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/provider/lxd (interfaces: Server,ServerFactory,InterfaceAddress)

// Package lxd is a generated GoMock package.
package lxd

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	lxd "github.com/juju/juju/container/lxd"
	environs "github.com/juju/juju/environs"
	network "github.com/juju/juju/network"
	client "github.com/lxc/lxd/client"
	api "github.com/lxc/lxd/shared/api"
)

// MockServer is a mock of Server interface
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// AliveContainers mocks base method
func (m *MockServer) AliveContainers(arg0 string) ([]lxd.Container, error) {
	ret := m.ctrl.Call(m, "AliveContainers", arg0)
	ret0, _ := ret[0].([]lxd.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AliveContainers indicates an expected call of AliveContainers
func (mr *MockServerMockRecorder) AliveContainers(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AliveContainers", reflect.TypeOf((*MockServer)(nil).AliveContainers), arg0)
}

// ContainerAddresses mocks base method
func (m *MockServer) ContainerAddresses(arg0 string) ([]network.Address, error) {
	ret := m.ctrl.Call(m, "ContainerAddresses", arg0)
	ret0, _ := ret[0].([]network.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerAddresses indicates an expected call of ContainerAddresses
func (mr *MockServerMockRecorder) ContainerAddresses(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerAddresses", reflect.TypeOf((*MockServer)(nil).ContainerAddresses), arg0)
}

// CreateCertificate mocks base method
func (m *MockServer) CreateCertificate(arg0 api.CertificatesPost) error {
	ret := m.ctrl.Call(m, "CreateCertificate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCertificate indicates an expected call of CreateCertificate
func (mr *MockServerMockRecorder) CreateCertificate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCertificate", reflect.TypeOf((*MockServer)(nil).CreateCertificate), arg0)
}

// CreateClientCertificate mocks base method
func (m *MockServer) CreateClientCertificate(arg0 *lxd.Certificate) error {
	ret := m.ctrl.Call(m, "CreateClientCertificate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateClientCertificate indicates an expected call of CreateClientCertificate
func (mr *MockServerMockRecorder) CreateClientCertificate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClientCertificate", reflect.TypeOf((*MockServer)(nil).CreateClientCertificate), arg0)
}

// CreateContainerFromSpec mocks base method
func (m *MockServer) CreateContainerFromSpec(arg0 lxd.ContainerSpec) (*lxd.Container, error) {
	ret := m.ctrl.Call(m, "CreateContainerFromSpec", arg0)
	ret0, _ := ret[0].(*lxd.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContainerFromSpec indicates an expected call of CreateContainerFromSpec
func (mr *MockServerMockRecorder) CreateContainerFromSpec(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainerFromSpec", reflect.TypeOf((*MockServer)(nil).CreateContainerFromSpec), arg0)
}

// CreatePool mocks base method
func (m *MockServer) CreatePool(arg0, arg1 string, arg2 map[string]string) error {
	ret := m.ctrl.Call(m, "CreatePool", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePool indicates an expected call of CreatePool
func (mr *MockServerMockRecorder) CreatePool(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePool", reflect.TypeOf((*MockServer)(nil).CreatePool), arg0, arg1, arg2)
}

// CreateProfile mocks base method
func (m *MockServer) CreateProfile(arg0 api.ProfilesPost) error {
	ret := m.ctrl.Call(m, "CreateProfile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProfile indicates an expected call of CreateProfile
func (mr *MockServerMockRecorder) CreateProfile(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfile", reflect.TypeOf((*MockServer)(nil).CreateProfile), arg0)
}

// CreateProfileWithConfig mocks base method
func (m *MockServer) CreateProfileWithConfig(arg0 string, arg1 map[string]string) error {
	ret := m.ctrl.Call(m, "CreateProfileWithConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProfileWithConfig indicates an expected call of CreateProfileWithConfig
func (mr *MockServerMockRecorder) CreateProfileWithConfig(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfileWithConfig", reflect.TypeOf((*MockServer)(nil).CreateProfileWithConfig), arg0, arg1)
}

// CreateVolume mocks base method
func (m *MockServer) CreateVolume(arg0, arg1 string, arg2 map[string]string) error {
	ret := m.ctrl.Call(m, "CreateVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVolume indicates an expected call of CreateVolume
func (mr *MockServerMockRecorder) CreateVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockServer)(nil).CreateVolume), arg0, arg1, arg2)
}

// DeleteCertificate mocks base method
func (m *MockServer) DeleteCertificate(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteCertificate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCertificate indicates an expected call of DeleteCertificate
func (mr *MockServerMockRecorder) DeleteCertificate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCertificate", reflect.TypeOf((*MockServer)(nil).DeleteCertificate), arg0)
}

// DeleteProfile mocks base method
func (m *MockServer) DeleteProfile(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteProfile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProfile indicates an expected call of DeleteProfile
func (mr *MockServerMockRecorder) DeleteProfile(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProfile", reflect.TypeOf((*MockServer)(nil).DeleteProfile), arg0)
}

// DeleteStoragePoolVolume mocks base method
func (m *MockServer) DeleteStoragePoolVolume(arg0, arg1, arg2 string) error {
	ret := m.ctrl.Call(m, "DeleteStoragePoolVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStoragePoolVolume indicates an expected call of DeleteStoragePoolVolume
func (mr *MockServerMockRecorder) DeleteStoragePoolVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStoragePoolVolume", reflect.TypeOf((*MockServer)(nil).DeleteStoragePoolVolume), arg0, arg1, arg2)
}

// EnableHTTPSListener mocks base method
func (m *MockServer) EnableHTTPSListener() error {
	ret := m.ctrl.Call(m, "EnableHTTPSListener")
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableHTTPSListener indicates an expected call of EnableHTTPSListener
func (mr *MockServerMockRecorder) EnableHTTPSListener() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableHTTPSListener", reflect.TypeOf((*MockServer)(nil).EnableHTTPSListener))
}

// EnsureDefaultStorage mocks base method
func (m *MockServer) EnsureDefaultStorage(arg0 *api.Profile, arg1 string) error {
	ret := m.ctrl.Call(m, "EnsureDefaultStorage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureDefaultStorage indicates an expected call of EnsureDefaultStorage
func (mr *MockServerMockRecorder) EnsureDefaultStorage(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureDefaultStorage", reflect.TypeOf((*MockServer)(nil).EnsureDefaultStorage), arg0, arg1)
}

// FilterContainers mocks base method
func (m *MockServer) FilterContainers(arg0 string, arg1 ...string) ([]lxd.Container, error) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FilterContainers", varargs...)
	ret0, _ := ret[0].([]lxd.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterContainers indicates an expected call of FilterContainers
func (mr *MockServerMockRecorder) FilterContainers(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterContainers", reflect.TypeOf((*MockServer)(nil).FilterContainers), varargs...)
}

// FindImage mocks base method
func (m *MockServer) FindImage(arg0, arg1 string, arg2 []lxd.ServerSpec, arg3 bool, arg4 environs.StatusCallbackFunc) (lxd.SourcedImage, error) {
	ret := m.ctrl.Call(m, "FindImage", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(lxd.SourcedImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindImage indicates an expected call of FindImage
func (mr *MockServerMockRecorder) FindImage(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindImage", reflect.TypeOf((*MockServer)(nil).FindImage), arg0, arg1, arg2, arg3, arg4)
}

// GetCertificate mocks base method
func (m *MockServer) GetCertificate(arg0 string) (*api.Certificate, string, error) {
	ret := m.ctrl.Call(m, "GetCertificate", arg0)
	ret0, _ := ret[0].(*api.Certificate)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCertificate indicates an expected call of GetCertificate
func (mr *MockServerMockRecorder) GetCertificate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificate", reflect.TypeOf((*MockServer)(nil).GetCertificate), arg0)
}

// GetClusterMembers mocks base method
func (m *MockServer) GetClusterMembers() ([]api.ClusterMember, error) {
	ret := m.ctrl.Call(m, "GetClusterMembers")
	ret0, _ := ret[0].([]api.ClusterMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterMembers indicates an expected call of GetClusterMembers
func (mr *MockServerMockRecorder) GetClusterMembers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterMembers", reflect.TypeOf((*MockServer)(nil).GetClusterMembers))
}

// GetConnectionInfo mocks base method
func (m *MockServer) GetConnectionInfo() (*client.ConnectionInfo, error) {
	ret := m.ctrl.Call(m, "GetConnectionInfo")
	ret0, _ := ret[0].(*client.ConnectionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConnectionInfo indicates an expected call of GetConnectionInfo
func (mr *MockServerMockRecorder) GetConnectionInfo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnectionInfo", reflect.TypeOf((*MockServer)(nil).GetConnectionInfo))
}

// GetContainerProfiles mocks base method
func (m *MockServer) GetContainerProfiles(arg0 string) ([]string, error) {
	ret := m.ctrl.Call(m, "GetContainerProfiles", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainerProfiles indicates an expected call of GetContainerProfiles
func (mr *MockServerMockRecorder) GetContainerProfiles(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerProfiles", reflect.TypeOf((*MockServer)(nil).GetContainerProfiles), arg0)
}

// GetNICsFromProfile mocks base method
func (m *MockServer) GetNICsFromProfile(arg0 string) (map[string]map[string]string, error) {
	ret := m.ctrl.Call(m, "GetNICsFromProfile", arg0)
	ret0, _ := ret[0].(map[string]map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNICsFromProfile indicates an expected call of GetNICsFromProfile
func (mr *MockServerMockRecorder) GetNICsFromProfile(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNICsFromProfile", reflect.TypeOf((*MockServer)(nil).GetNICsFromProfile), arg0)
}

// GetProfile mocks base method
func (m *MockServer) GetProfile(arg0 string) (*api.Profile, string, error) {
	ret := m.ctrl.Call(m, "GetProfile", arg0)
	ret0, _ := ret[0].(*api.Profile)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProfile indicates an expected call of GetProfile
func (mr *MockServerMockRecorder) GetProfile(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockServer)(nil).GetProfile), arg0)
}

// GetServer mocks base method
func (m *MockServer) GetServer() (*api.Server, string, error) {
	ret := m.ctrl.Call(m, "GetServer")
	ret0, _ := ret[0].(*api.Server)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetServer indicates an expected call of GetServer
func (mr *MockServerMockRecorder) GetServer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServer", reflect.TypeOf((*MockServer)(nil).GetServer))
}

// GetStoragePool mocks base method
func (m *MockServer) GetStoragePool(arg0 string) (*api.StoragePool, string, error) {
	ret := m.ctrl.Call(m, "GetStoragePool", arg0)
	ret0, _ := ret[0].(*api.StoragePool)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStoragePool indicates an expected call of GetStoragePool
func (mr *MockServerMockRecorder) GetStoragePool(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePool", reflect.TypeOf((*MockServer)(nil).GetStoragePool), arg0)
}

// GetStoragePoolVolume mocks base method
func (m *MockServer) GetStoragePoolVolume(arg0, arg1, arg2 string) (*api.StorageVolume, string, error) {
	ret := m.ctrl.Call(m, "GetStoragePoolVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.StorageVolume)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStoragePoolVolume indicates an expected call of GetStoragePoolVolume
func (mr *MockServerMockRecorder) GetStoragePoolVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePoolVolume", reflect.TypeOf((*MockServer)(nil).GetStoragePoolVolume), arg0, arg1, arg2)
}

// GetStoragePoolVolumes mocks base method
func (m *MockServer) GetStoragePoolVolumes(arg0 string) ([]api.StorageVolume, error) {
	ret := m.ctrl.Call(m, "GetStoragePoolVolumes", arg0)
	ret0, _ := ret[0].([]api.StorageVolume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoragePoolVolumes indicates an expected call of GetStoragePoolVolumes
func (mr *MockServerMockRecorder) GetStoragePoolVolumes(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePoolVolumes", reflect.TypeOf((*MockServer)(nil).GetStoragePoolVolumes), arg0)
}

// GetStoragePools mocks base method
func (m *MockServer) GetStoragePools() ([]api.StoragePool, error) {
	ret := m.ctrl.Call(m, "GetStoragePools")
	ret0, _ := ret[0].([]api.StoragePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoragePools indicates an expected call of GetStoragePools
func (mr *MockServerMockRecorder) GetStoragePools() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePools", reflect.TypeOf((*MockServer)(nil).GetStoragePools))
}

// HasProfile mocks base method
func (m *MockServer) HasProfile(arg0 string) (bool, error) {
	ret := m.ctrl.Call(m, "HasProfile", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasProfile indicates an expected call of HasProfile
func (mr *MockServerMockRecorder) HasProfile(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasProfile", reflect.TypeOf((*MockServer)(nil).HasProfile), arg0)
}

// HostArch mocks base method
func (m *MockServer) HostArch() string {
	ret := m.ctrl.Call(m, "HostArch")
	ret0, _ := ret[0].(string)
	return ret0
}

// HostArch indicates an expected call of HostArch
func (mr *MockServerMockRecorder) HostArch() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostArch", reflect.TypeOf((*MockServer)(nil).HostArch))
}

// IsClustered mocks base method
func (m *MockServer) IsClustered() bool {
	ret := m.ctrl.Call(m, "IsClustered")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsClustered indicates an expected call of IsClustered
func (mr *MockServerMockRecorder) IsClustered() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClustered", reflect.TypeOf((*MockServer)(nil).IsClustered))
}

// LocalBridgeName mocks base method
func (m *MockServer) LocalBridgeName() string {
	ret := m.ctrl.Call(m, "LocalBridgeName")
	ret0, _ := ret[0].(string)
	return ret0
}

// LocalBridgeName indicates an expected call of LocalBridgeName
func (mr *MockServerMockRecorder) LocalBridgeName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalBridgeName", reflect.TypeOf((*MockServer)(nil).LocalBridgeName))
}

// Name mocks base method
func (m *MockServer) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockServerMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockServer)(nil).Name))
}

// RemoveContainer mocks base method
func (m *MockServer) RemoveContainer(arg0 string) error {
	ret := m.ctrl.Call(m, "RemoveContainer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContainer indicates an expected call of RemoveContainer
func (mr *MockServerMockRecorder) RemoveContainer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveContainer", reflect.TypeOf((*MockServer)(nil).RemoveContainer), arg0)
}

// RemoveContainers mocks base method
func (m *MockServer) RemoveContainers(arg0 []string) error {
	ret := m.ctrl.Call(m, "RemoveContainers", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContainers indicates an expected call of RemoveContainers
func (mr *MockServerMockRecorder) RemoveContainers(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveContainers", reflect.TypeOf((*MockServer)(nil).RemoveContainers), arg0)
}

// ReplaceOrAddContainerProfile mocks base method
func (m *MockServer) ReplaceOrAddContainerProfile(arg0, arg1, arg2 string) error {
	ret := m.ctrl.Call(m, "ReplaceOrAddContainerProfile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplaceOrAddContainerProfile indicates an expected call of ReplaceOrAddContainerProfile
func (mr *MockServerMockRecorder) ReplaceOrAddContainerProfile(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceOrAddContainerProfile", reflect.TypeOf((*MockServer)(nil).ReplaceOrAddContainerProfile), arg0, arg1, arg2)
}

// ServerCertificate mocks base method
func (m *MockServer) ServerCertificate() string {
	ret := m.ctrl.Call(m, "ServerCertificate")
	ret0, _ := ret[0].(string)
	return ret0
}

// ServerCertificate indicates an expected call of ServerCertificate
func (mr *MockServerMockRecorder) ServerCertificate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerCertificate", reflect.TypeOf((*MockServer)(nil).ServerCertificate))
}

// ServerVersion mocks base method
func (m *MockServer) ServerVersion() string {
	ret := m.ctrl.Call(m, "ServerVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// ServerVersion indicates an expected call of ServerVersion
func (mr *MockServerMockRecorder) ServerVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerVersion", reflect.TypeOf((*MockServer)(nil).ServerVersion))
}

// StorageSupported mocks base method
func (m *MockServer) StorageSupported() bool {
	ret := m.ctrl.Call(m, "StorageSupported")
	ret0, _ := ret[0].(bool)
	return ret0
}

// StorageSupported indicates an expected call of StorageSupported
func (mr *MockServerMockRecorder) StorageSupported() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageSupported", reflect.TypeOf((*MockServer)(nil).StorageSupported))
}

// UpdateContainerConfig mocks base method
func (m *MockServer) UpdateContainerConfig(arg0 string, arg1 map[string]string) error {
	ret := m.ctrl.Call(m, "UpdateContainerConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContainerConfig indicates an expected call of UpdateContainerConfig
func (mr *MockServerMockRecorder) UpdateContainerConfig(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContainerConfig", reflect.TypeOf((*MockServer)(nil).UpdateContainerConfig), arg0, arg1)
}

// UpdateServerConfig mocks base method
func (m *MockServer) UpdateServerConfig(arg0 map[string]string) error {
	ret := m.ctrl.Call(m, "UpdateServerConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateServerConfig indicates an expected call of UpdateServerConfig
func (mr *MockServerMockRecorder) UpdateServerConfig(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServerConfig", reflect.TypeOf((*MockServer)(nil).UpdateServerConfig), arg0)
}

// UpdateStoragePoolVolume mocks base method
func (m *MockServer) UpdateStoragePoolVolume(arg0, arg1, arg2 string, arg3 api.StorageVolumePut, arg4 string) error {
	ret := m.ctrl.Call(m, "UpdateStoragePoolVolume", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStoragePoolVolume indicates an expected call of UpdateStoragePoolVolume
func (mr *MockServerMockRecorder) UpdateStoragePoolVolume(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStoragePoolVolume", reflect.TypeOf((*MockServer)(nil).UpdateStoragePoolVolume), arg0, arg1, arg2, arg3, arg4)
}

// UseTargetServer mocks base method
func (m *MockServer) UseTargetServer(arg0 string) (*lxd.Server, error) {
	ret := m.ctrl.Call(m, "UseTargetServer", arg0)
	ret0, _ := ret[0].(*lxd.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UseTargetServer indicates an expected call of UseTargetServer
func (mr *MockServerMockRecorder) UseTargetServer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseTargetServer", reflect.TypeOf((*MockServer)(nil).UseTargetServer), arg0)
}

// VerifyNetworkDevice mocks base method
func (m *MockServer) VerifyNetworkDevice(arg0 *api.Profile, arg1 string) error {
	ret := m.ctrl.Call(m, "VerifyNetworkDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyNetworkDevice indicates an expected call of VerifyNetworkDevice
func (mr *MockServerMockRecorder) VerifyNetworkDevice(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyNetworkDevice", reflect.TypeOf((*MockServer)(nil).VerifyNetworkDevice), arg0, arg1)
}

// WriteContainer mocks base method
func (m *MockServer) WriteContainer(arg0 *lxd.Container) error {
	ret := m.ctrl.Call(m, "WriteContainer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteContainer indicates an expected call of WriteContainer
func (mr *MockServerMockRecorder) WriteContainer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteContainer", reflect.TypeOf((*MockServer)(nil).WriteContainer), arg0)
}

// MockServerFactory is a mock of ServerFactory interface
type MockServerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockServerFactoryMockRecorder
}

// MockServerFactoryMockRecorder is the mock recorder for MockServerFactory
type MockServerFactoryMockRecorder struct {
	mock *MockServerFactory
}

// NewMockServerFactory creates a new mock instance
func NewMockServerFactory(ctrl *gomock.Controller) *MockServerFactory {
	mock := &MockServerFactory{ctrl: ctrl}
	mock.recorder = &MockServerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerFactory) EXPECT() *MockServerFactoryMockRecorder {
	return m.recorder
}

// InsecureRemoteServer mocks base method
func (m *MockServerFactory) InsecureRemoteServer(arg0 environs.CloudSpec) (Server, error) {
	ret := m.ctrl.Call(m, "InsecureRemoteServer", arg0)
	ret0, _ := ret[0].(Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsecureRemoteServer indicates an expected call of InsecureRemoteServer
func (mr *MockServerFactoryMockRecorder) InsecureRemoteServer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsecureRemoteServer", reflect.TypeOf((*MockServerFactory)(nil).InsecureRemoteServer), arg0)
}

// LocalServer mocks base method
func (m *MockServerFactory) LocalServer() (Server, error) {
	ret := m.ctrl.Call(m, "LocalServer")
	ret0, _ := ret[0].(Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocalServer indicates an expected call of LocalServer
func (mr *MockServerFactoryMockRecorder) LocalServer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalServer", reflect.TypeOf((*MockServerFactory)(nil).LocalServer))
}

// LocalServerAddress mocks base method
func (m *MockServerFactory) LocalServerAddress() (string, error) {
	ret := m.ctrl.Call(m, "LocalServerAddress")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocalServerAddress indicates an expected call of LocalServerAddress
func (mr *MockServerFactoryMockRecorder) LocalServerAddress() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalServerAddress", reflect.TypeOf((*MockServerFactory)(nil).LocalServerAddress))
}

// RemoteServer mocks base method
func (m *MockServerFactory) RemoteServer(arg0 environs.CloudSpec) (Server, error) {
	ret := m.ctrl.Call(m, "RemoteServer", arg0)
	ret0, _ := ret[0].(Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteServer indicates an expected call of RemoteServer
func (mr *MockServerFactoryMockRecorder) RemoteServer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteServer", reflect.TypeOf((*MockServerFactory)(nil).RemoteServer), arg0)
}

// MockInterfaceAddress is a mock of InterfaceAddress interface
type MockInterfaceAddress struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceAddressMockRecorder
}

// MockInterfaceAddressMockRecorder is the mock recorder for MockInterfaceAddress
type MockInterfaceAddressMockRecorder struct {
	mock *MockInterfaceAddress
}

// NewMockInterfaceAddress creates a new mock instance
func NewMockInterfaceAddress(ctrl *gomock.Controller) *MockInterfaceAddress {
	mock := &MockInterfaceAddress{ctrl: ctrl}
	mock.recorder = &MockInterfaceAddressMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterfaceAddress) EXPECT() *MockInterfaceAddressMockRecorder {
	return m.recorder
}

// InterfaceAddress mocks base method
func (m *MockInterfaceAddress) InterfaceAddress(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "InterfaceAddress", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InterfaceAddress indicates an expected call of InterfaceAddress
func (mr *MockInterfaceAddressMockRecorder) InterfaceAddress(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InterfaceAddress", reflect.TypeOf((*MockInterfaceAddress)(nil).InterfaceAddress), arg0)
}
