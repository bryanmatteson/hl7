package layout

import "hl7"

var msh = Segment("MSH")
var sft = Segment("SFT")
var uac = Segment("UAC")
var dsc = Segment("DSC")
var pid = Segment("PID")
var pd1 = Segment("PD1")
var nte = Segment("NTE")
var nk1 = Segment("NK1")
var pv1 = Segment("PV1")
var pv2 = Segment("PV2")
var orc = Segment("ORC")
var obr = Segment("OBR")
var ctd = Segment("CTD")
var spm = Segment("SPM")
var obx = Segment("OBX")
var cti = Segment("CTI")
var ft1 = Segment("FT1")
var tq1 = Segment("TQ1")
var tq2 = Segment("TQ2")
var rol = Segment("ROL")

// ORUR01Layout ...
var orur01 = Group("ORU_R01").
	With(msh).
	Maybe(List(sft)).
	Maybe(uac).
	List(Group("PATIENT_RESULT").
		Maybe(Group("PATIENT").
			With(pid).
			Maybe(pd1).
			Maybe(List(nte)).
			Maybe(List(nk1)).
			Maybe(List(obx)).
			Maybe(Group("PATIENT_VISIT").
				With(pv1).
				Maybe(pv2))).
		List(Group("ORDER_OBSERVATION").
			Maybe(orc).
			With(obr).
			Maybe(List(nte)).
			Maybe(List(rol)).
			Maybe(List(Group("TIMING_QTY").
				With(tq1).
				Maybe(List(tq2)))).
			Maybe(ctd).
			Maybe(List(Group("OBSERVATION_GROUP").
				With(obx).
				Maybe(List(nte)))).
			Maybe(List(ft1)).
			Maybe(List(cti)).
			Maybe(List(Group("SPECIMEN").
				With(spm).
				Maybe(List(obx)))))).
	Maybe(dsc)

func NewORUR01(msg hl7.Message) (hl7.Group, error) {
	result := orur01.Parse(NewInput(msg))
	if !result.Success {
		return nil, result.Error
	}
	return result.Item.(hl7.Group), nil
}
