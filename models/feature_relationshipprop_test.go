package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeatureRelationshipprops(t *testing.T) {
	t.Parallel()

	query := FeatureRelationshipprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeatureRelationshippropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipprop := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureRelationshipprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureRelationshippropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipprop := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureRelationshipprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureRelationshippropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipprop := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureRelationshippropSlice{featureRelationshipprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeatureRelationshippropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipprop := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeatureRelationshippropExists(tx, featureRelationshipprop.FeatureRelationshippropID)
	if err != nil {
		t.Errorf("Unable to check if FeatureRelationshipprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeatureRelationshippropExistsG to return true, but got false.")
	}
}
func testFeatureRelationshippropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipprop := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipprop.Insert(tx); err != nil {
		t.Error(err)
	}

	featureRelationshippropFound, err := FindFeatureRelationshipprop(tx, featureRelationshipprop.FeatureRelationshippropID)
	if err != nil {
		t.Error(err)
	}

	if featureRelationshippropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeatureRelationshippropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipprop := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureRelationshipprops(tx).Bind(featureRelationshipprop); err != nil {
		t.Error(err)
	}
}

func testFeatureRelationshippropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipprop := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeatureRelationshipprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeatureRelationshippropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropOne := &FeatureRelationshipprop{}
	featureRelationshippropTwo := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshippropOne, featureRelationshippropDBTypes, false, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}
	if err = randomize.Struct(seed, featureRelationshippropTwo, featureRelationshippropDBTypes, false, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureRelationshippropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureRelationshipprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeatureRelationshippropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featureRelationshippropOne := &FeatureRelationshipprop{}
	featureRelationshippropTwo := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshippropOne, featureRelationshippropDBTypes, false, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}
	if err = randomize.Struct(seed, featureRelationshippropTwo, featureRelationshippropDBTypes, false, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureRelationshippropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featureRelationshippropBeforeInsertHook(e boil.Executor, o *FeatureRelationshipprop) error {
	*o = FeatureRelationshipprop{}
	return nil
}

func featureRelationshippropAfterInsertHook(e boil.Executor, o *FeatureRelationshipprop) error {
	*o = FeatureRelationshipprop{}
	return nil
}

func featureRelationshippropAfterSelectHook(e boil.Executor, o *FeatureRelationshipprop) error {
	*o = FeatureRelationshipprop{}
	return nil
}

func featureRelationshippropBeforeUpdateHook(e boil.Executor, o *FeatureRelationshipprop) error {
	*o = FeatureRelationshipprop{}
	return nil
}

func featureRelationshippropAfterUpdateHook(e boil.Executor, o *FeatureRelationshipprop) error {
	*o = FeatureRelationshipprop{}
	return nil
}

func featureRelationshippropBeforeDeleteHook(e boil.Executor, o *FeatureRelationshipprop) error {
	*o = FeatureRelationshipprop{}
	return nil
}

func featureRelationshippropAfterDeleteHook(e boil.Executor, o *FeatureRelationshipprop) error {
	*o = FeatureRelationshipprop{}
	return nil
}

func featureRelationshippropBeforeUpsertHook(e boil.Executor, o *FeatureRelationshipprop) error {
	*o = FeatureRelationshipprop{}
	return nil
}

func featureRelationshippropAfterUpsertHook(e boil.Executor, o *FeatureRelationshipprop) error {
	*o = FeatureRelationshipprop{}
	return nil
}

func testFeatureRelationshippropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeatureRelationshipprop{}
	o := &FeatureRelationshipprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featureRelationshippropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop object: %s", err)
	}

	AddFeatureRelationshippropHook(boil.BeforeInsertHook, featureRelationshippropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropBeforeInsertHooks = []FeatureRelationshippropHook{}

	AddFeatureRelationshippropHook(boil.AfterInsertHook, featureRelationshippropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropAfterInsertHooks = []FeatureRelationshippropHook{}

	AddFeatureRelationshippropHook(boil.AfterSelectHook, featureRelationshippropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropAfterSelectHooks = []FeatureRelationshippropHook{}

	AddFeatureRelationshippropHook(boil.BeforeUpdateHook, featureRelationshippropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropBeforeUpdateHooks = []FeatureRelationshippropHook{}

	AddFeatureRelationshippropHook(boil.AfterUpdateHook, featureRelationshippropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropAfterUpdateHooks = []FeatureRelationshippropHook{}

	AddFeatureRelationshippropHook(boil.BeforeDeleteHook, featureRelationshippropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropBeforeDeleteHooks = []FeatureRelationshippropHook{}

	AddFeatureRelationshippropHook(boil.AfterDeleteHook, featureRelationshippropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropAfterDeleteHooks = []FeatureRelationshippropHook{}

	AddFeatureRelationshippropHook(boil.BeforeUpsertHook, featureRelationshippropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropBeforeUpsertHooks = []FeatureRelationshippropHook{}

	AddFeatureRelationshippropHook(boil.AfterUpsertHook, featureRelationshippropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropAfterUpsertHooks = []FeatureRelationshippropHook{}
}
func testFeatureRelationshippropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipprop := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureRelationshippropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipprop := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipprop.Insert(tx, featureRelationshippropColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureRelationshippropOneToOneFeatureRelationshippropPubUsingFeatureRelationshippropPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureRelationshippropPub
	var local FeatureRelationshipprop

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureRelationshippropPubDBTypes, true, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureRelationshippropDBTypes, true, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeatureRelationshippropID = local.FeatureRelationshippropID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureRelationshippropPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureRelationshippropID != foreign.FeatureRelationshippropID {
		t.Errorf("want: %v, got %v", foreign.FeatureRelationshippropID, check.FeatureRelationshippropID)
	}

	slice := FeatureRelationshippropSlice{&local}
	if err = local.L.LoadFeatureRelationshippropPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureRelationshippropPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureRelationshippropPub = nil
	if err = local.L.LoadFeatureRelationshippropPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureRelationshippropPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureRelationshippropOneToOneSetOpFeatureRelationshippropPubUsingFeatureRelationshippropPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureRelationshipprop
	var b, c FeatureRelationshippropPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureRelationshippropDBTypes, false, strmangle.SetComplement(featureRelationshippropPrimaryKeyColumns, featureRelationshippropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureRelationshippropPubDBTypes, false, strmangle.SetComplement(featureRelationshippropPubPrimaryKeyColumns, featureRelationshippropPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureRelationshippropPubDBTypes, false, strmangle.SetComplement(featureRelationshippropPubPrimaryKeyColumns, featureRelationshippropPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureRelationshippropPub{&b, &c} {
		err = a.SetFeatureRelationshippropPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureRelationshippropPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.FeatureRelationshipprop != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureRelationshippropID != x.FeatureRelationshippropID {
			t.Error("foreign key was wrong value", a.FeatureRelationshippropID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeatureRelationshippropID))
		reflect.Indirect(reflect.ValueOf(&x.FeatureRelationshippropID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureRelationshippropID != x.FeatureRelationshippropID {
			t.Error("foreign key was wrong value", a.FeatureRelationshippropID, x.FeatureRelationshippropID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testFeatureRelationshippropToOneFeatureRelationshipUsingFeatureRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureRelationshipprop
	var foreign FeatureRelationship

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureRelationshippropDBTypes, true, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.FeatureRelationshipID = foreign.FeatureRelationshipID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureRelationshipID != foreign.FeatureRelationshipID {
		t.Errorf("want: %v, got %v", foreign.FeatureRelationshipID, check.FeatureRelationshipID)
	}

	slice := FeatureRelationshippropSlice{&local}
	if err = local.L.LoadFeatureRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureRelationship = nil
	if err = local.L.LoadFeatureRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureRelationshippropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureRelationshipprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureRelationshippropDBTypes, true, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
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

	slice := FeatureRelationshippropSlice{&local}
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

func testFeatureRelationshippropToOneSetOpFeatureRelationshipUsingFeatureRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureRelationshipprop
	var b, c FeatureRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureRelationshippropDBTypes, false, strmangle.SetComplement(featureRelationshippropPrimaryKeyColumns, featureRelationshippropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureRelationshipDBTypes, false, strmangle.SetComplement(featureRelationshipPrimaryKeyColumns, featureRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureRelationshipDBTypes, false, strmangle.SetComplement(featureRelationshipPrimaryKeyColumns, featureRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureRelationship{&b, &c} {
		err = a.SetFeatureRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureRelationship != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FeatureRelationshipprop != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FeatureRelationshipID != x.FeatureRelationshipID {
			t.Error("foreign key was wrong value", a.FeatureRelationshipID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FeatureRelationshipID))
		reflect.Indirect(reflect.ValueOf(&a.FeatureRelationshipID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureRelationshipID != x.FeatureRelationshipID {
			t.Error("foreign key was wrong value", a.FeatureRelationshipID, x.FeatureRelationshipID)
		}
	}
}
func testFeatureRelationshippropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureRelationshipprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureRelationshippropDBTypes, false, strmangle.SetComplement(featureRelationshippropPrimaryKeyColumns, featureRelationshippropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeFeatureRelationshipprop != &a {
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
func testFeatureRelationshippropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipprop := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureRelationshipprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureRelationshippropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipprop := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureRelationshippropSlice{featureRelationshipprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeatureRelationshippropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipprop := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureRelationshipprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featureRelationshippropDBTypes = map[string]string{"FeatureRelationshipID": "integer", "FeatureRelationshippropID": "integer", "Rank": "integer", "TypeID": "integer", "Value": "text"}
	_                              = bytes.MinRead
)

func testFeatureRelationshippropsUpdate(t *testing.T) {
	t.Parallel()

	if len(featureRelationshippropColumns) == len(featureRelationshippropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureRelationshipprop := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	if err = featureRelationshipprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureRelationshippropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featureRelationshippropColumns) == len(featureRelationshippropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureRelationshipprop := &FeatureRelationshipprop{}
	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureRelationshipprop, featureRelationshippropDBTypes, true, featureRelationshippropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featureRelationshippropColumns, featureRelationshippropPrimaryKeyColumns) {
		fields = featureRelationshippropColumns
	} else {
		fields = strmangle.SetComplement(
			featureRelationshippropColumns,
			featureRelationshippropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featureRelationshipprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeatureRelationshippropSlice{featureRelationshipprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeatureRelationshippropsUpsert(t *testing.T) {
	t.Parallel()

	if len(featureRelationshippropColumns) == len(featureRelationshippropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featureRelationshipprop := FeatureRelationshipprop{}
	if err = randomize.Struct(seed, &featureRelationshipprop, featureRelationshippropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureRelationshipprop: %s", err)
	}

	count, err := FeatureRelationshipprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featureRelationshipprop, featureRelationshippropDBTypes, false, featureRelationshippropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	if err = featureRelationshipprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureRelationshipprop: %s", err)
	}

	count, err = FeatureRelationshipprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
