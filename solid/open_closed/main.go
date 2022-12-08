package main

import "fmt"

type Color int

const (
	Red Color = iota
	Green
	Blue
)

type Size int

const (
	Small Size = iota
	Medium
	Large
)

type Product struct {
	Name  string
	Color Color
	Size  Size
}

// Before we implement open closed principle, we have to create new method for every filter
type FilterWithoutOCP struct{}

func (f *FilterWithoutOCP) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.Color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *FilterWithoutOCP) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.Size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

// Spec interface to manage any filter spec
type Specification interface {
	IsSatisfied(p *Product) bool
}

// Create new spec for color
type ColorSpec struct {
	Color Color
}

func (c *ColorSpec) IsSatisfied(p *Product) bool {
	return c.Color == p.Color
}

// Create new spec for size
type SizeSpec struct {
	Size Size
}

func (s *SizeSpec) IsSatisfied(p *Product) bool {
	return s.Size == p.Size
}

// Create new spec for size and color
type AndSpec struct {
	first, second Specification
}

func (cs *AndSpec) IsSatisfied(p *Product) bool {
	return cs.first.IsSatisfied(p) && cs.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (b *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	res := make([]*Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			res = append(res, &products[i])
		}
	}

	return res
}

func main() {
	apple := Product{"apple", Red, Small}
	tree := Product{"tree", Green, Medium}
	house := Product{"house", Green, Large}
	products := []Product{apple, tree, house}

	fmt.Printf("Filter green products before OCP\n")
	filter := FilterWithoutOCP{}
	for _, v := range filter.FilterByColor(products, Green) {
		fmt.Printf("- %s is green\n", v.Name)
	}

	fmt.Printf("Filter green products after OCP\n")
	betterFilter := BetterFilter{}
	greenSpec := ColorSpec{Green}
	for _, v := range betterFilter.Filter(products, &greenSpec) {
		fmt.Printf("- %s is green\n", v.Name)
	}

	fmt.Printf("Filter green AND medium products after OCP\n")
	mediumSpec := SizeSpec{Medium}
	andSpec := AndSpec{&greenSpec, &mediumSpec}
	for _, v := range betterFilter.Filter(products, &andSpec) {
		fmt.Printf("- %s is green and medium\n", v.Name)
	}
}
