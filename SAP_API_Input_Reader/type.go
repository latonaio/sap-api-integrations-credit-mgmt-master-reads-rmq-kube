package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string      `json:"connection_key"`
	Result        bool        `json:"result"`
	RedisKey      string      `json:"redis_key"`
	Filepath      string      `json:"filepath"`
	Document   struct {
		DocumentNo                     string      `json:"document_no"`
		DeliverTo                      string      `json:"deliver_to"`
		Quantity                       string      `json:"quantity"`
		PickedQuantity                 string      `json:"picked_quantity"`
		Price                          string      `json:"price"`
	    Batch                          string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             string      `json:"quantity"`
		CompletedQuantity    string      `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 string      `json:"quantity"`
			CompletedQuantity        string      `json:"completed_quantity"`
			ErroredQuantity          string      `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity string      `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema               string      `json:"api_schema"`
	MaterialCode            string      `json:"material_code"`
	Plant_Supplier          string      `json:"plant/supplier"`
	Stock                   string      `json:"stock"`
	DocumentType            string      `json:"document_type"`
	DocumentNo              string      `json:"document_no"`
	PlannedDate             string      `json:"planned_date"`
	ValidatedDate           string      `json:"validated_date"`
	Deleted                 bool        `json:"deleted"`
}

type SDC struct {
	ConnectionKey          string `json:"connection_key"`
	Result                 bool   `json:"result"`
	RedisKey               string `json:"redis_key"`
	Filepath               string `json:"filepath"`
	CreditManagementMaster struct {
		BusinessPartner                string `json:"BusinessPartner"`
		CrdtMgmtBusinessPartnerGroup   string `json:"CrdtMgmtBusinessPartnerGroup"`
		CreditWorthinessScoreValue     string `json:"CreditWorthinessScoreValue"`
		CrdtWrthnssScoreValdtyEndDate  string `json:"CrdtWrthnssScoreValdtyEndDate"`
		CrdtWorthinessScoreLastChgDate string `json:"CrdtWorthinessScoreLastChgDate"`
		CalcdCrdtWorthinessScoreValue  string `json:"CalcdCrdtWorthinessScoreValue"`
		CreditRiskClass                string `json:"CreditRiskClass"`
		CalculatedCreditRiskClass      string `json:"CalculatedCreditRiskClass"`
		CreditRiskClassLastChangeDate  string `json:"CreditRiskClassLastChangeDate"`
		CreditCheckRule                string `json:"CreditCheckRule"`
		CreditScoreAndLimitCalcRule    string `json:"CreditScoreAndLimitCalcRule"`
		CreditAccount                  struct {
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
		} `json:"CreditAccount"`
	} `json:"CreditManagementMaster"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	BusinessPartnerCode string   `json:"business_partner_code"`
	Deleted             bool     `json:"deleted"`
}
