package jhanda_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/jhanda"
	"github.com/pivotal-cf/jhanda/internal/fakes"
)

var _ = Describe("Global Usage", func() {
	It("returns a formatted version of the flag set usage", func() {
		commandBake := &fakes.Command{}
		commandBake.UsageReturns(jhanda.Usage{ShortDescription: "it bakes"})
		commandClean := &fakes.Command{}
		commandClean.UsageReturns(jhanda.Usage{ShortDescription: "it cleans"})

		commandSet := jhanda.CommandSet{
			"bake":  commandBake,
			"clean": commandClean,
		}

		usage, err := jhanda.PrintGlobalUsage(commandSet)
		Expect(err).NotTo(HaveOccurred())
		Expect(usage).To(Equal(strings.TrimSpace(`
bake   it bakes
clean  it cleans
`)))
	})
})
