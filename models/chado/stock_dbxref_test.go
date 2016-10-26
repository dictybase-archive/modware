package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockDbxrefs(t *testing.T) {
	t.Parallel()

	query := StockDbxrefs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockDbxrefsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxref := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockDbxref.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockDbxrefsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxref := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockDbxrefs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockDbxrefsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxref := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockDbxrefSlice{stockDbxref}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockDbxrefsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxref := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockDbxrefExists(tx, stockDbxref.StockDbxrefID)
	if err != nil {
		t.Errorf("Unable to check if StockDbxref exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockDbxrefExistsG to return true, but got false.")
	}
}
func testStockDbxrefsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxref := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	stockDbxrefFound, err := FindStockDbxref(tx, stockDbxref.StockDbxrefID)
	if err != nil {
		t.Error(err)
	}

	if stockDbxrefFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockDbxrefsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxref := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockDbxrefs(tx).Bind(stockDbxref); err != nil {
		t.Error(err)
	}
}

func testStockDbxrefsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxref := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := StockDbxrefs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockDbxrefsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxrefOne := &StockDbxref{}
	stockDbxrefTwo := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxrefOne, stockDbxrefDBTypes, false, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}
	if err = randomize.Struct(seed, stockDbxrefTwo, stockDbxrefDBTypes, false, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockDbxrefTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockDbxrefs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockDbxrefsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockDbxrefOne := &StockDbxref{}
	stockDbxrefTwo := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxrefOne, stockDbxrefDBTypes, false, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}
	if err = randomize.Struct(seed, stockDbxrefTwo, stockDbxrefDBTypes, false, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockDbxrefTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockDbxrefBeforeInsertHook(e boil.Executor, o *StockDbxref) error {
	*o = StockDbxref{}
	return nil
}

func stockDbxrefAfterInsertHook(e boil.Executor, o *StockDbxref) error {
	*o = StockDbxref{}
	return nil
}

func stockDbxrefAfterSelectHook(e boil.Executor, o *StockDbxref) error {
	*o = StockDbxref{}
	return nil
}

func stockDbxrefBeforeUpdateHook(e boil.Executor, o *StockDbxref) error {
	*o = StockDbxref{}
	return nil
}

func stockDbxrefAfterUpdateHook(e boil.Executor, o *StockDbxref) error {
	*o = StockDbxref{}
	return nil
}

func stockDbxrefBeforeDeleteHook(e boil.Executor, o *StockDbxref) error {
	*o = StockDbxref{}
	return nil
}

func stockDbxrefAfterDeleteHook(e boil.Executor, o *StockDbxref) error {
	*o = StockDbxref{}
	return nil
}

func stockDbxrefBeforeUpsertHook(e boil.Executor, o *StockDbxref) error {
	*o = StockDbxref{}
	return nil
}

func stockDbxrefAfterUpsertHook(e boil.Executor, o *StockDbxref) error {
	*o = StockDbxref{}
	return nil
}

func testStockDbxrefsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &StockDbxref{}
	o := &StockDbxref{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockDbxrefDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StockDbxref object: %s", err)
	}

	AddStockDbxrefHook(boil.BeforeInsertHook, stockDbxrefBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockDbxrefBeforeInsertHooks = []StockDbxrefHook{}

	AddStockDbxrefHook(boil.AfterInsertHook, stockDbxrefAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockDbxrefAfterInsertHooks = []StockDbxrefHook{}

	AddStockDbxrefHook(boil.AfterSelectHook, stockDbxrefAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockDbxrefAfterSelectHooks = []StockDbxrefHook{}

	AddStockDbxrefHook(boil.BeforeUpdateHook, stockDbxrefBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockDbxrefBeforeUpdateHooks = []StockDbxrefHook{}

	AddStockDbxrefHook(boil.AfterUpdateHook, stockDbxrefAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockDbxrefAfterUpdateHooks = []StockDbxrefHook{}

	AddStockDbxrefHook(boil.BeforeDeleteHook, stockDbxrefBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockDbxrefBeforeDeleteHooks = []StockDbxrefHook{}

	AddStockDbxrefHook(boil.AfterDeleteHook, stockDbxrefAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockDbxrefAfterDeleteHooks = []StockDbxrefHook{}

	AddStockDbxrefHook(boil.BeforeUpsertHook, stockDbxrefBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockDbxrefBeforeUpsertHooks = []StockDbxrefHook{}

	AddStockDbxrefHook(boil.AfterUpsertHook, stockDbxrefAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockDbxrefAfterUpsertHooks = []StockDbxrefHook{}
}
func testStockDbxrefsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxref := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockDbxrefsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxref := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxref.Insert(tx, stockDbxrefColumns...); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockDbxrefOneToOneStockDbxrefpropUsingStockDbxrefprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockDbxrefprop
	var local StockDbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockDbxrefpropDBTypes, true, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockDbxrefDBTypes, true, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.StockDbxrefID = local.StockDbxrefID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockDbxrefprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockDbxrefID != foreign.StockDbxrefID {
		t.Errorf("want: %v, got %v", foreign.StockDbxrefID, check.StockDbxrefID)
	}

	slice := StockDbxrefSlice{&local}
	if err = local.L.LoadStockDbxrefprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockDbxrefprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockDbxrefprop = nil
	if err = local.L.LoadStockDbxrefprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockDbxrefprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockDbxrefOneToOneSetOpStockDbxrefpropUsingStockDbxrefprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockDbxref
	var b, c StockDbxrefprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDbxrefDBTypes, false, strmangle.SetComplement(stockDbxrefPrimaryKeyColumns, stockDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockDbxrefpropDBTypes, false, strmangle.SetComplement(stockDbxrefpropPrimaryKeyColumns, stockDbxrefpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockDbxrefpropDBTypes, false, strmangle.SetComplement(stockDbxrefpropPrimaryKeyColumns, stockDbxrefpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockDbxrefprop{&b, &c} {
		err = a.SetStockDbxrefprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockDbxrefprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.StockDbxref != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockDbxrefID != x.StockDbxrefID {
			t.Error("foreign key was wrong value", a.StockDbxrefID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.StockDbxrefID))
		reflect.Indirect(reflect.ValueOf(&x.StockDbxrefID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockDbxrefID != x.StockDbxrefID {
			t.Error("foreign key was wrong value", a.StockDbxrefID, x.StockDbxrefID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testStockDbxrefToOneStockUsingStock(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockDbxref
	var foreign Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockDbxrefDBTypes, true, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.StockID = foreign.StockID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Stock(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockID != foreign.StockID {
		t.Errorf("want: %v, got %v", foreign.StockID, check.StockID)
	}

	slice := StockDbxrefSlice{&local}
	if err = local.L.LoadStock(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Stock == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Stock = nil
	if err = local.L.LoadStock(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Stock == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockDbxrefToOneDbxrefUsingDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockDbxref
	var foreign Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockDbxrefDBTypes, true, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
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

	slice := StockDbxrefSlice{&local}
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

func testStockDbxrefToOneSetOpStockUsingStock(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockDbxref
	var b, c Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDbxrefDBTypes, false, strmangle.SetComplement(stockDbxrefPrimaryKeyColumns, stockDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Stock{&b, &c} {
		err = a.SetStock(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Stock != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.StockDbxref != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.StockID))
		reflect.Indirect(reflect.ValueOf(&a.StockID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID, x.StockID)
		}
	}
}
func testStockDbxrefToOneSetOpDbxrefUsingDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockDbxref
	var b, c Dbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDbxrefDBTypes, false, strmangle.SetComplement(stockDbxrefPrimaryKeyColumns, stockDbxrefColumnsWithoutDefault)...); err != nil {
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

		if x.R.StockDbxref != &a {
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
func testStockDbxrefsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxref := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockDbxref.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockDbxrefsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxref := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockDbxrefSlice{stockDbxref}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockDbxrefsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxref := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockDbxrefs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockDbxrefDBTypes = map[string]string{"DbxrefID": "integer", "IsCurrent": "boolean", "StockDbxrefID": "integer", "StockID": "integer"}
	_                  = bytes.MinRead
)

func testStockDbxrefsUpdate(t *testing.T) {
	t.Parallel()

	if len(stockDbxrefColumns) == len(stockDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockDbxref := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	if err = stockDbxref.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockDbxrefsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockDbxrefColumns) == len(stockDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockDbxref := &StockDbxref{}
	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockDbxref, stockDbxrefDBTypes, true, stockDbxrefPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockDbxrefColumns, stockDbxrefPrimaryKeyColumns) {
		fields = stockDbxrefColumns
	} else {
		fields = strmangle.SetComplement(
			stockDbxrefColumns,
			stockDbxrefPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockDbxref))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockDbxrefSlice{stockDbxref}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockDbxrefsUpsert(t *testing.T) {
	t.Parallel()

	if len(stockDbxrefColumns) == len(stockDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockDbxref := StockDbxref{}
	if err = randomize.Struct(seed, &stockDbxref, stockDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxref.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockDbxref: %s", err)
	}

	count, err := StockDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockDbxref, stockDbxrefDBTypes, false, stockDbxrefPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	if err = stockDbxref.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockDbxref: %s", err)
	}

	count, err = StockDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
