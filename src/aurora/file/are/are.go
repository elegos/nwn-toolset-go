package are

import (
	"aurora/file/gff"
	"errors"
	"fmt"
)

// Header structure of the ARE file format's header
type Header gff.Header

// ARE Area file format
type ARE struct {
	Header         Header
	TopLevelStruct TopLevelStruct
}

type TopLevelStruct struct {
	ChanceLightning uint32 // Percent chance of lightning (0-100)
	ChanceRain      uint32 // Percent chance of rain (0-100)
	ChanceSnow      uint32 // Percent chance of snow (0-100)
	Comments        string // Module designer comments
	CreatorID       int32  // Deprecated; unused. Always -1.
	DayNightCycle   byte   // 1 if day/night transitions occur, 0 otherwise
}

// FromBytes read an ARE file from its bytes
func FromBytes(bytes []byte) (ARE, error) {
	var result = ARE{}
	var gffFile = gff.FromBytes(bytes)

	if gffFile.Header.FileType != "ARE " {
		return result, fmt.Errorf("Not an ARE file. Found: '%s'", gffFile.Header.FileType)
	}

	if gffFile.Header.StructCount == 0 {
		return result, errors.New("Not a valid GFF file: there are no structs")
	}

	var tls = gffFile.StructArray[0]
	fmt.Println(tls)

	result.Header = Header(gffFile.Header)
	// result.TopLevelStruct = TopLevelStruct{
	//   ChanceLightning: tls.
	// }

	return result, nil
}
