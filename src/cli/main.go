package main

import (
	"aurora/file/erf"
	"flag"
	"fmt"
)

func main() {
	fileName := flag.String("file", "", "the module file name")

	flag.Parse()

	// Some stats
	module := erf.FromFile(*fileName)

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
}
