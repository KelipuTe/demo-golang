package annotation

import (
	"go/ast"
)

// SingleFileEntryVisitor 这部分和课堂演示差不多，但是我建议你们自己试着写一些
type SingleFileEntryVisitor struct {
	file *fileVisitor
}

func (s *SingleFileEntryVisitor) Get() File {
	if nil != s.file {
		return s.file.Get()
	}
	return File{}
}

func (s *SingleFileEntryVisitor) Visit(node ast.Node) ast.Visitor {
	p7sf, ok := node.(*ast.File)
	if ok {
		s.file = &fileVisitor{
			ans:     newAnnotations(p7sf, p7sf.Doc),
			types:   make([]*typeVisitor, 0),
			visited: false,
		}
		return s.file
	}
	return s
}

type fileVisitor struct {
	ans     Annotations[*ast.File]
	types   []*typeVisitor
	visited bool
}

func (f *fileVisitor) Get() File {
	s5t := make([]Type, 0, len(f.types))
	for _, p7tv := range f.types {
		s5t = append(s5t, p7tv.Get())
	}
	return File{
		Annotations: f.ans,
		Types:       s5t,
	}
}

func (f *fileVisitor) Visit(node ast.Node) ast.Visitor {
	p7ats, ok := node.(*ast.TypeSpec)
	if ok {
		p7tv := &typeVisitor{
			ans:    newAnnotations(p7ats, p7ats.Doc),
			fields: make([]Field, 0),
		}
		f.types = append(f.types, p7tv)
		return p7tv
	}
	return f
}

type File struct {
	Annotations[*ast.File]
	Types []Type
}

type typeVisitor struct {
	ans    Annotations[*ast.TypeSpec]
	fields []Field
}

func (t *typeVisitor) Get() Type {
	return Type{
		Annotations: t.ans,
		Fields:      t.fields,
	}
}

func (t *typeVisitor) Visit(node ast.Node) (w ast.Visitor) {
	p7af, ok := node.(*ast.Field)
	if ok {
		t4f := Field{
			Annotations: newAnnotations(p7af, p7af.Doc),
		}
		t.fields = append(t.fields, t4f)
		return nil
	}
	return t
}

type Type struct {
	Annotations[*ast.TypeSpec]
	Fields []Field
}

type Field struct {
	Annotations[*ast.Field]
}