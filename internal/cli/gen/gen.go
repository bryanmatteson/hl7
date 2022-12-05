package gen

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/dave/jennifer/jen"
)

// Schema ...
type Schema struct {
	DataTypes map[string]*APIDataTypeItem `json:"dataTypes"`
	Segments  map[string]*APISegmentItem  `json:"segments"`
	Events    map[string]*APIEventItem    `json:"events"`
	Tables    map[string]*APITableItem    `json:"tables"`
	Fields    map[string]*APIFieldItem    `json:"fields"`
}

// Context ...
type Context struct {
	TypeMap     map[string]*GoType
	PackageName string
	GoPackage   string
	OutputDir   string
}

// Options ...
type Options struct {
	GoPackage string
}

// Generate ...
func Generate(version HL7Version, contextDir, outputDir string, options *Options) error {
	ctx, err := LoadContext(contextDir)
	if err != nil {
		return err
	}

	ctx.PackageName = "hl7v" + strings.ReplaceAll(string(version), ".", "")
	ctx.GoPackage = options.GoPackage
	ctx.OutputDir = path.Join(outputDir, ctx.PackageName)

	if err := os.MkdirAll(ctx.OutputDir, os.ModePerm); err != nil {
		return err
	}

	segments := jen.NewFile(ctx.PackageName)
	datatypes := jen.NewFile(ctx.PackageName)
	groups := jen.NewFile(ctx.PackageName)
	messages := jen.NewFile(ctx.PackageName)
	layouts := jen.NewFile(ctx.PackageName)

	for _, typ := range ctx.TypeMap {
		switch typ.Kind {
		case GoMessageKind:
			genMessage(ctx, typ, messages)
			genLayout(ctx, typ, layouts)
		case GoGroupKind:
			genGroup(ctx, typ, groups)
		case GoSegmentKind:
			genSegment(ctx, typ, segments)
		case GoDataTypeKind:
			genDataType(ctx, typ, datatypes)
		}
	}

	if err := groups.Save(path.Join(ctx.OutputDir, "groups.go")); err != nil {
		return err
	}

	if err := layouts.Save(path.Join(ctx.OutputDir, "layouts.go")); err != nil {
		return err
	}

	if err := messages.Save(path.Join(ctx.OutputDir, "messages.go")); err != nil {
		return err
	}

	if err := datatypes.Save(path.Join(ctx.OutputDir, "datatypes.go")); err != nil {
		return err
	}

	if err := segments.Save(path.Join(ctx.OutputDir, "segments.go")); err != nil {
		return err
	}

	return nil
}

func genLayout(ctx *Context, msg *GoType, cf *jen.File) {
	if msg.Kind != GoMessageKind {
		return
	}

	cf.Var().Id("layout" + msg.Name).Op("=").Add(genLayoutGroup(ctx, msg))
	genLayoutCtor(ctx, msg, cf)
}

func genLayoutGroup(ctx *Context, grp *GoType) jen.Code {
	code := jen.Qual(path.Join(ctx.GoPackage, "internal", "layout"), "Group").Call(jen.Lit(grp.Name))

	for _, fld := range grp.Fields {
		var stmt jen.Code
		if fld.IsGroup {
			stmt = genLayoutGroup(ctx, fld.Type)
		} else {
			stmt = jen.Qual(path.Join(ctx.GoPackage, "internal", "layout"), "Segment").Call(jen.Lit(fld.TypeName))
		}

		if fld.IsArray {
			stmt = jen.Qual(path.Join(ctx.GoPackage, "internal", "layout"), "List").Call(stmt)
		}

		if !fld.IsRequired() {
			stmt = jen.Qual(path.Join(ctx.GoPackage, "internal", "layout"), "Maybe").Call(stmt)
		}

		code.Dot("With").Call(stmt)
	}
	return code
}

// func NewORUR01(msg hl7.Message) (hl7.Group, error) {
// 	result := orur01.Parse(NewInput(msg))
// 	if !result.Success {
// 		return nil, result.Error
// 	}
// 	return result.Item.(hl7.Group), nil
// }

