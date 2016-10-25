package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testTableinfos(t *testing.T) {
	t.Parallel()

	query := Tableinfos(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testTableinfosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tableinfo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfo.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = tableinfo.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Tableinfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTableinfosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tableinfo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfo.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Tableinfos(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Tableinfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTableinfosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tableinfo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TableinfoSlice{tableinfo}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Tableinfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testTableinfosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tableinfo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true, tableinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfo.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := TableinfoExists(tx, tableinfo.TableinfoID)
	if err != nil {
		t.Errorf("Unable to check if Tableinfo exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TableinfoExistsG to return true, but got false.")
	}
}
func testTableinfosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tableinfo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true, tableinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfo.Insert(tx); err != nil {
		t.Error(err)
	}

	tableinfoFound, err := FindTableinfo(tx, tableinfo.TableinfoID)
	if err != nil {
		t.Error(err)
	}

	if tableinfoFound == nil {
		t.Error("want a record, got nil")
	}
}
func testTableinfosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tableinfo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true, tableinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfo.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Tableinfos(tx).Bind(tableinfo); err != nil {
		t.Error(err)
	}
}

func testTableinfosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tableinfo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true, tableinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfo.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Tableinfos(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTableinfosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tableinfoOne := &Tableinfo{}
	tableinfoTwo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfoOne, tableinfoDBTypes, false, tableinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}
	if err = randomize.Struct(seed, tableinfoTwo, tableinfoDBTypes, false, tableinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfoOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = tableinfoTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Tableinfos(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTableinfosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tableinfoOne := &Tableinfo{}
	tableinfoTwo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfoOne, tableinfoDBTypes, false, tableinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}
	if err = randomize.Struct(seed, tableinfoTwo, tableinfoDBTypes, false, tableinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfoOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = tableinfoTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tableinfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func tableinfoBeforeInsertHook(e boil.Executor, o *Tableinfo) error {
	*o = Tableinfo{}
	return nil
}

func tableinfoAfterInsertHook(e boil.Executor, o *Tableinfo) error {
	*o = Tableinfo{}
	return nil
}

func tableinfoAfterSelectHook(e boil.Executor, o *Tableinfo) error {
	*o = Tableinfo{}
	return nil
}

func tableinfoBeforeUpdateHook(e boil.Executor, o *Tableinfo) error {
	*o = Tableinfo{}
	return nil
}

func tableinfoAfterUpdateHook(e boil.Executor, o *Tableinfo) error {
	*o = Tableinfo{}
	return nil
}

func tableinfoBeforeDeleteHook(e boil.Executor, o *Tableinfo) error {
	*o = Tableinfo{}
	return nil
}

func tableinfoAfterDeleteHook(e boil.Executor, o *Tableinfo) error {
	*o = Tableinfo{}
	return nil
}

func tableinfoBeforeUpsertHook(e boil.Executor, o *Tableinfo) error {
	*o = Tableinfo{}
	return nil
}

func tableinfoAfterUpsertHook(e boil.Executor, o *Tableinfo) error {
	*o = Tableinfo{}
	return nil
}

func testTableinfosHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Tableinfo{}
	o := &Tableinfo{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, tableinfoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Tableinfo object: %s", err)
	}

	AddTableinfoHook(boil.BeforeInsertHook, tableinfoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	tableinfoBeforeInsertHooks = []TableinfoHook{}

	AddTableinfoHook(boil.AfterInsertHook, tableinfoAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	tableinfoAfterInsertHooks = []TableinfoHook{}

	AddTableinfoHook(boil.AfterSelectHook, tableinfoAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	tableinfoAfterSelectHooks = []TableinfoHook{}

	AddTableinfoHook(boil.BeforeUpdateHook, tableinfoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	tableinfoBeforeUpdateHooks = []TableinfoHook{}

	AddTableinfoHook(boil.AfterUpdateHook, tableinfoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	tableinfoAfterUpdateHooks = []TableinfoHook{}

	AddTableinfoHook(boil.BeforeDeleteHook, tableinfoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	tableinfoBeforeDeleteHooks = []TableinfoHook{}

	AddTableinfoHook(boil.AfterDeleteHook, tableinfoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	tableinfoAfterDeleteHooks = []TableinfoHook{}

	AddTableinfoHook(boil.BeforeUpsertHook, tableinfoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	tableinfoBeforeUpsertHooks = []TableinfoHook{}

	AddTableinfoHook(boil.AfterUpsertHook, tableinfoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	tableinfoAfterUpsertHooks = []TableinfoHook{}
}
func testTableinfosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tableinfo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true, tableinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tableinfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTableinfosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tableinfo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfo.Insert(tx, tableinfoColumns...); err != nil {
		t.Error(err)
	}

	count, err := Tableinfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTableinfosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tableinfo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true, tableinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfo.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = tableinfo.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testTableinfosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tableinfo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true, tableinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TableinfoSlice{tableinfo}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testTableinfosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tableinfo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true, tableinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Tableinfos(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	tableinfoDBTypes = map[string]string{"IsUpdateable": "integer", "IsView": "integer", "ModificationDate": "date", "Name": "character varying", "PrimaryKeyColumn": "character varying", "SuperclassTableID": "integer", "TableinfoID": "integer", "ViewOnTableID": "integer"}
	_                = bytes.MinRead
)

func testTableinfosUpdate(t *testing.T) {
	t.Parallel()

	if len(tableinfoColumns) == len(tableinfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	tableinfo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tableinfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true, tableinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	if err = tableinfo.Update(tx); err != nil {
		t.Error(err)
	}
}

func testTableinfosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(tableinfoColumns) == len(tableinfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	tableinfo := &Tableinfo{}
	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tableinfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, tableinfo, tableinfoDBTypes, true, tableinfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(tableinfoColumns, tableinfoPrimaryKeyColumns) {
		fields = tableinfoColumns
	} else {
		fields = strmangle.SetComplement(
			tableinfoColumns,
			tableinfoPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(tableinfo))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := TableinfoSlice{tableinfo}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testTableinfosUpsert(t *testing.T) {
	t.Parallel()

	if len(tableinfoColumns) == len(tableinfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	tableinfo := Tableinfo{}
	if err = randomize.Struct(seed, &tableinfo, tableinfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tableinfo.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Tableinfo: %s", err)
	}

	count, err := Tableinfos(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &tableinfo, tableinfoDBTypes, false, tableinfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Tableinfo struct: %s", err)
	}

	if err = tableinfo.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Tableinfo: %s", err)
	}

	count, err = Tableinfos(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
