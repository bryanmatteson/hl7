package hl7v28

import (
	hl7 "hl7"

	layout "hl7/internal/layout"
)

var layoutQSBZ83 = layout.Group("QSBZ83").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQSBZ83Layout(msg hl7.Message) (*QSBZ83, error) {
	result := layoutQSBZ83.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQSBZ83(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCRMC07 = layout.Group("CRMC07").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("CRMC07Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("CRMC07PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("CSR")).With(layout.Maybe(layout.List(layout.Segment("CSP"))))))

func NewCRMC07Layout(msg hl7.Message) (*CRMC07, error) {
	result := layoutCRMC07.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCRMC07(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORUR40 = layout.Group("ORUR40").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("ORUR40PatientResult").With(layout.Maybe(layout.Group("ORUR40Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Group("ORUR40PatientObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.Group("ORUR40Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))).With(layout.List(layout.Group("ORUR40OrderObservation").With(layout.Maybe(layout.Group("ORUR40CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("ORUR40OrderDocument").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("TXA")))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORUR40TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Segment("CTD"))).With(layout.Maybe(layout.List(layout.Group("ORUR40Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Segment("FT1")))).With(layout.Maybe(layout.List(layout.Segment("CTI")))).With(layout.Maybe(layout.List(layout.Group("ORUR40Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("ORUR40SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))))))).With(layout.Maybe(layout.Segment("DSC")))

func NewORUR40Layout(msg hl7.Message) (*ORUR40, error) {
	result := layoutORUR40.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORUR40(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSIUS21 = layout.Group("SIUS21").With(layout.Segment("MSH")).With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SIUS21Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SIUS21Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SIUS21Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS21GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS21LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS21PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSIUS21Layout(msg hl7.Message) (*SIUS21, error) {
	result := layoutSIUS21.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSIUS21(result.Item.(hl7.Group))
	return &v, nil
}

var layoutBRPO30 = layout.Group("BRPO30").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("BRPO30Response").With(layout.Maybe(layout.Group("BRPO30Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Group("BRPO30Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("BRPO30Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Segment("BPO"))).With(layout.Maybe(layout.List(layout.Segment("BPX")))))))))))

func NewBRPO30Layout(msg hl7.Message) (*BRPO30, error) {
	result := layoutBRPO30.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewBRPO30(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA11 = layout.Group("ADTA11").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))

func NewADTA11Layout(msg hl7.Message) (*ADTA11, error) {
	result := layoutADTA11.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA11(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA43 = layout.Group("ADTA43").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.List(layout.Group("ADTA43Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("MRG"))))

func NewADTA43Layout(msg hl7.Message) (*ADTA43, error) {
	result := layoutADTA43.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA43(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOMIO23 = layout.Group("OMIO23").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMIO23Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMIO23PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Group("OMIO23Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("GT1"))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("OMIO23Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OMIO23Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Segment("CTD"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Group("OMIO23Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.List(layout.Segment("IPC")))))

func NewOMIO23Layout(msg hl7.Message) (*OMIO23, error) {
	result := layoutOMIO23.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOMIO23(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRRII14 = layout.Group("RRII14").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("MSA"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("RRII14AuthorizationContact").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("RRII14ProviderContact").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Group("RRII14Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("RRII14AuthorizationContact").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD")))))))).With(layout.Maybe(layout.List(layout.Group("RRII14Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("RRII14ResultsNotes").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("RRII14PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRRII14Layout(msg hl7.Message) (*RRII14, error) {
	result := layoutRRII14.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRRII14(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPZ90 = layout.Group("RSPZ90").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Segment("RCP")).With(layout.List(layout.Group("RSPZ90QueryResponse").With(layout.Maybe(layout.Group("RSPZ90Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("RSPZ90Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))))).With(layout.List(layout.Group("RSPZ90CommonOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("RSPZ90Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Segment("CTD"))).With(layout.List(layout.Group("RSPZ90Observation").With(layout.Maybe(layout.Segment("OBX"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))).With(layout.Maybe(layout.List(layout.Group("RSPZ90Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Segment("OBX"))))))))).With(layout.Segment("DSC"))

func NewRSPZ90Layout(msg hl7.Message) (*RSPZ90, error) {
	result := layoutRSPZ90.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPZ90(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPQ31 = layout.Group("QBPQ31").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("QBPQ31Qbp").With(layout.Maybe(layout.Segment("Hxx"))))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPQ31Layout(msg hl7.Message) (*QBPQ31, error) {
	result := layoutQBPQ31.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPQ31(result.Item.(hl7.Group))
	return &v, nil
}

var layoutEHCE10 = layout.Group("EHCE10").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.List(layout.Segment("UAC")))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.List(layout.Group("EHCE10InvoiceProcessingResultsInfo").With(layout.Segment("IPR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Segment("PYE")).With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Segment("IVC")).With(layout.List(layout.Group("EHCE10ProductServiceSection").With(layout.Segment("PSS")).With(layout.List(layout.Group("EHCE10ProductServiceGroup").With(layout.Segment("PSG")).With(layout.List(layout.Group("EHCE10ProductServiceLineInfo").With(layout.Segment("PSL")).With(layout.Maybe(layout.List(layout.Segment("ADJ"))))))))))))

func NewEHCE10Layout(msg hl7.Message) (*EHCE10, error) {
	result := layoutEHCE10.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewEHCE10(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPMUB04 = layout.Group("PMUB04").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("STF")).With(layout.Maybe(layout.List(layout.Segment("PRA")))).With(layout.Maybe(layout.List(layout.Segment("ORG"))))

func NewPMUB04Layout(msg hl7.Message) (*PMUB04, error) {
	result := layoutPMUB04.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPMUB04(result.Item.(hl7.Group))
	return &v, nil
}

var layoutDELO46 = layout.Group("DELO46").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("STF")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Group("DELO46Donnor").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("DELO46DonorRegistration").With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))).With(layout.Segment("DON")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewDELO46Layout(msg hl7.Message) (*DELO46, error) {
	result := layoutDELO46.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewDELO46(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA28 = layout.Group("ADTA28").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Group("ADTA28Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("ADTA28Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("AUT")))).With(layout.Maybe(layout.List(layout.Segment("RF1"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("UB1"))).With(layout.Maybe(layout.Segment("UB2")))

func NewADTA28Layout(msg hl7.Message) (*ADTA28, error) {
	result := layoutADTA28.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA28(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCCRI18 = layout.Group("CCRI18").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Segment("RF1"))).With(layout.List(layout.Group("CCRI18ProviderContact").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Maybe(layout.List(layout.Group("CCRI18ClinicalOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("CCRI18ClinicalOrderTiming").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.List(layout.Group("CCRI18ClinicalOrderDetail").With(layout.Group("CCRI18ClinicalOrderObject").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("RXO"))).With(layout.Maybe(layout.Segment("ODS"))).With(layout.Maybe(layout.Segment("PR1")))).With(layout.Maybe(layout.List(layout.Group("CCRI18ClinicalOrderObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.List(layout.Group("CCRI18Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Group("CCRI18Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.List(layout.Group("CCRI18AppointmentHistory").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Group("CCRI18Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("CCRI18ResourceDetail").With(layout.Group("CCRI18ResourceObject").With(layout.Maybe(layout.Segment("AIS"))).With(layout.Maybe(layout.Segment("AIG"))).With(layout.Maybe(layout.Segment("AIL"))).With(layout.Maybe(layout.Segment("AIP")))).With(layout.Maybe(layout.List(layout.Group("CCRI18ResourceObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))))))))).With(layout.Maybe(layout.List(layout.Group("CCRI18ClinicalHistory").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("CCRI18ClinicalHistoryDetail").With(layout.Group("CCRI18ClinicalHistoryObject").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("ODS"))).With(layout.Maybe(layout.Segment("PR1"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Segment("AL1"))).With(layout.Maybe(layout.Segment("IAM"))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("RMI"))).With(layout.Maybe(layout.Segment("DB1"))).With(layout.Maybe(layout.Segment("DG1"))).With(layout.Maybe(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Group("CCRI18ClinicalHistoryObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CCRI18RoleClinicalHistory").With(layout.Group("CCRI18RoleClinicalHistoryObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.List(layout.Group("CCRI18PatientVisits").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Group("CCRI18MedicationHistory").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("CCRI18MedicationOrderDetail").With(layout.Segment("RXO")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Group("CCRI18MedicationOrderObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.Group("CCRI18MedicationEncodingDetail").With(layout.Segment("RXE")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Group("CCRI18MedicationEncodingObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.List(layout.Group("CCRI18MedicationAdministrationDetail").With(layout.List(layout.Segment("RXA"))).With(layout.Segment("RXR")).With(layout.Maybe(layout.List(layout.Group("CCRI18MedicationAdministrationObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.Maybe(layout.List(layout.Group("CCRI18Problem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCRI18RoleProblem").With(layout.Group("CCRI18RoleProblemObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CCRI18ProblemObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CCRI18Goal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCRI18RoleGoal").With(layout.Group("CCRI18RoleGoalObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CCRI18GoalObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CCRI18Pathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCRI18RolePathway").With(layout.Group("CCRI18RolePathwayObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CCRI18PathwayObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Segment("REL"))))

func NewCCRI18Layout(msg hl7.Message) (*CCRI18, error) {
	result := layoutCCRI18.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCCRI18(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORAR33 = layout.Group("ORAR33").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.Segment("ORC")))

func NewORAR33Layout(msg hl7.Message) (*ORAR33, error) {
	result := layoutORAR33.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORAR33(result.Item.(hl7.Group))
	return &v, nil
}

var layoutEACU07 = layout.Group("EACU07").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EQU")).With(layout.List(layout.Group("EACU07Command").With(layout.Segment("ECD")).With(layout.Maybe(layout.Segment("TQ1"))).With(layout.Maybe(layout.Group("EACU07SpecimenContainer").With(layout.Segment("SAC")).With(layout.Maybe(layout.List(layout.Segment("OBR")))).With(layout.Maybe(layout.List(layout.Segment("SPM")))))).With(layout.Maybe(layout.Segment("CNS")))))

func NewEACU07Layout(msg hl7.Message) (*EACU07, error) {
	result := layoutEACU07.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewEACU07(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRMS09 = layout.Group("SRMS09").With(layout.Segment("MSH")).With(layout.Segment("ARQ")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRMS09Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Group("SRMS09Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRMS09Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRMS09Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS09GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS09LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS09PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSRMS09Layout(msg hl7.Message) (*SRMS09, error) {
	result := layoutSRMS09.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRMS09(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMDMT06 = layout.Group("MDMT06").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Group("MDMT06CommonOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("MDMT06Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Segment("TXA")).With(layout.Maybe(layout.List(layout.Segment("CON")))).With(layout.List(layout.Group("MDMT06Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))

func NewMDMT06Layout(msg hl7.Message) (*MDMT06, error) {
	result := layoutMDMT06.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMDMT06(result.Item.(hl7.Group))
	return &v, nil
}

var layoutDFTP03 = layout.Group("DFTP03").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Group("DFTP03CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Group("DFTP03TimingQuantity").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("DFTP03Order").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.Maybe(layout.List(layout.Group("DFTP03Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.List(layout.Group("DFTP03Financial").With(layout.Segment("FT1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.Segment("NTE"))).With(layout.Maybe(layout.List(layout.Group("DFTP03FinancialProcedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Group("DFTP03FinancialCommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Group("DFTP03FinancialTimingQuantity").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("DFTP03FinancialOrder").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.Maybe(layout.List(layout.Group("DFTP03FinancialObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("DFTP03Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.Segment("ACC")))

func NewDFTP03Layout(msg hl7.Message) (*DFTP03, error) {
	result := layoutDFTP03.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewDFTP03(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOMLO21 = layout.Group("OMLO21").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMLO21Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.Segment("ARV"))).With(layout.Maybe(layout.Group("OMLO21PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Group("OMLO21Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("GT1"))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("OMLO21Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OMLO21Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("OMLO21ObservationRequest").With(layout.Segment("OBR")).With(layout.Maybe(layout.Segment("TCD"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Segment("CTD"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Group("OMLO21Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Segment("TCD"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("OMLO21Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("OMLO21SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Group("OMLO21Container").With(layout.Segment("SAC")).With(layout.Maybe(layout.List(layout.Group("OMLO21ContainerObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))))))).With(layout.Maybe(layout.List(layout.Group("OMLO21PriorResult").With(layout.Maybe(layout.Group("OMLO21PatientPrior").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))))).With(layout.Maybe(layout.Group("OMLO21PatientVisitPrior").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.List(layout.Group("OMLO21OrderPrior").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OMLO21TimingPrior").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.List(layout.Group("OMLO21ObservationPrior").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))))))).With(layout.Maybe(layout.List(layout.Segment("FT1")))).With(layout.Maybe(layout.List(layout.Segment("CTI")))).With(layout.Maybe(layout.Segment("BLG")))))

func NewOMLO21Layout(msg hl7.Message) (*OMLO21, error) {
	result := layoutOMLO21.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOMLO21(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRASO17 = layout.Group("RASO17").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("RASO17Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Group("RASO17AdditionalDemographics").With(layout.Segment("PD1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.Group("RASO17PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))))))).With(layout.List(layout.Group("RASO17Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("RASO17Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("RASO17OrderDetail").With(layout.Segment("RXO")).With(layout.Maybe(layout.Group("RASO17OrderDetailSupplement").With(layout.List(layout.Segment("NTE"))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Group("RASO17Components").With(layout.Segment("RXC")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))))).With(layout.Maybe(layout.Group("RASO17Encoding").With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Group("RASO17TimingEncoded").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Segment("CDO")))))).With(layout.List(layout.Group("RASO17Administration").With(layout.List(layout.Segment("RXA"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("RXR")).With(layout.Maybe(layout.List(layout.Group("RASO17Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))

func NewRASO17Layout(msg hl7.Message) (*RASO17, error) {
	result := layoutRASO17.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRASO17(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCCII22 = layout.Group("CCII22").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Group("CCII22Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.List(layout.Group("CCII22AppointmentHistory").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Group("CCII22Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("CCII22ResourceDetail").With(layout.Maybe(layout.Group("CCII22ResourceObject").With(layout.Maybe(layout.Segment("AIS"))).With(layout.Maybe(layout.Segment("AIG"))).With(layout.Maybe(layout.Segment("AIL"))).With(layout.Maybe(layout.Segment("AIP"))))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))))))))))).With(layout.Maybe(layout.List(layout.Group("CCII22ClinicalHistory").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("CCII22ClinicalHistoryDetail").With(layout.Maybe(layout.Group("CCII22ClinicalHistoryObject").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("ODS"))).With(layout.Maybe(layout.Segment("PR1"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Segment("AL1"))).With(layout.Maybe(layout.Segment("IAM"))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("RMI"))).With(layout.Maybe(layout.Segment("DB1"))).With(layout.Maybe(layout.Segment("DG1"))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.Segment("PDA"))))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))))).With(layout.Maybe(layout.List(layout.Group("CCII22RoleClinicalHistory").With(layout.Maybe(layout.Group("CCII22RoleClinicalHistoryObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD"))))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.List(layout.Group("CCII22PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Group("CCII22MedicationHistory").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("CCII22MedicationOrderDetail").With(layout.Segment("RXO")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))))).With(layout.Maybe(layout.Group("CCII22MedicationEncodingDetail").With(layout.Segment("RXE")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))))).With(layout.Maybe(layout.List(layout.Group("CCII22MedicationAdministrationDetail").With(layout.List(layout.Segment("RXA"))).With(layout.Segment("RXR")).With(layout.Maybe(layout.List(layout.Segment("OBX"))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.Maybe(layout.List(layout.Group("CCII22Problem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCII22RoleProblem").With(layout.Maybe(layout.Group("CCII22RoleProblemObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD"))))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))))).With(layout.Maybe(layout.List(layout.Group("CCII22Goal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCII22RoleGoal").With(layout.Maybe(layout.Group("CCII22RoalGoalObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD"))))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))))).With(layout.Maybe(layout.List(layout.Group("CCII22Pathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCII22RolePathway").With(layout.Maybe(layout.Group("CCII22RolePathwayObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD"))))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))))).With(layout.Maybe(layout.List(layout.Segment("REL"))))

func NewCCII22Layout(msg hl7.Message) (*CCII22, error) {
	result := layoutCCII22.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCCII22(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFNM16 = layout.Group("MFNM16").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MFI")).With(layout.List(layout.Group("MFNM16MaterialItemRecord").With(layout.Segment("MFE")).With(layout.Segment("ITM")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("MFNM16Sterilization").With(layout.Segment("STZ")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("MFNM16PurchasingVendor").With(layout.Segment("VND")).With(layout.Maybe(layout.List(layout.Group("MFNM16Packaging").With(layout.Segment("PKG")).With(layout.Maybe(layout.List(layout.Segment("PCE")))))))))).With(layout.Maybe(layout.List(layout.Group("MFNM16MaterialLocation").With(layout.Segment("IVT")).With(layout.Maybe(layout.List(layout.Segment("ILT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewMFNM16Layout(msg hl7.Message) (*MFNM16, error) {
	result := layoutMFNM16.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFNM16(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFKM11 = layout.Group("MFKM11").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("MFI"))

func NewMFKM11Layout(msg hl7.Message) (*MFKM11, error) {
	result := layoutMFKM11.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFKM11(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPZ87 = layout.Group("QBPZ87").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("QBPZ87Qbp").With(layout.Maybe(layout.Segment("Hxx"))))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPZ87Layout(msg hl7.Message) (*QBPZ87, error) {
	result := layoutQBPZ87.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPZ87(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPQ15 = layout.Group("QBPQ15").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Segment("Hxx"))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPQ15Layout(msg hl7.Message) (*QBPQ15, error) {
	result := layoutQBPQ15.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPQ15(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA50 = layout.Group("ADTA50").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("MRG")).With(layout.Segment("PV1"))

func NewADTA50Layout(msg hl7.Message) (*ADTA50, error) {
	result := layoutADTA50.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA50(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPQ25 = layout.Group("QBPQ25").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPQ25Layout(msg hl7.Message) (*QBPQ25, error) {
	result := layoutQBPQ25.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPQ25(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRGVO15 = layout.Group("RGVO15").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("RGVO15Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("RGVO15PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))))))).With(layout.List(layout.Group("RGVO15Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("RGVO15Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("RGVO15OrderDetail").With(layout.Segment("RXO")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("RGVO15OrderDetailSupplement").With(layout.List(layout.Segment("NTE"))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Group("RGVO15Components").With(layout.Segment("RXC")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))))).With(layout.Maybe(layout.Group("RGVO15Encoding").With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Group("RGVO15TimingEncoded").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))))).With(layout.List(layout.Group("RGVO15Give").With(layout.Segment("RXG")).With(layout.List(layout.Group("RGVO15TimingGive").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.Segment("CDO"))).With(layout.Maybe(layout.List(layout.Group("RGVO15Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))))

func NewRGVO15Layout(msg hl7.Message) (*RGVO15, error) {
	result := layoutRGVO15.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRGVO15(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFKM17 = layout.Group("MFKM17").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("MFI")).With(layout.Maybe(layout.List(layout.Segment("MFA"))))

func NewMFKM17Layout(msg hl7.Message) (*MFKM17, error) {
	result := layoutMFKM17.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFKM17(result.Item.(hl7.Group))
	return &v, nil
}

var layoutINRU06 = layout.Group("INRU06").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EQU")).With(layout.List(layout.Segment("INV")))

func NewINRU06Layout(msg hl7.Message) (*INRU06, error) {
	result := layoutINRU06.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewINRU06(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORLO22 = layout.Group("ORLO22").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("ORLO22Response").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Group("ORLO22Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORLO22Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("ORLO22ObservationRequest").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORLO22Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Segment("SAC"))))))))))))))

func NewORLO22Layout(msg hl7.Message) (*ORLO22, error) {
	result := layoutORLO22.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORLO22(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRPII01 = layout.Group("RPII01").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("MSA"))).With(layout.List(layout.Group("RPII01Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.Group("RPII01GuarantorInsurance").With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.List(layout.Group("RPII01Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3"))))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRPII01Layout(msg hl7.Message) (*RPII01, error) {
	result := layoutRPII01.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRPII01(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMDMT08 = layout.Group("MDMT08").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Group("MDMT08CommonOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("MDMT08Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Segment("TXA")).With(layout.Maybe(layout.List(layout.Segment("CON")))).With(layout.List(layout.Group("MDMT08Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))

func NewMDMT08Layout(msg hl7.Message) (*MDMT08, error) {
	result := layoutMDMT08.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMDMT08(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPPPPCD = layout.Group("PPPPCD").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Group("PPPPCDPatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.List(layout.Group("PPPPCDPathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPPPCDPathwayRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPPPCDProblem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPPPCDProblemRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPPPCDProblemObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("PPPPCDGoal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPPPCDGoalRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPPPCDGoalObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.List(layout.Group("PPPPCDOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("PPPPCDOrderDetail").With(layout.Group("PPPPCDChoice").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("Hxx")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPPPCDOrderObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))))))))))))))))

func NewPPPPCDLayout(msg hl7.Message) (*PPPPCD, error) {
	result := layoutPPPPCD.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPPPPCD(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCCRI16 = layout.Group("CCRI16").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Segment("RF1"))).With(layout.List(layout.Group("CCRI16ProviderContact").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Maybe(layout.List(layout.Group("CCRI16ClinicalOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("CCRI16ClinicalOrderTiming").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.List(layout.Group("CCRI16ClinicalOrderDetail").With(layout.Group("CCRI16ClinicalOrderObject").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("RXO"))).With(layout.Maybe(layout.Segment("ODS"))).With(layout.Maybe(layout.Segment("PR1")))).With(layout.Maybe(layout.List(layout.Group("CCRI16ClinicalOrderObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.List(layout.Group("CCRI16Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Group("CCRI16Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.List(layout.Group("CCRI16AppointmentHistory").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Group("CCRI16Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("CCRI16ResourceDetail").With(layout.Group("CCRI16ResourceObject").With(layout.Maybe(layout.Segment("AIS"))).With(layout.Maybe(layout.Segment("AIG"))).With(layout.Maybe(layout.Segment("AIL"))).With(layout.Maybe(layout.Segment("AIP")))).With(layout.Maybe(layout.List(layout.Group("CCRI16ResourceObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))))))))).With(layout.Maybe(layout.List(layout.Group("CCRI16ClinicalHistory").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("CCRI16ClinicalHistoryDetail").With(layout.Group("CCRI16ClinicalHistoryObject").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("ODS"))).With(layout.Maybe(layout.Segment("PR1"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Segment("AL1"))).With(layout.Maybe(layout.Segment("IAM"))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("RMI"))).With(layout.Maybe(layout.Segment("DB1"))).With(layout.Maybe(layout.Segment("DG1"))).With(layout.Maybe(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Group("CCRI16ClinicalHistoryObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CCRI16RoleClinicalHistory").With(layout.Group("CCRI16RoleClinicalHistoryObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.List(layout.Group("CCRI16PatientVisits").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Group("CCRI16MedicationHistory").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("CCRI16MedicationOrderDetail").With(layout.Segment("RXO")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Group("CCRI16MedicationOrderObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.Group("CCRI16MedicationEncodingDetail").With(layout.Segment("RXE")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Group("CCRI16MedicationEncodingObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.List(layout.Group("CCRI16MedicationAdministrationDetail").With(layout.List(layout.Segment("RXA"))).With(layout.Segment("RXR")).With(layout.Maybe(layout.List(layout.Group("CCRI16MedicationAdministrationObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.Maybe(layout.List(layout.Group("CCRI16Problem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCRI16RoleProblem").With(layout.Group("CCRI16RoleProblemObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CCRI16ProblemObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CCRI16Goal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCRI16RoleGoal").With(layout.Group("CCRI16RoleGoalObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CCRI16GoalObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CCRI16Pathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCRI16RolePathway").With(layout.Group("CCRI16RolePathwayObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CCRI16PathwayObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Segment("REL"))))

func NewCCRI16Layout(msg hl7.Message) (*CCRI16, error) {
	result := layoutCCRI16.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCCRI16(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA33 = layout.Group("ADTA33").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))

func NewADTA33Layout(msg hl7.Message) (*ADTA33, error) {
	result := layoutADTA33.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA33(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPQ34 = layout.Group("QBPQ34").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPQ34Layout(msg hl7.Message) (*QBPQ34, error) {
	result := layoutQBPQ34.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPQ34(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRTBZ74 = layout.Group("RTBZ74").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("RTBZ74RowDefinition").With(layout.Segment("RDF")).With(layout.Maybe(layout.List(layout.Segment("RDT")))))).With(layout.Maybe(layout.Segment("DSC")))

func NewRTBZ74Layout(msg hl7.Message) (*RTBZ74, error) {
	result := layoutRTBZ74.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRTBZ74(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPQ24 = layout.Group("QBPQ24").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPQ24Layout(msg hl7.Message) (*QBPQ24, error) {
	result := layoutQBPQ24.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPQ24(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPPRPC2 = layout.Group("PPRPC2").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Group("PPRPC2PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.List(layout.Group("PPRPC2Problem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPRPC2ProblemRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPRPC2Pathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPRPC2ProblemObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("PPRPC2Goal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPRPC2GoalRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPRPC2GoalObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.List(layout.Group("PPRPC2Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("PPRPC2OrderDetail").With(layout.Group("PPRPC2Choice").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("Hxx")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPRPC2OrderObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))))))))))

func NewPPRPC2Layout(msg hl7.Message) (*PPRPC2, error) {
	result := layoutPPRPC2.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPPRPC2(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFKM13 = layout.Group("MFKM13").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("MFI")).With(layout.Maybe(layout.List(layout.Segment("MFA"))))

func NewMFKM13Layout(msg hl7.Message) (*MFKM13, error) {
	result := layoutMFKM13.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFKM13(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCQUI19 = layout.Group("CQUI19").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("CQUI19ProviderContact").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Maybe(layout.List(layout.Group("CQUI19Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1")))))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Group("CQUI19Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.List(layout.Group("CQUI19AppointmentHistory").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Group("CQUI19Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("CQUI19ResourceDetail").With(layout.Group("CQUI19ResourceObject").With(layout.Maybe(layout.Segment("AIS"))).With(layout.Maybe(layout.Segment("AIG"))).With(layout.Maybe(layout.Segment("AIL"))).With(layout.Maybe(layout.Segment("AIP")))).With(layout.Maybe(layout.List(layout.Group("CQUI19ResourceObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))))))))).With(layout.Maybe(layout.List(layout.Group("CQUI19ClinicalHistory").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("CQUI19ClinicalHistoryDetail").With(layout.Maybe(layout.Group("CQUI19ClinicalHistoryObject").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("ODS"))).With(layout.Maybe(layout.Segment("PR1"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Segment("AL1"))).With(layout.Maybe(layout.Segment("IAM"))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("RMI"))).With(layout.Maybe(layout.Segment("DB1"))).With(layout.Maybe(layout.Segment("DG1"))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.Segment("PDA"))))).With(layout.Maybe(layout.List(layout.Group("CQUI19ClinicalHistoryObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CQUI19RoleClinicalHistory").With(layout.Maybe(layout.Group("CQUI19RoleClinicalHistoryObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD"))))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.List(layout.Group("CQUI19PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Group("CQUI19MedicationHistory").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("CQUI19MedicationOrderDetail").With(layout.Segment("RXO")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Group("CQUI19MedicationOrderObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.Group("CQUI19MedicationEncodingDetail").With(layout.Segment("RXE")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Group("CQUI19MedicationEncodingObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.List(layout.Group("CQUI19MedicationAdministrationDetail").With(layout.List(layout.Segment("RXA"))).With(layout.Segment("RXR")).With(layout.Maybe(layout.List(layout.Group("CQUI19MedicationAdministrationObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.Maybe(layout.List(layout.Group("CQUI19Problem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CQUI19RoleProblem").With(layout.Group("CQUI19RoleProblemObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CQUI19ProblemObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CQUI19Goal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CQUI19RoleGoal").With(layout.Group("CQUI19RoleGoalObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CQUI19GoalObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CQUI19Pathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CQUI19RolePathway").With(layout.Group("CQUI19RolePathwayObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CQUI19PathwayObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Segment("REL"))))

func NewCQUI19Layout(msg hl7.Message) (*CQUI19, error) {
	result := layoutCQUI19.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCQUI19(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRMS07 = layout.Group("SRMS07").With(layout.Segment("MSH")).With(layout.Segment("ARQ")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRMS07Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Group("SRMS07Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRMS07Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRMS07Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS07GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS07LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS07PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSRMS07Layout(msg hl7.Message) (*SRMS07, error) {
	result := layoutSRMS07.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRMS07(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSDRS31 = layout.Group("SDRS31").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Group("SDRS31AntiMicrobialDeviceData").With(layout.Segment("SDD")).With(layout.Maybe(layout.List(layout.Segment("SCD")))))

func NewSDRS31Layout(msg hl7.Message) (*SDRS31, error) {
	result := layoutSDRS31.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSDRS31(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQSBQ16 = layout.Group("QSBQ16").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQSBQ16Layout(msg hl7.Message) (*QSBQ16, error) {
	result := layoutQSBQ16.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQSBQ16(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCRMC08 = layout.Group("CRMC08").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("CRMC08Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("CRMC08PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("CSR")).With(layout.Maybe(layout.List(layout.Segment("CSP"))))))

func NewCRMC08Layout(msg hl7.Message) (*CRMC08, error) {
	result := layoutCRMC08.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCRMC08(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRRS02 = layout.Group("SRRS02").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.Group("SRRS02Schedule").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRRS02Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRRS02Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRRS02Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS02GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS02LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS02PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))))

func NewSRRS02Layout(msg hl7.Message) (*SRRS02, error) {
	result := layoutSRRS02.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRRS02(result.Item.(hl7.Group))
	return &v, nil
}

var layoutLSRU13 = layout.Group("LSRU13").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EQU")).With(layout.List(layout.Segment("EQP")))

func NewLSRU13Layout(msg hl7.Message) (*LSRU13, error) {
	result := layoutLSRU13.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewLSRU13(result.Item.(hl7.Group))
	return &v, nil
}

var layoutBPSO29 = layout.Group("BPSO29").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("BPSO29Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("BPSO29PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))).With(layout.List(layout.Group("BPSO29Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("BPSO29Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("BPO")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("BPSO29Product").With(layout.Segment("BPX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewBPSO29Layout(msg hl7.Message) (*BPSO29, error) {
	result := layoutBPSO29.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewBPSO29(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPK24 = layout.Group("RSPK24").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("RSPK24QueryResponse").With(layout.Segment("PID")))).With(layout.Maybe(layout.Segment("DSC")))

func NewRSPK24Layout(msg hl7.Message) (*RSPK24, error) {
	result := layoutRSPK24.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPK24(result.Item.(hl7.Group))
	return &v, nil
}

var layoutINUU05 = layout.Group("INUU05").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EQU")).With(layout.List(layout.Segment("INV")))

func NewINUU05Layout(msg hl7.Message) (*INUU05, error) {
	result := layoutINUU05.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewINUU05(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFNM07 = layout.Group("MFNM07").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MFI")).With(layout.List(layout.Group("MFNM07MfClinStudySched").With(layout.Segment("MFE")).With(layout.Segment("CM0")).With(layout.Maybe(layout.List(layout.Segment("CM2"))))))

func NewMFNM07Layout(msg hl7.Message) (*MFNM07, error) {
	result := layoutMFNM07.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFNM07(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSSUU03 = layout.Group("SSUU03").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EQU")).With(layout.List(layout.Group("SSUU03SpecimenContainer").With(layout.Segment("SAC")).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SSUU03Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Segment("OBX")))))))))

func NewSSUU03Layout(msg hl7.Message) (*SSUU03, error) {
	result := layoutSSUU03.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSSUU03(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRQAI08 = layout.Group("RQAI08").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("RQAI08Authorization").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("RQAI08Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.Group("RQAI08GuarantorInsurance").With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.List(layout.Group("RQAI08Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Group("RQAI08Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("RQAI08Authorization2").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD")))))))).With(layout.Maybe(layout.List(layout.Group("RQAI08Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("RQAI08Results").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("RQAI08Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRQAI08Layout(msg hl7.Message) (*RQAI08, error) {
	result := layoutRQAI08.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRQAI08(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPPGPCH = layout.Group("PPGPCH").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Group("PPGPCHPatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.List(layout.Group("PPGPCHPathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPGPCHPathwayRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPGPCHGoal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPGPCHGoalRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPGPCHGoalObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("PPGPCHProblem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPGPCHProblemRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPGPCHProblemObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.List(layout.Group("PPGPCHOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("PPGPCHOrderDetail").With(layout.Group("PPGPCHChoice").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("Hxx")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPGPCHOrderObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))))))))))))))))

func NewPPGPCHLayout(msg hl7.Message) (*PPGPCH, error) {
	result := layoutPPGPCH.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPPGPCH(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA26 = layout.Group("ADTA26").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))

func NewADTA26Layout(msg hl7.Message) (*ADTA26, error) {
	result := layoutADTA26.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA26(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPK33 = layout.Group("RSPK33").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.List(layout.Group("RSPK33Donnor").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("ARV"))))))).With(layout.Maybe(layout.Segment("DSC")))

func NewRSPK33Layout(msg hl7.Message) (*RSPK33, error) {
	result := layoutRSPK33.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPK33(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPQ22 = layout.Group("QBPQ22").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPQ22Layout(msg hl7.Message) (*QBPQ22, error) {
	result := layoutQBPQ22.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPQ22(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORUR31 = layout.Group("ORUR31").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Group("ORUR31PatientObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.Group("ORUR31Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORUR31TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.List(layout.Group("ORUR31Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))

func NewORUR31Layout(msg hl7.Message) (*ORUR31, error) {
	result := layoutORUR31.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORUR31(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRRS07 = layout.Group("SRRS07").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.Group("SRRS07Schedule").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRRS07Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRRS07Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRRS07Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS07GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS07LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS07PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))))

func NewSRRS07Layout(msg hl7.Message) (*SRRS07, error) {
	result := layoutSRRS07.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRRS07(result.Item.(hl7.Group))
	return &v, nil
}

var layoutUDMQ05 = layout.Group("UDMQ05").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("URD")).With(layout.Maybe(layout.Segment("URS"))).With(layout.List(layout.Segment("DSP"))).With(layout.Maybe(layout.Segment("DSC")))

func NewUDMQ05Layout(msg hl7.Message) (*UDMQ05, error) {
	result := layoutUDMQ05.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewUDMQ05(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPZ84 = layout.Group("RSPZ84").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("RSPZ84SegmentPattern").With(layout.Segment("Hxx")))).With(layout.Maybe(layout.Segment("DSC")))

func NewRSPZ84Layout(msg hl7.Message) (*RSPZ84, error) {
	result := layoutRSPZ84.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPZ84(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA42 = layout.Group("ADTA42").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.List(layout.Group("ADTA42Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("MRG")).With(layout.Maybe(layout.Segment("PV1")))))

func NewADTA42Layout(msg hl7.Message) (*ADTA42, error) {
	result := layoutADTA42.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA42(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPK21 = layout.Group("RSPK21").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("RSPK21QueryResponse").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Segment("QRI")))).With(layout.Maybe(layout.Segment("DSC")))

func NewRSPK21Layout(msg hl7.Message) (*RSPK21, error) {
	result := layoutRSPK21.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPK21(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSLRS29 = layout.Group("SLRS29").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Segment("SLT")))

func NewSLRS29Layout(msg hl7.Message) (*SLRS29, error) {
	result := layoutSLRS29.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSLRS29(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORUR30 = layout.Group("ORUR30").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Group("ORUR30PatientObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.Group("ORUR30Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORUR30TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.List(layout.Group("ORUR30Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))

func NewORUR30Layout(msg hl7.Message) (*ORUR30, error) {
	result := layoutORUR30.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORUR30(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOPLO37 = layout.Group("OPLO37").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Segment("PRT"))).With(layout.Maybe(layout.Group("OPLO37Guarantor").With(layout.Segment("GT1")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.List(layout.Group("OPLO37Order").With(layout.List(layout.Segment("NK1"))).With(layout.Maybe(layout.Group("OPLO37Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Group("OPLO37ObservationsOnPatient").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Group("OPLO37Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("OPLO37Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("OPLO37SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Group("OPLO37Container").With(layout.Segment("SAC")).With(layout.Maybe(layout.List(layout.Group("OPLO37ContainerObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.List(layout.Group("OPLO37ObservationRequest").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OPLO37Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Segment("TCD"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Group("OPLO37OrderRelatedObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))))).With(layout.Maybe(layout.Group("OPLO37PriorResult").With(layout.List(layout.Segment("NK1"))).With(layout.Maybe(layout.Group("OPLO37PatientPrior").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))))).With(layout.Maybe(layout.Group("OPLO37PatientVisitPrior").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.Segment("AL1"))).With(layout.List(layout.Group("OPLO37OrderPrior").With(layout.Segment("OBR")).With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("OPLO37Timing2").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))).With(layout.List(layout.Group("OPLO37ObservationResultGroup").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Segment("FT1")))).With(layout.Maybe(layout.List(layout.Segment("CTI")))).With(layout.Maybe(layout.Segment("BLG")))))

func NewOPLO37Layout(msg hl7.Message) (*OPLO37, error) {
	result := layoutOPLO37.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOPLO37(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRTBZ78 = layout.Group("RTBZ78").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("RTBZ78RowDefinition").With(layout.Segment("RDF")).With(layout.Maybe(layout.List(layout.Segment("RDT")))))).With(layout.Maybe(layout.Segment("DSC")))

func NewRTBZ78Layout(msg hl7.Message) (*RTBZ78, error) {
	result := layoutRTBZ78.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRTBZ78(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRPAI08 = layout.Group("RPAI08").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("MSA"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("RPAI08Authorization").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("RPAI08Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.List(layout.Group("RPAI08Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3"))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Group("RPAI08Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("RPAI08Authorization2").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD")))))))).With(layout.Maybe(layout.List(layout.Group("RPAI08Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("RPAI08Results").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("RPAI08Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRPAI08Layout(msg hl7.Message) (*RPAI08, error) {
	result := layoutRPAI08.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRPAI08(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPMUB06 = layout.Group("PMUB06").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("STF")).With(layout.Maybe(layout.List(layout.Segment("PRA")))).With(layout.Maybe(layout.List(layout.Segment("ORG"))))

func NewPMUB06Layout(msg hl7.Message) (*PMUB06, error) {
	result := layoutPMUB06.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPMUB06(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRMS02 = layout.Group("SRMS02").With(layout.Segment("MSH")).With(layout.Segment("ARQ")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRMS02Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Group("SRMS02Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRMS02Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRMS02Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS02GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS02LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS02PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSRMS02Layout(msg hl7.Message) (*SRMS02, error) {
	result := layoutSRMS02.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRMS02(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOPUR25 = layout.Group("OPUR25").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("NTE"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OPUR25PatientVisitObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.List(layout.Group("OPUR25AccessionDetail").With(layout.List(layout.Segment("NK1"))).With(layout.Maybe(layout.Group("OPUR25Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Group("OPUR25PatientObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))).With(layout.List(layout.Group("OPUR25Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("OPUR25SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("OPUR25Container").With(layout.Segment("SAC")).With(layout.Maybe(layout.Segment("INV")))))).With(layout.List(layout.Group("OPUR25Order").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("OPUR25CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OPUR25TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.List(layout.Group("OPUR25Result").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))))))

func NewOPUR25Layout(msg hl7.Message) (*OPUR25, error) {
	result := layoutOPUR25.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOPUR25(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPZnn = layout.Group("QBPZnn").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Segment("RDF"))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPZnnLayout(msg hl7.Message) (*QBPZnn, error) {
	result := layoutQBPZnn.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPZnn(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRREO12 = layout.Group("RREO12").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("RREO12Response").With(layout.Maybe(layout.Group("RREO12Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.List(layout.Group("RREO12Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("RREO12Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("RREO12Encoding").With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Group("RREO12TimingEncoded").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC"))))))))))

func NewRREO12Layout(msg hl7.Message) (*RREO12, error) {
	result := layoutRREO12.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRREO12(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA13 = layout.Group("ADTA13").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Group("ADTA13Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("ADTA13Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("AUT")))).With(layout.Maybe(layout.List(layout.Segment("RF1"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("UB1"))).With(layout.Maybe(layout.Segment("UB2"))).With(layout.Maybe(layout.Segment("PDA")))

func NewADTA13Layout(msg hl7.Message) (*ADTA13, error) {
	result := layoutADTA13.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA13(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCCFI22 = layout.Group("CCFI22").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID"))

func NewCCFI22Layout(msg hl7.Message) (*CCFI22, error) {
	result := layoutCCFI22.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCCFI22(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMDMT01 = layout.Group("MDMT01").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Group("MDMT01CommonOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("MDMT01Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Segment("TXA")).With(layout.Maybe(layout.List(layout.Segment("CON"))))

func NewMDMT01Layout(msg hl7.Message) (*MDMT01, error) {
	result := layoutMDMT01.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMDMT01(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRRGO16 = layout.Group("RRGO16").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("RRGO16Response").With(layout.Maybe(layout.Group("RRGO16Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.List(layout.Group("RRGO16Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("RRGO16Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("RRGO16Give").With(layout.Segment("RXG")).With(layout.List(layout.Group("RRGO16TimingGive").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC"))))))))))

func NewRRGO16Layout(msg hl7.Message) (*RRGO16, error) {
	result := layoutRRGO16.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRRGO16(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA55 = layout.Group("ADTA55").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))

func NewADTA55Layout(msg hl7.Message) (*ADTA55, error) {
	result := layoutADTA55.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA55(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQCNJ01 = layout.Group("QCNJ01").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QID"))

func NewQCNJ01Layout(msg hl7.Message) (*QCNJ01, error) {
	result := layoutQCNJ01.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQCNJ01(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRRS06 = layout.Group("SRRS06").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.Group("SRRS06Schedule").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRRS06Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRRS06Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRRS06Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS06GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS06LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS06PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))))

func NewSRRS06Layout(msg hl7.Message) (*SRRS06, error) {
	result := layoutSRRS06.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRRS06(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRQII01 = layout.Group("RQII01").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("RQII01Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.Group("RQII01GuarantorInsurance").With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.List(layout.Group("RQII01Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3"))))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRQII01Layout(msg hl7.Message) (*RQII01, error) {
	result := layoutRQII01.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRQII01(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFKM04 = layout.Group("MFKM04").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("MFI")).With(layout.Maybe(layout.List(layout.Segment("MFA"))))

func NewMFKM04Layout(msg hl7.Message) (*MFKM04, error) {
	result := layoutMFKM04.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFKM04(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSTCS33 = layout.Group("STCS33").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Segment("SCP")))

func NewSTCS33Layout(msg hl7.Message) (*STCS33, error) {
	result := layoutSTCS33.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSTCS33(result.Item.(hl7.Group))
	return &v, nil
}

var layoutBARP01 = layout.Group("BARP01").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.List(layout.Group("BARP01Visit").With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Group("BARP01Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Group("BARP01Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("UB1"))).With(layout.Maybe(layout.Segment("UB2")))))

func NewBARP01Layout(msg hl7.Message) (*BARP01, error) {
	result := layoutBARP01.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewBARP01(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRMS06 = layout.Group("SRMS06").With(layout.Segment("MSH")).With(layout.Segment("ARQ")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRMS06Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Group("SRMS06Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRMS06Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRMS06Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS06GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS06LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS06PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSRMS06Layout(msg hl7.Message) (*SRMS06, error) {
	result := layoutSRMS06.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRMS06(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFNM13 = layout.Group("MFNM13").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MFI")).With(layout.List(layout.Segment("MFE")))

func NewMFNM13Layout(msg hl7.Message) (*MFNM13, error) {
	result := layoutMFNM13.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFNM13(result.Item.(hl7.Group))
	return &v, nil
}

var layoutDPRO48 = layout.Group("DPRO48").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("STF")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Group("DPRO48Donnor").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("DPRO48DonorRegistration").With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))).With(layout.List(layout.Group("DPRO48DonnorOrder").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.Maybe(layout.Group("DPRO48Donation").With(layout.Segment("DON")).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("DPRO48BloodUnit").With(layout.Maybe(layout.List(layout.Segment("BUI")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))

func NewDPRO48Layout(msg hl7.Message) (*DPRO48, error) {
	result := layoutDPRO48.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewDPRO48(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRQII02 = layout.Group("RQII02").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("RQII02Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.Group("RQII02GuarantorInsurance").With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.List(layout.Group("RQII02Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3"))))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRQII02Layout(msg hl7.Message) (*RQII02, error) {
	result := layoutRQII02.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRQII02(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPZ93 = layout.Group("QBPZ93").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("QBPZ93Qbp").With(layout.Maybe(layout.Segment("Hxx"))))).With(layout.Maybe(layout.Segment("RDF"))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPZ93Layout(msg hl7.Message) (*QBPZ93, error) {
	result := layoutQBPZ93.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPZ93(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPE03 = layout.Group("RSPE03").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.List(layout.Segment("UAC")))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.Group("RSPE03QueryAckIpr").With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.List(layout.Segment("IPR"))))))

func NewRSPE03Layout(msg hl7.Message) (*RSPE03, error) {
	result := layoutRSPE03.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPE03(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA24 = layout.Group("ADTA24").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.List(layout.Segment("DB1"))))

func NewADTA24Layout(msg hl7.Message) (*ADTA24, error) {
	result := layoutADTA24.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA24(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFKM12 = layout.Group("MFKM12").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("MFI")).With(layout.Maybe(layout.List(layout.Segment("MFA"))))

func NewMFKM12Layout(msg hl7.Message) (*MFKM12, error) {
	result := layoutMFKM12.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFKM12(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPZ75 = layout.Group("QBPZ75").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("QBPZ75Qbp").With(layout.Maybe(layout.Segment("Hxx"))))).With(layout.Maybe(layout.Segment("RDF"))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPZ75Layout(msg hl7.Message) (*QBPZ75, error) {
	result := layoutQBPZ75.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPZ75(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSIUS18 = layout.Group("SIUS18").With(layout.Segment("MSH")).With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SIUS18Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SIUS18Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SIUS18Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS18GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS18LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS18PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSIUS18Layout(msg hl7.Message) (*SIUS18, error) {
	result := layoutSIUS18.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSIUS18(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPZ99 = layout.Group("QBPZ99").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("QBPZ99Qbp").With(layout.Maybe(layout.Segment("Hxx"))))).With(layout.Maybe(layout.Segment("RDF"))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPZ99Layout(msg hl7.Message) (*QBPZ99, error) {
	result := layoutQBPZ99.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPZ99(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRRII13 = layout.Group("RRII13").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("MSA"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("RRII13AuthorizationContact").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("RRII13ProviderContact").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Group("RRII13Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("RRII13AuthorizationContact").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD")))))))).With(layout.Maybe(layout.List(layout.Group("RRII13Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("RRII13ResultsNotes").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("RRII13PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRRII13Layout(msg hl7.Message) (*RRII13, error) {
	result := layoutRRII13.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRRII13(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCCMI21 = layout.Group("CCMI21").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Group("CCMI21Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.List(layout.Group("CCMI21AppointmentHistory").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Group("CCMI21Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("CCMI21ResourceDetail").With(layout.Group("CCMI21ResourceObject").With(layout.Maybe(layout.Segment("AIS"))).With(layout.Maybe(layout.Segment("AIG"))).With(layout.Maybe(layout.Segment("AIL"))).With(layout.Maybe(layout.Segment("AIP")))).With(layout.Maybe(layout.List(layout.Group("CCMI21ResourceObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))))))))).With(layout.Maybe(layout.List(layout.Group("CCMI21ClinicalHistory").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("CCMI21ClinicalHistoryDetail").With(layout.Group("CCMI21ClinicalHistoryObject").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("ODS"))).With(layout.Maybe(layout.Segment("PR1"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Segment("AL1"))).With(layout.Maybe(layout.Segment("IAM"))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("RMI"))).With(layout.Maybe(layout.Segment("DB1"))).With(layout.Maybe(layout.Segment("DG1"))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.Segment("PDA")))).With(layout.Maybe(layout.List(layout.Group("CCMI21ClinicalHistoryObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CCMI21RoleClinicalHistory").With(layout.Group("CCMI21RoleClinicalHistoryObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.List(layout.Group("CCMI21PatientVisits").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Group("CCMI21MedicationHistory").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("CCMI21MedicationOrderDetail").With(layout.Segment("RXO")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Group("CCMI21MedicationOrderObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.Group("CCMI21MedicationEncodingDetail").With(layout.Segment("RXE")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Group("CCMI21MedicationEncodingObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.List(layout.Group("CCMI21MedicationAdministrationDetail").With(layout.List(layout.Segment("RXA"))).With(layout.Segment("RXR")).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Group("CCMI21MedicationAdministrationObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.Maybe(layout.List(layout.Group("CCMI21Problem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCMI21RoleProblem").With(layout.Group("CCMI21RoleProblemObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CCMI21ProblemObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CCMI21Goal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCMI21RoleGoal").With(layout.Group("CCMI21RoleGoalObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CCMI21GoalObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CCMI21Pathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCMI21RolePathway").With(layout.Group("CCMI21RolePathwayObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CCMI21PathwayObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Segment("REL"))))

func NewCCMI21Layout(msg hl7.Message) (*CCMI21, error) {
	result := layoutCCMI21.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCCMI21(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA10 = layout.Group("ADTA10").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))

func NewADTA10Layout(msg hl7.Message) (*ADTA10, error) {
	result := layoutADTA10.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA10(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA29 = layout.Group("ADTA29").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))

func NewADTA29Layout(msg hl7.Message) (*ADTA29, error) {
	result := layoutADTA29.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA29(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPQ33 = layout.Group("QBPQ33").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPQ33Layout(msg hl7.Message) (*QBPQ33, error) {
	result := layoutQBPQ33.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPQ33(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORAR41 = layout.Group("ORAR41").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))

func NewORAR41Layout(msg hl7.Message) (*ORAR41, error) {
	result := layoutORAR41.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORAR41(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFKM16 = layout.Group("MFKM16").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("MFI")).With(layout.Maybe(layout.List(layout.Segment("MFA"))))

func NewMFKM16Layout(msg hl7.Message) (*MFKM16, error) {
	result := layoutMFKM16.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFKM16(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFNM17 = layout.Group("MFNM17").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MFI")).With(layout.List(layout.Group("MFNM17MfDrg").With(layout.Segment("MFE")).With(layout.Segment("DMI"))))

func NewMFNM17Layout(msg hl7.Message) (*MFNM17, error) {
	result := layoutMFNM17.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFNM17(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRMS08 = layout.Group("SRMS08").With(layout.Segment("MSH")).With(layout.Segment("ARQ")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRMS08Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Group("SRMS08Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRMS08Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRMS08Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS08GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS08LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS08PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSRMS08Layout(msg hl7.Message) (*SRMS08, error) {
	result := layoutSRMS08.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRMS08(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSIUS14 = layout.Group("SIUS14").With(layout.Segment("MSH")).With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SIUS14Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SIUS14Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SIUS14Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS14GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS14LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS14PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSIUS14Layout(msg hl7.Message) (*SIUS14, error) {
	result := layoutSIUS14.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSIUS14(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRRS03 = layout.Group("SRRS03").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.Group("SRRS03Schedule").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRRS03Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRRS03Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRRS03Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS03GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS03LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS03PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))))

func NewSRRS03Layout(msg hl7.Message) (*SRRS03, error) {
	result := layoutSRRS03.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRRS03(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPMUB08 = layout.Group("PMUB08").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("STF")).With(layout.Maybe(layout.Segment("PRA"))).With(layout.Maybe(layout.List(layout.Segment("CER"))))

func NewPMUB08Layout(msg hl7.Message) (*PMUB08, error) {
	result := layoutPMUB08.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPMUB08(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOMLO39 = layout.Group("OMLO39").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMLO39Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("OMLO39PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Group("OMLO39Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("GT1"))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("OMLO39Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OMLO39Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("OMLO39ObservationRequest").With(layout.Segment("OBR")).With(layout.Maybe(layout.Segment("TCD"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Segment("CTD"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Group("OMLO39Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Segment("TCD"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("OMLO39SpecimenShipment").With(layout.Segment("SHP")).With(layout.Maybe(layout.List(layout.Group("OMLO39ShipmentObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.List(layout.Group("OMLO39Package").With(layout.Segment("PAC")).With(layout.Maybe(layout.List(layout.Group("OMLO39SpecimenInPackage").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("OMLO39SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Group("OMLO39SpecimenContainerInPackage").With(layout.Segment("SAC")).With(layout.Maybe(layout.List(layout.Group("OMLO39ContainerObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))))))))))))).With(layout.Maybe(layout.List(layout.Segment("FT1")))).With(layout.Maybe(layout.List(layout.Segment("CTI")))).With(layout.Maybe(layout.Segment("BLG")))))

func NewOMLO39Layout(msg hl7.Message) (*OMLO39, error) {
	result := layoutOMLO39.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOMLO39(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA61 = layout.Group("ADTA61").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.Segment("PV2")))

func NewADTA61Layout(msg hl7.Message) (*ADTA61, error) {
	result := layoutADTA61.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA61(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA49 = layout.Group("ADTA49").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.List(layout.Group("ADTA49Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("MRG"))))

func NewADTA49Layout(msg hl7.Message) (*ADTA49, error) {
	result := layoutADTA49.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA49(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSIUS27 = layout.Group("SIUS27").With(layout.Segment("MSH")).With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SIUS27Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SIUS27Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SIUS27Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS27GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS27LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS27PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSIUS27Layout(msg hl7.Message) (*SIUS27, error) {
	result := layoutSIUS27.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSIUS27(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSIUS24 = layout.Group("SIUS24").With(layout.Segment("MSH")).With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SIUS24Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SIUS24Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SIUS24Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS24GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS24LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS24PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSIUS24Layout(msg hl7.Message) (*SIUS24, error) {
	result := layoutSIUS24.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSIUS24(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCRMC06 = layout.Group("CRMC06").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("CRMC06Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("CRMC06PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("CSR")).With(layout.Maybe(layout.List(layout.Segment("CSP"))))))

func NewCRMC06Layout(msg hl7.Message) (*CRMC06, error) {
	result := layoutCRMC06.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCRMC06(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFKM05 = layout.Group("MFKM05").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("MFI")).With(layout.Maybe(layout.List(layout.Segment("MFA"))))

func NewMFKM05Layout(msg hl7.Message) (*MFKM05, error) {
	result := layoutMFKM05.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFKM05(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFKM02 = layout.Group("MFKM02").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("MFI")).With(layout.Maybe(layout.List(layout.Segment("MFA"))))

func NewMFKM02Layout(msg hl7.Message) (*MFKM02, error) {
	result := layoutMFKM02.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFKM02(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRRS01 = layout.Group("SRRS01").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.Group("SRRS01Schedule").With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRRS01Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRRS01Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRRS01Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS01GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS01LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS01PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))))

func NewSRRS01Layout(msg hl7.Message) (*SRRS01, error) {
	result := layoutSRRS01.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRRS01(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPK23 = layout.Group("RSPK23").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("RSPK23QueryResponse").With(layout.Segment("PID")))).With(layout.Maybe(layout.Segment("DSC")))

func NewRSPK23Layout(msg hl7.Message) (*RSPK23, error) {
	result := layoutRSPK23.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPK23(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFKM15 = layout.Group("MFKM15").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("MFI")).With(layout.Maybe(layout.List(layout.Segment("MFA"))))

func NewMFKM15Layout(msg hl7.Message) (*MFKM15, error) {
	result := layoutMFKM15.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFKM15(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRPAI10 = layout.Group("RPAI10").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("RPAI10Authorization1").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("RPAI10Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("RPAI10Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.List(layout.Group("RPAI10Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("RPAI10Authorization2").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))))).With(layout.Maybe(layout.List(layout.Group("RPAI10Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("RPAI10Results").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("RPAI10Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRPAI10Layout(msg hl7.Message) (*RPAI10, error) {
	result := layoutRPAI10.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRPAI10(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPINI07 = layout.Group("PINI07").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("PINI07Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.Group("PINI07GuarantorInsurance").With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.List(layout.Group("PINI07Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3"))))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewPINI07Layout(msg hl7.Message) (*PINI07, error) {
	result := layoutPINI07.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPINI07(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRDYZ98 = layout.Group("RDYZ98").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.List(layout.Segment("DSP")))).With(layout.Maybe(layout.Segment("DSC")))

func NewRDYZ98Layout(msg hl7.Message) (*RDYZ98, error) {
	result := layoutRDYZ98.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRDYZ98(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA01 = layout.Group("ADTA01").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Group("ADTA01Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("ADTA01Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("AUT")))).With(layout.Maybe(layout.List(layout.Segment("RF1"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("UB1"))).With(layout.Maybe(layout.Segment("UB2"))).With(layout.Maybe(layout.Segment("PDA")))

func NewADTA01Layout(msg hl7.Message) (*ADTA01, error) {
	result := layoutADTA01.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA01(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRDSO13 = layout.Group("RDSO13").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("RDSO13Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Group("RDSO13AdditionalDemographics").With(layout.Segment("PD1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.Group("RDSO13PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))))))).With(layout.List(layout.Group("RDSO13Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("RDSO13Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("RDSO13OrderDetail").With(layout.Segment("RXO")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("RDSO13OrderDetailSupplement").With(layout.List(layout.Segment("NTE"))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Group("RDSO13Component").With(layout.Segment("RXC")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))))).With(layout.Maybe(layout.Group("RDSO13Encoding").With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Group("RDSO13TimingEncoded").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))))).With(layout.Segment("RXD")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Segment("CDO")))).With(layout.Maybe(layout.List(layout.Group("RDSO13Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Segment("FT1"))))))

func NewRDSO13Layout(msg hl7.Message) (*RDSO13, error) {
	result := layoutRDSO13.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRDSO13(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRTBZ96 = layout.Group("RTBZ96").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("RTBZ96RowDefinition").With(layout.Segment("RDF")).With(layout.Maybe(layout.List(layout.Segment("RDT")))))).With(layout.Maybe(layout.Segment("DSC")))

func NewRTBZ96Layout(msg hl7.Message) (*RTBZ96, error) {
	result := layoutRTBZ96.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRTBZ96(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCCRI17 = layout.Group("CCRI17").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Segment("RF1"))).With(layout.List(layout.Group("CCRI17ProviderContact").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Maybe(layout.List(layout.Group("CCRI17ClinicalOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("CCRI17ClinicalOrderTiming").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.List(layout.Group("CCRI17ClinicalOrderDetail").With(layout.Group("CCRI17ClinicalOrderObject").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("RXO"))).With(layout.Maybe(layout.Segment("ODS"))).With(layout.Maybe(layout.Segment("PR1")))).With(layout.Maybe(layout.List(layout.Group("CCRI17ClinicalOrderObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.List(layout.Group("CCRI17Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Group("CCRI17Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.List(layout.Group("CCRI17AppointmentHistory").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Group("CCRI17Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("CCRI17ResourceDetail").With(layout.Group("CCRI17ResourceObject").With(layout.Maybe(layout.Segment("AIS"))).With(layout.Maybe(layout.Segment("AIG"))).With(layout.Maybe(layout.Segment("AIL"))).With(layout.Maybe(layout.Segment("AIP")))).With(layout.Maybe(layout.List(layout.Group("CCRI17ResourceObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))))))))).With(layout.Maybe(layout.List(layout.Group("CCRI17ClinicalHistory").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("CCRI17ClinicalHistoryDetail").With(layout.Group("CCRI17ClinicalHistoryObject").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("ODS"))).With(layout.Maybe(layout.Segment("PR1"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Segment("AL1"))).With(layout.Maybe(layout.Segment("IAM"))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("RMI"))).With(layout.Maybe(layout.Segment("DB1"))).With(layout.Maybe(layout.Segment("DG1"))).With(layout.Maybe(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Group("CCRI17ClinicalHistoryObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CCRI17RoleClinicalHistory").With(layout.Group("CCRI17RoleClinicalHistoryObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.List(layout.Group("CCRI17PatientVisits").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Group("CCRI17MedicationHistory").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("CCRI17MedicationOrderDetail").With(layout.Segment("RXO")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Group("CCRI17MedicationOrderObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.Group("CCRI17MedicationEncodingDetail").With(layout.Segment("RXE")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Group("CCRI17MedicationEncodingObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.List(layout.Group("CCRI17MedicationAdministrationDetail").With(layout.List(layout.Segment("RXA"))).With(layout.Segment("RXR")).With(layout.Maybe(layout.List(layout.Group("CCRI17MedicationAdministrationObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.Maybe(layout.List(layout.Group("CCRI17Problem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCRI17RoleProblem").With(layout.Group("CCRI17RoleProblemObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CCRI17ProblemObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CCRI17Goal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCRI17RoleGoal").With(layout.Group("CCRI17RoleGoalObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CCRI17GoalObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CCRI17Pathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCRI17RolePathway").With(layout.Group("CCRI17RolePathwayObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CCRI17PathwayObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Segment("REL"))))

func NewCCRI17Layout(msg hl7.Message) (*CCRI17, error) {
	result := layoutCCRI17.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCCRI17(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSIUS16 = layout.Group("SIUS16").With(layout.Segment("MSH")).With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SIUS16Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SIUS16Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SIUS16Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS16GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS16LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS16PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSIUS16Layout(msg hl7.Message) (*SIUS16, error) {
	result := layoutSIUS16.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSIUS16(result.Item.(hl7.Group))
	return &v, nil
}

var layoutEHCE20 = layout.Group("EHCE20").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.List(layout.Segment("UAC")))).With(layout.Group("EHCE20AuthorizationRequest").With(layout.Maybe(layout.List(layout.Segment("IVC")))).With(layout.List(layout.Segment("CTD"))).With(layout.Maybe(layout.List(layout.Segment("LOC")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.List(layout.Group("EHCE20PatInfo").With(layout.Maybe(layout.Segment("PID"))).With(layout.Maybe(layout.List(layout.Segment("ACC")))).With(layout.List(layout.Group("EHCE20Insurance").With(layout.Maybe(layout.Segment("IN1"))).With(layout.Maybe(layout.Segment("IN2"))))).With(layout.Maybe(layout.Group("EHCE20Diagnosis").With(layout.Segment("DG1")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.Maybe(layout.List(layout.Segment("OBX")))))).With(layout.List(layout.Group("EHCE20PslItemInfo").With(layout.Segment("PSL")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("ADJ")))).With(layout.Maybe(layout.List(layout.Segment("LOC")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))))))

func NewEHCE20Layout(msg hl7.Message) (*EHCE20, error) {
	result := layoutEHCE20.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewEHCE20(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRRAO18 = layout.Group("RRAO18").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("RRAO18Response").With(layout.Maybe(layout.Group("RRAO18Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.List(layout.Group("RRAO18Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("RRAO18Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("RRAO18Administration").With(layout.List(layout.Group("RRAO18Treatment").With(layout.Segment("RXA")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("RXR"))))))))

func NewRRAO18Layout(msg hl7.Message) (*RRAO18, error) {
	result := layoutRRAO18.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRRAO18(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSLNS34 = layout.Group("SLNS34").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Segment("SLT")))

func NewSLNS34Layout(msg hl7.Message) (*SLNS34, error) {
	result := layoutSLNS34.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSLNS34(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFNM10 = layout.Group("MFNM10").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MFI")).With(layout.List(layout.Group("MFNM10MfTestBatteries").With(layout.Segment("MFE")).With(layout.Segment("OM1")).With(layout.Maybe(layout.Group("MFNM10MfTestBattDetail").With(layout.Segment("OM5")).With(layout.Maybe(layout.List(layout.Segment("OM4"))))))))

func NewMFNM10Layout(msg hl7.Message) (*MFNM10, error) {
	result := layoutMFNM10.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFNM10(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRQII03 = layout.Group("RQII03").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("RQII03Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.Group("RQII03GuarantorInsurance").With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.List(layout.Group("RQII03Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3"))))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRQII03Layout(msg hl7.Message) (*RQII03, error) {
	result := layoutRQII03.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRQII03(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPPPPCB = layout.Group("PPPPCB").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Group("PPPPCBPatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.List(layout.Group("PPPPCBPathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPPPCBPathwayRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPPPCBProblem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPPPCBProblemRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPPPCBProblemObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("PPPPCBGoal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPPPCBGoalRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPPPCBGoalObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.List(layout.Group("PPPPCBOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("PPPPCBOrderDetail").With(layout.Group("PPPPCBChoice").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("Hxx")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPPPCBOrderObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))))))))))))))))

func NewPPPPCBLayout(msg hl7.Message) (*PPPPCB, error) {
	result := layoutPPPPCB.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPPPPCB(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSIUS20 = layout.Group("SIUS20").With(layout.Segment("MSH")).With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SIUS20Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SIUS20Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SIUS20Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS20GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS20LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS20PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSIUS20Layout(msg hl7.Message) (*SIUS20, error) {
	result := layoutSIUS20.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSIUS20(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMDMT04 = layout.Group("MDMT04").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Group("MDMT04CommonOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("MDMT04Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Segment("TXA")).With(layout.Maybe(layout.List(layout.Segment("CON")))).With(layout.List(layout.Group("MDMT04Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))

func NewMDMT04Layout(msg hl7.Message) (*MDMT04, error) {
	result := layoutMDMT04.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMDMT04(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPQ21 = layout.Group("QBPQ21").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPQ21Layout(msg hl7.Message) (*QBPQ21, error) {
	result := layoutQBPQ21.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPQ21(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFNM11 = layout.Group("MFNM11").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MFI")).With(layout.List(layout.Group("MFNM11MfTestCalculated").With(layout.Segment("MFE")).With(layout.Segment("OM1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("MFNM11MfTestCalcDetail").With(layout.Segment("OM6")).With(layout.Segment("OM2"))))))

func NewMFNM11Layout(msg hl7.Message) (*MFNM11, error) {
	result := layoutMFNM11.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFNM11(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRQAI11 = layout.Group("RQAI11").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("RQAI11Authorization").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("RQAI11Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.Group("RQAI11GuarantorInsurance").With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.List(layout.Group("RQAI11Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Group("RQAI11Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("RQAI11Authorization2").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD")))))))).With(layout.Maybe(layout.List(layout.Group("RQAI11Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("RQAI11Results").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("RQAI11Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRQAI11Layout(msg hl7.Message) (*RQAI11, error) {
	result := layoutRQAI11.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRQAI11(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA62 = layout.Group("ADTA62").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.Segment("PV2")))

func NewADTA62Layout(msg hl7.Message) (*ADTA62, error) {
	result := layoutADTA62.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA62(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRMS11 = layout.Group("SRMS11").With(layout.Segment("MSH")).With(layout.Segment("ARQ")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRMS11Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Group("SRMS11Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRMS11Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRMS11Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS11GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS11LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS11PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSRMS11Layout(msg hl7.Message) (*SRMS11, error) {
	result := layoutSRMS11.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRMS11(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPEXP07 = layout.Group("PEXP07").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("PEXP07Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.List(layout.Group("PEXP07Experience").With(layout.Segment("PES")).With(layout.List(layout.Group("PEXP07PexObservation").With(layout.Segment("PEO")).With(layout.List(layout.Group("PEXP07PexCause").With(layout.Segment("PCR")).With(layout.Maybe(layout.Group("PEXP07RxOrder").With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.List(layout.Group("PEXP07TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))).With(layout.Maybe(layout.List(layout.Segment("RXR")))))).With(layout.Maybe(layout.List(layout.Group("PEXP07RxAdministration").With(layout.Segment("RXA")).With(layout.Maybe(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("PRB")))).With(layout.Maybe(layout.List(layout.Group("PEXP07Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("PEXP07AssociatedPerson").With(layout.Segment("NK1")).With(layout.Maybe(layout.Group("PEXP07AssociatedRxOrder").With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.List(layout.Group("PEXP07Nk1TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))).With(layout.Maybe(layout.List(layout.Segment("RXR")))))).With(layout.Maybe(layout.List(layout.Group("PEXP07AssociatedRxAdmin").With(layout.Segment("RXA")).With(layout.Maybe(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("PRB")))).With(layout.Maybe(layout.List(layout.Group("PEXP07AssociatedObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.List(layout.Group("PEXP07Study").With(layout.Segment("CSR")).With(layout.Maybe(layout.List(layout.Segment("CSP")))))))))))))

func NewPEXP07Layout(msg hl7.Message) (*PEXP07, error) {
	result := layoutPEXP07.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPEXP07(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCCUI20 = layout.Group("CCUI20").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("RF1")).With(layout.Maybe(layout.List(layout.Group("CCUI20ProviderContact").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD"))))))).With(layout.Maybe(layout.List(layout.Group("CCUI20Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1")))))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Group("CCUI20Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.List(layout.Group("CCUI20AppointmentHistory").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Group("CCUI20Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("CCUI20ResourceDetail").With(layout.Group("CCUI20ResourceObject").With(layout.Maybe(layout.Segment("AIS"))).With(layout.Maybe(layout.Segment("AIG"))).With(layout.Maybe(layout.Segment("AIL"))).With(layout.Maybe(layout.Segment("AIP")))).With(layout.Maybe(layout.List(layout.Group("CCUI20ResourceObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))))))))).With(layout.Maybe(layout.List(layout.Group("CCUI20ClinicalHistory").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("CCUI20ClinicalHistoryDetail").With(layout.Group("CCUI20ClinicalHistoryObject").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("ODS"))).With(layout.Maybe(layout.Segment("PR1"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Segment("AL1"))).With(layout.Maybe(layout.Segment("IAM"))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("RMI"))).With(layout.Maybe(layout.Segment("DB1"))).With(layout.Maybe(layout.Segment("DG1"))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.Segment("PDA")))).With(layout.Maybe(layout.List(layout.Group("CCUI20ClinicalHistoryObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CCUI20RoleClinicalHistory").With(layout.Group("CCUI20RoleClinicalHistoryObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.List(layout.Group("CCUI20PatientVisits").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Group("CCUI20MedicationHistory").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("CCUI20MedicationOrderDetail").With(layout.Segment("RXO")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Group("CCUI20MedicationOrderObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.Group("CCUI20MedicationEncodingDetail").With(layout.Segment("RXE")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Group("CCUI20MedicationEncodingObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.List(layout.Group("CCUI20MedicationAdministrationDetail").With(layout.List(layout.Segment("RXA"))).With(layout.Segment("RXR")).With(layout.Maybe(layout.List(layout.Group("CCUI20MedicationAdministrationObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))).With(layout.Maybe(layout.List(layout.Group("CCUI20Problem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCUI20RoleProblem").With(layout.Group("CCUI20RoleProblemObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CCUI20ProblemObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CCUI20Goal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCUI20RoleGoal").With(layout.Group("CCUI20RoleGoalObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CCUI20GoalObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Group("CCUI20Pathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("CCUI20RolePathway").With(layout.Group("CCUI20RolePathwayObject").With(layout.Maybe(layout.Segment("ROL"))).With(layout.Maybe(layout.Segment("PRD")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("CCUI20PathwayObservation").With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.List(layout.Segment("REL"))))

func NewCCUI20Layout(msg hl7.Message) (*CCUI20, error) {
	result := layoutCCUI20.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCCUI20(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPPRPC1 = layout.Group("PPRPC1").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Group("PPRPC1PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.List(layout.Group("PPRPC1Problem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPRPC1ProblemRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPRPC1Pathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPRPC1ProblemObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("PPRPC1Goal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPRPC1GoalRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPRPC1GoalObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.List(layout.Group("PPRPC1Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("PPRPC1OrderDetail").With(layout.Group("PPRPC1Choice").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("Hxx")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPRPC1OrderObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))))))))))

func NewPPRPC1Layout(msg hl7.Message) (*PPRPC1, error) {
	result := layoutPPRPC1.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPPRPC1(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRMS04 = layout.Group("SRMS04").With(layout.Segment("MSH")).With(layout.Segment("ARQ")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRMS04Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Group("SRMS04Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRMS04Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRMS04Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS04GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS04LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS04PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSRMS04Layout(msg hl7.Message) (*SRMS04, error) {
	result := layoutSRMS04.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRMS04(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPK25 = layout.Group("RSPK25").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Segment("RCP")).With(layout.List(layout.Group("RSPK25Staff").With(layout.Segment("STF")).With(layout.Maybe(layout.List(layout.Segment("PRA")))).With(layout.Maybe(layout.List(layout.Segment("ORG")))).With(layout.Maybe(layout.List(layout.Segment("AFF")))).With(layout.Maybe(layout.List(layout.Segment("LAN")))).With(layout.Maybe(layout.List(layout.Segment("EDU")))).With(layout.Maybe(layout.List(layout.Segment("CER")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))))).With(layout.Maybe(layout.Segment("DSC")))

func NewRSPK25Layout(msg hl7.Message) (*RSPK25, error) {
	result := layoutRSPK25.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPK25(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSSRU04 = layout.Group("SSRU04").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EQU")).With(layout.List(layout.Group("SSRU04SpecimenContainer").With(layout.Segment("SAC")).With(layout.Maybe(layout.List(layout.Segment("SPM"))))))

func NewSSRU04Layout(msg hl7.Message) (*SSRU04, error) {
	result := layoutSSRU04.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSSRU04(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSIUS26 = layout.Group("SIUS26").With(layout.Segment("MSH")).With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SIUS26Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SIUS26Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SIUS26Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS26GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS26LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS26PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSIUS26Layout(msg hl7.Message) (*SIUS26, error) {
	result := layoutSIUS26.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSIUS26(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORIO24 = layout.Group("ORIO24").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("ORIO24Response").With(layout.Maybe(layout.Group("ORIO24Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.List(layout.Group("ORIO24Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORIO24Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.List(layout.Segment("IPC")))))))

func NewORIO24Layout(msg hl7.Message) (*ORIO24, error) {
	result := layoutORIO24.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORIO24(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPZ86 = layout.Group("RSPZ86").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.List(layout.Group("RSPZ86QueryResponse").With(layout.Maybe(layout.Group("RSPZ86Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("RSPZ86CommonOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("RSPZ86Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("RSPZ86OrderDetail").With(layout.Segment("RXO")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))))).With(layout.Maybe(layout.Group("RSPZ86EncodedOrder").With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Group("RSPZ86TimingEncoded").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))))).With(layout.Maybe(layout.Group("RSPZ86Dispense").With(layout.Segment("RXD")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))))).With(layout.Maybe(layout.Group("RSPZ86Give").With(layout.Segment("RXG")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))))).With(layout.Maybe(layout.Group("RSPZ86Administration").With(layout.Segment("RXA")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))))).With(layout.List(layout.Group("RSPZ86Observation").With(layout.Maybe(layout.Segment("OBX"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Segment("DSC")))

func NewRSPZ86Layout(msg hl7.Message) (*RSPZ86, error) {
	result := layoutRSPZ86.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPZ86(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFKM07 = layout.Group("MFKM07").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("MFI")).With(layout.Maybe(layout.List(layout.Segment("MFA"))))

func NewMFKM07Layout(msg hl7.Message) (*MFKM07, error) {
	result := layoutMFKM07.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFKM07(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRDEO25 = layout.Group("RDEO25").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("RDEO25Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("RDEO25PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))))).With(layout.Maybe(layout.List(layout.Group("RDEO25Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("GT1"))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("RDEO25Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("RDEO25Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("RDEO25OrderDetail").With(layout.Segment("RXO")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Group("RDEO25Component").With(layout.Segment("RXC")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))).With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Group("RDEO25TimingEncoded").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))).With(layout.Maybe(layout.List(layout.Group("RDEO25InfusionOrder").With(layout.Segment("RXV")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Group("RDEO25TimingEncoded").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Segment("CDO")))).With(layout.Maybe(layout.List(layout.Group("RDEO25Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Segment("FT1")))).With(layout.Maybe(layout.Segment("BLG"))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))

func NewRDEO25Layout(msg hl7.Message) (*RDEO25, error) {
	result := layoutRDEO25.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRDEO25(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA23 = layout.Group("ADTA23").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))

func NewADTA23Layout(msg hl7.Message) (*ADTA23, error) {
	result := layoutADTA23.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA23(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSDNS36 = layout.Group("SDNS36").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Group("SDNS36AntiMicrobialDeviceData").With(layout.Segment("SDD")).With(layout.Maybe(layout.List(layout.Segment("SCD")))))

func NewSDNS36Layout(msg hl7.Message) (*SDNS36, error) {
	result := layoutSDNS36.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSDNS36(result.Item.(hl7.Group))
	return &v, nil
}

var layoutDBUO42 = layout.Group("DBUO42").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("STF")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Group("DBUO42Donnor").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("ARV"))))))

func NewDBUO42Layout(msg hl7.Message) (*DBUO42, error) {
	result := layoutDBUO42.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewDBUO42(result.Item.(hl7.Group))
	return &v, nil
}

var layoutDEOO45 = layout.Group("DEOO45").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("STF")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Group("DEOO45Donnor").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("DEOO45DonorRegistration").With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))).With(layout.List(layout.Group("DEOO45DonnorOrder").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("DEOO45DonationObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewDEOO45Layout(msg hl7.Message) (*DEOO45, error) {
	result := layoutDEOO45.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewDEOO45(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMDMT03 = layout.Group("MDMT03").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Group("MDMT03CommonOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("MDMT03Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Segment("TXA")).With(layout.Maybe(layout.List(layout.Segment("CON"))))

func NewMDMT03Layout(msg hl7.Message) (*MDMT03, error) {
	result := layoutMDMT03.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMDMT03(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSIUS13 = layout.Group("SIUS13").With(layout.Segment("MSH")).With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SIUS13Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SIUS13Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SIUS13Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS13GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS13LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS13PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSIUS13Layout(msg hl7.Message) (*SIUS13, error) {
	result := layoutSIUS13.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSIUS13(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPZ89 = layout.Group("QBPZ89").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("QBPZ89Qbp").With(layout.Maybe(layout.Segment("Hxx"))))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPZ89Layout(msg hl7.Message) (*QBPZ89, error) {
	result := layoutQBPZ89.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPZ89(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOSMR26 = layout.Group("OSMR26").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("OSMR26Shipment").With(layout.Segment("SHP")).With(layout.List(layout.Segment("PRT"))).With(layout.Maybe(layout.List(layout.Group("OSMR26ShippingObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.List(layout.Group("OSMR26Package").With(layout.Segment("PAC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OSMR26Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OSMR26SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Group("OSMR26Container").With(layout.Segment("SAC")).With(layout.Maybe(layout.List(layout.Group("OSMR26ContainerObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))).With(layout.Maybe(layout.Group("OSMR26SubjectPersonAnimalIdentification").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Group("OSMR26PatientObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("NK1")))))).With(layout.Maybe(layout.Group("OSMR26SubjectPopulationLocationIdentification").With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OSMR26PatientVisitObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.Segment("PID"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))))))))))))

func NewOSMR26Layout(msg hl7.Message) (*OSMR26, error) {
	result := layoutOSMR26.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOSMR26(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRMS03 = layout.Group("SRMS03").With(layout.Segment("MSH")).With(layout.Segment("ARQ")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRMS03Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Group("SRMS03Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRMS03Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRMS03Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS03GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS03LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS03PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSRMS03Layout(msg hl7.Message) (*SRMS03, error) {
	result := layoutSRMS03.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRMS03(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORPO10 = layout.Group("ORPO10").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("ORPO10Response").With(layout.Maybe(layout.Group("ORPO10Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.List(layout.Group("ORPO10Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORPO10Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("ORPO10OrderDetail").With(layout.Segment("RXO")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Group("ORPO10Component").With(layout.Segment("RXC")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))))))

func NewORPO10Layout(msg hl7.Message) (*ORPO10, error) {
	result := layoutORPO10.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORPO10(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOMPO09 = layout.Group("OMPO09").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMPO09Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Group("OMPO09AdditionalDemographics").With(layout.Segment("PD1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("OMPO09PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))))).With(layout.Maybe(layout.List(layout.Group("OMPO09Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("GT1"))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("OMPO09Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("OMPO09Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("RXO")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Group("OMPO09Component").With(layout.Segment("RXC")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Segment("CDO")))).With(layout.Maybe(layout.List(layout.Group("OMPO09Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Segment("FT1")))).With(layout.Maybe(layout.Segment("BLG")))))

func NewOMPO09Layout(msg hl7.Message) (*OMPO09, error) {
	result := layoutOMPO09.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOMPO09(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPPRPC3 = layout.Group("PPRPC3").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Group("PPRPC3PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.List(layout.Group("PPRPC3Problem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPRPC3ProblemRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPRPC3Pathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPRPC3ProblemObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("PPRPC3Goal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPRPC3GoalRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPRPC3GoalObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.List(layout.Group("PPRPC3Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("PPRPC3OrderDetail").With(layout.Group("PPRPC3Choice").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("Hxx")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPRPC3OrderObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))))))))))

func NewPPRPC3Layout(msg hl7.Message) (*PPRPC3, error) {
	result := layoutPPRPC3.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPPRPC3(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRRS05 = layout.Group("SRRS05").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.Group("SRRS05Schedule").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRRS05Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRRS05Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRRS05Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS05GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS05LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS05PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))))

func NewSRRS05Layout(msg hl7.Message) (*SRRS05, error) {
	result := layoutSRRS05.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRRS05(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA38 = layout.Group("ADTA38").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG")))

func NewADTA38Layout(msg hl7.Message) (*ADTA38, error) {
	result := layoutADTA38.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA38(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPMUB02 = layout.Group("PMUB02").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("STF")).With(layout.Maybe(layout.List(layout.Segment("PRA")))).With(layout.Maybe(layout.List(layout.Segment("ORG")))).With(layout.Maybe(layout.List(layout.Segment("AFF")))).With(layout.Maybe(layout.List(layout.Segment("LAN")))).With(layout.Maybe(layout.List(layout.Segment("EDU")))).With(layout.Maybe(layout.List(layout.Segment("CER")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))

func NewPMUB02Layout(msg hl7.Message) (*PMUB02, error) {
	result := layoutPMUB02.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPMUB02(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPK11 = layout.Group("RSPK11").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("RSPK11SegmentPattern").With(layout.Segment("Hxx")))).With(layout.Maybe(layout.Segment("DSC")))

func NewRSPK11Layout(msg hl7.Message) (*RSPK11, error) {
	result := layoutRSPK11.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPK11(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORUR01 = layout.Group("ORUR01").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("ORUR01PatientResult").With(layout.Maybe(layout.Group("ORUR01Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Group("ORUR01PatientObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.Group("ORUR01Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))).With(layout.List(layout.Group("ORUR01OrderObservation").With(layout.Maybe(layout.Group("ORUR01CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("ORUR01OrderDocument").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("TXA")))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORUR01TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Segment("CTD"))).With(layout.Maybe(layout.List(layout.Group("ORUR01Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Segment("FT1")))).With(layout.Maybe(layout.List(layout.Segment("CTI")))).With(layout.Maybe(layout.List(layout.Group("ORUR01Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("ORUR01SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))))))).With(layout.Maybe(layout.Segment("DSC")))

func NewORUR01Layout(msg hl7.Message) (*ORUR01, error) {
	result := layoutORUR01.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORUR01(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOPRO38 = layout.Group("OPRO38").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OPRO38Response").With(layout.List(layout.Group("OPRO38Order").With(layout.List(layout.Segment("NK1"))).With(layout.Maybe(layout.Segment("PID"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Group("OPRO38Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("OPRO38SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("SAC")))).With(layout.Maybe(layout.List(layout.Group("OPRO38ObservationRequest").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Group("OPRO38Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))))))))))

func NewOPRO38Layout(msg hl7.Message) (*OPRO38, error) {
	result := layoutOPRO38.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOPRO38(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPE22 = layout.Group("QBPE22").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.List(layout.Segment("UAC")))).With(layout.Group("QBPE22Query").With(layout.Segment("QPD")).With(layout.Segment("RCP")))

func NewQBPE22Layout(msg hl7.Message) (*QBPE22, error) {
	result := layoutQBPE22.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPE22(result.Item.(hl7.Group))
	return &v, nil
}

var layoutBARP05 = layout.Group("BARP05").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.List(layout.Group("BARP05Visit").With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Group("BARP05Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Group("BARP05Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("UB1"))).With(layout.Maybe(layout.Segment("UB2"))).With(layout.Maybe(layout.Segment("ABS"))).With(layout.Maybe(layout.List(layout.Segment("BLC")))).With(layout.Maybe(layout.Segment("RMI")))))

func NewBARP05Layout(msg hl7.Message) (*BARP05, error) {
	result := layoutBARP05.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewBARP05(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA47 = layout.Group("ADTA47").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.List(layout.Group("ADTA47Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Segment("MRG"))))

func NewADTA47Layout(msg hl7.Message) (*ADTA47, error) {
	result := layoutADTA47.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA47(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA40 = layout.Group("ADTA40").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.List(layout.Group("ADTA40Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("MRG")).With(layout.Maybe(layout.Segment("PV1")))))

func NewADTA40Layout(msg hl7.Message) (*ADTA40, error) {
	result := layoutADTA40.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA40(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMDMT10 = layout.Group("MDMT10").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Group("MDMT10CommonOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("MDMT10Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Segment("TXA")).With(layout.Maybe(layout.List(layout.Segment("CON")))).With(layout.List(layout.Group("MDMT10Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))

func NewMDMT10Layout(msg hl7.Message) (*MDMT10, error) {
	result := layoutMDMT10.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMDMT10(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRQAI10 = layout.Group("RQAI10").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("RQAI10Authorization").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("RQAI10Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.Group("RQAI10GuarantorInsurance").With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.List(layout.Group("RQAI10Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Group("RQAI10Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("RQAI10Authorization2").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD")))))))).With(layout.Maybe(layout.List(layout.Group("RQAI10Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("RQAI10Results").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("RQAI10Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRQAI10Layout(msg hl7.Message) (*RQAI10, error) {
	result := layoutRQAI10.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRQAI10(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORLO34 = layout.Group("ORLO34").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("ORLO34Response").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Segment("ARV"))).With(layout.List(layout.Group("ORLO34Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("ORLO34SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("SAC")))).With(layout.Maybe(layout.List(layout.Group("ORLO34Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORLO34Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("ORLO34ObservationRequest").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))))))

func NewORLO34Layout(msg hl7.Message) (*ORLO34, error) {
	result := layoutORLO34.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORLO34(result.Item.(hl7.Group))
	return &v, nil
}

var layoutDRGO43 = layout.Group("DRGO43").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("STF")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Group("DRGO43Donnor").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("DRGO43DonorRegistration").With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))

func NewDRGO43Layout(msg hl7.Message) (*DRGO43, error) {
	result := layoutDRGO43.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewDRGO43(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORBO28 = layout.Group("ORBO28").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("ORBO28Response").With(layout.Maybe(layout.Group("ORBO28Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Group("ORBO28Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORBO28Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Segment("BPO"))))))))))

func NewORBO28Layout(msg hl7.Message) (*ORBO28, error) {
	result := layoutORBO28.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORBO28(result.Item.(hl7.Group))
	return &v, nil
}

var layoutEHCE21 = layout.Group("EHCE21").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.List(layout.Segment("UAC")))).With(layout.Group("EHCE21AuthorizationRequest").With(layout.Segment("IVC")).With(layout.List(layout.Group("EHCE21PslItemInfo").With(layout.Segment("PSL")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Segment("AUT"))))))

func NewEHCE21Layout(msg hl7.Message) (*EHCE21, error) {
	result := layoutEHCE21.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewEHCE21(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOMBO27 = layout.Group("OMBO27").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMBO27Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMBO27PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Group("OMBO27Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("GT1"))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("OMBO27Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OMBO27Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("BPO")).With(layout.Maybe(layout.Segment("SPM"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Group("OMBO27Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Segment("FT1")))).With(layout.Maybe(layout.Segment("BLG")))))

func NewOMBO27Layout(msg hl7.Message) (*OMBO27, error) {
	result := layoutOMBO27.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOMBO27(result.Item.(hl7.Group))
	return &v, nil
}

var layoutACK = layout.Group("ACK").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR"))))

func NewACKLayout(msg hl7.Message) (*ACK, error) {
	result := layoutACK.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewACK(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA37 = layout.Group("ADTA37").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.List(layout.Segment("DB1"))))

func NewADTA37Layout(msg hl7.Message) (*ADTA37, error) {
	result := layoutADTA37.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA37(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA16 = layout.Group("ADTA16").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Group("ADTA16Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("ADTA16Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("AUT")))).With(layout.Maybe(layout.List(layout.Segment("RF1"))))))).With(layout.Maybe(layout.Segment("ACC")))

func NewADTA16Layout(msg hl7.Message) (*ADTA16, error) {
	result := layoutADTA16.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA16(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPZ73 = layout.Group("QBPZ73").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Segment("RCP"))

func NewQBPZ73Layout(msg hl7.Message) (*QBPZ73, error) {
	result := layoutQBPZ73.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPZ73(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRRS11 = layout.Group("SRRS11").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.Group("SRRS11Schedule").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRRS11Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRRS11Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRRS11Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS11GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS11LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS11PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))))

func NewSRRS11Layout(msg hl7.Message) (*SRRS11, error) {
	result := layoutSRRS11.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRRS11(result.Item.(hl7.Group))
	return &v, nil
}

var layoutEANU09 = layout.Group("EANU09").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EQU")).With(layout.List(layout.Group("EANU09Notification").With(layout.Segment("NDS")).With(layout.Maybe(layout.Segment("NTE")))))

func NewEANU09Layout(msg hl7.Message) (*EANU09, error) {
	result := layoutEANU09.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewEANU09(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSIUS22 = layout.Group("SIUS22").With(layout.Segment("MSH")).With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SIUS22Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SIUS22Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SIUS22Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS22GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS22LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS22PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSIUS22Layout(msg hl7.Message) (*SIUS22, error) {
	result := layoutSIUS22.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSIUS22(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRPAI11 = layout.Group("RPAI11").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("RPAI11Authorization1").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("RPAI11Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("RPAI11Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.List(layout.Group("RPAI11Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("RPAI11Authorization2").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))))).With(layout.Maybe(layout.List(layout.Group("RPAI11Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("RPAI11Results").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("RPAI11Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRPAI11Layout(msg hl7.Message) (*RPAI11, error) {
	result := layoutRPAI11.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRPAI11(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSLNS35 = layout.Group("SLNS35").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Segment("SLT")))

func NewSLNS35Layout(msg hl7.Message) (*SLNS35, error) {
	result := layoutSLNS35.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSLNS35(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPZ79 = layout.Group("QBPZ79").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Segment("Hxx"))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPZ79Layout(msg hl7.Message) (*QBPZ79, error) {
	result := layoutQBPZ79.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPZ79(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCSUC11 = layout.Group("CSUC11").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("CSUC11Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("CSUC11Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("CSR")).With(layout.List(layout.Group("CSUC11StudyPhase").With(layout.Maybe(layout.Segment("CSP"))).With(layout.List(layout.Group("CSUC11StudySchedule").With(layout.Maybe(layout.Segment("CSS"))).With(layout.List(layout.Group("CSUC11StudyObservation").With(layout.Maybe(layout.Group("CSUC11CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("CSUC11TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.List(layout.Group("CSUC11StudyPharm").With(layout.Maybe(layout.Group("CSUC11CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.List(layout.Group("CSUC11RxAdmin").With(layout.Segment("RXA")).With(layout.Segment("RXR")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))))))))

func NewCSUC11Layout(msg hl7.Message) (*CSUC11, error) {
	result := layoutCSUC11.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCSUC11(result.Item.(hl7.Group))
	return &v, nil
}

var layoutEHCE04 = layout.Group("EHCE04").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.List(layout.Segment("UAC")))).With(layout.Group("EHCE04ReassessmentRequestInfo").With(layout.Maybe(layout.List(layout.Segment("IVC")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("EHCE04ProductServiceSection").With(layout.Segment("PSS")).With(layout.Maybe(layout.List(layout.Group("EHCE04ProductServiceGroup").With(layout.Segment("PSG")).With(layout.Maybe(layout.List(layout.Segment("PSL")))))))))))

func NewEHCE04Layout(msg hl7.Message) (*EHCE04, error) {
	result := layoutEHCE04.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewEHCE04(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRPRI03 = layout.Group("RPRI03").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("MSA"))).With(layout.List(layout.Group("RPRI03Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRPRI03Layout(msg hl7.Message) (*RPRI03, error) {
	result := layoutRPRI03.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRPRI03(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRTBK13 = layout.Group("RTBK13").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("RTBK13RowDefinition").With(layout.Segment("RDF")).With(layout.Maybe(layout.List(layout.Segment("RDT")))))).With(layout.Maybe(layout.Segment("DSC")))

func NewRTBK13Layout(msg hl7.Message) (*RTBK13, error) {
	result := layoutRTBK13.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRTBK13(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA60 = layout.Group("ADTA60").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Group("ADTA60AdverseReactionGroup").With(layout.Segment("IAM")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("IAR")))))))

func NewADTA60Layout(msg hl7.Message) (*ADTA60, error) {
	result := layoutADTA60.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA60(result.Item.(hl7.Group))
	return &v, nil
}

var layoutBRTO32 = layout.Group("BRTO32").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("BRTO32Response").With(layout.Maybe(layout.Segment("PID"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Group("BRTO32Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("BRTO32Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Segment("BPO"))).With(layout.Maybe(layout.List(layout.Segment("BTX")))))))))

func NewBRTO32Layout(msg hl7.Message) (*BRTO32, error) {
	result := layoutBRTO32.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewBRTO32(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFNM05 = layout.Group("MFNM05").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MFI")).With(layout.List(layout.Group("MFNM05MfLocation").With(layout.Segment("MFE")).With(layout.Segment("LOC")).With(layout.Maybe(layout.List(layout.Segment("LCH")))).With(layout.Maybe(layout.List(layout.Segment("LRL")))).With(layout.List(layout.Group("MFNM05MfLocDept").With(layout.Segment("LDP")).With(layout.Maybe(layout.List(layout.Segment("LCH")))).With(layout.Maybe(layout.List(layout.Segment("LCC"))))))))

func NewMFNM05Layout(msg hl7.Message) (*MFNM05, error) {
	result := layoutMFNM05.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFNM05(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPK34 = layout.Group("RSPK34").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("STF")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("RSPK34Donnor").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("RSPK34DonnorRegistration").With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))).With(layout.Maybe(layout.Group("RSPK34Donnation").With(layout.Segment("DON")).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))

func NewRSPK34Layout(msg hl7.Message) (*RSPK34, error) {
	result := layoutRSPK34.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPK34(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPMUB03 = layout.Group("PMUB03").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("STF"))

func NewPMUB03Layout(msg hl7.Message) (*PMUB03, error) {
	result := layoutPMUB03.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPMUB03(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFKM14 = layout.Group("MFKM14").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("MFI")).With(layout.Maybe(layout.List(layout.Segment("MFA"))))

func NewMFKM14Layout(msg hl7.Message) (*MFKM14, error) {
	result := layoutMFKM14.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFKM14(result.Item.(hl7.Group))
	return &v, nil
}

var layoutLSUU12 = layout.Group("LSUU12").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EQU")).With(layout.List(layout.Segment("EQP")))

func NewLSUU12Layout(msg hl7.Message) (*LSUU12, error) {
	result := layoutLSUU12.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewLSUU12(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPZ82 = layout.Group("RSPZ82").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Segment("RCP")).With(layout.List(layout.Group("RSPZ82QueryResponse").With(layout.Maybe(layout.Group("RSPZ82Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("RSPZ82Visit").With(layout.List(layout.Segment("AL1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))))).With(layout.List(layout.Group("RSPZ82CommonOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("RSPZ82Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("RSPZ82OrderDetail").With(layout.Segment("RXO")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.Group("RSPZ82Treatment").With(layout.List(layout.Segment("RXC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))).With(layout.Maybe(layout.Group("RSPZ82EncodedOrder").With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Group("RSPZ82TimingEncoded").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))))).With(layout.Segment("RXD")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.List(layout.Group("RSPZ82Observation").With(layout.Maybe(layout.Segment("OBX"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Segment("DSC")))

func NewRSPZ82Layout(msg hl7.Message) (*RSPZ82, error) {
	result := layoutRSPZ82.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPZ82(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA22 = layout.Group("ADTA22").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))

func NewADTA22Layout(msg hl7.Message) (*ADTA22, error) {
	result := layoutADTA22.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA22(result.Item.(hl7.Group))
	return &v, nil
}

var layoutDBCO41 = layout.Group("DBCO41").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("STF")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Group("DBCO41Donnor").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("AL1"))))))

func NewDBCO41Layout(msg hl7.Message) (*DBCO41, error) {
	result := layoutDBCO41.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewDBCO41(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSLRS28 = layout.Group("SLRS28").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Segment("SLT")))

func NewSLRS28Layout(msg hl7.Message) (*SLRS28, error) {
	result := layoutSLRS28.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSLRS28(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA07 = layout.Group("ADTA07").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.Segment("MRG"))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Group("ADTA07Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("ADTA07Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("UB1"))).With(layout.Maybe(layout.Segment("UB2")))

func NewADTA07Layout(msg hl7.Message) (*ADTA07, error) {
	result := layoutADTA07.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA07(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQVRQ17 = layout.Group("QVRQ17").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("QVRQ17Qbp").With(layout.Maybe(layout.Segment("Hxx"))))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQVRQ17Layout(msg hl7.Message) (*QVRQ17, error) {
	result := layoutQVRQ17.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQVRQ17(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPGLPC7 = layout.Group("PGLPC7").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Group("PGLPC7PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.List(layout.Group("PGLPC7Goal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PGLPC7GoalRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PGLPC7Pathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PGLPC7Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("PGLPC7Problem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PGLPC7ProblemRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PGLPC7ProblemObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.List(layout.Group("PGLPC7Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("PGLPC7OrderDetail").With(layout.Group("PGLPC7Choice").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("Hxx")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PGLPC7OrderObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))))))))))

func NewPGLPC7Layout(msg hl7.Message) (*PGLPC7, error) {
	result := layoutPGLPC7.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPGLPC7(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORDO04 = layout.Group("ORDO04").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("ORDO04Response").With(layout.Maybe(layout.Group("ORDO04Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.List(layout.Group("ORDO04OrderDiet").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORDO04TimingDiet").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.List(layout.Segment("ODS")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.Maybe(layout.List(layout.Group("ORDO04OrderTray").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORDO04TimingTray").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.List(layout.Segment("ODT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewORDO04Layout(msg hl7.Message) (*ORDO04, error) {
	result := layoutORDO04.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORDO04(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPZ91 = layout.Group("QBPZ91").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("QBPZ91Qbp").With(layout.Maybe(layout.Segment("Hxx"))))).With(layout.Maybe(layout.Segment("RDF"))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPZ91Layout(msg hl7.Message) (*QBPZ91, error) {
	result := layoutQBPZ91.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPZ91(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRRS09 = layout.Group("SRRS09").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.Group("SRRS09Schedule").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRRS09Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRRS09Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRRS09Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS09GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS09LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS09PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))))

func NewSRRS09Layout(msg hl7.Message) (*SRRS09, error) {
	result := layoutSRRS09.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRRS09(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPZ97 = layout.Group("QBPZ97").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Segment("Hxx"))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPZ97Layout(msg hl7.Message) (*QBPZ97, error) {
	result := layoutQBPZ97.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPZ97(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA04 = layout.Group("ADTA04").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Group("ADTA04Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("ADTA04Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("AUT")))).With(layout.Maybe(layout.List(layout.Segment("RF1"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("UB1"))).With(layout.Maybe(layout.Segment("UB2"))).With(layout.Maybe(layout.Segment("PDA")))

func NewADTA04Layout(msg hl7.Message) (*ADTA04, error) {
	result := layoutADTA04.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA04(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA52 = layout.Group("ADTA52").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2")))

func NewADTA52Layout(msg hl7.Message) (*ADTA52, error) {
	result := layoutADTA52.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA52(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPZ85 = layout.Group("QBPZ85").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("QBPZ85Qbp").With(layout.Maybe(layout.Segment("Hxx"))))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPZ85Layout(msg hl7.Message) (*QBPZ85, error) {
	result := layoutQBPZ85.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPZ85(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOMSO05 = layout.Group("OMSO05").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMSO05Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMSO05PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Group("OMSO05Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("GT1"))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("OMSO05Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.Segment("PRT"))).With(layout.Maybe(layout.List(layout.Group("OMSO05Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("RQD")).With(layout.Maybe(layout.Segment("RQ1"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("OMSO05Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.Segment("BLG")))))

func NewOMSO05Layout(msg hl7.Message) (*OMSO05, error) {
	result := layoutOMSO05.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOMSO05(result.Item.(hl7.Group))
	return &v, nil
}

var layoutESUU01 = layout.Group("ESUU01").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EQU")).With(layout.Maybe(layout.List(layout.Segment("ISD"))))

func NewESUU01Layout(msg hl7.Message) (*ESUU01, error) {
	result := layoutESUU01.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewESUU01(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRTBZ94 = layout.Group("RTBZ94").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("RTBZ94RowDefinition").With(layout.Segment("RDF")).With(layout.Maybe(layout.List(layout.Segment("RDT")))))).With(layout.Maybe(layout.Segment("DSC")))

func NewRTBZ94Layout(msg hl7.Message) (*RTBZ94, error) {
	result := layoutRTBZ94.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRTBZ94(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPK22 = layout.Group("RSPK22").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.List(layout.Group("RSPK22QueryResponse").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.Segment("QRI")))))).With(layout.Maybe(layout.Segment("DSC")))

func NewRSPK22Layout(msg hl7.Message) (*RSPK22, error) {
	result := layoutRSPK22.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPK22(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRQAI09 = layout.Group("RQAI09").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("RQAI09Authorization").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("RQAI09Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.Group("RQAI09GuarantorInsurance").With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.List(layout.Group("RQAI09Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Group("RQAI09Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("RQAI09Authorization2").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD")))))))).With(layout.Maybe(layout.List(layout.Group("RQAI09Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("RQAI09Results").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("RQAI09Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRQAI09Layout(msg hl7.Message) (*RQAI09, error) {
	result := layoutRQAI09.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRQAI09(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFNM04 = layout.Group("MFNM04").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MFI")).With(layout.List(layout.Group("MFNM04MfCdm").With(layout.Segment("MFE")).With(layout.Segment("CDM")).With(layout.Maybe(layout.List(layout.Segment("PRC"))))))

func NewMFNM04Layout(msg hl7.Message) (*MFNM04, error) {
	result := layoutMFNM04.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFNM04(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQSXJ02 = layout.Group("QSXJ02").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QID"))

func NewQSXJ02Layout(msg hl7.Message) (*QSXJ02, error) {
	result := layoutQSXJ02.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQSXJ02(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCSUC12 = layout.Group("CSUC12").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("CSUC12Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("CSUC12Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("CSR")).With(layout.List(layout.Group("CSUC12StudyPhase").With(layout.Maybe(layout.Segment("CSP"))).With(layout.List(layout.Group("CSUC12StudySchedule").With(layout.Maybe(layout.Segment("CSS"))).With(layout.List(layout.Group("CSUC12StudyObservation").With(layout.Maybe(layout.Group("CSUC12CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("CSUC12TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.List(layout.Group("CSUC12StudyPharm").With(layout.Maybe(layout.Group("CSUC12CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.List(layout.Group("CSUC12RxAdmin").With(layout.Segment("RXA")).With(layout.Segment("RXR")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))))))))

func NewCSUC12Layout(msg hl7.Message) (*CSUC12, error) {
	result := layoutCSUC12.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCSUC12(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOMQO42 = layout.Group("OMQO42").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMQO42Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("OMQO42PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Group("OMQO42Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("GT1"))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("OMQO42Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("TXA")).With(layout.Maybe(layout.Segment("CTD"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Group("OMQO42Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("OMQO42Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("OMQO42SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Group("OMQO42Container").With(layout.Segment("SAC")).With(layout.Maybe(layout.List(layout.Group("OMQO42ContainerObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))))))).With(layout.Maybe(layout.List(layout.Group("OMQO42PriorResult").With(layout.Maybe(layout.Group("OMQO42PatientPrior").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))))).With(layout.Maybe(layout.Group("OMQO42PatientVisitPrior").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.List(layout.Group("OMQO42OrderPrior").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Group("OMQO42TimingPrior").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Segment("CTD"))).With(layout.List(layout.Group("OMQO42ObservationPrior").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))))).With(layout.Maybe(layout.List(layout.Segment("FT1")))).With(layout.Maybe(layout.List(layout.Segment("CTI")))).With(layout.Maybe(layout.Segment("BLG")))))

func NewOMQO42Layout(msg hl7.Message) (*OMQO42, error) {
	result := layoutOMQO42.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOMQO42(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA54 = layout.Group("ADTA54").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))

func NewADTA54Layout(msg hl7.Message) (*ADTA54, error) {
	result := layoutADTA54.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA54(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRPAI09 = layout.Group("RPAI09").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("RPAI09Authorization1").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("RPAI09Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("RPAI09Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.List(layout.Group("RPAI09Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("RPAI09Authorization2").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))))).With(layout.Maybe(layout.List(layout.Group("RPAI09Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("RPAI09Results").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("RPAI09Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRPAI09Layout(msg hl7.Message) (*RPAI09, error) {
	result := layoutRPAI09.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRPAI09(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA12 = layout.Group("ADTA12").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))

func NewADTA12Layout(msg hl7.Message) (*ADTA12, error) {
	result := layoutADTA12.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA12(result.Item.(hl7.Group))
	return &v, nil
}

var layoutVXUV04 = layout.Group("VXUV04").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("VXUV04PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("VXUV04Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.List(layout.Group("VXUV04PersonObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("VXUV04Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("VXUV04Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("RXA")).With(layout.Maybe(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Group("VXUV04Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))))

func NewVXUV04Layout(msg hl7.Message) (*VXUV04, error) {
	result := layoutVXUV04.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewVXUV04(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA14 = layout.Group("ADTA14").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Group("ADTA14Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("ADTA14Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("AUT")))).With(layout.Maybe(layout.List(layout.Segment("RF1"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("UB1"))).With(layout.Maybe(layout.Segment("UB2")))

func NewADTA14Layout(msg hl7.Message) (*ADTA14, error) {
	result := layoutADTA14.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA14(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFNM06 = layout.Group("MFNM06").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MFI")).With(layout.List(layout.Group("MFNM06MfClinStudy").With(layout.Segment("MFE")).With(layout.Segment("CM0")).With(layout.Maybe(layout.List(layout.Group("MFNM06MfPhaseSchedDetail").With(layout.Segment("CM1")).With(layout.Maybe(layout.List(layout.Segment("CM2")))))))))

func NewMFNM06Layout(msg hl7.Message) (*MFNM06, error) {
	result := layoutMFNM06.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFNM06(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFNM02 = layout.Group("MFNM02").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MFI")).With(layout.List(layout.Group("MFNM02MfStaff").With(layout.Segment("MFE")).With(layout.Segment("STF")).With(layout.Maybe(layout.List(layout.Segment("PRA")))).With(layout.Maybe(layout.List(layout.Segment("ORG")))).With(layout.Maybe(layout.List(layout.Segment("AFF")))).With(layout.Maybe(layout.List(layout.Segment("LAN")))).With(layout.Maybe(layout.List(layout.Segment("EDU")))).With(layout.Maybe(layout.List(layout.Segment("CER")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))

func NewMFNM02Layout(msg hl7.Message) (*MFNM02, error) {
	result := layoutMFNM02.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFNM02(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPZ77 = layout.Group("QBPZ77").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("QBPZ77Qbp").With(layout.Maybe(layout.Segment("Hxx"))))).With(layout.Maybe(layout.Segment("RDF"))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPZ77Layout(msg hl7.Message) (*QBPZ77, error) {
	result := layoutQBPZ77.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPZ77(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSIUS12 = layout.Group("SIUS12").With(layout.Segment("MSH")).With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SIUS12Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SIUS12Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SIUS12Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS12GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS12LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS12PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSIUS12Layout(msg hl7.Message) (*SIUS12, error) {
	result := layoutSIUS12.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSIUS12(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA15 = layout.Group("ADTA15").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))

func NewADTA15Layout(msg hl7.Message) (*ADTA15, error) {
	result := layoutADTA15.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA15(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPQ11 = layout.Group("QBPQ11").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("QBPQ11Qbp").With(layout.Maybe(layout.Segment("Hxx"))))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPQ11Layout(msg hl7.Message) (*QBPQ11, error) {
	result := layoutQBPQ11.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPQ11(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCRMC02 = layout.Group("CRMC02").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("CRMC02Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("CRMC02PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("CSR")).With(layout.Maybe(layout.List(layout.Segment("CSP"))))))

func NewCRMC02Layout(msg hl7.Message) (*CRMC02, error) {
	result := layoutCRMC02.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCRMC02(result.Item.(hl7.Group))
	return &v, nil
}

var layoutREFI14 = layout.Group("REFI14").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("REFI14AuthorizationContact").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("REFI14ProviderContact").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("REFI14Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Group("REFI14Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("REFI14AuthorizationContact2").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD")))))))).With(layout.Maybe(layout.List(layout.Group("REFI14Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("REFI14ResultsNotes").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("REFI14PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewREFI14Layout(msg hl7.Message) (*REFI14, error) {
	result := layoutREFI14.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewREFI14(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOMGO19 = layout.Group("OMGO19").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMGO19Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("OMGO19PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Group("OMGO19Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("GT1"))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("OMGO19Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OMGO19Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Segment("CTD"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Group("OMGO19Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("OMGO19Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("OMGO19SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Group("OMGO19Container").With(layout.Segment("SAC")).With(layout.Maybe(layout.List(layout.Group("OMGO19ContainerObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))))))).With(layout.Maybe(layout.List(layout.Group("OMGO19PriorResult").With(layout.Maybe(layout.Group("OMGO19PatientPrior").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.Group("OMGO19PatientVisitPrior").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.List(layout.Group("OMGO19OrderPrior").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Group("OMGO19TimingPrior").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Segment("CTD"))).With(layout.List(layout.Group("OMGO19ObservationPrior").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))))).With(layout.Maybe(layout.List(layout.Segment("FT1")))).With(layout.Maybe(layout.List(layout.Segment("CTI")))).With(layout.Maybe(layout.Segment("BLG")))))

func NewOMGO19Layout(msg hl7.Message) (*OMGO19, error) {
	result := layoutOMGO19.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOMGO19(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSIUS19 = layout.Group("SIUS19").With(layout.Segment("MSH")).With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SIUS19Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SIUS19Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SIUS19Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS19GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS19LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS19PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSIUS19Layout(msg hl7.Message) (*SIUS19, error) {
	result := layoutSIUS19.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSIUS19(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRDRRDR = layout.Group("RDRRDR").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.Segment("SFT"))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("RDRRDRDefinition").With(layout.Segment("QRD")).With(layout.Maybe(layout.Segment("QRF"))).With(layout.Maybe(layout.Group("RDRRDRPatient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.List(layout.Group("RDRRDROrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("RDRRDRTiming").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("RDRRDREncoding").With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Group("RDRRDRTimingEncoded").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))))).With(layout.List(layout.Group("RDRRDRDispense").With(layout.Segment("RXD")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))))))))).With(layout.Maybe(layout.Segment("DSC")))

func NewRDRRDRLayout(msg hl7.Message) (*RDRRDR, error) {
	result := layoutRDRRDR.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRDRRDR(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPMUB01 = layout.Group("PMUB01").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("STF")).With(layout.Maybe(layout.List(layout.Segment("PRA")))).With(layout.Maybe(layout.List(layout.Segment("ORG")))).With(layout.Maybe(layout.List(layout.Segment("AFF")))).With(layout.Maybe(layout.List(layout.Segment("LAN")))).With(layout.Maybe(layout.List(layout.Segment("EDU")))).With(layout.Maybe(layout.List(layout.Segment("CER")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))

func NewPMUB01Layout(msg hl7.Message) (*PMUB01, error) {
	result := layoutPMUB01.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPMUB01(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRPLI02 = layout.Group("RPLI02").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.List(layout.Group("RPLI02Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("DSP")))).With(layout.Maybe(layout.Segment("DSC")))

func NewRPLI02Layout(msg hl7.Message) (*RPLI02, error) {
	result := layoutRPLI02.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRPLI02(result.Item.(hl7.Group))
	return &v, nil
}

var layoutEHCE24 = layout.Group("EHCE24").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.List(layout.Segment("UAC")))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Group("EHCE24AuthorizationResponseInfo").With(layout.Segment("IVC")).With(layout.List(layout.Group("EHCE24PslItemInfo").With(layout.Segment("PSL")).With(layout.Maybe(layout.Segment("AUT"))).With(layout.Maybe(layout.List(layout.Segment("ADJ")))))))

func NewEHCE24Layout(msg hl7.Message) (*EHCE24, error) {
	result := layoutEHCE24.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewEHCE24(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA32 = layout.Group("ADTA32").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))

func NewADTA32Layout(msg hl7.Message) (*ADTA32, error) {
	result := layoutADTA32.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA32(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOULR23 = layout.Group("OULR23").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("NTE"))).With(layout.Maybe(layout.Group("OULR23Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("OULR23PatientObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.Group("OULR23Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.List(layout.Group("OULR23Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("OULR23SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.List(layout.Group("OULR23Container").With(layout.Segment("SAC")).With(layout.Maybe(layout.Segment("INV"))).With(layout.List(layout.Group("OULR23Order").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("OULR23CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("OULR23OrderDocument").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("TXA")))))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OULR23TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.List(layout.Group("OULR23Result").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Segment("TCD"))).With(layout.Maybe(layout.List(layout.Segment("SID")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Segment("CTI")))))))))).With(layout.Maybe(layout.Segment("DSC")))

func NewOULR23Layout(msg hl7.Message) (*OULR23, error) {
	result := layoutOULR23.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOULR23(result.Item.(hl7.Group))
	return &v, nil
}

var layoutBARP12 = layout.Group("BARP12").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Group("BARP12Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.Segment("OBX")))

func NewBARP12Layout(msg hl7.Message) (*BARP12, error) {
	result := layoutBARP12.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewBARP12(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRRII12 = layout.Group("RRII12").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("MSA"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("RRII12AuthorizationContact").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("RRII12ProviderContact").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Group("RRII12Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("RRII12AuthorizationContact2").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD")))))))).With(layout.Maybe(layout.List(layout.Group("RRII12Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("RRII12ResultsNotes").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("RRII12PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRRII12Layout(msg hl7.Message) (*RRII12, error) {
	result := layoutRRII12.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRRII12(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORUR32 = layout.Group("ORUR32").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Group("ORUR32PatientObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.Group("ORUR32Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORUR32TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.List(layout.Group("ORUR32Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))

func NewORUR32Layout(msg hl7.Message) (*ORUR32, error) {
	result := layoutORUR32.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORUR32(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOMNO07 = layout.Group("OMNO07").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMNO07Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMNO07PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Group("OMNO07Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("GT1"))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("OMNO07Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OMNO07Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("RQD")).With(layout.Maybe(layout.Segment("RQ1"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("OMNO07Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.Segment("BLG")))))

func NewOMNO07Layout(msg hl7.Message) (*OMNO07, error) {
	result := layoutOMNO07.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOMNO07(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPQ13 = layout.Group("QBPQ13").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("QBPQ13Qbp").With(layout.Maybe(layout.Segment("Hxx"))))).With(layout.Maybe(layout.Segment("RDF"))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPQ13Layout(msg hl7.Message) (*QBPQ13, error) {
	result := layoutQBPQ13.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPQ13(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCRMC04 = layout.Group("CRMC04").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("CRMC04Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("CRMC04PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("CSR")).With(layout.Maybe(layout.List(layout.Segment("CSP"))))))

func NewCRMC04Layout(msg hl7.Message) (*CRMC04, error) {
	result := layoutCRMC04.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCRMC04(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCRMC01 = layout.Group("CRMC01").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("CRMC01Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("CRMC01PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("CSR")).With(layout.Maybe(layout.List(layout.Segment("CSP"))))))

func NewCRMC01Layout(msg hl7.Message) (*CRMC01, error) {
	result := layoutCRMC01.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCRMC01(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA09 = layout.Group("ADTA09").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))

func NewADTA09Layout(msg hl7.Message) (*ADTA09, error) {
	result := layoutADTA09.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA09(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCSUC10 = layout.Group("CSUC10").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("CSUC10Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("CSUC10Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("CSR")).With(layout.List(layout.Group("CSUC10StudyPhase").With(layout.Maybe(layout.Segment("CSP"))).With(layout.List(layout.Group("CSUC10StudySchedule").With(layout.Maybe(layout.Segment("CSS"))).With(layout.List(layout.Group("CSUC10StudyObservation").With(layout.Maybe(layout.Group("CSUC10CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("CSUC10TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.List(layout.Group("CSUC10StudyPharm").With(layout.Maybe(layout.Group("CSUC10CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.List(layout.Group("CSUC10RxAdmin").With(layout.Segment("RXA")).With(layout.Segment("RXR")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))))))))

func NewCSUC10Layout(msg hl7.Message) (*CSUC10, error) {
	result := layoutCSUC10.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCSUC10(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORLO36 = layout.Group("ORLO36").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("ORLO36Response").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.List(layout.Group("ORLO36Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("ORLO36SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Group("ORLO36SpecimenContainer").With(layout.Segment("SAC")).With(layout.Maybe(layout.List(layout.Group("ORLO36Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORLO36Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("ORLO36ObservationRequest").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))))))))))

func NewORLO36Layout(msg hl7.Message) (*ORLO36, error) {
	result := layoutORLO36.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORLO36(result.Item.(hl7.Group))
	return &v, nil
}

var layoutNMDN02 = layout.Group("NMDN02").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("NMDN02ClockAndStatsWithNotes").With(layout.Maybe(layout.Group("NMDN02Clock").With(layout.Segment("NCK")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.Maybe(layout.Group("NMDN02AppStats").With(layout.Segment("NST")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.Maybe(layout.Group("NMDN02AppStatus").With(layout.Segment("NSC")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))

func NewNMDN02Layout(msg hl7.Message) (*NMDN02, error) {
	result := layoutNMDN02.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewNMDN02(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRRS10 = layout.Group("SRRS10").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.Group("SRRS10Schedule").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRRS10Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRRS10Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRRS10Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS10GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS10LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS10PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))))

func NewSRRS10Layout(msg hl7.Message) (*SRRS10, error) {
	result := layoutSRRS10.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRRS10(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFKM10 = layout.Group("MFKM10").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("MFI")).With(layout.Maybe(layout.List(layout.Segment("MFA"))))

func NewMFKM10Layout(msg hl7.Message) (*MFKM10, error) {
	result := layoutMFKM10.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFKM10(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPMUB05 = layout.Group("PMUB05").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("STF")).With(layout.Maybe(layout.List(layout.Segment("PRA")))).With(layout.Maybe(layout.List(layout.Segment("ORG"))))

func NewPMUB05Layout(msg hl7.Message) (*PMUB05, error) {
	result := layoutPMUB05.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPMUB05(result.Item.(hl7.Group))
	return &v, nil
}

var layoutBARP02 = layout.Group("BARP02").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.List(layout.Group("BARP02Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.List(layout.Segment("DB1"))))))

func NewBARP02Layout(msg hl7.Message) (*BARP02, error) {
	result := layoutBARP02.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewBARP02(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRMS10 = layout.Group("SRMS10").With(layout.Segment("MSH")).With(layout.Segment("ARQ")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRMS10Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Group("SRMS10Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRMS10Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRMS10Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS10GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS10LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS10PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSRMS10Layout(msg hl7.Message) (*SRMS10, error) {
	result := layoutSRMS10.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRMS10(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFKM06 = layout.Group("MFKM06").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("MFI")).With(layout.Maybe(layout.List(layout.Segment("MFA"))))

func NewMFKM06Layout(msg hl7.Message) (*MFKM06, error) {
	result := layoutMFKM06.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFKM06(result.Item.(hl7.Group))
	return &v, nil
}

var layoutREFI15 = layout.Group("REFI15").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("REFI15AuthorizationContact").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("REFI15ProviderContact").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("REFI15Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Group("REFI15Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("REFI15AuthorizationContact2").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD")))))))).With(layout.Maybe(layout.List(layout.Group("REFI15Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("REFI15ResultsNotes").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("REFI15PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewREFI15Layout(msg hl7.Message) (*REFI15, error) {
	result := layoutREFI15.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewREFI15(result.Item.(hl7.Group))
	return &v, nil
}

var layoutBARP10 = layout.Group("BARP10").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Segment("GP1")).With(layout.Maybe(layout.List(layout.Group("BARP10Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Segment("GP2"))))))

func NewBARP10Layout(msg hl7.Message) (*BARP10, error) {
	result := layoutBARP10.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewBARP10(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRDYZ80 = layout.Group("RDYZ80").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("RDYZ80SegmentPattern").With(layout.Segment("Hxx")))).With(layout.Maybe(layout.Segment("DSC")))

func NewRDYZ80Layout(msg hl7.Message) (*RDYZ80, error) {
	result := layoutRDYZ80.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRDYZ80(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOULR24 = layout.Group("OULR24").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("NTE"))).With(layout.Maybe(layout.Group("OULR24Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("OULR24PatientObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.Group("OULR24Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.List(layout.Group("OULR24Order").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("OULR24CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("OULR24OrderDocument").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("TXA")))))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OULR24TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.List(layout.Group("OULR24Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("OULR24SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Group("OULR24Container").With(layout.Segment("SAC")).With(layout.Maybe(layout.Segment("INV"))))))))).With(layout.Maybe(layout.List(layout.Group("OULR24Result").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Segment("TCD"))).With(layout.Maybe(layout.List(layout.Segment("SID")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Segment("CTI")))).With(layout.Maybe(layout.Segment("DSC")))))

func NewOULR24Layout(msg hl7.Message) (*OULR24, error) {
	result := layoutOULR24.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOULR24(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA25 = layout.Group("ADTA25").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))

func NewADTA25Layout(msg hl7.Message) (*ADTA25, error) {
	result := layoutADTA25.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA25(result.Item.(hl7.Group))
	return &v, nil
}

var layoutREFI13 = layout.Group("REFI13").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("REFI13AuthorizationContact").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("REFI13ProviderContact").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("REFI13Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Group("REFI13Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("REFI13AuthorizationContact2").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD")))))))).With(layout.Maybe(layout.List(layout.Group("REFI13Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("REFI13ResultsNotes").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("REFI13PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewREFI13Layout(msg hl7.Message) (*REFI13, error) {
	result := layoutREFI13.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewREFI13(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA45 = layout.Group("ADTA45").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.List(layout.Group("ADTA45MergeInfo").With(layout.Segment("MRG")).With(layout.Segment("PV1"))))

func NewADTA45Layout(msg hl7.Message) (*ADTA45, error) {
	result := layoutADTA45.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA45(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRMS01 = layout.Group("SRMS01").With(layout.Segment("MSH")).With(layout.Segment("ARQ")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRMS01Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Group("SRMS01Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRMS01Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRMS01Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS01GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS01LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS01PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSRMS01Layout(msg hl7.Message) (*SRMS01, error) {
	result := layoutSRMS01.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRMS01(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRRS08 = layout.Group("SRRS08").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.Group("SRRS08Schedule").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRRS08Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRRS08Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRRS08Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS08GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS08LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS08PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))))

func NewSRRS08Layout(msg hl7.Message) (*SRRS08, error) {
	result := layoutSRRS08.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRRS08(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPPGPCJ = layout.Group("PPGPCJ").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Group("PPGPCJPatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.List(layout.Group("PPGPCJPathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPGPCJPathwayRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPGPCJGoal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPGPCJGoalRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPGPCJGoalObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("PPGPCJProblem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPGPCJProblemRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPGPCJProblemObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.List(layout.Group("PPGPCJOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("PPGPCJOrderDetail").With(layout.Group("PPGPCJChoice").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("Hxx")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPGPCJOrderObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))))))))))))))))

func NewPPGPCJLayout(msg hl7.Message) (*PPGPCJ, error) {
	result := layoutPPGPCJ.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPPGPCJ(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFNM14 = layout.Group("MFNM14").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MFI")).With(layout.List(layout.Group("MFNM14MfSiteDefined").With(layout.Segment("MFE")).With(layout.Segment("Hxx"))))

func NewMFNM14Layout(msg hl7.Message) (*MFNM14, error) {
	result := layoutMFNM14.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFNM14(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSMDS32 = layout.Group("SMDS32").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Group("SMDS32AntiMicrobialDeviceCycleData").With(layout.Segment("SDD")).With(layout.Maybe(layout.List(layout.Segment("SCD")))))

func NewSMDS32Layout(msg hl7.Message) (*SMDS32, error) {
	result := layoutSMDS32.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSMDS32(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCRMC05 = layout.Group("CRMC05").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("CRMC05Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("CRMC05PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("CSR")).With(layout.Maybe(layout.List(layout.Segment("CSP"))))))

func NewCRMC05Layout(msg hl7.Message) (*CRMC05, error) {
	result := layoutCRMC05.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCRMC05(result.Item.(hl7.Group))
	return &v, nil
}

var layoutTCUU10 = layout.Group("TCUU10").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EQU")).With(layout.List(layout.Group("TCUU10TestConfiguration").With(layout.Maybe(layout.Segment("SPM"))).With(layout.List(layout.Segment("TCC")))))

func NewTCUU10Layout(msg hl7.Message) (*TCUU10, error) {
	result := layoutTCUU10.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewTCUU10(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRREO26 = layout.Group("RREO26").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("RREO26Response").With(layout.Maybe(layout.Group("RREO26Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.List(layout.Group("RREO26Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("RREO26Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("RREO26Encoding").With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Group("RREO26TimingEncoded").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))

func NewRREO26Layout(msg hl7.Message) (*RREO26, error) {
	result := layoutRREO26.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRREO26(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA41 = layout.Group("ADTA41").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.List(layout.Group("ADTA41Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("MRG")).With(layout.Maybe(layout.Segment("PV1")))))

func NewADTA41Layout(msg hl7.Message) (*ADTA41, error) {
	result := layoutADTA41.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA41(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA08 = layout.Group("ADTA08").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Group("ADTA08Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("ADTA08Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("AUT")))).With(layout.Maybe(layout.List(layout.Segment("RF1"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("UB1"))).With(layout.Maybe(layout.Segment("UB2"))).With(layout.Maybe(layout.Segment("PDA")))

func NewADTA08Layout(msg hl7.Message) (*ADTA08, error) {
	result := layoutADTA08.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA08(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORSO06 = layout.Group("ORSO06").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("ORSO06Response").With(layout.Maybe(layout.Group("ORSO06Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.List(layout.Segment("ARV"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.List(layout.Group("ORSO06Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORSO06Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("RQD")).With(layout.Maybe(layout.Segment("RQ1"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))

func NewORSO06Layout(msg hl7.Message) (*ORSO06, error) {
	result := layoutORSO06.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORSO06(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPQ32 = layout.Group("QBPQ32").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPQ32Layout(msg hl7.Message) (*QBPQ32, error) {
	result := layoutQBPQ32.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPQ32(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORLO40 = layout.Group("ORLO40").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("ORLO40Response").With(layout.Maybe(layout.Group("ORLO40Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Group("ORLO40Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORLO40Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("ORLO40ObservationRequest").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORLO40SpecimenShipment").With(layout.Segment("SHP")).With(layout.List(layout.Group("ORLO40Package").With(layout.Segment("PAC")).With(layout.Maybe(layout.List(layout.Group("ORLO40SpecimenInPackage").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("ORLO40SpecimenContainerInPackage").With(layout.Segment("SAC"))))))))))))))))))))))

func NewORLO40Layout(msg hl7.Message) (*ORLO40, error) {
	result := layoutORLO40.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORLO40(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRPII04 = layout.Group("RPII04").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.List(layout.Group("RPII04Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.Group("RPII04GuarantorInsurance").With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.List(layout.Group("RPII04Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3"))))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRPII04Layout(msg hl7.Message) (*RPII04, error) {
	result := layoutRPII04.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRPII04(result.Item.(hl7.Group))
	return &v, nil
}

var layoutEHCE13 = layout.Group("EHCE13").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.List(layout.Segment("UAC")))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("RFI")).With(layout.Maybe(layout.List(layout.Segment("CTD")))).With(layout.Segment("IVC")).With(layout.Segment("PSS")).With(layout.Segment("PSG")).With(layout.Maybe(layout.Segment("PID"))).With(layout.Maybe(layout.Segment("PSL"))).With(layout.List(layout.Group("EHCE13Request").With(layout.Maybe(layout.Segment("CTD"))).With(layout.Segment("OBR")).With(layout.Maybe(layout.Segment("NTE"))).With(layout.List(layout.Group("EHCE13Response").With(layout.Segment("OBX")).With(layout.Maybe(layout.Segment("NTE"))).With(layout.Maybe(layout.Segment("TXA")))))))

func NewEHCE13Layout(msg hl7.Message) (*EHCE13, error) {
	result := layoutEHCE13.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewEHCE13(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOMLO33 = layout.Group("OMLO33").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMLO33Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("OMLO33PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Group("OMLO33Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("GT1"))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("OMLO33Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("OMLO33SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("SAC")))).With(layout.List(layout.Group("OMLO33Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OMLO33Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("OMLO33ObservationRequest").With(layout.Segment("OBR")).With(layout.Maybe(layout.Segment("TCD"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Group("OMLO33Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Segment("TCD"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("OMLO33PriorResult").With(layout.Maybe(layout.Group("OMLO33PatientPrior").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))))).With(layout.Maybe(layout.Group("OMLO33PatientVisitPrior").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.List(layout.Group("OMLO33OrderPrior").With(layout.Segment("ORC")).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OMLO33TimingPrior").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.List(layout.Group("OMLO33ObservationPrior").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))))))).With(layout.Maybe(layout.List(layout.Segment("FT1")))).With(layout.Maybe(layout.List(layout.Segment("CTI")))).With(layout.Maybe(layout.Segment("BLG")))))))

func NewOMLO33Layout(msg hl7.Message) (*OMLO33, error) {
	result := layoutOMLO33.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOMLO33(result.Item.(hl7.Group))
	return &v, nil
}

var layoutDFTP11 = layout.Group("DFTP11").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Group("DFTP11CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Group("DFTP11TimingQuantity").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("DFTP11Order").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.Maybe(layout.List(layout.Group("DFTP11Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("DFTP11Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.List(layout.Group("DFTP11Financial").With(layout.Segment("FT1")).With(layout.Maybe(layout.List(layout.Group("DFTP11FinancialProcedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Group("DFTP11FinancialCommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Group("DFTP11FinancialTimingQuantity").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("DFTP11FinancialOrder").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.Maybe(layout.List(layout.Group("DFTP11FinancialObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("DFTP11FinancialInsurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))))))))

func NewDFTP11Layout(msg hl7.Message) (*DFTP11, error) {
	result := layoutDFTP11.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewDFTP11(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCRMC03 = layout.Group("CRMC03").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("CRMC03Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("CRMC03PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("CSR")).With(layout.Maybe(layout.List(layout.Segment("CSP"))))))

func NewCRMC03Layout(msg hl7.Message) (*CRMC03, error) {
	result := layoutCRMC03.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCRMC03(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRDYK15 = layout.Group("RDYK15").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.List(layout.Segment("DSP")))).With(layout.Maybe(layout.Segment("DSC")))

func NewRDYK15Layout(msg hl7.Message) (*RDYK15, error) {
	result := layoutRDYK15.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRDYK15(result.Item.(hl7.Group))
	return &v, nil
}

var layoutEARU08 = layout.Group("EARU08").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EQU")).With(layout.List(layout.Group("EARU08CommandResponse").With(layout.Segment("ECD")).With(layout.Maybe(layout.Group("EARU08SpecimenContainer").With(layout.Segment("SAC")).With(layout.Maybe(layout.List(layout.Segment("SPM")))))).With(layout.Segment("ECR"))))

func NewEARU08Layout(msg hl7.Message) (*EARU08, error) {
	result := layoutEARU08.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewEARU08(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPZ88 = layout.Group("RSPZ88").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Segment("RCP")).With(layout.List(layout.Group("RSPZ88QueryResponse").With(layout.Maybe(layout.Group("RSPZ88Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("RSPZ88Allergy").With(layout.List(layout.Segment("AL1"))).With(layout.Maybe(layout.Group("RSPZ88Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))))))).With(layout.List(layout.Group("RSPZ88CommonOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("RSPZ88Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("RSPZ88OrderDetail").With(layout.Segment("RXO")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.Group("RSPZ88Component").With(layout.List(layout.Segment("RXC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))).With(layout.Maybe(layout.Group("RSPZ88OrderEncoded").With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Group("RSPZ88TimingEncoded").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))))).With(layout.Segment("RXD")).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.List(layout.Group("RSPZ88Observation").With(layout.Maybe(layout.Segment("OBX"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Segment("DSC"))

func NewRSPZ88Layout(msg hl7.Message) (*RSPZ88, error) {
	result := layoutRSPZ88.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPZ88(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORXO43 = layout.Group("ORXO43").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("ORXO43Response").With(layout.Maybe(layout.Group("ORXO43Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))))).With(layout.List(layout.Group("ORXO43Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("TXA")).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))))

func NewORXO43Layout(msg hl7.Message) (*ORXO43, error) {
	result := layoutORXO43.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORXO43(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPGLPC6 = layout.Group("PGLPC6").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Group("PGLPC6PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.List(layout.Group("PGLPC6Goal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PGLPC6GoalRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PGLPC6Pathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PGLPC6Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("PGLPC6Problem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PGLPC6ProblemRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PGLPC6ProblemObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.List(layout.Group("PGLPC6Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("PGLPC6OrderDetail").With(layout.Group("PGLPC6Choice").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("Hxx")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PGLPC6OrderObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))))))))))

func NewPGLPC6Layout(msg hl7.Message) (*PGLPC6, error) {
	result := layoutPGLPC6.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPGLPC6(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMDMT05 = layout.Group("MDMT05").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Group("MDMT05CommonOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("MDMT05Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Segment("TXA")).With(layout.Maybe(layout.List(layout.Segment("CON"))))

func NewMDMT05Layout(msg hl7.Message) (*MDMT05, error) {
	result := layoutMDMT05.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMDMT05(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPK32 = layout.Group("RSPK32").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.List(layout.Group("RSPK32QueryResponse").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.Segment("QRI")))))).With(layout.Maybe(layout.Segment("DSC")))

func NewRSPK32Layout(msg hl7.Message) (*RSPK32, error) {
	result := layoutRSPK32.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPK32(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMDMT02 = layout.Group("MDMT02").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Group("MDMT02CommonOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("MDMT02Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Segment("TXA")).With(layout.Maybe(layout.List(layout.Segment("CON")))).With(layout.List(layout.Group("MDMT02Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))

func NewMDMT02Layout(msg hl7.Message) (*MDMT02, error) {
	result := layoutMDMT02.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMDMT02(result.Item.(hl7.Group))
	return &v, nil
}

var layoutEHCE02 = layout.Group("EHCE02").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.List(layout.Segment("UAC")))).With(layout.Group("EHCE02InvoiceInformationCancel").With(layout.Segment("IVC")).With(layout.Segment("PYE")).With(layout.Maybe(layout.List(layout.Segment("CTD")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("EHCE02ProductServiceSection").With(layout.Segment("PSS")).With(layout.Maybe(layout.List(layout.Group("EHCE02Psg").With(layout.Segment("PSG")).With(layout.Maybe(layout.List(layout.Segment("PSL")))))))))))

func NewEHCE02Layout(msg hl7.Message) (*EHCE02, error) {
	result := layoutEHCE02.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewEHCE02(result.Item.(hl7.Group))
	return &v, nil
}

var layoutREFI12 = layout.Group("REFI12").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("REFI12AuthorizationContact").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("REFI12ProviderContact").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("REFI12Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Group("REFI12Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("REFI12AuthorizationContact2").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD")))))))).With(layout.Maybe(layout.List(layout.Group("REFI12Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("REFI12ResultsNotes").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("REFI12PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewREFI12Layout(msg hl7.Message) (*REFI12, error) {
	result := layoutREFI12.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewREFI12(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRRS04 = layout.Group("SRRS04").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.Group("SRRS04Schedule").With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRRS04Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRRS04Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRRS04Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS04GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS04LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRRS04PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))))

func NewSRRS04Layout(msg hl7.Message) (*SRRS04, error) {
	result := layoutSRRS04.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRRS04(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCCQI19 = layout.Group("CCQI19").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("RF1")).With(layout.Maybe(layout.List(layout.Group("CCQI19ProviderContact").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD"))))))).With(layout.Maybe(layout.List(layout.Segment("REL"))))

func NewCCQI19Layout(msg hl7.Message) (*CCQI19, error) {
	result := layoutCCQI19.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCCQI19(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPQ23 = layout.Group("QBPQ23").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPQ23Layout(msg hl7.Message) (*QBPQ23, error) {
	result := layoutQBPQ23.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPQ23(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSCNS37 = layout.Group("SCNS37").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Group("SCNS37AntiMicrobialDeviceCycleData").With(layout.Segment("SDD")).With(layout.Maybe(layout.List(layout.Segment("SCD")))))

func NewSCNS37Layout(msg hl7.Message) (*SCNS37, error) {
	result := layoutSCNS37.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSCNS37(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFNM09 = layout.Group("MFNM09").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MFI")).With(layout.List(layout.Group("MFNM09MfTestCategorical").With(layout.Segment("MFE")).With(layout.Segment("OM1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("MFNM09MfTestCatDetail").With(layout.Segment("OM3")).With(layout.Maybe(layout.List(layout.Segment("OM4"))))))))

func NewMFNM09Layout(msg hl7.Message) (*MFNM09, error) {
	result := layoutMFNM09.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFNM09(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA31 = layout.Group("ADTA31").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Group("ADTA31Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("ADTA31Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("AUT")))).With(layout.Maybe(layout.List(layout.Segment("RF1"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("UB1"))).With(layout.Maybe(layout.Segment("UB2")))

func NewADTA31Layout(msg hl7.Message) (*ADTA31, error) {
	result := layoutADTA31.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA31(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPPGPCG = layout.Group("PPGPCG").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Group("PPGPCGPatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.List(layout.Group("PPGPCGPathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPGPCGPathwayRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPGPCGGoal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPGPCGGoalRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPGPCGGoalObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("PPGPCGProblem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPGPCGProblemRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPGPCGProblemObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.List(layout.Group("PPGPCGOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("PPGPCGOrderDetail").With(layout.Group("PPGPCGChoice").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("Hxx")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPGPCGOrderObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))))))))))))))))

func NewPPGPCGLayout(msg hl7.Message) (*PPGPCG, error) {
	result := layoutPPGPCG.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPPGPCG(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORNO08 = layout.Group("ORNO08").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("ORNO08Response").With(layout.Maybe(layout.Group("ORNO08Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.List(layout.Group("ORNO08Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORNO08Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("RQD")).With(layout.Maybe(layout.Segment("RQ1"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))

func NewORNO08Layout(msg hl7.Message) (*ORNO08, error) {
	result := layoutORNO08.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORNO08(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRRDO14 = layout.Group("RRDO14").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("RRDO14Response").With(layout.Maybe(layout.Group("RRDO14Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))).With(layout.List(layout.Group("RRDO14Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("RRDO14Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("RRDO14Dispense").With(layout.Segment("RXD")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC"))))))))))

func NewRRDO14Layout(msg hl7.Message) (*RRDO14, error) {
	result := layoutRRDO14.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRRDO14(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOMLO35 = layout.Group("OMLO35").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMLO35Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("OMLO35PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Group("OMLO35Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("GT1"))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("OMLO35Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("OMLO35SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.List(layout.Group("OMLO35SpecimenContainer").With(layout.Segment("SAC")).With(layout.List(layout.Group("OMLO35Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OMLO35Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("OMLO35ObservationRequest").With(layout.Segment("OBR")).With(layout.Maybe(layout.Segment("TCD"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Group("OMLO35Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Segment("TCD"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("OMLO35PriorResult").With(layout.Maybe(layout.Group("OMLO35PatientPrior").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))))).With(layout.Maybe(layout.Group("OMLO35PatientVisitPrior").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.List(layout.Group("OMLO35OrderPrior").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OMLO35TimingPrior").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.List(layout.Group("OMLO35ObservationPrior").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))))))).With(layout.Maybe(layout.List(layout.Segment("FT1")))).With(layout.Maybe(layout.List(layout.Segment("CTI")))).With(layout.Maybe(layout.Segment("BLG")))))))))

func NewOMLO35Layout(msg hl7.Message) (*OMLO35, error) {
	result := layoutOMLO35.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOMLO35(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRRII15 = layout.Group("RRII15").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("MSA"))).With(layout.Maybe(layout.Segment("RF1"))).With(layout.Maybe(layout.Group("RRII15AuthorizationContact").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD"))))).With(layout.List(layout.Group("RRII15ProviderContact").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.List(layout.Segment("DRG")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Group("RRII15Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.Group("RRII15AuthorizationContact").With(layout.Segment("AUT")).With(layout.Maybe(layout.Segment("CTD")))))))).With(layout.Maybe(layout.List(layout.Group("RRII15Observation").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("RRII15ResultsNotes").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.Group("RRII15PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRRII15Layout(msg hl7.Message) (*RRII15, error) {
	result := layoutRRII15.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRRII15(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPE03 = layout.Group("QBPE03").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.List(layout.Segment("UAC")))).With(layout.Group("QBPE03QueryInformation").With(layout.Segment("QPD")).With(layout.Segment("RCP")))

func NewQBPE03Layout(msg hl7.Message) (*QBPE03, error) {
	result := layoutQBPE03.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPE03(result.Item.(hl7.Group))
	return &v, nil
}

var layoutBTSO31 = layout.Group("BTSO31").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("BTSO31Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("BTSO31PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))).With(layout.List(layout.Group("BTSO31Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("BTSO31Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("BPO")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("BTSO31ProductStatus").With(layout.Segment("BTX")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewBTSO31Layout(msg hl7.Message) (*BTSO31, error) {
	result := layoutBTSO31.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewBTSO31(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSRMS05 = layout.Group("SRMS05").With(layout.Segment("MSH")).With(layout.Segment("ARQ")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SRMS05Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Group("SRMS05Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SRMS05Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SRMS05Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS05GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS05LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SRMS05PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.Segment("APR"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSRMS05Layout(msg hl7.Message) (*SRMS05, error) {
	result := layoutSRMS05.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSRMS05(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPK31 = layout.Group("RSPK31").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Segment("RCP")).With(layout.List(layout.Group("RSPK31Response").With(layout.Maybe(layout.Group("RSPK31Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Group("RSPK31AdditionalDemographics").With(layout.Segment("PD1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.Group("RSPK31PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))).With(layout.List(layout.Group("RSPK31Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("RSPK31Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("RSPK31OrderDetail").With(layout.Segment("RXO")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Group("RSPK31Components").With(layout.Segment("RXC")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))).With(layout.Maybe(layout.Group("RSPK31Encoding").With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Group("RSPK31TimingEncoded").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))))).With(layout.Segment("RXD")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Group("RSPK31Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))))).With(layout.Maybe(layout.Segment("DSC")))

func NewRSPK31Layout(msg hl7.Message) (*RSPK31, error) {
	result := layoutRSPK31.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPK31(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRTBZ76 = layout.Group("RTBZ76").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("RTBZ76RowDefinition").With(layout.Segment("RDF")).With(layout.Maybe(layout.List(layout.Segment("RDT")))))).With(layout.Maybe(layout.Segment("DSC")))

func NewRTBZ76Layout(msg hl7.Message) (*RTBZ76, error) {
	result := layoutRTBZ76.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRTBZ76(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFNM08 = layout.Group("MFNM08").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MFI")).With(layout.List(layout.Group("MFNM08MfTestNumeric").With(layout.Segment("MFE")).With(layout.Segment("OM1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Segment("OM2"))).With(layout.Maybe(layout.Segment("OM3"))).With(layout.Maybe(layout.List(layout.Segment("OM4"))))))

func NewMFNM08Layout(msg hl7.Message) (*MFNM08, error) {
	result := layoutMFNM08.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFNM08(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA02 = layout.Group("ADTA02").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.Segment("PDA")))

func NewADTA02Layout(msg hl7.Message) (*ADTA02, error) {
	result := layoutADTA02.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA02(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRDEO11 = layout.Group("RDEO11").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("RDEO11Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("RDEO11PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))))).With(layout.Maybe(layout.List(layout.Group("RDEO11Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("GT1"))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("RDEO11Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("RDEO11Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("RDEO11OrderDetail").With(layout.Segment("RXO")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Group("RDEO11Component").With(layout.Segment("RXC")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))).With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Group("RDEO11TimingEncoded").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))).With(layout.Maybe(layout.Group("RDEO11InfusionOrder").With(layout.Segment("RXV")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.List(layout.Group("RDEO11TimingEncoded").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))))).With(layout.List(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("RXC")))).With(layout.Maybe(layout.List(layout.Segment("CDO")))).With(layout.Maybe(layout.List(layout.Group("RDEO11Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Segment("FT1")))).With(layout.Maybe(layout.Segment("BLG"))).With(layout.Maybe(layout.List(layout.Segment("CTI"))))))

func NewRDEO11Layout(msg hl7.Message) (*RDEO11, error) {
	result := layoutRDEO11.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRDEO11(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA17 = layout.Group("ADTA17").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))

func NewADTA17Layout(msg hl7.Message) (*ADTA17, error) {
	result := layoutADTA17.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA17(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOULR22 = layout.Group("OULR22").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Segment("NTE"))).With(layout.Maybe(layout.Group("OULR22Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("OULR22PatientObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.Group("OULR22Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.List(layout.Group("OULR22Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Group("OULR22SpecimenObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Group("OULR22Container").With(layout.Segment("SAC")).With(layout.Maybe(layout.Segment("INV")))))).With(layout.List(layout.Group("OULR22Order").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("OULR22CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("OULR22OrderDocument").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Segment("TXA")))))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OULR22TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.List(layout.Group("OULR22Result").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Segment("TCD"))).With(layout.Maybe(layout.List(layout.Segment("SID")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Segment("CTI")))))))).With(layout.Maybe(layout.Segment("DSC")))

func NewOULR22Layout(msg hl7.Message) (*OULR22, error) {
	result := layoutOULR22.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOULR22(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMDMT11 = layout.Group("MDMT11").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Group("MDMT11CommonOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("MDMT11Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Segment("TXA")).With(layout.Maybe(layout.List(layout.Segment("CON"))))

func NewMDMT11Layout(msg hl7.Message) (*MDMT11, error) {
	result := layoutMDMT11.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMDMT11(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPPPPCC = layout.Group("PPPPCC").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Group("PPPPCCPatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.List(layout.Group("PPPPCCPathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPPPCCPathwayRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPPPCCProblem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPPPCCProblemRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPPPCCProblemObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("PPPPCCGoal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPPPCCGoalRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PPPPCCGoalObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.List(layout.Group("PPPPCCOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("PPPPCCOrderDetail").With(layout.Group("PPPPCCChoice").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("Hxx")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PPPPCCOrderObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))))))))))))))))

func NewPPPPCCLayout(msg hl7.Message) (*PPPPCC, error) {
	result := layoutPPPPCC.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPPPPCC(result.Item.(hl7.Group))
	return &v, nil
}

var layoutEHCE12 = layout.Group("EHCE12").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.List(layout.Segment("UAC")))).With(layout.Segment("RFI")).With(layout.Maybe(layout.List(layout.Segment("CTD")))).With(layout.Segment("IVC")).With(layout.Segment("PSS")).With(layout.Segment("PSG")).With(layout.Maybe(layout.Segment("PID"))).With(layout.Maybe(layout.List(layout.Segment("PSL")))).With(layout.List(layout.Group("EHCE12Request").With(layout.Maybe(layout.Segment("CTD"))).With(layout.Segment("OBR")).With(layout.Maybe(layout.Segment("NTE"))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))))

func NewEHCE12Layout(msg hl7.Message) (*EHCE12, error) {
	result := layoutEHCE12.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewEHCE12(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMDMT07 = layout.Group("MDMT07").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Group("MDMT07CommonOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("MDMT07Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Segment("TXA")).With(layout.Maybe(layout.List(layout.Segment("CON"))))

func NewMDMT07Layout(msg hl7.Message) (*MDMT07, error) {
	result := layoutMDMT07.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMDMT07(result.Item.(hl7.Group))
	return &v, nil
}

var layoutESRU02 = layout.Group("ESRU02").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EQU"))

func NewESRU02Layout(msg hl7.Message) (*ESRU02, error) {
	result := layoutESRU02.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewESRU02(result.Item.(hl7.Group))
	return &v, nil
}

var layoutBARP06 = layout.Group("BARP06").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.List(layout.Group("BARP06Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1")))))

func NewBARP06Layout(msg hl7.Message) (*BARP06, error) {
	result := layoutBARP06.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewBARP06(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA53 = layout.Group("ADTA53").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2")))

func NewADTA53Layout(msg hl7.Message) (*ADTA53, error) {
	result := layoutADTA53.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA53(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPZ81 = layout.Group("QBPZ81").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("QBPZ81Qbp").With(layout.Maybe(layout.Segment("Hxx"))))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPZ81Layout(msg hl7.Message) (*QBPZ81, error) {
	result := layoutQBPZ81.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPZ81(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA27 = layout.Group("ADTA27").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))

func NewADTA27Layout(msg hl7.Message) (*ADTA27, error) {
	result := layoutADTA27.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA27(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA21 = layout.Group("ADTA21").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))

func NewADTA21Layout(msg hl7.Message) (*ADTA21, error) {
	result := layoutADTA21.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA21(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRQPI04 = layout.Group("RQPI04").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("RQPI04Provider").With(layout.Segment("PRD")).With(layout.Maybe(layout.List(layout.Segment("CTD")))))).With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))

func NewRQPI04Layout(msg hl7.Message) (*RQPI04, error) {
	result := layoutRQPI04.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRQPI04(result.Item.(hl7.Group))
	return &v, nil
}

var layoutTCRU11 = layout.Group("TCRU11").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EQU")).With(layout.List(layout.Group("TCRU11TestConfiguration").With(layout.Maybe(layout.Segment("SPM"))).With(layout.List(layout.Segment("TCC")))))

func NewTCRU11Layout(msg hl7.Message) (*TCRU11, error) {
	result := layoutTCRU11.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewTCRU11(result.Item.(hl7.Group))
	return &v, nil
}

var layoutEHCE15 = layout.Group("EHCE15").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.List(layout.Segment("UAC")))).With(layout.Group("EHCE15PaymentRemittanceHeaderInfo").With(layout.Maybe(layout.List(layout.Segment("PMT")))).With(layout.Segment("PYE"))).With(layout.Maybe(layout.List(layout.Group("EHCE15PaymentRemittanceDetailInfo").With(layout.Segment("IPR")).With(layout.Segment("IVC")).With(layout.List(layout.Group("EHCE15ProductServiceSection").With(layout.Segment("PSS")).With(layout.List(layout.Group("EHCE15ProductServiceGroup").With(layout.Segment("PSG")).With(layout.List(layout.Group("EHCE15PslItemInfo").With(layout.Segment("PSL")).With(layout.Maybe(layout.List(layout.Segment("ADJ"))))))))))))).With(layout.Maybe(layout.List(layout.Group("EHCE15AdjustmentPayee").With(layout.Segment("ADJ")).With(layout.Maybe(layout.Segment("PRT"))).With(layout.Maybe(layout.Segment("ROL"))))))

func NewEHCE15Layout(msg hl7.Message) (*EHCE15, error) {
	result := layoutEHCE15.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewEHCE15(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRSPE22 = layout.Group("RSPE22").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.List(layout.Segment("UAC")))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Group("RSPE22QueryAck").With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("RSPE22AuthorizationInfo").With(layout.Segment("IVC")).With(layout.Segment("PSG")).With(layout.List(layout.Group("RSPE22PslItemInfo").With(layout.Segment("PSL")))))))

func NewRSPE22Layout(msg hl7.Message) (*RSPE22, error) {
	result := layoutRSPE22.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRSPE22(result.Item.(hl7.Group))
	return &v, nil
}

var layoutRTBZ92 = layout.Group("RTBZ92").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.Segment("ERR"))).With(layout.Segment("QAK")).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("RTBZ92RowDefinition").With(layout.Segment("RDF")).With(layout.Maybe(layout.List(layout.Segment("RDT")))))).With(layout.Maybe(layout.Segment("DSC")))

func NewRTBZ92Layout(msg hl7.Message) (*RTBZ92, error) {
	result := layoutRTBZ92.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewRTBZ92(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFKM08 = layout.Group("MFKM08").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("MFI")).With(layout.Maybe(layout.List(layout.Segment("MFA"))))

func NewMFKM08Layout(msg hl7.Message) (*MFKM08, error) {
	result := layoutMFKM08.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFKM08(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPGLPC8 = layout.Group("PGLPC8").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("PID")).With(layout.Maybe(layout.Group("PGLPC8PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))))).With(layout.List(layout.Group("PGLPC8Goal").With(layout.Segment("GOL")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PGLPC8GoalRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PGLPC8Pathway").With(layout.Segment("PTH")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PGLPC8Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("PGLPC8Problem").With(layout.Segment("PRB")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PGLPC8ProblemRole").With(layout.Segment("ROL")).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))).With(layout.Maybe(layout.List(layout.Group("PGLPC8ProblemObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))).With(layout.Maybe(layout.List(layout.Group("PGLPC8Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.Group("PGLPC8OrderDetail").With(layout.Group("PGLPC8Choice").With(layout.Maybe(layout.Segment("OBR"))).With(layout.Maybe(layout.Segment("Hxx")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR")))).With(layout.Maybe(layout.List(layout.Group("PGLPC8OrderObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("VAR"))))))))))))))

func NewPGLPC8Layout(msg hl7.Message) (*PGLPC8, error) {
	result := layoutPGLPC8.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPGLPC8(result.Item.(hl7.Group))
	return &v, nil
}

var layoutEHCE01 = layout.Group("EHCE01").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.List(layout.Segment("UAC")))).With(layout.Group("EHCE01InvoiceInformationSubmit").With(layout.Segment("IVC")).With(layout.Maybe(layout.Segment("PYE"))).With(layout.Maybe(layout.List(layout.Segment("CTD")))).With(layout.Maybe(layout.Segment("AUT"))).With(layout.Maybe(layout.List(layout.Segment("LOC")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.List(layout.Group("EHCE01ProductServiceSection").With(layout.Segment("PSS")).With(layout.List(layout.Group("EHCE01ProductServiceGroup").With(layout.Segment("PSG")).With(layout.Maybe(layout.List(layout.Segment("LOC")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Group("EHCE01PatientInfo").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ACC")))).With(layout.List(layout.Group("EHCE01Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))))).With(layout.Maybe(layout.List(layout.Group("EHCE01Diagnosis").With(layout.Segment("DG1")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Segment("OBX"))))))).With(layout.List(layout.Group("EHCE01ProductServiceLineItem").With(layout.Segment("PSL")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("ADJ")))).With(layout.Maybe(layout.List(layout.Segment("AUT")))).With(layout.Maybe(layout.List(layout.Segment("LOC")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))))).With(layout.Maybe(layout.List(layout.Group("EHCE01Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Segment("IPR")))))))))

func NewEHCE01Layout(msg hl7.Message) (*EHCE01, error) {
	result := layoutEHCE01.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewEHCE01(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA51 = layout.Group("ADTA51").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Segment("MRG")).With(layout.Segment("PV1"))

func NewADTA51Layout(msg hl7.Message) (*ADTA51, error) {
	result := layoutADTA51.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA51(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPMUB07 = layout.Group("PMUB07").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("STF")).With(layout.Maybe(layout.Segment("PRA"))).With(layout.Maybe(layout.List(layout.Group("PMUB07Certificate").With(layout.Segment("CER")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))))))

func NewPMUB07Layout(msg hl7.Message) (*PMUB07, error) {
	result := layoutPMUB07.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPMUB07(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMDMT09 = layout.Group("MDMT09").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Segment("PV1")).With(layout.Maybe(layout.List(layout.Group("MDMT09CommonOrder").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Group("MDMT09Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Segment("TXA")).With(layout.Maybe(layout.List(layout.Segment("CON"))))

func NewMDMT09Layout(msg hl7.Message) (*MDMT09, error) {
	result := layoutMDMT09.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMDMT09(result.Item.(hl7.Group))
	return &v, nil
}

var layoutDRCO47 = layout.Group("DRCO47").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("STF")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Group("DRCO47Donnor").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("DRCO47DonorRegistration").With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))).With(layout.List(layout.Group("DRCO47DonnorOrder").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))

func NewDRCO47Layout(msg hl7.Message) (*DRCO47, error) {
	result := layoutDRCO47.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewDRCO47(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSIUS23 = layout.Group("SIUS23").With(layout.Segment("MSH")).With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SIUS23Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SIUS23Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SIUS23Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS23GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS23LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS23PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSIUS23Layout(msg hl7.Message) (*SIUS23, error) {
	result := layoutSIUS23.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSIUS23(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFNM12 = layout.Group("MFNM12").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MFI")).With(layout.List(layout.Group("MFNM12MfObsAttributes").With(layout.Segment("MFE")).With(layout.Segment("OM1")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.Group("MFNM12MfObsOtherAttributes").With(layout.Segment("OM7")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))

func NewMFNM12Layout(msg hl7.Message) (*MFNM12, error) {
	result := layoutMFNM12.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFNM12(result.Item.(hl7.Group))
	return &v, nil
}

var layoutCSUC09 = layout.Group("CSUC09").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Group("CSUC09Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("CSUC09Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("CSR")).With(layout.List(layout.Group("CSUC09StudyPhase").With(layout.Maybe(layout.Segment("CSP"))).With(layout.List(layout.Group("CSUC09StudySchedule").With(layout.Maybe(layout.Segment("CSS"))).With(layout.List(layout.Group("CSUC09StudyObservation").With(layout.Maybe(layout.Group("CSUC09CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("CSUC09TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.List(layout.Group("CSUC09StudyPharm").With(layout.Maybe(layout.Group("CSUC09CommonOrder").With(layout.Maybe(layout.Segment("ORC"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.List(layout.Group("CSUC09RxAdmin").With(layout.Segment("RXA")).With(layout.Segment("RXR")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))))))))

func NewCSUC09Layout(msg hl7.Message) (*CSUC09, error) {
	result := layoutCSUC09.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewCSUC09(result.Item.(hl7.Group))
	return &v, nil
}

var layoutPEXP08 = layout.Group("PEXP08").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("PEXP08Visit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.List(layout.Group("PEXP08Experience").With(layout.Segment("PES")).With(layout.List(layout.Group("PEXP08PexObservation").With(layout.Segment("PEO")).With(layout.List(layout.Group("PEXP08PexCause").With(layout.Segment("PCR")).With(layout.Maybe(layout.Group("PEXP08RxOrder").With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.List(layout.Group("PEXP08TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))).With(layout.Maybe(layout.List(layout.Segment("RXR")))))).With(layout.Maybe(layout.List(layout.Group("PEXP08RxAdministration").With(layout.Segment("RXA")).With(layout.Maybe(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("PRB")))).With(layout.Maybe(layout.List(layout.Group("PEXP08Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("PEXP08AssociatedPerson").With(layout.Segment("NK1")).With(layout.Maybe(layout.Group("PEXP08AssociatedRxOrder").With(layout.Segment("RXE")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.List(layout.Group("PEXP08Nk1TimingQty").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2")))))).With(layout.Maybe(layout.List(layout.Segment("RXR")))))).With(layout.Maybe(layout.List(layout.Group("PEXP08AssociatedRxAdmin").With(layout.Segment("RXA")).With(layout.Maybe(layout.Segment("RXR"))).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))).With(layout.Maybe(layout.List(layout.Segment("PRB")))).With(layout.Maybe(layout.List(layout.Group("PEXP08AssociatedObservation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))))).With(layout.Maybe(layout.List(layout.Group("PEXP08Study").With(layout.Segment("CSR")).With(layout.Maybe(layout.List(layout.Segment("CSP")))))))))))))

func NewPEXP08Layout(msg hl7.Message) (*PEXP08, error) {
	result := layoutPEXP08.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewPEXP08(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFKM09 = layout.Group("MFKM09").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Segment("MFI")).With(layout.Maybe(layout.List(layout.Segment("MFA"))))

func NewMFKM09Layout(msg hl7.Message) (*MFKM09, error) {
	result := layoutMFKM09.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFKM09(result.Item.(hl7.Group))
	return &v, nil
}

var layoutORGO20 = layout.Group("ORGO20").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("ORGO20Response").With(layout.Maybe(layout.Group("ORGO20Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))))).With(layout.List(layout.Group("ORGO20Order").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("ORGO20Timing").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("ORGO20ObservationGroup").With(layout.Segment("OBR")))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("CTI")))).With(layout.Maybe(layout.List(layout.Group("ORGO20Specimen").With(layout.Segment("SPM")).With(layout.Maybe(layout.List(layout.Segment("SAC")))))))))))

func NewORGO20Layout(msg hl7.Message) (*ORGO20, error) {
	result := layoutORGO20.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewORGO20(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOMDO03 = layout.Group("OMDO03").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMDO03Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Group("OMDO03PatientVisit").With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("PRT")))))).With(layout.Maybe(layout.List(layout.Group("OMDO03Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.Segment("IN3")))))).With(layout.Maybe(layout.Segment("GT1"))).With(layout.Maybe(layout.List(layout.Segment("AL1")))))).With(layout.List(layout.Group("OMDO03OrderDiet").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OMDO03TimingDiet").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.Maybe(layout.Group("OMDO03Diet").With(layout.List(layout.Segment("ODS"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("OMDO03Observation").With(layout.Segment("OBX")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))))))).With(layout.Maybe(layout.List(layout.Group("OMDO03OrderTray").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT")))).With(layout.Maybe(layout.List(layout.Group("OMDO03TimingTray").With(layout.Segment("TQ1")).With(layout.Maybe(layout.List(layout.Segment("TQ2"))))))).With(layout.List(layout.Segment("ODT"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))

func NewOMDO03Layout(msg hl7.Message) (*OMDO03, error) {
	result := layoutOMDO03.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOMDO03(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA06 = layout.Group("ADTA06").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.Segment("MRG"))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Group("ADTA06Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("ADTA06Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("UB1"))).With(layout.Maybe(layout.Segment("UB2")))

func NewADTA06Layout(msg hl7.Message) (*ADTA06, error) {
	result := layoutADTA06.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA06(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSIUS15 = layout.Group("SIUS15").With(layout.Segment("MSH")).With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SIUS15Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SIUS15Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SIUS15Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS15GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS15LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS15PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSIUS15Layout(msg hl7.Message) (*SIUS15, error) {
	result := layoutSIUS15.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSIUS15(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA44 = layout.Group("ADTA44").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.List(layout.Group("ADTA44Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Segment("MRG"))))

func NewADTA44Layout(msg hl7.Message) (*ADTA44, error) {
	result := layoutADTA44.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA44(result.Item.(hl7.Group))
	return &v, nil
}

var layoutMFNM15 = layout.Group("MFNM15").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("MFI")).With(layout.List(layout.Group("MFNM15MfInvItem").With(layout.Segment("MFE")).With(layout.Segment("IIM"))))

func NewMFNM15Layout(msg hl7.Message) (*MFNM15, error) {
	result := layoutMFNM15.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewMFNM15(result.Item.(hl7.Group))
	return &v, nil
}

var layoutQBPZ95 = layout.Group("QBPZ95").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("QPD")).With(layout.Maybe(layout.Group("QBPZ95Qbp").With(layout.Maybe(layout.Segment("Hxx"))))).With(layout.Maybe(layout.Segment("RDF"))).With(layout.Segment("RCP")).With(layout.Maybe(layout.Segment("DSC")))

func NewQBPZ95Layout(msg hl7.Message) (*QBPZ95, error) {
	result := layoutQBPZ95.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewQBPZ95(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSIUS17 = layout.Group("SIUS17").With(layout.Segment("MSH")).With(layout.Segment("SCH")).With(layout.Maybe(layout.List(layout.Segment("TQ1")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Group("SIUS17Patient").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("DG1"))))))).With(layout.List(layout.Group("SIUS17Resources").With(layout.Segment("RGS")).With(layout.Maybe(layout.List(layout.Group("SIUS17Service").With(layout.Segment("AIS")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS17GeneralResource").With(layout.Segment("AIG")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS17LocationResource").With(layout.Segment("AIL")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))).With(layout.Maybe(layout.List(layout.Group("SIUS17PersonnelResource").With(layout.Segment("AIP")).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))))

func NewSIUS17Layout(msg hl7.Message) (*SIUS17, error) {
	result := layoutSIUS17.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSIUS17(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA03 = layout.Group("ADTA03").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Group("ADTA03Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("ADTA03Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("AUT")))).With(layout.Maybe(layout.List(layout.Segment("RF1"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("PDA")))

func NewADTA03Layout(msg hl7.Message) (*ADTA03, error) {
	result := layoutADTA03.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA03(result.Item.(hl7.Group))
	return &v, nil
}

var layoutOSUO41 = layout.Group("OSUO41").With(layout.Segment("MSH")).With(layout.Segment("MSA")).With(layout.Maybe(layout.List(layout.Segment("ERR")))).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.Segment("PID"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.List(layout.Group("OSUO41OrderStatus").With(layout.Segment("ORC")).With(layout.Maybe(layout.List(layout.Segment("PRT"))))))

func NewOSUO41Layout(msg hl7.Message) (*OSUO41, error) {
	result := layoutOSUO41.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewOSUO41(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA05 = layout.Group("ADTA05").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("NK1")))).With(layout.Segment("PV1")).With(layout.Maybe(layout.Segment("PV2"))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("DB1")))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("DG1")))).With(layout.Maybe(layout.Segment("DRG"))).With(layout.Maybe(layout.List(layout.Group("ADTA05Procedure").With(layout.Segment("PR1")).With(layout.Maybe(layout.List(layout.Segment("ROL"))))))).With(layout.Maybe(layout.List(layout.Segment("GT1")))).With(layout.Maybe(layout.List(layout.Group("ADTA05Insurance").With(layout.Segment("IN1")).With(layout.Maybe(layout.Segment("IN2"))).With(layout.Maybe(layout.List(layout.Segment("IN3")))).With(layout.Maybe(layout.List(layout.Segment("ROL")))).With(layout.Maybe(layout.List(layout.Segment("AUT")))).With(layout.Maybe(layout.List(layout.Segment("RF1"))))))).With(layout.Maybe(layout.Segment("ACC"))).With(layout.Maybe(layout.Segment("UB1"))).With(layout.Maybe(layout.Segment("UB2")))

func NewADTA05Layout(msg hl7.Message) (*ADTA05, error) {
	result := layoutADTA05.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA05(result.Item.(hl7.Group))
	return &v, nil
}

var layoutDERO44 = layout.Group("DERO44").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("STF")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Maybe(layout.Group("DERO44Donnor").With(layout.Segment("PID")).With(layout.Maybe(layout.Segment("PD1"))).With(layout.Maybe(layout.List(layout.Segment("OBX")))).With(layout.Maybe(layout.List(layout.Segment("NTE")))).With(layout.Maybe(layout.List(layout.Segment("AL1")))).With(layout.Maybe(layout.List(layout.Segment("ARV")))).With(layout.Maybe(layout.Group("DERO44DonorRegistration").With(layout.Maybe(layout.Segment("PV1"))).With(layout.Maybe(layout.List(layout.Segment("NTE")))))))).With(layout.List(layout.Group("DERO44DonnorOrder").With(layout.Segment("OBR")).With(layout.Maybe(layout.List(layout.Segment("NTE"))))))

func NewDERO44Layout(msg hl7.Message) (*DERO44, error) {
	result := layoutDERO44.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewDERO44(result.Item.(hl7.Group))
	return &v, nil
}

var layoutADTA20 = layout.Group("ADTA20").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.Segment("EVN")).With(layout.Segment("NPU"))

func NewADTA20Layout(msg hl7.Message) (*ADTA20, error) {
	result := layoutADTA20.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewADTA20(result.Item.(hl7.Group))
	return &v, nil
}

var layoutSTIS30 = layout.Group("STIS30").With(layout.Segment("MSH")).With(layout.Maybe(layout.List(layout.Segment("SFT")))).With(layout.Maybe(layout.Segment("UAC"))).With(layout.List(layout.Segment("SLT")))

func NewSTIS30Layout(msg hl7.Message) (*STIS30, error) {
	result := layoutSTIS30.Parse(layout.NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	v := NewSTIS30(result.Item.(hl7.Group))
	return &v, nil
}
