package localkube

import (
	"github.com/redspread/localkube"
	"github.com/supergiant/guber"
	"github.com/supergiant/supergiant/pkg/core"
	"github.com/supergiant/supergiant/pkg/model"
)

// Provider Holds DO account info.
type Provider struct {
	Core  *core.Core
	Token string
	LK    localkube.LocalKube
}

var (
	LK *localkube.LocalKube

	DNSDomain     = "cluster.local"
	ClusterDNSIP  = "10.1.30.3"
	DNSServerAddr = "172.17.0.1:1970"
)

// ValidateAccount Valitades DO account info.
func (p *Provider) ValidateAccount(m *model.CloudAccount) error {
	return nil
}

// CreateKube creates a new DO kubernetes cluster.
func (p *Provider) CreateKube(m *model.Kube, action *core.Action) error {
	p.LK.Start(m.Name)

	return nil
}

// DeleteKube deletes a DO kubernetes cluster.
func (p *Provider) DeleteKube(m *model.Kube) error {
	p.LK.Stop(m.Name)
	return nil
}

// CreateNode creates a new minion on DO kubernetes cluster.
func (p *Provider) CreateNode(m *model.Node, action *core.Action) error {
	return nil
}

// DeleteNode deletes a minsion on a DO kubernetes cluster.
func (p *Provider) DeleteNode(m *model.Node) error {
	return nil
}

// CreateVolume createss a Volume on DO for Kubernetes
func (p *Provider) CreateVolume(m *model.Volume, action *core.Action) error {
	return nil
}

func (p *Provider) KubernetesVolumeDefinition(m *model.Volume) *guber.Volume {
	return nil
}

// ResizeVolume re-sizes volume on DO kubernetes cluster.
func (p *Provider) ResizeVolume(m *model.Volume, action *core.Action) error {
	return nil
}

// WaitForVolumeAvailable waits for DO volume to become available.
func (p *Provider) WaitForVolumeAvailable(m *model.Volume, action *core.Action) error {
	return nil
}

// DeleteVolume deletes a DO volume.
func (p *Provider) DeleteVolume(m *model.Volume) error {
	return nil
}

// CreateEntrypoint creates a new Load Balancer for Kubernetes in DO
func (p *Provider) CreateEntrypoint(m *model.Entrypoint, action *core.Action) error {
	return nil
}

//AddPortToEntrypoint adds an external entrypoint to a Loadbalancer in DO.
func (p *Provider) AddPortToEntrypoint(m *model.Entrypoint, lbPort int64, nodePort int64) error {
	return nil
}

// RemovePortFromEntrypoint removes external entrypoint from Loadbalancer on DO.
func (p *Provider) RemovePortFromEntrypoint(m *model.Entrypoint, lbPort int64) error {
	return nil
}

// DeleteEntrypoint deletes load balancer from DO.
func (p *Provider) DeleteEntrypoint(m *model.Entrypoint) error {
	return nil
}
