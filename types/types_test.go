package types_test

import (
	"github.com/aamgayle/gotmdb/consts"
	"github.com/aamgayle/gotmdb/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	testConfig = types.Config{"testkey", "testbearer", "testID"}
	testProps  = types.TMDBReqProps{consts.BaseURL, "", types.GET, testConfig}
)

var _ = Describe("Types tests", func() {
	Context("When NewHTTPRequestFromProps is given props", func() {
		It("should return an error if props is nil", func() {
			_, err := types.NewHTTPRequestFromProps(nil)
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("no props provided"))
		})
	})
})
