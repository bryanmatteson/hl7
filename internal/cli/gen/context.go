package gen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// LoadContext ...
func LoadContext(contextPath string) (ctx *Context, err error) {
	if !pathExists(contextPath) {
		err = fmt.Errorf("context directory does not exist")
		return
	}

	ctx = &Context{TypeMap: make(map[string]*GoType)}

	dataTypes := make([]APIDataTypeItem, 0)
	if err = loadJSONFile(path.Join(contextPath, "datatypes.json"), &dataTypes); err != nil {
		return
	}
	for _, dt := range dataTypes {
		typeName := goName(dt.ID)
		if _, ok := ctx.TypeMap[typeName]; ok {
			continue
		}
		goType := &GoType{Name: typeName, Kind: GoDataTypeKind, Title: dt.Name, Description: dt.Description}
		ctx.TypeMap[typeName] = goType

		for _, f := range dt.Fields {
			field := genGoFieldFromFieldInfo(f)
			goType.Fields = append(goType.Fields, field)
		}
	}
	dataTypes = nil

	segments := make([]APISegmentItem, 0)
	if err = loadJSONFile(path.Join(contextPath, "segments.json"), &segments); err != nil {
		return
	}
	for _, seg := range segments {
		typeName := goName(seg.ID)
		if _, ok := ctx.TypeMap[typeName]; ok {
			continue
		}
		goType := &GoType{Name: typeName, Kind: GoSegmentKind, Title: seg.LongName, Description: seg.Description}

		ctx.TypeMap[typeName] = goType
		for _, f := range seg.Fields {
			field := genGoFieldFromFieldInfo(f)
			goType.Fields = append(goType.Fields, field)
		}
	}
	segments = nil

	events := make([]APIEventItem, 0)
	if err = loadJSONFile(path.Join(contextPath, "events.json"), &events); err != nil {
		return
	}
	for _, evt := range events {
		ns := goName(evt.ID)

		var grpGen func(s *APISegmentInfo, fld *GoField)
		grpGen = func(s *APISegmentInfo, fld *GoField) {
			fld.TypeName = ns + fld.TypeName
			if _, ok := ctx.TypeMap[fld.TypeName]; ok {
				return
			}
			grp := &GoType{Name: fld.TypeName, Kind: GoGroupKind, Title: "Group"}
			ctx.TypeMap[fld.TypeName] = grp
			for _, seg := range s.Segments {
				segField := genGoFieldFromSegInfo(seg)
				if seg.IsGroup {
					grpGen(seg, segField)
				}
				grp.Fields = append(grp.Fields, segField)
			}
		}

		if _, ok := ctx.TypeMap[ns]; ok {
			continue
		}
		goType := &GoType{Name: ns, Kind: GoMessageKind, Title: evt.EventDesc, Description: evt.Description}
		ctx.TypeMap[ns] = goType
		for _, seg := range evt.Segments {
			field := genGoFieldFromSegInfo(seg)
			if seg.IsGroup {
				grpGen(seg, field)
			}
			goType.Fields = append(goType.Fields, field)
		}
	}
	events = nil

	for _, typ := range ctx.TypeMap {
		for _, fld := range typ.Fields {
			if fld.Type == nil {
				fld.Type = ctx.TypeMap[fld.TypeName]
			}
		}
	}

	return
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// loadJSONFile ...
func loadJSONFile(path string, data interface{}) (err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, data)
	return
}
