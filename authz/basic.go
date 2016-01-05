package authz

import (
	"authz/core"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/pkg/authorization"
	"github.com/docker/docker/vendor/src/github.com/Sirupsen/logrus"
	"github.com/howeyc/fsnotify"
	"io/ioutil"
	"path"
	"strings"
)

// BasicPolicy represent a single policy object that is evaluated in the authorization flow.
// Each policy object consists of multiple users and docker actions, where each user belongs to a single policy.
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
	Actions []string `json:"actions"` // Actions are the docker actions (mapped to authz terminology) that are allowed according to this policy
	Users   []string `json:"users"`   // Users are the users for which this policy apply to
	Name    string   `json:"name"`    // Name is the policy name
}

type BasicAuthorizer struct {
	settings *BasicAuthorizerSettings
	policies []BasicPolicy
}

type BasicAuthorizerSettings struct {
	PolicyPath string
}

// NewBasicAuthZHandler creates a new basic authz handler
func NewBasicAuthZHandler(settings *BasicAuthorizerSettings) core.Authorizer {
	return &BasicAuthorizer{settings: settings}
}

// Init loads the basic authz plugin configuration from disk
func (f *BasicAuthorizer) Init() error {

	err := f.loadPolicies()
	if err != nil {
		return err
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if ev.IsModify() {
					err := f.loadPolicies()
					if err != nil {
						logrus.Errorf("Error refreshing policy %q", err.Error())
					}
				}
			case err := <-watcher.Error:
				logrus.Errorf("Settings watcher error '%v'", err)
			}
		}
	}()

	err = watcher.Watch(f.settings.PolicyPath)
	if err != nil {
		// Silently ignore watching error
		logrus.Errorf("Failed to start watching folder %q", err.Error())
	}

	return nil
}

func (f *BasicAuthorizer) loadPolicies() error {
	data, err := ioutil.ReadFile(path.Join(f.settings.PolicyPath))

	if err != nil {
		return err
	}

	var policies []BasicPolicy
	for _, l := range strings.Split(string(data), "\n") {

		if l == "" {
			continue
		}

		var policy BasicPolicy
		err := json.Unmarshal([]byte(l), &policy)
		if err != nil {
			logrus.Errorf("Failed to unmarshel policy entry %q %q", l, err.Error())
		}
		policies = append(policies, policy)
	}
	logrus.Infof("Loaded '%d' policies", len(policies))

	// Notify when user appears in duplicate policies
	userToPolicy := make(map[string]string)
	for _, policy := range policies {
		for _, u := range policy.Users {
			if userPolicy, ok := userToPolicy[u]; ok {
				logrus.Warnf("User %q already appears in policy %q. Only single policy applies. Undefined policy behavior %q",
					u,
					userPolicy,
					policy.Name)
			}
			userToPolicy[u] = policy.Name
		}

	}

	f.policies = policies
	return nil
}

func (f *BasicAuthorizer) AuthZReq(authZReq *authorization.Request) *authorization.Response {

	logrus.Debugf("Received AuthZ request, method: '%s', url: '%s'", authZReq.RequestMethod, authZReq.RequestURI)

	action := core.ParseRoute(authZReq.RequestMethod, authZReq.RequestURI)

	for _, policy := range f.policies {
		for _, user := range policy.Users {
			if user == authZReq.User {
				for _, policyAction := range policy.Actions {
					if policyAction == action {
						return &authorization.Response{
							Allow: true,
							Msg:   fmt.Sprintf("action '%s' allowed for user '%s' by policy '%s'", action, authZReq.User, policy.Name),
						}
					}
				}
				return &authorization.Response{
					Allow: false,
					Msg:   fmt.Sprintf("action '%s' denied for user '%s' by policy '%s'", action, authZReq.User, policy.Name),
				}
			}
		}
	}

	return &authorization.Response{
		Allow: false,
		Msg:   fmt.Sprintf("no policy applied (user: '%s' action: '%s')", authZReq.User, action),
	}
}

// AuthZRes always allow responses from server
func (f *BasicAuthorizer) AuthZRes(authZReq *authorization.Request) *authorization.Response {
	return &authorization.Response{Allow: true}
}

// basicAuditor audit requset/response directly to standard output
type basicAuditor struct {
}

func NewBasicAuditor() core.Auditor {
	return &basicAuditor{}
}

// TODO: Check nil
func (f *basicAuditor) AuditRequest(req *authorization.Request, pluginRes *authorization.Response) {
	logrus.Debugf("Request: method:%s' uri:'%s' user:'%s' allow:'%t' plugin msg:'%s'", req.RequestMethod, req.RequestURI, req.User, pluginRes.Allow, pluginRes.Msg)
}

func (f *basicAuditor) AuditResponse(req *authorization.Request, pluginRes *authorization.Response) {
	logrus.Debugf("Response: method:%s' uri:'%s' user:'%s' daeomn status code '%d' allow:'%t' plugin msg:'%s'",
		req.RequestMethod,
		req.RequestURI,
		req.User,
		req.ResponseStatusCode,
		pluginRes.Allow,
		pluginRes.Msg)
}
