package hl7v28

import hl7 "hl7"

type XTN struct {
	TelephoneNumber                   ST
	TelecommunicationUseCode          ID
	TelecommunicationEquipmentType    ID
	CommunicationAddress              ST
	CountryCode                       SNM
	AreaCityCode                      SNM
	LocalNumber                       SNM
	Extension                         SNM
	AnyText                           ST
	ExtensionPrefix                   ST
	SpeedDialCode                     ST
	UnformattedTelephoneNumber        ST
	EffectiveStartDate                DTM
	ExpirationDate                    DTM
	ExpirationReason                  CWE
	ProtectionCode                    CWE
	SharedTelecommunicationIdentifier EI
	PreferenceOrder                   NM
}

// NewXTN creates an implementation of XTN
func NewXTN(input hl7.HL7Type) XTN {
	v := XTN{}
	v.TelephoneNumber = NewST(input.Index(0))
	v.TelecommunicationUseCode = NewID(input.Index(1))
	v.TelecommunicationEquipmentType = NewID(input.Index(2))
	v.CommunicationAddress = NewST(input.Index(3))
	v.CountryCode = NewSNM(input.Index(4))
	v.AreaCityCode = NewSNM(input.Index(5))
	v.LocalNumber = NewSNM(input.Index(6))
	v.Extension = NewSNM(input.Index(7))
	v.AnyText = NewST(input.Index(8))
	v.ExtensionPrefix = NewST(input.Index(9))
	v.SpeedDialCode = NewST(input.Index(10))
	v.UnformattedTelephoneNumber = NewST(input.Index(11))
	v.EffectiveStartDate = NewDTM(input.Index(12))
	v.ExpirationDate = NewDTM(input.Index(13))
	v.ExpirationReason = NewCWE(input.Index(14))
	v.ProtectionCode = NewCWE(input.Index(15))
	v.SharedTelecommunicationIdentifier = NewEI(input.Index(16))
	v.PreferenceOrder = NewNM(input.Index(17))
	return v
}

// NewXTNSlice will construct a slice of type XTN
func NewXTNSlice(input hl7.HL7Type) []XTN {
	values := make([]XTN, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewXTN(input.Index(i))
	}
	return values
}

type XCN struct {
	PersonIdentifier                            ST
	FamilyName                                  FN
	GivenName                                   ST
	SecondAndFurtherGivenNamesOrInitialsThereof ST
	SuffixeGJrOrIii                             ST
	PrefixeGDr                                  ST
	DegreeeGMd                                  ST
	SourceTable                                 CWE
	AssigningAuthority                          HD
	NameTypeCode                                ID
	IdentifierCheckDigit                        ST
	CheckDigitScheme                            ID
	IdentifierTypeCode                          ID
	AssigningFacility                           HD
	NameRepresentationCode                      ID
	NameContext                                 CWE
	NameValidityRange                           ST
	NameAssemblyOrder                           ID
	EffectiveDate                               DTM
	ExpirationDate                              DTM
	ProfessionalSuffix                          ST
	AssigningJurisdiction                       CWE
	AssigningAgencyOrDepartment                 CWE
	SecurityCheck                               ST
	SecurityCheckScheme                         ID
}

// NewXCN creates an implementation of XCN
func NewXCN(input hl7.HL7Type) XCN {
	v := XCN{}
	v.PersonIdentifier = NewST(input.Index(0))
	v.FamilyName = NewFN(input.Index(1))
	v.GivenName = NewST(input.Index(2))
	v.SecondAndFurtherGivenNamesOrInitialsThereof = NewST(input.Index(3))
	v.SuffixeGJrOrIii = NewST(input.Index(4))
	v.PrefixeGDr = NewST(input.Index(5))
	v.DegreeeGMd = NewST(input.Index(6))
	v.SourceTable = NewCWE(input.Index(7))
	v.AssigningAuthority = NewHD(input.Index(8))
	v.NameTypeCode = NewID(input.Index(9))
	v.IdentifierCheckDigit = NewST(input.Index(10))
	v.CheckDigitScheme = NewID(input.Index(11))
	v.IdentifierTypeCode = NewID(input.Index(12))
	v.AssigningFacility = NewHD(input.Index(13))
	v.NameRepresentationCode = NewID(input.Index(14))
	v.NameContext = NewCWE(input.Index(15))
	v.NameValidityRange = NewST(input.Index(16))
	v.NameAssemblyOrder = NewID(input.Index(17))
	v.EffectiveDate = NewDTM(input.Index(18))
	v.ExpirationDate = NewDTM(input.Index(19))
	v.ProfessionalSuffix = NewST(input.Index(20))
	v.AssigningJurisdiction = NewCWE(input.Index(21))
	v.AssigningAgencyOrDepartment = NewCWE(input.Index(22))
	v.SecurityCheck = NewST(input.Index(23))
	v.SecurityCheckScheme = NewID(input.Index(24))
	return v
}

// NewXCNSlice will construct a slice of type XCN
func NewXCNSlice(input hl7.HL7Type) []XCN {
	values := make([]XCN, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewXCN(input.Index(i))
	}
	return values
}

type XPN struct {
	FamilyName                                  FN
	GivenName                                   ST
	SecondAndFurtherGivenNamesOrInitialsThereof ST
	SuffixeGJrOrIii                             ST
	PrefixeGDr                                  ST
	DegreeeGMd                                  ST
	NameTypeCode                                ID
	NameRepresentationCode                      ID
	NameContext                                 CWE
	NameValidityRange                           ST
	NameAssemblyOrder                           ID
	EffectiveDate                               DTM
	ExpirationDate                              DTM
	ProfessionalSuffix                          ST
	CalledBy                                    ST
}

// NewXPN creates an implementation of XPN
func NewXPN(input hl7.HL7Type) XPN {
	v := XPN{}
	v.FamilyName = NewFN(input.Index(0))
	v.GivenName = NewST(input.Index(1))
	v.SecondAndFurtherGivenNamesOrInitialsThereof = NewST(input.Index(2))
	v.SuffixeGJrOrIii = NewST(input.Index(3))
	v.PrefixeGDr = NewST(input.Index(4))
	v.DegreeeGMd = NewST(input.Index(5))
	v.NameTypeCode = NewID(input.Index(6))
	v.NameRepresentationCode = NewID(input.Index(7))
	v.NameContext = NewCWE(input.Index(8))
	v.NameValidityRange = NewST(input.Index(9))
	v.NameAssemblyOrder = NewID(input.Index(10))
	v.EffectiveDate = NewDTM(input.Index(11))
	v.ExpirationDate = NewDTM(input.Index(12))
	v.ProfessionalSuffix = NewST(input.Index(13))
	v.CalledBy = NewST(input.Index(14))
	return v
}

// NewXPNSlice will construct a slice of type XPN
func NewXPNSlice(input hl7.HL7Type) []XPN {
	values := make([]XPN, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewXPN(input.Index(i))
	}
	return values
}

type EI struct {
	EntityIdentifier ST
	NamespaceID      IS
	UniversalID      ST
	UniversalIDType  ID
}

// NewEI creates an implementation of EI
func NewEI(input hl7.HL7Type) EI {
	v := EI{}
	v.EntityIdentifier = NewST(input.Index(0))
	v.NamespaceID = NewIS(input.Index(1))
	v.UniversalID = NewST(input.Index(2))
	v.UniversalIDType = NewID(input.Index(3))
	return v
}

// NewEISlice will construct a slice of type EI
func NewEISlice(input hl7.HL7Type) []EI {
	values := make([]EI, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEI(input.Index(i))
	}
	return values
}

type MOP struct {
	MoneyOrPercentageIndicator ID
	MoneyOrPercentageQuantity  NM
	MonetaryDenomination       ID
}

// NewMOP creates an implementation of MOP
func NewMOP(input hl7.HL7Type) MOP {
	v := MOP{}
	v.MoneyOrPercentageIndicator = NewID(input.Index(0))
	v.MoneyOrPercentageQuantity = NewNM(input.Index(1))
	v.MonetaryDenomination = NewID(input.Index(2))
	return v
}

// NewMOPSlice will construct a slice of type MOP
func NewMOPSlice(input hl7.HL7Type) []MOP {
	values := make([]MOP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMOP(input.Index(i))
	}
	return values
}

type CCD struct {
	InvocationEvent ID
	DateTime        DTM
}

// NewCCD creates an implementation of CCD
func NewCCD(input hl7.HL7Type) CCD {
	v := CCD{}
	v.InvocationEvent = NewID(input.Index(0))
	v.DateTime = NewDTM(input.Index(1))
	return v
}

// NewCCDSlice will construct a slice of type CCD
func NewCCDSlice(input hl7.HL7Type) []CCD {
	values := make([]CCD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCCD(input.Index(i))
	}
	return values
}

type RFR struct {
	NumericRange        NR
	AdministrativeSex   CWE
	AgeRange            NR
	GestationalAgeRange NR
	Species             ST
	RaceSubspecies      ST
	Conditions          TX
}

// NewRFR creates an implementation of RFR
func NewRFR(input hl7.HL7Type) RFR {
	v := RFR{}
	v.NumericRange = NewNR(input.Index(0))
	v.AdministrativeSex = NewCWE(input.Index(1))
	v.AgeRange = NewNR(input.Index(2))
	v.GestationalAgeRange = NewNR(input.Index(3))
	v.Species = NewST(input.Index(4))
	v.RaceSubspecies = NewST(input.Index(5))
	v.Conditions = NewTX(input.Index(6))
	return v
}

// NewRFRSlice will construct a slice of type RFR
func NewRFRSlice(input hl7.HL7Type) []RFR {
	values := make([]RFR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRFR(input.Index(i))
	}
	return values
}

type SAD struct {
	StreetOrMailingAddress ST
	StreetName             ST
	DwellingNumber         ST
}

// NewSAD creates an implementation of SAD
func NewSAD(input hl7.HL7Type) SAD {
	v := SAD{}
	v.StreetOrMailingAddress = NewST(input.Index(0))
	v.StreetName = NewST(input.Index(1))
	v.DwellingNumber = NewST(input.Index(2))
	return v
}

// NewSADSlice will construct a slice of type SAD
func NewSADSlice(input hl7.HL7Type) []SAD {
	values := make([]SAD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSAD(input.Index(i))
	}
	return values
}

type WVI struct {
	ChannelNumber NM
	ChannelName   ST
}

// NewWVI creates an implementation of WVI
func NewWVI(input hl7.HL7Type) WVI {
	v := WVI{}
	v.ChannelNumber = NewNM(input.Index(0))
	v.ChannelName = NewST(input.Index(1))
	return v
}

