// File path: csp-report-api/pkg/domain/csp_report/repository.go

package csp_report

import "context"

type ReportRepository interface {
	Save(ctx context.Context, report *Report) error
}
