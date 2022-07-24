package iso20022

// Describes the type of product and the assets to be transferred.
type ISATransfer16 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Identification of the confirmation assigned by the transferor to the transfer.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId,omitempty"`

	// Requested date at which the assets should be transferred.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio1Choice `xml:"Prtfl,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument34 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer16) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer16) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer16) SetTransferConfirmationIdentification(value string) {
	i.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer16) AddRequestedTransferDate() *DateFormat1Choice {
	i.RequestedTransferDate = new(DateFormat1Choice)
	return i.RequestedTransferDate
}

func (i *ISATransfer16) AddPortfolio() *ISAPortfolio1Choice {
	i.Portfolio = new(ISAPortfolio1Choice)
	return i.Portfolio
}

func (i *ISATransfer16) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer16) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer16) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument34 {
	newValue := new(FinancialInstrument34)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}