// NewWVISlice will construct a slice of type WVI
func NewWVISlice(input hl7.HL7Type) []WVI {
	values := make([]WVI, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewWVI(input.Index(i))
	}
	return values
}

type NR struct {
	LowValue  NM
	HighValue NM
}

// NewNR creates an implementation of NR
func NewNR(input hl7.HL7Type) NR {
	v := NR{}
	v.LowValue = NewNM(input.Index(0))
	v.HighValue = NewNM(input.Index(1))
	return v
}

// NewNRSlice will construct a slice of type NR
func NewNRSlice(input hl7.HL7Type) []NR {
	values := make([]NR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewNR(input.Index(i))
	}
	return values
}

type CD struct {
	ChannelIdentifier            WVI
	WaveformSource               WVS
	ChannelSensitivityAndUnits   CSU
	ChannelCalibrationParameters CCP
	ChannelSamplingFrequency     NM
	MinimumAndMaximumDataValues  NR
}

// NewCD creates an implementation of CD
func NewCD(input hl7.HL7Type) CD {
	v := CD{}
	v.ChannelIdentifier = NewWVI(input.Index(0))
	v.WaveformSource = NewWVS(input.Index(1))
	v.ChannelSensitivityAndUnits = NewCSU(input.Index(2))
	v.ChannelCalibrationParameters = NewCCP(input.Index(3))
	v.ChannelSamplingFrequency = NewNM(input.Index(4))
	v.MinimumAndMaximumDataValues = NewNR(input.Index(5))
	return v
}

// NewCDSlice will construct a slice of type CD
func NewCDSlice(input hl7.HL7Type) []CD {
	values := make([]CD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCD(input.Index(i))
	}
	return values
}

type ED struct {
	SourceApplication HD
	TypeOfData        ID
	DataSubtype       ID
	Encoding          ID
	Data              TX
}

// NewED creates an implementation of ED
func NewED(input hl7.HL7Type) ED {
	v := ED{}
	v.SourceApplication = NewHD(input.Index(0))
	v.TypeOfData = NewID(input.Index(1))
	v.DataSubtype = NewID(input.Index(2))
	v.Encoding = NewID(input.Index(3))
	v.Data = NewTX(input.Index(4))
	return v
}

// NewEDSlice will construct a slice of type ED
func NewEDSlice(input hl7.HL7Type) []ED {
	values := make([]ED, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewED(input.Index(i))
	}
	return values
}

type MOC struct {
	MonetaryAmount MO
	ChargeCode     CWE
}

// NewMOC creates an implementation of MOC
func NewMOC(input hl7.HL7Type) MOC {
	v := MOC{}
	v.MonetaryAmount = NewMO(input.Index(0))
	v.ChargeCode = NewCWE(input.Index(1))
	return v
}

// NewMOCSlice will construct a slice of type MOC
func NewMOCSlice(input hl7.HL7Type) []MOC {
	values := make([]MOC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMOC(input.Index(i))
	}
	return values
}

/*
ST - String Data
The String data type is used for text data when the appearance of text does not bear meaning. This is true for formalized text, symbols and formal expressions, and all kinds of names intended for machine processing (e.g., sorting, querying, indexing, etc.).

String data is left justified (i.e., no leading blank space) with trailing blanks optional. Any displayable (printable) ACSII characters (hexadecimal values between 20 and 7E, inclusive, or ASCII decimal values between 32 and 126), except the defined escape characters and defined delimiter characters.

Example 1: A textual ST field:
|almost any data at all|

Example 2:  URL encoded in an ST component:
^http://www.pacs.poupon.edu/wado.jsp^

Example 3:  ISO OID encoded in an ST subcomponent:
&2.16.840.1.113883.1.1&

To include any HL7 delimiter character (except the segment terminator) within a string data field, use the appropriate HL7 escape sequence (see Section 2.7.1, "Formatting Codes”).

Minimum Length: Not specified for the type. May be specified in the context of use. Defaults to 1
Maximum Length: Not specified for the type. May be specified in the context of use

ST has no inbuilt semantics – these are assigned where the ST is used. In each case where ST is used, minimum, maximum, and conformance lengths may be specified.  Unless specified in the context of use, values of type ST may not be truncated.

Usage note: The ST data type is intended for short strings (e.g., less than 1000 characters). For longer strings the TX or FT data types should be used (see Sections 2.A.79, “TX - text data” or 2.A.31, “FT - formatted text data”).

Alternate character set note: ST - string data may also be used to express other character sets. See Section 2.15.9.18, "Character set," and Section 2.15.9.20, "Alternate character set handling" for details.
*/
type ST string

// NewST creates an implementation of ST
func NewST(input hl7.HL7Type) ST {
	return ST(input.String())
}

// NewSTSlice will construct a slice of type ST
func NewSTSlice(input hl7.HL7Type) []ST {
	values := make([]ST, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewST(input.Index(i))
	}
	return values
}

type CF struct {
	Identifier                           ST
	FormattedText                        FT
	NameOfCodingSystem                   ID
	AlternateIdentifier                  ST
	AlternateFormattedText               FT
	NameOfAlternateCodingSystem          ID
	CodingSystemVersionID                ST
	AlternateCodingSystemVersionID       ST
	OriginalText                         ST
	SecondAlternateIdentifier            ST
	SecondAlternateFormattedText         FT
	NameOfSecondAlternateCodingSystem    ID
	SecondAlternateCodingSystemVersionID ST
	CodingSystemOID                      ST
	ValueSetOID                          ST
	ValueSetVersionID                    DTM
	AlternateCodingSystemOID             ST
	AlternateValueSetOID                 ST
	AlternateValueSetVersionID           DTM
	SecondAlternateCodingSystemOID       ST
	SecondAlternateValueSetOID           ST
	SecondAlternateValueSetVersionID     DTM
}

// NewCF creates an implementation of CF
func NewCF(input hl7.HL7Type) CF {
	v := CF{}
	v.Identifier = NewST(input.Index(0))
	v.FormattedText = NewFT(input.Index(1))
	v.NameOfCodingSystem = NewID(input.Index(2))
	v.AlternateIdentifier = NewST(input.Index(3))
	v.AlternateFormattedText = NewFT(input.Index(4))
	v.NameOfAlternateCodingSystem = NewID(input.Index(5))
	v.CodingSystemVersionID = NewST(input.Index(6))
	v.AlternateCodingSystemVersionID = NewST(input.Index(7))
	v.OriginalText = NewST(input.Index(8))
	v.SecondAlternateIdentifier = NewST(input.Index(9))
	v.SecondAlternateFormattedText = NewFT(input.Index(10))
	v.NameOfSecondAlternateCodingSystem = NewID(input.Index(11))
	v.SecondAlternateCodingSystemVersionID = NewST(input.Index(12))
	v.CodingSystemOID = NewST(input.Index(13))
	v.ValueSetOID = NewST(input.Index(14))
	v.ValueSetVersionID = NewDTM(input.Index(15))
	v.AlternateCodingSystemOID = NewST(input.Index(16))
	v.AlternateValueSetOID = NewST(input.Index(17))
	v.AlternateValueSetVersionID = NewDTM(input.Index(18))
	v.SecondAlternateCodingSystemOID = NewST(input.Index(19))
	v.SecondAlternateValueSetOID = NewST(input.Index(20))
	v.SecondAlternateValueSetVersionID = NewDTM(input.Index(21))
	return v
}

// NewCFSlice will construct a slice of type CF
func NewCFSlice(input hl7.HL7Type) []CF {
	values := make([]CF, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCF(input.Index(i))
	}
	return values
}

type DR struct {
	RangeStartDateTime DTM
	RangeEndDateTime   DTM
}

// NewDR creates an implementation of DR
func NewDR(input hl7.HL7Type) DR {
	v := DR{}
	v.RangeStartDateTime = NewDTM(input.Index(0))
	v.RangeEndDateTime = NewDTM(input.Index(1))
	return v
}

// NewDRSlice will construct a slice of type DR
func NewDRSlice(input hl7.HL7Type) []DR {
	values := make([]DR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDR(input.Index(i))
	}
	return values
}

/*
SNM - String Of Telephone Number Digits
A string whose characters are limited to "+" and/or the decimal digits 0  through 9.  As a string, leading zeros are always considered significant.

Used only in the XTN data type as of v2.7.

Minimum Length: 1
Maximum Length: Not specified for the type. May be specified in the context of use

SNM is used for telephone numbers, so it is never appropriate to truncate values of type SNM.
*/
type SNM string

// NewSNM creates an implementation of SNM
func NewSNM(input hl7.HL7Type) SNM {
	return SNM(input.String())
}

// NewSNMSlice will construct a slice of type SNM
func NewSNMSlice(input hl7.HL7Type) []SNM {
	values := make([]SNM, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSNM(input.Index(i))
	}
	return values
}

type SN struct {
	Comparator      ST
	Num1            NM
	SeparatorSuffix ST
	Num2            NM
}

// NewSN creates an implementation of SN
func NewSN(input hl7.HL7Type) SN {
	v := SN{}
	v.Comparator = NewST(input.Index(0))
	v.Num1 = NewNM(input.Index(1))
	v.SeparatorSuffix = NewST(input.Index(2))
	v.Num2 = NewNM(input.Index(3))
	return v
}

// NewSNSlice will construct a slice of type SN
func NewSNSlice(input hl7.HL7Type) []SN {
	values := make([]SN, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSN(input.Index(i))
	}
	return values
}

type RI struct {
	RepeatPattern        CWE
	ExplicitTimeInterval ST
}

// NewRI creates an implementation of RI
func NewRI(input hl7.HL7Type) RI {
	v := RI{}
	v.RepeatPattern = NewCWE(input.Index(0))
	v.ExplicitTimeInterval = NewST(input.Index(1))
	return v
}

// NewRISlice will construct a slice of type RI
func NewRISlice(input hl7.HL7Type) []RI {
	values := make([]RI, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRI(input.Index(i))
	}
	return values
}

