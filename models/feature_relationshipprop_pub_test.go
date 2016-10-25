package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeatureRelationshippropPubs(t *testing.T) {
	t.Parallel()

	query := FeatureRelationshippropPubs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeatureRelationshippropPubsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropPub := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureRelationshippropPub.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshippropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureRelationshippropPubsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropPub := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureRelationshippropPubs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshippropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureRelationshippropPubsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropPub := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureRelationshippropPubSlice{featureRelationshippropPub}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshippropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeatureRelationshippropPubsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropPub := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeatureRelationshippropPubExists(tx, featureRelationshippropPub.FeatureRelationshippropPubID)
	if err != nil {
		t.Errorf("Unable to check if FeatureRelationshippropPub exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeatureRelationshippropPubExistsG to return true, but got false.")
	}
}
func testFeatureRelationshippropPubsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropPub := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	featureRelationshippropPubFound, err := FindFeatureRelationshippropPub(tx, featureRelationshippropPub.FeatureRelationshippropPubID)
	if err != nil {
		t.Error(err)
	}

	if featureRelationshippropPubFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeatureRelationshippropPubsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropPub := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureRelationshippropPubs(tx).Bind(featureRelationshippropPub); err != nil {
		t.Error(err)
	}
}

func testFeatureRelationshippropPubsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropPub := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeatureRelationshippropPubs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeatureRelationshippropPubsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropPubOne := &FeatureRelationshippropPub{}
	featureRelationshippropPubTwo := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPubOne, featureRelationshippropPubDBTypes, false, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}
	if err = randomize.Struct(seed, featureRelationshippropPubTwo, featureRelationshippropPubDBTypes, false, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureRelationshippropPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureRelationshippropPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeatureRelationshippropPubsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featureRelationshippropPubOne := &FeatureRelationshippropPub{}
	featureRelationshippropPubTwo := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPubOne, featureRelationshippropPubDBTypes, false, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}
	if err = randomize.Struct(seed, featureRelationshippropPubTwo, featureRelationshippropPubDBTypes, false, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureRelationshippropPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshippropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featureRelationshippropPubBeforeInsertHook(e boil.Executor, o *FeatureRelationshippropPub) error {
	*o = FeatureRelationshippropPub{}
	return nil
}

func featureRelationshippropPubAfterInsertHook(e boil.Executor, o *FeatureRelationshippropPub) error {
	*o = FeatureRelationshippropPub{}
	return nil
}

func featureRelationshippropPubAfterSelectHook(e boil.Executor, o *FeatureRelationshippropPub) error {
	*o = FeatureRelationshippropPub{}
	return nil
}

func featureRelationshippropPubBeforeUpdateHook(e boil.Executor, o *FeatureRelationshippropPub) error {
	*o = FeatureRelationshippropPub{}
	return nil
}

func featureRelationshippropPubAfterUpdateHook(e boil.Executor, o *FeatureRelationshippropPub) error {
	*o = FeatureRelationshippropPub{}
	return nil
}

func featureRelationshippropPubBeforeDeleteHook(e boil.Executor, o *FeatureRelationshippropPub) error {
	*o = FeatureRelationshippropPub{}
	return nil
}

func featureRelationshippropPubAfterDeleteHook(e boil.Executor, o *FeatureRelationshippropPub) error {
	*o = FeatureRelationshippropPub{}
	return nil
}

func featureRelationshippropPubBeforeUpsertHook(e boil.Executor, o *FeatureRelationshippropPub) error {
	*o = FeatureRelationshippropPub{}
	return nil
}

func featureRelationshippropPubAfterUpsertHook(e boil.Executor, o *FeatureRelationshippropPub) error {
	*o = FeatureRelationshippropPub{}
	return nil
}

func testFeatureRelationshippropPubsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeatureRelationshippropPub{}
	o := &FeatureRelationshippropPub{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featureRelationshippropPubDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub object: %s", err)
	}

	AddFeatureRelationshippropPubHook(boil.BeforeInsertHook, featureRelationshippropPubBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropPubBeforeInsertHooks = []FeatureRelationshippropPubHook{}

	AddFeatureRelationshippropPubHook(boil.AfterInsertHook, featureRelationshippropPubAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropPubAfterInsertHooks = []FeatureRelationshippropPubHook{}

	AddFeatureRelationshippropPubHook(boil.AfterSelectHook, featureRelationshippropPubAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropPubAfterSelectHooks = []FeatureRelationshippropPubHook{}

	AddFeatureRelationshippropPubHook(boil.BeforeUpdateHook, featureRelationshippropPubBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropPubBeforeUpdateHooks = []FeatureRelationshippropPubHook{}

	AddFeatureRelationshippropPubHook(boil.AfterUpdateHook, featureRelationshippropPubAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropPubAfterUpdateHooks = []FeatureRelationshippropPubHook{}

	AddFeatureRelationshippropPubHook(boil.BeforeDeleteHook, featureRelationshippropPubBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropPubBeforeDeleteHooks = []FeatureRelationshippropPubHook{}

	AddFeatureRelationshippropPubHook(boil.AfterDeleteHook, featureRelationshippropPubAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropPubAfterDeleteHooks = []FeatureRelationshippropPubHook{}

	AddFeatureRelationshippropPubHook(boil.BeforeUpsertHook, featureRelationshippropPubBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropPubBeforeUpsertHooks = []FeatureRelationshippropPubHook{}

	AddFeatureRelationshippropPubHook(boil.AfterUpsertHook, featureRelationshippropPubAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshippropPubAfterUpsertHooks = []FeatureRelationshippropPubHook{}
}
func testFeatureRelationshippropPubsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropPub := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshippropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureRelationshippropPubsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropPub := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPub.Insert(tx, featureRelationshippropPubColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshippropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureRelationshippropPubToOneFeatureRelationshippropUsingFeatureRelationshipprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureRelationshippropPub
	var foreign FeatureRelationshipprop

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureRelationshippropPubDBTypes, true, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, featureRelationshippropDBTypes, true, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.FeatureRelationshippropID = foreign.FeatureRelationshippropID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureRelationshipprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureRelationshippropID != foreign.FeatureRelationshippropID {
		t.Errorf("want: %v, got %v", foreign.FeatureRelationshippropID, check.FeatureRelationshippropID)
	}

	slice := FeatureRelationshippropPubSlice{&local}
	if err = local.L.LoadFeatureRelationshipprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureRelationshipprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureRelationshipprop = nil
	if err = local.L.LoadFeatureRelationshipprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureRelationshipprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureRelationshippropPubToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureRelationshippropPub
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureRelationshippropPubDBTypes, true, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
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

	slice := FeatureRelationshippropPubSlice{&local}
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

func testFeatureRelationshippropPubToOneSetOpFeatureRelationshippropUsingFeatureRelationshipprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureRelationshippropPub
	var b, c FeatureRelationshipprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureRelationshippropPubDBTypes, false, strmangle.SetComplement(featureRelationshippropPubPrimaryKeyColumns, featureRelationshippropPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureRelationshippropDBTypes, false, strmangle.SetComplement(featureRelationshippropPrimaryKeyColumns, featureRelationshippropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureRelationshippropDBTypes, false, strmangle.SetComplement(featureRelationshippropPrimaryKeyColumns, featureRelationshippropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureRelationshipprop{&b, &c} {
		err = a.SetFeatureRelationshipprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureRelationshipprop != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FeatureRelationshippropPub != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FeatureRelationshippropID != x.FeatureRelationshippropID {
			t.Error("foreign key was wrong value", a.FeatureRelationshippropID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FeatureRelationshippropID))
		reflect.Indirect(reflect.ValueOf(&a.FeatureRelationshippropID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureRelationshippropID != x.FeatureRelationshippropID {
			t.Error("foreign key was wrong value", a.FeatureRelationshippropID, x.FeatureRelationshippropID)
		}
	}
}
func testFeatureRelationshippropPubToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureRelationshippropPub
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureRelationshippropPubDBTypes, false, strmangle.SetComplement(featureRelationshippropPubPrimaryKeyColumns, featureRelationshippropPubColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeatureRelationshippropPub != &a {
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
func testFeatureRelationshippropPubsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropPub := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureRelationshippropPub.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureRelationshippropPubsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropPub := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureRelationshippropPubSlice{featureRelationshippropPub}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeatureRelationshippropPubsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropPub := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureRelationshippropPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featureRelationshippropPubDBTypes = map[string]string{"FeatureRelationshippropID": "integer", "FeatureRelationshippropPubID": "integer", "PubID": "integer"}
	_                                 = bytes.MinRead
)

func testFeatureRelationshippropPubsUpdate(t *testing.T) {
	t.Parallel()

	if len(featureRelationshippropPubColumns) == len(featureRelationshippropPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropPub := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshippropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	if err = featureRelationshippropPub.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureRelationshippropPubsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featureRelationshippropPubColumns) == len(featureRelationshippropPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureRelationshippropPub := &FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshippropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureRelationshippropPub, featureRelationshippropPubDBTypes, true, featureRelationshippropPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featureRelationshippropPubColumns, featureRelationshippropPubPrimaryKeyColumns) {
		fields = featureRelationshippropPubColumns
	} else {
		fields = strmangle.SetComplement(
			featureRelationshippropPubColumns,
			featureRelationshippropPubPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featureRelationshippropPub))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeatureRelationshippropPubSlice{featureRelationshippropPub}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeatureRelationshippropPubsUpsert(t *testing.T) {
	t.Parallel()

	if len(featureRelationshippropPubColumns) == len(featureRelationshippropPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featureRelationshippropPub := FeatureRelationshippropPub{}
	if err = randomize.Struct(seed, &featureRelationshippropPub, featureRelationshippropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshippropPub.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureRelationshippropPub: %s", err)
	}

	count, err := FeatureRelationshippropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featureRelationshippropPub, featureRelationshippropPubDBTypes, false, featureRelationshippropPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}

	if err = featureRelationshippropPub.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureRelationshippropPub: %s", err)
	}

	count, err = FeatureRelationshippropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
