// File path: csp-report-api/pkg/domain/csp_report/service.go

package csp_report

import "context"

type ReportService struct {
	repository ReportRepository
}

func NewReportService(repository ReportRepository) *ReportService {
	return &ReportService{repository: repository}
}

func (s *ReportService) SaveReport(ctx context.Context, report *Report) error {
	return s.repository.Save(ctx, report)
}