/*
TM - Time
Specifies the hour of the day with optional minutes, seconds, fraction of second using a 24-hour clock notation and time zone.

As of v 2.3, the number of characters populated (excluding the time zone specification) specifies the precision.

Format: HH[MM[SS[.S[S[S[S]]]]]][+/-ZZZZ]

The fractional seconds could be sent by a transmitter who requires greater precision than whole seconds. Fractional representations of minutes, hours or other higher-order units of time are not permitted.

Note: The time zone [+/-ZZZZ], when used, is restricted to legally-defined time zones and is represented in HHMM format.

The time zone of the sender may be sent optionally as an offset from the coordinated universal time (previously known as Greenwich Mean Time). Where the time zone is not present in a particular TM field but is included as part of the date/time field in the MSH segment, the MSH value will be used as the default time zone. Otherwise, the time is understood to refer to the local time of the sender.

Prior to v 2.3, this data type was specified in the format HHMM[SS[.SSSS]][+/-ZZZZ]. As of v 2.3 minutes are no longer required. By site-specific agreement, HHMM[SS[.SSSS]][+/-ZZZZ] may be used where backward compatibility must be maintained.This corresponds a minimum length of 4.

The TM data type does not follow the normal truncation pattern, and the truncation character is never valid in the TM data type. Instead, the truncation behavior is based on the semantics of times.

Unless otherwise specified in the context where the DTM type is used, the DTM type may be truncated to a particular minute. When a TM is truncated, the truncated form SHALL still be a valid TM type. Refer to Chapter 2, section 2.5.5.2, "Truncation Pattern", for further information.
*/
type TM string

// NewTM creates an implementation of TM
func NewTM(input hl7.HL7Type) TM {
	return TM(input.String())
}

// NewTMSlice will construct a slice of type TM
func NewTMSlice(input hl7.HL7Type) []TM {
	values := make([]TM, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewTM(input.Index(i))
	}
	return values
}

type RMC struct {
	RoomType          CWE
	AmountType        CWE
	CoverageAmount    ST
	MoneyOrPercentage MOP
}

// NewRMC creates an implementation of RMC
func NewRMC(input hl7.HL7Type) RMC {
	v := RMC{}
	v.RoomType = NewCWE(input.Index(0))
	v.AmountType = NewCWE(input.Index(1))
	v.CoverageAmount = NewST(input.Index(2))
	v.MoneyOrPercentage = NewMOP(input.Index(3))
	return v
}

// NewRMCSlice will construct a slice of type RMC
func NewRMCSlice(input hl7.HL7Type) []RMC {
	values := make([]RMC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRMC(input.Index(i))
	}
	return values
}

type AD struct {
	StreetAddress              ST
	OtherDesignation           ST
	City                       ST
	StateOrProvince            ST
	ZipOrPostalCode            ST
	Country                    ID
	AddressType                ID
	OtherGeographicDesignation ST
}

// NewAD creates an implementation of AD
func NewAD(input hl7.HL7Type) AD {
	v := AD{}
	v.StreetAddress = NewST(input.Index(0))
	v.OtherDesignation = NewST(input.Index(1))
	v.City = NewST(input.Index(2))
	v.StateOrProvince = NewST(input.Index(3))
	v.ZipOrPostalCode = NewST(input.Index(4))
	v.Country = NewID(input.Index(5))
	v.AddressType = NewID(input.Index(6))
	v.OtherGeographicDesignation = NewST(input.Index(7))
	return v
}

// NewADSlice will construct a slice of type AD
func NewADSlice(input hl7.HL7Type) []AD {
	values := make([]AD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewAD(input.Index(i))
	}
	return values
}

type OCD struct {
	OccurrenceCode CNE
	OccurrenceDate DT
}

// NewOCD creates an implementation of OCD
func NewOCD(input hl7.HL7Type) OCD {
	v := OCD{}
	v.OccurrenceCode = NewCNE(input.Index(0))
	v.OccurrenceDate = NewDT(input.Index(1))
	return v
}

// NewOCDSlice will construct a slice of type OCD
func NewOCDSlice(input hl7.HL7Type) []OCD {
	values := make([]OCD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOCD(input.Index(i))
	}
	return values
}

type VID struct {
	VersionID                ID
	InternationalizationCode CWE
	InternationalVersionID   CWE
}

// NewVID creates an implementation of VID
func NewVID(input hl7.HL7Type) VID {
	v := VID{}
	v.VersionID = NewID(input.Index(0))
	v.InternationalizationCode = NewCWE(input.Index(1))
	v.InternationalVersionID = NewCWE(input.Index(2))
	return v
}

// NewVIDSlice will construct a slice of type VID
func NewVIDSlice(input hl7.HL7Type) []VID {
	values := make([]VID, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewVID(input.Index(i))
	}
	return values
}

type NDL struct {
	Name                CNN
	StartDateTime       DTM
	EndDateTime         DTM
	PointOfCare         IS
	Room                IS
	Bed                 IS
	Facility            HD
	LocationStatus      IS
	PatientLocationType IS
	Building            IS
	Floor               IS
}

// NewNDL creates an implementation of NDL
func NewNDL(input hl7.HL7Type) NDL {
	v := NDL{}
	v.Name = NewCNN(input.Index(0))
	v.StartDateTime = NewDTM(input.Index(1))
	v.EndDateTime = NewDTM(input.Index(2))
	v.PointOfCare = NewIS(input.Index(3))
	v.Room = NewIS(input.Index(4))
	v.Bed = NewIS(input.Index(5))
	v.Facility = NewHD(input.Index(6))
	v.LocationStatus = NewIS(input.Index(7))
	v.PatientLocationType = NewIS(input.Index(8))
	v.Building = NewIS(input.Index(9))
	v.Floor = NewIS(input.Index(10))
	return v
}

// NewNDLSlice will construct a slice of type NDL
func NewNDLSlice(input hl7.HL7Type) []NDL {
	values := make([]NDL, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewNDL(input.Index(i))
	}
	return values
}

type DLT struct {
	NormalRange       NR
	NumericThreshold  NM
	ChangeComputation ID
	DaysRetained      NM
}

// NewDLT creates an implementation of DLT
func NewDLT(input hl7.HL7Type) DLT {
	v := DLT{}
	v.NormalRange = NewNR(input.Index(0))
	v.NumericThreshold = NewNM(input.Index(1))
	v.ChangeComputation = NewID(input.Index(2))
	v.DaysRetained = NewNM(input.Index(3))
	return v
}

// NewDLTSlice will construct a slice of type DLT
func NewDLTSlice(input hl7.HL7Type) []DLT {
	values := make([]DLT, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDLT(input.Index(i))
	}
	return values
}

type EIP struct {
	PlacerAssignedIdentifier EI
	FillerAssignedIdentifier EI
}

// NewEIP creates an implementation of EIP
func NewEIP(input hl7.HL7Type) EIP {
	v := EIP{}
	v.PlacerAssignedIdentifier = NewEI(input.Index(0))
	v.FillerAssignedIdentifier = NewEI(input.Index(1))
	return v
}

// NewEIPSlice will construct a slice of type EIP
func NewEIPSlice(input hl7.HL7Type) []EIP {
	values := make([]EIP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewEIP(input.Index(i))
	}
	return values
}

type WVS struct {
	SourceOneName ST
	SourceTwoName ST
}

// NewWVS creates an implementation of WVS
func NewWVS(input hl7.HL7Type) WVS {
	v := WVS{}
	v.SourceOneName = NewST(input.Index(0))
	v.SourceTwoName = NewST(input.Index(1))
	return v
}

// NewWVSSlice will construct a slice of type WVS
func NewWVSSlice(input hl7.HL7Type) []WVS {
	values := make([]WVS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewWVS(input.Index(i))
	}
	return values
}

type CP struct {
	Price      MO
	PriceType  ID
	FromValue  NM
	ToValue    NM
	RangeUnits CWE
	RangeType  ID
}

// NewCP creates an implementation of CP
func NewCP(input hl7.HL7Type) CP {
	v := CP{}
	v.Price = NewMO(input.Index(0))
	v.PriceType = NewID(input.Index(1))
	v.FromValue = NewNM(input.Index(2))
	v.ToValue = NewNM(input.Index(3))
	v.RangeUnits = NewCWE(input.Index(4))
	v.RangeType = NewID(input.Index(5))
	return v
}

// NewCPSlice will construct a slice of type CP
func NewCPSlice(input hl7.HL7Type) []CP {
	values := make([]CP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCP(input.Index(i))
	}
	return values
}

type OSP struct {
	OccurrenceSpanCode      CNE
	OccurrenceSpanStartDate DT
	OccurrenceSpanStopDate  DT
}

// NewOSP creates an implementation of OSP
func NewOSP(input hl7.HL7Type) OSP {
	v := OSP{}
	v.OccurrenceSpanCode = NewCNE(input.Index(0))
	v.OccurrenceSpanStartDate = NewDT(input.Index(1))
	v.OccurrenceSpanStopDate = NewDT(input.Index(2))
	return v
}

// NewOSPSlice will construct a slice of type OSP
func NewOSPSlice(input hl7.HL7Type) []OSP {
	values := make([]OSP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewOSP(input.Index(i))
	}
	return values
}

type PRL struct {
	ParentObservationIdentifier      CWE
	ParentObservationSubIdentifier   ST
	ParentObservationValueDescriptor TX
}

// NewPRL creates an implementation of PRL
func NewPRL(input hl7.HL7Type) PRL {
	v := PRL{}
	v.ParentObservationIdentifier = NewCWE(input.Index(0))
	v.ParentObservationSubIdentifier = NewST(input.Index(1))
	v.ParentObservationValueDescriptor = NewTX(input.Index(2))
	return v
}

// NewPRLSlice will construct a slice of type PRL
func NewPRLSlice(input hl7.HL7Type) []PRL {
	values := make([]PRL, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPRL(input.Index(i))
	}
	return values
}

type CWE struct {
	Identifier                           ST
	Text                                 ST
	NameOfCodingSystem                   ID
	AlternateIdentifier                  ST
	AlternateText                        ST
	NameOfAlternateCodingSystem          ID
	CodingSystemVersionID                ST
	AlternateCodingSystemVersionID       ST
	OriginalText                         ST
	SecondAlternateIdentifier            ST
	SecondAlternateText                  ST
	NameOfSecondAlternateCodingSystem    ID
	SecondAlternateCodingSystemVersionID ST
	CodingSystemOID                      ST
	ValueSetOID                          ST
	ValueSetVersionID                    DTM
	AlternateCodingSystemOID             ST
	AlternateValueSetOID                 ST
	AlternateValueSetVersionID           DTM
	SecondAlternateCodingSystemOID       ST
	SecondAlternateValueSetOID           ST
	SecondAlternateValueSetVersionID     DTM
}

