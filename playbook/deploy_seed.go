package playbook

import (
	"fmt"
	"path"
)

type DeploySeed map[string]*Component

func NewDeploySeed(c *Cluster, workDir string) *DeploySeed {
	cs := DeploySeed(make(map[string]*Component))
	for _, cp := range c.Components {
		cs[cp.Name] = &Component{
			MetaComponent: cp.MetaComponent,
			Hosts:         ConvertHosts(c.ID, cp.Hosts),
		}

		getInherentProperties(
			path.Join(workDir, cp.Name+PlaybookSuffix, cp.Version),
			cs[cp.Name],
		)
	}
	return &cs
}

func (ds *DeploySeed) AllHosts() map[string]*ClusterHost {
	hosts := make(map[string]*ClusterHost)
	for _, v := range map[string]*Component(*ds) {
		for _, hv := range v.Hosts {
			for _, h := range hv {
				hosts[h.IP] = h
			}
		}
	}
	return hosts
}

type Component struct {
	database.MetaComponent
	Inherent map[string]interface{}
	Hosts    map[string][]*ClusterHost
}

func ConvertHosts(
	clusterID string,
	sourceHosts map[string][]string,
) map[string][]*ClusterHost {
	destHost := make(map[string][]*ClusterHost)
	for k, v := range sourceHosts {
		hosts := make([]*ClusterHost, 0, len(v))

		for _, h := range v {
			hh, err := database.Instance().RetrieveHost(clusterID, h)
			if err != nil {
				panic(fmt.Sprintf("find the host %s error: %v", h, err))
			}
			hosts = append(hosts, hh)
		}

		destHost[k] = hosts
	}

	return destHost
}
