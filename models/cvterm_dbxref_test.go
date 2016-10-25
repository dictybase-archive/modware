package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testCvtermDbxrefs(t *testing.T) {
	t.Parallel()

	query := CvtermDbxrefs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testCvtermDbxrefsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermDbxref := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cvtermDbxref.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := CvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCvtermDbxrefsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermDbxref := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = CvtermDbxrefs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := CvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCvtermDbxrefsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermDbxref := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CvtermDbxrefSlice{cvtermDbxref}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := CvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testCvtermDbxrefsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermDbxref := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := CvtermDbxrefExists(tx, cvtermDbxref.CvtermDbxrefID)
	if err != nil {
		t.Errorf("Unable to check if CvtermDbxref exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CvtermDbxrefExistsG to return true, but got false.")
	}
}
func testCvtermDbxrefsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermDbxref := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	cvtermDbxrefFound, err := FindCvtermDbxref(tx, cvtermDbxref.CvtermDbxrefID)
	if err != nil {
		t.Error(err)
	}

	if cvtermDbxrefFound == nil {
		t.Error("want a record, got nil")
	}
}
func testCvtermDbxrefsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermDbxref := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = CvtermDbxrefs(tx).Bind(cvtermDbxref); err != nil {
		t.Error(err)
	}
}

func testCvtermDbxrefsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermDbxref := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := CvtermDbxrefs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCvtermDbxrefsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermDbxrefOne := &CvtermDbxref{}
	cvtermDbxrefTwo := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxrefOne, cvtermDbxrefDBTypes, false, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}
	if err = randomize.Struct(seed, cvtermDbxrefTwo, cvtermDbxrefDBTypes, false, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxrefOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvtermDbxrefTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := CvtermDbxrefs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCvtermDbxrefsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	cvtermDbxrefOne := &CvtermDbxref{}
	cvtermDbxrefTwo := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxrefOne, cvtermDbxrefDBTypes, false, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}
	if err = randomize.Struct(seed, cvtermDbxrefTwo, cvtermDbxrefDBTypes, false, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxrefOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvtermDbxrefTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func cvtermDbxrefBeforeInsertHook(e boil.Executor, o *CvtermDbxref) error {
	*o = CvtermDbxref{}
	return nil
}

func cvtermDbxrefAfterInsertHook(e boil.Executor, o *CvtermDbxref) error {
	*o = CvtermDbxref{}
	return nil
}

func cvtermDbxrefAfterSelectHook(e boil.Executor, o *CvtermDbxref) error {
	*o = CvtermDbxref{}
	return nil
}

func cvtermDbxrefBeforeUpdateHook(e boil.Executor, o *CvtermDbxref) error {
	*o = CvtermDbxref{}
	return nil
}

func cvtermDbxrefAfterUpdateHook(e boil.Executor, o *CvtermDbxref) error {
	*o = CvtermDbxref{}
	return nil
}

func cvtermDbxrefBeforeDeleteHook(e boil.Executor, o *CvtermDbxref) error {
	*o = CvtermDbxref{}
	return nil
}

func cvtermDbxrefAfterDeleteHook(e boil.Executor, o *CvtermDbxref) error {
	*o = CvtermDbxref{}
	return nil
}

func cvtermDbxrefBeforeUpsertHook(e boil.Executor, o *CvtermDbxref) error {
	*o = CvtermDbxref{}
	return nil
}

func cvtermDbxrefAfterUpsertHook(e boil.Executor, o *CvtermDbxref) error {
	*o = CvtermDbxref{}
	return nil
}

func testCvtermDbxrefsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &CvtermDbxref{}
	o := &CvtermDbxref{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, cvtermDbxrefDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref object: %s", err)
	}

	AddCvtermDbxrefHook(boil.BeforeInsertHook, cvtermDbxrefBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	cvtermDbxrefBeforeInsertHooks = []CvtermDbxrefHook{}

	AddCvtermDbxrefHook(boil.AfterInsertHook, cvtermDbxrefAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	cvtermDbxrefAfterInsertHooks = []CvtermDbxrefHook{}

	AddCvtermDbxrefHook(boil.AfterSelectHook, cvtermDbxrefAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	cvtermDbxrefAfterSelectHooks = []CvtermDbxrefHook{}

	AddCvtermDbxrefHook(boil.BeforeUpdateHook, cvtermDbxrefBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	cvtermDbxrefBeforeUpdateHooks = []CvtermDbxrefHook{}

	AddCvtermDbxrefHook(boil.AfterUpdateHook, cvtermDbxrefAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	cvtermDbxrefAfterUpdateHooks = []CvtermDbxrefHook{}

	AddCvtermDbxrefHook(boil.BeforeDeleteHook, cvtermDbxrefBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	cvtermDbxrefBeforeDeleteHooks = []CvtermDbxrefHook{}

	AddCvtermDbxrefHook(boil.AfterDeleteHook, cvtermDbxrefAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	cvtermDbxrefAfterDeleteHooks = []CvtermDbxrefHook{}

	AddCvtermDbxrefHook(boil.BeforeUpsertHook, cvtermDbxrefBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	cvtermDbxrefBeforeUpsertHooks = []CvtermDbxrefHook{}

	AddCvtermDbxrefHook(boil.AfterUpsertHook, cvtermDbxrefAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	cvtermDbxrefAfterUpsertHooks = []CvtermDbxrefHook{}
}
func testCvtermDbxrefsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermDbxref := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCvtermDbxrefsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermDbxref := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxref.Insert(tx, cvtermDbxrefColumns...); err != nil {
		t.Error(err)
	}

	count, err := CvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCvtermDbxrefToOneCvtermUsingCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local CvtermDbxref
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvtermDbxrefDBTypes, true, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
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

	slice := CvtermDbxrefSlice{&local}
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

func testCvtermDbxrefToOneDbxrefUsingDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local CvtermDbxref
	var foreign Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvtermDbxrefDBTypes, true, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
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

	slice := CvtermDbxrefSlice{&local}
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

func testCvtermDbxrefToOneSetOpCvtermUsingCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a CvtermDbxref
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDbxrefDBTypes, false, strmangle.SetComplement(cvtermDbxrefPrimaryKeyColumns, cvtermDbxrefColumnsWithoutDefault)...); err != nil {
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

		if x.R.CvtermDbxref != &a {
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
func testCvtermDbxrefToOneSetOpDbxrefUsingDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a CvtermDbxref
	var b, c Dbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDbxrefDBTypes, false, strmangle.SetComplement(cvtermDbxrefPrimaryKeyColumns, cvtermDbxrefColumnsWithoutDefault)...); err != nil {
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

		if x.R.CvtermDbxref != &a {
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
func testCvtermDbxrefsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermDbxref := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cvtermDbxref.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCvtermDbxrefsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermDbxref := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CvtermDbxrefSlice{cvtermDbxref}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testCvtermDbxrefsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermDbxref := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := CvtermDbxrefs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	cvtermDbxrefDBTypes = map[string]string{"CvtermDbxrefID": "integer", "CvtermID": "integer", "DbxrefID": "integer", "IsForDefinition": "integer"}
	_                   = bytes.MinRead
)

func testCvtermDbxrefsUpdate(t *testing.T) {
	t.Parallel()

	if len(cvtermDbxrefColumns) == len(cvtermDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cvtermDbxref := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	if err = cvtermDbxref.Update(tx); err != nil {
		t.Error(err)
	}
}

func testCvtermDbxrefsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(cvtermDbxrefColumns) == len(cvtermDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cvtermDbxref := &CvtermDbxref{}
	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cvtermDbxref, cvtermDbxrefDBTypes, true, cvtermDbxrefPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(cvtermDbxrefColumns, cvtermDbxrefPrimaryKeyColumns) {
		fields = cvtermDbxrefColumns
	} else {
		fields = strmangle.SetComplement(
			cvtermDbxrefColumns,
			cvtermDbxrefPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(cvtermDbxref))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := CvtermDbxrefSlice{cvtermDbxref}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testCvtermDbxrefsUpsert(t *testing.T) {
	t.Parallel()

	if len(cvtermDbxrefColumns) == len(cvtermDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	cvtermDbxref := CvtermDbxref{}
	if err = randomize.Struct(seed, &cvtermDbxref, cvtermDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermDbxref.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert CvtermDbxref: %s", err)
	}

	count, err := CvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &cvtermDbxref, cvtermDbxrefDBTypes, false, cvtermDbxrefPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}

	if err = cvtermDbxref.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert CvtermDbxref: %s", err)
	}

	count, err = CvtermDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
