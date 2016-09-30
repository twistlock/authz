# Twistlock AuthZ Broker

[![Travis](https://travis-ci.org/twistlock/authz.svg)](https://travis-ci.org/twistlock/authz)
[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](https://godoc.org/github.com/twistlock/authz)
[![Go Report Card](https://goreportcard.com/badge/github.com/twistlock/authz)](https://goreportcard.com/report/github.com/twistlock/authz)


 
A basic extendable Docker [authorization plugin] (https://github.com/docker/docker/blob/master/docs/extend/authorization.md) that runs directly on the host or inside a container. The framework depends on [docker authentication plugin support] (https://github.com/docker/docker/pull/18514). Basic authorization is provided when Docker daemon is started with `--tlsverify` flag (username is extracted from the certificate common name).
Provided by [Twistlock] (https://www.twistlock.com).

![Twistlock Logo](https://www.twistlock.com/wp-content/uploads/2015/12/Twistlock-Logo.png)

## Basic policy enforcement 

The authorization broker is delivered as a reference implementation of a basic authorization mechanism, which consists of simple evaluation of user policies. The behavior of the plugin in the basic authorization flow is determined by the policy object:

```go
// BasicPolicy represent a single policy object that is evaluated in the authorization flow.
// Each policy object consists of multiple users and Docker actions, where each user belongs to a single policy.
//
// The policies are evaluated according to the following flow:
//   For each policy object check
//      If the user belongs to the policy
//         If action in request in policy allow otherwise deny
//   If no appropriate policy found, return deny
//
// Remark: In basic flow, each user must have a unique policy.
// If a user is used by more than one policy, the results may be inconsistent
type BasicPolicy struct {
	Actions []string `json:"actions"`  // Actions are the docker actions (mapped to authz terminology) that are allowed according to this policy
	                                   // Action are are specified as regular expressions
	Users   []string `json:"users"`    // Users are the users for which this policy apply to
	Name    string   `json:"name"`     // Name is the policy name
	Readonly bool    `json:"readonly"` // Readonly indicates this policy only allow get commands
}
```

For basic authorization flows, all policies reside in a single policy file under `/var/lib/authz-broker/policy.json`. The file  is continuously monitored and no restart is required upon changes.
The file format is [one policy JSON object per line](http://jsonlines.org/).  There should be no enclosing list or map, just one map per line.

The conversation between [Docker remote API] (https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/) (the URI and method that are passed Docker daemon to AuthZ plugin) to internal action parameters is defined by the [route parser] (https://github.com/twistlock/authz/blob/master/core/route_parser.go).
All requests and their associated authorization responses are logged to the standard output. Additional hooks such as syslog and log file is also available. To add additional [logrus hooks] (https://github.com/Sirupsen/logrus#hooks), see [extending the authorization plugin].

### Examples

Below are some examples for basic policy scenarios:
 1. Alice can run all Docker commands:                     `{"name":"policy_1","users":["alice"],"actions":[""]}`
 2. All users can run all Docker commands:                    `{"name":"policy_2","users":[""],"actions":[""]}`
 3. Alice and Bob can create new containers:              `{"name":"policy_3","users":["alice","bob"],"actions":["container_create"]}`
 4. Service account can read logs and run container top:  `{"name":"policy_4","users":["service_account"],"actions":["container_logs","container_top"]}` 
 5. Alice can perform anything on containers: `{"name":"policy_5","users":["alice"],"actions":["container"]}` 
 6. Alice can only perform get operations on containers:  `{"name":"policy_5","users":["alice"],"actions":["container"], "readonly":true }` 

## Installing the plugin

The authorization plugin can run as a container application or as a host service.

### Running inside a container 

 1. Install the containerized version of the Twistlock authorization plugin: 
```bash
 $ docker run -d  --restart=always -v /var/lib/authz-broker/policy.json:/var/lib/authz-broker/policy.json -v /run/docker/plugins/:/run/docker/plugins twistlock/authz-broker
```
    For auditing using syslog hook add the following settings to the docker command:<code>-e AUDITOR-HOOK:syslog -v /dev/log:/dev/log</code>
    For auditing using file add the following settings to the docker command:<code>-e AUDITOR-HOOK:file -v PATH_TO_LOCAL_LOG_FILE:/var/log/authz.log</code>

 2. Update Docker daemon to run with authorization enabled.
    For example, if Docker is installed as a systemd service:
```bash
 $ sudo systemctl edit --full docker.service 
```
 3. Add authz broker plugin parameter to ExecStart parameter
```bash
  ExecStart=/usr/bin/docker daemon -H fd:// --authorization-plugin=authz-broker
```
### Running as a stand-alone service

 *  Download Twistlock authZ binary (todo:link)
 *  Install Twistlock as a service 
```bash
   $ wget xxx | sudo sh
```
 * Update Docker daemon to run with authorization enabled.
     For example, if Docker is installed as a systemd service:
```bash
  $ sudo systemctl edit --full docker.service 
```
  add authz-plugin parameter to ExecStart parameter
```bash
   ExecStart=/usr/bin/docker daemon -H fd:// --authorization-plugin=authz-broker
``` 
  
# Dev environment
  
## Setting up local dev environment

  * Install [go 1.5](https://golang.org/dl/) and [docker](https://docs.docker.com/linux/step_one/).
  * Install [godep](https://github.com/tools/godep).
  * Clone the project.
  * Restore go dependencies:
```go
  $ godep restore
```
  * Build the binary and image:
```go
  $ make all
```

## Extending the authorization plugin

The framework consists of two extendable interfaces: the Authorizer, 
which handles the authorization flow; and the Auditor, which audits the request and response in the authorization flow.

```go
// Authorizer handles the authorization of docker requests and responses
type Authorizer interface {
	Init() error                                                 // Init initialize the handler
	AuthZReq(req *authorization.Request) *authorization.Response // AuthZReq handles the request from docker client
	// to docker daemon
	AuthZRes(req *authorization.Request) *authorization.Response // AuthZRes handles the response from docker deamon to docker client
}
```

```go
// Auditor audits the request and response sent from/to docker daemon
type Auditor interface {
	// AuditRequest audit the request sent from docker client and the associated authorization response
	// Docker client -> authorization -> audit -> Docker daemon
	AuditRequest(req *authorization.Request, pluginRes *authorization.Response)
	// AuditRequest audit the response sent from docker daemon and the associated authorization response
	// Docker daemon -> authorization  -> audit -> Docker client
	AuditResponse(req *authorization.Request, pluginRes *authorization.Response)
}
```

## Licensing

Twistlock authorization plugin is licensed under the Apache License, Version 2.0. See [LICENSE](https://github.com/twistlock/authz/blob/master/LICENSE) for the full license text.
