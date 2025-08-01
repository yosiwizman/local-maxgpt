package e2e_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	localAIURL = os.Getenv("LOCALAI_API")
)

func TestMaxGPT(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MaxGPT E2E test suite")
}
