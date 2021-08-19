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
	Identity Identity `json:"identity"`
}

type ReqUploadObjectRecord struct {
	UserAccessGrant string `json:"userAccessGrant"`
	ObjectKey string `json:"objectKey"`
	Record Record `json:"record"`
}

type ReqUploadObjectRecords struct {
	UserAccessGrant string `json:"userAccessGrant"`
	ObjectKey string `json:"objectKey"`
	Record []Record `json:"record"`
}

type ReqDownloadObject struct {
	UserAccessGrant string `json:"userAccessGrant"`
	ObjectKey string `json:"objectKey"`
}

type ReqDownloadObjects struct {
	UserAccessGrant string `json:"userAccessGrant"`
	ObjectKey []string `json:"objectKey"`
}

