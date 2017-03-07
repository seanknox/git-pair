package user_parser_test

import (
	"fmt"

	. "github.com/seanknox/git-pair/user_parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserParser", func() {
	Context("without at least one username provided", func() {

		It("should return an error", func() {
			username, err := ParseUsername(nil)
			Expect(username).To(Equal([]string{""}))
			Expect(err).To(Equal(fmt.Errorf("Please supply at least one GitHub username")))
		})
	})

	Context("with at least one username provided", func() {

		It("should return the username", func() {
			var args = []string{"seanknox"}
			Expect(ParseUsername(args)).To(Equal([]string{"seanknox"}))
		})
	})

	Context("with two usernames provided", func() {
		It("should return both usernames", func() {
			var args = []string{"seanknox", "jdumars"}
			Expect(ParseUsername(args)).To(Equal([]string{"seanknox", "jdumars"}))
		})
	})
})
