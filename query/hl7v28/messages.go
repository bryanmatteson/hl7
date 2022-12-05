package hl7v28

import hl7 "hl7"

// NewQSBZ83 creates an implementation of QSBZ83
func NewQSBZ83(input hl7.HL7Type) QSBZ83 {
	v := QSBZ83{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.ResponseControlParameter = NewRCP(input.Index(4))
	v.ContinuationPointer = NewDSC(input.Index(5))
	return v
}

type QSBZ83 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQSBZ83Slice will construct a slice of type QSBZ83
func NewQSBZ83Slice(input hl7.HL7Type) []QSBZ83 {
	values := make([]QSBZ83, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQSBZ83(input.Index(i))
	}
	return values
}

// NewCRMC07 creates an implementation of CRMC07
func NewCRMC07(input hl7.HL7Type) CRMC07 {
	v := CRMC07{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Patient = NewCRMC07PatientSlice(input.Index(3))
	return v
}

type CRMC07 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Patient                       []CRMC07Patient
}

// NewCRMC07Slice will construct a slice of type CRMC07
func NewCRMC07Slice(input hl7.HL7Type) []CRMC07 {
	values := make([]CRMC07, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCRMC07(input.Index(i))
	}
	return values
}

// NewORUR40 creates an implementation of ORUR40
func NewORUR40(input hl7.HL7Type) ORUR40 {
	v := ORUR40{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientResult = NewORUR40PatientResultSlice(input.Index(3))
	v.ContinuationPointer = NewDSC(input.Index(4))
	return v
}

type ORUR40 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientResult                 []ORUR40PatientResult
	ContinuationPointer           DSC
}

// NewORUR40Slice will construct a slice of type ORUR40
func NewORUR40Slice(input hl7.HL7Type) []ORUR40 {
	values := make([]ORUR40, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORUR40(input.Index(i))
	}
	return values
}

// NewSIUS21 creates an implementation of SIUS21
func NewSIUS21(input hl7.HL7Type) SIUS21 {
	v := SIUS21{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SchedulingActivityInformation = NewSCH(input.Index(1))
	v.TimingQuantity = NewTQ1Slice(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSIUS21PatientSlice(input.Index(4))
	v.Resources = NewSIUS21ResourcesSlice(input.Index(5))
	return v
}

type SIUS21 struct {
	MessageHeader                 MSH
	SchedulingActivityInformation SCH
	TimingQuantity                []TQ1
	NotesAndComments              []NTE
	Patient                       []SIUS21Patient
	Resources                     []SIUS21Resources
}

// NewSIUS21Slice will construct a slice of type SIUS21
func NewSIUS21Slice(input hl7.HL7Type) []SIUS21 {
	values := make([]SIUS21, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSIUS21(input.Index(i))
	}
	return values
}

// NewBRPO30 creates an implementation of BRPO30
func NewBRPO30(input hl7.HL7Type) BRPO30 {
	v := BRPO30{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewBRPO30Response(input.Index(6))
	return v
}

type BRPO30 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      BRPO30Response
}

// NewBRPO30Slice will construct a slice of type BRPO30
func NewBRPO30Slice(input hl7.HL7Type) []BRPO30 {
	values := make([]BRPO30, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBRPO30(input.Index(i))
	}
	return values
}

// NewADTA11 creates an implementation of ADTA11
func NewADTA11(input hl7.HL7Type) ADTA11 {
	v := ADTA11{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.Disability = NewDB1Slice(input.Index(8))
	v.ObservationResult = NewOBXSlice(input.Index(9))
	return v
}

type ADTA11 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	Disability                        []DB1
	ObservationResult                 []OBX
}

// NewADTA11Slice will construct a slice of type ADTA11
func NewADTA11Slice(input hl7.HL7Type) []ADTA11 {
	values := make([]ADTA11, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA11(input.Index(i))
	}
	return values
}

// NewADTA43 creates an implementation of ADTA43
func NewADTA43(input hl7.HL7Type) ADTA43 {
	v := ADTA43{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.Patient = NewADTA43PatientSlice(input.Index(4))
	return v
}

type ADTA43 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	Patient                       []ADTA43Patient
}

// NewADTA43Slice will construct a slice of type ADTA43
func NewADTA43Slice(input hl7.HL7Type) []ADTA43 {
	values := make([]ADTA43, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA43(input.Index(i))
	}
	return values
}

// NewOMIO23 creates an implementation of OMIO23
func NewOMIO23(input hl7.HL7Type) OMIO23 {
	v := OMIO23{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewOMIO23Patient(input.Index(4))
	v.Order = NewOMIO23OrderSlice(input.Index(5))
	return v
}

type OMIO23 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       OMIO23Patient
	Order                         []OMIO23Order
}

// NewOMIO23Slice will construct a slice of type OMIO23
func NewOMIO23Slice(input hl7.HL7Type) []OMIO23 {
	values := make([]OMIO23, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOMIO23(input.Index(i))
	}
	return values
}

// NewRRII14 creates an implementation of RRII14
func NewRRII14(input hl7.HL7Type) RRII14 {
	v := RRII14{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.ReferralInformation = NewRF1(input.Index(4))
	v.AuthorizationContact = NewRRII14AuthorizationContact(input.Index(5))
	v.ProviderContact = NewRRII14ProviderContactSlice(input.Index(6))
	v.PatientIdentification = NewPID(input.Index(7))
	v.Accident = NewACC(input.Index(8))
	v.Diagnosis = NewDG1Slice(input.Index(9))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(10))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(11))
	v.Procedure = NewRRII14ProcedureSlice(input.Index(12))
	v.Observation = NewRRII14ObservationSlice(input.Index(13))
	v.PatientVisit = NewRRII14PatientVisit(input.Index(14))
	v.NotesAndComments = NewNTESlice(input.Index(15))
	return v
}

type RRII14 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	ReferralInformation           RF1
	AuthorizationContact          RRII14AuthorizationContact
	ProviderContact               []RRII14ProviderContact
	PatientIdentification         PID
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []RRII14Procedure
	Observation                   []RRII14Observation
	PatientVisit                  RRII14PatientVisit
	NotesAndComments              []NTE
}

// NewRRII14Slice will construct a slice of type RRII14
func NewRRII14Slice(input hl7.HL7Type) []RRII14 {
	values := make([]RRII14, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRRII14(input.Index(i))
	}
	return values
}

// NewRSPZ90 creates an implementation of RSPZ90
func NewRSPZ90(input hl7.HL7Type) RSPZ90 {
	v := RSPZ90{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.ResponseControlParameter = NewRCP(input.Index(7))
	v.QueryResponse = NewRSPZ90QueryResponseSlice(input.Index(8))
	v.ContinuationPointer = NewDSC(input.Index(9))
	return v
}

type RSPZ90 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
	QueryResponse                 []RSPZ90QueryResponse
	ContinuationPointer           DSC
}

// NewRSPZ90Slice will construct a slice of type RSPZ90
func NewRSPZ90Slice(input hl7.HL7Type) []RSPZ90 {
	values := make([]RSPZ90, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPZ90(input.Index(i))
	}
	return values
}

// NewQBPQ31 creates an implementation of QBPQ31
func NewQBPQ31(input hl7.HL7Type) QBPQ31 {
	v := QBPQ31{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.Qbp = NewQBPQ31Qbp(input.Index(4))
	v.ResponseControlParameter = NewRCP(input.Index(5))
	v.ContinuationPointer = NewDSC(input.Index(6))
	return v
}

type QBPQ31 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	Qbp                           QBPQ31Qbp
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPQ31Slice will construct a slice of type QBPQ31
func NewQBPQ31Slice(input hl7.HL7Type) []QBPQ31 {
	values := make([]QBPQ31, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPQ31(input.Index(i))
	}
	return values
}

// NewEHCE10 creates an implementation of EHCE10
func NewEHCE10(input hl7.HL7Type) EHCE10 {
	v := EHCE10{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUACSlice(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.InvoiceProcessingResultsInfo = NewEHCE10InvoiceProcessingResultsInfoSlice(input.Index(5))
	return v
}

type EHCE10 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials []UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	InvoiceProcessingResultsInfo  []EHCE10InvoiceProcessingResultsInfo
}

// NewEHCE10Slice will construct a slice of type EHCE10
func NewEHCE10Slice(input hl7.HL7Type) []EHCE10 {
	values := make([]EHCE10, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEHCE10(input.Index(i))
	}
	return values
}

// NewPMUB04 creates an implementation of PMUB04
func NewPMUB04(input hl7.HL7Type) PMUB04 {
	v := PMUB04{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.StaffIdentification = NewSTF(input.Index(4))
	v.PractitionerDetail = NewPRASlice(input.Index(5))
	v.PractitionerOrganizationUnit = NewORGSlice(input.Index(6))
	return v
}

type PMUB04 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	StaffIdentification           STF
	PractitionerDetail            []PRA
	PractitionerOrganizationUnit  []ORG
}

// NewPMUB04Slice will construct a slice of type PMUB04
func NewPMUB04Slice(input hl7.HL7Type) []PMUB04 {
	values := make([]PMUB04, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPMUB04(input.Index(i))
	}
	return values
}

// NewDELO46 creates an implementation of DELO46
func NewDELO46(input hl7.HL7Type) DELO46 {
	v := DELO46{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.StaffIdentification = NewSTFSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Donnor = NewDELO46Donnor(input.Index(3))
	v.DonationSegment = NewDON(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	return v
}

type DELO46 struct {
	MessageHeader                 MSH
	StaffIdentification           []STF
	UserAuthenticationCredentials UAC
	Donnor                        DELO46Donnor
	DonationSegment               DON
	NotesAndComments              []NTE
}

// NewDELO46Slice will construct a slice of type DELO46
func NewDELO46Slice(input hl7.HL7Type) []DELO46 {
	values := make([]DELO46, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDELO46(input.Index(i))
	}
	return values
}

// NewADTA28 creates an implementation of ADTA28
func NewADTA28(input hl7.HL7Type) ADTA28 {
	v := ADTA28{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(8))
	v.PatientVisit = NewPV1(input.Index(9))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(10))
	v.AdditionalAccessRestriction = NewARVSlice(input.Index(11))
	v.AdditionalRole = NewROLSlice(input.Index(12))
	v.Disability = NewDB1Slice(input.Index(13))
	v.ObservationResult = NewOBXSlice(input.Index(14))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(15))
	v.Diagnosis = NewDG1Slice(input.Index(16))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(17))
	v.Procedure = NewADTA28ProcedureSlice(input.Index(18))
	v.Guarantor = NewGT1Slice(input.Index(19))
	v.Insurance = NewADTA28InsuranceSlice(input.Index(20))
	v.Accident = NewACC(input.Index(21))
	v.Ub82 = NewUB1(input.Index(22))
	v.UniformBillingData = NewUB2(input.Index(23))
	return v
}

type ADTA28 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	AccessRestriction                 []ARV
	Role                              []ROL
	NextOfKinAssociatedParties        []NK1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalAccessRestriction       []ARV
	AdditionalRole                    []ROL
	Disability                        []DB1
	ObservationResult                 []OBX
	PatientAllergyInformation         []AL1
	Diagnosis                         []DG1
	DiagnosisRelatedGroup             DRG
	Procedure                         []ADTA28Procedure
	Guarantor                         []GT1
	Insurance                         []ADTA28Insurance
	Accident                          ACC
	Ub82                              UB1
	UniformBillingData                UB2
}

// NewADTA28Slice will construct a slice of type ADTA28
func NewADTA28Slice(input hl7.HL7Type) []ADTA28 {
	values := make([]ADTA28, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA28(input.Index(i))
	}
	return values
}

// NewCCRI18 creates an implementation of CCRI18
func NewCCRI18(input hl7.HL7Type) CCRI18 {
	v := CCRI18{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.ReferralInformation = NewRF1Slice(input.Index(3))
	v.ProviderContact = NewCCRI18ProviderContactSlice(input.Index(4))
	v.ClinicalOrder = NewCCRI18ClinicalOrderSlice(input.Index(5))
	v.Patient = NewCCRI18PatientSlice(input.Index(6))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(7))
	v.Insurance = NewCCRI18InsuranceSlice(input.Index(8))
	v.AppointmentHistory = NewCCRI18AppointmentHistorySlice(input.Index(9))
	v.ClinicalHistory = NewCCRI18ClinicalHistorySlice(input.Index(10))
	v.PatientVisits = NewCCRI18PatientVisitsSlice(input.Index(11))
	v.MedicationHistory = NewCCRI18MedicationHistorySlice(input.Index(12))
	v.Problem = NewCCRI18ProblemSlice(input.Index(13))
	v.Goal = NewCCRI18GoalSlice(input.Index(14))
	v.Pathway = NewCCRI18PathwaySlice(input.Index(15))
	v.ClinicalRelationshipSegment = NewRELSlice(input.Index(16))
	return v
}

type CCRI18 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	ReferralInformation           []RF1
	ProviderContact               []CCRI18ProviderContact
	ClinicalOrder                 []CCRI18ClinicalOrder
	Patient                       []CCRI18Patient
	NextOfKinAssociatedParties    []NK1
	Insurance                     []CCRI18Insurance
	AppointmentHistory            []CCRI18AppointmentHistory
	ClinicalHistory               []CCRI18ClinicalHistory
	PatientVisits                 []CCRI18PatientVisits
	MedicationHistory             []CCRI18MedicationHistory
	Problem                       []CCRI18Problem
	Goal                          []CCRI18Goal
	Pathway                       []CCRI18Pathway
	ClinicalRelationshipSegment   []REL
}

// NewCCRI18Slice will construct a slice of type CCRI18
func NewCCRI18Slice(input hl7.HL7Type) []CCRI18 {
	values := make([]CCRI18, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCCRI18(input.Index(i))
	}
	return values
}

// NewORAR33 creates an implementation of ORAR33
func NewORAR33(input hl7.HL7Type) ORAR33 {
	v := ORAR33{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.CommonOrder = NewORC(input.Index(5))
	return v
}

type ORAR33 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	CommonOrder                   ORC
}

// NewORAR33Slice will construct a slice of type ORAR33
func NewORAR33Slice(input hl7.HL7Type) []ORAR33 {
	values := make([]ORAR33, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORAR33(input.Index(i))
	}
	return values
}

// NewEACU07 creates an implementation of EACU07
func NewEACU07(input hl7.HL7Type) EACU07 {
	v := EACU07{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EquipmentDetail = NewEQU(input.Index(3))
	v.Command = NewEACU07CommandSlice(input.Index(4))
	return v
}

type EACU07 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EquipmentDetail               EQU
	Command                       []EACU07Command
}

// NewEACU07Slice will construct a slice of type EACU07
func NewEACU07Slice(input hl7.HL7Type) []EACU07 {
	values := make([]EACU07, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEACU07(input.Index(i))
	}
	return values
}

// NewSRMS09 creates an implementation of SRMS09
func NewSRMS09(input hl7.HL7Type) SRMS09 {
	v := SRMS09{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.AppointmentRequest = NewARQ(input.Index(1))
	v.AppointmentPreferences = NewAPR(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSRMS09PatientSlice(input.Index(4))
	v.Resources = NewSRMS09ResourcesSlice(input.Index(5))
	return v
}

type SRMS09 struct {
	MessageHeader          MSH
	AppointmentRequest     ARQ
	AppointmentPreferences APR
	NotesAndComments       []NTE
	Patient                []SRMS09Patient
	Resources              []SRMS09Resources
}

// NewSRMS09Slice will construct a slice of type SRMS09
func NewSRMS09Slice(input hl7.HL7Type) []SRMS09 {
	values := make([]SRMS09, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRMS09(input.Index(i))
	}
	return values
}

// NewMDMT06 creates an implementation of MDMT06
func NewMDMT06(input hl7.HL7Type) MDMT06 {
	v := MDMT06{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientVisit = NewPV1(input.Index(5))
	v.CommonOrder = NewMDMT06CommonOrderSlice(input.Index(6))
	v.TranscriptionDocumentHeader = NewTXA(input.Index(7))
	v.ConsentSegment = NewCONSlice(input.Index(8))
	v.Observation = NewMDMT06ObservationSlice(input.Index(9))
	return v
}

type MDMT06 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientVisit                  PV1
	CommonOrder                   []MDMT06CommonOrder
	TranscriptionDocumentHeader   TXA
	ConsentSegment                []CON
	Observation                   []MDMT06Observation
}

// NewMDMT06Slice will construct a slice of type MDMT06
func NewMDMT06Slice(input hl7.HL7Type) []MDMT06 {
	values := make([]MDMT06, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMDMT06(input.Index(i))
	}
	return values
}

// NewDFTP03 creates an implementation of DFTP03
func NewDFTP03(input hl7.HL7Type) DFTP03 {
	v := DFTP03{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.ParticipationInformation = NewPRTSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.PatientVisit = NewPV1(input.Index(8))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(9))
	v.AdditionalParticipationInformation = NewPRTSlice(input.Index(10))
	v.AdditionalRole = NewROLSlice(input.Index(11))
	v.Disability = NewDB1Slice(input.Index(12))
	v.CommonOrder = NewDFTP03CommonOrderSlice(input.Index(13))
	v.Financial = NewDFTP03FinancialSlice(input.Index(14))
	v.Diagnosis = NewDG1Slice(input.Index(15))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(16))
	v.Guarantor = NewGT1Slice(input.Index(17))
	v.Insurance = NewDFTP03InsuranceSlice(input.Index(18))
	v.Accident = NewACC(input.Index(19))
	return v
}

type DFTP03 struct {
	MessageHeader                      MSH
	SoftwareSegment                    []SFT
	UserAuthenticationCredentials      UAC
	EventType                          EVN
	PatientIdentification              PID
	PatientAdditionalDemographic       PD1
	ParticipationInformation           []PRT
	Role                               []ROL
	PatientVisit                       PV1
	PatientVisitAdditionalInformation  PV2
	AdditionalParticipationInformation []PRT
	AdditionalRole                     []ROL
	Disability                         []DB1
	CommonOrder                        []DFTP03CommonOrder
	Financial                          []DFTP03Financial
	Diagnosis                          []DG1
	DiagnosisRelatedGroup              DRG
	Guarantor                          []GT1
	Insurance                          []DFTP03Insurance
	Accident                           ACC
}

// NewDFTP03Slice will construct a slice of type DFTP03
func NewDFTP03Slice(input hl7.HL7Type) []DFTP03 {
	values := make([]DFTP03, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDFTP03(input.Index(i))
	}
	return values
}

// NewOMLO21 creates an implementation of OMLO21
func NewOMLO21(input hl7.HL7Type) OMLO21 {
	v := OMLO21{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewOMLO21Patient(input.Index(4))
	v.Order = NewOMLO21OrderSlice(input.Index(5))
	return v
}

type OMLO21 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       OMLO21Patient
	Order                         []OMLO21Order
}

// NewOMLO21Slice will construct a slice of type OMLO21
func NewOMLO21Slice(input hl7.HL7Type) []OMLO21 {
	values := make([]OMLO21, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOMLO21(input.Index(i))
	}
	return values
}

// NewRASO17 creates an implementation of RASO17
func NewRASO17(input hl7.HL7Type) RASO17 {
	v := RASO17{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewRASO17Patient(input.Index(4))
	v.Order = NewRASO17OrderSlice(input.Index(5))
	return v
}

type RASO17 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       RASO17Patient
	Order                         []RASO17Order
}

// NewRASO17Slice will construct a slice of type RASO17
func NewRASO17Slice(input hl7.HL7Type) []RASO17 {
	values := make([]RASO17, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRASO17(input.Index(i))
	}
	return values
}

// NewCCII22 creates an implementation of CCII22
func NewCCII22(input hl7.HL7Type) CCII22 {
	v := CCII22{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.PatientIdentification = NewPID(input.Index(5))
	v.PatientAdditionalDemographic = NewPD1(input.Index(6))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(7))
	v.Insurance = NewCCII22InsuranceSlice(input.Index(8))
	v.AppointmentHistory = NewCCII22AppointmentHistorySlice(input.Index(9))
	v.ClinicalHistory = NewCCII22ClinicalHistorySlice(input.Index(10))
	v.PatientVisit = NewCCII22PatientVisitSlice(input.Index(11))
	v.MedicationHistory = NewCCII22MedicationHistorySlice(input.Index(12))
	v.Problem = NewCCII22ProblemSlice(input.Index(13))
	v.Goal = NewCCII22GoalSlice(input.Index(14))
	v.Pathway = NewCCII22PathwaySlice(input.Index(15))
	v.ClinicalRelationshipSegment = NewRELSlice(input.Index(16))
	return v
}

type CCII22 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	PatientIdentification         PID
	PatientAdditionalDemographic  PD1
	NextOfKinAssociatedParties    []NK1
	Insurance                     []CCII22Insurance
	AppointmentHistory            []CCII22AppointmentHistory
	ClinicalHistory               []CCII22ClinicalHistory
	PatientVisit                  []CCII22PatientVisit
	MedicationHistory             []CCII22MedicationHistory
	Problem                       []CCII22Problem
	Goal                          []CCII22Goal
	Pathway                       []CCII22Pathway
	ClinicalRelationshipSegment   []REL
}

// NewCCII22Slice will construct a slice of type CCII22
func NewCCII22Slice(input hl7.HL7Type) []CCII22 {
	values := make([]CCII22, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCCII22(input.Index(i))
	}
	return values
}

// NewMFNM16 creates an implementation of MFNM16
func NewMFNM16(input hl7.HL7Type) MFNM16 {
	v := MFNM16{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MasterFileIdentification = NewMFI(input.Index(3))
	v.MaterialItemRecord = NewMFNM16MaterialItemRecordSlice(input.Index(4))
	return v
}

type MFNM16 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MasterFileIdentification      MFI
	MaterialItemRecord            []MFNM16MaterialItemRecord
}

// NewMFNM16Slice will construct a slice of type MFNM16
func NewMFNM16Slice(input hl7.HL7Type) []MFNM16 {
	values := make([]MFNM16, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFNM16(input.Index(i))
	}
	return values
}

// NewMFKM11 creates an implementation of MFKM11
func NewMFKM11(input hl7.HL7Type) MFKM11 {
	v := MFKM11{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.MasterFileIdentification = NewMFI(input.Index(5))
	return v
}

type MFKM11 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	MasterFileIdentification      MFI
}

// NewMFKM11Slice will construct a slice of type MFKM11
func NewMFKM11Slice(input hl7.HL7Type) []MFKM11 {
	values := make([]MFKM11, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFKM11(input.Index(i))
	}
	return values
}

