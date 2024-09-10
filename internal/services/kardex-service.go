package services

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	repo "github.com/henrybravo/micro-report/internal/repositories"
	"github.com/henrybravo/micro-report/pkg/validate"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
	"log"
)

type KardexServer struct {
	KardexRepo *repo.KardexRepository
}

func (k *KardexServer) RetrieveKardexValued(
	_ context.Context,
	req *connect.Request[v1.RetrieveKardexValuedRequest],
) (*connect.Response[v1.RetrieveKardexValuedResponse], error) {
	log.Println("Request headers: ", req.Header())
	localID := req.Msg.GetLocalId()
	isNotes := req.Msg.GetIncludeNotes()
	productID := req.Msg.GetProductId()
	startDate := req.Msg.GetStartDate()
	endDate := req.Msg.GetEndDate()
	period := req.Msg.GetPeriod()
	perPeriod := true
	if !validate.IsValidUUID(localID) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid business ID"))
	}
	if !validate.IsValidPeriod(period) {
		perPeriod = false
		if !validate.IsValidDate(startDate) || !validate.IsValidDate(endDate) {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid date, need period or dates to retrieve kardex"))
		}
	} else {
		startDate = period
		endDate = period
	}
	if !validate.IsValidUUID(productID) {
		productID = ""
	}
	kardex, err := k.KardexRepo.GetReportKardex(localID, startDate, endDate, productID, isNotes, perPeriod)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&v1.RetrieveKardexValuedResponse{
		Data: kardex,
	})
	res.Header().Set("Report-Version", "v1")
	return res, nil

}
