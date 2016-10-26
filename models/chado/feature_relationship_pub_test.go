package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeatureRelationshipPubs(t *testing.T) {
	t.Parallel()

	query := FeatureRelationshipPubs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeatureRelationshipPubsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipPub := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureRelationshipPub.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureRelationshipPubsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipPub := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureRelationshipPubs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureRelationshipPubsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipPub := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureRelationshipPubSlice{featureRelationshipPub}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeatureRelationshipPubsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipPub := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeatureRelationshipPubExists(tx, featureRelationshipPub.FeatureRelationshipPubID)
	if err != nil {
		t.Errorf("Unable to check if FeatureRelationshipPub exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeatureRelationshipPubExistsG to return true, but got false.")
	}
}
func testFeatureRelationshipPubsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipPub := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	featureRelationshipPubFound, err := FindFeatureRelationshipPub(tx, featureRelationshipPub.FeatureRelationshipPubID)
	if err != nil {
		t.Error(err)
	}

	if featureRelationshipPubFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeatureRelationshipPubsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipPub := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureRelationshipPubs(tx).Bind(featureRelationshipPub); err != nil {
		t.Error(err)
	}
}

func testFeatureRelationshipPubsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipPub := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeatureRelationshipPubs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeatureRelationshipPubsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipPubOne := &FeatureRelationshipPub{}
	featureRelationshipPubTwo := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPubOne, featureRelationshipPubDBTypes, false, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}
	if err = randomize.Struct(seed, featureRelationshipPubTwo, featureRelationshipPubDBTypes, false, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureRelationshipPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureRelationshipPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeatureRelationshipPubsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featureRelationshipPubOne := &FeatureRelationshipPub{}
	featureRelationshipPubTwo := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPubOne, featureRelationshipPubDBTypes, false, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}
	if err = randomize.Struct(seed, featureRelationshipPubTwo, featureRelationshipPubDBTypes, false, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureRelationshipPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featureRelationshipPubBeforeInsertHook(e boil.Executor, o *FeatureRelationshipPub) error {
	*o = FeatureRelationshipPub{}
	return nil
}

func featureRelationshipPubAfterInsertHook(e boil.Executor, o *FeatureRelationshipPub) error {
	*o = FeatureRelationshipPub{}
	return nil
}

func featureRelationshipPubAfterSelectHook(e boil.Executor, o *FeatureRelationshipPub) error {
	*o = FeatureRelationshipPub{}
	return nil
}

func featureRelationshipPubBeforeUpdateHook(e boil.Executor, o *FeatureRelationshipPub) error {
	*o = FeatureRelationshipPub{}
	return nil
}

func featureRelationshipPubAfterUpdateHook(e boil.Executor, o *FeatureRelationshipPub) error {
	*o = FeatureRelationshipPub{}
	return nil
}

func featureRelationshipPubBeforeDeleteHook(e boil.Executor, o *FeatureRelationshipPub) error {
	*o = FeatureRelationshipPub{}
	return nil
}

func featureRelationshipPubAfterDeleteHook(e boil.Executor, o *FeatureRelationshipPub) error {
	*o = FeatureRelationshipPub{}
	return nil
}

func featureRelationshipPubBeforeUpsertHook(e boil.Executor, o *FeatureRelationshipPub) error {
	*o = FeatureRelationshipPub{}
	return nil
}

func featureRelationshipPubAfterUpsertHook(e boil.Executor, o *FeatureRelationshipPub) error {
	*o = FeatureRelationshipPub{}
	return nil
}

func testFeatureRelationshipPubsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeatureRelationshipPub{}
	o := &FeatureRelationshipPub{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featureRelationshipPubDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub object: %s", err)
	}

	AddFeatureRelationshipPubHook(boil.BeforeInsertHook, featureRelationshipPubBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshipPubBeforeInsertHooks = []FeatureRelationshipPubHook{}

	AddFeatureRelationshipPubHook(boil.AfterInsertHook, featureRelationshipPubAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshipPubAfterInsertHooks = []FeatureRelationshipPubHook{}

	AddFeatureRelationshipPubHook(boil.AfterSelectHook, featureRelationshipPubAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featureRelationshipPubAfterSelectHooks = []FeatureRelationshipPubHook{}

	AddFeatureRelationshipPubHook(boil.BeforeUpdateHook, featureRelationshipPubBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featureRelationshipPubBeforeUpdateHooks = []FeatureRelationshipPubHook{}

	AddFeatureRelationshipPubHook(boil.AfterUpdateHook, featureRelationshipPubAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featureRelationshipPubAfterUpdateHooks = []FeatureRelationshipPubHook{}

	AddFeatureRelationshipPubHook(boil.BeforeDeleteHook, featureRelationshipPubBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featureRelationshipPubBeforeDeleteHooks = []FeatureRelationshipPubHook{}

	AddFeatureRelationshipPubHook(boil.AfterDeleteHook, featureRelationshipPubAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featureRelationshipPubAfterDeleteHooks = []FeatureRelationshipPubHook{}

	AddFeatureRelationshipPubHook(boil.BeforeUpsertHook, featureRelationshipPubBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshipPubBeforeUpsertHooks = []FeatureRelationshipPubHook{}

	AddFeatureRelationshipPubHook(boil.AfterUpsertHook, featureRelationshipPubAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshipPubAfterUpsertHooks = []FeatureRelationshipPubHook{}
}
func testFeatureRelationshipPubsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipPub := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureRelationshipPubsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipPub := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPub.Insert(tx, featureRelationshipPubColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureRelationshipPubToOneFeatureRelationshipUsingFeatureRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureRelationshipPub
	var foreign FeatureRelationship

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureRelationshipPubDBTypes, true, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
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

	slice := FeatureRelationshipPubSlice{&local}
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

func testFeatureRelationshipPubToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureRelationshipPub
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureRelationshipPubDBTypes, true, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
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

	slice := FeatureRelationshipPubSlice{&local}
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

func testFeatureRelationshipPubToOneSetOpFeatureRelationshipUsingFeatureRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureRelationshipPub
	var b, c FeatureRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureRelationshipPubDBTypes, false, strmangle.SetComplement(featureRelationshipPubPrimaryKeyColumns, featureRelationshipPubColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeatureRelationshipPub != &a {
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
func testFeatureRelationshipPubToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureRelationshipPub
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureRelationshipPubDBTypes, false, strmangle.SetComplement(featureRelationshipPubPrimaryKeyColumns, featureRelationshipPubColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeatureRelationshipPub != &a {
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
func testFeatureRelationshipPubsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipPub := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureRelationshipPub.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureRelationshipPubsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipPub := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureRelationshipPubSlice{featureRelationshipPub}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeatureRelationshipPubsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipPub := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureRelationshipPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featureRelationshipPubDBTypes = map[string]string{"FeatureRelationshipID": "integer", "FeatureRelationshipPubID": "integer", "PubID": "integer"}
	_                             = bytes.MinRead
)

func testFeatureRelationshipPubsUpdate(t *testing.T) {
	t.Parallel()

	if len(featureRelationshipPubColumns) == len(featureRelationshipPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureRelationshipPub := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	if err = featureRelationshipPub.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureRelationshipPubsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featureRelationshipPubColumns) == len(featureRelationshipPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureRelationshipPub := &FeatureRelationshipPub{}
	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureRelationshipPub, featureRelationshipPubDBTypes, true, featureRelationshipPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featureRelationshipPubColumns, featureRelationshipPubPrimaryKeyColumns) {
		fields = featureRelationshipPubColumns
	} else {
		fields = strmangle.SetComplement(
			featureRelationshipPubColumns,
			featureRelationshipPubPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featureRelationshipPub))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeatureRelationshipPubSlice{featureRelationshipPub}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeatureRelationshipPubsUpsert(t *testing.T) {
	t.Parallel()

	if len(featureRelationshipPubColumns) == len(featureRelationshipPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featureRelationshipPub := FeatureRelationshipPub{}
	if err = randomize.Struct(seed, &featureRelationshipPub, featureRelationshipPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipPub.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureRelationshipPub: %s", err)
	}

	count, err := FeatureRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featureRelationshipPub, featureRelationshipPubDBTypes, false, featureRelationshipPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}

	if err = featureRelationshipPub.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureRelationshipPub: %s", err)
	}

	count, err = FeatureRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
