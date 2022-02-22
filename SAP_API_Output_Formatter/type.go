package sap_api_output_formatter

type CreditManagementMaster struct {
	ConnectionKey    string `json:"connection_key"`
	Result           bool   `json:"result"`
	RedisKey         string `json:"redis_key"`
	Filepath         string `json:"filepath"`
	APISchema        string `json:"api_schema"`
	BusinessPartner  string `json:"business_partner"`
	Deleted          bool   `json:"deleted"`
}

type BusinessPartner struct {
	BusinessPartner                string      `json:"BusinessPartner"`
	CrdtMgmtBusinessPartnerGroup   string      `json:"CrdtMgmtBusinessPartnerGroup"`
	CreditWorthinessScoreValue     string      `json:"CreditWorthinessScoreValue"`
	CrdtWrthnssScoreValdtyEndDate  string      `json:"CrdtWrthnssScoreValdtyEndDate"`
	CrdtWorthinessScoreLastChgDate string      `json:"CrdtWorthinessScoreLastChgDate"`
	CalcdCrdtWorthinessScoreValue  string      `json:"CalcdCrdtWorthinessScoreValue"`
	CreditRiskClass                string      `json:"CreditRiskClass"`
	CalculatedCreditRiskClass      string      `json:"CalculatedCreditRiskClass"`
	CreditRiskClassLastChangeDate  string      `json:"CreditRiskClassLastChangeDate"`
	CreditCheckRule                string      `json:"CreditCheckRule"`
	CreditScoreAndLimitCalcRule    string      `json:"CreditScoreAndLimitCalcRule"`
	ToCreditAccount                string      `json:"to_CreditMgmtAccountTP"`
}

type CreditAccount struct {
	BusinessPartner                string      `json:"BusinessPartner"`
	CreditSegment                  string      `json:"CreditSegment"`
	BusinessPartnerIsCritical      bool        `json:"BusinessPartnerIsCritical"`
	CreditAccountIsBlocked         bool        `json:"CreditAccountIsBlocked"`
	CreditAccountBlockReason       string      `json:"CreditAccountBlockReason"`
	CreditAccountResubmissionDate  string      `json:"CreditAccountResubmissionDate"`
	CreditLimitAmount              string      `json:"CreditLimitAmount"`
	CreditLimitValidityEndDate     string      `json:"CreditLimitValidityEndDate"`
	CreditLimitLastChangeDate      string      `json:"CreditLimitLastChangeDate"`
	CreditLimitCalculatedAmount    string      `json:"CreditLimitCalculatedAmount"`
	CreditLimitIsZero              bool        `json:"CreditLimitIsZero"`
	CreditLimitRequestedAmount     string      `json:"CreditLimitRequestedAmount"`
	CrdtLmtIsReqdFrmAutomCalc      bool        `json:"CrdtLmtIsReqdFrmAutomCalc"`
	CreditLimitReqdValidityEndDate string      `json:"CreditLimitReqdValidityEndDate"`
	CreditLimitRequestDate         string      `json:"CreditLimitRequestDate"`
	CreditSegmentCurrency          string      `json:"CreditSegmentCurrency"`
}
    
type ToCreditAccount struct {
	BusinessPartner                string      `json:"BusinessPartner"`
	CreditSegment                  string      `json:"CreditSegment"`
	BusinessPartnerIsCritical      bool        `json:"BusinessPartnerIsCritical"`
	CreditAccountIsBlocked         bool        `json:"CreditAccountIsBlocked"`
	CreditAccountBlockReason       string      `json:"CreditAccountBlockReason"`
	CreditAccountResubmissionDate  string      `json:"CreditAccountResubmissionDate"`
	CreditLimitAmount              string      `json:"CreditLimitAmount"`
	CreditLimitValidityEndDate     string      `json:"CreditLimitValidityEndDate"`
	CreditLimitLastChangeDate      string      `json:"CreditLimitLastChangeDate"`
	CreditLimitCalculatedAmount    string      `json:"CreditLimitCalculatedAmount"`
	CreditLimitIsZero              bool        `json:"CreditLimitIsZero"`
	CreditLimitRequestedAmount     string      `json:"CreditLimitRequestedAmount"`
	CrdtLmtIsReqdFrmAutomCalc      bool        `json:"CrdtLmtIsReqdFrmAutomCalc"`
	CreditLimitReqdValidityEndDate string      `json:"CreditLimitReqdValidityEndDate"`
	CreditLimitRequestDate         string      `json:"CreditLimitRequestDate"`
	CreditSegmentCurrency          string      `json:"CreditSegmentCurrency"`
}
