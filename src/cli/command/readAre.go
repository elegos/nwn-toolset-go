package command

import (
	"aurora/file/are"
	"cli/tools"
	"fmt"
	"io/ioutil"
)

// ReadAreFromFile read an ARE file's information
func ReadAreFromFile(file string) {
	bytes, err := ioutil.ReadFile(file)
	tools.EasyPanic(err)

	areFile, err := are.FromBytes(bytes)
	tools.EasyPanic(err)

	fmt.Println(fmt.Sprintf("%19s: %d", "ChanceLightning", areFile.Data.ChanceLightning))
	fmt.Println(fmt.Sprintf("%19s: %d", "ChanceRain", areFile.Data.ChanceRain))
	fmt.Println(fmt.Sprintf("%19s: %d", "ChanceSnow", areFile.Data.ChanceSnow))
	fmt.Println(fmt.Sprintf("%19s: %s", "Comments", areFile.Data.Comments))
	fmt.Println(fmt.Sprintf("%19s: %d", "CreatorID (unused)", areFile.Data.CreatorID))
	fmt.Println(fmt.Sprintf("%19s: %d", "DayNightCycle", areFile.Data.DayNightCycle))
	fmt.Print(fmt.Sprintf("%19s: ", "Flags (interior)"))
	if areFile.Data.Flags&uint32(are.FlagInterior) == uint32(are.FlagInterior) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	fmt.Print(fmt.Sprintf("%19s: ", "Flags (natural)"))
	if areFile.Data.Flags&uint32(are.FlagNatural) == uint32(are.FlagNatural) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	fmt.Print(fmt.Sprintf("%19s: ", "Flags (underground)"))
	if areFile.Data.Flags&uint32(are.FlagUnderground) == uint32(are.FlagUnderground) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	fmt.Println(fmt.Sprintf("%19s: %v", "Height", areFile.Data.Height))
	fmt.Println(fmt.Sprintf("%19s: %v", "ID", areFile.Data.ID))
	fmt.Println(fmt.Sprintf("%19s: %v", "IsNight", areFile.Data.IsNight))
	fmt.Println(fmt.Sprintf("%19s: %v", "LightingScheme", areFile.Data.LightingScheme))
	fmt.Println(fmt.Sprintf("%19s: %v", "LoadScreenID", areFile.Data.LoadScreenID))
	fmt.Println(fmt.Sprintf("%19s: %v", "ModListenCheck", areFile.Data.ModListenCheck))
	fmt.Println(fmt.Sprintf("%19s: %v", "ModSpotCheck", areFile.Data.ModSpotCheck))
	fmt.Println(fmt.Sprintf("%19s: %v", "MoonAmbientColor", areFile.Data.MoonAmbientColor))
	fmt.Println(fmt.Sprintf("%19s: %v", "MoonDiffuseColor", areFile.Data.MoonDiffuseColor))
	fmt.Println(fmt.Sprintf("%19s: %v", "MoonFogAmount", areFile.Data.MoonFogAmount))
	fmt.Println(fmt.Sprintf("%19s: %v", "MoonFogColor", areFile.Data.MoonFogColor))
	fmt.Println(fmt.Sprintf("%19s: %v", "MoonShadows", areFile.Data.MoonShadows))
	fmt.Println(fmt.Sprintf("%19s: %v", "Name", areFile.Data.Name))
	fmt.Println(fmt.Sprintf("%19s: %v", "NoRest", areFile.Data.NoRest))
	fmt.Println(fmt.Sprintf("%19s: %v", "OnEnter", areFile.Data.OnEnter))
	fmt.Println(fmt.Sprintf("%19s: %v", "OnExit", areFile.Data.OnExit))
	fmt.Println(fmt.Sprintf("%19s: %v", "OnHeartbeat", areFile.Data.OnHeartbeat))
	fmt.Println(fmt.Sprintf("%19s: %v", "OnUserDefined", areFile.Data.OnUserDefined))
	fmt.Println(fmt.Sprintf("%19s: %v", "PlayerVsPlayer", areFile.Data.PlayerVsPlayer))
	fmt.Println(fmt.Sprintf("%19s: %v", "ResRef", areFile.Data.ResRef))
	fmt.Println(fmt.Sprintf("%19s: %v", "SkyBox", areFile.Data.SkyBox))
	fmt.Println(fmt.Sprintf("%19s: %v", "ShadowOpacity", areFile.Data.ShadowOpacity))
	fmt.Println(fmt.Sprintf("%19s: %v", "SunAmbientColor", areFile.Data.SunAmbientColor))
	fmt.Println(fmt.Sprintf("%19s: %v", "SunDiffuseColor", areFile.Data.SunDiffuseColor))
	fmt.Println(fmt.Sprintf("%19s: %v", "SunFogAmount", areFile.Data.SunFogAmount))
	fmt.Println(fmt.Sprintf("%19s: %v", "SunFogColor", areFile.Data.SunFogColor))
	fmt.Println(fmt.Sprintf("%19s: %v", "SunShadows", areFile.Data.SunShadows))
	fmt.Println(fmt.Sprintf("%19s: %v", "Tag", areFile.Data.Tag))
	fmt.Println(fmt.Sprintf("%19s: %v", "TileList (length)", len(areFile.Data.TileList)))
	fmt.Println(fmt.Sprintf("%19s: %v", "TileSet", areFile.Data.Tileset))
	fmt.Println(fmt.Sprintf("%19s: %v", "Version", areFile.Data.Version))
	fmt.Println(fmt.Sprintf("%19s: %v", "Width", areFile.Data.Width))
	fmt.Println(fmt.Sprintf("%19s: %v", "WindPower", areFile.Data.WindPower))
}
