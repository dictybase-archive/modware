package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeatureCvterms(t *testing.T) {
	t.Parallel()

	query := FeatureCvterms(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeatureCvtermsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvterm := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureCvterm.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureCvtermsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvterm := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureCvterms(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureCvtermsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvterm := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureCvtermSlice{featureCvterm}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeatureCvtermsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvterm := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeatureCvtermExists(tx, featureCvterm.FeatureCvtermID)
	if err != nil {
		t.Errorf("Unable to check if FeatureCvterm exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeatureCvtermExistsG to return true, but got false.")
	}
}
func testFeatureCvtermsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvterm := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	featureCvtermFound, err := FindFeatureCvterm(tx, featureCvterm.FeatureCvtermID)
	if err != nil {
		t.Error(err)
	}

	if featureCvtermFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeatureCvtermsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvterm := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureCvterms(tx).Bind(featureCvterm); err != nil {
		t.Error(err)
	}
}

func testFeatureCvtermsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvterm := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeatureCvterms(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeatureCvtermsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermOne := &FeatureCvterm{}
	featureCvtermTwo := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvtermOne, featureCvtermDBTypes, false, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}
	if err = randomize.Struct(seed, featureCvtermTwo, featureCvtermDBTypes, false, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureCvtermTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureCvterms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeatureCvtermsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featureCvtermOne := &FeatureCvterm{}
	featureCvtermTwo := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvtermOne, featureCvtermDBTypes, false, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}
	if err = randomize.Struct(seed, featureCvtermTwo, featureCvtermDBTypes, false, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureCvtermTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featureCvtermBeforeInsertHook(e boil.Executor, o *FeatureCvterm) error {
	*o = FeatureCvterm{}
	return nil
}

func featureCvtermAfterInsertHook(e boil.Executor, o *FeatureCvterm) error {
	*o = FeatureCvterm{}
	return nil
}

func featureCvtermAfterSelectHook(e boil.Executor, o *FeatureCvterm) error {
	*o = FeatureCvterm{}
	return nil
}

func featureCvtermBeforeUpdateHook(e boil.Executor, o *FeatureCvterm) error {
	*o = FeatureCvterm{}
	return nil
}

func featureCvtermAfterUpdateHook(e boil.Executor, o *FeatureCvterm) error {
	*o = FeatureCvterm{}
	return nil
}

func featureCvtermBeforeDeleteHook(e boil.Executor, o *FeatureCvterm) error {
	*o = FeatureCvterm{}
	return nil
}

func featureCvtermAfterDeleteHook(e boil.Executor, o *FeatureCvterm) error {
	*o = FeatureCvterm{}
	return nil
}

func featureCvtermBeforeUpsertHook(e boil.Executor, o *FeatureCvterm) error {
	*o = FeatureCvterm{}
	return nil
}

func featureCvtermAfterUpsertHook(e boil.Executor, o *FeatureCvterm) error {
	*o = FeatureCvterm{}
	return nil
}

func testFeatureCvtermsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeatureCvterm{}
	o := &FeatureCvterm{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featureCvtermDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm object: %s", err)
	}

	AddFeatureCvtermHook(boil.BeforeInsertHook, featureCvtermBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermBeforeInsertHooks = []FeatureCvtermHook{}

	AddFeatureCvtermHook(boil.AfterInsertHook, featureCvtermAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermAfterInsertHooks = []FeatureCvtermHook{}

	AddFeatureCvtermHook(boil.AfterSelectHook, featureCvtermAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featureCvtermAfterSelectHooks = []FeatureCvtermHook{}

	AddFeatureCvtermHook(boil.BeforeUpdateHook, featureCvtermBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featureCvtermBeforeUpdateHooks = []FeatureCvtermHook{}

	AddFeatureCvtermHook(boil.AfterUpdateHook, featureCvtermAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featureCvtermAfterUpdateHooks = []FeatureCvtermHook{}

	AddFeatureCvtermHook(boil.BeforeDeleteHook, featureCvtermBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featureCvtermBeforeDeleteHooks = []FeatureCvtermHook{}

	AddFeatureCvtermHook(boil.AfterDeleteHook, featureCvtermAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featureCvtermAfterDeleteHooks = []FeatureCvtermHook{}

	AddFeatureCvtermHook(boil.BeforeUpsertHook, featureCvtermBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermBeforeUpsertHooks = []FeatureCvtermHook{}

	AddFeatureCvtermHook(boil.AfterUpsertHook, featureCvtermAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermAfterUpsertHooks = []FeatureCvtermHook{}
}
func testFeatureCvtermsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvterm := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureCvtermsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvterm := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvterm.Insert(tx, featureCvtermColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureCvtermOneToOneFeatureCvtermPubUsingFeatureCvtermPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureCvtermPub
	var local FeatureCvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureCvtermPubDBTypes, true, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeatureCvtermID = local.FeatureCvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureCvtermPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureCvtermID != foreign.FeatureCvtermID {
		t.Errorf("want: %v, got %v", foreign.FeatureCvtermID, check.FeatureCvtermID)
	}

	slice := FeatureCvtermSlice{&local}
	if err = local.L.LoadFeatureCvtermPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureCvtermPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureCvtermPub = nil
	if err = local.L.LoadFeatureCvtermPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureCvtermPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureCvtermOneToOneFeatureCvtermDbxrefUsingFeatureCvtermDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureCvtermDbxref
	var local FeatureCvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureCvtermDbxrefDBTypes, true, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeatureCvtermID = local.FeatureCvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureCvtermDbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureCvtermID != foreign.FeatureCvtermID {
		t.Errorf("want: %v, got %v", foreign.FeatureCvtermID, check.FeatureCvtermID)
	}

	slice := FeatureCvtermSlice{&local}
	if err = local.L.LoadFeatureCvtermDbxref(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureCvtermDbxref == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureCvtermDbxref = nil
	if err = local.L.LoadFeatureCvtermDbxref(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureCvtermDbxref == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureCvtermOneToOneFeatureCvtermpropUsingFeatureCvtermprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureCvtermprop
	var local FeatureCvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureCvtermpropDBTypes, true, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeatureCvtermID = local.FeatureCvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureCvtermprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureCvtermID != foreign.FeatureCvtermID {
		t.Errorf("want: %v, got %v", foreign.FeatureCvtermID, check.FeatureCvtermID)
	}

	slice := FeatureCvtermSlice{&local}
	if err = local.L.LoadFeatureCvtermprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureCvtermprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureCvtermprop = nil
	if err = local.L.LoadFeatureCvtermprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureCvtermprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureCvtermOneToOneSetOpFeatureCvtermPubUsingFeatureCvtermPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureCvterm
	var b, c FeatureCvtermPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureCvtermDBTypes, false, strmangle.SetComplement(featureCvtermPrimaryKeyColumns, featureCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureCvtermPubDBTypes, false, strmangle.SetComplement(featureCvtermPubPrimaryKeyColumns, featureCvtermPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureCvtermPubDBTypes, false, strmangle.SetComplement(featureCvtermPubPrimaryKeyColumns, featureCvtermPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureCvtermPub{&b, &c} {
		err = a.SetFeatureCvtermPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureCvtermPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.FeatureCvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureCvtermID != x.FeatureCvtermID {
			t.Error("foreign key was wrong value", a.FeatureCvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeatureCvtermID))
		reflect.Indirect(reflect.ValueOf(&x.FeatureCvtermID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureCvtermID != x.FeatureCvtermID {
			t.Error("foreign key was wrong value", a.FeatureCvtermID, x.FeatureCvtermID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testFeatureCvtermOneToOneSetOpFeatureCvtermDbxrefUsingFeatureCvtermDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureCvterm
	var b, c FeatureCvtermDbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureCvtermDBTypes, false, strmangle.SetComplement(featureCvtermPrimaryKeyColumns, featureCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureCvtermDbxrefDBTypes, false, strmangle.SetComplement(featureCvtermDbxrefPrimaryKeyColumns, featureCvtermDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureCvtermDbxrefDBTypes, false, strmangle.SetComplement(featureCvtermDbxrefPrimaryKeyColumns, featureCvtermDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureCvtermDbxref{&b, &c} {
		err = a.SetFeatureCvtermDbxref(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureCvtermDbxref != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.FeatureCvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureCvtermID != x.FeatureCvtermID {
			t.Error("foreign key was wrong value", a.FeatureCvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeatureCvtermID))
		reflect.Indirect(reflect.ValueOf(&x.FeatureCvtermID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureCvtermID != x.FeatureCvtermID {
			t.Error("foreign key was wrong value", a.FeatureCvtermID, x.FeatureCvtermID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testFeatureCvtermOneToOneSetOpFeatureCvtermpropUsingFeatureCvtermprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureCvterm
	var b, c FeatureCvtermprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureCvtermDBTypes, false, strmangle.SetComplement(featureCvtermPrimaryKeyColumns, featureCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureCvtermpropDBTypes, false, strmangle.SetComplement(featureCvtermpropPrimaryKeyColumns, featureCvtermpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureCvtermpropDBTypes, false, strmangle.SetComplement(featureCvtermpropPrimaryKeyColumns, featureCvtermpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureCvtermprop{&b, &c} {
		err = a.SetFeatureCvtermprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureCvtermprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.FeatureCvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureCvtermID != x.FeatureCvtermID {
			t.Error("foreign key was wrong value", a.FeatureCvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeatureCvtermID))
		reflect.Indirect(reflect.ValueOf(&x.FeatureCvtermID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureCvtermID != x.FeatureCvtermID {
			t.Error("foreign key was wrong value", a.FeatureCvtermID, x.FeatureCvtermID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testFeatureCvtermToOneCvtermUsingCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureCvterm
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.CvtermID = foreign.CvtermID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Cvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := FeatureCvtermSlice{&local}
	if err = local.L.LoadCvterm(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Cvterm == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Cvterm = nil
	if err = local.L.LoadCvterm(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Cvterm == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureCvtermToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureCvterm
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
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

	slice := FeatureCvtermSlice{&local}
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

func testFeatureCvtermToOneFeatureUsingFeature(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureCvterm
	var foreign Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
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

	slice := FeatureCvtermSlice{&local}
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

func testFeatureCvtermToOneSetOpCvtermUsingCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureCvterm
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureCvtermDBTypes, false, strmangle.SetComplement(featureCvtermPrimaryKeyColumns, featureCvtermColumnsWithoutDefault)...); err != nil {
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
		err = a.SetCvterm(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Cvterm != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FeatureCvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CvtermID))
		reflect.Indirect(reflect.ValueOf(&a.CvtermID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID, x.CvtermID)
		}
	}
}
func testFeatureCvtermToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureCvterm
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureCvtermDBTypes, false, strmangle.SetComplement(featureCvtermPrimaryKeyColumns, featureCvtermColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeatureCvterm != &a {
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
func testFeatureCvtermToOneSetOpFeatureUsingFeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureCvterm
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureCvtermDBTypes, false, strmangle.SetComplement(featureCvtermPrimaryKeyColumns, featureCvtermColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeatureCvterm != &a {
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
func testFeatureCvtermsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvterm := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureCvterm.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureCvtermsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvterm := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureCvtermSlice{featureCvterm}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeatureCvtermsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvterm := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureCvterms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featureCvtermDBTypes = map[string]string{"CvtermID": "integer", "FeatureCvtermID": "integer", "FeatureID": "integer", "IsNot": "boolean", "PubID": "integer", "Rank": "integer"}
	_                    = bytes.MinRead
)

func testFeatureCvtermsUpdate(t *testing.T) {
	t.Parallel()

	if len(featureCvtermColumns) == len(featureCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureCvterm := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	if err = featureCvterm.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureCvtermsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featureCvtermColumns) == len(featureCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureCvterm := &FeatureCvterm{}
	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureCvterm, featureCvtermDBTypes, true, featureCvtermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featureCvtermColumns, featureCvtermPrimaryKeyColumns) {
		fields = featureCvtermColumns
	} else {
		fields = strmangle.SetComplement(
			featureCvtermColumns,
			featureCvtermPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featureCvterm))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeatureCvtermSlice{featureCvterm}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeatureCvtermsUpsert(t *testing.T) {
	t.Parallel()

	if len(featureCvtermColumns) == len(featureCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featureCvterm := FeatureCvterm{}
	if err = randomize.Struct(seed, &featureCvterm, featureCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvterm.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureCvterm: %s", err)
	}

	count, err := FeatureCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featureCvterm, featureCvtermDBTypes, false, featureCvtermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}

	if err = featureCvterm.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureCvterm: %s", err)
	}

	count, err = FeatureCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
