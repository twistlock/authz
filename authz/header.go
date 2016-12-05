package authz

import (
	"fmt"
	"log/syslog"
	"os"
	"path"

	"github.com/Sirupsen/logrus"
	logrus_syslog "github.com/Sirupsen/logrus/hooks/syslog"
	"github.com/docker/docker/pkg/authorization"
	"github.com/twistlock/authz/core"
)

type headerAuthorizer struct {
	header string
}

// NewBasicAuthZAuthorizer creates a new basic authorizer
func NewHeaderAuthZAuthorizer() core.Authorizer {
	return &headerAuthorizer{header: "X-Docker-Unprivileged"}
}

func (f *headerAuthorizer) Init() error {
	return nil
}

func (f *headerAuthorizer) AuthZReq(authZReq *authorization.Request) *authorization.Response {

	logrus.Debugf("Received AuthZ request, method: '%s', url: '%s'", authZReq.RequestMethod, authZReq.RequestURI)

	action := core.ParseRoute(authZReq.RequestMethod, authZReq.RequestURI)
	if _, ok := authZReq.RequestHeaders[f.header]; !ok {
		// no restricted/unprivileged user header; all requests ok
		return &authorization.Response{
			Allow: true,
			Msg:   fmt.Sprintf("action '%s' allowed; all privileges OK", action),
		}
	}
	// header exists; limit to only "GET" (read) API access
	if authZReq.RequestMethod == "GET" {
		return &authorization.Response{
			Allow: true,
			Msg:   fmt.Sprintf("action '%s' allowed; privileges limited to read/GET operations", action),
		}
	}
	return &authorization.Response{
		Allow: false,
		Msg:   fmt.Sprintf("action '%s' not allowed due to unprivileged API access header", action),
	}
}

// AuthZRes always allow responses from server
func (f *headerAuthorizer) AuthZRes(authZReq *authorization.Request) *authorization.Response {
	return &authorization.Response{Allow: true}
}

// headerAuditor audit requset/response directly to standard output
type headerAuditor struct {
	logger   *logrus.Logger
	settings *HeaderAuditorSettings
}

// NewHeaderAuditor returns a new authz auditor that uses the specified logging hook (e.g., syslog or stdout)
func NewHeaderAuditor(settings *HeaderAuditorSettings) core.Auditor {
	b := &headerAuditor{settings: settings}
	return b
}

// HeaderAuditorSettings are settings used by the basic auditor
type HeaderAuditorSettings struct {
	LogHook string // LogHook is the log hook used to audit authorization data
	LogPath string // LogPath is the path to audit log file (if file hook is specified)
}

func (b *headerAuditor) AuditRequest(req *authorization.Request, pluginRes *authorization.Response) error {

	if req == nil {
		return fmt.Errorf("Authorization request is nil")
	}

	if pluginRes == nil {
		return fmt.Errorf("Authorization response is nil")
	}

	err := b.init()
	if err != nil {
		return err
	}
	// Default - file
	fields := logrus.Fields{
		"method": req.RequestMethod,
		"uri":    req.RequestURI,
		"user":   req.User,
		"allow":  pluginRes.Allow,
		"msg":    pluginRes.Msg,
	}

	if pluginRes != nil || pluginRes.Err != "" {
		fields["err"] = pluginRes.Err
	}

	b.logger.WithFields(fields).Info("Request")
	return nil
}

func (b *headerAuditor) AuditResponse(req *authorization.Request, pluginRes *authorization.Response) error {
	// Only log requests
	return nil
}

// init inits the auditor logger
func (b *headerAuditor) init() error {

	if b.settings == nil {
		return fmt.Errorf("Settings are not defined")
	}

	if b.logger != nil {
		return nil
	}

	b.logger = logrus.New()
	b.logger.Formatter = &logrus.JSONFormatter{}

	switch b.settings.LogHook {
	case AuditHookSyslog:
		{
			hook, err := logrus_syslog.NewSyslogHook("", "", syslog.LOG_ERR, "authz")
			if err != nil {
				return err
			}
			b.logger.Hooks.Add(hook)
		}
	case AuditHookFile:
		{
			logPath := b.settings.LogPath
			if logPath == "" {
				logrus.Infof("Using default log file path '%s'", logPath)
				logPath = defaultAuditLogPath
			}

			os.MkdirAll(path.Dir(logPath), 0700)
			f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0750)
			if err != nil {
				return err
			}
			b.logger.Out = f
		}
	case AuditHookStdout:
		{
			// Default - stdout
		}
	default:
		return fmt.Errorf("Wrong log hook value '%s'", b.settings.LogHook)
	}

	return nil
}
