package user_parser_test

import (
	"os"

	. "github.com/seanknox/git-pair/user_parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserParser", func() {
	Context("without at least one username provided", func() {

		It("should return an error", func() {
			os.Args = []string{""}
			Expect(ParseUsername()).To(Equal([]string{""}))
		})
	})

	Context("with at least one username provided", func() {

		It("should return the username", func() {
			os.Args = []string{"seanknox"}
			Expect(ParseUsername()).To(Equal([]string{"seanknox"}))
		})
	})

	Context("with two usernames provided", func() {
		It("should return both usernames", func() {
			os.Args = []string{"seanknox", "jdumars"}
			Expect(ParseUsername()).To(Equal([]string{"seanknox", "jdumars"}))
		})
	})
})
