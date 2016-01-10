package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRouteParser(t *testing.T) {

	tests := []struct {
		method         string
		url            string
		expectedAction string
	}{
		{"GET", "/v1.21/version", ActionDockerVersion},
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
		{"POST", "/v.1.21/containers/id/attach", ActionContainerAttach},
		{"GET", "/v.1.21/containers/id/attach/ws", ActionContainerAttachWs},
		{"GET", "/v.1.21/containers/id/json", ActionContainerInspect},
		{"POST", "/v.1.21/containers/id/rename", ActionContainerRename},
		{"POST", "/v.1.21/containers/id/unpause", ActionContainerUnpause},
		{"GET", "/v.1.21/containers/json", ActionContainerList},
		{"DELETE", "/v.1.21/containers/id", ActionContainerDelete},
		{"GET", "/v.1.21/containers/id/stats", ActionContainerStats},
		{"GET", "/v.1.21/containers/id/changes", ActionContainerChanges},
		{"GET", "/v.1.21/containers/id/top", ActionContainerTop},
		{"POST", "/v.1.21/containers/create", ActionContainerCreate},
		{"GET", "/v.1.21/images/id/get", ActionImageArchive},
		{"POST", "/v.1.21/images/id/tag", ActionImageTag},
		{"GET", "/v.1.21/images/id/history", ActionImageHistory},
		{"POST", "/v.1.21/images/id/push", ActionImagePush},
		{"POST", "/v.1.21/images/create", ActionImageCreate},
		{"POST", "/v.1.21/images/load", ActionImageLoad},
		{"GET", "/v.1.21/images/json", ActionImageList},
		{"POST", "/v.1.21/images/build", ActionImageBuild},
		{"GET", "/v.1.21/images/id/json", ActionImageInspect},
		{"DELETE", "/v.1.21/images/id", ActionImageDelete},
		{"GET", "/v.1.21/_ping", ActionDockerPing},
		{"GET", "/v.1.21/info", ActionDockerInfo},
		{"GET", "/v.1.21/images/search", ActionImagesSearch},
		{"GET", "/v.1.21/networks", ActionNetworkList},
		{"GET", "/v.1.21/networks/id", ActionNetworkInspect},
		{"POST", "/v.1.21/networks/id/disconnect", ActionNetworkDisconnect},
		{"POST", "/v.1.21/networks/id/connect", ActionNetworkConnect},
		{"DELETE", "/v.1.21/networks/id", ActionNetworkRemove},
		{"DELETE", "/v.1.21/volumes/id", ActionVolumeRemove},
		{"POST", "/v.1.21/volumes/create", ActionVolumeCreate},
		{"GET", "/v.1.21/volumes/id", ActionVolumeInspect},
		{"GET", "/v.1.21/volumes", ActionVolumeList},
		{"GET", "/v.1.21/images/non_existing", ActionNone},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedAction, ParseRoute(test.method, test.url))
	}
}
