package compliancemanager

import (
	clusterDatastore "github.com/stackrox/rox/central/cluster/datastore"
	compIntegration "github.com/stackrox/rox/central/complianceoperator/v2/integration/datastore"
	profileDatastore "github.com/stackrox/rox/central/complianceoperator/v2/profiles/datastore"
	compScanSetting "github.com/stackrox/rox/central/complianceoperator/v2/scanconfigurations/datastore"
	"github.com/stackrox/rox/central/sensor/service/connection"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	manager Manager
	once    sync.Once
)

// Singleton returns the compliance operator manager
func Singleton() Manager {
	once.Do(func() {
		manager = New(connection.ManagerSingleton(), compIntegration.Singleton(), compScanSetting.Singleton(), clusterDatastore.Singleton(), profileDatastore.Singleton())
	})
	return manager
}
