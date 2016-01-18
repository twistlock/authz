package core

import "github.com/docker/docker/pkg/authorization"

// Authorizer handles the authorization of docker requests and responses
type Authorizer interface {
	// Init initialize the authorizer
	Init() error                                                 // Init initialize the handler
	AuthZReq(req *authorization.Request) *authorization.Response // AuthZReq handles the request from docker client
	// to docker daemon
	AuthZRes(req *authorization.Request) *authorization.Response // AuthZRes handles the response from docker daemon to docker client
}

// Auditor audits the request and response sent from/to docker daemon
type Auditor interface {
	// AuditRequest audit the request sent from docker client and the associated authorization response
	// Docker client -> authorization -> audit -> Docker daemon
	AuditRequest(req *authorization.Request, pluginRes *authorization.Response) error
	// AuditRequest audit the response sent from docker daemon and the associated authorization response
	// Docker daemon -> authorization  -> audit -> Docker client
	AuditResponse(req *authorization.Request, pluginRes *authorization.Response) error
}
