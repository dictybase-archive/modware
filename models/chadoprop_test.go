package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testChadoprops(t *testing.T) {
	t.Parallel()

	query := Chadoprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testChadopropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chadoprop := &Chadoprop{}
	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadoprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = chadoprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Chadoprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChadopropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chadoprop := &Chadoprop{}
	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadoprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Chadoprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Chadoprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChadopropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chadoprop := &Chadoprop{}
	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadoprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ChadopropSlice{chadoprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Chadoprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testChadopropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chadoprop := &Chadoprop{}
	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true, chadopropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadoprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := ChadopropExists(tx, chadoprop.ChadopropID)
	if err != nil {
		t.Errorf("Unable to check if Chadoprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ChadopropExistsG to return true, but got false.")
	}
}
func testChadopropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chadoprop := &Chadoprop{}
	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true, chadopropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadoprop.Insert(tx); err != nil {
		t.Error(err)
	}

	chadopropFound, err := FindChadoprop(tx, chadoprop.ChadopropID)
	if err != nil {
		t.Error(err)
	}

	if chadopropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testChadopropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chadoprop := &Chadoprop{}
	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true, chadopropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadoprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Chadoprops(tx).Bind(chadoprop); err != nil {
		t.Error(err)
	}
}

func testChadopropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chadoprop := &Chadoprop{}
	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true, chadopropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadoprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Chadoprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testChadopropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chadopropOne := &Chadoprop{}
	chadopropTwo := &Chadoprop{}
	if err = randomize.Struct(seed, chadopropOne, chadopropDBTypes, false, chadopropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}
	if err = randomize.Struct(seed, chadopropTwo, chadopropDBTypes, false, chadopropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadopropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = chadopropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Chadoprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testChadopropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	chadopropOne := &Chadoprop{}
	chadopropTwo := &Chadoprop{}
	if err = randomize.Struct(seed, chadopropOne, chadopropDBTypes, false, chadopropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}
	if err = randomize.Struct(seed, chadopropTwo, chadopropDBTypes, false, chadopropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadopropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = chadopropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Chadoprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func chadopropBeforeInsertHook(e boil.Executor, o *Chadoprop) error {
	*o = Chadoprop{}
	return nil
}

func chadopropAfterInsertHook(e boil.Executor, o *Chadoprop) error {
	*o = Chadoprop{}
	return nil
}

func chadopropAfterSelectHook(e boil.Executor, o *Chadoprop) error {
	*o = Chadoprop{}
	return nil
}

func chadopropBeforeUpdateHook(e boil.Executor, o *Chadoprop) error {
	*o = Chadoprop{}
	return nil
}

func chadopropAfterUpdateHook(e boil.Executor, o *Chadoprop) error {
	*o = Chadoprop{}
	return nil
}

func chadopropBeforeDeleteHook(e boil.Executor, o *Chadoprop) error {
	*o = Chadoprop{}
	return nil
}

func chadopropAfterDeleteHook(e boil.Executor, o *Chadoprop) error {
	*o = Chadoprop{}
	return nil
}

func chadopropBeforeUpsertHook(e boil.Executor, o *Chadoprop) error {
	*o = Chadoprop{}
	return nil
}

func chadopropAfterUpsertHook(e boil.Executor, o *Chadoprop) error {
	*o = Chadoprop{}
	return nil
}

func testChadopropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Chadoprop{}
	o := &Chadoprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, chadopropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Chadoprop object: %s", err)
	}

	AddChadopropHook(boil.BeforeInsertHook, chadopropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	chadopropBeforeInsertHooks = []ChadopropHook{}

	AddChadopropHook(boil.AfterInsertHook, chadopropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	chadopropAfterInsertHooks = []ChadopropHook{}

	AddChadopropHook(boil.AfterSelectHook, chadopropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	chadopropAfterSelectHooks = []ChadopropHook{}

	AddChadopropHook(boil.BeforeUpdateHook, chadopropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	chadopropBeforeUpdateHooks = []ChadopropHook{}

	AddChadopropHook(boil.AfterUpdateHook, chadopropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	chadopropAfterUpdateHooks = []ChadopropHook{}

	AddChadopropHook(boil.BeforeDeleteHook, chadopropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	chadopropBeforeDeleteHooks = []ChadopropHook{}

	AddChadopropHook(boil.AfterDeleteHook, chadopropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	chadopropAfterDeleteHooks = []ChadopropHook{}

	AddChadopropHook(boil.BeforeUpsertHook, chadopropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	chadopropBeforeUpsertHooks = []ChadopropHook{}

	AddChadopropHook(boil.AfterUpsertHook, chadopropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	chadopropAfterUpsertHooks = []ChadopropHook{}
}
func testChadopropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chadoprop := &Chadoprop{}
	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true, chadopropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadoprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Chadoprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testChadopropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chadoprop := &Chadoprop{}
	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadoprop.Insert(tx, chadopropColumns...); err != nil {
		t.Error(err)
	}

	count, err := Chadoprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testChadopropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Chadoprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, chadopropDBTypes, true, chadopropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.TypeID = foreign.CvtermID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Type(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := ChadopropSlice{&local}
	if err = local.L.LoadType(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Type == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Type = nil
	if err = local.L.LoadType(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Type == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testChadopropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Chadoprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chadopropDBTypes, false, strmangle.SetComplement(chadopropPrimaryKeyColumns, chadopropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Cvterm{&b, &c} {
		err = a.SetType(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Type != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TypeChadoprop != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TypeID != x.CvtermID {
			t.Error("foreign key was wrong value", a.TypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TypeID))
		reflect.Indirect(reflect.ValueOf(&a.TypeID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TypeID != x.CvtermID {
			t.Error("foreign key was wrong value", a.TypeID, x.CvtermID)
		}
	}
}
func testChadopropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chadoprop := &Chadoprop{}
	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true, chadopropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadoprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = chadoprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testChadopropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chadoprop := &Chadoprop{}
	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true, chadopropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadoprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ChadopropSlice{chadoprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testChadopropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chadoprop := &Chadoprop{}
	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true, chadopropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadoprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Chadoprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	chadopropDBTypes = map[string]string{"ChadopropID": "integer", "Rank": "integer", "TypeID": "integer", "Value": "text"}
	_                = bytes.MinRead
)

func testChadopropsUpdate(t *testing.T) {
	t.Parallel()

	if len(chadopropColumns) == len(chadopropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	chadoprop := &Chadoprop{}
	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadoprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Chadoprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true, chadopropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	if err = chadoprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testChadopropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(chadopropColumns) == len(chadopropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	chadoprop := &Chadoprop{}
	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadoprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Chadoprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, chadoprop, chadopropDBTypes, true, chadopropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(chadopropColumns, chadopropPrimaryKeyColumns) {
		fields = chadopropColumns
	} else {
		fields = strmangle.SetComplement(
			chadopropColumns,
			chadopropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(chadoprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := ChadopropSlice{chadoprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testChadopropsUpsert(t *testing.T) {
	t.Parallel()

	if len(chadopropColumns) == len(chadopropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	chadoprop := Chadoprop{}
	if err = randomize.Struct(seed, &chadoprop, chadopropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = chadoprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Chadoprop: %s", err)
	}

	count, err := Chadoprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &chadoprop, chadopropDBTypes, false, chadopropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}

	if err = chadoprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Chadoprop: %s", err)
	}

	count, err = Chadoprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
