package responses

type BusinessPartner struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
			ToCreditAccount                struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_CreditMgmtAccountTP"`
		} `json:"results"`
	} `json:"d"`
}
