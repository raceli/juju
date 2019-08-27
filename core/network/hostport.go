// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package network

import (
	"net"
	"sort"
	"strconv"
	"strings"

	"github.com/juju/collections/set"
	"github.com/juju/errors"
)

// HostPort describes methods on an object that
// represents a network connection endpoint.
type HostPort interface {
	Host() string
	Port() int
	AddressScope() Scope
	Address() Address
}

// HostPorts derives from a slice of HostPort
// and allows bulk operations on its members.
type HostPorts []HostPort

// FilterUnusable returns a copy of the receiver HostPorts after removing
// any addresses unlikely to be usable (ScopeMachineLocal or ScopeLinkLocal).
func (hps HostPorts) FilterUnusable() HostPorts {
	filtered := make(HostPorts, 0, len(hps))
	for _, addr := range hps {
		switch addr.AddressScope() {
		case ScopeMachineLocal, ScopeLinkLocal:
			continue
		}
		filtered = append(filtered, addr)
	}
	return filtered
}

func (hps HostPorts) Strings() []string {
	result := make([]string, len(hps))
	for i, addr := range hps {
		result[i] = DialAddress(addr)
	}
	return result
}

// Unique returns a copy of the receiver HostPorts with duplicate endpoints
// removed. Note that this only applies to dial addresses; spaces are ignored.
func (hps HostPorts) Unique() HostPorts {
	results := make([]HostPort, 0, len(hps))
	seen := set.NewStrings()

	for _, addr := range hps {
		da := DialAddress(addr)
		if seen.Contains(da) {
			continue
		}

		seen.Add(da)
		results = append(results, addr)
	}
	return results
}

// PrioritizeInternal orders the HostPorts by best match for use as an endpoint
// for juju internal communication and returns them in NetAddr form.
// If there are no suitable addresses then an empty slice is returned.
func (hps HostPorts) PrioritizeInternal(machineLocal bool) []string {
	indexes := prioritizedAddressIndexes(len(hps), func(i int) Address {
		return hps[i].Address()
	}, internalAddressMatcher(machineLocal))

	out := make([]string, 0, len(indexes))
	for _, index := range indexes {
		out = append(out, DialAddress(hps[index]))
	}
	return out
}

func DialAddress(a HostPort) string {
	return net.JoinHostPort(a.Host(), strconv.Itoa(a.Port()))
}

// TODO (manadart 2019-08-15): Finish deprecation of `Port` and use that name.
type NetPort int

func (p NetPort) Port() int {
	return int(p)
}

// MachineHostPort associates a space-unaware address with a port.
type MachineHostPort struct {
	MachineAddress
	NetPort
}

var _ HostPort = MachineHostPort{}

// String implements Stringer.
func (hp MachineHostPort) String() string {
	return DialAddress(hp)
}

// GoString implements fmt.GoStringer.
func (hp MachineHostPort) GoString() string {
	return hp.String()
}

type MachineHostPorts []MachineHostPort

func (hp MachineHostPorts) HostPorts() HostPorts {
	addrs := make(HostPorts, len(hp))
	for i, hp := range hp {
		addrs[i] = hp
	}
	return addrs
}

// NewMachineHostPorts creates a list of MachineHostPorts
// from each given string address and port.
func NewMachineHostPorts(port int, addresses ...string) MachineHostPorts {
	hps := make([]MachineHostPort, len(addresses))
	for i, addr := range addresses {
		hps[i] = MachineHostPort{
			MachineAddress: NewMachineAddress(addr),
			NetPort:        NetPort(port),
		}
	}
	return hps
}

