package ram

import (
	"encoding/json"
	//	"fmt"
	"testing"
)

var (
	policy_username string
	policy_name     string
	policy_document = PolicyDocument{
		Statement: []PolicyItem{
			PolicyItem{
				Action:   "*",
				Effect:   "Allow",
				Resource: "*",
			},
		},
		Version: "1",
	}
	policy_req = PolicyRequest{
		PolicyName:  "unit_tesst_policy",
		Description: "nothing",
		PolicyType:  "Custom",
	}
)

/*
	TODO maybe I need import ginko to my project
	BeforeEach and AfterClean is needed
*/

func TestCreatePolicy(t *testing.T) {
	var policyReq = policy_req
	document, err := json.Marshal(policy_document)
	if err != nil {
		t.Errorf("Failed to marshal document %v", err)
	}
	policyReq.PolicyDocument = string(document)
	client := NewTestClient()
	resp, err := client.CreatePolicy(policyReq)
	if err != nil {
		t.Errorf("Failed to CreatePolicy %v", err)
	}
	policy_name = resp.Policy.PolicyName
	t.Logf("pass CreatePolicy %+++v", resp)
}

func TestGetPolicy(t *testing.T) {
	var policyReq = policy_req
	policyReq.PolicyName = policy_name
	client := NewTestClient()
	resp, err := client.GetPolicy(policyReq)
	if err != nil {
		t.Errorf("Failed to GetPolicy %v", err)
	}
	t.Logf("pass GetPolicy %v", resp)
}

func TestAttachPolicyToUser(t *testing.T) {
	client := NewTestClient()
	listParams := ListUserRequest{}
	resp, err := client.ListUsers(listParams)
	if err != nil {
		t.Errorf("Failed to ListUser %v", err)
		return
	}
	policy_username = resp.Users.User[0].UserName
	attachPolicyRequest := AttachPolicyRequest{
		PolicyRequest: PolicyRequest{
			PolicyType: "Custom",
			PolicyName: policy_name,
		},
		UserName: policy_username,
	}
	attachResp, err := client.AttachPolicyToUser(attachPolicyRequest)
	if err != nil {
		t.Errorf("Failed to AttachPolicyToUser %v", err)
		return
	}
	t.Logf("pass AttachPolicyToUser %++v", attachResp)
}

func TestListPoliciesForUser(t *testing.T) {
	client := NewTestClient()
	userQuery := UserQueryRequest{
		UserName: policy_username,
	}
	resp, err := client.ListPoliciesForUser(userQuery)
	if err != nil {
		t.Errorf("Failed to ListPoliciesForUser %v", err)
		return
	}
	t.Logf("pass ListPoliciesForUser %++v", resp)
}

func TestDetachPolicyFromUser(t *testing.T) {
	client := NewTestClient()
	detachPolicyRequest := AttachPolicyRequest{
		PolicyRequest: PolicyRequest{
			PolicyType: "Custom",
			PolicyName: policy_name,
		},
		UserName: policy_username,
	}
	resp, err := client.DetachPolicyFromUser(detachPolicyRequest)
	if err != nil {
		t.Errorf("Failed to DetachPolicyFromUser %++v", err)
		return
	}
	t.Logf("pass DetachPolicyFromUser %++v", resp)
}

func TestDeletePolicy(t *testing.T) {
	client := NewTestClient()
	policyReq := policy_req
	resp, err := client.DeletePolicy(policyReq)
	if err != nil {
		t.Errorf("Failed to DeletePolicy %v", err)
		return
	}
	t.Logf("pass DeletePolicy %++v", resp)
}