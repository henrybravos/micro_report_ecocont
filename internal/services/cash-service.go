package services

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	repo "github.com/henrybravo/micro-report/internal/repositories"
	"github.com/henrybravo/micro-report/pkg/validate"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
	"strings"
)

type CashServer struct {
	CashRepo *repo.CashRepository
}

func (c *CashServer) RetrieveCashBook(_ context.Context, req *connect.Request[v1.RetrieveCashBookRequest]) (*connect.Response[v1.RetrieveCashBookResponse], error) {
	businessID := req.Msg.GetBusinessId()
	period := req.Msg.GetPeriod()
	accountsIDs := req.Msg.GetAccountIds()
	if !validate.IsValidUUID(businessID) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid business ID"))
	}
	if !validate.IsValidPeriod(period) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid period"))
	}
	if len(accountsIDs) == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid account IDs"))
	}
	partsPeriod := strings.Split(period, "-")
	year := partsPeriod[0]
	month := partsPeriod[1]
	accounts, err := c.CashRepo.GetCashBalance(businessID, year, month, accountsIDs)
	fmt.Println(accounts)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	cashList, err := c.CashRepo.GetLCash(businessID, year, month, accountsIDs)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(&v1.RetrieveCashBookResponse{
		CashBooks:       cashList,
		AccountBalances: accounts,
	})
	res.Header().Set("Report-Version", "v1")
	return res, nil
}
