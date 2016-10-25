package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testPubauthors(t *testing.T) {
	t.Parallel()

	query := Pubauthors(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testPubauthorsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubauthor := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthor.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = pubauthor.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubauthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPubauthorsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubauthor := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthor.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Pubauthors(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Pubauthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPubauthorsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubauthor := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthor.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PubauthorSlice{pubauthor}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubauthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testPubauthorsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubauthor := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true, pubauthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthor.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := PubauthorExists(tx, pubauthor.PubauthorID)
	if err != nil {
		t.Errorf("Unable to check if Pubauthor exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PubauthorExistsG to return true, but got false.")
	}
}
func testPubauthorsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubauthor := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true, pubauthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthor.Insert(tx); err != nil {
		t.Error(err)
	}

	pubauthorFound, err := FindPubauthor(tx, pubauthor.PubauthorID)
	if err != nil {
		t.Error(err)
	}

	if pubauthorFound == nil {
		t.Error("want a record, got nil")
	}
}
func testPubauthorsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubauthor := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true, pubauthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthor.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Pubauthors(tx).Bind(pubauthor); err != nil {
		t.Error(err)
	}
}

func testPubauthorsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubauthor := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true, pubauthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthor.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Pubauthors(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPubauthorsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubauthorOne := &Pubauthor{}
	pubauthorTwo := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthorOne, pubauthorDBTypes, false, pubauthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}
	if err = randomize.Struct(seed, pubauthorTwo, pubauthorDBTypes, false, pubauthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthorOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = pubauthorTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Pubauthors(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPubauthorsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	pubauthorOne := &Pubauthor{}
	pubauthorTwo := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthorOne, pubauthorDBTypes, false, pubauthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}
	if err = randomize.Struct(seed, pubauthorTwo, pubauthorDBTypes, false, pubauthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthorOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = pubauthorTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubauthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func pubauthorBeforeInsertHook(e boil.Executor, o *Pubauthor) error {
	*o = Pubauthor{}
	return nil
}

func pubauthorAfterInsertHook(e boil.Executor, o *Pubauthor) error {
	*o = Pubauthor{}
	return nil
}

func pubauthorAfterSelectHook(e boil.Executor, o *Pubauthor) error {
	*o = Pubauthor{}
	return nil
}

func pubauthorBeforeUpdateHook(e boil.Executor, o *Pubauthor) error {
	*o = Pubauthor{}
	return nil
}

func pubauthorAfterUpdateHook(e boil.Executor, o *Pubauthor) error {
	*o = Pubauthor{}
	return nil
}

func pubauthorBeforeDeleteHook(e boil.Executor, o *Pubauthor) error {
	*o = Pubauthor{}
	return nil
}

func pubauthorAfterDeleteHook(e boil.Executor, o *Pubauthor) error {
	*o = Pubauthor{}
	return nil
}

func pubauthorBeforeUpsertHook(e boil.Executor, o *Pubauthor) error {
	*o = Pubauthor{}
	return nil
}

func pubauthorAfterUpsertHook(e boil.Executor, o *Pubauthor) error {
	*o = Pubauthor{}
	return nil
}

func testPubauthorsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Pubauthor{}
	o := &Pubauthor{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, pubauthorDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Pubauthor object: %s", err)
	}

	AddPubauthorHook(boil.BeforeInsertHook, pubauthorBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	pubauthorBeforeInsertHooks = []PubauthorHook{}

	AddPubauthorHook(boil.AfterInsertHook, pubauthorAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	pubauthorAfterInsertHooks = []PubauthorHook{}

	AddPubauthorHook(boil.AfterSelectHook, pubauthorAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	pubauthorAfterSelectHooks = []PubauthorHook{}

	AddPubauthorHook(boil.BeforeUpdateHook, pubauthorBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	pubauthorBeforeUpdateHooks = []PubauthorHook{}

	AddPubauthorHook(boil.AfterUpdateHook, pubauthorAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	pubauthorAfterUpdateHooks = []PubauthorHook{}

	AddPubauthorHook(boil.BeforeDeleteHook, pubauthorBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	pubauthorBeforeDeleteHooks = []PubauthorHook{}

	AddPubauthorHook(boil.AfterDeleteHook, pubauthorAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	pubauthorAfterDeleteHooks = []PubauthorHook{}

	AddPubauthorHook(boil.BeforeUpsertHook, pubauthorBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	pubauthorBeforeUpsertHooks = []PubauthorHook{}

	AddPubauthorHook(boil.AfterUpsertHook, pubauthorAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	pubauthorAfterUpsertHooks = []PubauthorHook{}
}
func testPubauthorsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubauthor := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true, pubauthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthor.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubauthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPubauthorsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubauthor := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthor.Insert(tx, pubauthorColumns...); err != nil {
		t.Error(err)
	}

	count, err := Pubauthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPubauthorToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Pubauthor
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, pubauthorDBTypes, true, pubauthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
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

	slice := PubauthorSlice{&local}
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

func testPubauthorToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pubauthor
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubauthorDBTypes, false, strmangle.SetComplement(pubauthorPrimaryKeyColumns, pubauthorColumnsWithoutDefault)...); err != nil {
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

		if x.R.Pubauthor != &a {
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
func testPubauthorsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubauthor := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true, pubauthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthor.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = pubauthor.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPubauthorsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubauthor := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true, pubauthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthor.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PubauthorSlice{pubauthor}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testPubauthorsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubauthor := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true, pubauthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthor.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Pubauthors(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	pubauthorDBTypes = map[string]string{"Editor": "boolean", "Givennames": "character varying", "PubID": "integer", "PubauthorID": "integer", "Rank": "integer", "Suffix": "character varying", "Surname": "character varying"}
	_                = bytes.MinRead
)

func testPubauthorsUpdate(t *testing.T) {
	t.Parallel()

	if len(pubauthorColumns) == len(pubauthorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	pubauthor := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthor.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubauthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true, pubauthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	if err = pubauthor.Update(tx); err != nil {
		t.Error(err)
	}
}

func testPubauthorsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(pubauthorColumns) == len(pubauthorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	pubauthor := &Pubauthor{}
	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthor.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubauthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, pubauthor, pubauthorDBTypes, true, pubauthorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(pubauthorColumns, pubauthorPrimaryKeyColumns) {
		fields = pubauthorColumns
	} else {
		fields = strmangle.SetComplement(
			pubauthorColumns,
			pubauthorPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(pubauthor))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := PubauthorSlice{pubauthor}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testPubauthorsUpsert(t *testing.T) {
	t.Parallel()

	if len(pubauthorColumns) == len(pubauthorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	pubauthor := Pubauthor{}
	if err = randomize.Struct(seed, &pubauthor, pubauthorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubauthor.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Pubauthor: %s", err)
	}

	count, err := Pubauthors(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &pubauthor, pubauthorDBTypes, false, pubauthorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}

	if err = pubauthor.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Pubauthor: %s", err)
	}

	count, err = Pubauthors(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
