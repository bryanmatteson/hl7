package hl7v28

import hl7 "hl7"

type UAC struct {
	UserAuthenticationCredentialTypeCode CWE
	UserAuthenticationCredential         ED
}

// NewUAC creates an implementation of UAC
func NewUAC(input hl7.HL7Type) UAC {
	v := UAC{}
	v.UserAuthenticationCredentialTypeCode = NewCWE(input.Index(1).Index(0))
	v.UserAuthenticationCredential = NewED(input.Index(2).Index(0))
	return v
}

// NewUACSlice will construct a slice of type UAC
func NewUACSlice(input hl7.HL7Type) []UAC {
	values := make([]UAC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewUAC(input.Index(i))
	}
	return values
}

type LCH struct {
	PrimaryKeyValueLch             PL
	SegmentActionCode              ID
	SegmentUniqueKey               EI
	LocationCharacteristicID       CWE
	LocationCharacteristicValueLch CWE
}

// NewLCH creates an implementation of LCH
func NewLCH(input hl7.HL7Type) LCH {
	v := LCH{}
	v.PrimaryKeyValueLch = NewPL(input.Index(1).Index(0))
	v.SegmentActionCode = NewID(input.Index(2).Index(0))
	v.SegmentUniqueKey = NewEI(input.Index(3).Index(0))
	v.LocationCharacteristicID = NewCWE(input.Index(4).Index(0))
	v.LocationCharacteristicValueLch = NewCWE(input.Index(5).Index(0))
	return v
}

// NewLCHSlice will construct a slice of type LCH
func NewLCHSlice(input hl7.HL7Type) []LCH {
	values := make([]LCH, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewLCH(input.Index(i))
	}
	return values
}

type PDC struct {
	ManufacturerDistributor []XON
	Country                 CWE
	BrandName               ST
	DeviceFamilyName        ST
	GenericName             CWE
	ModelIdentifier         []ST
	CatalogueIdentifier     ST
	OtherIdentifier         []ST
	ProductCode             CWE
	MarketingBasis          ID
	MarketingApprovalID     ST
	LabeledShelfLife        CQ
	ExpectedShelfLife       CQ
	DateFirstMarketed       DTM
	DateLastMarketed        DTM
}

// NewPDC creates an implementation of PDC
func NewPDC(input hl7.HL7Type) PDC {
	v := PDC{}
	v.ManufacturerDistributor = NewXONSlice(input.Index(1))
	v.Country = NewCWE(input.Index(2).Index(0))
	v.BrandName = NewST(input.Index(3).Index(0))
	v.DeviceFamilyName = NewST(input.Index(4).Index(0))
	v.GenericName = NewCWE(input.Index(5).Index(0))
	v.ModelIdentifier = NewSTSlice(input.Index(6))
	v.CatalogueIdentifier = NewST(input.Index(7).Index(0))
	v.OtherIdentifier = NewSTSlice(input.Index(8))
	v.ProductCode = NewCWE(input.Index(9).Index(0))
	v.MarketingBasis = NewID(input.Index(10).Index(0))
	v.MarketingApprovalID = NewST(input.Index(11).Index(0))
	v.LabeledShelfLife = NewCQ(input.Index(12).Index(0))
	v.ExpectedShelfLife = NewCQ(input.Index(13).Index(0))
	v.DateFirstMarketed = NewDTM(input.Index(14).Index(0))
	v.DateLastMarketed = NewDTM(input.Index(15).Index(0))
	return v
}

// NewPDCSlice will construct a slice of type PDC
func NewPDCSlice(input hl7.HL7Type) []PDC {
	values := make([]PDC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPDC(input.Index(i))
	}
	return values
}

type ADJ struct {
	ProviderAdjustmentNumber                           EI
	PayerAdjustmentNumber                              EI
	AdjustmentSequenceNumber                           SI
	AdjustmentCategory                                 CWE
	AdjustmentAmount                                   CP
	AdjustmentQuantity                                 CQ
	AdjustmentReasonPa                                 CWE
	AdjustmentDescription                              ST
	OriginalValue                                      NM
	SubstituteValue                                    NM
	AdjustmentAction                                   CWE
	ProviderAdjustmentNumberCrossReference             EI
	ProviderProductServiceLineItemNumberCrossReference EI
	AdjustmentDate                                     DTM
	ResponsibleOrganization                            XON
}

// NewADJ creates an implementation of ADJ
func NewADJ(input hl7.HL7Type) ADJ {
	v := ADJ{}
	v.ProviderAdjustmentNumber = NewEI(input.Index(1).Index(0))
	v.PayerAdjustmentNumber = NewEI(input.Index(2).Index(0))
	v.AdjustmentSequenceNumber = NewSI(input.Index(3).Index(0))
	v.AdjustmentCategory = NewCWE(input.Index(4).Index(0))
	v.AdjustmentAmount = NewCP(input.Index(5).Index(0))
	v.AdjustmentQuantity = NewCQ(input.Index(6).Index(0))
	v.AdjustmentReasonPa = NewCWE(input.Index(7).Index(0))
	v.AdjustmentDescription = NewST(input.Index(8).Index(0))
	v.OriginalValue = NewNM(input.Index(9).Index(0))
	v.SubstituteValue = NewNM(input.Index(10).Index(0))
	v.AdjustmentAction = NewCWE(input.Index(11).Index(0))
	v.ProviderAdjustmentNumberCrossReference = NewEI(input.Index(12).Index(0))
	v.ProviderProductServiceLineItemNumberCrossReference = NewEI(input.Index(13).Index(0))
	v.AdjustmentDate = NewDTM(input.Index(14).Index(0))
	v.ResponsibleOrganization = NewXON(input.Index(15).Index(0))
	return v
}

// NewADJSlice will construct a slice of type ADJ
func NewADJSlice(input hl7.HL7Type) []ADJ {
	values := make([]ADJ, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADJ(input.Index(i))
	}
	return values
}

type AL1 struct {
	SetID                           SI
	AllergenTypeCode                CWE
	AllergenCodeMnemonicDescription CWE
	AllergySeverityCode             CWE
	AllergyReactionCode             []ST
	IdentificationDate              ST
}

// NewAL1 creates an implementation of AL1
func NewAL1(input hl7.HL7Type) AL1 {
	v := AL1{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.AllergenTypeCode = NewCWE(input.Index(2).Index(0))
	v.AllergenCodeMnemonicDescription = NewCWE(input.Index(3).Index(0))
	v.AllergySeverityCode = NewCWE(input.Index(4).Index(0))
	v.AllergyReactionCode = NewSTSlice(input.Index(5))
	v.IdentificationDate = NewST(input.Index(6).Index(0))
	return v
}

// NewAL1Slice will construct a slice of type AL1
func NewAL1Slice(input hl7.HL7Type) []AL1 {
	values := make([]AL1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewAL1(input.Index(i))
	}
	return values
}

type LCC struct {
	PrimaryKeyValueLcc PL
	LocationDepartment CWE
	AccommodationType  []CWE
	ChargeCode         []CWE
}

// NewLCC creates an implementation of LCC
func NewLCC(input hl7.HL7Type) LCC {
	v := LCC{}
	v.PrimaryKeyValueLcc = NewPL(input.Index(1).Index(0))
	v.LocationDepartment = NewCWE(input.Index(2).Index(0))
	v.AccommodationType = NewCWESlice(input.Index(3))
	v.ChargeCode = NewCWESlice(input.Index(4))
	return v
}

// NewLCCSlice will construct a slice of type LCC
func NewLCCSlice(input hl7.HL7Type) []LCC {
	values := make([]LCC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewLCC(input.Index(i))
	}
	return values
}

type RXE struct {
	QuantityTiming                                          ST
	GiveCode                                                CWE
	GiveAmountMinimum                                       NM
	GiveAmountMaximum                                       NM
	GiveUnits                                               CWE
	GiveDosageForm                                          CWE
	ProvidersAdministrationInstructions                     []CWE
	DeliverToLocation                                       ST
	SubstitutionStatus                                      ID
	DispenseAmount                                          NM
	DispenseUnits                                           CWE
	NumberOfRefills                                         NM
	OrderingProvidersDeaNumber                              []XCN
	PharmacistTreatmentSuppliersVerifierID                  []XCN
	PrescriptionNumber                                      ST
	NumberOfRefillsRemaining                                NM
	NumberOfRefillsDosesDispensed                           NM
	DTOfMostRecentRefillOrDoseDispensed                     DTM
	TotalDailyDose                                          CQ
	NeedsHumanReview                                        ID
	PharmacyTreatmentSuppliersSpecialDispensingInstructions []CWE
	GivePertimeUnit                                         ST
	GiveRateAmount                                          ST
	GiveRateUnits                                           CWE
	GiveStrength                                            NM
	GiveStrengthUnits                                       CWE
	GiveIndication                                          []CWE
	DispensePackageSize                                     NM
	DispensePackageSizeUnit                                 CWE
	DispensePackageMethod                                   ID
	SupplementaryCode                                       []CWE
	OriginalOrderDateTime                                   DTM
	GiveDrugStrengthVolume                                  NM
	GiveDrugStrengthVolumeUnits                             CWE
	ControlledSubstanceSchedule                             CWE
	FormularyStatus                                         ID
	PharmaceuticalSubstanceAlternative                      []CWE
	PharmacyOfMostRecentFill                                CWE
	InitialDispenseAmount                                   NM
	DispensingPharmacy                                      CWE
	DispensingPharmacyAddress                               XAD
	DeliverToPatientLocation                                PL
	DeliverToAddress                                        XAD
	PharmacyOrderType                                       ID
	PharmacyPhoneNumber                                     []XTN
}

// NewRXE creates an implementation of RXE
func NewRXE(input hl7.HL7Type) RXE {
	v := RXE{}
	v.QuantityTiming = NewST(input.Index(1).Index(0))
	v.GiveCode = NewCWE(input.Index(2).Index(0))
	v.GiveAmountMinimum = NewNM(input.Index(3).Index(0))
	v.GiveAmountMaximum = NewNM(input.Index(4).Index(0))
	v.GiveUnits = NewCWE(input.Index(5).Index(0))
	v.GiveDosageForm = NewCWE(input.Index(6).Index(0))
	v.ProvidersAdministrationInstructions = NewCWESlice(input.Index(7))
	v.DeliverToLocation = NewST(input.Index(8).Index(0))
	v.SubstitutionStatus = NewID(input.Index(9).Index(0))
	v.DispenseAmount = NewNM(input.Index(10).Index(0))
	v.DispenseUnits = NewCWE(input.Index(11).Index(0))
	v.NumberOfRefills = NewNM(input.Index(12).Index(0))
	v.OrderingProvidersDeaNumber = NewXCNSlice(input.Index(13))
	v.PharmacistTreatmentSuppliersVerifierID = NewXCNSlice(input.Index(14))
	v.PrescriptionNumber = NewST(input.Index(15).Index(0))
	v.NumberOfRefillsRemaining = NewNM(input.Index(16).Index(0))
	v.NumberOfRefillsDosesDispensed = NewNM(input.Index(17).Index(0))
	v.DTOfMostRecentRefillOrDoseDispensed = NewDTM(input.Index(18).Index(0))
	v.TotalDailyDose = NewCQ(input.Index(19).Index(0))
	v.NeedsHumanReview = NewID(input.Index(20).Index(0))
	v.PharmacyTreatmentSuppliersSpecialDispensingInstructions = NewCWESlice(input.Index(21))
	v.GivePertimeUnit = NewST(input.Index(22).Index(0))
	v.GiveRateAmount = NewST(input.Index(23).Index(0))
	v.GiveRateUnits = NewCWE(input.Index(24).Index(0))
	v.GiveStrength = NewNM(input.Index(25).Index(0))
	v.GiveStrengthUnits = NewCWE(input.Index(26).Index(0))
	v.GiveIndication = NewCWESlice(input.Index(27))
	v.DispensePackageSize = NewNM(input.Index(28).Index(0))
	v.DispensePackageSizeUnit = NewCWE(input.Index(29).Index(0))
	v.DispensePackageMethod = NewID(input.Index(30).Index(0))
	v.SupplementaryCode = NewCWESlice(input.Index(31))
	v.OriginalOrderDateTime = NewDTM(input.Index(32).Index(0))
	v.GiveDrugStrengthVolume = NewNM(input.Index(33).Index(0))
	v.GiveDrugStrengthVolumeUnits = NewCWE(input.Index(34).Index(0))
	v.ControlledSubstanceSchedule = NewCWE(input.Index(35).Index(0))
	v.FormularyStatus = NewID(input.Index(36).Index(0))
	v.PharmaceuticalSubstanceAlternative = NewCWESlice(input.Index(37))
	v.PharmacyOfMostRecentFill = NewCWE(input.Index(38).Index(0))
	v.InitialDispenseAmount = NewNM(input.Index(39).Index(0))
	v.DispensingPharmacy = NewCWE(input.Index(40).Index(0))
	v.DispensingPharmacyAddress = NewXAD(input.Index(41).Index(0))
	v.DeliverToPatientLocation = NewPL(input.Index(42).Index(0))
	v.DeliverToAddress = NewXAD(input.Index(43).Index(0))
	v.PharmacyOrderType = NewID(input.Index(44).Index(0))
	v.PharmacyPhoneNumber = NewXTNSlice(input.Index(45))
	return v
}

// NewRXESlice will construct a slice of type RXE
func NewRXESlice(input hl7.HL7Type) []RXE {
	values := make([]RXE, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRXE(input.Index(i))
	}
	return values
}

type Zxx struct {
	Zxx1 ST
}

// NewZxx creates an implementation of Zxx
func NewZxx(input hl7.HL7Type) Zxx {
	v := Zxx{}
	v.Zxx1 = NewST(input.Index(1).Index(0))
	return v
}

// NewZxxSlice will construct a slice of type Zxx
func NewZxxSlice(input hl7.HL7Type) []Zxx {
	values := make([]Zxx, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewZxx(input.Index(i))
	}
	return values
}

type CDM struct {
	PrimaryKeyValueCdm           CWE
	ChargeCodeAlias              []CWE
	ChargeDescriptionShort       ST
	ChargeDescriptionLong        ST
	DescriptionOverrideIndicator CWE
	ExplodingCharges             []CWE
	ProcedureCode                []CNE
	ActiveInactiveFlag           ID
	InventoryNumber              []CWE
	ResourceLoad                 NM
	ContractNumber               []CX
	ContractOrganization         []XON
	RoomFeeIndicator             ID
}

// NewCDM creates an implementation of CDM
func NewCDM(input hl7.HL7Type) CDM {
	v := CDM{}
	v.PrimaryKeyValueCdm = NewCWE(input.Index(1).Index(0))
	v.ChargeCodeAlias = NewCWESlice(input.Index(2))
	v.ChargeDescriptionShort = NewST(input.Index(3).Index(0))
	v.ChargeDescriptionLong = NewST(input.Index(4).Index(0))
	v.DescriptionOverrideIndicator = NewCWE(input.Index(5).Index(0))
	v.ExplodingCharges = NewCWESlice(input.Index(6))
	v.ProcedureCode = NewCNESlice(input.Index(7))
	v.ActiveInactiveFlag = NewID(input.Index(8).Index(0))
	v.InventoryNumber = NewCWESlice(input.Index(9))
	v.ResourceLoad = NewNM(input.Index(10).Index(0))
	v.ContractNumber = NewCXSlice(input.Index(11))
	v.ContractOrganization = NewXONSlice(input.Index(12))
	v.RoomFeeIndicator = NewID(input.Index(13).Index(0))
	return v
}

// NewCDMSlice will construct a slice of type CDM
func NewCDMSlice(input hl7.HL7Type) []CDM {
	values := make([]CDM, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCDM(input.Index(i))
	}
	return values
}

type RXA struct {
	GiveSubIDCounter                    NM
	AdministrationSubIDCounter          NM
	DateTimeStartOfAdministration       DTM
	DateTimeEndOfAdministration         DTM
	AdministeredCode                    CWE
	AdministeredAmount                  NM
	AdministeredUnits                   CWE
	AdministeredDosageForm              CWE
	AdministrationNotes                 []CWE
	AdministeringProvider               []XCN
	AdministeredAtLocation              ST
	AdministeredPertimeUnit             ST
	AdministeredStrength                NM
	AdministeredStrengthUnits           CWE
	SubstanceLotNumber                  []ST
	SubstanceExpirationDate             []DTM
	SubstanceManufacturerName           []CWE
	SubstanceTreatmentRefusalReason     []CWE
	Indication                          []CWE
	CompletionStatus                    ID
	ActionCodeRxa                       ID
	SystemEntryDateTime                 DTM
	AdministeredDrugStrengthVolume      NM
	AdministeredDrugStrengthVolumeUnits CWE
	AdministeredBarcodeIdentifier       CWE
	PharmacyOrderType                   ID
	AdministerAt                        PL
	AdministeredAtAddress               XAD
	AdministeredTagIdentifier           []EI
}

// NewRXA creates an implementation of RXA
func NewRXA(input hl7.HL7Type) RXA {
	v := RXA{}
	v.GiveSubIDCounter = NewNM(input.Index(1).Index(0))
	v.AdministrationSubIDCounter = NewNM(input.Index(2).Index(0))
	v.DateTimeStartOfAdministration = NewDTM(input.Index(3).Index(0))
	v.DateTimeEndOfAdministration = NewDTM(input.Index(4).Index(0))
	v.AdministeredCode = NewCWE(input.Index(5).Index(0))
	v.AdministeredAmount = NewNM(input.Index(6).Index(0))
	v.AdministeredUnits = NewCWE(input.Index(7).Index(0))
	v.AdministeredDosageForm = NewCWE(input.Index(8).Index(0))
	v.AdministrationNotes = NewCWESlice(input.Index(9))
	v.AdministeringProvider = NewXCNSlice(input.Index(10))
	v.AdministeredAtLocation = NewST(input.Index(11).Index(0))
	v.AdministeredPertimeUnit = NewST(input.Index(12).Index(0))
	v.AdministeredStrength = NewNM(input.Index(13).Index(0))
	v.AdministeredStrengthUnits = NewCWE(input.Index(14).Index(0))
	v.SubstanceLotNumber = NewSTSlice(input.Index(15))
	v.SubstanceExpirationDate = NewDTMSlice(input.Index(16))
	v.SubstanceManufacturerName = NewCWESlice(input.Index(17))
	v.SubstanceTreatmentRefusalReason = NewCWESlice(input.Index(18))
	v.Indication = NewCWESlice(input.Index(19))
	v.CompletionStatus = NewID(input.Index(20).Index(0))
	v.ActionCodeRxa = NewID(input.Index(21).Index(0))
	v.SystemEntryDateTime = NewDTM(input.Index(22).Index(0))
	v.AdministeredDrugStrengthVolume = NewNM(input.Index(23).Index(0))
	v.AdministeredDrugStrengthVolumeUnits = NewCWE(input.Index(24).Index(0))
	v.AdministeredBarcodeIdentifier = NewCWE(input.Index(25).Index(0))
	v.PharmacyOrderType = NewID(input.Index(26).Index(0))
	v.AdministerAt = NewPL(input.Index(27).Index(0))
	v.AdministeredAtAddress = NewXAD(input.Index(28).Index(0))
	v.AdministeredTagIdentifier = NewEISlice(input.Index(29))
	return v
}

// NewRXASlice will construct a slice of type RXA
func NewRXASlice(input hl7.HL7Type) []RXA {
	values := make([]RXA, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRXA(input.Index(i))
	}
	return values
}

type RXO struct {
	RequestedGiveCode                      CWE
	RequestedGiveAmountMinimum             NM
	RequestedGiveAmountMaximum             NM
	RequestedGiveUnits                     CWE
	RequestedDosageForm                    CWE
	ProvidersPharmacyTreatmentInstructions []CWE
	ProvidersAdministrationInstructions    []CWE
	DeliverToLocation                      ST
	AllowSubstitutions                     ID
	RequestedDispenseCode                  CWE
	RequestedDispenseAmount                NM
	RequestedDispenseUnits                 CWE
	NumberOfRefills                        NM
	OrderingProvidersDeaNumber             []XCN
	PharmacistTreatmentSuppliersVerifierID []XCN
	NeedsHumanReview                       ID
	RequestedGivePertimeUnit               ST
	RequestedGiveStrength                  NM
	RequestedGiveStrengthUnits             CWE
	Indication                             []CWE
	RequestedGiveRateAmount                ST
	RequestedGiveRateUnits                 CWE
	TotalDailyDose                         CQ
	SupplementaryCode                      []CWE
	RequestedDrugStrengthVolume            NM
	RequestedDrugStrengthVolumeUnits       CWE
	PharmacyOrderType                      ID
	DispensingInterval                     NM
	MedicationInstanceIdentifier           EI
	SegmentInstanceIdentifier              EI
	MoodCode                               CNE
	DispensingPharmacy                     CWE
	DispensingPharmacyAddress              XAD
	DeliverToPatientLocation               PL
	DeliverToAddress                       XAD
	PharmacyPhoneNumber                    []XTN
}

// NewRXO creates an implementation of RXO
func NewRXO(input hl7.HL7Type) RXO {
	v := RXO{}
	v.RequestedGiveCode = NewCWE(input.Index(1).Index(0))
	v.RequestedGiveAmountMinimum = NewNM(input.Index(2).Index(0))
	v.RequestedGiveAmountMaximum = NewNM(input.Index(3).Index(0))
	v.RequestedGiveUnits = NewCWE(input.Index(4).Index(0))
	v.RequestedDosageForm = NewCWE(input.Index(5).Index(0))
	v.ProvidersPharmacyTreatmentInstructions = NewCWESlice(input.Index(6))
	v.ProvidersAdministrationInstructions = NewCWESlice(input.Index(7))
	v.DeliverToLocation = NewST(input.Index(8).Index(0))
	v.AllowSubstitutions = NewID(input.Index(9).Index(0))
	v.RequestedDispenseCode = NewCWE(input.Index(10).Index(0))
	v.RequestedDispenseAmount = NewNM(input.Index(11).Index(0))
	v.RequestedDispenseUnits = NewCWE(input.Index(12).Index(0))
	v.NumberOfRefills = NewNM(input.Index(13).Index(0))
	v.OrderingProvidersDeaNumber = NewXCNSlice(input.Index(14))
	v.PharmacistTreatmentSuppliersVerifierID = NewXCNSlice(input.Index(15))
	v.NeedsHumanReview = NewID(input.Index(16).Index(0))
	v.RequestedGivePertimeUnit = NewST(input.Index(17).Index(0))
	v.RequestedGiveStrength = NewNM(input.Index(18).Index(0))
	v.RequestedGiveStrengthUnits = NewCWE(input.Index(19).Index(0))
	v.Indication = NewCWESlice(input.Index(20))
	v.RequestedGiveRateAmount = NewST(input.Index(21).Index(0))
	v.RequestedGiveRateUnits = NewCWE(input.Index(22).Index(0))
	v.TotalDailyDose = NewCQ(input.Index(23).Index(0))
	v.SupplementaryCode = NewCWESlice(input.Index(24))
	v.RequestedDrugStrengthVolume = NewNM(input.Index(25).Index(0))
	v.RequestedDrugStrengthVolumeUnits = NewCWE(input.Index(26).Index(0))
	v.PharmacyOrderType = NewID(input.Index(27).Index(0))
	v.DispensingInterval = NewNM(input.Index(28).Index(0))
	v.MedicationInstanceIdentifier = NewEI(input.Index(29).Index(0))
	v.SegmentInstanceIdentifier = NewEI(input.Index(30).Index(0))
	v.MoodCode = NewCNE(input.Index(31).Index(0))
	v.DispensingPharmacy = NewCWE(input.Index(32).Index(0))
	v.DispensingPharmacyAddress = NewXAD(input.Index(33).Index(0))
	v.DeliverToPatientLocation = NewPL(input.Index(34).Index(0))
	v.DeliverToAddress = NewXAD(input.Index(35).Index(0))
	v.PharmacyPhoneNumber = NewXTNSlice(input.Index(36))
	return v
}

// NewRXOSlice will construct a slice of type RXO
func NewRXOSlice(input hl7.HL7Type) []RXO {
	values := make([]RXO, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRXO(input.Index(i))
	}
	return values
}

type PRD struct {
	ProviderRole                                 []CWE
	ProviderName                                 []XPN
	ProviderAddress                              []XAD
	ProviderLocation                             PL
	ProviderCommunicationInformation             []XTN
	PreferredMethodOfContact                     CWE
	ProviderIdentifiers                          []PLN
	EffectiveStartDateOfProviderRole             DTM
	EffectiveEndDateOfProviderRole               []DTM
	ProviderOrganizationNameAndIdentifier        XON
	ProviderOrganizationAddress                  []XAD
	ProviderOrganizationLocationInformation      []PL
	ProviderOrganizationCommunicationInformation []XTN
	ProviderOrganizationMethodOfContact          CWE
}

// NewPRD creates an implementation of PRD
func NewPRD(input hl7.HL7Type) PRD {
	v := PRD{}
	v.ProviderRole = NewCWESlice(input.Index(1))
	v.ProviderName = NewXPNSlice(input.Index(2))
	v.ProviderAddress = NewXADSlice(input.Index(3))
	v.ProviderLocation = NewPL(input.Index(4).Index(0))
	v.ProviderCommunicationInformation = NewXTNSlice(input.Index(5))
	v.PreferredMethodOfContact = NewCWE(input.Index(6).Index(0))
	v.ProviderIdentifiers = NewPLNSlice(input.Index(7))
	v.EffectiveStartDateOfProviderRole = NewDTM(input.Index(8).Index(0))
	v.EffectiveEndDateOfProviderRole = NewDTMSlice(input.Index(9))
	v.ProviderOrganizationNameAndIdentifier = NewXON(input.Index(10).Index(0))
	v.ProviderOrganizationAddress = NewXADSlice(input.Index(11))
	v.ProviderOrganizationLocationInformation = NewPLSlice(input.Index(12))
	v.ProviderOrganizationCommunicationInformation = NewXTNSlice(input.Index(13))
	v.ProviderOrganizationMethodOfContact = NewCWE(input.Index(14).Index(0))
	return v
}

// NewPRDSlice will construct a slice of type PRD
func NewPRDSlice(input hl7.HL7Type) []PRD {
	values := make([]PRD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPRD(input.Index(i))
	}
	return values
}

type PMT struct {
	PaymentRemittanceAdviceNumber       EI
	PaymentRemittanceEffectiveDateTime  DTM
	PaymentRemittanceExpirationDateTime DTM
	PaymentMethod                       CWE
	PaymentRemittanceDateTime           DTM
	PaymentRemittanceAmount             CP
	CheckNumber                         EI
	PayeeBankIdentification             XON
	PayeeTransitNumber                  ST
	PayeeBankAccountID                  CX
	PaymentOrganization                 XON
	EsrCodeLine                         ST
}

// NewPMT creates an implementation of PMT
func NewPMT(input hl7.HL7Type) PMT {
	v := PMT{}
	v.PaymentRemittanceAdviceNumber = NewEI(input.Index(1).Index(0))
	v.PaymentRemittanceEffectiveDateTime = NewDTM(input.Index(2).Index(0))
	v.PaymentRemittanceExpirationDateTime = NewDTM(input.Index(3).Index(0))
	v.PaymentMethod = NewCWE(input.Index(4).Index(0))
	v.PaymentRemittanceDateTime = NewDTM(input.Index(5).Index(0))
	v.PaymentRemittanceAmount = NewCP(input.Index(6).Index(0))
	v.CheckNumber = NewEI(input.Index(7).Index(0))
	v.PayeeBankIdentification = NewXON(input.Index(8).Index(0))
	v.PayeeTransitNumber = NewST(input.Index(9).Index(0))
	v.PayeeBankAccountID = NewCX(input.Index(10).Index(0))
	v.PaymentOrganization = NewXON(input.Index(11).Index(0))
	v.EsrCodeLine = NewST(input.Index(12).Index(0))
	return v
}

// NewPMTSlice will construct a slice of type PMT
func NewPMTSlice(input hl7.HL7Type) []PMT {
	values := make([]PMT, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPMT(input.Index(i))
	}
	return values
}

type PV2 struct {
	PriorPendingLocation                PL
	AccommodationCode                   CWE
	AdmitReason                         CWE
	TransferReason                      CWE
	PatientValuables                    []ST
	PatientValuablesLocation            ST
	VisitUserCode                       []CWE
	ExpectedAdmitDateTime               DTM
	ExpectedDischargeDateTime           DTM
	EstimatedLengthOfInpatientStay      NM
	ActualLengthOfInpatientStay         NM
	VisitDescription                    ST
	ReferralSourceCode                  []XCN
	PreviousServiceDate                 DT
	EmploymentIllnessRelatedIndicator   ID
	PurgeStatusCode                     CWE
	PurgeStatusDate                     DT
	SpecialProgramCode                  CWE
	RetentionIndicator                  ID
	ExpectedNumberOfInsurancePlans      NM
	VisitPublicityCode                  CWE
	VisitProtectionIndicator            ID
	ClinicOrganizationName              []XON
	PatientStatusCode                   CWE
	VisitPriorityCode                   CWE
	PreviousTreatmentDate               DT
	ExpectedDischargeDisposition        CWE
	SignatureOnFileDate                 DT
	FirstSimilarIllnessDate             DT
	PatientChargeAdjustmentCode         CWE
	RecurringServiceCode                CWE
	BillingMediaCode                    ID
	ExpectedSurgeryDateAndTime          DTM
	MilitaryPartnershipCode             ID
	MilitaryNonAvailabilityCode         ID
	NewbornBabyIndicator                ID
	BabyDetainedIndicator               ID
	ModeOfArrivalCode                   CWE
	RecreationalDrugUseCode             []CWE
	AdmissionLevelOfCareCode            CWE
	PrecautionCode                      []CWE
	PatientConditionCode                CWE
	LivingWillCode                      CWE
	OrganDonorCode                      CWE
	AdvanceDirectiveCode                []CWE
	PatientStatusEffectiveDate          DT
	ExpectedLoaReturnDateTime           DTM
	ExpectedPreAdmissionTestingDateTime DTM
	NotifyClergyCode                    []CWE
	AdvanceDirectiveLastVerifiedDate    DT
}

// NewPV2 creates an implementation of PV2
func NewPV2(input hl7.HL7Type) PV2 {
	v := PV2{}
	v.PriorPendingLocation = NewPL(input.Index(1).Index(0))
	v.AccommodationCode = NewCWE(input.Index(2).Index(0))
	v.AdmitReason = NewCWE(input.Index(3).Index(0))
	v.TransferReason = NewCWE(input.Index(4).Index(0))
	v.PatientValuables = NewSTSlice(input.Index(5))
	v.PatientValuablesLocation = NewST(input.Index(6).Index(0))
	v.VisitUserCode = NewCWESlice(input.Index(7))
	v.ExpectedAdmitDateTime = NewDTM(input.Index(8).Index(0))
	v.ExpectedDischargeDateTime = NewDTM(input.Index(9).Index(0))
	v.EstimatedLengthOfInpatientStay = NewNM(input.Index(10).Index(0))
	v.ActualLengthOfInpatientStay = NewNM(input.Index(11).Index(0))
	v.VisitDescription = NewST(input.Index(12).Index(0))
	v.ReferralSourceCode = NewXCNSlice(input.Index(13))
	v.PreviousServiceDate = NewDT(input.Index(14).Index(0))
	v.EmploymentIllnessRelatedIndicator = NewID(input.Index(15).Index(0))
	v.PurgeStatusCode = NewCWE(input.Index(16).Index(0))
	v.PurgeStatusDate = NewDT(input.Index(17).Index(0))
	v.SpecialProgramCode = NewCWE(input.Index(18).Index(0))
	v.RetentionIndicator = NewID(input.Index(19).Index(0))
	v.ExpectedNumberOfInsurancePlans = NewNM(input.Index(20).Index(0))
	v.VisitPublicityCode = NewCWE(input.Index(21).Index(0))
	v.VisitProtectionIndicator = NewID(input.Index(22).Index(0))
	v.ClinicOrganizationName = NewXONSlice(input.Index(23))
	v.PatientStatusCode = NewCWE(input.Index(24).Index(0))
	v.VisitPriorityCode = NewCWE(input.Index(25).Index(0))
	v.PreviousTreatmentDate = NewDT(input.Index(26).Index(0))
	v.ExpectedDischargeDisposition = NewCWE(input.Index(27).Index(0))
	v.SignatureOnFileDate = NewDT(input.Index(28).Index(0))
	v.FirstSimilarIllnessDate = NewDT(input.Index(29).Index(0))
	v.PatientChargeAdjustmentCode = NewCWE(input.Index(30).Index(0))
	v.RecurringServiceCode = NewCWE(input.Index(31).Index(0))
	v.BillingMediaCode = NewID(input.Index(32).Index(0))
	v.ExpectedSurgeryDateAndTime = NewDTM(input.Index(33).Index(0))
	v.MilitaryPartnershipCode = NewID(input.Index(34).Index(0))
	v.MilitaryNonAvailabilityCode = NewID(input.Index(35).Index(0))
	v.NewbornBabyIndicator = NewID(input.Index(36).Index(0))
	v.BabyDetainedIndicator = NewID(input.Index(37).Index(0))
	v.ModeOfArrivalCode = NewCWE(input.Index(38).Index(0))
	v.RecreationalDrugUseCode = NewCWESlice(input.Index(39))
	v.AdmissionLevelOfCareCode = NewCWE(input.Index(40).Index(0))
	v.PrecautionCode = NewCWESlice(input.Index(41))
	v.PatientConditionCode = NewCWE(input.Index(42).Index(0))
	v.LivingWillCode = NewCWE(input.Index(43).Index(0))
	v.OrganDonorCode = NewCWE(input.Index(44).Index(0))
	v.AdvanceDirectiveCode = NewCWESlice(input.Index(45))
	v.PatientStatusEffectiveDate = NewDT(input.Index(46).Index(0))
	v.ExpectedLoaReturnDateTime = NewDTM(input.Index(47).Index(0))
	v.ExpectedPreAdmissionTestingDateTime = NewDTM(input.Index(48).Index(0))
	v.NotifyClergyCode = NewCWESlice(input.Index(49))
	v.AdvanceDirectiveLastVerifiedDate = NewDT(input.Index(50).Index(0))
	return v
}

// NewPV2Slice will construct a slice of type PV2
func NewPV2Slice(input hl7.HL7Type) []PV2 {
	values := make([]PV2, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPV2(input.Index(i))
	}
	return values
}

type LOC struct {
	PrimaryKeyValueLoc  PL
	LocationDescription ST
	LocationTypeLoc     []CWE
	OrganizationNameLoc []XON
	LocationAddress     []XAD
	LocationPhone       []XTN
	LicenseNumber       []CWE
	LocationEquipment   []CWE
	LocationServiceCode CWE
}

// NewLOC creates an implementation of LOC
func NewLOC(input hl7.HL7Type) LOC {
	v := LOC{}
	v.PrimaryKeyValueLoc = NewPL(input.Index(1).Index(0))
	v.LocationDescription = NewST(input.Index(2).Index(0))
	v.LocationTypeLoc = NewCWESlice(input.Index(3))
	v.OrganizationNameLoc = NewXONSlice(input.Index(4))
	v.LocationAddress = NewXADSlice(input.Index(5))
	v.LocationPhone = NewXTNSlice(input.Index(6))
	v.LicenseNumber = NewCWESlice(input.Index(7))
	v.LocationEquipment = NewCWESlice(input.Index(8))
	v.LocationServiceCode = NewCWE(input.Index(9).Index(0))
	return v
}

// NewLOCSlice will construct a slice of type LOC
func NewLOCSlice(input hl7.HL7Type) []LOC {
	values := make([]LOC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewLOC(input.Index(i))
	}
	return values
}

type IIM struct {
	PrimaryKeyValueIim            CWE
	ServiceItemCode               CWE
	InventoryLotNumber            ST
	InventoryExpirationDate       DTM
	InventoryManufacturerName     CWE
	InventoryLocation             CWE
	InventoryReceivedDate         DTM
	InventoryReceivedQuantity     NM
	InventoryReceivedQuantityUnit CWE
	InventoryReceivedItemCost     MO
	InventoryOnHandDate           DTM
	InventoryOnHandQuantity       NM
	InventoryOnHandQuantityUnit   CWE
	ProcedureCode                 CNE
	ProcedureCodeModifier         []CNE
}

// NewIIM creates an implementation of IIM
func NewIIM(input hl7.HL7Type) IIM {
	v := IIM{}
	v.PrimaryKeyValueIim = NewCWE(input.Index(1).Index(0))
	v.ServiceItemCode = NewCWE(input.Index(2).Index(0))
	v.InventoryLotNumber = NewST(input.Index(3).Index(0))
	v.InventoryExpirationDate = NewDTM(input.Index(4).Index(0))
	v.InventoryManufacturerName = NewCWE(input.Index(5).Index(0))
	v.InventoryLocation = NewCWE(input.Index(6).Index(0))
	v.InventoryReceivedDate = NewDTM(input.Index(7).Index(0))
	v.InventoryReceivedQuantity = NewNM(input.Index(8).Index(0))
	v.InventoryReceivedQuantityUnit = NewCWE(input.Index(9).Index(0))
	v.InventoryReceivedItemCost = NewMO(input.Index(10).Index(0))
	v.InventoryOnHandDate = NewDTM(input.Index(11).Index(0))
	v.InventoryOnHandQuantity = NewNM(input.Index(12).Index(0))
	v.InventoryOnHandQuantityUnit = NewCWE(input.Index(13).Index(0))
	v.ProcedureCode = NewCNE(input.Index(14).Index(0))
	v.ProcedureCodeModifier = NewCNESlice(input.Index(15))
	return v
}

// NewIIMSlice will construct a slice of type IIM
func NewIIMSlice(input hl7.HL7Type) []IIM {
	values := make([]IIM, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewIIM(input.Index(i))
	}
	return values
}

type BUI struct {
	SetID                     SI
	BloodUnitIdentifier       EI
	BloodUnitType             CWE
	BloodUnitWeight           NM
	WeightUnits               CNE
	BloodUnitVolume           NM
	VolumeUnits               CNE
	ContainerCatalogNumber    ST
	ContainerLotNumber        ST
	ContainerManufacturer     XON
	TransportTemperature      NR
	TransportTemperatureUnits []CNE
}

// NewBUI creates an implementation of BUI
func NewBUI(input hl7.HL7Type) BUI {
	v := BUI{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.BloodUnitIdentifier = NewEI(input.Index(2).Index(0))
	v.BloodUnitType = NewCWE(input.Index(3).Index(0))
	v.BloodUnitWeight = NewNM(input.Index(4).Index(0))
	v.WeightUnits = NewCNE(input.Index(5).Index(0))
	v.BloodUnitVolume = NewNM(input.Index(6).Index(0))
	v.VolumeUnits = NewCNE(input.Index(7).Index(0))
	v.ContainerCatalogNumber = NewST(input.Index(8).Index(0))
	v.ContainerLotNumber = NewST(input.Index(9).Index(0))
	v.ContainerManufacturer = NewXON(input.Index(10).Index(0))
	v.TransportTemperature = NewNR(input.Index(11).Index(0))
	v.TransportTemperatureUnits = NewCNESlice(input.Index(12))
	return v
}

// NewBUISlice will construct a slice of type BUI
func NewBUISlice(input hl7.HL7Type) []BUI {
	values := make([]BUI, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBUI(input.Index(i))
	}
	return values
}

type TCD struct {
	UniversalServiceIdentifier            CWE
	AutoDilutionFactor                    SN
	RerunDilutionFactor                   SN
	PreDilutionFactor                     SN
	EndogenousContentOfPreDilutionDiluent SN
	AutomaticRepeatAllowed                ID
	ReflexAllowed                         ID
	AnalyteRepeatStatus                   CWE
}

// NewTCD creates an implementation of TCD
func NewTCD(input hl7.HL7Type) TCD {
	v := TCD{}
	v.UniversalServiceIdentifier = NewCWE(input.Index(1).Index(0))
	v.AutoDilutionFactor = NewSN(input.Index(2).Index(0))
	v.RerunDilutionFactor = NewSN(input.Index(3).Index(0))
	v.PreDilutionFactor = NewSN(input.Index(4).Index(0))
	v.EndogenousContentOfPreDilutionDiluent = NewSN(input.Index(5).Index(0))
	v.AutomaticRepeatAllowed = NewID(input.Index(6).Index(0))
	v.ReflexAllowed = NewID(input.Index(7).Index(0))
	v.AnalyteRepeatStatus = NewCWE(input.Index(8).Index(0))
	return v
}

// NewTCDSlice will construct a slice of type TCD
func NewTCDSlice(input hl7.HL7Type) []TCD {
	values := make([]TCD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewTCD(input.Index(i))
	}
	return values
}

type NDS struct {
	NotificationReferenceNumber NM
	NotificationDateTime        DTM
	NotificationAlertSeverity   CWE
	NotificationCode            CWE
}

// NewNDS creates an implementation of NDS
func NewNDS(input hl7.HL7Type) NDS {
	v := NDS{}
	v.NotificationReferenceNumber = NewNM(input.Index(1).Index(0))
	v.NotificationDateTime = NewDTM(input.Index(2).Index(0))
	v.NotificationAlertSeverity = NewCWE(input.Index(3).Index(0))
	v.NotificationCode = NewCWE(input.Index(4).Index(0))
	return v
}

// NewNDSSlice will construct a slice of type NDS
func NewNDSSlice(input hl7.HL7Type) []NDS {
	values := make([]NDS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewNDS(input.Index(i))
	}
	return values
}

type RXV struct {
	SetID                    SI
	BolusType                ID
	BolusDoseAmount          NM
	BolusDoseAmountUnits     CWE
	BolusDoseVolume          NM
	BolusDoseVolumeUnits     CWE
	PCAType                  ID
	PCADoseAmount            NM
	PCADoseAmountUnits       CWE
	PCADoseAmountVolume      NM
	PCADoseAmountVolumeUnits CWE
	MaxDoseAmount            NM
	MaxDoseAmountUnits       CWE
	MaxDoseAmountVolume      NM
	MaxDoseAmountVolumeUnits CWE
	MaxDosePerTime           CQ
	LockoutInterval          CQ
	SyringeManufacturer      CWE
	SyringeModelNumber       CWE
	SyringeSize              NM
	SyringeSizeUnits         CWE
	ActionCode               ID
}

// NewRXV creates an implementation of RXV
func NewRXV(input hl7.HL7Type) RXV {
	v := RXV{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.BolusType = NewID(input.Index(2).Index(0))
	v.BolusDoseAmount = NewNM(input.Index(3).Index(0))
	v.BolusDoseAmountUnits = NewCWE(input.Index(4).Index(0))
	v.BolusDoseVolume = NewNM(input.Index(5).Index(0))
	v.BolusDoseVolumeUnits = NewCWE(input.Index(6).Index(0))
	v.PCAType = NewID(input.Index(7).Index(0))
	v.PCADoseAmount = NewNM(input.Index(8).Index(0))
	v.PCADoseAmountUnits = NewCWE(input.Index(9).Index(0))
	v.PCADoseAmountVolume = NewNM(input.Index(10).Index(0))
	v.PCADoseAmountVolumeUnits = NewCWE(input.Index(11).Index(0))
	v.MaxDoseAmount = NewNM(input.Index(12).Index(0))
	v.MaxDoseAmountUnits = NewCWE(input.Index(13).Index(0))
	v.MaxDoseAmountVolume = NewNM(input.Index(14).Index(0))
	v.MaxDoseAmountVolumeUnits = NewCWE(input.Index(15).Index(0))
	v.MaxDosePerTime = NewCQ(input.Index(16).Index(0))
	v.LockoutInterval = NewCQ(input.Index(17).Index(0))
	v.SyringeManufacturer = NewCWE(input.Index(18).Index(0))
	v.SyringeModelNumber = NewCWE(input.Index(19).Index(0))
	v.SyringeSize = NewNM(input.Index(20).Index(0))
	v.SyringeSizeUnits = NewCWE(input.Index(21).Index(0))
	v.ActionCode = NewID(input.Index(22).Index(0))
	return v
}

// NewRXVSlice will construct a slice of type RXV
func NewRXVSlice(input hl7.HL7Type) []RXV {
	values := make([]RXV, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRXV(input.Index(i))
	}
	return values
}

type TXA struct {
	SetID                                    SI
	DocumentType                             CWE
	DocumentContentPresentation              ID
	ActivityDateTime                         DTM
	PrimaryActivityProviderCodeName          []XCN
	OriginationDateTime                      DTM
	TranscriptionDateTime                    DTM
	EditDateTime                             []DTM
	OriginatorCodeName                       []XCN
	AssignedDocumentAuthenticator            []XCN
	TranscriptionistCodeName                 []XCN
	UniqueDocumentNumber                     EI
	ParentDocumentNumber                     EI
	PlacerOrderNumber                        []EI
	FillerOrderNumber                        EI
	UniqueDocumentFileName                   ST
	DocumentCompletionStatus                 ID
	DocumentConfidentialityStatus            ID
	DocumentAvailabilityStatus               ID
	DocumentStorageStatus                    ID
	DocumentChangeReason                     ST
	AuthenticationPersonTimeStampset         []PPN
	DistributedCopiescodeAndNameOfRecipients []XCN
	FolderAssignment                         []CWE
	DocumentTitle                            []ST
	AgreedDueDateTime                        DTM
}

// NewTXA creates an implementation of TXA
func NewTXA(input hl7.HL7Type) TXA {
	v := TXA{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.DocumentType = NewCWE(input.Index(2).Index(0))
	v.DocumentContentPresentation = NewID(input.Index(3).Index(0))
	v.ActivityDateTime = NewDTM(input.Index(4).Index(0))
	v.PrimaryActivityProviderCodeName = NewXCNSlice(input.Index(5))
	v.OriginationDateTime = NewDTM(input.Index(6).Index(0))
	v.TranscriptionDateTime = NewDTM(input.Index(7).Index(0))
	v.EditDateTime = NewDTMSlice(input.Index(8))
	v.OriginatorCodeName = NewXCNSlice(input.Index(9))
	v.AssignedDocumentAuthenticator = NewXCNSlice(input.Index(10))
	v.TranscriptionistCodeName = NewXCNSlice(input.Index(11))
	v.UniqueDocumentNumber = NewEI(input.Index(12).Index(0))
	v.ParentDocumentNumber = NewEI(input.Index(13).Index(0))
	v.PlacerOrderNumber = NewEISlice(input.Index(14))
	v.FillerOrderNumber = NewEI(input.Index(15).Index(0))
	v.UniqueDocumentFileName = NewST(input.Index(16).Index(0))
	v.DocumentCompletionStatus = NewID(input.Index(17).Index(0))
	v.DocumentConfidentialityStatus = NewID(input.Index(18).Index(0))
	v.DocumentAvailabilityStatus = NewID(input.Index(19).Index(0))
	v.DocumentStorageStatus = NewID(input.Index(20).Index(0))
	v.DocumentChangeReason = NewST(input.Index(21).Index(0))
	v.AuthenticationPersonTimeStampset = NewPPNSlice(input.Index(22))
	v.DistributedCopiescodeAndNameOfRecipients = NewXCNSlice(input.Index(23))
	v.FolderAssignment = NewCWESlice(input.Index(24))
	v.DocumentTitle = NewSTSlice(input.Index(25))
	v.AgreedDueDateTime = NewDTM(input.Index(26).Index(0))
	return v
}

// NewTXASlice will construct a slice of type TXA
func NewTXASlice(input hl7.HL7Type) []TXA {
	values := make([]TXA, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewTXA(input.Index(i))
	}
	return values
}

type PYE struct {
	SetID                             SI
	PayeeType                         CWE
	PayeeRelationshipToInvoicepatient CWE
	PayeeIdentificationList           XON
	PayeePersonName                   XPN
	PayeeAddress                      XAD
	PaymentMethod                     CWE
}

// NewPYE creates an implementation of PYE
func NewPYE(input hl7.HL7Type) PYE {
	v := PYE{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.PayeeType = NewCWE(input.Index(2).Index(0))
	v.PayeeRelationshipToInvoicepatient = NewCWE(input.Index(3).Index(0))
	v.PayeeIdentificationList = NewXON(input.Index(4).Index(0))
	v.PayeePersonName = NewXPN(input.Index(5).Index(0))
	v.PayeeAddress = NewXAD(input.Index(6).Index(0))
	v.PaymentMethod = NewCWE(input.Index(7).Index(0))
	return v
}

// NewPYESlice will construct a slice of type PYE
func NewPYESlice(input hl7.HL7Type) []PYE {
	values := make([]PYE, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPYE(input.Index(i))
	}
	return values
}

type RF1 struct {
	ReferralStatus                        CWE
	ReferralPriority                      CWE
	ReferralType                          CWE
	ReferralDisposition                   []CWE
	ReferralCategory                      CWE
	OriginatingReferralIdentifier         EI
	EffectiveDate                         DTM
	ExpirationDate                        DTM
	ProcessDate                           DTM
	ReferralReason                        []CWE
	ExternalReferralIdentifier            []EI
	ReferralDocumentationCompletionStatus CWE
	PlannedTreatmentStopDate              DTM
	ReferralReasonText                    ST
	NumberOfAuthorizedTreatmentsUnits     CQ
	NumberOfUsedTreatmentsUnits           CQ
	NumberOfScheduleTreatmentsUnits       CQ
	RemainingBenefitAmount                MO
	AuthorizedProvider                    XON
	AuthorizedHealthProfessional          XCN
	SourceText                            ST
	SourceDate                            DTM
	SourcePhone                           XTN
	Comment                               ST
	ActionCode                            ID
}

// NewRF1 creates an implementation of RF1
func NewRF1(input hl7.HL7Type) RF1 {
	v := RF1{}
	v.ReferralStatus = NewCWE(input.Index(1).Index(0))
	v.ReferralPriority = NewCWE(input.Index(2).Index(0))
	v.ReferralType = NewCWE(input.Index(3).Index(0))
	v.ReferralDisposition = NewCWESlice(input.Index(4))
	v.ReferralCategory = NewCWE(input.Index(5).Index(0))
	v.OriginatingReferralIdentifier = NewEI(input.Index(6).Index(0))
	v.EffectiveDate = NewDTM(input.Index(7).Index(0))
	v.ExpirationDate = NewDTM(input.Index(8).Index(0))
	v.ProcessDate = NewDTM(input.Index(9).Index(0))
	v.ReferralReason = NewCWESlice(input.Index(10))
	v.ExternalReferralIdentifier = NewEISlice(input.Index(11))
	v.ReferralDocumentationCompletionStatus = NewCWE(input.Index(12).Index(0))
	v.PlannedTreatmentStopDate = NewDTM(input.Index(13).Index(0))
	v.ReferralReasonText = NewST(input.Index(14).Index(0))
	v.NumberOfAuthorizedTreatmentsUnits = NewCQ(input.Index(15).Index(0))
	v.NumberOfUsedTreatmentsUnits = NewCQ(input.Index(16).Index(0))
	v.NumberOfScheduleTreatmentsUnits = NewCQ(input.Index(17).Index(0))
	v.RemainingBenefitAmount = NewMO(input.Index(18).Index(0))
	v.AuthorizedProvider = NewXON(input.Index(19).Index(0))
	v.AuthorizedHealthProfessional = NewXCN(input.Index(20).Index(0))
	v.SourceText = NewST(input.Index(21).Index(0))
	v.SourceDate = NewDTM(input.Index(22).Index(0))
	v.SourcePhone = NewXTN(input.Index(23).Index(0))
	v.Comment = NewST(input.Index(24).Index(0))
	v.ActionCode = NewID(input.Index(25).Index(0))
	return v
}

// NewRF1Slice will construct a slice of type RF1
func NewRF1Slice(input hl7.HL7Type) []RF1 {
	values := make([]RF1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRF1(input.Index(i))
	}
	return values
}

type ARV struct {
	SetID                                SI
	AccessRestrictionActionCode          CNE
	AccessRestrictionValue               CWE
	AccessRestrictionReason              []CWE
	SpecialAccessRestrictionInstructions []ST
	AccessRestrictionDateRange           DR
}

// NewARV creates an implementation of ARV
func NewARV(input hl7.HL7Type) ARV {
	v := ARV{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.AccessRestrictionActionCode = NewCNE(input.Index(2).Index(0))
	v.AccessRestrictionValue = NewCWE(input.Index(3).Index(0))
	v.AccessRestrictionReason = NewCWESlice(input.Index(4))
	v.SpecialAccessRestrictionInstructions = NewSTSlice(input.Index(5))
	v.AccessRestrictionDateRange = NewDR(input.Index(6).Index(0))
	return v
}

// NewARVSlice will construct a slice of type ARV
func NewARVSlice(input hl7.HL7Type) []ARV {
	values := make([]ARV, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewARV(input.Index(i))
	}
	return values
}

type IAM struct {
	SetID                                SI
	AllergenTypeCode                     CWE
	AllergenCodeMnemonicDescription      CWE
	AllergySeverityCode                  CWE
	AllergyReactionCode                  []ST
	AllergyActionCode                    CNE
	AllergyUniqueIdentifier              EI
	ActionReason                         ST
	SensitivityToCausativeAgentCode      CWE
	AllergenGroupCodeMnemonicDescription CWE
	OnsetDate                            DT
	OnsetDateText                        ST
	ReportedDateTime                     DTM
	ReportedBy                           XPN
	RelationshipToPatientCode            CWE
	AlertDeviceCode                      CWE
	AllergyClinicalStatusCode            CWE
	StatusedByPerson                     XCN
	StatusedByOrganization               XON
	StatusedAtDateTime                   DTM
	InactivatedByPerson                  XCN
	InactivatedDateTime                  DTM
	InitiallyRecordedByPerson            XCN
	InitiallyRecordedDateTime            DTM
	ModifiedByPerson                     XCN
	ModifiedDateTime                     DTM
	ClinicianIdentifiedCode              CWE
	InitiallyRecordedByOrganization      XON
	ModifiedByOrganization               XON
	InactivatedByOrganization            XON
}

// NewIAM creates an implementation of IAM
func NewIAM(input hl7.HL7Type) IAM {
	v := IAM{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.AllergenTypeCode = NewCWE(input.Index(2).Index(0))
	v.AllergenCodeMnemonicDescription = NewCWE(input.Index(3).Index(0))
	v.AllergySeverityCode = NewCWE(input.Index(4).Index(0))
	v.AllergyReactionCode = NewSTSlice(input.Index(5))
	v.AllergyActionCode = NewCNE(input.Index(6).Index(0))
	v.AllergyUniqueIdentifier = NewEI(input.Index(7).Index(0))
	v.ActionReason = NewST(input.Index(8).Index(0))
	v.SensitivityToCausativeAgentCode = NewCWE(input.Index(9).Index(0))
	v.AllergenGroupCodeMnemonicDescription = NewCWE(input.Index(10).Index(0))
	v.OnsetDate = NewDT(input.Index(11).Index(0))
	v.OnsetDateText = NewST(input.Index(12).Index(0))
	v.ReportedDateTime = NewDTM(input.Index(13).Index(0))
	v.ReportedBy = NewXPN(input.Index(14).Index(0))
	v.RelationshipToPatientCode = NewCWE(input.Index(15).Index(0))
	v.AlertDeviceCode = NewCWE(input.Index(16).Index(0))
	v.AllergyClinicalStatusCode = NewCWE(input.Index(17).Index(0))
	v.StatusedByPerson = NewXCN(input.Index(18).Index(0))
	v.StatusedByOrganization = NewXON(input.Index(19).Index(0))
	v.StatusedAtDateTime = NewDTM(input.Index(20).Index(0))
	v.InactivatedByPerson = NewXCN(input.Index(21).Index(0))
	v.InactivatedDateTime = NewDTM(input.Index(22).Index(0))
	v.InitiallyRecordedByPerson = NewXCN(input.Index(23).Index(0))
	v.InitiallyRecordedDateTime = NewDTM(input.Index(24).Index(0))
	v.ModifiedByPerson = NewXCN(input.Index(25).Index(0))
	v.ModifiedDateTime = NewDTM(input.Index(26).Index(0))
	v.ClinicianIdentifiedCode = NewCWE(input.Index(27).Index(0))
	v.InitiallyRecordedByOrganization = NewXON(input.Index(28).Index(0))
	v.ModifiedByOrganization = NewXON(input.Index(29).Index(0))
	v.InactivatedByOrganization = NewXON(input.Index(30).Index(0))
	return v
}

// NewIAMSlice will construct a slice of type IAM
func NewIAMSlice(input hl7.HL7Type) []IAM {
	values := make([]IAM, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewIAM(input.Index(i))
	}
	return values
}

type URS struct {
	URS1 ST
}

// NewURS creates an implementation of URS
func NewURS(input hl7.HL7Type) URS {
	v := URS{}
	v.URS1 = NewST(input.Index(1).Index(0))
	return v
}

// NewURSSlice will construct a slice of type URS
func NewURSSlice(input hl7.HL7Type) []URS {
	values := make([]URS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewURS(input.Index(i))
	}
	return values
}

type PSH struct {
	ReportType                                         ST
	ReportFormIdentifier                               ST
	ReportDate                                         DTM
	ReportIntervalStartDate                            DTM
	ReportIntervalEndDate                              DTM
	QuantityManufactured                               CQ
	QuantityDistributed                                CQ
	QuantityDistributedMethod                          ID
	QuantityDistributedComment                         FT
	QuantityInUse                                      CQ
	QuantityInUseMethod                                ID
	QuantityInUseComment                               FT
	NumberOfProductExperienceReportsFiledByFacility    []NM
	NumberOfProductExperienceReportsFiledByDistributor []NM
}

// NewPSH creates an implementation of PSH
func NewPSH(input hl7.HL7Type) PSH {
	v := PSH{}
	v.ReportType = NewST(input.Index(1).Index(0))
	v.ReportFormIdentifier = NewST(input.Index(2).Index(0))
	v.ReportDate = NewDTM(input.Index(3).Index(0))
	v.ReportIntervalStartDate = NewDTM(input.Index(4).Index(0))
	v.ReportIntervalEndDate = NewDTM(input.Index(5).Index(0))
	v.QuantityManufactured = NewCQ(input.Index(6).Index(0))
	v.QuantityDistributed = NewCQ(input.Index(7).Index(0))
	v.QuantityDistributedMethod = NewID(input.Index(8).Index(0))
	v.QuantityDistributedComment = NewFT(input.Index(9).Index(0))
	v.QuantityInUse = NewCQ(input.Index(10).Index(0))
	v.QuantityInUseMethod = NewID(input.Index(11).Index(0))
	v.QuantityInUseComment = NewFT(input.Index(12).Index(0))
	v.NumberOfProductExperienceReportsFiledByFacility = NewNMSlice(input.Index(13))
	v.NumberOfProductExperienceReportsFiledByDistributor = NewNMSlice(input.Index(14))
	return v
}

// NewPSHSlice will construct a slice of type PSH
func NewPSHSlice(input hl7.HL7Type) []PSH {
	values := make([]PSH, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPSH(input.Index(i))
	}
	return values
}

type DSC struct {
	ContinuationPointer ST
	ContinuationStyle   ID
}

// NewDSC creates an implementation of DSC
func NewDSC(input hl7.HL7Type) DSC {
	v := DSC{}
	v.ContinuationPointer = NewST(input.Index(1).Index(0))
	v.ContinuationStyle = NewID(input.Index(2).Index(0))
	return v
}

// NewDSCSlice will construct a slice of type DSC
func NewDSCSlice(input hl7.HL7Type) []DSC {
	values := make([]DSC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDSC(input.Index(i))
	}
	return values
}

type PD1 struct {
	LivingDependency                        []CWE
	LivingArrangement                       CWE
	PatientPrimaryFacility                  []XON
	PatientPrimaryCareProviderNameIDNo      ST
	StudentIndicator                        CWE
	Handicap                                CWE
	LivingWillCode                          CWE
	OrganDonorCode                          CWE
	SeparateBill                            ID
	DuplicatePatient                        []CX
	PublicityCode                           CWE
	ProtectionIndicator                     ID
	ProtectionIndicatorEffectiveDate        DT
	PlaceOfWorship                          []XON
	AdvanceDirectiveCode                    []CWE
	ImmunizationRegistryStatus              CWE
	ImmunizationRegistryStatusEffectiveDate DT
	PublicityCodeEffectiveDate              DT
	MilitaryBranch                          CWE
	MilitaryRankGrade                       CWE
	MilitaryStatus                          CWE
	AdvanceDirectiveLastVerifiedDate        DT
}

// NewPD1 creates an implementation of PD1
func NewPD1(input hl7.HL7Type) PD1 {
	v := PD1{}
	v.LivingDependency = NewCWESlice(input.Index(1))
	v.LivingArrangement = NewCWE(input.Index(2).Index(0))
	v.PatientPrimaryFacility = NewXONSlice(input.Index(3))
	v.PatientPrimaryCareProviderNameIDNo = NewST(input.Index(4).Index(0))
	v.StudentIndicator = NewCWE(input.Index(5).Index(0))
	v.Handicap = NewCWE(input.Index(6).Index(0))
	v.LivingWillCode = NewCWE(input.Index(7).Index(0))
	v.OrganDonorCode = NewCWE(input.Index(8).Index(0))
	v.SeparateBill = NewID(input.Index(9).Index(0))
	v.DuplicatePatient = NewCXSlice(input.Index(10))
	v.PublicityCode = NewCWE(input.Index(11).Index(0))
	v.ProtectionIndicator = NewID(input.Index(12).Index(0))
	v.ProtectionIndicatorEffectiveDate = NewDT(input.Index(13).Index(0))
	v.PlaceOfWorship = NewXONSlice(input.Index(14))
	v.AdvanceDirectiveCode = NewCWESlice(input.Index(15))
	v.ImmunizationRegistryStatus = NewCWE(input.Index(16).Index(0))
	v.ImmunizationRegistryStatusEffectiveDate = NewDT(input.Index(17).Index(0))
	v.PublicityCodeEffectiveDate = NewDT(input.Index(18).Index(0))
	v.MilitaryBranch = NewCWE(input.Index(19).Index(0))
	v.MilitaryRankGrade = NewCWE(input.Index(20).Index(0))
	v.MilitaryStatus = NewCWE(input.Index(21).Index(0))
	v.AdvanceDirectiveLastVerifiedDate = NewDT(input.Index(22).Index(0))
	return v
}

// NewPD1Slice will construct a slice of type PD1
func NewPD1Slice(input hl7.HL7Type) []PD1 {
	values := make([]PD1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPD1(input.Index(i))
	}
	return values
}

type SAC struct {
	ExternalAccessionIdentifier       EI
	AccessionIdentifier               EI
	ContainerIdentifier               EI
	PrimaryparentContainerIdentifier  EI
	EquipmentContainerIdentifier      EI
	SpecimenSource                    ST
	RegistrationDateTime              DTM
	ContainerStatus                   CWE
	CarrierType                       CWE
	CarrierIdentifier                 EI
	PositionInCarrier                 NA
	TrayTypeSac                       CWE
	TrayIdentifier                    EI
	PositionInTray                    NA
	Location                          []CWE
	ContainerHeight                   NM
	ContainerDiameter                 NM
	BarrierDelta                      NM
	BottomDelta                       NM
	ContainerHeightDiameterDeltaUnits CWE
	ContainerVolume                   NM
	AvailableSpecimenVolume           NM
	InitialSpecimenVolume             NM
	VolumeUnits                       CWE
	SeparatorType                     CWE
	CapType                           CWE
	Additive                          []CWE
	SpecimenComponent                 CWE
	DilutionFactor                    SN
	Treatment                         CWE
	Temperature                       SN
	HemolysisIndex                    NM
	HemolysisIndexUnits               CWE
	LipemiaIndex                      NM
	LipemiaIndexUnits                 CWE
	IcterusIndex                      NM
	IcterusIndexUnits                 CWE
	FibrinIndex                       NM
	FibrinIndexUnits                  CWE
	SystemInducedContaminants         []CWE
	DrugInterference                  []CWE
	ArtificialBlood                   CWE
	SpecialHandlingCode               []CWE
	OtherEnvironmentalFactors         []CWE
}

// NewSAC creates an implementation of SAC
func NewSAC(input hl7.HL7Type) SAC {
	v := SAC{}
	v.ExternalAccessionIdentifier = NewEI(input.Index(1).Index(0))
	v.AccessionIdentifier = NewEI(input.Index(2).Index(0))
	v.ContainerIdentifier = NewEI(input.Index(3).Index(0))
	v.PrimaryparentContainerIdentifier = NewEI(input.Index(4).Index(0))
	v.EquipmentContainerIdentifier = NewEI(input.Index(5).Index(0))
	v.SpecimenSource = NewST(input.Index(6).Index(0))
	v.RegistrationDateTime = NewDTM(input.Index(7).Index(0))
	v.ContainerStatus = NewCWE(input.Index(8).Index(0))
	v.CarrierType = NewCWE(input.Index(9).Index(0))
	v.CarrierIdentifier = NewEI(input.Index(10).Index(0))
	v.PositionInCarrier = NewNA(input.Index(11).Index(0))
	v.TrayTypeSac = NewCWE(input.Index(12).Index(0))
	v.TrayIdentifier = NewEI(input.Index(13).Index(0))
	v.PositionInTray = NewNA(input.Index(14).Index(0))
	v.Location = NewCWESlice(input.Index(15))
	v.ContainerHeight = NewNM(input.Index(16).Index(0))
	v.ContainerDiameter = NewNM(input.Index(17).Index(0))
	v.BarrierDelta = NewNM(input.Index(18).Index(0))
	v.BottomDelta = NewNM(input.Index(19).Index(0))
	v.ContainerHeightDiameterDeltaUnits = NewCWE(input.Index(20).Index(0))
	v.ContainerVolume = NewNM(input.Index(21).Index(0))
	v.AvailableSpecimenVolume = NewNM(input.Index(22).Index(0))
	v.InitialSpecimenVolume = NewNM(input.Index(23).Index(0))
	v.VolumeUnits = NewCWE(input.Index(24).Index(0))
	v.SeparatorType = NewCWE(input.Index(25).Index(0))
	v.CapType = NewCWE(input.Index(26).Index(0))
	v.Additive = NewCWESlice(input.Index(27))
	v.SpecimenComponent = NewCWE(input.Index(28).Index(0))
	v.DilutionFactor = NewSN(input.Index(29).Index(0))
	v.Treatment = NewCWE(input.Index(30).Index(0))
	v.Temperature = NewSN(input.Index(31).Index(0))
	v.HemolysisIndex = NewNM(input.Index(32).Index(0))
	v.HemolysisIndexUnits = NewCWE(input.Index(33).Index(0))
	v.LipemiaIndex = NewNM(input.Index(34).Index(0))
	v.LipemiaIndexUnits = NewCWE(input.Index(35).Index(0))
	v.IcterusIndex = NewNM(input.Index(36).Index(0))
	v.IcterusIndexUnits = NewCWE(input.Index(37).Index(0))
	v.FibrinIndex = NewNM(input.Index(38).Index(0))
	v.FibrinIndexUnits = NewCWE(input.Index(39).Index(0))
	v.SystemInducedContaminants = NewCWESlice(input.Index(40))
	v.DrugInterference = NewCWESlice(input.Index(41))
	v.ArtificialBlood = NewCWE(input.Index(42).Index(0))
	v.SpecialHandlingCode = NewCWESlice(input.Index(43))
	v.OtherEnvironmentalFactors = NewCWESlice(input.Index(44))
	return v
}

// NewSACSlice will construct a slice of type SAC
func NewSACSlice(input hl7.HL7Type) []SAC {
	values := make([]SAC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSAC(input.Index(i))
	}
	return values
}

type SCH struct {
	PlacerAppointmentID       EI
	FillerAppointmentID       EI
	OccurrenceNumber          NM
	PlacerGroupNumber         EI
	ScheduleID                CWE
	EventReason               CWE
	AppointmentReason         CWE
	AppointmentType           CWE
	AppointmentDuration       NM
	AppointmentDurationUnits  CNE
	AppointmentTimingQuantity ST
	PlacerContactPerson       []XCN
	PlacerContactPhoneNumber  XTN
	PlacerContactAddress      []XAD
	PlacerContactLocation     PL
	FillerContactPerson       []XCN
	FillerContactPhoneNumber  XTN
	FillerContactAddress      []XAD
	FillerContactLocation     PL
	EnteredByPerson           []XCN
	EnteredByPhoneNumber      []XTN
	EnteredByLocation         PL
	ParentPlacerAppointmentID EI
	ParentFillerAppointmentID EI
	FillerStatusCode          CWE
	PlacerOrderNumber         []EI
	FillerOrderNumber         []EI
}

// NewSCH creates an implementation of SCH
func NewSCH(input hl7.HL7Type) SCH {
	v := SCH{}
	v.PlacerAppointmentID = NewEI(input.Index(1).Index(0))
	v.FillerAppointmentID = NewEI(input.Index(2).Index(0))
	v.OccurrenceNumber = NewNM(input.Index(3).Index(0))
	v.PlacerGroupNumber = NewEI(input.Index(4).Index(0))
	v.ScheduleID = NewCWE(input.Index(5).Index(0))
	v.EventReason = NewCWE(input.Index(6).Index(0))
	v.AppointmentReason = NewCWE(input.Index(7).Index(0))
	v.AppointmentType = NewCWE(input.Index(8).Index(0))
	v.AppointmentDuration = NewNM(input.Index(9).Index(0))
	v.AppointmentDurationUnits = NewCNE(input.Index(10).Index(0))
	v.AppointmentTimingQuantity = NewST(input.Index(11).Index(0))
	v.PlacerContactPerson = NewXCNSlice(input.Index(12))
	v.PlacerContactPhoneNumber = NewXTN(input.Index(13).Index(0))
	v.PlacerContactAddress = NewXADSlice(input.Index(14))
	v.PlacerContactLocation = NewPL(input.Index(15).Index(0))
	v.FillerContactPerson = NewXCNSlice(input.Index(16))
	v.FillerContactPhoneNumber = NewXTN(input.Index(17).Index(0))
	v.FillerContactAddress = NewXADSlice(input.Index(18))
	v.FillerContactLocation = NewPL(input.Index(19).Index(0))
	v.EnteredByPerson = NewXCNSlice(input.Index(20))
	v.EnteredByPhoneNumber = NewXTNSlice(input.Index(21))
	v.EnteredByLocation = NewPL(input.Index(22).Index(0))
	v.ParentPlacerAppointmentID = NewEI(input.Index(23).Index(0))
	v.ParentFillerAppointmentID = NewEI(input.Index(24).Index(0))
	v.FillerStatusCode = NewCWE(input.Index(25).Index(0))
	v.PlacerOrderNumber = NewEISlice(input.Index(26))
	v.FillerOrderNumber = NewEISlice(input.Index(27))
	return v
}

// NewSCHSlice will construct a slice of type SCH
func NewSCHSlice(input hl7.HL7Type) []SCH {
	values := make([]SCH, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSCH(input.Index(i))
	}
	return values
}

type EQU struct {
	EquipmentInstanceIdentifier EI
	EventDateTime               DTM
	EquipmentState              CWE
	LocalRemoteControlState     CWE
	AlertLevel                  CWE
}

// NewEQU creates an implementation of EQU
func NewEQU(input hl7.HL7Type) EQU {
	v := EQU{}
	v.EquipmentInstanceIdentifier = NewEI(input.Index(1).Index(0))
	v.EventDateTime = NewDTM(input.Index(2).Index(0))
	v.EquipmentState = NewCWE(input.Index(3).Index(0))
	v.LocalRemoteControlState = NewCWE(input.Index(4).Index(0))
	v.AlertLevel = NewCWE(input.Index(5).Index(0))
	return v
}

// NewEQUSlice will construct a slice of type EQU
func NewEQUSlice(input hl7.HL7Type) []EQU {
	values := make([]EQU, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEQU(input.Index(i))
	}
	return values
}

type FHS struct {
	FileFieldSeparator          ST
	FileEncodingCharacters      ST
	FileSendingApplication      HD
	FileSendingFacility         HD
	FileReceivingApplication    HD
	FileReceivingFacility       HD
	FileCreationDateTime        DTM
	FileSecurity                ST
	FileNameID                  ST
	FileHeaderComment           ST
	FileControlID               ST
	ReferenceFileControlID      ST
	FileSendingNetworkAddress   HD
	FileReceivingNetworkAddress HD
}

// NewFHS creates an implementation of FHS
func NewFHS(input hl7.HL7Type) FHS {
	v := FHS{}
	v.FileFieldSeparator = NewST(input.Index(1).Index(0))
	v.FileEncodingCharacters = NewST(input.Index(2).Index(0))
	v.FileSendingApplication = NewHD(input.Index(3).Index(0))
	v.FileSendingFacility = NewHD(input.Index(4).Index(0))
	v.FileReceivingApplication = NewHD(input.Index(5).Index(0))
	v.FileReceivingFacility = NewHD(input.Index(6).Index(0))
	v.FileCreationDateTime = NewDTM(input.Index(7).Index(0))
	v.FileSecurity = NewST(input.Index(8).Index(0))
	v.FileNameID = NewST(input.Index(9).Index(0))
	v.FileHeaderComment = NewST(input.Index(10).Index(0))
	v.FileControlID = NewST(input.Index(11).Index(0))
	v.ReferenceFileControlID = NewST(input.Index(12).Index(0))
	v.FileSendingNetworkAddress = NewHD(input.Index(13).Index(0))
	v.FileReceivingNetworkAddress = NewHD(input.Index(14).Index(0))
	return v
}

// NewFHSSlice will construct a slice of type FHS
func NewFHSSlice(input hl7.HL7Type) []FHS {
	values := make([]FHS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewFHS(input.Index(i))
	}
	return values
}

type NK1 struct {
	SetID                                      SI
	Name                                       []XPN
	Relationship                               CWE
	Address                                    []XAD
	PhoneNumber                                []XTN
	BusinessPhoneNumber                        []XTN
	ContactRole                                CWE
	StartDate                                  DT
	EndDate                                    DT
	NextOfKinAssociatedPartiesJobTitle         ST
	NextOfKinAssociatedPartiesJobCodeClass     JCC
	NextOfKinAssociatedPartiesEmployeeNumber   CX
	OrganizationNameNk1                        []XON
	MaritalStatus                              CWE
	AdministrativeSex                          CWE
	DateTimeOfBirth                            DTM
	LivingDependency                           []CWE
	AmbulatoryStatus                           []CWE
	Citizenship                                []CWE
	PrimaryLanguage                            CWE
	LivingArrangement                          CWE
	PublicityCode                              CWE
	ProtectionIndicator                        ID
	StudentIndicator                           CWE
	Religion                                   CWE
	MothersMaidenName                          []XPN
	Nationality                                CWE
	EthnicGroup                                []CWE
	ContactReason                              []CWE
	ContactPersonsName                         []XPN
	ContactPersonsTelephoneNumber              []XTN
	ContactPersonsAddress                      []XAD
	NextOfKinAssociatedPartysIdentifiers       []CX
	JobStatus                                  CWE
	Race                                       []CWE
	Handicap                                   CWE
	ContactPersonSocialSecurityNumber          ST
	NextOfKinBirthPlace                        ST
	VipIndicator                               CWE
	NextOfKinTelecommunicationInformation      XTN
	ContactPersonsTelecommunicationInformation XTN
}

// NewNK1 creates an implementation of NK1
func NewNK1(input hl7.HL7Type) NK1 {
	v := NK1{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.Name = NewXPNSlice(input.Index(2))
	v.Relationship = NewCWE(input.Index(3).Index(0))
	v.Address = NewXADSlice(input.Index(4))
	v.PhoneNumber = NewXTNSlice(input.Index(5))
	v.BusinessPhoneNumber = NewXTNSlice(input.Index(6))
	v.ContactRole = NewCWE(input.Index(7).Index(0))
	v.StartDate = NewDT(input.Index(8).Index(0))
	v.EndDate = NewDT(input.Index(9).Index(0))
	v.NextOfKinAssociatedPartiesJobTitle = NewST(input.Index(10).Index(0))
	v.NextOfKinAssociatedPartiesJobCodeClass = NewJCC(input.Index(11).Index(0))
	v.NextOfKinAssociatedPartiesEmployeeNumber = NewCX(input.Index(12).Index(0))
	v.OrganizationNameNk1 = NewXONSlice(input.Index(13))
	v.MaritalStatus = NewCWE(input.Index(14).Index(0))
	v.AdministrativeSex = NewCWE(input.Index(15).Index(0))
	v.DateTimeOfBirth = NewDTM(input.Index(16).Index(0))
	v.LivingDependency = NewCWESlice(input.Index(17))
	v.AmbulatoryStatus = NewCWESlice(input.Index(18))
	v.Citizenship = NewCWESlice(input.Index(19))
	v.PrimaryLanguage = NewCWE(input.Index(20).Index(0))
	v.LivingArrangement = NewCWE(input.Index(21).Index(0))
	v.PublicityCode = NewCWE(input.Index(22).Index(0))
	v.ProtectionIndicator = NewID(input.Index(23).Index(0))
	v.StudentIndicator = NewCWE(input.Index(24).Index(0))
	v.Religion = NewCWE(input.Index(25).Index(0))
	v.MothersMaidenName = NewXPNSlice(input.Index(26))
	v.Nationality = NewCWE(input.Index(27).Index(0))
	v.EthnicGroup = NewCWESlice(input.Index(28))
	v.ContactReason = NewCWESlice(input.Index(29))
	v.ContactPersonsName = NewXPNSlice(input.Index(30))
	v.ContactPersonsTelephoneNumber = NewXTNSlice(input.Index(31))
	v.ContactPersonsAddress = NewXADSlice(input.Index(32))
	v.NextOfKinAssociatedPartysIdentifiers = NewCXSlice(input.Index(33))
	v.JobStatus = NewCWE(input.Index(34).Index(0))
	v.Race = NewCWESlice(input.Index(35))
	v.Handicap = NewCWE(input.Index(36).Index(0))
	v.ContactPersonSocialSecurityNumber = NewST(input.Index(37).Index(0))
	v.NextOfKinBirthPlace = NewST(input.Index(38).Index(0))
	v.VipIndicator = NewCWE(input.Index(39).Index(0))
	v.NextOfKinTelecommunicationInformation = NewXTN(input.Index(40).Index(0))
	v.ContactPersonsTelecommunicationInformation = NewXTN(input.Index(41).Index(0))
	return v
}

// NewNK1Slice will construct a slice of type NK1
func NewNK1Slice(input hl7.HL7Type) []NK1 {
	values := make([]NK1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewNK1(input.Index(i))
	}
	return values
}

type VAR struct {
	VarianceInstanceID     EI
	DocumentedDateTime     DTM
	StatedVarianceDateTime DTM
	VarianceOriginator     []XCN
	VarianceClassification CWE
	VarianceDescription    []ST
}

// NewVAR creates an implementation of VAR
func NewVAR(input hl7.HL7Type) VAR {
	v := VAR{}
	v.VarianceInstanceID = NewEI(input.Index(1).Index(0))
	v.DocumentedDateTime = NewDTM(input.Index(2).Index(0))
	v.StatedVarianceDateTime = NewDTM(input.Index(3).Index(0))
	v.VarianceOriginator = NewXCNSlice(input.Index(4))
	v.VarianceClassification = NewCWE(input.Index(5).Index(0))
	v.VarianceDescription = NewSTSlice(input.Index(6))
	return v
}

// NewVARSlice will construct a slice of type VAR
func NewVARSlice(input hl7.HL7Type) []VAR {
	values := make([]VAR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewVAR(input.Index(i))
	}
	return values
}

type CM1 struct {
	SetID                   SI
	StudyPhaseIdentifier    CWE
	DescriptionOfStudyPhase ST
}

// NewCM1 creates an implementation of CM1
func NewCM1(input hl7.HL7Type) CM1 {
	v := CM1{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.StudyPhaseIdentifier = NewCWE(input.Index(2).Index(0))
	v.DescriptionOfStudyPhase = NewST(input.Index(3).Index(0))
	return v
}

// NewCM1Slice will construct a slice of type CM1
func NewCM1Slice(input hl7.HL7Type) []CM1 {
	values := make([]CM1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCM1(input.Index(i))
	}
	return values
}

type MFI struct {
	MasterFileIdentifier            CWE
	MasterFileApplicationIdentifier []HD
	FileLevelEventCode              ID
	EnteredDateTime                 DTM
	EffectiveDateTime               DTM
	ResponseLevelCode               ID
}

// NewMFI creates an implementation of MFI
func NewMFI(input hl7.HL7Type) MFI {
	v := MFI{}
	v.MasterFileIdentifier = NewCWE(input.Index(1).Index(0))
	v.MasterFileApplicationIdentifier = NewHDSlice(input.Index(2))
	v.FileLevelEventCode = NewID(input.Index(3).Index(0))
	v.EnteredDateTime = NewDTM(input.Index(4).Index(0))
	v.EffectiveDateTime = NewDTM(input.Index(5).Index(0))
	v.ResponseLevelCode = NewID(input.Index(6).Index(0))
	return v
}

// NewMFISlice will construct a slice of type MFI
func NewMFISlice(input hl7.HL7Type) []MFI {
	values := make([]MFI, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFI(input.Index(i))
	}
	return values
}

type INV struct {
	SubstanceIdentifier          CWE
	SubstanceStatus              []CWE
	SubstanceType                CWE
	InventoryContainerIdentifier CWE
	ContainerCarrierIdentifier   CWE
	PositionOnCarrier            CWE
	InitialQuantity              NM
	CurrentQuantity              NM
	AvailableQuantity            NM
	ConsumptionQuantity          NM
	QuantityUnits                CWE
	ExpirationDateTime           DTM
	FirstUsedDateTime            DTM
	OnBoardStabilityDuration     ST
	TestFluidIdentifiers         []CWE
	ManufacturerLotNumber        ST
	ManufacturerIdentifier       CWE
	SupplierIdentifier           CWE
	OnBoardStabilityTime         CQ
	TargetValue                  CQ
}

// NewINV creates an implementation of INV
func NewINV(input hl7.HL7Type) INV {
	v := INV{}
	v.SubstanceIdentifier = NewCWE(input.Index(1).Index(0))
	v.SubstanceStatus = NewCWESlice(input.Index(2))
	v.SubstanceType = NewCWE(input.Index(3).Index(0))
	v.InventoryContainerIdentifier = NewCWE(input.Index(4).Index(0))
	v.ContainerCarrierIdentifier = NewCWE(input.Index(5).Index(0))
	v.PositionOnCarrier = NewCWE(input.Index(6).Index(0))
	v.InitialQuantity = NewNM(input.Index(7).Index(0))
	v.CurrentQuantity = NewNM(input.Index(8).Index(0))
	v.AvailableQuantity = NewNM(input.Index(9).Index(0))
	v.ConsumptionQuantity = NewNM(input.Index(10).Index(0))
	v.QuantityUnits = NewCWE(input.Index(11).Index(0))
	v.ExpirationDateTime = NewDTM(input.Index(12).Index(0))
	v.FirstUsedDateTime = NewDTM(input.Index(13).Index(0))
	v.OnBoardStabilityDuration = NewST(input.Index(14).Index(0))
	v.TestFluidIdentifiers = NewCWESlice(input.Index(15))
	v.ManufacturerLotNumber = NewST(input.Index(16).Index(0))
	v.ManufacturerIdentifier = NewCWE(input.Index(17).Index(0))
	v.SupplierIdentifier = NewCWE(input.Index(18).Index(0))
	v.OnBoardStabilityTime = NewCQ(input.Index(19).Index(0))
	v.TargetValue = NewCQ(input.Index(20).Index(0))
	return v
}

// NewINVSlice will construct a slice of type INV
func NewINVSlice(input hl7.HL7Type) []INV {
	values := make([]INV, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewINV(input.Index(i))
	}
	return values
}

type PR1 struct {
	SetID                           SI
	ProcedureCodingMethod           ST
	ProcedureCode                   CNE
	ProcedureDescription            ST
	ProcedureDateTime               DTM
	ProcedureFunctionalType         CWE
	ProcedureMinutes                NM
	Anesthesiologist                ST
	AnesthesiaCode                  CWE
	AnesthesiaMinutes               NM
	Surgeon                         ST
	ProcedurePractitioner           ST
	ConsentCode                     CWE
	ProcedurePriority               NM
	AssociatedDiagnosisCode         CWE
	ProcedureCodeModifier           []CNE
	ProcedureDrgType                CWE
	TissueTypeCode                  []CWE
	ProcedureIdentifier             EI
	ProcedureActionCode             ID
	DrgProcedureDeterminationStatus CWE
	DrgProcedureRelevance           CWE
	TreatingOrganizationalUnit      []PL
	RespiratoryWithinSurgery        ID
	ParentProcedureID               EI
}

// NewPR1 creates an implementation of PR1
func NewPR1(input hl7.HL7Type) PR1 {
	v := PR1{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.ProcedureCodingMethod = NewST(input.Index(2).Index(0))
	v.ProcedureCode = NewCNE(input.Index(3).Index(0))
	v.ProcedureDescription = NewST(input.Index(4).Index(0))
	v.ProcedureDateTime = NewDTM(input.Index(5).Index(0))
	v.ProcedureFunctionalType = NewCWE(input.Index(6).Index(0))
	v.ProcedureMinutes = NewNM(input.Index(7).Index(0))
	v.Anesthesiologist = NewST(input.Index(8).Index(0))
	v.AnesthesiaCode = NewCWE(input.Index(9).Index(0))
	v.AnesthesiaMinutes = NewNM(input.Index(10).Index(0))
	v.Surgeon = NewST(input.Index(11).Index(0))
	v.ProcedurePractitioner = NewST(input.Index(12).Index(0))
	v.ConsentCode = NewCWE(input.Index(13).Index(0))
	v.ProcedurePriority = NewNM(input.Index(14).Index(0))
	v.AssociatedDiagnosisCode = NewCWE(input.Index(15).Index(0))
	v.ProcedureCodeModifier = NewCNESlice(input.Index(16))
	v.ProcedureDrgType = NewCWE(input.Index(17).Index(0))
	v.TissueTypeCode = NewCWESlice(input.Index(18))
	v.ProcedureIdentifier = NewEI(input.Index(19).Index(0))
	v.ProcedureActionCode = NewID(input.Index(20).Index(0))
	v.DrgProcedureDeterminationStatus = NewCWE(input.Index(21).Index(0))
	v.DrgProcedureRelevance = NewCWE(input.Index(22).Index(0))
	v.TreatingOrganizationalUnit = NewPLSlice(input.Index(23))
	v.RespiratoryWithinSurgery = NewID(input.Index(24).Index(0))
	v.ParentProcedureID = NewEI(input.Index(25).Index(0))
	return v
}

// NewPR1Slice will construct a slice of type PR1
func NewPR1Slice(input hl7.HL7Type) []PR1 {
	values := make([]PR1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPR1(input.Index(i))
	}
	return values
}

type DON struct {
	DonationIdentificationNumberDIN   EI
	DonationType                      CNE
	PhlebotomyStartDateTime           DTM
	PhlebotomyEndDateTime             DTM
	DonationDuration                  NM
	DonationDurationUnits             CNE
	IntendedProcedureType             []CNE
	ActualProcedureType               []CNE
	DonorEligibilityFlag              ID
	DonorEligibilityProcedureType     []CNE
	DonorEligibilityDate              DTM
	ProcessInterruption               CNE
	ProcessInterruptionReason         CNE
	PhlebotomyIssue                   []CNE
	IntendedRecipientBloodRelative    ID
	IntendedRecipientName             XPN
	IntendedRecipientDOB              DTM
	IntendedRecipientFacility         XON
	IntendedRecipientProcedureDate    DTM
	IntendedRecipientOrderingProvider XPN
	PhlebotomyStatus                  CNE
	ArmStick                          CNE
	BleedStartPhlebotomist            XPN
	BleedEndPhlebotomist              XPN
	AphaeresisTypeMachine             ST
	AphaeresisMachineSerialNumber     ST
	DonorReaction                     ID
	FinalReviewStaffID                XPN
	FinalReviewDateTime               DTM
	NumberOfTubesCollected            NM
	DonationSampleIdentifier          []EI
	DonationAcceptStaff               XCN
	DonationMaterialReviewStaff       []XCN
}

// NewDON creates an implementation of DON
func NewDON(input hl7.HL7Type) DON {
	v := DON{}
	v.DonationIdentificationNumberDIN = NewEI(input.Index(1).Index(0))
	v.DonationType = NewCNE(input.Index(2).Index(0))
	v.PhlebotomyStartDateTime = NewDTM(input.Index(3).Index(0))
	v.PhlebotomyEndDateTime = NewDTM(input.Index(4).Index(0))
	v.DonationDuration = NewNM(input.Index(5).Index(0))
	v.DonationDurationUnits = NewCNE(input.Index(6).Index(0))
	v.IntendedProcedureType = NewCNESlice(input.Index(7))
	v.ActualProcedureType = NewCNESlice(input.Index(8))
	v.DonorEligibilityFlag = NewID(input.Index(9).Index(0))
	v.DonorEligibilityProcedureType = NewCNESlice(input.Index(10))
	v.DonorEligibilityDate = NewDTM(input.Index(11).Index(0))
	v.ProcessInterruption = NewCNE(input.Index(12).Index(0))
	v.ProcessInterruptionReason = NewCNE(input.Index(13).Index(0))
	v.PhlebotomyIssue = NewCNESlice(input.Index(14))
	v.IntendedRecipientBloodRelative = NewID(input.Index(15).Index(0))
	v.IntendedRecipientName = NewXPN(input.Index(16).Index(0))
	v.IntendedRecipientDOB = NewDTM(input.Index(17).Index(0))
	v.IntendedRecipientFacility = NewXON(input.Index(18).Index(0))
	v.IntendedRecipientProcedureDate = NewDTM(input.Index(19).Index(0))
	v.IntendedRecipientOrderingProvider = NewXPN(input.Index(20).Index(0))
	v.PhlebotomyStatus = NewCNE(input.Index(21).Index(0))
	v.ArmStick = NewCNE(input.Index(22).Index(0))
	v.BleedStartPhlebotomist = NewXPN(input.Index(23).Index(0))
	v.BleedEndPhlebotomist = NewXPN(input.Index(24).Index(0))
	v.AphaeresisTypeMachine = NewST(input.Index(25).Index(0))
	v.AphaeresisMachineSerialNumber = NewST(input.Index(26).Index(0))
	v.DonorReaction = NewID(input.Index(27).Index(0))
	v.FinalReviewStaffID = NewXPN(input.Index(28).Index(0))
	v.FinalReviewDateTime = NewDTM(input.Index(29).Index(0))
	v.NumberOfTubesCollected = NewNM(input.Index(30).Index(0))
	v.DonationSampleIdentifier = NewEISlice(input.Index(31))
	v.DonationAcceptStaff = NewXCN(input.Index(32).Index(0))
	v.DonationMaterialReviewStaff = NewXCNSlice(input.Index(33))
	return v
}

// NewDONSlice will construct a slice of type DON
func NewDONSlice(input hl7.HL7Type) []DON {
	values := make([]DON, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDON(input.Index(i))
	}
	return values
}

type GP1 struct {
	TypeOfBillCode              CWE
	RevenueCode                 []CWE
	OverallClaimDispositionCode CWE
	OceEditsPerVisitCode        []CWE
	OutlierCost                 CP
}

// NewGP1 creates an implementation of GP1
func NewGP1(input hl7.HL7Type) GP1 {
	v := GP1{}
	v.TypeOfBillCode = NewCWE(input.Index(1).Index(0))
	v.RevenueCode = NewCWESlice(input.Index(2))
	v.OverallClaimDispositionCode = NewCWE(input.Index(3).Index(0))
	v.OceEditsPerVisitCode = NewCWESlice(input.Index(4))
	v.OutlierCost = NewCP(input.Index(5).Index(0))
	return v
}

// NewGP1Slice will construct a slice of type GP1
func NewGP1Slice(input hl7.HL7Type) []GP1 {
	values := make([]GP1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewGP1(input.Index(i))
	}
	return values
}

type TCC struct {
	UniversalServiceIdentifier            CWE
	EquipmentTestApplicationIdentifier    EI
	SpecimenSource                        ST
	AutoDilutionFactorDefault             SN
	RerunDilutionFactorDefault            SN
	PreDilutionFactorDefault              SN
	EndogenousContentOfPreDilutionDiluent SN
	InventoryLimitsWarningLevel           NM
	AutomaticRerunAllowed                 ID
	AutomaticRepeatAllowed                ID
	AutomaticReflexAllowed                ID
	EquipmentDynamicRange                 SN
	Units                                 CWE
	ProcessingType                        CWE
	TestCriticality                       CWE
}

// NewTCC creates an implementation of TCC
func NewTCC(input hl7.HL7Type) TCC {
	v := TCC{}
	v.UniversalServiceIdentifier = NewCWE(input.Index(1).Index(0))
	v.EquipmentTestApplicationIdentifier = NewEI(input.Index(2).Index(0))
	v.SpecimenSource = NewST(input.Index(3).Index(0))
	v.AutoDilutionFactorDefault = NewSN(input.Index(4).Index(0))
	v.RerunDilutionFactorDefault = NewSN(input.Index(5).Index(0))
	v.PreDilutionFactorDefault = NewSN(input.Index(6).Index(0))
	v.EndogenousContentOfPreDilutionDiluent = NewSN(input.Index(7).Index(0))
	v.InventoryLimitsWarningLevel = NewNM(input.Index(8).Index(0))
	v.AutomaticRerunAllowed = NewID(input.Index(9).Index(0))
	v.AutomaticRepeatAllowed = NewID(input.Index(10).Index(0))
	v.AutomaticReflexAllowed = NewID(input.Index(11).Index(0))
	v.EquipmentDynamicRange = NewSN(input.Index(12).Index(0))
	v.Units = NewCWE(input.Index(13).Index(0))
	v.ProcessingType = NewCWE(input.Index(14).Index(0))
	v.TestCriticality = NewCWE(input.Index(15).Index(0))
	return v
}

// NewTCCSlice will construct a slice of type TCC
func NewTCCSlice(input hl7.HL7Type) []TCC {
	values := make([]TCC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewTCC(input.Index(i))
	}
	return values
}

type PAC struct {
	SetID                   SI
	PackageID               EI
	ParentPackageID         EI
	PositionInParentPackage NA
	PackageType             CWE
	PackageCondition        []CWE
	PackageHandlingCode     []CWE
	PackageRiskCode         []CWE
}

// NewPAC creates an implementation of PAC
func NewPAC(input hl7.HL7Type) PAC {
	v := PAC{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.PackageID = NewEI(input.Index(2).Index(0))
	v.ParentPackageID = NewEI(input.Index(3).Index(0))
	v.PositionInParentPackage = NewNA(input.Index(4).Index(0))
	v.PackageType = NewCWE(input.Index(5).Index(0))
	v.PackageCondition = NewCWESlice(input.Index(6))
	v.PackageHandlingCode = NewCWESlice(input.Index(7))
	v.PackageRiskCode = NewCWESlice(input.Index(8))
	return v
}

// NewPACSlice will construct a slice of type PAC
func NewPACSlice(input hl7.HL7Type) []PAC {
	values := make([]PAC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPAC(input.Index(i))
	}
	return values
}

type DG1 struct {
	SetID                           SI
	DiagnosisCodingMethod           ST
	DiagnosisCodeDg1                CWE
	DiagnosisDescription            ST
	DiagnosisDateTime               DTM
	DiagnosisType                   CWE
	MajorDiagnosticCategory         ST
	DiagnosticRelatedGroup          ST
	DrgApprovalIndicator            ST
	DrgGrouperReviewCode            ST
	OutlierType                     ST
	OutlierDays                     ST
	OutlierCost                     ST
	GrouperVersionAndType           ST
	DiagnosisPriority               NM
	DiagnosingClinician             []XCN
	DiagnosisClassification         CWE
	ConfidentialIndicator           ID
	AttestationDateTime             DTM
	DiagnosisIdentifier             EI
	DiagnosisActionCode             ID
	ParentDiagnosis                 EI
	DrgCclValueCode                 CWE
	DrgGroupingUsage                ID
	DrgDiagnosisDeterminationStatus CWE
	PresentOnAdmissionpoaIndicator  CWE
}

// NewDG1 creates an implementation of DG1
func NewDG1(input hl7.HL7Type) DG1 {
	v := DG1{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.DiagnosisCodingMethod = NewST(input.Index(2).Index(0))
	v.DiagnosisCodeDg1 = NewCWE(input.Index(3).Index(0))
	v.DiagnosisDescription = NewST(input.Index(4).Index(0))
	v.DiagnosisDateTime = NewDTM(input.Index(5).Index(0))
	v.DiagnosisType = NewCWE(input.Index(6).Index(0))
	v.MajorDiagnosticCategory = NewST(input.Index(7).Index(0))
	v.DiagnosticRelatedGroup = NewST(input.Index(8).Index(0))
	v.DrgApprovalIndicator = NewST(input.Index(9).Index(0))
	v.DrgGrouperReviewCode = NewST(input.Index(10).Index(0))
	v.OutlierType = NewST(input.Index(11).Index(0))
	v.OutlierDays = NewST(input.Index(12).Index(0))
	v.OutlierCost = NewST(input.Index(13).Index(0))
	v.GrouperVersionAndType = NewST(input.Index(14).Index(0))
	v.DiagnosisPriority = NewNM(input.Index(15).Index(0))
	v.DiagnosingClinician = NewXCNSlice(input.Index(16))
	v.DiagnosisClassification = NewCWE(input.Index(17).Index(0))
	v.ConfidentialIndicator = NewID(input.Index(18).Index(0))
	v.AttestationDateTime = NewDTM(input.Index(19).Index(0))
	v.DiagnosisIdentifier = NewEI(input.Index(20).Index(0))
	v.DiagnosisActionCode = NewID(input.Index(21).Index(0))
	v.ParentDiagnosis = NewEI(input.Index(22).Index(0))
	v.DrgCclValueCode = NewCWE(input.Index(23).Index(0))
	v.DrgGroupingUsage = NewID(input.Index(24).Index(0))
	v.DrgDiagnosisDeterminationStatus = NewCWE(input.Index(25).Index(0))
	v.PresentOnAdmissionpoaIndicator = NewCWE(input.Index(26).Index(0))
	return v
}

// NewDG1Slice will construct a slice of type DG1
func NewDG1Slice(input hl7.HL7Type) []DG1 {
	values := make([]DG1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDG1(input.Index(i))
	}
	return values
}

type ODT struct {
	TrayType        CWE
	ServicePeriod   []CWE
	TextInstruction ST
}

// NewODT creates an implementation of ODT
func NewODT(input hl7.HL7Type) ODT {
	v := ODT{}
	v.TrayType = NewCWE(input.Index(1).Index(0))
	v.ServicePeriod = NewCWESlice(input.Index(2))
	v.TextInstruction = NewST(input.Index(3).Index(0))
	return v
}

// NewODTSlice will construct a slice of type ODT
func NewODTSlice(input hl7.HL7Type) []ODT {
	values := make([]ODT, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewODT(input.Index(i))
	}
	return values
}

type ADD struct {
	AddendumContinuationPointer ST
}

// NewADD creates an implementation of ADD
func NewADD(input hl7.HL7Type) ADD {
	v := ADD{}
	v.AddendumContinuationPointer = NewST(input.Index(1).Index(0))
	return v
}

// NewADDSlice will construct a slice of type ADD
func NewADDSlice(input hl7.HL7Type) []ADD {
	values := make([]ADD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADD(input.Index(i))
	}
	return values
}

type IPR struct {
	IprIdentifier                    EI
	ProviderCrossReferenceIdentifier EI
	PayerCrossReferenceIdentifier    EI
	IprStatus                        CWE
	IprDateTime                      DTM
	AdjudicatedPaidAmount            CP
	ExpectedPaymentDateTime          DTM
	IprChecksum                      ST
}

// NewIPR creates an implementation of IPR
func NewIPR(input hl7.HL7Type) IPR {
	v := IPR{}
	v.IprIdentifier = NewEI(input.Index(1).Index(0))
	v.ProviderCrossReferenceIdentifier = NewEI(input.Index(2).Index(0))
	v.PayerCrossReferenceIdentifier = NewEI(input.Index(3).Index(0))
	v.IprStatus = NewCWE(input.Index(4).Index(0))
	v.IprDateTime = NewDTM(input.Index(5).Index(0))
	v.AdjudicatedPaidAmount = NewCP(input.Index(6).Index(0))
	v.ExpectedPaymentDateTime = NewDTM(input.Index(7).Index(0))
	v.IprChecksum = NewST(input.Index(8).Index(0))
	return v
}

// NewIPRSlice will construct a slice of type IPR
func NewIPRSlice(input hl7.HL7Type) []IPR {
	values := make([]IPR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewIPR(input.Index(i))
	}
	return values
}

type CDO struct {
	SetID                             SI
	ActionCode                        ID
	CumulativeDosageLimit             CQ
	CumulativeDosageLimitTimeInterval CQ
}

// NewCDO creates an implementation of CDO
func NewCDO(input hl7.HL7Type) CDO {
	v := CDO{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.ActionCode = NewID(input.Index(2).Index(0))
	v.CumulativeDosageLimit = NewCQ(input.Index(3).Index(0))
	v.CumulativeDosageLimitTimeInterval = NewCQ(input.Index(4).Index(0))
	return v
}

// NewCDOSlice will construct a slice of type CDO
func NewCDOSlice(input hl7.HL7Type) []CDO {
	values := make([]CDO, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCDO(input.Index(i))
	}
	return values
}

type APR struct {
	TimeSelectionCriteria     []SCV
	ResourceSelectionCriteria []SCV
	LocationSelectionCriteria []SCV
	SlotSpacingCriteria       NM
	FillerOverrideCriteria    []SCV
}

// NewAPR creates an implementation of APR
func NewAPR(input hl7.HL7Type) APR {
	v := APR{}
	v.TimeSelectionCriteria = NewSCVSlice(input.Index(1))
	v.ResourceSelectionCriteria = NewSCVSlice(input.Index(2))
	v.LocationSelectionCriteria = NewSCVSlice(input.Index(3))
	v.SlotSpacingCriteria = NewNM(input.Index(4).Index(0))
	v.FillerOverrideCriteria = NewSCVSlice(input.Index(5))
	return v
}

// NewAPRSlice will construct a slice of type APR
func NewAPRSlice(input hl7.HL7Type) []APR {
	values := make([]APR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewAPR(input.Index(i))
	}
	return values
}

type MFA struct {
	RecordLevelEventCode      ID
	MfnControlID              ST
	EventCompletionDateTime   DTM
	MfnRecordLevelErrorReturn CWE
	PrimaryKeyValueMfa        []Varies
	PrimaryKeyValueTypeMfa    []ID
}

// NewMFA creates an implementation of MFA
func NewMFA(input hl7.HL7Type) MFA {
	v := MFA{}
	v.RecordLevelEventCode = NewID(input.Index(1).Index(0))
	v.MfnControlID = NewST(input.Index(2).Index(0))
	v.EventCompletionDateTime = NewDTM(input.Index(3).Index(0))
	v.MfnRecordLevelErrorReturn = NewCWE(input.Index(4).Index(0))
	v.PrimaryKeyValueMfa = NewVariesSlice(input.Index(5))
	v.PrimaryKeyValueTypeMfa = NewIDSlice(input.Index(6))
	return v
}

// NewMFASlice will construct a slice of type MFA
func NewMFASlice(input hl7.HL7Type) []MFA {
	values := make([]MFA, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFA(input.Index(i))
	}
	return values
}

type VND struct {
	SetID                  SI
	VendorIdentifier       EI
	VendorName             ST
	VendorCatalogNumber    EI
	PrimaryVendorIndicator CNE
}

// NewVND creates an implementation of VND
func NewVND(input hl7.HL7Type) VND {
	v := VND{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.VendorIdentifier = NewEI(input.Index(2).Index(0))
	v.VendorName = NewST(input.Index(3).Index(0))
	v.VendorCatalogNumber = NewEI(input.Index(4).Index(0))
	v.PrimaryVendorIndicator = NewCNE(input.Index(5).Index(0))
	return v
}

// NewVNDSlice will construct a slice of type VND
func NewVNDSlice(input hl7.HL7Type) []VND {
	values := make([]VND, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewVND(input.Index(i))
	}
	return values
}

type RXG struct {
	GiveSubIDCounter                                            NM
	DispenseSubIDCounter                                        NM
	QuantityTiming                                              ST
	GiveCode                                                    CWE
	GiveAmountMinimum                                           NM
	GiveAmountMaximum                                           NM
	GiveUnits                                                   CWE
	GiveDosageForm                                              CWE
	AdministrationNotes                                         []CWE
	SubstitutionStatus                                          ID
	DispenseToLocation                                          ST
	NeedsHumanReview                                            ID
	PharmacyTreatmentSuppliersSpecialAdministrationInstructions []CWE
	GivePertimeUnit                                             ST
	GiveRateAmount                                              ST
	GiveRateUnits                                               CWE
	GiveStrength                                                NM
	GiveStrengthUnits                                           CWE
	SubstanceLotNumber                                          []ST
	SubstanceExpirationDate                                     []DTM
	SubstanceManufacturerName                                   []CWE
	Indication                                                  []CWE
	GiveDrugStrengthVolume                                      NM
	GiveDrugStrengthVolumeUnits                                 CWE
	GiveBarcodeIdentifier                                       CWE
	PharmacyOrderType                                           ID
	DispenseToPharmacy                                          CWE
	DispenseToPharmacyAddress                                   XAD
	DeliverToPatientLocation                                    PL
	DeliverToAddress                                            XAD
	GiveTagIdentifier                                           []EI
	DispenseAmount                                              NM
	DispenseUnits                                               CWE
}

// NewRXG creates an implementation of RXG
func NewRXG(input hl7.HL7Type) RXG {
	v := RXG{}
	v.GiveSubIDCounter = NewNM(input.Index(1).Index(0))
	v.DispenseSubIDCounter = NewNM(input.Index(2).Index(0))
	v.QuantityTiming = NewST(input.Index(3).Index(0))
	v.GiveCode = NewCWE(input.Index(4).Index(0))
	v.GiveAmountMinimum = NewNM(input.Index(5).Index(0))
	v.GiveAmountMaximum = NewNM(input.Index(6).Index(0))
	v.GiveUnits = NewCWE(input.Index(7).Index(0))
	v.GiveDosageForm = NewCWE(input.Index(8).Index(0))
	v.AdministrationNotes = NewCWESlice(input.Index(9))
	v.SubstitutionStatus = NewID(input.Index(10).Index(0))
	v.DispenseToLocation = NewST(input.Index(11).Index(0))
	v.NeedsHumanReview = NewID(input.Index(12).Index(0))
	v.PharmacyTreatmentSuppliersSpecialAdministrationInstructions = NewCWESlice(input.Index(13))
	v.GivePertimeUnit = NewST(input.Index(14).Index(0))
	v.GiveRateAmount = NewST(input.Index(15).Index(0))
	v.GiveRateUnits = NewCWE(input.Index(16).Index(0))
	v.GiveStrength = NewNM(input.Index(17).Index(0))
	v.GiveStrengthUnits = NewCWE(input.Index(18).Index(0))
	v.SubstanceLotNumber = NewSTSlice(input.Index(19))
	v.SubstanceExpirationDate = NewDTMSlice(input.Index(20))
	v.SubstanceManufacturerName = NewCWESlice(input.Index(21))
	v.Indication = NewCWESlice(input.Index(22))
	v.GiveDrugStrengthVolume = NewNM(input.Index(23).Index(0))
	v.GiveDrugStrengthVolumeUnits = NewCWE(input.Index(24).Index(0))
	v.GiveBarcodeIdentifier = NewCWE(input.Index(25).Index(0))
	v.PharmacyOrderType = NewID(input.Index(26).Index(0))
	v.DispenseToPharmacy = NewCWE(input.Index(27).Index(0))
	v.DispenseToPharmacyAddress = NewXAD(input.Index(28).Index(0))
	v.DeliverToPatientLocation = NewPL(input.Index(29).Index(0))
	v.DeliverToAddress = NewXAD(input.Index(30).Index(0))
	v.GiveTagIdentifier = NewEISlice(input.Index(31))
	v.DispenseAmount = NewNM(input.Index(32).Index(0))
	v.DispenseUnits = NewCWE(input.Index(33).Index(0))
	return v
}

// NewRXGSlice will construct a slice of type RXG
func NewRXGSlice(input hl7.HL7Type) []RXG {
	values := make([]RXG, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRXG(input.Index(i))
	}
	return values
}

type EDU struct {
	SetID                                       SI
	AcademicDegree                              CWE
	AcademicDegreeProgramDateRange              DR
	AcademicDegreeProgramParticipationDateRange DR
	AcademicDegreeGrantedDate                   DT
	School                                      XON
	SchoolTypeCode                              CWE
	SchoolAddress                               XAD
	MajorFieldOfStudy                           []CWE
}

// NewEDU creates an implementation of EDU
func NewEDU(input hl7.HL7Type) EDU {
	v := EDU{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.AcademicDegree = NewCWE(input.Index(2).Index(0))
	v.AcademicDegreeProgramDateRange = NewDR(input.Index(3).Index(0))
	v.AcademicDegreeProgramParticipationDateRange = NewDR(input.Index(4).Index(0))
	v.AcademicDegreeGrantedDate = NewDT(input.Index(5).Index(0))
	v.School = NewXON(input.Index(6).Index(0))
	v.SchoolTypeCode = NewCWE(input.Index(7).Index(0))
	v.SchoolAddress = NewXAD(input.Index(8).Index(0))
	v.MajorFieldOfStudy = NewCWESlice(input.Index(9))
	return v
}

// NewEDUSlice will construct a slice of type EDU
func NewEDUSlice(input hl7.HL7Type) []EDU {
	values := make([]EDU, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEDU(input.Index(i))
	}
	return values
}

type CTD struct {
	ContactRole                     []CWE
	ContactName                     []XPN
	ContactAddress                  []XAD
	ContactLocation                 PL
	ContactCommunicationInformation []XTN
	PreferredMethodOfContact        CWE
	ContactIdentifiers              []PLN
}

// NewCTD creates an implementation of CTD
func NewCTD(input hl7.HL7Type) CTD {
	v := CTD{}
	v.ContactRole = NewCWESlice(input.Index(1))
	v.ContactName = NewXPNSlice(input.Index(2))
	v.ContactAddress = NewXADSlice(input.Index(3))
	v.ContactLocation = NewPL(input.Index(4).Index(0))
	v.ContactCommunicationInformation = NewXTNSlice(input.Index(5))
	v.PreferredMethodOfContact = NewCWE(input.Index(6).Index(0))
	v.ContactIdentifiers = NewPLNSlice(input.Index(7))
	return v
}

// NewCTDSlice will construct a slice of type CTD
func NewCTDSlice(input hl7.HL7Type) []CTD {
	values := make([]CTD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCTD(input.Index(i))
	}
	return values
}

type Hxx struct {
	Hxx1 ST
}

// NewHxx creates an implementation of Hxx
func NewHxx(input hl7.HL7Type) Hxx {
	v := Hxx{}
	v.Hxx1 = NewST(input.Index(1).Index(0))
	return v
}

// NewHxxSlice will construct a slice of type Hxx
func NewHxxSlice(input hl7.HL7Type) []Hxx {
	values := make([]Hxx, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewHxx(input.Index(i))
	}
	return values
}

type FT1 struct {
	SetID                                      SI
	TransactionID                              ST
	TransactionBatchID                         ST
	TransactionDate                            DR
	TransactionPostingDate                     DTM
	TransactionType                            CWE
	TransactionCode                            CWE
	TransactionDescription                     ST
	TransactionDescriptionAlt                  ST
	TransactionQuantity                        NM
	TransactionAmountExtended                  CP
	TransactionAmountUnit                      CP
	DepartmentCode                             CWE
	HealthPlanID                               CWE
	InsuranceAmount                            CP
	AssignedPatientLocation                    PL
	FeeSchedule                                CWE
	PatientType                                CWE
	DiagnosisCodeFt1                           []CWE
	PerformedByCode                            []XCN
	OrderedByCode                              []XCN
	UnitCost                                   CP
	FillerOrderNumber                          EI
	EnteredByCode                              []XCN
	ProcedureCode                              CNE
	ProcedureCodeModifier                      []CNE
	AdvancedBeneficiaryNoticeCode              CWE
	MedicallyNecessaryDuplicateProcedureReason CWE
	NdcCode                                    CWE
	PaymentReferenceID                         CX
	TransactionReferenceKey                    []SI
	PerformingFacility                         []XON
	OrderingFacility                           XON
	ItemNumber                                 CWE
	ModelNumber                                ST
	SpecialProcessingCode                      []CWE
	ClinicCode                                 CWE
	ReferralNumber                             CX
	AuthorizationNumber                        CX
	ServiceProviderTaxonomyCode                CWE
	RevenueCode                                CWE
	PrescriptionNumber                         ST
	NdcQtyAndUom                               CQ
}

// NewFT1 creates an implementation of FT1
func NewFT1(input hl7.HL7Type) FT1 {
	v := FT1{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.TransactionID = NewST(input.Index(2).Index(0))
	v.TransactionBatchID = NewST(input.Index(3).Index(0))
	v.TransactionDate = NewDR(input.Index(4).Index(0))
	v.TransactionPostingDate = NewDTM(input.Index(5).Index(0))
	v.TransactionType = NewCWE(input.Index(6).Index(0))
	v.TransactionCode = NewCWE(input.Index(7).Index(0))
	v.TransactionDescription = NewST(input.Index(8).Index(0))
	v.TransactionDescriptionAlt = NewST(input.Index(9).Index(0))
	v.TransactionQuantity = NewNM(input.Index(10).Index(0))
	v.TransactionAmountExtended = NewCP(input.Index(11).Index(0))
	v.TransactionAmountUnit = NewCP(input.Index(12).Index(0))
	v.DepartmentCode = NewCWE(input.Index(13).Index(0))
	v.HealthPlanID = NewCWE(input.Index(14).Index(0))
	v.InsuranceAmount = NewCP(input.Index(15).Index(0))
	v.AssignedPatientLocation = NewPL(input.Index(16).Index(0))
	v.FeeSchedule = NewCWE(input.Index(17).Index(0))
	v.PatientType = NewCWE(input.Index(18).Index(0))
	v.DiagnosisCodeFt1 = NewCWESlice(input.Index(19))
	v.PerformedByCode = NewXCNSlice(input.Index(20))
	v.OrderedByCode = NewXCNSlice(input.Index(21))
	v.UnitCost = NewCP(input.Index(22).Index(0))
	v.FillerOrderNumber = NewEI(input.Index(23).Index(0))
	v.EnteredByCode = NewXCNSlice(input.Index(24))
	v.ProcedureCode = NewCNE(input.Index(25).Index(0))
	v.ProcedureCodeModifier = NewCNESlice(input.Index(26))
	v.AdvancedBeneficiaryNoticeCode = NewCWE(input.Index(27).Index(0))
	v.MedicallyNecessaryDuplicateProcedureReason = NewCWE(input.Index(28).Index(0))
	v.NdcCode = NewCWE(input.Index(29).Index(0))
	v.PaymentReferenceID = NewCX(input.Index(30).Index(0))
	v.TransactionReferenceKey = NewSISlice(input.Index(31))
	v.PerformingFacility = NewXONSlice(input.Index(32))
	v.OrderingFacility = NewXON(input.Index(33).Index(0))
	v.ItemNumber = NewCWE(input.Index(34).Index(0))
	v.ModelNumber = NewST(input.Index(35).Index(0))
	v.SpecialProcessingCode = NewCWESlice(input.Index(36))
	v.ClinicCode = NewCWE(input.Index(37).Index(0))
	v.ReferralNumber = NewCX(input.Index(38).Index(0))
	v.AuthorizationNumber = NewCX(input.Index(39).Index(0))
	v.ServiceProviderTaxonomyCode = NewCWE(input.Index(40).Index(0))
	v.RevenueCode = NewCWE(input.Index(41).Index(0))
	v.PrescriptionNumber = NewST(input.Index(42).Index(0))
	v.NdcQtyAndUom = NewCQ(input.Index(43).Index(0))
	return v
}

// NewFT1Slice will construct a slice of type FT1
func NewFT1Slice(input hl7.HL7Type) []FT1 {
	values := make([]FT1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewFT1(input.Index(i))
	}
	return values
}

type EQP struct {
	EventType       CWE
	FileName        ST
	StartDateTime   DTM
	EndDateTime     DTM
	TransactionData FT
}

// NewEQP creates an implementation of EQP
func NewEQP(input hl7.HL7Type) EQP {
	v := EQP{}
	v.EventType = NewCWE(input.Index(1).Index(0))
	v.FileName = NewST(input.Index(2).Index(0))
	v.StartDateTime = NewDTM(input.Index(3).Index(0))
	v.EndDateTime = NewDTM(input.Index(4).Index(0))
	v.TransactionData = NewFT(input.Index(5).Index(0))
	return v
}

// NewEQPSlice will construct a slice of type EQP
func NewEQPSlice(input hl7.HL7Type) []EQP {
	values := make([]EQP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEQP(input.Index(i))
	}
	return values
}

type NTE struct {
	SetID              SI
	SourceOfComment    ID
	Comment            []FT
	CommentType        CWE
	EnteredBy          XCN
	EnteredDateTime    DTM
	EffectiveStartDate DTM
	ExpirationDate     DTM
}

// NewNTE creates an implementation of NTE
func NewNTE(input hl7.HL7Type) NTE {
	v := NTE{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.SourceOfComment = NewID(input.Index(2).Index(0))
	v.Comment = NewFTSlice(input.Index(3))
	v.CommentType = NewCWE(input.Index(4).Index(0))
	v.EnteredBy = NewXCN(input.Index(5).Index(0))
	v.EnteredDateTime = NewDTM(input.Index(6).Index(0))
	v.EffectiveStartDate = NewDTM(input.Index(7).Index(0))
	v.ExpirationDate = NewDTM(input.Index(8).Index(0))
	return v
}

// NewNTESlice will construct a slice of type NTE
func NewNTESlice(input hl7.HL7Type) []NTE {
	values := make([]NTE, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewNTE(input.Index(i))
	}
	return values
}

type BPO struct {
	SetID                          SI
	BpUniversalServiceIdentifier   CWE
	BpProcessingRequirements       []CWE
	BpQuantity                     NM
	BpAmount                       NM
	BpUnits                        CWE
	BpIntendedUseDateTime          DTM
	BpIntendedDispenseFromLocation PL
	BpIntendedDispenseFromAddress  XAD
	BpRequestedDispenseDateTime    DTM
	BpRequestedDispenseToLocation  PL
	BpRequestedDispenseToAddress   XAD
	BpIndicationForUse             []CWE
	BpInformedConsentIndicator     ID
}

// NewBPO creates an implementation of BPO
func NewBPO(input hl7.HL7Type) BPO {
	v := BPO{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.BpUniversalServiceIdentifier = NewCWE(input.Index(2).Index(0))
	v.BpProcessingRequirements = NewCWESlice(input.Index(3))
	v.BpQuantity = NewNM(input.Index(4).Index(0))
	v.BpAmount = NewNM(input.Index(5).Index(0))
	v.BpUnits = NewCWE(input.Index(6).Index(0))
	v.BpIntendedUseDateTime = NewDTM(input.Index(7).Index(0))
	v.BpIntendedDispenseFromLocation = NewPL(input.Index(8).Index(0))
	v.BpIntendedDispenseFromAddress = NewXAD(input.Index(9).Index(0))
	v.BpRequestedDispenseDateTime = NewDTM(input.Index(10).Index(0))
	v.BpRequestedDispenseToLocation = NewPL(input.Index(11).Index(0))
	v.BpRequestedDispenseToAddress = NewXAD(input.Index(12).Index(0))
	v.BpIndicationForUse = NewCWESlice(input.Index(13))
	v.BpInformedConsentIndicator = NewID(input.Index(14).Index(0))
	return v
}

// NewBPOSlice will construct a slice of type BPO
func NewBPOSlice(input hl7.HL7Type) []BPO {
	values := make([]BPO, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBPO(input.Index(i))
	}
	return values
}

type CM0 struct {
	SetID                   SI
	SponsorStudyID          EI
	AlternateStudyID        []EI
	TitleOfStudy            ST
	ChairmanOfStudy         []XCN
	LastIrbApprovalDate     DT
	TotalAccrualToDate      NM
	LastAccrualDate         DT
	ContactForStudy         []XCN
	ContactsTelephoneNumber XTN
	ContactsAddress         []XAD
}

// NewCM0 creates an implementation of CM0
func NewCM0(input hl7.HL7Type) CM0 {
	v := CM0{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.SponsorStudyID = NewEI(input.Index(2).Index(0))
	v.AlternateStudyID = NewEISlice(input.Index(3))
	v.TitleOfStudy = NewST(input.Index(4).Index(0))
	v.ChairmanOfStudy = NewXCNSlice(input.Index(5))
	v.LastIrbApprovalDate = NewDT(input.Index(6).Index(0))
	v.TotalAccrualToDate = NewNM(input.Index(7).Index(0))
	v.LastAccrualDate = NewDT(input.Index(8).Index(0))
	v.ContactForStudy = NewXCNSlice(input.Index(9))
	v.ContactsTelephoneNumber = NewXTN(input.Index(10).Index(0))
	v.ContactsAddress = NewXADSlice(input.Index(11))
	return v
}

// NewCM0Slice will construct a slice of type CM0
func NewCM0Slice(input hl7.HL7Type) []CM0 {
	values := make([]CM0, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCM0(input.Index(i))
	}
	return values
}

type DMI struct {
	DiagnosticRelatedGroup  CNE
	MajorDiagnosticCategory CNE
	LowerAndUpperTrimPoints NR
	AverageLengthOfStay     NM
	RelativeWeight          NM
}

// NewDMI creates an implementation of DMI
func NewDMI(input hl7.HL7Type) DMI {
	v := DMI{}
	v.DiagnosticRelatedGroup = NewCNE(input.Index(1).Index(0))
	v.MajorDiagnosticCategory = NewCNE(input.Index(2).Index(0))
	v.LowerAndUpperTrimPoints = NewNR(input.Index(3).Index(0))
	v.AverageLengthOfStay = NewNM(input.Index(4).Index(0))
	v.RelativeWeight = NewNM(input.Index(5).Index(0))
	return v
}

// NewDMISlice will construct a slice of type DMI
func NewDMISlice(input hl7.HL7Type) []DMI {
	values := make([]DMI, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDMI(input.Index(i))
	}
	return values
}

type MSA struct {
	AcknowledgmentCode        ID
	MessageControlID          ST
	TextMessage               ST
	ExpectedSequenceNumber    NM
	DelayedAcknowledgmentType ST
	ErrorCondition            ST
	MessageWaitingNumber      NM
	MessageWaitingPriority    ID
}

// NewMSA creates an implementation of MSA
func NewMSA(input hl7.HL7Type) MSA {
	v := MSA{}
	v.AcknowledgmentCode = NewID(input.Index(1).Index(0))
	v.MessageControlID = NewST(input.Index(2).Index(0))
	v.TextMessage = NewST(input.Index(3).Index(0))
	v.ExpectedSequenceNumber = NewNM(input.Index(4).Index(0))
	v.DelayedAcknowledgmentType = NewST(input.Index(5).Index(0))
	v.ErrorCondition = NewST(input.Index(6).Index(0))
	v.MessageWaitingNumber = NewNM(input.Index(7).Index(0))
	v.MessageWaitingPriority = NewID(input.Index(8).Index(0))
	return v
}

// NewMSASlice will construct a slice of type MSA
func NewMSASlice(input hl7.HL7Type) []MSA {
	values := make([]MSA, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMSA(input.Index(i))
	}
	return values
}

type BHS struct {
	BatchFieldSeparator          ST
	BatchEncodingCharacters      ST
	BatchSendingApplication      HD
	BatchSendingFacility         HD
	BatchReceivingApplication    HD
	BatchReceivingFacility       HD
	BatchCreationDateTime        DTM
	BatchSecurity                ST
	BatchNameIDType              ST
	BatchComment                 ST
	BatchControlID               ST
	ReferenceBatchControlID      ST
	BatchSendingNetworkAddress   HD
	BatchReceivingNetworkAddress HD
}

// NewBHS creates an implementation of BHS
func NewBHS(input hl7.HL7Type) BHS {
	v := BHS{}
	v.BatchFieldSeparator = NewST(input.Index(1).Index(0))
	v.BatchEncodingCharacters = NewST(input.Index(2).Index(0))
	v.BatchSendingApplication = NewHD(input.Index(3).Index(0))
	v.BatchSendingFacility = NewHD(input.Index(4).Index(0))
	v.BatchReceivingApplication = NewHD(input.Index(5).Index(0))
	v.BatchReceivingFacility = NewHD(input.Index(6).Index(0))
	v.BatchCreationDateTime = NewDTM(input.Index(7).Index(0))
	v.BatchSecurity = NewST(input.Index(8).Index(0))
	v.BatchNameIDType = NewST(input.Index(9).Index(0))
	v.BatchComment = NewST(input.Index(10).Index(0))
	v.BatchControlID = NewST(input.Index(11).Index(0))
	v.ReferenceBatchControlID = NewST(input.Index(12).Index(0))
	v.BatchSendingNetworkAddress = NewHD(input.Index(13).Index(0))
	v.BatchReceivingNetworkAddress = NewHD(input.Index(14).Index(0))
	return v
}

// NewBHSSlice will construct a slice of type BHS
func NewBHSSlice(input hl7.HL7Type) []BHS {
	values := make([]BHS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBHS(input.Index(i))
	}
	return values
}

type UB2 struct {
	SetID                     SI
	CoInsuranceDays9          ST
	ConditionCode2430         []CWE
	CoveredDays7              ST
	NonCoveredDays8           ST
	ValueAmountCode           []UVC
	OccurrenceCodeDate3235    []OCD
	OccurrenceSpanCodeDates36 []OSP
	UB92Locator2state         []ST
	UB92Locator11state        []ST
	UB92Locator31national     ST
	DocumentControlNumber     []ST
	UB92Locator49national     []ST
	UB92Locator56state        []ST
	UB92Locator57national     ST
	UB92Locator78state        []ST
	SpecialVisitCount         NM
}

// NewUB2 creates an implementation of UB2
func NewUB2(input hl7.HL7Type) UB2 {
	v := UB2{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.CoInsuranceDays9 = NewST(input.Index(2).Index(0))
	v.ConditionCode2430 = NewCWESlice(input.Index(3))
	v.CoveredDays7 = NewST(input.Index(4).Index(0))
	v.NonCoveredDays8 = NewST(input.Index(5).Index(0))
	v.ValueAmountCode = NewUVCSlice(input.Index(6))
	v.OccurrenceCodeDate3235 = NewOCDSlice(input.Index(7))
	v.OccurrenceSpanCodeDates36 = NewOSPSlice(input.Index(8))
	v.UB92Locator2state = NewSTSlice(input.Index(9))
	v.UB92Locator11state = NewSTSlice(input.Index(10))
	v.UB92Locator31national = NewST(input.Index(11).Index(0))
	v.DocumentControlNumber = NewSTSlice(input.Index(12))
	v.UB92Locator49national = NewSTSlice(input.Index(13))
	v.UB92Locator56state = NewSTSlice(input.Index(14))
	v.UB92Locator57national = NewST(input.Index(15).Index(0))
	v.UB92Locator78state = NewSTSlice(input.Index(16))
	v.SpecialVisitCount = NewNM(input.Index(17).Index(0))
	return v
}

// NewUB2Slice will construct a slice of type UB2
func NewUB2Slice(input hl7.HL7Type) []UB2 {
	values := make([]UB2, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewUB2(input.Index(i))
	}
	return values
}

type ISD struct {
	ReferenceInteractionNumber NM
	InteractionTypeIdentifier  CWE
	InteractionActiveState     CWE
}

// NewISD creates an implementation of ISD
func NewISD(input hl7.HL7Type) ISD {
	v := ISD{}
	v.ReferenceInteractionNumber = NewNM(input.Index(1).Index(0))
	v.InteractionTypeIdentifier = NewCWE(input.Index(2).Index(0))
	v.InteractionActiveState = NewCWE(input.Index(3).Index(0))
	return v
}

// NewISDSlice will construct a slice of type ISD
func NewISDSlice(input hl7.HL7Type) []ISD {
	values := make([]ISD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewISD(input.Index(i))
	}
	return values
}

type SPM struct {
	SetID                      SI
	SpecimenID                 EIP
	SpecimenParentIDs          []EIP
	SpecimenType               CWE
	SpecimenTypeModifier       []CWE
	SpecimenAdditives          []CWE
	SpecimenCollectionMethod   CWE
	SpecimenSourceSite         CWE
	SpecimenSourceSiteModifier []CWE
	SpecimenCollectionSite     CWE
	SpecimenRole               []CWE
	SpecimenCollectionAmount   CQ
	GroupedSpecimenCount       NM
	SpecimenDescription        []ST
	SpecimenHandlingCode       []CWE
	SpecimenRiskCode           []CWE
	SpecimenCollectionDateTime DR
	SpecimenReceivedDateTime   DTM
	SpecimenExpirationDateTime DTM
	SpecimenAvailability       ID
	SpecimenRejectReason       []CWE
	SpecimenQuality            CWE
	SpecimenAppropriateness    CWE
	SpecimenCondition          []CWE
	SpecimenCurrentQuantity    CQ
	NumberOfSpecimenContainers NM
	ContainerType              CWE
	ContainerCondition         CWE
	SpecimenChildRole          CWE
	AccessionID                []CX
	OtherSpecimenID            []CX
	ShipmentID                 EI
}

// NewSPM creates an implementation of SPM
func NewSPM(input hl7.HL7Type) SPM {
	v := SPM{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.SpecimenID = NewEIP(input.Index(2).Index(0))
	v.SpecimenParentIDs = NewEIPSlice(input.Index(3))
	v.SpecimenType = NewCWE(input.Index(4).Index(0))
	v.SpecimenTypeModifier = NewCWESlice(input.Index(5))
	v.SpecimenAdditives = NewCWESlice(input.Index(6))
	v.SpecimenCollectionMethod = NewCWE(input.Index(7).Index(0))
	v.SpecimenSourceSite = NewCWE(input.Index(8).Index(0))
	v.SpecimenSourceSiteModifier = NewCWESlice(input.Index(9))
	v.SpecimenCollectionSite = NewCWE(input.Index(10).Index(0))
	v.SpecimenRole = NewCWESlice(input.Index(11))
	v.SpecimenCollectionAmount = NewCQ(input.Index(12).Index(0))
	v.GroupedSpecimenCount = NewNM(input.Index(13).Index(0))
	v.SpecimenDescription = NewSTSlice(input.Index(14))
	v.SpecimenHandlingCode = NewCWESlice(input.Index(15))
	v.SpecimenRiskCode = NewCWESlice(input.Index(16))
	v.SpecimenCollectionDateTime = NewDR(input.Index(17).Index(0))
	v.SpecimenReceivedDateTime = NewDTM(input.Index(18).Index(0))
	v.SpecimenExpirationDateTime = NewDTM(input.Index(19).Index(0))
	v.SpecimenAvailability = NewID(input.Index(20).Index(0))
	v.SpecimenRejectReason = NewCWESlice(input.Index(21))
	v.SpecimenQuality = NewCWE(input.Index(22).Index(0))
	v.SpecimenAppropriateness = NewCWE(input.Index(23).Index(0))
	v.SpecimenCondition = NewCWESlice(input.Index(24))
	v.SpecimenCurrentQuantity = NewCQ(input.Index(25).Index(0))
	v.NumberOfSpecimenContainers = NewNM(input.Index(26).Index(0))
	v.ContainerType = NewCWE(input.Index(27).Index(0))
	v.ContainerCondition = NewCWE(input.Index(28).Index(0))
	v.SpecimenChildRole = NewCWE(input.Index(29).Index(0))
	v.AccessionID = NewCXSlice(input.Index(30))
	v.OtherSpecimenID = NewCXSlice(input.Index(31))
	v.ShipmentID = NewEI(input.Index(32).Index(0))
	return v
}

// NewSPMSlice will construct a slice of type SPM
func NewSPMSlice(input hl7.HL7Type) []SPM {
	values := make([]SPM, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSPM(input.Index(i))
	}
	return values
}

type BTX struct {
	SetID                              SI
	BcDonationID                       EI
	BcComponent                        CNE
	BcBloodGroup                       CNE
	CpCommercialProduct                CWE
	CpManufacturer                     XON
	CpLotNumber                        EI
	BpQuantity                         NM
	BpAmount                           NM
	BpUnits                            CWE
	BpTransfusionDispositionStatus     CWE
	BpMessageStatus                    ID
	BpDateTimeOfStatus                 DTM
	BpTransfusionAdministrator         XCN
	BpTransfusionVerifier              XCN
	BpTransfusionStartDateTimeOfStatus DTM
	BpTransfusionEndDateTimeOfStatus   DTM
	BpAdverseReactionType              []CWE
	BpTransfusionInterruptedReason     CWE
	BPUniqueID                         EI
}

// NewBTX creates an implementation of BTX
func NewBTX(input hl7.HL7Type) BTX {
	v := BTX{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.BcDonationID = NewEI(input.Index(2).Index(0))
	v.BcComponent = NewCNE(input.Index(3).Index(0))
	v.BcBloodGroup = NewCNE(input.Index(4).Index(0))
	v.CpCommercialProduct = NewCWE(input.Index(5).Index(0))
	v.CpManufacturer = NewXON(input.Index(6).Index(0))
	v.CpLotNumber = NewEI(input.Index(7).Index(0))
	v.BpQuantity = NewNM(input.Index(8).Index(0))
	v.BpAmount = NewNM(input.Index(9).Index(0))
	v.BpUnits = NewCWE(input.Index(10).Index(0))
	v.BpTransfusionDispositionStatus = NewCWE(input.Index(11).Index(0))
	v.BpMessageStatus = NewID(input.Index(12).Index(0))
	v.BpDateTimeOfStatus = NewDTM(input.Index(13).Index(0))
	v.BpTransfusionAdministrator = NewXCN(input.Index(14).Index(0))
	v.BpTransfusionVerifier = NewXCN(input.Index(15).Index(0))
	v.BpTransfusionStartDateTimeOfStatus = NewDTM(input.Index(16).Index(0))
	v.BpTransfusionEndDateTimeOfStatus = NewDTM(input.Index(17).Index(0))
	v.BpAdverseReactionType = NewCWESlice(input.Index(18))
	v.BpTransfusionInterruptedReason = NewCWE(input.Index(19).Index(0))
	v.BPUniqueID = NewEI(input.Index(20).Index(0))
	return v
}

// NewBTXSlice will construct a slice of type BTX
func NewBTXSlice(input hl7.HL7Type) []BTX {
	values := make([]BTX, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBTX(input.Index(i))
	}
	return values
}

type ARQ struct {
	PlacerAppointmentID         EI
	FillerAppointmentID         EI
	OccurrenceNumber            NM
	PlacerGroupNumber           EI
	ScheduleID                  CWE
	RequestEventReason          CWE
	AppointmentReason           CWE
	AppointmentType             CWE
	AppointmentDuration         NM
	AppointmentDurationUnits    CNE
	RequestedStartDateTimeRange []DR
	PriorityArq                 ST
	RepeatingInterval           RI
	RepeatingIntervalDuration   ST
	PlacerContactPerson         []XCN
	PlacerContactPhoneNumber    []XTN
	PlacerContactAddress        []XAD
	PlacerContactLocation       PL
	EnteredByPerson             []XCN
	EnteredByPhoneNumber        []XTN
	EnteredByLocation           PL
	ParentPlacerAppointmentID   EI
	ParentFillerAppointmentID   EI
	PlacerOrderNumber           []EI
	FillerOrderNumber           []EI
}

// NewARQ creates an implementation of ARQ
func NewARQ(input hl7.HL7Type) ARQ {
	v := ARQ{}
	v.PlacerAppointmentID = NewEI(input.Index(1).Index(0))
	v.FillerAppointmentID = NewEI(input.Index(2).Index(0))
	v.OccurrenceNumber = NewNM(input.Index(3).Index(0))
	v.PlacerGroupNumber = NewEI(input.Index(4).Index(0))
	v.ScheduleID = NewCWE(input.Index(5).Index(0))
	v.RequestEventReason = NewCWE(input.Index(6).Index(0))
	v.AppointmentReason = NewCWE(input.Index(7).Index(0))
	v.AppointmentType = NewCWE(input.Index(8).Index(0))
	v.AppointmentDuration = NewNM(input.Index(9).Index(0))
	v.AppointmentDurationUnits = NewCNE(input.Index(10).Index(0))
	v.RequestedStartDateTimeRange = NewDRSlice(input.Index(11))
	v.PriorityArq = NewST(input.Index(12).Index(0))
	v.RepeatingInterval = NewRI(input.Index(13).Index(0))
	v.RepeatingIntervalDuration = NewST(input.Index(14).Index(0))
	v.PlacerContactPerson = NewXCNSlice(input.Index(15))
	v.PlacerContactPhoneNumber = NewXTNSlice(input.Index(16))
	v.PlacerContactAddress = NewXADSlice(input.Index(17))
	v.PlacerContactLocation = NewPL(input.Index(18).Index(0))
	v.EnteredByPerson = NewXCNSlice(input.Index(19))
	v.EnteredByPhoneNumber = NewXTNSlice(input.Index(20))
	v.EnteredByLocation = NewPL(input.Index(21).Index(0))
	v.ParentPlacerAppointmentID = NewEI(input.Index(22).Index(0))
	v.ParentFillerAppointmentID = NewEI(input.Index(23).Index(0))
	v.PlacerOrderNumber = NewEISlice(input.Index(24))
	v.FillerOrderNumber = NewEISlice(input.Index(25))
	return v
}

// NewARQSlice will construct a slice of type ARQ
func NewARQSlice(input hl7.HL7Type) []ARQ {
	values := make([]ARQ, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewARQ(input.Index(i))
	}
	return values
}

type PRA struct {
	PrimaryKeyValuePra                        CWE
	PractitionerGroup                         []CWE
	PractitionerCategory                      []CWE
	ProviderBilling                           ID
	Specialty                                 []SPD
	PractitionerIDNumbers                     []PLN
	Privileges                                []PIP
	DateEnteredPractice                       DT
	Institution                               CWE
	DateLeftPractice                          DT
	GovernmentReimbursementBillingEligibility []CWE
	SetID                                     SI
}

// NewPRA creates an implementation of PRA
func NewPRA(input hl7.HL7Type) PRA {
	v := PRA{}
	v.PrimaryKeyValuePra = NewCWE(input.Index(1).Index(0))
	v.PractitionerGroup = NewCWESlice(input.Index(2))
	v.PractitionerCategory = NewCWESlice(input.Index(3))
	v.ProviderBilling = NewID(input.Index(4).Index(0))
	v.Specialty = NewSPDSlice(input.Index(5))
	v.PractitionerIDNumbers = NewPLNSlice(input.Index(6))
	v.Privileges = NewPIPSlice(input.Index(7))
	v.DateEnteredPractice = NewDT(input.Index(8).Index(0))
	v.Institution = NewCWE(input.Index(9).Index(0))
	v.DateLeftPractice = NewDT(input.Index(10).Index(0))
	v.GovernmentReimbursementBillingEligibility = NewCWESlice(input.Index(11))
	v.SetID = NewSI(input.Index(12).Index(0))
	return v
}

// NewPRASlice will construct a slice of type PRA
func NewPRASlice(input hl7.HL7Type) []PRA {
	values := make([]PRA, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPRA(input.Index(i))
	}
	return values
}

type RCP struct {
	QueryPriority            ID
	QuantityLimitedRequest   CQ
	ResponseModality         CNE
	ExecutionAndDeliveryTime DTM
	ModifyIndicator          ID
	SortByField              []SRT
	SegmentGroupInclusion    []ID
}

// NewRCP creates an implementation of RCP
func NewRCP(input hl7.HL7Type) RCP {
	v := RCP{}
	v.QueryPriority = NewID(input.Index(1).Index(0))
	v.QuantityLimitedRequest = NewCQ(input.Index(2).Index(0))
	v.ResponseModality = NewCNE(input.Index(3).Index(0))
	v.ExecutionAndDeliveryTime = NewDTM(input.Index(4).Index(0))
	v.ModifyIndicator = NewID(input.Index(5).Index(0))
	v.SortByField = NewSRTSlice(input.Index(6))
	v.SegmentGroupInclusion = NewIDSlice(input.Index(7))
	return v
}

// NewRCPSlice will construct a slice of type RCP
func NewRCPSlice(input hl7.HL7Type) []RCP {
	values := make([]RCP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRCP(input.Index(i))
	}
	return values
}

type CM2 struct {
	SetID                        SI
	ScheduledTimePoint           CWE
	DescriptionOfTimePoint       ST
	EventsScheduledThisTimePoint []CWE
}

// NewCM2 creates an implementation of CM2
func NewCM2(input hl7.HL7Type) CM2 {
	v := CM2{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.ScheduledTimePoint = NewCWE(input.Index(2).Index(0))
	v.DescriptionOfTimePoint = NewST(input.Index(3).Index(0))
	v.EventsScheduledThisTimePoint = NewCWESlice(input.Index(4))
	return v
}

// NewCM2Slice will construct a slice of type CM2
func NewCM2Slice(input hl7.HL7Type) []CM2 {
	values := make([]CM2, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCM2(input.Index(i))
	}
	return values
}

type PSG struct {
	ProviderProductServiceGroupNumber EI
	PayerProductServiceGroupNumber    EI
	ProductServiceGroupSequenceNumber SI
	AdjudicateAsGroup                 ID
	ProductServiceGroupBilledAmount   CP
	ProductServiceGroupDescription    ST
}

// NewPSG creates an implementation of PSG
func NewPSG(input hl7.HL7Type) PSG {
	v := PSG{}
	v.ProviderProductServiceGroupNumber = NewEI(input.Index(1).Index(0))
	v.PayerProductServiceGroupNumber = NewEI(input.Index(2).Index(0))
	v.ProductServiceGroupSequenceNumber = NewSI(input.Index(3).Index(0))
	v.AdjudicateAsGroup = NewID(input.Index(4).Index(0))
	v.ProductServiceGroupBilledAmount = NewCP(input.Index(5).Index(0))
	v.ProductServiceGroupDescription = NewST(input.Index(6).Index(0))
	return v
}

// NewPSGSlice will construct a slice of type PSG
func NewPSGSlice(input hl7.HL7Type) []PSG {
	values := make([]PSG, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPSG(input.Index(i))
	}
	return values
}

type DSP struct {
	SetID             SI
	DisplayLevel      SI
	DataLine          TX
	LogicalBreakPoint ST
	ResultID          TX
}

// NewDSP creates an implementation of DSP
func NewDSP(input hl7.HL7Type) DSP {
	v := DSP{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.DisplayLevel = NewSI(input.Index(2).Index(0))
	v.DataLine = NewTX(input.Index(3).Index(0))
	v.LogicalBreakPoint = NewST(input.Index(4).Index(0))
	v.ResultID = NewTX(input.Index(5).Index(0))
	return v
}

// NewDSPSlice will construct a slice of type DSP
func NewDSPSlice(input hl7.HL7Type) []DSP {
	values := make([]DSP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDSP(input.Index(i))
	}
	return values
}

type QRD struct {
	QRD1 ST
}

// NewQRD creates an implementation of QRD
func NewQRD(input hl7.HL7Type) QRD {
	v := QRD{}
	v.QRD1 = NewST(input.Index(1).Index(0))
	return v
}

// NewQRDSlice will construct a slice of type QRD
func NewQRDSlice(input hl7.HL7Type) []QRD {
	values := make([]QRD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQRD(input.Index(i))
	}
	return values
}

type RFI struct {
	RequestDate                           DTM
	ResponseDueDate                       DTM
	PatientConsent                        ID
	DateAdditionalInformationWasSubmitted DTM
}

// NewRFI creates an implementation of RFI
func NewRFI(input hl7.HL7Type) RFI {
	v := RFI{}
	v.RequestDate = NewDTM(input.Index(1).Index(0))
	v.ResponseDueDate = NewDTM(input.Index(2).Index(0))
	v.PatientConsent = NewID(input.Index(3).Index(0))
	v.DateAdditionalInformationWasSubmitted = NewDTM(input.Index(4).Index(0))
	return v
}

// NewRFISlice will construct a slice of type RFI
func NewRFISlice(input hl7.HL7Type) []RFI {
	values := make([]RFI, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRFI(input.Index(i))
	}
	return values
}

type PTH struct {
	ActionCode                           ID
	PathwayID                            CWE
	PathwayInstanceID                    EI
	PathwayEstablishedDateTime           DTM
	PathwayLifeCycleStatus               CWE
	ChangePathwayLifeCycleStatusDateTime DTM
	MoodCode                             CNE
}

// NewPTH creates an implementation of PTH
func NewPTH(input hl7.HL7Type) PTH {
	v := PTH{}
	v.ActionCode = NewID(input.Index(1).Index(0))
	v.PathwayID = NewCWE(input.Index(2).Index(0))
	v.PathwayInstanceID = NewEI(input.Index(3).Index(0))
	v.PathwayEstablishedDateTime = NewDTM(input.Index(4).Index(0))
	v.PathwayLifeCycleStatus = NewCWE(input.Index(5).Index(0))
	v.ChangePathwayLifeCycleStatusDateTime = NewDTM(input.Index(6).Index(0))
	v.MoodCode = NewCNE(input.Index(7).Index(0))
	return v
}

// NewPTHSlice will construct a slice of type PTH
func NewPTHSlice(input hl7.HL7Type) []PTH {
	values := make([]PTH, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPTH(input.Index(i))
	}
	return values
}

type ORC struct {
	OrderControl                            ID
	PlacerOrderNumber                       EI
	FillerOrderNumber                       EI
	PlacerGroupNumber                       EI
	OrderStatus                             ID
	ResponseFlag                            ID
	QuantityTiming                          []ST
	Parent                                  EIP
	DateTimeOfTransaction                   DTM
	EnteredBy                               []XCN
	VerifiedBy                              []XCN
	OrderingProvider                        []XCN
	EnterersLocation                        PL
	CallBackPhoneNumber                     []XTN
	OrderEffectiveDateTime                  DTM
	OrderControlCodeReason                  CWE
	EnteringOrganization                    CWE
	EnteringDevice                          CWE
	ActionBy                                []XCN
	AdvancedBeneficiaryNoticeCode           CWE
	OrderingFacilityName                    []XON
	OrderingFacilityAddress                 []XAD
	OrderingFacilityPhoneNumber             []XTN
	OrderingProviderAddress                 []XAD
	OrderStatusModifier                     CWE
	AdvancedBeneficiaryNoticeOverrideReason CWE
	FillersExpectedAvailabilityDateTime     DTM
	ConfidentialityCode                     CWE
	OrderType                               CWE
	EntererAuthorizationMode                CNE
	ParentUniversalServiceIdentifier        CWE
	AdvancedBeneficiaryNoticeDate           DT
	AlternatePlacerOrderNumber              []CX
	OrderWorkflowProfile                    []EI
}

// NewORC creates an implementation of ORC
func NewORC(input hl7.HL7Type) ORC {
	v := ORC{}
	v.OrderControl = NewID(input.Index(1).Index(0))
	v.PlacerOrderNumber = NewEI(input.Index(2).Index(0))
	v.FillerOrderNumber = NewEI(input.Index(3).Index(0))
	v.PlacerGroupNumber = NewEI(input.Index(4).Index(0))
	v.OrderStatus = NewID(input.Index(5).Index(0))
	v.ResponseFlag = NewID(input.Index(6).Index(0))
	v.QuantityTiming = NewSTSlice(input.Index(7))
	v.Parent = NewEIP(input.Index(8).Index(0))
	v.DateTimeOfTransaction = NewDTM(input.Index(9).Index(0))
	v.EnteredBy = NewXCNSlice(input.Index(10))
	v.VerifiedBy = NewXCNSlice(input.Index(11))
	v.OrderingProvider = NewXCNSlice(input.Index(12))
	v.EnterersLocation = NewPL(input.Index(13).Index(0))
	v.CallBackPhoneNumber = NewXTNSlice(input.Index(14))
	v.OrderEffectiveDateTime = NewDTM(input.Index(15).Index(0))
	v.OrderControlCodeReason = NewCWE(input.Index(16).Index(0))
	v.EnteringOrganization = NewCWE(input.Index(17).Index(0))
	v.EnteringDevice = NewCWE(input.Index(18).Index(0))
	v.ActionBy = NewXCNSlice(input.Index(19))
	v.AdvancedBeneficiaryNoticeCode = NewCWE(input.Index(20).Index(0))
	v.OrderingFacilityName = NewXONSlice(input.Index(21))
	v.OrderingFacilityAddress = NewXADSlice(input.Index(22))
	v.OrderingFacilityPhoneNumber = NewXTNSlice(input.Index(23))
	v.OrderingProviderAddress = NewXADSlice(input.Index(24))
	v.OrderStatusModifier = NewCWE(input.Index(25).Index(0))
	v.AdvancedBeneficiaryNoticeOverrideReason = NewCWE(input.Index(26).Index(0))
	v.FillersExpectedAvailabilityDateTime = NewDTM(input.Index(27).Index(0))
	v.ConfidentialityCode = NewCWE(input.Index(28).Index(0))
	v.OrderType = NewCWE(input.Index(29).Index(0))
	v.EntererAuthorizationMode = NewCNE(input.Index(30).Index(0))
	v.ParentUniversalServiceIdentifier = NewCWE(input.Index(31).Index(0))
	v.AdvancedBeneficiaryNoticeDate = NewDT(input.Index(32).Index(0))
	v.AlternatePlacerOrderNumber = NewCXSlice(input.Index(33))
	v.OrderWorkflowProfile = NewEISlice(input.Index(34))
	return v
}

// NewORCSlice will construct a slice of type ORC
func NewORCSlice(input hl7.HL7Type) []ORC {
	values := make([]ORC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORC(input.Index(i))
	}
	return values
}

type QID struct {
	QueryTag         ST
	MessageQueryName CWE
}

// NewQID creates an implementation of QID
func NewQID(input hl7.HL7Type) QID {
	v := QID{}
	v.QueryTag = NewST(input.Index(1).Index(0))
	v.MessageQueryName = NewCWE(input.Index(2).Index(0))
	return v
}

// NewQIDSlice will construct a slice of type QID
func NewQIDSlice(input hl7.HL7Type) []QID {
	values := make([]QID, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQID(input.Index(i))
	}
	return values
}

type OBR struct {
	SetID                                      SI
	PlacerOrderNumber                          EI
	FillerOrderNumber                          EI
	UniversalServiceIdentifier                 CWE
	Priority                                   ST
	RequestedDateTime                          ST
	ObservationDateTime                        DTM
	ObservationEndDateTime                     DTM
	CollectionVolume                           CQ
	CollectorIdentifier                        []XCN
	SpecimenActionCode                         ID
	DangerCode                                 CWE
	RelevantClinicalInformation                ST
	SpecimenReceivedDateTime                   ST
	SpecimenSource                             ST
	OrderingProvider                           []XCN
	OrderCallbackPhoneNumber                   []XTN
	PlacerField1                               ST
	PlacerField2                               ST
	FillerField1                               ST
	FillerField2                               ST
	ResultsRptStatusChngDateTime               DTM
	ChargeToPractice                           MOC
	DiagnosticServSectID                       ID
	ResultStatus                               ID
	ParentResult                               PRL
	QuantityTiming                             []ST
	ResultCopiesTo                             []XCN
	Parent                                     EIP
	TransportationMode                         ID
	ReasonForStudy                             []CWE
	PrincipalResultInterpreter                 NDL
	AssistantResultInterpreter                 []NDL
	Technician                                 []NDL
	Transcriptionist                           []NDL
	ScheduledDateTime                          DTM
	NumberOfSampleContainers                   NM
	TransportLogisticsOfCollectedSample        []CWE
	CollectorsComment                          []CWE
	TransportArrangementResponsibility         CWE
	TransportArranged                          ID
	EscortRequired                             ID
	PlannedPatientTransportComment             []CWE
	ProcedureCode                              CNE
	ProcedureCodeModifier                      []CNE
	PlacerSupplementalServiceInformation       []CWE
	FillerSupplementalServiceInformation       []CWE
	MedicallyNecessaryDuplicateProcedureReason CWE
	ResultHandling                             CWE
	ParentUniversalServiceIdentifier           CWE
	ObservationGroupID                         EI
	ParentObservationGroupID                   EI
	AlternatePlacerOrderNumber                 []CX
	ParentOrder                                EIP
}

// NewOBR creates an implementation of OBR
func NewOBR(input hl7.HL7Type) OBR {
	v := OBR{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.PlacerOrderNumber = NewEI(input.Index(2).Index(0))
	v.FillerOrderNumber = NewEI(input.Index(3).Index(0))
	v.UniversalServiceIdentifier = NewCWE(input.Index(4).Index(0))
	v.Priority = NewST(input.Index(5).Index(0))
	v.RequestedDateTime = NewST(input.Index(6).Index(0))
	v.ObservationDateTime = NewDTM(input.Index(7).Index(0))
	v.ObservationEndDateTime = NewDTM(input.Index(8).Index(0))
	v.CollectionVolume = NewCQ(input.Index(9).Index(0))
	v.CollectorIdentifier = NewXCNSlice(input.Index(10))
	v.SpecimenActionCode = NewID(input.Index(11).Index(0))
	v.DangerCode = NewCWE(input.Index(12).Index(0))
	v.RelevantClinicalInformation = NewST(input.Index(13).Index(0))
	v.SpecimenReceivedDateTime = NewST(input.Index(14).Index(0))
	v.SpecimenSource = NewST(input.Index(15).Index(0))
	v.OrderingProvider = NewXCNSlice(input.Index(16))
	v.OrderCallbackPhoneNumber = NewXTNSlice(input.Index(17))
	v.PlacerField1 = NewST(input.Index(18).Index(0))
	v.PlacerField2 = NewST(input.Index(19).Index(0))
	v.FillerField1 = NewST(input.Index(20).Index(0))
	v.FillerField2 = NewST(input.Index(21).Index(0))
	v.ResultsRptStatusChngDateTime = NewDTM(input.Index(22).Index(0))
	v.ChargeToPractice = NewMOC(input.Index(23).Index(0))
	v.DiagnosticServSectID = NewID(input.Index(24).Index(0))
	v.ResultStatus = NewID(input.Index(25).Index(0))
	v.ParentResult = NewPRL(input.Index(26).Index(0))
	v.QuantityTiming = NewSTSlice(input.Index(27))
	v.ResultCopiesTo = NewXCNSlice(input.Index(28))
	v.Parent = NewEIP(input.Index(29).Index(0))
	v.TransportationMode = NewID(input.Index(30).Index(0))
	v.ReasonForStudy = NewCWESlice(input.Index(31))
	v.PrincipalResultInterpreter = NewNDL(input.Index(32).Index(0))
	v.AssistantResultInterpreter = NewNDLSlice(input.Index(33))
	v.Technician = NewNDLSlice(input.Index(34))
	v.Transcriptionist = NewNDLSlice(input.Index(35))
	v.ScheduledDateTime = NewDTM(input.Index(36).Index(0))
	v.NumberOfSampleContainers = NewNM(input.Index(37).Index(0))
	v.TransportLogisticsOfCollectedSample = NewCWESlice(input.Index(38))
	v.CollectorsComment = NewCWESlice(input.Index(39))
	v.TransportArrangementResponsibility = NewCWE(input.Index(40).Index(0))
	v.TransportArranged = NewID(input.Index(41).Index(0))
	v.EscortRequired = NewID(input.Index(42).Index(0))
	v.PlannedPatientTransportComment = NewCWESlice(input.Index(43))
	v.ProcedureCode = NewCNE(input.Index(44).Index(0))
	v.ProcedureCodeModifier = NewCNESlice(input.Index(45))
	v.PlacerSupplementalServiceInformation = NewCWESlice(input.Index(46))
	v.FillerSupplementalServiceInformation = NewCWESlice(input.Index(47))
	v.MedicallyNecessaryDuplicateProcedureReason = NewCWE(input.Index(48).Index(0))
	v.ResultHandling = NewCWE(input.Index(49).Index(0))
	v.ParentUniversalServiceIdentifier = NewCWE(input.Index(50).Index(0))
	v.ObservationGroupID = NewEI(input.Index(51).Index(0))
	v.ParentObservationGroupID = NewEI(input.Index(52).Index(0))
	v.AlternatePlacerOrderNumber = NewCXSlice(input.Index(53))
	v.ParentOrder = NewEIP(input.Index(54).Index(0))
	return v
}

// NewOBRSlice will construct a slice of type OBR
func NewOBRSlice(input hl7.HL7Type) []OBR {
	values := make([]OBR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOBR(input.Index(i))
	}
	return values
}

type PDA struct {
	DeathCauseCode                 []CWE
	DeathLocation                  PL
	DeathCertifiedIndicator        ID
	DeathCertificateSignedDateTime DTM
	DeathCertifiedBy               XCN
	AutopsyIndicator               ID
	AutopsyStartAndEndDateTime     DR
	AutopsyPerformedBy             XCN
	CoronerIndicator               ID
}

// NewPDA creates an implementation of PDA
func NewPDA(input hl7.HL7Type) PDA {
	v := PDA{}
	v.DeathCauseCode = NewCWESlice(input.Index(1))
	v.DeathLocation = NewPL(input.Index(2).Index(0))
	v.DeathCertifiedIndicator = NewID(input.Index(3).Index(0))
	v.DeathCertificateSignedDateTime = NewDTM(input.Index(4).Index(0))
	v.DeathCertifiedBy = NewXCN(input.Index(5).Index(0))
	v.AutopsyIndicator = NewID(input.Index(6).Index(0))
	v.AutopsyStartAndEndDateTime = NewDR(input.Index(7).Index(0))
	v.AutopsyPerformedBy = NewXCN(input.Index(8).Index(0))
	v.CoronerIndicator = NewID(input.Index(9).Index(0))
	return v
}

// NewPDASlice will construct a slice of type PDA
func NewPDASlice(input hl7.HL7Type) []PDA {
	values := make([]PDA, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPDA(input.Index(i))
	}
	return values
}

type IN2 struct {
	InsuredsEmployeeID                   []CX
	InsuredsSocialSecurityNumber         ST
	InsuredsEmployersNameAndID           []XCN
	EmployerInformationData              CWE
	MailClaimParty                       []CWE
	MedicareHealthInsCardNumber          ST
	MedicaidCaseName                     []XPN
	MedicaidCaseNumber                   ST
	MilitarySponsorName                  []XPN
	MilitaryIDNumber                     ST
	DependentOfMilitaryRecipient         CWE
	MilitaryOrganization                 ST
	MilitaryStation                      ST
	MilitaryService                      CWE
	MilitaryRankGrade                    CWE
	MilitaryStatus                       CWE
	MilitaryRetireDate                   DT
	MilitaryNonAvailCertOnFile           ID
	BabyCoverage                         ID
	CombineBabyBill                      ID
	BloodDeductible                      ST
	SpecialCoverageApprovalName          []XPN
	SpecialCoverageApprovalTitle         ST
	NonCoveredInsuranceCode              []CWE
	PayorID                              []CX
	PayorSubscriberID                    []CX
	EligibilitySource                    CWE
	RoomCoverageTypeAmount               []RMC
	PolicyTypeAmount                     []PTA
	DailyDeductible                      DDI
	LivingDependency                     CWE
	AmbulatoryStatus                     []CWE
	Citizenship                          []CWE
	PrimaryLanguage                      CWE
	LivingArrangement                    CWE
	PublicityCode                        CWE
	ProtectionIndicator                  ID
	StudentIndicator                     CWE
	Religion                             CWE
	MothersMaidenName                    []XPN
	Nationality                          CWE
	EthnicGroup                          []CWE
	MaritalStatus                        []CWE
	InsuredsEmploymentStartDate          DT
	EmploymentStopDate                   DT
	JobTitle                             ST
	JobCodeClass                         JCC
	JobStatus                            CWE
	EmployerContactPersonName            []XPN
	EmployerContactPersonPhoneNumber     []XTN
	EmployerContactReason                CWE
	InsuredsContactPersonsName           []XPN
	InsuredsContactPersonPhoneNumber     []XTN
	InsuredsContactPersonReason          []CWE
	RelationshipToThePatientStartDate    DT
	RelationshipToThePatientStopDate     []DT
	InsuranceCoContactReason             CWE
	InsuranceCoContactPhoneNumber        []XTN
	PolicyScope                          CWE
	PolicySource                         CWE
	PatientMemberNumber                  CX
	GuarantorsRelationshipToInsured      CWE
	InsuredsPhoneNumberHome              []XTN
	InsuredsEmployerPhoneNumber          []XTN
	MilitaryHandicappedProgram           CWE
	SuspendFlag                          ID
	CopayLimitFlag                       ID
	StoplossLimitFlag                    ID
	InsuredOrganizationNameAndID         []XON
	InsuredEmployerOrganizationNameAndID []XON
	Race                                 []CWE
	PatientsRelationshipToInsured        CWE
}

// NewIN2 creates an implementation of IN2
func NewIN2(input hl7.HL7Type) IN2 {
	v := IN2{}
	v.InsuredsEmployeeID = NewCXSlice(input.Index(1))
	v.InsuredsSocialSecurityNumber = NewST(input.Index(2).Index(0))
	v.InsuredsEmployersNameAndID = NewXCNSlice(input.Index(3))
	v.EmployerInformationData = NewCWE(input.Index(4).Index(0))
	v.MailClaimParty = NewCWESlice(input.Index(5))
	v.MedicareHealthInsCardNumber = NewST(input.Index(6).Index(0))
	v.MedicaidCaseName = NewXPNSlice(input.Index(7))
	v.MedicaidCaseNumber = NewST(input.Index(8).Index(0))
	v.MilitarySponsorName = NewXPNSlice(input.Index(9))
	v.MilitaryIDNumber = NewST(input.Index(10).Index(0))
	v.DependentOfMilitaryRecipient = NewCWE(input.Index(11).Index(0))
	v.MilitaryOrganization = NewST(input.Index(12).Index(0))
	v.MilitaryStation = NewST(input.Index(13).Index(0))
	v.MilitaryService = NewCWE(input.Index(14).Index(0))
	v.MilitaryRankGrade = NewCWE(input.Index(15).Index(0))
	v.MilitaryStatus = NewCWE(input.Index(16).Index(0))
	v.MilitaryRetireDate = NewDT(input.Index(17).Index(0))
	v.MilitaryNonAvailCertOnFile = NewID(input.Index(18).Index(0))
	v.BabyCoverage = NewID(input.Index(19).Index(0))
	v.CombineBabyBill = NewID(input.Index(20).Index(0))
	v.BloodDeductible = NewST(input.Index(21).Index(0))
	v.SpecialCoverageApprovalName = NewXPNSlice(input.Index(22))
	v.SpecialCoverageApprovalTitle = NewST(input.Index(23).Index(0))
	v.NonCoveredInsuranceCode = NewCWESlice(input.Index(24))
	v.PayorID = NewCXSlice(input.Index(25))
	v.PayorSubscriberID = NewCXSlice(input.Index(26))
	v.EligibilitySource = NewCWE(input.Index(27).Index(0))
	v.RoomCoverageTypeAmount = NewRMCSlice(input.Index(28))
	v.PolicyTypeAmount = NewPTASlice(input.Index(29))
	v.DailyDeductible = NewDDI(input.Index(30).Index(0))
	v.LivingDependency = NewCWE(input.Index(31).Index(0))
	v.AmbulatoryStatus = NewCWESlice(input.Index(32))
	v.Citizenship = NewCWESlice(input.Index(33))
	v.PrimaryLanguage = NewCWE(input.Index(34).Index(0))
	v.LivingArrangement = NewCWE(input.Index(35).Index(0))
	v.PublicityCode = NewCWE(input.Index(36).Index(0))
	v.ProtectionIndicator = NewID(input.Index(37).Index(0))
	v.StudentIndicator = NewCWE(input.Index(38).Index(0))
	v.Religion = NewCWE(input.Index(39).Index(0))
	v.MothersMaidenName = NewXPNSlice(input.Index(40))
	v.Nationality = NewCWE(input.Index(41).Index(0))
	v.EthnicGroup = NewCWESlice(input.Index(42))
	v.MaritalStatus = NewCWESlice(input.Index(43))
	v.InsuredsEmploymentStartDate = NewDT(input.Index(44).Index(0))
	v.EmploymentStopDate = NewDT(input.Index(45).Index(0))
	v.JobTitle = NewST(input.Index(46).Index(0))
	v.JobCodeClass = NewJCC(input.Index(47).Index(0))
	v.JobStatus = NewCWE(input.Index(48).Index(0))
	v.EmployerContactPersonName = NewXPNSlice(input.Index(49))
	v.EmployerContactPersonPhoneNumber = NewXTNSlice(input.Index(50))
	v.EmployerContactReason = NewCWE(input.Index(51).Index(0))
	v.InsuredsContactPersonsName = NewXPNSlice(input.Index(52))
	v.InsuredsContactPersonPhoneNumber = NewXTNSlice(input.Index(53))
	v.InsuredsContactPersonReason = NewCWESlice(input.Index(54))
	v.RelationshipToThePatientStartDate = NewDT(input.Index(55).Index(0))
	v.RelationshipToThePatientStopDate = NewDTSlice(input.Index(56))
	v.InsuranceCoContactReason = NewCWE(input.Index(57).Index(0))
	v.InsuranceCoContactPhoneNumber = NewXTNSlice(input.Index(58))
	v.PolicyScope = NewCWE(input.Index(59).Index(0))
	v.PolicySource = NewCWE(input.Index(60).Index(0))
	v.PatientMemberNumber = NewCX(input.Index(61).Index(0))
	v.GuarantorsRelationshipToInsured = NewCWE(input.Index(62).Index(0))
	v.InsuredsPhoneNumberHome = NewXTNSlice(input.Index(63))
	v.InsuredsEmployerPhoneNumber = NewXTNSlice(input.Index(64))
	v.MilitaryHandicappedProgram = NewCWE(input.Index(65).Index(0))
	v.SuspendFlag = NewID(input.Index(66).Index(0))
	v.CopayLimitFlag = NewID(input.Index(67).Index(0))
	v.StoplossLimitFlag = NewID(input.Index(68).Index(0))
	v.InsuredOrganizationNameAndID = NewXONSlice(input.Index(69))
	v.InsuredEmployerOrganizationNameAndID = NewXONSlice(input.Index(70))
	v.Race = NewCWESlice(input.Index(71))
	v.PatientsRelationshipToInsured = NewCWE(input.Index(72).Index(0))
	return v
}

// NewIN2Slice will construct a slice of type IN2
func NewIN2Slice(input hl7.HL7Type) []IN2 {
	values := make([]IN2, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewIN2(input.Index(i))
	}
	return values
}

type IN3 struct {
	SetID                              SI
	CertificationNumber                CX
	CertifiedBy                        []XCN
	CertificationRequired              ID
	Penalty                            MOP
	CertificationDateTime              DTM
	CertificationModifyDateTime        DTM
	Operator                           []XCN
	CertificationBeginDate             DT
	CertificationEndDate               DT
	Days                               DTN
	NonConcurCodeDescription           CWE
	NonConcurEffectiveDateTime         DTM
	PhysicianReviewer                  []XCN
	CertificationContact               ST
	CertificationContactPhoneNumber    []XTN
	AppealReason                       CWE
	CertificationAgency                CWE
	CertificationAgencyPhoneNumber     []XTN
	PreCertificationRequirement        []ICD
	CaseManager                        ST
	SecondOpinionDate                  DT
	SecondOpinionStatus                CWE
	SecondOpinionDocumentationReceived []CWE
	SecondOpinionPhysician             []XCN
	CertificationType                  CWE
	CertificationCategory              CWE
}

// NewIN3 creates an implementation of IN3
func NewIN3(input hl7.HL7Type) IN3 {
	v := IN3{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.CertificationNumber = NewCX(input.Index(2).Index(0))
	v.CertifiedBy = NewXCNSlice(input.Index(3))
	v.CertificationRequired = NewID(input.Index(4).Index(0))
	v.Penalty = NewMOP(input.Index(5).Index(0))
	v.CertificationDateTime = NewDTM(input.Index(6).Index(0))
	v.CertificationModifyDateTime = NewDTM(input.Index(7).Index(0))
	v.Operator = NewXCNSlice(input.Index(8))
	v.CertificationBeginDate = NewDT(input.Index(9).Index(0))
	v.CertificationEndDate = NewDT(input.Index(10).Index(0))
	v.Days = NewDTN(input.Index(11).Index(0))
	v.NonConcurCodeDescription = NewCWE(input.Index(12).Index(0))
	v.NonConcurEffectiveDateTime = NewDTM(input.Index(13).Index(0))
	v.PhysicianReviewer = NewXCNSlice(input.Index(14))
	v.CertificationContact = NewST(input.Index(15).Index(0))
	v.CertificationContactPhoneNumber = NewXTNSlice(input.Index(16))
	v.AppealReason = NewCWE(input.Index(17).Index(0))
	v.CertificationAgency = NewCWE(input.Index(18).Index(0))
	v.CertificationAgencyPhoneNumber = NewXTNSlice(input.Index(19))
	v.PreCertificationRequirement = NewICDSlice(input.Index(20))
	v.CaseManager = NewST(input.Index(21).Index(0))
	v.SecondOpinionDate = NewDT(input.Index(22).Index(0))
	v.SecondOpinionStatus = NewCWE(input.Index(23).Index(0))
	v.SecondOpinionDocumentationReceived = NewCWESlice(input.Index(24))
	v.SecondOpinionPhysician = NewXCNSlice(input.Index(25))
	v.CertificationType = NewCWE(input.Index(26).Index(0))
	v.CertificationCategory = NewCWE(input.Index(27).Index(0))
	return v
}

// NewIN3Slice will construct a slice of type IN3
func NewIN3Slice(input hl7.HL7Type) []IN3 {
	values := make([]IN3, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewIN3(input.Index(i))
	}
	return values
}

type CER struct {
	SetID                              SI
	SerialNumber                       ST
	Version                            ST
	GrantingAuthority                  XON
	IssuingAuthority                   XCN
	Signature                          ED
	GrantingCountry                    ID
	GrantingStateProvince              CWE
	GrantingCountyParish               CWE
	CertificateType                    CWE
	CertificateDomain                  CWE
	SubjectID                          EI
	SubjectName                        ST
	SubjectDirectoryAttributeExtension []CWE
	SubjectPublicKeyInfo               CWE
	AuthorityKeyIdentifier             CWE
	BasicConstraint                    ID
	CrlDistributionPoint               []CWE
	JurisdictionCountry                ID
	JurisdictionStateProvince          CWE
	JurisdictionCountyParish           CWE
	JurisdictionBreadth                []CWE
	GrantingDate                       DTM
	IssuingDate                        DTM
	ActivationDate                     DTM
	InactivationDate                   DTM
	ExpirationDate                     DTM
	RenewalDate                        DTM
	RevocationDate                     DTM
	RevocationReasonCode               CWE
	CertificateStatusCode              CWE
}

// NewCER creates an implementation of CER
func NewCER(input hl7.HL7Type) CER {
	v := CER{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.SerialNumber = NewST(input.Index(2).Index(0))
	v.Version = NewST(input.Index(3).Index(0))
	v.GrantingAuthority = NewXON(input.Index(4).Index(0))
	v.IssuingAuthority = NewXCN(input.Index(5).Index(0))
	v.Signature = NewED(input.Index(6).Index(0))
	v.GrantingCountry = NewID(input.Index(7).Index(0))
	v.GrantingStateProvince = NewCWE(input.Index(8).Index(0))
	v.GrantingCountyParish = NewCWE(input.Index(9).Index(0))
	v.CertificateType = NewCWE(input.Index(10).Index(0))
	v.CertificateDomain = NewCWE(input.Index(11).Index(0))
	v.SubjectID = NewEI(input.Index(12).Index(0))
	v.SubjectName = NewST(input.Index(13).Index(0))
	v.SubjectDirectoryAttributeExtension = NewCWESlice(input.Index(14))
	v.SubjectPublicKeyInfo = NewCWE(input.Index(15).Index(0))
	v.AuthorityKeyIdentifier = NewCWE(input.Index(16).Index(0))
	v.BasicConstraint = NewID(input.Index(17).Index(0))
	v.CrlDistributionPoint = NewCWESlice(input.Index(18))
	v.JurisdictionCountry = NewID(input.Index(19).Index(0))
	v.JurisdictionStateProvince = NewCWE(input.Index(20).Index(0))
	v.JurisdictionCountyParish = NewCWE(input.Index(21).Index(0))
	v.JurisdictionBreadth = NewCWESlice(input.Index(22))
	v.GrantingDate = NewDTM(input.Index(23).Index(0))
	v.IssuingDate = NewDTM(input.Index(24).Index(0))
	v.ActivationDate = NewDTM(input.Index(25).Index(0))
	v.InactivationDate = NewDTM(input.Index(26).Index(0))
	v.ExpirationDate = NewDTM(input.Index(27).Index(0))
	v.RenewalDate = NewDTM(input.Index(28).Index(0))
	v.RevocationDate = NewDTM(input.Index(29).Index(0))
	v.RevocationReasonCode = NewCWE(input.Index(30).Index(0))
	v.CertificateStatusCode = NewCWE(input.Index(31).Index(0))
	return v
}

// NewCERSlice will construct a slice of type CER
func NewCERSlice(input hl7.HL7Type) []CER {
	values := make([]CER, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCER(input.Index(i))
	}
	return values
}

type SCP struct {
	NumberOfDecontaminationSterilizationDevices NM
	LaborCalculationType                        CWE
	DateFormat                                  CWE
	DeviceNumber                                EI
	DeviceName                                  ST
	DeviceModelName                             ST
	DeviceType                                  CWE
	LotControl                                  CWE
}

// NewSCP creates an implementation of SCP
func NewSCP(input hl7.HL7Type) SCP {
	v := SCP{}
	v.NumberOfDecontaminationSterilizationDevices = NewNM(input.Index(1).Index(0))
	v.LaborCalculationType = NewCWE(input.Index(2).Index(0))
	v.DateFormat = NewCWE(input.Index(3).Index(0))
	v.DeviceNumber = NewEI(input.Index(4).Index(0))
	v.DeviceName = NewST(input.Index(5).Index(0))
	v.DeviceModelName = NewST(input.Index(6).Index(0))
	v.DeviceType = NewCWE(input.Index(7).Index(0))
	v.LotControl = NewCWE(input.Index(8).Index(0))
	return v
}

// NewSCPSlice will construct a slice of type SCP
func NewSCPSlice(input hl7.HL7Type) []SCP {
	values := make([]SCP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSCP(input.Index(i))
	}
	return values
}

type BPX struct {
	SetID                       SI
	BpDispenseStatus            CWE
	BpStatus                    ID
	BpDateTimeOfStatus          DTM
	BcDonationID                EI
	BcComponent                 CNE
	BcDonationTypeIntendedUse   CNE
	CpCommercialProduct         CWE
	CpManufacturer              XON
	CpLotNumber                 EI
	BpBloodGroup                CNE
	BcSpecialTesting            []CNE
	BpExpirationDateTime        DTM
	BpQuantity                  NM
	BpAmount                    NM
	BpUnits                     CWE
	BpUniqueID                  EI
	BpActualDispensedToLocation PL
	BpActualDispensedToAddress  XAD
	BpDispensedToReceiver       XCN
	BpDispensingIndividual      XCN
}

// NewBPX creates an implementation of BPX
func NewBPX(input hl7.HL7Type) BPX {
	v := BPX{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.BpDispenseStatus = NewCWE(input.Index(2).Index(0))
	v.BpStatus = NewID(input.Index(3).Index(0))
	v.BpDateTimeOfStatus = NewDTM(input.Index(4).Index(0))
	v.BcDonationID = NewEI(input.Index(5).Index(0))
	v.BcComponent = NewCNE(input.Index(6).Index(0))
	v.BcDonationTypeIntendedUse = NewCNE(input.Index(7).Index(0))
	v.CpCommercialProduct = NewCWE(input.Index(8).Index(0))
	v.CpManufacturer = NewXON(input.Index(9).Index(0))
	v.CpLotNumber = NewEI(input.Index(10).Index(0))
	v.BpBloodGroup = NewCNE(input.Index(11).Index(0))
	v.BcSpecialTesting = NewCNESlice(input.Index(12))
	v.BpExpirationDateTime = NewDTM(input.Index(13).Index(0))
	v.BpQuantity = NewNM(input.Index(14).Index(0))
	v.BpAmount = NewNM(input.Index(15).Index(0))
	v.BpUnits = NewCWE(input.Index(16).Index(0))
	v.BpUniqueID = NewEI(input.Index(17).Index(0))
	v.BpActualDispensedToLocation = NewPL(input.Index(18).Index(0))
	v.BpActualDispensedToAddress = NewXAD(input.Index(19).Index(0))
	v.BpDispensedToReceiver = NewXCN(input.Index(20).Index(0))
	v.BpDispensingIndividual = NewXCN(input.Index(21).Index(0))
	return v
}

// NewBPXSlice will construct a slice of type BPX
func NewBPXSlice(input hl7.HL7Type) []BPX {
	values := make([]BPX, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBPX(input.Index(i))
	}
	return values
}

type QRI struct {
	CandidateConfidence NM
	MatchReasonCode     []CWE
	AlgorithmDescriptor CWE
}

// NewQRI creates an implementation of QRI
func NewQRI(input hl7.HL7Type) QRI {
	v := QRI{}
	v.CandidateConfidence = NewNM(input.Index(1).Index(0))
	v.MatchReasonCode = NewCWESlice(input.Index(2))
	v.AlgorithmDescriptor = NewCWE(input.Index(3).Index(0))
	return v
}

// NewQRISlice will construct a slice of type QRI
func NewQRISlice(input hl7.HL7Type) []QRI {
	values := make([]QRI, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQRI(input.Index(i))
	}
	return values
}

type ODS struct {
	Type                           ID
	ServicePeriod                  []CWE
	DietSupplementOrPreferenceCode []CWE
	TextInstruction                []ST
}

// NewODS creates an implementation of ODS
func NewODS(input hl7.HL7Type) ODS {
	v := ODS{}
	v.Type = NewID(input.Index(1).Index(0))
	v.ServicePeriod = NewCWESlice(input.Index(2))
	v.DietSupplementOrPreferenceCode = NewCWESlice(input.Index(3))
	v.TextInstruction = NewSTSlice(input.Index(4))
	return v
}

// NewODSSlice will construct a slice of type ODS
func NewODSSlice(input hl7.HL7Type) []ODS {
	values := make([]ODS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewODS(input.Index(i))
	}
	return values
}

type GP2 struct {
	RevenueCode                         CWE
	NumberOfServiceUnits                NM
	Charge                              CP
	ReimbursementActionCode             CWE
	DenialOrRejectionCode               CWE
	OceEditCode                         []CWE
	AmbulatoryPaymentClassificationCode CWE
	ModifierEditCode                    []CWE
	PaymentAdjustmentCode               CWE
	PackagingStatusCode                 CWE
	ExpectedCmsPaymentAmount            CP
	ReimbursementTypeCode               CWE
	CoPayAmount                         CP
	PayRatePerServiceUnit               NM
}

// NewGP2 creates an implementation of GP2
func NewGP2(input hl7.HL7Type) GP2 {
	v := GP2{}
	v.RevenueCode = NewCWE(input.Index(1).Index(0))
	v.NumberOfServiceUnits = NewNM(input.Index(2).Index(0))
	v.Charge = NewCP(input.Index(3).Index(0))
	v.ReimbursementActionCode = NewCWE(input.Index(4).Index(0))
	v.DenialOrRejectionCode = NewCWE(input.Index(5).Index(0))
	v.OceEditCode = NewCWESlice(input.Index(6))
	v.AmbulatoryPaymentClassificationCode = NewCWE(input.Index(7).Index(0))
	v.ModifierEditCode = NewCWESlice(input.Index(8))
	v.PaymentAdjustmentCode = NewCWE(input.Index(9).Index(0))
	v.PackagingStatusCode = NewCWE(input.Index(10).Index(0))
	v.ExpectedCmsPaymentAmount = NewCP(input.Index(11).Index(0))
	v.ReimbursementTypeCode = NewCWE(input.Index(12).Index(0))
	v.CoPayAmount = NewCP(input.Index(13).Index(0))
	v.PayRatePerServiceUnit = NewNM(input.Index(14).Index(0))
	return v
}

// NewGP2Slice will construct a slice of type GP2
func NewGP2Slice(input hl7.HL7Type) []GP2 {
	values := make([]GP2, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewGP2(input.Index(i))
	}
	return values
}

type FAC struct {
	FacilityIDFac                       EI
	FacilityType                        ID
	FacilityAddress                     []XAD
	FacilityTelecommunication           XTN
	ContactPerson                       []XCN
	ContactTitle                        []ST
	ContactAddress                      []XAD
	ContactTelecommunication            []XTN
	SignatureAuthority                  []XCN
	SignatureAuthorityTitle             ST
	SignatureAuthorityAddress           []XAD
	SignatureAuthorityTelecommunication XTN
}

// NewFAC creates an implementation of FAC
func NewFAC(input hl7.HL7Type) FAC {
	v := FAC{}
	v.FacilityIDFac = NewEI(input.Index(1).Index(0))
	v.FacilityType = NewID(input.Index(2).Index(0))
	v.FacilityAddress = NewXADSlice(input.Index(3))
	v.FacilityTelecommunication = NewXTN(input.Index(4).Index(0))
	v.ContactPerson = NewXCNSlice(input.Index(5))
	v.ContactTitle = NewSTSlice(input.Index(6))
	v.ContactAddress = NewXADSlice(input.Index(7))
	v.ContactTelecommunication = NewXTNSlice(input.Index(8))
	v.SignatureAuthority = NewXCNSlice(input.Index(9))
	v.SignatureAuthorityTitle = NewST(input.Index(10).Index(0))
	v.SignatureAuthorityAddress = NewXADSlice(input.Index(11))
	v.SignatureAuthorityTelecommunication = NewXTN(input.Index(12).Index(0))
	return v
}

// NewFACSlice will construct a slice of type FAC
func NewFACSlice(input hl7.HL7Type) []FAC {
	values := make([]FAC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewFAC(input.Index(i))
	}
	return values
}

type NSC struct {
	ApplicationChangeType CWE
	CurrentCPU            ST
	CurrentFileserver     ST
	CurrentApplication    HD
	CurrentFacility       HD
	NewCPU                ST
	NewFileserver         ST
	NewApplication        HD
	NewFacility           HD
}

// NewNSC creates an implementation of NSC
func NewNSC(input hl7.HL7Type) NSC {
	v := NSC{}
	v.ApplicationChangeType = NewCWE(input.Index(1).Index(0))
	v.CurrentCPU = NewST(input.Index(2).Index(0))
	v.CurrentFileserver = NewST(input.Index(3).Index(0))
	v.CurrentApplication = NewHD(input.Index(4).Index(0))
	v.CurrentFacility = NewHD(input.Index(5).Index(0))
	v.NewCPU = NewST(input.Index(6).Index(0))
	v.NewFileserver = NewST(input.Index(7).Index(0))
	v.NewApplication = NewHD(input.Index(8).Index(0))
	v.NewFacility = NewHD(input.Index(9).Index(0))
	return v
}

// NewNSCSlice will construct a slice of type NSC
func NewNSCSlice(input hl7.HL7Type) []NSC {
	values := make([]NSC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewNSC(input.Index(i))
	}
	return values
}

type UB1 struct {
	SetID                   ST
	BloodDeductible         ST
	BloodFurnishedPints     ST
	BloodReplacedPints      ST
	BloodNotReplacedPints   ST
	CoInsuranceDays         ST
	ConditionCode           []ST
	CoveredDays             ST
	NonCoveredDays          ST
	ValueAmountCode         []ST
	NumberOfGraceDays       ST
	SpecialProgramIndicator ST
	PsroUrApprovalIndicator ST
	PsroUrApprovedStayFm    ST
	PsroUrApprovedStayTo    ST
	Occurrence              []ST
	OccurrenceSpan          ST
	OccurSpanStartDate      ST
	OccurSpanEndDate        ST
	Ub82Locator2            ST
	Ub82Locator9            ST
	Ub82Locator27           ST
	Ub82Locator45           ST
}

// NewUB1 creates an implementation of UB1
func NewUB1(input hl7.HL7Type) UB1 {
	v := UB1{}
	v.SetID = NewST(input.Index(1).Index(0))
	v.BloodDeductible = NewST(input.Index(2).Index(0))
	v.BloodFurnishedPints = NewST(input.Index(3).Index(0))
	v.BloodReplacedPints = NewST(input.Index(4).Index(0))
	v.BloodNotReplacedPints = NewST(input.Index(5).Index(0))
	v.CoInsuranceDays = NewST(input.Index(6).Index(0))
	v.ConditionCode = NewSTSlice(input.Index(7))
	v.CoveredDays = NewST(input.Index(8).Index(0))
	v.NonCoveredDays = NewST(input.Index(9).Index(0))
	v.ValueAmountCode = NewSTSlice(input.Index(10))
	v.NumberOfGraceDays = NewST(input.Index(11).Index(0))
	v.SpecialProgramIndicator = NewST(input.Index(12).Index(0))
	v.PsroUrApprovalIndicator = NewST(input.Index(13).Index(0))
	v.PsroUrApprovedStayFm = NewST(input.Index(14).Index(0))
	v.PsroUrApprovedStayTo = NewST(input.Index(15).Index(0))
	v.Occurrence = NewSTSlice(input.Index(16))
	v.OccurrenceSpan = NewST(input.Index(17).Index(0))
	v.OccurSpanStartDate = NewST(input.Index(18).Index(0))
	v.OccurSpanEndDate = NewST(input.Index(19).Index(0))
	v.Ub82Locator2 = NewST(input.Index(20).Index(0))
	v.Ub82Locator9 = NewST(input.Index(21).Index(0))
	v.Ub82Locator27 = NewST(input.Index(22).Index(0))
	v.Ub82Locator45 = NewST(input.Index(23).Index(0))
	return v
}

// NewUB1Slice will construct a slice of type UB1
func NewUB1Slice(input hl7.HL7Type) []UB1 {
	values := make([]UB1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewUB1(input.Index(i))
	}
	return values
}

type TQ1 struct {
	SetID                SI
	Quantity             CQ
	RepeatPattern        []RPT
	ExplicitTime         []TM
	RelativeTimeAndUnits []CQ
	ServiceDuration      CQ
	StartDateTime        DTM
	EndDateTime          DTM
	Priority             []CWE
	ConditionText        TX
	TextInstruction      TX
	Conjunction          ID
	OccurrenceDuration   CQ
	TotalOccurrences     NM
}

// NewTQ1 creates an implementation of TQ1
func NewTQ1(input hl7.HL7Type) TQ1 {
	v := TQ1{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.Quantity = NewCQ(input.Index(2).Index(0))
	v.RepeatPattern = NewRPTSlice(input.Index(3))
	v.ExplicitTime = NewTMSlice(input.Index(4))
	v.RelativeTimeAndUnits = NewCQSlice(input.Index(5))
	v.ServiceDuration = NewCQ(input.Index(6).Index(0))
	v.StartDateTime = NewDTM(input.Index(7).Index(0))
	v.EndDateTime = NewDTM(input.Index(8).Index(0))
	v.Priority = NewCWESlice(input.Index(9))
	v.ConditionText = NewTX(input.Index(10).Index(0))
	v.TextInstruction = NewTX(input.Index(11).Index(0))
	v.Conjunction = NewID(input.Index(12).Index(0))
	v.OccurrenceDuration = NewCQ(input.Index(13).Index(0))
	v.TotalOccurrences = NewNM(input.Index(14).Index(0))
	return v
}

// NewTQ1Slice will construct a slice of type TQ1
func NewTQ1Slice(input hl7.HL7Type) []TQ1 {
	values := make([]TQ1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewTQ1(input.Index(i))
	}
	return values
}

type PEO struct {
	EventIdentifiersUsed                  []CWE
	EventSymptomDiagnosisCode             []CWE
	EventOnsetDateTime                    DTM
	EventExacerbationDateTime             DTM
	EventImprovedDateTime                 DTM
	EventEndedDataTime                    DTM
	EventLocationOccurredAddress          []XAD
	EventQualification                    []ID
	EventSerious                          ID
	EventExpected                         ID
	EventOutcome                          []ID
	PatientOutcome                        ID
	EventDescriptionFromOthers            []FT
	EventDescriptionFromOriginalReporter  []FT
	EventDescriptionFromPatient           []FT
	EventDescriptionFromPractitioner      []FT
	EventDescriptionFromAutopsy           []FT
	CauseOfDeath                          []CWE
	PrimaryObserverName                   []XPN
	PrimaryObserverAddress                []XAD
	PrimaryObserverTelephone              []XTN
	PrimaryObserversQualification         ID
	ConfirmationProvidedBy                ID
	PrimaryObserverAwareDateTime          DTM
	PrimaryObserversIdentityMayBeDivulged ID
}

// NewPEO creates an implementation of PEO
func NewPEO(input hl7.HL7Type) PEO {
	v := PEO{}
	v.EventIdentifiersUsed = NewCWESlice(input.Index(1))
	v.EventSymptomDiagnosisCode = NewCWESlice(input.Index(2))
	v.EventOnsetDateTime = NewDTM(input.Index(3).Index(0))
	v.EventExacerbationDateTime = NewDTM(input.Index(4).Index(0))
	v.EventImprovedDateTime = NewDTM(input.Index(5).Index(0))
	v.EventEndedDataTime = NewDTM(input.Index(6).Index(0))
	v.EventLocationOccurredAddress = NewXADSlice(input.Index(7))
	v.EventQualification = NewIDSlice(input.Index(8))
	v.EventSerious = NewID(input.Index(9).Index(0))
	v.EventExpected = NewID(input.Index(10).Index(0))
	v.EventOutcome = NewIDSlice(input.Index(11))
	v.PatientOutcome = NewID(input.Index(12).Index(0))
	v.EventDescriptionFromOthers = NewFTSlice(input.Index(13))
	v.EventDescriptionFromOriginalReporter = NewFTSlice(input.Index(14))
	v.EventDescriptionFromPatient = NewFTSlice(input.Index(15))
	v.EventDescriptionFromPractitioner = NewFTSlice(input.Index(16))
	v.EventDescriptionFromAutopsy = NewFTSlice(input.Index(17))
	v.CauseOfDeath = NewCWESlice(input.Index(18))
	v.PrimaryObserverName = NewXPNSlice(input.Index(19))
	v.PrimaryObserverAddress = NewXADSlice(input.Index(20))
	v.PrimaryObserverTelephone = NewXTNSlice(input.Index(21))
	v.PrimaryObserversQualification = NewID(input.Index(22).Index(0))
	v.ConfirmationProvidedBy = NewID(input.Index(23).Index(0))
	v.PrimaryObserverAwareDateTime = NewDTM(input.Index(24).Index(0))
	v.PrimaryObserversIdentityMayBeDivulged = NewID(input.Index(25).Index(0))
	return v
}

// NewPEOSlice will construct a slice of type PEO
func NewPEOSlice(input hl7.HL7Type) []PEO {
	values := make([]PEO, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPEO(input.Index(i))
	}
	return values
}

type RMI struct {
	RiskManagementIncidentCode CWE
	DateTimeIncident           DTM
	IncidentTypeCode           CWE
}

// NewRMI creates an implementation of RMI
func NewRMI(input hl7.HL7Type) RMI {
	v := RMI{}
	v.RiskManagementIncidentCode = NewCWE(input.Index(1).Index(0))
	v.DateTimeIncident = NewDTM(input.Index(2).Index(0))
	v.IncidentTypeCode = NewCWE(input.Index(3).Index(0))
	return v
}

// NewRMISlice will construct a slice of type RMI
func NewRMISlice(input hl7.HL7Type) []RMI {
	values := make([]RMI, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRMI(input.Index(i))
	}
	return values
}

type ECR struct {
	CommandResponse           CWE
	DateTimeCompleted         DTM
	CommandResponseParameters []TX
}

// NewECR creates an implementation of ECR
func NewECR(input hl7.HL7Type) ECR {
	v := ECR{}
	v.CommandResponse = NewCWE(input.Index(1).Index(0))
	v.DateTimeCompleted = NewDTM(input.Index(2).Index(0))
	v.CommandResponseParameters = NewTXSlice(input.Index(3))
	return v
}

// NewECRSlice will construct a slice of type ECR
func NewECRSlice(input hl7.HL7Type) []ECR {
	values := make([]ECR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewECR(input.Index(i))
	}
	return values
}

type RGS struct {
	SetID             SI
	SegmentActionCode ID
	ResourceGroupID   CWE
}

// NewRGS creates an implementation of RGS
func NewRGS(input hl7.HL7Type) RGS {
	v := RGS{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.SegmentActionCode = NewID(input.Index(2).Index(0))
	v.ResourceGroupID = NewCWE(input.Index(3).Index(0))
	return v
}

// NewRGSSlice will construct a slice of type RGS
func NewRGSSlice(input hl7.HL7Type) []RGS {
	values := make([]RGS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRGS(input.Index(i))
	}
	return values
}

type RXC struct {
	RxComponentType                  ID
	ComponentCode                    CWE
	ComponentAmount                  NM
	ComponentUnits                   CWE
	ComponentStrength                NM
	ComponentStrengthUnits           CWE
	SupplementaryCode                []CWE
	ComponentDrugStrengthVolume      NM
	ComponentDrugStrengthVolumeUnits CWE
}

// NewRXC creates an implementation of RXC
func NewRXC(input hl7.HL7Type) RXC {
	v := RXC{}
	v.RxComponentType = NewID(input.Index(1).Index(0))
	v.ComponentCode = NewCWE(input.Index(2).Index(0))
	v.ComponentAmount = NewNM(input.Index(3).Index(0))
	v.ComponentUnits = NewCWE(input.Index(4).Index(0))
	v.ComponentStrength = NewNM(input.Index(5).Index(0))
	v.ComponentStrengthUnits = NewCWE(input.Index(6).Index(0))
	v.SupplementaryCode = NewCWESlice(input.Index(7))
	v.ComponentDrugStrengthVolume = NewNM(input.Index(8).Index(0))
	v.ComponentDrugStrengthVolumeUnits = NewCWE(input.Index(9).Index(0))
	return v
}

// NewRXCSlice will construct a slice of type RXC
func NewRXCSlice(input hl7.HL7Type) []RXC {
	values := make([]RXC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRXC(input.Index(i))
	}
	return values
}

type REL struct {
	SetID                                           SI
	RelationshipType                                CWE
	ThisRelationshipInstanceIdentifier              EI
	SourceInformationInstanceIdentifier             EI
	TargetInformationInstanceIdentifier             EI
	AssertingEntityInstanceID                       EI
	AssertingPerson                                 XCN
	AssertingOrganization                           XON
	AssertorAddress                                 XAD
	AssertorContact                                 XTN
	AssertionDateRange                              DR
	NegationIndicator                               ID
	CertaintyOfRelationship                         CWE
	PriorityNo                                      NM
	PrioritySequenceNorelPreferenceForConsideration NM
	SeparabilityIndicator                           ID
}

// NewREL creates an implementation of REL
func NewREL(input hl7.HL7Type) REL {
	v := REL{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.RelationshipType = NewCWE(input.Index(2).Index(0))
	v.ThisRelationshipInstanceIdentifier = NewEI(input.Index(3).Index(0))
	v.SourceInformationInstanceIdentifier = NewEI(input.Index(4).Index(0))
	v.TargetInformationInstanceIdentifier = NewEI(input.Index(5).Index(0))
	v.AssertingEntityInstanceID = NewEI(input.Index(6).Index(0))
	v.AssertingPerson = NewXCN(input.Index(7).Index(0))
	v.AssertingOrganization = NewXON(input.Index(8).Index(0))
	v.AssertorAddress = NewXAD(input.Index(9).Index(0))
	v.AssertorContact = NewXTN(input.Index(10).Index(0))
	v.AssertionDateRange = NewDR(input.Index(11).Index(0))
	v.NegationIndicator = NewID(input.Index(12).Index(0))
	v.CertaintyOfRelationship = NewCWE(input.Index(13).Index(0))
	v.PriorityNo = NewNM(input.Index(14).Index(0))
	v.PrioritySequenceNorelPreferenceForConsideration = NewNM(input.Index(15).Index(0))
	v.SeparabilityIndicator = NewID(input.Index(16).Index(0))
	return v
}

// NewRELSlice will construct a slice of type REL
func NewRELSlice(input hl7.HL7Type) []REL {
	values := make([]REL, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewREL(input.Index(i))
	}
	return values
}

type SID struct {
	ApplicationMethodIdentifier     CWE
	SubstanceLotNumber              ST
	SubstanceContainerIdentifier    ST
	SubstanceManufacturerIdentifier CWE
}

// NewSID creates an implementation of SID
func NewSID(input hl7.HL7Type) SID {
	v := SID{}
	v.ApplicationMethodIdentifier = NewCWE(input.Index(1).Index(0))
	v.SubstanceLotNumber = NewST(input.Index(2).Index(0))
	v.SubstanceContainerIdentifier = NewST(input.Index(3).Index(0))
	v.SubstanceManufacturerIdentifier = NewCWE(input.Index(4).Index(0))
	return v
}

// NewSIDSlice will construct a slice of type SID
func NewSIDSlice(input hl7.HL7Type) []SID {
	values := make([]SID, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSID(input.Index(i))
	}
	return values
}

type AFF struct {
	SetID                                        SI
	ProfessionalOrganization                     XON
	ProfessionalOrganizationAddress              XAD
	ProfessionalOrganizationAffiliationDateRange []DR
	ProfessionalAffiliationAdditionalInformation ST
}

// NewAFF creates an implementation of AFF
func NewAFF(input hl7.HL7Type) AFF {
	v := AFF{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.ProfessionalOrganization = NewXON(input.Index(2).Index(0))
	v.ProfessionalOrganizationAddress = NewXAD(input.Index(3).Index(0))
	v.ProfessionalOrganizationAffiliationDateRange = NewDRSlice(input.Index(4))
	v.ProfessionalAffiliationAdditionalInformation = NewST(input.Index(5).Index(0))
	return v
}

// NewAFFSlice will construct a slice of type AFF
func NewAFFSlice(input hl7.HL7Type) []AFF {
	values := make([]AFF, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewAFF(input.Index(i))
	}
	return values
}

type IVT struct {
	SetID                             SI
	InventoryLocationIdentifier       EI
	InventoryLocationName             ST
	SourceLocationIdentifier          EI
	SourceLocationName                ST
	ItemStatus                        CWE
	BinLocationIdentifier             []EI
	OrderPackaging                    CWE
	IssuePackaging                    CWE
	DefaultInventoryAssetAccount      EI
	PatientChargeableIndicator        CNE
	TransactionCode                   CWE
	TransactionAmountUnit             CP
	ItemImportanceCode                CWE
	StockedItemIndicator              CNE
	ConsignmentItemIndicator          CNE
	ReusableItemIndicator             CNE
	ReusableCost                      CP
	SubstituteItemIdentifier          []EI
	LatexFreeSubstituteItemIdentifier EI
	RecommendedReorderTheory          CWE
	RecommendedSafetyStockDays        NM
	RecommendedMaximumDaysInventory   NM
	RecommendedOrderPoint             NM
	RecommendedOrderAmount            NM
	OperatingRoomParLevelIndicator    CNE
}

// NewIVT creates an implementation of IVT
func NewIVT(input hl7.HL7Type) IVT {
	v := IVT{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.InventoryLocationIdentifier = NewEI(input.Index(2).Index(0))
	v.InventoryLocationName = NewST(input.Index(3).Index(0))
	v.SourceLocationIdentifier = NewEI(input.Index(4).Index(0))
	v.SourceLocationName = NewST(input.Index(5).Index(0))
	v.ItemStatus = NewCWE(input.Index(6).Index(0))
	v.BinLocationIdentifier = NewEISlice(input.Index(7))
	v.OrderPackaging = NewCWE(input.Index(8).Index(0))
	v.IssuePackaging = NewCWE(input.Index(9).Index(0))
	v.DefaultInventoryAssetAccount = NewEI(input.Index(10).Index(0))
	v.PatientChargeableIndicator = NewCNE(input.Index(11).Index(0))
	v.TransactionCode = NewCWE(input.Index(12).Index(0))
	v.TransactionAmountUnit = NewCP(input.Index(13).Index(0))
	v.ItemImportanceCode = NewCWE(input.Index(14).Index(0))
	v.StockedItemIndicator = NewCNE(input.Index(15).Index(0))
	v.ConsignmentItemIndicator = NewCNE(input.Index(16).Index(0))
	v.ReusableItemIndicator = NewCNE(input.Index(17).Index(0))
	v.ReusableCost = NewCP(input.Index(18).Index(0))
	v.SubstituteItemIdentifier = NewEISlice(input.Index(19))
	v.LatexFreeSubstituteItemIdentifier = NewEI(input.Index(20).Index(0))
	v.RecommendedReorderTheory = NewCWE(input.Index(21).Index(0))
	v.RecommendedSafetyStockDays = NewNM(input.Index(22).Index(0))
	v.RecommendedMaximumDaysInventory = NewNM(input.Index(23).Index(0))
	v.RecommendedOrderPoint = NewNM(input.Index(24).Index(0))
	v.RecommendedOrderAmount = NewNM(input.Index(25).Index(0))
	v.OperatingRoomParLevelIndicator = NewCNE(input.Index(26).Index(0))
	return v
}

// NewIVTSlice will construct a slice of type IVT
func NewIVTSlice(input hl7.HL7Type) []IVT {
	values := make([]IVT, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewIVT(input.Index(i))
	}
	return values
}

type RXD struct {
	DispenseSubIDCounter                                    NM
	DispenseGiveCode                                        CWE
	DateTimeDispensed                                       DTM
	ActualDispenseAmount                                    NM
	ActualDispenseUnits                                     CWE
	ActualDosageForm                                        CWE
	PrescriptionNumber                                      ST
	NumberOfRefillsRemaining                                NM
	DispenseNotes                                           []ST
	DispensingProvider                                      []XCN
	SubstitutionStatus                                      ID
	TotalDailyDose                                          CQ
	DispenseToLocation                                      ST
	NeedsHumanReview                                        ID
	PharmacyTreatmentSuppliersSpecialDispensingInstructions []CWE
	ActualStrength                                          NM
	ActualStrengthUnit                                      CWE
	SubstanceLotNumber                                      []ST
	SubstanceExpirationDate                                 []DTM
	SubstanceManufacturerName                               []CWE
	Indication                                              []CWE
	DispensePackageSize                                     NM
	DispensePackageSizeUnit                                 CWE
	DispensePackageMethod                                   ID
	SupplementaryCode                                       []CWE
	InitiatingLocation                                      CWE
	PackagingAssemblyLocation                               CWE
	ActualDrugStrengthVolume                                NM
	ActualDrugStrengthVolumeUnits                           CWE
	DispenseToPharmacy                                      CWE
	DispenseToPharmacyAddress                               XAD
	PharmacyOrderType                                       ID
	DispenseType                                            CWE
	PharmacyPhoneNumber                                     []XTN
	DispenseTagIdentifier                                   []EI
}

// NewRXD creates an implementation of RXD
func NewRXD(input hl7.HL7Type) RXD {
	v := RXD{}
	v.DispenseSubIDCounter = NewNM(input.Index(1).Index(0))
	v.DispenseGiveCode = NewCWE(input.Index(2).Index(0))
	v.DateTimeDispensed = NewDTM(input.Index(3).Index(0))
	v.ActualDispenseAmount = NewNM(input.Index(4).Index(0))
	v.ActualDispenseUnits = NewCWE(input.Index(5).Index(0))
	v.ActualDosageForm = NewCWE(input.Index(6).Index(0))
	v.PrescriptionNumber = NewST(input.Index(7).Index(0))
	v.NumberOfRefillsRemaining = NewNM(input.Index(8).Index(0))
	v.DispenseNotes = NewSTSlice(input.Index(9))
	v.DispensingProvider = NewXCNSlice(input.Index(10))
	v.SubstitutionStatus = NewID(input.Index(11).Index(0))
	v.TotalDailyDose = NewCQ(input.Index(12).Index(0))
	v.DispenseToLocation = NewST(input.Index(13).Index(0))
	v.NeedsHumanReview = NewID(input.Index(14).Index(0))
	v.PharmacyTreatmentSuppliersSpecialDispensingInstructions = NewCWESlice(input.Index(15))
	v.ActualStrength = NewNM(input.Index(16).Index(0))
	v.ActualStrengthUnit = NewCWE(input.Index(17).Index(0))
	v.SubstanceLotNumber = NewSTSlice(input.Index(18))
	v.SubstanceExpirationDate = NewDTMSlice(input.Index(19))
	v.SubstanceManufacturerName = NewCWESlice(input.Index(20))
	v.Indication = NewCWESlice(input.Index(21))
	v.DispensePackageSize = NewNM(input.Index(22).Index(0))
	v.DispensePackageSizeUnit = NewCWE(input.Index(23).Index(0))
	v.DispensePackageMethod = NewID(input.Index(24).Index(0))
	v.SupplementaryCode = NewCWESlice(input.Index(25))
	v.InitiatingLocation = NewCWE(input.Index(26).Index(0))
	v.PackagingAssemblyLocation = NewCWE(input.Index(27).Index(0))
	v.ActualDrugStrengthVolume = NewNM(input.Index(28).Index(0))
	v.ActualDrugStrengthVolumeUnits = NewCWE(input.Index(29).Index(0))
	v.DispenseToPharmacy = NewCWE(input.Index(30).Index(0))
	v.DispenseToPharmacyAddress = NewXAD(input.Index(31).Index(0))
	v.PharmacyOrderType = NewID(input.Index(32).Index(0))
	v.DispenseType = NewCWE(input.Index(33).Index(0))
	v.PharmacyPhoneNumber = NewXTNSlice(input.Index(34))
	v.DispenseTagIdentifier = NewEISlice(input.Index(35))
	return v
}

// NewRXDSlice will construct a slice of type RXD
func NewRXDSlice(input hl7.HL7Type) []RXD {
	values := make([]RXD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRXD(input.Index(i))
	}
	return values
}

type RDT struct {
	ColumnValue Varies
}

// NewRDT creates an implementation of RDT
func NewRDT(input hl7.HL7Type) RDT {
	v := RDT{}
	v.ColumnValue = NewVaries(input.Index(1).Index(0))
	return v
}

// NewRDTSlice will construct a slice of type RDT
func NewRDTSlice(input hl7.HL7Type) []RDT {
	values := make([]RDT, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRDT(input.Index(i))
	}
	return values
}

type PID struct {
	SetID                               SI
	PatientID                           ST
	PatientIdentifierList               []CX
	AlternatePatientIDPid               ST
	PatientName                         []XPN
	MothersMaidenName                   []XPN
	DateTimeOfBirth                     DTM
	AdministrativeSex                   CWE
	PatientAlias                        ST
	Race                                []CWE
	PatientAddress                      []XAD
	CountyCode                          ST
	PhoneNumberHome                     []XTN
	PhoneNumberBusiness                 []XTN
	PrimaryLanguage                     CWE
	MaritalStatus                       CWE
	Religion                            CWE
	PatientAccountNumber                CX
	SsnNumberPatient                    ST
	DriversLicenseNumberPatient         ST
	MothersIdentifier                   []CX
	EthnicGroup                         []CWE
	BirthPlace                          ST
	MultipleBirthIndicator              ID
	BirthOrder                          NM
	Citizenship                         []CWE
	VeteransMilitaryStatus              CWE
	Nationality                         ST
	PatientDeathDateAndTime             DTM
	PatientDeathIndicator               ID
	IdentityUnknownIndicator            ID
	IdentityReliabilityCode             []CWE
	LastUpdateDateTime                  DTM
	LastUpdateFacility                  HD
	TaxonomicClassificationCode         CWE
	BreedCode                           CWE
	Strain                              ST
	ProductionClassCode                 []CWE
	TribalCitizenship                   []CWE
	PatientTelecommunicationInformation []XTN
}

// NewPID creates an implementation of PID
func NewPID(input hl7.HL7Type) PID {
	v := PID{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.PatientID = NewST(input.Index(2).Index(0))
	v.PatientIdentifierList = NewCXSlice(input.Index(3))
	v.AlternatePatientIDPid = NewST(input.Index(4).Index(0))
	v.PatientName = NewXPNSlice(input.Index(5))
	v.MothersMaidenName = NewXPNSlice(input.Index(6))
	v.DateTimeOfBirth = NewDTM(input.Index(7).Index(0))
	v.AdministrativeSex = NewCWE(input.Index(8).Index(0))
	v.PatientAlias = NewST(input.Index(9).Index(0))
	v.Race = NewCWESlice(input.Index(10))
	v.PatientAddress = NewXADSlice(input.Index(11))
	v.CountyCode = NewST(input.Index(12).Index(0))
	v.PhoneNumberHome = NewXTNSlice(input.Index(13))
	v.PhoneNumberBusiness = NewXTNSlice(input.Index(14))
	v.PrimaryLanguage = NewCWE(input.Index(15).Index(0))
	v.MaritalStatus = NewCWE(input.Index(16).Index(0))
	v.Religion = NewCWE(input.Index(17).Index(0))
	v.PatientAccountNumber = NewCX(input.Index(18).Index(0))
	v.SsnNumberPatient = NewST(input.Index(19).Index(0))
	v.DriversLicenseNumberPatient = NewST(input.Index(20).Index(0))
	v.MothersIdentifier = NewCXSlice(input.Index(21))
	v.EthnicGroup = NewCWESlice(input.Index(22))
	v.BirthPlace = NewST(input.Index(23).Index(0))
	v.MultipleBirthIndicator = NewID(input.Index(24).Index(0))
	v.BirthOrder = NewNM(input.Index(25).Index(0))
	v.Citizenship = NewCWESlice(input.Index(26))
	v.VeteransMilitaryStatus = NewCWE(input.Index(27).Index(0))
	v.Nationality = NewST(input.Index(28).Index(0))
	v.PatientDeathDateAndTime = NewDTM(input.Index(29).Index(0))
	v.PatientDeathIndicator = NewID(input.Index(30).Index(0))
	v.IdentityUnknownIndicator = NewID(input.Index(31).Index(0))
	v.IdentityReliabilityCode = NewCWESlice(input.Index(32))
	v.LastUpdateDateTime = NewDTM(input.Index(33).Index(0))
	v.LastUpdateFacility = NewHD(input.Index(34).Index(0))
	v.TaxonomicClassificationCode = NewCWE(input.Index(35).Index(0))
	v.BreedCode = NewCWE(input.Index(36).Index(0))
	v.Strain = NewST(input.Index(37).Index(0))
	v.ProductionClassCode = NewCWESlice(input.Index(38))
	v.TribalCitizenship = NewCWESlice(input.Index(39))
	v.PatientTelecommunicationInformation = NewXTNSlice(input.Index(40))
	return v
}

// NewPIDSlice will construct a slice of type PID
func NewPIDSlice(input hl7.HL7Type) []PID {
	values := make([]PID, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPID(input.Index(i))
	}
	return values
}

type ROL struct {
	RoleInstanceID              EI
	ActionCode                  ID
	RoleRol                     CWE
	RolePerson                  []XCN
	RoleBeginDateTime           DTM
	RoleEndDateTime             DTM
	RoleDuration                CWE
	RoleActionReason            CWE
	ProviderType                []CWE
	OrganizationUnitType        CWE
	OfficeHomeAddressBirthplace []XAD
	Phone                       []XTN
	PersonsLocation             PL
	Organization                XON
}

// NewROL creates an implementation of ROL
func NewROL(input hl7.HL7Type) ROL {
	v := ROL{}
	v.RoleInstanceID = NewEI(input.Index(1).Index(0))
	v.ActionCode = NewID(input.Index(2).Index(0))
	v.RoleRol = NewCWE(input.Index(3).Index(0))
	v.RolePerson = NewXCNSlice(input.Index(4))
	v.RoleBeginDateTime = NewDTM(input.Index(5).Index(0))
	v.RoleEndDateTime = NewDTM(input.Index(6).Index(0))
	v.RoleDuration = NewCWE(input.Index(7).Index(0))
	v.RoleActionReason = NewCWE(input.Index(8).Index(0))
	v.ProviderType = NewCWESlice(input.Index(9))
	v.OrganizationUnitType = NewCWE(input.Index(10).Index(0))
	v.OfficeHomeAddressBirthplace = NewXADSlice(input.Index(11))
	v.Phone = NewXTNSlice(input.Index(12))
	v.PersonsLocation = NewPL(input.Index(13).Index(0))
	v.Organization = NewXON(input.Index(14).Index(0))
	return v
}

// NewROLSlice will construct a slice of type ROL
func NewROLSlice(input hl7.HL7Type) []ROL {
	values := make([]ROL, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewROL(input.Index(i))
	}
	return values
}

type LAN struct {
	SetID                   SI
	LanguageCode            CWE
	LanguageAbilityCode     []CWE
	LanguageProficiencyCode CWE
}

// NewLAN creates an implementation of LAN
func NewLAN(input hl7.HL7Type) LAN {
	v := LAN{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.LanguageCode = NewCWE(input.Index(2).Index(0))
	v.LanguageAbilityCode = NewCWESlice(input.Index(3))
	v.LanguageProficiencyCode = NewCWE(input.Index(4).Index(0))
	return v
}

// NewLANSlice will construct a slice of type LAN
func NewLANSlice(input hl7.HL7Type) []LAN {
	values := make([]LAN, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewLAN(input.Index(i))
	}
	return values
}

type RQ1 struct {
	AnticipatedPrice       ST
	ManufacturerIdentifier CWE
	ManufacturersCatalog   ST
	VendorID               CWE
	VendorCatalog          ST
	Taxable                ID
	SubstituteAllowed      ID
}

// NewRQ1 creates an implementation of RQ1
func NewRQ1(input hl7.HL7Type) RQ1 {
	v := RQ1{}
	v.AnticipatedPrice = NewST(input.Index(1).Index(0))
	v.ManufacturerIdentifier = NewCWE(input.Index(2).Index(0))
	v.ManufacturersCatalog = NewST(input.Index(3).Index(0))
	v.VendorID = NewCWE(input.Index(4).Index(0))
	v.VendorCatalog = NewST(input.Index(5).Index(0))
	v.Taxable = NewID(input.Index(6).Index(0))
	v.SubstituteAllowed = NewID(input.Index(7).Index(0))
	return v
}

// NewRQ1Slice will construct a slice of type RQ1
func NewRQ1Slice(input hl7.HL7Type) []RQ1 {
	values := make([]RQ1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRQ1(input.Index(i))
	}
	return values
}

type RDF struct {
	NumberOfColumnsPerRow NM
	ColumnDescription     []RCD
}

// NewRDF creates an implementation of RDF
func NewRDF(input hl7.HL7Type) RDF {
	v := RDF{}
	v.NumberOfColumnsPerRow = NewNM(input.Index(1).Index(0))
	v.ColumnDescription = NewRCDSlice(input.Index(2))
	return v
}

// NewRDFSlice will construct a slice of type RDF
func NewRDFSlice(input hl7.HL7Type) []RDF {
	values := make([]RDF, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRDF(input.Index(i))
	}
	return values
}

type AIL struct {
	SetID                    SI
	SegmentActionCode        ID
	LocationResourceID       []PL
	LocationTypeAil          CWE
	LocationGroup            CWE
	StartDateTime            DTM
	StartDateTimeOffset      NM
	StartDateTimeOffsetUnits CNE
	Duration                 NM
	DurationUnits            CNE
	AllowSubstitutionCode    CWE
	FillerStatusCode         CWE
}

// NewAIL creates an implementation of AIL
func NewAIL(input hl7.HL7Type) AIL {
	v := AIL{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.SegmentActionCode = NewID(input.Index(2).Index(0))
	v.LocationResourceID = NewPLSlice(input.Index(3))
	v.LocationTypeAil = NewCWE(input.Index(4).Index(0))
	v.LocationGroup = NewCWE(input.Index(5).Index(0))
	v.StartDateTime = NewDTM(input.Index(6).Index(0))
	v.StartDateTimeOffset = NewNM(input.Index(7).Index(0))
	v.StartDateTimeOffsetUnits = NewCNE(input.Index(8).Index(0))
	v.Duration = NewNM(input.Index(9).Index(0))
	v.DurationUnits = NewCNE(input.Index(10).Index(0))
	v.AllowSubstitutionCode = NewCWE(input.Index(11).Index(0))
	v.FillerStatusCode = NewCWE(input.Index(12).Index(0))
	return v
}

// NewAILSlice will construct a slice of type AIL
func NewAILSlice(input hl7.HL7Type) []AIL {
	values := make([]AIL, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewAIL(input.Index(i))
	}
	return values
}

type TQ2 struct {
	SetID                             SI
	SequenceResultsFlag               ID
	RelatedPlacerNumber               []EI
	RelatedFillerNumber               []EI
	RelatedPlacerGroupNumber          []EI
	SequenceConditionCode             ID
	CyclicEntryExitIndicator          ID
	SequenceConditionTimeInterval     CQ
	CyclicGroupMaximumNumberOfRepeats NM
	SpecialServiceRequestRelationship ID
}

// NewTQ2 creates an implementation of TQ2
func NewTQ2(input hl7.HL7Type) TQ2 {
	v := TQ2{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.SequenceResultsFlag = NewID(input.Index(2).Index(0))
	v.RelatedPlacerNumber = NewEISlice(input.Index(3))
	v.RelatedFillerNumber = NewEISlice(input.Index(4))
	v.RelatedPlacerGroupNumber = NewEISlice(input.Index(5))
	v.SequenceConditionCode = NewID(input.Index(6).Index(0))
	v.CyclicEntryExitIndicator = NewID(input.Index(7).Index(0))
	v.SequenceConditionTimeInterval = NewCQ(input.Index(8).Index(0))
	v.CyclicGroupMaximumNumberOfRepeats = NewNM(input.Index(9).Index(0))
	v.SpecialServiceRequestRelationship = NewID(input.Index(10).Index(0))
	return v
}

// NewTQ2Slice will construct a slice of type TQ2
func NewTQ2Slice(input hl7.HL7Type) []TQ2 {
	values := make([]TQ2, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewTQ2(input.Index(i))
	}
	return values
}

type ITM struct {
	ItemIdentifier                                     EI
	ItemDescription                                    ST
	ItemStatus                                         CWE
	ItemType                                           CWE
	ItemCategory                                       CWE
	SubjectToExpirationIndicator                       CNE
	ManufacturerIdentifier                             EI
	ManufacturerName                                   ST
	ManufacturerCatalogNumber                          ST
	ManufacturerLabelerIdentificationCode              CWE
	PatientChargeableIndicator                         CNE
	TransactionCode                                    CWE
	TransactionAmountUnit                              CP
	StockedItemIndicator                               CNE
	SupplyRiskCodes                                    CWE
	ApprovingRegulatoryAgency                          []XON
	LatexIndicator                                     CNE
	RulingAct                                          []CWE
	ItemNaturalAccountCode                             CWE
	ApprovedToBuyQuantity                              NM
	ApprovedToBuyPrice                                 MO
	TaxableItemIndicator                               CNE
	FreightChargeIndicator                             CNE
	ItemSetIndicator                                   CNE
	ItemSetIdentifier                                  EI
	TrackDepartmentUsageIndicator                      CNE
	ProcedureCode                                      CNE
	ProcedureCodeModifier                              []CNE
	SpecialHandlingCode                                CWE
	HazardousIndicator                                 CNE
	SterileIndicator                                   CNE
	MaterialDataSafetySheetNumber                      EI
	UnitedNationsStandardProductsAndServicesCodeUNSPSC CWE
}

// NewITM creates an implementation of ITM
func NewITM(input hl7.HL7Type) ITM {
	v := ITM{}
	v.ItemIdentifier = NewEI(input.Index(1).Index(0))
	v.ItemDescription = NewST(input.Index(2).Index(0))
	v.ItemStatus = NewCWE(input.Index(3).Index(0))
	v.ItemType = NewCWE(input.Index(4).Index(0))
	v.ItemCategory = NewCWE(input.Index(5).Index(0))
	v.SubjectToExpirationIndicator = NewCNE(input.Index(6).Index(0))
	v.ManufacturerIdentifier = NewEI(input.Index(7).Index(0))
	v.ManufacturerName = NewST(input.Index(8).Index(0))
	v.ManufacturerCatalogNumber = NewST(input.Index(9).Index(0))
	v.ManufacturerLabelerIdentificationCode = NewCWE(input.Index(10).Index(0))
	v.PatientChargeableIndicator = NewCNE(input.Index(11).Index(0))
	v.TransactionCode = NewCWE(input.Index(12).Index(0))
	v.TransactionAmountUnit = NewCP(input.Index(13).Index(0))
	v.StockedItemIndicator = NewCNE(input.Index(14).Index(0))
	v.SupplyRiskCodes = NewCWE(input.Index(15).Index(0))
	v.ApprovingRegulatoryAgency = NewXONSlice(input.Index(16))
	v.LatexIndicator = NewCNE(input.Index(17).Index(0))
	v.RulingAct = NewCWESlice(input.Index(18))
	v.ItemNaturalAccountCode = NewCWE(input.Index(19).Index(0))
	v.ApprovedToBuyQuantity = NewNM(input.Index(20).Index(0))
	v.ApprovedToBuyPrice = NewMO(input.Index(21).Index(0))
	v.TaxableItemIndicator = NewCNE(input.Index(22).Index(0))
	v.FreightChargeIndicator = NewCNE(input.Index(23).Index(0))
	v.ItemSetIndicator = NewCNE(input.Index(24).Index(0))
	v.ItemSetIdentifier = NewEI(input.Index(25).Index(0))
	v.TrackDepartmentUsageIndicator = NewCNE(input.Index(26).Index(0))
	v.ProcedureCode = NewCNE(input.Index(27).Index(0))
	v.ProcedureCodeModifier = NewCNESlice(input.Index(28))
	v.SpecialHandlingCode = NewCWE(input.Index(29).Index(0))
	v.HazardousIndicator = NewCNE(input.Index(30).Index(0))
	v.SterileIndicator = NewCNE(input.Index(31).Index(0))
	v.MaterialDataSafetySheetNumber = NewEI(input.Index(32).Index(0))
	v.UnitedNationsStandardProductsAndServicesCodeUNSPSC = NewCWE(input.Index(33).Index(0))
	return v
}

// NewITMSlice will construct a slice of type ITM
func NewITMSlice(input hl7.HL7Type) []ITM {
	values := make([]ITM, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewITM(input.Index(i))
	}
	return values
}

type EVN struct {
	EventTypeCode        ST
	RecordedDateTime     DTM
	DateTimePlannedEvent DTM
	EventReasonCode      CWE
	OperatorID           []XCN
	EventOccurred        DTM
	EventFacility        HD
}

// NewEVN creates an implementation of EVN
func NewEVN(input hl7.HL7Type) EVN {
	v := EVN{}
	v.EventTypeCode = NewST(input.Index(1).Index(0))
	v.RecordedDateTime = NewDTM(input.Index(2).Index(0))
	v.DateTimePlannedEvent = NewDTM(input.Index(3).Index(0))
	v.EventReasonCode = NewCWE(input.Index(4).Index(0))
	v.OperatorID = NewXCNSlice(input.Index(5))
	v.EventOccurred = NewDTM(input.Index(6).Index(0))
	v.EventFacility = NewHD(input.Index(7).Index(0))
	return v
}

// NewEVNSlice will construct a slice of type EVN
func NewEVNSlice(input hl7.HL7Type) []EVN {
	values := make([]EVN, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEVN(input.Index(i))
	}
	return values
}

type MFE struct {
	RecordLevelEventCode ID
	MfnControlID         ST
	EffectiveDateTime    DTM
	PrimaryKeyValueMfe   []Varies
	PrimaryKeyValueType  []ID
	EnteredDateTime      DTM
	EnteredBy            XCN
}

// NewMFE creates an implementation of MFE
func NewMFE(input hl7.HL7Type) MFE {
	v := MFE{}
	v.RecordLevelEventCode = NewID(input.Index(1).Index(0))
	v.MfnControlID = NewST(input.Index(2).Index(0))
	v.EffectiveDateTime = NewDTM(input.Index(3).Index(0))
	v.PrimaryKeyValueMfe = NewVariesSlice(input.Index(4))
	v.PrimaryKeyValueType = NewIDSlice(input.Index(5))
	v.EnteredDateTime = NewDTM(input.Index(6).Index(0))
	v.EnteredBy = NewXCN(input.Index(7).Index(0))
	return v
}

// NewMFESlice will construct a slice of type MFE
func NewMFESlice(input hl7.HL7Type) []MFE {
	values := make([]MFE, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFE(input.Index(i))
	}
	return values
}

type IAR struct {
	AllergyReactionCode             CWE
	AllergySeverityCode             CWE
	SensitivityToCausativeAgentCode CWE
	Management                      ST
}

// NewIAR creates an implementation of IAR
func NewIAR(input hl7.HL7Type) IAR {
	v := IAR{}
	v.AllergyReactionCode = NewCWE(input.Index(1).Index(0))
	v.AllergySeverityCode = NewCWE(input.Index(2).Index(0))
	v.SensitivityToCausativeAgentCode = NewCWE(input.Index(3).Index(0))
	v.Management = NewST(input.Index(4).Index(0))
	return v
}

// NewIARSlice will construct a slice of type IAR
func NewIARSlice(input hl7.HL7Type) []IAR {
	values := make([]IAR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewIAR(input.Index(i))
	}
	return values
}

type PRT struct {
	ParticipationInstanceID               EI
	ActionCode                            ID
	ActionReason                          CWE
	Participation                         CWE
	ParticipationPerson                   []XCN
	ParticipationPersonProviderType       CWE
	ParticipantOrganizationUnitType       CWE
	ParticipationOrganization             []XON
	ParticipantLocation                   []PL
	ParticipationDevice                   []EI
	ParticipationBeginDateTimearrivalTime DTM
	ParticipationEndDateTimedepartureTime DTM
	ParticipationQualitativeDuration      CWE
	ParticipationAddress                  []XAD
	ParticipantTelecommunicationAddress   []XTN
}

// NewPRT creates an implementation of PRT
func NewPRT(input hl7.HL7Type) PRT {
	v := PRT{}
	v.ParticipationInstanceID = NewEI(input.Index(1).Index(0))
	v.ActionCode = NewID(input.Index(2).Index(0))
	v.ActionReason = NewCWE(input.Index(3).Index(0))
	v.Participation = NewCWE(input.Index(4).Index(0))
	v.ParticipationPerson = NewXCNSlice(input.Index(5))
	v.ParticipationPersonProviderType = NewCWE(input.Index(6).Index(0))
	v.ParticipantOrganizationUnitType = NewCWE(input.Index(7).Index(0))
	v.ParticipationOrganization = NewXONSlice(input.Index(8))
	v.ParticipantLocation = NewPLSlice(input.Index(9))
	v.ParticipationDevice = NewEISlice(input.Index(10))
	v.ParticipationBeginDateTimearrivalTime = NewDTM(input.Index(11).Index(0))
	v.ParticipationEndDateTimedepartureTime = NewDTM(input.Index(12).Index(0))
	v.ParticipationQualitativeDuration = NewCWE(input.Index(13).Index(0))
	v.ParticipationAddress = NewXADSlice(input.Index(14))
	v.ParticipantTelecommunicationAddress = NewXTNSlice(input.Index(15))
	return v
}

// NewPRTSlice will construct a slice of type PRT
func NewPRTSlice(input hl7.HL7Type) []PRT {
	values := make([]PRT, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPRT(input.Index(i))
	}
	return values
}

type LDP struct {
	PrimaryKeyValueLdp  PL
	LocationDepartment  CWE
	LocationService     []CWE
	SpecialtyType       []CWE
	ValidPatientClasses []CWE
	ActiveInactiveFlag  ID
	ActivationDateLdp   DTM
	InactivationDateLdp DTM
	InactivatedReason   ST
	VisitingHours       []VH
	ContactPhone        XTN
	LocationCostCenter  CWE
}

// NewLDP creates an implementation of LDP
func NewLDP(input hl7.HL7Type) LDP {
	v := LDP{}
	v.PrimaryKeyValueLdp = NewPL(input.Index(1).Index(0))
	v.LocationDepartment = NewCWE(input.Index(2).Index(0))
	v.LocationService = NewCWESlice(input.Index(3))
	v.SpecialtyType = NewCWESlice(input.Index(4))
	v.ValidPatientClasses = NewCWESlice(input.Index(5))
	v.ActiveInactiveFlag = NewID(input.Index(6).Index(0))
	v.ActivationDateLdp = NewDTM(input.Index(7).Index(0))
	v.InactivationDateLdp = NewDTM(input.Index(8).Index(0))
	v.InactivatedReason = NewST(input.Index(9).Index(0))
	v.VisitingHours = NewVHSlice(input.Index(10))
	v.ContactPhone = NewXTN(input.Index(11).Index(0))
	v.LocationCostCenter = NewCWE(input.Index(12).Index(0))
	return v
}

// NewLDPSlice will construct a slice of type LDP
func NewLDPSlice(input hl7.HL7Type) []LDP {
	values := make([]LDP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewLDP(input.Index(i))
	}
	return values
}

type CSS struct {
	StudyScheduledTimePoint        CWE
	StudyScheduledPatientTimePoint DTM
	StudyQualityControlCodes       []CWE
}

// NewCSS creates an implementation of CSS
func NewCSS(input hl7.HL7Type) CSS {
	v := CSS{}
	v.StudyScheduledTimePoint = NewCWE(input.Index(1).Index(0))
	v.StudyScheduledPatientTimePoint = NewDTM(input.Index(2).Index(0))
	v.StudyQualityControlCodes = NewCWESlice(input.Index(3))
	return v
}

// NewCSSSlice will construct a slice of type CSS
func NewCSSSlice(input hl7.HL7Type) []CSS {
	values := make([]CSS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCSS(input.Index(i))
	}
	return values
}

type ABS struct {
	DischargeCareProvider      XCN
	TransferMedicalServiceCode CWE
	SeverityOfIllnessCode      CWE
	DateTimeOfAttestation      DTM
	AttestedBy                 XCN
	TriageCode                 CWE
	AbstractCompletionDateTime DTM
	AbstractedBy               XCN
	CaseCategoryCode           CWE
	CaesarianSectionIndicator  ID
	GestationCategoryCode      CWE
	GestationPeriodWeeks       NM
	NewbornCode                CWE
	StillbornIndicator         ID
}

// NewABS creates an implementation of ABS
func NewABS(input hl7.HL7Type) ABS {
	v := ABS{}
	v.DischargeCareProvider = NewXCN(input.Index(1).Index(0))
	v.TransferMedicalServiceCode = NewCWE(input.Index(2).Index(0))
	v.SeverityOfIllnessCode = NewCWE(input.Index(3).Index(0))
	v.DateTimeOfAttestation = NewDTM(input.Index(4).Index(0))
	v.AttestedBy = NewXCN(input.Index(5).Index(0))
	v.TriageCode = NewCWE(input.Index(6).Index(0))
	v.AbstractCompletionDateTime = NewDTM(input.Index(7).Index(0))
	v.AbstractedBy = NewXCN(input.Index(8).Index(0))
	v.CaseCategoryCode = NewCWE(input.Index(9).Index(0))
	v.CaesarianSectionIndicator = NewID(input.Index(10).Index(0))
	v.GestationCategoryCode = NewCWE(input.Index(11).Index(0))
	v.GestationPeriodWeeks = NewNM(input.Index(12).Index(0))
	v.NewbornCode = NewCWE(input.Index(13).Index(0))
	v.StillbornIndicator = NewID(input.Index(14).Index(0))
	return v
}

// NewABSSlice will construct a slice of type ABS
func NewABSSlice(input hl7.HL7Type) []ABS {
	values := make([]ABS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewABS(input.Index(i))
	}
	return values
}

type QAK struct {
	QueryTag            ST
	QueryResponseStatus ID
	MessageQueryName    CWE
	HitCountTotal       NM
	ThisPayload         NM
	HitsRemaining       NM
}

// NewQAK creates an implementation of QAK
func NewQAK(input hl7.HL7Type) QAK {
	v := QAK{}
	v.QueryTag = NewST(input.Index(1).Index(0))
	v.QueryResponseStatus = NewID(input.Index(2).Index(0))
	v.MessageQueryName = NewCWE(input.Index(3).Index(0))
	v.HitCountTotal = NewNM(input.Index(4).Index(0))
	v.ThisPayload = NewNM(input.Index(5).Index(0))
	v.HitsRemaining = NewNM(input.Index(6).Index(0))
	return v
}

// NewQAKSlice will construct a slice of type QAK
func NewQAKSlice(input hl7.HL7Type) []QAK {
	values := make([]QAK, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQAK(input.Index(i))
	}
	return values
}

type MSH struct {
	FieldSeparator                      ST
	EncodingCharacters                  ST
	SendingApplication                  HD
	SendingFacility                     HD
	ReceivingApplication                HD
	ReceivingFacility                   HD
	DateTimeOfMessage                   DTM
	Security                            ST
	MessageType                         MSG
	MessageControlID                    ST
	ProcessingID                        PT
	VersionID                           VID
	SequenceNumber                      NM
	ContinuationPointer                 ST
	AcceptAcknowledgmentType            ID
	ApplicationAcknowledgmentType       ID
	CountryCode                         ID
	CharacterSet                        []ID
	PrincipalLanguageOfMessage          CWE
	AlternateCharacterSetHandlingScheme ID
	MessageProfileIdentifier            []EI
	SendingResponsibleOrganization      XON
	ReceivingResponsibleOrganization    XON
	SendingNetworkAddress               HD
	ReceivingNetworkAddress             HD
}

// NewMSH creates an implementation of MSH
func NewMSH(input hl7.HL7Type) MSH {
	v := MSH{}
	v.FieldSeparator = NewST(input.Index(1).Index(0))
	v.EncodingCharacters = NewST(input.Index(2).Index(0))
	v.SendingApplication = NewHD(input.Index(3).Index(0))
	v.SendingFacility = NewHD(input.Index(4).Index(0))
	v.ReceivingApplication = NewHD(input.Index(5).Index(0))
	v.ReceivingFacility = NewHD(input.Index(6).Index(0))
	v.DateTimeOfMessage = NewDTM(input.Index(7).Index(0))
	v.Security = NewST(input.Index(8).Index(0))
	v.MessageType = NewMSG(input.Index(9).Index(0))
	v.MessageControlID = NewST(input.Index(10).Index(0))
	v.ProcessingID = NewPT(input.Index(11).Index(0))
	v.VersionID = NewVID(input.Index(12).Index(0))
	v.SequenceNumber = NewNM(input.Index(13).Index(0))
	v.ContinuationPointer = NewST(input.Index(14).Index(0))
	v.AcceptAcknowledgmentType = NewID(input.Index(15).Index(0))
	v.ApplicationAcknowledgmentType = NewID(input.Index(16).Index(0))
	v.CountryCode = NewID(input.Index(17).Index(0))
	v.CharacterSet = NewIDSlice(input.Index(18))
	v.PrincipalLanguageOfMessage = NewCWE(input.Index(19).Index(0))
	v.AlternateCharacterSetHandlingScheme = NewID(input.Index(20).Index(0))
	v.MessageProfileIdentifier = NewEISlice(input.Index(21))
	v.SendingResponsibleOrganization = NewXON(input.Index(22).Index(0))
	v.ReceivingResponsibleOrganization = NewXON(input.Index(23).Index(0))
	v.SendingNetworkAddress = NewHD(input.Index(24).Index(0))
	v.ReceivingNetworkAddress = NewHD(input.Index(25).Index(0))
	return v
}

// NewMSHSlice will construct a slice of type MSH
func NewMSHSlice(input hl7.HL7Type) []MSH {
	values := make([]MSH, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMSH(input.Index(i))
	}
	return values
}

type GT1 struct {
	SetID                              SI
	GuarantorNumber                    []CX
	GuarantorName                      []XPN
	GuarantorSpouseName                []XPN
	GuarantorAddress                   []XAD
	GuarantorPhNumHome                 []XTN
	GuarantorPhNumBusiness             []XTN
	GuarantorDateTimeOfBirth           DTM
	GuarantorAdministrativeSex         CWE
	GuarantorType                      CWE
	GuarantorRelationship              CWE
	GuarantorSsn                       ST
	GuarantorDateBegin                 DT
	GuarantorDateEnd                   DT
	GuarantorPriority                  NM
	GuarantorEmployerName              []XPN
	GuarantorEmployerAddress           []XAD
	GuarantorEmployerPhoneNumber       []XTN
	GuarantorEmployeeIDNumber          []CX
	GuarantorEmploymentStatus          CWE
	GuarantorOrganizationName          []XON
	GuarantorBillingHoldFlag           ID
	GuarantorCreditRatingCode          CWE
	GuarantorDeathDateAndTime          DTM
	GuarantorDeathFlag                 ID
	GuarantorChargeAdjustmentCode      CWE
	GuarantorHouseholdAnnualIncome     CP
	GuarantorHouseholdSize             NM
	GuarantorEmployerIDNumber          []CX
	GuarantorMaritalStatusCode         CWE
	GuarantorHireEffectiveDate         DT
	EmploymentStopDate                 DT
	LivingDependency                   CWE
	AmbulatoryStatus                   []CWE
	Citizenship                        []CWE
	PrimaryLanguage                    CWE
	LivingArrangement                  CWE
	PublicityCode                      CWE
	ProtectionIndicator                ID
	StudentIndicator                   CWE
	Religion                           CWE
	MothersMaidenName                  []XPN
	Nationality                        CWE
	EthnicGroup                        []CWE
	ContactPersonsName                 []XPN
	ContactPersonsTelephoneNumber      []XTN
	ContactReason                      CWE
	ContactRelationship                CWE
	JobTitle                           ST
	JobCodeClass                       JCC
	GuarantorEmployersOrganizationName []XON
	Handicap                           CWE
	JobStatus                          CWE
	GuarantorFinancialClass            FC
	GuarantorRace                      []CWE
	GuarantorBirthPlace                ST
	VipIndicator                       CWE
}

// NewGT1 creates an implementation of GT1
func NewGT1(input hl7.HL7Type) GT1 {
	v := GT1{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.GuarantorNumber = NewCXSlice(input.Index(2))
	v.GuarantorName = NewXPNSlice(input.Index(3))
	v.GuarantorSpouseName = NewXPNSlice(input.Index(4))
	v.GuarantorAddress = NewXADSlice(input.Index(5))
	v.GuarantorPhNumHome = NewXTNSlice(input.Index(6))
	v.GuarantorPhNumBusiness = NewXTNSlice(input.Index(7))
	v.GuarantorDateTimeOfBirth = NewDTM(input.Index(8).Index(0))
	v.GuarantorAdministrativeSex = NewCWE(input.Index(9).Index(0))
	v.GuarantorType = NewCWE(input.Index(10).Index(0))
	v.GuarantorRelationship = NewCWE(input.Index(11).Index(0))
	v.GuarantorSsn = NewST(input.Index(12).Index(0))
	v.GuarantorDateBegin = NewDT(input.Index(13).Index(0))
	v.GuarantorDateEnd = NewDT(input.Index(14).Index(0))
	v.GuarantorPriority = NewNM(input.Index(15).Index(0))
	v.GuarantorEmployerName = NewXPNSlice(input.Index(16))
	v.GuarantorEmployerAddress = NewXADSlice(input.Index(17))
	v.GuarantorEmployerPhoneNumber = NewXTNSlice(input.Index(18))
	v.GuarantorEmployeeIDNumber = NewCXSlice(input.Index(19))
	v.GuarantorEmploymentStatus = NewCWE(input.Index(20).Index(0))
	v.GuarantorOrganizationName = NewXONSlice(input.Index(21))
	v.GuarantorBillingHoldFlag = NewID(input.Index(22).Index(0))
	v.GuarantorCreditRatingCode = NewCWE(input.Index(23).Index(0))
	v.GuarantorDeathDateAndTime = NewDTM(input.Index(24).Index(0))
	v.GuarantorDeathFlag = NewID(input.Index(25).Index(0))
	v.GuarantorChargeAdjustmentCode = NewCWE(input.Index(26).Index(0))
	v.GuarantorHouseholdAnnualIncome = NewCP(input.Index(27).Index(0))
	v.GuarantorHouseholdSize = NewNM(input.Index(28).Index(0))
	v.GuarantorEmployerIDNumber = NewCXSlice(input.Index(29))
	v.GuarantorMaritalStatusCode = NewCWE(input.Index(30).Index(0))
	v.GuarantorHireEffectiveDate = NewDT(input.Index(31).Index(0))
	v.EmploymentStopDate = NewDT(input.Index(32).Index(0))
	v.LivingDependency = NewCWE(input.Index(33).Index(0))
	v.AmbulatoryStatus = NewCWESlice(input.Index(34))
	v.Citizenship = NewCWESlice(input.Index(35))
	v.PrimaryLanguage = NewCWE(input.Index(36).Index(0))
	v.LivingArrangement = NewCWE(input.Index(37).Index(0))
	v.PublicityCode = NewCWE(input.Index(38).Index(0))
	v.ProtectionIndicator = NewID(input.Index(39).Index(0))
	v.StudentIndicator = NewCWE(input.Index(40).Index(0))
	v.Religion = NewCWE(input.Index(41).Index(0))
	v.MothersMaidenName = NewXPNSlice(input.Index(42))
	v.Nationality = NewCWE(input.Index(43).Index(0))
	v.EthnicGroup = NewCWESlice(input.Index(44))
	v.ContactPersonsName = NewXPNSlice(input.Index(45))
	v.ContactPersonsTelephoneNumber = NewXTNSlice(input.Index(46))
	v.ContactReason = NewCWE(input.Index(47).Index(0))
	v.ContactRelationship = NewCWE(input.Index(48).Index(0))
	v.JobTitle = NewST(input.Index(49).Index(0))
	v.JobCodeClass = NewJCC(input.Index(50).Index(0))
	v.GuarantorEmployersOrganizationName = NewXONSlice(input.Index(51))
	v.Handicap = NewCWE(input.Index(52).Index(0))
	v.JobStatus = NewCWE(input.Index(53).Index(0))
	v.GuarantorFinancialClass = NewFC(input.Index(54).Index(0))
	v.GuarantorRace = NewCWESlice(input.Index(55))
	v.GuarantorBirthPlace = NewST(input.Index(56).Index(0))
	v.VipIndicator = NewCWE(input.Index(57).Index(0))
	return v
}

// NewGT1Slice will construct a slice of type GT1
func NewGT1Slice(input hl7.HL7Type) []GT1 {
	values := make([]GT1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewGT1(input.Index(i))
	}
	return values
}

type CSP struct {
	StudyPhaseIdentifier    CWE
	DateTimeStudyPhaseBegan DTM
	DateTimeStudyPhaseEnded DTM
	StudyPhaseEvaluability  CWE
}

// NewCSP creates an implementation of CSP
func NewCSP(input hl7.HL7Type) CSP {
	v := CSP{}
	v.StudyPhaseIdentifier = NewCWE(input.Index(1).Index(0))
	v.DateTimeStudyPhaseBegan = NewDTM(input.Index(2).Index(0))
	v.DateTimeStudyPhaseEnded = NewDTM(input.Index(3).Index(0))
	v.StudyPhaseEvaluability = NewCWE(input.Index(4).Index(0))
	return v
}

// NewCSPSlice will construct a slice of type CSP
func NewCSPSlice(input hl7.HL7Type) []CSP {
	values := make([]CSP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCSP(input.Index(i))
	}
	return values
}

type OM7 struct {
	SequenceNumberTestObservationMasterFile NM
	UniversalServiceIdentifier              CWE
	CategoryIdentifier                      []CWE
	CategoryDescription                     TX
	CategorySynonym                         []ST
	EffectiveTestServiceStartDateTime       DTM
	EffectiveTestServiceEndDateTime         DTM
	TestServiceDefaultDurationQuantity      NM
	TestServiceDefaultDurationUnits         CWE
	TestServiceDefaultFrequency             CWE
	ConsentIndicator                        ID
	ConsentIdentifier                       CWE
	ConsentEffectiveStartDateTime           DTM
	ConsentEffectiveEndDateTime             DTM
	ConsentIntervalQuantity                 NM
	ConsentIntervalUnits                    CWE
	ConsentWaitingPeriodQuantity            NM
	ConsentWaitingPeriodUnits               CWE
	EffectiveDateTimeOfChange               DTM
	EnteredBy                               XCN
	OrderableAtLocation                     []PL
	FormularyStatus                         CWE
	SpecialOrderIndicator                   ID
	PrimaryKeyValueCdm                      []CWE
}

// NewOM7 creates an implementation of OM7
func NewOM7(input hl7.HL7Type) OM7 {
	v := OM7{}
	v.SequenceNumberTestObservationMasterFile = NewNM(input.Index(1).Index(0))
	v.UniversalServiceIdentifier = NewCWE(input.Index(2).Index(0))
	v.CategoryIdentifier = NewCWESlice(input.Index(3))
	v.CategoryDescription = NewTX(input.Index(4).Index(0))
	v.CategorySynonym = NewSTSlice(input.Index(5))
	v.EffectiveTestServiceStartDateTime = NewDTM(input.Index(6).Index(0))
	v.EffectiveTestServiceEndDateTime = NewDTM(input.Index(7).Index(0))
	v.TestServiceDefaultDurationQuantity = NewNM(input.Index(8).Index(0))
	v.TestServiceDefaultDurationUnits = NewCWE(input.Index(9).Index(0))
	v.TestServiceDefaultFrequency = NewCWE(input.Index(10).Index(0))
	v.ConsentIndicator = NewID(input.Index(11).Index(0))
	v.ConsentIdentifier = NewCWE(input.Index(12).Index(0))
	v.ConsentEffectiveStartDateTime = NewDTM(input.Index(13).Index(0))
	v.ConsentEffectiveEndDateTime = NewDTM(input.Index(14).Index(0))
	v.ConsentIntervalQuantity = NewNM(input.Index(15).Index(0))
	v.ConsentIntervalUnits = NewCWE(input.Index(16).Index(0))
	v.ConsentWaitingPeriodQuantity = NewNM(input.Index(17).Index(0))
	v.ConsentWaitingPeriodUnits = NewCWE(input.Index(18).Index(0))
	v.EffectiveDateTimeOfChange = NewDTM(input.Index(19).Index(0))
	v.EnteredBy = NewXCN(input.Index(20).Index(0))
	v.OrderableAtLocation = NewPLSlice(input.Index(21))
	v.FormularyStatus = NewCWE(input.Index(22).Index(0))
	v.SpecialOrderIndicator = NewID(input.Index(23).Index(0))
	v.PrimaryKeyValueCdm = NewCWESlice(input.Index(24))
	return v
}

// NewOM7Slice will construct a slice of type OM7
func NewOM7Slice(input hl7.HL7Type) []OM7 {
	values := make([]OM7, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOM7(input.Index(i))
	}
	return values
}

type BTS struct {
	BatchMessageCount ST
	BatchComment      ST
	BatchTotals       []NM
}

// NewBTS creates an implementation of BTS
func NewBTS(input hl7.HL7Type) BTS {
	v := BTS{}
	v.BatchMessageCount = NewST(input.Index(1).Index(0))
	v.BatchComment = NewST(input.Index(2).Index(0))
	v.BatchTotals = NewNMSlice(input.Index(3))
	return v
}

// NewBTSSlice will construct a slice of type BTS
func NewBTSSlice(input hl7.HL7Type) []BTS {
	values := make([]BTS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBTS(input.Index(i))
	}
	return values
}

type GOL struct {
	ActionCode                  ID
	ActionDateTime              DTM
	GoalID                      CWE
	GoalInstanceID              EI
	EpisodeOfCareID             EI
	GoalListPriority            NM
	GoalEstablishedDateTime     DTM
	ExpectedGoalAchieveDateTime DTM
	GoalClassification          CWE
	GoalManagementDiscipline    CWE
	CurrentGoalReviewStatus     CWE
	CurrentGoalReviewDateTime   DTM
	NextGoalReviewDateTime      DTM
	PreviousGoalReviewDateTime  DTM
	GoalReviewInterval          ST
	GoalEvaluation              CWE
	GoalEvaluationComment       []ST
	GoalLifeCycleStatus         CWE
	GoalLifeCycleStatusDateTime DTM
	GoalTargetType              []CWE
	GoalTargetName              []XPN
	MoodCode                    CNE
}

// NewGOL creates an implementation of GOL
func NewGOL(input hl7.HL7Type) GOL {
	v := GOL{}
	v.ActionCode = NewID(input.Index(1).Index(0))
	v.ActionDateTime = NewDTM(input.Index(2).Index(0))
	v.GoalID = NewCWE(input.Index(3).Index(0))
	v.GoalInstanceID = NewEI(input.Index(4).Index(0))
	v.EpisodeOfCareID = NewEI(input.Index(5).Index(0))
	v.GoalListPriority = NewNM(input.Index(6).Index(0))
	v.GoalEstablishedDateTime = NewDTM(input.Index(7).Index(0))
	v.ExpectedGoalAchieveDateTime = NewDTM(input.Index(8).Index(0))
	v.GoalClassification = NewCWE(input.Index(9).Index(0))
	v.GoalManagementDiscipline = NewCWE(input.Index(10).Index(0))
	v.CurrentGoalReviewStatus = NewCWE(input.Index(11).Index(0))
	v.CurrentGoalReviewDateTime = NewDTM(input.Index(12).Index(0))
	v.NextGoalReviewDateTime = NewDTM(input.Index(13).Index(0))
	v.PreviousGoalReviewDateTime = NewDTM(input.Index(14).Index(0))
	v.GoalReviewInterval = NewST(input.Index(15).Index(0))
	v.GoalEvaluation = NewCWE(input.Index(16).Index(0))
	v.GoalEvaluationComment = NewSTSlice(input.Index(17))
	v.GoalLifeCycleStatus = NewCWE(input.Index(18).Index(0))
	v.GoalLifeCycleStatusDateTime = NewDTM(input.Index(19).Index(0))
	v.GoalTargetType = NewCWESlice(input.Index(20))
	v.GoalTargetName = NewXPNSlice(input.Index(21))
	v.MoodCode = NewCNE(input.Index(22).Index(0))
	return v
}

// NewGOLSlice will construct a slice of type GOL
func NewGOLSlice(input hl7.HL7Type) []GOL {
	values := make([]GOL, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewGOL(input.Index(i))
	}
	return values
}

type OM4 struct {
	SequenceNumberTestObservationMasterFile NM
	DerivedSpecimen                         ID
	ContainerDescription                    TX
	ContainerVolume                         NM
	ContainerUnits                          CWE
	Specimen                                CWE
	Additive                                CWE
	Preparation                             TX
	SpecialHandlingRequirements             TX
	NormalCollectionVolume                  CQ
	MinimumCollectionVolume                 CQ
	SpecimenRequirements                    TX
	SpecimenPriorities                      []ID
	SpecimenRetentionTime                   CQ
	SpecimenHandlingCode                    []CWE
	SpecimenPreference                      ID
	PreferredSpecimenAttribtureSequenceID   NM
	TaxonomicClassificationCode             []CWE
}

// NewOM4 creates an implementation of OM4
func NewOM4(input hl7.HL7Type) OM4 {
	v := OM4{}
	v.SequenceNumberTestObservationMasterFile = NewNM(input.Index(1).Index(0))
	v.DerivedSpecimen = NewID(input.Index(2).Index(0))
	v.ContainerDescription = NewTX(input.Index(3).Index(0))
	v.ContainerVolume = NewNM(input.Index(4).Index(0))
	v.ContainerUnits = NewCWE(input.Index(5).Index(0))
	v.Specimen = NewCWE(input.Index(6).Index(0))
	v.Additive = NewCWE(input.Index(7).Index(0))
	v.Preparation = NewTX(input.Index(8).Index(0))
	v.SpecialHandlingRequirements = NewTX(input.Index(9).Index(0))
	v.NormalCollectionVolume = NewCQ(input.Index(10).Index(0))
	v.MinimumCollectionVolume = NewCQ(input.Index(11).Index(0))
	v.SpecimenRequirements = NewTX(input.Index(12).Index(0))
	v.SpecimenPriorities = NewIDSlice(input.Index(13))
	v.SpecimenRetentionTime = NewCQ(input.Index(14).Index(0))
	v.SpecimenHandlingCode = NewCWESlice(input.Index(15))
	v.SpecimenPreference = NewID(input.Index(16).Index(0))
	v.PreferredSpecimenAttribtureSequenceID = NewNM(input.Index(17).Index(0))
	v.TaxonomicClassificationCode = NewCWESlice(input.Index(18))
	return v
}

// NewOM4Slice will construct a slice of type OM4
func NewOM4Slice(input hl7.HL7Type) []OM4 {
	values := make([]OM4, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOM4(input.Index(i))
	}
	return values
}

type PKG struct {
	SetID                              SI
	PackagingUnits                     CWE
	DefaultOrderUnitOfMeasureIndicator CNE
	PackageQuantity                    NM
	Price                              CP
	FutureItemPrice                    CP
	FutureItemPriceEffectiveDate       DTM
	GlobalTradeItemNumber              CWE
}

// NewPKG creates an implementation of PKG
func NewPKG(input hl7.HL7Type) PKG {
	v := PKG{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.PackagingUnits = NewCWE(input.Index(2).Index(0))
	v.DefaultOrderUnitOfMeasureIndicator = NewCNE(input.Index(3).Index(0))
	v.PackageQuantity = NewNM(input.Index(4).Index(0))
	v.Price = NewCP(input.Index(5).Index(0))
	v.FutureItemPrice = NewCP(input.Index(6).Index(0))
	v.FutureItemPriceEffectiveDate = NewDTM(input.Index(7).Index(0))
	v.GlobalTradeItemNumber = NewCWE(input.Index(8).Index(0))
	return v
}

// NewPKGSlice will construct a slice of type PKG
func NewPKGSlice(input hl7.HL7Type) []PKG {
	values := make([]PKG, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPKG(input.Index(i))
	}
	return values
}

type NST struct {
	StatisticsAvailable           ID
	SourceIdentifier              ST
	SourceType                    ID
	StatisticsStart               DTM
	StatisticsEnd                 DTM
	ReceiveCharacterCount         NM
	SendCharacterCount            NM
	MessagesReceived              NM
	MessagesSent                  NM
	ChecksumErrorsReceived        NM
	LengthErrorsReceived          NM
	OtherErrorsReceived           NM
	ConnectTimeouts               NM
	ReceiveTimeouts               NM
	ApplicationControlLevelErrors NM
}

// NewNST creates an implementation of NST
func NewNST(input hl7.HL7Type) NST {
	v := NST{}
	v.StatisticsAvailable = NewID(input.Index(1).Index(0))
	v.SourceIdentifier = NewST(input.Index(2).Index(0))
	v.SourceType = NewID(input.Index(3).Index(0))
	v.StatisticsStart = NewDTM(input.Index(4).Index(0))
	v.StatisticsEnd = NewDTM(input.Index(5).Index(0))
	v.ReceiveCharacterCount = NewNM(input.Index(6).Index(0))
	v.SendCharacterCount = NewNM(input.Index(7).Index(0))
	v.MessagesReceived = NewNM(input.Index(8).Index(0))
	v.MessagesSent = NewNM(input.Index(9).Index(0))
	v.ChecksumErrorsReceived = NewNM(input.Index(10).Index(0))
	v.LengthErrorsReceived = NewNM(input.Index(11).Index(0))
	v.OtherErrorsReceived = NewNM(input.Index(12).Index(0))
	v.ConnectTimeouts = NewNM(input.Index(13).Index(0))
	v.ReceiveTimeouts = NewNM(input.Index(14).Index(0))
	v.ApplicationControlLevelErrors = NewNM(input.Index(15).Index(0))
	return v
}

// NewNSTSlice will construct a slice of type NST
func NewNSTSlice(input hl7.HL7Type) []NST {
	values := make([]NST, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewNST(input.Index(i))
	}
	return values
}

type IPC struct {
	AccessionIdentifier            EI
	RequestedProcedureID           EI
	StudyInstanceUID               EI
	ScheduledProcedureStepID       EI
	Modality                       CWE
	ProtocolCode                   []CWE
	ScheduledStationName           EI
	ScheduledProcedureStepLocation []CWE
	ScheduledStationAeTitle        ST
}

// NewIPC creates an implementation of IPC
func NewIPC(input hl7.HL7Type) IPC {
	v := IPC{}
	v.AccessionIdentifier = NewEI(input.Index(1).Index(0))
	v.RequestedProcedureID = NewEI(input.Index(2).Index(0))
	v.StudyInstanceUID = NewEI(input.Index(3).Index(0))
	v.ScheduledProcedureStepID = NewEI(input.Index(4).Index(0))
	v.Modality = NewCWE(input.Index(5).Index(0))
	v.ProtocolCode = NewCWESlice(input.Index(6))
	v.ScheduledStationName = NewEI(input.Index(7).Index(0))
	v.ScheduledProcedureStepLocation = NewCWESlice(input.Index(8))
	v.ScheduledStationAeTitle = NewST(input.Index(9).Index(0))
	return v
}

// NewIPCSlice will construct a slice of type IPC
func NewIPCSlice(input hl7.HL7Type) []IPC {
	values := make([]IPC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewIPC(input.Index(i))
	}
	return values
}

type IN1 struct {
	SetID                         SI
	HealthPlanID                  CWE
	InsuranceCompanyID            []CX
	InsuranceCompanyName          []XON
	InsuranceCompanyAddress       []XAD
	InsuranceCoContactPerson      []XPN
	InsuranceCoPhoneNumber        []XTN
	GroupNumber                   ST
	GroupName                     []XON
	InsuredsGroupEmpID            []CX
	InsuredsGroupEmpName          []XON
	PlanEffectiveDate             DT
	PlanExpirationDate            DT
	AuthorizationInformation      AUI
	PlanType                      CWE
	NameOfInsured                 []XPN
	InsuredsRelationshipToPatient CWE
	InsuredsDateOfBirth           DTM
	InsuredsAddress               []XAD
	AssignmentOfBenefits          CWE
	CoordinationOfBenefits        CWE
	CoordOfBenPriority            ST
	NoticeOfAdmissionFlag         ID
	NoticeOfAdmissionDate         DT
	ReportOfEligibilityFlag       ID
	ReportOfEligibilityDate       DT
	ReleaseInformationCode        CWE
	PreAdmitCertpac               ST
	VerificationDateTime          DTM
	VerificationBy                []XCN
	TypeOfAgreementCode           CWE
	BillingStatus                 CWE
	LifetimeReserveDays           NM
	DelayBeforeLRDay              NM
	CompanyPlanCode               CWE
	PolicyNumber                  ST
	PolicyDeductible              CP
	PolicyLimitAmount             ST
	PolicyLimitDays               NM
	RoomRateSemiPrivate           ST
	RoomRatePrivate               ST
	InsuredsEmploymentStatus      CWE
	InsuredsAdministrativeSex     CWE
	InsuredsEmployersAddress      []XAD
	VerificationStatus            ST
	PriorInsurancePlanID          CWE
	CoverageType                  CWE
	Handicap                      CWE
	InsuredsIDNumber              []CX
	SignatureCode                 CWE
	SignatureCodeDate             DT
	InsuredsBirthPlace            ST
	VipIndicator                  CWE
	ExternalHealthPlanIdentifiers []CX
	InsuranceActionCode           ID
}

// NewIN1 creates an implementation of IN1
func NewIN1(input hl7.HL7Type) IN1 {
	v := IN1{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.HealthPlanID = NewCWE(input.Index(2).Index(0))
	v.InsuranceCompanyID = NewCXSlice(input.Index(3))
	v.InsuranceCompanyName = NewXONSlice(input.Index(4))
	v.InsuranceCompanyAddress = NewXADSlice(input.Index(5))
	v.InsuranceCoContactPerson = NewXPNSlice(input.Index(6))
	v.InsuranceCoPhoneNumber = NewXTNSlice(input.Index(7))
	v.GroupNumber = NewST(input.Index(8).Index(0))
	v.GroupName = NewXONSlice(input.Index(9))
	v.InsuredsGroupEmpID = NewCXSlice(input.Index(10))
	v.InsuredsGroupEmpName = NewXONSlice(input.Index(11))
	v.PlanEffectiveDate = NewDT(input.Index(12).Index(0))
	v.PlanExpirationDate = NewDT(input.Index(13).Index(0))
	v.AuthorizationInformation = NewAUI(input.Index(14).Index(0))
	v.PlanType = NewCWE(input.Index(15).Index(0))
	v.NameOfInsured = NewXPNSlice(input.Index(16))
	v.InsuredsRelationshipToPatient = NewCWE(input.Index(17).Index(0))
	v.InsuredsDateOfBirth = NewDTM(input.Index(18).Index(0))
	v.InsuredsAddress = NewXADSlice(input.Index(19))
	v.AssignmentOfBenefits = NewCWE(input.Index(20).Index(0))
	v.CoordinationOfBenefits = NewCWE(input.Index(21).Index(0))
	v.CoordOfBenPriority = NewST(input.Index(22).Index(0))
	v.NoticeOfAdmissionFlag = NewID(input.Index(23).Index(0))
	v.NoticeOfAdmissionDate = NewDT(input.Index(24).Index(0))
	v.ReportOfEligibilityFlag = NewID(input.Index(25).Index(0))
	v.ReportOfEligibilityDate = NewDT(input.Index(26).Index(0))
	v.ReleaseInformationCode = NewCWE(input.Index(27).Index(0))
	v.PreAdmitCertpac = NewST(input.Index(28).Index(0))
	v.VerificationDateTime = NewDTM(input.Index(29).Index(0))
	v.VerificationBy = NewXCNSlice(input.Index(30))
	v.TypeOfAgreementCode = NewCWE(input.Index(31).Index(0))
	v.BillingStatus = NewCWE(input.Index(32).Index(0))
	v.LifetimeReserveDays = NewNM(input.Index(33).Index(0))
	v.DelayBeforeLRDay = NewNM(input.Index(34).Index(0))
	v.CompanyPlanCode = NewCWE(input.Index(35).Index(0))
	v.PolicyNumber = NewST(input.Index(36).Index(0))
	v.PolicyDeductible = NewCP(input.Index(37).Index(0))
	v.PolicyLimitAmount = NewST(input.Index(38).Index(0))
	v.PolicyLimitDays = NewNM(input.Index(39).Index(0))
	v.RoomRateSemiPrivate = NewST(input.Index(40).Index(0))
	v.RoomRatePrivate = NewST(input.Index(41).Index(0))
	v.InsuredsEmploymentStatus = NewCWE(input.Index(42).Index(0))
	v.InsuredsAdministrativeSex = NewCWE(input.Index(43).Index(0))
	v.InsuredsEmployersAddress = NewXADSlice(input.Index(44))
	v.VerificationStatus = NewST(input.Index(45).Index(0))
	v.PriorInsurancePlanID = NewCWE(input.Index(46).Index(0))
	v.CoverageType = NewCWE(input.Index(47).Index(0))
	v.Handicap = NewCWE(input.Index(48).Index(0))
	v.InsuredsIDNumber = NewCXSlice(input.Index(49))
	v.SignatureCode = NewCWE(input.Index(50).Index(0))
	v.SignatureCodeDate = NewDT(input.Index(51).Index(0))
	v.InsuredsBirthPlace = NewST(input.Index(52).Index(0))
	v.VipIndicator = NewCWE(input.Index(53).Index(0))
	v.ExternalHealthPlanIdentifiers = NewCXSlice(input.Index(54))
	v.InsuranceActionCode = NewID(input.Index(55).Index(0))
	return v
}

// NewIN1Slice will construct a slice of type IN1
func NewIN1Slice(input hl7.HL7Type) []IN1 {
	values := make([]IN1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewIN1(input.Index(i))
	}
	return values
}

type FTS struct {
	FileBatchCount     NM
	FileTrailerComment ST
}

// NewFTS creates an implementation of FTS
func NewFTS(input hl7.HL7Type) FTS {
	v := FTS{}
	v.FileBatchCount = NewNM(input.Index(1).Index(0))
	v.FileTrailerComment = NewST(input.Index(2).Index(0))
	return v
}

// NewFTSSlice will construct a slice of type FTS
func NewFTSSlice(input hl7.HL7Type) []FTS {
	values := make([]FTS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewFTS(input.Index(i))
	}
	return values
}

type ACC struct {
	AccidentDateTime            DTM
	AccidentCode                CWE
	AccidentLocation            ST
	AutoAccidentState           CWE
	AccidentJobRelatedIndicator ID
	AccidentDeathIndicator      ID
	EnteredBy                   XCN
	AccidentDescription         ST
	BroughtInBy                 ST
	PoliceNotifiedIndicator     ID
	AccidentAddress             XAD
	DegreeOfPatientLiability    NM
	AccidentIdentifier          []EI
}

// NewACC creates an implementation of ACC
func NewACC(input hl7.HL7Type) ACC {
	v := ACC{}
	v.AccidentDateTime = NewDTM(input.Index(1).Index(0))
	v.AccidentCode = NewCWE(input.Index(2).Index(0))
	v.AccidentLocation = NewST(input.Index(3).Index(0))
	v.AutoAccidentState = NewCWE(input.Index(4).Index(0))
	v.AccidentJobRelatedIndicator = NewID(input.Index(5).Index(0))
	v.AccidentDeathIndicator = NewID(input.Index(6).Index(0))
	v.EnteredBy = NewXCN(input.Index(7).Index(0))
	v.AccidentDescription = NewST(input.Index(8).Index(0))
	v.BroughtInBy = NewST(input.Index(9).Index(0))
	v.PoliceNotifiedIndicator = NewID(input.Index(10).Index(0))
	v.AccidentAddress = NewXAD(input.Index(11).Index(0))
	v.DegreeOfPatientLiability = NewNM(input.Index(12).Index(0))
	v.AccidentIdentifier = NewEISlice(input.Index(13))
	return v
}

// NewACCSlice will construct a slice of type ACC
func NewACCSlice(input hl7.HL7Type) []ACC {
	values := make([]ACC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewACC(input.Index(i))
	}
	return values
}

type OM5 struct {
	SequenceNumberTestObservationMasterFile            NM
	TestObservationsIncludedWithinAnOrderedTestBattery []CWE
	ObservationIDSuffixes                              ST
}

// NewOM5 creates an implementation of OM5
func NewOM5(input hl7.HL7Type) OM5 {
	v := OM5{}
	v.SequenceNumberTestObservationMasterFile = NewNM(input.Index(1).Index(0))
	v.TestObservationsIncludedWithinAnOrderedTestBattery = NewCWESlice(input.Index(2))
	v.ObservationIDSuffixes = NewST(input.Index(3).Index(0))
	return v
}

// NewOM5Slice will construct a slice of type OM5
func NewOM5Slice(input hl7.HL7Type) []OM5 {
	values := make([]OM5, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOM5(input.Index(i))
	}
	return values
}

type OM6 struct {
	SequenceNumberTestObservationMasterFile NM
	DerivationRule                          TX
}

// NewOM6 creates an implementation of OM6
func NewOM6(input hl7.HL7Type) OM6 {
	v := OM6{}
	v.SequenceNumberTestObservationMasterFile = NewNM(input.Index(1).Index(0))
	v.DerivationRule = NewTX(input.Index(2).Index(0))
	return v
}

// NewOM6Slice will construct a slice of type OM6
func NewOM6Slice(input hl7.HL7Type) []OM6 {
	values := make([]OM6, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOM6(input.Index(i))
	}
	return values
}

type CON struct {
	SetID                                  SI
	ConsentType                            CWE
	ConsentFormIDAndVersion                ST
	ConsentFormNumber                      EI
	ConsentText                            []FT
	SubjectSpecificConsentText             []FT
	ConsentBackgroundInformation           []FT
	SubjectSpecificConsentBackgroundText   []FT
	ConsenterImposedLimitations            []FT
	ConsentMode                            CNE
	ConsentStatus                          CNE
	ConsentDiscussionDateTime              DTM
	ConsentDecisionDateTime                DTM
	ConsentEffectiveDateTime               DTM
	ConsentEndDateTime                     DTM
	SubjectCompetenceIndicator             ID
	TranslatorAssistanceIndicator          ID
	LanguageTranslatedTo                   CWE
	InformationalMaterialSuppliedIndicator ID
	ConsentBypassReason                    CWE
	ConsentDisclosureLevel                 ID
	ConsentNonDisclosureReason             CWE
	NonSubjectConsenterReason              CWE
	ConsenterID                            []XPN
	RelationshipToSubject                  []CWE
}

// NewCON creates an implementation of CON
func NewCON(input hl7.HL7Type) CON {
	v := CON{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.ConsentType = NewCWE(input.Index(2).Index(0))
	v.ConsentFormIDAndVersion = NewST(input.Index(3).Index(0))
	v.ConsentFormNumber = NewEI(input.Index(4).Index(0))
	v.ConsentText = NewFTSlice(input.Index(5))
	v.SubjectSpecificConsentText = NewFTSlice(input.Index(6))
	v.ConsentBackgroundInformation = NewFTSlice(input.Index(7))
	v.SubjectSpecificConsentBackgroundText = NewFTSlice(input.Index(8))
	v.ConsenterImposedLimitations = NewFTSlice(input.Index(9))
	v.ConsentMode = NewCNE(input.Index(10).Index(0))
	v.ConsentStatus = NewCNE(input.Index(11).Index(0))
	v.ConsentDiscussionDateTime = NewDTM(input.Index(12).Index(0))
	v.ConsentDecisionDateTime = NewDTM(input.Index(13).Index(0))
	v.ConsentEffectiveDateTime = NewDTM(input.Index(14).Index(0))
	v.ConsentEndDateTime = NewDTM(input.Index(15).Index(0))
	v.SubjectCompetenceIndicator = NewID(input.Index(16).Index(0))
	v.TranslatorAssistanceIndicator = NewID(input.Index(17).Index(0))
	v.LanguageTranslatedTo = NewCWE(input.Index(18).Index(0))
	v.InformationalMaterialSuppliedIndicator = NewID(input.Index(19).Index(0))
	v.ConsentBypassReason = NewCWE(input.Index(20).Index(0))
	v.ConsentDisclosureLevel = NewID(input.Index(21).Index(0))
	v.ConsentNonDisclosureReason = NewCWE(input.Index(22).Index(0))
	v.NonSubjectConsenterReason = NewCWE(input.Index(23).Index(0))
	v.ConsenterID = NewXPNSlice(input.Index(24))
	v.RelationshipToSubject = NewCWESlice(input.Index(25))
	return v
}

// NewCONSlice will construct a slice of type CON
func NewCONSlice(input hl7.HL7Type) []CON {
	values := make([]CON, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCON(input.Index(i))
	}
	return values
}

type OM2 struct {
	SequenceNumberTestObservationMasterFile                 NM
	UnitsOfMeasure                                          CWE
	RangeOfDecimalPrecision                                 []NM
	CorrespondingSiUnitsOfMeasure                           CWE
	SiConversionFactor                                      TX
	ReferencenormalRangeForOrdinalAndContinuousObservations []RFR
	CriticalRangeForOrdinalAndContinuousObservations        []RFR
	AbsoluteRangeForOrdinalAndContinuousObservations        RFR
	DeltaCheckCriteria                                      []DLT
	MinimumMeaningfulIncrements                             NM
}

// NewOM2 creates an implementation of OM2
func NewOM2(input hl7.HL7Type) OM2 {
	v := OM2{}
	v.SequenceNumberTestObservationMasterFile = NewNM(input.Index(1).Index(0))
	v.UnitsOfMeasure = NewCWE(input.Index(2).Index(0))
	v.RangeOfDecimalPrecision = NewNMSlice(input.Index(3))
	v.CorrespondingSiUnitsOfMeasure = NewCWE(input.Index(4).Index(0))
	v.SiConversionFactor = NewTX(input.Index(5).Index(0))
	v.ReferencenormalRangeForOrdinalAndContinuousObservations = NewRFRSlice(input.Index(6))
	v.CriticalRangeForOrdinalAndContinuousObservations = NewRFRSlice(input.Index(7))
	v.AbsoluteRangeForOrdinalAndContinuousObservations = NewRFR(input.Index(8).Index(0))
	v.DeltaCheckCriteria = NewDLTSlice(input.Index(9))
	v.MinimumMeaningfulIncrements = NewNM(input.Index(10).Index(0))
	return v
}

// NewOM2Slice will construct a slice of type OM2
func NewOM2Slice(input hl7.HL7Type) []OM2 {
	values := make([]OM2, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOM2(input.Index(i))
	}
	return values
}

type SLT struct {
	DeviceNumber   EI
	DeviceName     ST
	LotNumber      EI
	ItemIdentifier EI
	BarCode        ST
}

// NewSLT creates an implementation of SLT
func NewSLT(input hl7.HL7Type) SLT {
	v := SLT{}
	v.DeviceNumber = NewEI(input.Index(1).Index(0))
	v.DeviceName = NewST(input.Index(2).Index(0))
	v.LotNumber = NewEI(input.Index(3).Index(0))
	v.ItemIdentifier = NewEI(input.Index(4).Index(0))
	v.BarCode = NewST(input.Index(5).Index(0))
	return v
}

// NewSLTSlice will construct a slice of type SLT
func NewSLTSlice(input hl7.HL7Type) []SLT {
	values := make([]SLT, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSLT(input.Index(i))
	}
	return values
}

type STF struct {
	PrimaryKeyValueStf              CWE
	StaffIdentifierList             []CX
	StaffName                       []XPN
	StaffType                       []CWE
	AdministrativeSex               CWE
	DateTimeOfBirth                 DTM
	ActiveInactiveFlag              ID
	Department                      []CWE
	HospitalServiceStf              []CWE
	Phone                           []XTN
	OfficeHomeAddressBirthplace     []XAD
	InstitutionActivationDate       []DIN
	InstitutionInactivationDate     []DIN
	BackupPersonID                  []CWE
	EMailAddress                    []ST
	PreferredMethodOfContact        CWE
	MaritalStatus                   CWE
	JobTitle                        ST
	JobCodeClass                    JCC
	EmploymentStatusCode            CWE
	AdditionalInsuredOnAuto         ID
	DriversLicenseNumberStaff       DLN
	CopyAutoIns                     ID
	AutoInsExpires                  DT
	DateLastDmvReview               DT
	DateNextDmvReview               DT
	Race                            CWE
	EthnicGroup                     CWE
	ReActivationApprovalIndicator   ID
	Citizenship                     []CWE
	DateTimeOfDeath                 DTM
	DeathIndicator                  ID
	InstitutionRelationshipTypeCode CWE
	InstitutionRelationshipPeriod   DR
	ExpectedReturnDate              DT
	CostCenterCode                  []CWE
	GenericClassificationIndicator  ID
	InactiveReasonCode              CWE
	GenericResourceTypeOrCategory   []CWE
	Religion                        CWE
	Signature                       ED
}

// NewSTF creates an implementation of STF
func NewSTF(input hl7.HL7Type) STF {
	v := STF{}
	v.PrimaryKeyValueStf = NewCWE(input.Index(1).Index(0))
	v.StaffIdentifierList = NewCXSlice(input.Index(2))
	v.StaffName = NewXPNSlice(input.Index(3))
	v.StaffType = NewCWESlice(input.Index(4))
	v.AdministrativeSex = NewCWE(input.Index(5).Index(0))
	v.DateTimeOfBirth = NewDTM(input.Index(6).Index(0))
	v.ActiveInactiveFlag = NewID(input.Index(7).Index(0))
	v.Department = NewCWESlice(input.Index(8))
	v.HospitalServiceStf = NewCWESlice(input.Index(9))
	v.Phone = NewXTNSlice(input.Index(10))
	v.OfficeHomeAddressBirthplace = NewXADSlice(input.Index(11))
	v.InstitutionActivationDate = NewDINSlice(input.Index(12))
	v.InstitutionInactivationDate = NewDINSlice(input.Index(13))
	v.BackupPersonID = NewCWESlice(input.Index(14))
	v.EMailAddress = NewSTSlice(input.Index(15))
	v.PreferredMethodOfContact = NewCWE(input.Index(16).Index(0))
	v.MaritalStatus = NewCWE(input.Index(17).Index(0))
	v.JobTitle = NewST(input.Index(18).Index(0))
	v.JobCodeClass = NewJCC(input.Index(19).Index(0))
	v.EmploymentStatusCode = NewCWE(input.Index(20).Index(0))
	v.AdditionalInsuredOnAuto = NewID(input.Index(21).Index(0))
	v.DriversLicenseNumberStaff = NewDLN(input.Index(22).Index(0))
	v.CopyAutoIns = NewID(input.Index(23).Index(0))
	v.AutoInsExpires = NewDT(input.Index(24).Index(0))
	v.DateLastDmvReview = NewDT(input.Index(25).Index(0))
	v.DateNextDmvReview = NewDT(input.Index(26).Index(0))
	v.Race = NewCWE(input.Index(27).Index(0))
	v.EthnicGroup = NewCWE(input.Index(28).Index(0))
	v.ReActivationApprovalIndicator = NewID(input.Index(29).Index(0))
	v.Citizenship = NewCWESlice(input.Index(30))
	v.DateTimeOfDeath = NewDTM(input.Index(31).Index(0))
	v.DeathIndicator = NewID(input.Index(32).Index(0))
	v.InstitutionRelationshipTypeCode = NewCWE(input.Index(33).Index(0))
	v.InstitutionRelationshipPeriod = NewDR(input.Index(34).Index(0))
	v.ExpectedReturnDate = NewDT(input.Index(35).Index(0))
	v.CostCenterCode = NewCWESlice(input.Index(36))
	v.GenericClassificationIndicator = NewID(input.Index(37).Index(0))
	v.InactiveReasonCode = NewCWE(input.Index(38).Index(0))
	v.GenericResourceTypeOrCategory = NewCWESlice(input.Index(39))
	v.Religion = NewCWE(input.Index(40).Index(0))
	v.Signature = NewED(input.Index(41).Index(0))
	return v
}

// NewSTFSlice will construct a slice of type STF
func NewSTFSlice(input hl7.HL7Type) []STF {
	values := make([]STF, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSTF(input.Index(i))
	}
	return values
}

type SFT struct {
	SoftwareVendorOrganization              XON
	SoftwareCertifiedVersionOrReleaseNumber ST
	SoftwareProductName                     ST
	SoftwareBinaryID                        ST
	SoftwareProductInformation              TX
	SoftwareInstallDate                     DTM
}

// NewSFT creates an implementation of SFT
func NewSFT(input hl7.HL7Type) SFT {
	v := SFT{}
	v.SoftwareVendorOrganization = NewXON(input.Index(1).Index(0))
	v.SoftwareCertifiedVersionOrReleaseNumber = NewST(input.Index(2).Index(0))
	v.SoftwareProductName = NewST(input.Index(3).Index(0))
	v.SoftwareBinaryID = NewST(input.Index(4).Index(0))
	v.SoftwareProductInformation = NewTX(input.Index(5).Index(0))
	v.SoftwareInstallDate = NewDTM(input.Index(6).Index(0))
	return v
}

// NewSFTSlice will construct a slice of type SFT
func NewSFTSlice(input hl7.HL7Type) []SFT {
	values := make([]SFT, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSFT(input.Index(i))
	}
	return values
}

type BLC struct {
	BloodProductCode CWE
	BloodAmount      CQ
}

// NewBLC creates an implementation of BLC
func NewBLC(input hl7.HL7Type) BLC {
	v := BLC{}
	v.BloodProductCode = NewCWE(input.Index(1).Index(0))
	v.BloodAmount = NewCQ(input.Index(2).Index(0))
	return v
}

// NewBLCSlice will construct a slice of type BLC
func NewBLCSlice(input hl7.HL7Type) []BLC {
	values := make([]BLC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBLC(input.Index(i))
	}
	return values
}

type URD struct {
	URD1 ST
}

// NewURD creates an implementation of URD
func NewURD(input hl7.HL7Type) URD {
	v := URD{}
	v.URD1 = NewST(input.Index(1).Index(0))
	return v
}

// NewURDSlice will construct a slice of type URD
func NewURDSlice(input hl7.HL7Type) []URD {
	values := make([]URD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewURD(input.Index(i))
	}
	return values
}

type OM1 struct {
	SequenceNumberTestObservationMasterFile                NM
	ProducersServiceTestObservationID                      CWE
	PermittedDataTypes                                     []ID
	SpecimenRequired                                       ID
	ProducerID                                             CWE
	ObservationDescription                                 TX
	OtherServiceTestObservationIDsForTheObservation        CWE
	OtherNames                                             []ST
	PreferredReportNameForTheObservation                   ST
	PreferredShortNameOrMnemonicForTheObservation          ST
	PreferredLongNameForTheObservation                     ST
	Orderability                                           ID
	IdentityOfInstrumentUsedToPerformThisStudy             []CWE
	CodedRepresentationOfMethod                            []CWE
	PortableDeviceIndicator                                ID
	ObservationProducingDepartmentSection                  []CWE
	TelephoneNumberOfSection                               XTN
	NatureOfServiceTestObservation                         CWE
	ReportSubheader                                        CWE
	ReportDisplayOrder                                     ST
	DateTimeStampForAnyChangeInDefinitionForTheObservation DTM
	EffectiveDateTimeOfChange                              DTM
	TypicalTurnAroundTime                                  NM
	ProcessingTime                                         NM
	ProcessingPriority                                     []ID
	ReportingPriority                                      ID
	OutsideSitesWhereObservationMayBePerformed             []CWE
	AddressOfOutsideSites                                  []XAD
	PhoneNumberOfOutsideSite                               XTN
	ConfidentialityCode                                    CWE
	ObservationsRequiredToInterpretThisObservation         CWE
	InterpretationOfObservations                           TX
	ContraindicationsToObservations                        CWE
	ReflexTestsObservations                                []CWE
	RulesThatTriggerReflexTesting                          TX
	FixedCannedMessage                                     CWE
	PatientPreparation                                     TX
	ProcedureMedication                                    CWE
	FactorsThatMayAffectTheObservation                     TX
	ServiceTestObservationPerformanceSchedule              []ST
	DescriptionOfTestMethods                               TX
	KindOfQuantityObserved                                 CWE
	PointVersusInterval                                    CWE
	ChallengeInformation                                   TX
	RelationshipModifier                                   CWE
	TargetAnatomicSiteOfTest                               CWE
	ModalityOfImagingMeasurement                           CWE
	ExclusiveTest                                          ID
	DiagnosticServSectID                                   ID
	TaxonomicClassificationCode                            CWE
	AdditionalNames                                        []ST
}

// NewOM1 creates an implementation of OM1
func NewOM1(input hl7.HL7Type) OM1 {
	v := OM1{}
	v.SequenceNumberTestObservationMasterFile = NewNM(input.Index(1).Index(0))
	v.ProducersServiceTestObservationID = NewCWE(input.Index(2).Index(0))
	v.PermittedDataTypes = NewIDSlice(input.Index(3))
	v.SpecimenRequired = NewID(input.Index(4).Index(0))
	v.ProducerID = NewCWE(input.Index(5).Index(0))
	v.ObservationDescription = NewTX(input.Index(6).Index(0))
	v.OtherServiceTestObservationIDsForTheObservation = NewCWE(input.Index(7).Index(0))
	v.OtherNames = NewSTSlice(input.Index(8))
	v.PreferredReportNameForTheObservation = NewST(input.Index(9).Index(0))
	v.PreferredShortNameOrMnemonicForTheObservation = NewST(input.Index(10).Index(0))
	v.PreferredLongNameForTheObservation = NewST(input.Index(11).Index(0))
	v.Orderability = NewID(input.Index(12).Index(0))
	v.IdentityOfInstrumentUsedToPerformThisStudy = NewCWESlice(input.Index(13))
	v.CodedRepresentationOfMethod = NewCWESlice(input.Index(14))
	v.PortableDeviceIndicator = NewID(input.Index(15).Index(0))
	v.ObservationProducingDepartmentSection = NewCWESlice(input.Index(16))
	v.TelephoneNumberOfSection = NewXTN(input.Index(17).Index(0))
	v.NatureOfServiceTestObservation = NewCWE(input.Index(18).Index(0))
	v.ReportSubheader = NewCWE(input.Index(19).Index(0))
	v.ReportDisplayOrder = NewST(input.Index(20).Index(0))
	v.DateTimeStampForAnyChangeInDefinitionForTheObservation = NewDTM(input.Index(21).Index(0))
	v.EffectiveDateTimeOfChange = NewDTM(input.Index(22).Index(0))
	v.TypicalTurnAroundTime = NewNM(input.Index(23).Index(0))
	v.ProcessingTime = NewNM(input.Index(24).Index(0))
	v.ProcessingPriority = NewIDSlice(input.Index(25))
	v.ReportingPriority = NewID(input.Index(26).Index(0))
	v.OutsideSitesWhereObservationMayBePerformed = NewCWESlice(input.Index(27))
	v.AddressOfOutsideSites = NewXADSlice(input.Index(28))
	v.PhoneNumberOfOutsideSite = NewXTN(input.Index(29).Index(0))
	v.ConfidentialityCode = NewCWE(input.Index(30).Index(0))
	v.ObservationsRequiredToInterpretThisObservation = NewCWE(input.Index(31).Index(0))
	v.InterpretationOfObservations = NewTX(input.Index(32).Index(0))
	v.ContraindicationsToObservations = NewCWE(input.Index(33).Index(0))
	v.ReflexTestsObservations = NewCWESlice(input.Index(34))
	v.RulesThatTriggerReflexTesting = NewTX(input.Index(35).Index(0))
	v.FixedCannedMessage = NewCWE(input.Index(36).Index(0))
	v.PatientPreparation = NewTX(input.Index(37).Index(0))
	v.ProcedureMedication = NewCWE(input.Index(38).Index(0))
	v.FactorsThatMayAffectTheObservation = NewTX(input.Index(39).Index(0))
	v.ServiceTestObservationPerformanceSchedule = NewSTSlice(input.Index(40))
	v.DescriptionOfTestMethods = NewTX(input.Index(41).Index(0))
	v.KindOfQuantityObserved = NewCWE(input.Index(42).Index(0))
	v.PointVersusInterval = NewCWE(input.Index(43).Index(0))
	v.ChallengeInformation = NewTX(input.Index(44).Index(0))
	v.RelationshipModifier = NewCWE(input.Index(45).Index(0))
	v.TargetAnatomicSiteOfTest = NewCWE(input.Index(46).Index(0))
	v.ModalityOfImagingMeasurement = NewCWE(input.Index(47).Index(0))
	v.ExclusiveTest = NewID(input.Index(48).Index(0))
	v.DiagnosticServSectID = NewID(input.Index(49).Index(0))
	v.TaxonomicClassificationCode = NewCWE(input.Index(50).Index(0))
	v.AdditionalNames = NewSTSlice(input.Index(51))
	return v
}

// NewOM1Slice will construct a slice of type OM1
func NewOM1Slice(input hl7.HL7Type) []OM1 {
	values := make([]OM1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOM1(input.Index(i))
	}
	return values
}

type PCR struct {
	ImplicatedProduct                 CWE
	GenericProduct                    IS
	ProductClass                      CWE
	TotalDurationOfTherapy            CQ
	ProductManufactureDate            DTM
	ProductExpirationDate             DTM
	ProductImplantationDate           DTM
	ProductExplantationDate           DTM
	SingleUseDevice                   CWE
	IndicationForProductUse           CWE
	ProductProblem                    CWE
	ProductSerialLotNumber            []ST
	ProductAvailableForInspection     CWE
	ProductEvaluationPerformed        CWE
	ProductEvaluationStatus           CWE
	ProductEvaluationResults          CWE
	EvaluatedProductSource            ID
	DateProductReturnedToManufacturer DTM
	DeviceOperatorQualifications      ID
	RelatednessAssessment             ID
	ActionTakenInResponseToTheEvent   []ID
	EventCausalityObservations        []ID
	IndirectExposureMechanism         []ID
}

// NewPCR creates an implementation of PCR
func NewPCR(input hl7.HL7Type) PCR {
	v := PCR{}
	v.ImplicatedProduct = NewCWE(input.Index(1).Index(0))
	v.GenericProduct = NewIS(input.Index(2).Index(0))
	v.ProductClass = NewCWE(input.Index(3).Index(0))
	v.TotalDurationOfTherapy = NewCQ(input.Index(4).Index(0))
	v.ProductManufactureDate = NewDTM(input.Index(5).Index(0))
	v.ProductExpirationDate = NewDTM(input.Index(6).Index(0))
	v.ProductImplantationDate = NewDTM(input.Index(7).Index(0))
	v.ProductExplantationDate = NewDTM(input.Index(8).Index(0))
	v.SingleUseDevice = NewCWE(input.Index(9).Index(0))
	v.IndicationForProductUse = NewCWE(input.Index(10).Index(0))
	v.ProductProblem = NewCWE(input.Index(11).Index(0))
	v.ProductSerialLotNumber = NewSTSlice(input.Index(12))
	v.ProductAvailableForInspection = NewCWE(input.Index(13).Index(0))
	v.ProductEvaluationPerformed = NewCWE(input.Index(14).Index(0))
	v.ProductEvaluationStatus = NewCWE(input.Index(15).Index(0))
	v.ProductEvaluationResults = NewCWE(input.Index(16).Index(0))
	v.EvaluatedProductSource = NewID(input.Index(17).Index(0))
	v.DateProductReturnedToManufacturer = NewDTM(input.Index(18).Index(0))
	v.DeviceOperatorQualifications = NewID(input.Index(19).Index(0))
	v.RelatednessAssessment = NewID(input.Index(20).Index(0))
	v.ActionTakenInResponseToTheEvent = NewIDSlice(input.Index(21))
	v.EventCausalityObservations = NewIDSlice(input.Index(22))
	v.IndirectExposureMechanism = NewIDSlice(input.Index(23))
	return v
}

// NewPCRSlice will construct a slice of type PCR
func NewPCRSlice(input hl7.HL7Type) []PCR {
	values := make([]PCR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPCR(input.Index(i))
	}
	return values
}

type CTI struct {
	SponsorStudyID          EI
	StudyPhaseIdentifier    CWE
	StudyScheduledTimePoint CWE
}

// NewCTI creates an implementation of CTI
func NewCTI(input hl7.HL7Type) CTI {
	v := CTI{}
	v.SponsorStudyID = NewEI(input.Index(1).Index(0))
	v.StudyPhaseIdentifier = NewCWE(input.Index(2).Index(0))
	v.StudyScheduledTimePoint = NewCWE(input.Index(3).Index(0))
	return v
}

// NewCTISlice will construct a slice of type CTI
func NewCTISlice(input hl7.HL7Type) []CTI {
	values := make([]CTI, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCTI(input.Index(i))
	}
	return values
}

type PRB struct {
	ActionCode                                        ID
	ActionDateTime                                    DTM
	ProblemID                                         CWE
	ProblemInstanceID                                 EI
	EpisodeOfCareID                                   EI
	ProblemListPriority                               NM
	ProblemEstablishedDateTime                        DTM
	AnticipatedProblemResolutionDateTime              DTM
	ActualProblemResolutionDateTime                   DTM
	ProblemClassification                             CWE
	ProblemManagementDiscipline                       []CWE
	ProblemPersistence                                CWE
	ProblemConfirmationStatus                         CWE
	ProblemLifeCycleStatus                            CWE
	ProblemLifeCycleStatusDateTime                    DTM
	ProblemDateOfOnset                                DTM
	ProblemOnsetText                                  ST
	ProblemRanking                                    CWE
	CertaintyOfProblem                                CWE
	ProbabilityOfProblem01                            NM
	IndividualAwarenessOfProblem                      CWE
	ProblemPrognosis                                  CWE
	IndividualAwarenessOfPrognosis                    CWE
	FamilySignificantOtherAwarenessOfProblemPrognosis ST
	SecuritySensitivity                               CWE
	ProblemSeverity                                   CWE
	ProblemPerspective                                CWE
	MoodCode                                          CNE
}

// NewPRB creates an implementation of PRB
func NewPRB(input hl7.HL7Type) PRB {
	v := PRB{}
	v.ActionCode = NewID(input.Index(1).Index(0))
	v.ActionDateTime = NewDTM(input.Index(2).Index(0))
	v.ProblemID = NewCWE(input.Index(3).Index(0))
	v.ProblemInstanceID = NewEI(input.Index(4).Index(0))
	v.EpisodeOfCareID = NewEI(input.Index(5).Index(0))
	v.ProblemListPriority = NewNM(input.Index(6).Index(0))
	v.ProblemEstablishedDateTime = NewDTM(input.Index(7).Index(0))
	v.AnticipatedProblemResolutionDateTime = NewDTM(input.Index(8).Index(0))
	v.ActualProblemResolutionDateTime = NewDTM(input.Index(9).Index(0))
	v.ProblemClassification = NewCWE(input.Index(10).Index(0))
	v.ProblemManagementDiscipline = NewCWESlice(input.Index(11))
	v.ProblemPersistence = NewCWE(input.Index(12).Index(0))
	v.ProblemConfirmationStatus = NewCWE(input.Index(13).Index(0))
	v.ProblemLifeCycleStatus = NewCWE(input.Index(14).Index(0))
	v.ProblemLifeCycleStatusDateTime = NewDTM(input.Index(15).Index(0))
	v.ProblemDateOfOnset = NewDTM(input.Index(16).Index(0))
	v.ProblemOnsetText = NewST(input.Index(17).Index(0))
	v.ProblemRanking = NewCWE(input.Index(18).Index(0))
	v.CertaintyOfProblem = NewCWE(input.Index(19).Index(0))
	v.ProbabilityOfProblem01 = NewNM(input.Index(20).Index(0))
	v.IndividualAwarenessOfProblem = NewCWE(input.Index(21).Index(0))
	v.ProblemPrognosis = NewCWE(input.Index(22).Index(0))
	v.IndividualAwarenessOfPrognosis = NewCWE(input.Index(23).Index(0))
	v.FamilySignificantOtherAwarenessOfProblemPrognosis = NewST(input.Index(24).Index(0))
	v.SecuritySensitivity = NewCWE(input.Index(25).Index(0))
	v.ProblemSeverity = NewCWE(input.Index(26).Index(0))
	v.ProblemPerspective = NewCWE(input.Index(27).Index(0))
	v.MoodCode = NewCNE(input.Index(28).Index(0))
	return v
}

// NewPRBSlice will construct a slice of type PRB
func NewPRBSlice(input hl7.HL7Type) []PRB {
	values := make([]PRB, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPRB(input.Index(i))
	}
	return values
}

type ORG struct {
	SetID                                      SI
	OrganizationUnitCode                       CWE
	OrganizationUnitTypeCode                   CWE
	PrimaryOrgUnitIndicator                    ID
	PractitionerOrgUnitIdentifier              CX
	HealthCareProviderTypeCode                 CWE
	HealthCareProviderClassificationCode       CWE
	HealthCareProviderAreaOfSpecializationCode CWE
	EffectiveDateRange                         DR
	EmploymentStatusCode                       CWE
	BoardApprovalIndicator                     ID
	PrimaryCarePhysicianIndicator              ID
	CostCenterCode                             []CWE
}

// NewORG creates an implementation of ORG
func NewORG(input hl7.HL7Type) ORG {
	v := ORG{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.OrganizationUnitCode = NewCWE(input.Index(2).Index(0))
	v.OrganizationUnitTypeCode = NewCWE(input.Index(3).Index(0))
	v.PrimaryOrgUnitIndicator = NewID(input.Index(4).Index(0))
	v.PractitionerOrgUnitIdentifier = NewCX(input.Index(5).Index(0))
	v.HealthCareProviderTypeCode = NewCWE(input.Index(6).Index(0))
	v.HealthCareProviderClassificationCode = NewCWE(input.Index(7).Index(0))
	v.HealthCareProviderAreaOfSpecializationCode = NewCWE(input.Index(8).Index(0))
	v.EffectiveDateRange = NewDR(input.Index(9).Index(0))
	v.EmploymentStatusCode = NewCWE(input.Index(10).Index(0))
	v.BoardApprovalIndicator = NewID(input.Index(11).Index(0))
	v.PrimaryCarePhysicianIndicator = NewID(input.Index(12).Index(0))
	v.CostCenterCode = NewCWESlice(input.Index(13))
	return v
}

// NewORGSlice will construct a slice of type ORG
func NewORGSlice(input hl7.HL7Type) []ORG {
	values := make([]ORG, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORG(input.Index(i))
	}
	return values
}

type MRG struct {
	PriorPatientIdentifierList []CX
	PriorAlternatePatientID    ST
	PriorPatientAccountNumber  CX
	PriorPatientID             ST
	PriorVisitNumber           CX
	PriorAlternateVisitID      CX
	PriorPatientName           []XPN
}

// NewMRG creates an implementation of MRG
func NewMRG(input hl7.HL7Type) MRG {
	v := MRG{}
	v.PriorPatientIdentifierList = NewCXSlice(input.Index(1))
	v.PriorAlternatePatientID = NewST(input.Index(2).Index(0))
	v.PriorPatientAccountNumber = NewCX(input.Index(3).Index(0))
	v.PriorPatientID = NewST(input.Index(4).Index(0))
	v.PriorVisitNumber = NewCX(input.Index(5).Index(0))
	v.PriorAlternateVisitID = NewCX(input.Index(6).Index(0))
	v.PriorPatientName = NewXPNSlice(input.Index(7))
	return v
}

// NewMRGSlice will construct a slice of type MRG
func NewMRGSlice(input hl7.HL7Type) []MRG {
	values := make([]MRG, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMRG(input.Index(i))
	}
	return values
}

type CSR struct {
	SponsorStudyID                     EI
	AlternateStudyID                   EI
	InstitutionRegisteringThePatient   CWE
	SponsorPatientID                   CX
	AlternatePatientIDCsr              CX
	DateTimeOfPatientStudyRegistration DTM
	PersonPerformingStudyRegistration  []XCN
	StudyAuthorizingProvider           []XCN
	DateTimePatientStudyConsentSigned  DTM
	PatientStudyEligibilityStatus      CWE
	StudyRandomizationDateTime         []DTM
	RandomizedStudyArm                 []CWE
	StratumForStudyRandomization       []CWE
	PatientEvaluabilityStatus          CWE
	DateTimeEndedStudy                 DTM
	ReasonEndedStudy                   CWE
}

// NewCSR creates an implementation of CSR
func NewCSR(input hl7.HL7Type) CSR {
	v := CSR{}
	v.SponsorStudyID = NewEI(input.Index(1).Index(0))
	v.AlternateStudyID = NewEI(input.Index(2).Index(0))
	v.InstitutionRegisteringThePatient = NewCWE(input.Index(3).Index(0))
	v.SponsorPatientID = NewCX(input.Index(4).Index(0))
	v.AlternatePatientIDCsr = NewCX(input.Index(5).Index(0))
	v.DateTimeOfPatientStudyRegistration = NewDTM(input.Index(6).Index(0))
	v.PersonPerformingStudyRegistration = NewXCNSlice(input.Index(7))
	v.StudyAuthorizingProvider = NewXCNSlice(input.Index(8))
	v.DateTimePatientStudyConsentSigned = NewDTM(input.Index(9).Index(0))
	v.PatientStudyEligibilityStatus = NewCWE(input.Index(10).Index(0))
	v.StudyRandomizationDateTime = NewDTMSlice(input.Index(11))
	v.RandomizedStudyArm = NewCWESlice(input.Index(12))
	v.StratumForStudyRandomization = NewCWESlice(input.Index(13))
	v.PatientEvaluabilityStatus = NewCWE(input.Index(14).Index(0))
	v.DateTimeEndedStudy = NewDTM(input.Index(15).Index(0))
	v.ReasonEndedStudy = NewCWE(input.Index(16).Index(0))
	return v
}

// NewCSRSlice will construct a slice of type CSR
func NewCSRSlice(input hl7.HL7Type) []CSR {
	values := make([]CSR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCSR(input.Index(i))
	}
	return values
}

type SDD struct {
	LotNumber       EI
	DeviceNumber    EI
	DeviceName      ST
	DeviceDataState CWE
	LoadStatus      CWE
	ControlCode     NM
	OperatorName    ST
}

// NewSDD creates an implementation of SDD
func NewSDD(input hl7.HL7Type) SDD {
	v := SDD{}
	v.LotNumber = NewEI(input.Index(1).Index(0))
	v.DeviceNumber = NewEI(input.Index(2).Index(0))
	v.DeviceName = NewST(input.Index(3).Index(0))
	v.DeviceDataState = NewCWE(input.Index(4).Index(0))
	v.LoadStatus = NewCWE(input.Index(5).Index(0))
	v.ControlCode = NewNM(input.Index(6).Index(0))
	v.OperatorName = NewST(input.Index(7).Index(0))
	return v
}

// NewSDDSlice will construct a slice of type SDD
func NewSDDSlice(input hl7.HL7Type) []SDD {
	values := make([]SDD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSDD(input.Index(i))
	}
	return values
}

type CNS struct {
	StartingNotificationReferenceNumber NM
	EndingNotificationReferenceNumber   NM
	StartingNotificationDateTime        DTM
	EndingNotificationDateTime          DTM
	StartingNotificationCode            CWE
	EndingNotificationCode              CWE
}

// NewCNS creates an implementation of CNS
func NewCNS(input hl7.HL7Type) CNS {
	v := CNS{}
	v.StartingNotificationReferenceNumber = NewNM(input.Index(1).Index(0))
	v.EndingNotificationReferenceNumber = NewNM(input.Index(2).Index(0))
	v.StartingNotificationDateTime = NewDTM(input.Index(3).Index(0))
	v.EndingNotificationDateTime = NewDTM(input.Index(4).Index(0))
	v.StartingNotificationCode = NewCWE(input.Index(5).Index(0))
	v.EndingNotificationCode = NewCWE(input.Index(6).Index(0))
	return v
}

// NewCNSSlice will construct a slice of type CNS
func NewCNSSlice(input hl7.HL7Type) []CNS {
	values := make([]CNS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCNS(input.Index(i))
	}
	return values
}

type PCE struct {
	SetID                   SI
	CostCenterAccountNumber CX
	TransactionCode         CWE
	TransactionAmountUnit   CP
}

// NewPCE creates an implementation of PCE
func NewPCE(input hl7.HL7Type) PCE {
	v := PCE{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.CostCenterAccountNumber = NewCX(input.Index(2).Index(0))
	v.TransactionCode = NewCWE(input.Index(3).Index(0))
	v.TransactionAmountUnit = NewCP(input.Index(4).Index(0))
	return v
}

// NewPCESlice will construct a slice of type PCE
func NewPCESlice(input hl7.HL7Type) []PCE {
	values := make([]PCE, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPCE(input.Index(i))
	}
	return values
}

type ZL7 struct {
	ZL71 ST
}

// NewZL7 creates an implementation of ZL7
func NewZL7(input hl7.HL7Type) ZL7 {
	v := ZL7{}
	v.ZL71 = NewST(input.Index(1).Index(0))
	return v
}

// NewZL7Slice will construct a slice of type ZL7
func NewZL7Slice(input hl7.HL7Type) []ZL7 {
	values := make([]ZL7, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewZL7(input.Index(i))
	}
	return values
}

type BLG struct {
	WhenToCharge     CCD
	ChargeType       ID
	AccountID        CX
	ChargeTypeReason CWE
}

// NewBLG creates an implementation of BLG
func NewBLG(input hl7.HL7Type) BLG {
	v := BLG{}
	v.WhenToCharge = NewCCD(input.Index(1).Index(0))
	v.ChargeType = NewID(input.Index(2).Index(0))
	v.AccountID = NewCX(input.Index(3).Index(0))
	v.ChargeTypeReason = NewCWE(input.Index(4).Index(0))
	return v
}

// NewBLGSlice will construct a slice of type BLG
func NewBLGSlice(input hl7.HL7Type) []BLG {
	values := make([]BLG, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBLG(input.Index(i))
	}
	return values
}

type PES struct {
	SenderOrganizationName []XON
	SenderIndividualName   []XCN
	SenderAddress          []XAD
	SenderTelephone        []XTN
	SenderEventIdentifier  EI
	SenderSequenceNumber   NM
	SenderEventDescription []FT
	SenderComment          FT
	SenderAwareDateTime    DTM
	EventReportDate        DTM
	EventReportTimingType  []ID
	EventReportSource      ID
	EventReportedTo        []ID
}

// NewPES creates an implementation of PES
func NewPES(input hl7.HL7Type) PES {
	v := PES{}
	v.SenderOrganizationName = NewXONSlice(input.Index(1))
	v.SenderIndividualName = NewXCNSlice(input.Index(2))
	v.SenderAddress = NewXADSlice(input.Index(3))
	v.SenderTelephone = NewXTNSlice(input.Index(4))
	v.SenderEventIdentifier = NewEI(input.Index(5).Index(0))
	v.SenderSequenceNumber = NewNM(input.Index(6).Index(0))
	v.SenderEventDescription = NewFTSlice(input.Index(7))
	v.SenderComment = NewFT(input.Index(8).Index(0))
	v.SenderAwareDateTime = NewDTM(input.Index(9).Index(0))
	v.EventReportDate = NewDTM(input.Index(10).Index(0))
	v.EventReportTimingType = NewIDSlice(input.Index(11))
	v.EventReportSource = NewID(input.Index(12).Index(0))
	v.EventReportedTo = NewIDSlice(input.Index(13))
	return v
}

// NewPESSlice will construct a slice of type PES
func NewPESSlice(input hl7.HL7Type) []PES {
	values := make([]PES, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPES(input.Index(i))
	}
	return values
}

type IVC struct {
	ProviderInvoiceNumber                  EI
	PayerInvoiceNumber                     EI
	ContractAgreementNumber                EI
	InvoiceControl                         CWE
	InvoiceReason                          CWE
	InvoiceType                            CWE
	InvoiceDateTime                        DTM
	InvoiceAmount                          CP
	PaymentTerms                           ST
	ProviderOrganization                   XON
	PayerOrganization                      XON
	Attention                              XCN
	LastInvoiceIndicator                   ID
	InvoiceBookingPeriod                   DTM
	Origin                                 ST
	InvoiceFixedAmount                     CP
	SpecialCosts                           CP
	AmountForDoctorsTreatment              CP
	ResponsiblePhysician                   XCN
	CostCenter                             CX
	InvoicePrepaidAmount                   CP
	TotalInvoiceAmountWithoutPrepaidAmount CP
	TotalAmountOfVat                       CP
	VatRatesApplied                        []NM
	BenefitGroup                           CWE
	ProviderTaxID                          ST
	PayerTaxID                             ST
	ProviderTaxStatus                      CWE
	PayerTaxStatus                         CWE
	SalesTaxID                             ST
}

// NewIVC creates an implementation of IVC
func NewIVC(input hl7.HL7Type) IVC {
	v := IVC{}
	v.ProviderInvoiceNumber = NewEI(input.Index(1).Index(0))
	v.PayerInvoiceNumber = NewEI(input.Index(2).Index(0))
	v.ContractAgreementNumber = NewEI(input.Index(3).Index(0))
	v.InvoiceControl = NewCWE(input.Index(4).Index(0))
	v.InvoiceReason = NewCWE(input.Index(5).Index(0))
	v.InvoiceType = NewCWE(input.Index(6).Index(0))
	v.InvoiceDateTime = NewDTM(input.Index(7).Index(0))
	v.InvoiceAmount = NewCP(input.Index(8).Index(0))
	v.PaymentTerms = NewST(input.Index(9).Index(0))
	v.ProviderOrganization = NewXON(input.Index(10).Index(0))
	v.PayerOrganization = NewXON(input.Index(11).Index(0))
	v.Attention = NewXCN(input.Index(12).Index(0))
	v.LastInvoiceIndicator = NewID(input.Index(13).Index(0))
	v.InvoiceBookingPeriod = NewDTM(input.Index(14).Index(0))
	v.Origin = NewST(input.Index(15).Index(0))
	v.InvoiceFixedAmount = NewCP(input.Index(16).Index(0))
	v.SpecialCosts = NewCP(input.Index(17).Index(0))
	v.AmountForDoctorsTreatment = NewCP(input.Index(18).Index(0))
	v.ResponsiblePhysician = NewXCN(input.Index(19).Index(0))
	v.CostCenter = NewCX(input.Index(20).Index(0))
	v.InvoicePrepaidAmount = NewCP(input.Index(21).Index(0))
	v.TotalInvoiceAmountWithoutPrepaidAmount = NewCP(input.Index(22).Index(0))
	v.TotalAmountOfVat = NewCP(input.Index(23).Index(0))
	v.VatRatesApplied = NewNMSlice(input.Index(24))
	v.BenefitGroup = NewCWE(input.Index(25).Index(0))
	v.ProviderTaxID = NewST(input.Index(26).Index(0))
	v.PayerTaxID = NewST(input.Index(27).Index(0))
	v.ProviderTaxStatus = NewCWE(input.Index(28).Index(0))
	v.PayerTaxStatus = NewCWE(input.Index(29).Index(0))
	v.SalesTaxID = NewST(input.Index(30).Index(0))
	return v
}

// NewIVCSlice will construct a slice of type IVC
func NewIVCSlice(input hl7.HL7Type) []IVC {
	values := make([]IVC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewIVC(input.Index(i))
	}
	return values
}

type RXR struct {
	Route                      CWE
	AdministrationSite         CWE
	AdministrationDevice       CWE
	AdministrationMethod       CWE
	RoutingInstruction         CWE
	AdministrationSiteModifier CWE
}

// NewRXR creates an implementation of RXR
func NewRXR(input hl7.HL7Type) RXR {
	v := RXR{}
	v.Route = NewCWE(input.Index(1).Index(0))
	v.AdministrationSite = NewCWE(input.Index(2).Index(0))
	v.AdministrationDevice = NewCWE(input.Index(3).Index(0))
	v.AdministrationMethod = NewCWE(input.Index(4).Index(0))
	v.RoutingInstruction = NewCWE(input.Index(5).Index(0))
	v.AdministrationSiteModifier = NewCWE(input.Index(6).Index(0))
	return v
}

// NewRXRSlice will construct a slice of type RXR
func NewRXRSlice(input hl7.HL7Type) []RXR {
	values := make([]RXR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRXR(input.Index(i))
	}
	return values
}

type AIG struct {
	SetID                    SI
	SegmentActionCode        ID
	ResourceID               CWE
	ResourceType             CWE
	ResourceGroup            []CWE
	ResourceQuantity         NM
	ResourceQuantityUnits    CNE
	StartDateTime            DTM
	StartDateTimeOffset      NM
	StartDateTimeOffsetUnits CNE
	Duration                 NM
	DurationUnits            CNE
	AllowSubstitutionCode    CWE
	FillerStatusCode         CWE
}

// NewAIG creates an implementation of AIG
func NewAIG(input hl7.HL7Type) AIG {
	v := AIG{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.SegmentActionCode = NewID(input.Index(2).Index(0))
	v.ResourceID = NewCWE(input.Index(3).Index(0))
	v.ResourceType = NewCWE(input.Index(4).Index(0))
	v.ResourceGroup = NewCWESlice(input.Index(5))
	v.ResourceQuantity = NewNM(input.Index(6).Index(0))
	v.ResourceQuantityUnits = NewCNE(input.Index(7).Index(0))
	v.StartDateTime = NewDTM(input.Index(8).Index(0))
	v.StartDateTimeOffset = NewNM(input.Index(9).Index(0))
	v.StartDateTimeOffsetUnits = NewCNE(input.Index(10).Index(0))
	v.Duration = NewNM(input.Index(11).Index(0))
	v.DurationUnits = NewCNE(input.Index(12).Index(0))
	v.AllowSubstitutionCode = NewCWE(input.Index(13).Index(0))
	v.FillerStatusCode = NewCWE(input.Index(14).Index(0))
	return v
}

// NewAIGSlice will construct a slice of type AIG
func NewAIGSlice(input hl7.HL7Type) []AIG {
	values := make([]AIG, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewAIG(input.Index(i))
	}
	return values
}

type SHP struct {
	ShipmentID                 EI
	InternalShipmentID         []EI
	ShipmentStatus             CWE
	ShipmentStatusDateTime     DTM
	ShipmentStatusReason       TX
	ShipmentPriority           CWE
	ShipmentConfidentiality    []CWE
	NumberOfPackagesInShipment NM
	ShipmentCondition          []CWE
	ShipmentHandlingCode       []CWE
	ShipmentRiskCode           []CWE
}

// NewSHP creates an implementation of SHP
func NewSHP(input hl7.HL7Type) SHP {
	v := SHP{}
	v.ShipmentID = NewEI(input.Index(1).Index(0))
	v.InternalShipmentID = NewEISlice(input.Index(2))
	v.ShipmentStatus = NewCWE(input.Index(3).Index(0))
	v.ShipmentStatusDateTime = NewDTM(input.Index(4).Index(0))
	v.ShipmentStatusReason = NewTX(input.Index(5).Index(0))
	v.ShipmentPriority = NewCWE(input.Index(6).Index(0))
	v.ShipmentConfidentiality = NewCWESlice(input.Index(7))
	v.NumberOfPackagesInShipment = NewNM(input.Index(8).Index(0))
	v.ShipmentCondition = NewCWESlice(input.Index(9))
	v.ShipmentHandlingCode = NewCWESlice(input.Index(10))
	v.ShipmentRiskCode = NewCWESlice(input.Index(11))
	return v
}

// NewSHPSlice will construct a slice of type SHP
func NewSHPSlice(input hl7.HL7Type) []SHP {
	values := make([]SHP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSHP(input.Index(i))
	}
	return values
}

type LRL struct {
	PrimaryKeyValueLrl                      PL
	SegmentActionCode                       ID
	SegmentUniqueKey                        EI
	LocationRelationshipID                  CWE
	OrganizationalLocationRelationshipValue []XON
	PatientLocationRelationshipValue        PL
}

// NewLRL creates an implementation of LRL
func NewLRL(input hl7.HL7Type) LRL {
	v := LRL{}
	v.PrimaryKeyValueLrl = NewPL(input.Index(1).Index(0))
	v.SegmentActionCode = NewID(input.Index(2).Index(0))
	v.SegmentUniqueKey = NewEI(input.Index(3).Index(0))
	v.LocationRelationshipID = NewCWE(input.Index(4).Index(0))
	v.OrganizationalLocationRelationshipValue = NewXONSlice(input.Index(5))
	v.PatientLocationRelationshipValue = NewPL(input.Index(6).Index(0))
	return v
}

// NewLRLSlice will construct a slice of type LRL
func NewLRLSlice(input hl7.HL7Type) []LRL {
	values := make([]LRL, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewLRL(input.Index(i))
	}
	return values
}

type ERR struct {
	ErrorCodeAndLocation      ST
	ErrorLocation             []ERL
	Hl7ErrorCode              CWE
	Severity                  ID
	ApplicationErrorCode      CWE
	ApplicationErrorParameter []ST
	DiagnosticInformation     TX
	UserMessage               TX
	InformPersonIndicator     []CWE
	OverrideType              CWE
	OverrideReasonCode        []CWE
	HelpDeskContactPoint      []XTN
}

// NewERR creates an implementation of ERR
func NewERR(input hl7.HL7Type) ERR {
	v := ERR{}
	v.ErrorCodeAndLocation = NewST(input.Index(1).Index(0))
	v.ErrorLocation = NewERLSlice(input.Index(2))
	v.Hl7ErrorCode = NewCWE(input.Index(3).Index(0))
	v.Severity = NewID(input.Index(4).Index(0))
	v.ApplicationErrorCode = NewCWE(input.Index(5).Index(0))
	v.ApplicationErrorParameter = NewSTSlice(input.Index(6))
	v.DiagnosticInformation = NewTX(input.Index(7).Index(0))
	v.UserMessage = NewTX(input.Index(8).Index(0))
	v.InformPersonIndicator = NewCWESlice(input.Index(9))
	v.OverrideType = NewCWE(input.Index(10).Index(0))
	v.OverrideReasonCode = NewCWESlice(input.Index(11))
	v.HelpDeskContactPoint = NewXTNSlice(input.Index(12))
	return v
}

// NewERRSlice will construct a slice of type ERR
func NewERRSlice(input hl7.HL7Type) []ERR {
	values := make([]ERR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewERR(input.Index(i))
	}
	return values
}

type AIS struct {
	SetID                                SI
	SegmentActionCode                    ID
	UniversalServiceIdentifier           CWE
	StartDateTime                        DTM
	StartDateTimeOffset                  NM
	StartDateTimeOffsetUnits             CNE
	Duration                             NM
	DurationUnits                        CNE
	AllowSubstitutionCode                CWE
	FillerStatusCode                     CWE
	PlacerSupplementalServiceInformation []CWE
	FillerSupplementalServiceInformation []CWE
}

// NewAIS creates an implementation of AIS
func NewAIS(input hl7.HL7Type) AIS {
	v := AIS{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.SegmentActionCode = NewID(input.Index(2).Index(0))
	v.UniversalServiceIdentifier = NewCWE(input.Index(3).Index(0))
	v.StartDateTime = NewDTM(input.Index(4).Index(0))
	v.StartDateTimeOffset = NewNM(input.Index(5).Index(0))
	v.StartDateTimeOffsetUnits = NewCNE(input.Index(6).Index(0))
	v.Duration = NewNM(input.Index(7).Index(0))
	v.DurationUnits = NewCNE(input.Index(8).Index(0))
	v.AllowSubstitutionCode = NewCWE(input.Index(9).Index(0))
	v.FillerStatusCode = NewCWE(input.Index(10).Index(0))
	v.PlacerSupplementalServiceInformation = NewCWESlice(input.Index(11))
	v.FillerSupplementalServiceInformation = NewCWESlice(input.Index(12))
	return v
}

// NewAISSlice will construct a slice of type AIS
func NewAISSlice(input hl7.HL7Type) []AIS {
	values := make([]AIS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewAIS(input.Index(i))
	}
	return values
}

type AUT struct {
	AuthorizingPayorPlanID            CWE
	AuthorizingPayorCompanyID         CWE
	AuthorizingPayorCompanyName       ST
	AuthorizationEffectiveDate        DTM
	AuthorizationExpirationDate       DTM
	AuthorizationIdentifier           EI
	ReimbursementLimit                CP
	RequestedNumberOfTreatments       CQ
	AuthorizedNumberOfTreatments      CQ
	ProcessDate                       DTM
	RequestedDisciplines              []CWE
	AuthorizedDisciplines             []CWE
	AuthorizationReferralType         CWE
	ApprovalStatus                    CWE
	PlannedTreatmentStopDate          DTM
	ClinicalService                   CWE
	ReasonText                        ST
	NumberOfAuthorizedTreatmentsUnits CQ
	NumberOfUsedTreatmentsUnits       CQ
	NumberOfScheduleTreatmentsUnits   CQ
	EncounterType                     CWE
	RemainingBenefitAmount            MO
	AuthorizedProvider                XON
	AuthorizedHealthProfessional      XCN
	SourceText                        ST
	SourceDate                        DTM
	SourcePhone                       XTN
	Comment                           ST
	ActionCode                        ID
}

// NewAUT creates an implementation of AUT
func NewAUT(input hl7.HL7Type) AUT {
	v := AUT{}
	v.AuthorizingPayorPlanID = NewCWE(input.Index(1).Index(0))
	v.AuthorizingPayorCompanyID = NewCWE(input.Index(2).Index(0))
	v.AuthorizingPayorCompanyName = NewST(input.Index(3).Index(0))
	v.AuthorizationEffectiveDate = NewDTM(input.Index(4).Index(0))
	v.AuthorizationExpirationDate = NewDTM(input.Index(5).Index(0))
	v.AuthorizationIdentifier = NewEI(input.Index(6).Index(0))
	v.ReimbursementLimit = NewCP(input.Index(7).Index(0))
	v.RequestedNumberOfTreatments = NewCQ(input.Index(8).Index(0))
	v.AuthorizedNumberOfTreatments = NewCQ(input.Index(9).Index(0))
	v.ProcessDate = NewDTM(input.Index(10).Index(0))
	v.RequestedDisciplines = NewCWESlice(input.Index(11))
	v.AuthorizedDisciplines = NewCWESlice(input.Index(12))
	v.AuthorizationReferralType = NewCWE(input.Index(13).Index(0))
	v.ApprovalStatus = NewCWE(input.Index(14).Index(0))
	v.PlannedTreatmentStopDate = NewDTM(input.Index(15).Index(0))
	v.ClinicalService = NewCWE(input.Index(16).Index(0))
	v.ReasonText = NewST(input.Index(17).Index(0))
	v.NumberOfAuthorizedTreatmentsUnits = NewCQ(input.Index(18).Index(0))
	v.NumberOfUsedTreatmentsUnits = NewCQ(input.Index(19).Index(0))
	v.NumberOfScheduleTreatmentsUnits = NewCQ(input.Index(20).Index(0))
	v.EncounterType = NewCWE(input.Index(21).Index(0))
	v.RemainingBenefitAmount = NewMO(input.Index(22).Index(0))
	v.AuthorizedProvider = NewXON(input.Index(23).Index(0))
	v.AuthorizedHealthProfessional = NewXCN(input.Index(24).Index(0))
	v.SourceText = NewST(input.Index(25).Index(0))
	v.SourceDate = NewDTM(input.Index(26).Index(0))
	v.SourcePhone = NewXTN(input.Index(27).Index(0))
	v.Comment = NewST(input.Index(28).Index(0))
	v.ActionCode = NewID(input.Index(29).Index(0))
	return v
}

// NewAUTSlice will construct a slice of type AUT
func NewAUTSlice(input hl7.HL7Type) []AUT {
	values := make([]AUT, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewAUT(input.Index(i))
	}
	return values
}

type NPU struct {
	BedLocation PL
	BedStatus   CWE
}

// NewNPU creates an implementation of NPU
func NewNPU(input hl7.HL7Type) NPU {
	v := NPU{}
	v.BedLocation = NewPL(input.Index(1).Index(0))
	v.BedStatus = NewCWE(input.Index(2).Index(0))
	return v
}

// NewNPUSlice will construct a slice of type NPU
func NewNPUSlice(input hl7.HL7Type) []NPU {
	values := make([]NPU, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewNPU(input.Index(i))
	}
	return values
}

type NCK struct {
	SystemDateTime DTM
}

// NewNCK creates an implementation of NCK
func NewNCK(input hl7.HL7Type) NCK {
	v := NCK{}
	v.SystemDateTime = NewDTM(input.Index(1).Index(0))
	return v
}

// NewNCKSlice will construct a slice of type NCK
func NewNCKSlice(input hl7.HL7Type) []NCK {
	values := make([]NCK, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewNCK(input.Index(i))
	}
	return values
}

type QPD struct {
	MessageQueryName                 CWE
	QueryTag                         ST
	UserParametersinSuccessiveFields Varies
}

// NewQPD creates an implementation of QPD
func NewQPD(input hl7.HL7Type) QPD {
	v := QPD{}
	v.MessageQueryName = NewCWE(input.Index(1).Index(0))
	v.QueryTag = NewST(input.Index(2).Index(0))
	v.UserParametersinSuccessiveFields = NewVaries(input.Index(3).Index(0))
	return v
}

// NewQPDSlice will construct a slice of type QPD
func NewQPDSlice(input hl7.HL7Type) []QPD {
	values := make([]QPD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQPD(input.Index(i))
	}
	return values
}

type OM3 struct {
	SequenceNumberTestObservationMasterFile     NM
	PreferredCodingSystem                       CWE
	ValidCodedanswers                           []CWE
	NormalTextCodesForCategoricalObservations   []CWE
	AbnormalTextCodesForCategoricalObservations []CWE
	CriticalTextCodesForCategoricalObservations []CWE
	ValueType                                   ID
}

// NewOM3 creates an implementation of OM3
func NewOM3(input hl7.HL7Type) OM3 {
	v := OM3{}
	v.SequenceNumberTestObservationMasterFile = NewNM(input.Index(1).Index(0))
	v.PreferredCodingSystem = NewCWE(input.Index(2).Index(0))
	v.ValidCodedanswers = NewCWESlice(input.Index(3))
	v.NormalTextCodesForCategoricalObservations = NewCWESlice(input.Index(4))
	v.AbnormalTextCodesForCategoricalObservations = NewCWESlice(input.Index(5))
	v.CriticalTextCodesForCategoricalObservations = NewCWESlice(input.Index(6))
	v.ValueType = NewID(input.Index(7).Index(0))
	return v
}

// NewOM3Slice will construct a slice of type OM3
func NewOM3Slice(input hl7.HL7Type) []OM3 {
	values := make([]OM3, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOM3(input.Index(i))
	}
	return values
}

type PV1 struct {
	SetID                     SI
	PatientClass              CWE
	AssignedPatientLocation   PL
	AdmissionType             CWE
	PreadmitNumber            CX
	PriorPatientLocation      PL
	AttendingDoctor           []XCN
	ReferringDoctor           []XCN
	ConsultingDoctor          []XCN
	HospitalService           CWE
	TemporaryLocation         PL
	PreadmitTestIndicator     CWE
	ReAdmissionIndicator      CWE
	AdmitSource               CWE
	AmbulatoryStatus          []CWE
	VipIndicator              CWE
	AdmittingDoctor           []XCN
	PatientType               CWE
	VisitNumber               CX
	FinancialClass            []FC
	ChargePriceIndicator      CWE
	CourtesyCode              CWE
	CreditRating              CWE
	ContractCode              []CWE
	ContractEffectiveDate     []DT
	ContractAmount            []NM
	ContractPeriod            []NM
	InterestCode              CWE
	TransferToBadDebtCode     CWE
	TransferToBadDebtDate     DT
	BadDebtAgencyCode         CWE
	BadDebtTransferAmount     NM
	BadDebtRecoveryAmount     NM
	DeleteAccountIndicator    CWE
	DeleteAccountDate         DT
	DischargeDisposition      CWE
	DischargedToLocation      DLD
	DietType                  CWE
	ServicingFacility         CWE
	BedStatus                 ST
	AccountStatus             CWE
	PendingLocation           PL
	PriorTemporaryLocation    PL
	AdmitDateTime             DTM
	DischargeDateTime         DTM
	CurrentPatientBalance     NM
	TotalCharges              NM
	TotalAdjustments          NM
	TotalPayments             NM
	AlternateVisitID          CX
	VisitIndicator            CWE
	OtherHealthcareProvider   ST
	ServiceEpisodeDescription ST
	ServiceEpisodeIdentifier  CX
}

// NewPV1 creates an implementation of PV1
func NewPV1(input hl7.HL7Type) PV1 {
	v := PV1{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.PatientClass = NewCWE(input.Index(2).Index(0))
	v.AssignedPatientLocation = NewPL(input.Index(3).Index(0))
	v.AdmissionType = NewCWE(input.Index(4).Index(0))
	v.PreadmitNumber = NewCX(input.Index(5).Index(0))
	v.PriorPatientLocation = NewPL(input.Index(6).Index(0))
	v.AttendingDoctor = NewXCNSlice(input.Index(7))
	v.ReferringDoctor = NewXCNSlice(input.Index(8))
	v.ConsultingDoctor = NewXCNSlice(input.Index(9))
	v.HospitalService = NewCWE(input.Index(10).Index(0))
	v.TemporaryLocation = NewPL(input.Index(11).Index(0))
	v.PreadmitTestIndicator = NewCWE(input.Index(12).Index(0))
	v.ReAdmissionIndicator = NewCWE(input.Index(13).Index(0))
	v.AdmitSource = NewCWE(input.Index(14).Index(0))
	v.AmbulatoryStatus = NewCWESlice(input.Index(15))
	v.VipIndicator = NewCWE(input.Index(16).Index(0))
	v.AdmittingDoctor = NewXCNSlice(input.Index(17))
	v.PatientType = NewCWE(input.Index(18).Index(0))
	v.VisitNumber = NewCX(input.Index(19).Index(0))
	v.FinancialClass = NewFCSlice(input.Index(20))
	v.ChargePriceIndicator = NewCWE(input.Index(21).Index(0))
	v.CourtesyCode = NewCWE(input.Index(22).Index(0))
	v.CreditRating = NewCWE(input.Index(23).Index(0))
	v.ContractCode = NewCWESlice(input.Index(24))
	v.ContractEffectiveDate = NewDTSlice(input.Index(25))
	v.ContractAmount = NewNMSlice(input.Index(26))
	v.ContractPeriod = NewNMSlice(input.Index(27))
	v.InterestCode = NewCWE(input.Index(28).Index(0))
	v.TransferToBadDebtCode = NewCWE(input.Index(29).Index(0))
	v.TransferToBadDebtDate = NewDT(input.Index(30).Index(0))
	v.BadDebtAgencyCode = NewCWE(input.Index(31).Index(0))
	v.BadDebtTransferAmount = NewNM(input.Index(32).Index(0))
	v.BadDebtRecoveryAmount = NewNM(input.Index(33).Index(0))
	v.DeleteAccountIndicator = NewCWE(input.Index(34).Index(0))
	v.DeleteAccountDate = NewDT(input.Index(35).Index(0))
	v.DischargeDisposition = NewCWE(input.Index(36).Index(0))
	v.DischargedToLocation = NewDLD(input.Index(37).Index(0))
	v.DietType = NewCWE(input.Index(38).Index(0))
	v.ServicingFacility = NewCWE(input.Index(39).Index(0))
	v.BedStatus = NewST(input.Index(40).Index(0))
	v.AccountStatus = NewCWE(input.Index(41).Index(0))
	v.PendingLocation = NewPL(input.Index(42).Index(0))
	v.PriorTemporaryLocation = NewPL(input.Index(43).Index(0))
	v.AdmitDateTime = NewDTM(input.Index(44).Index(0))
	v.DischargeDateTime = NewDTM(input.Index(45).Index(0))
	v.CurrentPatientBalance = NewNM(input.Index(46).Index(0))
	v.TotalCharges = NewNM(input.Index(47).Index(0))
	v.TotalAdjustments = NewNM(input.Index(48).Index(0))
	v.TotalPayments = NewNM(input.Index(49).Index(0))
	v.AlternateVisitID = NewCX(input.Index(50).Index(0))
	v.VisitIndicator = NewCWE(input.Index(51).Index(0))
	v.OtherHealthcareProvider = NewST(input.Index(52).Index(0))
	v.ServiceEpisodeDescription = NewST(input.Index(53).Index(0))
	v.ServiceEpisodeIdentifier = NewCX(input.Index(54).Index(0))
	return v
}

// NewPV1Slice will construct a slice of type PV1
func NewPV1Slice(input hl7.HL7Type) []PV1 {
	values := make([]PV1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPV1(input.Index(i))
	}
	return values
}

type DB1 struct {
	SetID                      SI
	DisabledPersonCode         CWE
	DisabledPersonIdentifier   []CX
	DisabilityIndicator        ID
	DisabilityStartDate        DT
	DisabilityEndDate          DT
	DisabilityReturnToWorkDate DT
	DisabilityUnableToWorkDate DT
}

// NewDB1 creates an implementation of DB1
func NewDB1(input hl7.HL7Type) DB1 {
	v := DB1{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.DisabledPersonCode = NewCWE(input.Index(2).Index(0))
	v.DisabledPersonIdentifier = NewCXSlice(input.Index(3))
	v.DisabilityIndicator = NewID(input.Index(4).Index(0))
	v.DisabilityStartDate = NewDT(input.Index(5).Index(0))
	v.DisabilityEndDate = NewDT(input.Index(6).Index(0))
	v.DisabilityReturnToWorkDate = NewDT(input.Index(7).Index(0))
	v.DisabilityUnableToWorkDate = NewDT(input.Index(8).Index(0))
	return v
}

// NewDB1Slice will construct a slice of type DB1
func NewDB1Slice(input hl7.HL7Type) []DB1 {
	values := make([]DB1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDB1(input.Index(i))
	}
	return values
}

type STZ struct {
	SterilizationType  CWE
	SterilizationCycle CWE
	MaintenanceCycle   CWE
	MaintenanceType    CWE
}

// NewSTZ creates an implementation of STZ
func NewSTZ(input hl7.HL7Type) STZ {
	v := STZ{}
	v.SterilizationType = NewCWE(input.Index(1).Index(0))
	v.SterilizationCycle = NewCWE(input.Index(2).Index(0))
	v.MaintenanceCycle = NewCWE(input.Index(3).Index(0))
	v.MaintenanceType = NewCWE(input.Index(4).Index(0))
	return v
}

// NewSTZSlice will construct a slice of type STZ
func NewSTZSlice(input hl7.HL7Type) []STZ {
	values := make([]STZ, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSTZ(input.Index(i))
	}
	return values
}

type RQD struct {
	RequisitionLineNumber    SI
	ItemCodeInternal         CWE
	ItemCodeExternal         CWE
	HospitalItemCode         CWE
	RequisitionQuantity      NM
	RequisitionUnitOfMeasure CWE
	CostCenterAccountNumber  CX
	ItemNaturalAccountCode   CWE
	DeliverToID              CWE
	DateNeeded               DT
}

// NewRQD creates an implementation of RQD
func NewRQD(input hl7.HL7Type) RQD {
	v := RQD{}
	v.RequisitionLineNumber = NewSI(input.Index(1).Index(0))
	v.ItemCodeInternal = NewCWE(input.Index(2).Index(0))
	v.ItemCodeExternal = NewCWE(input.Index(3).Index(0))
	v.HospitalItemCode = NewCWE(input.Index(4).Index(0))
	v.RequisitionQuantity = NewNM(input.Index(5).Index(0))
	v.RequisitionUnitOfMeasure = NewCWE(input.Index(6).Index(0))
	v.CostCenterAccountNumber = NewCX(input.Index(7).Index(0))
	v.ItemNaturalAccountCode = NewCWE(input.Index(8).Index(0))
	v.DeliverToID = NewCWE(input.Index(9).Index(0))
	v.DateNeeded = NewDT(input.Index(10).Index(0))
	return v
}

// NewRQDSlice will construct a slice of type RQD
func NewRQDSlice(input hl7.HL7Type) []RQD {
	values := make([]RQD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRQD(input.Index(i))
	}
	return values
}

type ILT struct {
	SetID                         SI
	InventoryLotNumber            ST
	InventoryExpirationDate       DTM
	InventoryReceivedDate         DTM
	InventoryReceivedQuantity     NM
	InventoryReceivedQuantityUnit CWE
	InventoryReceivedItemCost     MO
	InventoryOnHandDate           DTM
	InventoryOnHandQuantity       NM
	InventoryOnHandQuantityUnit   CWE
}

// NewILT creates an implementation of ILT
func NewILT(input hl7.HL7Type) ILT {
	v := ILT{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.InventoryLotNumber = NewST(input.Index(2).Index(0))
	v.InventoryExpirationDate = NewDTM(input.Index(3).Index(0))
	v.InventoryReceivedDate = NewDTM(input.Index(4).Index(0))
	v.InventoryReceivedQuantity = NewNM(input.Index(5).Index(0))
	v.InventoryReceivedQuantityUnit = NewCWE(input.Index(6).Index(0))
	v.InventoryReceivedItemCost = NewMO(input.Index(7).Index(0))
	v.InventoryOnHandDate = NewDTM(input.Index(8).Index(0))
	v.InventoryOnHandQuantity = NewNM(input.Index(9).Index(0))
	v.InventoryOnHandQuantityUnit = NewCWE(input.Index(10).Index(0))
	return v
}

// NewILTSlice will construct a slice of type ILT
func NewILTSlice(input hl7.HL7Type) []ILT {
	values := make([]ILT, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewILT(input.Index(i))
	}
	return values
}

type AIP struct {
	SetID                    SI
	SegmentActionCode        ID
	PersonnelResourceID      []XCN
	ResourceType             CWE
	ResourceGroup            CWE
	StartDateTime            DTM
	StartDateTimeOffset      NM
	StartDateTimeOffsetUnits CNE
	Duration                 NM
	DurationUnits            CNE
	AllowSubstitutionCode    CWE
	FillerStatusCode         CWE
}

// NewAIP creates an implementation of AIP
func NewAIP(input hl7.HL7Type) AIP {
	v := AIP{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.SegmentActionCode = NewID(input.Index(2).Index(0))
	v.PersonnelResourceID = NewXCNSlice(input.Index(3))
	v.ResourceType = NewCWE(input.Index(4).Index(0))
	v.ResourceGroup = NewCWE(input.Index(5).Index(0))
	v.StartDateTime = NewDTM(input.Index(6).Index(0))
	v.StartDateTimeOffset = NewNM(input.Index(7).Index(0))
	v.StartDateTimeOffsetUnits = NewCNE(input.Index(8).Index(0))
	v.Duration = NewNM(input.Index(9).Index(0))
	v.DurationUnits = NewCNE(input.Index(10).Index(0))
	v.AllowSubstitutionCode = NewCWE(input.Index(11).Index(0))
	v.FillerStatusCode = NewCWE(input.Index(12).Index(0))
	return v
}

// NewAIPSlice will construct a slice of type AIP
func NewAIPSlice(input hl7.HL7Type) []AIP {
	values := make([]AIP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewAIP(input.Index(i))
	}
	return values
}

type PSS struct {
	ProviderProductServiceSectionNumber EI
	PayerProductServiceSectionNumber    EI
	ProductServiceSectionSequenceNumber SI
	BilledAmount                        CP
	SectionDescriptionOrHeading         ST
}

// NewPSS creates an implementation of PSS
func NewPSS(input hl7.HL7Type) PSS {
	v := PSS{}
	v.ProviderProductServiceSectionNumber = NewEI(input.Index(1).Index(0))
	v.PayerProductServiceSectionNumber = NewEI(input.Index(2).Index(0))
	v.ProductServiceSectionSequenceNumber = NewSI(input.Index(3).Index(0))
	v.BilledAmount = NewCP(input.Index(4).Index(0))
	v.SectionDescriptionOrHeading = NewST(input.Index(5).Index(0))
	return v
}

// NewPSSSlice will construct a slice of type PSS
func NewPSSSlice(input hl7.HL7Type) []PSS {
	values := make([]PSS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPSS(input.Index(i))
	}
	return values
}

type ECD struct {
	ReferenceCommandNumber  NM
	RemoteControlCommand    CWE
	ResponseRequired        ID
	RequestedCompletionTime ST
	Parameters              []TX
}

// NewECD creates an implementation of ECD
func NewECD(input hl7.HL7Type) ECD {
	v := ECD{}
	v.ReferenceCommandNumber = NewNM(input.Index(1).Index(0))
	v.RemoteControlCommand = NewCWE(input.Index(2).Index(0))
	v.ResponseRequired = NewID(input.Index(3).Index(0))
	v.RequestedCompletionTime = NewST(input.Index(4).Index(0))
	v.Parameters = NewTXSlice(input.Index(5))
	return v
}

// NewECDSlice will construct a slice of type ECD
func NewECDSlice(input hl7.HL7Type) []ECD {
	values := make([]ECD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewECD(input.Index(i))
	}
	return values
}

type OVR struct {
	BusinessRuleOverrideType CWE
	BusinessRuleOverrideCode CWE
	OverrideComments         TX
	OverrideEnteredBy        XCN
	OverrideAuthorizedBy     XCN
}

// NewOVR creates an implementation of OVR
func NewOVR(input hl7.HL7Type) OVR {
	v := OVR{}
	v.BusinessRuleOverrideType = NewCWE(input.Index(1).Index(0))
	v.BusinessRuleOverrideCode = NewCWE(input.Index(2).Index(0))
	v.OverrideComments = NewTX(input.Index(3).Index(0))
	v.OverrideEnteredBy = NewXCN(input.Index(4).Index(0))
	v.OverrideAuthorizedBy = NewXCN(input.Index(5).Index(0))
	return v
}

// NewOVRSlice will construct a slice of type OVR
func NewOVRSlice(input hl7.HL7Type) []OVR {
	values := make([]OVR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOVR(input.Index(i))
	}
	return values
}

type PSL struct {
	ProviderProductServiceLineItemNumber     EI
	PayerProductServiceLineItemNumber        EI
	ProductServiceLineItemSequenceNumber     SI
	ProviderTrackingID                       EI
	PayerTrackingID                          EI
	ProductServiceLineItemStatus             CWE
	ProductServiceCode                       CWE
	ProductServiceCodeModifier               []CWE
	ProductServiceCodeDescription            ST
	ProductServiceEffectiveDate              DTM
	ProductServiceExpirationDate             DTM
	ProductServiceQuantity                   CQ
	ProductServiceUnitCost                   CP
	NumberOfItemsPerUnit                     NM
	ProductServiceGrossAmount                CP
	ProductServiceBilledAmount               CP
	ProductServiceClarificationCodeType      []CWE
	ProductServiceClarificationCodeValue     []ST
	HealthDocumentReferenceIdentifier        []EI
	ProcessingConsiderationCode              []CWE
	RestrictedDisclosureIndicator            ID
	RelatedProductServiceCodeIndicator       CWE
	ProductServiceAmountForPhysician         CP
	ProductServiceCostFactor                 NM
	CostCenter                               CX
	BillingPeriod                            DR
	DaysWithoutBilling                       NM
	SessionNo                                NM
	ExecutingPhysicianID                     XCN
	ResponsiblePhysicianID                   XCN
	RoleExecutingPhysician                   CWE
	MedicalRoleExecutingPhysician            CWE
	SideOfBody                               CWE
	NumberOfTpsPp                            NM
	TpValuePp                                CP
	InternalScalingFactorPp                  NM
	ExternalScalingFactorPp                  NM
	AmountPp                                 CP
	NumberOfTpsTechnicalPart                 NM
	TpValueTechnicalPart                     CP
	InternalScalingFactorTechnicalPart       NM
	ExternalScalingFactorTechnicalPart       NM
	AmountTechnicalPart                      CP
	TotalAmountProfessionalPartTechnicalPart CP
	VatRate                                  NM
	MainService                              ID
	Validation                               ID
	Comment                                  ST
}

// NewPSL creates an implementation of PSL
func NewPSL(input hl7.HL7Type) PSL {
	v := PSL{}
	v.ProviderProductServiceLineItemNumber = NewEI(input.Index(1).Index(0))
	v.PayerProductServiceLineItemNumber = NewEI(input.Index(2).Index(0))
	v.ProductServiceLineItemSequenceNumber = NewSI(input.Index(3).Index(0))
	v.ProviderTrackingID = NewEI(input.Index(4).Index(0))
	v.PayerTrackingID = NewEI(input.Index(5).Index(0))
	v.ProductServiceLineItemStatus = NewCWE(input.Index(6).Index(0))
	v.ProductServiceCode = NewCWE(input.Index(7).Index(0))
	v.ProductServiceCodeModifier = NewCWESlice(input.Index(8))
	v.ProductServiceCodeDescription = NewST(input.Index(9).Index(0))
	v.ProductServiceEffectiveDate = NewDTM(input.Index(10).Index(0))
	v.ProductServiceExpirationDate = NewDTM(input.Index(11).Index(0))
	v.ProductServiceQuantity = NewCQ(input.Index(12).Index(0))
	v.ProductServiceUnitCost = NewCP(input.Index(13).Index(0))
	v.NumberOfItemsPerUnit = NewNM(input.Index(14).Index(0))
	v.ProductServiceGrossAmount = NewCP(input.Index(15).Index(0))
	v.ProductServiceBilledAmount = NewCP(input.Index(16).Index(0))
	v.ProductServiceClarificationCodeType = NewCWESlice(input.Index(17))
	v.ProductServiceClarificationCodeValue = NewSTSlice(input.Index(18))
	v.HealthDocumentReferenceIdentifier = NewEISlice(input.Index(19))
	v.ProcessingConsiderationCode = NewCWESlice(input.Index(20))
	v.RestrictedDisclosureIndicator = NewID(input.Index(21).Index(0))
	v.RelatedProductServiceCodeIndicator = NewCWE(input.Index(22).Index(0))
	v.ProductServiceAmountForPhysician = NewCP(input.Index(23).Index(0))
	v.ProductServiceCostFactor = NewNM(input.Index(24).Index(0))
	v.CostCenter = NewCX(input.Index(25).Index(0))
	v.BillingPeriod = NewDR(input.Index(26).Index(0))
	v.DaysWithoutBilling = NewNM(input.Index(27).Index(0))
	v.SessionNo = NewNM(input.Index(28).Index(0))
	v.ExecutingPhysicianID = NewXCN(input.Index(29).Index(0))
	v.ResponsiblePhysicianID = NewXCN(input.Index(30).Index(0))
	v.RoleExecutingPhysician = NewCWE(input.Index(31).Index(0))
	v.MedicalRoleExecutingPhysician = NewCWE(input.Index(32).Index(0))
	v.SideOfBody = NewCWE(input.Index(33).Index(0))
	v.NumberOfTpsPp = NewNM(input.Index(34).Index(0))
	v.TpValuePp = NewCP(input.Index(35).Index(0))
	v.InternalScalingFactorPp = NewNM(input.Index(36).Index(0))
	v.ExternalScalingFactorPp = NewNM(input.Index(37).Index(0))
	v.AmountPp = NewCP(input.Index(38).Index(0))
	v.NumberOfTpsTechnicalPart = NewNM(input.Index(39).Index(0))
	v.TpValueTechnicalPart = NewCP(input.Index(40).Index(0))
	v.InternalScalingFactorTechnicalPart = NewNM(input.Index(41).Index(0))
	v.ExternalScalingFactorTechnicalPart = NewNM(input.Index(42).Index(0))
	v.AmountTechnicalPart = NewCP(input.Index(43).Index(0))
	v.TotalAmountProfessionalPartTechnicalPart = NewCP(input.Index(44).Index(0))
	v.VatRate = NewNM(input.Index(45).Index(0))
	v.MainService = NewID(input.Index(46).Index(0))
	v.Validation = NewID(input.Index(47).Index(0))
	v.Comment = NewST(input.Index(48).Index(0))
	return v
}

// NewPSLSlice will construct a slice of type PSL
func NewPSLSlice(input hl7.HL7Type) []PSL {
	values := make([]PSL, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPSL(input.Index(i))
	}
	return values
}

type PRC struct {
	PrimaryKeyValuePrc  CWE
	FacilityIDPrc       []CWE
	Department          []CWE
	ValidPatientClasses []CWE
	Price               []CP
	Formula             []ST
	MinimumQuantity     NM
	MaximumQuantity     NM
	MinimumPrice        MO
	MaximumPrice        MO
	EffectiveStartDate  DTM
	EffectiveEndDate    DTM
	PriceOverrideFlag   CWE
	BillingCategory     []CWE
	ChargeableFlag      ID
	ActiveInactiveFlag  ID
	Cost                MO
	ChargeOnIndicator   CWE
}

// NewPRC creates an implementation of PRC
func NewPRC(input hl7.HL7Type) PRC {
	v := PRC{}
	v.PrimaryKeyValuePrc = NewCWE(input.Index(1).Index(0))
	v.FacilityIDPrc = NewCWESlice(input.Index(2))
	v.Department = NewCWESlice(input.Index(3))
	v.ValidPatientClasses = NewCWESlice(input.Index(4))
	v.Price = NewCPSlice(input.Index(5))
	v.Formula = NewSTSlice(input.Index(6))
	v.MinimumQuantity = NewNM(input.Index(7).Index(0))
	v.MaximumQuantity = NewNM(input.Index(8).Index(0))
	v.MinimumPrice = NewMO(input.Index(9).Index(0))
	v.MaximumPrice = NewMO(input.Index(10).Index(0))
	v.EffectiveStartDate = NewDTM(input.Index(11).Index(0))
	v.EffectiveEndDate = NewDTM(input.Index(12).Index(0))
	v.PriceOverrideFlag = NewCWE(input.Index(13).Index(0))
	v.BillingCategory = NewCWESlice(input.Index(14))
	v.ChargeableFlag = NewID(input.Index(15).Index(0))
	v.ActiveInactiveFlag = NewID(input.Index(16).Index(0))
	v.Cost = NewMO(input.Index(17).Index(0))
	v.ChargeOnIndicator = NewCWE(input.Index(18).Index(0))
	return v
}

// NewPRCSlice will construct a slice of type PRC
func NewPRCSlice(input hl7.HL7Type) []PRC {
	values := make([]PRC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPRC(input.Index(i))
	}
	return values
}

type QRF struct {
	QRF1 ST
}

// NewQRF creates an implementation of QRF
func NewQRF(input hl7.HL7Type) QRF {
	v := QRF{}
	v.QRF1 = NewST(input.Index(1).Index(0))
	return v
}

// NewQRFSlice will construct a slice of type QRF
func NewQRFSlice(input hl7.HL7Type) []QRF {
	values := make([]QRF, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQRF(input.Index(i))
	}
	return values
}

type OBX struct {
	SetID                                 SI
	ValueType                             ID
	ObservationIdentifier                 CWE
	ObservationSubID                      ST
	ObservationValue                      []Varies
	Units                                 CWE
	ReferencesRange                       ST
	InterpretationCodes                   []CWE
	Probability                           NM
	NatureOfAbnormalTest                  []ID
	ObservationResultStatus               ID
	EffectiveDateOfReferenceRange         DTM
	UserDefinedAccessChecks               ST
	DateTimeOfTheObservation              DTM
	ProducersID                           CWE
	ResponsibleObserver                   []XCN
	ObservationMethod                     []CWE
	EquipmentInstanceIdentifier           []EI
	DateTimeOfTheAnalysis                 DTM
	ObservationSite                       []CWE
	ObservationInstanceIdentifier         EI
	MoodCode                              CNE
	PerformingOrganizationName            XON
	PerformingOrganizationAddress         XAD
	PerformingOrganizationMedicalDirector XCN
	PatientResultsReleaseCategory         ID
	RootCause                             CWE
	LocalProcessControl                   []CWE
}

// NewOBX creates an implementation of OBX
func NewOBX(input hl7.HL7Type) OBX {
	v := OBX{}
	v.SetID = NewSI(input.Index(1).Index(0))
	v.ValueType = NewID(input.Index(2).Index(0))
	v.ObservationIdentifier = NewCWE(input.Index(3).Index(0))
	v.ObservationSubID = NewST(input.Index(4).Index(0))
	v.ObservationValue = NewVariesSlice(input.Index(5))
	v.Units = NewCWE(input.Index(6).Index(0))
	v.ReferencesRange = NewST(input.Index(7).Index(0))
	v.InterpretationCodes = NewCWESlice(input.Index(8))
	v.Probability = NewNM(input.Index(9).Index(0))
	v.NatureOfAbnormalTest = NewIDSlice(input.Index(10))
	v.ObservationResultStatus = NewID(input.Index(11).Index(0))
	v.EffectiveDateOfReferenceRange = NewDTM(input.Index(12).Index(0))
	v.UserDefinedAccessChecks = NewST(input.Index(13).Index(0))
	v.DateTimeOfTheObservation = NewDTM(input.Index(14).Index(0))
	v.ProducersID = NewCWE(input.Index(15).Index(0))
	v.ResponsibleObserver = NewXCNSlice(input.Index(16))
	v.ObservationMethod = NewCWESlice(input.Index(17))
	v.EquipmentInstanceIdentifier = NewEISlice(input.Index(18))
	v.DateTimeOfTheAnalysis = NewDTM(input.Index(19).Index(0))
	v.ObservationSite = NewCWESlice(input.Index(20))
	v.ObservationInstanceIdentifier = NewEI(input.Index(21).Index(0))
	v.MoodCode = NewCNE(input.Index(22).Index(0))
	v.PerformingOrganizationName = NewXON(input.Index(23).Index(0))
	v.PerformingOrganizationAddress = NewXAD(input.Index(24).Index(0))
	v.PerformingOrganizationMedicalDirector = NewXCN(input.Index(25).Index(0))
	v.PatientResultsReleaseCategory = NewID(input.Index(26).Index(0))
	v.RootCause = NewCWE(input.Index(27).Index(0))
	v.LocalProcessControl = NewCWESlice(input.Index(28))
	return v
}

// NewOBXSlice will construct a slice of type OBX
func NewOBXSlice(input hl7.HL7Type) []OBX {
	values := make([]OBX, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOBX(input.Index(i))
	}
	return values
}

type DRG struct {
	DiagnosticRelatedGroup     CNE
	DrgAssignedDateTime        DTM
	DrgApprovalIndicator       ID
	DrgGrouperReviewCode       CWE
	OutlierType                CWE
	OutlierDays                NM
	OutlierCost                CP
	DrgPayor                   CWE
	OutlierReimbursement       CP
	ConfidentialIndicator      ID
	DrgTransferType            CWE
	NameOfCoder                XPN
	GrouperStatus              CWE
	PcclValueCode              CWE
	EffectiveWeight            NM
	MonetaryAmount             MO
	StatusPatient              CWE
	GrouperSoftwareName        ST
	GrouperSoftwareVersion     ST
	StatusFinancialCalculation CWE
	RelativeDiscountSurcharge  MO
	BasicCharge                MO
	TotalCharge                MO
	DiscountSurcharge          MO
	CalculatedDays             NM
	StatusGender               CWE
	StatusAge                  CWE
	StatusLengthOfStay         CWE
	StatusSameDayFlag          CWE
	StatusSeparationMode       CWE
	StatusWeightAtBirth        CWE
	StatusRespirationMinutes   CWE
	StatusAdmission            CWE
}

// NewDRG creates an implementation of DRG
func NewDRG(input hl7.HL7Type) DRG {
	v := DRG{}
	v.DiagnosticRelatedGroup = NewCNE(input.Index(1).Index(0))
	v.DrgAssignedDateTime = NewDTM(input.Index(2).Index(0))
	v.DrgApprovalIndicator = NewID(input.Index(3).Index(0))
	v.DrgGrouperReviewCode = NewCWE(input.Index(4).Index(0))
	v.OutlierType = NewCWE(input.Index(5).Index(0))
	v.OutlierDays = NewNM(input.Index(6).Index(0))
	v.OutlierCost = NewCP(input.Index(7).Index(0))
	v.DrgPayor = NewCWE(input.Index(8).Index(0))
	v.OutlierReimbursement = NewCP(input.Index(9).Index(0))
	v.ConfidentialIndicator = NewID(input.Index(10).Index(0))
	v.DrgTransferType = NewCWE(input.Index(11).Index(0))
	v.NameOfCoder = NewXPN(input.Index(12).Index(0))
	v.GrouperStatus = NewCWE(input.Index(13).Index(0))
	v.PcclValueCode = NewCWE(input.Index(14).Index(0))
	v.EffectiveWeight = NewNM(input.Index(15).Index(0))
	v.MonetaryAmount = NewMO(input.Index(16).Index(0))
	v.StatusPatient = NewCWE(input.Index(17).Index(0))
	v.GrouperSoftwareName = NewST(input.Index(18).Index(0))
	v.GrouperSoftwareVersion = NewST(input.Index(19).Index(0))
	v.StatusFinancialCalculation = NewCWE(input.Index(20).Index(0))
	v.RelativeDiscountSurcharge = NewMO(input.Index(21).Index(0))
	v.BasicCharge = NewMO(input.Index(22).Index(0))
	v.TotalCharge = NewMO(input.Index(23).Index(0))
	v.DiscountSurcharge = NewMO(input.Index(24).Index(0))
	v.CalculatedDays = NewNM(input.Index(25).Index(0))
	v.StatusGender = NewCWE(input.Index(26).Index(0))
	v.StatusAge = NewCWE(input.Index(27).Index(0))
	v.StatusLengthOfStay = NewCWE(input.Index(28).Index(0))
	v.StatusSameDayFlag = NewCWE(input.Index(29).Index(0))
	v.StatusSeparationMode = NewCWE(input.Index(30).Index(0))
	v.StatusWeightAtBirth = NewCWE(input.Index(31).Index(0))
	v.StatusRespirationMinutes = NewCWE(input.Index(32).Index(0))
	v.StatusAdmission = NewCWE(input.Index(33).Index(0))
	return v
}

// NewDRGSlice will construct a slice of type DRG
func NewDRGSlice(input hl7.HL7Type) []DRG {
	values := make([]DRG, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDRG(input.Index(i))
	}
	return values
}

type SCD struct {
	CycleStartTime         TM
	CycleCount             NM
	TempMax                CQ
	TempMin                CQ
	LoadNumber             NM
	ConditionTime          CQ
	SterilizeTime          CQ
	ExhaustTime            CQ
	TotalCycleTime         CQ
	DeviceStatus           CWE
	CycleStartDateTime     DTM
	DryTime                CQ
	LeakRate               CQ
	ControlTemperature     CQ
	SterilizerTemperature  CQ
	CycleCompleteTime      TM
	UnderTemperature       CQ
	OverTemperature        CQ
	AbortCycle             CNE
	Alarm                  CNE
	LongInChargePhase      CNE
	LongInExhaustPhase     CNE
	LongInFastExhaustPhase CNE
	Reset                  CNE
	OperatorUnload         XCN
	DoorOpen               CNE
	ReadingFailure         CNE
	CycleType              CWE
	ThermalRinseTime       CQ
	WashTime               CQ
	InjectionRate          CQ
	ProcedureCode          CNE
	PatientIdentifierList  []CX
	AttendingDoctor        XCN
	DilutionFactor         SN
	FillTime               CQ
	InletTemperature       CQ
}

// NewSCD creates an implementation of SCD
func NewSCD(input hl7.HL7Type) SCD {
	v := SCD{}
	v.CycleStartTime = NewTM(input.Index(1).Index(0))
	v.CycleCount = NewNM(input.Index(2).Index(0))
	v.TempMax = NewCQ(input.Index(3).Index(0))
	v.TempMin = NewCQ(input.Index(4).Index(0))
	v.LoadNumber = NewNM(input.Index(5).Index(0))
	v.ConditionTime = NewCQ(input.Index(6).Index(0))
	v.SterilizeTime = NewCQ(input.Index(7).Index(0))
	v.ExhaustTime = NewCQ(input.Index(8).Index(0))
	v.TotalCycleTime = NewCQ(input.Index(9).Index(0))
	v.DeviceStatus = NewCWE(input.Index(10).Index(0))
	v.CycleStartDateTime = NewDTM(input.Index(11).Index(0))
	v.DryTime = NewCQ(input.Index(12).Index(0))
	v.LeakRate = NewCQ(input.Index(13).Index(0))
	v.ControlTemperature = NewCQ(input.Index(14).Index(0))
	v.SterilizerTemperature = NewCQ(input.Index(15).Index(0))
	v.CycleCompleteTime = NewTM(input.Index(16).Index(0))
	v.UnderTemperature = NewCQ(input.Index(17).Index(0))
	v.OverTemperature = NewCQ(input.Index(18).Index(0))
	v.AbortCycle = NewCNE(input.Index(19).Index(0))
	v.Alarm = NewCNE(input.Index(20).Index(0))
	v.LongInChargePhase = NewCNE(input.Index(21).Index(0))
	v.LongInExhaustPhase = NewCNE(input.Index(22).Index(0))
	v.LongInFastExhaustPhase = NewCNE(input.Index(23).Index(0))
	v.Reset = NewCNE(input.Index(24).Index(0))
	v.OperatorUnload = NewXCN(input.Index(25).Index(0))
	v.DoorOpen = NewCNE(input.Index(26).Index(0))
	v.ReadingFailure = NewCNE(input.Index(27).Index(0))
	v.CycleType = NewCWE(input.Index(28).Index(0))
	v.ThermalRinseTime = NewCQ(input.Index(29).Index(0))
	v.WashTime = NewCQ(input.Index(30).Index(0))
	v.InjectionRate = NewCQ(input.Index(31).Index(0))
	v.ProcedureCode = NewCNE(input.Index(32).Index(0))
	v.PatientIdentifierList = NewCXSlice(input.Index(33))
	v.AttendingDoctor = NewXCN(input.Index(34).Index(0))
	v.DilutionFactor = NewSN(input.Index(35).Index(0))
	v.FillTime = NewCQ(input.Index(36).Index(0))
	v.InletTemperature = NewCQ(input.Index(37).Index(0))
	return v
}

// NewSCDSlice will construct a slice of type SCD
func NewSCDSlice(input hl7.HL7Type) []SCD {
	values := make([]SCD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSCD(input.Index(i))
	}
	return values
}
