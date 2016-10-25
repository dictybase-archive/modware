package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeatureCvtermPubs(t *testing.T) {
	t.Parallel()

	query := FeatureCvtermPubs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeatureCvtermPubsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermPub := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureCvtermPub.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureCvtermPubsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermPub := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureCvtermPubs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureCvtermPubsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermPub := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureCvtermPubSlice{featureCvtermPub}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeatureCvtermPubsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermPub := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPub.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeatureCvtermPubExists(tx, featureCvtermPub.FeatureCvtermPubID)
	if err != nil {
		t.Errorf("Unable to check if FeatureCvtermPub exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeatureCvtermPubExistsG to return true, but got false.")
	}
}
func testFeatureCvtermPubsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermPub := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPub.Insert(tx); err != nil {
		t.Error(err)
	}

	featureCvtermPubFound, err := FindFeatureCvtermPub(tx, featureCvtermPub.FeatureCvtermPubID)
	if err != nil {
		t.Error(err)
	}

	if featureCvtermPubFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeatureCvtermPubsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermPub := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureCvtermPubs(tx).Bind(featureCvtermPub); err != nil {
		t.Error(err)
	}
}

func testFeatureCvtermPubsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermPub := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeatureCvtermPubs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeatureCvtermPubsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermPubOne := &FeatureCvtermPub{}
	featureCvtermPubTwo := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPubOne, featureCvtermPubDBTypes, false, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}
	if err = randomize.Struct(seed, featureCvtermPubTwo, featureCvtermPubDBTypes, false, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureCvtermPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureCvtermPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeatureCvtermPubsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featureCvtermPubOne := &FeatureCvtermPub{}
	featureCvtermPubTwo := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPubOne, featureCvtermPubDBTypes, false, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}
	if err = randomize.Struct(seed, featureCvtermPubTwo, featureCvtermPubDBTypes, false, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureCvtermPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featureCvtermPubBeforeInsertHook(e boil.Executor, o *FeatureCvtermPub) error {
	*o = FeatureCvtermPub{}
	return nil
}

func featureCvtermPubAfterInsertHook(e boil.Executor, o *FeatureCvtermPub) error {
	*o = FeatureCvtermPub{}
	return nil
}

func featureCvtermPubAfterSelectHook(e boil.Executor, o *FeatureCvtermPub) error {
	*o = FeatureCvtermPub{}
	return nil
}

func featureCvtermPubBeforeUpdateHook(e boil.Executor, o *FeatureCvtermPub) error {
	*o = FeatureCvtermPub{}
	return nil
}

func featureCvtermPubAfterUpdateHook(e boil.Executor, o *FeatureCvtermPub) error {
	*o = FeatureCvtermPub{}
	return nil
}

func featureCvtermPubBeforeDeleteHook(e boil.Executor, o *FeatureCvtermPub) error {
	*o = FeatureCvtermPub{}
	return nil
}

func featureCvtermPubAfterDeleteHook(e boil.Executor, o *FeatureCvtermPub) error {
	*o = FeatureCvtermPub{}
	return nil
}

func featureCvtermPubBeforeUpsertHook(e boil.Executor, o *FeatureCvtermPub) error {
	*o = FeatureCvtermPub{}
	return nil
}

func featureCvtermPubAfterUpsertHook(e boil.Executor, o *FeatureCvtermPub) error {
	*o = FeatureCvtermPub{}
	return nil
}

func testFeatureCvtermPubsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeatureCvtermPub{}
	o := &FeatureCvtermPub{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featureCvtermPubDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub object: %s", err)
	}

	AddFeatureCvtermPubHook(boil.BeforeInsertHook, featureCvtermPubBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermPubBeforeInsertHooks = []FeatureCvtermPubHook{}

	AddFeatureCvtermPubHook(boil.AfterInsertHook, featureCvtermPubAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermPubAfterInsertHooks = []FeatureCvtermPubHook{}

	AddFeatureCvtermPubHook(boil.AfterSelectHook, featureCvtermPubAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featureCvtermPubAfterSelectHooks = []FeatureCvtermPubHook{}

	AddFeatureCvtermPubHook(boil.BeforeUpdateHook, featureCvtermPubBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featureCvtermPubBeforeUpdateHooks = []FeatureCvtermPubHook{}

	AddFeatureCvtermPubHook(boil.AfterUpdateHook, featureCvtermPubAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featureCvtermPubAfterUpdateHooks = []FeatureCvtermPubHook{}

	AddFeatureCvtermPubHook(boil.BeforeDeleteHook, featureCvtermPubBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featureCvtermPubBeforeDeleteHooks = []FeatureCvtermPubHook{}

	AddFeatureCvtermPubHook(boil.AfterDeleteHook, featureCvtermPubAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featureCvtermPubAfterDeleteHooks = []FeatureCvtermPubHook{}

	AddFeatureCvtermPubHook(boil.BeforeUpsertHook, featureCvtermPubBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermPubBeforeUpsertHooks = []FeatureCvtermPubHook{}

	AddFeatureCvtermPubHook(boil.AfterUpsertHook, featureCvtermPubAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermPubAfterUpsertHooks = []FeatureCvtermPubHook{}
}
func testFeatureCvtermPubsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermPub := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureCvtermPubsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermPub := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPub.Insert(tx, featureCvtermPubColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureCvtermPubToOneFeatureCvtermUsingFeatureCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureCvtermPub
	var foreign FeatureCvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureCvtermPubDBTypes, true, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
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

	slice := FeatureCvtermPubSlice{&local}
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

func testFeatureCvtermPubToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureCvtermPub
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureCvtermPubDBTypes, true, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
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

	slice := FeatureCvtermPubSlice{&local}
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

func testFeatureCvtermPubToOneSetOpFeatureCvtermUsingFeatureCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureCvtermPub
	var b, c FeatureCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureCvtermPubDBTypes, false, strmangle.SetComplement(featureCvtermPubPrimaryKeyColumns, featureCvtermPubColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeatureCvtermPub != &a {
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
func testFeatureCvtermPubToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureCvtermPub
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureCvtermPubDBTypes, false, strmangle.SetComplement(featureCvtermPubPrimaryKeyColumns, featureCvtermPubColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeatureCvtermPub != &a {
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
func testFeatureCvtermPubsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermPub := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureCvtermPub.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureCvtermPubsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermPub := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureCvtermPubSlice{featureCvtermPub}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeatureCvtermPubsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermPub := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureCvtermPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featureCvtermPubDBTypes = map[string]string{"FeatureCvtermID": "integer", "FeatureCvtermPubID": "integer", "PubID": "integer"}
	_                       = bytes.MinRead
)

func testFeatureCvtermPubsUpdate(t *testing.T) {
	t.Parallel()

	if len(featureCvtermPubColumns) == len(featureCvtermPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureCvtermPub := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	if err = featureCvtermPub.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureCvtermPubsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featureCvtermPubColumns) == len(featureCvtermPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureCvtermPub := &FeatureCvtermPub{}
	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureCvtermPub, featureCvtermPubDBTypes, true, featureCvtermPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featureCvtermPubColumns, featureCvtermPubPrimaryKeyColumns) {
		fields = featureCvtermPubColumns
	} else {
		fields = strmangle.SetComplement(
			featureCvtermPubColumns,
			featureCvtermPubPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featureCvtermPub))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeatureCvtermPubSlice{featureCvtermPub}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeatureCvtermPubsUpsert(t *testing.T) {
	t.Parallel()

	if len(featureCvtermPubColumns) == len(featureCvtermPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featureCvtermPub := FeatureCvtermPub{}
	if err = randomize.Struct(seed, &featureCvtermPub, featureCvtermPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermPub.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureCvtermPub: %s", err)
	}

	count, err := FeatureCvtermPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featureCvtermPub, featureCvtermPubDBTypes, false, featureCvtermPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}

	if err = featureCvtermPub.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureCvtermPub: %s", err)
	}

	count, err = FeatureCvtermPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
