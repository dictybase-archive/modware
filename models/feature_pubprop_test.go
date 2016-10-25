package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeaturePubprops(t *testing.T) {
	t.Parallel()

	query := FeaturePubprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeaturePubpropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePubprop := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featurePubprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturePubpropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePubprop := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeaturePubprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturePubpropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePubprop := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeaturePubpropSlice{featurePubprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeaturePubpropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePubprop := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeaturePubpropExists(tx, featurePubprop.FeaturePubpropID)
	if err != nil {
		t.Errorf("Unable to check if FeaturePubprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeaturePubpropExistsG to return true, but got false.")
	}
}
func testFeaturePubpropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePubprop := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	featurePubpropFound, err := FindFeaturePubprop(tx, featurePubprop.FeaturePubpropID)
	if err != nil {
		t.Error(err)
	}

	if featurePubpropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeaturePubpropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePubprop := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeaturePubprops(tx).Bind(featurePubprop); err != nil {
		t.Error(err)
	}
}

func testFeaturePubpropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePubprop := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeaturePubprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeaturePubpropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePubpropOne := &FeaturePubprop{}
	featurePubpropTwo := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubpropOne, featurePubpropDBTypes, false, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}
	if err = randomize.Struct(seed, featurePubpropTwo, featurePubpropDBTypes, false, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featurePubpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeaturePubprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeaturePubpropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featurePubpropOne := &FeaturePubprop{}
	featurePubpropTwo := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubpropOne, featurePubpropDBTypes, false, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}
	if err = randomize.Struct(seed, featurePubpropTwo, featurePubpropDBTypes, false, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featurePubpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featurePubpropBeforeInsertHook(e boil.Executor, o *FeaturePubprop) error {
	*o = FeaturePubprop{}
	return nil
}

func featurePubpropAfterInsertHook(e boil.Executor, o *FeaturePubprop) error {
	*o = FeaturePubprop{}
	return nil
}

func featurePubpropAfterSelectHook(e boil.Executor, o *FeaturePubprop) error {
	*o = FeaturePubprop{}
	return nil
}

func featurePubpropBeforeUpdateHook(e boil.Executor, o *FeaturePubprop) error {
	*o = FeaturePubprop{}
	return nil
}

func featurePubpropAfterUpdateHook(e boil.Executor, o *FeaturePubprop) error {
	*o = FeaturePubprop{}
	return nil
}

func featurePubpropBeforeDeleteHook(e boil.Executor, o *FeaturePubprop) error {
	*o = FeaturePubprop{}
	return nil
}

func featurePubpropAfterDeleteHook(e boil.Executor, o *FeaturePubprop) error {
	*o = FeaturePubprop{}
	return nil
}

func featurePubpropBeforeUpsertHook(e boil.Executor, o *FeaturePubprop) error {
	*o = FeaturePubprop{}
	return nil
}

func featurePubpropAfterUpsertHook(e boil.Executor, o *FeaturePubprop) error {
	*o = FeaturePubprop{}
	return nil
}

func testFeaturePubpropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeaturePubprop{}
	o := &FeaturePubprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featurePubpropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop object: %s", err)
	}

	AddFeaturePubpropHook(boil.BeforeInsertHook, featurePubpropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featurePubpropBeforeInsertHooks = []FeaturePubpropHook{}

	AddFeaturePubpropHook(boil.AfterInsertHook, featurePubpropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featurePubpropAfterInsertHooks = []FeaturePubpropHook{}

	AddFeaturePubpropHook(boil.AfterSelectHook, featurePubpropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featurePubpropAfterSelectHooks = []FeaturePubpropHook{}

	AddFeaturePubpropHook(boil.BeforeUpdateHook, featurePubpropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featurePubpropBeforeUpdateHooks = []FeaturePubpropHook{}

	AddFeaturePubpropHook(boil.AfterUpdateHook, featurePubpropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featurePubpropAfterUpdateHooks = []FeaturePubpropHook{}

	AddFeaturePubpropHook(boil.BeforeDeleteHook, featurePubpropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featurePubpropBeforeDeleteHooks = []FeaturePubpropHook{}

	AddFeaturePubpropHook(boil.AfterDeleteHook, featurePubpropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featurePubpropAfterDeleteHooks = []FeaturePubpropHook{}

	AddFeaturePubpropHook(boil.BeforeUpsertHook, featurePubpropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featurePubpropBeforeUpsertHooks = []FeaturePubpropHook{}

	AddFeaturePubpropHook(boil.AfterUpsertHook, featurePubpropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featurePubpropAfterUpsertHooks = []FeaturePubpropHook{}
}
func testFeaturePubpropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePubprop := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeaturePubpropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePubprop := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubprop.Insert(tx, featurePubpropColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeaturePubpropToOneFeaturePubUsingFeaturePub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeaturePubprop
	var foreign FeaturePub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featurePubpropDBTypes, true, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, featurePubDBTypes, true, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.FeaturePubID = foreign.FeaturePubID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeaturePub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeaturePubID != foreign.FeaturePubID {
		t.Errorf("want: %v, got %v", foreign.FeaturePubID, check.FeaturePubID)
	}

	slice := FeaturePubpropSlice{&local}
	if err = local.L.LoadFeaturePub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturePub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeaturePub = nil
	if err = local.L.LoadFeaturePub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturePub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeaturePubpropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeaturePubprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featurePubpropDBTypes, true, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
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

	slice := FeaturePubpropSlice{&local}
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

func testFeaturePubpropToOneSetOpFeaturePubUsingFeaturePub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeaturePubprop
	var b, c FeaturePub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurePubpropDBTypes, false, strmangle.SetComplement(featurePubpropPrimaryKeyColumns, featurePubpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featurePubDBTypes, false, strmangle.SetComplement(featurePubPrimaryKeyColumns, featurePubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featurePubDBTypes, false, strmangle.SetComplement(featurePubPrimaryKeyColumns, featurePubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeaturePub{&b, &c} {
		err = a.SetFeaturePub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeaturePub != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FeaturePubprop != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FeaturePubID != x.FeaturePubID {
			t.Error("foreign key was wrong value", a.FeaturePubID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FeaturePubID))
		reflect.Indirect(reflect.ValueOf(&a.FeaturePubID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeaturePubID != x.FeaturePubID {
			t.Error("foreign key was wrong value", a.FeaturePubID, x.FeaturePubID)
		}
	}
}
func testFeaturePubpropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeaturePubprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurePubpropDBTypes, false, strmangle.SetComplement(featurePubpropPrimaryKeyColumns, featurePubpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeFeaturePubprop != &a {
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
func testFeaturePubpropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePubprop := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featurePubprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturePubpropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePubprop := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeaturePubpropSlice{featurePubprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeaturePubpropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePubprop := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeaturePubprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featurePubpropDBTypes = map[string]string{"FeaturePubID": "integer", "FeaturePubpropID": "integer", "Rank": "integer", "TypeID": "integer", "Value": "text"}
	_                     = bytes.MinRead
)

func testFeaturePubpropsUpdate(t *testing.T) {
	t.Parallel()

	if len(featurePubpropColumns) == len(featurePubpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featurePubprop := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	if err = featurePubprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturePubpropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featurePubpropColumns) == len(featurePubpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featurePubprop := &FeaturePubprop{}
	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featurePubprop, featurePubpropDBTypes, true, featurePubpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featurePubpropColumns, featurePubpropPrimaryKeyColumns) {
		fields = featurePubpropColumns
	} else {
		fields = strmangle.SetComplement(
			featurePubpropColumns,
			featurePubpropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featurePubprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeaturePubpropSlice{featurePubprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeaturePubpropsUpsert(t *testing.T) {
	t.Parallel()

	if len(featurePubpropColumns) == len(featurePubpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featurePubprop := FeaturePubprop{}
	if err = randomize.Struct(seed, &featurePubprop, featurePubpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeaturePubprop: %s", err)
	}

	count, err := FeaturePubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featurePubprop, featurePubpropDBTypes, false, featurePubpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}

	if err = featurePubprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeaturePubprop: %s", err)
	}

	count, err = FeaturePubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
