package iso20022

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson23 struct {

	// Term used to address a person.
	NamePrefix *NamePrefix1Choice `xml:"NmPrfx,omitempty"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm"`

	// Second name of a person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Additional information about a person that follows a person's name, for example, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Specifies the gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Language in which a person communicates.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt"`

	// Country where a person was born.
	CountryOfBirth *CountryCode `xml:"CtryOfBirth,omitempty"`

	// Province where a person was born.
	ProvinceOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty"`

	// City where a person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth,omitempty"`

	// Name of the occupation or job of a person.
	Profession *Max35Text `xml:"Prfssn,omitempty"`

	// Country of taxation of an individual person.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country and residential status of an individual, for example, non-permanent resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType1 `xml:"CtryAndResdtlSts,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress []*PostalAddress3 `xml:"PstlAdr"`

	// Nationality and legal status (minor or major) or rights that an individual may possess.
	Citizenship []*CitizenshipInformation `xml:"Ctznsh"`

	// Organisation represented by a person, or for which a person works.
	EmployingCompany *Max140Text `xml:"EmplngCpny,omitempty"`

	// Title of the function.
	BusinessFunction *Max35Text `xml:"BizFctn,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification55 `xml:"OthrId,omitempty"`

	// Additional regulatory information about the investor that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`

	// Specifies if due diligence checks on the political exposure of the investor have been carried out and whether these checks are national or foreign. (A politically exposed person is someone who has been entrusted with a prominent public function, or an individual who is closely related to such a person.)
	PoliticallyExposedPersonType *PoliticalExposureType1Choice `xml:"PltclyXpsdPrsnTp,omitempty"`
}

func (i *IndividualPerson23) AddNamePrefix() *NamePrefix1Choice {
	i.NamePrefix = new(NamePrefix1Choice)
	return i.NamePrefix
}

func (i *IndividualPerson23) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson23) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson23) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson23) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson23) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson23) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *IndividualPerson23) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson23) SetCountryOfBirth(value string) {
	i.CountryOfBirth = (*CountryCode)(&value)
}

func (i *IndividualPerson23) SetProvinceOfBirth(value string) {
	i.ProvinceOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson23) SetCityOfBirth(value string) {
	i.CityOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson23) SetProfession(value string) {
	i.Profession = (*Max35Text)(&value)
}

func (i *IndividualPerson23) SetTaxationCountry(value string) {
	i.TaxationCountry = (*CountryCode)(&value)
}

func (i *IndividualPerson23) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType1 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType1)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson23) AddPostalAddress() *PostalAddress3 {
	newValue := new(PostalAddress3)
	i.PostalAddress = append(i.PostalAddress, newValue)
	return newValue
}

func (i *IndividualPerson23) AddCitizenship() *CitizenshipInformation {
	newValue := new(CitizenshipInformation)
	i.Citizenship = append(i.Citizenship, newValue)
	return newValue
}

func (i *IndividualPerson23) SetEmployingCompany(value string) {
	i.EmployingCompany = (*Max140Text)(&value)
}

func (i *IndividualPerson23) SetBusinessFunction(value string) {
	i.BusinessFunction = (*Max35Text)(&value)
}

func (i *IndividualPerson23) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	i.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return i.PrimaryCommunicationAddress
}

func (i *IndividualPerson23) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	i.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return i.SecondaryCommunicationAddress
}

func (i *IndividualPerson23) AddOtherIdentification() *GenericIdentification55 {
	newValue := new(GenericIdentification55)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}

func (i *IndividualPerson23) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	i.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return i.AdditionalRegulatoryInformation
}

func (i *IndividualPerson23) AddPoliticallyExposedPersonType() *PoliticalExposureType1Choice {
	i.PoliticallyExposedPersonType = new(PoliticalExposureType1Choice)
	return i.PoliticallyExposedPersonType
}