// NewCWE creates an implementation of CWE
func NewCWE(input hl7.HL7Type) CWE {
	v := CWE{}
	v.Identifier = NewST(input.Index(0))
	v.Text = NewST(input.Index(1))
	v.NameOfCodingSystem = NewID(input.Index(2))
	v.AlternateIdentifier = NewST(input.Index(3))
	v.AlternateText = NewST(input.Index(4))
	v.NameOfAlternateCodingSystem = NewID(input.Index(5))
	v.CodingSystemVersionID = NewST(input.Index(6))
	v.AlternateCodingSystemVersionID = NewST(input.Index(7))
	v.OriginalText = NewST(input.Index(8))
	v.SecondAlternateIdentifier = NewST(input.Index(9))
	v.SecondAlternateText = NewST(input.Index(10))
	v.NameOfSecondAlternateCodingSystem = NewID(input.Index(11))
	v.SecondAlternateCodingSystemVersionID = NewST(input.Index(12))
	v.CodingSystemOID = NewST(input.Index(13))
	v.ValueSetOID = NewST(input.Index(14))
	v.ValueSetVersionID = NewDTM(input.Index(15))
	v.AlternateCodingSystemOID = NewST(input.Index(16))
	v.AlternateValueSetOID = NewST(input.Index(17))
	v.AlternateValueSetVersionID = NewDTM(input.Index(18))
	v.SecondAlternateCodingSystemOID = NewST(input.Index(19))
	v.SecondAlternateValueSetOID = NewST(input.Index(20))
	v.SecondAlternateValueSetVersionID = NewDTM(input.Index(21))
	return v
}

// NewCWESlice will construct a slice of type CWE
func NewCWESlice(input hl7.HL7Type) []CWE {
	values := make([]CWE, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCWE(input.Index(i))
	}
	return values
}

type RP struct {
	Pointer       ST
	ApplicationID HD
	TypeOfData    ID
	Subtype       ID
}

// NewRP creates an implementation of RP
func NewRP(input hl7.HL7Type) RP {
	v := RP{}
	v.Pointer = NewST(input.Index(0))
	v.ApplicationID = NewHD(input.Index(1))
	v.TypeOfData = NewID(input.Index(2))
	v.Subtype = NewID(input.Index(3))
	return v
}

// NewRPSlice will construct a slice of type RP
func NewRPSlice(input hl7.HL7Type) []RP {
	values := make([]RP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRP(input.Index(i))
	}
	return values
}

type SRT struct {
	SortByField ST
	Sequencing  ID
}

// NewSRT creates an implementation of SRT
func NewSRT(input hl7.HL7Type) SRT {
	v := SRT{}
	v.SortByField = NewST(input.Index(0))
	v.Sequencing = NewID(input.Index(1))
	return v
}

// NewSRTSlice will construct a slice of type SRT
func NewSRTSlice(input hl7.HL7Type) []SRT {
	values := make([]SRT, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSRT(input.Index(i))
	}
	return values
}

type CX struct {
	IDNumber                    ST
	IdentifierCheckDigit        ST
	CheckDigitScheme            ID
	AssigningAuthority          HD
	IdentifierTypeCode          ID
	AssigningFacility           HD
	EffectiveDate               DT
	ExpirationDate              DT
	AssigningJurisdiction       CWE
	AssigningAgencyOrDepartment CWE
	SecurityCheck               ST
	SecurityCheckScheme         ID
}

// NewCX creates an implementation of CX
func NewCX(input hl7.HL7Type) CX {
	v := CX{}
	v.IDNumber = NewST(input.Index(0))
	v.IdentifierCheckDigit = NewST(input.Index(1))
	v.CheckDigitScheme = NewID(input.Index(2))
	v.AssigningAuthority = NewHD(input.Index(3))
	v.IdentifierTypeCode = NewID(input.Index(4))
	v.AssigningFacility = NewHD(input.Index(5))
	v.EffectiveDate = NewDT(input.Index(6))
	v.ExpirationDate = NewDT(input.Index(7))
	v.AssigningJurisdiction = NewCWE(input.Index(8))
	v.AssigningAgencyOrDepartment = NewCWE(input.Index(9))
	v.SecurityCheck = NewST(input.Index(10))
	v.SecurityCheckScheme = NewID(input.Index(11))
	return v
}

// NewCXSlice will construct a slice of type CX
func NewCXSlice(input hl7.HL7Type) []CX {
	values := make([]CX, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCX(input.Index(i))
	}
	return values
}

// Varies - Varies
type Varies string

// NewVaries creates an implementation of Varies
func NewVaries(input hl7.HL7Type) Varies {
	return Varies(input.String())
}

// NewVariesSlice will construct a slice of type Varies
func NewVariesSlice(input hl7.HL7Type) []Varies {
	values := make([]Varies, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewVaries(input.Index(i))
	}
	return values
}

type FC struct {
	FinancialClassCode CWE
	EffectiveDate      DTM
}

// NewFC creates an implementation of FC
func NewFC(input hl7.HL7Type) FC {
	v := FC{}
	v.FinancialClassCode = NewCWE(input.Index(0))
	v.EffectiveDate = NewDTM(input.Index(1))
	return v
}

// NewFCSlice will construct a slice of type FC
func NewFCSlice(input hl7.HL7Type) []FC {
	values := make([]FC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewFC(input.Index(i))
	}
	return values
}

type VR struct {
	FirstDataCodeValue ST
	LastDataCodeValue  ST
}

// NewVR creates an implementation of VR
func NewVR(input hl7.HL7Type) VR {
	v := VR{}
	v.FirstDataCodeValue = NewST(input.Index(0))
	v.LastDataCodeValue = NewST(input.Index(1))
	return v
}

// NewVRSlice will construct a slice of type VR
func NewVRSlice(input hl7.HL7Type) []VR {
	values := make([]VR, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewVR(input.Index(i))
	}
	return values
}

type ERL struct {
	SegmentID          ST
	SegmentSequence    NM
	FieldPosition      NM
	FieldRepetition    NM
	ComponentNumber    NM
	SubComponentNumber NM
}

// NewERL creates an implementation of ERL
func NewERL(input hl7.HL7Type) ERL {
	v := ERL{}
	v.SegmentID = NewST(input.Index(0))
	v.SegmentSequence = NewNM(input.Index(1))
	v.FieldPosition = NewNM(input.Index(2))
	v.FieldRepetition = NewNM(input.Index(3))
	v.ComponentNumber = NewNM(input.Index(4))
	v.SubComponentNumber = NewNM(input.Index(5))
	return v
}

// NewERLSlice will construct a slice of type ERL
func NewERLSlice(input hl7.HL7Type) []ERL {
	values := make([]ERL, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewERL(input.Index(i))
	}
	return values
}

type SPD struct {
	SpecialtyName       ST
	GoverningBoard      ST
	EligibleOrCertified ID
	DateOfCertification DT
}

// NewSPD creates an implementation of SPD
func NewSPD(input hl7.HL7Type) SPD {
	v := SPD{}
	v.SpecialtyName = NewST(input.Index(0))
	v.GoverningBoard = NewST(input.Index(1))
	v.EligibleOrCertified = NewID(input.Index(2))
	v.DateOfCertification = NewDT(input.Index(3))
	return v
}

// NewSPDSlice will construct a slice of type SPD
func NewSPDSlice(input hl7.HL7Type) []SPD {
	values := make([]SPD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSPD(input.Index(i))
	}
	return values
}

/*
NM - Numeric
A number represented as a series of ASCII numeric characters consisting of an optional leading sign (+ or -), the digits and an optional decimal point. In the absence of a sign, the number is assumed to be positive. If there is no decimal point the number is assumed to be an integer.

Values of this data type shall contain at least one digit to the left of the decimal point. This means that 0.1 is a valid representation, while .1 is not. Leading zeros, or trailing zeros after a decimal point, are not significant. For example, the following two values with different representations, "01.20" and "1.2," are identical. Except for the optional leading sign (+ or -) and the optional decimal point (.), no non-numeric ASCII characters are allowed. Thus, the value <12 should be encoded as a structured numeric (SN) (preferred) or as a string (ST) (allowed, but not preferred) data type.

The NM data type does not follow the normal truncation pattern, and the truncation character is never valid in the NM data type. Instead, the truncation behavior is based on the semantics of numbers.

Values of type NM may always have leading zeros truncated. Note that HL7 recommends that leading zeros not be used. Unless NM is used to represent a monetary amount, implementations may truncate trailing zeros after the decimal point up to the first non-zero digit or the decimal point, which ever comes first. Any digits to the left of the decimal point may never be truncated (other than leading zeros).

Example: 1.0200 may be truncated to 1.02, but not to 1.0.
*/
type NM string

// NewNM creates an implementation of NM
func NewNM(input hl7.HL7Type) NM {
	return NM(input.String())
}

// NewNMSlice will construct a slice of type NM
func NewNMSlice(input hl7.HL7Type) []NM {
	values := make([]NM, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewNM(input.Index(i))
	}
	return values
}

type PPN struct {
	PersonIdentifier                            ST
	FamilyName                                  FN
	GivenName                                   ST
	SecondAndFurtherGivenNamesOrInitialsThereof ST
	SuffixeGJrOrIii                             ST
	PrefixeGDr                                  ST
	DegreeeGMd                                  ST
	SourceTable                                 CWE
	AssigningAuthority                          HD
	NameTypeCode                                ID
	IdentifierCheckDigit                        ST
	CheckDigitScheme                            ID
	IdentifierTypeCode                          ID
	AssigningFacility                           HD
	DateTimeActionPerformed                     DTM
	NameRepresentationCode                      ID
	NameContext                                 CWE
	NameValidityRange                           ST
	NameAssemblyOrder                           ID
	EffectiveDate                               DTM
	ExpirationDate                              DTM
	ProfessionalSuffix                          ST
	AssigningJurisdiction                       CWE
	AssigningAgencyOrDepartment                 CWE
	SecurityCheck                               ST
	SecurityCheckScheme                         ID
}

