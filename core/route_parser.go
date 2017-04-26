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

	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/create-a-new-image-from-a-container-s-changes
	{pattern: "/commit", method: "POST", action: ActionContainerCommit},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/monitor-docker-s-events
	{pattern: "/events", method: "POST", action: ActionDockerEvents},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/show-the-docker-version-information
	{pattern: "/version", method: "GET", action: ActionDockerVersion},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/check-auth-configuration
	{pattern: "/auth", method: "POST", action: ActionDockerCheckAuth},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/wait-a-container
	{pattern: "/containers/.+/wait", method: "POST", action: ActionContainerWait},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/resize-a-container-tty
	{pattern: "/containers/.+/resize", method: "POST", action: ActionContainerResize},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/export-a-container
	{pattern: "/containers/.+/export", method: "GET", action: ActionContainerExport},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/stop-a-container
	{pattern: "/containers/.+/stop", method: "POST", action: ActionContainerStop},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/kill-a-container
	{pattern: "/containers/.*/kill", method: "POST", action: ActionContainerKill},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/restart-a-container
	{pattern: "/containers/.+/restart", method: "POST", action: ActionContainerRestart},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/start-a-container
	{pattern: "/containers/.+/start", method: "POST", action: ActionContainerStart},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/exec-create
	{pattern: "/containers/.+/exec", method: "POST", action: ActionContainerExecCreate},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/unpause-a-container
	{pattern: "/containers/.+/unpause", method: "POST", action: ActionContainerUnpause},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/pause-a-container
	{pattern: "/containers/.+/pause", method: "POST", action: ActionContainerPause},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/extract-an-archive-of-files-or-folders-to-a-directory-in-a-container
	{pattern: "/containers/.+/archive", method: "PUT", action: ActionContainerArchiveExtract},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/retrieving-information-about-files-and-folders-in-a-container
	{pattern: "/containers/.+/archive", method: "HEAD", action: ActionContainerArchiveInfo},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/get-an-archive-of-a-filesystem-resource-in-a-container
	{pattern: "/containers/.+/archive", method: "GET", action: ActionContainerArchive},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/attach-to-a-container-websocket
	{pattern: "/containers/.+/attach/ws", method: "GET", action: ActionContainerAttachWs},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/attach-to-a-container
	{pattern: "/containers/.+/attach", method: "POST", action: ActionContainerAttach},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/list-containers
	{pattern: "/containers/json", method: "GET", action: ActionContainerList},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/inspect-a-container
	{pattern: "/containers/.+/json", method: "GET", action: ActionContainerInspect},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/remove-a-container
	{pattern: "/containers/.+", method: "DELETE", action: ActionContainerDelete},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/rename-a-container
	{pattern: "/containers/.+/rename", method: "POST", action: ActionContainerRename},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/get-container-stats-based-on-resource-usage
	{pattern: "/containers/.+/stats", method: "GET", action: ActionContainerStats},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/inspect-changes-on-a-container-s-filesystem
	{pattern: "/containers/.+/changes", method: "GET", action: ActionContainerChanges},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/list-processes-running-inside-a-container
	{pattern: "/containers/.+/top", method: "GET", action: ActionContainerTop},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/get-container-logs
	{pattern: "/containers/.+/logs", method: "GET", action: ActionContainerLogs},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/create-a-container
	{pattern: "/containers/create", method: "POST", action: ActionContainerCreate},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/update-a-container
	{pattern: "/containers/.+/update", method: "POST", action: ActionContainerUpdate},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/get-a-tarball-containing-all-images-in-a-repository
	{pattern: "/images/.+./get", method: "GET", action: ActionImageArchive},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/search-images
	{pattern: "/images/search", method: "GET", action: ActionImagesSearch},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/tag-an-image-into-a-repository
	{pattern: "/images/.+/tag", method: "POST", action: ActionImageTag},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/inspect-an-image
	{pattern: "/images/.+/json", method: "GET", action: ActionImageInspect},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/remove-an-image
	{pattern: "/images/.+", method: "DELETE", action: ActionImageDelete},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/get-the-history-of-an-image
	{pattern: "/images/.+/history", method: "GET", action: ActionImageHistory},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/push-an-image-on-the-registry
	{pattern: "/images/.+/push", method: "POST", action: ActionImagePush},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/create-an-image
	{pattern: "/images/create", method: "POST", action: ActionImageCreate},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/load-a-tarball-with-a-set-of-images-and-tags-into-docker
	{pattern: "/images/load", method: "POST", action: ActionImageLoad},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/build-image-from-a-dockerfile
	{pattern: "/build", method: "POST", action: ActionImageBuild},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/list-images
	{pattern: "/images/json", method: "GET", action: ActionImageList},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/ping-the-docker-server
	{pattern: "/_ping", method: "GET", action: ActionDockerPing},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/display-system-wide-information
	{pattern: "/info", method: "GET", action: ActionDockerInfo},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/exec-inspect
	{pattern: "/exec/.+/json", method: "GET", action: ActionContainerExecInspect},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/exec-start
	{pattern: "/exec/.+/start", method: "POST", action: ActionContainerExecStart},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/exec-resize
	{pattern: "/exec/.+/resize", method: "POST", action: ActionContainerExecResize},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/inspect-a-volume
	{pattern: "/volumes/.+", method: "GET", action: ActionVolumeInspect},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/list-volumes
	{pattern: "/volumes", method: "GET", action: ActionVolumeList},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/create-a-volume
	{pattern: "/volumes/create", method: "POST", action: ActionVolumeCreate},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/remove-a-volume
	{pattern: "/volumes/.+", method: "DELETE", action: ActionVolumeRemove},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/inspect-network
	{pattern: "/networks/.+", method: "GET", action: ActionNetworkInspect},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/list-networks
	{pattern: "/networks", method: "GET", action: ActionNetworkList},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/create-a-network
	{pattern: "/networks/create", method: "POST", action: ActionNetworkCreate},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/connect-a-container-to-a-network
	{pattern: "/networks/.+/connect", method: "POST", action: ActionNetworkConnect},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/disconnect-a-container-from-a-network
	{pattern: "/networks/.+/disconnect", method: "POST", action: ActionNetworkDisconnect},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/remove-a-network
	{pattern: "/networks/.+", method: "DELETE", action: ActionNetworkRemove},
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
