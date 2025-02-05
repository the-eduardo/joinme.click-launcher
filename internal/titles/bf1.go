package titles

import (
	"github.com/cetteup/joinme.click-launcher/internal/domain"
	"github.com/cetteup/joinme.click-launcher/internal/platforms"
	"github.com/cetteup/joinme.click-launcher/internal/titles/internal"
	"github.com/cetteup/joinme.click-launcher/pkg/game_launcher"
	"github.com/cetteup/joinme.click-launcher/pkg/software_finder"
)

const (
	bf1Exe = "bf1.exe"
)

var Bf1 = domain.GameTitle{
	Name:           "Battlefield 1",
	ProtocolScheme: "bf1",
	PlatformClient: &platforms.EaClient,
	FinderConfigs: []software_finder.Config{
		{
			ForType:           software_finder.RegistryFinder,
			RegistryKey:       software_finder.RegistryKeyLocalMachine,
			RegistryPath:      "SOFTWARE\\WOW6432Node\\EA Games\\Battlefield 1",
			RegistryValueName: "Install Dir",
		},
	},
	LauncherConfig: game_launcher.Config{
		ExecutableName: bf1Exe,
		HookConfigs: []game_launcher.HookConfig{
			{
				Handler:     internal.HookKillProcess,
				When:        game_launcher.HookWhenPreLaunch,
				ExitOnError: true,
			},
		},
	},
	URLValidator: internal.MakePatternURLValidator(internal.Frostbite3GameIdPattern),
	CmdBuilder:   internal.MakeOriginCmdBuilder("1026023"),
	HookHandlers: []game_launcher.HookHandler{
		internal.MakeKillProcessHookHandler(false, bf1Exe), // Launcher config executable name will be "Origin.exe", which we don't want to kill
	},
}
