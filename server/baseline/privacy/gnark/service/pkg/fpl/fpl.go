package fpl

import (
	"github.com/consensys/gnark/frontend"
)

//Note that:
//Each element of an array is an ASCII character in decimal representation

// Raw FPL Plan

// 7. AIRCRAFT IDENTIFICATION
// Assume that the aircraft identifier has been decomposed into ASCII, the ASCII concatenated, and then split at each integer to form an array of 14 digits
type AircraftIdentifer = string

// 8. FLIGHT RULES (IFR, VFR or combination of the two) and TYPE OF FLIGHT (Scheduled Air Service, General Aviation etc.)
type FlightRule = string
type TypeOfFlight = string

// 9. NUMBER, TYPE OF AIRCRAFT, WAKE TURBULENCE CATEGORY
type NumberOfAircraft = string
type TypeOfAircraft = string
type WakeTurbulenceCategory = string

// 10. EQUIPMENT AND CAPABILITIES ( RADIO COMMUNICATION, NAVIGATION AND APPROACH AID EQUIPMENT AND CAPABILITIES, SURVEILLANCE EQUIPMENT)
type Serviceable = string
type Available = string

type C struct {
	A Available `json:"a"'`
}

type B struct {
	A Available `json:"a"'`
}

type ADS struct {
	C C `json:"c"'`
	B B `json:"b"'`
}

type ModeAC struct {
	A Available `json:"a"'`
}

type ModeS struct {
	A Available `json:"a"'`
}

type SSR struct {
	ModeAC ModeAC `json:"mode_a_c"'`
	ModeS  ModeS  `json:"mode_s"'`
}

type Surveillance struct {
	S   Serviceable `json:"s"'`
	SSR SSR         `json:"ssr"'`
	ADS ADS         `json:"ads"'`
}

type RCNAAE struct {
	S Serviceable `json:"s"'`
	A Available   `json:"a"'`
}

type EquipmentAndCapabilites struct {
	RCNAAE       RCNAAE       `json:"rc_n_aae"'`
	Surveillance Surveillance `json:"surveillance"'`
}

// 13. DEPARTURE AERODROME AND TIME (8 CHARACTERS)
type Time = string
type DepartureAerodrome = string

// 15. CRUISING SPEED LEVEL AND ROUTE
type CruisingSpeed = string
type CruisingLevel = string

//Allow for a route of upto 152 characters
type Route = string

// 16. DESTINATION, TOTAL ELAPSED TIME AND DESTINATION ALTERNATE AERODROMES
type DestinationAerodrome = string
type TotalEstimatedElapsedTime = string
type AlternateDestinationAerodrome = string
type AlternateDestinationAerodromes = [2]AlternateDestinationAerodrome

// 18. OTHER INFORMATION
type OtherInformation = string

// 19. SUPPLEMENTARY INFORMATION
type Endurance = string
type PersonsOnBoard = string         //allow for hundreds of passengers
type EmergencyRadio = string         //precede letter to be crossed out by 'x'
type SurvivalEquipment = string      //precede letter to be crossed out by 'x'
type Jackets = string                //precede letters to be crossed out by 'x'
type Dinghies = string               //HashSegment[2] // precede letters to be crossed out by 'x' or do as otherwise
type AircraftColourMarkings = string //HashSegment[2]
type Remarks = string                //HashSegment[2]
type PilotName = string              //HashSegment

