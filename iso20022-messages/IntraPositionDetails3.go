package iso20022

// Details of the intra-position movement.
type IntraPositionDetails3 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Balance from which the securities were moved.
	BalanceFrom *SecuritiesBalanceType3Choice `xml:"BalFr"`

	// Intra-position movement(s) having been performed.
	IntraPositionMovement []*IntraPositionMovementDetails1 `xml:"IntraPosMvmnt"`
}

func (i *IntraPositionDetails3) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	i.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionDetails3) AddBalanceFrom() *SecuritiesBalanceType3Choice {
	i.BalanceFrom = new(SecuritiesBalanceType3Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails3) AddIntraPositionMovement() *IntraPositionMovementDetails1 {
	newValue := new(IntraPositionMovementDetails1)
	i.IntraPositionMovement = append(i.IntraPositionMovement, newValue)
	return newValue
}
