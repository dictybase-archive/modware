package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeatureCvtermDbxrefs(t *testing.T) {
	t.Parallel()

	query := FeatureCvtermDbxrefs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeatureCvtermDbxrefsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermDbxref := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureCvtermDbxref.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureCvtermDbxrefsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermDbxref := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureCvtermDbxrefs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureCvtermDbxrefsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermDbxref := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureCvtermDbxrefSlice{featureCvtermDbxref}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeatureCvtermDbxrefsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermDbxref := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeatureCvtermDbxrefExists(tx, featureCvtermDbxref.FeatureCvtermDbxrefID)
	if err != nil {
		t.Errorf("Unable to check if FeatureCvtermDbxref exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeatureCvtermDbxrefExistsG to return true, but got false.")
	}
}
func testFeatureCvtermDbxrefsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermDbxref := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	featureCvtermDbxrefFound, err := FindFeatureCvtermDbxref(tx, featureCvtermDbxref.FeatureCvtermDbxrefID)
	if err != nil {
		t.Error(err)
	}

	if featureCvtermDbxrefFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeatureCvtermDbxrefsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermDbxref := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureCvtermDbxrefs(tx).Bind(featureCvtermDbxref); err != nil {
		t.Error(err)
	}
}

func testFeatureCvtermDbxrefsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermDbxref := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeatureCvtermDbxrefs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeatureCvtermDbxrefsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermDbxrefOne := &FeatureCvtermDbxref{}
	featureCvtermDbxrefTwo := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxrefOne, featureCvtermDbxrefDBTypes, false, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}
	if err = randomize.Struct(seed, featureCvtermDbxrefTwo, featureCvtermDbxrefDBTypes, false, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxrefOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureCvtermDbxrefTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureCvtermDbxrefs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeatureCvtermDbxrefsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featureCvtermDbxrefOne := &FeatureCvtermDbxref{}
	featureCvtermDbxrefTwo := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxrefOne, featureCvtermDbxrefDBTypes, false, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}
	if err = randomize.Struct(seed, featureCvtermDbxrefTwo, featureCvtermDbxrefDBTypes, false, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxrefOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureCvtermDbxrefTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featureCvtermDbxrefBeforeInsertHook(e boil.Executor, o *FeatureCvtermDbxref) error {
	*o = FeatureCvtermDbxref{}
	return nil
}

func featureCvtermDbxrefAfterInsertHook(e boil.Executor, o *FeatureCvtermDbxref) error {
	*o = FeatureCvtermDbxref{}
	return nil
}

func featureCvtermDbxrefAfterSelectHook(e boil.Executor, o *FeatureCvtermDbxref) error {
	*o = FeatureCvtermDbxref{}
	return nil
}

func featureCvtermDbxrefBeforeUpdateHook(e boil.Executor, o *FeatureCvtermDbxref) error {
	*o = FeatureCvtermDbxref{}
	return nil
}

func featureCvtermDbxrefAfterUpdateHook(e boil.Executor, o *FeatureCvtermDbxref) error {
	*o = FeatureCvtermDbxref{}
	return nil
}

func featureCvtermDbxrefBeforeDeleteHook(e boil.Executor, o *FeatureCvtermDbxref) error {
	*o = FeatureCvtermDbxref{}
	return nil
}

func featureCvtermDbxrefAfterDeleteHook(e boil.Executor, o *FeatureCvtermDbxref) error {
	*o = FeatureCvtermDbxref{}
	return nil
}

func featureCvtermDbxrefBeforeUpsertHook(e boil.Executor, o *FeatureCvtermDbxref) error {
	*o = FeatureCvtermDbxref{}
	return nil
}

func featureCvtermDbxrefAfterUpsertHook(e boil.Executor, o *FeatureCvtermDbxref) error {
	*o = FeatureCvtermDbxref{}
	return nil
}

func testFeatureCvtermDbxrefsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeatureCvtermDbxref{}
	o := &FeatureCvtermDbxref{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featureCvtermDbxrefDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref object: %s", err)
	}

	AddFeatureCvtermDbxrefHook(boil.BeforeInsertHook, featureCvtermDbxrefBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermDbxrefBeforeInsertHooks = []FeatureCvtermDbxrefHook{}

	AddFeatureCvtermDbxrefHook(boil.AfterInsertHook, featureCvtermDbxrefAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermDbxrefAfterInsertHooks = []FeatureCvtermDbxrefHook{}

	AddFeatureCvtermDbxrefHook(boil.AfterSelectHook, featureCvtermDbxrefAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featureCvtermDbxrefAfterSelectHooks = []FeatureCvtermDbxrefHook{}

	AddFeatureCvtermDbxrefHook(boil.BeforeUpdateHook, featureCvtermDbxrefBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featureCvtermDbxrefBeforeUpdateHooks = []FeatureCvtermDbxrefHook{}

	AddFeatureCvtermDbxrefHook(boil.AfterUpdateHook, featureCvtermDbxrefAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featureCvtermDbxrefAfterUpdateHooks = []FeatureCvtermDbxrefHook{}

	AddFeatureCvtermDbxrefHook(boil.BeforeDeleteHook, featureCvtermDbxrefBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featureCvtermDbxrefBeforeDeleteHooks = []FeatureCvtermDbxrefHook{}

	AddFeatureCvtermDbxrefHook(boil.AfterDeleteHook, featureCvtermDbxrefAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featureCvtermDbxrefAfterDeleteHooks = []FeatureCvtermDbxrefHook{}

	AddFeatureCvtermDbxrefHook(boil.BeforeUpsertHook, featureCvtermDbxrefBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermDbxrefBeforeUpsertHooks = []FeatureCvtermDbxrefHook{}

	AddFeatureCvtermDbxrefHook(boil.AfterUpsertHook, featureCvtermDbxrefAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featureCvtermDbxrefAfterUpsertHooks = []FeatureCvtermDbxrefHook{}
}
func testFeatureCvtermDbxrefsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermDbxref := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureCvtermDbxrefsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermDbxref := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxref.Insert(tx, featureCvtermDbxrefColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureCvtermDbxrefToOneFeatureCvtermUsingFeatureCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureCvtermDbxref
	var foreign FeatureCvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureCvtermDbxrefDBTypes, true, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
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

	slice := FeatureCvtermDbxrefSlice{&local}
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

func testFeatureCvtermDbxrefToOneDbxrefUsingDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureCvtermDbxref
	var foreign Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureCvtermDbxrefDBTypes, true, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.DbxrefID = foreign.DbxrefID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Dbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.DbxrefID != foreign.DbxrefID {
		t.Errorf("want: %v, got %v", foreign.DbxrefID, check.DbxrefID)
	}

	slice := FeatureCvtermDbxrefSlice{&local}
	if err = local.L.LoadDbxref(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Dbxref == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Dbxref = nil
	if err = local.L.LoadDbxref(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Dbxref == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureCvtermDbxrefToOneSetOpFeatureCvtermUsingFeatureCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureCvtermDbxref
	var b, c FeatureCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureCvtermDbxrefDBTypes, false, strmangle.SetComplement(featureCvtermDbxrefPrimaryKeyColumns, featureCvtermDbxrefColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeatureCvtermDbxref != &a {
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
func testFeatureCvtermDbxrefToOneSetOpDbxrefUsingDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureCvtermDbxref
	var b, c Dbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureCvtermDbxrefDBTypes, false, strmangle.SetComplement(featureCvtermDbxrefPrimaryKeyColumns, featureCvtermDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Dbxref{&b, &c} {
		err = a.SetDbxref(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Dbxref != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FeatureCvtermDbxref != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DbxrefID))
		reflect.Indirect(reflect.ValueOf(&a.DbxrefID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID, x.DbxrefID)
		}
	}
}
func testFeatureCvtermDbxrefsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermDbxref := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureCvtermDbxref.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureCvtermDbxrefsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermDbxref := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureCvtermDbxrefSlice{featureCvtermDbxref}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeatureCvtermDbxrefsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureCvtermDbxref := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureCvtermDbxrefs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featureCvtermDbxrefDBTypes = map[string]string{"DbxrefID": "integer", "FeatureCvtermDbxrefID": "integer", "FeatureCvtermID": "integer"}
	_                          = bytes.MinRead
)

func testFeatureCvtermDbxrefsUpdate(t *testing.T) {
	t.Parallel()

	if len(featureCvtermDbxrefColumns) == len(featureCvtermDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureCvtermDbxref := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	if err = featureCvtermDbxref.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureCvtermDbxrefsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featureCvtermDbxrefColumns) == len(featureCvtermDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureCvtermDbxref := &FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureCvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureCvtermDbxref, featureCvtermDbxrefDBTypes, true, featureCvtermDbxrefPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featureCvtermDbxrefColumns, featureCvtermDbxrefPrimaryKeyColumns) {
		fields = featureCvtermDbxrefColumns
	} else {
		fields = strmangle.SetComplement(
			featureCvtermDbxrefColumns,
			featureCvtermDbxrefPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featureCvtermDbxref))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeatureCvtermDbxrefSlice{featureCvtermDbxref}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeatureCvtermDbxrefsUpsert(t *testing.T) {
	t.Parallel()

	if len(featureCvtermDbxrefColumns) == len(featureCvtermDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featureCvtermDbxref := FeatureCvtermDbxref{}
	if err = randomize.Struct(seed, &featureCvtermDbxref, featureCvtermDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureCvtermDbxref.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureCvtermDbxref: %s", err)
	}

	count, err := FeatureCvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featureCvtermDbxref, featureCvtermDbxrefDBTypes, false, featureCvtermDbxrefPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}

	if err = featureCvtermDbxref.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureCvtermDbxref: %s", err)
	}

	count, err = FeatureCvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
