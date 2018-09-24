package integration_test

import (
	"github.com/cloudfoundry/libbuildpack/cutlass"

	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CF NodeJS Buildpack", func() {
	var app *cutlass.App
	AfterEach(func() {
		if app != nil {
			app.Destroy()
		}
		app = nil

	})
	Context("deploying a Node.js app that uses yarn workspaces", func() {
		BeforeEach(func() {
			app = cutlass.New(filepath.Join(bpDir, "fixtures", "yarn_with_workspaces"))
		})

		FIt("outputs config contents when queried", func() {
			PushAppAndConfirm(app)
			// Eventually(app.Stdout.String).Should(MatchRegexp("Installing node modules (yarn.lock)"))
			Expect(app.GetBody("/check")).To(ContainSubstring(
				`"config":{"prop1":"Package A value 1","prop2":"Package A value 2"}`,
			))
		})
	})
})
