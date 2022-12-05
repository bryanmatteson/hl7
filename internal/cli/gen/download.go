package gen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// DownloadContext ...
func DownloadContext(version HL7Version, contextPath string) error {
	if stat, err := os.Stat(contextPath); err != nil || !stat.IsDir() {
		return fmt.Errorf("context path must be a directory: %s", contextPath)
	}

	client := NewCaristixClient(version)

	// get data types
	items, err := client.GetDataTypesList()
	if err != nil {
		return err
	}

	var dataTypes []*APIDataTypeItem
	for _, item := range items {
		dt, err := client.GetDataTypeDetail(item.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		dataTypes = append(dataTypes, &dt)
	}

	bytes, err := json.Marshal(dataTypes)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(contextPath, "datatypes.json"), bytes, os.ModePerm); err != nil {
		return err
	}

	// get segments
	items, err = client.GetSegmentsList()
	if err != nil {
		return err
	}

	var segments []*APISegmentItem
	for _, item := range items {
		detail, err := client.GetSegmentDetail(item.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		segments = append(segments, &detail)
	}

	bytes, err = json.Marshal(segments)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(contextPath, "segments.json"), bytes, os.ModePerm); err != nil {
		return err
	}

	// get events
	items, err = client.GetTriggerEventsList()
	if err != nil {
		return err
	}

	var events []*APIEventItem
	for _, item := range items {
		detail, err := client.GetTriggerEventDetail(item.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		events = append(events, &detail)
	}

	bytes, err = json.Marshal(events)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(contextPath, "events.json"), bytes, os.ModePerm); err != nil {
		return err
	}

	// get tables
	items, err = client.GetTablesList()
	if err != nil {
		return err
	}

	var tables []*APITableItem
	for _, item := range items {
		detail, err := client.GetTableDetail(item.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		tables = append(tables, &detail)
	}

	bytes, err = json.Marshal(tables)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(contextPath, "tables.json"), bytes, os.ModePerm); err != nil {
		return err
	}

	fieldsMap := map[string]*APIFieldItem{}
	// get fields from data types
	for _, dt := range dataTypes {
		for _, f := range dt.Fields {
			if _, ok := fieldsMap[f.ID]; !ok {
				fld, err := client.GetDataTypeFieldDetail(f.ID)
				if err != nil {
					log.Println(err)
					continue
				}
				fieldsMap[f.ID] = &fld
			}
		}
	}

	// get fields from segments
	for _, seg := range segments {
		for _, f := range seg.Fields {
			if _, ok := fieldsMap[f.ID]; !ok {
				fld, err := client.GetFieldDetail(f.ID)
				if err != nil {
					log.Println(err)
					continue
				}
				fieldsMap[f.ID] = &fld
			}
		}
	}

	var fields []*APIFieldItem
	for _, f := range fieldsMap {
		fields = append(fields, f)
	}

	bytes, err = json.Marshal(fields)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(contextPath, "fields.json"), bytes, os.ModePerm); err != nil {
		return err
	}

	return nil
}