type FPLForm struct {
	AI  AircraftIdentifer `json:"ai"`
	FR  FlightRule        `json:"fr"`
	TOF TypeOfFlight      `json:"tof"`

	NOA NumberOfAircraft       `json:"noa"`
	TOA TypeOfAircraft         `json:"toa"`
	WTC WakeTurbulenceCategory `json:"wtc"`

	EANDC EquipmentAndCapabilites `json:"eandc"`

	T  Time               `json:"t"`
	DA DepartureAerodrome `json:"da"`

	CS CruisingSpeed `json:"cs"`
	CL CruisingLevel `json:"cl"`

	R Route `json:"r"`

	DE   DestinationAerodrome           `json:"de"`
	TEET TotalEstimatedElapsedTime      `json:"teet"`
	ADA  AlternateDestinationAerodrome  `json:"ada"`
	ADAS AlternateDestinationAerodromes `json:"adas"`

	OI OtherInformation `json:"oi"`

	E   Endurance         `json:"e"`
	POB PersonsOnBoard    `json:"pob"`
	ER  EmergencyRadio    `json:"er"`
	SE  SurvivalEquipment `json:"se"`
	J   Jackets           `json:"j"`

	D   Dinghies               `json:"d"`
	ACM AircraftColourMarkings `json:"acm"`

	RE Remarks   `json:"re"`
	PN PilotName `json:"pn"`
}

// Parsed FPL Plan

// 7. AIRCRAFT IDENTIFICATION
// Assume that the aircraft identifier has been decomposed into ASCII
type ParsedAircraftIdentifer = [7]int

// 8. FLIGHT RULES (IFR, VFR or combination of the two) and TYPE OF FLIGHT (Scheduled Air Service, General Aviation etc.)
type ParsedFlightRule = [1]int
type ParsedTypeOfFlight = [1]int

// 9. NUMBER, TYPE OF AIRCRAFT, WAKE TURBULENCE CATEGORY
type ParsedNumberOfAircraft = [2]int
type ParsedTypeOfAircraft = [4]int
type ParsedWakeTurbulenceCategory = [1]int

// 10. EQUIPMENT AND CAPABILITIES ( RADIO COMMUNICATION, NAVIGATION AND APPROACH AID EQUIPMENT AND CAPABILITIES, SURVEILLENCE EQUIPMENT)
type ParsedServiceable = int
type ParsedAvailable = int

type ParsedC struct {
	A [4]ParsedAvailable
}

type ParsedB struct {
	A [12]ParsedAvailable
}

type ParsedADS struct {
	C ParsedC
	B ParsedB
}

type ParsedModeAC struct {
	A [2]ParsedAvailable
}

type ParsedModeS struct {
	A [7]ParsedAvailable
}

type ParsedSSR struct {
	ModeAC ParsedModeAC
	ModeS  ParsedModeS
}

type ParsedSurveillance struct {
	S   [1]ParsedServiceable
	SSR ParsedSSR
	ADS ParsedADS
}

type ParsedRCNAAE struct {
	S [2]ParsedServiceable
	A [49]ParsedAvailable
}

type ParsedEquipmentAndCapabilites struct {
	RCNAAE       ParsedRCNAAE
	Surveillance ParsedSurveillance
}

// 13. DEPARTURE AERODROME AND TIME (8 CHARACTERS)
type ParsedTime = [4]int
type ParsedDepartureAerodrome = [4]int

// 15. CRUISING SPEED LEVEL AND ROUTE
type ParsedCruisingSpeed = [5]int
type ParsedCruisingLevel = [5]int

//Allow for a route of upto 148 characters
type ParsedRoute = [148]int //HashSegment[8]

// 16. DESTINATION, TOTAL ELAPSED TIME AND DESTINATION ALTERNATE AERODROMES
type ParsedDestinationAerodrome = [4]int
type ParsedTotalEstimatedElapsedTime = [4]int
type ParsedAlternateDestinationAerodrome = [4]int
type ParsedAlternateDestinationAerodromes = [2]ParsedAlternateDestinationAerodrome

// 18. OTHER INFORMATION
type ParsedOtherInformation = [74]int //HashSegment[4]

// 19. SUPPLEMENTARY INFORMATION
type ParsedEndurance = [6]int
type ParsedPersonsOnBoard = [5]int          //allow for hundreds of passengers
type ParsedEmergencyRadio = [4]int          //precede letter to be crossed out by 'x'
type ParsedSurvivalEquipment = [4]int       //precede letter to be crossed out by 'x'
type ParsedJackets = [10]int                //precede letters to be crossed out by 'x'
type ParsedDinghies = [37]int               //HashSegment[2] // precede letters to be crossed out by 'x' or do as otherwise
type ParsedAircraftColourMarkings = [37]int //HashSegment[2]
type ParsedRemarks = [37]int                //HashSegment[2]
type ParsedPilotName = [19]int              //HashSegment

