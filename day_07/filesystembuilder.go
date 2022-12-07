package main

import (
	"andurian/adventofcode/2022/util"
	"strings"
)

type FilesystemBuilder struct {
	root    *Directory
	current *Directory
}

func (b *FilesystemBuilder) ChangeDirectory(dir string) {
	switch dir {
	case "/":
		b.current = b.root
	case "..":
		b.current = b.current.parent
	default:
		b.current = b.current.children[dir]
	}
}

func (b *FilesystemBuilder) MakeFile(name string, size int) {
	b.current.addFile(&File{name, size})
}

func (b *FilesystemBuilder) MakeDirectory(name string) {
	b.current.addSubdirectory(name)
}

func (b *FilesystemBuilder) ExecuteCommand(cmd string) {
	if strings.HasPrefix(cmd, "cd") {
		b.ChangeDirectory(strings.Split(cmd, " ")[1])
	}
}

func (b *FilesystemBuilder) MakeThing(line string) {
	tokens := strings.Split(line, " ")
	if tokens[0] == "dir" {
		b.MakeDirectory(tokens[1])
	} else {
		b.MakeFile(tokens[1], util.AtoiSafe(tokens[0]))
	}
}

func NewFilesystemBuilder() *FilesystemBuilder {
	ret := &FilesystemBuilder{root: NewEmptyDirectory("/", nil)}
	ret.current = ret.root
	return ret
}

func MakeDirectory(input string) *Directory {
	builder := NewFilesystemBuilder()
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "$") {
			builder.ExecuteCommand(line[2:])
		} else {
			builder.MakeThing(line)
		}
	}
	return builder.root
}
