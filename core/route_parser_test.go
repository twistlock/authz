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
		{"HEAD", "/v1.21/containers/id/archive", ActionContainerArchiveHead},
		{"PUT", "/v1.21/containers/id/archive", ActionContainerArchiveExtract},
		{"POST", "/v.1.21/containers/id/attach", ActionContainerAttach},
		{"GET", "/v.1.21/containers/id/attach/ws", ActionContainerAttachWs},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedAction, ParseRoute(test.method, test.url))
	}
}
