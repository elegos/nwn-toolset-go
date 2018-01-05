package are

import (
	"aurora/file/gff"
	"errors"
	"fmt"
)

// Header structure of the ARE file format's header
type Header gff.Header

/*
	These flags affect game behaviour with respect to
	ability to hear things behind walls, map exploration
	visibility, and whether certain feats are active, though
	not necessarily in that order. They do not affect how the
	toolset presents the area to the user.
*/
const (
	FlagInterior    uint8 = 1 // Exterior if unset
	FlagUnderground uint8 = 2 // Aboveground if unsed
	FlagNatural     uint8 = 4 // Urban if unset
)

// AreaTile a tile element in the area
type AreaTile struct {
	AnimLoop1   int8  // Boolean value to indicate if the animation in the tile model should play (1) or not
	AnimLoop2   int8  // Boolean value to indicate if the animation in the tile model should play (1) or not
	AnimLoop3   int8  // Boolean value to indicate if the animation in the tile model should play (1) or not
	Height      int32 // Number of height transitions that this tile is located at. Should never be negative.
	ID          int32 // Index into the tileset file's list of tiles, to specify what tile to use
	MainLight1  int8  // Index into lightcolor.2da to specify mainlight 1 color on the tile. 0 = off
	MainLight2  int8  // Index into lightcolor.2da to specify mainlight 2 color on the tile. 0 = off
	Orientation int32 // Counterclockwise: 0 = 0°; 1 = 90°, 2 = 180°, 3 = 270°
	SrcLight1   int8  // 0 if SourceLight is off or does not exist. 1-15 to specify color and animation of sourcelight.
	SrcLight2   int8  // 0 if SourceLight is off or does not exist. 1-15 to specify color and animation of sourcelight.
}

// Data the ARE's data struct (Top-level struct)
type Data struct {
	ID        int32  // Deprecated; unused. Always -1.
	CreatorID int32  // Deprecated; unused. Always -1.
	Version   uint32 // Revision number of the area, starting from 1 and incrementing every time the ARE file is saved.

	ChanceLightning  int32      // Percent chance of lightning (0-100)
	ChanceRain       int32      // Percent chance of rain (0-100)
	ChanceSnow       int32      // Percent chance of snow (0-100)
	DayNightCycle    int8       // 1 if day/night transitions occur, 0 otherwise
	Flags            uint32     // Set of bit flags specifying area terrain type: 0x0001 (interior)
	FogClipDist      float32    // Fog Clip Distance (m)
	Height           int32      // Area size in the y-direction (north-sourh direction) measured in number of tiles
	IsNight          int8       // 1 if the area is always night, 0 if area is always day. Meaningful only if DayNightCycle is 0.
	LightingScheme   int8       // Index into environment.2da
	LoadScreenID     uint16     // Index into loadscreens.2da. Default loading screen to use when loading this area.
	ModListenCheck   int32      // Modifier to Listen akill checks made in area
	ModSpotCheck     int32      // Modifier to Spot skill checks made in area
	MoonAmbientColor uint32     // Nighttime ambient light color (BGR format)
	MoonDiffuseColor uint32     // Nighttime diffuse light color (BGR format)
	MoonFogAmount    int8       // Nighttime fog amount (0-15)
	MoonFogColor     uint32     // Nighttime fog color (BGR format)
	MoonShadows      int8       // 1 if shadows appear at night, 0 otherwise
	NoRest           int8       // 1 if resting is not allowed, 0 otherwise
	OnEnter          string     // OnEnter event
	OnExit           string     // OnExit event
	OnHeartbeat      string     // OnHeartbeat event
	OnUserDefined    string     // OnUserDefined event
	PlayerVsPlayer   int8       // Index into pvpsettings.2da. Data is hard-coded in-game, so changes to 2da make no difference.
	ResRef           string     // Should be identical to the filename of the area
	ShadowOpacity    int8       // Opacity of shadows (0-100)
	SkyBox           int8       // Index into skyboxes.2da (0-255). 0 means no skybox.
	SunAmbientColor  uint32     // Daytime ambient light color (BGR format)
	SunDiffuseColor  uint32     // Daytime diffuse light color (BGR format)
	SunFogAmount     int8       // Daytime fog amount (0-15)
	SunFogColor      uint32     // Daytime fog color (BGR format)
	SunShadows       int8       // 1 if shadows appear during the day, 0 otherwise
	Tileset          string     // ResRef of the tileset (.SET) file used by the area
	Width            int32      // Area size in the x-direction (west-east direction) measured in number of tiles
	WindPower        int32      // Strength of the wind in the area. None, weak or strong (0-2)
	TileList         []AreaTile // List of AreaTiles used in the area

	Tag      string // Tag of the area, used for scripting
	Comments string // Module designer comments
	// Name of area as seen in game and in left-hand module contents treeview in toolset.
	// If there is a colon (:) in the name, then the game does
	// not show any of the text up to and including the first colon.
	Name gff.CExoLocString
}

// ARE Area file format
type ARE struct {
	Header Header
	Data   Data
}

