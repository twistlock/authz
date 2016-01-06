package authz

import (
	"github.com/docker/docker/pkg/authorization"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestPolicyApply(t *testing.T) {

	policy := `{"name":"policy_1","users":["user_1","user_2"],"actions":["container_create","docker_version"]}
	           {"name":"policy_2","users":["user_3","user_4"],"actions":["container_create","container_exec"]}
	           {"name":"policy_3","users":["user_5"],"actions":["container"]}
	           {"name":"policy_4","users":["user_6"],"actions":["container"], "readonly":true }` // User can do anything with containers

	const policyFileName = "/tmp/policy.json"
	err := ioutil.WriteFile(policyFileName, []byte(policy), 0755)
	assert.NoError(t, err)

	tests := []struct {
		method         string
		uri            string
		user           string // user is the user in the request
		allow          bool   // allow is the allow/deny response from the policy plugin
		expectedPolicy string // expectedPolicy is the expected policy name that should appear in the message
	}{
		{"GET", "/v1.21/version", "user_1", true, "policy_1"},                // User and command allowed
		{"GET", "/v1.21/version", "user_3", false, "policy_2"},               // User and command not allowed
		{"GET", "/v1.21/version", "user_5", false, ""},                       // Non existing user (no policy found)
		{"GET", "/v1.21/containers/id/json", "user_5", true, "policy_3"},     // All containers action allowed
		{"GET", "/v1.21/containers/id/json", "user_6", true, "policy_4"},     // Readonly policy - GET allowed
		{"POST", "/v1.21/containers/id/rename", "user_6", false, "policy_4"}, // Readonly policy - GET denied
	}

	handler := NewBasicAuthZHandler(&BasicAuthorizerSettings{PolicyPath: policyFileName})

	assert.NoError(t, handler.Init(), "Initialization must be succesfull")

	for _, test := range tests {
		res := handler.AuthZReq(&authorization.Request{RequestMethod: test.method, RequestURI: test.uri, User: test.user})
		assert.Equal(t, res.Allow, test.allow, "Request must be allowed/denied based on policy")
		assert.Contains(t, res.Msg, test.expectedPolicy, "Policy name must appear in the response")
	}
}
