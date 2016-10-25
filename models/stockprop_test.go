package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockprops(t *testing.T) {
	t.Parallel()

	query := Stockprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockpropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockprop := &Stockprop{}
	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockpropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockprop := &Stockprop{}
	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Stockprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Stockprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockpropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockprop := &Stockprop{}
	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockpropSlice{stockprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockpropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockprop := &Stockprop{}
	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockpropExists(tx, stockprop.StockpropID)
	if err != nil {
		t.Errorf("Unable to check if Stockprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockpropExistsG to return true, but got false.")
	}
}
func testStockpropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockprop := &Stockprop{}
	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockprop.Insert(tx); err != nil {
		t.Error(err)
	}

	stockpropFound, err := FindStockprop(tx, stockprop.StockpropID)
	if err != nil {
		t.Error(err)
	}

	if stockpropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockpropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockprop := &Stockprop{}
	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Stockprops(tx).Bind(stockprop); err != nil {
		t.Error(err)
	}
}

func testStockpropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockprop := &Stockprop{}
	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Stockprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockpropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockpropOne := &Stockprop{}
	stockpropTwo := &Stockprop{}
	if err = randomize.Struct(seed, stockpropOne, stockpropDBTypes, false, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}
	if err = randomize.Struct(seed, stockpropTwo, stockpropDBTypes, false, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Stockprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockpropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockpropOne := &Stockprop{}
	stockpropTwo := &Stockprop{}
	if err = randomize.Struct(seed, stockpropOne, stockpropDBTypes, false, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}
	if err = randomize.Struct(seed, stockpropTwo, stockpropDBTypes, false, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockpropBeforeInsertHook(e boil.Executor, o *Stockprop) error {
	*o = Stockprop{}
	return nil
}

func stockpropAfterInsertHook(e boil.Executor, o *Stockprop) error {
	*o = Stockprop{}
	return nil
}

func stockpropAfterSelectHook(e boil.Executor, o *Stockprop) error {
	*o = Stockprop{}
	return nil
}

func stockpropBeforeUpdateHook(e boil.Executor, o *Stockprop) error {
	*o = Stockprop{}
	return nil
}

func stockpropAfterUpdateHook(e boil.Executor, o *Stockprop) error {
	*o = Stockprop{}
	return nil
}

func stockpropBeforeDeleteHook(e boil.Executor, o *Stockprop) error {
	*o = Stockprop{}
	return nil
}

func stockpropAfterDeleteHook(e boil.Executor, o *Stockprop) error {
	*o = Stockprop{}
	return nil
}

func stockpropBeforeUpsertHook(e boil.Executor, o *Stockprop) error {
	*o = Stockprop{}
	return nil
}

func stockpropAfterUpsertHook(e boil.Executor, o *Stockprop) error {
	*o = Stockprop{}
	return nil
}

func testStockpropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Stockprop{}
	o := &Stockprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockpropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Stockprop object: %s", err)
	}

	AddStockpropHook(boil.BeforeInsertHook, stockpropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockpropBeforeInsertHooks = []StockpropHook{}

	AddStockpropHook(boil.AfterInsertHook, stockpropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockpropAfterInsertHooks = []StockpropHook{}

	AddStockpropHook(boil.AfterSelectHook, stockpropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockpropAfterSelectHooks = []StockpropHook{}

	AddStockpropHook(boil.BeforeUpdateHook, stockpropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockpropBeforeUpdateHooks = []StockpropHook{}

	AddStockpropHook(boil.AfterUpdateHook, stockpropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockpropAfterUpdateHooks = []StockpropHook{}

	AddStockpropHook(boil.BeforeDeleteHook, stockpropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockpropBeforeDeleteHooks = []StockpropHook{}

	AddStockpropHook(boil.AfterDeleteHook, stockpropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockpropAfterDeleteHooks = []StockpropHook{}

	AddStockpropHook(boil.BeforeUpsertHook, stockpropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockpropBeforeUpsertHooks = []StockpropHook{}

	AddStockpropHook(boil.AfterUpsertHook, stockpropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockpropAfterUpsertHooks = []StockpropHook{}
}
func testStockpropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockprop := &Stockprop{}
	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockpropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockprop := &Stockprop{}
	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockprop.Insert(tx, stockpropColumns...); err != nil {
		t.Error(err)
	}

	count, err := Stockprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockpropOneToOneStockpropPubUsingStockpropPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockpropPub
	var local Stockprop

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockpropPubDBTypes, true, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockpropDBTypes, true, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.StockpropID = local.StockpropID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockpropPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockpropID != foreign.StockpropID {
		t.Errorf("want: %v, got %v", foreign.StockpropID, check.StockpropID)
	}

	slice := StockpropSlice{&local}
	if err = local.L.LoadStockpropPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockpropPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockpropPub = nil
	if err = local.L.LoadStockpropPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockpropPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockpropOneToOneSetOpStockpropPubUsingStockpropPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stockprop
	var b, c StockpropPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockpropDBTypes, false, strmangle.SetComplement(stockpropPrimaryKeyColumns, stockpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockpropPubDBTypes, false, strmangle.SetComplement(stockpropPubPrimaryKeyColumns, stockpropPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockpropPubDBTypes, false, strmangle.SetComplement(stockpropPubPrimaryKeyColumns, stockpropPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockpropPub{&b, &c} {
		err = a.SetStockpropPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockpropPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Stockprop != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockpropID != x.StockpropID {
			t.Error("foreign key was wrong value", a.StockpropID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.StockpropID))
		reflect.Indirect(reflect.ValueOf(&x.StockpropID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockpropID != x.StockpropID {
			t.Error("foreign key was wrong value", a.StockpropID, x.StockpropID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testStockpropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Stockprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockpropDBTypes, true, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
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

	slice := StockpropSlice{&local}
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

func testStockpropToOneStockUsingStock(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Stockprop
	var foreign Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockpropDBTypes, true, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
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

	slice := StockpropSlice{&local}
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

func testStockpropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stockprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockpropDBTypes, false, strmangle.SetComplement(stockpropPrimaryKeyColumns, stockpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeStockprop != &a {
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
func testStockpropToOneSetOpStockUsingStock(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stockprop
	var b, c Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockpropDBTypes, false, strmangle.SetComplement(stockpropPrimaryKeyColumns, stockpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.Stockprop != &a {
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
func testStockpropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockprop := &Stockprop{}
	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockpropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockprop := &Stockprop{}
	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockpropSlice{stockprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockpropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockprop := &Stockprop{}
	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Stockprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockpropDBTypes = map[string]string{"Rank": "integer", "StockID": "integer", "StockpropID": "integer", "TypeID": "integer", "Value": "text"}
	_                = bytes.MinRead
)

func testStockpropsUpdate(t *testing.T) {
	t.Parallel()

	if len(stockpropColumns) == len(stockpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockprop := &Stockprop{}
	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	if err = stockprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockpropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockpropColumns) == len(stockpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockprop := &Stockprop{}
	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockprop, stockpropDBTypes, true, stockpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockpropColumns, stockpropPrimaryKeyColumns) {
		fields = stockpropColumns
	} else {
		fields = strmangle.SetComplement(
			stockpropColumns,
			stockpropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockpropSlice{stockprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockpropsUpsert(t *testing.T) {
	t.Parallel()

	if len(stockpropColumns) == len(stockpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockprop := Stockprop{}
	if err = randomize.Struct(seed, &stockprop, stockpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Stockprop: %s", err)
	}

	count, err := Stockprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockprop, stockpropDBTypes, false, stockpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	if err = stockprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Stockprop: %s", err)
	}

	count, err = Stockprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