// ParseMachineHostPort converts a string containing a
// single host and port value to a MachineHostPort.
func ParseMachineHostPort(hp string) (*MachineHostPort, error) {
	host, port, err := net.SplitHostPort(hp)
	if err != nil {
		return nil, errors.Annotatef(err, "cannot parse %q as address:port", hp)
	}
	numPort, err := strconv.Atoi(port)
	if err != nil {
		return nil, errors.Annotatef(err, "cannot parse %q port", hp)
	}
	return &MachineHostPort{
		MachineAddress: NewMachineAddress(host),
		NetPort:        NetPort(numPort),
	}, nil
}

// CollapseToHostPorts returns the input nested slice of MachineHostPort
// as a flat slice of HostPort, preserving the order.
func CollapseToHostPorts(serversHostPorts []MachineHostPorts) HostPorts {
	var collapsed HostPorts
	for _, hps := range serversHostPorts {
		for _, hp := range hps {
			collapsed = append(collapsed, hp)
		}
	}
	return collapsed
}

// ProviderHostPort associates a provider/space aware address with a port.
type ProviderHostPort struct {
	ProviderAddress
	NetPort
}

var _ HostPort = ProviderHostPort{}

// String implements Stringer.
func (hp ProviderHostPort) String() string {
	return DialAddress(hp)
}

// GoString implements fmt.GoStringer.
func (hp ProviderHostPort) GoString() string {
	return hp.String()
}

type ProviderHostPorts []ProviderHostPort

func (hp ProviderHostPorts) Addresses() []ProviderAddress {
	addrs := make([]ProviderAddress, len(hp))
	for i, hp := range hp {
		addrs[i] = hp.ProviderAddress
	}
	return addrs
}

func (hp ProviderHostPorts) HostPorts() HostPorts {
	addrs := make(HostPorts, len(hp))
	for i, hp := range hp {
		addrs[i] = hp
	}
	return addrs
}

// ParseMachineHostPorts creates a slice of MachineHostPorts parsing each given
// string containing address:port.
// An error is returned if any string cannot be parsed as a MachineHostPort.
func ParseProviderHostPorts(hostPorts ...string) (ProviderHostPorts, error) {
	hps := make(ProviderHostPorts, len(hostPorts))
	for i, hp := range hostPorts {
		mhp, err := ParseMachineHostPort(hp)
		if err != nil {
			return nil, errors.Trace(err)
		}
		hps[i] = ProviderHostPort{
			ProviderAddress: ProviderAddress{MachineAddress: mhp.MachineAddress},
			NetPort:         mhp.NetPort,
		}
	}
	return hps, nil
}

// SpaceHostPort associates an address with a port.
type SpaceHostPort struct {
	SpaceAddress
	NetPort
}

var _ HostPort = SpaceHostPort{}

// String implements Stringer.
func (hp SpaceHostPort) String() string {
	return DialAddress(hp)
}

// GoString implements fmt.GoStringer.
func (hp SpaceHostPort) GoString() string {
	return hp.String()
}

type SpaceHostPorts []SpaceHostPort

func (hps SpaceHostPorts) HostPorts() HostPorts {
	addrs := make(HostPorts, len(hps))
	for i, hp := range hps {
		addrs[i] = hp
	}
	return addrs
}

// InSpaces returns the SpaceHostPorts that are in the input spaces.
func (hps SpaceHostPorts) InSpaces(spaces ...SpaceInfo) (SpaceHostPorts, bool) {
	if len(spaces) == 0 {
		logger.Errorf("host ports not filtered - no spaces given.")
		return hps, false
	}

	spaceInfos := SpaceInfos(spaces)
	var selectedHostPorts SpaceHostPorts
	for _, hp := range hps {
		if space := spaceInfos.Space(hp.SpaceID); space != nil {
			logger.Debugf("selected %q as a hostPort in space %q", hp.Value, space.Name)
			selectedHostPorts = append(selectedHostPorts, hp)
		}
	}

	if len(selectedHostPorts) > 0 {
		return selectedHostPorts, true
	}

	logger.Errorf("no hostPorts found in spaces %s", spaceInfos)
	return hps, false
}

