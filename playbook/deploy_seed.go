package playbook

import (
	"path"

	"github.com/ygqbasic/nuwa/models"
)

type DeploySeed map[string]*models.ClusterComponent

func NewDeploySeed(c *models.Cluster, workDir string) *DeploySeed {
	comps := DeploySeed(make(map[string]*models.ClusterComponent))
	cs, _ := models.RetrieveClusterComponents(c.Id)

	for _, cp := range cs {
		comps[cp.ComponentName] = cp
		getInherentProperties(
			path.Join(workDir, cp.ComponentName+PlaybookSuffix, cp.Version),
			cp,
		)
	}
	return &comps
}

func (ds *DeploySeed) AllHosts() map[string]*models.ClusterHost {

	hosts := make(map[string]*models.ClusterHost)
	for _, v := range map[string]*models.ClusterComponent(*ds) {
		hv, _ := models.RetrieveComponentHosts(v.Hosts)

		for _, h := range hv {
			hosts[h.Ip] = h
		}
	}
	return hosts
}

type Component struct {
	models.MetaComponent
	Inherent map[string]interface{}
	Hosts    map[string][]*models.ClusterHost
}
