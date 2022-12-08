package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Create Journal struct
type Journals struct {
	Entries    []string
	EntryCount int
}

func (j *Journals) AddEntry(name string) int {
	j.Entries = append(j.Entries, name)
	j.EntryCount++
	return j.EntryCount
}

func (j *Journals) RemoveEntry(index int) {
	j.Entries = append(j.Entries[:index], j.Entries[index+1:]...)
}

func (j *Journals) GetEntries() []string {
	return j.Entries
}

// Create books struct
type Books struct {
	Entries    []string
	EntryCount int
}

func (b *Books) AddEntry(name string) int {
	b.Entries = append(b.Entries, name)
	b.EntryCount++
	return b.EntryCount
}

func (b *Books) RemoveEntry(index int) {
	b.Entries = append(b.Entries[:index], b.Entries[index+1:]...)
}

func (b *Books) GetEntries() []string {
	return b.Entries
}

// Create Entries interface as a polymorphism
type Entries interface {
	AddEntry(string) int
	RemoveEntry(int)
	GetEntries() []string
}

// Type persistence to save to file
type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile(e Entries, filename string) {
	f, err := os.Create(filename + ".txt")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = f.WriteString(strings.Join(e.GetEntries(), p.LineSeparator))
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	j := &Journals{}
	j.AddEntry("file 1")
	j.AddEntry("file 2")

	b := &Books{}
	b.AddEntry("book 1")
	b.AddEntry("book 2")

	p := &Persistence{
		LineSeparator: "\n",
	}

	p.SaveToFile(j, "Journals")
	p.SaveToFile(b, "Books")

	fmt.Println(j.Entries)
}
