package iso20022

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount55 struct {

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification70Choice `xml:"OwnrId,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Form, that is, ownership, of the security, that is, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Indicates whether a security exists only as an electronic record, that is, there is no physical document representing the security.
	DematerialisedIndicator *YesNoIndicator `xml:"DmtrlsdInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Beneficial owner or its designated agent certifies that it complies with any holding or investment restrictions or requirements of the fund.
	BeneficiaryCertificationCompletion *BeneficiaryCertificationCompletion1Code `xml:"BnfcryCertfctnCmpltn,omitempty"`

	// Institution that maintains the records where the account is held.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount5 `xml:"SubAcctDtls,omitempty"`
}

func (i *InvestmentAccount55) AddOwnerIdentification() *PartyIdentification70Choice {
	newValue := new(PartyIdentification70Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount55) SetAccountIdentification(value string) {
	i.AccountIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccount55) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount55) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount55) SetSecuritiesForm(value string) {
	i.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentAccount55) SetDematerialisedIndicator(value string) {
	i.DematerialisedIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount55) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentAccount55) SetBeneficiaryCertificationCompletion(value string) {
	i.BeneficiaryCertificationCompletion = (*BeneficiaryCertificationCompletion1Code)(&value)
}

func (i *InvestmentAccount55) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount55) AddSubAccountDetails() *SubAccount5 {
	i.SubAccountDetails = new(SubAccount5)
	return i.SubAccountDetails
}
