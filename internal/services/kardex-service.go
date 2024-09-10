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
	period := req.Msg.GetPeriod()
	if !validate.IsValidUUID(localID) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid business ID"))
	}
	if !validate.IsValidPeriod(period) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid period"))
	}

	kardex, err := k.KardexRepo.GetReportKardex(localID, period, period, "", isNotes, true)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&v1.RetrieveKardexValuedResponse{
		Data: kardex,
	})
	res.Header().Set("Report-Version", "v1")
	return res, nil

}