// NewPPN creates an implementation of PPN
func NewPPN(input hl7.HL7Type) PPN {
	v := PPN{}
	v.PersonIdentifier = NewST(input.Index(0))
	v.FamilyName = NewFN(input.Index(1))
	v.GivenName = NewST(input.Index(2))
	v.SecondAndFurtherGivenNamesOrInitialsThereof = NewST(input.Index(3))
	v.SuffixeGJrOrIii = NewST(input.Index(4))
	v.PrefixeGDr = NewST(input.Index(5))
	v.DegreeeGMd = NewST(input.Index(6))
	v.SourceTable = NewCWE(input.Index(7))
	v.AssigningAuthority = NewHD(input.Index(8))
	v.NameTypeCode = NewID(input.Index(9))
	v.IdentifierCheckDigit = NewST(input.Index(10))
	v.CheckDigitScheme = NewID(input.Index(11))
	v.IdentifierTypeCode = NewID(input.Index(12))
	v.AssigningFacility = NewHD(input.Index(13))
	v.DateTimeActionPerformed = NewDTM(input.Index(14))
	v.NameRepresentationCode = NewID(input.Index(15))
	v.NameContext = NewCWE(input.Index(16))
	v.NameValidityRange = NewST(input.Index(17))
	v.NameAssemblyOrder = NewID(input.Index(18))
	v.EffectiveDate = NewDTM(input.Index(19))
	v.ExpirationDate = NewDTM(input.Index(20))
	v.ProfessionalSuffix = NewST(input.Index(21))
	v.AssigningJurisdiction = NewCWE(input.Index(22))
	v.AssigningAgencyOrDepartment = NewCWE(input.Index(23))
	v.SecurityCheck = NewST(input.Index(24))
	v.SecurityCheckScheme = NewID(input.Index(25))
	return v
}

// NewPPNSlice will construct a slice of type PPN
func NewPPNSlice(input hl7.HL7Type) []PPN {
	values := make([]PPN, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPPN(input.Index(i))
	}
	return values
}

type CSU struct {
	ChannelSensitivity                          NM
	UnitOfMeasureIdentifier                     ST
	UnitOfMeasureDescription                    ST
	UnitOfMeasureCodingSystem                   ID
	AlternateUnitOfMeasureIdentifier            ST
	AlternateUnitOfMeasureDescription           ST
	AlternateUnitOfMeasureCodingSystem          ID
	UnitOfMeasureCodingSystemVersionID          ST
	AlternateUnitOfMeasureCodingSystemVersionID ST
	OriginalText                                ST
	SecondAlternateUnitOfMeasureIdentifier      ST
	SecondAlternateUnitOfMeasureText            ST
	NameOfSecondAlternateUnitOfMeasureCodingSy  ID
	SecondAlternateUnitOfMeasureCodingSystemVer ST
	UnitOfMeasureCodingSystemOID                ST
	UnitOfMeasureValueSetOID                    ST
	UnitOfMeasureValueSetVersionID              DTM
	SecondUnitOfMeasureCodingSystemOID          ST
	SecondUnitOfMeasureValueSetOID              ST
	SecondUnitOfMeasureValueSetVersionID        DTM
	ThirdUnitOfMeasureCodingSystemOID           ST
	ThirdUnitOfMeasureValueSetOID               ST
	ThirdUnitOfMeasureValueSetVersionID         ST
}

// NewCSU creates an implementation of CSU
func NewCSU(input hl7.HL7Type) CSU {
	v := CSU{}
	v.ChannelSensitivity = NewNM(input.Index(0))
	v.UnitOfMeasureIdentifier = NewST(input.Index(1))
	v.UnitOfMeasureDescription = NewST(input.Index(2))
	v.UnitOfMeasureCodingSystem = NewID(input.Index(3))
	v.AlternateUnitOfMeasureIdentifier = NewST(input.Index(4))
	v.AlternateUnitOfMeasureDescription = NewST(input.Index(5))
	v.AlternateUnitOfMeasureCodingSystem = NewID(input.Index(6))
	v.UnitOfMeasureCodingSystemVersionID = NewST(input.Index(7))
	v.AlternateUnitOfMeasureCodingSystemVersionID = NewST(input.Index(8))
	v.OriginalText = NewST(input.Index(9))
	v.SecondAlternateUnitOfMeasureIdentifier = NewST(input.Index(10))
	v.SecondAlternateUnitOfMeasureText = NewST(input.Index(11))
	v.NameOfSecondAlternateUnitOfMeasureCodingSy = NewID(input.Index(12))
	v.SecondAlternateUnitOfMeasureCodingSystemVer = NewST(input.Index(13))
	v.UnitOfMeasureCodingSystemOID = NewST(input.Index(14))
	v.UnitOfMeasureValueSetOID = NewST(input.Index(15))
	v.UnitOfMeasureValueSetVersionID = NewDTM(input.Index(16))
	v.SecondUnitOfMeasureCodingSystemOID = NewST(input.Index(17))
	v.SecondUnitOfMeasureValueSetOID = NewST(input.Index(18))
	v.SecondUnitOfMeasureValueSetVersionID = NewDTM(input.Index(19))
	v.ThirdUnitOfMeasureCodingSystemOID = NewST(input.Index(20))
	v.ThirdUnitOfMeasureValueSetOID = NewST(input.Index(21))
	v.ThirdUnitOfMeasureValueSetVersionID = NewST(input.Index(22))
	return v
}

// NewCSUSlice will construct a slice of type CSU
func NewCSUSlice(input hl7.HL7Type) []CSU {
	values := make([]CSU, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCSU(input.Index(i))
	}
	return values
}

/*
DT - Date
Specifies the century and year with optional precision to month and day

The number of digits populated specifies the precision using the format specification YYYY[MM[DD]]

Examples:
|19880704|
|199503|

The DT data type does not follow the normal truncation pattern, and the truncation character is never valid in the DT data type. Instead, the truncation behavior is based on the semantics of dates.

Unless specified in the context where the DT type  is used, the DT type may not be truncated. When a DT is truncated, the truncated form SHALL still be a valid DT type. Systems should always be able to persist full dates. Refer to Chapter 2, section 2.5.5.2 "Truncation Pattern" for further information.

Note: Prior to v2.3, this data type was specified in the format YYYYMMDD. As of v2.3, month and days are no longer required. By site-specific agreement, YYYYMMDD may be used where backward compatibility must be maintained.
*/
type DT string

// NewDT creates an implementation of DT
func NewDT(input hl7.HL7Type) DT {
	return DT(input.String())
}

// NewDTSlice will construct a slice of type DT
func NewDTSlice(input hl7.HL7Type) []DT {
	values := make([]DT, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDT(input.Index(i))
	}
	return values
}

type RPT struct {
	RepeatPatternCode          CWE
	CalendarAlignment          ID
	PhaseRangeBeginValue       NM
	PhaseRangeEndValue         NM
	PeriodQuantity             NM
	PeriodUnits                CWE
	InstitutionSpecifiedTime   ID
	Event                      ID
	EventOffsetQuantity        NM
	EventOffsetUnits           CWE
	GeneralTimingSpecification GTS
}

// NewRPT creates an implementation of RPT
func NewRPT(input hl7.HL7Type) RPT {
	v := RPT{}
	v.RepeatPatternCode = NewCWE(input.Index(0))
	v.CalendarAlignment = NewID(input.Index(1))
	v.PhaseRangeBeginValue = NewNM(input.Index(2))
	v.PhaseRangeEndValue = NewNM(input.Index(3))
	v.PeriodQuantity = NewNM(input.Index(4))
	v.PeriodUnits = NewCWE(input.Index(5))
	v.InstitutionSpecifiedTime = NewID(input.Index(6))
	v.Event = NewID(input.Index(7))
	v.EventOffsetQuantity = NewNM(input.Index(8))
	v.EventOffsetUnits = NewCWE(input.Index(9))
	v.GeneralTimingSpecification = NewGTS(input.Index(10))
	return v
}

// NewRPTSlice will construct a slice of type RPT
func NewRPTSlice(input hl7.HL7Type) []RPT {
	values := make([]RPT, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRPT(input.Index(i))
	}
	return values
}

/*
DTM - Date/time
Specifies a point in time using a 24-hour clock notation.

The number of characters populated (excluding the time zone specification) specifies the precision.

Format: YYYY[MM[DD[HH[MM[SS[.S[S[S[S]]]]]]]]][+/-ZZZZ].

The time zone (+/-ZZZZ) is represented as +/-HHMM offset from Coordinated Universal Time (UTC) (formerly Greenwich Mean Time (GMT)), where +0000 or -0000 both represent UTC (without offset). The specific data representations used in the HL7 encoding rules are compatible with ISO 8824-1987(E).
Note that if the time zone is not included, the time zone defaults to that of the local time zone of the sender. Also note that a DTM valued field with the HHMM part set to "0000" represents midnight of the night extending from the previous day to the day given by the YYYYMMDD part.

The HL7 Standard strongly recommends that all systems routinely send the time zone offset but does not require it. All HL7 systems are required to accept the time zone offset, but its implementation is application specific. For many applications the time of interest is the local time of the sender. For example, an application in the Eastern Standard Time zone receiving notification of an admission that takes place at 11:00 PM in San Francisco on December 11 would prefer to treat the admission as having occurred on December 11 rather than advancing the date to December 12.

Note: The time zone [+/-ZZZZ], when used, is restricted to legally-defined time zones and is represented in HHMM format.

One exception to this rule would be a clinical system that processed patient data collected in a clinic and a nearby hospital that happens to be in a different time zone. Such applications may choose to convert the data to a common representation. Similar concerns apply to the transitions to and from daylight saving time. HL7 supports such requirements by requiring that the time zone information be present when the information is sent. It does not, however, specify which of the treatments discussed here will be applied by the receiving system.

The DTM data type does not follow the normal truncation pattern, and the truncation character is never valid in the DTM data type. Instead, the truncation behavior is based on the semantics of dates and times.

Unless otherwise specified in the context where the DTM type is used, the DTM type may be truncated to a day. When a DTM is truncated, the truncated form SHALL still be a valid DTM type. Systems should always be able to persist full date / time information including the timezone. Refer to Chapter 2, section 2.5.5.2 "Truncation Pattern" for further information.
*/
type DTM string

// NewDTM creates an implementation of DTM
func NewDTM(input hl7.HL7Type) DTM {
	return DTM(input.String())
}

// NewDTMSlice will construct a slice of type DTM
func NewDTMSlice(input hl7.HL7Type) []DTM {
	values := make([]DTM, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDTM(input.Index(i))
	}
	return values
}

type PTA struct {
	PolicyType                CWE
	AmountClass               CWE
	MoneyOrPercentageQuantity ST
	MoneyOrPercentage         MOP
}

// NewPTA creates an implementation of PTA
func NewPTA(input hl7.HL7Type) PTA {
	v := PTA{}
	v.PolicyType = NewCWE(input.Index(0))
	v.AmountClass = NewCWE(input.Index(1))
	v.MoneyOrPercentageQuantity = NewST(input.Index(2))
	v.MoneyOrPercentage = NewMOP(input.Index(3))
	return v
}

// NewPTASlice will construct a slice of type PTA
func NewPTASlice(input hl7.HL7Type) []PTA {
	values := make([]PTA, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPTA(input.Index(i))
	}
	return values
}

