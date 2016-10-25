package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testPubDbxrefs(t *testing.T) {
	t.Parallel()

	query := PubDbxrefs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testPubDbxrefsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubDbxref := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = pubDbxref.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := PubDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPubDbxrefsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubDbxref := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = PubDbxrefs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := PubDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPubDbxrefsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubDbxref := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PubDbxrefSlice{pubDbxref}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := PubDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testPubDbxrefsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubDbxref := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := PubDbxrefExists(tx, pubDbxref.PubDbxrefID)
	if err != nil {
		t.Errorf("Unable to check if PubDbxref exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PubDbxrefExistsG to return true, but got false.")
	}
}
func testPubDbxrefsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubDbxref := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	pubDbxrefFound, err := FindPubDbxref(tx, pubDbxref.PubDbxrefID)
	if err != nil {
		t.Error(err)
	}

	if pubDbxrefFound == nil {
		t.Error("want a record, got nil")
	}
}
func testPubDbxrefsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubDbxref := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = PubDbxrefs(tx).Bind(pubDbxref); err != nil {
		t.Error(err)
	}
}

func testPubDbxrefsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubDbxref := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := PubDbxrefs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPubDbxrefsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubDbxrefOne := &PubDbxref{}
	pubDbxrefTwo := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxrefOne, pubDbxrefDBTypes, false, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}
	if err = randomize.Struct(seed, pubDbxrefTwo, pubDbxrefDBTypes, false, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxrefOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = pubDbxrefTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := PubDbxrefs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPubDbxrefsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	pubDbxrefOne := &PubDbxref{}
	pubDbxrefTwo := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxrefOne, pubDbxrefDBTypes, false, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}
	if err = randomize.Struct(seed, pubDbxrefTwo, pubDbxrefDBTypes, false, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxrefOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = pubDbxrefTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PubDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func pubDbxrefBeforeInsertHook(e boil.Executor, o *PubDbxref) error {
	*o = PubDbxref{}
	return nil
}

func pubDbxrefAfterInsertHook(e boil.Executor, o *PubDbxref) error {
	*o = PubDbxref{}
	return nil
}

func pubDbxrefAfterSelectHook(e boil.Executor, o *PubDbxref) error {
	*o = PubDbxref{}
	return nil
}

func pubDbxrefBeforeUpdateHook(e boil.Executor, o *PubDbxref) error {
	*o = PubDbxref{}
	return nil
}

func pubDbxrefAfterUpdateHook(e boil.Executor, o *PubDbxref) error {
	*o = PubDbxref{}
	return nil
}

func pubDbxrefBeforeDeleteHook(e boil.Executor, o *PubDbxref) error {
	*o = PubDbxref{}
	return nil
}

func pubDbxrefAfterDeleteHook(e boil.Executor, o *PubDbxref) error {
	*o = PubDbxref{}
	return nil
}

func pubDbxrefBeforeUpsertHook(e boil.Executor, o *PubDbxref) error {
	*o = PubDbxref{}
	return nil
}

func pubDbxrefAfterUpsertHook(e boil.Executor, o *PubDbxref) error {
	*o = PubDbxref{}
	return nil
}

func testPubDbxrefsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &PubDbxref{}
	o := &PubDbxref{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, pubDbxrefDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PubDbxref object: %s", err)
	}

	AddPubDbxrefHook(boil.BeforeInsertHook, pubDbxrefBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	pubDbxrefBeforeInsertHooks = []PubDbxrefHook{}

	AddPubDbxrefHook(boil.AfterInsertHook, pubDbxrefAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	pubDbxrefAfterInsertHooks = []PubDbxrefHook{}

	AddPubDbxrefHook(boil.AfterSelectHook, pubDbxrefAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	pubDbxrefAfterSelectHooks = []PubDbxrefHook{}

	AddPubDbxrefHook(boil.BeforeUpdateHook, pubDbxrefBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	pubDbxrefBeforeUpdateHooks = []PubDbxrefHook{}

	AddPubDbxrefHook(boil.AfterUpdateHook, pubDbxrefAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	pubDbxrefAfterUpdateHooks = []PubDbxrefHook{}

	AddPubDbxrefHook(boil.BeforeDeleteHook, pubDbxrefBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	pubDbxrefBeforeDeleteHooks = []PubDbxrefHook{}

	AddPubDbxrefHook(boil.AfterDeleteHook, pubDbxrefAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	pubDbxrefAfterDeleteHooks = []PubDbxrefHook{}

	AddPubDbxrefHook(boil.BeforeUpsertHook, pubDbxrefBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	pubDbxrefBeforeUpsertHooks = []PubDbxrefHook{}

	AddPubDbxrefHook(boil.AfterUpsertHook, pubDbxrefAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	pubDbxrefAfterUpsertHooks = []PubDbxrefHook{}
}
func testPubDbxrefsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubDbxref := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PubDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPubDbxrefsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubDbxref := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxref.Insert(tx, pubDbxrefColumns...); err != nil {
		t.Error(err)
	}

	count, err := PubDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPubDbxrefToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PubDbxref
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, pubDbxrefDBTypes, true, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
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

	slice := PubDbxrefSlice{&local}
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

func testPubDbxrefToOneDbxrefUsingDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PubDbxref
	var foreign Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, pubDbxrefDBTypes, true, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
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

	slice := PubDbxrefSlice{&local}
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

func testPubDbxrefToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PubDbxref
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDbxrefDBTypes, false, strmangle.SetComplement(pubDbxrefPrimaryKeyColumns, pubDbxrefColumnsWithoutDefault)...); err != nil {
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

		if x.R.PubDbxref != &a {
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
func testPubDbxrefToOneSetOpDbxrefUsingDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PubDbxref
	var b, c Dbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDbxrefDBTypes, false, strmangle.SetComplement(pubDbxrefPrimaryKeyColumns, pubDbxrefColumnsWithoutDefault)...); err != nil {
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

		if x.R.PubDbxref != &a {
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
func testPubDbxrefsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubDbxref := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = pubDbxref.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPubDbxrefsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubDbxref := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PubDbxrefSlice{pubDbxref}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testPubDbxrefsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubDbxref := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := PubDbxrefs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	pubDbxrefDBTypes = map[string]string{"DbxrefID": "integer", "IsCurrent": "boolean", "PubDbxrefID": "integer", "PubID": "integer"}
	_                = bytes.MinRead
)

func testPubDbxrefsUpdate(t *testing.T) {
	t.Parallel()

	if len(pubDbxrefColumns) == len(pubDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	pubDbxref := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PubDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	if err = pubDbxref.Update(tx); err != nil {
		t.Error(err)
	}
}

func testPubDbxrefsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(pubDbxrefColumns) == len(pubDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	pubDbxref := &PubDbxref{}
	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PubDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, pubDbxref, pubDbxrefDBTypes, true, pubDbxrefPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(pubDbxrefColumns, pubDbxrefPrimaryKeyColumns) {
		fields = pubDbxrefColumns
	} else {
		fields = strmangle.SetComplement(
			pubDbxrefColumns,
			pubDbxrefPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(pubDbxref))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := PubDbxrefSlice{pubDbxref}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testPubDbxrefsUpsert(t *testing.T) {
	t.Parallel()

	if len(pubDbxrefColumns) == len(pubDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	pubDbxref := PubDbxref{}
	if err = randomize.Struct(seed, &pubDbxref, pubDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubDbxref.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert PubDbxref: %s", err)
	}

	count, err := PubDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &pubDbxref, pubDbxrefDBTypes, false, pubDbxrefPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}

	if err = pubDbxref.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert PubDbxref: %s", err)
	}

	count, err = PubDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
