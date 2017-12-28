package command

import (
	auroraFile "aurora/file"
	"aurora/file/erf"
	"aurora/file/gff"
	"fmt"
	"strconv"
)

// ReadGffFromErf read GFF files from ERF
func ReadGffFromErf(file *erf.ERF) {
	// Known GFF files
	var gffResTypes = []auroraFile.ResourceType{
		auroraFile.Are, auroraFile.Bic, auroraFile.Dlg,
		auroraFile.Fac, auroraFile.Gff, auroraFile.Gic,
		auroraFile.Git, auroraFile.Gui, auroraFile.Ifo,
		auroraFile.Itp, auroraFile.Jrl, auroraFile.Ptm,
		auroraFile.Ptt, auroraFile.Utc, auroraFile.Utd,
		auroraFile.Ute, auroraFile.Uti, auroraFile.Utp,
		auroraFile.Uts, auroraFile.Utt, auroraFile.Utw,
	}

	var resourceDataOffset = file.Header.OffsetToResourceList + file.Header.EntryCount*8

	for index, key := range file.KeyList {
		// Skip if not a GFF file
		for _, resType := range gffResTypes {
			if resType != key.ResType {
				continue
			}
		}

		var listElement = file.ResourceList[index]
		var offset = listElement.OffsetToResource - resourceDataOffset

		var fileContent = file.ResourceData[offset : offset+listElement.ResourceSize-1]

		var data = gff.FromBytes(fileContent)
		fmt.Println(fmt.Sprintf(
			"Header: {FileType: %s, FileVersion: %s, StructCount: %3d, FieldCount: %4d, LabelCount: %4d, FieldDataCount: %5db, FieldIndicesCount: %5db, ListIndicesCount: %5db}",
			data.Header.FileType,
			data.Header.FileVersion,
			data.Header.StructCount,
			data.Header.FieldCount,
			data.Header.LabelCount,
			data.Header.FieldDataCount,
			data.Header.FieldIndicesCount,
			data.Header.ListIndicesCount,
		))
		var structArray = []string{}
		for _, element := range data.StructArray {
			var eType = strconv.Itoa(int(element.Type))
			if element.Type == 0xFFFFFFFF {
				eType = "Top-Level struct"
			}
			structArray = append(structArray, fmt.Sprintf(
				"\n  {Type: %16s, DataOrDataOffset: %3d, FieldCount: %3d}",
				eType,
				element.DataOrDataOffset,
				element.FieldCount,
			))
		}
		fmt.Println(fmt.Sprintf("StructArray: %v", structArray))
		fmt.Println("")

		var fieldArray = []string{}
		for _, element := range data.FieldArray {
			fieldArray = append(fieldArray, fmt.Sprintf(
				"\n  {Type: %12d, LabelIndex: %12d, DataOrDataOffset: %12d}",
				element.Type,
				element.LabelIndex,
				element.DataOrDataOffset,
			))
		}
		fmt.Println(fmt.Sprintf("FieldArray: %v", fieldArray))
		fmt.Println("")

		fmt.Println("LabelArray:")
		for index, element := range data.LabelArray {
			fmt.Println(fmt.Sprintf("%3d: %s", index, element))
		}
		fmt.Println("")

		fmt.Println(fmt.Sprintf("Field data block size: %d bytes", len(data.FieldDataBlock)))
		fmt.Println("")

		fmt.Println("FieldIndicesArray:")
		for index, element := range data.FieldIndicesArray {
			fmt.Println(fmt.Sprintf("%3d: %d", index, element))
		}
		fmt.Println("")

		var listIndicesArray = []string{}
		for _, element := range data.ListIndicesArray {
			listIndicesArray = append(listIndicesArray, fmt.Sprintf(
				"\n  %v", element,
			))
		}
		fmt.Println(fmt.Sprintf("ListIndicesArray: %v", listIndicesArray))
		fmt.Println("")

		fmt.Println("================================================")
		fmt.Println("")
	}
}
