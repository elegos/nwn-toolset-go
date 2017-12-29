package erf_test

import (
	"aurora/file"
	"aurora/file/erf"
	"aurora/tools/test"
	"fmt"
	"os"
	"testing"
)

func checkKeyListElement(t *testing.T, erfData erf.ERF, index int, resRef string, resId uint32, resType file.ResourceType) {
	test.ExpectString(t, fmt.Sprintf("KeyList[%d].ResRef", index), resRef, erfData.KeyList[index].ResRef)
	test.ExpectUint32(t, fmt.Sprintf("KeyList[%d].ResID", index), resId, erfData.KeyList[index].ResID)
	test.ExpectUint32(t, fmt.Sprintf("KeyList[%d].ResType", index), uint32(resType), uint32(erfData.KeyList[index].ResType))
}

func TestFromFile(t *testing.T) {
	var filePath = "test/module.mod"
	erfData, err := erf.FromFile(filePath)

	if err != nil {
		t.Error("Unexpected error")
	}

	// header
	test.ExpectString(t, "Header.FileType", "MOD ", erfData.Header.FileType)
	test.ExpectString(t, "Header.Version", "V1.0", erfData.Header.Version)
	test.ExpectUint32(t, "Header.LanguageCount", 0, erfData.Header.LanguageCount)
	test.ExpectUint32(t, "Header.LocalizedStringSize", 0, erfData.Header.LocalizedStringSize)
	test.ExpectUint32(t, "Header.EntryCount", 17, erfData.Header.EntryCount)
	test.ExpectUint32(t, "Header.OffsetToLocalizedString", 160, erfData.Header.OffsetToLocalizedString)
	test.ExpectUint32(t, "Header.OffsetToKeyList", 160, erfData.Header.OffsetToKeyList)
	test.ExpectUint32(t, "Header.OffsetToResourceList", 704, erfData.Header.OffsetToResourceList)
	test.ExpectUint32(t, "Header.DescriptionStrRef", 0xffffffff, erfData.Header.DescriptionStrRef)
	test.ExpectUint32(t, "Header.Reserved (bytes)", 116, uint32(len(erfData.Header.Reserved)))

	// localized string list
	test.ExpectUint32(t, "LocalizedStringList (elements)", 0, uint32(len(erfData.LocalizedStringList)))

	// key list
	test.ExpectUint32(t, "KeyList (elements)", 17, uint32(len(erfData.KeyList)))
	checkKeyListElement(t, erfData, 0, "barrowsinterior8", 0, file.Are)
	checkKeyListElement(t, erfData, 1, "barrowsinterior8", 1, file.Gic)
	checkKeyListElement(t, erfData, 2, "barrowsinterior8", 2, file.Git)
	checkKeyListElement(t, erfData, 3, "beholdercaves10x", 3, file.Are)
	checkKeyListElement(t, erfData, 4, "beholdercaves10x", 4, file.Gic)
	checkKeyListElement(t, erfData, 5, "beholdercaves10x", 5, file.Git)
	checkKeyListElement(t, erfData, 6, "creaturepalcus", 6, file.Itp)
	checkKeyListElement(t, erfData, 7, "doorpalcus", 7, file.Itp)
	checkKeyListElement(t, erfData, 8, "encounterpalcus", 8, file.Itp)
	checkKeyListElement(t, erfData, 9, "itempalcus", 9, file.Itp)
	checkKeyListElement(t, erfData, 10, "module", 10, file.Ifo)
	checkKeyListElement(t, erfData, 11, "placeablepalcus", 11, file.Itp)
	checkKeyListElement(t, erfData, 12, "Repute", 12, file.Fac)
	checkKeyListElement(t, erfData, 13, "soundpalcus", 13, file.Itp)
	checkKeyListElement(t, erfData, 14, "storepalcus", 14, file.Itp)
	checkKeyListElement(t, erfData, 15, "triggerpalcus", 15, file.Itp)
	checkKeyListElement(t, erfData, 16, "waypointpalcus", 16, file.Itp)

	// resource list
	test.ExpectUint32(t, "ResourceList (bytes)", erfData.Header.EntryCount, uint32(len(erfData.ResourceList)))

	// resource data
	file, _ := os.Open(filePath)
	defer file.Close()

	stat, _ := file.Stat()
	toSkip := erfData.Header.OffsetToResourceList + erfData.Header.EntryCount*8

	test.ExpectUint32(
		t,
		"ResourceData (bytes)",
		uint32(stat.Size()-int64(toSkip)),
		uint32(len(erfData.ResourceData)))
}
