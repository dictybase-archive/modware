package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeatureprops(t *testing.T) {
	t.Parallel()

	query := Featureprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeaturepropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureprop := &Featureprop{}
	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Featureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturepropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureprop := &Featureprop{}
	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Featureprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Featureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturepropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureprop := &Featureprop{}
	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeaturepropSlice{featureprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Featureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeaturepropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureprop := &Featureprop{}
	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeaturepropExists(tx, featureprop.FeaturepropID)
	if err != nil {
		t.Errorf("Unable to check if Featureprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeaturepropExistsG to return true, but got false.")
	}
}
func testFeaturepropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureprop := &Featureprop{}
	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	featurepropFound, err := FindFeatureprop(tx, featureprop.FeaturepropID)
	if err != nil {
		t.Error(err)
	}

	if featurepropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeaturepropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureprop := &Featureprop{}
	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Featureprops(tx).Bind(featureprop); err != nil {
		t.Error(err)
	}
}

func testFeaturepropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureprop := &Featureprop{}
	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Featureprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeaturepropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurepropOne := &Featureprop{}
	featurepropTwo := &Featureprop{}
	if err = randomize.Struct(seed, featurepropOne, featurepropDBTypes, false, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}
	if err = randomize.Struct(seed, featurepropTwo, featurepropDBTypes, false, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featurepropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Featureprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeaturepropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featurepropOne := &Featureprop{}
	featurepropTwo := &Featureprop{}
	if err = randomize.Struct(seed, featurepropOne, featurepropDBTypes, false, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}
	if err = randomize.Struct(seed, featurepropTwo, featurepropDBTypes, false, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featurepropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Featureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featurepropBeforeInsertHook(e boil.Executor, o *Featureprop) error {
	*o = Featureprop{}
	return nil
}

func featurepropAfterInsertHook(e boil.Executor, o *Featureprop) error {
	*o = Featureprop{}
	return nil
}

func featurepropAfterSelectHook(e boil.Executor, o *Featureprop) error {
	*o = Featureprop{}
	return nil
}

func featurepropBeforeUpdateHook(e boil.Executor, o *Featureprop) error {
	*o = Featureprop{}
	return nil
}

func featurepropAfterUpdateHook(e boil.Executor, o *Featureprop) error {
	*o = Featureprop{}
	return nil
}

func featurepropBeforeDeleteHook(e boil.Executor, o *Featureprop) error {
	*o = Featureprop{}
	return nil
}

func featurepropAfterDeleteHook(e boil.Executor, o *Featureprop) error {
	*o = Featureprop{}
	return nil
}

func featurepropBeforeUpsertHook(e boil.Executor, o *Featureprop) error {
	*o = Featureprop{}
	return nil
}

func featurepropAfterUpsertHook(e boil.Executor, o *Featureprop) error {
	*o = Featureprop{}
	return nil
}

func testFeaturepropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Featureprop{}
	o := &Featureprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featurepropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Featureprop object: %s", err)
	}

	AddFeaturepropHook(boil.BeforeInsertHook, featurepropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featurepropBeforeInsertHooks = []FeaturepropHook{}

	AddFeaturepropHook(boil.AfterInsertHook, featurepropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featurepropAfterInsertHooks = []FeaturepropHook{}

	AddFeaturepropHook(boil.AfterSelectHook, featurepropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featurepropAfterSelectHooks = []FeaturepropHook{}

	AddFeaturepropHook(boil.BeforeUpdateHook, featurepropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featurepropBeforeUpdateHooks = []FeaturepropHook{}

	AddFeaturepropHook(boil.AfterUpdateHook, featurepropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featurepropAfterUpdateHooks = []FeaturepropHook{}

	AddFeaturepropHook(boil.BeforeDeleteHook, featurepropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featurepropBeforeDeleteHooks = []FeaturepropHook{}

	AddFeaturepropHook(boil.AfterDeleteHook, featurepropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featurepropAfterDeleteHooks = []FeaturepropHook{}

	AddFeaturepropHook(boil.BeforeUpsertHook, featurepropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featurepropBeforeUpsertHooks = []FeaturepropHook{}

	AddFeaturepropHook(boil.AfterUpsertHook, featurepropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featurepropAfterUpsertHooks = []FeaturepropHook{}
}
func testFeaturepropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureprop := &Featureprop{}
	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Featureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeaturepropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureprop := &Featureprop{}
	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureprop.Insert(tx, featurepropColumns...); err != nil {
		t.Error(err)
	}

	count, err := Featureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeaturepropOneToOneFeaturepropPubUsingFeaturepropPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeaturepropPub
	var local Featureprop

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featurepropPubDBTypes, true, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featurepropDBTypes, true, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeaturepropID = local.FeaturepropID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeaturepropPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeaturepropID != foreign.FeaturepropID {
		t.Errorf("want: %v, got %v", foreign.FeaturepropID, check.FeaturepropID)
	}

	slice := FeaturepropSlice{&local}
	if err = local.L.LoadFeaturepropPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturepropPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeaturepropPub = nil
	if err = local.L.LoadFeaturepropPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturepropPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeaturepropOneToOneSetOpFeaturepropPubUsingFeaturepropPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Featureprop
	var b, c FeaturepropPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurepropDBTypes, false, strmangle.SetComplement(featurepropPrimaryKeyColumns, featurepropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featurepropPubDBTypes, false, strmangle.SetComplement(featurepropPubPrimaryKeyColumns, featurepropPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featurepropPubDBTypes, false, strmangle.SetComplement(featurepropPubPrimaryKeyColumns, featurepropPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeaturepropPub{&b, &c} {
		err = a.SetFeaturepropPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeaturepropPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Featureprop != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeaturepropID != x.FeaturepropID {
			t.Error("foreign key was wrong value", a.FeaturepropID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeaturepropID))
		reflect.Indirect(reflect.ValueOf(&x.FeaturepropID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeaturepropID != x.FeaturepropID {
			t.Error("foreign key was wrong value", a.FeaturepropID, x.FeaturepropID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testFeaturepropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Featureprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featurepropDBTypes, true, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
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

	slice := FeaturepropSlice{&local}
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

func testFeaturepropToOneFeatureUsingFeature(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Featureprop
	var foreign Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featurepropDBTypes, true, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.FeatureID = foreign.FeatureID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Feature(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureID != foreign.FeatureID {
		t.Errorf("want: %v, got %v", foreign.FeatureID, check.FeatureID)
	}

	slice := FeaturepropSlice{&local}
	if err = local.L.LoadFeature(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Feature == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Feature = nil
	if err = local.L.LoadFeature(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Feature == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeaturepropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Featureprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurepropDBTypes, false, strmangle.SetComplement(featurepropPrimaryKeyColumns, featurepropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeFeatureprop != &a {
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
func testFeaturepropToOneSetOpFeatureUsingFeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Featureprop
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurepropDBTypes, false, strmangle.SetComplement(featurepropPrimaryKeyColumns, featurepropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Feature{&b, &c} {
		err = a.SetFeature(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Feature != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Featureprop != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FeatureID))
		reflect.Indirect(reflect.ValueOf(&a.FeatureID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID, x.FeatureID)
		}
	}
}
func testFeaturepropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureprop := &Featureprop{}
	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturepropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureprop := &Featureprop{}
	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeaturepropSlice{featureprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeaturepropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureprop := &Featureprop{}
	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Featureprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featurepropDBTypes = map[string]string{"FeatureID": "integer", "FeaturepropID": "integer", "Rank": "integer", "TypeID": "integer", "Value": "text"}
	_                  = bytes.MinRead
)

func testFeaturepropsUpdate(t *testing.T) {
	t.Parallel()

	if len(featurepropColumns) == len(featurepropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureprop := &Featureprop{}
	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Featureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	if err = featureprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturepropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featurepropColumns) == len(featurepropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureprop := &Featureprop{}
	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Featureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureprop, featurepropDBTypes, true, featurepropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featurepropColumns, featurepropPrimaryKeyColumns) {
		fields = featurepropColumns
	} else {
		fields = strmangle.SetComplement(
			featurepropColumns,
			featurepropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featureprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeaturepropSlice{featureprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeaturepropsUpsert(t *testing.T) {
	t.Parallel()

	if len(featurepropColumns) == len(featurepropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featureprop := Featureprop{}
	if err = randomize.Struct(seed, &featureprop, featurepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Featureprop: %s", err)
	}

	count, err := Featureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featureprop, featurepropDBTypes, false, featurepropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	if err = featureprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Featureprop: %s", err)
	}

	count, err = Featureprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
