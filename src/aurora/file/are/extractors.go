package are

import (
	"aurora/file/gff"
	"aurora/tools"
)

func extractAreaTilesList(byteOffset uint32, gffFile gff.GFF, errorBag *tools.ErrorBag) []AreaTile {
	var result = []AreaTile{}

	if errorBag.Error != nil {
		return result
	}

	var indexes, err = gff.ExtractListStructArrayIndexes(byteOffset, gffFile)

	if err != nil {
		errorBag.Error = err
		return result
	}

	for _, index := range indexes {
		var fields, err = gff.ExtractStructFields(gffFile.StructArray[index], gffFile)

		if err != nil {
			errorBag.Error = err

			return result
		}

		var element = AreaTile{
			AnimLoop1:   gff.GetFieldByteValue("Tile_AnimLoop1", fields, errorBag),
			AnimLoop2:   gff.GetFieldByteValue("Tile_AnimLoop2", fields, errorBag),
			AnimLoop3:   gff.GetFieldByteValue("Tile_AnimLoop3", fields, errorBag),
			Height:      gff.GetFieldInt32Value("Tile_Height", fields, errorBag),
			ID:          gff.GetFieldInt32Value("Tile_ID", fields, errorBag),
			MainLight1:  gff.GetFieldByteValue("Tile_MainLight1", fields, errorBag),
			MainLight2:  gff.GetFieldByteValue("Tile_MainLight2", fields, errorBag),
			Orientation: gff.GetFieldInt32Value("Tile_Orientation", fields, errorBag),
			SrcLight1:   gff.GetFieldByteValue("Tile_SrcLight1", fields, errorBag),
			SrcLight2:   gff.GetFieldByteValue("Tile_SrcLight2", fields, errorBag),
		}

		result = append(result, element)
	}

	return result
}