// SelectInternal picks the best matching HostPorts that can be used as an
// endpoint for juju internal communication and returns them in NetAddr form.
// If there are no suitable addresses, an empty slice is returned.
func (hps SpaceHostPorts) SelectInternal(machineLocal bool) []string {
	indexes := bestAddressIndexes(len(hps), func(i int) Address {
		return hps[i].SpaceAddress
	}, internalAddressMatcher(machineLocal))

	out := make([]string, 0, len(indexes))
	for _, index := range indexes {
		out = append(out, DialAddress(hps[index]))
	}
	return out
}

// Less reports whether hp is ordered before hp2
// according to the criteria used by SortHostPorts.
func (hp SpaceHostPort) Less(hp2 SpaceHostPort) bool {
	order1 := hp.sortOrder()
	order2 := hp2.sortOrder()
	if order1 == order2 {
		if hp.SpaceAddress.Value == hp2.SpaceAddress.Value {
			return hp.Port() < hp2.Port()
		}
		return hp.SpaceAddress.Value < hp2.SpaceAddress.Value
	}
	return order1 < order2
}

// SpaceAddressesWithPort returns the input SpaceAddresses
// all associated with the given port.
func SpaceAddressesWithPort(addrs []SpaceAddress, port int) SpaceHostPorts {
	hps := make([]SpaceHostPort, len(addrs))
	for i, addr := range addrs {
		hps[i] = SpaceHostPort{
			SpaceAddress: addr,
			NetPort:      NetPort(port),
		}
	}
	return hps
}

// NewSpaceHostPorts creates a list of SpaceHostPorts
// from each input string address and port.
func NewSpaceHostPorts(port int, addresses ...string) SpaceHostPorts {
	hps := make([]SpaceHostPort, len(addresses))
	for i, addr := range addresses {
		hps[i] = SpaceHostPort{
			SpaceAddress: NewSpaceAddress(addr),
			NetPort:      NetPort(port),
		}
	}
	return hps
}

type hostPortsPreferringIPv4Slice []SpaceHostPort

func (hp hostPortsPreferringIPv4Slice) Len() int      { return len(hp) }
func (hp hostPortsPreferringIPv4Slice) Swap(i, j int) { hp[i], hp[j] = hp[j], hp[i] }
func (hp hostPortsPreferringIPv4Slice) Less(i, j int) bool {
	return hp[i].Less(hp[j])
}

// SortHostPorts sorts the given SpaceHostPort slice according to the sortOrder of
// each SpaceHostPort's embedded Address. See Address.sortOrder() for more info.
func SortHostPorts(hps []SpaceHostPort) {
	sort.Sort(hostPortsPreferringIPv4Slice(hps))
}

// APIHostPortsToNoProxyString converts list of lists of NetAddrs() to
// a NoProxy-like comma separated string, ignoring local addresses
func APIHostPortsToNoProxyString(ahp []SpaceHostPorts) string {
	noProxySet := set.NewStrings()
	for _, host := range ahp {
		for _, hp := range host {
			if hp.SpaceAddress.Scope == ScopeMachineLocal ||
				hp.SpaceAddress.Scope == ScopeLinkLocal {
				continue
			}
			noProxySet.Add(hp.SpaceAddress.Value)
		}
	}
	return strings.Join(noProxySet.SortedValues(), ",")
}

// EnsureFirstHostPort scans the given list of HostPorts and if
// "first" is found, it moved to index 0. Otherwise, if "first" is not
// in the list, it's inserted at index 0.
func EnsureFirstHostPort(first SpaceHostPort, hps []SpaceHostPort) []SpaceHostPort {
	var result []SpaceHostPort
	found := false
	for _, hp := range hps {
		if hp.String() == first.String() && !found {
			// Found, so skip it.
			found = true
			continue
		}
		result = append(result, hp)
	}
	// Insert it at the top.
	result = append([]SpaceHostPort{first}, result...)
	return result
}
