package core

import (
	"regexp"
)

type route struct {
	pattern string
	method  string
	action  string
}

var routes = []route{
	// https://docs.docker.com/reference/api/docker_remote_api_v1.20/#create-a-new-image-from-a-container-s-changes
	route{pattern: "/commit", method: "POST", action: ActionContainerCommit},

	// https://docs.docker.com/reference/api/docker_remote_api_v1.20/#monitor-docker-s-events
	route{pattern: "/events", method: "POST", action: ActionDockerEvents},

	// https://docs.docker.com/reference/api/docker_remote_api_v1.20/#show-the-docker-version-information
	route{pattern: "/version", method: "GET", action: ActionDockerVersion},

	// https://docs.docker.com/reference/api/docker_remote_api_v1.20/#check-auth-configuration
	route{pattern: "/auth", method: "POST", action: ActionDockerCheckAuth},

	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#wait-a-container
	route{pattern: "/containers/.*/wait", method: "POST", action: ActionContainerWait},

	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#resize-a-container-tty
	route{pattern: "/containers/.*/resize", method: "POST", action: ActionContainerResize},

	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#export-a-container
	route{pattern: "/containers/.*/export", method: "POST", action: ActionContainerExport},

	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#export-a-container
	route{pattern: "/containers/.*/stop", method: "POST", action: ActionContainerStop},

	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#kill-a-container
	route{pattern: "/containers/.*/kill", method: "POST", action: ActionContainerKill},

	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#restart-a-container
	route{pattern: "/containers/.*/restart", method: "POST", action: ActionContainerRestart},

	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#start-a-container
	route{pattern: "/containers/.*/start", method: "POST", action: ActionContainerStart},

	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#exec-create
	route{pattern: "/containers/.*/exec", method: "POST", action: ActionContainerExecCreate},

	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#unpause-a-container
	route{pattern: "/containers/.*/unpause", method: "POST", action: ActionContainerUnpause},

	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#pause-a-container
	route{pattern: "/containers/.*/pause", method: "POST", action: ActionContainerPause},

	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#copy-files-or-folders-from-a-container
	route{pattern: "/containers/.*/copy", method: "POST", action: ActionContainerCopyFiles},

	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#extract-an-archive-of-files-or-folders-to-a-directory-in-a-container
	route{pattern: "/containers/.*/archive", method: "PUT", action: ActionContainerArchiveExtract},

	route{pattern: "/containers/.*/archive", method: "HEAD", action: ActionContainerArchiveHead},

	route{pattern: "/containers/.*/archive", method: "GET", action: ActionContainerArchive},

	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#attach-to-a-container-websocket
	route{pattern: "/containers/.*/attach/ws", method: "GET", action: ActionContainerAttachWs},

	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#attach-to-a-container
	route{pattern: "/containers/.*/attach", method: "POST", action: ActionContainerAttach},
}

func ParseRoute(method, url string) string {

	for _, route := range routes {
		if route.method == method {
			match, err := regexp.MatchString(route.pattern, url)
			if err == nil && match {
				return route.action
			}

		}
	}

	return ""
}