// FromBytes read an ARE file from its bytes
func FromBytes(bytes []byte) (ARE, error) {
	var result = ARE{}
	var gffFile = gff.FromBytes(bytes)

	if gffFile.Header.FileType != "ARE " {
		return result, fmt.Errorf("Not an ARE file. Found: '%s'", gffFile.Header.FileType)
	}

	if gffFile.Header.FileVersion != "V3.2" {
		return result, fmt.Errorf("Unsupported file version, expected 'V3.2', found '%s'", gffFile.Header.FileVersion)
	}

	if gffFile.Header.StructCount == 0 {
		return result, errors.New("Not a valid ARE file, there should be at least one struct")
	}

	result.Header = Header(gffFile.Header)

	// Read the top-level struct
	var topLevelStruct = gffFile.StructArray[0]
	var fields, err = gff.ExtractStructFields(topLevelStruct, gffFile)

	if err != nil {
		return result, err
	}

	var errorBag = gff.ErrorBag{}
	result.Data = Data{
		ID:        gff.GetFieldInt32Value("ID", fields, &errorBag),
		CreatorID: gff.GetFieldInt32Value("Creator_ID", fields, &errorBag),

		// uint32
		Flags:            gff.GetFieldUint32Value("Flags", fields, &errorBag),
		MoonAmbientColor: gff.GetFieldUint32Value("MoonAmbientColor", fields, &errorBag),
		MoonDiffuseColor: gff.GetFieldUint32Value("MoonDiffuseColor", fields, &errorBag),
		MoonFogColor:     gff.GetFieldUint32Value("MoonFogColor", fields, &errorBag),
		SunAmbientColor:  gff.GetFieldUint32Value("SunAmbientColor", fields, &errorBag),
		SunDiffuseColor:  gff.GetFieldUint32Value("SunDiffuseColor", fields, &errorBag),
		SunFogColor:      gff.GetFieldUint32Value("SunFogColor", fields, &errorBag),
		Version:          gff.GetFieldUint32Value("Version", fields, &errorBag),

		// int32
		ChanceLightning: gff.GetFieldInt32Value("ChanceLightning", fields, &errorBag),
		ChanceRain:      gff.GetFieldInt32Value("ChanceRain", fields, &errorBag),
		ChanceSnow:      gff.GetFieldInt32Value("ChanceSnow", fields, &errorBag),
		Height:          gff.GetFieldInt32Value("Height", fields, &errorBag),
		ModListenCheck:  gff.GetFieldInt32Value("ModListenCheck", fields, &errorBag),
		ModSpotCheck:    gff.GetFieldInt32Value("ModSpotCheck", fields, &errorBag),
		Width:           gff.GetFieldInt32Value("Width", fields, &errorBag),
		WindPower:       gff.GetFieldInt32Value("WindPower", fields, &errorBag),

		// uint16
		LoadScreenID: gff.GetFieldUint16Value("LoadScreenID", fields, &errorBag),

		// int8
		DayNightCycle:  gff.GetFieldByteValue("DayNightCycle", fields, &errorBag),
		IsNight:        gff.GetFieldByteValue("IsNight", fields, &errorBag),
		LightingScheme: gff.GetFieldByteValue("LightingScheme", fields, &errorBag),
		MoonFogAmount:  gff.GetFieldByteValue("MoonFogAmount", fields, &errorBag),
		MoonShadows:    gff.GetFieldByteValue("MoonShadows", fields, &errorBag),
		NoRest:         gff.GetFieldByteValue("NoRest", fields, &errorBag),
		PlayerVsPlayer: gff.GetFieldByteValue("PlayerVsPlayer", fields, &errorBag),
		ShadowOpacity:  gff.GetFieldByteValue("ShadowOpacity", fields, &errorBag),
		SkyBox:         gff.GetFieldByteValue("SkyBox", fields, &errorBag),
		SunFogAmount:   gff.GetFieldByteValue("SunFogAmount", fields, &errorBag),
		SunShadows:     gff.GetFieldByteValue("SunShadows", fields, &errorBag),

		// float32
		FogClipDist: gff.GetFieldFloatValue("FogClipDist", fields, &errorBag),

		// CExoString
		Tag:      gff.GetFieldCExoStringValue("Tag", fields, &errorBag),
		Comments: gff.GetFieldCExoStringValue("Comments", fields, &errorBag),

		// CExoLocString
		Name: gff.GetFieldCExoLocStringValue("Name", fields, &errorBag),

		// ResRef
		ResRef:        gff.GetFieldCResRefValue("ResRef", fields, &errorBag),
		OnEnter:       gff.GetFieldCResRefValue("OnEnter", fields, &errorBag),
		OnExit:        gff.GetFieldCResRefValue("OnExit", fields, &errorBag),
		OnHeartbeat:   gff.GetFieldCResRefValue("OnHeartbeat", fields, &errorBag),
		OnUserDefined: gff.GetFieldCResRefValue("OnUserDefined", fields, &errorBag),
		Tileset:       gff.GetFieldCResRefValue("Tileset", fields, &errorBag),
	}

	var tileListByteOffset = gff.GetFieldListByteOffsetValue("Tile_List", fields, &errorBag)
	result.Data.TileList = extractAreaTilesList(tileListByteOffset, gffFile, &errorBag)

	return result, errorBag.Error
}
