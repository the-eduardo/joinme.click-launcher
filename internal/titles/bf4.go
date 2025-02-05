package titles

import (
	"github.com/cetteup/joinme.click-launcher/internal/domain"
	"github.com/cetteup/joinme.click-launcher/internal/platforms"
	"github.com/cetteup/joinme.click-launcher/internal/titles/internal"
	"github.com/cetteup/joinme.click-launcher/pkg/game_launcher"
	"github.com/cetteup/joinme.click-launcher/pkg/software_finder"
)

const (
	bf4Exe = "bf4.exe"
)

var Bf4 = domain.GameTitle{
	Name:           "Battlefield 4",
	ProtocolScheme: "bf4",
	PlatformClient: &platforms.OriginClient,
	FinderConfigs: []software_finder.Config{
		{
			ForType:           software_finder.RegistryFinder,
			RegistryKey:       software_finder.RegistryKeyLocalMachine,
			RegistryPath:      "SOFTWARE\\WOW6432Node\\EA Games\\Battlefield 4",
			RegistryValueName: "Install Dir",
		},
	},
	LauncherConfig: game_launcher.Config{
		ExecutableName: bf4Exe,
		HookConfigs: []game_launcher.HookConfig{
			{
				Handler:     internal.HookKillProcess,
				When:        game_launcher.HookWhenPreLaunch,
				ExitOnError: true,
			},
		},
	},
	URLValidator: internal.MakePatternURLValidator(internal.Frostbite3GameIdPattern),
	CmdBuilder:   internal.MakeOriginCmdBuilder("1007968", "1011575", "1011576", "1011577", "1010268", "1010269", "1010270", "1010271", "1010958", "1010959", "1010960", "1010961", "1007077", "1016751", "1016757", "1016754", "1015365", "1015364", "1015363", "1015362"),
	HookHandlers: []game_launcher.HookHandler{
		internal.MakeKillProcessHookHandler(false, bf4Exe), // Launcher config executable name will be "Origin.exe", which we don't want to kill
	},
}