type ParsedFPLForm struct {
	AI  ParsedAircraftIdentifer //14 h1
	FR  ParsedFlightRule        //2 h1
	TOF ParsedTypeOfFlight      //2 h1

	NOA ParsedNumberOfAircraft       //4 h1
	TOA ParsedTypeOfAircraft         //8 h1
	WTC ParsedWakeTurbulenceCategory //2 h1

	EANDC ParsedEquipmentAndCapabilites //4 98 2 4 14 8 24 h1 | 98 h2

	T  ParsedTime               //8 h1
	DA ParsedDepartureAerodrome //8 h1

	CS ParsedCruisingSpeed //10 h1
	CL ParsedCruisingLevel //10 h1

	R ParsedRoute //38 * 8 -> 304 h3 h4

	DE   ParsedDestinationAerodrome           //8 h1
	TEET ParsedTotalEstimatedElapsedTime      //8 h1
	ADA  ParsedAlternateDestinationAerodrome  //8 h1
	ADAS ParsedAlternateDestinationAerodromes //8 * 2 h7

	OI ParsedOtherInformation //38*4 -> 152 h5

	E   ParsedEndurance         //12 h7
	POB ParsedPersonsOnBoard    //10 h7
	ER  ParsedEmergencyRadio    //8 h7
	SE  ParsedSurvivalEquipment //8 h7
	J   ParsedJackets           //20 h8

	D   ParsedDinghies               //38*2 -> 76 h6
	ACM ParsedAircraftColourMarkings //38*2 -> 76 h6

	RE ParsedRemarks   //38*2 -> 76 h8
	PN ParsedPilotName //38 h7
}

// Circuit FPL Plan

// 7. AIRCRAFT IDENTIFICATION
// Assume that the aircraft identifier has been decomposed into ASCII, the ASCII concatenated, and then split at each integer to form an array of 14 digits
type CircuitAircraftIdentifer = [7]frontend.Variable

// 8. FLIGHT RULES (IFR, VFR or combination of the two) and TYPE OF FLIGHT (Scheduled Air Service, General Aviation etc.)
type CircuitFlightRule = [1]frontend.Variable
type CircuitTypeOfFlight = [1]frontend.Variable

// 9. NUMBER, TYPE OF AIRCRAFT, WAKE TURBULENCE CATEGORY
type CircuitNumberOfAircraft = [2]frontend.Variable
type CircuitTypeOfAircraft = [4]frontend.Variable
type CircuitWakeTurbulenceCategory = [1]frontend.Variable

// 10. EQUIPMENT AND CAPABILITIES ( RADIO COMMUNICATION, NAVIGATION AND APPROACH AID EQUIPMENT AND CAPABILITIES, SURVEILLENCE EQUIPMENT)
type CircuitServiceable = frontend.Variable
type CircuitAvailable = frontend.Variable

type CircuitC struct {
	A [4]CircuitAvailable
}

type CircuitB struct {
	A [12]CircuitAvailable
}

type CircuitADS struct {
	C CircuitC
	B CircuitB
}

type CircuitModeAC struct {
	A [2]CircuitAvailable
}

type CircuitModeS struct {
	A [7]CircuitAvailable
}

type CircuitSSR struct {
	ModeAC CircuitModeAC
	ModeS  CircuitModeS
}

type CircuitSurveillance struct {
	S   [1]CircuitServiceable
	SSR CircuitSSR
	ADS CircuitADS
}

type CircuitRCNAAE struct {
	S [2]CircuitServiceable
	A [49]CircuitAvailable
}

type CircuitEquipmentAndCapabilites struct {
	RCNAAE       CircuitRCNAAE
	Surveillance CircuitSurveillance
}

