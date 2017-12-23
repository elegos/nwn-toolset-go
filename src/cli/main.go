package main

import (
	"aurora/file"
	"aurora/file/erf"
	"flag"
	"fmt"
)

func main() {
	fileName := flag.String("file", "", "the module file name")

	flag.Parse()

	// Some stats
	module := erf.FromFile(*fileName)

	fmt.Println("Header")
	fmt.Println("")
	fmt.Println(fmt.Sprintf("File type:               '%s'", module.Header.FileType))
	fmt.Println(fmt.Sprintf("Version:                 '%s'", module.Header.Version))
	fmt.Println(fmt.Sprintf("LanguageCount:           %d", module.Header.LanguageCount))
	fmt.Println(fmt.Sprintf("LocalizedStringSize:     %d", module.Header.LocalizedStringSize))
	fmt.Println(fmt.Sprintf("EntryCount:              %d", module.Header.EntryCount))
	fmt.Println(fmt.Sprintf("OffsetToLocalizedString: %d", module.Header.OffsetToLocalizedString))
	fmt.Println(fmt.Sprintf("OffsetToKeyList:         %d", module.Header.OffsetToKeyList))
	fmt.Println(fmt.Sprintf("OffsetToResourceList:    %d", module.Header.OffsetToResourceList))
	fmt.Println(fmt.Sprintf("BuildYear:               %d (%d)", module.Header.BuildYear, module.Header.BuildYear+1900))
	fmt.Println(fmt.Sprintf("BuildDay:                %d", module.Header.BuildDay))
	fmt.Println(fmt.Sprintf("DescriptionStrRef:       %d", module.Header.DescriptionStrRef))

	fmt.Println("===================================")
	fmt.Println("")
	fmt.Println("Localized strings list")
	fmt.Println("")
	for _, element := range module.LocalizedStringList {
		fmt.Println(fmt.Sprintf("Language: %s, Size: %d: %s", file.LanguageLookup[element.LanguageID], element.StringSize, element.String))
	}

	fmt.Println("===================================")
	fmt.Println("")
	fmt.Println("Keys list")
	fmt.Println("")
	for _, element := range module.KeyList {
		fmt.Println(fmt.Sprintf("ResRef: %s, ResID: %d,\tResType: %s", element.ResRef, element.ResID, file.ResourceTypeLookup[element.ResType]))
	}

	fmt.Println("===================================")
	fmt.Println("")
	var resourceDataBytes = len(module.ResourceData)
	var resourceDataKB = resourceDataBytes / 1024
	var resourceDataMB = resourceDataKB / 1024
	fmt.Println(fmt.Sprintf("Resource data size: %d MB (%d KB, %d bytes)", resourceDataMB, resourceDataKB, resourceDataBytes))
}
