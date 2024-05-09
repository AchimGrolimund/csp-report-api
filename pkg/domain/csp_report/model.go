// File path: csp-report-api/pkg/domain/csp_report/model.go

package csp_report

type Report struct {
	Report CSPReport `json:"csp-report"`
}

type CSPReport struct {
	DocumentURI        string `json:"document-uri"`
	Referrer           string `json:"referrer"`
	ViolatedDirective  string `json:"violated-directive"`
	EffectiveDirective string `json:"effective-directive"`
	OriginalPolicy     string `json:"original-policy"`
	Disposition        string `json:"disposition"`
	BlockedURI         string `json:"blocked-uri"`
	LineNumber         int    `json:"line-number"`
	SourceFile         string `json:"source-file"`
	StatusCode         int    `json:"status-code"`
	ScriptSample       string `json:"script-sample"`
	ClientIP           string `json:"clientIP"`
	UserAgent          string `json:"userAgent"`
}
