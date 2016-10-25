package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockDbxrefprops(t *testing.T) {
	t.Parallel()

	query := StockDbxrefprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockDbxrefpropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxrefprop := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockDbxrefprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockDbxrefpropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxrefprop := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockDbxrefprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockDbxrefpropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxrefprop := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockDbxrefpropSlice{stockDbxrefprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockDbxrefpropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxrefprop := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockDbxrefpropExists(tx, stockDbxrefprop.StockDbxrefpropID)
	if err != nil {
		t.Errorf("Unable to check if StockDbxrefprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockDbxrefpropExistsG to return true, but got false.")
	}
}
func testStockDbxrefpropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxrefprop := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	stockDbxrefpropFound, err := FindStockDbxrefprop(tx, stockDbxrefprop.StockDbxrefpropID)
	if err != nil {
		t.Error(err)
	}

	if stockDbxrefpropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockDbxrefpropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxrefprop := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockDbxrefprops(tx).Bind(stockDbxrefprop); err != nil {
		t.Error(err)
	}
}

func testStockDbxrefpropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxrefprop := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := StockDbxrefprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockDbxrefpropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxrefpropOne := &StockDbxrefprop{}
	stockDbxrefpropTwo := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefpropOne, stockDbxrefpropDBTypes, false, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}
	if err = randomize.Struct(seed, stockDbxrefpropTwo, stockDbxrefpropDBTypes, false, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockDbxrefpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockDbxrefprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockDbxrefpropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockDbxrefpropOne := &StockDbxrefprop{}
	stockDbxrefpropTwo := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefpropOne, stockDbxrefpropDBTypes, false, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}
	if err = randomize.Struct(seed, stockDbxrefpropTwo, stockDbxrefpropDBTypes, false, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockDbxrefpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockDbxrefpropBeforeInsertHook(e boil.Executor, o *StockDbxrefprop) error {
	*o = StockDbxrefprop{}
	return nil
}

func stockDbxrefpropAfterInsertHook(e boil.Executor, o *StockDbxrefprop) error {
	*o = StockDbxrefprop{}
	return nil
}

func stockDbxrefpropAfterSelectHook(e boil.Executor, o *StockDbxrefprop) error {
	*o = StockDbxrefprop{}
	return nil
}

func stockDbxrefpropBeforeUpdateHook(e boil.Executor, o *StockDbxrefprop) error {
	*o = StockDbxrefprop{}
	return nil
}

func stockDbxrefpropAfterUpdateHook(e boil.Executor, o *StockDbxrefprop) error {
	*o = StockDbxrefprop{}
	return nil
}

func stockDbxrefpropBeforeDeleteHook(e boil.Executor, o *StockDbxrefprop) error {
	*o = StockDbxrefprop{}
	return nil
}

func stockDbxrefpropAfterDeleteHook(e boil.Executor, o *StockDbxrefprop) error {
	*o = StockDbxrefprop{}
	return nil
}

func stockDbxrefpropBeforeUpsertHook(e boil.Executor, o *StockDbxrefprop) error {
	*o = StockDbxrefprop{}
	return nil
}

func stockDbxrefpropAfterUpsertHook(e boil.Executor, o *StockDbxrefprop) error {
	*o = StockDbxrefprop{}
	return nil
}

func testStockDbxrefpropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &StockDbxrefprop{}
	o := &StockDbxrefprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockDbxrefpropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop object: %s", err)
	}

	AddStockDbxrefpropHook(boil.BeforeInsertHook, stockDbxrefpropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockDbxrefpropBeforeInsertHooks = []StockDbxrefpropHook{}

	AddStockDbxrefpropHook(boil.AfterInsertHook, stockDbxrefpropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockDbxrefpropAfterInsertHooks = []StockDbxrefpropHook{}

	AddStockDbxrefpropHook(boil.AfterSelectHook, stockDbxrefpropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockDbxrefpropAfterSelectHooks = []StockDbxrefpropHook{}

	AddStockDbxrefpropHook(boil.BeforeUpdateHook, stockDbxrefpropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockDbxrefpropBeforeUpdateHooks = []StockDbxrefpropHook{}

	AddStockDbxrefpropHook(boil.AfterUpdateHook, stockDbxrefpropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockDbxrefpropAfterUpdateHooks = []StockDbxrefpropHook{}

	AddStockDbxrefpropHook(boil.BeforeDeleteHook, stockDbxrefpropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockDbxrefpropBeforeDeleteHooks = []StockDbxrefpropHook{}

	AddStockDbxrefpropHook(boil.AfterDeleteHook, stockDbxrefpropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockDbxrefpropAfterDeleteHooks = []StockDbxrefpropHook{}

	AddStockDbxrefpropHook(boil.BeforeUpsertHook, stockDbxrefpropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockDbxrefpropBeforeUpsertHooks = []StockDbxrefpropHook{}

	AddStockDbxrefpropHook(boil.AfterUpsertHook, stockDbxrefpropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockDbxrefpropAfterUpsertHooks = []StockDbxrefpropHook{}
}
func testStockDbxrefpropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxrefprop := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockDbxrefpropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxrefprop := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefprop.Insert(tx, stockDbxrefpropColumns...); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockDbxrefpropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockDbxrefprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockDbxrefpropDBTypes, true, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
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

	slice := StockDbxrefpropSlice{&local}
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

func testStockDbxrefpropToOneStockDbxrefUsingStockDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockDbxrefprop
	var foreign StockDbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockDbxrefpropDBTypes, true, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, stockDbxrefDBTypes, true, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.StockDbxrefID = foreign.StockDbxrefID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockDbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockDbxrefID != foreign.StockDbxrefID {
		t.Errorf("want: %v, got %v", foreign.StockDbxrefID, check.StockDbxrefID)
	}

	slice := StockDbxrefpropSlice{&local}
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

func testStockDbxrefpropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockDbxrefprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDbxrefpropDBTypes, false, strmangle.SetComplement(stockDbxrefpropPrimaryKeyColumns, stockDbxrefpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeStockDbxrefprop != &a {
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
func testStockDbxrefpropToOneSetOpStockDbxrefUsingStockDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockDbxrefprop
	var b, c StockDbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDbxrefpropDBTypes, false, strmangle.SetComplement(stockDbxrefpropPrimaryKeyColumns, stockDbxrefpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.StockDbxrefprop != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StockDbxrefID != x.StockDbxrefID {
			t.Error("foreign key was wrong value", a.StockDbxrefID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.StockDbxrefID))
		reflect.Indirect(reflect.ValueOf(&a.StockDbxrefID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockDbxrefID != x.StockDbxrefID {
			t.Error("foreign key was wrong value", a.StockDbxrefID, x.StockDbxrefID)
		}
	}
}
func testStockDbxrefpropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxrefprop := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockDbxrefprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockDbxrefpropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxrefprop := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockDbxrefpropSlice{stockDbxrefprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockDbxrefpropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockDbxrefprop := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockDbxrefprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockDbxrefpropDBTypes = map[string]string{"Rank": "integer", "StockDbxrefID": "integer", "StockDbxrefpropID": "integer", "TypeID": "integer", "Value": "text"}
	_                      = bytes.MinRead
)

func testStockDbxrefpropsUpdate(t *testing.T) {
	t.Parallel()

	if len(stockDbxrefpropColumns) == len(stockDbxrefpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockDbxrefprop := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	if err = stockDbxrefprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockDbxrefpropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockDbxrefpropColumns) == len(stockDbxrefpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockDbxrefprop := &StockDbxrefprop{}
	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockDbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockDbxrefprop, stockDbxrefpropDBTypes, true, stockDbxrefpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockDbxrefpropColumns, stockDbxrefpropPrimaryKeyColumns) {
		fields = stockDbxrefpropColumns
	} else {
		fields = strmangle.SetComplement(
			stockDbxrefpropColumns,
			stockDbxrefpropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockDbxrefprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockDbxrefpropSlice{stockDbxrefprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockDbxrefpropsUpsert(t *testing.T) {
	t.Parallel()

	if len(stockDbxrefpropColumns) == len(stockDbxrefpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockDbxrefprop := StockDbxrefprop{}
	if err = randomize.Struct(seed, &stockDbxrefprop, stockDbxrefpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockDbxrefprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockDbxrefprop: %s", err)
	}

	count, err := StockDbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockDbxrefprop, stockDbxrefpropDBTypes, false, stockDbxrefpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}

	if err = stockDbxrefprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockDbxrefprop: %s", err)
	}

	count, err = StockDbxrefprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
