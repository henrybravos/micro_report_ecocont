package services

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	repo "github.com/henrybravo/micro-report/internal/repositories"
	"github.com/henrybravo/micro-report/pkg/validate"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
	"strings"
)

type BookServer struct {
	BookRepo *repo.BankBookRepository
}

func (b *BookServer) RetrieveBankBook(_ context.Context, req *connect.Request[v1.RetrieveBankBookRequest]) (*connect.Response[v1.RetrieveBankBookResponse], error) {

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
	balances, err := b.BookRepo.GetLBankBalance(businessID, year, month, accountsIDs)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	accounts, err := b.BookRepo.GetLBanks(businessID, year, month, accountsIDs)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&v1.RetrieveBankBookResponse{
		BankBooks:    accounts,
		BankBalances: balances,
	})
	return res, nil
}
