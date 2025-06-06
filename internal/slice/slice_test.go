package slice_test

import (
	"slices"
	"strconv"
	"testing"

	"github.com/bitmagnet-io/bitmagnet/internal/slice"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSlice(t *testing.T) {
	t.Parallel()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Slice Suite")
}

var _ = Describe("Slice Utils", func() {
	Describe("Map", func() {
		It("returns empty slice for an empty input", func() {
			mapFunc := func(v int) string { return strconv.Itoa(v * 2) }
			result := slice.Map([]int{}, mapFunc)
			Expect(result).To(BeEmpty())
		})

		It("returns a new slice with elements mapped", func() {
			mapFunc := func(v int) string { return strconv.Itoa(v * 2) }
			result := slice.Map([]int{1, 2, 3, 4}, mapFunc)
			Expect(result).To(ConsistOf("2", "4", "6", "8"))
		})
	})

	Describe("MapWithArg", func() {
		It("returns empty slice for an empty input", func() {
			mapFunc := func(a int, v int) string { return strconv.Itoa(a + v) }
			result := slice.MapWithArg([]int{}, 10, mapFunc)
			Expect(result).To(BeEmpty())
		})

		It("returns a new slice with elements mapped", func() {
			mapFunc := func(a int, v int) string { return strconv.Itoa(a + v) }
			result := slice.MapWithArg([]int{1, 2, 3, 4}, 10, mapFunc)
			Expect(result).To(ConsistOf("11", "12", "13", "14"))
		})
	})

	Describe("Group", func() {
		It("returns empty map for an empty input", func() {
			keyFunc := func(v int) int { return v % 2 }
			result := slice.Group([]int{}, keyFunc)
			Expect(result).To(BeEmpty())
		})

		It("groups by the result of the key function", func() {
			keyFunc := func(v int) int { return v % 2 }
			result := slice.Group([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, keyFunc)
			Expect(result).To(HaveLen(2))
			Expect(result[0]).To(ConsistOf(2, 4, 6, 8, 10))
			Expect(result[1]).To(ConsistOf(1, 3, 5, 7, 9, 11))
		})
	})

	Describe("ToMap", func() {
		It("returns empty map for an empty input", func() {
			transformFunc := func(v int) (int, string) { return v, strconv.Itoa(v) }
			result := slice.ToMap([]int{}, transformFunc)
			Expect(result).To(BeEmpty())
		})

		It("returns a map with the result of the transform function", func() {
			transformFunc := func(v int) (int, string) { return v * 2, strconv.Itoa(v * 2) }
			result := slice.ToMap([]int{1, 2, 3, 4}, transformFunc)
			Expect(result).To(HaveLen(4))
			Expect(result).To(HaveKeyWithValue(2, "2"))
			Expect(result).To(HaveKeyWithValue(4, "4"))
			Expect(result).To(HaveKeyWithValue(6, "6"))
			Expect(result).To(HaveKeyWithValue(8, "8"))
		})
	})

	Describe("Unique", func() {
		It("returns empty slice for an empty input", func() {
			Expect(slice.Unique([]int{})).To(BeEmpty())
		})

		It("returns the unique elements", func() {
			Expect(slice.Unique([]int{1, 2, 1, 2, 3, 2})).To(HaveExactElements(1, 2, 3))
		})
	})

	DescribeTable("CollectChunks",
		func(input []int, n int, expected [][]int) {
			var result [][]int
			for chunks := range slice.CollectChunks(slices.Values(input), n) {
				result = append(result, chunks)
			}
			Expect(result).To(Equal(expected))
		},
		Entry("returns empty slice (nil) for an empty input", []int{}, 1, nil),
		Entry("returns the slice in one chunk if len < chunkSize", []int{1, 2, 3}, 10, [][]int{{1, 2, 3}}),
		Entry("breaks up the slice if len > chunkSize", []int{1, 2, 3, 4, 5}, 3, [][]int{{1, 2, 3}, {4, 5}}),
	)

	Describe("SeqFunc", func() {
		It("returns empty slice for an empty input", func() {
			it := slice.SeqFunc([]int{}, func(v int) int { return v })

			result := slices.Collect(it)
			Expect(result).To(BeEmpty())
		})

		It("returns a new slice with mapped elements", func() {
			it := slice.SeqFunc([]int{1, 2, 3, 4}, func(v int) string { return strconv.Itoa(v * 2) })

			result := slices.Collect(it)
			Expect(result).To(ConsistOf("2", "4", "6", "8"))
		})
	})
})
