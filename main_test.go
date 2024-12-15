package main_test

import (
	repository "a21hc3NpZ25tZW50/repository"
	fileservice "a21hc3NpZ25tZW50/service"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("FileService", func() {
	var fileService fileservice.FileServiceInterface

	BeforeEach(func() {
		fileService = fileservice.NewFileService(repository.FileRepository{})
	})

	Describe("ProcessFile", func() {
		It("should return the correct result for valid CSV data", func() {
			fileContent := `header1,header2
value1,value2
value3,value4`
			expected := map[string][]string{
				"header1": {"value1", "value3"},
				"header2": {"value2", "value4"},
			}

			result, err := fileService.ProcessFile(fileContent)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(expected))
		})

		It("should return an error for empty CSV data", func() {
			fileContent := ``

			result, err := fileService.ProcessFile(fileContent)
			Expect(err).To(HaveOccurred())
			Expect(result).To(BeNil())
		})

		It("should return an error for invalid CSV data", func() {
			fileContent := `header1,header2
value1,value2
value3`

			result, err := fileService.ProcessFile(fileContent)
			Expect(err).To(HaveOccurred())
			Expect(result).To(BeNil())
		})
	})
})
