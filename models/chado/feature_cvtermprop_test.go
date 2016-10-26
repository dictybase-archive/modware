package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeatureCvtermprops(t *testing.T) {
	t.Parallel()

	query := FeatureCvtermprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeatureCvtermpropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermprop := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureCvtermprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureCvtermpropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermprop := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureCvtermprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureCvtermpropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermprop := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureCvtermpropSlice{featureCvtermprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeatureCvtermpropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermprop := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeatureCvtermpropExists(tx, featureCvtermprop.FeatureCvtermpropID)
	if err != nil {
		t.Errorf("Unable to check if FeatureCvtermprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeatureCvtermpropExistsG to return true, but got false.")
	}
}
func testFeatureCvtermpropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermprop := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	featureCvtermpropFound, err := FindFeatureCvtermprop(tx, featureCvtermprop.FeatureCvtermpropID)
	if err != nil {
		t.Error(err)
	}

	if featureCvtermpropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeatureCvtermpropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermprop := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureCvtermprops(tx).Bind(featureCvtermprop); err != nil {
		t.Error(err)
	}
}

func testFeatureCvtermpropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermprop := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeatureCvtermprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeatureCvtermpropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermpropOne := &FeatureCvtermprop{}
	featureCvtermpropTwo := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermpropOne, featureCvtermpropDBTypes, false, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}
	if err = randomize.Struct(seed, featureCvtermpropTwo, featureCvtermpropDBTypes, false, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureCvtermpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureCvtermprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeatureCvtermpropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featureCvtermpropOne := &FeatureCvtermprop{}
	featureCvtermpropTwo := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermpropOne, featureCvtermpropDBTypes, false, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}
	if err = randomize.Struct(seed, featureCvtermpropTwo, featureCvtermpropDBTypes, false, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureCvtermpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featureCvtermpropBeforeInsertHook(e boil.Executor, o *FeatureCvtermprop) error {
	*o = FeatureCvtermprop{}
	return nil
}

func featureCvtermpropAfterInsertHook(e boil.Executor, o *FeatureCvtermprop) error {
	*o = FeatureCvtermprop{}
	return nil
}

func featureCvtermpropAfterSelectHook(e boil.Executor, o *FeatureCvtermprop) error {
	*o = FeatureCvtermprop{}
	return nil
}

func featureCvtermpropBeforeUpdateHook(e boil.Executor, o *FeatureCvtermprop) error {
	*o = FeatureCvtermprop{}
	return nil
}

func featureCvtermpropAfterUpdateHook(e boil.Executor, o *FeatureCvtermprop) error {
	*o = FeatureCvtermprop{}
	return nil
}

func featureCvtermpropBeforeDeleteHook(e boil.Executor, o *FeatureCvtermprop) error {
	*o = FeatureCvtermprop{}
	return nil
}

func featureCvtermpropAfterDeleteHook(e boil.Executor, o *FeatureCvtermprop) error {
	*o = FeatureCvtermprop{}
	return nil
}

func featureCvtermpropBeforeUpsertHook(e boil.Executor, o *FeatureCvtermprop) error {
	*o = FeatureCvtermprop{}
	return nil
}

func featureCvtermpropAfterUpsertHook(e boil.Executor, o *FeatureCvtermprop) error {
	*o = FeatureCvtermprop{}
	return nil
}

func testFeatureCvtermpropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeatureCvtermprop{}
	o := &FeatureCvtermprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featureCvtermpropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop object: %s", err)
	}

	AddFeatureCvtermpropHook(boil.BeforeInsertHook, featureCvtermpropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermpropBeforeInsertHooks = []FeatureCvtermpropHook{}

	AddFeatureCvtermpropHook(boil.AfterInsertHook, featureCvtermpropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermpropAfterInsertHooks = []FeatureCvtermpropHook{}

	AddFeatureCvtermpropHook(boil.AfterSelectHook, featureCvtermpropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featureCvtermpropAfterSelectHooks = []FeatureCvtermpropHook{}

	AddFeatureCvtermpropHook(boil.BeforeUpdateHook, featureCvtermpropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featureCvtermpropBeforeUpdateHooks = []FeatureCvtermpropHook{}

	AddFeatureCvtermpropHook(boil.AfterUpdateHook, featureCvtermpropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featureCvtermpropAfterUpdateHooks = []FeatureCvtermpropHook{}

	AddFeatureCvtermpropHook(boil.BeforeDeleteHook, featureCvtermpropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featureCvtermpropBeforeDeleteHooks = []FeatureCvtermpropHook{}

	AddFeatureCvtermpropHook(boil.AfterDeleteHook, featureCvtermpropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featureCvtermpropAfterDeleteHooks = []FeatureCvtermpropHook{}

	AddFeatureCvtermpropHook(boil.BeforeUpsertHook, featureCvtermpropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermpropBeforeUpsertHooks = []FeatureCvtermpropHook{}

	AddFeatureCvtermpropHook(boil.AfterUpsertHook, featureCvtermpropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermpropAfterUpsertHooks = []FeatureCvtermpropHook{}
}
func testFeatureCvtermpropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermprop := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureCvtermpropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermprop := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermprop.Insert(tx, featureCvtermpropColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureCvtermpropToOneFeatureCvtermUsingFeatureCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureCvtermprop
	var foreign FeatureCvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureCvtermpropDBTypes, true, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.FeatureCvtermID = foreign.FeatureCvtermID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureCvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureCvtermID != foreign.FeatureCvtermID {
		t.Errorf("want: %v, got %v", foreign.FeatureCvtermID, check.FeatureCvtermID)
	}

	slice := FeatureCvtermpropSlice{&local}
	if err = local.L.LoadFeatureCvterm(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureCvterm == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureCvterm = nil
	if err = local.L.LoadFeatureCvterm(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureCvterm == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureCvtermpropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureCvtermprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureCvtermpropDBTypes, true, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
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

	slice := FeatureCvtermpropSlice{&local}
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

func testFeatureCvtermpropToOneSetOpFeatureCvtermUsingFeatureCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureCvtermprop
	var b, c FeatureCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureCvtermpropDBTypes, false, strmangle.SetComplement(featureCvtermpropPrimaryKeyColumns, featureCvtermpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureCvtermDBTypes, false, strmangle.SetComplement(featureCvtermPrimaryKeyColumns, featureCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureCvtermDBTypes, false, strmangle.SetComplement(featureCvtermPrimaryKeyColumns, featureCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureCvterm{&b, &c} {
		err = a.SetFeatureCvterm(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureCvterm != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FeatureCvtermprop != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FeatureCvtermID != x.FeatureCvtermID {
			t.Error("foreign key was wrong value", a.FeatureCvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FeatureCvtermID))
		reflect.Indirect(reflect.ValueOf(&a.FeatureCvtermID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureCvtermID != x.FeatureCvtermID {
			t.Error("foreign key was wrong value", a.FeatureCvtermID, x.FeatureCvtermID)
		}
	}
}
func testFeatureCvtermpropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureCvtermprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureCvtermpropDBTypes, false, strmangle.SetComplement(featureCvtermpropPrimaryKeyColumns, featureCvtermpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeFeatureCvtermprop != &a {
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
func testFeatureCvtermpropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermprop := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureCvtermprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureCvtermpropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermprop := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureCvtermpropSlice{featureCvtermprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeatureCvtermpropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermprop := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureCvtermprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featureCvtermpropDBTypes = map[string]string{"FeatureCvtermID": "integer", "FeatureCvtermpropID": "integer", "Rank": "integer", "TypeID": "integer", "Value": "text"}
	_                        = bytes.MinRead
)

func testFeatureCvtermpropsUpdate(t *testing.T) {
	t.Parallel()

	if len(featureCvtermpropColumns) == len(featureCvtermpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureCvtermprop := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	if err = featureCvtermprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureCvtermpropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featureCvtermpropColumns) == len(featureCvtermpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureCvtermprop := &FeatureCvtermprop{}
	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureCvtermprop, featureCvtermpropDBTypes, true, featureCvtermpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featureCvtermpropColumns, featureCvtermpropPrimaryKeyColumns) {
		fields = featureCvtermpropColumns
	} else {
		fields = strmangle.SetComplement(
			featureCvtermpropColumns,
			featureCvtermpropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featureCvtermprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeatureCvtermpropSlice{featureCvtermprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeatureCvtermpropsUpsert(t *testing.T) {
	t.Parallel()

	if len(featureCvtermpropColumns) == len(featureCvtermpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featureCvtermprop := FeatureCvtermprop{}
	if err = randomize.Struct(seed, &featureCvtermprop, featureCvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureCvtermprop: %s", err)
	}

	count, err := FeatureCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featureCvtermprop, featureCvtermpropDBTypes, false, featureCvtermpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}

	if err = featureCvtermprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureCvtermprop: %s", err)
	}

	count, err = FeatureCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
