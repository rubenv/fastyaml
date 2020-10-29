package fastyaml

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/rubenv/fastyaml/parser"
)

type parserData struct {
	Package string
	Imports []string
	Name    string
	Types   []*parseType
}

type parseType struct {
	Name     string
	Type     string
	TypeInit string
	ErrType  string
	Fields   []*parseField
}

type parseField struct {
	Name          string
	Field         string
	Method        string
	Slice         bool
	AdvanceBefore bool
}

type generator struct{}

func Generate(pkg string, obj interface{}) (string, error) {
	g := &generator{}
	return g.Generate(pkg, obj)
}

func (g *generator) Generate(pkg string, obj interface{}) (string, error) {
	p := &parserData{
		Package: pkg,
	}

	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		return "", errors.New("Not a struct")
	}
	p.Name = t.Name()

	// Import path
	pt := reflect.TypeOf(parser.Parser{})
	p.Imports = append(p.Imports, pt.PkgPath())

	g.outputType(t, p, true)

	// Generate + format
	var buf bytes.Buffer
	err := fileTpl.Execute(&buf, p)
	if err != nil {
		return "", err
	}

	out, err := format.Source(buf.Bytes())
	if err != nil {
		return "", fmt.Errorf("%w\n\n%s", err, formatSource(buf.String()))
	}
	return string(out), nil
}

func (g *generator) outputType(t reflect.Type, p *parserData, isPointer bool) error {
	// Skip if already present
	for _, ct := range p.Types {
		if ct.Name == t.Name() {
			return nil
		}
	}

	typeName := t.Name()
	fullType := typeName
	typeInit := fmt.Sprintf("%s{}", typeName)
	errType := typeInit
	if isPointer {
		fullType = fmt.Sprintf("*%s", typeName)
		typeInit = fmt.Sprintf("&%s{}", typeName)
		errType = "nil"
	}
	pt := &parseType{
		Name:     typeName,
		Type:     fullType,
		TypeInit: typeInit,
		ErrType:  errType,
	}
	p.Types = append(p.Types, pt)

	for i := 0; i < t.NumField(); i++ {
		f := t.FieldByIndex([]int{i})
		tag := f.Tag.Get("yaml")
		if tag == "-" {
			continue
		}

		pf := &parseField{
			Name:  tag,
			Field: f.Name,
		}
		pt.Fields = append(pt.Fields, pf)

		switch f.Type.Kind() {
		case reflect.Struct:
			pf.Method = fmt.Sprintf("p.parse%s", f.Type.Name())
			pf.AdvanceBefore = true
			g.outputType(f.Type, p, false)
		case reflect.String:
			pf.Method = "p.ReadString"
		case reflect.Slice:
			elType := f.Type.Elem()
			pf.Slice = true
			pf.Method = fmt.Sprintf("p.parse%s", elType.Name())
			pf.AdvanceBefore = true
			g.outputType(elType, p, false)
		case reflect.Int:
			pf.Method = "p.ReadInt"
		case reflect.Float32:
			pf.Method = "p.ReadFloat32"
		case reflect.Float64:
			pf.Method = "p.ReadFloat64"
		default:
			return fmt.Errorf("Unsupported type: %v", f.Type.Kind())
		}
	}
	return nil
}

func formatSource(src string) string {
	lines := strings.Split(src, "\n")
	for i, l := range lines {
		lines[i] = fmt.Sprintf("%d\t%s", i+1, l)
	}
	return strings.Join(lines, "\n")
}

func GenerateFile(pkg string, obj interface{}, filename string) error {
	src, err := Generate(pkg, obj)
	if err != nil {
		return err
	}

	fp, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = io.WriteString(fp, src)
	return err
}
