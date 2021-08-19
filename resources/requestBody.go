package resources

type ReqNewUser struct {
	UserPassphrase string `json:"userPassphrase"`
}

type ReqListObjects struct {
	UserAccessGrant string `json:"userAccessGrant"`
	UserPrefix string `json:"userPrefix"`
}

type ReqUploadObjectIdentity struct {
	UserAccessGrant string `json:"userAccessGrant"`
	ObjectKey string `json:"objectKey"`
	Data Identity `json:"identity"`
}

type ReqUploadObjectRecord struct {
	UserAccessGrant string `json:"userAccessGrant"`
	ObjectKey string `json:"objectKey"`
	Data Record `json:"record"`
}

type ReqDownloadObject struct {
	UserAccessGrant string `json:"userAccessGrant"`
	ObjectKey string `json:"objectKey"`
}

type ReqDownloadObjects struct {
	UserAccessGrant string `json:"userAccessGrant"`
	ObjectKeys []string `json:"objectKeys"`
}

