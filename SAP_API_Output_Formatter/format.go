package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-credit-mgmt-master-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToBusinessPartner(raw []byte, l *logger.Logger) ([]BusinessPartner, error) {
	pm := &responses.BusinessPartner{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to BusinessPartner. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	businessPartner := make([]BusinessPartner, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		businessPartner = append(businessPartner, BusinessPartner{
			BusinessPartner:                data.BusinessPartner,
			CrdtMgmtBusinessPartnerGroup:   data.CrdtMgmtBusinessPartnerGroup,
			CreditWorthinessScoreValue:     data.CreditWorthinessScoreValue,
			CrdtWrthnssScoreValdtyEndDate:  data.CrdtWrthnssScoreValdtyEndDate,
			CrdtWorthinessScoreLastChgDate: data.CrdtWorthinessScoreLastChgDate,
			CalcdCrdtWorthinessScoreValue:  data.CalcdCrdtWorthinessScoreValue,
			CreditRiskClass:                data.CreditRiskClass,
			CalculatedCreditRiskClass:      data.CalculatedCreditRiskClass,
			CreditRiskClassLastChangeDate:  data.CreditRiskClassLastChangeDate,
			CreditCheckRule:                data.CreditCheckRule,
			CreditScoreAndLimitCalcRule:    data.CreditScoreAndLimitCalcRule,
			ToCreditAccount:                data.ToCreditAccount.Deferred.URI,
		})
	}

	return businessPartner, nil
}

func ConvertToCreditAccount(raw []byte, l *logger.Logger) ([]CreditAccount, error) {
	pm := &responses.CreditAccount{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CreditAccount. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	creditAccount := make([]CreditAccount, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		creditAccount = append(creditAccount, CreditAccount{
			BusinessPartner:                data.BusinessPartner,
			CreditSegment:                  data.CreditSegment,
			BusinessPartnerIsCritical:      data.BusinessPartnerIsCritical,
			CreditAccountIsBlocked:         data.CreditAccountIsBlocked,
			CreditAccountBlockReason:       data.CreditAccountBlockReason,
			CreditAccountResubmissionDate:  data.CreditAccountResubmissionDate,
			CreditLimitAmount:              data.CreditLimitAmount,
			CreditLimitValidityEndDate:     data.CreditLimitValidityEndDate,
			CreditLimitLastChangeDate:      data.CreditLimitLastChangeDate,
			CreditLimitCalculatedAmount:    data.CreditLimitCalculatedAmount,
			CreditLimitIsZero:              data.CreditLimitIsZero,
			CreditLimitRequestedAmount:     data.CreditLimitRequestedAmount,
			CrdtLmtIsReqdFrmAutomCalc:      data.CrdtLmtIsReqdFrmAutomCalc,
			CreditLimitReqdValidityEndDate: data.CreditLimitReqdValidityEndDate,
			CreditLimitRequestDate:         data.CreditLimitRequestDate,
			CreditSegmentCurrency:          data.CreditSegmentCurrency,
		})
	}

	return creditAccount, nil
}

func ConvertToToCreditAccount(raw []byte, l *logger.Logger) ([]ToCreditAccount, error) {
	pm := &responses.ToCreditAccount{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToCreditAccount. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toCreditAccount := make([]ToCreditAccount, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toCreditAccount = append(toCreditAccount, ToCreditAccount{
			BusinessPartner:                data.BusinessPartner,
			CreditSegment:                  data.CreditSegment,
			BusinessPartnerIsCritical:      data.BusinessPartnerIsCritical,
			CreditAccountIsBlocked:         data.CreditAccountIsBlocked,
			CreditAccountBlockReason:       data.CreditAccountBlockReason,
			CreditAccountResubmissionDate:  data.CreditAccountResubmissionDate,
			CreditLimitAmount:              data.CreditLimitAmount,
			CreditLimitValidityEndDate:     data.CreditLimitValidityEndDate,
			CreditLimitLastChangeDate:      data.CreditLimitLastChangeDate,
			CreditLimitCalculatedAmount:    data.CreditLimitCalculatedAmount,
			CreditLimitIsZero:              data.CreditLimitIsZero,
			CreditLimitRequestedAmount:     data.CreditLimitRequestedAmount,
			CrdtLmtIsReqdFrmAutomCalc:      data.CrdtLmtIsReqdFrmAutomCalc,
			CreditLimitReqdValidityEndDate: data.CreditLimitReqdValidityEndDate,
			CreditLimitRequestDate:         data.CreditLimitRequestDate,
			CreditSegmentCurrency:          data.CreditSegmentCurrency,
		})
	}

	return toCreditAccount, nil
}
