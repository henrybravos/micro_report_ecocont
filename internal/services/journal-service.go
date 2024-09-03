package services

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	repo "github.com/henrybravo/micro-report/internal/repositories"
	"github.com/henrybravo/micro-report/pkg/validate"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
	"log"
	"strings"
)

type JournalServer struct {
	JournalRepo *repo.JournalRepository
}

func (s *JournalServer) RetrieveJournalReport(
	_ context.Context,
	req *connect.Request[v1.RetrieveJournalReportRequest],
) (*connect.Response[v1.RetrieveJournalReportResponse], error) {
	log.Println("Request headers: ", req.Header())
	businessID := req.Msg.GetBusinessId()
	period := req.Msg.GetPeriod()
	isConsolidated := req.Msg.GetIsConsolidated()
	includeClose := req.Msg.GetIncludeClose()
	includeCuBa := req.Msg.GetIncludeCuBa()
	if !validate.IsValidUUID(businessID) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid business ID"))
	}
	if !validate.IsValidPeriod(period) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid period"))
	}
	partsPeriod := strings.Split(period, "-")
	year := partsPeriod[0]
	month := partsPeriod[1]
	journalEntries, err := s.JournalRepo.GetJournalEntries(businessID, year, month, isConsolidated, includeClose, includeCuBa)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&v1.RetrieveJournalReportResponse{
		Journals: journalEntries,
	})
	res.Header().Set("Report-Version", "v1")
	return res, nil
}
