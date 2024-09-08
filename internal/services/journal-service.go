package services

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	repo "github.com/henrybravo/micro-report/internal/repositories"
	"github.com/henrybravo/micro-report/pkg/utils"
	"github.com/henrybravo/micro-report/pkg/validate"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
	"log"
	"strconv"
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

func (s *JournalServer) RetrieveGeneralJournal(_ context.Context, req *connect.Request[v1.RetrieveGeneralJournalRequest]) (*connect.Response[v1.RetrieveGeneralJournalResponse], error) {
	businessID := req.Msg.GetBusinessId()
	period := req.Msg.GetPeriod()
	if !validate.IsValidUUID(businessID) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid business ID"))
	}
	if !validate.IsValidPeriod(period) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid period"))
	}
	partsPeriod := strings.Split(period, "-")
	year := partsPeriod[0]
	month := partsPeriod[1]
	journalEntries, err := s.JournalRepo.GetLfJournals(businessID, year, month, false, false, false, "060100")
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	var journals []*v1.GeneralJournal

	for _, element := range journalEntries {
		cuo := element["i2"]
		observaciones := element["observaciones"]
		if observaciones == "" || observaciones == "null" {
			observaciones = ""
		}
		newCuo := strings.TrimSuffix(cuo.(string), "-")
		cuoParts := strings.Split(cuo.(string), "-")
		cuoSuffix := ""
		if len(cuoParts) > 1 {
			cuoSuffix = "-" + cuoParts[len(cuoParts)-1]
		}
		operacion := element["i15"]
		if strings.Contains(cuoSuffix, "-P") || strings.Contains(newCuo, "-DC") {
			operacion = element["i15pg"]
		}

		descripcion := element["i16o"]
		if strings.Contains(cuoSuffix, "-D") && !strings.Contains(newCuo, "-DC") {
			descripcion = element["i16"]
		}

		tipoCambio, _ := strconv.ParseFloat(fmt.Sprint(element["p_tipo_cambio"]), 64)
		if tipoCambio <= 1 {
			tipoCambio, _ = strconv.ParseFloat(fmt.Sprint(element["tipo_cambio"]), 64)
		}

		debe := element["i18"].(float64) * tipoCambio
		haber := element["i19"].(float64) * tipoCambio

		journals = append(journals, &v1.GeneralJournal{
			Id:           element["id"].(string),
			Cuo:          newCuo,
			Operacion:    utils.FormatYYYY_MM_DD(operacion.(string)),
			Descripcion:  descripcion.(string),
			Cuenta:       element["i4"].(string),
			Denominacion: element["denominacion"].(string),
			Debe:         float32(debe),
			Haber:        float32(haber),
		})
	}

	res := connect.NewResponse(&v1.RetrieveGeneralJournalResponse{
		GeneralJournals: journals,
	})
	res.Header().Set("Report-Version", "v1")
	return res, nil
}
