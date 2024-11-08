package models

type AuthData struct {
	OrgPID    string `json:"org_pid"`
	OrgAccPID string `json:"org_account_pid"`
	IsSandbox bool   `json:"is_sandbox"`
}