// 13. DEPARTURE AERODROME AND TIME (8 CHARACTERS)
type CircuitTime = [4]frontend.Variable
type CircuitDepartureAerodrome = [4]frontend.Variable

// 15. CRUISING SPEED LEVEL AND ROUTE
type CircuitCruisingSpeed = [5]frontend.Variable
type CircuitCruisingLevel = [5]frontend.Variable

//Allow for a route of upto 152 characters
type CircuitRoute = [148]frontend.Variable //HashSegment[8]

// 16. DESTINATION, TOTAL ELAPSED TIME AND DESTINATION ALTERNATE AERODROMES
type CircuitDestinationAerodrome = [4]frontend.Variable
type CircuitTotalEstimatedElapsedTime = [4]frontend.Variable
type CircuitAlternateDestinationAerodrome = [4]frontend.Variable
type CircuitAlternateDestinationAerodromes = [2]CircuitAlternateDestinationAerodrome

// 18. OTHER INFORMATION
type CircuitOtherInformation = [74]frontend.Variable //HashSegment[4]

// 19. SUPPLEMENTARY INFORMATION
type CircuitEndurance = [6]frontend.Variable
type CircuitPersonsOnBoard = [5]frontend.Variable          //allow for hundreds of passengers
type CircuitEmergencyRadio = [4]frontend.Variable          //precede letter to be crossed out by 'x'
type CircuitSurvivalEquipment = [4]frontend.Variable       //precede letter to be crossed out by 'x'
type CircuitJackets = [10]frontend.Variable                //precede letters to be crossed out by 'x'
type CircuitDinghies = [37]frontend.Variable               //HashSegment[2] // precede letters to be crossed out by 'x' or do as otherwise
type CircuitAircraftColourMarkings = [37]frontend.Variable //HashSegment[2]
type CircuitRemarks = [37]frontend.Variable                //HashSegment[2]
type CircuitPilotName = [19]frontend.Variable              //HashSegment

type CircuitFPLForm struct {
	AI  CircuitAircraftIdentifer //14 h1
	FR  CircuitFlightRule        //2 h1
	TOF CircuitTypeOfFlight      //2 h1

	NOA CircuitNumberOfAircraft       //4 h1
	TOA CircuitTypeOfAircraft         //8 h1
	WTC CircuitWakeTurbulenceCategory //2 h1

	EANDC CircuitEquipmentAndCapabilites //4 98 2 4 14 8 24 h1 | 98 h2

	T  CircuitTime               //8 h1
	DA CircuitDepartureAerodrome //8 h1

	CS CircuitCruisingSpeed //10 h1
	CL CircuitCruisingLevel //10 h1

	R CircuitRoute //38 * 8 -> 304 h3 h4

	DE   CircuitDestinationAerodrome           //8 h1
	TEET CircuitTotalEstimatedElapsedTime      //8 h1
	ADA  CircuitAlternateDestinationAerodrome  //8 h1
	ADAS CircuitAlternateDestinationAerodromes //8 * 2 h7

	OI CircuitOtherInformation //38*4 -> 152 h5

	E   CircuitEndurance         //12 h7
	POB CircuitPersonsOnBoard    //10 h7
	ER  CircuitEmergencyRadio    //8 h7
	SE  CircuitSurvivalEquipment //8 h7
	J   CircuitJackets           //20 h8

	D   CircuitDinghies               //38*2 -> 76 h6
	ACM CircuitAircraftColourMarkings //38*2 -> 76 h6

	RE CircuitRemarks   //38*2 -> 76 h8
	PN CircuitPilotName //38 h7
}

type HashedFPLForm = frontend.Variable
type HashedFPLFormACK = frontend.Variable

