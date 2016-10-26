package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDbxrefs(t *testing.T) {
	t.Parallel()

	query := Dbxrefs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDbxrefsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxref := &Dbxref{}
	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dbxref.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDbxrefsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxref := &Dbxref{}
	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Dbxrefs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDbxrefsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxref := &Dbxref{}
	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DbxrefSlice{dbxref}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDbxrefsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxref := &Dbxref{}
	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DbxrefExists(tx, dbxref.DbxrefID)
	if err != nil {
		t.Errorf("Unable to check if Dbxref exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DbxrefExistsG to return true, but got false.")
	}
}
func testDbxrefsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxref := &Dbxref{}
	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	dbxrefFound, err := FindDbxref(tx, dbxref.DbxrefID)
	if err != nil {
		t.Error(err)
	}

	if dbxrefFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDbxrefsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxref := &Dbxref{}
	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Dbxrefs(tx).Bind(dbxref); err != nil {
		t.Error(err)
	}
}

func testDbxrefsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxref := &Dbxref{}
	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Dbxrefs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDbxrefsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxrefOne := &Dbxref{}
	dbxrefTwo := &Dbxref{}
	if err = randomize.Struct(seed, dbxrefOne, dbxrefDBTypes, false, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}
	if err = randomize.Struct(seed, dbxrefTwo, dbxrefDBTypes, false, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dbxrefTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Dbxrefs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDbxrefsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dbxrefOne := &Dbxref{}
	dbxrefTwo := &Dbxref{}
	if err = randomize.Struct(seed, dbxrefOne, dbxrefDBTypes, false, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}
	if err = randomize.Struct(seed, dbxrefTwo, dbxrefDBTypes, false, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dbxrefTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func dbxrefBeforeInsertHook(e boil.Executor, o *Dbxref) error {
	*o = Dbxref{}
	return nil
}

func dbxrefAfterInsertHook(e boil.Executor, o *Dbxref) error {
	*o = Dbxref{}
	return nil
}

func dbxrefAfterSelectHook(e boil.Executor, o *Dbxref) error {
	*o = Dbxref{}
	return nil
}

func dbxrefBeforeUpdateHook(e boil.Executor, o *Dbxref) error {
	*o = Dbxref{}
	return nil
}

func dbxrefAfterUpdateHook(e boil.Executor, o *Dbxref) error {
	*o = Dbxref{}
	return nil
}

func dbxrefBeforeDeleteHook(e boil.Executor, o *Dbxref) error {
	*o = Dbxref{}
	return nil
}

func dbxrefAfterDeleteHook(e boil.Executor, o *Dbxref) error {
	*o = Dbxref{}
	return nil
}

func dbxrefBeforeUpsertHook(e boil.Executor, o *Dbxref) error {
	*o = Dbxref{}
	return nil
}

func dbxrefAfterUpsertHook(e boil.Executor, o *Dbxref) error {
	*o = Dbxref{}
	return nil
}

func testDbxrefsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Dbxref{}
	o := &Dbxref{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dbxrefDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Dbxref object: %s", err)
	}

	AddDbxrefHook(boil.BeforeInsertHook, dbxrefBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dbxrefBeforeInsertHooks = []DbxrefHook{}

	AddDbxrefHook(boil.AfterInsertHook, dbxrefAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dbxrefAfterInsertHooks = []DbxrefHook{}

	AddDbxrefHook(boil.AfterSelectHook, dbxrefAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dbxrefAfterSelectHooks = []DbxrefHook{}

	AddDbxrefHook(boil.BeforeUpdateHook, dbxrefBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dbxrefBeforeUpdateHooks = []DbxrefHook{}

	AddDbxrefHook(boil.AfterUpdateHook, dbxrefAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dbxrefAfterUpdateHooks = []DbxrefHook{}

	AddDbxrefHook(boil.BeforeDeleteHook, dbxrefBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dbxrefBeforeDeleteHooks = []DbxrefHook{}

	AddDbxrefHook(boil.AfterDeleteHook, dbxrefAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dbxrefAfterDeleteHooks = []DbxrefHook{}

	AddDbxrefHook(boil.BeforeUpsertHook, dbxrefBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dbxrefBeforeUpsertHooks = []DbxrefHook{}

	AddDbxrefHook(boil.AfterUpsertHook, dbxrefAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dbxrefAfterUpsertHooks = []DbxrefHook{}
}
func testDbxrefsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxref := &Dbxref{}
	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDbxrefsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxref := &Dbxref{}
	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxref.Insert(tx, dbxrefColumns...); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDbxrefOneToOnePubDbxrefUsingPubDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign PubDbxref
	var local Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, pubDbxrefDBTypes, true, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.DbxrefID = local.DbxrefID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.PubDbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.DbxrefID != foreign.DbxrefID {
		t.Errorf("want: %v, got %v", foreign.DbxrefID, check.DbxrefID)
	}

	slice := DbxrefSlice{&local}
	if err = local.L.LoadPubDbxref(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.PubDbxref == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.PubDbxref = nil
	if err = local.L.LoadPubDbxref(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.PubDbxref == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDbxrefOneToOneCvtermUsingCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Cvterm
	var local Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.DbxrefID = local.DbxrefID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Cvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.DbxrefID != foreign.DbxrefID {
		t.Errorf("want: %v, got %v", foreign.DbxrefID, check.DbxrefID)
	}

	slice := DbxrefSlice{&local}
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

func testDbxrefOneToOneFeatureDbxrefUsingFeatureDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureDbxref
	var local Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureDbxrefDBTypes, true, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.DbxrefID = local.DbxrefID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureDbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.DbxrefID != foreign.DbxrefID {
		t.Errorf("want: %v, got %v", foreign.DbxrefID, check.DbxrefID)
	}

	slice := DbxrefSlice{&local}
	if err = local.L.LoadFeatureDbxref(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureDbxref == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureDbxref = nil
	if err = local.L.LoadFeatureDbxref(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureDbxref == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDbxrefOneToOneDbxrefpropUsingDbxrefprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Dbxrefprop
	var local Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, dbxrefpropDBTypes, true, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.DbxrefID = local.DbxrefID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Dbxrefprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.DbxrefID != foreign.DbxrefID {
		t.Errorf("want: %v, got %v", foreign.DbxrefID, check.DbxrefID)
	}

	slice := DbxrefSlice{&local}
	if err = local.L.LoadDbxrefprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Dbxrefprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Dbxrefprop = nil
	if err = local.L.LoadDbxrefprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Dbxrefprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDbxrefOneToOneFeatureCvtermDbxrefUsingFeatureCvtermDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureCvtermDbxref
	var local Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureCvtermDbxrefDBTypes, true, featureCvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermDbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.DbxrefID = local.DbxrefID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureCvtermDbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.DbxrefID != foreign.DbxrefID {
		t.Errorf("want: %v, got %v", foreign.DbxrefID, check.DbxrefID)
	}

	slice := DbxrefSlice{&local}
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

func testDbxrefOneToOneCvtermDbxrefUsingCvtermDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign CvtermDbxref
	var local Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, cvtermDbxrefDBTypes, true, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.DbxrefID = local.DbxrefID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.CvtermDbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.DbxrefID != foreign.DbxrefID {
		t.Errorf("want: %v, got %v", foreign.DbxrefID, check.DbxrefID)
	}

	slice := DbxrefSlice{&local}
	if err = local.L.LoadCvtermDbxref(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.CvtermDbxref == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.CvtermDbxref = nil
	if err = local.L.LoadCvtermDbxref(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.CvtermDbxref == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDbxrefOneToOneStockDbxrefUsingStockDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockDbxref
	var local Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockDbxrefDBTypes, true, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.DbxrefID = local.DbxrefID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockDbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.DbxrefID != foreign.DbxrefID {
		t.Errorf("want: %v, got %v", foreign.DbxrefID, check.DbxrefID)
	}

	slice := DbxrefSlice{&local}
	if err = local.L.LoadStockDbxref(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockDbxref == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockDbxref = nil
	if err = local.L.LoadStockDbxref(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockDbxref == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDbxrefOneToOneOrganismDbxrefUsingOrganismDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign OrganismDbxref
	var local Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, organismDbxrefDBTypes, true, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.DbxrefID = local.DbxrefID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.OrganismDbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.DbxrefID != foreign.DbxrefID {
		t.Errorf("want: %v, got %v", foreign.DbxrefID, check.DbxrefID)
	}

	slice := DbxrefSlice{&local}
	if err = local.L.LoadOrganismDbxref(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.OrganismDbxref == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.OrganismDbxref = nil
	if err = local.L.LoadOrganismDbxref(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.OrganismDbxref == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDbxrefOneToOneSetOpPubDbxrefUsingPubDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c PubDbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pubDbxrefDBTypes, false, strmangle.SetComplement(pubDbxrefPrimaryKeyColumns, pubDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pubDbxrefDBTypes, false, strmangle.SetComplement(pubDbxrefPrimaryKeyColumns, pubDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*PubDbxref{&b, &c} {
		err = a.SetPubDbxref(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.PubDbxref != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Dbxref != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.DbxrefID))
		reflect.Indirect(reflect.ValueOf(&x.DbxrefID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID, x.DbxrefID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testDbxrefOneToOneSetOpCvtermUsingCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
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
		if x.R.Dbxref != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.DbxrefID))
		reflect.Indirect(reflect.ValueOf(&x.DbxrefID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID, x.DbxrefID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testDbxrefOneToOneSetOpFeatureDbxrefUsingFeatureDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c FeatureDbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureDbxrefDBTypes, false, strmangle.SetComplement(featureDbxrefPrimaryKeyColumns, featureDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureDbxrefDBTypes, false, strmangle.SetComplement(featureDbxrefPrimaryKeyColumns, featureDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureDbxref{&b, &c} {
		err = a.SetFeatureDbxref(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureDbxref != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Dbxref != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.DbxrefID))
		reflect.Indirect(reflect.ValueOf(&x.DbxrefID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID, x.DbxrefID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testDbxrefOneToOneSetOpDbxrefpropUsingDbxrefprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c Dbxrefprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dbxrefpropDBTypes, false, strmangle.SetComplement(dbxrefpropPrimaryKeyColumns, dbxrefpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dbxrefpropDBTypes, false, strmangle.SetComplement(dbxrefpropPrimaryKeyColumns, dbxrefpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Dbxrefprop{&b, &c} {
		err = a.SetDbxrefprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Dbxrefprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Dbxref != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.DbxrefID))
		reflect.Indirect(reflect.ValueOf(&x.DbxrefID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID, x.DbxrefID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testDbxrefOneToOneSetOpFeatureCvtermDbxrefUsingFeatureCvtermDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c FeatureCvtermDbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
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
		if x.R.Dbxref != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.DbxrefID))
		reflect.Indirect(reflect.ValueOf(&x.DbxrefID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID, x.DbxrefID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testDbxrefOneToOneSetOpCvtermDbxrefUsingCvtermDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c CvtermDbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermDbxrefDBTypes, false, strmangle.SetComplement(cvtermDbxrefPrimaryKeyColumns, cvtermDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvtermDbxrefDBTypes, false, strmangle.SetComplement(cvtermDbxrefPrimaryKeyColumns, cvtermDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*CvtermDbxref{&b, &c} {
		err = a.SetCvtermDbxref(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.CvtermDbxref != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Dbxref != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.DbxrefID))
		reflect.Indirect(reflect.ValueOf(&x.DbxrefID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID, x.DbxrefID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testDbxrefOneToOneSetOpStockDbxrefUsingStockDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c StockDbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockDbxrefDBTypes, false, strmangle.SetComplement(stockDbxrefPrimaryKeyColumns, stockDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockDbxrefDBTypes, false, strmangle.SetComplement(stockDbxrefPrimaryKeyColumns, stockDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockDbxref{&b, &c} {
		err = a.SetStockDbxref(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockDbxref != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Dbxref != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.DbxrefID))
		reflect.Indirect(reflect.ValueOf(&x.DbxrefID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID, x.DbxrefID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testDbxrefOneToOneSetOpOrganismDbxrefUsingOrganismDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c OrganismDbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organismDbxrefDBTypes, false, strmangle.SetComplement(organismDbxrefPrimaryKeyColumns, organismDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, organismDbxrefDBTypes, false, strmangle.SetComplement(organismDbxrefPrimaryKeyColumns, organismDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*OrganismDbxref{&b, &c} {
		err = a.SetOrganismDbxref(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.OrganismDbxref != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Dbxref != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.DbxrefID))
		reflect.Indirect(reflect.ValueOf(&x.DbxrefID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID, x.DbxrefID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testDbxrefToManyStocks(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, stockDBTypes, false, stockColumnsWithDefault...)
	randomize.Struct(seed, &c, stockDBTypes, false, stockColumnsWithDefault...)
	b.DbxrefID.Valid = true
	c.DbxrefID.Valid = true
	b.DbxrefID.Int = a.DbxrefID
	c.DbxrefID.Int = a.DbxrefID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	stock, err := a.Stocks(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range stock {
		if v.DbxrefID.Int == b.DbxrefID.Int {
			bFound = true
		}
		if v.DbxrefID.Int == c.DbxrefID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DbxrefSlice{&a}
	if err = a.L.LoadStocks(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Stocks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Stocks = nil
	if err = a.L.LoadStocks(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Stocks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", stock)
	}
}

func testDbxrefToManyFeatures(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, featureDBTypes, false, featureColumnsWithDefault...)
	randomize.Struct(seed, &c, featureDBTypes, false, featureColumnsWithDefault...)
	b.DbxrefID.Valid = true
	c.DbxrefID.Valid = true
	b.DbxrefID.Int = a.DbxrefID
	c.DbxrefID.Int = a.DbxrefID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	feature, err := a.Features(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range feature {
		if v.DbxrefID.Int == b.DbxrefID.Int {
			bFound = true
		}
		if v.DbxrefID.Int == c.DbxrefID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DbxrefSlice{&a}
	if err = a.L.LoadFeatures(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Features); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Features = nil
	if err = a.L.LoadFeatures(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Features); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", feature)
	}
}

func testDbxrefToManyAddOpStocks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c, d, e Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Stock{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Stock{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddStocks(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.DbxrefID != first.DbxrefID.Int {
			t.Error("foreign key was wrong value", a.DbxrefID, first.DbxrefID.Int)
		}
		if a.DbxrefID != second.DbxrefID.Int {
			t.Error("foreign key was wrong value", a.DbxrefID, second.DbxrefID.Int)
		}

		if first.R.Dbxref != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Dbxref != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Stocks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Stocks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Stocks(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDbxrefToManySetOpStocks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c, d, e Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Stock{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetStocks(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Stocks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetStocks(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Stocks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.DbxrefID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.DbxrefID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.DbxrefID != d.DbxrefID.Int {
		t.Error("foreign key was wrong value", a.DbxrefID, d.DbxrefID.Int)
	}
	if a.DbxrefID != e.DbxrefID.Int {
		t.Error("foreign key was wrong value", a.DbxrefID, e.DbxrefID.Int)
	}

	if b.R.Dbxref != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Dbxref != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Dbxref != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Dbxref != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Stocks[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Stocks[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testDbxrefToManyRemoveOpStocks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c, d, e Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Stock{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddStocks(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Stocks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveStocks(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Stocks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.DbxrefID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.DbxrefID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Dbxref != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Dbxref != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Dbxref != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Dbxref != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Stocks) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Stocks[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Stocks[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testDbxrefToManyAddOpFeatures(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c, d, e Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Feature{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Feature{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddFeatures(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.DbxrefID != first.DbxrefID.Int {
			t.Error("foreign key was wrong value", a.DbxrefID, first.DbxrefID.Int)
		}
		if a.DbxrefID != second.DbxrefID.Int {
			t.Error("foreign key was wrong value", a.DbxrefID, second.DbxrefID.Int)
		}

		if first.R.Dbxref != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Dbxref != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Features[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Features[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Features(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDbxrefToManySetOpFeatures(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c, d, e Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Feature{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetFeatures(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Features(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetFeatures(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Features(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.DbxrefID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.DbxrefID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.DbxrefID != d.DbxrefID.Int {
		t.Error("foreign key was wrong value", a.DbxrefID, d.DbxrefID.Int)
	}
	if a.DbxrefID != e.DbxrefID.Int {
		t.Error("foreign key was wrong value", a.DbxrefID, e.DbxrefID.Int)
	}

	if b.R.Dbxref != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Dbxref != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Dbxref != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Dbxref != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Features[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Features[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testDbxrefToManyRemoveOpFeatures(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c, d, e Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Feature{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddFeatures(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Features(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveFeatures(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Features(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.DbxrefID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.DbxrefID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Dbxref != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Dbxref != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Dbxref != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Dbxref != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Features) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Features[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Features[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testDbxrefToOneDBUsingDB(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Dbxref
	var foreign DB

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dbDBTypes, true, dbColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.DBID = foreign.DBID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.DB(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.DBID != foreign.DBID {
		t.Errorf("want: %v, got %v", foreign.DBID, check.DBID)
	}

	slice := DbxrefSlice{&local}
	if err = local.L.LoadDB(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.DB == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.DB = nil
	if err = local.L.LoadDB(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.DB == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDbxrefToOneSetOpDBUsingDB(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxref
	var b, c DB

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dbDBTypes, false, strmangle.SetComplement(dbPrimaryKeyColumns, dbColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dbDBTypes, false, strmangle.SetComplement(dbPrimaryKeyColumns, dbColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DB{&b, &c} {
		err = a.SetDB(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.DB != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Dbxref != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DBID != x.DBID {
			t.Error("foreign key was wrong value", a.DBID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DBID))
		reflect.Indirect(reflect.ValueOf(&a.DBID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DBID != x.DBID {
			t.Error("foreign key was wrong value", a.DBID, x.DBID)
		}
	}
}
func testDbxrefsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxref := &Dbxref{}
	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dbxref.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDbxrefsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxref := &Dbxref{}
	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DbxrefSlice{dbxref}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDbxrefsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxref := &Dbxref{}
	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Dbxrefs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dbxrefDBTypes = map[string]string{"Accession": "character varying", "DBID": "integer", "DbxrefID": "integer", "Description": "text", "Version": "character varying"}
	_             = bytes.MinRead
)

func testDbxrefsUpdate(t *testing.T) {
	t.Parallel()

	if len(dbxrefColumns) == len(dbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dbxref := &Dbxref{}
	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	if err = dbxref.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDbxrefsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dbxrefColumns) == len(dbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dbxref := &Dbxref{}
	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dbxref, dbxrefDBTypes, true, dbxrefPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dbxrefColumns, dbxrefPrimaryKeyColumns) {
		fields = dbxrefColumns
	} else {
		fields = strmangle.SetComplement(
			dbxrefColumns,
			dbxrefPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(dbxref))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DbxrefSlice{dbxref}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDbxrefsUpsert(t *testing.T) {
	t.Parallel()

	if len(dbxrefColumns) == len(dbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	dbxref := Dbxref{}
	if err = randomize.Struct(seed, &dbxref, dbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxref.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Dbxref: %s", err)
	}

	count, err := Dbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &dbxref, dbxrefDBTypes, false, dbxrefPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	if err = dbxref.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Dbxref: %s", err)
	}

	count, err = Dbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
