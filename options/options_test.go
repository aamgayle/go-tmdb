package options_test

import (
	"github.com/aamgayle/gotmdb/options"
	. "github.com/aamgayle/gotmdb/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Options tests", func() {
	Context("When given a different base url", func() {
		It("should return a url that doesn't match the default", func() {
			protocol := "https://"
			testHost := "www.testurl.com"
			testSudDirectory := "/1/test/"

			tmdbReq, err := NewTMDBRequest(options.WithBaseURL(protocol + testHost + testSudDirectory))
			Expect(err).To(BeNil())
			Expect(tmdbReq.Host).To(ContainSubstring(testHost))
		})
	})
})
