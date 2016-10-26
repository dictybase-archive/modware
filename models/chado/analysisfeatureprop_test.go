package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testAnalysisfeatureprops(t *testing.T) {
	t.Parallel()

	query := Analysisfeatureprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAnalysisfeaturepropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeatureprop := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = analysisfeatureprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnalysisfeaturepropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeatureprop := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Analysisfeatureprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnalysisfeaturepropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeatureprop := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AnalysisfeaturepropSlice{analysisfeatureprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAnalysisfeaturepropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeatureprop := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AnalysisfeaturepropExists(tx, analysisfeatureprop.AnalysisfeaturepropID)
	if err != nil {
		t.Errorf("Unable to check if Analysisfeatureprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AnalysisfeaturepropExistsG to return true, but got false.")
	}
}
func testAnalysisfeaturepropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeatureprop := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	analysisfeaturepropFound, err := FindAnalysisfeatureprop(tx, analysisfeatureprop.AnalysisfeaturepropID)
	if err != nil {
		t.Error(err)
	}

	if analysisfeaturepropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAnalysisfeaturepropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeatureprop := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Analysisfeatureprops(tx).Bind(analysisfeatureprop); err != nil {
		t.Error(err)
	}
}

func testAnalysisfeaturepropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeatureprop := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Analysisfeatureprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAnalysisfeaturepropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeaturepropOne := &Analysisfeatureprop{}
	analysisfeaturepropTwo := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeaturepropOne, analysisfeaturepropDBTypes, false, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}
	if err = randomize.Struct(seed, analysisfeaturepropTwo, analysisfeaturepropDBTypes, false, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeaturepropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = analysisfeaturepropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Analysisfeatureprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAnalysisfeaturepropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	analysisfeaturepropOne := &Analysisfeatureprop{}
	analysisfeaturepropTwo := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeaturepropOne, analysisfeaturepropDBTypes, false, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}
	if err = randomize.Struct(seed, analysisfeaturepropTwo, analysisfeaturepropDBTypes, false, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeaturepropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = analysisfeaturepropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func analysisfeaturepropBeforeInsertHook(e boil.Executor, o *Analysisfeatureprop) error {
	*o = Analysisfeatureprop{}
	return nil
}

func analysisfeaturepropAfterInsertHook(e boil.Executor, o *Analysisfeatureprop) error {
	*o = Analysisfeatureprop{}
	return nil
}

func analysisfeaturepropAfterSelectHook(e boil.Executor, o *Analysisfeatureprop) error {
	*o = Analysisfeatureprop{}
	return nil
}

func analysisfeaturepropBeforeUpdateHook(e boil.Executor, o *Analysisfeatureprop) error {
	*o = Analysisfeatureprop{}
	return nil
}

func analysisfeaturepropAfterUpdateHook(e boil.Executor, o *Analysisfeatureprop) error {
	*o = Analysisfeatureprop{}
	return nil
}

func analysisfeaturepropBeforeDeleteHook(e boil.Executor, o *Analysisfeatureprop) error {
	*o = Analysisfeatureprop{}
	return nil
}

func analysisfeaturepropAfterDeleteHook(e boil.Executor, o *Analysisfeatureprop) error {
	*o = Analysisfeatureprop{}
	return nil
}

func analysisfeaturepropBeforeUpsertHook(e boil.Executor, o *Analysisfeatureprop) error {
	*o = Analysisfeatureprop{}
	return nil
}

func analysisfeaturepropAfterUpsertHook(e boil.Executor, o *Analysisfeatureprop) error {
	*o = Analysisfeatureprop{}
	return nil
}

func testAnalysisfeaturepropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Analysisfeatureprop{}
	o := &Analysisfeatureprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, analysisfeaturepropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop object: %s", err)
	}

	AddAnalysisfeaturepropHook(boil.BeforeInsertHook, analysisfeaturepropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	analysisfeaturepropBeforeInsertHooks = []AnalysisfeaturepropHook{}

	AddAnalysisfeaturepropHook(boil.AfterInsertHook, analysisfeaturepropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	analysisfeaturepropAfterInsertHooks = []AnalysisfeaturepropHook{}

	AddAnalysisfeaturepropHook(boil.AfterSelectHook, analysisfeaturepropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	analysisfeaturepropAfterSelectHooks = []AnalysisfeaturepropHook{}

	AddAnalysisfeaturepropHook(boil.BeforeUpdateHook, analysisfeaturepropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	analysisfeaturepropBeforeUpdateHooks = []AnalysisfeaturepropHook{}

	AddAnalysisfeaturepropHook(boil.AfterUpdateHook, analysisfeaturepropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	analysisfeaturepropAfterUpdateHooks = []AnalysisfeaturepropHook{}

	AddAnalysisfeaturepropHook(boil.BeforeDeleteHook, analysisfeaturepropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	analysisfeaturepropBeforeDeleteHooks = []AnalysisfeaturepropHook{}

	AddAnalysisfeaturepropHook(boil.AfterDeleteHook, analysisfeaturepropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	analysisfeaturepropAfterDeleteHooks = []AnalysisfeaturepropHook{}

	AddAnalysisfeaturepropHook(boil.BeforeUpsertHook, analysisfeaturepropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	analysisfeaturepropBeforeUpsertHooks = []AnalysisfeaturepropHook{}

	AddAnalysisfeaturepropHook(boil.AfterUpsertHook, analysisfeaturepropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	analysisfeaturepropAfterUpsertHooks = []AnalysisfeaturepropHook{}
}
func testAnalysisfeaturepropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeatureprop := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAnalysisfeaturepropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeatureprop := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureprop.Insert(tx, analysisfeaturepropColumns...); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAnalysisfeaturepropToOneAnalysisfeatureUsingAnalysisfeature(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Analysisfeatureprop
	var foreign Analysisfeature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, analysisfeaturepropDBTypes, true, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, analysisfeatureDBTypes, true, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AnalysisfeatureID = foreign.AnalysisfeatureID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Analysisfeature(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.AnalysisfeatureID != foreign.AnalysisfeatureID {
		t.Errorf("want: %v, got %v", foreign.AnalysisfeatureID, check.AnalysisfeatureID)
	}

	slice := AnalysisfeaturepropSlice{&local}
	if err = local.L.LoadAnalysisfeature(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Analysisfeature == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Analysisfeature = nil
	if err = local.L.LoadAnalysisfeature(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Analysisfeature == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAnalysisfeaturepropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Analysisfeatureprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, analysisfeaturepropDBTypes, true, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
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

	slice := AnalysisfeaturepropSlice{&local}
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

func testAnalysisfeaturepropToOneSetOpAnalysisfeatureUsingAnalysisfeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Analysisfeatureprop
	var b, c Analysisfeature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, analysisfeaturepropDBTypes, false, strmangle.SetComplement(analysisfeaturepropPrimaryKeyColumns, analysisfeaturepropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, analysisfeatureDBTypes, false, strmangle.SetComplement(analysisfeaturePrimaryKeyColumns, analysisfeatureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, analysisfeatureDBTypes, false, strmangle.SetComplement(analysisfeaturePrimaryKeyColumns, analysisfeatureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Analysisfeature{&b, &c} {
		err = a.SetAnalysisfeature(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Analysisfeature != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Analysisfeatureprop != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AnalysisfeatureID != x.AnalysisfeatureID {
			t.Error("foreign key was wrong value", a.AnalysisfeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AnalysisfeatureID))
		reflect.Indirect(reflect.ValueOf(&a.AnalysisfeatureID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AnalysisfeatureID != x.AnalysisfeatureID {
			t.Error("foreign key was wrong value", a.AnalysisfeatureID, x.AnalysisfeatureID)
		}
	}
}
func testAnalysisfeaturepropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Analysisfeatureprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, analysisfeaturepropDBTypes, false, strmangle.SetComplement(analysisfeaturepropPrimaryKeyColumns, analysisfeaturepropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeAnalysisfeatureprop != &a {
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
func testAnalysisfeaturepropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeatureprop := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = analysisfeatureprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAnalysisfeaturepropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeatureprop := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AnalysisfeaturepropSlice{analysisfeatureprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAnalysisfeaturepropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeatureprop := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Analysisfeatureprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	analysisfeaturepropDBTypes = map[string]string{"AnalysisfeatureID": "integer", "AnalysisfeaturepropID": "integer", "Rank": "integer", "TypeID": "integer", "Value": "text"}
	_                          = bytes.MinRead
)

func testAnalysisfeaturepropsUpdate(t *testing.T) {
	t.Parallel()

	if len(analysisfeaturepropColumns) == len(analysisfeaturepropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	analysisfeatureprop := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	if err = analysisfeatureprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAnalysisfeaturepropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(analysisfeaturepropColumns) == len(analysisfeaturepropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	analysisfeatureprop := &Analysisfeatureprop{}
	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, analysisfeatureprop, analysisfeaturepropDBTypes, true, analysisfeaturepropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(analysisfeaturepropColumns, analysisfeaturepropPrimaryKeyColumns) {
		fields = analysisfeaturepropColumns
	} else {
		fields = strmangle.SetComplement(
			analysisfeaturepropColumns,
			analysisfeaturepropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(analysisfeatureprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AnalysisfeaturepropSlice{analysisfeatureprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAnalysisfeaturepropsUpsert(t *testing.T) {
	t.Parallel()

	if len(analysisfeaturepropColumns) == len(analysisfeaturepropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	analysisfeatureprop := Analysisfeatureprop{}
	if err = randomize.Struct(seed, &analysisfeatureprop, analysisfeaturepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Analysisfeatureprop: %s", err)
	}

	count, err := Analysisfeatureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &analysisfeatureprop, analysisfeaturepropDBTypes, false, analysisfeaturepropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}

	if err = analysisfeatureprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Analysisfeatureprop: %s", err)
	}

	count, err = Analysisfeatureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
