package flights

import (
	"fmt"
)

var airlineAcronyms = map[string]string{
	"AA": "American Airlines",
	"AS": "Alaska Airlines",
	"B6": "JetBlue Airways",
	"F9": "Frontier Airlines",
	"G4": "Allegiant Air",
	"HA": "Hawaiian Airlines",
	"NK": "Spirit Airlines",
	"SY": "Sun Country Airlines",
	"WN": "Southwest Airlines",
	"WS": "WestJet",
	"A3": "Aegean Airlines",
	"AB": "Air Berlin (defunct)",
	"AF": "Air France",
	"AY": "Finnair",
	"AZ": "Alitalia",
	"BA": "British Airways",
	"EI": "Aer Lingus",
	"EW": "Eurowings",
	"IB": "Iberia",
	"KL": "KLM Royal Dutch Airlines",
	"LH": "Lufthansa",
	"LX": "Swiss International Air Lines",
	"OS": "Austrian Airlines",
	"SK": "Scandinavian Airlines",
	"TK": "Turkish Airlines",
	"U2": "easyJet",
	"VY": "Vueling",
	"9W": "Jet Airways (defunct)",
	"AI": "Air India",
	"CA": "Air China",
	"CX": "Cathay Pacific",
	"EK": "Emirates",
	"EY": "Etihad Airways",
	"GA": "Garuda Indonesia",
	"JL": "Japan Airlines",
	"KE": "Korean Air",
	"MH": "Malaysia Airlines",
	"NH": "All Nippon Airways",
	"OZ": "Asiana Airlines",
	"QR": "Qatar Airways",
	"SQ": "Singapore Airlines",
	"TG": "Thai Airways",
	"VN": "Vietnam Airlines",
	"NZ": "Air New Zealand",
	"QF": "Qantas",
	"VA": "Virgin Australia",
	"4O": "Interjet (defunct)",
	"AR": "Aerolineas Argentinas",
	"AV": "Avianca",
	"CM": "COPA Airlines",
	"ET": "Ethiopian Airlines",
	"G3": "Gol Transportes Aéreos",
	"LA": "LATAM Airlines Group",
	"MS": "EgyptAir",
	"SA": "South African Airways",
	"SV": "Saudia",
	"W6": "Wizz Air",
	"BZ": "Blue Dart Aviation",
	"LY": "El Al Israel Airlines",
	"QS": "Smartwings",
	"RO": "Tarom",
	"5W": "Flyadeal",
	"W4": "Wizz Air Malta",
	"CY": "Cyprus Airways",
	"H4": "HiSky",
	"3F": "FlyOne Armenia",
	"U8": "TUS Airways",
}

func FindAirlineName(acronym string) (string, error) {
	name, exists := airlineAcronyms[acronym]
	if !exists {
		return "", fmt.Errorf("airline not found: %s", acronym)
	}
	return name, nil
}
