package resources

type ResNewUser struct {
	UserAccessGrant string `json:"userAccessGrant"`
	UserSalt string `json:"userSalt"`
}

type ResListObjects struct {
	Objects []string `json:"objects"`
}

