package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeaturePubs(t *testing.T) {
	t.Parallel()

	query := FeaturePubs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeaturePubsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePub := &FeaturePub{}
	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featurePub.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturePubsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePub := &FeaturePub{}
	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeaturePubs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturePubsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePub := &FeaturePub{}
	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeaturePubSlice{featurePub}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeaturePubsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePub := &FeaturePub{}
	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePub.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeaturePubExists(tx, featurePub.FeaturePubID)
	if err != nil {
		t.Errorf("Unable to check if FeaturePub exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeaturePubExistsG to return true, but got false.")
	}
}
func testFeaturePubsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePub := &FeaturePub{}
	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePub.Insert(tx); err != nil {
		t.Error(err)
	}

	featurePubFound, err := FindFeaturePub(tx, featurePub.FeaturePubID)
	if err != nil {
		t.Error(err)
	}

	if featurePubFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeaturePubsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePub := &FeaturePub{}
	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeaturePubs(tx).Bind(featurePub); err != nil {
		t.Error(err)
	}
}

func testFeaturePubsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePub := &FeaturePub{}
	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePub.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeaturePubs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeaturePubsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePubOne := &FeaturePub{}
	featurePubTwo := &FeaturePub{}
	if err = randomize.Struct(seed, featurePubOne, featurePubDBTypes, false, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}
	if err = randomize.Struct(seed, featurePubTwo, featurePubDBTypes, false, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featurePubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeaturePubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeaturePubsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featurePubOne := &FeaturePub{}
	featurePubTwo := &FeaturePub{}
	if err = randomize.Struct(seed, featurePubOne, featurePubDBTypes, false, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}
	if err = randomize.Struct(seed, featurePubTwo, featurePubDBTypes, false, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featurePubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featurePubBeforeInsertHook(e boil.Executor, o *FeaturePub) error {
	*o = FeaturePub{}
	return nil
}

func featurePubAfterInsertHook(e boil.Executor, o *FeaturePub) error {
	*o = FeaturePub{}
	return nil
}

func featurePubAfterSelectHook(e boil.Executor, o *FeaturePub) error {
	*o = FeaturePub{}
	return nil
}

func featurePubBeforeUpdateHook(e boil.Executor, o *FeaturePub) error {
	*o = FeaturePub{}
	return nil
}

func featurePubAfterUpdateHook(e boil.Executor, o *FeaturePub) error {
	*o = FeaturePub{}
	return nil
}

func featurePubBeforeDeleteHook(e boil.Executor, o *FeaturePub) error {
	*o = FeaturePub{}
	return nil
}

func featurePubAfterDeleteHook(e boil.Executor, o *FeaturePub) error {
	*o = FeaturePub{}
	return nil
}

func featurePubBeforeUpsertHook(e boil.Executor, o *FeaturePub) error {
	*o = FeaturePub{}
	return nil
}

func featurePubAfterUpsertHook(e boil.Executor, o *FeaturePub) error {
	*o = FeaturePub{}
	return nil
}

func testFeaturePubsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeaturePub{}
	o := &FeaturePub{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featurePubDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeaturePub object: %s", err)
	}

	AddFeaturePubHook(boil.BeforeInsertHook, featurePubBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featurePubBeforeInsertHooks = []FeaturePubHook{}

	AddFeaturePubHook(boil.AfterInsertHook, featurePubAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featurePubAfterInsertHooks = []FeaturePubHook{}

	AddFeaturePubHook(boil.AfterSelectHook, featurePubAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featurePubAfterSelectHooks = []FeaturePubHook{}

	AddFeaturePubHook(boil.BeforeUpdateHook, featurePubBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featurePubBeforeUpdateHooks = []FeaturePubHook{}

	AddFeaturePubHook(boil.AfterUpdateHook, featurePubAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featurePubAfterUpdateHooks = []FeaturePubHook{}

	AddFeaturePubHook(boil.BeforeDeleteHook, featurePubBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featurePubBeforeDeleteHooks = []FeaturePubHook{}

	AddFeaturePubHook(boil.AfterDeleteHook, featurePubAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featurePubAfterDeleteHooks = []FeaturePubHook{}

	AddFeaturePubHook(boil.BeforeUpsertHook, featurePubBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featurePubBeforeUpsertHooks = []FeaturePubHook{}

	AddFeaturePubHook(boil.AfterUpsertHook, featurePubAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featurePubAfterUpsertHooks = []FeaturePubHook{}
}
func testFeaturePubsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePub := &FeaturePub{}
	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeaturePubsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePub := &FeaturePub{}
	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePub.Insert(tx, featurePubColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeaturePubOneToOneFeaturePubpropUsingFeaturePubprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeaturePubprop
	var local FeaturePub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featurePubpropDBTypes, true, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featurePubDBTypes, true, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeaturePubID = local.FeaturePubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeaturePubprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeaturePubID != foreign.FeaturePubID {
		t.Errorf("want: %v, got %v", foreign.FeaturePubID, check.FeaturePubID)
	}

	slice := FeaturePubSlice{&local}
	if err = local.L.LoadFeaturePubprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturePubprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeaturePubprop = nil
	if err = local.L.LoadFeaturePubprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturePubprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeaturePubOneToOneSetOpFeaturePubpropUsingFeaturePubprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeaturePub
	var b, c FeaturePubprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurePubDBTypes, false, strmangle.SetComplement(featurePubPrimaryKeyColumns, featurePubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featurePubpropDBTypes, false, strmangle.SetComplement(featurePubpropPrimaryKeyColumns, featurePubpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featurePubpropDBTypes, false, strmangle.SetComplement(featurePubpropPrimaryKeyColumns, featurePubpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeaturePubprop{&b, &c} {
		err = a.SetFeaturePubprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeaturePubprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.FeaturePub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeaturePubID != x.FeaturePubID {
			t.Error("foreign key was wrong value", a.FeaturePubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeaturePubID))
		reflect.Indirect(reflect.ValueOf(&x.FeaturePubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeaturePubID != x.FeaturePubID {
			t.Error("foreign key was wrong value", a.FeaturePubID, x.FeaturePubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testFeaturePubToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeaturePub
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featurePubDBTypes, true, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.PubID = foreign.PubID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Pub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := FeaturePubSlice{&local}
	if err = local.L.LoadPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Pub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Pub = nil
	if err = local.L.LoadPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Pub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeaturePubToOneFeatureUsingFeature(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeaturePub
	var foreign Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featurePubDBTypes, true, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
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

	slice := FeaturePubSlice{&local}
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

func testFeaturePubToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeaturePub
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurePubDBTypes, false, strmangle.SetComplement(featurePubPrimaryKeyColumns, featurePubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Pub{&b, &c} {
		err = a.SetPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Pub != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FeaturePub != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PubID))
		reflect.Indirect(reflect.ValueOf(&a.PubID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}
	}
}
func testFeaturePubToOneSetOpFeatureUsingFeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeaturePub
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurePubDBTypes, false, strmangle.SetComplement(featurePubPrimaryKeyColumns, featurePubColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeaturePub != &a {
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
func testFeaturePubsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePub := &FeaturePub{}
	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featurePub.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturePubsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePub := &FeaturePub{}
	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeaturePubSlice{featurePub}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeaturePubsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePub := &FeaturePub{}
	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeaturePubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featurePubDBTypes = map[string]string{"FeatureID": "integer", "FeaturePubID": "integer", "PubID": "integer"}
	_                 = bytes.MinRead
)

func testFeaturePubsUpdate(t *testing.T) {
	t.Parallel()

	if len(featurePubColumns) == len(featurePubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featurePub := &FeaturePub{}
	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	if err = featurePub.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturePubsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featurePubColumns) == len(featurePubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featurePub := &FeaturePub{}
	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featurePub, featurePubDBTypes, true, featurePubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featurePubColumns, featurePubPrimaryKeyColumns) {
		fields = featurePubColumns
	} else {
		fields = strmangle.SetComplement(
			featurePubColumns,
			featurePubPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featurePub))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeaturePubSlice{featurePub}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeaturePubsUpsert(t *testing.T) {
	t.Parallel()

	if len(featurePubColumns) == len(featurePubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featurePub := FeaturePub{}
	if err = randomize.Struct(seed, &featurePub, featurePubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePub.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeaturePub: %s", err)
	}

	count, err := FeaturePubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featurePub, featurePubDBTypes, false, featurePubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}

	if err = featurePub.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeaturePub: %s", err)
	}

	count, err = FeaturePubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