func parseFPLField(field string) []int {
	fieldAsASCII := []rune(field)
	//var s string
	//for _, character := range fieldAsASCII {
	//	s += strconv.FormatInt(int64(character), 10)
	//}
	//sSplit := strings.Split(s, "")
	//digits := make([]int, 0)
	//for _, char := range sSplit {
	//	i, _ := strconv.ParseInt(char, 10, 8)
	//	digits = append(digits, int(i))
	//}
	digits := make([]int, 0)
	for _, character := range fieldAsASCII {
		digits = append(digits, int(character))
	}
	return digits
}

func ParseFPL(form FPLForm) ParsedFPLForm {
	parsedForm := ParsedFPLForm{}
	copy(parsedForm.AI[:], parseFPLField(form.AI))
	copy(parsedForm.FR[:], parseFPLField(form.FR))
	copy(parsedForm.TOF[:], parseFPLField(form.TOF))

	copy(parsedForm.NOA[:], parseFPLField(form.NOA))
	copy(parsedForm.TOA[:], parseFPLField(form.TOA))
	copy(parsedForm.WTC[:], parseFPLField(form.WTC))

	copy(parsedForm.EANDC.RCNAAE.S[:], parseFPLField(form.EANDC.RCNAAE.S))
	copy(parsedForm.EANDC.RCNAAE.A[:], parseFPLField(form.EANDC.RCNAAE.A))
	copy(parsedForm.EANDC.Surveillance.S[:], parseFPLField(form.EANDC.Surveillance.S))
	copy(parsedForm.EANDC.Surveillance.SSR.ModeAC.A[:], parseFPLField(form.EANDC.Surveillance.SSR.ModeAC.A))
	copy(parsedForm.EANDC.Surveillance.SSR.ModeS.A[:], parseFPLField(form.EANDC.Surveillance.SSR.ModeS.A))
	copy(parsedForm.EANDC.Surveillance.ADS.C.A[:], parseFPLField(form.EANDC.Surveillance.ADS.C.A))
	copy(parsedForm.EANDC.Surveillance.ADS.B.A[:], parseFPLField(form.EANDC.Surveillance.ADS.B.A))

	copy(parsedForm.T[:], parseFPLField(form.T))
	copy(parsedForm.DA[:], parseFPLField(form.DA))

	copy(parsedForm.CS[:], parseFPLField(form.CS))
	copy(parsedForm.CL[:], parseFPLField(form.CL))

	copy(parsedForm.R[:], parseFPLField(form.R))

	copy(parsedForm.DE[:], parseFPLField(form.DE))
	copy(parsedForm.TEET[:], parseFPLField(form.TEET))
	copy(parsedForm.ADA[:], parseFPLField(form.ADA))
	copy(parsedForm.ADAS[0][:], parseFPLField(form.ADAS[0]))
	copy(parsedForm.ADAS[1][:], parseFPLField(form.ADAS[1]))

	copy(parsedForm.OI[:], parseFPLField(form.OI))

	copy(parsedForm.E[:], parseFPLField(form.E))
	copy(parsedForm.POB[:], parseFPLField(form.POB))
	copy(parsedForm.ER[:], parseFPLField(form.ER))
	copy(parsedForm.SE[:], parseFPLField(form.SE))
	copy(parsedForm.J[:], parseFPLField(form.J))

	copy(parsedForm.D[:], parseFPLField(form.D))
	copy(parsedForm.ACM[:], parseFPLField(form.ACM))

	copy(parsedForm.RE[:], parseFPLField(form.RE))
	copy(parsedForm.PN[:], parseFPLField(form.PN))

	return parsedForm
}

func circuitFieldFromParsedField(parsedField []int) []frontend.Variable {
	//for i, j := 0, len(parsedField)-1; i < j; i, j = i+1, j-1 {
	//	parsedField[i], parsedField[j] = parsedField[j], parsedField[i]
	//}
	circuitField := make([]frontend.Variable, len(parsedField))
	for i, _ := range parsedField {
		circuitField[i] = frontend.Variable(parsedField[i])
	}
	return circuitField
}

