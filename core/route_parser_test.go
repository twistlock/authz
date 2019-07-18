package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouteParser(t *testing.T) {

	tests := []struct {
		method         string
		url            string
		expectedAction string
	}{
		{"GET", "/v1.21/version", ActionDockerVersion},
		{"POST", "/v1.21/build", ActionImageBuild},
		{"POST", "/v1.21/containers/id/wait", ActionContainerWait},
		{"POST", "/v1.21/containers/id/wait", ActionContainerWait},
		{"POST", "/v1.21/containers/id/resize", ActionContainerResize},
		{"POST", "/v1.21/containers/id/stop", ActionContainerStop},
		{"POST", "/v1.21/containers/id/kill", ActionContainerKill},
		{"POST", "/v1.21/containers/id/restart", ActionContainerRestart},
		{"POST", "/v1.21/containers/id/start", ActionContainerStart},
		{"POST", "/v1.21/containers/id/exec", ActionContainerExecCreate},
		{"GET", "/v1.21/containers/id/archive", ActionContainerArchive},
		{"GET", "/v1.21/exec/id/json", ActionContainerExecInspect},
		{"POST", "/v1.21/exec/id/start", ActionContainerExecStart},
		{"HEAD", "/v1.21/containers/id/archive", ActionContainerArchiveInfo},
		{"PUT", "/v1.21/containers/id/archive", ActionContainerArchiveExtract},
		{"POST", "/v1.21/containers/id/export", ActionContainerExport},
		{"POST", "/v1.21/containers/id/attach", ActionContainerAttach},
		{"GET", "/v1.21/containers/id/attach/ws", ActionContainerAttachWs},
		{"GET", "/v1.21/containers/id/json", ActionContainerInspect},
		{"POST", "/v1.21/containers/id/rename", ActionContainerRename},
		{"POST", "/v1.21/containers/id/unpause", ActionContainerUnpause},
		{"GET", "/v1.21/containers/json", ActionContainerList},
		{"DELETE", "/v1.21/containers/id", ActionContainerDelete},
		{"GET", "/v1.21/containers/id/stats", ActionContainerStats},
		{"GET", "/v1.21/containers/id/changes", ActionContainerChanges},
		{"GET", "/v1.21/containers/id/top", ActionContainerTop},
		{"POST", "/v1.21/containers/create", ActionContainerCreate},
		{"GET", "/v1.21/images/id/get", ActionImageArchive},
		{"POST", "/v1.21/images/id/tag", ActionImageTag},
		{"GET", "/v1.21/images/id/history", ActionImageHistory},
		{"POST", "/v1.21/images/id/push", ActionImagePush},
		{"POST", "/v1.21/images/create", ActionImageCreate},
		{"POST", "/v1.21/images/load", ActionImageLoad},
		{"GET", "/v1.21/images/json", ActionImageList},
		{"POST", "/v1.21/images/build", ActionImageBuild},
		{"GET", "/v1.21/images/id/json", ActionImageInspect},
		{"DELETE", "/v1.21/images/id", ActionImageDelete},
		{"POST", "/v1.37/images/prune", ActionImagePrune},
		{"GET", "/v1.21/_ping", ActionDockerPing},
		{"GET", "/v1.21/info", ActionDockerInfo},
		{"GET", "/v1.21/images/search", ActionImagesSearch},
		{"GET", "/v1.21/networks", ActionNetworkList},
		{"GET", "/v1.21/networks/id", ActionNetworkInspect},
		{"POST", "/v1.21/networks/id/disconnect", ActionNetworkDisconnect},
		{"POST", "/v1.21/networks/id/connect", ActionNetworkConnect},
		{"DELETE", "/v1.21/networks/id", ActionNetworkRemove},
		{"DELETE", "/v1.21/volumes/id", ActionVolumeRemove},
		{"POST", "/v1.21/volumes/create", ActionVolumeCreate},
		{"GET", "/v1.21/volumes/id", ActionVolumeInspect},
		{"GET", "/v1.21/volumes", ActionVolumeList},
		{"GET", "/v1.21/images/non_existing", ActionNone},
		{"GET", "/v1.39/swarm", ActionSwarmInspect},
		{"POST", "/v1.39/swarm/init", ActionSwarmInit},
		{"POST", "/v1.39/swarm/join", ActionSwarmJoin},
		{"POST", "/v1.39/swarm/leave", ActionSwarmLeave},
		{"POST", "/v1.39/swarm/update", ActionSwarmUpdate},
		{"GET", "/v1.39/swarm/unlockkey", ActionSwarmUnlockKey},
		{"POST", "/v1.39/swarm/unlock", ActionSwarmUnlock},
		{"GET", "/v1.39/nodes", ActionNodeList},
		{"GET", "/v1.39/nodes/id", ActionNodeInspect},
		{"DELETE", "/v1.39/nodes/id", ActionNodeDelete},
		{"POST", "/v1.39/nodes/id/update", ActionNodeUpdate},
		{"GET", "/v1.39/services", ActionServiceList},
		{"POST", "/v1.39/services/create", ActionServiceCreate},
		{"GET", "/v1.39/services/id", ActionServiceInspect},
		{"DELETE", "/v1.39/services/id", ActionServiceDelete},
		{"POST", "/v1.39/services/id/update", ActionServiceUpdate},
		{"GET", "/v1.39/services/id/logs", ActionServiceLogs},
		{"GET", "/v1.39/tasks", ActionTaskList},
		{"GET", "/v1.39/tasks/id", ActionTaskInspect},
		{"GET", "/v1.39/secrets", ActionSecretList},
		{"POST", "/v1.39/secrets/create", ActionSecretCreate},
		{"GET", "/v1.39/secrets/id", ActionSecretInspect},
		{"DELETE", "/v1.39/secrets/id", ActionSecretDelete},
		{"POST", "/v1.39/secrets/id/update", ActionSecretUpdate},
		{"POST", "/v1.39/configs/create", ActionConfigCreate},
		{"GET", "/v1.39/configs/id", ActionConfigInspect},
		{"DELETE", "/v1.39/configs/id", ActionConfigDelete},
		{"POST", "/v1.39/configs/id/update", ActionConfigUpdate},
		{"GET", "/v1.39/distribution/twistlock/authz-broker:latest/json", ActionDistributionInspect},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedAction, ParseRoute(test.method, test.url))
	}
}
