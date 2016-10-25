package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeaturepropPubs(t *testing.T) {
	t.Parallel()

	query := FeaturepropPubs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeaturepropPubsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurepropPub := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featurepropPub.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturepropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturepropPubsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurepropPub := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeaturepropPubs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeaturepropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturepropPubsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurepropPub := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeaturepropPubSlice{featurepropPub}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturepropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeaturepropPubsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurepropPub := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeaturepropPubExists(tx, featurepropPub.FeaturepropPubID)
	if err != nil {
		t.Errorf("Unable to check if FeaturepropPub exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeaturepropPubExistsG to return true, but got false.")
	}
}
func testFeaturepropPubsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurepropPub := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	featurepropPubFound, err := FindFeaturepropPub(tx, featurepropPub.FeaturepropPubID)
	if err != nil {
		t.Error(err)
	}

	if featurepropPubFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeaturepropPubsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurepropPub := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeaturepropPubs(tx).Bind(featurepropPub); err != nil {
		t.Error(err)
	}
}

func testFeaturepropPubsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurepropPub := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeaturepropPubs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeaturepropPubsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurepropPubOne := &FeaturepropPub{}
	featurepropPubTwo := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPubOne, featurepropPubDBTypes, false, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}
	if err = randomize.Struct(seed, featurepropPubTwo, featurepropPubDBTypes, false, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featurepropPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeaturepropPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeaturepropPubsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featurepropPubOne := &FeaturepropPub{}
	featurepropPubTwo := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPubOne, featurepropPubDBTypes, false, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}
	if err = randomize.Struct(seed, featurepropPubTwo, featurepropPubDBTypes, false, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featurepropPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturepropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featurepropPubBeforeInsertHook(e boil.Executor, o *FeaturepropPub) error {
	*o = FeaturepropPub{}
	return nil
}

func featurepropPubAfterInsertHook(e boil.Executor, o *FeaturepropPub) error {
	*o = FeaturepropPub{}
	return nil
}

func featurepropPubAfterSelectHook(e boil.Executor, o *FeaturepropPub) error {
	*o = FeaturepropPub{}
	return nil
}

func featurepropPubBeforeUpdateHook(e boil.Executor, o *FeaturepropPub) error {
	*o = FeaturepropPub{}
	return nil
}

func featurepropPubAfterUpdateHook(e boil.Executor, o *FeaturepropPub) error {
	*o = FeaturepropPub{}
	return nil
}

func featurepropPubBeforeDeleteHook(e boil.Executor, o *FeaturepropPub) error {
	*o = FeaturepropPub{}
	return nil
}

func featurepropPubAfterDeleteHook(e boil.Executor, o *FeaturepropPub) error {
	*o = FeaturepropPub{}
	return nil
}

func featurepropPubBeforeUpsertHook(e boil.Executor, o *FeaturepropPub) error {
	*o = FeaturepropPub{}
	return nil
}

func featurepropPubAfterUpsertHook(e boil.Executor, o *FeaturepropPub) error {
	*o = FeaturepropPub{}
	return nil
}

func testFeaturepropPubsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeaturepropPub{}
	o := &FeaturepropPub{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featurepropPubDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub object: %s", err)
	}

	AddFeaturepropPubHook(boil.BeforeInsertHook, featurepropPubBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featurepropPubBeforeInsertHooks = []FeaturepropPubHook{}

	AddFeaturepropPubHook(boil.AfterInsertHook, featurepropPubAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featurepropPubAfterInsertHooks = []FeaturepropPubHook{}

	AddFeaturepropPubHook(boil.AfterSelectHook, featurepropPubAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featurepropPubAfterSelectHooks = []FeaturepropPubHook{}

	AddFeaturepropPubHook(boil.BeforeUpdateHook, featurepropPubBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featurepropPubBeforeUpdateHooks = []FeaturepropPubHook{}

	AddFeaturepropPubHook(boil.AfterUpdateHook, featurepropPubAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featurepropPubAfterUpdateHooks = []FeaturepropPubHook{}

	AddFeaturepropPubHook(boil.BeforeDeleteHook, featurepropPubBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featurepropPubBeforeDeleteHooks = []FeaturepropPubHook{}

	AddFeaturepropPubHook(boil.AfterDeleteHook, featurepropPubAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featurepropPubAfterDeleteHooks = []FeaturepropPubHook{}

	AddFeaturepropPubHook(boil.BeforeUpsertHook, featurepropPubBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featurepropPubBeforeUpsertHooks = []FeaturepropPubHook{}

	AddFeaturepropPubHook(boil.AfterUpsertHook, featurepropPubAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featurepropPubAfterUpsertHooks = []FeaturepropPubHook{}
}
func testFeaturepropPubsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurepropPub := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturepropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeaturepropPubsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurepropPub := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPub.Insert(tx, featurepropPubColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeaturepropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeaturepropPubToOneFeaturepropUsingFeatureprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeaturepropPub
	var foreign Featureprop

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featurepropPubDBTypes, true, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, featurepropDBTypes, true, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.FeaturepropID = foreign.FeaturepropID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Featureprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeaturepropID != foreign.FeaturepropID {
		t.Errorf("want: %v, got %v", foreign.FeaturepropID, check.FeaturepropID)
	}

	slice := FeaturepropPubSlice{&local}
	if err = local.L.LoadFeatureprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Featureprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Featureprop = nil
	if err = local.L.LoadFeatureprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Featureprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeaturepropPubToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeaturepropPub
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featurepropPubDBTypes, true, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
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

	slice := FeaturepropPubSlice{&local}
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

func testFeaturepropPubToOneSetOpFeaturepropUsingFeatureprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeaturepropPub
	var b, c Featureprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurepropPubDBTypes, false, strmangle.SetComplement(featurepropPubPrimaryKeyColumns, featurepropPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featurepropDBTypes, false, strmangle.SetComplement(featurepropPrimaryKeyColumns, featurepropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featurepropDBTypes, false, strmangle.SetComplement(featurepropPrimaryKeyColumns, featurepropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Featureprop{&b, &c} {
		err = a.SetFeatureprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Featureprop != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FeaturepropPub != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FeaturepropID != x.FeaturepropID {
			t.Error("foreign key was wrong value", a.FeaturepropID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FeaturepropID))
		reflect.Indirect(reflect.ValueOf(&a.FeaturepropID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeaturepropID != x.FeaturepropID {
			t.Error("foreign key was wrong value", a.FeaturepropID, x.FeaturepropID)
		}
	}
}
func testFeaturepropPubToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeaturepropPub
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurepropPubDBTypes, false, strmangle.SetComplement(featurepropPubPrimaryKeyColumns, featurepropPubColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeaturepropPub != &a {
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
func testFeaturepropPubsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurepropPub := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featurepropPub.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturepropPubsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurepropPub := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeaturepropPubSlice{featurepropPub}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeaturepropPubsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurepropPub := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeaturepropPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featurepropPubDBTypes = map[string]string{"FeaturepropID": "integer", "FeaturepropPubID": "integer", "PubID": "integer"}
	_                     = bytes.MinRead
)

func testFeaturepropPubsUpdate(t *testing.T) {
	t.Parallel()

	if len(featurepropPubColumns) == len(featurepropPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featurepropPub := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturepropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	if err = featurepropPub.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturepropPubsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featurepropPubColumns) == len(featurepropPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featurepropPub := &FeaturepropPub{}
	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturepropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featurepropPub, featurepropPubDBTypes, true, featurepropPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featurepropPubColumns, featurepropPubPrimaryKeyColumns) {
		fields = featurepropPubColumns
	} else {
		fields = strmangle.SetComplement(
			featurepropPubColumns,
			featurepropPubPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featurepropPub))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeaturepropPubSlice{featurepropPub}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeaturepropPubsUpsert(t *testing.T) {
	t.Parallel()

	if len(featurepropPubColumns) == len(featurepropPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featurepropPub := FeaturepropPub{}
	if err = randomize.Struct(seed, &featurepropPub, featurepropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurepropPub.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeaturepropPub: %s", err)
	}

	count, err := FeaturepropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featurepropPub, featurepropPubDBTypes, false, featurepropPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}

	if err = featurepropPub.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeaturepropPub: %s", err)
	}

	count, err = FeaturepropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