// NewQBPZ87 creates an implementation of QBPZ87
func NewQBPZ87(input hl7.HL7Type) QBPZ87 {
	v := QBPZ87{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.Qbp = NewQBPZ87Qbp(input.Index(4))
	v.ResponseControlParameter = NewRCP(input.Index(5))
	v.ContinuationPointer = NewDSC(input.Index(6))
	return v
}

type QBPZ87 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	Qbp                           QBPZ87Qbp
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPZ87Slice will construct a slice of type QBPZ87
func NewQBPZ87Slice(input hl7.HL7Type) []QBPZ87 {
	values := make([]QBPZ87, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPZ87(input.Index(i))
	}
	return values
}

// NewQBPQ15 creates an implementation of QBPQ15
func NewQBPQ15(input hl7.HL7Type) QBPQ15 {
	v := QBPQ15{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.AnyHl7Segment = NewHxx(input.Index(4))
	v.ResponseControlParameter = NewRCP(input.Index(5))
	v.ContinuationPointer = NewDSC(input.Index(6))
	return v
}

type QBPQ15 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	AnyHl7Segment                 Hxx
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPQ15Slice will construct a slice of type QBPQ15
func NewQBPQ15Slice(input hl7.HL7Type) []QBPQ15 {
	values := make([]QBPQ15, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPQ15(input.Index(i))
	}
	return values
}

// NewADTA50 creates an implementation of ADTA50
func NewADTA50(input hl7.HL7Type) ADTA50 {
	v := ADTA50{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.MergePatientInformation = NewMRG(input.Index(6))
	v.PatientVisit = NewPV1(input.Index(7))
	return v
}

type ADTA50 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientAdditionalDemographic  PD1
	MergePatientInformation       MRG
	PatientVisit                  PV1
}

// NewADTA50Slice will construct a slice of type ADTA50
func NewADTA50Slice(input hl7.HL7Type) []ADTA50 {
	values := make([]ADTA50, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA50(input.Index(i))
	}
	return values
}

// NewQBPQ25 creates an implementation of QBPQ25
func NewQBPQ25(input hl7.HL7Type) QBPQ25 {
	v := QBPQ25{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.ResponseControlParameter = NewRCP(input.Index(4))
	v.ContinuationPointer = NewDSC(input.Index(5))
	return v
}

type QBPQ25 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPQ25Slice will construct a slice of type QBPQ25
func NewQBPQ25Slice(input hl7.HL7Type) []QBPQ25 {
	values := make([]QBPQ25, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPQ25(input.Index(i))
	}
	return values
}

// NewRGVO15 creates an implementation of RGVO15
func NewRGVO15(input hl7.HL7Type) RGVO15 {
	v := RGVO15{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewRGVO15Patient(input.Index(4))
	v.Order = NewRGVO15OrderSlice(input.Index(5))
	return v
}

type RGVO15 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       RGVO15Patient
	Order                         []RGVO15Order
}

// NewRGVO15Slice will construct a slice of type RGVO15
func NewRGVO15Slice(input hl7.HL7Type) []RGVO15 {
	values := make([]RGVO15, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRGVO15(input.Index(i))
	}
	return values
}

// NewMFKM17 creates an implementation of MFKM17
func NewMFKM17(input hl7.HL7Type) MFKM17 {
	v := MFKM17{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.MasterFileIdentification = NewMFI(input.Index(5))
	v.MasterFileAcknowledgment = NewMFASlice(input.Index(6))
	return v
}

type MFKM17 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	MasterFileIdentification      MFI
	MasterFileAcknowledgment      []MFA
}

// NewMFKM17Slice will construct a slice of type MFKM17
func NewMFKM17Slice(input hl7.HL7Type) []MFKM17 {
	values := make([]MFKM17, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFKM17(input.Index(i))
	}
	return values
}

// NewINRU06 creates an implementation of INRU06
func NewINRU06(input hl7.HL7Type) INRU06 {
	v := INRU06{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EquipmentDetail = NewEQU(input.Index(3))
	v.InventoryDetail = NewINVSlice(input.Index(4))
	return v
}

type INRU06 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EquipmentDetail               EQU
	InventoryDetail               []INV
}

// NewINRU06Slice will construct a slice of type INRU06
func NewINRU06Slice(input hl7.HL7Type) []INRU06 {
	values := make([]INRU06, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewINRU06(input.Index(i))
	}
	return values
}

// NewORLO22 creates an implementation of ORLO22
func NewORLO22(input hl7.HL7Type) ORLO22 {
	v := ORLO22{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewORLO22Response(input.Index(6))
	return v
}

type ORLO22 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      ORLO22Response
}

// NewORLO22Slice will construct a slice of type ORLO22
func NewORLO22Slice(input hl7.HL7Type) []ORLO22 {
	values := make([]ORLO22, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORLO22(input.Index(i))
	}
	return values
}

// NewRPII01 creates an implementation of RPII01
func NewRPII01(input hl7.HL7Type) RPII01 {
	v := RPII01{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Provider = NewRPII01ProviderSlice(input.Index(4))
	v.PatientIdentification = NewPID(input.Index(5))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(6))
	v.GuarantorInsurance = NewRPII01GuarantorInsurance(input.Index(7))
	v.NotesAndComments = NewNTESlice(input.Index(8))
	return v
}

type RPII01 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Provider                      []RPII01Provider
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	GuarantorInsurance            RPII01GuarantorInsurance
	NotesAndComments              []NTE
}

// NewRPII01Slice will construct a slice of type RPII01
func NewRPII01Slice(input hl7.HL7Type) []RPII01 {
	values := make([]RPII01, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRPII01(input.Index(i))
	}
	return values
}

// NewMDMT08 creates an implementation of MDMT08
func NewMDMT08(input hl7.HL7Type) MDMT08 {
	v := MDMT08{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientVisit = NewPV1(input.Index(5))
	v.CommonOrder = NewMDMT08CommonOrderSlice(input.Index(6))
	v.TranscriptionDocumentHeader = NewTXA(input.Index(7))
	v.ConsentSegment = NewCONSlice(input.Index(8))
	v.Observation = NewMDMT08ObservationSlice(input.Index(9))
	return v
}

type MDMT08 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientVisit                  PV1
	CommonOrder                   []MDMT08CommonOrder
	TranscriptionDocumentHeader   TXA
	ConsentSegment                []CON
	Observation                   []MDMT08Observation
}

// NewMDMT08Slice will construct a slice of type MDMT08
func NewMDMT08Slice(input hl7.HL7Type) []MDMT08 {
	values := make([]MDMT08, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMDMT08(input.Index(i))
	}
	return values
}

// NewPPPPCD creates an implementation of PPPPCD
func NewPPPPCD(input hl7.HL7Type) PPPPCD {
	v := PPPPCD{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientVisit = NewPPPPCDPatientVisit(input.Index(4))
	v.Pathway = NewPPPPCDPathwaySlice(input.Index(5))
	return v
}

type PPPPCD struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientIdentification         PID
	PatientVisit                  PPPPCDPatientVisit
	Pathway                       []PPPPCDPathway
}

// NewPPPPCDSlice will construct a slice of type PPPPCD
func NewPPPPCDSlice(input hl7.HL7Type) []PPPPCD {
	values := make([]PPPPCD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPPPPCD(input.Index(i))
	}
	return values
}

// NewCCRI16 creates an implementation of CCRI16
func NewCCRI16(input hl7.HL7Type) CCRI16 {
	v := CCRI16{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.ReferralInformation = NewRF1Slice(input.Index(3))
	v.ProviderContact = NewCCRI16ProviderContactSlice(input.Index(4))
	v.ClinicalOrder = NewCCRI16ClinicalOrderSlice(input.Index(5))
	v.Patient = NewCCRI16PatientSlice(input.Index(6))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(7))
	v.Insurance = NewCCRI16InsuranceSlice(input.Index(8))
	v.AppointmentHistory = NewCCRI16AppointmentHistorySlice(input.Index(9))
	v.ClinicalHistory = NewCCRI16ClinicalHistorySlice(input.Index(10))
	v.PatientVisits = NewCCRI16PatientVisitsSlice(input.Index(11))
	v.MedicationHistory = NewCCRI16MedicationHistorySlice(input.Index(12))
	v.Problem = NewCCRI16ProblemSlice(input.Index(13))
	v.Goal = NewCCRI16GoalSlice(input.Index(14))
	v.Pathway = NewCCRI16PathwaySlice(input.Index(15))
	v.ClinicalRelationshipSegment = NewRELSlice(input.Index(16))
	return v
}

type CCRI16 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	ReferralInformation           []RF1
	ProviderContact               []CCRI16ProviderContact
	ClinicalOrder                 []CCRI16ClinicalOrder
	Patient                       []CCRI16Patient
	NextOfKinAssociatedParties    []NK1
	Insurance                     []CCRI16Insurance
	AppointmentHistory            []CCRI16AppointmentHistory
	ClinicalHistory               []CCRI16ClinicalHistory
	PatientVisits                 []CCRI16PatientVisits
	MedicationHistory             []CCRI16MedicationHistory
	Problem                       []CCRI16Problem
	Goal                          []CCRI16Goal
	Pathway                       []CCRI16Pathway
	ClinicalRelationshipSegment   []REL
}

// NewCCRI16Slice will construct a slice of type CCRI16
func NewCCRI16Slice(input hl7.HL7Type) []CCRI16 {
	values := make([]CCRI16, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCCRI16(input.Index(i))
	}
	return values
}

// NewADTA33 creates an implementation of ADTA33
func NewADTA33(input hl7.HL7Type) ADTA33 {
	v := ADTA33{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.Disability = NewDB1Slice(input.Index(8))
	v.ObservationResult = NewOBXSlice(input.Index(9))
	return v
}

type ADTA33 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	Disability                        []DB1
	ObservationResult                 []OBX
}

// NewADTA33Slice will construct a slice of type ADTA33
func NewADTA33Slice(input hl7.HL7Type) []ADTA33 {
	values := make([]ADTA33, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA33(input.Index(i))
	}
	return values
}

// NewQBPQ34 creates an implementation of QBPQ34
func NewQBPQ34(input hl7.HL7Type) QBPQ34 {
	v := QBPQ34{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.ResponseControlParameter = NewRCP(input.Index(4))
	v.ContinuationPointer = NewDSC(input.Index(5))
	return v
}

type QBPQ34 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPQ34Slice will construct a slice of type QBPQ34
func NewQBPQ34Slice(input hl7.HL7Type) []QBPQ34 {
	values := make([]QBPQ34, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPQ34(input.Index(i))
	}
	return values
}

// NewRTBZ74 creates an implementation of RTBZ74
func NewRTBZ74(input hl7.HL7Type) RTBZ74 {
	v := RTBZ74{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.RowDefinition = NewRTBZ74RowDefinition(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RTBZ74 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	RowDefinition                 RTBZ74RowDefinition
	ContinuationPointer           DSC
}

// NewRTBZ74Slice will construct a slice of type RTBZ74
func NewRTBZ74Slice(input hl7.HL7Type) []RTBZ74 {
	values := make([]RTBZ74, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRTBZ74(input.Index(i))
	}
	return values
}

// NewQBPQ24 creates an implementation of QBPQ24
func NewQBPQ24(input hl7.HL7Type) QBPQ24 {
	v := QBPQ24{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.ResponseControlParameter = NewRCP(input.Index(4))
	v.ContinuationPointer = NewDSC(input.Index(5))
	return v
}

type QBPQ24 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPQ24Slice will construct a slice of type QBPQ24
func NewQBPQ24Slice(input hl7.HL7Type) []QBPQ24 {
	values := make([]QBPQ24, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPQ24(input.Index(i))
	}
	return values
}

// NewPPRPC2 creates an implementation of PPRPC2
func NewPPRPC2(input hl7.HL7Type) PPRPC2 {
	v := PPRPC2{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientVisit = NewPPRPC2PatientVisit(input.Index(4))
	v.Problem = NewPPRPC2ProblemSlice(input.Index(5))
	return v
}

type PPRPC2 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientIdentification         PID
	PatientVisit                  PPRPC2PatientVisit
	Problem                       []PPRPC2Problem
}

// NewPPRPC2Slice will construct a slice of type PPRPC2
func NewPPRPC2Slice(input hl7.HL7Type) []PPRPC2 {
	values := make([]PPRPC2, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPPRPC2(input.Index(i))
	}
	return values
}

// NewMFKM13 creates an implementation of MFKM13
func NewMFKM13(input hl7.HL7Type) MFKM13 {
	v := MFKM13{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.MasterFileIdentification = NewMFI(input.Index(5))
	v.MasterFileAcknowledgment = NewMFASlice(input.Index(6))
	return v
}

type MFKM13 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	MasterFileIdentification      MFI
	MasterFileAcknowledgment      []MFA
}

// NewMFKM13Slice will construct a slice of type MFKM13
func NewMFKM13Slice(input hl7.HL7Type) []MFKM13 {
	values := make([]MFKM13, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFKM13(input.Index(i))
	}
	return values
}

// NewCQUI19 creates an implementation of CQUI19
func NewCQUI19(input hl7.HL7Type) CQUI19 {
	v := CQUI19{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.ReferralInformation = NewRF1(input.Index(5))
	v.ProviderContact = NewCQUI19ProviderContact(input.Index(6))
	v.Patient = NewCQUI19PatientSlice(input.Index(7))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(8))
	v.Insurance = NewCQUI19InsuranceSlice(input.Index(9))
	v.AppointmentHistory = NewCQUI19AppointmentHistorySlice(input.Index(10))
	v.ClinicalHistory = NewCQUI19ClinicalHistorySlice(input.Index(11))
	v.PatientVisit = NewCQUI19PatientVisitSlice(input.Index(12))
	v.MedicationHistory = NewCQUI19MedicationHistorySlice(input.Index(13))
	v.Problem = NewCQUI19ProblemSlice(input.Index(14))
	v.Goal = NewCQUI19GoalSlice(input.Index(15))
	v.Pathway = NewCQUI19PathwaySlice(input.Index(16))
	v.ClinicalRelationshipSegment = NewRELSlice(input.Index(17))
	return v
}

type CQUI19 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	ReferralInformation           RF1
	ProviderContact               CQUI19ProviderContact
	Patient                       []CQUI19Patient
	NextOfKinAssociatedParties    []NK1
	Insurance                     []CQUI19Insurance
	AppointmentHistory            []CQUI19AppointmentHistory
	ClinicalHistory               []CQUI19ClinicalHistory
	PatientVisit                  []CQUI19PatientVisit
	MedicationHistory             []CQUI19MedicationHistory
	Problem                       []CQUI19Problem
	Goal                          []CQUI19Goal
	Pathway                       []CQUI19Pathway
	ClinicalRelationshipSegment   []REL
}

// NewCQUI19Slice will construct a slice of type CQUI19
func NewCQUI19Slice(input hl7.HL7Type) []CQUI19 {
	values := make([]CQUI19, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCQUI19(input.Index(i))
	}
	return values
}

// NewSRMS07 creates an implementation of SRMS07
func NewSRMS07(input hl7.HL7Type) SRMS07 {
	v := SRMS07{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.AppointmentRequest = NewARQ(input.Index(1))
	v.AppointmentPreferences = NewAPR(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSRMS07PatientSlice(input.Index(4))
	v.Resources = NewSRMS07ResourcesSlice(input.Index(5))
	return v
}

type SRMS07 struct {
	MessageHeader          MSH
	AppointmentRequest     ARQ
	AppointmentPreferences APR
	NotesAndComments       []NTE
	Patient                []SRMS07Patient
	Resources              []SRMS07Resources
}

// NewSRMS07Slice will construct a slice of type SRMS07
func NewSRMS07Slice(input hl7.HL7Type) []SRMS07 {
	values := make([]SRMS07, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRMS07(input.Index(i))
	}
	return values
}

// NewSDRS31 creates an implementation of SDRS31
func NewSDRS31(input hl7.HL7Type) SDRS31 {
	v := SDRS31{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.AntiMicrobialDeviceData = NewSDRS31AntiMicrobialDeviceData(input.Index(3))
	return v
}

type SDRS31 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	AntiMicrobialDeviceData       SDRS31AntiMicrobialDeviceData
}

// NewSDRS31Slice will construct a slice of type SDRS31
func NewSDRS31Slice(input hl7.HL7Type) []SDRS31 {
	values := make([]SDRS31, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSDRS31(input.Index(i))
	}
	return values
}

// NewQSBQ16 creates an implementation of QSBQ16
func NewQSBQ16(input hl7.HL7Type) QSBQ16 {
	v := QSBQ16{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.ResponseControlParameter = NewRCP(input.Index(4))
	v.ContinuationPointer = NewDSC(input.Index(5))
	return v
}

type QSBQ16 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQSBQ16Slice will construct a slice of type QSBQ16
func NewQSBQ16Slice(input hl7.HL7Type) []QSBQ16 {
	values := make([]QSBQ16, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQSBQ16(input.Index(i))
	}
	return values
}

// NewCRMC08 creates an implementation of CRMC08
func NewCRMC08(input hl7.HL7Type) CRMC08 {
	v := CRMC08{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Patient = NewCRMC08PatientSlice(input.Index(3))
	return v
}

type CRMC08 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Patient                       []CRMC08Patient
}

// NewCRMC08Slice will construct a slice of type CRMC08
func NewCRMC08Slice(input hl7.HL7Type) []CRMC08 {
	values := make([]CRMC08, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCRMC08(input.Index(i))
	}
	return values
}

// NewSRRS02 creates an implementation of SRRS02
func NewSRRS02(input hl7.HL7Type) SRRS02 {
	v := SRRS02{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.Schedule = NewSRRS02Schedule(input.Index(3))
	return v
}

type SRRS02 struct {
	MessageHeader         MSH
	MessageAcknowledgment MSA
	Error                 []ERR
	Schedule              SRRS02Schedule
}

// NewSRRS02Slice will construct a slice of type SRRS02
func NewSRRS02Slice(input hl7.HL7Type) []SRRS02 {
	values := make([]SRRS02, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRRS02(input.Index(i))
	}
	return values
}

// NewLSRU13 creates an implementation of LSRU13
func NewLSRU13(input hl7.HL7Type) LSRU13 {
	v := LSRU13{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EquipmentDetail = NewEQU(input.Index(3))
	v.EquipmentLogService = NewEQPSlice(input.Index(4))
	return v
}

type LSRU13 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EquipmentDetail               EQU
	EquipmentLogService           []EQP
}

// NewLSRU13Slice will construct a slice of type LSRU13
func NewLSRU13Slice(input hl7.HL7Type) []LSRU13 {
	values := make([]LSRU13, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewLSRU13(input.Index(i))
	}
	return values
}

// NewBPSO29 creates an implementation of BPSO29
func NewBPSO29(input hl7.HL7Type) BPSO29 {
	v := BPSO29{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewBPSO29Patient(input.Index(4))
	v.Order = NewBPSO29OrderSlice(input.Index(5))
	return v
}

type BPSO29 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       BPSO29Patient
	Order                         []BPSO29Order
}

// NewBPSO29Slice will construct a slice of type BPSO29
func NewBPSO29Slice(input hl7.HL7Type) []BPSO29 {
	values := make([]BPSO29, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBPSO29(input.Index(i))
	}
	return values
}

// NewRSPK24 creates an implementation of RSPK24
func NewRSPK24(input hl7.HL7Type) RSPK24 {
	v := RSPK24{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.QueryResponse = NewRSPK24QueryResponse(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RSPK24 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	QueryResponse                 RSPK24QueryResponse
	ContinuationPointer           DSC
}

// NewRSPK24Slice will construct a slice of type RSPK24
func NewRSPK24Slice(input hl7.HL7Type) []RSPK24 {
	values := make([]RSPK24, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPK24(input.Index(i))
	}
	return values
}

// NewINUU05 creates an implementation of INUU05
func NewINUU05(input hl7.HL7Type) INUU05 {
	v := INUU05{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EquipmentDetail = NewEQU(input.Index(3))
	v.InventoryDetail = NewINVSlice(input.Index(4))
	return v
}

type INUU05 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EquipmentDetail               EQU
	InventoryDetail               []INV
}

// NewINUU05Slice will construct a slice of type INUU05
func NewINUU05Slice(input hl7.HL7Type) []INUU05 {
	values := make([]INUU05, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewINUU05(input.Index(i))
	}
	return values
}

// NewMFNM07 creates an implementation of MFNM07
func NewMFNM07(input hl7.HL7Type) MFNM07 {
	v := MFNM07{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MasterFileIdentification = NewMFI(input.Index(3))
	v.MfClinStudySched = NewMFNM07MfClinStudySchedSlice(input.Index(4))
	return v
}

type MFNM07 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MasterFileIdentification      MFI
	MfClinStudySched              []MFNM07MfClinStudySched
}

// NewMFNM07Slice will construct a slice of type MFNM07
func NewMFNM07Slice(input hl7.HL7Type) []MFNM07 {
	values := make([]MFNM07, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFNM07(input.Index(i))
	}
	return values
}

// NewSSUU03 creates an implementation of SSUU03
func NewSSUU03(input hl7.HL7Type) SSUU03 {
	v := SSUU03{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EquipmentDetail = NewEQU(input.Index(3))
	v.SpecimenContainer = NewSSUU03SpecimenContainerSlice(input.Index(4))
	return v
}

type SSUU03 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EquipmentDetail               EQU
	SpecimenContainer             []SSUU03SpecimenContainer
}

// NewSSUU03Slice will construct a slice of type SSUU03
func NewSSUU03Slice(input hl7.HL7Type) []SSUU03 {
	values := make([]SSUU03, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSSUU03(input.Index(i))
	}
	return values
}

// NewRQAI08 creates an implementation of RQAI08
func NewRQAI08(input hl7.HL7Type) RQAI08 {
	v := RQAI08{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.ReferralInformation = NewRF1(input.Index(3))
	v.Authorization = NewRQAI08Authorization(input.Index(4))
	v.Provider = NewRQAI08ProviderSlice(input.Index(5))
	v.PatientIdentification = NewPID(input.Index(6))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(7))
	v.GuarantorInsurance = NewRQAI08GuarantorInsurance(input.Index(8))
	v.Accident = NewACC(input.Index(9))
	v.Diagnosis = NewDG1Slice(input.Index(10))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(11))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(12))
	v.Procedure = NewRQAI08ProcedureSlice(input.Index(13))
	v.Observation = NewRQAI08ObservationSlice(input.Index(14))
	v.Visit = NewRQAI08Visit(input.Index(15))
	v.NotesAndComments = NewNTESlice(input.Index(16))
	return v
}

type RQAI08 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	ReferralInformation           RF1
	Authorization                 RQAI08Authorization
	Provider                      []RQAI08Provider
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	GuarantorInsurance            RQAI08GuarantorInsurance
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []RQAI08Procedure
	Observation                   []RQAI08Observation
	Visit                         RQAI08Visit
	NotesAndComments              []NTE
}

// NewRQAI08Slice will construct a slice of type RQAI08
func NewRQAI08Slice(input hl7.HL7Type) []RQAI08 {
	values := make([]RQAI08, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRQAI08(input.Index(i))
	}
	return values
}

// NewPPGPCH creates an implementation of PPGPCH
func NewPPGPCH(input hl7.HL7Type) PPGPCH {
	v := PPGPCH{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientVisit = NewPPGPCHPatientVisit(input.Index(4))
	v.Pathway = NewPPGPCHPathwaySlice(input.Index(5))
	return v
}

type PPGPCH struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientIdentification         PID
	PatientVisit                  PPGPCHPatientVisit
	Pathway                       []PPGPCHPathway
}

// NewPPGPCHSlice will construct a slice of type PPGPCH
func NewPPGPCHSlice(input hl7.HL7Type) []PPGPCH {
	values := make([]PPGPCH, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPPGPCH(input.Index(i))
	}
	return values
}

// NewADTA26 creates an implementation of ADTA26
func NewADTA26(input hl7.HL7Type) ADTA26 {
	v := ADTA26{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.Disability = NewDB1Slice(input.Index(8))
	v.ObservationResult = NewOBXSlice(input.Index(9))
	return v
}

type ADTA26 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	Disability                        []DB1
	ObservationResult                 []OBX
}

// NewADTA26Slice will construct a slice of type ADTA26
func NewADTA26Slice(input hl7.HL7Type) []ADTA26 {
	values := make([]ADTA26, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA26(input.Index(i))
	}
	return values
}

// NewRSPK33 creates an implementation of RSPK33
func NewRSPK33(input hl7.HL7Type) RSPK33 {
	v := RSPK33{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.MessageAcknowledgment = NewMSA(input.Index(2))
	v.Error = NewERR(input.Index(3))
	v.QueryAcknowledgment = NewQAK(input.Index(4))
	v.QueryParameterDefinition = NewQPD(input.Index(5))
	v.Donnor = NewRSPK33DonnorSlice(input.Index(6))
	v.ContinuationPointer = NewDSC(input.Index(7))
	return v
}

type RSPK33 struct {
	MessageHeader            MSH
	SoftwareSegment          []SFT
	MessageAcknowledgment    MSA
	Error                    ERR
	QueryAcknowledgment      QAK
	QueryParameterDefinition QPD
	Donnor                   []RSPK33Donnor
	ContinuationPointer      DSC
}

// NewRSPK33Slice will construct a slice of type RSPK33
func NewRSPK33Slice(input hl7.HL7Type) []RSPK33 {
	values := make([]RSPK33, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPK33(input.Index(i))
	}
	return values
}

// NewQBPQ22 creates an implementation of QBPQ22
func NewQBPQ22(input hl7.HL7Type) QBPQ22 {
	v := QBPQ22{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.ResponseControlParameter = NewRCP(input.Index(4))
	v.ContinuationPointer = NewDSC(input.Index(5))
	return v
}

type QBPQ22 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPQ22Slice will construct a slice of type QBPQ22
func NewQBPQ22Slice(input hl7.HL7Type) []QBPQ22 {
	values := make([]QBPQ22, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPQ22(input.Index(i))
	}
	return values
}

// NewORUR31 creates an implementation of ORUR31
func NewORUR31(input hl7.HL7Type) ORUR31 {
	v := ORUR31{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientAdditionalDemographic = NewPD1(input.Index(4))
	v.ParticipationInformation = NewPRTSlice(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.PatientObservation = NewORUR31PatientObservationSlice(input.Index(7))
	v.Visit = NewORUR31Visit(input.Index(8))
	v.CommonOrder = NewORC(input.Index(9))
	v.AdditionalParticipationInformation = NewPRTSlice(input.Index(10))
	v.ObservationRequest = NewOBR(input.Index(11))
	v.NotesAndComments = NewNTESlice(input.Index(12))
	v.MoreAdditionalParticipationInformation = NewPRTSlice(input.Index(13))
	v.TimingQty = NewORUR31TimingQtySlice(input.Index(14))
	v.Observation = NewORUR31ObservationSlice(input.Index(15))
	return v
}

type ORUR31 struct {
	MessageHeader                          MSH
	SoftwareSegment                        []SFT
	UserAuthenticationCredentials          UAC
	PatientIdentification                  PID
	PatientAdditionalDemographic           PD1
	ParticipationInformation               []PRT
	AccessRestriction                      []ARV
	PatientObservation                     []ORUR31PatientObservation
	Visit                                  ORUR31Visit
	CommonOrder                            ORC
	AdditionalParticipationInformation     []PRT
	ObservationRequest                     OBR
	NotesAndComments                       []NTE
	MoreAdditionalParticipationInformation []PRT
	TimingQty                              []ORUR31TimingQty
	Observation                            []ORUR31Observation
}

// NewORUR31Slice will construct a slice of type ORUR31
func NewORUR31Slice(input hl7.HL7Type) []ORUR31 {
	values := make([]ORUR31, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORUR31(input.Index(i))
	}
	return values
}

// NewSRRS07 creates an implementation of SRRS07
func NewSRRS07(input hl7.HL7Type) SRRS07 {
	v := SRRS07{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.Schedule = NewSRRS07Schedule(input.Index(3))
	return v
}

type SRRS07 struct {
	MessageHeader         MSH
	MessageAcknowledgment MSA
	Error                 []ERR
	Schedule              SRRS07Schedule
}

// NewSRRS07Slice will construct a slice of type SRRS07
func NewSRRS07Slice(input hl7.HL7Type) []SRRS07 {
	values := make([]SRRS07, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRRS07(input.Index(i))
	}
	return values
}

// NewUDMQ05 creates an implementation of UDMQ05
func NewUDMQ05(input hl7.HL7Type) UDMQ05 {
	v := UDMQ05{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.ResultsUpdateDefinition = NewURD(input.Index(3))
	v.ResultsUpdateSelectionCriteria = NewURS(input.Index(4))
	v.DisplayData = NewDSPSlice(input.Index(5))
	v.ContinuationPointer = NewDSC(input.Index(6))
	return v
}

type UDMQ05 struct {
	MessageHeader                  MSH
	SoftwareSegment                []SFT
	UserAuthenticationCredentials  UAC
	ResultsUpdateDefinition        URD
	ResultsUpdateSelectionCriteria URS
	DisplayData                    []DSP
	ContinuationPointer            DSC
}

// NewUDMQ05Slice will construct a slice of type UDMQ05
func NewUDMQ05Slice(input hl7.HL7Type) []UDMQ05 {
	values := make([]UDMQ05, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewUDMQ05(input.Index(i))
	}
	return values
}

// NewRSPZ84 creates an implementation of RSPZ84
func NewRSPZ84(input hl7.HL7Type) RSPZ84 {
	v := RSPZ84{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.SegmentPattern = NewRSPZ84SegmentPattern(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RSPZ84 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	SegmentPattern                RSPZ84SegmentPattern
	ContinuationPointer           DSC
}

// NewRSPZ84Slice will construct a slice of type RSPZ84
func NewRSPZ84Slice(input hl7.HL7Type) []RSPZ84 {
	values := make([]RSPZ84, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPZ84(input.Index(i))
	}
	return values
}

// NewADTA42 creates an implementation of ADTA42
func NewADTA42(input hl7.HL7Type) ADTA42 {
	v := ADTA42{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.Patient = NewADTA42PatientSlice(input.Index(4))
	return v
}

type ADTA42 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	Patient                       []ADTA42Patient
}

// NewADTA42Slice will construct a slice of type ADTA42
func NewADTA42Slice(input hl7.HL7Type) []ADTA42 {
	values := make([]ADTA42, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA42(input.Index(i))
	}
	return values
}

// NewRSPK21 creates an implementation of RSPK21
func NewRSPK21(input hl7.HL7Type) RSPK21 {
	v := RSPK21{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.QueryResponse = NewRSPK21QueryResponse(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RSPK21 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	QueryResponse                 RSPK21QueryResponse
	ContinuationPointer           DSC
}

// NewRSPK21Slice will construct a slice of type RSPK21
func NewRSPK21Slice(input hl7.HL7Type) []RSPK21 {
	values := make([]RSPK21, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPK21(input.Index(i))
	}
	return values
}

// NewSLRS29 creates an implementation of SLRS29
func NewSLRS29(input hl7.HL7Type) SLRS29 {
	v := SLRS29{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.SterilizationLot = NewSLTSlice(input.Index(3))
	return v
}

type SLRS29 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	SterilizationLot              []SLT
}

// NewSLRS29Slice will construct a slice of type SLRS29
func NewSLRS29Slice(input hl7.HL7Type) []SLRS29 {
	values := make([]SLRS29, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSLRS29(input.Index(i))
	}
	return values
}

// NewORUR30 creates an implementation of ORUR30
func NewORUR30(input hl7.HL7Type) ORUR30 {
	v := ORUR30{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientAdditionalDemographic = NewPD1(input.Index(4))
	v.ParticipationInformation = NewPRTSlice(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.PatientObservation = NewORUR30PatientObservationSlice(input.Index(7))
	v.Visit = NewORUR30Visit(input.Index(8))
	v.CommonOrder = NewORC(input.Index(9))
	v.SecondParticipationInformation = NewPRTSlice(input.Index(10))
	v.ObservationRequest = NewOBR(input.Index(11))
	v.NotesAndComments = NewNTESlice(input.Index(12))
	v.ThirdParticipationInformation = NewPRTSlice(input.Index(13))
	v.TimingQty = NewORUR30TimingQtySlice(input.Index(14))
	v.Observation = NewORUR30ObservationSlice(input.Index(15))
	return v
}

type ORUR30 struct {
	MessageHeader                  MSH
	SoftwareSegment                []SFT
	UserAuthenticationCredentials  UAC
	PatientIdentification          PID
	PatientAdditionalDemographic   PD1
	ParticipationInformation       []PRT
	AccessRestriction              []ARV
	PatientObservation             []ORUR30PatientObservation
	Visit                          ORUR30Visit
	CommonOrder                    ORC
	SecondParticipationInformation []PRT
	ObservationRequest             OBR
	NotesAndComments               []NTE
	ThirdParticipationInformation  []PRT
	TimingQty                      []ORUR30TimingQty
	Observation                    []ORUR30Observation
}

// NewORUR30Slice will construct a slice of type ORUR30
func NewORUR30Slice(input hl7.HL7Type) []ORUR30 {
	values := make([]ORUR30, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORUR30(input.Index(i))
	}
	return values
}

// NewOPLO37 creates an implementation of OPLO37
func NewOPLO37(input hl7.HL7Type) OPLO37 {
	v := OPLO37{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.ParticipationInformation = NewPRTSlice(input.Index(4))
	v.Guarantor = NewOPLO37Guarantor(input.Index(5))
	v.Order = NewOPLO37OrderSlice(input.Index(6))
	return v
}

type OPLO37 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	ParticipationInformation      []PRT
	Guarantor                     OPLO37Guarantor
	Order                         []OPLO37Order
}

// NewOPLO37Slice will construct a slice of type OPLO37
func NewOPLO37Slice(input hl7.HL7Type) []OPLO37 {
	values := make([]OPLO37, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOPLO37(input.Index(i))
	}
	return values
}

// NewRTBZ78 creates an implementation of RTBZ78
func NewRTBZ78(input hl7.HL7Type) RTBZ78 {
	v := RTBZ78{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.RowDefinition = NewRTBZ78RowDefinition(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RTBZ78 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	RowDefinition                 RTBZ78RowDefinition
	ContinuationPointer           DSC
}

// NewRTBZ78Slice will construct a slice of type RTBZ78
func NewRTBZ78Slice(input hl7.HL7Type) []RTBZ78 {
	values := make([]RTBZ78, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRTBZ78(input.Index(i))
	}
	return values
}

// NewRPAI08 creates an implementation of RPAI08
func NewRPAI08(input hl7.HL7Type) RPAI08 {
	v := RPAI08{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.ReferralInformation = NewRF1(input.Index(4))
	v.Authorization = NewRPAI08Authorization(input.Index(5))
	v.Provider = NewRPAI08ProviderSlice(input.Index(6))
	v.PatientIdentification = NewPID(input.Index(7))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(8))
	v.Guarantor = NewGT1Slice(input.Index(9))
	v.Insurance = NewRPAI08InsuranceSlice(input.Index(10))
	v.Accident = NewACC(input.Index(11))
	v.Diagnosis = NewDG1Slice(input.Index(12))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(13))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(14))
	v.Procedure = NewRPAI08ProcedureSlice(input.Index(15))
	v.Observation = NewRPAI08ObservationSlice(input.Index(16))
	v.Visit = NewRPAI08Visit(input.Index(17))
	v.NotesAndComments = NewNTESlice(input.Index(18))
	return v
}

type RPAI08 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	ReferralInformation           RF1
	Authorization                 RPAI08Authorization
	Provider                      []RPAI08Provider
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	Guarantor                     []GT1
	Insurance                     []RPAI08Insurance
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []RPAI08Procedure
	Observation                   []RPAI08Observation
	Visit                         RPAI08Visit
	NotesAndComments              []NTE
}

// NewRPAI08Slice will construct a slice of type RPAI08
func NewRPAI08Slice(input hl7.HL7Type) []RPAI08 {
	values := make([]RPAI08, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRPAI08(input.Index(i))
	}
	return values
}

// NewPMUB06 creates an implementation of PMUB06
func NewPMUB06(input hl7.HL7Type) PMUB06 {
	v := PMUB06{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.StaffIdentification = NewSTF(input.Index(4))
	v.PractitionerDetail = NewPRASlice(input.Index(5))
	v.PractitionerOrganizationUnit = NewORGSlice(input.Index(6))
	return v
}

type PMUB06 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	StaffIdentification           STF
	PractitionerDetail            []PRA
	PractitionerOrganizationUnit  []ORG
}

// NewPMUB06Slice will construct a slice of type PMUB06
func NewPMUB06Slice(input hl7.HL7Type) []PMUB06 {
	values := make([]PMUB06, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPMUB06(input.Index(i))
	}
	return values
}

// NewSRMS02 creates an implementation of SRMS02
func NewSRMS02(input hl7.HL7Type) SRMS02 {
	v := SRMS02{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.AppointmentRequest = NewARQ(input.Index(1))
	v.AppointmentPreferences = NewAPR(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSRMS02PatientSlice(input.Index(4))
	v.Resources = NewSRMS02ResourcesSlice(input.Index(5))
	return v
}

type SRMS02 struct {
	MessageHeader          MSH
	AppointmentRequest     ARQ
	AppointmentPreferences APR
	NotesAndComments       []NTE
	Patient                []SRMS02Patient
	Resources              []SRMS02Resources
}

// NewSRMS02Slice will construct a slice of type SRMS02
func NewSRMS02Slice(input hl7.HL7Type) []SRMS02 {
	values := make([]SRMS02, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRMS02(input.Index(i))
	}
	return values
}

// NewOPUR25 creates an implementation of OPUR25
func NewOPUR25(input hl7.HL7Type) OPUR25 {
	v := OPUR25{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTE(input.Index(3))
	v.PatientVisit = NewPV1(input.Index(4))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(5))
	v.ParticipationInformation = NewPRTSlice(input.Index(6))
	v.PatientVisitObservation = NewOPUR25PatientVisitObservationSlice(input.Index(7))
	v.AccessionDetail = NewOPUR25AccessionDetailSlice(input.Index(8))
	return v
}

type OPUR25 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	NotesAndComments                  NTE
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	ParticipationInformation          []PRT
	PatientVisitObservation           []OPUR25PatientVisitObservation
	AccessionDetail                   []OPUR25AccessionDetail
}

// NewOPUR25Slice will construct a slice of type OPUR25
func NewOPUR25Slice(input hl7.HL7Type) []OPUR25 {
	values := make([]OPUR25, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOPUR25(input.Index(i))
	}
	return values
}

// NewQBPZnn creates an implementation of QBPZnn
func NewQBPZnn(input hl7.HL7Type) QBPZnn {
	v := QBPZnn{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.TableRowDefinition = NewRDF(input.Index(4))
	v.ResponseControlParameter = NewRCP(input.Index(5))
	v.ContinuationPointer = NewDSC(input.Index(6))
	return v
}

type QBPZnn struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	TableRowDefinition            RDF
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPZnnSlice will construct a slice of type QBPZnn
func NewQBPZnnSlice(input hl7.HL7Type) []QBPZnn {
	values := make([]QBPZnn, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPZnn(input.Index(i))
	}
	return values
}

// NewRREO12 creates an implementation of RREO12
func NewRREO12(input hl7.HL7Type) RREO12 {
	v := RREO12{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewRREO12Response(input.Index(6))
	return v
}

type RREO12 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      RREO12Response
}

// NewRREO12Slice will construct a slice of type RREO12
func NewRREO12Slice(input hl7.HL7Type) []RREO12 {
	values := make([]RREO12, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRREO12(input.Index(i))
	}
	return values
}

// NewADTA13 creates an implementation of ADTA13
func NewADTA13(input hl7.HL7Type) ADTA13 {
	v := ADTA13{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(8))
	v.PatientVisit = NewPV1(input.Index(9))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(10))
	v.AdditionalAccessRestriction = NewARVSlice(input.Index(11))
	v.AdditionalRole = NewROLSlice(input.Index(12))
	v.Disability = NewDB1Slice(input.Index(13))
	v.ObservationResult = NewOBXSlice(input.Index(14))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(15))
	v.Diagnosis = NewDG1Slice(input.Index(16))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(17))
	v.Procedure = NewADTA13ProcedureSlice(input.Index(18))
	v.Guarantor = NewGT1Slice(input.Index(19))
	v.Insurance = NewADTA13InsuranceSlice(input.Index(20))
	v.Accident = NewACC(input.Index(21))
	v.Ub82 = NewUB1(input.Index(22))
	v.UniformBillingData = NewUB2(input.Index(23))
	v.PatientDeathAndAutopsy = NewPDA(input.Index(24))
	return v
}

type ADTA13 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	AccessRestriction                 []ARV
	Role                              []ROL
	NextOfKinAssociatedParties        []NK1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalAccessRestriction       []ARV
	AdditionalRole                    []ROL
	Disability                        []DB1
	ObservationResult                 []OBX
	PatientAllergyInformation         []AL1
	Diagnosis                         []DG1
	DiagnosisRelatedGroup             DRG
	Procedure                         []ADTA13Procedure
	Guarantor                         []GT1
	Insurance                         []ADTA13Insurance
	Accident                          ACC
	Ub82                              UB1
	UniformBillingData                UB2
	PatientDeathAndAutopsy            PDA
}

// NewADTA13Slice will construct a slice of type ADTA13
func NewADTA13Slice(input hl7.HL7Type) []ADTA13 {
	values := make([]ADTA13, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA13(input.Index(i))
	}
	return values
}

// NewCCFI22 creates an implementation of CCFI22
func NewCCFI22(input hl7.HL7Type) CCFI22 {
	v := CCFI22{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	return v
}

type CCFI22 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientIdentification         PID
}

// NewCCFI22Slice will construct a slice of type CCFI22
func NewCCFI22Slice(input hl7.HL7Type) []CCFI22 {
	values := make([]CCFI22, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCCFI22(input.Index(i))
	}
	return values
}

// NewMDMT01 creates an implementation of MDMT01
func NewMDMT01(input hl7.HL7Type) MDMT01 {
	v := MDMT01{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientVisit = NewPV1(input.Index(5))
	v.CommonOrder = NewMDMT01CommonOrderSlice(input.Index(6))
	v.TranscriptionDocumentHeader = NewTXA(input.Index(7))
	v.ConsentSegment = NewCONSlice(input.Index(8))
	return v
}

type MDMT01 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientVisit                  PV1
	CommonOrder                   []MDMT01CommonOrder
	TranscriptionDocumentHeader   TXA
	ConsentSegment                []CON
}

// NewMDMT01Slice will construct a slice of type MDMT01
func NewMDMT01Slice(input hl7.HL7Type) []MDMT01 {
	values := make([]MDMT01, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMDMT01(input.Index(i))
	}
	return values
}

// NewRRGO16 creates an implementation of RRGO16
func NewRRGO16(input hl7.HL7Type) RRGO16 {
	v := RRGO16{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewRRGO16Response(input.Index(6))
	return v
}

type RRGO16 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      RRGO16Response
}

// NewRRGO16Slice will construct a slice of type RRGO16
func NewRRGO16Slice(input hl7.HL7Type) []RRGO16 {
	values := make([]RRGO16, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRRGO16(input.Index(i))
	}
	return values
}

// NewADTA55 creates an implementation of ADTA55
func NewADTA55(input hl7.HL7Type) ADTA55 {
	v := ADTA55{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.Role = NewROLSlice(input.Index(6))
	v.PatientVisit = NewPV1(input.Index(7))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(8))
	v.AdditionalRole = NewROLSlice(input.Index(9))
	return v
}

type ADTA55 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	Role                              []ROL
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalRole                    []ROL
}

// NewADTA55Slice will construct a slice of type ADTA55
func NewADTA55Slice(input hl7.HL7Type) []ADTA55 {
	values := make([]ADTA55, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA55(input.Index(i))
	}
	return values
}

// NewQCNJ01 creates an implementation of QCNJ01
func NewQCNJ01(input hl7.HL7Type) QCNJ01 {
	v := QCNJ01{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryIdentification = NewQID(input.Index(3))
	return v
}

type QCNJ01 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryIdentification           QID
}

// NewQCNJ01Slice will construct a slice of type QCNJ01
func NewQCNJ01Slice(input hl7.HL7Type) []QCNJ01 {
	values := make([]QCNJ01, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQCNJ01(input.Index(i))
	}
	return values
}

// NewSRRS06 creates an implementation of SRRS06
func NewSRRS06(input hl7.HL7Type) SRRS06 {
	v := SRRS06{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.Schedule = NewSRRS06Schedule(input.Index(3))
	return v
}

type SRRS06 struct {
	MessageHeader         MSH
	MessageAcknowledgment MSA
	Error                 []ERR
	Schedule              SRRS06Schedule
}

// NewSRRS06Slice will construct a slice of type SRRS06
func NewSRRS06Slice(input hl7.HL7Type) []SRRS06 {
	values := make([]SRRS06, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRRS06(input.Index(i))
	}
	return values
}

// NewRQII01 creates an implementation of RQII01
func NewRQII01(input hl7.HL7Type) RQII01 {
	v := RQII01{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Provider = NewRQII01ProviderSlice(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(5))
	v.GuarantorInsurance = NewRQII01GuarantorInsurance(input.Index(6))
	v.NotesAndComments = NewNTESlice(input.Index(7))
	return v
}

type RQII01 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Provider                      []RQII01Provider
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	GuarantorInsurance            RQII01GuarantorInsurance
	NotesAndComments              []NTE
}

// NewRQII01Slice will construct a slice of type RQII01
func NewRQII01Slice(input hl7.HL7Type) []RQII01 {
	values := make([]RQII01, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRQII01(input.Index(i))
	}
	return values
}

// NewMFKM04 creates an implementation of MFKM04
func NewMFKM04(input hl7.HL7Type) MFKM04 {
	v := MFKM04{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.MasterFileIdentification = NewMFI(input.Index(5))
	v.MasterFileAcknowledgment = NewMFASlice(input.Index(6))
	return v
}

type MFKM04 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	MasterFileIdentification      MFI
	MasterFileAcknowledgment      []MFA
}

// NewMFKM04Slice will construct a slice of type MFKM04
func NewMFKM04Slice(input hl7.HL7Type) []MFKM04 {
	values := make([]MFKM04, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFKM04(input.Index(i))
	}
	return values
}

// NewSTCS33 creates an implementation of STCS33
func NewSTCS33(input hl7.HL7Type) STCS33 {
	v := STCS33{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.SterilizerConfigurationantiMicrobialDevices = NewSCPSlice(input.Index(3))
	return v
}

type STCS33 struct {
	MessageHeader                               MSH
	SoftwareSegment                             []SFT
	UserAuthenticationCredentials               UAC
	SterilizerConfigurationantiMicrobialDevices []SCP
}

// NewSTCS33Slice will construct a slice of type STCS33
func NewSTCS33Slice(input hl7.HL7Type) []STCS33 {
	values := make([]STCS33, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSTCS33(input.Index(i))
	}
	return values
}

// NewBARP01 creates an implementation of BARP01
func NewBARP01(input hl7.HL7Type) BARP01 {
	v := BARP01{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.ParticipationInformation = NewPRTSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.Visit = NewBARP01VisitSlice(input.Index(8))
	return v
}

type BARP01 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientAdditionalDemographic  PD1
	ParticipationInformation      []PRT
	Role                          []ROL
	Visit                         []BARP01Visit
}

// NewBARP01Slice will construct a slice of type BARP01
func NewBARP01Slice(input hl7.HL7Type) []BARP01 {
	values := make([]BARP01, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBARP01(input.Index(i))
	}
	return values
}

// NewSRMS06 creates an implementation of SRMS06
func NewSRMS06(input hl7.HL7Type) SRMS06 {
	v := SRMS06{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.AppointmentRequest = NewARQ(input.Index(1))
	v.AppointmentPreferences = NewAPR(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSRMS06PatientSlice(input.Index(4))
	v.Resources = NewSRMS06ResourcesSlice(input.Index(5))
	return v
}

type SRMS06 struct {
	MessageHeader          MSH
	AppointmentRequest     ARQ
	AppointmentPreferences APR
	NotesAndComments       []NTE
	Patient                []SRMS06Patient
	Resources              []SRMS06Resources
}

// NewSRMS06Slice will construct a slice of type SRMS06
func NewSRMS06Slice(input hl7.HL7Type) []SRMS06 {
	values := make([]SRMS06, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRMS06(input.Index(i))
	}
	return values
}

// NewMFNM13 creates an implementation of MFNM13
func NewMFNM13(input hl7.HL7Type) MFNM13 {
	v := MFNM13{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MasterFileIdentification = NewMFI(input.Index(3))
	v.MasterFileEntry = NewMFESlice(input.Index(4))
	return v
}

type MFNM13 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MasterFileIdentification      MFI
	MasterFileEntry               []MFE
}

// NewMFNM13Slice will construct a slice of type MFNM13
func NewMFNM13Slice(input hl7.HL7Type) []MFNM13 {
	values := make([]MFNM13, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFNM13(input.Index(i))
	}
	return values
}

// NewDPRO48 creates an implementation of DPRO48
func NewDPRO48(input hl7.HL7Type) DPRO48 {
	v := DPRO48{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.StaffIdentification = NewSTFSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Donnor = NewDPRO48Donnor(input.Index(3))
	v.DonnorOrder = NewDPRO48DonnorOrderSlice(input.Index(4))
	v.Donation = NewDPRO48Donation(input.Index(5))
	return v
}

type DPRO48 struct {
	MessageHeader                 MSH
	StaffIdentification           []STF
	UserAuthenticationCredentials UAC
	Donnor                        DPRO48Donnor
	DonnorOrder                   []DPRO48DonnorOrder
	Donation                      DPRO48Donation
}

// NewDPRO48Slice will construct a slice of type DPRO48
func NewDPRO48Slice(input hl7.HL7Type) []DPRO48 {
	values := make([]DPRO48, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDPRO48(input.Index(i))
	}
	return values
}

// NewRQII02 creates an implementation of RQII02
func NewRQII02(input hl7.HL7Type) RQII02 {
	v := RQII02{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Provider = NewRQII02ProviderSlice(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(5))
	v.GuarantorInsurance = NewRQII02GuarantorInsurance(input.Index(6))
	v.NotesAndComments = NewNTESlice(input.Index(7))
	return v
}

type RQII02 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Provider                      []RQII02Provider
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	GuarantorInsurance            RQII02GuarantorInsurance
	NotesAndComments              []NTE
}

// NewRQII02Slice will construct a slice of type RQII02
func NewRQII02Slice(input hl7.HL7Type) []RQII02 {
	values := make([]RQII02, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRQII02(input.Index(i))
	}
	return values
}

// NewQBPZ93 creates an implementation of QBPZ93
func NewQBPZ93(input hl7.HL7Type) QBPZ93 {
	v := QBPZ93{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.Qbp = NewQBPZ93Qbp(input.Index(4))
	v.TableRowDefinition = NewRDF(input.Index(5))
	v.ResponseControlParameter = NewRCP(input.Index(6))
	v.ContinuationPointer = NewDSC(input.Index(7))
	return v
}

type QBPZ93 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	Qbp                           QBPZ93Qbp
	TableRowDefinition            RDF
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPZ93Slice will construct a slice of type QBPZ93
func NewQBPZ93Slice(input hl7.HL7Type) []QBPZ93 {
	values := make([]QBPZ93, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPZ93(input.Index(i))
	}
	return values
}

// NewRSPE03 creates an implementation of RSPE03
func NewRSPE03(input hl7.HL7Type) RSPE03 {
	v := RSPE03{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUACSlice(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.QueryAckIpr = NewRSPE03QueryAckIpr(input.Index(5))
	return v
}

type RSPE03 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials []UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	QueryAckIpr                   RSPE03QueryAckIpr
}

// NewRSPE03Slice will construct a slice of type RSPE03
func NewRSPE03Slice(input hl7.HL7Type) []RSPE03 {
	values := make([]RSPE03, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPE03(input.Index(i))
	}
	return values
}

// NewADTA24 creates an implementation of ADTA24
func NewADTA24(input hl7.HL7Type) ADTA24 {
	v := ADTA24{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.Disability = NewDB1Slice(input.Index(7))
	v.AdditionalPatientIdentification = NewPID(input.Index(8))
	v.AdditionalPatientAdditionalDemographic = NewPD1(input.Index(9))
	v.AdditionalPatientVisit = NewPV1(input.Index(10))
	v.AdditionalDisability = NewDB1Slice(input.Index(11))
	return v
}

type ADTA24 struct {
	MessageHeader                          MSH
	SoftwareSegment                        []SFT
	UserAuthenticationCredentials          UAC
	EventType                              EVN
	PatientIdentification                  PID
	PatientAdditionalDemographic           PD1
	PatientVisit                           PV1
	Disability                             []DB1
	AdditionalPatientIdentification        PID
	AdditionalPatientAdditionalDemographic PD1
	AdditionalPatientVisit                 PV1
	AdditionalDisability                   []DB1
}

// NewADTA24Slice will construct a slice of type ADTA24
func NewADTA24Slice(input hl7.HL7Type) []ADTA24 {
	values := make([]ADTA24, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA24(input.Index(i))
	}
	return values
}

// NewMFKM12 creates an implementation of MFKM12
func NewMFKM12(input hl7.HL7Type) MFKM12 {
	v := MFKM12{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.MasterFileIdentification = NewMFI(input.Index(5))
	v.MasterFileAcknowledgment = NewMFASlice(input.Index(6))
	return v
}

type MFKM12 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	MasterFileIdentification      MFI
	MasterFileAcknowledgment      []MFA
}

// NewMFKM12Slice will construct a slice of type MFKM12
func NewMFKM12Slice(input hl7.HL7Type) []MFKM12 {
	values := make([]MFKM12, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFKM12(input.Index(i))
	}
	return values
}

// NewQBPZ75 creates an implementation of QBPZ75
func NewQBPZ75(input hl7.HL7Type) QBPZ75 {
	v := QBPZ75{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.Qbp = NewQBPZ75Qbp(input.Index(4))
	v.TableRowDefinition = NewRDF(input.Index(5))
	v.ResponseControlParameter = NewRCP(input.Index(6))
	v.ContinuationPointer = NewDSC(input.Index(7))
	return v
}

type QBPZ75 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	Qbp                           QBPZ75Qbp
	TableRowDefinition            RDF
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPZ75Slice will construct a slice of type QBPZ75
func NewQBPZ75Slice(input hl7.HL7Type) []QBPZ75 {
	values := make([]QBPZ75, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPZ75(input.Index(i))
	}
	return values
}

// NewSIUS18 creates an implementation of SIUS18
func NewSIUS18(input hl7.HL7Type) SIUS18 {
	v := SIUS18{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SchedulingActivityInformation = NewSCH(input.Index(1))
	v.TimingQuantity = NewTQ1Slice(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSIUS18PatientSlice(input.Index(4))
	v.Resources = NewSIUS18ResourcesSlice(input.Index(5))
	return v
}

type SIUS18 struct {
	MessageHeader                 MSH
	SchedulingActivityInformation SCH
	TimingQuantity                []TQ1
	NotesAndComments              []NTE
	Patient                       []SIUS18Patient
	Resources                     []SIUS18Resources
}

// NewSIUS18Slice will construct a slice of type SIUS18
func NewSIUS18Slice(input hl7.HL7Type) []SIUS18 {
	values := make([]SIUS18, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSIUS18(input.Index(i))
	}
	return values
}

// NewQBPZ99 creates an implementation of QBPZ99
func NewQBPZ99(input hl7.HL7Type) QBPZ99 {
	v := QBPZ99{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.Qbp = NewQBPZ99Qbp(input.Index(4))
	v.TableRowDefinition = NewRDF(input.Index(5))
	v.ResponseControlParameter = NewRCP(input.Index(6))
	v.ContinuationPointer = NewDSC(input.Index(7))
	return v
}

type QBPZ99 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	Qbp                           QBPZ99Qbp
	TableRowDefinition            RDF
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPZ99Slice will construct a slice of type QBPZ99
func NewQBPZ99Slice(input hl7.HL7Type) []QBPZ99 {
	values := make([]QBPZ99, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPZ99(input.Index(i))
	}
	return values
}

// NewRRII13 creates an implementation of RRII13
func NewRRII13(input hl7.HL7Type) RRII13 {
	v := RRII13{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.ReferralInformation = NewRF1(input.Index(4))
	v.AuthorizationContact = NewRRII13AuthorizationContact(input.Index(5))
	v.ProviderContact = NewRRII13ProviderContactSlice(input.Index(6))
	v.PatientIdentification = NewPID(input.Index(7))
	v.Accident = NewACC(input.Index(8))
	v.Diagnosis = NewDG1Slice(input.Index(9))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(10))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(11))
	v.Procedure = NewRRII13ProcedureSlice(input.Index(12))
	v.Observation = NewRRII13ObservationSlice(input.Index(13))
	v.PatientVisit = NewRRII13PatientVisit(input.Index(14))
	v.NotesAndComments = NewNTESlice(input.Index(15))
	return v
}

type RRII13 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	ReferralInformation           RF1
	AuthorizationContact          RRII13AuthorizationContact
	ProviderContact               []RRII13ProviderContact
	PatientIdentification         PID
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []RRII13Procedure
	Observation                   []RRII13Observation
	PatientVisit                  RRII13PatientVisit
	NotesAndComments              []NTE
}

// NewRRII13Slice will construct a slice of type RRII13
func NewRRII13Slice(input hl7.HL7Type) []RRII13 {
	values := make([]RRII13, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRRII13(input.Index(i))
	}
	return values
}

// NewCCMI21 creates an implementation of CCMI21
func NewCCMI21(input hl7.HL7Type) CCMI21 {
	v := CCMI21{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientAdditionalDemographic = NewPD1(input.Index(4))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(5))
	v.Insurance = NewCCMI21InsuranceSlice(input.Index(6))
	v.AppointmentHistory = NewCCMI21AppointmentHistorySlice(input.Index(7))
	v.ClinicalHistory = NewCCMI21ClinicalHistorySlice(input.Index(8))
	v.PatientVisits = NewCCMI21PatientVisitsSlice(input.Index(9))
	v.MedicationHistory = NewCCMI21MedicationHistorySlice(input.Index(10))
	v.Problem = NewCCMI21ProblemSlice(input.Index(11))
	v.Goal = NewCCMI21GoalSlice(input.Index(12))
	v.Pathway = NewCCMI21PathwaySlice(input.Index(13))
	v.ClinicalRelationshipSegment = NewRELSlice(input.Index(14))
	return v
}

type CCMI21 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientIdentification         PID
	PatientAdditionalDemographic  PD1
	NextOfKinAssociatedParties    []NK1
	Insurance                     []CCMI21Insurance
	AppointmentHistory            []CCMI21AppointmentHistory
	ClinicalHistory               []CCMI21ClinicalHistory
	PatientVisits                 []CCMI21PatientVisits
	MedicationHistory             []CCMI21MedicationHistory
	Problem                       []CCMI21Problem
	Goal                          []CCMI21Goal
	Pathway                       []CCMI21Pathway
	ClinicalRelationshipSegment   []REL
}

// NewCCMI21Slice will construct a slice of type CCMI21
func NewCCMI21Slice(input hl7.HL7Type) []CCMI21 {
	values := make([]CCMI21, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCCMI21(input.Index(i))
	}
	return values
}

// NewADTA10 creates an implementation of ADTA10
func NewADTA10(input hl7.HL7Type) ADTA10 {
	v := ADTA10{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.Disability = NewDB1Slice(input.Index(8))
	v.ObservationResult = NewOBXSlice(input.Index(9))
	return v
}

type ADTA10 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	Disability                        []DB1
	ObservationResult                 []OBX
}

// NewADTA10Slice will construct a slice of type ADTA10
func NewADTA10Slice(input hl7.HL7Type) []ADTA10 {
	values := make([]ADTA10, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA10(input.Index(i))
	}
	return values
}

// NewADTA29 creates an implementation of ADTA29
func NewADTA29(input hl7.HL7Type) ADTA29 {
	v := ADTA29{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.Disability = NewDB1Slice(input.Index(8))
	v.ObservationResult = NewOBXSlice(input.Index(9))
	return v
}

type ADTA29 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	Disability                        []DB1
	ObservationResult                 []OBX
}

// NewADTA29Slice will construct a slice of type ADTA29
func NewADTA29Slice(input hl7.HL7Type) []ADTA29 {
	values := make([]ADTA29, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA29(input.Index(i))
	}
	return values
}

// NewQBPQ33 creates an implementation of QBPQ33
func NewQBPQ33(input hl7.HL7Type) QBPQ33 {
	v := QBPQ33{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.ResponseControlParameter = NewRCP(input.Index(4))
	v.ContinuationPointer = NewDSC(input.Index(5))
	return v
}

type QBPQ33 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPQ33Slice will construct a slice of type QBPQ33
func NewQBPQ33Slice(input hl7.HL7Type) []QBPQ33 {
	values := make([]QBPQ33, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPQ33(input.Index(i))
	}
	return values
}

// NewORAR41 creates an implementation of ORAR41
func NewORAR41(input hl7.HL7Type) ORAR41 {
	v := ORAR41{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.ParticipationInformation = NewPRTSlice(input.Index(5))
	return v
}

type ORAR41 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	ParticipationInformation      []PRT
}

// NewORAR41Slice will construct a slice of type ORAR41
func NewORAR41Slice(input hl7.HL7Type) []ORAR41 {
	values := make([]ORAR41, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORAR41(input.Index(i))
	}
	return values
}

// NewMFKM16 creates an implementation of MFKM16
func NewMFKM16(input hl7.HL7Type) MFKM16 {
	v := MFKM16{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.MasterFileIdentification = NewMFI(input.Index(5))
	v.MasterFileAcknowledgment = NewMFASlice(input.Index(6))
	return v
}

type MFKM16 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	MasterFileIdentification      MFI
	MasterFileAcknowledgment      []MFA
}

// NewMFKM16Slice will construct a slice of type MFKM16
func NewMFKM16Slice(input hl7.HL7Type) []MFKM16 {
	values := make([]MFKM16, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFKM16(input.Index(i))
	}
	return values
}

// NewMFNM17 creates an implementation of MFNM17
func NewMFNM17(input hl7.HL7Type) MFNM17 {
	v := MFNM17{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MasterFileIdentification = NewMFI(input.Index(3))
	v.MfDrg = NewMFNM17MfDrgSlice(input.Index(4))
	return v
}

type MFNM17 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MasterFileIdentification      MFI
	MfDrg                         []MFNM17MfDrg
}

// NewMFNM17Slice will construct a slice of type MFNM17
func NewMFNM17Slice(input hl7.HL7Type) []MFNM17 {
	values := make([]MFNM17, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFNM17(input.Index(i))
	}
	return values
}

// NewSRMS08 creates an implementation of SRMS08
func NewSRMS08(input hl7.HL7Type) SRMS08 {
	v := SRMS08{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.AppointmentRequest = NewARQ(input.Index(1))
	v.AppointmentPreferences = NewAPR(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSRMS08PatientSlice(input.Index(4))
	v.Resources = NewSRMS08ResourcesSlice(input.Index(5))
	return v
}

type SRMS08 struct {
	MessageHeader          MSH
	AppointmentRequest     ARQ
	AppointmentPreferences APR
	NotesAndComments       []NTE
	Patient                []SRMS08Patient
	Resources              []SRMS08Resources
}

// NewSRMS08Slice will construct a slice of type SRMS08
func NewSRMS08Slice(input hl7.HL7Type) []SRMS08 {
	values := make([]SRMS08, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRMS08(input.Index(i))
	}
	return values
}

// NewSIUS14 creates an implementation of SIUS14
func NewSIUS14(input hl7.HL7Type) SIUS14 {
	v := SIUS14{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SchedulingActivityInformation = NewSCH(input.Index(1))
	v.TimingQuantity = NewTQ1Slice(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSIUS14PatientSlice(input.Index(4))
	v.Resources = NewSIUS14ResourcesSlice(input.Index(5))
	return v
}

type SIUS14 struct {
	MessageHeader                 MSH
	SchedulingActivityInformation SCH
	TimingQuantity                []TQ1
	NotesAndComments              []NTE
	Patient                       []SIUS14Patient
	Resources                     []SIUS14Resources
}

// NewSIUS14Slice will construct a slice of type SIUS14
func NewSIUS14Slice(input hl7.HL7Type) []SIUS14 {
	values := make([]SIUS14, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSIUS14(input.Index(i))
	}
	return values
}

// NewSRRS03 creates an implementation of SRRS03
func NewSRRS03(input hl7.HL7Type) SRRS03 {
	v := SRRS03{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.Schedule = NewSRRS03Schedule(input.Index(3))
	return v
}

type SRRS03 struct {
	MessageHeader         MSH
	MessageAcknowledgment MSA
	Error                 []ERR
	Schedule              SRRS03Schedule
}

// NewSRRS03Slice will construct a slice of type SRRS03
func NewSRRS03Slice(input hl7.HL7Type) []SRRS03 {
	values := make([]SRRS03, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRRS03(input.Index(i))
	}
	return values
}

// NewPMUB08 creates an implementation of PMUB08
func NewPMUB08(input hl7.HL7Type) PMUB08 {
	v := PMUB08{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.StaffIdentification = NewSTF(input.Index(4))
	v.PractitionerDetail = NewPRA(input.Index(5))
	v.CertificateDetail = NewCERSlice(input.Index(6))
	return v
}

type PMUB08 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	StaffIdentification           STF
	PractitionerDetail            PRA
	CertificateDetail             []CER
}

// NewPMUB08Slice will construct a slice of type PMUB08
func NewPMUB08Slice(input hl7.HL7Type) []PMUB08 {
	values := make([]PMUB08, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPMUB08(input.Index(i))
	}
	return values
}

// NewOMLO39 creates an implementation of OMLO39
func NewOMLO39(input hl7.HL7Type) OMLO39 {
	v := OMLO39{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewOMLO39Patient(input.Index(4))
	v.Order = NewOMLO39OrderSlice(input.Index(5))
	return v
}

type OMLO39 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       OMLO39Patient
	Order                         []OMLO39Order
}

// NewOMLO39Slice will construct a slice of type OMLO39
func NewOMLO39Slice(input hl7.HL7Type) []OMLO39 {
	values := make([]OMLO39, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOMLO39(input.Index(i))
	}
	return values
}

// NewADTA61 creates an implementation of ADTA61
func NewADTA61(input hl7.HL7Type) ADTA61 {
	v := ADTA61{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.Role = NewROLSlice(input.Index(6))
	v.PatientVisit = NewPV1(input.Index(7))
	v.AdditionalRole = NewROLSlice(input.Index(8))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(9))
	return v
}

type ADTA61 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	Role                              []ROL
	PatientVisit                      PV1
	AdditionalRole                    []ROL
	PatientVisitAdditionalInformation PV2
}

// NewADTA61Slice will construct a slice of type ADTA61
func NewADTA61Slice(input hl7.HL7Type) []ADTA61 {
	values := make([]ADTA61, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA61(input.Index(i))
	}
	return values
}

// NewADTA49 creates an implementation of ADTA49
func NewADTA49(input hl7.HL7Type) ADTA49 {
	v := ADTA49{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.Patient = NewADTA49PatientSlice(input.Index(4))
	return v
}

type ADTA49 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	Patient                       []ADTA49Patient
}

// NewADTA49Slice will construct a slice of type ADTA49
func NewADTA49Slice(input hl7.HL7Type) []ADTA49 {
	values := make([]ADTA49, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA49(input.Index(i))
	}
	return values
}

// NewSIUS27 creates an implementation of SIUS27
func NewSIUS27(input hl7.HL7Type) SIUS27 {
	v := SIUS27{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SchedulingActivityInformation = NewSCH(input.Index(1))
	v.TimingQuantity = NewTQ1Slice(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSIUS27PatientSlice(input.Index(4))
	v.Resources = NewSIUS27ResourcesSlice(input.Index(5))
	return v
}

type SIUS27 struct {
	MessageHeader                 MSH
	SchedulingActivityInformation SCH
	TimingQuantity                []TQ1
	NotesAndComments              []NTE
	Patient                       []SIUS27Patient
	Resources                     []SIUS27Resources
}

// NewSIUS27Slice will construct a slice of type SIUS27
func NewSIUS27Slice(input hl7.HL7Type) []SIUS27 {
	values := make([]SIUS27, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSIUS27(input.Index(i))
	}
	return values
}

// NewSIUS24 creates an implementation of SIUS24
func NewSIUS24(input hl7.HL7Type) SIUS24 {
	v := SIUS24{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SchedulingActivityInformation = NewSCH(input.Index(1))
	v.TimingQuantity = NewTQ1Slice(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSIUS24PatientSlice(input.Index(4))
	v.Resources = NewSIUS24ResourcesSlice(input.Index(5))
	return v
}

type SIUS24 struct {
	MessageHeader                 MSH
	SchedulingActivityInformation SCH
	TimingQuantity                []TQ1
	NotesAndComments              []NTE
	Patient                       []SIUS24Patient
	Resources                     []SIUS24Resources
}

// NewSIUS24Slice will construct a slice of type SIUS24
func NewSIUS24Slice(input hl7.HL7Type) []SIUS24 {
	values := make([]SIUS24, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSIUS24(input.Index(i))
	}
	return values
}

// NewCRMC06 creates an implementation of CRMC06
func NewCRMC06(input hl7.HL7Type) CRMC06 {
	v := CRMC06{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Patient = NewCRMC06PatientSlice(input.Index(3))
	return v
}

type CRMC06 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Patient                       []CRMC06Patient
}

// NewCRMC06Slice will construct a slice of type CRMC06
func NewCRMC06Slice(input hl7.HL7Type) []CRMC06 {
	values := make([]CRMC06, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCRMC06(input.Index(i))
	}
	return values
}

// NewMFKM05 creates an implementation of MFKM05
func NewMFKM05(input hl7.HL7Type) MFKM05 {
	v := MFKM05{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.MasterFileIdentification = NewMFI(input.Index(5))
	v.MasterFileAcknowledgment = NewMFASlice(input.Index(6))
	return v
}

type MFKM05 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	MasterFileIdentification      MFI
	MasterFileAcknowledgment      []MFA
}

// NewMFKM05Slice will construct a slice of type MFKM05
func NewMFKM05Slice(input hl7.HL7Type) []MFKM05 {
	values := make([]MFKM05, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFKM05(input.Index(i))
	}
	return values
}

// NewMFKM02 creates an implementation of MFKM02
func NewMFKM02(input hl7.HL7Type) MFKM02 {
	v := MFKM02{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.MasterFileIdentification = NewMFI(input.Index(5))
	v.MasterFileAcknowledgment = NewMFASlice(input.Index(6))
	return v
}

type MFKM02 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	MasterFileIdentification      MFI
	MasterFileAcknowledgment      []MFA
}

// NewMFKM02Slice will construct a slice of type MFKM02
func NewMFKM02Slice(input hl7.HL7Type) []MFKM02 {
	values := make([]MFKM02, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFKM02(input.Index(i))
	}
	return values
}

// NewSRRS01 creates an implementation of SRRS01
func NewSRRS01(input hl7.HL7Type) SRRS01 {
	v := SRRS01{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.Schedule = NewSRRS01Schedule(input.Index(3))
	return v
}

type SRRS01 struct {
	MessageHeader         MSH
	MessageAcknowledgment MSA
	Error                 []ERR
	Schedule              SRRS01Schedule
}

// NewSRRS01Slice will construct a slice of type SRRS01
func NewSRRS01Slice(input hl7.HL7Type) []SRRS01 {
	values := make([]SRRS01, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRRS01(input.Index(i))
	}
	return values
}

// NewRSPK23 creates an implementation of RSPK23
func NewRSPK23(input hl7.HL7Type) RSPK23 {
	v := RSPK23{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.QueryResponse = NewRSPK23QueryResponse(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RSPK23 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	QueryResponse                 RSPK23QueryResponse
	ContinuationPointer           DSC
}

// NewRSPK23Slice will construct a slice of type RSPK23
func NewRSPK23Slice(input hl7.HL7Type) []RSPK23 {
	values := make([]RSPK23, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPK23(input.Index(i))
	}
	return values
}

// NewMFKM15 creates an implementation of MFKM15
func NewMFKM15(input hl7.HL7Type) MFKM15 {
	v := MFKM15{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.MasterFileIdentification = NewMFI(input.Index(5))
	v.MasterFileAcknowledgment = NewMFASlice(input.Index(6))
	return v
}

type MFKM15 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	MasterFileIdentification      MFI
	MasterFileAcknowledgment      []MFA
}

// NewMFKM15Slice will construct a slice of type MFKM15
func NewMFKM15Slice(input hl7.HL7Type) []MFKM15 {
	values := make([]MFKM15, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFKM15(input.Index(i))
	}
	return values
}

// NewRPAI10 creates an implementation of RPAI10
func NewRPAI10(input hl7.HL7Type) RPAI10 {
	v := RPAI10{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.ReferralInformation = NewRF1(input.Index(4))
	v.Authorization1 = NewRPAI10Authorization1(input.Index(5))
	v.Provider = NewRPAI10ProviderSlice(input.Index(6))
	v.PatientIdentification = NewPID(input.Index(7))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(8))
	v.Guarantor = NewGT1Slice(input.Index(9))
	v.Insurance = NewRPAI10InsuranceSlice(input.Index(10))
	v.Accident = NewACC(input.Index(11))
	v.Diagnosis = NewDG1Slice(input.Index(12))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(13))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(14))
	v.Procedure = NewRPAI10ProcedureSlice(input.Index(15))
	v.Observation = NewRPAI10ObservationSlice(input.Index(16))
	v.Visit = NewRPAI10Visit(input.Index(17))
	v.NotesAndComments = NewNTESlice(input.Index(18))
	return v
}

type RPAI10 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	ReferralInformation           RF1
	Authorization1                RPAI10Authorization1
	Provider                      []RPAI10Provider
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	Guarantor                     []GT1
	Insurance                     []RPAI10Insurance
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []RPAI10Procedure
	Observation                   []RPAI10Observation
	Visit                         RPAI10Visit
	NotesAndComments              []NTE
}

// NewRPAI10Slice will construct a slice of type RPAI10
func NewRPAI10Slice(input hl7.HL7Type) []RPAI10 {
	values := make([]RPAI10, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRPAI10(input.Index(i))
	}
	return values
}

// NewPINI07 creates an implementation of PINI07
func NewPINI07(input hl7.HL7Type) PINI07 {
	v := PINI07{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Provider = NewPINI07ProviderSlice(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(5))
	v.GuarantorInsurance = NewPINI07GuarantorInsurance(input.Index(6))
	v.NotesAndComments = NewNTESlice(input.Index(7))
	return v
}

type PINI07 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Provider                      []PINI07Provider
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	GuarantorInsurance            PINI07GuarantorInsurance
	NotesAndComments              []NTE
}

// NewPINI07Slice will construct a slice of type PINI07
func NewPINI07Slice(input hl7.HL7Type) []PINI07 {
	values := make([]PINI07, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPINI07(input.Index(i))
	}
	return values
}

// NewRDYZ98 creates an implementation of RDYZ98
func NewRDYZ98(input hl7.HL7Type) RDYZ98 {
	v := RDYZ98{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.DisplayData = NewDSPSlice(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RDYZ98 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	DisplayData                   []DSP
	ContinuationPointer           DSC
}

// NewRDYZ98Slice will construct a slice of type RDYZ98
func NewRDYZ98Slice(input hl7.HL7Type) []RDYZ98 {
	values := make([]RDYZ98, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRDYZ98(input.Index(i))
	}
	return values
}

// NewADTA01 creates an implementation of ADTA01
func NewADTA01(input hl7.HL7Type) ADTA01 {
	v := ADTA01{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(8))
	v.PatientVisit = NewPV1(input.Index(9))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(10))
	v.AdditionalAccessRestriction = NewARVSlice(input.Index(11))
	v.AdditionalRole = NewROLSlice(input.Index(12))
	v.Disability = NewDB1Slice(input.Index(13))
	v.ObservationResult = NewOBXSlice(input.Index(14))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(15))
	v.Diagnosis = NewDG1Slice(input.Index(16))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(17))
	v.Procedure = NewADTA01ProcedureSlice(input.Index(18))
	v.Guarantor = NewGT1Slice(input.Index(19))
	v.Insurance = NewADTA01InsuranceSlice(input.Index(20))
	v.Accident = NewACC(input.Index(21))
	v.Ub82 = NewUB1(input.Index(22))
	v.UniformBillingData = NewUB2(input.Index(23))
	v.PatientDeathAndAutopsy = NewPDA(input.Index(24))
	return v
}

type ADTA01 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	AccessRestriction                 []ARV
	Role                              []ROL
	NextOfKinAssociatedParties        []NK1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalAccessRestriction       []ARV
	AdditionalRole                    []ROL
	Disability                        []DB1
	ObservationResult                 []OBX
	PatientAllergyInformation         []AL1
	Diagnosis                         []DG1
	DiagnosisRelatedGroup             DRG
	Procedure                         []ADTA01Procedure
	Guarantor                         []GT1
	Insurance                         []ADTA01Insurance
	Accident                          ACC
	Ub82                              UB1
	UniformBillingData                UB2
	PatientDeathAndAutopsy            PDA
}

// NewADTA01Slice will construct a slice of type ADTA01
func NewADTA01Slice(input hl7.HL7Type) []ADTA01 {
	values := make([]ADTA01, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA01(input.Index(i))
	}
	return values
}

// NewRDSO13 creates an implementation of RDSO13
func NewRDSO13(input hl7.HL7Type) RDSO13 {
	v := RDSO13{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewRDSO13Patient(input.Index(4))
	v.Order = NewRDSO13OrderSlice(input.Index(5))
	return v
}

type RDSO13 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       RDSO13Patient
	Order                         []RDSO13Order
}

// NewRDSO13Slice will construct a slice of type RDSO13
func NewRDSO13Slice(input hl7.HL7Type) []RDSO13 {
	values := make([]RDSO13, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRDSO13(input.Index(i))
	}
	return values
}

// NewRTBZ96 creates an implementation of RTBZ96
func NewRTBZ96(input hl7.HL7Type) RTBZ96 {
	v := RTBZ96{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.RowDefinition = NewRTBZ96RowDefinition(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RTBZ96 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	RowDefinition                 RTBZ96RowDefinition
	ContinuationPointer           DSC
}

// NewRTBZ96Slice will construct a slice of type RTBZ96
func NewRTBZ96Slice(input hl7.HL7Type) []RTBZ96 {
	values := make([]RTBZ96, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRTBZ96(input.Index(i))
	}
	return values
}

// NewCCRI17 creates an implementation of CCRI17
func NewCCRI17(input hl7.HL7Type) CCRI17 {
	v := CCRI17{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.ReferralInformation = NewRF1Slice(input.Index(3))
	v.ProviderContact = NewCCRI17ProviderContactSlice(input.Index(4))
	v.ClinicalOrder = NewCCRI17ClinicalOrderSlice(input.Index(5))
	v.Patient = NewCCRI17PatientSlice(input.Index(6))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(7))
	v.Insurance = NewCCRI17InsuranceSlice(input.Index(8))
	v.AppointmentHistory = NewCCRI17AppointmentHistorySlice(input.Index(9))
	v.ClinicalHistory = NewCCRI17ClinicalHistorySlice(input.Index(10))
	v.PatientVisits = NewCCRI17PatientVisitsSlice(input.Index(11))
	v.MedicationHistory = NewCCRI17MedicationHistorySlice(input.Index(12))
	v.Problem = NewCCRI17ProblemSlice(input.Index(13))
	v.Goal = NewCCRI17GoalSlice(input.Index(14))
	v.Pathway = NewCCRI17PathwaySlice(input.Index(15))
	v.ClinicalRelationshipSegment = NewRELSlice(input.Index(16))
	return v
}

type CCRI17 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	ReferralInformation           []RF1
	ProviderContact               []CCRI17ProviderContact
	ClinicalOrder                 []CCRI17ClinicalOrder
	Patient                       []CCRI17Patient
	NextOfKinAssociatedParties    []NK1
	Insurance                     []CCRI17Insurance
	AppointmentHistory            []CCRI17AppointmentHistory
	ClinicalHistory               []CCRI17ClinicalHistory
	PatientVisits                 []CCRI17PatientVisits
	MedicationHistory             []CCRI17MedicationHistory
	Problem                       []CCRI17Problem
	Goal                          []CCRI17Goal
	Pathway                       []CCRI17Pathway
	ClinicalRelationshipSegment   []REL
}

// NewCCRI17Slice will construct a slice of type CCRI17
func NewCCRI17Slice(input hl7.HL7Type) []CCRI17 {
	values := make([]CCRI17, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCCRI17(input.Index(i))
	}
	return values
}

// NewSIUS16 creates an implementation of SIUS16
func NewSIUS16(input hl7.HL7Type) SIUS16 {
	v := SIUS16{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SchedulingActivityInformation = NewSCH(input.Index(1))
	v.TimingQuantity = NewTQ1Slice(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSIUS16PatientSlice(input.Index(4))
	v.Resources = NewSIUS16ResourcesSlice(input.Index(5))
	return v
}

type SIUS16 struct {
	MessageHeader                 MSH
	SchedulingActivityInformation SCH
	TimingQuantity                []TQ1
	NotesAndComments              []NTE
	Patient                       []SIUS16Patient
	Resources                     []SIUS16Resources
}

// NewSIUS16Slice will construct a slice of type SIUS16
func NewSIUS16Slice(input hl7.HL7Type) []SIUS16 {
	values := make([]SIUS16, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSIUS16(input.Index(i))
	}
	return values
}

// NewEHCE20 creates an implementation of EHCE20
func NewEHCE20(input hl7.HL7Type) EHCE20 {
	v := EHCE20{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUACSlice(input.Index(2))
	v.AuthorizationRequest = NewEHCE20AuthorizationRequest(input.Index(3))
	return v
}

type EHCE20 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials []UAC
	AuthorizationRequest          EHCE20AuthorizationRequest
}

// NewEHCE20Slice will construct a slice of type EHCE20
func NewEHCE20Slice(input hl7.HL7Type) []EHCE20 {
	values := make([]EHCE20, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEHCE20(input.Index(i))
	}
	return values
}

// NewRRAO18 creates an implementation of RRAO18
func NewRRAO18(input hl7.HL7Type) RRAO18 {
	v := RRAO18{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewRRAO18Response(input.Index(6))
	return v
}

type RRAO18 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      RRAO18Response
}

// NewRRAO18Slice will construct a slice of type RRAO18
func NewRRAO18Slice(input hl7.HL7Type) []RRAO18 {
	values := make([]RRAO18, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRRAO18(input.Index(i))
	}
	return values
}

// NewSLNS34 creates an implementation of SLNS34
func NewSLNS34(input hl7.HL7Type) SLNS34 {
	v := SLNS34{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.SterilizationLot = NewSLTSlice(input.Index(3))
	return v
}

type SLNS34 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	SterilizationLot              []SLT
}

// NewSLNS34Slice will construct a slice of type SLNS34
func NewSLNS34Slice(input hl7.HL7Type) []SLNS34 {
	values := make([]SLNS34, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSLNS34(input.Index(i))
	}
	return values
}

// NewMFNM10 creates an implementation of MFNM10
func NewMFNM10(input hl7.HL7Type) MFNM10 {
	v := MFNM10{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MasterFileIdentification = NewMFI(input.Index(3))
	v.MfTestBatteries = NewMFNM10MfTestBatteriesSlice(input.Index(4))
	return v
}

type MFNM10 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MasterFileIdentification      MFI
	MfTestBatteries               []MFNM10MfTestBatteries
}

// NewMFNM10Slice will construct a slice of type MFNM10
func NewMFNM10Slice(input hl7.HL7Type) []MFNM10 {
	values := make([]MFNM10, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFNM10(input.Index(i))
	}
	return values
}

// NewRQII03 creates an implementation of RQII03
func NewRQII03(input hl7.HL7Type) RQII03 {
	v := RQII03{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Provider = NewRQII03ProviderSlice(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(5))
	v.GuarantorInsurance = NewRQII03GuarantorInsurance(input.Index(6))
	v.NotesAndComments = NewNTESlice(input.Index(7))
	return v
}

type RQII03 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Provider                      []RQII03Provider
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	GuarantorInsurance            RQII03GuarantorInsurance
	NotesAndComments              []NTE
}

// NewRQII03Slice will construct a slice of type RQII03
func NewRQII03Slice(input hl7.HL7Type) []RQII03 {
	values := make([]RQII03, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRQII03(input.Index(i))
	}
	return values
}

// NewPPPPCB creates an implementation of PPPPCB
func NewPPPPCB(input hl7.HL7Type) PPPPCB {
	v := PPPPCB{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientVisit = NewPPPPCBPatientVisit(input.Index(4))
	v.Pathway = NewPPPPCBPathwaySlice(input.Index(5))
	return v
}

type PPPPCB struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientIdentification         PID
	PatientVisit                  PPPPCBPatientVisit
	Pathway                       []PPPPCBPathway
}

// NewPPPPCBSlice will construct a slice of type PPPPCB
func NewPPPPCBSlice(input hl7.HL7Type) []PPPPCB {
	values := make([]PPPPCB, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPPPPCB(input.Index(i))
	}
	return values
}

// NewSIUS20 creates an implementation of SIUS20
func NewSIUS20(input hl7.HL7Type) SIUS20 {
	v := SIUS20{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SchedulingActivityInformation = NewSCH(input.Index(1))
	v.TimingQuantity = NewTQ1Slice(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSIUS20PatientSlice(input.Index(4))
	v.Resources = NewSIUS20ResourcesSlice(input.Index(5))
	return v
}

type SIUS20 struct {
	MessageHeader                 MSH
	SchedulingActivityInformation SCH
	TimingQuantity                []TQ1
	NotesAndComments              []NTE
	Patient                       []SIUS20Patient
	Resources                     []SIUS20Resources
}

// NewSIUS20Slice will construct a slice of type SIUS20
func NewSIUS20Slice(input hl7.HL7Type) []SIUS20 {
	values := make([]SIUS20, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSIUS20(input.Index(i))
	}
	return values
}

// NewMDMT04 creates an implementation of MDMT04
func NewMDMT04(input hl7.HL7Type) MDMT04 {
	v := MDMT04{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientVisit = NewPV1(input.Index(5))
	v.CommonOrder = NewMDMT04CommonOrderSlice(input.Index(6))
	v.TranscriptionDocumentHeader = NewTXA(input.Index(7))
	v.ConsentSegment = NewCONSlice(input.Index(8))
	v.Observation = NewMDMT04ObservationSlice(input.Index(9))
	return v
}

type MDMT04 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientVisit                  PV1
	CommonOrder                   []MDMT04CommonOrder
	TranscriptionDocumentHeader   TXA
	ConsentSegment                []CON
	Observation                   []MDMT04Observation
}

// NewMDMT04Slice will construct a slice of type MDMT04
func NewMDMT04Slice(input hl7.HL7Type) []MDMT04 {
	values := make([]MDMT04, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMDMT04(input.Index(i))
	}
	return values
}

// NewQBPQ21 creates an implementation of QBPQ21
func NewQBPQ21(input hl7.HL7Type) QBPQ21 {
	v := QBPQ21{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.ResponseControlParameter = NewRCP(input.Index(4))
	v.ContinuationPointer = NewDSC(input.Index(5))
	return v
}

type QBPQ21 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPQ21Slice will construct a slice of type QBPQ21
func NewQBPQ21Slice(input hl7.HL7Type) []QBPQ21 {
	values := make([]QBPQ21, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPQ21(input.Index(i))
	}
	return values
}

// NewMFNM11 creates an implementation of MFNM11
func NewMFNM11(input hl7.HL7Type) MFNM11 {
	v := MFNM11{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MasterFileIdentification = NewMFI(input.Index(3))
	v.MfTestCalculated = NewMFNM11MfTestCalculatedSlice(input.Index(4))
	return v
}

type MFNM11 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MasterFileIdentification      MFI
	MfTestCalculated              []MFNM11MfTestCalculated
}

// NewMFNM11Slice will construct a slice of type MFNM11
func NewMFNM11Slice(input hl7.HL7Type) []MFNM11 {
	values := make([]MFNM11, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFNM11(input.Index(i))
	}
	return values
}

// NewRQAI11 creates an implementation of RQAI11
func NewRQAI11(input hl7.HL7Type) RQAI11 {
	v := RQAI11{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.ReferralInformation = NewRF1(input.Index(3))
	v.Authorization = NewRQAI11Authorization(input.Index(4))
	v.Provider = NewRQAI11ProviderSlice(input.Index(5))
	v.PatientIdentification = NewPID(input.Index(6))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(7))
	v.GuarantorInsurance = NewRQAI11GuarantorInsurance(input.Index(8))
	v.Accident = NewACC(input.Index(9))
	v.Diagnosis = NewDG1Slice(input.Index(10))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(11))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(12))
	v.Procedure = NewRQAI11ProcedureSlice(input.Index(13))
	v.Observation = NewRQAI11ObservationSlice(input.Index(14))
	v.Visit = NewRQAI11Visit(input.Index(15))
	v.NotesAndComments = NewNTESlice(input.Index(16))
	return v
}

type RQAI11 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	ReferralInformation           RF1
	Authorization                 RQAI11Authorization
	Provider                      []RQAI11Provider
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	GuarantorInsurance            RQAI11GuarantorInsurance
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []RQAI11Procedure
	Observation                   []RQAI11Observation
	Visit                         RQAI11Visit
	NotesAndComments              []NTE
}

// NewRQAI11Slice will construct a slice of type RQAI11
func NewRQAI11Slice(input hl7.HL7Type) []RQAI11 {
	values := make([]RQAI11, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRQAI11(input.Index(i))
	}
	return values
}

// NewADTA62 creates an implementation of ADTA62
func NewADTA62(input hl7.HL7Type) ADTA62 {
	v := ADTA62{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.Role = NewROLSlice(input.Index(6))
	v.PatientVisit = NewPV1(input.Index(7))
	v.AdditionalRole = NewROLSlice(input.Index(8))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(9))
	return v
}

type ADTA62 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	Role                              []ROL
	PatientVisit                      PV1
	AdditionalRole                    []ROL
	PatientVisitAdditionalInformation PV2
}

// NewADTA62Slice will construct a slice of type ADTA62
func NewADTA62Slice(input hl7.HL7Type) []ADTA62 {
	values := make([]ADTA62, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA62(input.Index(i))
	}
	return values
}

// NewSRMS11 creates an implementation of SRMS11
func NewSRMS11(input hl7.HL7Type) SRMS11 {
	v := SRMS11{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.AppointmentRequest = NewARQ(input.Index(1))
	v.AppointmentPreferences = NewAPR(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSRMS11PatientSlice(input.Index(4))
	v.Resources = NewSRMS11ResourcesSlice(input.Index(5))
	return v
}

type SRMS11 struct {
	MessageHeader          MSH
	AppointmentRequest     ARQ
	AppointmentPreferences APR
	NotesAndComments       []NTE
	Patient                []SRMS11Patient
	Resources              []SRMS11Resources
}

// NewSRMS11Slice will construct a slice of type SRMS11
func NewSRMS11Slice(input hl7.HL7Type) []SRMS11 {
	values := make([]SRMS11, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRMS11(input.Index(i))
	}
	return values
}

// NewPEXP07 creates an implementation of PEXP07
func NewPEXP07(input hl7.HL7Type) PEXP07 {
	v := PEXP07{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.ParticipationInformation = NewPRTSlice(input.Index(6))
	v.AccessRestriction = NewARVSlice(input.Index(7))
	v.NotesAndComments = NewNTESlice(input.Index(8))
	v.Visit = NewPEXP07Visit(input.Index(9))
	v.Experience = NewPEXP07ExperienceSlice(input.Index(10))
	return v
}

type PEXP07 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientAdditionalDemographic  PD1
	ParticipationInformation      []PRT
	AccessRestriction             []ARV
	NotesAndComments              []NTE
	Visit                         PEXP07Visit
	Experience                    []PEXP07Experience
}

// NewPEXP07Slice will construct a slice of type PEXP07
func NewPEXP07Slice(input hl7.HL7Type) []PEXP07 {
	values := make([]PEXP07, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPEXP07(input.Index(i))
	}
	return values
}

// NewCCUI20 creates an implementation of CCUI20
func NewCCUI20(input hl7.HL7Type) CCUI20 {
	v := CCUI20{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.ReferralInformation = NewRF1(input.Index(3))
	v.ProviderContact = NewCCUI20ProviderContactSlice(input.Index(4))
	v.Patient = NewCCUI20PatientSlice(input.Index(5))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(6))
	v.Insurance = NewCCUI20InsuranceSlice(input.Index(7))
	v.AppointmentHistory = NewCCUI20AppointmentHistorySlice(input.Index(8))
	v.ClinicalHistory = NewCCUI20ClinicalHistorySlice(input.Index(9))
	v.PatientVisits = NewCCUI20PatientVisitsSlice(input.Index(10))
	v.MedicationHistory = NewCCUI20MedicationHistorySlice(input.Index(11))
	v.Problem = NewCCUI20ProblemSlice(input.Index(12))
	v.Goal = NewCCUI20GoalSlice(input.Index(13))
	v.Pathway = NewCCUI20PathwaySlice(input.Index(14))
	v.ClinicalRelationshipSegment = NewRELSlice(input.Index(15))
	return v
}

type CCUI20 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	ReferralInformation           RF1
	ProviderContact               []CCUI20ProviderContact
	Patient                       []CCUI20Patient
	NextOfKinAssociatedParties    []NK1
	Insurance                     []CCUI20Insurance
	AppointmentHistory            []CCUI20AppointmentHistory
	ClinicalHistory               []CCUI20ClinicalHistory
	PatientVisits                 []CCUI20PatientVisits
	MedicationHistory             []CCUI20MedicationHistory
	Problem                       []CCUI20Problem
	Goal                          []CCUI20Goal
	Pathway                       []CCUI20Pathway
	ClinicalRelationshipSegment   []REL
}

// NewCCUI20Slice will construct a slice of type CCUI20
func NewCCUI20Slice(input hl7.HL7Type) []CCUI20 {
	values := make([]CCUI20, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCCUI20(input.Index(i))
	}
	return values
}

// NewPPRPC1 creates an implementation of PPRPC1
func NewPPRPC1(input hl7.HL7Type) PPRPC1 {
	v := PPRPC1{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientVisit = NewPPRPC1PatientVisit(input.Index(4))
	v.Problem = NewPPRPC1ProblemSlice(input.Index(5))
	return v
}

type PPRPC1 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientIdentification         PID
	PatientVisit                  PPRPC1PatientVisit
	Problem                       []PPRPC1Problem
}

// NewPPRPC1Slice will construct a slice of type PPRPC1
func NewPPRPC1Slice(input hl7.HL7Type) []PPRPC1 {
	values := make([]PPRPC1, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPPRPC1(input.Index(i))
	}
	return values
}

// NewSRMS04 creates an implementation of SRMS04
func NewSRMS04(input hl7.HL7Type) SRMS04 {
	v := SRMS04{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.AppointmentRequest = NewARQ(input.Index(1))
	v.AppointmentPreferences = NewAPR(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSRMS04PatientSlice(input.Index(4))
	v.Resources = NewSRMS04ResourcesSlice(input.Index(5))
	return v
}

type SRMS04 struct {
	MessageHeader          MSH
	AppointmentRequest     ARQ
	AppointmentPreferences APR
	NotesAndComments       []NTE
	Patient                []SRMS04Patient
	Resources              []SRMS04Resources
}

// NewSRMS04Slice will construct a slice of type SRMS04
func NewSRMS04Slice(input hl7.HL7Type) []SRMS04 {
	values := make([]SRMS04, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRMS04(input.Index(i))
	}
	return values
}

// NewRSPK25 creates an implementation of RSPK25
func NewRSPK25(input hl7.HL7Type) RSPK25 {
	v := RSPK25{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.ResponseControlParameter = NewRCP(input.Index(7))
	v.Staff = NewRSPK25StaffSlice(input.Index(8))
	v.ContinuationPointer = NewDSC(input.Index(9))
	return v
}

type RSPK25 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
	Staff                         []RSPK25Staff
	ContinuationPointer           DSC
}

// NewRSPK25Slice will construct a slice of type RSPK25
func NewRSPK25Slice(input hl7.HL7Type) []RSPK25 {
	values := make([]RSPK25, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPK25(input.Index(i))
	}
	return values
}

// NewSSRU04 creates an implementation of SSRU04
func NewSSRU04(input hl7.HL7Type) SSRU04 {
	v := SSRU04{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EquipmentDetail = NewEQU(input.Index(3))
	v.SpecimenContainer = NewSSRU04SpecimenContainerSlice(input.Index(4))
	return v
}

type SSRU04 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EquipmentDetail               EQU
	SpecimenContainer             []SSRU04SpecimenContainer
}

// NewSSRU04Slice will construct a slice of type SSRU04
func NewSSRU04Slice(input hl7.HL7Type) []SSRU04 {
	values := make([]SSRU04, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSSRU04(input.Index(i))
	}
	return values
}

// NewSIUS26 creates an implementation of SIUS26
func NewSIUS26(input hl7.HL7Type) SIUS26 {
	v := SIUS26{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SchedulingActivityInformation = NewSCH(input.Index(1))
	v.TimingQuantity = NewTQ1Slice(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSIUS26PatientSlice(input.Index(4))
	v.Resources = NewSIUS26ResourcesSlice(input.Index(5))
	return v
}

type SIUS26 struct {
	MessageHeader                 MSH
	SchedulingActivityInformation SCH
	TimingQuantity                []TQ1
	NotesAndComments              []NTE
	Patient                       []SIUS26Patient
	Resources                     []SIUS26Resources
}

// NewSIUS26Slice will construct a slice of type SIUS26
func NewSIUS26Slice(input hl7.HL7Type) []SIUS26 {
	values := make([]SIUS26, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSIUS26(input.Index(i))
	}
	return values
}

// NewORIO24 creates an implementation of ORIO24
func NewORIO24(input hl7.HL7Type) ORIO24 {
	v := ORIO24{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewORIO24Response(input.Index(6))
	return v
}

type ORIO24 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      ORIO24Response
}

// NewORIO24Slice will construct a slice of type ORIO24
func NewORIO24Slice(input hl7.HL7Type) []ORIO24 {
	values := make([]ORIO24, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORIO24(input.Index(i))
	}
	return values
}

// NewRSPZ86 creates an implementation of RSPZ86
func NewRSPZ86(input hl7.HL7Type) RSPZ86 {
	v := RSPZ86{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.QueryResponse = NewRSPZ86QueryResponseSlice(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RSPZ86 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	QueryResponse                 []RSPZ86QueryResponse
	ContinuationPointer           DSC
}

// NewRSPZ86Slice will construct a slice of type RSPZ86
func NewRSPZ86Slice(input hl7.HL7Type) []RSPZ86 {
	values := make([]RSPZ86, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPZ86(input.Index(i))
	}
	return values
}

// NewMFKM07 creates an implementation of MFKM07
func NewMFKM07(input hl7.HL7Type) MFKM07 {
	v := MFKM07{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.MasterFileIdentification = NewMFI(input.Index(5))
	v.MasterFileAcknowledgment = NewMFASlice(input.Index(6))
	return v
}

type MFKM07 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	MasterFileIdentification      MFI
	MasterFileAcknowledgment      []MFA
}

// NewMFKM07Slice will construct a slice of type MFKM07
func NewMFKM07Slice(input hl7.HL7Type) []MFKM07 {
	values := make([]MFKM07, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFKM07(input.Index(i))
	}
	return values
}

// NewRDEO25 creates an implementation of RDEO25
func NewRDEO25(input hl7.HL7Type) RDEO25 {
	v := RDEO25{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewRDEO25Patient(input.Index(4))
	v.Order = NewRDEO25OrderSlice(input.Index(5))
	return v
}

type RDEO25 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       RDEO25Patient
	Order                         []RDEO25Order
}

// NewRDEO25Slice will construct a slice of type RDEO25
func NewRDEO25Slice(input hl7.HL7Type) []RDEO25 {
	values := make([]RDEO25, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRDEO25(input.Index(i))
	}
	return values
}

// NewADTA23 creates an implementation of ADTA23
func NewADTA23(input hl7.HL7Type) ADTA23 {
	v := ADTA23{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.Disability = NewDB1Slice(input.Index(8))
	v.ObservationResult = NewOBXSlice(input.Index(9))
	return v
}

type ADTA23 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	Disability                        []DB1
	ObservationResult                 []OBX
}

// NewADTA23Slice will construct a slice of type ADTA23
func NewADTA23Slice(input hl7.HL7Type) []ADTA23 {
	values := make([]ADTA23, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA23(input.Index(i))
	}
	return values
}

// NewSDNS36 creates an implementation of SDNS36
func NewSDNS36(input hl7.HL7Type) SDNS36 {
	v := SDNS36{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.AntiMicrobialDeviceData = NewSDNS36AntiMicrobialDeviceData(input.Index(3))
	return v
}

type SDNS36 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	AntiMicrobialDeviceData       SDNS36AntiMicrobialDeviceData
}

// NewSDNS36Slice will construct a slice of type SDNS36
func NewSDNS36Slice(input hl7.HL7Type) []SDNS36 {
	values := make([]SDNS36, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSDNS36(input.Index(i))
	}
	return values
}

// NewDBUO42 creates an implementation of DBUO42
func NewDBUO42(input hl7.HL7Type) DBUO42 {
	v := DBUO42{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.StaffIdentification = NewSTFSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Donnor = NewDBUO42Donnor(input.Index(3))
	return v
}

type DBUO42 struct {
	MessageHeader                 MSH
	StaffIdentification           []STF
	UserAuthenticationCredentials UAC
	Donnor                        DBUO42Donnor
}

// NewDBUO42Slice will construct a slice of type DBUO42
func NewDBUO42Slice(input hl7.HL7Type) []DBUO42 {
	values := make([]DBUO42, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDBUO42(input.Index(i))
	}
	return values
}

// NewDEOO45 creates an implementation of DEOO45
func NewDEOO45(input hl7.HL7Type) DEOO45 {
	v := DEOO45{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.StaffIdentification = NewSTFSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Donnor = NewDEOO45Donnor(input.Index(3))
	v.DonnorOrder = NewDEOO45DonnorOrderSlice(input.Index(4))
	return v
}

type DEOO45 struct {
	MessageHeader                 MSH
	StaffIdentification           []STF
	UserAuthenticationCredentials UAC
	Donnor                        DEOO45Donnor
	DonnorOrder                   []DEOO45DonnorOrder
}

// NewDEOO45Slice will construct a slice of type DEOO45
func NewDEOO45Slice(input hl7.HL7Type) []DEOO45 {
	values := make([]DEOO45, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDEOO45(input.Index(i))
	}
	return values
}

// NewMDMT03 creates an implementation of MDMT03
func NewMDMT03(input hl7.HL7Type) MDMT03 {
	v := MDMT03{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientVisit = NewPV1(input.Index(5))
	v.CommonOrder = NewMDMT03CommonOrderSlice(input.Index(6))
	v.TranscriptionDocumentHeader = NewTXA(input.Index(7))
	v.ConsentSegment = NewCONSlice(input.Index(8))
	return v
}

type MDMT03 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientVisit                  PV1
	CommonOrder                   []MDMT03CommonOrder
	TranscriptionDocumentHeader   TXA
	ConsentSegment                []CON
}

// NewMDMT03Slice will construct a slice of type MDMT03
func NewMDMT03Slice(input hl7.HL7Type) []MDMT03 {
	values := make([]MDMT03, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMDMT03(input.Index(i))
	}
	return values
}

// NewSIUS13 creates an implementation of SIUS13
func NewSIUS13(input hl7.HL7Type) SIUS13 {
	v := SIUS13{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SchedulingActivityInformation = NewSCH(input.Index(1))
	v.TimingQuantity = NewTQ1Slice(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSIUS13PatientSlice(input.Index(4))
	v.Resources = NewSIUS13ResourcesSlice(input.Index(5))
	return v
}

type SIUS13 struct {
	MessageHeader                 MSH
	SchedulingActivityInformation SCH
	TimingQuantity                []TQ1
	NotesAndComments              []NTE
	Patient                       []SIUS13Patient
	Resources                     []SIUS13Resources
}

// NewSIUS13Slice will construct a slice of type SIUS13
func NewSIUS13Slice(input hl7.HL7Type) []SIUS13 {
	values := make([]SIUS13, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSIUS13(input.Index(i))
	}
	return values
}

// NewQBPZ89 creates an implementation of QBPZ89
func NewQBPZ89(input hl7.HL7Type) QBPZ89 {
	v := QBPZ89{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.Qbp = NewQBPZ89Qbp(input.Index(4))
	v.ResponseControlParameter = NewRCP(input.Index(5))
	v.ContinuationPointer = NewDSC(input.Index(6))
	return v
}

type QBPZ89 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	Qbp                           QBPZ89Qbp
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPZ89Slice will construct a slice of type QBPZ89
func NewQBPZ89Slice(input hl7.HL7Type) []QBPZ89 {
	values := make([]QBPZ89, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPZ89(input.Index(i))
	}
	return values
}

// NewOSMR26 creates an implementation of OSMR26
func NewOSMR26(input hl7.HL7Type) OSMR26 {
	v := OSMR26{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Shipment = NewOSMR26ShipmentSlice(input.Index(3))
	return v
}

type OSMR26 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Shipment                      []OSMR26Shipment
}

// NewOSMR26Slice will construct a slice of type OSMR26
func NewOSMR26Slice(input hl7.HL7Type) []OSMR26 {
	values := make([]OSMR26, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOSMR26(input.Index(i))
	}
	return values
}

// NewSRMS03 creates an implementation of SRMS03
func NewSRMS03(input hl7.HL7Type) SRMS03 {
	v := SRMS03{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.AppointmentRequest = NewARQ(input.Index(1))
	v.AppointmentPreferences = NewAPR(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSRMS03PatientSlice(input.Index(4))
	v.Resources = NewSRMS03ResourcesSlice(input.Index(5))
	return v
}

type SRMS03 struct {
	MessageHeader          MSH
	AppointmentRequest     ARQ
	AppointmentPreferences APR
	NotesAndComments       []NTE
	Patient                []SRMS03Patient
	Resources              []SRMS03Resources
}

// NewSRMS03Slice will construct a slice of type SRMS03
func NewSRMS03Slice(input hl7.HL7Type) []SRMS03 {
	values := make([]SRMS03, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRMS03(input.Index(i))
	}
	return values
}

// NewORPO10 creates an implementation of ORPO10
func NewORPO10(input hl7.HL7Type) ORPO10 {
	v := ORPO10{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewORPO10Response(input.Index(6))
	return v
}

type ORPO10 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      ORPO10Response
}

// NewORPO10Slice will construct a slice of type ORPO10
func NewORPO10Slice(input hl7.HL7Type) []ORPO10 {
	values := make([]ORPO10, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORPO10(input.Index(i))
	}
	return values
}

// NewOMPO09 creates an implementation of OMPO09
func NewOMPO09(input hl7.HL7Type) OMPO09 {
	v := OMPO09{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewOMPO09Patient(input.Index(4))
	v.Order = NewOMPO09OrderSlice(input.Index(5))
	return v
}

type OMPO09 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       OMPO09Patient
	Order                         []OMPO09Order
}

// NewOMPO09Slice will construct a slice of type OMPO09
func NewOMPO09Slice(input hl7.HL7Type) []OMPO09 {
	values := make([]OMPO09, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOMPO09(input.Index(i))
	}
	return values
}

// NewPPRPC3 creates an implementation of PPRPC3
func NewPPRPC3(input hl7.HL7Type) PPRPC3 {
	v := PPRPC3{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientVisit = NewPPRPC3PatientVisit(input.Index(4))
	v.Problem = NewPPRPC3ProblemSlice(input.Index(5))
	return v
}

type PPRPC3 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientIdentification         PID
	PatientVisit                  PPRPC3PatientVisit
	Problem                       []PPRPC3Problem
}

// NewPPRPC3Slice will construct a slice of type PPRPC3
func NewPPRPC3Slice(input hl7.HL7Type) []PPRPC3 {
	values := make([]PPRPC3, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPPRPC3(input.Index(i))
	}
	return values
}

// NewSRRS05 creates an implementation of SRRS05
func NewSRRS05(input hl7.HL7Type) SRRS05 {
	v := SRRS05{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.Schedule = NewSRRS05Schedule(input.Index(3))
	return v
}

type SRRS05 struct {
	MessageHeader         MSH
	MessageAcknowledgment MSA
	Error                 []ERR
	Schedule              SRRS05Schedule
}

// NewSRRS05Slice will construct a slice of type SRRS05
func NewSRRS05Slice(input hl7.HL7Type) []SRRS05 {
	values := make([]SRRS05, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRRS05(input.Index(i))
	}
	return values
}

// NewADTA38 creates an implementation of ADTA38
func NewADTA38(input hl7.HL7Type) ADTA38 {
	v := ADTA38{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.Disability = NewDB1Slice(input.Index(8))
	v.ObservationResult = NewOBXSlice(input.Index(9))
	v.Diagnosis = NewDG1Slice(input.Index(10))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(11))
	return v
}

type ADTA38 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	Disability                        []DB1
	ObservationResult                 []OBX
	Diagnosis                         []DG1
	DiagnosisRelatedGroup             DRG
}

// NewADTA38Slice will construct a slice of type ADTA38
func NewADTA38Slice(input hl7.HL7Type) []ADTA38 {
	values := make([]ADTA38, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA38(input.Index(i))
	}
	return values
}

// NewPMUB02 creates an implementation of PMUB02
func NewPMUB02(input hl7.HL7Type) PMUB02 {
	v := PMUB02{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.StaffIdentification = NewSTF(input.Index(4))
	v.PractitionerDetail = NewPRASlice(input.Index(5))
	v.PractitionerOrganizationUnit = NewORGSlice(input.Index(6))
	v.ProfessionalAffiliation = NewAFFSlice(input.Index(7))
	v.LanguageDetail = NewLANSlice(input.Index(8))
	v.EducationalDetail = NewEDUSlice(input.Index(9))
	v.CertificateDetail = NewCERSlice(input.Index(10))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(11))
	v.ParticipationInformation = NewPRTSlice(input.Index(12))
	v.Role = NewROLSlice(input.Index(13))
	return v
}

type PMUB02 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	StaffIdentification           STF
	PractitionerDetail            []PRA
	PractitionerOrganizationUnit  []ORG
	ProfessionalAffiliation       []AFF
	LanguageDetail                []LAN
	EducationalDetail             []EDU
	CertificateDetail             []CER
	NextOfKinAssociatedParties    []NK1
	ParticipationInformation      []PRT
	Role                          []ROL
}

// NewPMUB02Slice will construct a slice of type PMUB02
func NewPMUB02Slice(input hl7.HL7Type) []PMUB02 {
	values := make([]PMUB02, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPMUB02(input.Index(i))
	}
	return values
}

// NewRSPK11 creates an implementation of RSPK11
func NewRSPK11(input hl7.HL7Type) RSPK11 {
	v := RSPK11{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.SegmentPattern = NewRSPK11SegmentPattern(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RSPK11 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	SegmentPattern                RSPK11SegmentPattern
	ContinuationPointer           DSC
}

// NewRSPK11Slice will construct a slice of type RSPK11
func NewRSPK11Slice(input hl7.HL7Type) []RSPK11 {
	values := make([]RSPK11, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPK11(input.Index(i))
	}
	return values
}

// NewORUR01 creates an implementation of ORUR01
func NewORUR01(input hl7.HL7Type) ORUR01 {
	v := ORUR01{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.Software = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientResult = NewORUR01PatientResultSlice(input.Index(3))
	v.ContinuationPointer = NewDSC(input.Index(4))
	return v
}

type ORUR01 struct {
	MessageHeader                 MSH
	Software                      []SFT
	UserAuthenticationCredentials UAC
	PatientResult                 []ORUR01PatientResult
	ContinuationPointer           DSC
}

// NewORUR01Slice will construct a slice of type ORUR01
func NewORUR01Slice(input hl7.HL7Type) []ORUR01 {
	values := make([]ORUR01, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORUR01(input.Index(i))
	}
	return values
}

// NewOPRO38 creates an implementation of OPRO38
func NewOPRO38(input hl7.HL7Type) OPRO38 {
	v := OPRO38{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewOPRO38Response(input.Index(6))
	return v
}

type OPRO38 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      OPRO38Response
}

// NewOPRO38Slice will construct a slice of type OPRO38
func NewOPRO38Slice(input hl7.HL7Type) []OPRO38 {
	values := make([]OPRO38, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOPRO38(input.Index(i))
	}
	return values
}

// NewQBPE22 creates an implementation of QBPE22
func NewQBPE22(input hl7.HL7Type) QBPE22 {
	v := QBPE22{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUACSlice(input.Index(2))
	v.Query = NewQBPE22Query(input.Index(3))
	return v
}

type QBPE22 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials []UAC
	Query                         QBPE22Query
}

// NewQBPE22Slice will construct a slice of type QBPE22
func NewQBPE22Slice(input hl7.HL7Type) []QBPE22 {
	values := make([]QBPE22, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPE22(input.Index(i))
	}
	return values
}

// NewBARP05 creates an implementation of BARP05
func NewBARP05(input hl7.HL7Type) BARP05 {
	v := BARP05{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.ParticipationInformation = NewPRTSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.Visit = NewBARP05VisitSlice(input.Index(8))
	return v
}

type BARP05 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientAdditionalDemographic  PD1
	ParticipationInformation      []PRT
	Role                          []ROL
	Visit                         []BARP05Visit
}

// NewBARP05Slice will construct a slice of type BARP05
func NewBARP05Slice(input hl7.HL7Type) []BARP05 {
	values := make([]BARP05, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBARP05(input.Index(i))
	}
	return values
}

// NewADTA47 creates an implementation of ADTA47
func NewADTA47(input hl7.HL7Type) ADTA47 {
	v := ADTA47{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.Patient = NewADTA47PatientSlice(input.Index(4))
	return v
}

type ADTA47 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	Patient                       []ADTA47Patient
}

// NewADTA47Slice will construct a slice of type ADTA47
func NewADTA47Slice(input hl7.HL7Type) []ADTA47 {
	values := make([]ADTA47, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA47(input.Index(i))
	}
	return values
}

// NewADTA40 creates an implementation of ADTA40
func NewADTA40(input hl7.HL7Type) ADTA40 {
	v := ADTA40{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.Patient = NewADTA40PatientSlice(input.Index(4))
	return v
}

type ADTA40 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	Patient                       []ADTA40Patient
}

// NewADTA40Slice will construct a slice of type ADTA40
func NewADTA40Slice(input hl7.HL7Type) []ADTA40 {
	values := make([]ADTA40, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA40(input.Index(i))
	}
	return values
}

// NewMDMT10 creates an implementation of MDMT10
func NewMDMT10(input hl7.HL7Type) MDMT10 {
	v := MDMT10{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientVisit = NewPV1(input.Index(5))
	v.CommonOrder = NewMDMT10CommonOrderSlice(input.Index(6))
	v.TranscriptionDocumentHeader = NewTXA(input.Index(7))
	v.ConsentSegment = NewCONSlice(input.Index(8))
	v.Observation = NewMDMT10ObservationSlice(input.Index(9))
	return v
}

type MDMT10 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientVisit                  PV1
	CommonOrder                   []MDMT10CommonOrder
	TranscriptionDocumentHeader   TXA
	ConsentSegment                []CON
	Observation                   []MDMT10Observation
}

// NewMDMT10Slice will construct a slice of type MDMT10
func NewMDMT10Slice(input hl7.HL7Type) []MDMT10 {
	values := make([]MDMT10, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMDMT10(input.Index(i))
	}
	return values
}

// NewRQAI10 creates an implementation of RQAI10
func NewRQAI10(input hl7.HL7Type) RQAI10 {
	v := RQAI10{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.ReferralInformation = NewRF1(input.Index(3))
	v.Authorization = NewRQAI10Authorization(input.Index(4))
	v.Provider = NewRQAI10ProviderSlice(input.Index(5))
	v.PatientIdentification = NewPID(input.Index(6))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(7))
	v.GuarantorInsurance = NewRQAI10GuarantorInsurance(input.Index(8))
	v.Accident = NewACC(input.Index(9))
	v.Diagnosis = NewDG1Slice(input.Index(10))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(11))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(12))
	v.Procedure = NewRQAI10ProcedureSlice(input.Index(13))
	v.Observation = NewRQAI10ObservationSlice(input.Index(14))
	v.Visit = NewRQAI10Visit(input.Index(15))
	v.NotesAndComments = NewNTESlice(input.Index(16))
	return v
}

type RQAI10 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	ReferralInformation           RF1
	Authorization                 RQAI10Authorization
	Provider                      []RQAI10Provider
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	GuarantorInsurance            RQAI10GuarantorInsurance
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []RQAI10Procedure
	Observation                   []RQAI10Observation
	Visit                         RQAI10Visit
	NotesAndComments              []NTE
}

// NewRQAI10Slice will construct a slice of type RQAI10
func NewRQAI10Slice(input hl7.HL7Type) []RQAI10 {
	values := make([]RQAI10, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRQAI10(input.Index(i))
	}
	return values
}

// NewORLO34 creates an implementation of ORLO34
func NewORLO34(input hl7.HL7Type) ORLO34 {
	v := ORLO34{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewORLO34Response(input.Index(6))
	return v
}

type ORLO34 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      ORLO34Response
}

// NewORLO34Slice will construct a slice of type ORLO34
func NewORLO34Slice(input hl7.HL7Type) []ORLO34 {
	values := make([]ORLO34, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORLO34(input.Index(i))
	}
	return values
}

// NewDRGO43 creates an implementation of DRGO43
func NewDRGO43(input hl7.HL7Type) DRGO43 {
	v := DRGO43{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.StaffIdentification = NewSTFSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Donnor = NewDRGO43Donnor(input.Index(3))
	return v
}

type DRGO43 struct {
	MessageHeader                 MSH
	StaffIdentification           []STF
	UserAuthenticationCredentials UAC
	Donnor                        DRGO43Donnor
}

// NewDRGO43Slice will construct a slice of type DRGO43
func NewDRGO43Slice(input hl7.HL7Type) []DRGO43 {
	values := make([]DRGO43, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDRGO43(input.Index(i))
	}
	return values
}

// NewORBO28 creates an implementation of ORBO28
func NewORBO28(input hl7.HL7Type) ORBO28 {
	v := ORBO28{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewORBO28Response(input.Index(6))
	return v
}

type ORBO28 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      ORBO28Response
}

// NewORBO28Slice will construct a slice of type ORBO28
func NewORBO28Slice(input hl7.HL7Type) []ORBO28 {
	values := make([]ORBO28, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORBO28(input.Index(i))
	}
	return values
}

// NewEHCE21 creates an implementation of EHCE21
func NewEHCE21(input hl7.HL7Type) EHCE21 {
	v := EHCE21{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUACSlice(input.Index(2))
	v.AuthorizationRequest = NewEHCE21AuthorizationRequest(input.Index(3))
	return v
}

type EHCE21 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials []UAC
	AuthorizationRequest          EHCE21AuthorizationRequest
}

// NewEHCE21Slice will construct a slice of type EHCE21
func NewEHCE21Slice(input hl7.HL7Type) []EHCE21 {
	values := make([]EHCE21, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEHCE21(input.Index(i))
	}
	return values
}

// NewOMBO27 creates an implementation of OMBO27
func NewOMBO27(input hl7.HL7Type) OMBO27 {
	v := OMBO27{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewOMBO27Patient(input.Index(4))
	v.Order = NewOMBO27OrderSlice(input.Index(5))
	return v
}

type OMBO27 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       OMBO27Patient
	Order                         []OMBO27Order
}

// NewOMBO27Slice will construct a slice of type OMBO27
func NewOMBO27Slice(input hl7.HL7Type) []OMBO27 {
	values := make([]OMBO27, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOMBO27(input.Index(i))
	}
	return values
}

// NewACK creates an implementation of ACK
func NewACK(input hl7.HL7Type) ACK {
	v := ACK{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	return v
}

type ACK struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
}

// NewACKSlice will construct a slice of type ACK
func NewACKSlice(input hl7.HL7Type) []ACK {
	values := make([]ACK, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewACK(input.Index(i))
	}
	return values
}

// NewADTA37 creates an implementation of ADTA37
func NewADTA37(input hl7.HL7Type) ADTA37 {
	v := ADTA37{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.Disability = NewDB1Slice(input.Index(7))
	v.AdditionalPatientIdentification = NewPID(input.Index(8))
	v.AdditionalPatientAdditionalDemographic = NewPD1(input.Index(9))
	v.AdditionalPatientVisit = NewPV1(input.Index(10))
	v.AdditionalDisability = NewDB1Slice(input.Index(11))
	return v
}

type ADTA37 struct {
	MessageHeader                          MSH
	SoftwareSegment                        []SFT
	UserAuthenticationCredentials          UAC
	EventType                              EVN
	PatientIdentification                  PID
	PatientAdditionalDemographic           PD1
	PatientVisit                           PV1
	Disability                             []DB1
	AdditionalPatientIdentification        PID
	AdditionalPatientAdditionalDemographic PD1
	AdditionalPatientVisit                 PV1
	AdditionalDisability                   []DB1
}

// NewADTA37Slice will construct a slice of type ADTA37
func NewADTA37Slice(input hl7.HL7Type) []ADTA37 {
	values := make([]ADTA37, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA37(input.Index(i))
	}
	return values
}

// NewADTA16 creates an implementation of ADTA16
func NewADTA16(input hl7.HL7Type) ADTA16 {
	v := ADTA16{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(8))
	v.PatientVisit = NewPV1(input.Index(9))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(10))
	v.AdditionalAccessRestriction = NewARVSlice(input.Index(11))
	v.AdditionalRole = NewROLSlice(input.Index(12))
	v.Disability = NewDB1Slice(input.Index(13))
	v.ObservationResult = NewOBXSlice(input.Index(14))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(15))
	v.Diagnosis = NewDG1Slice(input.Index(16))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(17))
	v.Procedure = NewADTA16ProcedureSlice(input.Index(18))
	v.Guarantor = NewGT1Slice(input.Index(19))
	v.Insurance = NewADTA16InsuranceSlice(input.Index(20))
	v.Accident = NewACC(input.Index(21))
	return v
}

type ADTA16 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	AccessRestriction                 []ARV
	Role                              []ROL
	NextOfKinAssociatedParties        []NK1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalAccessRestriction       []ARV
	AdditionalRole                    []ROL
	Disability                        []DB1
	ObservationResult                 []OBX
	PatientAllergyInformation         []AL1
	Diagnosis                         []DG1
	DiagnosisRelatedGroup             DRG
	Procedure                         []ADTA16Procedure
	Guarantor                         []GT1
	Insurance                         []ADTA16Insurance
	Accident                          ACC
}

// NewADTA16Slice will construct a slice of type ADTA16
func NewADTA16Slice(input hl7.HL7Type) []ADTA16 {
	values := make([]ADTA16, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA16(input.Index(i))
	}
	return values
}

// NewQBPZ73 creates an implementation of QBPZ73
func NewQBPZ73(input hl7.HL7Type) QBPZ73 {
	v := QBPZ73{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.ResponseControlParameter = NewRCP(input.Index(4))
	return v
}

type QBPZ73 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
}

// NewQBPZ73Slice will construct a slice of type QBPZ73
func NewQBPZ73Slice(input hl7.HL7Type) []QBPZ73 {
	values := make([]QBPZ73, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPZ73(input.Index(i))
	}
	return values
}

// NewSRRS11 creates an implementation of SRRS11
func NewSRRS11(input hl7.HL7Type) SRRS11 {
	v := SRRS11{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.Schedule = NewSRRS11Schedule(input.Index(3))
	return v
}

type SRRS11 struct {
	MessageHeader         MSH
	MessageAcknowledgment MSA
	Error                 []ERR
	Schedule              SRRS11Schedule
}

// NewSRRS11Slice will construct a slice of type SRRS11
func NewSRRS11Slice(input hl7.HL7Type) []SRRS11 {
	values := make([]SRRS11, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRRS11(input.Index(i))
	}
	return values
}

// NewEANU09 creates an implementation of EANU09
func NewEANU09(input hl7.HL7Type) EANU09 {
	v := EANU09{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EquipmentDetail = NewEQU(input.Index(3))
	v.Notification = NewEANU09NotificationSlice(input.Index(4))
	return v
}

type EANU09 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EquipmentDetail               EQU
	Notification                  []EANU09Notification
}

// NewEANU09Slice will construct a slice of type EANU09
func NewEANU09Slice(input hl7.HL7Type) []EANU09 {
	values := make([]EANU09, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEANU09(input.Index(i))
	}
	return values
}

// NewSIUS22 creates an implementation of SIUS22
func NewSIUS22(input hl7.HL7Type) SIUS22 {
	v := SIUS22{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SchedulingActivityInformation = NewSCH(input.Index(1))
	v.TimingQuantity = NewTQ1Slice(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSIUS22PatientSlice(input.Index(4))
	v.Resources = NewSIUS22ResourcesSlice(input.Index(5))
	return v
}

type SIUS22 struct {
	MessageHeader                 MSH
	SchedulingActivityInformation SCH
	TimingQuantity                []TQ1
	NotesAndComments              []NTE
	Patient                       []SIUS22Patient
	Resources                     []SIUS22Resources
}

// NewSIUS22Slice will construct a slice of type SIUS22
func NewSIUS22Slice(input hl7.HL7Type) []SIUS22 {
	values := make([]SIUS22, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSIUS22(input.Index(i))
	}
	return values
}

// NewRPAI11 creates an implementation of RPAI11
func NewRPAI11(input hl7.HL7Type) RPAI11 {
	v := RPAI11{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.ReferralInformation = NewRF1(input.Index(4))
	v.Authorization1 = NewRPAI11Authorization1(input.Index(5))
	v.Provider = NewRPAI11ProviderSlice(input.Index(6))
	v.PatientIdentification = NewPID(input.Index(7))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(8))
	v.Guarantor = NewGT1Slice(input.Index(9))
	v.Insurance = NewRPAI11InsuranceSlice(input.Index(10))
	v.Accident = NewACC(input.Index(11))
	v.Diagnosis = NewDG1Slice(input.Index(12))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(13))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(14))
	v.Procedure = NewRPAI11ProcedureSlice(input.Index(15))
	v.Observation = NewRPAI11ObservationSlice(input.Index(16))
	v.Visit = NewRPAI11Visit(input.Index(17))
	v.NotesAndComments = NewNTESlice(input.Index(18))
	return v
}

type RPAI11 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	ReferralInformation           RF1
	Authorization1                RPAI11Authorization1
	Provider                      []RPAI11Provider
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	Guarantor                     []GT1
	Insurance                     []RPAI11Insurance
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []RPAI11Procedure
	Observation                   []RPAI11Observation
	Visit                         RPAI11Visit
	NotesAndComments              []NTE
}

// NewRPAI11Slice will construct a slice of type RPAI11
func NewRPAI11Slice(input hl7.HL7Type) []RPAI11 {
	values := make([]RPAI11, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRPAI11(input.Index(i))
	}
	return values
}

// NewSLNS35 creates an implementation of SLNS35
func NewSLNS35(input hl7.HL7Type) SLNS35 {
	v := SLNS35{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.SterilizationLot = NewSLTSlice(input.Index(3))
	return v
}

type SLNS35 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	SterilizationLot              []SLT
}

// NewSLNS35Slice will construct a slice of type SLNS35
func NewSLNS35Slice(input hl7.HL7Type) []SLNS35 {
	values := make([]SLNS35, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSLNS35(input.Index(i))
	}
	return values
}

// NewQBPZ79 creates an implementation of QBPZ79
func NewQBPZ79(input hl7.HL7Type) QBPZ79 {
	v := QBPZ79{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.AnyHl7Segment = NewHxx(input.Index(4))
	v.ResponseControlParameter = NewRCP(input.Index(5))
	v.ContinuationPointer = NewDSC(input.Index(6))
	return v
}

type QBPZ79 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	AnyHl7Segment                 Hxx
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPZ79Slice will construct a slice of type QBPZ79
func NewQBPZ79Slice(input hl7.HL7Type) []QBPZ79 {
	values := make([]QBPZ79, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPZ79(input.Index(i))
	}
	return values
}

// NewCSUC11 creates an implementation of CSUC11
func NewCSUC11(input hl7.HL7Type) CSUC11 {
	v := CSUC11{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Patient = NewCSUC11PatientSlice(input.Index(3))
	return v
}

type CSUC11 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Patient                       []CSUC11Patient
}

// NewCSUC11Slice will construct a slice of type CSUC11
func NewCSUC11Slice(input hl7.HL7Type) []CSUC11 {
	values := make([]CSUC11, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCSUC11(input.Index(i))
	}
	return values
}

// NewEHCE04 creates an implementation of EHCE04
func NewEHCE04(input hl7.HL7Type) EHCE04 {
	v := EHCE04{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUACSlice(input.Index(2))
	v.ReassessmentRequestInfo = NewEHCE04ReassessmentRequestInfo(input.Index(3))
	return v
}

type EHCE04 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials []UAC
	ReassessmentRequestInfo       EHCE04ReassessmentRequestInfo
}

// NewEHCE04Slice will construct a slice of type EHCE04
func NewEHCE04Slice(input hl7.HL7Type) []EHCE04 {
	values := make([]EHCE04, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEHCE04(input.Index(i))
	}
	return values
}

// NewRPRI03 creates an implementation of RPRI03
func NewRPRI03(input hl7.HL7Type) RPRI03 {
	v := RPRI03{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Provider = NewRPRI03ProviderSlice(input.Index(4))
	v.PatientIdentification = NewPID(input.Index(5))
	v.NotesAndComments = NewNTESlice(input.Index(6))
	return v
}

type RPRI03 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Provider                      []RPRI03Provider
	PatientIdentification         PID
	NotesAndComments              []NTE
}

// NewRPRI03Slice will construct a slice of type RPRI03
func NewRPRI03Slice(input hl7.HL7Type) []RPRI03 {
	values := make([]RPRI03, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRPRI03(input.Index(i))
	}
	return values
}

// NewRTBK13 creates an implementation of RTBK13
func NewRTBK13(input hl7.HL7Type) RTBK13 {
	v := RTBK13{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.RowDefinition = NewRTBK13RowDefinition(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RTBK13 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	RowDefinition                 RTBK13RowDefinition
	ContinuationPointer           DSC
}

// NewRTBK13Slice will construct a slice of type RTBK13
func NewRTBK13Slice(input hl7.HL7Type) []RTBK13 {
	values := make([]RTBK13, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRTBK13(input.Index(i))
	}
	return values
}

// NewADTA60 creates an implementation of ADTA60
func NewADTA60(input hl7.HL7Type) ADTA60 {
	v := ADTA60{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.AccessRestriction = NewARVSlice(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.AdditionalAccessRestriction = NewARVSlice(input.Index(8))
	v.AdverseReactionGroup = NewADTA60AdverseReactionGroupSlice(input.Index(9))
	return v
}

type ADTA60 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	AccessRestriction                 []ARV
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalAccessRestriction       []ARV
	AdverseReactionGroup              []ADTA60AdverseReactionGroup
}

// NewADTA60Slice will construct a slice of type ADTA60
func NewADTA60Slice(input hl7.HL7Type) []ADTA60 {
	values := make([]ADTA60, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA60(input.Index(i))
	}
	return values
}

// NewBRTO32 creates an implementation of BRTO32
func NewBRTO32(input hl7.HL7Type) BRTO32 {
	v := BRTO32{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewBRTO32Response(input.Index(6))
	return v
}

type BRTO32 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      BRTO32Response
}

// NewBRTO32Slice will construct a slice of type BRTO32
func NewBRTO32Slice(input hl7.HL7Type) []BRTO32 {
	values := make([]BRTO32, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBRTO32(input.Index(i))
	}
	return values
}

// NewMFNM05 creates an implementation of MFNM05
func NewMFNM05(input hl7.HL7Type) MFNM05 {
	v := MFNM05{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MasterFileIdentification = NewMFI(input.Index(3))
	v.MfLocation = NewMFNM05MfLocationSlice(input.Index(4))
	return v
}

type MFNM05 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MasterFileIdentification      MFI
	MfLocation                    []MFNM05MfLocation
}

// NewMFNM05Slice will construct a slice of type MFNM05
func NewMFNM05Slice(input hl7.HL7Type) []MFNM05 {
	values := make([]MFNM05, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFNM05(input.Index(i))
	}
	return values
}

// NewRSPK34 creates an implementation of RSPK34
func NewRSPK34(input hl7.HL7Type) RSPK34 {
	v := RSPK34{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.StaffIdentification = NewSTFSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.Donnor = NewRSPK34Donnor(input.Index(7))
	v.Donnation = NewRSPK34Donnation(input.Index(8))
	return v
}

type RSPK34 struct {
	MessageHeader                 MSH
	StaffIdentification           []STF
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	Donnor                        RSPK34Donnor
	Donnation                     RSPK34Donnation
}

// NewRSPK34Slice will construct a slice of type RSPK34
func NewRSPK34Slice(input hl7.HL7Type) []RSPK34 {
	values := make([]RSPK34, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPK34(input.Index(i))
	}
	return values
}

// NewPMUB03 creates an implementation of PMUB03
func NewPMUB03(input hl7.HL7Type) PMUB03 {
	v := PMUB03{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.StaffIdentification = NewSTF(input.Index(4))
	return v
}

type PMUB03 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	StaffIdentification           STF
}

// NewPMUB03Slice will construct a slice of type PMUB03
func NewPMUB03Slice(input hl7.HL7Type) []PMUB03 {
	values := make([]PMUB03, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPMUB03(input.Index(i))
	}
	return values
}

// NewMFKM14 creates an implementation of MFKM14
func NewMFKM14(input hl7.HL7Type) MFKM14 {
	v := MFKM14{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.MasterFileIdentification = NewMFI(input.Index(5))
	v.MasterFileAcknowledgment = NewMFASlice(input.Index(6))
	return v
}

type MFKM14 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	MasterFileIdentification      MFI
	MasterFileAcknowledgment      []MFA
}

// NewMFKM14Slice will construct a slice of type MFKM14
func NewMFKM14Slice(input hl7.HL7Type) []MFKM14 {
	values := make([]MFKM14, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFKM14(input.Index(i))
	}
	return values
}

// NewLSUU12 creates an implementation of LSUU12
func NewLSUU12(input hl7.HL7Type) LSUU12 {
	v := LSUU12{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EquipmentDetail = NewEQU(input.Index(3))
	v.EquipmentLogService = NewEQPSlice(input.Index(4))
	return v
}

type LSUU12 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EquipmentDetail               EQU
	EquipmentLogService           []EQP
}

// NewLSUU12Slice will construct a slice of type LSUU12
func NewLSUU12Slice(input hl7.HL7Type) []LSUU12 {
	values := make([]LSUU12, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewLSUU12(input.Index(i))
	}
	return values
}

// NewRSPZ82 creates an implementation of RSPZ82
func NewRSPZ82(input hl7.HL7Type) RSPZ82 {
	v := RSPZ82{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.ResponseControlParameter = NewRCP(input.Index(7))
	v.QueryResponse = NewRSPZ82QueryResponseSlice(input.Index(8))
	v.ContinuationPointer = NewDSC(input.Index(9))
	return v
}

type RSPZ82 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
	QueryResponse                 []RSPZ82QueryResponse
	ContinuationPointer           DSC
}

// NewRSPZ82Slice will construct a slice of type RSPZ82
func NewRSPZ82Slice(input hl7.HL7Type) []RSPZ82 {
	values := make([]RSPZ82, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPZ82(input.Index(i))
	}
	return values
}

// NewADTA22 creates an implementation of ADTA22
func NewADTA22(input hl7.HL7Type) ADTA22 {
	v := ADTA22{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.Disability = NewDB1Slice(input.Index(8))
	v.ObservationResult = NewOBXSlice(input.Index(9))
	return v
}

type ADTA22 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	Disability                        []DB1
	ObservationResult                 []OBX
}

// NewADTA22Slice will construct a slice of type ADTA22
func NewADTA22Slice(input hl7.HL7Type) []ADTA22 {
	values := make([]ADTA22, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA22(input.Index(i))
	}
	return values
}

// NewDBCO41 creates an implementation of DBCO41
func NewDBCO41(input hl7.HL7Type) DBCO41 {
	v := DBCO41{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.StaffIdentification = NewSTFSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Donnor = NewDBCO41Donnor(input.Index(3))
	return v
}

type DBCO41 struct {
	MessageHeader                 MSH
	StaffIdentification           []STF
	UserAuthenticationCredentials UAC
	Donnor                        DBCO41Donnor
}

// NewDBCO41Slice will construct a slice of type DBCO41
func NewDBCO41Slice(input hl7.HL7Type) []DBCO41 {
	values := make([]DBCO41, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDBCO41(input.Index(i))
	}
	return values
}

// NewSLRS28 creates an implementation of SLRS28
func NewSLRS28(input hl7.HL7Type) SLRS28 {
	v := SLRS28{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.SterilizationLot = NewSLTSlice(input.Index(3))
	return v
}

type SLRS28 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	SterilizationLot              []SLT
}

// NewSLRS28Slice will construct a slice of type SLRS28
func NewSLRS28Slice(input hl7.HL7Type) []SLRS28 {
	values := make([]SLRS28, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSLRS28(input.Index(i))
	}
	return values
}

// NewADTA07 creates an implementation of ADTA07
func NewADTA07(input hl7.HL7Type) ADTA07 {
	v := ADTA07{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.MergePatientInformation = NewMRG(input.Index(8))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(9))
	v.PatientVisit = NewPV1(input.Index(10))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(11))
	v.AdditionalAccessRestriction = NewARVSlice(input.Index(12))
	v.AdditionalRole = NewROLSlice(input.Index(13))
	v.Disability = NewDB1Slice(input.Index(14))
	v.ObservationResult = NewOBXSlice(input.Index(15))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(16))
	v.Diagnosis = NewDG1Slice(input.Index(17))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(18))
	v.Procedure = NewADTA07ProcedureSlice(input.Index(19))
	v.Guarantor = NewGT1Slice(input.Index(20))
	v.Insurance = NewADTA07InsuranceSlice(input.Index(21))
	v.Accident = NewACC(input.Index(22))
	v.Ub82 = NewUB1(input.Index(23))
	v.UniformBillingData = NewUB2(input.Index(24))
	return v
}

type ADTA07 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	AccessRestriction                 []ARV
	Role                              []ROL
	MergePatientInformation           MRG
	NextOfKinAssociatedParties        []NK1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalAccessRestriction       []ARV
	AdditionalRole                    []ROL
	Disability                        []DB1
	ObservationResult                 []OBX
	PatientAllergyInformation         []AL1
	Diagnosis                         []DG1
	DiagnosisRelatedGroup             DRG
	Procedure                         []ADTA07Procedure
	Guarantor                         []GT1
	Insurance                         []ADTA07Insurance
	Accident                          ACC
	Ub82                              UB1
	UniformBillingData                UB2
}

// NewADTA07Slice will construct a slice of type ADTA07
func NewADTA07Slice(input hl7.HL7Type) []ADTA07 {
	values := make([]ADTA07, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA07(input.Index(i))
	}
	return values
}

// NewQVRQ17 creates an implementation of QVRQ17
func NewQVRQ17(input hl7.HL7Type) QVRQ17 {
	v := QVRQ17{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.Qbp = NewQVRQ17Qbp(input.Index(4))
	v.ResponseControlParameter = NewRCP(input.Index(5))
	v.ContinuationPointer = NewDSC(input.Index(6))
	return v
}

type QVRQ17 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	Qbp                           QVRQ17Qbp
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQVRQ17Slice will construct a slice of type QVRQ17
func NewQVRQ17Slice(input hl7.HL7Type) []QVRQ17 {
	values := make([]QVRQ17, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQVRQ17(input.Index(i))
	}
	return values
}

// NewPGLPC7 creates an implementation of PGLPC7
func NewPGLPC7(input hl7.HL7Type) PGLPC7 {
	v := PGLPC7{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientVisit = NewPGLPC7PatientVisit(input.Index(4))
	v.Goal = NewPGLPC7GoalSlice(input.Index(5))
	return v
}

type PGLPC7 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientIdentification         PID
	PatientVisit                  PGLPC7PatientVisit
	Goal                          []PGLPC7Goal
}

// NewPGLPC7Slice will construct a slice of type PGLPC7
func NewPGLPC7Slice(input hl7.HL7Type) []PGLPC7 {
	values := make([]PGLPC7, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPGLPC7(input.Index(i))
	}
	return values
}

// NewORDO04 creates an implementation of ORDO04
func NewORDO04(input hl7.HL7Type) ORDO04 {
	v := ORDO04{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewORDO04Response(input.Index(6))
	return v
}

type ORDO04 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      ORDO04Response
}

// NewORDO04Slice will construct a slice of type ORDO04
func NewORDO04Slice(input hl7.HL7Type) []ORDO04 {
	values := make([]ORDO04, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORDO04(input.Index(i))
	}
	return values
}

// NewQBPZ91 creates an implementation of QBPZ91
func NewQBPZ91(input hl7.HL7Type) QBPZ91 {
	v := QBPZ91{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.Qbp = NewQBPZ91Qbp(input.Index(4))
	v.TableRowDefinition = NewRDF(input.Index(5))
	v.ResponseControlParameter = NewRCP(input.Index(6))
	v.ContinuationPointer = NewDSC(input.Index(7))
	return v
}

type QBPZ91 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	Qbp                           QBPZ91Qbp
	TableRowDefinition            RDF
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPZ91Slice will construct a slice of type QBPZ91
func NewQBPZ91Slice(input hl7.HL7Type) []QBPZ91 {
	values := make([]QBPZ91, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPZ91(input.Index(i))
	}
	return values
}

// NewSRRS09 creates an implementation of SRRS09
func NewSRRS09(input hl7.HL7Type) SRRS09 {
	v := SRRS09{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.Schedule = NewSRRS09Schedule(input.Index(3))
	return v
}

type SRRS09 struct {
	MessageHeader         MSH
	MessageAcknowledgment MSA
	Error                 []ERR
	Schedule              SRRS09Schedule
}

// NewSRRS09Slice will construct a slice of type SRRS09
func NewSRRS09Slice(input hl7.HL7Type) []SRRS09 {
	values := make([]SRRS09, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRRS09(input.Index(i))
	}
	return values
}

// NewQBPZ97 creates an implementation of QBPZ97
func NewQBPZ97(input hl7.HL7Type) QBPZ97 {
	v := QBPZ97{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.AnyHl7Segment = NewHxx(input.Index(4))
	v.ResponseControlParameter = NewRCP(input.Index(5))
	v.ContinuationPointer = NewDSC(input.Index(6))
	return v
}

type QBPZ97 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	AnyHl7Segment                 Hxx
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPZ97Slice will construct a slice of type QBPZ97
func NewQBPZ97Slice(input hl7.HL7Type) []QBPZ97 {
	values := make([]QBPZ97, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPZ97(input.Index(i))
	}
	return values
}

// NewADTA04 creates an implementation of ADTA04
func NewADTA04(input hl7.HL7Type) ADTA04 {
	v := ADTA04{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(8))
	v.PatientVisit = NewPV1(input.Index(9))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(10))
	v.AdditionalAccessRestriction = NewARVSlice(input.Index(11))
	v.AdditionalRole = NewROLSlice(input.Index(12))
	v.Disability = NewDB1Slice(input.Index(13))
	v.ObservationResult = NewOBXSlice(input.Index(14))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(15))
	v.Diagnosis = NewDG1Slice(input.Index(16))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(17))
	v.Procedure = NewADTA04ProcedureSlice(input.Index(18))
	v.Guarantor = NewGT1Slice(input.Index(19))
	v.Insurance = NewADTA04InsuranceSlice(input.Index(20))
	v.Accident = NewACC(input.Index(21))
	v.Ub82 = NewUB1(input.Index(22))
	v.UniformBillingData = NewUB2(input.Index(23))
	v.PatientDeathAndAutopsy = NewPDA(input.Index(24))
	return v
}

type ADTA04 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	AccessRestriction                 []ARV
	Role                              []ROL
	NextOfKinAssociatedParties        []NK1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalAccessRestriction       []ARV
	AdditionalRole                    []ROL
	Disability                        []DB1
	ObservationResult                 []OBX
	PatientAllergyInformation         []AL1
	Diagnosis                         []DG1
	DiagnosisRelatedGroup             DRG
	Procedure                         []ADTA04Procedure
	Guarantor                         []GT1
	Insurance                         []ADTA04Insurance
	Accident                          ACC
	Ub82                              UB1
	UniformBillingData                UB2
	PatientDeathAndAutopsy            PDA
}

// NewADTA04Slice will construct a slice of type ADTA04
func NewADTA04Slice(input hl7.HL7Type) []ADTA04 {
	values := make([]ADTA04, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA04(input.Index(i))
	}
	return values
}

// NewADTA52 creates an implementation of ADTA52
func NewADTA52(input hl7.HL7Type) ADTA52 {
	v := ADTA52{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	return v
}

type ADTA52 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
}

// NewADTA52Slice will construct a slice of type ADTA52
func NewADTA52Slice(input hl7.HL7Type) []ADTA52 {
	values := make([]ADTA52, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA52(input.Index(i))
	}
	return values
}

// NewQBPZ85 creates an implementation of QBPZ85
func NewQBPZ85(input hl7.HL7Type) QBPZ85 {
	v := QBPZ85{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.Qbp = NewQBPZ85Qbp(input.Index(4))
	v.ResponseControlParameter = NewRCP(input.Index(5))
	v.ContinuationPointer = NewDSC(input.Index(6))
	return v
}

type QBPZ85 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	Qbp                           QBPZ85Qbp
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPZ85Slice will construct a slice of type QBPZ85
func NewQBPZ85Slice(input hl7.HL7Type) []QBPZ85 {
	values := make([]QBPZ85, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPZ85(input.Index(i))
	}
	return values
}

// NewOMSO05 creates an implementation of OMSO05
func NewOMSO05(input hl7.HL7Type) OMSO05 {
	v := OMSO05{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewOMSO05Patient(input.Index(4))
	v.Order = NewOMSO05OrderSlice(input.Index(5))
	return v
}

type OMSO05 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       OMSO05Patient
	Order                         []OMSO05Order
}

// NewOMSO05Slice will construct a slice of type OMSO05
func NewOMSO05Slice(input hl7.HL7Type) []OMSO05 {
	values := make([]OMSO05, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOMSO05(input.Index(i))
	}
	return values
}

// NewESUU01 creates an implementation of ESUU01
func NewESUU01(input hl7.HL7Type) ESUU01 {
	v := ESUU01{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EquipmentDetail = NewEQU(input.Index(3))
	v.InteractionStatusDetail = NewISDSlice(input.Index(4))
	return v
}

type ESUU01 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EquipmentDetail               EQU
	InteractionStatusDetail       []ISD
}

// NewESUU01Slice will construct a slice of type ESUU01
func NewESUU01Slice(input hl7.HL7Type) []ESUU01 {
	values := make([]ESUU01, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewESUU01(input.Index(i))
	}
	return values
}

// NewRTBZ94 creates an implementation of RTBZ94
func NewRTBZ94(input hl7.HL7Type) RTBZ94 {
	v := RTBZ94{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.RowDefinition = NewRTBZ94RowDefinition(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RTBZ94 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	RowDefinition                 RTBZ94RowDefinition
	ContinuationPointer           DSC
}

// NewRTBZ94Slice will construct a slice of type RTBZ94
func NewRTBZ94Slice(input hl7.HL7Type) []RTBZ94 {
	values := make([]RTBZ94, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRTBZ94(input.Index(i))
	}
	return values
}

// NewRSPK22 creates an implementation of RSPK22
func NewRSPK22(input hl7.HL7Type) RSPK22 {
	v := RSPK22{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.QueryResponse = NewRSPK22QueryResponseSlice(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RSPK22 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	QueryResponse                 []RSPK22QueryResponse
	ContinuationPointer           DSC
}

// NewRSPK22Slice will construct a slice of type RSPK22
func NewRSPK22Slice(input hl7.HL7Type) []RSPK22 {
	values := make([]RSPK22, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPK22(input.Index(i))
	}
	return values
}

// NewRQAI09 creates an implementation of RQAI09
func NewRQAI09(input hl7.HL7Type) RQAI09 {
	v := RQAI09{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.ReferralInformation = NewRF1(input.Index(3))
	v.Authorization = NewRQAI09Authorization(input.Index(4))
	v.Provider = NewRQAI09ProviderSlice(input.Index(5))
	v.PatientIdentification = NewPID(input.Index(6))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(7))
	v.GuarantorInsurance = NewRQAI09GuarantorInsurance(input.Index(8))
	v.Accident = NewACC(input.Index(9))
	v.Diagnosis = NewDG1Slice(input.Index(10))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(11))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(12))
	v.Procedure = NewRQAI09ProcedureSlice(input.Index(13))
	v.Observation = NewRQAI09ObservationSlice(input.Index(14))
	v.Visit = NewRQAI09Visit(input.Index(15))
	v.NotesAndComments = NewNTESlice(input.Index(16))
	return v
}

type RQAI09 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	ReferralInformation           RF1
	Authorization                 RQAI09Authorization
	Provider                      []RQAI09Provider
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	GuarantorInsurance            RQAI09GuarantorInsurance
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []RQAI09Procedure
	Observation                   []RQAI09Observation
	Visit                         RQAI09Visit
	NotesAndComments              []NTE
}

// NewRQAI09Slice will construct a slice of type RQAI09
func NewRQAI09Slice(input hl7.HL7Type) []RQAI09 {
	values := make([]RQAI09, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRQAI09(input.Index(i))
	}
	return values
}

// NewMFNM04 creates an implementation of MFNM04
func NewMFNM04(input hl7.HL7Type) MFNM04 {
	v := MFNM04{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MasterFileIdentification = NewMFI(input.Index(3))
	v.MfCdm = NewMFNM04MfCdmSlice(input.Index(4))
	return v
}

type MFNM04 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MasterFileIdentification      MFI
	MfCdm                         []MFNM04MfCdm
}

// NewMFNM04Slice will construct a slice of type MFNM04
func NewMFNM04Slice(input hl7.HL7Type) []MFNM04 {
	values := make([]MFNM04, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFNM04(input.Index(i))
	}
	return values
}

// NewQSXJ02 creates an implementation of QSXJ02
func NewQSXJ02(input hl7.HL7Type) QSXJ02 {
	v := QSXJ02{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryIdentification = NewQID(input.Index(3))
	return v
}

type QSXJ02 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryIdentification           QID
}

// NewQSXJ02Slice will construct a slice of type QSXJ02
func NewQSXJ02Slice(input hl7.HL7Type) []QSXJ02 {
	values := make([]QSXJ02, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQSXJ02(input.Index(i))
	}
	return values
}

// NewCSUC12 creates an implementation of CSUC12
func NewCSUC12(input hl7.HL7Type) CSUC12 {
	v := CSUC12{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Patient = NewCSUC12PatientSlice(input.Index(3))
	return v
}

type CSUC12 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Patient                       []CSUC12Patient
}

// NewCSUC12Slice will construct a slice of type CSUC12
func NewCSUC12Slice(input hl7.HL7Type) []CSUC12 {
	values := make([]CSUC12, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCSUC12(input.Index(i))
	}
	return values
}

// NewOMQO42 creates an implementation of OMQO42
func NewOMQO42(input hl7.HL7Type) OMQO42 {
	v := OMQO42{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewOMQO42Patient(input.Index(4))
	v.Order = NewOMQO42OrderSlice(input.Index(5))
	return v
}

type OMQO42 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       OMQO42Patient
	Order                         []OMQO42Order
}

// NewOMQO42Slice will construct a slice of type OMQO42
func NewOMQO42Slice(input hl7.HL7Type) []OMQO42 {
	values := make([]OMQO42, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOMQO42(input.Index(i))
	}
	return values
}

// NewADTA54 creates an implementation of ADTA54
func NewADTA54(input hl7.HL7Type) ADTA54 {
	v := ADTA54{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.Role = NewROLSlice(input.Index(6))
	v.PatientVisit = NewPV1(input.Index(7))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(8))
	v.AdditonalRole = NewROLSlice(input.Index(9))
	return v
}

type ADTA54 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	Role                              []ROL
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditonalRole                     []ROL
}

// NewADTA54Slice will construct a slice of type ADTA54
func NewADTA54Slice(input hl7.HL7Type) []ADTA54 {
	values := make([]ADTA54, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA54(input.Index(i))
	}
	return values
}

// NewRPAI09 creates an implementation of RPAI09
func NewRPAI09(input hl7.HL7Type) RPAI09 {
	v := RPAI09{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.ReferralInformation = NewRF1(input.Index(4))
	v.Authorization1 = NewRPAI09Authorization1(input.Index(5))
	v.Provider = NewRPAI09ProviderSlice(input.Index(6))
	v.PatientIdentification = NewPID(input.Index(7))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(8))
	v.Guarantor = NewGT1Slice(input.Index(9))
	v.Insurance = NewRPAI09InsuranceSlice(input.Index(10))
	v.Accident = NewACC(input.Index(11))
	v.Diagnosis = NewDG1Slice(input.Index(12))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(13))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(14))
	v.Procedure = NewRPAI09ProcedureSlice(input.Index(15))
	v.Observation = NewRPAI09ObservationSlice(input.Index(16))
	v.Visit = NewRPAI09Visit(input.Index(17))
	v.NotesAndComments = NewNTESlice(input.Index(18))
	return v
}

type RPAI09 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	ReferralInformation           RF1
	Authorization1                RPAI09Authorization1
	Provider                      []RPAI09Provider
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	Guarantor                     []GT1
	Insurance                     []RPAI09Insurance
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []RPAI09Procedure
	Observation                   []RPAI09Observation
	Visit                         RPAI09Visit
	NotesAndComments              []NTE
}

// NewRPAI09Slice will construct a slice of type RPAI09
func NewRPAI09Slice(input hl7.HL7Type) []RPAI09 {
	values := make([]RPAI09, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRPAI09(input.Index(i))
	}
	return values
}

// NewADTA12 creates an implementation of ADTA12
func NewADTA12(input hl7.HL7Type) ADTA12 {
	v := ADTA12{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.Disability = NewDB1Slice(input.Index(8))
	v.ObservationResult = NewOBXSlice(input.Index(9))
	return v
}

type ADTA12 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	Disability                        []DB1
	ObservationResult                 []OBX
}

// NewADTA12Slice will construct a slice of type ADTA12
func NewADTA12Slice(input hl7.HL7Type) []ADTA12 {
	values := make([]ADTA12, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA12(input.Index(i))
	}
	return values
}

// NewVXUV04 creates an implementation of VXUV04
func NewVXUV04(input hl7.HL7Type) VXUV04 {
	v := VXUV04{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientAdditionalDemographic = NewPD1(input.Index(4))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.PatientVisit = NewVXUV04PatientVisit(input.Index(7))
	v.Guarantor = NewGT1Slice(input.Index(8))
	v.Insurance = NewVXUV04InsuranceSlice(input.Index(9))
	v.PersonObservation = NewVXUV04PersonObservationSlice(input.Index(10))
	v.Order = NewVXUV04OrderSlice(input.Index(11))
	return v
}

type VXUV04 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientIdentification         PID
	PatientAdditionalDemographic  PD1
	NextOfKinAssociatedParties    []NK1
	AccessRestriction             []ARV
	PatientVisit                  VXUV04PatientVisit
	Guarantor                     []GT1
	Insurance                     []VXUV04Insurance
	PersonObservation             []VXUV04PersonObservation
	Order                         []VXUV04Order
}

// NewVXUV04Slice will construct a slice of type VXUV04
func NewVXUV04Slice(input hl7.HL7Type) []VXUV04 {
	values := make([]VXUV04, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewVXUV04(input.Index(i))
	}
	return values
}

// NewADTA14 creates an implementation of ADTA14
func NewADTA14(input hl7.HL7Type) ADTA14 {
	v := ADTA14{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(8))
	v.PatientVisit = NewPV1(input.Index(9))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(10))
	v.AdditionalAccessRestriction = NewARVSlice(input.Index(11))
	v.AdditionalRole = NewROLSlice(input.Index(12))
	v.Disability = NewDB1Slice(input.Index(13))
	v.ObservationResult = NewOBXSlice(input.Index(14))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(15))
	v.Diagnosis = NewDG1Slice(input.Index(16))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(17))
	v.Procedure = NewADTA14ProcedureSlice(input.Index(18))
	v.Guarantor = NewGT1Slice(input.Index(19))
	v.Insurance = NewADTA14InsuranceSlice(input.Index(20))
	v.Accident = NewACC(input.Index(21))
	v.Ub82 = NewUB1(input.Index(22))
	v.UniformBillingData = NewUB2(input.Index(23))
	return v
}

type ADTA14 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	AccessRestriction                 []ARV
	Role                              []ROL
	NextOfKinAssociatedParties        []NK1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalAccessRestriction       []ARV
	AdditionalRole                    []ROL
	Disability                        []DB1
	ObservationResult                 []OBX
	PatientAllergyInformation         []AL1
	Diagnosis                         []DG1
	DiagnosisRelatedGroup             DRG
	Procedure                         []ADTA14Procedure
	Guarantor                         []GT1
	Insurance                         []ADTA14Insurance
	Accident                          ACC
	Ub82                              UB1
	UniformBillingData                UB2
}

// NewADTA14Slice will construct a slice of type ADTA14
func NewADTA14Slice(input hl7.HL7Type) []ADTA14 {
	values := make([]ADTA14, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA14(input.Index(i))
	}
	return values
}

// NewMFNM06 creates an implementation of MFNM06
func NewMFNM06(input hl7.HL7Type) MFNM06 {
	v := MFNM06{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MasterFileIdentification = NewMFI(input.Index(3))
	v.MfClinStudy = NewMFNM06MfClinStudySlice(input.Index(4))
	return v
}

type MFNM06 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MasterFileIdentification      MFI
	MfClinStudy                   []MFNM06MfClinStudy
}

// NewMFNM06Slice will construct a slice of type MFNM06
func NewMFNM06Slice(input hl7.HL7Type) []MFNM06 {
	values := make([]MFNM06, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFNM06(input.Index(i))
	}
	return values
}

// NewMFNM02 creates an implementation of MFNM02
func NewMFNM02(input hl7.HL7Type) MFNM02 {
	v := MFNM02{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MasterFileIdentification = NewMFI(input.Index(3))
	v.MfStaff = NewMFNM02MfStaffSlice(input.Index(4))
	return v
}

type MFNM02 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MasterFileIdentification      MFI
	MfStaff                       []MFNM02MfStaff
}

// NewMFNM02Slice will construct a slice of type MFNM02
func NewMFNM02Slice(input hl7.HL7Type) []MFNM02 {
	values := make([]MFNM02, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFNM02(input.Index(i))
	}
	return values
}

// NewQBPZ77 creates an implementation of QBPZ77
func NewQBPZ77(input hl7.HL7Type) QBPZ77 {
	v := QBPZ77{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.Qbp = NewQBPZ77Qbp(input.Index(4))
	v.TableRowDefinition = NewRDF(input.Index(5))
	v.ResponseControlParameter = NewRCP(input.Index(6))
	v.ContinuationPointer = NewDSC(input.Index(7))
	return v
}

type QBPZ77 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	Qbp                           QBPZ77Qbp
	TableRowDefinition            RDF
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPZ77Slice will construct a slice of type QBPZ77
func NewQBPZ77Slice(input hl7.HL7Type) []QBPZ77 {
	values := make([]QBPZ77, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPZ77(input.Index(i))
	}
	return values
}

// NewSIUS12 creates an implementation of SIUS12
func NewSIUS12(input hl7.HL7Type) SIUS12 {
	v := SIUS12{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SchedulingActivityInformation = NewSCH(input.Index(1))
	v.TimingQuantity = NewTQ1Slice(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSIUS12PatientSlice(input.Index(4))
	v.Resources = NewSIUS12ResourcesSlice(input.Index(5))
	return v
}

type SIUS12 struct {
	MessageHeader                 MSH
	SchedulingActivityInformation SCH
	TimingQuantity                []TQ1
	NotesAndComments              []NTE
	Patient                       []SIUS12Patient
	Resources                     []SIUS12Resources
}

// NewSIUS12Slice will construct a slice of type SIUS12
func NewSIUS12Slice(input hl7.HL7Type) []SIUS12 {
	values := make([]SIUS12, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSIUS12(input.Index(i))
	}
	return values
}

// NewADTA15 creates an implementation of ADTA15
func NewADTA15(input hl7.HL7Type) ADTA15 {
	v := ADTA15{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.PatientVisit = NewPV1(input.Index(8))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(9))
	v.AdditionalAccessRestriction = NewARVSlice(input.Index(10))
	v.AdditionalRole = NewROLSlice(input.Index(11))
	v.Disability = NewDB1Slice(input.Index(12))
	v.ObservationResult = NewOBXSlice(input.Index(13))
	return v
}

type ADTA15 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	AccessRestriction                 []ARV
	Role                              []ROL
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalAccessRestriction       []ARV
	AdditionalRole                    []ROL
	Disability                        []DB1
	ObservationResult                 []OBX
}

// NewADTA15Slice will construct a slice of type ADTA15
func NewADTA15Slice(input hl7.HL7Type) []ADTA15 {
	values := make([]ADTA15, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA15(input.Index(i))
	}
	return values
}

// NewQBPQ11 creates an implementation of QBPQ11
func NewQBPQ11(input hl7.HL7Type) QBPQ11 {
	v := QBPQ11{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.Qbp = NewQBPQ11Qbp(input.Index(4))
	v.ResponseControlParameter = NewRCP(input.Index(5))
	v.ContinuationPointer = NewDSC(input.Index(6))
	return v
}

type QBPQ11 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	Qbp                           QBPQ11Qbp
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPQ11Slice will construct a slice of type QBPQ11
func NewQBPQ11Slice(input hl7.HL7Type) []QBPQ11 {
	values := make([]QBPQ11, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPQ11(input.Index(i))
	}
	return values
}

// NewCRMC02 creates an implementation of CRMC02
func NewCRMC02(input hl7.HL7Type) CRMC02 {
	v := CRMC02{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Patient = NewCRMC02PatientSlice(input.Index(3))
	return v
}

type CRMC02 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Patient                       []CRMC02Patient
}

// NewCRMC02Slice will construct a slice of type CRMC02
func NewCRMC02Slice(input hl7.HL7Type) []CRMC02 {
	values := make([]CRMC02, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCRMC02(input.Index(i))
	}
	return values
}

// NewREFI14 creates an implementation of REFI14
func NewREFI14(input hl7.HL7Type) REFI14 {
	v := REFI14{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.ReferralInformation = NewRF1(input.Index(3))
	v.AuthorizationContact = NewREFI14AuthorizationContact(input.Index(4))
	v.ProviderContact = NewREFI14ProviderContactSlice(input.Index(5))
	v.PatientIdentification = NewPID(input.Index(6))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(7))
	v.Guarantor = NewGT1Slice(input.Index(8))
	v.Insurance = NewREFI14InsuranceSlice(input.Index(9))
	v.Accident = NewACC(input.Index(10))
	v.Diagnosis = NewDG1Slice(input.Index(11))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(12))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(13))
	v.Procedure = NewREFI14ProcedureSlice(input.Index(14))
	v.Observation = NewREFI14ObservationSlice(input.Index(15))
	v.PatientVisit = NewREFI14PatientVisit(input.Index(16))
	v.NotesAndComments = NewNTESlice(input.Index(17))
	return v
}

type REFI14 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	ReferralInformation           RF1
	AuthorizationContact          REFI14AuthorizationContact
	ProviderContact               []REFI14ProviderContact
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	Guarantor                     []GT1
	Insurance                     []REFI14Insurance
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []REFI14Procedure
	Observation                   []REFI14Observation
	PatientVisit                  REFI14PatientVisit
	NotesAndComments              []NTE
}

// NewREFI14Slice will construct a slice of type REFI14
func NewREFI14Slice(input hl7.HL7Type) []REFI14 {
	values := make([]REFI14, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewREFI14(input.Index(i))
	}
	return values
}

// NewOMGO19 creates an implementation of OMGO19
func NewOMGO19(input hl7.HL7Type) OMGO19 {
	v := OMGO19{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewOMGO19Patient(input.Index(4))
	v.Order = NewOMGO19OrderSlice(input.Index(5))
	return v
}

type OMGO19 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       OMGO19Patient
	Order                         []OMGO19Order
}

// NewOMGO19Slice will construct a slice of type OMGO19
func NewOMGO19Slice(input hl7.HL7Type) []OMGO19 {
	values := make([]OMGO19, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOMGO19(input.Index(i))
	}
	return values
}

// NewSIUS19 creates an implementation of SIUS19
func NewSIUS19(input hl7.HL7Type) SIUS19 {
	v := SIUS19{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SchedulingActivityInformation = NewSCH(input.Index(1))
	v.TimingQuantity = NewTQ1Slice(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSIUS19PatientSlice(input.Index(4))
	v.Resources = NewSIUS19ResourcesSlice(input.Index(5))
	return v
}

type SIUS19 struct {
	MessageHeader                 MSH
	SchedulingActivityInformation SCH
	TimingQuantity                []TQ1
	NotesAndComments              []NTE
	Patient                       []SIUS19Patient
	Resources                     []SIUS19Resources
}

// NewSIUS19Slice will construct a slice of type SIUS19
func NewSIUS19Slice(input hl7.HL7Type) []SIUS19 {
	values := make([]SIUS19, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSIUS19(input.Index(i))
	}
	return values
}

// NewRDRRDR creates an implementation of RDRRDR
func NewRDRRDR(input hl7.HL7Type) RDRRDR {
	v := RDRRDR{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFT(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.Definition = NewRDRRDRDefinitionSlice(input.Index(5))
	v.ContinuationPointer = NewDSC(input.Index(6))
	return v
}

type RDRRDR struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               SFT
	UserAuthenticationCredentials UAC
	Definition                    []RDRRDRDefinition
	ContinuationPointer           DSC
}

// NewRDRRDRSlice will construct a slice of type RDRRDR
func NewRDRRDRSlice(input hl7.HL7Type) []RDRRDR {
	values := make([]RDRRDR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRDRRDR(input.Index(i))
	}
	return values
}

// NewPMUB01 creates an implementation of PMUB01
func NewPMUB01(input hl7.HL7Type) PMUB01 {
	v := PMUB01{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.StaffIdentification = NewSTF(input.Index(4))
	v.PractitionerDetail = NewPRASlice(input.Index(5))
	v.PractitionerOrganizationUnit = NewORGSlice(input.Index(6))
	v.ProfessionalAffiliation = NewAFFSlice(input.Index(7))
	v.LanguageDetail = NewLANSlice(input.Index(8))
	v.EducationalDetail = NewEDUSlice(input.Index(9))
	v.CertificateDetail = NewCERSlice(input.Index(10))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(11))
	v.ParticipationInformation = NewPRTSlice(input.Index(12))
	v.Role = NewROLSlice(input.Index(13))
	return v
}

type PMUB01 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	StaffIdentification           STF
	PractitionerDetail            []PRA
	PractitionerOrganizationUnit  []ORG
	ProfessionalAffiliation       []AFF
	LanguageDetail                []LAN
	EducationalDetail             []EDU
	CertificateDetail             []CER
	NextOfKinAssociatedParties    []NK1
	ParticipationInformation      []PRT
	Role                          []ROL
}

// NewPMUB01Slice will construct a slice of type PMUB01
func NewPMUB01Slice(input hl7.HL7Type) []PMUB01 {
	values := make([]PMUB01, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPMUB01(input.Index(i))
	}
	return values
}

// NewRPLI02 creates an implementation of RPLI02
func NewRPLI02(input hl7.HL7Type) RPLI02 {
	v := RPLI02{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Provider = NewRPLI02ProviderSlice(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.DisplayData = NewDSPSlice(input.Index(6))
	v.ContinuationPointer = NewDSC(input.Index(7))
	return v
}

type RPLI02 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Provider                      []RPLI02Provider
	NotesAndComments              []NTE
	DisplayData                   []DSP
	ContinuationPointer           DSC
}

// NewRPLI02Slice will construct a slice of type RPLI02
func NewRPLI02Slice(input hl7.HL7Type) []RPLI02 {
	values := make([]RPLI02, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRPLI02(input.Index(i))
	}
	return values
}

// NewEHCE24 creates an implementation of EHCE24
func NewEHCE24(input hl7.HL7Type) EHCE24 {
	v := EHCE24{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUACSlice(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.AuthorizationResponseInfo = NewEHCE24AuthorizationResponseInfo(input.Index(5))
	return v
}

type EHCE24 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials []UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	AuthorizationResponseInfo     EHCE24AuthorizationResponseInfo
}

// NewEHCE24Slice will construct a slice of type EHCE24
func NewEHCE24Slice(input hl7.HL7Type) []EHCE24 {
	values := make([]EHCE24, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEHCE24(input.Index(i))
	}
	return values
}

// NewADTA32 creates an implementation of ADTA32
func NewADTA32(input hl7.HL7Type) ADTA32 {
	v := ADTA32{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.Disability = NewDB1Slice(input.Index(8))
	v.ObservationResult = NewOBXSlice(input.Index(9))
	return v
}

type ADTA32 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	Disability                        []DB1
	ObservationResult                 []OBX
}

// NewADTA32Slice will construct a slice of type ADTA32
func NewADTA32Slice(input hl7.HL7Type) []ADTA32 {
	values := make([]ADTA32, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA32(input.Index(i))
	}
	return values
}

// NewOULR23 creates an implementation of OULR23
func NewOULR23(input hl7.HL7Type) OULR23 {
	v := OULR23{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTE(input.Index(3))
	v.Patient = NewOULR23Patient(input.Index(4))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(5))
	v.Specimen = NewOULR23SpecimenSlice(input.Index(6))
	v.ContinuationPointer = NewDSC(input.Index(7))
	return v
}

type OULR23 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              NTE
	Patient                       OULR23Patient
	NextOfKinAssociatedParties    []NK1
	Specimen                      []OULR23Specimen
	ContinuationPointer           DSC
}

// NewOULR23Slice will construct a slice of type OULR23
func NewOULR23Slice(input hl7.HL7Type) []OULR23 {
	values := make([]OULR23, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOULR23(input.Index(i))
	}
	return values
}

// NewBARP12 creates an implementation of BARP12
func NewBARP12(input hl7.HL7Type) BARP12 {
	v := BARP12{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientVisit = NewPV1(input.Index(5))
	v.Diagnosis = NewDG1Slice(input.Index(6))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(7))
	v.Procedure = NewBARP12ProcedureSlice(input.Index(8))
	v.ObservationResult = NewOBX(input.Index(9))
	return v
}

type BARP12 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientVisit                  PV1
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         DRG
	Procedure                     []BARP12Procedure
	ObservationResult             OBX
}

// NewBARP12Slice will construct a slice of type BARP12
func NewBARP12Slice(input hl7.HL7Type) []BARP12 {
	values := make([]BARP12, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBARP12(input.Index(i))
	}
	return values
}

// NewRRII12 creates an implementation of RRII12
func NewRRII12(input hl7.HL7Type) RRII12 {
	v := RRII12{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.ReferralInformation = NewRF1(input.Index(4))
	v.AuthorizationContact = NewRRII12AuthorizationContact(input.Index(5))
	v.ProviderContact = NewRRII12ProviderContactSlice(input.Index(6))
	v.PatientIdentification = NewPID(input.Index(7))
	v.Accident = NewACC(input.Index(8))
	v.Diagnosis = NewDG1Slice(input.Index(9))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(10))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(11))
	v.Procedure = NewRRII12ProcedureSlice(input.Index(12))
	v.Observation = NewRRII12ObservationSlice(input.Index(13))
	v.PatientVisit = NewRRII12PatientVisit(input.Index(14))
	v.NotesAndComments = NewNTESlice(input.Index(15))
	return v
}

type RRII12 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	ReferralInformation           RF1
	AuthorizationContact          RRII12AuthorizationContact
	ProviderContact               []RRII12ProviderContact
	PatientIdentification         PID
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []RRII12Procedure
	Observation                   []RRII12Observation
	PatientVisit                  RRII12PatientVisit
	NotesAndComments              []NTE
}

// NewRRII12Slice will construct a slice of type RRII12
func NewRRII12Slice(input hl7.HL7Type) []RRII12 {
	values := make([]RRII12, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRRII12(input.Index(i))
	}
	return values
}

// NewORUR32 creates an implementation of ORUR32
func NewORUR32(input hl7.HL7Type) ORUR32 {
	v := ORUR32{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientAdditionalDemographic = NewPD1(input.Index(4))
	v.SecondParticipationInformation = NewPRTSlice(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.PatientObservation = NewORUR32PatientObservationSlice(input.Index(7))
	v.Visit = NewORUR32Visit(input.Index(8))
	v.CommonOrder = NewORC(input.Index(9))
	v.ThirdParticipationInformation = NewPRTSlice(input.Index(10))
	v.ObservationRequest = NewOBR(input.Index(11))
	v.NotesAndComments = NewNTESlice(input.Index(12))
	v.FourthParticipationInformation = NewPRTSlice(input.Index(13))
	v.TimingQty = NewORUR32TimingQtySlice(input.Index(14))
	v.Observation = NewORUR32ObservationSlice(input.Index(15))
	return v
}

type ORUR32 struct {
	MessageHeader                  MSH
	SoftwareSegment                []SFT
	UserAuthenticationCredentials  UAC
	PatientIdentification          PID
	PatientAdditionalDemographic   PD1
	SecondParticipationInformation []PRT
	AccessRestriction              []ARV
	PatientObservation             []ORUR32PatientObservation
	Visit                          ORUR32Visit
	CommonOrder                    ORC
	ThirdParticipationInformation  []PRT
	ObservationRequest             OBR
	NotesAndComments               []NTE
	FourthParticipationInformation []PRT
	TimingQty                      []ORUR32TimingQty
	Observation                    []ORUR32Observation
}

// NewORUR32Slice will construct a slice of type ORUR32
func NewORUR32Slice(input hl7.HL7Type) []ORUR32 {
	values := make([]ORUR32, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORUR32(input.Index(i))
	}
	return values
}

// NewOMNO07 creates an implementation of OMNO07
func NewOMNO07(input hl7.HL7Type) OMNO07 {
	v := OMNO07{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewOMNO07Patient(input.Index(4))
	v.Order = NewOMNO07OrderSlice(input.Index(5))
	return v
}

type OMNO07 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       OMNO07Patient
	Order                         []OMNO07Order
}

// NewOMNO07Slice will construct a slice of type OMNO07
func NewOMNO07Slice(input hl7.HL7Type) []OMNO07 {
	values := make([]OMNO07, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOMNO07(input.Index(i))
	}
	return values
}

// NewQBPQ13 creates an implementation of QBPQ13
func NewQBPQ13(input hl7.HL7Type) QBPQ13 {
	v := QBPQ13{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.Qbp = NewQBPQ13Qbp(input.Index(4))
	v.TableRowDefinition = NewRDF(input.Index(5))
	v.ResponseControlParameter = NewRCP(input.Index(6))
	v.ContinuationPointer = NewDSC(input.Index(7))
	return v
}

type QBPQ13 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	Qbp                           QBPQ13Qbp
	TableRowDefinition            RDF
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPQ13Slice will construct a slice of type QBPQ13
func NewQBPQ13Slice(input hl7.HL7Type) []QBPQ13 {
	values := make([]QBPQ13, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPQ13(input.Index(i))
	}
	return values
}

// NewCRMC04 creates an implementation of CRMC04
func NewCRMC04(input hl7.HL7Type) CRMC04 {
	v := CRMC04{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Patient = NewCRMC04PatientSlice(input.Index(3))
	return v
}

type CRMC04 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Patient                       []CRMC04Patient
}

// NewCRMC04Slice will construct a slice of type CRMC04
func NewCRMC04Slice(input hl7.HL7Type) []CRMC04 {
	values := make([]CRMC04, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCRMC04(input.Index(i))
	}
	return values
}

// NewCRMC01 creates an implementation of CRMC01
func NewCRMC01(input hl7.HL7Type) CRMC01 {
	v := CRMC01{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Patient = NewCRMC01PatientSlice(input.Index(3))
	return v
}

type CRMC01 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Patient                       []CRMC01Patient
}

// NewCRMC01Slice will construct a slice of type CRMC01
func NewCRMC01Slice(input hl7.HL7Type) []CRMC01 {
	values := make([]CRMC01, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCRMC01(input.Index(i))
	}
	return values
}

// NewADTA09 creates an implementation of ADTA09
func NewADTA09(input hl7.HL7Type) ADTA09 {
	v := ADTA09{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.Disability = NewDB1Slice(input.Index(8))
	v.ObservationResult = NewOBXSlice(input.Index(9))
	return v
}

type ADTA09 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	Disability                        []DB1
	ObservationResult                 []OBX
}

// NewADTA09Slice will construct a slice of type ADTA09
func NewADTA09Slice(input hl7.HL7Type) []ADTA09 {
	values := make([]ADTA09, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA09(input.Index(i))
	}
	return values
}

// NewCSUC10 creates an implementation of CSUC10
func NewCSUC10(input hl7.HL7Type) CSUC10 {
	v := CSUC10{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Patient = NewCSUC10PatientSlice(input.Index(3))
	return v
}

type CSUC10 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Patient                       []CSUC10Patient
}

// NewCSUC10Slice will construct a slice of type CSUC10
func NewCSUC10Slice(input hl7.HL7Type) []CSUC10 {
	values := make([]CSUC10, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCSUC10(input.Index(i))
	}
	return values
}

// NewORLO36 creates an implementation of ORLO36
func NewORLO36(input hl7.HL7Type) ORLO36 {
	v := ORLO36{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewORLO36Response(input.Index(6))
	return v
}

type ORLO36 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      ORLO36Response
}

// NewORLO36Slice will construct a slice of type ORLO36
func NewORLO36Slice(input hl7.HL7Type) []ORLO36 {
	values := make([]ORLO36, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORLO36(input.Index(i))
	}
	return values
}

// NewNMDN02 creates an implementation of NMDN02
func NewNMDN02(input hl7.HL7Type) NMDN02 {
	v := NMDN02{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.ClockAndStatsWithNotes = NewNMDN02ClockAndStatsWithNotesSlice(input.Index(3))
	return v
}

type NMDN02 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	ClockAndStatsWithNotes        []NMDN02ClockAndStatsWithNotes
}

// NewNMDN02Slice will construct a slice of type NMDN02
func NewNMDN02Slice(input hl7.HL7Type) []NMDN02 {
	values := make([]NMDN02, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewNMDN02(input.Index(i))
	}
	return values
}

// NewSRRS10 creates an implementation of SRRS10
func NewSRRS10(input hl7.HL7Type) SRRS10 {
	v := SRRS10{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.Schedule = NewSRRS10Schedule(input.Index(3))
	return v
}

type SRRS10 struct {
	MessageHeader         MSH
	MessageAcknowledgment MSA
	Error                 []ERR
	Schedule              SRRS10Schedule
}

// NewSRRS10Slice will construct a slice of type SRRS10
func NewSRRS10Slice(input hl7.HL7Type) []SRRS10 {
	values := make([]SRRS10, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRRS10(input.Index(i))
	}
	return values
}

// NewMFKM10 creates an implementation of MFKM10
func NewMFKM10(input hl7.HL7Type) MFKM10 {
	v := MFKM10{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.MasterFileIdentification = NewMFI(input.Index(5))
	v.MasterFileAcknowledgment = NewMFASlice(input.Index(6))
	return v
}

type MFKM10 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	MasterFileIdentification      MFI
	MasterFileAcknowledgment      []MFA
}

// NewMFKM10Slice will construct a slice of type MFKM10
func NewMFKM10Slice(input hl7.HL7Type) []MFKM10 {
	values := make([]MFKM10, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFKM10(input.Index(i))
	}
	return values
}

// NewPMUB05 creates an implementation of PMUB05
func NewPMUB05(input hl7.HL7Type) PMUB05 {
	v := PMUB05{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.StaffIdentification = NewSTF(input.Index(4))
	v.PractitionerDetail = NewPRASlice(input.Index(5))
	v.PractitionerOrganizationUnit = NewORGSlice(input.Index(6))
	return v
}

type PMUB05 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	StaffIdentification           STF
	PractitionerDetail            []PRA
	PractitionerOrganizationUnit  []ORG
}

// NewPMUB05Slice will construct a slice of type PMUB05
func NewPMUB05Slice(input hl7.HL7Type) []PMUB05 {
	values := make([]PMUB05, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPMUB05(input.Index(i))
	}
	return values
}

// NewBARP02 creates an implementation of BARP02
func NewBARP02(input hl7.HL7Type) BARP02 {
	v := BARP02{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.Patient = NewBARP02PatientSlice(input.Index(4))
	return v
}

type BARP02 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	Patient                       []BARP02Patient
}

// NewBARP02Slice will construct a slice of type BARP02
func NewBARP02Slice(input hl7.HL7Type) []BARP02 {
	values := make([]BARP02, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBARP02(input.Index(i))
	}
	return values
}

// NewSRMS10 creates an implementation of SRMS10
func NewSRMS10(input hl7.HL7Type) SRMS10 {
	v := SRMS10{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.AppointmentRequest = NewARQ(input.Index(1))
	v.AppointmentPreferences = NewAPR(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSRMS10PatientSlice(input.Index(4))
	v.Resources = NewSRMS10ResourcesSlice(input.Index(5))
	return v
}

type SRMS10 struct {
	MessageHeader          MSH
	AppointmentRequest     ARQ
	AppointmentPreferences APR
	NotesAndComments       []NTE
	Patient                []SRMS10Patient
	Resources              []SRMS10Resources
}

// NewSRMS10Slice will construct a slice of type SRMS10
func NewSRMS10Slice(input hl7.HL7Type) []SRMS10 {
	values := make([]SRMS10, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRMS10(input.Index(i))
	}
	return values
}

// NewMFKM06 creates an implementation of MFKM06
func NewMFKM06(input hl7.HL7Type) MFKM06 {
	v := MFKM06{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.MasterFileIdentification = NewMFI(input.Index(5))
	v.MasterFileAcknowledgment = NewMFASlice(input.Index(6))
	return v
}

type MFKM06 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	MasterFileIdentification      MFI
	MasterFileAcknowledgment      []MFA
}

// NewMFKM06Slice will construct a slice of type MFKM06
func NewMFKM06Slice(input hl7.HL7Type) []MFKM06 {
	values := make([]MFKM06, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFKM06(input.Index(i))
	}
	return values
}

// NewREFI15 creates an implementation of REFI15
func NewREFI15(input hl7.HL7Type) REFI15 {
	v := REFI15{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.ReferralInformation = NewRF1(input.Index(3))
	v.AuthorizationContact = NewREFI15AuthorizationContact(input.Index(4))
	v.ProviderContact = NewREFI15ProviderContactSlice(input.Index(5))
	v.PatientIdentification = NewPID(input.Index(6))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(7))
	v.Guarantor = NewGT1Slice(input.Index(8))
	v.Insurance = NewREFI15InsuranceSlice(input.Index(9))
	v.Accident = NewACC(input.Index(10))
	v.Diagnosis = NewDG1Slice(input.Index(11))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(12))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(13))
	v.Procedure = NewREFI15ProcedureSlice(input.Index(14))
	v.Observation = NewREFI15ObservationSlice(input.Index(15))
	v.PatientVisit = NewREFI15PatientVisit(input.Index(16))
	v.NotesAndComments = NewNTESlice(input.Index(17))
	return v
}

type REFI15 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	ReferralInformation           RF1
	AuthorizationContact          REFI15AuthorizationContact
	ProviderContact               []REFI15ProviderContact
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	Guarantor                     []GT1
	Insurance                     []REFI15Insurance
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []REFI15Procedure
	Observation                   []REFI15Observation
	PatientVisit                  REFI15PatientVisit
	NotesAndComments              []NTE
}

// NewREFI15Slice will construct a slice of type REFI15
func NewREFI15Slice(input hl7.HL7Type) []REFI15 {
	values := make([]REFI15, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewREFI15(input.Index(i))
	}
	return values
}

// NewBARP10 creates an implementation of BARP10
func NewBARP10(input hl7.HL7Type) BARP10 {
	v := BARP10{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientVisit = NewPV1(input.Index(5))
	v.Diagnosis = NewDG1Slice(input.Index(6))
	v.GroupingReimbursementVisit = NewGP1(input.Index(7))
	v.Procedure = NewBARP10ProcedureSlice(input.Index(8))
	return v
}

type BARP10 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientVisit                  PV1
	Diagnosis                     []DG1
	GroupingReimbursementVisit    GP1
	Procedure                     []BARP10Procedure
}

// NewBARP10Slice will construct a slice of type BARP10
func NewBARP10Slice(input hl7.HL7Type) []BARP10 {
	values := make([]BARP10, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBARP10(input.Index(i))
	}
	return values
}

// NewRDYZ80 creates an implementation of RDYZ80
func NewRDYZ80(input hl7.HL7Type) RDYZ80 {
	v := RDYZ80{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.SegmentPattern = NewRDYZ80SegmentPattern(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RDYZ80 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	SegmentPattern                RDYZ80SegmentPattern
	ContinuationPointer           DSC
}

// NewRDYZ80Slice will construct a slice of type RDYZ80
func NewRDYZ80Slice(input hl7.HL7Type) []RDYZ80 {
	values := make([]RDYZ80, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRDYZ80(input.Index(i))
	}
	return values
}

// NewOULR24 creates an implementation of OULR24
func NewOULR24(input hl7.HL7Type) OULR24 {
	v := OULR24{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTE(input.Index(3))
	v.Patient = NewOULR24Patient(input.Index(4))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(5))
	v.Order = NewOULR24OrderSlice(input.Index(6))
	return v
}

type OULR24 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              NTE
	Patient                       OULR24Patient
	NextOfKinAssociatedParties    []NK1
	Order                         []OULR24Order
}

// NewOULR24Slice will construct a slice of type OULR24
func NewOULR24Slice(input hl7.HL7Type) []OULR24 {
	values := make([]OULR24, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOULR24(input.Index(i))
	}
	return values
}

// NewADTA25 creates an implementation of ADTA25
func NewADTA25(input hl7.HL7Type) ADTA25 {
	v := ADTA25{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.Disability = NewDB1Slice(input.Index(8))
	v.ObservationResult = NewOBXSlice(input.Index(9))
	return v
}

type ADTA25 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	Disability                        []DB1
	ObservationResult                 []OBX
}

// NewADTA25Slice will construct a slice of type ADTA25
func NewADTA25Slice(input hl7.HL7Type) []ADTA25 {
	values := make([]ADTA25, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA25(input.Index(i))
	}
	return values
}

// NewREFI13 creates an implementation of REFI13
func NewREFI13(input hl7.HL7Type) REFI13 {
	v := REFI13{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.ReferralInformation = NewRF1(input.Index(3))
	v.AuthorizationContact = NewREFI13AuthorizationContact(input.Index(4))
	v.ProviderContact = NewREFI13ProviderContactSlice(input.Index(5))
	v.PatientIdentification = NewPID(input.Index(6))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(7))
	v.Guarantor = NewGT1Slice(input.Index(8))
	v.Insurance = NewREFI13InsuranceSlice(input.Index(9))
	v.Accident = NewACC(input.Index(10))
	v.Diagnosis = NewDG1Slice(input.Index(11))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(12))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(13))
	v.Procedure = NewREFI13ProcedureSlice(input.Index(14))
	v.Observation = NewREFI13ObservationSlice(input.Index(15))
	v.PatientVisit = NewREFI13PatientVisit(input.Index(16))
	v.NotesAndComments = NewNTESlice(input.Index(17))
	return v
}

type REFI13 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	ReferralInformation           RF1
	AuthorizationContact          REFI13AuthorizationContact
	ProviderContact               []REFI13ProviderContact
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	Guarantor                     []GT1
	Insurance                     []REFI13Insurance
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []REFI13Procedure
	Observation                   []REFI13Observation
	PatientVisit                  REFI13PatientVisit
	NotesAndComments              []NTE
}

// NewREFI13Slice will construct a slice of type REFI13
func NewREFI13Slice(input hl7.HL7Type) []REFI13 {
	values := make([]REFI13, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewREFI13(input.Index(i))
	}
	return values
}

// NewADTA45 creates an implementation of ADTA45
func NewADTA45(input hl7.HL7Type) ADTA45 {
	v := ADTA45{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.MergeInfo = NewADTA45MergeInfoSlice(input.Index(6))
	return v
}

type ADTA45 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientAdditionalDemographic  PD1
	MergeInfo                     []ADTA45MergeInfo
}

// NewADTA45Slice will construct a slice of type ADTA45
func NewADTA45Slice(input hl7.HL7Type) []ADTA45 {
	values := make([]ADTA45, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA45(input.Index(i))
	}
	return values
}

// NewSRMS01 creates an implementation of SRMS01
func NewSRMS01(input hl7.HL7Type) SRMS01 {
	v := SRMS01{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.AppointmentRequest = NewARQ(input.Index(1))
	v.AppointmentPreferences = NewAPR(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSRMS01PatientSlice(input.Index(4))
	v.Resources = NewSRMS01ResourcesSlice(input.Index(5))
	return v
}

type SRMS01 struct {
	MessageHeader          MSH
	AppointmentRequest     ARQ
	AppointmentPreferences APR
	NotesAndComments       []NTE
	Patient                []SRMS01Patient
	Resources              []SRMS01Resources
}

// NewSRMS01Slice will construct a slice of type SRMS01
func NewSRMS01Slice(input hl7.HL7Type) []SRMS01 {
	values := make([]SRMS01, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRMS01(input.Index(i))
	}
	return values
}

// NewSRRS08 creates an implementation of SRRS08
func NewSRRS08(input hl7.HL7Type) SRRS08 {
	v := SRRS08{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.Schedule = NewSRRS08Schedule(input.Index(3))
	return v
}

type SRRS08 struct {
	MessageHeader         MSH
	MessageAcknowledgment MSA
	Error                 []ERR
	Schedule              SRRS08Schedule
}

// NewSRRS08Slice will construct a slice of type SRRS08
func NewSRRS08Slice(input hl7.HL7Type) []SRRS08 {
	values := make([]SRRS08, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRRS08(input.Index(i))
	}
	return values
}

// NewPPGPCJ creates an implementation of PPGPCJ
func NewPPGPCJ(input hl7.HL7Type) PPGPCJ {
	v := PPGPCJ{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientVisit = NewPPGPCJPatientVisit(input.Index(4))
	v.Pathway = NewPPGPCJPathwaySlice(input.Index(5))
	return v
}

type PPGPCJ struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientIdentification         PID
	PatientVisit                  PPGPCJPatientVisit
	Pathway                       []PPGPCJPathway
}

// NewPPGPCJSlice will construct a slice of type PPGPCJ
func NewPPGPCJSlice(input hl7.HL7Type) []PPGPCJ {
	values := make([]PPGPCJ, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPPGPCJ(input.Index(i))
	}
	return values
}

// NewMFNM14 creates an implementation of MFNM14
func NewMFNM14(input hl7.HL7Type) MFNM14 {
	v := MFNM14{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MasterFileIdentification = NewMFI(input.Index(3))
	v.MfSiteDefined = NewMFNM14MfSiteDefinedSlice(input.Index(4))
	return v
}

type MFNM14 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MasterFileIdentification      MFI
	MfSiteDefined                 []MFNM14MfSiteDefined
}

// NewMFNM14Slice will construct a slice of type MFNM14
func NewMFNM14Slice(input hl7.HL7Type) []MFNM14 {
	values := make([]MFNM14, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFNM14(input.Index(i))
	}
	return values
}

// NewSMDS32 creates an implementation of SMDS32
func NewSMDS32(input hl7.HL7Type) SMDS32 {
	v := SMDS32{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.AntiMicrobialDeviceCycleData = NewSMDS32AntiMicrobialDeviceCycleData(input.Index(3))
	return v
}

type SMDS32 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	AntiMicrobialDeviceCycleData  SMDS32AntiMicrobialDeviceCycleData
}

// NewSMDS32Slice will construct a slice of type SMDS32
func NewSMDS32Slice(input hl7.HL7Type) []SMDS32 {
	values := make([]SMDS32, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSMDS32(input.Index(i))
	}
	return values
}

// NewCRMC05 creates an implementation of CRMC05
func NewCRMC05(input hl7.HL7Type) CRMC05 {
	v := CRMC05{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.AccessRestriction = NewARVSlice(input.Index(2))
	v.UserAuthenticationCredentials = NewUAC(input.Index(3))
	v.Patient = NewCRMC05PatientSlice(input.Index(4))
	return v
}

type CRMC05 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	AccessRestriction             []ARV
	UserAuthenticationCredentials UAC
	Patient                       []CRMC05Patient
}

// NewCRMC05Slice will construct a slice of type CRMC05
func NewCRMC05Slice(input hl7.HL7Type) []CRMC05 {
	values := make([]CRMC05, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCRMC05(input.Index(i))
	}
	return values
}

// NewTCUU10 creates an implementation of TCUU10
func NewTCUU10(input hl7.HL7Type) TCUU10 {
	v := TCUU10{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EquipmentDetail = NewEQU(input.Index(3))
	v.TestConfiguration = NewTCUU10TestConfigurationSlice(input.Index(4))
	return v
}

type TCUU10 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EquipmentDetail               EQU
	TestConfiguration             []TCUU10TestConfiguration
}

// NewTCUU10Slice will construct a slice of type TCUU10
func NewTCUU10Slice(input hl7.HL7Type) []TCUU10 {
	values := make([]TCUU10, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewTCUU10(input.Index(i))
	}
	return values
}

// NewRREO26 creates an implementation of RREO26
func NewRREO26(input hl7.HL7Type) RREO26 {
	v := RREO26{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewRREO26Response(input.Index(6))
	return v
}

type RREO26 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      RREO26Response
}

// NewRREO26Slice will construct a slice of type RREO26
func NewRREO26Slice(input hl7.HL7Type) []RREO26 {
	values := make([]RREO26, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRREO26(input.Index(i))
	}
	return values
}

// NewADTA41 creates an implementation of ADTA41
func NewADTA41(input hl7.HL7Type) ADTA41 {
	v := ADTA41{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.Patient = NewADTA41PatientSlice(input.Index(4))
	return v
}

type ADTA41 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	Patient                       []ADTA41Patient
}

// NewADTA41Slice will construct a slice of type ADTA41
func NewADTA41Slice(input hl7.HL7Type) []ADTA41 {
	values := make([]ADTA41, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA41(input.Index(i))
	}
	return values
}

// NewADTA08 creates an implementation of ADTA08
func NewADTA08(input hl7.HL7Type) ADTA08 {
	v := ADTA08{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(8))
	v.PatientVisit = NewPV1(input.Index(9))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(10))
	v.AdditionalAccessRestriction = NewARVSlice(input.Index(11))
	v.AdditionalRole = NewROLSlice(input.Index(12))
	v.Disability = NewDB1Slice(input.Index(13))
	v.ObservationResult = NewOBXSlice(input.Index(14))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(15))
	v.Diagnosis = NewDG1Slice(input.Index(16))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(17))
	v.Procedure = NewADTA08ProcedureSlice(input.Index(18))
	v.Guarantor = NewGT1Slice(input.Index(19))
	v.Insurance = NewADTA08InsuranceSlice(input.Index(20))
	v.Accident = NewACC(input.Index(21))
	v.Ub82 = NewUB1(input.Index(22))
	v.UniformBillingData = NewUB2(input.Index(23))
	v.PatientDeathAndAutopsy = NewPDA(input.Index(24))
	return v
}

type ADTA08 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	AccessRestriction                 []ARV
	Role                              []ROL
	NextOfKinAssociatedParties        []NK1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalAccessRestriction       []ARV
	AdditionalRole                    []ROL
	Disability                        []DB1
	ObservationResult                 []OBX
	PatientAllergyInformation         []AL1
	Diagnosis                         []DG1
	DiagnosisRelatedGroup             DRG
	Procedure                         []ADTA08Procedure
	Guarantor                         []GT1
	Insurance                         []ADTA08Insurance
	Accident                          ACC
	Ub82                              UB1
	UniformBillingData                UB2
	PatientDeathAndAutopsy            PDA
}

// NewADTA08Slice will construct a slice of type ADTA08
func NewADTA08Slice(input hl7.HL7Type) []ADTA08 {
	values := make([]ADTA08, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA08(input.Index(i))
	}
	return values
}

// NewORSO06 creates an implementation of ORSO06
func NewORSO06(input hl7.HL7Type) ORSO06 {
	v := ORSO06{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewORSO06Response(input.Index(6))
	return v
}

type ORSO06 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      ORSO06Response
}

// NewORSO06Slice will construct a slice of type ORSO06
func NewORSO06Slice(input hl7.HL7Type) []ORSO06 {
	values := make([]ORSO06, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORSO06(input.Index(i))
	}
	return values
}

// NewQBPQ32 creates an implementation of QBPQ32
func NewQBPQ32(input hl7.HL7Type) QBPQ32 {
	v := QBPQ32{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.ResponseControlParameter = NewRCP(input.Index(4))
	v.ContinuationPointer = NewDSC(input.Index(5))
	return v
}

type QBPQ32 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPQ32Slice will construct a slice of type QBPQ32
func NewQBPQ32Slice(input hl7.HL7Type) []QBPQ32 {
	values := make([]QBPQ32, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPQ32(input.Index(i))
	}
	return values
}

// NewORLO40 creates an implementation of ORLO40
func NewORLO40(input hl7.HL7Type) ORLO40 {
	v := ORLO40{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewORLO40Response(input.Index(6))
	return v
}

type ORLO40 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      ORLO40Response
}

// NewORLO40Slice will construct a slice of type ORLO40
func NewORLO40Slice(input hl7.HL7Type) []ORLO40 {
	values := make([]ORLO40, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORLO40(input.Index(i))
	}
	return values
}

// NewRPII04 creates an implementation of RPII04
func NewRPII04(input hl7.HL7Type) RPII04 {
	v := RPII04{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Provider = NewRPII04ProviderSlice(input.Index(4))
	v.PatientIdentification = NewPID(input.Index(5))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(6))
	v.GuarantorInsurance = NewRPII04GuarantorInsurance(input.Index(7))
	v.NotesAndComments = NewNTESlice(input.Index(8))
	return v
}

type RPII04 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Provider                      []RPII04Provider
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	GuarantorInsurance            RPII04GuarantorInsurance
	NotesAndComments              []NTE
}

// NewRPII04Slice will construct a slice of type RPII04
func NewRPII04Slice(input hl7.HL7Type) []RPII04 {
	values := make([]RPII04, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRPII04(input.Index(i))
	}
	return values
}

// NewEHCE13 creates an implementation of EHCE13
func NewEHCE13(input hl7.HL7Type) EHCE13 {
	v := EHCE13{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUACSlice(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.RequestForInformation = NewRFI(input.Index(5))
	v.ContactData = NewCTDSlice(input.Index(6))
	v.InvoiceSegment = NewIVC(input.Index(7))
	v.ProductServiceSection = NewPSS(input.Index(8))
	v.ProductServiceGroup = NewPSG(input.Index(9))
	v.PatientIdentification = NewPID(input.Index(10))
	v.ProductServiceLineItem = NewPSL(input.Index(11))
	v.Request = NewEHCE13RequestSlice(input.Index(12))
	return v
}

type EHCE13 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials []UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	RequestForInformation         RFI
	ContactData                   []CTD
	InvoiceSegment                IVC
	ProductServiceSection         PSS
	ProductServiceGroup           PSG
	PatientIdentification         PID
	ProductServiceLineItem        PSL
	Request                       []EHCE13Request
}

// NewEHCE13Slice will construct a slice of type EHCE13
func NewEHCE13Slice(input hl7.HL7Type) []EHCE13 {
	values := make([]EHCE13, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEHCE13(input.Index(i))
	}
	return values
}

// NewOMLO33 creates an implementation of OMLO33
func NewOMLO33(input hl7.HL7Type) OMLO33 {
	v := OMLO33{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewOMLO33Patient(input.Index(4))
	v.Specimen = NewOMLO33SpecimenSlice(input.Index(5))
	return v
}

type OMLO33 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       OMLO33Patient
	Specimen                      []OMLO33Specimen
}

// NewOMLO33Slice will construct a slice of type OMLO33
func NewOMLO33Slice(input hl7.HL7Type) []OMLO33 {
	values := make([]OMLO33, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOMLO33(input.Index(i))
	}
	return values
}

// NewDFTP11 creates an implementation of DFTP11
func NewDFTP11(input hl7.HL7Type) DFTP11 {
	v := DFTP11{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.ParticipationInformation = NewPRTSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.PatientVisit = NewPV1(input.Index(8))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(9))
	v.AdditionalParticipationInformation = NewPRTSlice(input.Index(10))
	v.AdditionalRole = NewROLSlice(input.Index(11))
	v.Disability = NewDB1Slice(input.Index(12))
	v.CommonOrder = NewDFTP11CommonOrderSlice(input.Index(13))
	v.Diagnosis = NewDG1Slice(input.Index(14))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(15))
	v.Guarantor = NewGT1Slice(input.Index(16))
	v.Insurance = NewDFTP11InsuranceSlice(input.Index(17))
	v.Accident = NewACC(input.Index(18))
	v.Financial = NewDFTP11FinancialSlice(input.Index(19))
	return v
}

type DFTP11 struct {
	MessageHeader                      MSH
	SoftwareSegment                    []SFT
	UserAuthenticationCredentials      UAC
	EventType                          EVN
	PatientIdentification              PID
	PatientAdditionalDemographic       PD1
	ParticipationInformation           []PRT
	Role                               []ROL
	PatientVisit                       PV1
	PatientVisitAdditionalInformation  PV2
	AdditionalParticipationInformation []PRT
	AdditionalRole                     []ROL
	Disability                         []DB1
	CommonOrder                        []DFTP11CommonOrder
	Diagnosis                          []DG1
	DiagnosisRelatedGroup              DRG
	Guarantor                          []GT1
	Insurance                          []DFTP11Insurance
	Accident                           ACC
	Financial                          []DFTP11Financial
}

// NewDFTP11Slice will construct a slice of type DFTP11
func NewDFTP11Slice(input hl7.HL7Type) []DFTP11 {
	values := make([]DFTP11, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDFTP11(input.Index(i))
	}
	return values
}

// NewCRMC03 creates an implementation of CRMC03
func NewCRMC03(input hl7.HL7Type) CRMC03 {
	v := CRMC03{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Patient = NewCRMC03PatientSlice(input.Index(3))
	return v
}

type CRMC03 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Patient                       []CRMC03Patient
}

// NewCRMC03Slice will construct a slice of type CRMC03
func NewCRMC03Slice(input hl7.HL7Type) []CRMC03 {
	values := make([]CRMC03, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCRMC03(input.Index(i))
	}
	return values
}

// NewRDYK15 creates an implementation of RDYK15
func NewRDYK15(input hl7.HL7Type) RDYK15 {
	v := RDYK15{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.DisplayData = NewDSPSlice(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RDYK15 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	DisplayData                   []DSP
	ContinuationPointer           DSC
}

// NewRDYK15Slice will construct a slice of type RDYK15
func NewRDYK15Slice(input hl7.HL7Type) []RDYK15 {
	values := make([]RDYK15, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRDYK15(input.Index(i))
	}
	return values
}

// NewEARU08 creates an implementation of EARU08
func NewEARU08(input hl7.HL7Type) EARU08 {
	v := EARU08{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EquipmentDetail = NewEQU(input.Index(3))
	v.CommandResponse = NewEARU08CommandResponseSlice(input.Index(4))
	return v
}

type EARU08 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EquipmentDetail               EQU
	CommandResponse               []EARU08CommandResponse
}

// NewEARU08Slice will construct a slice of type EARU08
func NewEARU08Slice(input hl7.HL7Type) []EARU08 {
	values := make([]EARU08, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEARU08(input.Index(i))
	}
	return values
}

// NewRSPZ88 creates an implementation of RSPZ88
func NewRSPZ88(input hl7.HL7Type) RSPZ88 {
	v := RSPZ88{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.ResponseControlParameter = NewRCP(input.Index(7))
	v.QueryResponse = NewRSPZ88QueryResponseSlice(input.Index(8))
	v.ContinuationPointer = NewDSC(input.Index(9))
	return v
}

type RSPZ88 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
	QueryResponse                 []RSPZ88QueryResponse
	ContinuationPointer           DSC
}

// NewRSPZ88Slice will construct a slice of type RSPZ88
func NewRSPZ88Slice(input hl7.HL7Type) []RSPZ88 {
	values := make([]RSPZ88, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPZ88(input.Index(i))
	}
	return values
}

// NewORXO43 creates an implementation of ORXO43
func NewORXO43(input hl7.HL7Type) ORXO43 {
	v := ORXO43{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewORXO43Response(input.Index(6))
	return v
}

type ORXO43 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      ORXO43Response
}

// NewORXO43Slice will construct a slice of type ORXO43
func NewORXO43Slice(input hl7.HL7Type) []ORXO43 {
	values := make([]ORXO43, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORXO43(input.Index(i))
	}
	return values
}

// NewPGLPC6 creates an implementation of PGLPC6
func NewPGLPC6(input hl7.HL7Type) PGLPC6 {
	v := PGLPC6{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientVisit = NewPGLPC6PatientVisit(input.Index(4))
	v.Goal = NewPGLPC6GoalSlice(input.Index(5))
	return v
}

type PGLPC6 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientIdentification         PID
	PatientVisit                  PGLPC6PatientVisit
	Goal                          []PGLPC6Goal
}

// NewPGLPC6Slice will construct a slice of type PGLPC6
func NewPGLPC6Slice(input hl7.HL7Type) []PGLPC6 {
	values := make([]PGLPC6, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPGLPC6(input.Index(i))
	}
	return values
}

// NewMDMT05 creates an implementation of MDMT05
func NewMDMT05(input hl7.HL7Type) MDMT05 {
	v := MDMT05{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientVisit = NewPV1(input.Index(5))
	v.CommonOrder = NewMDMT05CommonOrderSlice(input.Index(6))
	v.TranscriptionDocumentHeader = NewTXA(input.Index(7))
	v.ConsentSegment = NewCONSlice(input.Index(8))
	return v
}

type MDMT05 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientVisit                  PV1
	CommonOrder                   []MDMT05CommonOrder
	TranscriptionDocumentHeader   TXA
	ConsentSegment                []CON
}

// NewMDMT05Slice will construct a slice of type MDMT05
func NewMDMT05Slice(input hl7.HL7Type) []MDMT05 {
	values := make([]MDMT05, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMDMT05(input.Index(i))
	}
	return values
}

// NewRSPK32 creates an implementation of RSPK32
func NewRSPK32(input hl7.HL7Type) RSPK32 {
	v := RSPK32{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.MessageAcknowledgment = NewMSA(input.Index(2))
	v.Error = NewERR(input.Index(3))
	v.QueryAcknowledgment = NewQAK(input.Index(4))
	v.QueryParameterDefinition = NewQPD(input.Index(5))
	v.QueryResponse = NewRSPK32QueryResponseSlice(input.Index(6))
	v.ContinuationPointer = NewDSC(input.Index(7))
	return v
}

type RSPK32 struct {
	MessageHeader            MSH
	SoftwareSegment          []SFT
	MessageAcknowledgment    MSA
	Error                    ERR
	QueryAcknowledgment      QAK
	QueryParameterDefinition QPD
	QueryResponse            []RSPK32QueryResponse
	ContinuationPointer      DSC
}

// NewRSPK32Slice will construct a slice of type RSPK32
func NewRSPK32Slice(input hl7.HL7Type) []RSPK32 {
	values := make([]RSPK32, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPK32(input.Index(i))
	}
	return values
}

// NewMDMT02 creates an implementation of MDMT02
func NewMDMT02(input hl7.HL7Type) MDMT02 {
	v := MDMT02{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientVisit = NewPV1(input.Index(5))
	v.CommonOrder = NewMDMT02CommonOrderSlice(input.Index(6))
	v.TranscriptionDocumentHeader = NewTXA(input.Index(7))
	v.ConsentSegment = NewCONSlice(input.Index(8))
	v.Observation = NewMDMT02ObservationSlice(input.Index(9))
	return v
}

type MDMT02 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientVisit                  PV1
	CommonOrder                   []MDMT02CommonOrder
	TranscriptionDocumentHeader   TXA
	ConsentSegment                []CON
	Observation                   []MDMT02Observation
}

// NewMDMT02Slice will construct a slice of type MDMT02
func NewMDMT02Slice(input hl7.HL7Type) []MDMT02 {
	values := make([]MDMT02, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMDMT02(input.Index(i))
	}
	return values
}

// NewEHCE02 creates an implementation of EHCE02
func NewEHCE02(input hl7.HL7Type) EHCE02 {
	v := EHCE02{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUACSlice(input.Index(2))
	v.InvoiceInformationCancel = NewEHCE02InvoiceInformationCancel(input.Index(3))
	return v
}

type EHCE02 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials []UAC
	InvoiceInformationCancel      EHCE02InvoiceInformationCancel
}

// NewEHCE02Slice will construct a slice of type EHCE02
func NewEHCE02Slice(input hl7.HL7Type) []EHCE02 {
	values := make([]EHCE02, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEHCE02(input.Index(i))
	}
	return values
}

// NewREFI12 creates an implementation of REFI12
func NewREFI12(input hl7.HL7Type) REFI12 {
	v := REFI12{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.ReferralInformation = NewRF1(input.Index(3))
	v.AuthorizationContact = NewREFI12AuthorizationContact(input.Index(4))
	v.ProviderContact = NewREFI12ProviderContactSlice(input.Index(5))
	v.PatientIdentification = NewPID(input.Index(6))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(7))
	v.Guarantor = NewGT1Slice(input.Index(8))
	v.Insurance = NewREFI12InsuranceSlice(input.Index(9))
	v.Accident = NewACC(input.Index(10))
	v.Diagnosis = NewDG1Slice(input.Index(11))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(12))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(13))
	v.Procedure = NewREFI12ProcedureSlice(input.Index(14))
	v.Observation = NewREFI12ObservationSlice(input.Index(15))
	v.PatientVisit = NewREFI12PatientVisit(input.Index(16))
	v.NotesAndComments = NewNTESlice(input.Index(17))
	return v
}

type REFI12 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	ReferralInformation           RF1
	AuthorizationContact          REFI12AuthorizationContact
	ProviderContact               []REFI12ProviderContact
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	Guarantor                     []GT1
	Insurance                     []REFI12Insurance
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []REFI12Procedure
	Observation                   []REFI12Observation
	PatientVisit                  REFI12PatientVisit
	NotesAndComments              []NTE
}

// NewREFI12Slice will construct a slice of type REFI12
func NewREFI12Slice(input hl7.HL7Type) []REFI12 {
	values := make([]REFI12, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewREFI12(input.Index(i))
	}
	return values
}

// NewSRRS04 creates an implementation of SRRS04
func NewSRRS04(input hl7.HL7Type) SRRS04 {
	v := SRRS04{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.Schedule = NewSRRS04Schedule(input.Index(3))
	return v
}

type SRRS04 struct {
	MessageHeader         MSH
	MessageAcknowledgment MSA
	Error                 []ERR
	Schedule              SRRS04Schedule
}

// NewSRRS04Slice will construct a slice of type SRRS04
func NewSRRS04Slice(input hl7.HL7Type) []SRRS04 {
	values := make([]SRRS04, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRRS04(input.Index(i))
	}
	return values
}

// NewCCQI19 creates an implementation of CCQI19
func NewCCQI19(input hl7.HL7Type) CCQI19 {
	v := CCQI19{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.ReferralInformation = NewRF1(input.Index(3))
	v.ProviderContact = NewCCQI19ProviderContactSlice(input.Index(4))
	v.ClinicalRelationshipSegment = NewRELSlice(input.Index(5))
	return v
}

type CCQI19 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	ReferralInformation           RF1
	ProviderContact               []CCQI19ProviderContact
	ClinicalRelationshipSegment   []REL
}

// NewCCQI19Slice will construct a slice of type CCQI19
func NewCCQI19Slice(input hl7.HL7Type) []CCQI19 {
	values := make([]CCQI19, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCCQI19(input.Index(i))
	}
	return values
}

// NewQBPQ23 creates an implementation of QBPQ23
func NewQBPQ23(input hl7.HL7Type) QBPQ23 {
	v := QBPQ23{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.ResponseControlParameter = NewRCP(input.Index(4))
	v.ContinuationPointer = NewDSC(input.Index(5))
	return v
}

type QBPQ23 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPQ23Slice will construct a slice of type QBPQ23
func NewQBPQ23Slice(input hl7.HL7Type) []QBPQ23 {
	values := make([]QBPQ23, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPQ23(input.Index(i))
	}
	return values
}

// NewSCNS37 creates an implementation of SCNS37
func NewSCNS37(input hl7.HL7Type) SCNS37 {
	v := SCNS37{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.AntiMicrobialDeviceCycleData = NewSCNS37AntiMicrobialDeviceCycleData(input.Index(3))
	return v
}

type SCNS37 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	AntiMicrobialDeviceCycleData  SCNS37AntiMicrobialDeviceCycleData
}

// NewSCNS37Slice will construct a slice of type SCNS37
func NewSCNS37Slice(input hl7.HL7Type) []SCNS37 {
	values := make([]SCNS37, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSCNS37(input.Index(i))
	}
	return values
}

// NewMFNM09 creates an implementation of MFNM09
func NewMFNM09(input hl7.HL7Type) MFNM09 {
	v := MFNM09{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MasterFileIdentification = NewMFI(input.Index(3))
	v.MfTestCategorical = NewMFNM09MfTestCategoricalSlice(input.Index(4))
	return v
}

type MFNM09 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MasterFileIdentification      MFI
	MfTestCategorical             []MFNM09MfTestCategorical
}

// NewMFNM09Slice will construct a slice of type MFNM09
func NewMFNM09Slice(input hl7.HL7Type) []MFNM09 {
	values := make([]MFNM09, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFNM09(input.Index(i))
	}
	return values
}

// NewADTA31 creates an implementation of ADTA31
func NewADTA31(input hl7.HL7Type) ADTA31 {
	v := ADTA31{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(8))
	v.PatientVisit = NewPV1(input.Index(9))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(10))
	v.AdditionalAccessRestriction = NewARVSlice(input.Index(11))
	v.AdditionalRole = NewROLSlice(input.Index(12))
	v.Disability = NewDB1Slice(input.Index(13))
	v.ObservationResult = NewOBXSlice(input.Index(14))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(15))
	v.Diagnosis = NewDG1Slice(input.Index(16))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(17))
	v.Procedure = NewADTA31ProcedureSlice(input.Index(18))
	v.Guarantor = NewGT1Slice(input.Index(19))
	v.Insurance = NewADTA31InsuranceSlice(input.Index(20))
	v.Accident = NewACC(input.Index(21))
	v.Ub82 = NewUB1(input.Index(22))
	v.UniformBillingData = NewUB2(input.Index(23))
	return v
}

type ADTA31 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	AccessRestriction                 []ARV
	Role                              []ROL
	NextOfKinAssociatedParties        []NK1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalAccessRestriction       []ARV
	AdditionalRole                    []ROL
	Disability                        []DB1
	ObservationResult                 []OBX
	PatientAllergyInformation         []AL1
	Diagnosis                         []DG1
	DiagnosisRelatedGroup             DRG
	Procedure                         []ADTA31Procedure
	Guarantor                         []GT1
	Insurance                         []ADTA31Insurance
	Accident                          ACC
	Ub82                              UB1
	UniformBillingData                UB2
}

// NewADTA31Slice will construct a slice of type ADTA31
func NewADTA31Slice(input hl7.HL7Type) []ADTA31 {
	values := make([]ADTA31, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA31(input.Index(i))
	}
	return values
}

// NewPPGPCG creates an implementation of PPGPCG
func NewPPGPCG(input hl7.HL7Type) PPGPCG {
	v := PPGPCG{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientVisit = NewPPGPCGPatientVisit(input.Index(4))
	v.Pathway = NewPPGPCGPathwaySlice(input.Index(5))
	return v
}

type PPGPCG struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientIdentification         PID
	PatientVisit                  PPGPCGPatientVisit
	Pathway                       []PPGPCGPathway
}

// NewPPGPCGSlice will construct a slice of type PPGPCG
func NewPPGPCGSlice(input hl7.HL7Type) []PPGPCG {
	values := make([]PPGPCG, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPPGPCG(input.Index(i))
	}
	return values
}

// NewORNO08 creates an implementation of ORNO08
func NewORNO08(input hl7.HL7Type) ORNO08 {
	v := ORNO08{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewORNO08Response(input.Index(6))
	return v
}

type ORNO08 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      ORNO08Response
}

// NewORNO08Slice will construct a slice of type ORNO08
func NewORNO08Slice(input hl7.HL7Type) []ORNO08 {
	values := make([]ORNO08, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORNO08(input.Index(i))
	}
	return values
}

// NewRRDO14 creates an implementation of RRDO14
func NewRRDO14(input hl7.HL7Type) RRDO14 {
	v := RRDO14{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewRRDO14Response(input.Index(6))
	return v
}

type RRDO14 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      RRDO14Response
}

// NewRRDO14Slice will construct a slice of type RRDO14
func NewRRDO14Slice(input hl7.HL7Type) []RRDO14 {
	values := make([]RRDO14, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRRDO14(input.Index(i))
	}
	return values
}

// NewOMLO35 creates an implementation of OMLO35
func NewOMLO35(input hl7.HL7Type) OMLO35 {
	v := OMLO35{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewOMLO35Patient(input.Index(4))
	v.Specimen = NewOMLO35SpecimenSlice(input.Index(5))
	return v
}

type OMLO35 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       OMLO35Patient
	Specimen                      []OMLO35Specimen
}

// NewOMLO35Slice will construct a slice of type OMLO35
func NewOMLO35Slice(input hl7.HL7Type) []OMLO35 {
	values := make([]OMLO35, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOMLO35(input.Index(i))
	}
	return values
}

// NewRRII15 creates an implementation of RRII15
func NewRRII15(input hl7.HL7Type) RRII15 {
	v := RRII15{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.ReferralInformation = NewRF1(input.Index(4))
	v.AuthorizationContact = NewRRII15AuthorizationContact(input.Index(5))
	v.ProviderContact = NewRRII15ProviderContactSlice(input.Index(6))
	v.PatientIdentification = NewPID(input.Index(7))
	v.Accident = NewACC(input.Index(8))
	v.Diagnosis = NewDG1Slice(input.Index(9))
	v.DiagnosisRelatedGroup = NewDRGSlice(input.Index(10))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(11))
	v.Procedure = NewRRII15ProcedureSlice(input.Index(12))
	v.Observation = NewRRII15ObservationSlice(input.Index(13))
	v.PatientVisit = NewRRII15PatientVisit(input.Index(14))
	v.NotesAndComments = NewNTESlice(input.Index(15))
	return v
}

type RRII15 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	ReferralInformation           RF1
	AuthorizationContact          RRII15AuthorizationContact
	ProviderContact               []RRII15ProviderContact
	PatientIdentification         PID
	Accident                      ACC
	Diagnosis                     []DG1
	DiagnosisRelatedGroup         []DRG
	PatientAllergyInformation     []AL1
	Procedure                     []RRII15Procedure
	Observation                   []RRII15Observation
	PatientVisit                  RRII15PatientVisit
	NotesAndComments              []NTE
}

// NewRRII15Slice will construct a slice of type RRII15
func NewRRII15Slice(input hl7.HL7Type) []RRII15 {
	values := make([]RRII15, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRRII15(input.Index(i))
	}
	return values
}

// NewQBPE03 creates an implementation of QBPE03
func NewQBPE03(input hl7.HL7Type) QBPE03 {
	v := QBPE03{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUACSlice(input.Index(2))
	v.QueryInformation = NewQBPE03QueryInformation(input.Index(3))
	return v
}

type QBPE03 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials []UAC
	QueryInformation              QBPE03QueryInformation
}

// NewQBPE03Slice will construct a slice of type QBPE03
func NewQBPE03Slice(input hl7.HL7Type) []QBPE03 {
	values := make([]QBPE03, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPE03(input.Index(i))
	}
	return values
}

// NewBTSO31 creates an implementation of BTSO31
func NewBTSO31(input hl7.HL7Type) BTSO31 {
	v := BTSO31{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewBTSO31Patient(input.Index(4))
	v.Order = NewBTSO31OrderSlice(input.Index(5))
	return v
}

type BTSO31 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       BTSO31Patient
	Order                         []BTSO31Order
}

// NewBTSO31Slice will construct a slice of type BTSO31
func NewBTSO31Slice(input hl7.HL7Type) []BTSO31 {
	values := make([]BTSO31, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBTSO31(input.Index(i))
	}
	return values
}

// NewSRMS05 creates an implementation of SRMS05
func NewSRMS05(input hl7.HL7Type) SRMS05 {
	v := SRMS05{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.AppointmentRequest = NewARQ(input.Index(1))
	v.AppointmentPreferences = NewAPR(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSRMS05PatientSlice(input.Index(4))
	v.Resources = NewSRMS05ResourcesSlice(input.Index(5))
	return v
}

type SRMS05 struct {
	MessageHeader          MSH
	AppointmentRequest     ARQ
	AppointmentPreferences APR
	NotesAndComments       []NTE
	Patient                []SRMS05Patient
	Resources              []SRMS05Resources
}

// NewSRMS05Slice will construct a slice of type SRMS05
func NewSRMS05Slice(input hl7.HL7Type) []SRMS05 {
	values := make([]SRMS05, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRMS05(input.Index(i))
	}
	return values
}

// NewRSPK31 creates an implementation of RSPK31
func NewRSPK31(input hl7.HL7Type) RSPK31 {
	v := RSPK31{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.ResponseControlParameter = NewRCP(input.Index(7))
	v.Response = NewRSPK31ResponseSlice(input.Index(8))
	v.ContinuationPointer = NewDSC(input.Index(9))
	return v
}

type RSPK31 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	ResponseControlParameter      RCP
	Response                      []RSPK31Response
	ContinuationPointer           DSC
}

// NewRSPK31Slice will construct a slice of type RSPK31
func NewRSPK31Slice(input hl7.HL7Type) []RSPK31 {
	values := make([]RSPK31, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPK31(input.Index(i))
	}
	return values
}

// NewRTBZ76 creates an implementation of RTBZ76
func NewRTBZ76(input hl7.HL7Type) RTBZ76 {
	v := RTBZ76{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.RowDefinition = NewRTBZ76RowDefinition(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RTBZ76 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	RowDefinition                 RTBZ76RowDefinition
	ContinuationPointer           DSC
}

// NewRTBZ76Slice will construct a slice of type RTBZ76
func NewRTBZ76Slice(input hl7.HL7Type) []RTBZ76 {
	values := make([]RTBZ76, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRTBZ76(input.Index(i))
	}
	return values
}

// NewMFNM08 creates an implementation of MFNM08
func NewMFNM08(input hl7.HL7Type) MFNM08 {
	v := MFNM08{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MasterFileIdentification = NewMFI(input.Index(3))
	v.MfTestNumeric = NewMFNM08MfTestNumericSlice(input.Index(4))
	return v
}

type MFNM08 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MasterFileIdentification      MFI
	MfTestNumeric                 []MFNM08MfTestNumeric
}

// NewMFNM08Slice will construct a slice of type MFNM08
func NewMFNM08Slice(input hl7.HL7Type) []MFNM08 {
	values := make([]MFNM08, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFNM08(input.Index(i))
	}
	return values
}

// NewADTA02 creates an implementation of ADTA02
func NewADTA02(input hl7.HL7Type) ADTA02 {
	v := ADTA02{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.PatientVisit = NewPV1(input.Index(8))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(9))
	v.AdditionalAccessRestriction = NewARVSlice(input.Index(10))
	v.AdditionalRole = NewROLSlice(input.Index(11))
	v.Disability = NewDB1Slice(input.Index(12))
	v.ObservationResult = NewOBXSlice(input.Index(13))
	v.PatientDeathAndAutopsy = NewPDA(input.Index(14))
	return v
}

type ADTA02 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	AccessRestriction                 []ARV
	Role                              []ROL
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalAccessRestriction       []ARV
	AdditionalRole                    []ROL
	Disability                        []DB1
	ObservationResult                 []OBX
	PatientDeathAndAutopsy            PDA
}

// NewADTA02Slice will construct a slice of type ADTA02
func NewADTA02Slice(input hl7.HL7Type) []ADTA02 {
	values := make([]ADTA02, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA02(input.Index(i))
	}
	return values
}

// NewRDEO11 creates an implementation of RDEO11
func NewRDEO11(input hl7.HL7Type) RDEO11 {
	v := RDEO11{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewRDEO11Patient(input.Index(4))
	v.Order = NewRDEO11OrderSlice(input.Index(5))
	return v
}

type RDEO11 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       RDEO11Patient
	Order                         []RDEO11Order
}

// NewRDEO11Slice will construct a slice of type RDEO11
func NewRDEO11Slice(input hl7.HL7Type) []RDEO11 {
	values := make([]RDEO11, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRDEO11(input.Index(i))
	}
	return values
}

// NewADTA17 creates an implementation of ADTA17
func NewADTA17(input hl7.HL7Type) ADTA17 {
	v := ADTA17{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.Disability = NewDB1Slice(input.Index(8))
	v.ObservationResult = NewOBXSlice(input.Index(9))
	v.AdditionalPatientIdentification = NewPID(input.Index(10))
	v.AdditionalPatientAdditionalDemographic = NewPD1(input.Index(11))
	v.AdditionalPatientVisit = NewPV1(input.Index(12))
	v.AdditionalPatientVisitAdditionalInformation = NewPV2(input.Index(13))
	v.AdditionalDisability = NewDB1Slice(input.Index(14))
	v.AdditionalObservationResult = NewOBXSlice(input.Index(15))
	return v
}

type ADTA17 struct {
	MessageHeader                               MSH
	SoftwareSegment                             []SFT
	UserAuthenticationCredentials               UAC
	EventType                                   EVN
	PatientIdentification                       PID
	PatientAdditionalDemographic                PD1
	PatientVisit                                PV1
	PatientVisitAdditionalInformation           PV2
	Disability                                  []DB1
	ObservationResult                           []OBX
	AdditionalPatientIdentification             PID
	AdditionalPatientAdditionalDemographic      PD1
	AdditionalPatientVisit                      PV1
	AdditionalPatientVisitAdditionalInformation PV2
	AdditionalDisability                        []DB1
	AdditionalObservationResult                 []OBX
}

// NewADTA17Slice will construct a slice of type ADTA17
func NewADTA17Slice(input hl7.HL7Type) []ADTA17 {
	values := make([]ADTA17, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA17(input.Index(i))
	}
	return values
}

// NewOULR22 creates an implementation of OULR22
func NewOULR22(input hl7.HL7Type) OULR22 {
	v := OULR22{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTE(input.Index(3))
	v.Patient = NewOULR22Patient(input.Index(4))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(5))
	v.Specimen = NewOULR22SpecimenSlice(input.Index(6))
	v.ContinuationPointer = NewDSC(input.Index(7))
	return v
}

type OULR22 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              NTE
	Patient                       OULR22Patient
	NextOfKinAssociatedParties    []NK1
	Specimen                      []OULR22Specimen
	ContinuationPointer           DSC
}

// NewOULR22Slice will construct a slice of type OULR22
func NewOULR22Slice(input hl7.HL7Type) []OULR22 {
	values := make([]OULR22, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOULR22(input.Index(i))
	}
	return values
}

// NewMDMT11 creates an implementation of MDMT11
func NewMDMT11(input hl7.HL7Type) MDMT11 {
	v := MDMT11{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientVisit = NewPV1(input.Index(5))
	v.CommonOrder = NewMDMT11CommonOrderSlice(input.Index(6))
	v.TranscriptionDocumentHeader = NewTXA(input.Index(7))
	v.ConsentSegment = NewCONSlice(input.Index(8))
	return v
}

type MDMT11 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientVisit                  PV1
	CommonOrder                   []MDMT11CommonOrder
	TranscriptionDocumentHeader   TXA
	ConsentSegment                []CON
}

// NewMDMT11Slice will construct a slice of type MDMT11
func NewMDMT11Slice(input hl7.HL7Type) []MDMT11 {
	values := make([]MDMT11, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMDMT11(input.Index(i))
	}
	return values
}

// NewPPPPCC creates an implementation of PPPPCC
func NewPPPPCC(input hl7.HL7Type) PPPPCC {
	v := PPPPCC{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientVisit = NewPPPPCCPatientVisit(input.Index(4))
	v.Pathway = NewPPPPCCPathwaySlice(input.Index(5))
	return v
}

type PPPPCC struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientIdentification         PID
	PatientVisit                  PPPPCCPatientVisit
	Pathway                       []PPPPCCPathway
}

// NewPPPPCCSlice will construct a slice of type PPPPCC
func NewPPPPCCSlice(input hl7.HL7Type) []PPPPCC {
	values := make([]PPPPCC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPPPPCC(input.Index(i))
	}
	return values
}

// NewEHCE12 creates an implementation of EHCE12
func NewEHCE12(input hl7.HL7Type) EHCE12 {
	v := EHCE12{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUACSlice(input.Index(2))
	v.RequestForInformation = NewRFI(input.Index(3))
	v.ContactData = NewCTDSlice(input.Index(4))
	v.InvoiceSegment = NewIVC(input.Index(5))
	v.ProductServiceSection = NewPSS(input.Index(6))
	v.ProductServiceGroup = NewPSG(input.Index(7))
	v.PatientIdentification = NewPID(input.Index(8))
	v.ProductServiceLineItem = NewPSLSlice(input.Index(9))
	v.Request = NewEHCE12RequestSlice(input.Index(10))
	return v
}

type EHCE12 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials []UAC
	RequestForInformation         RFI
	ContactData                   []CTD
	InvoiceSegment                IVC
	ProductServiceSection         PSS
	ProductServiceGroup           PSG
	PatientIdentification         PID
	ProductServiceLineItem        []PSL
	Request                       []EHCE12Request
}

// NewEHCE12Slice will construct a slice of type EHCE12
func NewEHCE12Slice(input hl7.HL7Type) []EHCE12 {
	values := make([]EHCE12, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEHCE12(input.Index(i))
	}
	return values
}

// NewMDMT07 creates an implementation of MDMT07
func NewMDMT07(input hl7.HL7Type) MDMT07 {
	v := MDMT07{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientVisit = NewPV1(input.Index(5))
	v.CommonOrder = NewMDMT07CommonOrderSlice(input.Index(6))
	v.TranscriptionDocumentHeader = NewTXA(input.Index(7))
	v.ConsentSegment = NewCONSlice(input.Index(8))
	return v
}

type MDMT07 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientVisit                  PV1
	CommonOrder                   []MDMT07CommonOrder
	TranscriptionDocumentHeader   TXA
	ConsentSegment                []CON
}

// NewMDMT07Slice will construct a slice of type MDMT07
func NewMDMT07Slice(input hl7.HL7Type) []MDMT07 {
	values := make([]MDMT07, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMDMT07(input.Index(i))
	}
	return values
}

// NewESRU02 creates an implementation of ESRU02
func NewESRU02(input hl7.HL7Type) ESRU02 {
	v := ESRU02{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EquipmentDetail = NewEQU(input.Index(3))
	return v
}

type ESRU02 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EquipmentDetail               EQU
}

// NewESRU02Slice will construct a slice of type ESRU02
func NewESRU02Slice(input hl7.HL7Type) []ESRU02 {
	values := make([]ESRU02, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewESRU02(input.Index(i))
	}
	return values
}

// NewBARP06 creates an implementation of BARP06
func NewBARP06(input hl7.HL7Type) BARP06 {
	v := BARP06{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.Patient = NewBARP06PatientSlice(input.Index(4))
	return v
}

type BARP06 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	Patient                       []BARP06Patient
}

// NewBARP06Slice will construct a slice of type BARP06
func NewBARP06Slice(input hl7.HL7Type) []BARP06 {
	values := make([]BARP06, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewBARP06(input.Index(i))
	}
	return values
}

// NewADTA53 creates an implementation of ADTA53
func NewADTA53(input hl7.HL7Type) ADTA53 {
	v := ADTA53{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	return v
}

type ADTA53 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
}

// NewADTA53Slice will construct a slice of type ADTA53
func NewADTA53Slice(input hl7.HL7Type) []ADTA53 {
	values := make([]ADTA53, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA53(input.Index(i))
	}
	return values
}

// NewQBPZ81 creates an implementation of QBPZ81
func NewQBPZ81(input hl7.HL7Type) QBPZ81 {
	v := QBPZ81{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.Qbp = NewQBPZ81Qbp(input.Index(4))
	v.ResponseControlParameter = NewRCP(input.Index(5))
	v.ContinuationPointer = NewDSC(input.Index(6))
	return v
}

type QBPZ81 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	Qbp                           QBPZ81Qbp
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPZ81Slice will construct a slice of type QBPZ81
func NewQBPZ81Slice(input hl7.HL7Type) []QBPZ81 {
	values := make([]QBPZ81, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPZ81(input.Index(i))
	}
	return values
}

// NewADTA27 creates an implementation of ADTA27
func NewADTA27(input hl7.HL7Type) ADTA27 {
	v := ADTA27{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.Disability = NewDB1Slice(input.Index(8))
	v.ObservationResult = NewOBXSlice(input.Index(9))
	return v
}

type ADTA27 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	Disability                        []DB1
	ObservationResult                 []OBX
}

// NewADTA27Slice will construct a slice of type ADTA27
func NewADTA27Slice(input hl7.HL7Type) []ADTA27 {
	values := make([]ADTA27, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA27(input.Index(i))
	}
	return values
}

// NewADTA21 creates an implementation of ADTA21
func NewADTA21(input hl7.HL7Type) ADTA21 {
	v := ADTA21{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.PatientVisit = NewPV1(input.Index(6))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(7))
	v.Disability = NewDB1Slice(input.Index(8))
	v.ObservationResult = NewOBXSlice(input.Index(9))
	return v
}

type ADTA21 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	Disability                        []DB1
	ObservationResult                 []OBX
}

// NewADTA21Slice will construct a slice of type ADTA21
func NewADTA21Slice(input hl7.HL7Type) []ADTA21 {
	values := make([]ADTA21, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA21(input.Index(i))
	}
	return values
}

// NewRQPI04 creates an implementation of RQPI04
func NewRQPI04(input hl7.HL7Type) RQPI04 {
	v := RQPI04{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Provider = NewRQPI04ProviderSlice(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(5))
	v.Guarantor = NewGT1Slice(input.Index(6))
	v.NotesAndComments = NewNTESlice(input.Index(7))
	return v
}

type RQPI04 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Provider                      []RQPI04Provider
	PatientIdentification         PID
	NextOfKinAssociatedParties    []NK1
	Guarantor                     []GT1
	NotesAndComments              []NTE
}

// NewRQPI04Slice will construct a slice of type RQPI04
func NewRQPI04Slice(input hl7.HL7Type) []RQPI04 {
	values := make([]RQPI04, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRQPI04(input.Index(i))
	}
	return values
}

// NewTCRU11 creates an implementation of TCRU11
func NewTCRU11(input hl7.HL7Type) TCRU11 {
	v := TCRU11{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EquipmentDetail = NewEQU(input.Index(3))
	v.TestConfiguration = NewTCRU11TestConfigurationSlice(input.Index(4))
	return v
}

type TCRU11 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EquipmentDetail               EQU
	TestConfiguration             []TCRU11TestConfiguration
}

// NewTCRU11Slice will construct a slice of type TCRU11
func NewTCRU11Slice(input hl7.HL7Type) []TCRU11 {
	values := make([]TCRU11, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewTCRU11(input.Index(i))
	}
	return values
}

// NewEHCE15 creates an implementation of EHCE15
func NewEHCE15(input hl7.HL7Type) EHCE15 {
	v := EHCE15{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUACSlice(input.Index(2))
	v.PaymentRemittanceHeaderInfo = NewEHCE15PaymentRemittanceHeaderInfo(input.Index(3))
	v.PaymentRemittanceDetailInfo = NewEHCE15PaymentRemittanceDetailInfoSlice(input.Index(4))
	v.AdjustmentPayee = NewEHCE15AdjustmentPayeeSlice(input.Index(5))
	return v
}

type EHCE15 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials []UAC
	PaymentRemittanceHeaderInfo   EHCE15PaymentRemittanceHeaderInfo
	PaymentRemittanceDetailInfo   []EHCE15PaymentRemittanceDetailInfo
	AdjustmentPayee               []EHCE15AdjustmentPayee
}

// NewEHCE15Slice will construct a slice of type EHCE15
func NewEHCE15Slice(input hl7.HL7Type) []EHCE15 {
	values := make([]EHCE15, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEHCE15(input.Index(i))
	}
	return values
}

// NewRSPE22 creates an implementation of RSPE22
func NewRSPE22(input hl7.HL7Type) RSPE22 {
	v := RSPE22{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUACSlice(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.QueryAck = NewRSPE22QueryAck(input.Index(5))
	return v
}

type RSPE22 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials []UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	QueryAck                      RSPE22QueryAck
}

// NewRSPE22Slice will construct a slice of type RSPE22
func NewRSPE22Slice(input hl7.HL7Type) []RSPE22 {
	values := make([]RSPE22, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRSPE22(input.Index(i))
	}
	return values
}

// NewRTBZ92 creates an implementation of RTBZ92
func NewRTBZ92(input hl7.HL7Type) RTBZ92 {
	v := RTBZ92{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERR(input.Index(4))
	v.QueryAcknowledgment = NewQAK(input.Index(5))
	v.QueryParameterDefinition = NewQPD(input.Index(6))
	v.RowDefinition = NewRTBZ92RowDefinition(input.Index(7))
	v.ContinuationPointer = NewDSC(input.Index(8))
	return v
}

type RTBZ92 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         ERR
	QueryAcknowledgment           QAK
	QueryParameterDefinition      QPD
	RowDefinition                 RTBZ92RowDefinition
	ContinuationPointer           DSC
}

// NewRTBZ92Slice will construct a slice of type RTBZ92
func NewRTBZ92Slice(input hl7.HL7Type) []RTBZ92 {
	values := make([]RTBZ92, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRTBZ92(input.Index(i))
	}
	return values
}

// NewMFKM08 creates an implementation of MFKM08
func NewMFKM08(input hl7.HL7Type) MFKM08 {
	v := MFKM08{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.MasterFileIdentification = NewMFI(input.Index(5))
	v.MasterFileAcknowledgment = NewMFASlice(input.Index(6))
	return v
}

type MFKM08 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	MasterFileIdentification      MFI
	MasterFileAcknowledgment      []MFA
}

// NewMFKM08Slice will construct a slice of type MFKM08
func NewMFKM08Slice(input hl7.HL7Type) []MFKM08 {
	values := make([]MFKM08, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFKM08(input.Index(i))
	}
	return values
}

// NewPGLPC8 creates an implementation of PGLPC8
func NewPGLPC8(input hl7.HL7Type) PGLPC8 {
	v := PGLPC8{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.PatientIdentification = NewPID(input.Index(3))
	v.PatientVisit = NewPGLPC8PatientVisit(input.Index(4))
	v.Goal = NewPGLPC8GoalSlice(input.Index(5))
	return v
}

type PGLPC8 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	PatientIdentification         PID
	PatientVisit                  PGLPC8PatientVisit
	Goal                          []PGLPC8Goal
}

// NewPGLPC8Slice will construct a slice of type PGLPC8
func NewPGLPC8Slice(input hl7.HL7Type) []PGLPC8 {
	values := make([]PGLPC8, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPGLPC8(input.Index(i))
	}
	return values
}

// NewEHCE01 creates an implementation of EHCE01
func NewEHCE01(input hl7.HL7Type) EHCE01 {
	v := EHCE01{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUACSlice(input.Index(2))
	v.InvoiceInformationSubmit = NewEHCE01InvoiceInformationSubmit(input.Index(3))
	return v
}

type EHCE01 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials []UAC
	InvoiceInformationSubmit      EHCE01InvoiceInformationSubmit
}

// NewEHCE01Slice will construct a slice of type EHCE01
func NewEHCE01Slice(input hl7.HL7Type) []EHCE01 {
	values := make([]EHCE01, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEHCE01(input.Index(i))
	}
	return values
}

// NewADTA51 creates an implementation of ADTA51
func NewADTA51(input hl7.HL7Type) ADTA51 {
	v := ADTA51{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.MergePatientInformation = NewMRG(input.Index(6))
	v.PatientVisit = NewPV1(input.Index(7))
	return v
}

type ADTA51 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientAdditionalDemographic  PD1
	MergePatientInformation       MRG
	PatientVisit                  PV1
}

// NewADTA51Slice will construct a slice of type ADTA51
func NewADTA51Slice(input hl7.HL7Type) []ADTA51 {
	values := make([]ADTA51, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA51(input.Index(i))
	}
	return values
}

// NewPMUB07 creates an implementation of PMUB07
func NewPMUB07(input hl7.HL7Type) PMUB07 {
	v := PMUB07{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.StaffIdentification = NewSTF(input.Index(4))
	v.PractitionerDetail = NewPRA(input.Index(5))
	v.Certificate = NewPMUB07CertificateSlice(input.Index(6))
	return v
}

type PMUB07 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	StaffIdentification           STF
	PractitionerDetail            PRA
	Certificate                   []PMUB07Certificate
}

// NewPMUB07Slice will construct a slice of type PMUB07
func NewPMUB07Slice(input hl7.HL7Type) []PMUB07 {
	values := make([]PMUB07, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPMUB07(input.Index(i))
	}
	return values
}

// NewMDMT09 creates an implementation of MDMT09
func NewMDMT09(input hl7.HL7Type) MDMT09 {
	v := MDMT09{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientVisit = NewPV1(input.Index(5))
	v.CommonOrder = NewMDMT09CommonOrderSlice(input.Index(6))
	v.TranscriptionDocumentHeader = NewTXA(input.Index(7))
	v.ConsentSegment = NewCONSlice(input.Index(8))
	return v
}

type MDMT09 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientVisit                  PV1
	CommonOrder                   []MDMT09CommonOrder
	TranscriptionDocumentHeader   TXA
	ConsentSegment                []CON
}

// NewMDMT09Slice will construct a slice of type MDMT09
func NewMDMT09Slice(input hl7.HL7Type) []MDMT09 {
	values := make([]MDMT09, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMDMT09(input.Index(i))
	}
	return values
}

// NewDRCO47 creates an implementation of DRCO47
func NewDRCO47(input hl7.HL7Type) DRCO47 {
	v := DRCO47{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.StaffIdentification = NewSTFSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Donnor = NewDRCO47Donnor(input.Index(3))
	v.DonnorOrder = NewDRCO47DonnorOrderSlice(input.Index(4))
	return v
}

type DRCO47 struct {
	MessageHeader                 MSH
	StaffIdentification           []STF
	UserAuthenticationCredentials UAC
	Donnor                        DRCO47Donnor
	DonnorOrder                   []DRCO47DonnorOrder
}

// NewDRCO47Slice will construct a slice of type DRCO47
func NewDRCO47Slice(input hl7.HL7Type) []DRCO47 {
	values := make([]DRCO47, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDRCO47(input.Index(i))
	}
	return values
}

// NewSIUS23 creates an implementation of SIUS23
func NewSIUS23(input hl7.HL7Type) SIUS23 {
	v := SIUS23{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SchedulingActivityInformation = NewSCH(input.Index(1))
	v.TimingQuantity = NewTQ1Slice(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSIUS23PatientSlice(input.Index(4))
	v.Resources = NewSIUS23ResourcesSlice(input.Index(5))
	return v
}

type SIUS23 struct {
	MessageHeader                 MSH
	SchedulingActivityInformation SCH
	TimingQuantity                []TQ1
	NotesAndComments              []NTE
	Patient                       []SIUS23Patient
	Resources                     []SIUS23Resources
}

// NewSIUS23Slice will construct a slice of type SIUS23
func NewSIUS23Slice(input hl7.HL7Type) []SIUS23 {
	values := make([]SIUS23, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSIUS23(input.Index(i))
	}
	return values
}

// NewMFNM12 creates an implementation of MFNM12
func NewMFNM12(input hl7.HL7Type) MFNM12 {
	v := MFNM12{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MasterFileIdentification = NewMFI(input.Index(3))
	v.MfObsAttributes = NewMFNM12MfObsAttributesSlice(input.Index(4))
	return v
}

type MFNM12 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MasterFileIdentification      MFI
	MfObsAttributes               []MFNM12MfObsAttributes
}

// NewMFNM12Slice will construct a slice of type MFNM12
func NewMFNM12Slice(input hl7.HL7Type) []MFNM12 {
	values := make([]MFNM12, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFNM12(input.Index(i))
	}
	return values
}

// NewCSUC09 creates an implementation of CSUC09
func NewCSUC09(input hl7.HL7Type) CSUC09 {
	v := CSUC09{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Patient = NewCSUC09PatientSlice(input.Index(3))
	return v
}

type CSUC09 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	Patient                       []CSUC09Patient
}

// NewCSUC09Slice will construct a slice of type CSUC09
func NewCSUC09Slice(input hl7.HL7Type) []CSUC09 {
	values := make([]CSUC09, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCSUC09(input.Index(i))
	}
	return values
}

// NewPEXP08 creates an implementation of PEXP08
func NewPEXP08(input hl7.HL7Type) PEXP08 {
	v := PEXP08{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.ParticipationInformation = NewPRTSlice(input.Index(6))
	v.AccessRestriction = NewARVSlice(input.Index(7))
	v.NotesAndComments = NewNTESlice(input.Index(8))
	v.Visit = NewPEXP08Visit(input.Index(9))
	v.Experience = NewPEXP08ExperienceSlice(input.Index(10))
	return v
}

type PEXP08 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	PatientIdentification         PID
	PatientAdditionalDemographic  PD1
	ParticipationInformation      []PRT
	AccessRestriction             []ARV
	NotesAndComments              []NTE
	Visit                         PEXP08Visit
	Experience                    []PEXP08Experience
}

// NewPEXP08Slice will construct a slice of type PEXP08
func NewPEXP08Slice(input hl7.HL7Type) []PEXP08 {
	values := make([]PEXP08, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPEXP08(input.Index(i))
	}
	return values
}

// NewMFKM09 creates an implementation of MFKM09
func NewMFKM09(input hl7.HL7Type) MFKM09 {
	v := MFKM09{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MessageAcknowledgment = NewMSA(input.Index(3))
	v.Error = NewERRSlice(input.Index(4))
	v.MasterFileIdentification = NewMFI(input.Index(5))
	v.MasterFileAcknowledgment = NewMFASlice(input.Index(6))
	return v
}

type MFKM09 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MessageAcknowledgment         MSA
	Error                         []ERR
	MasterFileIdentification      MFI
	MasterFileAcknowledgment      []MFA
}

// NewMFKM09Slice will construct a slice of type MFKM09
func NewMFKM09Slice(input hl7.HL7Type) []MFKM09 {
	values := make([]MFKM09, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFKM09(input.Index(i))
	}
	return values
}

// NewORGO20 creates an implementation of ORGO20
func NewORGO20(input hl7.HL7Type) ORGO20 {
	v := ORGO20{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.Response = NewORGO20Response(input.Index(6))
	return v
}

type ORGO20 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Response                      ORGO20Response
}

// NewORGO20Slice will construct a slice of type ORGO20
func NewORGO20Slice(input hl7.HL7Type) []ORGO20 {
	values := make([]ORGO20, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewORGO20(input.Index(i))
	}
	return values
}

// NewOMDO03 creates an implementation of OMDO03
func NewOMDO03(input hl7.HL7Type) OMDO03 {
	v := OMDO03{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewOMDO03Patient(input.Index(4))
	v.OrderDiet = NewOMDO03OrderDietSlice(input.Index(5))
	v.OrderTray = NewOMDO03OrderTraySlice(input.Index(6))
	return v
}

type OMDO03 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	Patient                       OMDO03Patient
	OrderDiet                     []OMDO03OrderDiet
	OrderTray                     []OMDO03OrderTray
}

// NewOMDO03Slice will construct a slice of type OMDO03
func NewOMDO03Slice(input hl7.HL7Type) []OMDO03 {
	values := make([]OMDO03, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOMDO03(input.Index(i))
	}
	return values
}

// NewADTA06 creates an implementation of ADTA06
func NewADTA06(input hl7.HL7Type) ADTA06 {
	v := ADTA06{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.MergePatientInformation = NewMRG(input.Index(8))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(9))
	v.PatientVisit = NewPV1(input.Index(10))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(11))
	v.AdditionalAccessRestriction = NewARVSlice(input.Index(12))
	v.AdditionalRole = NewROLSlice(input.Index(13))
	v.Disability = NewDB1Slice(input.Index(14))
	v.ObservationResult = NewOBXSlice(input.Index(15))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(16))
	v.Diagnosis = NewDG1Slice(input.Index(17))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(18))
	v.Procedure = NewADTA06ProcedureSlice(input.Index(19))
	v.Guarantor = NewGT1Slice(input.Index(20))
	v.Insurance = NewADTA06InsuranceSlice(input.Index(21))
	v.Accident = NewACC(input.Index(22))
	v.Ub82 = NewUB1(input.Index(23))
	v.UniformBillingData = NewUB2(input.Index(24))
	return v
}

type ADTA06 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	AccessRestriction                 []ARV
	Role                              []ROL
	MergePatientInformation           MRG
	NextOfKinAssociatedParties        []NK1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalAccessRestriction       []ARV
	AdditionalRole                    []ROL
	Disability                        []DB1
	ObservationResult                 []OBX
	PatientAllergyInformation         []AL1
	Diagnosis                         []DG1
	DiagnosisRelatedGroup             DRG
	Procedure                         []ADTA06Procedure
	Guarantor                         []GT1
	Insurance                         []ADTA06Insurance
	Accident                          ACC
	Ub82                              UB1
	UniformBillingData                UB2
}

// NewADTA06Slice will construct a slice of type ADTA06
func NewADTA06Slice(input hl7.HL7Type) []ADTA06 {
	values := make([]ADTA06, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA06(input.Index(i))
	}
	return values
}

// NewSIUS15 creates an implementation of SIUS15
func NewSIUS15(input hl7.HL7Type) SIUS15 {
	v := SIUS15{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SchedulingActivityInformation = NewSCH(input.Index(1))
	v.TimingQuantity = NewTQ1Slice(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSIUS15PatientSlice(input.Index(4))
	v.Resources = NewSIUS15ResourcesSlice(input.Index(5))
	return v
}

type SIUS15 struct {
	MessageHeader                 MSH
	SchedulingActivityInformation SCH
	TimingQuantity                []TQ1
	NotesAndComments              []NTE
	Patient                       []SIUS15Patient
	Resources                     []SIUS15Resources
}

// NewSIUS15Slice will construct a slice of type SIUS15
func NewSIUS15Slice(input hl7.HL7Type) []SIUS15 {
	values := make([]SIUS15, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSIUS15(input.Index(i))
	}
	return values
}

// NewADTA44 creates an implementation of ADTA44
func NewADTA44(input hl7.HL7Type) ADTA44 {
	v := ADTA44{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.Patient = NewADTA44PatientSlice(input.Index(4))
	return v
}

type ADTA44 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	Patient                       []ADTA44Patient
}

// NewADTA44Slice will construct a slice of type ADTA44
func NewADTA44Slice(input hl7.HL7Type) []ADTA44 {
	values := make([]ADTA44, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA44(input.Index(i))
	}
	return values
}

// NewMFNM15 creates an implementation of MFNM15
func NewMFNM15(input hl7.HL7Type) MFNM15 {
	v := MFNM15{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.MasterFileIdentification = NewMFI(input.Index(3))
	v.MfInvItem = NewMFNM15MfInvItemSlice(input.Index(4))
	return v
}

type MFNM15 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	MasterFileIdentification      MFI
	MfInvItem                     []MFNM15MfInvItem
}

// NewMFNM15Slice will construct a slice of type MFNM15
func NewMFNM15Slice(input hl7.HL7Type) []MFNM15 {
	values := make([]MFNM15, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMFNM15(input.Index(i))
	}
	return values
}

// NewQBPZ95 creates an implementation of QBPZ95
func NewQBPZ95(input hl7.HL7Type) QBPZ95 {
	v := QBPZ95{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.QueryParameterDefinition = NewQPD(input.Index(3))
	v.Qbp = NewQBPZ95Qbp(input.Index(4))
	v.TableRowDefinition = NewRDF(input.Index(5))
	v.ResponseControlParameter = NewRCP(input.Index(6))
	v.ContinuationPointer = NewDSC(input.Index(7))
	return v
}

type QBPZ95 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	QueryParameterDefinition      QPD
	Qbp                           QBPZ95Qbp
	TableRowDefinition            RDF
	ResponseControlParameter      RCP
	ContinuationPointer           DSC
}

// NewQBPZ95Slice will construct a slice of type QBPZ95
func NewQBPZ95Slice(input hl7.HL7Type) []QBPZ95 {
	values := make([]QBPZ95, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQBPZ95(input.Index(i))
	}
	return values
}

// NewSIUS17 creates an implementation of SIUS17
func NewSIUS17(input hl7.HL7Type) SIUS17 {
	v := SIUS17{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SchedulingActivityInformation = NewSCH(input.Index(1))
	v.TimingQuantity = NewTQ1Slice(input.Index(2))
	v.NotesAndComments = NewNTESlice(input.Index(3))
	v.Patient = NewSIUS17PatientSlice(input.Index(4))
	v.Resources = NewSIUS17ResourcesSlice(input.Index(5))
	return v
}

type SIUS17 struct {
	MessageHeader                 MSH
	SchedulingActivityInformation SCH
	TimingQuantity                []TQ1
	NotesAndComments              []NTE
	Patient                       []SIUS17Patient
	Resources                     []SIUS17Resources
}

// NewSIUS17Slice will construct a slice of type SIUS17
func NewSIUS17Slice(input hl7.HL7Type) []SIUS17 {
	values := make([]SIUS17, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSIUS17(input.Index(i))
	}
	return values
}

// NewADTA03 creates an implementation of ADTA03
func NewADTA03(input hl7.HL7Type) ADTA03 {
	v := ADTA03{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(8))
	v.PatientVisit = NewPV1(input.Index(9))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(10))
	v.AdditionalAccessRestriction = NewARVSlice(input.Index(11))
	v.AdditionalRole = NewROLSlice(input.Index(12))
	v.Disability = NewDB1Slice(input.Index(13))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(14))
	v.Diagnosis = NewDG1Slice(input.Index(15))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(16))
	v.Procedure = NewADTA03ProcedureSlice(input.Index(17))
	v.ObservationResult = NewOBXSlice(input.Index(18))
	v.Guarantor = NewGT1Slice(input.Index(19))
	v.Insurance = NewADTA03InsuranceSlice(input.Index(20))
	v.Accident = NewACC(input.Index(21))
	v.PatientDeathAndAutopsy = NewPDA(input.Index(22))
	return v
}

type ADTA03 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	AccessRestriction                 []ARV
	Role                              []ROL
	NextOfKinAssociatedParties        []NK1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalAccessRestriction       []ARV
	AdditionalRole                    []ROL
	Disability                        []DB1
	PatientAllergyInformation         []AL1
	Diagnosis                         []DG1
	DiagnosisRelatedGroup             DRG
	Procedure                         []ADTA03Procedure
	ObservationResult                 []OBX
	Guarantor                         []GT1
	Insurance                         []ADTA03Insurance
	Accident                          ACC
	PatientDeathAndAutopsy            PDA
}

// NewADTA03Slice will construct a slice of type ADTA03
func NewADTA03Slice(input hl7.HL7Type) []ADTA03 {
	values := make([]ADTA03, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA03(input.Index(i))
	}
	return values
}

// NewOSUO41 creates an implementation of OSUO41
func NewOSUO41(input hl7.HL7Type) OSUO41 {
	v := OSUO41{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.MessageAcknowledgment = NewMSA(input.Index(1))
	v.Error = NewERRSlice(input.Index(2))
	v.SoftwareSegment = NewSFTSlice(input.Index(3))
	v.UserAuthenticationCredentials = NewUAC(input.Index(4))
	v.NotesAndComments = NewNTESlice(input.Index(5))
	v.PatientIdentification = NewPID(input.Index(6))
	v.AccessRestriction = NewARVSlice(input.Index(7))
	v.OrderStatus = NewOSUO41OrderStatusSlice(input.Index(8))
	return v
}

type OSUO41 struct {
	MessageHeader                 MSH
	MessageAcknowledgment         MSA
	Error                         []ERR
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	NotesAndComments              []NTE
	PatientIdentification         PID
	AccessRestriction             []ARV
	OrderStatus                   []OSUO41OrderStatus
}

// NewOSUO41Slice will construct a slice of type OSUO41
func NewOSUO41Slice(input hl7.HL7Type) []OSUO41 {
	values := make([]OSUO41, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOSUO41(input.Index(i))
	}
	return values
}

// NewADTA05 creates an implementation of ADTA05
func NewADTA05(input hl7.HL7Type) ADTA05 {
	v := ADTA05{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.PatientIdentification = NewPID(input.Index(4))
	v.PatientAdditionalDemographic = NewPD1(input.Index(5))
	v.AccessRestriction = NewARVSlice(input.Index(6))
	v.Role = NewROLSlice(input.Index(7))
	v.NextOfKinAssociatedParties = NewNK1Slice(input.Index(8))
	v.PatientVisit = NewPV1(input.Index(9))
	v.PatientVisitAdditionalInformation = NewPV2(input.Index(10))
	v.AdditionalAccessRestriction = NewARVSlice(input.Index(11))
	v.AdditionalRole = NewROLSlice(input.Index(12))
	v.Disability = NewDB1Slice(input.Index(13))
	v.ObservationResult = NewOBXSlice(input.Index(14))
	v.PatientAllergyInformation = NewAL1Slice(input.Index(15))
	v.Diagnosis = NewDG1Slice(input.Index(16))
	v.DiagnosisRelatedGroup = NewDRG(input.Index(17))
	v.Procedure = NewADTA05ProcedureSlice(input.Index(18))
	v.Guarantor = NewGT1Slice(input.Index(19))
	v.Insurance = NewADTA05InsuranceSlice(input.Index(20))
	v.Accident = NewACC(input.Index(21))
	v.Ub82 = NewUB1(input.Index(22))
	v.UniformBillingData = NewUB2(input.Index(23))
	return v
}

type ADTA05 struct {
	MessageHeader                     MSH
	SoftwareSegment                   []SFT
	UserAuthenticationCredentials     UAC
	EventType                         EVN
	PatientIdentification             PID
	PatientAdditionalDemographic      PD1
	AccessRestriction                 []ARV
	Role                              []ROL
	NextOfKinAssociatedParties        []NK1
	PatientVisit                      PV1
	PatientVisitAdditionalInformation PV2
	AdditionalAccessRestriction       []ARV
	AdditionalRole                    []ROL
	Disability                        []DB1
	ObservationResult                 []OBX
	PatientAllergyInformation         []AL1
	Diagnosis                         []DG1
	DiagnosisRelatedGroup             DRG
	Procedure                         []ADTA05Procedure
	Guarantor                         []GT1
	Insurance                         []ADTA05Insurance
	Accident                          ACC
	Ub82                              UB1
	UniformBillingData                UB2
}

// NewADTA05Slice will construct a slice of type ADTA05
func NewADTA05Slice(input hl7.HL7Type) []ADTA05 {
	values := make([]ADTA05, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA05(input.Index(i))
	}
	return values
}

// NewDERO44 creates an implementation of DERO44
func NewDERO44(input hl7.HL7Type) DERO44 {
	v := DERO44{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.StaffIdentification = NewSTFSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.Donnor = NewDERO44Donnor(input.Index(3))
	v.DonnorOrder = NewDERO44DonnorOrderSlice(input.Index(4))
	return v
}

type DERO44 struct {
	MessageHeader                 MSH
	StaffIdentification           []STF
	UserAuthenticationCredentials UAC
	Donnor                        DERO44Donnor
	DonnorOrder                   []DERO44DonnorOrder
}

// NewDERO44Slice will construct a slice of type DERO44
func NewDERO44Slice(input hl7.HL7Type) []DERO44 {
	values := make([]DERO44, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDERO44(input.Index(i))
	}
	return values
}

// NewADTA20 creates an implementation of ADTA20
func NewADTA20(input hl7.HL7Type) ADTA20 {
	v := ADTA20{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.EventType = NewEVN(input.Index(3))
	v.BedStatusUpdate = NewNPU(input.Index(4))
	return v
}

type ADTA20 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	EventType                     EVN
	BedStatusUpdate               NPU
}

// NewADTA20Slice will construct a slice of type ADTA20
func NewADTA20Slice(input hl7.HL7Type) []ADTA20 {
	values := make([]ADTA20, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewADTA20(input.Index(i))
	}
	return values
}

// NewSTIS30 creates an implementation of STIS30
func NewSTIS30(input hl7.HL7Type) STIS30 {
	v := STIS30{}
	v.MessageHeader = NewMSH(input.Index(0))
	v.SoftwareSegment = NewSFTSlice(input.Index(1))
	v.UserAuthenticationCredentials = NewUAC(input.Index(2))
	v.SterilizationLot = NewSLTSlice(input.Index(3))
	return v
}

type STIS30 struct {
	MessageHeader                 MSH
	SoftwareSegment               []SFT
	UserAuthenticationCredentials UAC
	SterilizationLot              []SLT
}

// NewSTIS30Slice will construct a slice of type STIS30
func NewSTIS30Slice(input hl7.HL7Type) []STIS30 {
	values := make([]STIS30, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSTIS30(input.Index(i))
	}
	return values
}
