// File path: pkg/infrastructure/persistence/mongodb/csp_report_repository.go

package mongodb

import (
	"context"
	"github.com/achimgrolimund/csp-report-api/pkg/domain/csp_report"
	"go.mongodb.org/mongo-driver/mongo"
)

type CSPReportRepository struct {
	collection *mongo.Collection
}

func NewCSPReportRepository(client *mongo.Client, dbName, collectionName string) *CSPReportRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &CSPReportRepository{collection: collection}
}

func (r *CSPReportRepository) Save(ctx context.Context, report *csp_report.Report) error {
	_, err := r.collection.InsertOne(ctx, report)
	return err
}
