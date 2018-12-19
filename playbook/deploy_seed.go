package playbook

import (
	"encoding/json"
	"path"

	"github.com/ygqbasic/nuwa/models"
)

type DeploySeed map[string]*Component

func NewDeploySeed(c *models.Cluster, workDir string) *DeploySeed {
	cs := DeploySeed(make(map[string]*Component))
	for _, cp := range c.Components {

		cs[cp.ComponentName] = json.Unmarshal([]byte(cp.Component), &Component)

		getInherentProperties(
			path.Join(workDir, cp.ComponentName+PlaybookSuffix, cp.Version),
			cs[cp.ComponentName],
		)
	}
	return &cs
}

func (ds *DeploySeed) AllHosts() map[string]*models.ClusterHost {
	hosts := make(map[string]*models.ClusterHost)
	for _, v := range map[string]*Component(*ds) {
		for _, hv := range v.Hosts {
			for _, h := range hv {
				hosts[h.Ip] = h
			}
		}
	}
	return hosts
}

type Component struct {
	models.MetaComponent
	Inherent map[string]interface{}
	Hosts    map[string][]*models.ClusterHost
}
