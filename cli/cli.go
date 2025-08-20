package cli

import (
	"flag"
	"fmt"
	"github.com/unknown321/fuse/playerresource"
	"github.com/unknown321/fuse/server"
	"github.com/unknown321/fuse/sessionmanager"
	"github.com/unknown321/fuse/util"
	"log/slog"
	"os"
	"runtime/debug"
	"strings"
)

func Start() {
	var err error

	var standalone bool
	var baseURL string
	var listenAddress string
	var platform string
	var writeLog bool
	var passThrough bool
	var clientURL string
	flag.StringVar(&baseURL, "baseURL", "http://127.0.0.1:6667/", "base url")
	flag.StringVar(&listenAddress, "listenAddr", "127.0.0.1:6667", "listen addr")
	flag.StringVar(&platform, "platform", "tppstm", "server platform")
	flag.BoolVar(&writeLog, "writeLog", false, "save requests and responses to files")
	flag.BoolVar(&passThrough, "passThrough", false, "work as a proxy between client and original master server")
	flag.BoolVar(&standalone, "standalone", false, "standalone mode (do not attempt to patch and backup executable)")
	flag.StringVar(&clientURL, "clientURL", "http://127.0.0.1:6667/", "user URL (will be patched into the exe)")
	flag.Parse()

	opts := &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, opts)))
	slog.Info("starting")

	bf, ok := debug.ReadBuildInfo()
	if !ok {
		slog.Error("cannot read build info")
		os.Exit(1)
	}

	commit := ""
	commitDate := ""
	dirty := false

	for _, k := range bf.Settings {
		switch k.Key {
		case "vcs.revision":
			commit = k.Value
		case "vcs.time":
			commitDate = k.Value
		case "vcs.modified":
			dirty = k.Value == "true"
		default:
		}
	}

	slog.Info("version", "commit", commit, "date", commitDate, "dirty", dirty)

	if !standalone {
		if err = util.Backup(); err != nil {
			slog.Error("backup", "error", err.Error())
			os.Exit(1)
		}

		customURL := false
		if clientURL != "http://127.0.0.1:6667/" {
			customURL = true
		}

		if err = util.Patch(clientURL); err != nil {
			slog.Error("patch", "error", err.Error())
			os.Exit(1)
		}

		slog.Info("patched executable", "server url", strings.TrimSpace(clientURL))

		if customURL {
			slog.Info("custom server url provided, not starting local server")
			slog.Info("You can start the game now")
			slog.Info("Press Enter to close this window")

			var a []byte
			_, _ = fmt.Scanln(&a)
			os.Exit(0)
		}
	}

	slog.Info("starting server", "listen address", listenAddress)

	bonus := sessionmanager.SignupBonus{
		GMP: 25_000_000,
		Resources: playerresource.PlayerResource{
			Raw: playerresource.Raw{
				Fuel:          500_000,
				Bio:           500_000,
				CommonMetal:   500_000,
				MinorMetal:    500_000,
				PreciousMetal: 500_000,
			},
			Processed: playerresource.Processed{
				Fuel:          500_000,
				Bio:           500_000,
				CommonMetal:   500_000,
				MinorMetal:    500_000,
				PreciousMetal: 500_000,
			},
			Plants: playerresource.Plants{
				Wormwood:          10_000,
				BlackCarrot:       10_000,
				GoldenCrescent:    10_000,
				Tarragon:          10_000,
				AfricanPeach:      10_000,
				DigitalisPurpurea: 10_000,
				DigitalisLutea:    10_000,
				Haoma:             10_000,
			},
			Vehicles: playerresource.Vehicles{
				ZaAZS84:       100,
				APET41LV:      100,
				ZiGRA6T:       100,
				Boar53CT:      100,
				ZhukBr3:       100,
				StoutIfvSc:    100,
				ZhukRsZo:      100,
				StoutIfvFs:    100,
				TT77Nosorog:   100,
				M84AMagloader: 100,
			},
			WalkerGears: playerresource.WalkerGears{
				WGPP:        50,
				CCCPWGTypeC: 50,
				CCCPWGTypeA: 50,
				CFAWGTypeC:  50,
				CFAWGTypeA:  50,
			},
			Nuclear: playerresource.Nuclear{
				Weapon: 10,
				Waste:  10,
			},
			Parasites: playerresource.Parasites{
				Mist:       20,
				Camouflage: 20,
				Armor:      20,
			},
			Mystery01: 0,
			Placed: playerresource.Placed{
				VolgaK12:     100,
				HMG3Wingate:  100,
				M2A304Mortar: 100,
				Zhizdra45:    100,
				M276AAGGun:   100,
			},
			Mystery02: [7]int{},
			Unused:    [13]int{},
		},
	}

	server.Start(baseURL, listenAddress, platform, writeLog, passThrough, server.DsnURIDefault, bonus)
}
