package core

var (
	// ActionContainerArchive describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#get-an-archive-of-a-filesystem-resource-in-a-container
	ActionContainerArchive = "container_archive"
	// ActionContainerArchiveExtract describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#extract-an-archive-of-files-or-folders-to-a-directory-in-a-container
	ActionContainerArchiveExtract = "container_archive_extract"
	// ActionContainerArchiveInfo describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#retrieving-information-about-files-and-folders-in-a-container
	ActionContainerArchiveInfo = "container_archive_info"
	// ActionContainerAttach describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#attach-to-a-container
	ActionContainerAttach = "container_attach"
	// ActionContainerAttachWs describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#attach-to-a-container-websocket
	ActionContainerAttachWs = "container_attach_websocket"

	// ActionContainerChanges describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#inspect-changes-on-a-container-s-filesystem
	ActionContainerChanges = "container_changes"
	// ActionContainerCommit describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#create-a-new-image-from-a-container-s-changes
	ActionContainerCommit = "container_commit"
	// ActionContainerCopyFiles describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#copy-files-or-folders-from-a-container
	ActionContainerCopyFiles = "container_copyfiles"
	// ActionContainerCreate describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#create-a-container
	ActionContainerCreate = "container_create"
	ActionContainerDelete = "container_delete"
	// ActionContainerExecCreate describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#exec-create
	ActionContainerExecCreate  = "container_exec_create"
	ActionContainerExecInspect = "container_exec_inspect"
	ActionContainerExecStart   = "container_exec_start"
	ActionContainerExport      = "container_export"
	ActionContainerInspect     = "container_inspect"
	ActionContainerKill        = "container_kill"
	// ActionContainerList describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#list-containers
	ActionContainerList = "container_list"
	// ActionContainerLogs describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#get-container-logs
	ActionContainerLogs = "container_logs"
	// ActionContainerPause describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#pause-a-container
	ActionContainerPause = "container_pause"
	// ActionContainerRename describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#rename-a-container
	ActionContainerRename = "container_rename"
	// ActionContainerResize describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#resize-a-container-tty
	ActionContainerResize = "container_resize"
	// ActionContainerRestart describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#restart-a-container
	ActionContainerRestart = "container_restart"
	ActionContainerStart   = "container_start"
	ActionContainerStats   = "container_stats"
	ActionContainerStop    = "container_stop"
	ActionContainerTop     = "container_top"
	ActionContainerUnpause = "container_unpause"
	// ActionContainerWait describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#wait-a-container
	ActionContainerWait = "container_wait"
	// ActionDockerCheckAuth describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#check-auth-configuration
	ActionDockerCheckAuth = "docker_auth"
	// ActionDockerEvents describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#monitor-docker-s-events
	ActionDockerEvents = "docker_events"
	// ActionDockerInfo describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#display-system-wide-information
	ActionDockerInfo    = "docker_info"
	ActionDockerPing    = "docker_ping"
	ActionDockerVersion = "docker_version"
	ActionImageArchive  = "images_archive"
	ActionImageBuild    = "image_build"
	// ActionImageCreate describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#create-an-image
	ActionImageCreate  = "image_create"
	ActionImageDelete  = "image_delete"
	ActionImageHistory = "image_history"
	ActionImageInspect = "image_inspect"
	ActionImageList    = "image_list"
	ActionImageLoad    = "images_load"
	ActionImagePush    = "image_push"
	ActionImagesSearch = "images_search"
	ActionImageTag     = "image_tag"
	ActionVolumeList   = "volume_list"
	// ActionVolumeCreate describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#create-a-volume
	ActionVolumeCreate  = "volume_create"
	ActionVolumeInspect = "volume_inspect"
	// ActionVolumeRemove describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#remove-a-volume
	ActionVolumeRemove = "volume_remove"
	// ActionNetworkList describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#list-networks
	ActionNetworkList    = "network_list"
	ActionNetworkInspect = "network_inspect"
	// ActionNetworkCreate describes https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#create-a-network
	ActionNetworkCreate     = "network_create"
	ActionNetworkConnect    = "network_connect"
	ActionNetworkDisconnect = "network_disconnect"
	ActionNetworkRemove     = "network_remove"
	ActionNone              = ""
)