func genLayoutCtor(ctx *Context, msg *GoType, cf *jen.File) {
	cf.Func().Id("New" + msg.Name + "Layout").Params(jen.Id("msg").Qual(ctx.GoPackage, "Message")).Parens(jen.List(jen.Op("*").Id(msg.Name), jen.Error())).BlockFunc(func(g *jen.Group) {
		g.Id("result").Op(":=").Id("layout" + msg.Name).Dot("Parse").Call(jen.Qual(path.Join(ctx.GoPackage, "internal", "layout"), "NewInput").Call(jen.Id("msg")))
		g.If(jen.Op("!").Id("result").Dot("Success")).BlockFunc(func(g *jen.Group) {
			g.Return(jen.Nil(), jen.Id("result").Dot("Error"))
		})
		g.Id("v").Op(":=").Id("New" + msg.Name).Call(jen.Id("result").Dot("Item").Assert(jen.Qual(ctx.GoPackage, "Group")))
		g.Return(jen.Op("&").Id("v"), jen.Nil())
	})
}

func genGroup(ctx *Context, grp *GoType, cf *jen.File) {
	if grp.Kind != GoGroupKind {
		return
	}

	genConstructor(ctx, grp, cf)
	genConcreteImplementationType(ctx, grp, cf)
	genListType(ctx, grp, cf)
}

func genMessage(ctx *Context, msg *GoType, cf *jen.File) {
	if msg.Kind != GoMessageKind {
		return
	}

	genConstructor(ctx, msg, cf)
	genConcreteImplementationType(ctx, msg, cf)
	genListType(ctx, msg, cf)
}

func genSegment(ctx *Context, seg *GoType, cf *jen.File) {
	if seg.Kind != GoSegmentKind {
		return
	}

	genConcreteImplementationType(ctx, seg, cf)
	genConstructor(ctx, seg, cf)
	genListType(ctx, seg, cf)
}

func genDataType(ctx *Context, dt *GoType, cf *jen.File) {
	if dt.Kind != GoDataTypeKind {
		return
	}

	if len(dt.Fields) == 0 {
		cf.Comment(dt.Comment())
		cf.Type().Id(dt.Name).String()
		genConstructor(ctx, dt, cf)
	} else {
		genConcreteImplementationType(ctx, dt, cf)
		genConstructor(ctx, dt, cf)
	}
	genListType(ctx, dt, cf)
}

func genListType(ctx *Context, typ *GoType, cf *jen.File) {
	name := typ.Name

	constructorName := "New" + name + "Slice"
	cf.Comment(fmt.Sprintf("%s will construct a slice of type %s", constructorName, name))
	cf.Func().Id(constructorName).Params(jen.Id("input").Qual(ctx.GoPackage, "HL7Type")).Index().Id(name).BlockFunc(func(g *jen.Group) {
		g.Id("values").Op(":=").Make(jen.Index().Id(name), jen.Id("input").Dot("Len").Call())
		g.For(jen.Id("i").Op(":=").Lit(0), jen.Id("i").Op("<").Id("input").Dot("Len").Call(), jen.Id("i").Op("++")).BlockFunc(func(g *jen.Group) {
			g.Id("values").Index(jen.Id("i")).Op("=").Id("New" + name).Call(jen.Id("input").Dot("Index").Call(jen.Id("i")))
		})
		g.Return(jen.Id("values"))
	})
}

func genConstructor(ctx *Context, typ *GoType, cf *jen.File) {
	fieldCount := len(typ.Fields)

	name := typ.ConcreteName()
	input := jen.Id("input").Qual(ctx.GoPackage, "HL7Type")

	constructorName := "New" + name
	cf.Comment(fmt.Sprintf("%s creates an implementation of %s", constructorName, name))
	cf.Func().Id(constructorName).Params(input).Id(name).BlockFunc(func(g *jen.Group) {
		if fieldCount > 0 {
			g.Id("v").Op(":=").Id(typ.ConcreteName()).Values()
			for idx, fld := range typ.Fields {
				index := idx
				if typ.Kind == GoSegmentKind {
					index++
				}

				item := jen.Id("input").Dot("Index").Call(jen.Lit(index))
				if !fld.IsArray && (typ.Kind == GoSegmentKind) {
					item = item.Dot("Index").Call(jen.Lit(0))
				}
				ctor := "New" + fld.TypeName
				if fld.IsArray {
					ctor += "Slice"
				}
				g.Id("v").Dot(fld.Name).Op("=").Id(ctor).Call(item)
			}
			g.Return(jen.Id("v"))
		} else {
			g.Return(jen.Id(name).Call(jen.Id("input").Dot("String").Call()))
		}
	})
}

func genConcreteImplementationType(ctx *Context, typ *GoType, cf *jen.File) {
	cf.Type().Id(typ.ConcreteName()).StructFunc(func(g *jen.Group) {
		for _, fld := range typ.Fields {
			def := g.Id(fld.Name)
			if fld.IsArray {
				def = def.Index()
			}
			def.Id(fld.ConcreteTypeName())
		}
	})
}
