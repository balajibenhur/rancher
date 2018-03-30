package workload

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/rancher/pkg/api/customization/workload"
	"github.com/rancher/rancher/pkg/clustermanager"
	"github.com/rancher/types/apis/project.cattle.io/v3/schema"
)

func ConfigureStore(schemas *types.Schemas, manager *clustermanager.Manager) {
	workloadSchema := schemas.Schema(&schema.Version, "workload")
	store := NewAggregateStore(schemas.Schema(&schema.Version, "deployment"),
		schemas.Schema(&schema.Version, "replicaSet"),
		schemas.Schema(&schema.Version, "replicationController"),
		schemas.Schema(&schema.Version, "daemonSet"),
		schemas.Schema(&schema.Version, "statefulSet"),
		schemas.Schema(&schema.Version, "job"),
		schemas.Schema(&schema.Version, "cronJob"))

	for _, s := range store.Schemas {
		if s.ID == "deployment" {
			s.Formatter = workload.DeploymentFormatter
		} else {
			s.Formatter = workload.Formatter
		}
	}

	workloadSchema.Store = store
	actionWrapper := workload.ActionWrapper{
		ClusterManager: manager,
	}
	workloadSchema.ActionHandler = actionWrapper.ActionHandler
	workloadSchema.LinkHandler = workload.Handler{}.LinkHandler
}
