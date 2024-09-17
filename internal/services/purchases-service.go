package services

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	repo "github.com/henrybravo/micro-report/internal/repositories"
	"github.com/henrybravo/micro-report/pkg/validate"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
)

type PurchaseService struct {
	PurchaseRepo *repo.PurchaseRepository
}

func (p *PurchaseService) RetrievePurchaseReport(
	_ context.Context,
	req *connect.Request[v1.RetrievePurchaseReportRequest],
) (*connect.Response[v1.RetrievePurchaseReportResponse], error) {
	businessId := req.Msg.GetBusinessId()
	period := req.Msg.GetPeriod()
	if !validate.IsValidUUID(businessId) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid business ID"))
	}
	if !validate.IsValidPeriod(period) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid period"))
	}
	purchases, err := p.PurchaseRepo.GetPurchasesByBusinessID(businessId, period)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp := connect.NewResponse(&v1.RetrievePurchaseReportResponse{
		Data: purchases,
	})
	return resp, nil
}