type CCP struct {
	ChannelCalibrationSensitivityCorrectionFactor NM
	ChannelCalibrationBaseline                    NM
	ChannelCalibrationTimeSkew                    NM
}

// NewCCP creates an implementation of CCP
func NewCCP(input hl7.HL7Type) CCP {
	v := CCP{}
	v.ChannelCalibrationSensitivityCorrectionFactor = NewNM(input.Index(0))
	v.ChannelCalibrationBaseline = NewNM(input.Index(1))
	v.ChannelCalibrationTimeSkew = NewNM(input.Index(2))
	return v
}

// NewCCPSlice will construct a slice of type CCP
func NewCCPSlice(input hl7.HL7Type) []CCP {
	values := make([]CCP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCCP(input.Index(i))
	}
	return values
}

type UVC struct {
	ValueCode                      CWE
	ValueAmount                    MO
	NonMonetaryValueAmountQuantity NM
	NonMonetaryValueAmountUnits    CWE
}

// NewUVC creates an implementation of UVC
func NewUVC(input hl7.HL7Type) UVC {
	v := UVC{}
	v.ValueCode = NewCWE(input.Index(0))
	v.ValueAmount = NewMO(input.Index(1))
	v.NonMonetaryValueAmountQuantity = NewNM(input.Index(2))
	v.NonMonetaryValueAmountUnits = NewCWE(input.Index(3))
	return v
}

// NewUVCSlice will construct a slice of type UVC
func NewUVCSlice(input hl7.HL7Type) []UVC {
	values := make([]UVC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewUVC(input.Index(i))
	}
	return values
}

/*
ID - Coded Value For Hl7 Defined Tables
The value of such a field follows the formatting rules for an ST field except that it is drawn from a table of legal values. There shall be an HL7 table number associated with ID data types. An example of an ID field is OBR-25 Result Status. This data type should be used only for HL7 tables (see Chapter 2C, section 2.C.1.2, "HL7 Tables"). The reverse is not true, since in some circumstances it is more appropriate to use the CNE or CWE data type for HL7 tables.

The minimum and maximum lengths are specified in the context in which the ID data type is used. The longest HL7 defined legal value is 15 characters, but there are a few circumstances where the legal values are taken from code systems defined by other bodies (such as IANA mime types). In these cases, a different conformance length may be specified where the ID data type is used. It is never acceptable to truncate an ID value.
*/
type ID string

// NewID creates an implementation of ID
func NewID(input hl7.HL7Type) ID {
	return ID(input.String())
}

// NewIDSlice will construct a slice of type ID
func NewIDSlice(input hl7.HL7Type) []ID {
	values := make([]ID, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewID(input.Index(i))
	}
	return values
}

type HD struct {
	NamespaceID     IS
	UniversalID     ST
	UniversalIDType ID
}

// NewHD creates an implementation of HD
func NewHD(input hl7.HL7Type) HD {
	v := HD{}
	v.NamespaceID = NewIS(input.Index(0))
	v.UniversalID = NewST(input.Index(1))
	v.UniversalIDType = NewID(input.Index(2))
	return v
}

// NewHDSlice will construct a slice of type HD
func NewHDSlice(input hl7.HL7Type) []HD {
	values := make([]HD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewHD(input.Index(i))
	}
	return values
}

type QIP struct {
	SegmentFieldName ST
	Values           ST
}

// NewQIP creates an implementation of QIP
func NewQIP(input hl7.HL7Type) QIP {
	v := QIP{}
	v.SegmentFieldName = NewST(input.Index(0))
	v.Values = NewST(input.Index(1))
	return v
}

// NewQIPSlice will construct a slice of type QIP
func NewQIPSlice(input hl7.HL7Type) []QIP {
	values := make([]QIP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQIP(input.Index(i))
	}
	return values
}

type SCV struct {
	ParameterClass CWE
	ParameterValue ST
}

// NewSCV creates an implementation of SCV
func NewSCV(input hl7.HL7Type) SCV {
	v := SCV{}
	v.ParameterClass = NewCWE(input.Index(0))
	v.ParameterValue = NewST(input.Index(1))
	return v
}

// NewSCVSlice will construct a slice of type SCV
func NewSCVSlice(input hl7.HL7Type) []SCV {
	values := make([]SCV, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSCV(input.Index(i))
	}
	return values
}

type PLN struct {
	IDNumber                        ST
	TypeOfIDNumber                  CWE
	StateOtherQualifyingInformation ST
	ExpirationDate                  DT
}

// NewPLN creates an implementation of PLN
func NewPLN(input hl7.HL7Type) PLN {
	v := PLN{}
	v.IDNumber = NewST(input.Index(0))
	v.TypeOfIDNumber = NewCWE(input.Index(1))
	v.StateOtherQualifyingInformation = NewST(input.Index(2))
	v.ExpirationDate = NewDT(input.Index(3))
	return v
}

// NewPLNSlice will construct a slice of type PLN
func NewPLNSlice(input hl7.HL7Type) []PLN {
	values := make([]PLN, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPLN(input.Index(i))
	}
	return values
}

type JCC struct {
	JobCode            CWE
	JobClass           CWE
	JobDescriptionText TX
}

// NewJCC creates an implementation of JCC
func NewJCC(input hl7.HL7Type) JCC {
	v := JCC{}
	v.JobCode = NewCWE(input.Index(0))
	v.JobClass = NewCWE(input.Index(1))
	v.JobDescriptionText = NewTX(input.Index(2))
	return v
}

// NewJCCSlice will construct a slice of type JCC
func NewJCCSlice(input hl7.HL7Type) []JCC {
	values := make([]JCC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewJCC(input.Index(i))
	}
	return values
}

/*
FT - Formatted Text Data
This data type is derived from the TX data type by allowing the addition of embedded formatting instructions. These instructions are limited to those that are intrinsic and independent of the circumstances under which the field is being used. The actual instructions and their representation are described in section 2.7.6, “Usage and Examples of Formatted Text”. The FT field is of arbitrary length (up to 64k) and may contain formatting commands enclosed in escape characters.

Example:
|\.sp\(skip one vertical line)|

For additional examples of formatting commands see Section 2.7, "Use of Escape Sequences in Text Fields".

To include alternative character sets, use the appropriate escape sequence. See Chapter 2, "Control", section 2.15.9.18, "Character set" and section 2.14.9.20, "Alternate character set handling scheme".

This specification applies no limit to the length of the FT data type, either here where the data type is defined, or elsewhere where the data type is used. While there is no intrinsic reason to limit the length of this data type for semantic or syntactical reasons, it is expected that some sort of limitation will be imposed for technical reasons in implementations. HL7 recommends that implementation length limits are published in implementation profiles. The contents of an FT field may be truncated, but the truncation pattern does not apply.
*/
type FT string

// NewFT creates an implementation of FT
func NewFT(input hl7.HL7Type) FT {
	return FT(input.String())
}

// NewFTSlice will construct a slice of type FT
func NewFTSlice(input hl7.HL7Type) []FT {
	values := make([]FT, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewFT(input.Index(i))
	}
	return values
}

type ICD struct {
	CertificationPatientType      CWE
	CertificationRequired         ID
	DateTimeCertificationRequired DTM
}

// NewICD creates an implementation of ICD
func NewICD(input hl7.HL7Type) ICD {
	v := ICD{}
	v.CertificationPatientType = NewCWE(input.Index(0))
	v.CertificationRequired = NewID(input.Index(1))
	v.DateTimeCertificationRequired = NewDTM(input.Index(2))
	return v
}

// NewICDSlice will construct a slice of type ICD
func NewICDSlice(input hl7.HL7Type) []ICD {
	values := make([]ICD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewICD(input.Index(i))
	}
	return values
}

type DLD struct {
	DischargeToLocation CWE
	EffectiveDate       DTM
}

// NewDLD creates an implementation of DLD
func NewDLD(input hl7.HL7Type) DLD {
	v := DLD{}
	v.DischargeToLocation = NewCWE(input.Index(0))
	v.EffectiveDate = NewDTM(input.Index(1))
	return v
}

// NewDLDSlice will construct a slice of type DLD
func NewDLDSlice(input hl7.HL7Type) []DLD {
	values := make([]DLD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDLD(input.Index(i))
	}
	return values
}

type XON struct {
	OrganizationName         ST
	OrganizationNameTypeCode CWE
	IDNumber                 ST
	IdentifierCheckDigit     ST
	CheckDigitScheme         ST
	AssigningAuthority       HD
	IdentifierTypeCode       ID
	AssigningFacility        HD
	NameRepresentationCode   ID
	OrganizationIdentifier   ST
}

// NewXON creates an implementation of XON
func NewXON(input hl7.HL7Type) XON {
	v := XON{}
	v.OrganizationName = NewST(input.Index(0))
	v.OrganizationNameTypeCode = NewCWE(input.Index(1))
	v.IDNumber = NewST(input.Index(2))
	v.IdentifierCheckDigit = NewST(input.Index(3))
	v.CheckDigitScheme = NewST(input.Index(4))
	v.AssigningAuthority = NewHD(input.Index(5))
	v.IdentifierTypeCode = NewID(input.Index(6))
	v.AssigningFacility = NewHD(input.Index(7))
	v.NameRepresentationCode = NewID(input.Index(8))
	v.OrganizationIdentifier = NewST(input.Index(9))
	return v
}

// NewXONSlice will construct a slice of type XON
func NewXONSlice(input hl7.HL7Type) []XON {
	values := make([]XON, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewXON(input.Index(i))
	}
	return values
}

type MO struct {
	Quantity     NM
	Denomination ID
}

// NewMO creates an implementation of MO
func NewMO(input hl7.HL7Type) MO {
	v := MO{}
	v.Quantity = NewNM(input.Index(0))
	v.Denomination = NewID(input.Index(1))
	return v
}

// NewMOSlice will construct a slice of type MO
func NewMOSlice(input hl7.HL7Type) []MO {
	values := make([]MO, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMO(input.Index(i))
	}
	return values
}

type XAD struct {
	StreetAddress              SAD
	OtherDesignation           ST
	City                       ST
	StateOrProvince            ST
	ZipOrPostalCode            ST
	Country                    ID
	AddressType                ID
	OtherGeographicDesignation ST
	CountyParishCode           CWE
	CensusTract                CWE
	AddressRepresentationCode  ID
	AddressValidityRange       ST
	EffectiveDate              DTM
	ExpirationDate             DTM
	ExpirationReason           CWE
	TemporaryIndicator         ID
	BadAddressIndicator        ID
	AddressUsage               ID
	Addressee                  ST
	Comment                    ST
	PreferenceOrder            NM
	ProtectionCode             CWE
	AddressIdentifier          EI
}

