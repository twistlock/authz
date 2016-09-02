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
		{"POST", "/v1.24/commit", ActionContainerCommit},
		{"POST", "/v1.24/events", ActionDockerEvents},
		{"GET", "/v1.24/version", ActionDockerVersion},
		{"POST", "/v1.24/auth", ActionDockerCheckAuth},
		{"POST", "/v1.24/containers/id/wait", ActionContainerWait},
		{"POST", "/v1.24/containers/id/wait", ActionContainerWait},
		{"POST", "/v1.24/containers/id/resize", ActionContainerResize},
		{"POST", "/v1.24/containers/id/stop", ActionContainerStop},
		{"POST", "/v1.24/containers/id/kill", ActionContainerKill},
		{"POST", "/v1.24/containers/id/restart", ActionContainerRestart},
		{"POST", "/v1.24/containers/id/start", ActionContainerStart},
		{"POST", "/v1.24/containers/id/exec", ActionContainerExecCreate},
		{"GET", "/v1.24/containers/id/archive", ActionContainerArchive},
		{"GET", "/v1.24/exec/id/json", ActionContainerExecInspect},
		{"POST", "/v1.24/exec/id/start", ActionContainerExecStart},
		{"POST", "/v1.24/exec/id/resize", ActionContainerExecResize},
		{"HEAD", "/v1.24/containers/id/archive", ActionContainerArchiveInfo},
		{"PUT", "/v1.24/containers/id/archive", ActionContainerArchiveExtract},
		{"GET", "/v1.24/containers/id/export", ActionContainerExport},
		{"POST", "/v1.24/containers/id/attach", ActionContainerAttach},
		{"GET", "/v1.24/containers/id/attach/ws", ActionContainerAttachWs},
		{"GET", "/v1.24/containers/id/json", ActionContainerInspect},
		{"POST", "/v1.24/containers/id/rename", ActionContainerRename},
		{"POST", "/v1.24/containers/id/unpause", ActionContainerUnpause},
		{"GET", "/v1.24/containers/json", ActionContainerList},
		{"DELETE", "/v1.24/containers/id", ActionContainerDelete},
		{"GET", "/v1.24/containers/id/stats", ActionContainerStats},
		{"GET", "/v1.24/containers/id/changes", ActionContainerChanges},
		{"GET", "/v1.24/containers/id/top", ActionContainerTop},
		{"POST", "/v1.24/containers/create", ActionContainerCreate},
		{"POST", "/v1.24/containers/id/update", ActionContainerUpdate},
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
		{"GET", "/v1.24/networks", ActionNetworkList},
		{"GET", "/v1.24/networks/id", ActionNetworkInspect},
		{"POST", "/v1.24/networks/id/disconnect", ActionNetworkDisconnect},
		{"POST", "/v1.24/networks/id/connect", ActionNetworkConnect},
		{"DELETE", "/v1.24/networks/id", ActionNetworkRemove},
		{"DELETE", "/v1.24/volumes/id", ActionVolumeRemove},
		{"POST", "/v1.24/volumes/create", ActionVolumeCreate},
		{"GET", "/v1.24/volumes/id", ActionVolumeInspect},
		{"GET", "/v1.24/volumes", ActionVolumeList},
		{"GET", "/v.1.21/images/non_existing", ActionNone},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedAction, ParseRoute(test.method, test.url))
	}
}
