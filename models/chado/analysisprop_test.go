package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testAnalysisprops(t *testing.T) {
	t.Parallel()

	query := Analysisprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAnalysispropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisprop := &Analysisprop{}
	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = analysisprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnalysispropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisprop := &Analysisprop{}
	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Analysisprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Analysisprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnalysispropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisprop := &Analysisprop{}
	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AnalysispropSlice{analysisprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAnalysispropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisprop := &Analysisprop{}
	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AnalysispropExists(tx, analysisprop.AnalysispropID)
	if err != nil {
		t.Errorf("Unable to check if Analysisprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AnalysispropExistsG to return true, but got false.")
	}
}
func testAnalysispropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisprop := &Analysisprop{}
	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisprop.Insert(tx); err != nil {
		t.Error(err)
	}

	analysispropFound, err := FindAnalysisprop(tx, analysisprop.AnalysispropID)
	if err != nil {
		t.Error(err)
	}

	if analysispropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAnalysispropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisprop := &Analysisprop{}
	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Analysisprops(tx).Bind(analysisprop); err != nil {
		t.Error(err)
	}
}

func testAnalysispropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisprop := &Analysisprop{}
	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Analysisprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAnalysispropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysispropOne := &Analysisprop{}
	analysispropTwo := &Analysisprop{}
	if err = randomize.Struct(seed, analysispropOne, analysispropDBTypes, false, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}
	if err = randomize.Struct(seed, analysispropTwo, analysispropDBTypes, false, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysispropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = analysispropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Analysisprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAnalysispropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	analysispropOne := &Analysisprop{}
	analysispropTwo := &Analysisprop{}
	if err = randomize.Struct(seed, analysispropOne, analysispropDBTypes, false, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}
	if err = randomize.Struct(seed, analysispropTwo, analysispropDBTypes, false, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysispropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = analysispropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func analysispropBeforeInsertHook(e boil.Executor, o *Analysisprop) error {
	*o = Analysisprop{}
	return nil
}

func analysispropAfterInsertHook(e boil.Executor, o *Analysisprop) error {
	*o = Analysisprop{}
	return nil
}

func analysispropAfterSelectHook(e boil.Executor, o *Analysisprop) error {
	*o = Analysisprop{}
	return nil
}

func analysispropBeforeUpdateHook(e boil.Executor, o *Analysisprop) error {
	*o = Analysisprop{}
	return nil
}

func analysispropAfterUpdateHook(e boil.Executor, o *Analysisprop) error {
	*o = Analysisprop{}
	return nil
}

func analysispropBeforeDeleteHook(e boil.Executor, o *Analysisprop) error {
	*o = Analysisprop{}
	return nil
}

func analysispropAfterDeleteHook(e boil.Executor, o *Analysisprop) error {
	*o = Analysisprop{}
	return nil
}

func analysispropBeforeUpsertHook(e boil.Executor, o *Analysisprop) error {
	*o = Analysisprop{}
	return nil
}

func analysispropAfterUpsertHook(e boil.Executor, o *Analysisprop) error {
	*o = Analysisprop{}
	return nil
}

func testAnalysispropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Analysisprop{}
	o := &Analysisprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, analysispropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Analysisprop object: %s", err)
	}

	AddAnalysispropHook(boil.BeforeInsertHook, analysispropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	analysispropBeforeInsertHooks = []AnalysispropHook{}

	AddAnalysispropHook(boil.AfterInsertHook, analysispropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	analysispropAfterInsertHooks = []AnalysispropHook{}

	AddAnalysispropHook(boil.AfterSelectHook, analysispropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	analysispropAfterSelectHooks = []AnalysispropHook{}

	AddAnalysispropHook(boil.BeforeUpdateHook, analysispropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	analysispropBeforeUpdateHooks = []AnalysispropHook{}

	AddAnalysispropHook(boil.AfterUpdateHook, analysispropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	analysispropAfterUpdateHooks = []AnalysispropHook{}

	AddAnalysispropHook(boil.BeforeDeleteHook, analysispropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	analysispropBeforeDeleteHooks = []AnalysispropHook{}

	AddAnalysispropHook(boil.AfterDeleteHook, analysispropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	analysispropAfterDeleteHooks = []AnalysispropHook{}

	AddAnalysispropHook(boil.BeforeUpsertHook, analysispropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	analysispropBeforeUpsertHooks = []AnalysispropHook{}

	AddAnalysispropHook(boil.AfterUpsertHook, analysispropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	analysispropAfterUpsertHooks = []AnalysispropHook{}
}
func testAnalysispropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisprop := &Analysisprop{}
	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAnalysispropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisprop := &Analysisprop{}
	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisprop.Insert(tx, analysispropColumns...); err != nil {
		t.Error(err)
	}

	count, err := Analysisprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAnalysispropToOneAnalysiUsingAnalysi(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Analysisprop
	var foreign Analysi

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, analysispropDBTypes, true, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, analysiDBTypes, true, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AnalysisID = foreign.AnalysisID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Analysi(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.AnalysisID != foreign.AnalysisID {
		t.Errorf("want: %v, got %v", foreign.AnalysisID, check.AnalysisID)
	}

	slice := AnalysispropSlice{&local}
	if err = local.L.LoadAnalysi(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Analysi == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Analysi = nil
	if err = local.L.LoadAnalysi(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Analysi == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAnalysispropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Analysisprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, analysispropDBTypes, true, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
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

	slice := AnalysispropSlice{&local}
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

func testAnalysispropToOneSetOpAnalysiUsingAnalysi(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Analysisprop
	var b, c Analysi

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, analysispropDBTypes, false, strmangle.SetComplement(analysispropPrimaryKeyColumns, analysispropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, analysiDBTypes, false, strmangle.SetComplement(analysiPrimaryKeyColumns, analysiColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, analysiDBTypes, false, strmangle.SetComplement(analysiPrimaryKeyColumns, analysiColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Analysi{&b, &c} {
		err = a.SetAnalysi(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Analysi != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Analysisprop != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AnalysisID != x.AnalysisID {
			t.Error("foreign key was wrong value", a.AnalysisID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AnalysisID))
		reflect.Indirect(reflect.ValueOf(&a.AnalysisID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AnalysisID != x.AnalysisID {
			t.Error("foreign key was wrong value", a.AnalysisID, x.AnalysisID)
		}
	}
}
func testAnalysispropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Analysisprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, analysispropDBTypes, false, strmangle.SetComplement(analysispropPrimaryKeyColumns, analysispropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeAnalysisprop != &a {
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
func testAnalysispropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisprop := &Analysisprop{}
	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = analysisprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAnalysispropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisprop := &Analysisprop{}
	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AnalysispropSlice{analysisprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAnalysispropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisprop := &Analysisprop{}
	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Analysisprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	analysispropDBTypes = map[string]string{"AnalysisID": "integer", "AnalysispropID": "integer", "Rank": "integer", "TypeID": "integer", "Value": "text"}
	_                   = bytes.MinRead
)

func testAnalysispropsUpdate(t *testing.T) {
	t.Parallel()

	if len(analysispropColumns) == len(analysispropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	analysisprop := &Analysisprop{}
	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	if err = analysisprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAnalysispropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(analysispropColumns) == len(analysispropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	analysisprop := &Analysisprop{}
	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, analysisprop, analysispropDBTypes, true, analysispropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(analysispropColumns, analysispropPrimaryKeyColumns) {
		fields = analysispropColumns
	} else {
		fields = strmangle.SetComplement(
			analysispropColumns,
			analysispropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(analysisprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AnalysispropSlice{analysisprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAnalysispropsUpsert(t *testing.T) {
	t.Parallel()

	if len(analysispropColumns) == len(analysispropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	analysisprop := Analysisprop{}
	if err = randomize.Struct(seed, &analysisprop, analysispropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Analysisprop: %s", err)
	}

	count, err := Analysisprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &analysisprop, analysispropDBTypes, false, analysispropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}

	if err = analysisprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Analysisprop: %s", err)
	}

	count, err = Analysisprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
