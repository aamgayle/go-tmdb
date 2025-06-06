package options_test

import (
	"github.com/aamgayle/gotmdb/options"
	. "github.com/aamgayle/gotmdb/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	protocol         = "https://"
	testHost         = "www.testurl.com"
	testSudDirectory = "/1/test"
)

var _ = Describe("Options tests", func() {
	Context("When given a different base url", func() {
		It("should return a url that doesn't match the default", func() {
			options := []func(*TMDBReqProps){options.WithBaseURL(protocol + testHost + testSudDirectory)}
			tmdbReq, err := NewTMDBRequest(options...)
			Expect(err).To(BeNil())
			Expect(tmdbReq.Host).To(ContainSubstring(testHost))
		})
	})

	Context("When using the pages options function", func() {
		It("should give 1 page when given values under 1", func() {
			options := []func(*TMDBReqProps){
				options.WithBaseURL(protocol + testHost + testSudDirectory),
				options.WithPages(0)}

			tmdbReq, err := NewTMDBRequest(options...)
			Expect(err).To(BeNil())
			Expect(tmdbReq.URL.Path).To(ContainSubstring("&page=1"))
		})
		It("should give 1 page when given values under 1", func() {
			options := []func(*TMDBReqProps){
				options.WithBaseURL(protocol + testHost + testSudDirectory),
				options.WithPages(12)}

			tmdbReq, err := NewTMDBRequest(options...)
			Expect(err).To(BeNil())
			Expect(tmdbReq.URL.Path).To(ContainSubstring("&page=12"))
		})
		It("should add a keyword query with proper keyword", func() {
			options := []func(*TMDBReqProps){
				options.WithBaseURL(protocol + testHost + testSudDirectory),
				options.WithKeyword("keyword"),
				options.WithPages(0),
			}

			tmdbReq, err := NewTMDBRequest(options...)
			Expect(err).To(BeNil())
			Expect(tmdbReq.URL.RawQuery).To(ContainSubstring("&page=1"))
			Expect(tmdbReq.URL.RawQuery).To(ContainSubstring("query=keyword"))
		})
	})

	Context("When using the credit ID option function", func() {
		It("should add the credit ID ", func() {
			options := []func(*TMDBReqProps){
				options.WithBaseURL(protocol + testHost + testSudDirectory),
				options.WithCreditID("credit_id"),
			}

			tmdbReq, err := NewTMDBRequest(options...)
			Expect(err).To(BeNil())
			Expect(tmdbReq.URL.Path).To(ContainSubstring("/credit/credit_id"))
		})
	})
})
