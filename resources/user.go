package resources

type UserAccessRequest struct {
	UserPassphrase string `json: "userPassPhrase"`
}

type GeneratedUser struct {
	UserAccessGrant string `json:"userAccessGrant"`
	UserSalt        []byte `json:"userSalt"`
}

type ObjectsReq struct {
	UserAccessGrant string `json:"userAccessGrant"`
	UserPrefix string `json:"userPrefix"`
	ObjectKey []string `json:"objectKey"`
	AccessToRevoke string `json:"accessToRevoke"`
	Identity Identity `json:"identity"`
	Data []Record `json:"records"`
}

type ListData struct {
	Objects []string `json: "objects"`
}

type Record struct {
	PatientName string `json:"patient name"`
	RecordName string `json: "record name"`
	Provider string `json: "provider"`
	Date string `json: "date"`
	PatientAddress string `json: "patient address"`
	DoctorAddress string `json: "doctor address"`
	DoctorNote string `json: "doctor note"`
}


type Identity struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

