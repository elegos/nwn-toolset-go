package erf

import (
	"aurora/tools"
	"io/ioutil"
	"os"
)

// Language the language ids
type Language uint32

const (
	// LangEnglish English language id
	LangEnglish Language = 0
	// LangFrench French language id
	LangFrench Language = 1
	// LangGerman German language id
	LangGerman Language = 2
	// LangItalian Italian language id
	LangItalian Language = 3
	// LangSpanish Spanish language id
	LangSpanish Language = 4
	// LangPolish Polish language id
	LangPolish Language = 5
	// LangKorean Korean language id
	LangKorean Language = 128
	// LangChineseTraditional Chinese Trasitional language id
	LangChineseTraditional Language = 129
	// LangChineseSimplified Chinese Simplified language id
	LangChineseSimplified Language = 130
	// LangJapanese Japanese language id
	LangJapanese Language = 131
)

// Header structure of the ERF file format's header
type Header struct {
	FileType                string    // 4 bytes, "ERF ", "MOD ", "SAV ", "HACK "
	Version                 string    // 4 bytes, V1.0
	LanguageCount           uint32    // number of strings in the Localized String
	LocalizedStringSize     uint32    // total size (bytes) of Localized String Table
	EntryCount              uint32    // number of files packed into the ERF
	OffsetToLocalizedString uint32    // from beginning of file
	OffsetToKeyList         uint32    // from beginning of file
	OffsetToResourceList    uint32    // from beginning of file
	BuildYear               uint32    // since 1900
	BuildDay                uint32    // since January 1st
	DescriptionStrRef       uint32    // strref for file description
	Reserved                [116]byte // 116 bytes of reserved space for future ELF versions
}

// LocalizedStringElement the element of ERF's LocalizedStringList
type LocalizedStringElement struct {
	LanguageID Language // The language's id
	StringSize uint32   // Length of  the String
	String     string   // Variable size as specified by StringSize field
}

// KeyElement the element of ERF's KeyList
type KeyElement struct {
	ResRef  string  // Filename, 16 chars
	ResID   uint32  // Resource ID, starts at 0 and increments
	ResType uint16  // File type, TODO list
	Unused  [2]byte // unused 2 bytes
}

// ResourceElement the element of ERF's ResourceList
type ResourceElement struct {
	OffsetToResource uint32 // Offset to file data from the beginning of the file
	ResourceSize     uint32 // Number of bytes
}

// ERF Encapsulated Resource File Format
type ERF struct {
	Header              Header
	LocalizedStringList []LocalizedStringElement
	KeyList             []KeyElement
	ResourceList        []ResourceElement
	ResourceData        []byte
}

// FromFile read the file and return an ERF struct
func FromFile(fileName string) ERF {
	var result = ERF{}

	_, err := ioutil.ReadFile(fileName)
	tools.EasyPanic(err)

	file, err := os.Open(fileName)
	tools.EasyPanic(err)
	defer file.Close()

	result.Header = extractHeader(file)
	result.LocalizedStringList = extractLocalizedStringList(file, result.Header.LocalizedStringSize)
	result.KeyList = extractKeyList(file, int64(result.Header.OffsetToKeyList), result.Header.EntryCount)
	result.ResourceList = extractResourceList(file, int64(result.Header.OffsetToResourceList), result.Header.EntryCount)
	result.ResourceData = extractResourceData(file, result.Header.OffsetToResourceList, result.Header.EntryCount)

	return result
}
