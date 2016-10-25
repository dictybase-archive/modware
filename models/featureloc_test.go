package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeaturelocs(t *testing.T) {
	t.Parallel()

	query := Featurelocs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeaturelocsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureloc := &Featureloc{}
	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureloc.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureloc.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Featurelocs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturelocsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureloc := &Featureloc{}
	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureloc.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Featurelocs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Featurelocs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturelocsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureloc := &Featureloc{}
	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureloc.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeaturelocSlice{featureloc}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Featurelocs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeaturelocsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureloc := &Featureloc{}
	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureloc.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeaturelocExists(tx, featureloc.FeaturelocID)
	if err != nil {
		t.Errorf("Unable to check if Featureloc exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeaturelocExistsG to return true, but got false.")
	}
}
func testFeaturelocsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureloc := &Featureloc{}
	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureloc.Insert(tx); err != nil {
		t.Error(err)
	}

	featurelocFound, err := FindFeatureloc(tx, featureloc.FeaturelocID)
	if err != nil {
		t.Error(err)
	}

	if featurelocFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeaturelocsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureloc := &Featureloc{}
	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureloc.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Featurelocs(tx).Bind(featureloc); err != nil {
		t.Error(err)
	}
}

func testFeaturelocsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureloc := &Featureloc{}
	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureloc.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Featurelocs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeaturelocsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurelocOne := &Featureloc{}
	featurelocTwo := &Featureloc{}
	if err = randomize.Struct(seed, featurelocOne, featurelocDBTypes, false, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}
	if err = randomize.Struct(seed, featurelocTwo, featurelocDBTypes, false, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featurelocTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Featurelocs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeaturelocsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featurelocOne := &Featureloc{}
	featurelocTwo := &Featureloc{}
	if err = randomize.Struct(seed, featurelocOne, featurelocDBTypes, false, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}
	if err = randomize.Struct(seed, featurelocTwo, featurelocDBTypes, false, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featurelocTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Featurelocs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featurelocBeforeInsertHook(e boil.Executor, o *Featureloc) error {
	*o = Featureloc{}
	return nil
}

func featurelocAfterInsertHook(e boil.Executor, o *Featureloc) error {
	*o = Featureloc{}
	return nil
}

func featurelocAfterSelectHook(e boil.Executor, o *Featureloc) error {
	*o = Featureloc{}
	return nil
}

func featurelocBeforeUpdateHook(e boil.Executor, o *Featureloc) error {
	*o = Featureloc{}
	return nil
}

func featurelocAfterUpdateHook(e boil.Executor, o *Featureloc) error {
	*o = Featureloc{}
	return nil
}

func featurelocBeforeDeleteHook(e boil.Executor, o *Featureloc) error {
	*o = Featureloc{}
	return nil
}

func featurelocAfterDeleteHook(e boil.Executor, o *Featureloc) error {
	*o = Featureloc{}
	return nil
}

func featurelocBeforeUpsertHook(e boil.Executor, o *Featureloc) error {
	*o = Featureloc{}
	return nil
}

func featurelocAfterUpsertHook(e boil.Executor, o *Featureloc) error {
	*o = Featureloc{}
	return nil
}

func testFeaturelocsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Featureloc{}
	o := &Featureloc{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featurelocDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Featureloc object: %s", err)
	}

	AddFeaturelocHook(boil.BeforeInsertHook, featurelocBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featurelocBeforeInsertHooks = []FeaturelocHook{}

	AddFeaturelocHook(boil.AfterInsertHook, featurelocAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featurelocAfterInsertHooks = []FeaturelocHook{}

	AddFeaturelocHook(boil.AfterSelectHook, featurelocAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featurelocAfterSelectHooks = []FeaturelocHook{}

	AddFeaturelocHook(boil.BeforeUpdateHook, featurelocBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featurelocBeforeUpdateHooks = []FeaturelocHook{}

	AddFeaturelocHook(boil.AfterUpdateHook, featurelocAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featurelocAfterUpdateHooks = []FeaturelocHook{}

	AddFeaturelocHook(boil.BeforeDeleteHook, featurelocBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featurelocBeforeDeleteHooks = []FeaturelocHook{}

	AddFeaturelocHook(boil.AfterDeleteHook, featurelocAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featurelocAfterDeleteHooks = []FeaturelocHook{}

	AddFeaturelocHook(boil.BeforeUpsertHook, featurelocBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featurelocBeforeUpsertHooks = []FeaturelocHook{}

	AddFeaturelocHook(boil.AfterUpsertHook, featurelocAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featurelocAfterUpsertHooks = []FeaturelocHook{}
}
func testFeaturelocsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureloc := &Featureloc{}
	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureloc.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Featurelocs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeaturelocsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureloc := &Featureloc{}
	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureloc.Insert(tx, featurelocColumns...); err != nil {
		t.Error(err)
	}

	count, err := Featurelocs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeaturelocOneToOneFeaturelocPubUsingFeaturelocPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeaturelocPub
	var local Featureloc

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featurelocPubDBTypes, true, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featurelocDBTypes, true, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeaturelocID = local.FeaturelocID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeaturelocPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeaturelocID != foreign.FeaturelocID {
		t.Errorf("want: %v, got %v", foreign.FeaturelocID, check.FeaturelocID)
	}

	slice := FeaturelocSlice{&local}
	if err = local.L.LoadFeaturelocPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturelocPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeaturelocPub = nil
	if err = local.L.LoadFeaturelocPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturelocPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeaturelocOneToOneSetOpFeaturelocPubUsingFeaturelocPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Featureloc
	var b, c FeaturelocPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurelocDBTypes, false, strmangle.SetComplement(featurelocPrimaryKeyColumns, featurelocColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featurelocPubDBTypes, false, strmangle.SetComplement(featurelocPubPrimaryKeyColumns, featurelocPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featurelocPubDBTypes, false, strmangle.SetComplement(featurelocPubPrimaryKeyColumns, featurelocPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeaturelocPub{&b, &c} {
		err = a.SetFeaturelocPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeaturelocPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Featureloc != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeaturelocID != x.FeaturelocID {
			t.Error("foreign key was wrong value", a.FeaturelocID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeaturelocID))
		reflect.Indirect(reflect.ValueOf(&x.FeaturelocID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeaturelocID != x.FeaturelocID {
			t.Error("foreign key was wrong value", a.FeaturelocID, x.FeaturelocID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testFeaturelocToOneFeatureUsingFeature(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Featureloc
	var foreign Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featurelocDBTypes, true, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
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

	slice := FeaturelocSlice{&local}
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

func testFeaturelocToOneFeatureUsingSrcfeature(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Featureloc
	var foreign Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featurelocDBTypes, true, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	local.SrcfeatureID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.SrcfeatureID.Int = foreign.FeatureID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Srcfeature(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureID != foreign.FeatureID {
		t.Errorf("want: %v, got %v", foreign.FeatureID, check.FeatureID)
	}

	slice := FeaturelocSlice{&local}
	if err = local.L.LoadSrcfeature(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Srcfeature == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Srcfeature = nil
	if err = local.L.LoadSrcfeature(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Srcfeature == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeaturelocToOneSetOpFeatureUsingFeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Featureloc
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurelocDBTypes, false, strmangle.SetComplement(featurelocPrimaryKeyColumns, featurelocColumnsWithoutDefault)...); err != nil {
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

		if x.R.Featureloc != &a {
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
func testFeaturelocToOneSetOpFeatureUsingSrcfeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Featureloc
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurelocDBTypes, false, strmangle.SetComplement(featurelocPrimaryKeyColumns, featurelocColumnsWithoutDefault)...); err != nil {
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
		err = a.SetSrcfeature(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Srcfeature != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SrcfeatureFeaturelocs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SrcfeatureID.Int != x.FeatureID {
			t.Error("foreign key was wrong value", a.SrcfeatureID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SrcfeatureID.Int))
		reflect.Indirect(reflect.ValueOf(&a.SrcfeatureID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SrcfeatureID.Int != x.FeatureID {
			t.Error("foreign key was wrong value", a.SrcfeatureID.Int, x.FeatureID)
		}
	}
}

func testFeaturelocToOneRemoveOpFeatureUsingSrcfeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Featureloc
	var b Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurelocDBTypes, false, strmangle.SetComplement(featurelocPrimaryKeyColumns, featurelocColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetSrcfeature(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveSrcfeature(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Srcfeature(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Srcfeature != nil {
		t.Error("R struct entry should be nil")
	}

	if a.SrcfeatureID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.SrcfeatureFeaturelocs) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testFeaturelocsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureloc := &Featureloc{}
	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureloc.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureloc.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturelocsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureloc := &Featureloc{}
	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureloc.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeaturelocSlice{featureloc}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeaturelocsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureloc := &Featureloc{}
	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureloc.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Featurelocs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featurelocDBTypes = map[string]string{"FeatureID": "integer", "FeaturelocID": "integer", "Fmax": "integer", "Fmin": "integer", "IsFmaxPartial": "boolean", "IsFminPartial": "boolean", "Locgroup": "integer", "Phase": "integer", "Rank": "integer", "ResidueInfo": "text", "SrcfeatureID": "integer", "Strand": "smallint"}
	_                 = bytes.MinRead
)

func testFeaturelocsUpdate(t *testing.T) {
	t.Parallel()

	if len(featurelocColumns) == len(featurelocPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureloc := &Featureloc{}
	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureloc.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Featurelocs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	if err = featureloc.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturelocsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featurelocColumns) == len(featurelocPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureloc := &Featureloc{}
	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureloc.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Featurelocs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureloc, featurelocDBTypes, true, featurelocPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featurelocColumns, featurelocPrimaryKeyColumns) {
		fields = featurelocColumns
	} else {
		fields = strmangle.SetComplement(
			featurelocColumns,
			featurelocPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featureloc))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeaturelocSlice{featureloc}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeaturelocsUpsert(t *testing.T) {
	t.Parallel()

	if len(featurelocColumns) == len(featurelocPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featureloc := Featureloc{}
	if err = randomize.Struct(seed, &featureloc, featurelocDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureloc.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Featureloc: %s", err)
	}

	count, err := Featurelocs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featureloc, featurelocDBTypes, false, featurelocPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	if err = featureloc.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Featureloc: %s", err)
	}

	count, err = Featurelocs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
