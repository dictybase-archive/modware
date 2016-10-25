package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeatureDbxrefs(t *testing.T) {
	t.Parallel()

	query := FeatureDbxrefs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeatureDbxrefsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureDbxref := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureDbxref.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureDbxrefsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureDbxref := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureDbxrefs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeatureDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureDbxrefsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureDbxref := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureDbxrefSlice{featureDbxref}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeatureDbxrefsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureDbxref := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeatureDbxrefExists(tx, featureDbxref.FeatureDbxrefID)
	if err != nil {
		t.Errorf("Unable to check if FeatureDbxref exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeatureDbxrefExistsG to return true, but got false.")
	}
}
func testFeatureDbxrefsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureDbxref := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	featureDbxrefFound, err := FindFeatureDbxref(tx, featureDbxref.FeatureDbxrefID)
	if err != nil {
		t.Error(err)
	}

	if featureDbxrefFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeatureDbxrefsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureDbxref := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureDbxrefs(tx).Bind(featureDbxref); err != nil {
		t.Error(err)
	}
}

func testFeatureDbxrefsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureDbxref := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeatureDbxrefs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeatureDbxrefsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureDbxrefOne := &FeatureDbxref{}
	featureDbxrefTwo := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxrefOne, featureDbxrefDBTypes, false, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}
	if err = randomize.Struct(seed, featureDbxrefTwo, featureDbxrefDBTypes, false, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxrefOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureDbxrefTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureDbxrefs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeatureDbxrefsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featureDbxrefOne := &FeatureDbxref{}
	featureDbxrefTwo := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxrefOne, featureDbxrefDBTypes, false, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}
	if err = randomize.Struct(seed, featureDbxrefTwo, featureDbxrefDBTypes, false, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxrefOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureDbxrefTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featureDbxrefBeforeInsertHook(e boil.Executor, o *FeatureDbxref) error {
	*o = FeatureDbxref{}
	return nil
}

func featureDbxrefAfterInsertHook(e boil.Executor, o *FeatureDbxref) error {
	*o = FeatureDbxref{}
	return nil
}

func featureDbxrefAfterSelectHook(e boil.Executor, o *FeatureDbxref) error {
	*o = FeatureDbxref{}
	return nil
}

func featureDbxrefBeforeUpdateHook(e boil.Executor, o *FeatureDbxref) error {
	*o = FeatureDbxref{}
	return nil
}

func featureDbxrefAfterUpdateHook(e boil.Executor, o *FeatureDbxref) error {
	*o = FeatureDbxref{}
	return nil
}

func featureDbxrefBeforeDeleteHook(e boil.Executor, o *FeatureDbxref) error {
	*o = FeatureDbxref{}
	return nil
}

func featureDbxrefAfterDeleteHook(e boil.Executor, o *FeatureDbxref) error {
	*o = FeatureDbxref{}
	return nil
}

func featureDbxrefBeforeUpsertHook(e boil.Executor, o *FeatureDbxref) error {
	*o = FeatureDbxref{}
	return nil
}

func featureDbxrefAfterUpsertHook(e boil.Executor, o *FeatureDbxref) error {
	*o = FeatureDbxref{}
	return nil
}

func testFeatureDbxrefsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeatureDbxref{}
	o := &FeatureDbxref{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featureDbxrefDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref object: %s", err)
	}

	AddFeatureDbxrefHook(boil.BeforeInsertHook, featureDbxrefBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featureDbxrefBeforeInsertHooks = []FeatureDbxrefHook{}

	AddFeatureDbxrefHook(boil.AfterInsertHook, featureDbxrefAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featureDbxrefAfterInsertHooks = []FeatureDbxrefHook{}

	AddFeatureDbxrefHook(boil.AfterSelectHook, featureDbxrefAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featureDbxrefAfterSelectHooks = []FeatureDbxrefHook{}

	AddFeatureDbxrefHook(boil.BeforeUpdateHook, featureDbxrefBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featureDbxrefBeforeUpdateHooks = []FeatureDbxrefHook{}

	AddFeatureDbxrefHook(boil.AfterUpdateHook, featureDbxrefAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featureDbxrefAfterUpdateHooks = []FeatureDbxrefHook{}

	AddFeatureDbxrefHook(boil.BeforeDeleteHook, featureDbxrefBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featureDbxrefBeforeDeleteHooks = []FeatureDbxrefHook{}

	AddFeatureDbxrefHook(boil.AfterDeleteHook, featureDbxrefAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featureDbxrefAfterDeleteHooks = []FeatureDbxrefHook{}

	AddFeatureDbxrefHook(boil.BeforeUpsertHook, featureDbxrefBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featureDbxrefBeforeUpsertHooks = []FeatureDbxrefHook{}

	AddFeatureDbxrefHook(boil.AfterUpsertHook, featureDbxrefAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featureDbxrefAfterUpsertHooks = []FeatureDbxrefHook{}
}
func testFeatureDbxrefsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureDbxref := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureDbxrefsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureDbxref := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxref.Insert(tx, featureDbxrefColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeatureDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureDbxrefToOneDbxrefUsingDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureDbxref
	var foreign Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureDbxrefDBTypes, true, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
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

	slice := FeatureDbxrefSlice{&local}
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

func testFeatureDbxrefToOneFeatureUsingFeature(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureDbxref
	var foreign Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureDbxrefDBTypes, true, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
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

	slice := FeatureDbxrefSlice{&local}
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

func testFeatureDbxrefToOneSetOpDbxrefUsingDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureDbxref
	var b, c Dbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDbxrefDBTypes, false, strmangle.SetComplement(featureDbxrefPrimaryKeyColumns, featureDbxrefColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeatureDbxref != &a {
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
func testFeatureDbxrefToOneSetOpFeatureUsingFeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureDbxref
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDbxrefDBTypes, false, strmangle.SetComplement(featureDbxrefPrimaryKeyColumns, featureDbxrefColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeatureDbxref != &a {
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
func testFeatureDbxrefsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureDbxref := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureDbxref.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureDbxrefsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureDbxref := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureDbxrefSlice{featureDbxref}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeatureDbxrefsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureDbxref := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureDbxrefs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featureDbxrefDBTypes = map[string]string{"DbxrefID": "integer", "FeatureDbxrefID": "integer", "FeatureID": "integer", "IsCurrent": "boolean"}
	_                    = bytes.MinRead
)

func testFeatureDbxrefsUpdate(t *testing.T) {
	t.Parallel()

	if len(featureDbxrefColumns) == len(featureDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureDbxref := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	if err = featureDbxref.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureDbxrefsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featureDbxrefColumns) == len(featureDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureDbxref := &FeatureDbxref{}
	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureDbxref, featureDbxrefDBTypes, true, featureDbxrefPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featureDbxrefColumns, featureDbxrefPrimaryKeyColumns) {
		fields = featureDbxrefColumns
	} else {
		fields = strmangle.SetComplement(
			featureDbxrefColumns,
			featureDbxrefPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featureDbxref))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeatureDbxrefSlice{featureDbxref}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeatureDbxrefsUpsert(t *testing.T) {
	t.Parallel()

	if len(featureDbxrefColumns) == len(featureDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featureDbxref := FeatureDbxref{}
	if err = randomize.Struct(seed, &featureDbxref, featureDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureDbxref.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureDbxref: %s", err)
	}

	count, err := FeatureDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featureDbxref, featureDbxrefDBTypes, false, featureDbxrefPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}

	if err = featureDbxref.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureDbxref: %s", err)
	}

	count, err = FeatureDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
