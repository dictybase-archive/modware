package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeatureRelationships(t *testing.T) {
	t.Parallel()

	query := FeatureRelationships(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeatureRelationshipsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationship := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureRelationship.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureRelationshipsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationship := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureRelationships(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureRelationshipsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationship := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureRelationshipSlice{featureRelationship}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeatureRelationshipsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationship := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeatureRelationshipExists(tx, featureRelationship.FeatureRelationshipID)
	if err != nil {
		t.Errorf("Unable to check if FeatureRelationship exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeatureRelationshipExistsG to return true, but got false.")
	}
}
func testFeatureRelationshipsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationship := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	featureRelationshipFound, err := FindFeatureRelationship(tx, featureRelationship.FeatureRelationshipID)
	if err != nil {
		t.Error(err)
	}

	if featureRelationshipFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeatureRelationshipsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationship := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureRelationships(tx).Bind(featureRelationship); err != nil {
		t.Error(err)
	}
}

func testFeatureRelationshipsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationship := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeatureRelationships(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeatureRelationshipsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationshipOne := &FeatureRelationship{}
	featureRelationshipTwo := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationshipOne, featureRelationshipDBTypes, false, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}
	if err = randomize.Struct(seed, featureRelationshipTwo, featureRelationshipDBTypes, false, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureRelationshipTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureRelationships(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeatureRelationshipsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featureRelationshipOne := &FeatureRelationship{}
	featureRelationshipTwo := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationshipOne, featureRelationshipDBTypes, false, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}
	if err = randomize.Struct(seed, featureRelationshipTwo, featureRelationshipDBTypes, false, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationshipOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureRelationshipTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featureRelationshipBeforeInsertHook(e boil.Executor, o *FeatureRelationship) error {
	*o = FeatureRelationship{}
	return nil
}

func featureRelationshipAfterInsertHook(e boil.Executor, o *FeatureRelationship) error {
	*o = FeatureRelationship{}
	return nil
}

func featureRelationshipAfterSelectHook(e boil.Executor, o *FeatureRelationship) error {
	*o = FeatureRelationship{}
	return nil
}

func featureRelationshipBeforeUpdateHook(e boil.Executor, o *FeatureRelationship) error {
	*o = FeatureRelationship{}
	return nil
}

func featureRelationshipAfterUpdateHook(e boil.Executor, o *FeatureRelationship) error {
	*o = FeatureRelationship{}
	return nil
}

func featureRelationshipBeforeDeleteHook(e boil.Executor, o *FeatureRelationship) error {
	*o = FeatureRelationship{}
	return nil
}

func featureRelationshipAfterDeleteHook(e boil.Executor, o *FeatureRelationship) error {
	*o = FeatureRelationship{}
	return nil
}

func featureRelationshipBeforeUpsertHook(e boil.Executor, o *FeatureRelationship) error {
	*o = FeatureRelationship{}
	return nil
}

func featureRelationshipAfterUpsertHook(e boil.Executor, o *FeatureRelationship) error {
	*o = FeatureRelationship{}
	return nil
}

func testFeatureRelationshipsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeatureRelationship{}
	o := &FeatureRelationship{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featureRelationshipDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship object: %s", err)
	}

	AddFeatureRelationshipHook(boil.BeforeInsertHook, featureRelationshipBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshipBeforeInsertHooks = []FeatureRelationshipHook{}

	AddFeatureRelationshipHook(boil.AfterInsertHook, featureRelationshipAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshipAfterInsertHooks = []FeatureRelationshipHook{}

	AddFeatureRelationshipHook(boil.AfterSelectHook, featureRelationshipAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featureRelationshipAfterSelectHooks = []FeatureRelationshipHook{}

	AddFeatureRelationshipHook(boil.BeforeUpdateHook, featureRelationshipBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featureRelationshipBeforeUpdateHooks = []FeatureRelationshipHook{}

	AddFeatureRelationshipHook(boil.AfterUpdateHook, featureRelationshipAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featureRelationshipAfterUpdateHooks = []FeatureRelationshipHook{}

	AddFeatureRelationshipHook(boil.BeforeDeleteHook, featureRelationshipBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featureRelationshipBeforeDeleteHooks = []FeatureRelationshipHook{}

	AddFeatureRelationshipHook(boil.AfterDeleteHook, featureRelationshipAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featureRelationshipAfterDeleteHooks = []FeatureRelationshipHook{}

	AddFeatureRelationshipHook(boil.BeforeUpsertHook, featureRelationshipBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshipBeforeUpsertHooks = []FeatureRelationshipHook{}

	AddFeatureRelationshipHook(boil.AfterUpsertHook, featureRelationshipAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featureRelationshipAfterUpsertHooks = []FeatureRelationshipHook{}
}
func testFeatureRelationshipsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationship := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureRelationshipsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationship := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationship.Insert(tx, featureRelationshipColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureRelationshipOneToOneFeatureRelationshippropUsingFeatureRelationshipprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureRelationshipprop
	var local FeatureRelationship

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureRelationshippropDBTypes, true, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeatureRelationshipID = local.FeatureRelationshipID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureRelationshipprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureRelationshipID != foreign.FeatureRelationshipID {
		t.Errorf("want: %v, got %v", foreign.FeatureRelationshipID, check.FeatureRelationshipID)
	}

	slice := FeatureRelationshipSlice{&local}
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

func testFeatureRelationshipOneToOneFeatureRelationshipPubUsingFeatureRelationshipPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureRelationshipPub
	var local FeatureRelationship

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureRelationshipPubDBTypes, true, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeatureRelationshipID = local.FeatureRelationshipID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureRelationshipPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureRelationshipID != foreign.FeatureRelationshipID {
		t.Errorf("want: %v, got %v", foreign.FeatureRelationshipID, check.FeatureRelationshipID)
	}

	slice := FeatureRelationshipSlice{&local}
	if err = local.L.LoadFeatureRelationshipPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureRelationshipPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureRelationshipPub = nil
	if err = local.L.LoadFeatureRelationshipPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureRelationshipPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureRelationshipOneToOneSetOpFeatureRelationshippropUsingFeatureRelationshipprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureRelationship
	var b, c FeatureRelationshipprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureRelationshipDBTypes, false, strmangle.SetComplement(featureRelationshipPrimaryKeyColumns, featureRelationshipColumnsWithoutDefault)...); err != nil {
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
		if x.R.FeatureRelationship != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureRelationshipID != x.FeatureRelationshipID {
			t.Error("foreign key was wrong value", a.FeatureRelationshipID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeatureRelationshipID))
		reflect.Indirect(reflect.ValueOf(&x.FeatureRelationshipID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureRelationshipID != x.FeatureRelationshipID {
			t.Error("foreign key was wrong value", a.FeatureRelationshipID, x.FeatureRelationshipID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testFeatureRelationshipOneToOneSetOpFeatureRelationshipPubUsingFeatureRelationshipPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureRelationship
	var b, c FeatureRelationshipPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureRelationshipDBTypes, false, strmangle.SetComplement(featureRelationshipPrimaryKeyColumns, featureRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureRelationshipPubDBTypes, false, strmangle.SetComplement(featureRelationshipPubPrimaryKeyColumns, featureRelationshipPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureRelationshipPubDBTypes, false, strmangle.SetComplement(featureRelationshipPubPrimaryKeyColumns, featureRelationshipPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureRelationshipPub{&b, &c} {
		err = a.SetFeatureRelationshipPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureRelationshipPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.FeatureRelationship != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureRelationshipID != x.FeatureRelationshipID {
			t.Error("foreign key was wrong value", a.FeatureRelationshipID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeatureRelationshipID))
		reflect.Indirect(reflect.ValueOf(&x.FeatureRelationshipID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureRelationshipID != x.FeatureRelationshipID {
			t.Error("foreign key was wrong value", a.FeatureRelationshipID, x.FeatureRelationshipID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testFeatureRelationshipToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureRelationship
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
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

	slice := FeatureRelationshipSlice{&local}
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

func testFeatureRelationshipToOneFeatureUsingObject(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureRelationship
	var foreign Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ObjectID = foreign.FeatureID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Object(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureID != foreign.FeatureID {
		t.Errorf("want: %v, got %v", foreign.FeatureID, check.FeatureID)
	}

	slice := FeatureRelationshipSlice{&local}
	if err = local.L.LoadObject(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Object == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Object = nil
	if err = local.L.LoadObject(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Object == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureRelationshipToOneFeatureUsingSubject(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureRelationship
	var foreign Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.SubjectID = foreign.FeatureID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Subject(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureID != foreign.FeatureID {
		t.Errorf("want: %v, got %v", foreign.FeatureID, check.FeatureID)
	}

	slice := FeatureRelationshipSlice{&local}
	if err = local.L.LoadSubject(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Subject == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Subject = nil
	if err = local.L.LoadSubject(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Subject == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureRelationshipToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureRelationship
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureRelationshipDBTypes, false, strmangle.SetComplement(featureRelationshipPrimaryKeyColumns, featureRelationshipColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeFeatureRelationship != &a {
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
func testFeatureRelationshipToOneSetOpFeatureUsingObject(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureRelationship
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureRelationshipDBTypes, false, strmangle.SetComplement(featureRelationshipPrimaryKeyColumns, featureRelationshipColumnsWithoutDefault)...); err != nil {
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
		err = a.SetObject(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Object != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ObjectFeatureRelationship != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ObjectID != x.FeatureID {
			t.Error("foreign key was wrong value", a.ObjectID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ObjectID))
		reflect.Indirect(reflect.ValueOf(&a.ObjectID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ObjectID != x.FeatureID {
			t.Error("foreign key was wrong value", a.ObjectID, x.FeatureID)
		}
	}
}
func testFeatureRelationshipToOneSetOpFeatureUsingSubject(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureRelationship
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureRelationshipDBTypes, false, strmangle.SetComplement(featureRelationshipPrimaryKeyColumns, featureRelationshipColumnsWithoutDefault)...); err != nil {
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
		err = a.SetSubject(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Subject != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SubjectFeatureRelationship != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SubjectID != x.FeatureID {
			t.Error("foreign key was wrong value", a.SubjectID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SubjectID))
		reflect.Indirect(reflect.ValueOf(&a.SubjectID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SubjectID != x.FeatureID {
			t.Error("foreign key was wrong value", a.SubjectID, x.FeatureID)
		}
	}
}
func testFeatureRelationshipsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationship := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureRelationship.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureRelationshipsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationship := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureRelationshipSlice{featureRelationship}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeatureRelationshipsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureRelationship := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureRelationships(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featureRelationshipDBTypes = map[string]string{"FeatureRelationshipID": "integer", "ObjectID": "integer", "Rank": "integer", "SubjectID": "integer", "TypeID": "integer", "Value": "text"}
	_                          = bytes.MinRead
)

func testFeatureRelationshipsUpdate(t *testing.T) {
	t.Parallel()

	if len(featureRelationshipColumns) == len(featureRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureRelationship := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	if err = featureRelationship.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureRelationshipsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featureRelationshipColumns) == len(featureRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureRelationship := &FeatureRelationship{}
	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureRelationship, featureRelationshipDBTypes, true, featureRelationshipPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featureRelationshipColumns, featureRelationshipPrimaryKeyColumns) {
		fields = featureRelationshipColumns
	} else {
		fields = strmangle.SetComplement(
			featureRelationshipColumns,
			featureRelationshipPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featureRelationship))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeatureRelationshipSlice{featureRelationship}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeatureRelationshipsUpsert(t *testing.T) {
	t.Parallel()

	if len(featureRelationshipColumns) == len(featureRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featureRelationship := FeatureRelationship{}
	if err = randomize.Struct(seed, &featureRelationship, featureRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureRelationship.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureRelationship: %s", err)
	}

	count, err := FeatureRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featureRelationship, featureRelationshipDBTypes, false, featureRelationshipPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}

	if err = featureRelationship.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureRelationship: %s", err)
	}

	count, err = FeatureRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
