package main

import "fmt"

type Directory struct {
	name     string
	parent   *Directory
	children map[string]*Directory
	files    map[string]*File
	size     int
}

func (d *Directory) traverse(f func(d *Directory)) {
	f(d)
	for _, val := range d.children {
		val.traverse(f)
	}
}

func (d *Directory) String() string {
	return fmt.Sprintf("dir %q - size: %d", d.name, d.calculateSize())
}
func (d *Directory) addFile(file *File) {
	d.files[file.name] = file
	d.sizeIncreased(file.size)
}

func (d *Directory) addSubdirectory(name string) {
	d.children[name] = NewEmptyDirectory(name, d)
}

func (d *Directory) sizeIncreased(deltaSize int) {
	d.size += deltaSize
	if d.parent != nil {
		d.parent.sizeIncreased(deltaSize)
	}
}

func (d *Directory) calculateSize() int {
	totalSize := 0
	for _, value := range d.children {
		totalSize += value.calculateSize()
	}
	for _, value := range d.files {
		totalSize += value.size
	}
	return totalSize
}

func NewEmptyDirectory(name string, parent *Directory) *Directory {
	return &Directory{name, parent, make(map[string]*Directory), make(map[string]*File), 0}
}