// NewXAD creates an implementation of XAD
func NewXAD(input hl7.HL7Type) XAD {
	v := XAD{}
	v.StreetAddress = NewSAD(input.Index(0))
	v.OtherDesignation = NewST(input.Index(1))
	v.City = NewST(input.Index(2))
	v.StateOrProvince = NewST(input.Index(3))
	v.ZipOrPostalCode = NewST(input.Index(4))
	v.Country = NewID(input.Index(5))
	v.AddressType = NewID(input.Index(6))
	v.OtherGeographicDesignation = NewST(input.Index(7))
	v.CountyParishCode = NewCWE(input.Index(8))
	v.CensusTract = NewCWE(input.Index(9))
	v.AddressRepresentationCode = NewID(input.Index(10))
	v.AddressValidityRange = NewST(input.Index(11))
	v.EffectiveDate = NewDTM(input.Index(12))
	v.ExpirationDate = NewDTM(input.Index(13))
	v.ExpirationReason = NewCWE(input.Index(14))
	v.TemporaryIndicator = NewID(input.Index(15))
	v.BadAddressIndicator = NewID(input.Index(16))
	v.AddressUsage = NewID(input.Index(17))
	v.Addressee = NewST(input.Index(18))
	v.Comment = NewST(input.Index(19))
	v.PreferenceOrder = NewNM(input.Index(20))
	v.ProtectionCode = NewCWE(input.Index(21))
	v.AddressIdentifier = NewEI(input.Index(22))
	return v
}

// NewXADSlice will construct a slice of type XAD
func NewXADSlice(input hl7.HL7Type) []XAD {
	values := make([]XAD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewXAD(input.Index(i))
	}
	return values
}

type DLN struct {
	LicenseNumber               ST
	IssuingStateProvinceCountry CWE
	ExpirationDate              DT
}

// NewDLN creates an implementation of DLN
func NewDLN(input hl7.HL7Type) DLN {
	v := DLN{}
	v.LicenseNumber = NewST(input.Index(0))
	v.IssuingStateProvinceCountry = NewCWE(input.Index(1))
	v.ExpirationDate = NewDT(input.Index(2))
	return v
}

// NewDLNSlice will construct a slice of type DLN
func NewDLNSlice(input hl7.HL7Type) []DLN {
	values := make([]DLN, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDLN(input.Index(i))
	}
	return values
}

type PL struct {
	PointOfCare                     HD
	Room                            HD
	Bed                             HD
	Facility                        HD
	LocationStatus                  IS
	PersonLocationType              IS
	Building                        HD
	Floor                           HD
	LocationDescription             ST
	ComprehensiveLocationIdentifier EI
	AssigningAuthorityForLocation   HD
}

// NewPL creates an implementation of PL
func NewPL(input hl7.HL7Type) PL {
	v := PL{}
	v.PointOfCare = NewHD(input.Index(0))
	v.Room = NewHD(input.Index(1))
	v.Bed = NewHD(input.Index(2))
	v.Facility = NewHD(input.Index(3))
	v.LocationStatus = NewIS(input.Index(4))
	v.PersonLocationType = NewIS(input.Index(5))
	v.Building = NewHD(input.Index(6))
	v.Floor = NewHD(input.Index(7))
	v.LocationDescription = NewST(input.Index(8))
	v.ComprehensiveLocationIdentifier = NewEI(input.Index(9))
	v.AssigningAuthorityForLocation = NewHD(input.Index(10))
	return v
}

// NewPLSlice will construct a slice of type PL
func NewPLSlice(input hl7.HL7Type) []PL {
	values := make([]PL, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPL(input.Index(i))
	}
	return values
}

type DIN struct {
	Date            DTM
	InstitutionName CWE
}

// NewDIN creates an implementation of DIN
func NewDIN(input hl7.HL7Type) DIN {
	v := DIN{}
	v.Date = NewDTM(input.Index(0))
	v.InstitutionName = NewCWE(input.Index(1))
	return v
}

// NewDINSlice will construct a slice of type DIN
func NewDINSlice(input hl7.HL7Type) []DIN {
	values := make([]DIN, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDIN(input.Index(i))
	}
	return values
}

/*
SI - Sequence Id
A non-negative integer in the form of a NM field. The uses of this data type are defined in the chapters defining the segments and messages in which it appears.

This allows for a number between 0 and 9999 to be specified.
*/
type SI string

// NewSI creates an implementation of SI
func NewSI(input hl7.HL7Type) SI {
	return SI(input.String())
}

// NewSISlice will construct a slice of type SI
func NewSISlice(input hl7.HL7Type) []SI {
	values := make([]SI, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewSI(input.Index(i))
	}
	return values
}

type QSC struct {
	SegmentFieldName      ST
	RelationalOperator    ID
	Value                 ST
	RelationalConjunction ID
}

// NewQSC creates an implementation of QSC
func NewQSC(input hl7.HL7Type) QSC {
	v := QSC{}
	v.SegmentFieldName = NewST(input.Index(0))
	v.RelationalOperator = NewID(input.Index(1))
	v.Value = NewST(input.Index(2))
	v.RelationalConjunction = NewID(input.Index(3))
	return v
}

// NewQSCSlice will construct a slice of type QSC
func NewQSCSlice(input hl7.HL7Type) []QSC {
	values := make([]QSC, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewQSC(input.Index(i))
	}
	return values
}

type CNE struct {
	Identifier                           ST
	Text                                 ST
	NameOfCodingSystem                   ID
	AlternateIdentifier                  ST
	AlternateText                        ST
	NameOfAlternateCodingSystem          ID
	CodingSystemVersionID                ST
	AlternateCodingSystemVersionID       ST
	OriginalText                         ST
	SecondAlternateIdentifier            ST
	SecondAlternateText                  ST
	NameOfSecondAlternateCodingSystem    ID
	SecondAlternateCodingSystemVersionID ST
	CodingSystemOID                      ST
	ValueSetOID                          ST
	ValueSetVersionID                    DTM
	AlternateCodingSystemOID             ST
	AlternateValueSetOID                 ST
	AlternateValueSetVersionID           DTM
	SecondAlternateCodingSystemOID       ST
	SecondAlternateValueSetOID           ST
	SecondAlternateValueSetVersionID     DTM
}

// NewCNE creates an implementation of CNE
func NewCNE(input hl7.HL7Type) CNE {
	v := CNE{}
	v.Identifier = NewST(input.Index(0))
	v.Text = NewST(input.Index(1))
	v.NameOfCodingSystem = NewID(input.Index(2))
	v.AlternateIdentifier = NewST(input.Index(3))
	v.AlternateText = NewST(input.Index(4))
	v.NameOfAlternateCodingSystem = NewID(input.Index(5))
	v.CodingSystemVersionID = NewST(input.Index(6))
	v.AlternateCodingSystemVersionID = NewST(input.Index(7))
	v.OriginalText = NewST(input.Index(8))
	v.SecondAlternateIdentifier = NewST(input.Index(9))
	v.SecondAlternateText = NewST(input.Index(10))
	v.NameOfSecondAlternateCodingSystem = NewID(input.Index(11))
	v.SecondAlternateCodingSystemVersionID = NewST(input.Index(12))
	v.CodingSystemOID = NewST(input.Index(13))
	v.ValueSetOID = NewST(input.Index(14))
	v.ValueSetVersionID = NewDTM(input.Index(15))
	v.AlternateCodingSystemOID = NewST(input.Index(16))
	v.AlternateValueSetOID = NewST(input.Index(17))
	v.AlternateValueSetVersionID = NewDTM(input.Index(18))
	v.SecondAlternateCodingSystemOID = NewST(input.Index(19))
	v.SecondAlternateValueSetOID = NewST(input.Index(20))
	v.SecondAlternateValueSetVersionID = NewDTM(input.Index(21))
	return v
}

// NewCNESlice will construct a slice of type CNE
func NewCNESlice(input hl7.HL7Type) []CNE {
	values := make([]CNE, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCNE(input.Index(i))
	}
	return values
}

type MA struct {
	SampleYFromChannel1 NM
	SampleYFromChannel2 NM
	SampleYFromChannel3 NM
	SampleYFromChannel4 NM
}

// NewMA creates an implementation of MA
func NewMA(input hl7.HL7Type) MA {
	v := MA{}
	v.SampleYFromChannel1 = NewNM(input.Index(0))
	v.SampleYFromChannel2 = NewNM(input.Index(1))
	v.SampleYFromChannel3 = NewNM(input.Index(2))
	v.SampleYFromChannel4 = NewNM(input.Index(3))
	return v
}

// NewMASlice will construct a slice of type MA
func NewMASlice(input hl7.HL7Type) []MA {
	values := make([]MA, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMA(input.Index(i))
	}
	return values
}

type FN struct {
	Surname                        ST
	OwnSurnamePrefix               ST
	OwnSurname                     ST
	SurnamePrefixFromPartnerSpouse ST
	SurnameFromPartnerSpouse       ST
}

// NewFN creates an implementation of FN
func NewFN(input hl7.HL7Type) FN {
	v := FN{}
	v.Surname = NewST(input.Index(0))
	v.OwnSurnamePrefix = NewST(input.Index(1))
	v.OwnSurname = NewST(input.Index(2))
	v.SurnamePrefixFromPartnerSpouse = NewST(input.Index(3))
	v.SurnameFromPartnerSpouse = NewST(input.Index(4))
	return v
}

// NewFNSlice will construct a slice of type FN
func NewFNSlice(input hl7.HL7Type) []FN {
	values := make([]FN, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewFN(input.Index(i))
	}
	return values
}

/*
GTS - General Timing Specification
The General Timing Specification data type is used to communicate complex inter-related information Timing information. The value of such a field follows the formatting rules for a ST field. The string data will be structured according to the rules set forth in the "Version 3 Data Types Part II Unabridged Specification" for the General Timing Specification (GTS) data type.

There is no technical limit to the length of a GTS expression – the expression may be as long as logically required. The conformance length of 199 caters for all the common expressions. GTS expressions are not to be truncated.
*/
type GTS string

// NewGTS creates an implementation of GTS
func NewGTS(input hl7.HL7Type) GTS {
	return GTS(input.String())
}

// NewGTSSlice will construct a slice of type GTS
func NewGTSSlice(input hl7.HL7Type) []GTS {
	values := make([]GTS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewGTS(input.Index(i))
	}
	return values
}

