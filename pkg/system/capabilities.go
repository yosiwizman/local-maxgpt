package system

import (
	"os"
	"runtime"
	"strings"

	"github.com/mudler/LocalAI/pkg/xsysinfo"
	"github.com/rs/zerolog/log"
)

type SystemState struct {
	GPUVendor string
}

const (
	defaultCapability = "default"
	nvidiaL4T         = "nvidia-l4t"
	darwinX86         = "darwin-x86"
	metal             = "metal"
)

func (s *SystemState) Capability(capMap map[string]string) string {
	reportedCapability := s.getSystemCapabilities()

	// Check if the reported capability is in the map
	if _, exists := capMap[reportedCapability]; exists {
		log.Debug().Str("reportedCapability", reportedCapability).Any("capMap", capMap).Msg("Using reported capability")
		return reportedCapability
	}

	log.Debug().Str("reportedCapability", reportedCapability).Any("capMap", capMap).Msg("The requested capability was not found, using default capability")
	// Otherwise, return the default capability (catch-all)
	return defaultCapability
}

func (s *SystemState) getSystemCapabilities() string {
	if os.Getenv("MAXGPT_FORCE_META_BACKEND_CAPABILITY") != "" {
		log.Debug().Str("MAXGPT_FORCE_META_BACKEND_CAPABILITY", os.Getenv("MAXGPT_FORCE_META_BACKEND_CAPABILITY")).Msg("Using forced capability")
		return os.Getenv("MAXGPT_FORCE_META_BACKEND_CAPABILITY")
	}

	capabilityRunFile := "/run/maxgpt/capability"
	if os.Getenv("MAXGPT_FORCE_META_BACKEND_CAPABILITY_RUN_FILE") != "" {
		log.Debug().Str("MAXGPT_FORCE_META_BACKEND_CAPABILITY_RUN_FILE", os.Getenv("MAXGPT_FORCE_META_BACKEND_CAPABILITY_RUN_FILE")).Msg("Using forced capability run file")
		capabilityRunFile = os.Getenv("MAXGPT_FORCE_META_BACKEND_CAPABILITY_RUN_FILE")
	}

	// Check if /run/maxgpt/capability exists and use it
	// This might be used by e.g. container images to specify which
	// backends to pull in automatically when installing meta backends.
	if _, err := os.Stat(capabilityRunFile); err == nil {
		capability, err := os.ReadFile(capabilityRunFile)
		if err == nil {
			log.Debug().Str("capability", string(capability)).Msg("Using capability from run file")
			return strings.Trim(strings.TrimSpace(string(capability)), "\n")
		}
	}

	// If we are on mac and arm64, we will return metal
	if runtime.GOOS == "darwin" && runtime.GOARCH == "arm64" {
		log.Debug().Msg("Using metal capability")
		return metal
	}

	// If we are on mac and x86, we will return darwin-x86
	if runtime.GOOS == "darwin" && runtime.GOARCH == "amd64" {
		log.Debug().Msg("Using darwin-x86 capability")
		return darwinX86
	}

	// If arm64 on linux and a nvidia gpu is detected, we will return nvidia-l4t
	if runtime.GOOS == "linux" && runtime.GOARCH == "arm64" {
		if s.GPUVendor == "nvidia" {
			log.Debug().Msg("Using nvidia-l4t capability")
			return nvidiaL4T
		}
	}

	if s.GPUVendor == "" {
		log.Debug().Msg("Using default capability")
		return defaultCapability
	}

	log.Debug().Str("GPUVendor", s.GPUVendor).Msg("Using GPU vendor capability")
	return s.GPUVendor
}

func GetSystemState() (*SystemState, error) {
	gpuVendor, _ := detectGPUVendor()
	log.Debug().Str("gpuVendor", gpuVendor).Msg("GPU vendor")

	return &SystemState{
		GPUVendor: gpuVendor,
	}, nil
}

func detectGPUVendor() (string, error) {
	gpus, err := xsysinfo.GPUs()
	if err != nil {
		return "", err
	}

	for _, gpu := range gpus {
		if gpu.DeviceInfo != nil {
			if gpu.DeviceInfo.Vendor != nil {
				gpuVendorName := strings.ToUpper(gpu.DeviceInfo.Vendor.Name)
				if strings.Contains(gpuVendorName, "NVIDIA") {
					return "nvidia", nil
				}
				if strings.Contains(gpuVendorName, "AMD") {
					return "amd", nil
				}
				if strings.Contains(gpuVendorName, "INTEL") {
					return "intel", nil
				}
				return "nvidia", nil
			}
		}

	}

	return "", nil
}
