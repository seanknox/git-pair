package user_parser_test

import (
	"fmt"
	"os"

	. "github.com/seanknox/git-pair/user_parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserParser", func() {
	Context("without at least one username provided", func() {

		It("should return an error", func() {
			os.Args = nil
			username, err := ParseUsername()
			Expect(username).To(Equal([]string{""}))
			Expect(err).To(Equal(fmt.Errorf("Please supply at least one GitHub username")))
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
