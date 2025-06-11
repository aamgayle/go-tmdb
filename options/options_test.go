package options_test

import (
	"time"

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

	//Start and End Date option tests

	Context("When using the EndDate option function", func() {
		It("should add a properly formatted date to the query", func() {
			t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
			options := []func(*TMDBReqProps){
				options.WithBaseURL(protocol + testHost + testSudDirectory),
				options.WithEndDate(t),
			}
			tmdbReq, err := NewTMDBRequest(options...)
			Expect(err).To(BeNil())
			Expect(tmdbReq.URL.Path).To(ContainSubstring("&end_date=2009-11-10"))
		})

		It("should add add padding when single digit days or months are used", func() {
			t := time.Date(2009, time.January, 5, 23, 0, 0, 0, time.UTC)
			options := []func(*TMDBReqProps){
				options.WithBaseURL(protocol + testHost + testSudDirectory),
				options.WithEndDate(t),
			}
			tmdbReq, err := NewTMDBRequest(options...)
			Expect(err).To(BeNil())
			Expect(tmdbReq.URL.Path).To(ContainSubstring("&end_date=2009-01-05"))
		})
	})

	Context("When using the StartDate option function", func() {
		It("should add a properly formatted date to the query", func() {
			t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
			options := []func(*TMDBReqProps){
				options.WithBaseURL(protocol + testHost + testSudDirectory),
				options.WithStartDate(t),
			}
			tmdbReq, err := NewTMDBRequest(options...)
			Expect(err).To(BeNil())
			Expect(tmdbReq.URL.Path).To(ContainSubstring("&start_date=2009-11-10"))
		})

		It("should add add padding when single digit days or months are used", func() {
			t := time.Date(2009, time.January, 5, 23, 0, 0, 0, time.UTC)
			options := []func(*TMDBReqProps){
				options.WithBaseURL(protocol + testHost + testSudDirectory),
				options.WithStartDate(t),
			}
			tmdbReq, err := NewTMDBRequest(options...)
			Expect(err).To(BeNil())
			Expect(tmdbReq.URL.Path).To(ContainSubstring("&start_date=2009-01-05"))
		})
	})

})
