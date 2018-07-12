package jhanda_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/jhanda"
	"github.com/pivotal-cf/jhanda/internal/fakes"
)

const GLOBAL_USAGE = `
Usage:
  --query, -?     asks a question
  --surprise, -!  gives you a present

Commands:
  bake   it bakes
  clean  it cleans
`

var _ = Describe("Global Usage", func() {
	var (
		globalFlags string
		commandSet  jhanda.CommandSet
	)
	BeforeEach(func() {
		globalFlags = strings.TrimSpace(`
--query, -?     asks a question
--surprise, -!  gives you a present
`)

		commandBake := &fakes.Command{}
		commandBake.UsageReturns(jhanda.Usage{ShortDescription: "it bakes"})
		commandClean := &fakes.Command{}
		commandClean.UsageReturns(jhanda.Usage{ShortDescription: "it cleans"})

		commandSet = jhanda.CommandSet{
			"bake":  commandBake,
			"clean": commandClean,
		}
	})
	It("returns a formatted version of the flag set usage", func() {
		usage, err := jhanda.PrintGlobalUsage(globalFlags, commandSet)
		Expect(err).NotTo(HaveOccurred())

		fmt.Println(usage)

		Expect(usage).To(Equal(GLOBAL_USAGE))
	})
})
