package resource

import (
	_ "embed"
	"fyne.io/fyne/v2"
)

//go:embed logo.png
var Logo []byte

//go:embed HarmonyOS_Sans_SC_Regular.ttf
var HMttf []byte

func GetLogo() fyne.Resource {
	return &fyne.StaticResource{
		StaticName:    "logo.png",
		StaticContent: Logo,
	}
}
