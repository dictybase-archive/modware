package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDbxrefprops(t *testing.T) {
	t.Parallel()

	query := Dbxrefprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDbxrefpropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxrefprop := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dbxrefprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDbxrefpropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxrefprop := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Dbxrefprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDbxrefpropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxrefprop := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DbxrefpropSlice{dbxrefprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDbxrefpropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxrefprop := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DbxrefpropExists(tx, dbxrefprop.DbxrefpropID)
	if err != nil {
		t.Errorf("Unable to check if Dbxrefprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DbxrefpropExistsG to return true, but got false.")
	}
}
func testDbxrefpropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxrefprop := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	dbxrefpropFound, err := FindDbxrefprop(tx, dbxrefprop.DbxrefpropID)
	if err != nil {
		t.Error(err)
	}

	if dbxrefpropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDbxrefpropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxrefprop := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Dbxrefprops(tx).Bind(dbxrefprop); err != nil {
		t.Error(err)
	}
}

func testDbxrefpropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxrefprop := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Dbxrefprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDbxrefpropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxrefpropOne := &Dbxrefprop{}
	dbxrefpropTwo := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefpropOne, dbxrefpropDBTypes, false, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}
	if err = randomize.Struct(seed, dbxrefpropTwo, dbxrefpropDBTypes, false, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dbxrefpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Dbxrefprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDbxrefpropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dbxrefpropOne := &Dbxrefprop{}
	dbxrefpropTwo := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefpropOne, dbxrefpropDBTypes, false, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}
	if err = randomize.Struct(seed, dbxrefpropTwo, dbxrefpropDBTypes, false, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dbxrefpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func dbxrefpropBeforeInsertHook(e boil.Executor, o *Dbxrefprop) error {
	*o = Dbxrefprop{}
	return nil
}

func dbxrefpropAfterInsertHook(e boil.Executor, o *Dbxrefprop) error {
	*o = Dbxrefprop{}
	return nil
}

func dbxrefpropAfterSelectHook(e boil.Executor, o *Dbxrefprop) error {
	*o = Dbxrefprop{}
	return nil
}

func dbxrefpropBeforeUpdateHook(e boil.Executor, o *Dbxrefprop) error {
	*o = Dbxrefprop{}
	return nil
}

func dbxrefpropAfterUpdateHook(e boil.Executor, o *Dbxrefprop) error {
	*o = Dbxrefprop{}
	return nil
}

func dbxrefpropBeforeDeleteHook(e boil.Executor, o *Dbxrefprop) error {
	*o = Dbxrefprop{}
	return nil
}

func dbxrefpropAfterDeleteHook(e boil.Executor, o *Dbxrefprop) error {
	*o = Dbxrefprop{}
	return nil
}

func dbxrefpropBeforeUpsertHook(e boil.Executor, o *Dbxrefprop) error {
	*o = Dbxrefprop{}
	return nil
}

func dbxrefpropAfterUpsertHook(e boil.Executor, o *Dbxrefprop) error {
	*o = Dbxrefprop{}
	return nil
}

func testDbxrefpropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Dbxrefprop{}
	o := &Dbxrefprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dbxrefpropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop object: %s", err)
	}

	AddDbxrefpropHook(boil.BeforeInsertHook, dbxrefpropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dbxrefpropBeforeInsertHooks = []DbxrefpropHook{}

	AddDbxrefpropHook(boil.AfterInsertHook, dbxrefpropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dbxrefpropAfterInsertHooks = []DbxrefpropHook{}

	AddDbxrefpropHook(boil.AfterSelectHook, dbxrefpropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dbxrefpropAfterSelectHooks = []DbxrefpropHook{}

	AddDbxrefpropHook(boil.BeforeUpdateHook, dbxrefpropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dbxrefpropBeforeUpdateHooks = []DbxrefpropHook{}

	AddDbxrefpropHook(boil.AfterUpdateHook, dbxrefpropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dbxrefpropAfterUpdateHooks = []DbxrefpropHook{}

	AddDbxrefpropHook(boil.BeforeDeleteHook, dbxrefpropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dbxrefpropBeforeDeleteHooks = []DbxrefpropHook{}

	AddDbxrefpropHook(boil.AfterDeleteHook, dbxrefpropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dbxrefpropAfterDeleteHooks = []DbxrefpropHook{}

	AddDbxrefpropHook(boil.BeforeUpsertHook, dbxrefpropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dbxrefpropBeforeUpsertHooks = []DbxrefpropHook{}

	AddDbxrefpropHook(boil.AfterUpsertHook, dbxrefpropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dbxrefpropAfterUpsertHooks = []DbxrefpropHook{}
}
func testDbxrefpropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxrefprop := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDbxrefpropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxrefprop := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefprop.Insert(tx, dbxrefpropColumns...); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDbxrefpropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Dbxrefprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dbxrefpropDBTypes, true, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
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

	slice := DbxrefpropSlice{&local}
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

func testDbxrefpropToOneDbxrefUsingDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Dbxrefprop
	var foreign Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dbxrefpropDBTypes, true, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
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

	slice := DbxrefpropSlice{&local}
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

func testDbxrefpropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxrefprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefpropDBTypes, false, strmangle.SetComplement(dbxrefpropPrimaryKeyColumns, dbxrefpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeDbxrefprop != &a {
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
func testDbxrefpropToOneSetOpDbxrefUsingDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Dbxrefprop
	var b, c Dbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dbxrefpropDBTypes, false, strmangle.SetComplement(dbxrefpropPrimaryKeyColumns, dbxrefpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.Dbxrefprop != &a {
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
func testDbxrefpropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxrefprop := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dbxrefprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDbxrefpropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxrefprop := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DbxrefpropSlice{dbxrefprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDbxrefpropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dbxrefprop := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Dbxrefprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dbxrefpropDBTypes = map[string]string{"DbxrefID": "integer", "DbxrefpropID": "integer", "Rank": "integer", "TypeID": "integer", "Value": "text"}
	_                 = bytes.MinRead
)

func testDbxrefpropsUpdate(t *testing.T) {
	t.Parallel()

	if len(dbxrefpropColumns) == len(dbxrefpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dbxrefprop := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	if err = dbxrefprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDbxrefpropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dbxrefpropColumns) == len(dbxrefpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dbxrefprop := &Dbxrefprop{}
	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Dbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dbxrefprop, dbxrefpropDBTypes, true, dbxrefpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dbxrefpropColumns, dbxrefpropPrimaryKeyColumns) {
		fields = dbxrefpropColumns
	} else {
		fields = strmangle.SetComplement(
			dbxrefpropColumns,
			dbxrefpropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(dbxrefprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DbxrefpropSlice{dbxrefprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDbxrefpropsUpsert(t *testing.T) {
	t.Parallel()

	if len(dbxrefpropColumns) == len(dbxrefpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	dbxrefprop := Dbxrefprop{}
	if err = randomize.Struct(seed, &dbxrefprop, dbxrefpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dbxrefprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Dbxrefprop: %s", err)
	}

	count, err := Dbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &dbxrefprop, dbxrefpropDBTypes, false, dbxrefpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}

	if err = dbxrefprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Dbxrefprop: %s", err)
	}

	count, err = Dbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
