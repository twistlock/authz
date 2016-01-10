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
	route{pattern: "/containers/.+/wait", method: "POST", action: ActionContainerWait},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#resize-a-container-tty
	route{pattern: "/containers/.+/resize", method: "POST", action: ActionContainerResize},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#export-a-container
	route{pattern: "/containers/.+/export", method: "POST", action: ActionContainerExport},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#export-a-container
	route{pattern: "/containers/.+/stop", method: "POST", action: ActionContainerStop},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#kill-a-container
	route{pattern: "/containers/.*/kill", method: "POST", action: ActionContainerKill},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#restart-a-container
	route{pattern: "/containers/.+/restart", method: "POST", action: ActionContainerRestart},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#start-a-container
	route{pattern: "/containers/.+/start", method: "POST", action: ActionContainerStart},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#exec-create
	route{pattern: "/containers/.+/exec", method: "POST", action: ActionContainerExecCreate},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#unpause-a-container
	route{pattern: "/containers/.+/unpause", method: "POST", action: ActionContainerUnpause},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#pause-a-container
	route{pattern: "/containers/.+/pause", method: "POST", action: ActionContainerPause},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#copy-files-or-folders-from-a-container
	route{pattern: "/containers/.+/copy", method: "POST", action: ActionContainerCopyFiles},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#extract-an-archive-of-files-or-folders-to-a-directory-in-a-container
	route{pattern: "/containers/.+/archive", method: "PUT", action: ActionContainerArchiveExtract},
	route{pattern: "/containers/.+/archive", method: "HEAD", action: ActionContainerArchiveInfo},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#get-an-archive-of-a-filesystem-resource-in-a-container
	route{pattern: "/containers/.+/archive", method: "GET", action: ActionContainerArchive},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#attach-to-a-container-websocket
	route{pattern: "/containers/.+/attach/ws", method: "GET", action: ActionContainerAttachWs},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#attach-to-a-container
	route{pattern: "/containers/.+/attach", method: "POST", action: ActionContainerAttach},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#list-containers
	route{pattern: "/containers/json", method: "GET", action: ActionContainerList},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#inspect-a-container
	route{pattern: "/containers/.+/json", method: "GET", action: ActionContainerInspect},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#remove-a-container
	route{pattern: "/containers/.+", method: "DELETE", action: ActionContainerDelete},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#rename-a-container
	route{pattern: "/containers/.+/rename", method: "POST", action: ActionContainerRename},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#get-container-stats-based-on-resource-usage
	route{pattern: "/containers/.+/stats", method: "GET", action: ActionContainerStats},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#inspect-changes-on-a-container-s-filesystem
	route{pattern: "/containers/.+/changes", method: "GET", action: ActionContainerChanges},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#list-processes-running-inside-a-container
	route{pattern: "/containers/.+/top", method: "GET", action: ActionContainerTop},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#get-container-logs
	route{pattern: "/containers/.+/logs", method: "GET", action: ActionContainerLogs},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#create-a-container
	route{pattern: "/containers/create", method: "POST", action: ActionContainerCreate},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#get-a-tarball-containing-all-images
	route{pattern: "/images/.+./get", method: "GET", action: ActionImageArchive},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#search-images
	route{pattern: "/images/search", method: "GET", action: ActionImagesSearch},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#tag-an-image-into-a-repository
	route{pattern: "/images/.+/tag", method: "POST", action: ActionImageTag},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#inspect-an-image
	route{pattern: "/images/.+/json", method: "GET", action: ActionImageInspect},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.18/#inspect-an-image
	route{pattern: "/images/.+", method: "DELETE", action: ActionImageDelete},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#get-the-history-of-an-image
	route{pattern: "/images/.+/history", method: "GET", action: ActionImageHistory},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#push-an-image-on-the-registry
	route{pattern: "/images/.+/push", method: "POST", action: ActionImagePush},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#create-an-image
	route{pattern: "/images/create", method: "POST", action: ActionImageCreate},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#load-a-tarball-with-a-set-of-images-and-tags-into-docker
	route{pattern: "/images/load", method: "POST", action: ActionImageLoad},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#build-image-from-a-dockerfile
	route{pattern: "/images/build", method: "POST", action: ActionImageBuild},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#list-images
	route{pattern: "/images/json", method: "GET", action: ActionImageList},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#ping-the-docker-server
	route{pattern: "/_ping", method: "GET", action: ActionDockerPing},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#display-system-wide-information
	route{pattern: "/info", method: "GET", action: ActionDockerInfo},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#exec-inspect
	route{pattern: "/exec/.+/json", method: "GET", action: ActionContainerExecInspect},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#exec-start
	route{pattern: "/exec/.+/start", method: "POST", action: ActionContainerExecStart},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#inspect-a-volume
	route{pattern: "/volumes/.+", method: "GET", action: ActionVolumeInspect},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#list-volumes
	route{pattern: "/volumes", method: "GET", action: ActionVolumeList},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#create-a-volume
	route{pattern: "/volumes/create", method: "POST", action: ActionVolumeCreate},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#remove-a-volume
	route{pattern: "/volumes/.+", method: "DELETE", action: ActionVolumeRemove},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#inspect-network
	route{pattern: "/networks/.+", method: "GET", action: ActionNetworkInspect},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#list-networks
	route{pattern: "/networks", method: "GET", action: ActionNetworkList},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#create-a-network
	route{pattern: "/networks/create", method: "POST", action: ActionNetworkCreate},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#connect-a-container-to-a-network
	route{pattern: "/networks/.+/connect", method: "POST", action: ActionNetworkConnect},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#disconnect-a-container-from-a-network
	route{pattern: "/networks/.+/disconnect", method: "POST", action: ActionNetworkDisconnect},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#remove-a-network
	route{pattern: "/networks/.+", method: "DELETE", action: ActionNetworkRemove},
}

// ParseRoute convert a method/url pattern to corresponding docker action
func ParseRoute(method, url string) string {

	for _, route := range routes {
		if route.method == method {
			match, err := regexp.MatchString(route.pattern, url)
			if err == nil && match {
				return route.action
			}

		}
	}

	return ActionNone
}
