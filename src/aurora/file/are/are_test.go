package are_test

import (
	"aurora/file"
	"aurora/file/are"
	"aurora/tools/test"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestFromBytes_will_error_wrong_type(t *testing.T) {
	fileBytes, _ := ioutil.ReadFile("../erf/test/module.mod")
	_, err := are.FromBytes(fileBytes)

	if err == nil {
		t.Error("Expected error, none found")
	}
}

func TestFromBytes_will_success(t *testing.T) {
	fileBytes, _ := ioutil.ReadFile("./test/barrowsinterior8.are")
	areFile, err := are.FromBytes(fileBytes)

	if err != nil {
		t.Errorf("Unexpected error, found: %v", err)
	}

	test.ExpectInt32(t, "Data.ChanceLightning", 0, areFile.Data.ChanceLightning)
	test.ExpectInt32(t, "Data.ChanceRain", 0, areFile.Data.ChanceRain)
	test.ExpectInt32(t, "Data.ChanceSnow", 0, areFile.Data.ChanceSnow)
	test.ExpectString(t, "Data.Comments", "This is the comment of the area.", areFile.Data.Comments)
	test.ExpectInt32(t, "Data.CreatorID", -1, areFile.Data.CreatorID)
	test.ExpectInt8(t, "Data.DayNightCycle", 0, areFile.Data.DayNightCycle)
	test.ExpectUint32(t, "Data.Flags", uint32(are.FlagInterior|are.FlagNatural|are.FlagUnderground), areFile.Data.Flags)
	test.ExpectFloat32(t, "Data.FogClipDist", 45.0, areFile.Data.FogClipDist)
	test.ExpectInt32(t, "Data.Height", 8, areFile.Data.Height)
	test.ExpectInt32(t, "Data.ID", -1, areFile.Data.ID)
	test.ExpectInt8(t, "Data.IsNight", 1, areFile.Data.IsNight)
	test.ExpectInt8(t, "Data.LightingScheme", 25, areFile.Data.LightingScheme)
	test.ExpectUint16(t, "Data.LoadScreenID", 0, areFile.Data.LoadScreenID)
	test.ExpectInt32(t, "Data.ModListenCheck", 0, areFile.Data.ModListenCheck)
	test.ExpectInt32(t, "Data.ModSpotCheck", 0, areFile.Data.ModSpotCheck)
	test.ExpectUint32(t, "Data.MoonAmbientColor", 4679544, areFile.Data.MoonAmbientColor)
	test.ExpectUint32(t, "Data.MoonDiffuseColor", 6457991, areFile.Data.MoonDiffuseColor)
	test.ExpectInt8(t, "Data.MoonFogAmount", 8, areFile.Data.MoonFogAmount)
	test.ExpectUint32(t, "Data.MoonFogColor", 0, areFile.Data.MoonFogColor)
	test.ExpectInt8(t, "Data.MoonShadows", 0, areFile.Data.MoonShadows)
	test.ExpectUint32(t, "Data.Name.StringRef", 0xffffffff, areFile.Data.Name.StringRef)
	test.ExpectInt32(t, "Data.Name.SubStrings (len)", 1, int32(len(areFile.Data.Name.SubStrings)))
	test.ExpectString(t, "Data.Name.SubStrings[0].String", "area title", areFile.Data.Name.SubStrings[0].String)
	test.ExpectBool(t, "Data.Name.SubStrings[0].Masculine", false, areFile.Data.Name.SubStrings[0].Masculine)
	test.ExpectUint32(t, "Data.Name.SubStrings[0].Language", uint32(file.LangEnglish), uint32(areFile.Data.Name.SubStrings[0].LanguageID))
	test.ExpectInt8(t, "Data.NoRest", 0, areFile.Data.NoRest)
	test.ExpectString(t, "Data.OnEnter", "", areFile.Data.OnEnter)
	test.ExpectString(t, "Data.OnExit", "", areFile.Data.OnExit)
	test.ExpectString(t, "Data.OnHeartbeat", "", areFile.Data.OnHeartbeat)
	test.ExpectString(t, "Data.OnUserDefined", "", areFile.Data.OnUserDefined)
	test.ExpectInt8(t, "Data.PlayerVsPlayer", 3, areFile.Data.PlayerVsPlayer)
	test.ExpectString(t, "Data.ResRef", "barrowsinterior8", areFile.Data.ResRef)
	test.ExpectInt8(t, "Data.ShadowOpacity", 60, areFile.Data.ShadowOpacity)
	test.ExpectInt8(t, "Data.SkyBox", 0, areFile.Data.SkyBox)
	test.ExpectUint32(t, "Data.SunAmbientColor", 0, areFile.Data.SunAmbientColor)
	test.ExpectUint32(t, "Data.SunDiffuseColor", 0, areFile.Data.SunDiffuseColor)
	test.ExpectInt8(t, "Data.SunFogAmount", 0, areFile.Data.SunFogAmount)
	test.ExpectUint32(t, "Data.SunFogColor", 0, areFile.Data.SunFogColor)
	test.ExpectInt8(t, "Data.SunShadows", 0, areFile.Data.SunShadows)
	test.ExpectString(t, "Data.Tag", "barrowsinterior8x8", areFile.Data.Tag)
	test.ExpectUint32(t, "Data.TileList (len)", 64, uint32(len(areFile.Data.TileList)))
	test.ExpectString(t, "Data.Tileset", "tbw01", areFile.Data.Tileset)
	test.ExpectUint32(t, "Data.Version", 3, areFile.Data.Version)
	test.ExpectInt32(t, "Data.Width", 8, areFile.Data.Width)
	test.ExpectInt32(t, "Data.WindPower", 0, areFile.Data.WindPower)

	// Tiles tests
	var expectedTileList = []are.AreaTile{
		are.AreaTile{AnimLoop1: 0, AnimLoop2: 0, AnimLoop3: 0, Height: 0, ID: 0, MainLight1: 0, MainLight2: 13, Orientation: 2, SrcLight1: 3, SrcLight2: 3},
		are.AreaTile{AnimLoop1: 0, AnimLoop2: 0, AnimLoop3: 0, Height: 0, ID: 0, MainLight1: 4, MainLight2: 0, Orientation: 3, SrcLight1: 0, SrcLight2: 0},
		are.AreaTile{AnimLoop1: 0, AnimLoop2: 0, AnimLoop3: 0, Height: 0, ID: 0, MainLight1: 4, MainLight2: 0, Orientation: 3, SrcLight1: 0, SrcLight2: 0},
		are.AreaTile{AnimLoop1: 0, AnimLoop2: 0, AnimLoop3: 0, Height: 0, ID: 0, MainLight1: 0, MainLight2: 14, Orientation: 2, SrcLight1: 2, SrcLight2: 2},
		are.AreaTile{AnimLoop1: 0, AnimLoop2: 0, AnimLoop3: 0, Height: 0, ID: 0, MainLight1: 30, MainLight2: 13, Orientation: 3, SrcLight1: 3, SrcLight2: 3},
		// ...
	}

	for i := 0; i < len(expectedTileList); i++ {
		test.ExpectInt8(t, fmt.Sprintf("Data.TileList[%d].AnimLoop1", i), expectedTileList[i].AnimLoop1, areFile.Data.TileList[i].AnimLoop1)
		test.ExpectInt8(t, fmt.Sprintf("Data.TileList[%d].AnimLoop2", i), expectedTileList[i].AnimLoop3, areFile.Data.TileList[i].AnimLoop2)
		test.ExpectInt8(t, fmt.Sprintf("Data.TileList[%d].AnimLoop3", i), expectedTileList[i].AnimLoop2, areFile.Data.TileList[i].AnimLoop3)
		test.ExpectInt32(t, fmt.Sprintf("Data.TileList[%d].Height", i), expectedTileList[i].Height, areFile.Data.TileList[i].Height)
		test.ExpectInt32(t, fmt.Sprintf("Data.TileList[%d].ID", i), expectedTileList[i].ID, areFile.Data.TileList[i].ID)
		test.ExpectInt8(t, fmt.Sprintf("Data.TileList[%d].MainLight1", i), expectedTileList[i].MainLight1, areFile.Data.TileList[i].MainLight1)
		test.ExpectInt8(t, fmt.Sprintf("Data.TileList[%d].MainLight1", i), expectedTileList[i].MainLight2, areFile.Data.TileList[i].MainLight2)
		test.ExpectInt32(t, fmt.Sprintf("Data.TileList[%d].Orientation", i), expectedTileList[i].Orientation, areFile.Data.TileList[i].Orientation)
		test.ExpectInt8(t, fmt.Sprintf("Data.TileList[%d].SrcLight1", i), expectedTileList[i].SrcLight1, areFile.Data.TileList[i].SrcLight1)
		test.ExpectInt8(t, fmt.Sprintf("Data.TileList[%d].SrcLight2", i), expectedTileList[i].SrcLight2, areFile.Data.TileList[i].SrcLight2)
	}
}
