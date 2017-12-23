package file

// ResourceType file resource type (i.e. in erf files)
type ResourceType uint16

const (
	// Invalid invalid resource type
	Invalid ResourceType = 0xFFFF
	// Bmp (binary) Windows BMP file
	Bmp ResourceType = 1
	// Tga (binary) TGA image format
	Tga ResourceType = 3
	// Wav (binary) WAV sound file
	Wav ResourceType = 4
	// Plt (binary) Bioware Packed Layered Texture, used for player character skins,
	// allows for multiple color layers
	Plt ResourceType = 6
	// Ini (text, ini) Windows INI file format
	Ini ResourceType = 7
	// Txt (text) Text file
	Txt ResourceType = 10
	// Mdl (mdl) Aurora model
	Mdl ResourceType = 2002
	// Nss (text) NWScript Source
	Nss ResourceType = 2009
	// Ncs (binary) NWScript Compiled Script
	Ncs ResourceType = 2010
	// Are (gff) BioWare Aurora Engine Area file. Contains information on what tiles are
	// located in an area, as well as other static area properties that cannot change
	// via scripting. For each .are file in a .mod, there must also be a corresponding
	// .git and .gic file having the same ResRef.
	Are ResourceType = 2012
	// Set (text, ini) BioWare Aurora Engine Tileset
	Set ResourceType = 2013
	// Ifo (gff) Module Info File. See the IFO format document.
	Ifo ResourceType = 2014
	// Bic (gff) Character/Creature
	Bic ResourceType = 2015
	// Wok (mdl) Walkmesh
	Wok ResourceType = 2016
	// T2da (text) 2-D Array
	T2da ResourceType = 2017
	// Txi (text) Extra Texture Info
	Txi ResourceType = 2022
	// Git (gff) Game Instance File. Contains information for all object
	// instances in an aread, and all area properties that can change via scripting
	Git ResourceType = 2023
	// Uti (gff) Item Blueprint
	Uti ResourceType = 2025
	// Utc (gff) Creature Blueprint
	Utc ResourceType = 2027
	// Dlg (gff) Conversation File
	Dlg ResourceType = 2029
	// Itp (gff) Tile/Blueprint Palette File
	Itp ResourceType = 2030
	// Utt (gff) Trigger Blueprint
	Utt ResourceType = 2032
	// Dds (binary) Compressed Texture file
	Dds ResourceType = 2033
	// Uts (gff) Sound Blueprint
	Uts ResourceType = 2035
	// Ltr (binary) Letter-combo probability info for name generation
	Ltr ResourceType = 2036
	// Gff (gff) Generic File FOrmat. Used when undesirable to create a new file extension
	// for a resource, but the resource is a GFF (Examples of GFFs include itp, utc,
	// uti, ifo, are, git)
	Gff ResourceType = 2037
	// Fac (gff) Faction File
	Fac ResourceType = 2038
	// Ute (gff) Encounter Blueprint
	Ute ResourceType = 2040
	// Utd (gff) Door Blueprint
	Utd ResourceType = 2042
	// Utp (gff) Placeable Object Blueprint
	Utp ResourceType = 2044
	// Dft (text, ini) Default Values file. Used by area properties dialog
	Dft ResourceType = 2045
	// Gic (gff) Game Instance Comments. Comments on instances are not used by
	// the game, only the toolset, so they are stored in a gic instead of in
	// the git with the other instance properties..
	Gic ResourceType = 2046
	// Gui (gff) Graphical User Interface layout used by game
	Gui ResourceType = 2047
	// Utm (gff) Store/Merchant Blueprint
	Utm ResourceType = 2051
	// Dwk (mdl) Door walkmesh
	Dwk ResourceType = 2052
	// Pwk (mdl) Placeable Object walkmesh
	Pwk ResourceType = 2053
	// Jrl (gff) Journal File
	Jrl ResourceType = 2056
	// Utw (gff) Waypoint Blueprint. See Waypoint GFF document.
	Utw ResourceType = 2058
	// Ssf (binary) Sound Set File. See Sound Set File Format document
	Ssf ResourceType = 2060
	// Ndb (binary) Script Debugger File
	Ndb ResourceType = 2064
	// Ptm (gff) Plot Manager file/Plot Instance
	Ptm ResourceType = 2065
	// Ptt (gff) Plot Wizard Blueprint
	Ptt ResourceType = 2066
)

// ResourceTypeLookup from constant to name
var ResourceTypeLookup = map[ResourceType]string{
	Invalid: "Invalid",
	Bmp:     "bmp",
	Tga:     "tga",
	Wav:     "wav",
	Plt:     "plt",
	Ini:     "ini",
	Txt:     "txt",
	Mdl:     "mdl",
	Nss:     "nss",
	Ncs:     "ncs",
	Are:     "are",
	Set:     "set",
	Ifo:     "ifo",
	Bic:     "bic",
	Wok:     "wok",
	T2da:    "2da",
	Txi:     "txi",
	Git:     "git",
	Uti:     "uti",
	Utc:     "utc",
	Dlg:     "dlg",
	Itp:     "itp",
	Utt:     "utt",
	Dds:     "dds",
	Uts:     "uts",
	Ltr:     "ltr",
	Gff:     "gff",
	Fac:     "fac",
	Ute:     "ute",
	Utd:     "utd",
	Utp:     "utp",
	Dft:     "dft",
	Gic:     "gic",
	Gui:     "gui",
	Utm:     "utm",
	Dwk:     "dwk",
	Pwk:     "pwk",
	Jrl:     "jrl",
	Utw:     "utw",
	Ssf:     "ssf",
	Ndb:     "ndb",
	Ptm:     "ptm",
	Ptt:     "ptt",
}
