package iso20022

// Information about a securities account and its characteristics.
type InvestmentAccount62 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Status of the account.
	AccountStatus *AccountStatus2 `xml:"AcctSts,omitempty"`

	// Specifies the account is blocked and other factors for the blocked account.
	BlockedStatus *BlockedStatusReason2Choice `xml:"BlckdSts,omitempty"`

	// Date the status is assigned.
	StatusDate *DateAndDateTime1Choice `xml:"StsDt,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Type of account.
	Type *AccountType2Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType2Choice `xml:"OwnrshTp,omitempty"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason2Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason2Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Specifies, for income on the fund or security that is to be reinvested, parameters for the reinvestment. If the reinvestment percentage is less than one hundred percent, the remaining percentage will be invested according to the investor’s or account owner's subsequent instructions.
	ReinvestmentDetails []*Reinvestment3 `xml:"RinvstmtDtls,omitempty"`

	// Method by which the tax (withholding tax) is to be processed, that is,  either withheld at source or tax information is reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod3Code `xml:"TaxWhldgMtd,omitempty"`

	// Details for the reporting of tax, for example, the country of taxation.
	TaxReporting []*TaxReporting2 `xml:"TaxRptg,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund or security associated to the account.
	FinancialInstrumentDetails []*FinancialInstrument56 `xml:"FinInstrmDtls,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`

	// Specifies the type of usage of the account.
	AccountUsageType *AccountUsageType2Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor or account owner signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`

	// Means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Specifies the category of the account.
	InvestmentAccountCategory *InvestmentAccountCategory1Choice `xml:"InvstmtAcctCtgy,omitempty"`

	// Specifies whether the holdings in the account are eligible for pledging.
	Pledging *Eligible1Code `xml:"Pldgg,omitempty"`

	// Specifies whether the holdings in the account are used as collateral.
	Collateral *Collateral1Code `xml:"Coll,omitempty"`

	// Details of third party rights.
	ThirdPartyRights *ThirdPartyRights1 `xml:"ThrdPtyRghts,omitempty"`

	// Functionality permitted to a third party in the actions that may be taken on an account on behalf of the investor.
	PowerOfAttorneyLevelOfControl *LevelOfControl1Choice `xml:"PwrOfAttnyLvlOfCtrl,omitempty"`

	// Specifies if the account is regarded as domestic or non-domestic for reporting purposes.
	AccountingStatus *AccountingStatus1Choice `xml:"AcctgSts,omitempty"`

	// Legal opening date for the account.
	OpeningDate *DateAndDateTimeChoice `xml:"OpngDt,omitempty"`

	// Legal closing date for the account.
	ClosingDate *DateAndDateTimeChoice `xml:"ClsgDt,omitempty"`

	// Indicates whether the account can hold a negative position in a security.
	NegativeIndicator *YesNoIndicator `xml:"NegInd,omitempty"`

	// Order in which securities are moved from the account.
	ProcessingOrder *PositionEffect3Code `xml:"PrcgOrdr,omitempty"`

	// Specifies whether the investor assumes responsibility for the liability.
	Liability *Liability1Choice `xml:"Lblty,omitempty"`

	// Information used to determine fees and types of operations that can be carried out on the account.
	InvestorProfile []*InvestorProfile1 `xml:"InvstrPrfl,omitempty"`

	// Fiscal year, when not the same as the calendar year.
	FiscalYear *FiscalYear1Choice `xml:"FsclYr,omitempty"`
}

func (i *InvestmentAccount62) SetIdentification(value string) {
	i.Identification = (*Max35Text)(&value)
}

func (i *InvestmentAccount62) AddAccountStatus() *AccountStatus2 {
	i.AccountStatus = new(AccountStatus2)
	return i.AccountStatus
}

func (i *InvestmentAccount62) AddBlockedStatus() *BlockedStatusReason2Choice {
	i.BlockedStatus = new(BlockedStatusReason2Choice)
	return i.BlockedStatus
}

func (i *InvestmentAccount62) AddStatusDate() *DateAndDateTime1Choice {
	i.StatusDate = new(DateAndDateTime1Choice)
	return i.StatusDate
}

func (i *InvestmentAccount62) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount62) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount62) AddType() *AccountType2Choice {
	i.Type = new(AccountType2Choice)
	return i.Type
}

func (i *InvestmentAccount62) AddOwnershipType() *OwnershipType2Choice {
	i.OwnershipType = new(OwnershipType2Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount62) AddTaxExemption() *TaxExemptionReason2Choice {
	i.TaxExemption = new(TaxExemptionReason2Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount62) AddStatementFrequency() *StatementFrequencyReason2Choice {
	i.StatementFrequency = new(StatementFrequencyReason2Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount62) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount62) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount62) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentAccount62) AddReinvestmentDetails() *Reinvestment3 {
	newValue := new(Reinvestment3)
	i.ReinvestmentDetails = append(i.ReinvestmentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount62) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod3Code)(&value)
}

func (i *InvestmentAccount62) AddTaxReporting() *TaxReporting2 {
	newValue := new(TaxReporting2)
	i.TaxReporting = append(i.TaxReporting, newValue)
	return newValue
}

func (i *InvestmentAccount62) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount62) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount62) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount62) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount62) AddFinancialInstrumentDetails() *FinancialInstrument56 {
	newValue := new(FinancialInstrument56)
	i.FinancialInstrumentDetails = append(i.FinancialInstrumentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount62) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount62) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount62) AddAccountUsageType() *AccountUsageType2Choice {
	i.AccountUsageType = new(AccountUsageType2Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount62) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount62) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

func (i *InvestmentAccount62) AddTransactionChannelType() *TransactionChannelType1Choice {
	i.TransactionChannelType = new(TransactionChannelType1Choice)
	return i.TransactionChannelType
}

func (i *InvestmentAccount62) AddInvestmentAccountCategory() *InvestmentAccountCategory1Choice {
	i.InvestmentAccountCategory = new(InvestmentAccountCategory1Choice)
	return i.InvestmentAccountCategory
}

func (i *InvestmentAccount62) SetPledging(value string) {
	i.Pledging = (*Eligible1Code)(&value)
}

func (i *InvestmentAccount62) SetCollateral(value string) {
	i.Collateral = (*Collateral1Code)(&value)
}

func (i *InvestmentAccount62) AddThirdPartyRights() *ThirdPartyRights1 {
	i.ThirdPartyRights = new(ThirdPartyRights1)
	return i.ThirdPartyRights
}

func (i *InvestmentAccount62) AddPowerOfAttorneyLevelOfControl() *LevelOfControl1Choice {
	i.PowerOfAttorneyLevelOfControl = new(LevelOfControl1Choice)
	return i.PowerOfAttorneyLevelOfControl
}

func (i *InvestmentAccount62) AddAccountingStatus() *AccountingStatus1Choice {
	i.AccountingStatus = new(AccountingStatus1Choice)
	return i.AccountingStatus
}

func (i *InvestmentAccount62) AddOpeningDate() *DateAndDateTimeChoice {
	i.OpeningDate = new(DateAndDateTimeChoice)
	return i.OpeningDate
}

func (i *InvestmentAccount62) AddClosingDate() *DateAndDateTimeChoice {
	i.ClosingDate = new(DateAndDateTimeChoice)
	return i.ClosingDate
}

func (i *InvestmentAccount62) SetNegativeIndicator(value string) {
	i.NegativeIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount62) SetProcessingOrder(value string) {
	i.ProcessingOrder = (*PositionEffect3Code)(&value)
}

func (i *InvestmentAccount62) AddLiability() *Liability1Choice {
	i.Liability = new(Liability1Choice)
	return i.Liability
}

func (i *InvestmentAccount62) AddInvestorProfile() *InvestorProfile1 {
	newValue := new(InvestorProfile1)
	i.InvestorProfile = append(i.InvestorProfile, newValue)
	return newValue
}

func (i *InvestmentAccount62) AddFiscalYear() *FiscalYear1Choice {
	i.FiscalYear = new(FiscalYear1Choice)
	return i.FiscalYear
}
