package file

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

// LanguageLookup from constant to name
var LanguageLookup = map[Language]string{
	LangEnglish:            "English",
	LangFrench:             "French",
	LangGerman:             "German",
	LangItalian:            "Italian",
	LangSpanish:            "Spanish",
	LangPolish:             "Polish",
	LangKorean:             "Korean",
	LangChineseTraditional: "Chinese traditional",
	LangChineseSimplified:  "Chinese simplified",
	LangJapanese:           "Japanese",
}