func ParsedFPLToCircuitFPL(parsedForm ParsedFPLForm) CircuitFPLForm {
	circuitForm := CircuitFPLForm{}
	copy(circuitForm.AI[:], circuitFieldFromParsedField(parsedForm.AI[:]))
	copy(circuitForm.FR[:], circuitFieldFromParsedField(parsedForm.FR[:]))
	copy(circuitForm.TOF[:], circuitFieldFromParsedField(parsedForm.TOF[:]))

	copy(circuitForm.NOA[:], circuitFieldFromParsedField(parsedForm.NOA[:]))
	copy(circuitForm.TOA[:], circuitFieldFromParsedField(parsedForm.TOA[:]))
	copy(circuitForm.WTC[:], circuitFieldFromParsedField(parsedForm.WTC[:]))

	copy(circuitForm.EANDC.RCNAAE.S[:], circuitFieldFromParsedField(parsedForm.EANDC.RCNAAE.S[:]))
	copy(circuitForm.EANDC.RCNAAE.A[:], circuitFieldFromParsedField(parsedForm.EANDC.RCNAAE.A[:]))
	copy(circuitForm.EANDC.Surveillance.S[:], circuitFieldFromParsedField(parsedForm.EANDC.Surveillance.S[:]))
	copy(circuitForm.EANDC.Surveillance.SSR.ModeAC.A[:], circuitFieldFromParsedField(parsedForm.EANDC.Surveillance.SSR.ModeAC.A[:]))
	copy(circuitForm.EANDC.Surveillance.SSR.ModeS.A[:], circuitFieldFromParsedField(parsedForm.EANDC.Surveillance.SSR.ModeS.A[:]))
	copy(circuitForm.EANDC.Surveillance.ADS.C.A[:], circuitFieldFromParsedField(parsedForm.EANDC.Surveillance.ADS.C.A[:]))
	copy(circuitForm.EANDC.Surveillance.ADS.B.A[:], circuitFieldFromParsedField(parsedForm.EANDC.Surveillance.ADS.B.A[:]))

	copy(circuitForm.T[:], circuitFieldFromParsedField(parsedForm.T[:]))
	copy(circuitForm.DA[:], circuitFieldFromParsedField(parsedForm.DA[:]))

	copy(circuitForm.CS[:], circuitFieldFromParsedField(parsedForm.CS[:]))
	copy(circuitForm.CL[:], circuitFieldFromParsedField(parsedForm.CL[:]))

	copy(circuitForm.R[:], circuitFieldFromParsedField(parsedForm.R[:]))

	copy(circuitForm.DE[:], circuitFieldFromParsedField(parsedForm.DE[:]))
	copy(circuitForm.TEET[:], circuitFieldFromParsedField(parsedForm.TEET[:]))
	copy(circuitForm.ADA[:], circuitFieldFromParsedField(parsedForm.ADA[:]))
	copy(circuitForm.ADAS[0][:], circuitFieldFromParsedField(parsedForm.ADAS[0][:]))
	copy(circuitForm.ADAS[1][:], circuitFieldFromParsedField(parsedForm.ADAS[1][:]))

	copy(circuitForm.OI[:], circuitFieldFromParsedField(parsedForm.OI[:]))

	copy(circuitForm.E[:], circuitFieldFromParsedField(parsedForm.E[:]))
	copy(circuitForm.POB[:], circuitFieldFromParsedField(parsedForm.POB[:]))
	copy(circuitForm.ER[:], circuitFieldFromParsedField(parsedForm.ER[:]))
	copy(circuitForm.SE[:], circuitFieldFromParsedField(parsedForm.SE[:]))
	copy(circuitForm.J[:], circuitFieldFromParsedField(parsedForm.J[:]))

	copy(circuitForm.D[:], circuitFieldFromParsedField(parsedForm.D[:]))
	copy(circuitForm.ACM[:], circuitFieldFromParsedField(parsedForm.ACM[:]))

	copy(circuitForm.RE[:], circuitFieldFromParsedField(parsedForm.RE[:]))
	copy(circuitForm.PN[:], circuitFieldFromParsedField(parsedForm.PN[:]))

	return circuitForm
}
