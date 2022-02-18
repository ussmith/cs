package data

type ScanStatus struct {
	ScannerName string   `json:"scannerName"`
	Found       bool     `json:"found"`
	Viruses     []string `json:"viruses"`
	Err         error    `json:"error"`
}

type ScanPackage struct {
	Found      bool                `json:"found"`
	Viruses    map[string][]string `json:"viruses"`
	FailedJobs int                 `json:"failedJobs"`
}