type PIP struct {
	Privilege      CWE
	PrivilegeClass CWE
	ExpirationDate DT
	ActivationDate DT
	Facility       EI
}

// NewPIP creates an implementation of PIP
func NewPIP(input hl7.HL7Type) PIP {
	v := PIP{}
	v.Privilege = NewCWE(input.Index(0))
	v.PrivilegeClass = NewCWE(input.Index(1))
	v.ExpirationDate = NewDT(input.Index(2))
	v.ActivationDate = NewDT(input.Index(3))
	v.Facility = NewEI(input.Index(4))
	return v
}

// NewPIPSlice will construct a slice of type PIP
func NewPIPSlice(input hl7.HL7Type) []PIP {
	values := make([]PIP, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPIP(input.Index(i))
	}
	return values
}

type CNN struct {
	IDNumber                                    ST
	FamilyName                                  ST
	GivenName                                   ST
	SecondAndFurtherGivenNamesOrInitialsThereof ST
	SuffixeGJrOrIii                             ST
	PrefixeGDr                                  ST
	DegreeeGMd                                  IS
	SourceTable                                 IS
	AssigningAuthorityNamespaceID               IS
	AssigningAuthorityUniversalID               ST
	AssigningAuthorityUniversalIDType           ID
}

// NewCNN creates an implementation of CNN
func NewCNN(input hl7.HL7Type) CNN {
	v := CNN{}
	v.IDNumber = NewST(input.Index(0))
	v.FamilyName = NewST(input.Index(1))
	v.GivenName = NewST(input.Index(2))
	v.SecondAndFurtherGivenNamesOrInitialsThereof = NewST(input.Index(3))
	v.SuffixeGJrOrIii = NewST(input.Index(4))
	v.PrefixeGDr = NewST(input.Index(5))
	v.DegreeeGMd = NewIS(input.Index(6))
	v.SourceTable = NewIS(input.Index(7))
	v.AssigningAuthorityNamespaceID = NewIS(input.Index(8))
	v.AssigningAuthorityUniversalID = NewST(input.Index(9))
	v.AssigningAuthorityUniversalIDType = NewID(input.Index(10))
	return v
}

// NewCNNSlice will construct a slice of type CNN
func NewCNNSlice(input hl7.HL7Type) []CNN {
	values := make([]CNN, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCNN(input.Index(i))
	}
	return values
}

/*
IS - Coded Value For User-defined Tables
As of v2.7, the only approved use of the IS data type is in the HD.1, EI.2 and PL.6 plus a limited number of fields where a determination could not readily be made as to whether the item is an identifier or an actual coded item. Additionally, in accordance with chapter 2 rules, any field or data type component marked as "Retained for backward compatibility" will retain any IS data type.

The value of such a field follows the formatting rules for a ST field except that it is drawn from a site-defined (or user-defined) table of legal values. There shall be an HL7 table number associated with IS data types. An example of an IS field is the Event reason code defined in Chapter 3, "Patient Administration", section 3.4.1.4, "Event reason code". This data type should be used only for user-defined tables (see Chapter 2C, "Code Tables", section 2.C.1.1, "User-defined Tables"). The reverse is not true, since in some circumstances, it is more appropriate to use the CWE data type for user-defined tables.

It is never acceptable to truncate an IS value.
*/
type IS string

// NewIS creates an implementation of IS
func NewIS(input hl7.HL7Type) IS {
	return IS(input.String())
}

// NewISSlice will construct a slice of type IS
func NewISSlice(input hl7.HL7Type) []IS {
	values := make([]IS, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewIS(input.Index(i))
	}
	return values
}

type VH struct {
	StartDayRange  ID
	EndDayRange    ID
	StartHourRange TM
	EndHourRange   TM
}

// NewVH creates an implementation of VH
func NewVH(input hl7.HL7Type) VH {
	v := VH{}
	v.StartDayRange = NewID(input.Index(0))
	v.EndDayRange = NewID(input.Index(1))
	v.StartHourRange = NewTM(input.Index(2))
	v.EndHourRange = NewTM(input.Index(3))
	return v
}

// NewVHSlice will construct a slice of type VH
func NewVHSlice(input hl7.HL7Type) []VH {
	values := make([]VH, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewVH(input.Index(i))
	}
	return values
}

type MSG struct {
	MessageCode      ID
	TriggerEvent     ID
	MessageStructure ID
}

// NewMSG creates an implementation of MSG
func NewMSG(input hl7.HL7Type) MSG {
	v := MSG{}
	v.MessageCode = NewID(input.Index(0))
	v.TriggerEvent = NewID(input.Index(1))
	v.MessageStructure = NewID(input.Index(2))
	return v
}

// NewMSGSlice will construct a slice of type MSG
func NewMSGSlice(input hl7.HL7Type) []MSG {
	values := make([]MSG, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewMSG(input.Index(i))
	}
	return values
}

/*
TX - Text Data
String data meant for user display (on a terminal or printer). Such data would not necessarily be left justified since leading spaces may contribute greatly to the clarity of the presentation to the user. Because this type of data is intended for display, it may contain certain escape character sequences designed to control the display. Escape sequence formatting is defined in Section 2.7, "Use of escape sequences in text fields". Leading spaces should be included. Trailing spaces should be removed.

Example:
|  leading spaces are allowed.|

Since TX data is intended for display purposes, the repeat delimiter, when used with a TX data field, implies a series of repeating lines to be displayed on a printer or terminal. Therefore, the repeat delimiters are regarded as paragraph terminators or hard carriage returns (e.g., they would display as though a CR/LF were inserted in the text (DOS type system) or as though a LF were inserted into the text (UNIX style system)).

A receiving system would word wrap the text between repeat delimiters in order to fit it into an arbitrarily sized display window but start any line beginning with a repeat delimiter on a new line.

To include alternative character sets, use the appropriate escape sequence. See Chapter 2, section 2.14.9.18, "MSH-18 - Character Set" and section 2.14.9.20, "MSH-20 - Alternate Character Set Handling Scheme".

This specification applies no limit to the length of the TX data type, either here where the data type is defined, or elsewhere where the data type is used. While there is no intrinsic reason to limit the length of this data type for semantic or syntactical reasons, it is expected that some sort of limitation will be imposed for technical reasons in implementations. HL7 recommends that implementation length limits be published in implementation profiles.
*/
type TX string

// NewTX creates an implementation of TX
func NewTX(input hl7.HL7Type) TX {
	return TX(input.String())
}

// NewTXSlice will construct a slice of type TX
func NewTXSlice(input hl7.HL7Type) []TX {
	values := make([]TX, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewTX(input.Index(i))
	}
	return values
}

type RCD struct {
	SegmentFieldName   ST
	Hl7DataType        ID
	MaximumColumnWidth NM
}

// NewRCD creates an implementation of RCD
func NewRCD(input hl7.HL7Type) RCD {
	v := RCD{}
	v.SegmentFieldName = NewST(input.Index(0))
	v.Hl7DataType = NewID(input.Index(1))
	v.MaximumColumnWidth = NewNM(input.Index(2))
	return v
}

// NewRCDSlice will construct a slice of type RCD
func NewRCDSlice(input hl7.HL7Type) []RCD {
	values := make([]RCD, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewRCD(input.Index(i))
	}
	return values
}

type CQ struct {
	Quantity NM
	Units    CWE
}

// NewCQ creates an implementation of CQ
func NewCQ(input hl7.HL7Type) CQ {
	v := CQ{}
	v.Quantity = NewNM(input.Index(0))
	v.Units = NewCWE(input.Index(1))
	return v
}

// NewCQSlice will construct a slice of type CQ
func NewCQSlice(input hl7.HL7Type) []CQ {
	values := make([]CQ, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewCQ(input.Index(i))
	}
	return values
}

type DDI struct {
	DelayDays      NM
	MonetaryAmount MO
	NumberOfDays   NM
}

// NewDDI creates an implementation of DDI
func NewDDI(input hl7.HL7Type) DDI {
	v := DDI{}
	v.DelayDays = NewNM(input.Index(0))
	v.MonetaryAmount = NewMO(input.Index(1))
	v.NumberOfDays = NewNM(input.Index(2))
	return v
}

// NewDDISlice will construct a slice of type DDI
func NewDDISlice(input hl7.HL7Type) []DDI {
	values := make([]DDI, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDDI(input.Index(i))
	}
	return values
}

type AUI struct {
	AuthorizationNumber ST
	Date                DT
	Source              ST
}

// NewAUI creates an implementation of AUI
func NewAUI(input hl7.HL7Type) AUI {
	v := AUI{}
	v.AuthorizationNumber = NewST(input.Index(0))
	v.Date = NewDT(input.Index(1))
	v.Source = NewST(input.Index(2))
	return v
}

// NewAUISlice will construct a slice of type AUI
func NewAUISlice(input hl7.HL7Type) []AUI {
	values := make([]AUI, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewAUI(input.Index(i))
	}
	return values
}

type DTN struct {
	DayType      CWE
	NumberOfDays NM
}

// NewDTN creates an implementation of DTN
func NewDTN(input hl7.HL7Type) DTN {
	v := DTN{}
	v.DayType = NewCWE(input.Index(0))
	v.NumberOfDays = NewNM(input.Index(1))
	return v
}

// NewDTNSlice will construct a slice of type DTN
func NewDTNSlice(input hl7.HL7Type) []DTN {
	values := make([]DTN, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewDTN(input.Index(i))
	}
	return values
}

type PT struct {
	ProcessingID   ID
	ProcessingMode ID
}

// NewPT creates an implementation of PT
func NewPT(input hl7.HL7Type) PT {
	v := PT{}
	v.ProcessingID = NewID(input.Index(0))
	v.ProcessingMode = NewID(input.Index(1))
	return v
}

// NewPTSlice will construct a slice of type PT
func NewPTSlice(input hl7.HL7Type) []PT {
	values := make([]PT, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewPT(input.Index(i))
	}
	return values
}

type NA struct {
	Value1 NM
	Value2 NM
	Value3 NM
	Value4 NM
}

// NewNA creates an implementation of NA
func NewNA(input hl7.HL7Type) NA {
	v := NA{}
	v.Value1 = NewNM(input.Index(0))
	v.Value2 = NewNM(input.Index(1))
	v.Value3 = NewNM(input.Index(2))
	v.Value4 = NewNM(input.Index(3))
	return v
}

// NewNASlice will construct a slice of type NA
func NewNASlice(input hl7.HL7Type) []NA {
	values := make([]NA, input.Len())
	for i := 0; i < input.Len(); i++ {
		values[i] = NewNA(input.Index(i))
	}
	return values
}
