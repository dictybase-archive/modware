package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDBS(t *testing.T) {
	t.Parallel()

	query := DBS(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDBSDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	db := &DB{}
	if err = randomize.Struct(seed, db, dbDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = db.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = db.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDBSQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	db := &DB{}
	if err = randomize.Struct(seed, db, dbDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = db.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DBS(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDBSSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	db := &DB{}
	if err = randomize.Struct(seed, db, dbDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = db.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DBSlice{db}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDBSExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	db := &DB{}
	if err = randomize.Struct(seed, db, dbDBTypes, true, dbColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = db.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DBExists(tx, db.DBID)
	if err != nil {
		t.Errorf("Unable to check if DB exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DBExistsG to return true, but got false.")
	}
}
func testDBSFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	db := &DB{}
	if err = randomize.Struct(seed, db, dbDBTypes, true, dbColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = db.Insert(tx); err != nil {
		t.Error(err)
	}

	dbFound, err := FindDB(tx, db.DBID)
	if err != nil {
		t.Error(err)
	}

	if dbFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDBSBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	db := &DB{}
	if err = randomize.Struct(seed, db, dbDBTypes, true, dbColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = db.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DBS(tx).Bind(db); err != nil {
		t.Error(err)
	}
}

func testDBSOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	db := &DB{}
	if err = randomize.Struct(seed, db, dbDBTypes, true, dbColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = db.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DBS(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDBSAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbOne := &DB{}
	dbTwo := &DB{}
	if err = randomize.Struct(seed, dbOne, dbDBTypes, false, dbColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}
	if err = randomize.Struct(seed, dbTwo, dbDBTypes, false, dbColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dbTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DBS(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDBSCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dbOne := &DB{}
	dbTwo := &DB{}
	if err = randomize.Struct(seed, dbOne, dbDBTypes, false, dbColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}
	if err = randomize.Struct(seed, dbTwo, dbDBTypes, false, dbColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dbTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func dbBeforeInsertHook(e boil.Executor, o *DB) error {
	*o = DB{}
	return nil
}

func dbAfterInsertHook(e boil.Executor, o *DB) error {
	*o = DB{}
	return nil
}

func dbAfterSelectHook(e boil.Executor, o *DB) error {
	*o = DB{}
	return nil
}

func dbBeforeUpdateHook(e boil.Executor, o *DB) error {
	*o = DB{}
	return nil
}

func dbAfterUpdateHook(e boil.Executor, o *DB) error {
	*o = DB{}
	return nil
}

func dbBeforeDeleteHook(e boil.Executor, o *DB) error {
	*o = DB{}
	return nil
}

func dbAfterDeleteHook(e boil.Executor, o *DB) error {
	*o = DB{}
	return nil
}

func dbBeforeUpsertHook(e boil.Executor, o *DB) error {
	*o = DB{}
	return nil
}

func dbAfterUpsertHook(e boil.Executor, o *DB) error {
	*o = DB{}
	return nil
}

func testDBSHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &DB{}
	o := &DB{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dbDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DB object: %s", err)
	}

	AddDBHook(boil.BeforeInsertHook, dbBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dbBeforeInsertHooks = []DBHook{}

	AddDBHook(boil.AfterInsertHook, dbAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dbAfterInsertHooks = []DBHook{}

	AddDBHook(boil.AfterSelectHook, dbAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dbAfterSelectHooks = []DBHook{}

	AddDBHook(boil.BeforeUpdateHook, dbBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dbBeforeUpdateHooks = []DBHook{}

	AddDBHook(boil.AfterUpdateHook, dbAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dbAfterUpdateHooks = []DBHook{}

	AddDBHook(boil.BeforeDeleteHook, dbBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dbBeforeDeleteHooks = []DBHook{}

	AddDBHook(boil.AfterDeleteHook, dbAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dbAfterDeleteHooks = []DBHook{}

	AddDBHook(boil.BeforeUpsertHook, dbBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dbBeforeUpsertHooks = []DBHook{}

	AddDBHook(boil.AfterUpsertHook, dbAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dbAfterUpsertHooks = []DBHook{}
}
func testDBSInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	db := &DB{}
	if err = randomize.Struct(seed, db, dbDBTypes, true, dbColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = db.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDBSInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	db := &DB{}
	if err = randomize.Struct(seed, db, dbDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = db.Insert(tx, dbColumns...); err != nil {
		t.Error(err)
	}

	count, err := DBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDBOneToOneDbxrefUsingDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Dbxref
	var local DB

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, dbDBTypes, true, dbColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.DBID = local.DBID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Dbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.DBID != foreign.DBID {
		t.Errorf("want: %v, got %v", foreign.DBID, check.DBID)
	}

	slice := DBSlice{&local}
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

func testDBOneToOneSetOpDbxrefUsingDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DB
	var b, c Dbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbDBTypes, false, strmangle.SetComplement(dbPrimaryKeyColumns, dbColumnsWithoutDefault)...); err != nil {
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
		if x.R.DB != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.DBID != x.DBID {
			t.Error("foreign key was wrong value", a.DBID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.DBID))
		reflect.Indirect(reflect.ValueOf(&x.DBID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DBID != x.DBID {
			t.Error("foreign key was wrong value", a.DBID, x.DBID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testDBSReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	db := &DB{}
	if err = randomize.Struct(seed, db, dbDBTypes, true, dbColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = db.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = db.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDBSReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	db := &DB{}
	if err = randomize.Struct(seed, db, dbDBTypes, true, dbColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = db.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DBSlice{db}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDBSSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	db := &DB{}
	if err = randomize.Struct(seed, db, dbDBTypes, true, dbColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = db.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DBS(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dbDBTypes = map[string]string{"DBID": "integer", "Description": "character varying", "Name": "character varying", "Url": "character varying", "Urlprefix": "character varying"}
	_         = bytes.MinRead
)

func testDBSUpdate(t *testing.T) {
	t.Parallel()

	if len(dbColumns) == len(dbPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	db := &DB{}
	if err = randomize.Struct(seed, db, dbDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = db.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, db, dbDBTypes, true, dbColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	if err = db.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDBSSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dbColumns) == len(dbPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	db := &DB{}
	if err = randomize.Struct(seed, db, dbDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = db.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, db, dbDBTypes, true, dbPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dbColumns, dbPrimaryKeyColumns) {
		fields = dbColumns
	} else {
		fields = strmangle.SetComplement(
			dbColumns,
			dbPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(db))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DBSlice{db}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDBSUpsert(t *testing.T) {
	t.Parallel()

	if len(dbColumns) == len(dbPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	db := DB{}
	if err = randomize.Struct(seed, &db, dbDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = db.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DB: %s", err)
	}

	count, err := DBS(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &db, dbDBTypes, false, dbPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DB struct: %s", err)
	}

	if err = db.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DB: %s", err)
	}

	count, err = DBS(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
