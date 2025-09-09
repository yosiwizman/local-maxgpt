package cli

import (
	cliContext "github.com/mudler/LocalAI/core/cli/context"
	"github.com/mudler/LocalAI/core/cli/worker"
)

var CLI struct {
	cliContext.Context `embed:""`

	Run             RunCMD             `cmd:"" help:"Run MaxGPT, this the default command if no other command is specified. Run 'maxgpt run --help' for more information" default:"withargs"`
	Federated       FederatedCLI       `cmd:"" help:"Run MaxGPT in federated mode"`
	Models          ModelsCMD          `cmd:"" help:"Manage MaxGPT models and definitions"`
	Backends        BackendsCMD        `cmd:"" help:"Manage MaxGPT backends and definitions"`
	TTS             TTSCMD             `cmd:"" help:"Convert text to speech"`
	SoundGeneration SoundGenerationCMD `cmd:"" help:"Generates audio files from text or audio"`
	Transcript      TranscriptCMD      `cmd:"" help:"Convert audio to text"`
	Worker          worker.Worker      `cmd:"" help:"Run workers to distribute workload (llama.cpp-only)"`
	Util            UtilCMD            `cmd:"" help:"Utility commands"`
	Explorer        ExplorerCMD        `cmd:"" help:"Run p2p explorer"`
}
