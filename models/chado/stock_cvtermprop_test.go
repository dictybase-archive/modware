package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockCvtermprops(t *testing.T) {
	t.Parallel()

	query := StockCvtermprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockCvtermpropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvtermprop := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockCvtermprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := StockCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockCvtermpropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvtermprop := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockCvtermprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := StockCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockCvtermpropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvtermprop := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockCvtermpropSlice{stockCvtermprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := StockCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockCvtermpropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvtermprop := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockCvtermpropExists(tx, stockCvtermprop.StockCvtermpropID)
	if err != nil {
		t.Errorf("Unable to check if StockCvtermprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockCvtermpropExistsG to return true, but got false.")
	}
}
func testStockCvtermpropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvtermprop := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	stockCvtermpropFound, err := FindStockCvtermprop(tx, stockCvtermprop.StockCvtermpropID)
	if err != nil {
		t.Error(err)
	}

	if stockCvtermpropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockCvtermpropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvtermprop := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockCvtermprops(tx).Bind(stockCvtermprop); err != nil {
		t.Error(err)
	}
}

func testStockCvtermpropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvtermprop := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := StockCvtermprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockCvtermpropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvtermpropOne := &StockCvtermprop{}
	stockCvtermpropTwo := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermpropOne, stockCvtermpropDBTypes, false, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}
	if err = randomize.Struct(seed, stockCvtermpropTwo, stockCvtermpropDBTypes, false, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockCvtermpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockCvtermprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockCvtermpropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockCvtermpropOne := &StockCvtermprop{}
	stockCvtermpropTwo := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermpropOne, stockCvtermpropDBTypes, false, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}
	if err = randomize.Struct(seed, stockCvtermpropTwo, stockCvtermpropDBTypes, false, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockCvtermpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockCvtermpropBeforeInsertHook(e boil.Executor, o *StockCvtermprop) error {
	*o = StockCvtermprop{}
	return nil
}

func stockCvtermpropAfterInsertHook(e boil.Executor, o *StockCvtermprop) error {
	*o = StockCvtermprop{}
	return nil
}

func stockCvtermpropAfterSelectHook(e boil.Executor, o *StockCvtermprop) error {
	*o = StockCvtermprop{}
	return nil
}

func stockCvtermpropBeforeUpdateHook(e boil.Executor, o *StockCvtermprop) error {
	*o = StockCvtermprop{}
	return nil
}

func stockCvtermpropAfterUpdateHook(e boil.Executor, o *StockCvtermprop) error {
	*o = StockCvtermprop{}
	return nil
}

func stockCvtermpropBeforeDeleteHook(e boil.Executor, o *StockCvtermprop) error {
	*o = StockCvtermprop{}
	return nil
}

func stockCvtermpropAfterDeleteHook(e boil.Executor, o *StockCvtermprop) error {
	*o = StockCvtermprop{}
	return nil
}

func stockCvtermpropBeforeUpsertHook(e boil.Executor, o *StockCvtermprop) error {
	*o = StockCvtermprop{}
	return nil
}

func stockCvtermpropAfterUpsertHook(e boil.Executor, o *StockCvtermprop) error {
	*o = StockCvtermprop{}
	return nil
}

func testStockCvtermpropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &StockCvtermprop{}
	o := &StockCvtermprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockCvtermpropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop object: %s", err)
	}

	AddStockCvtermpropHook(boil.BeforeInsertHook, stockCvtermpropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockCvtermpropBeforeInsertHooks = []StockCvtermpropHook{}

	AddStockCvtermpropHook(boil.AfterInsertHook, stockCvtermpropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockCvtermpropAfterInsertHooks = []StockCvtermpropHook{}

	AddStockCvtermpropHook(boil.AfterSelectHook, stockCvtermpropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockCvtermpropAfterSelectHooks = []StockCvtermpropHook{}

	AddStockCvtermpropHook(boil.BeforeUpdateHook, stockCvtermpropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockCvtermpropBeforeUpdateHooks = []StockCvtermpropHook{}

	AddStockCvtermpropHook(boil.AfterUpdateHook, stockCvtermpropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockCvtermpropAfterUpdateHooks = []StockCvtermpropHook{}

	AddStockCvtermpropHook(boil.BeforeDeleteHook, stockCvtermpropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockCvtermpropBeforeDeleteHooks = []StockCvtermpropHook{}

	AddStockCvtermpropHook(boil.AfterDeleteHook, stockCvtermpropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockCvtermpropAfterDeleteHooks = []StockCvtermpropHook{}

	AddStockCvtermpropHook(boil.BeforeUpsertHook, stockCvtermpropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockCvtermpropBeforeUpsertHooks = []StockCvtermpropHook{}

	AddStockCvtermpropHook(boil.AfterUpsertHook, stockCvtermpropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockCvtermpropAfterUpsertHooks = []StockCvtermpropHook{}
}
func testStockCvtermpropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvtermprop := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockCvtermpropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvtermprop := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermprop.Insert(tx, stockCvtermpropColumns...); err != nil {
		t.Error(err)
	}

	count, err := StockCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockCvtermpropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockCvtermprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockCvtermpropDBTypes, true, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
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

	slice := StockCvtermpropSlice{&local}
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

func testStockCvtermpropToOneStockCvtermUsingStockCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockCvtermprop
	var foreign StockCvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockCvtermpropDBTypes, true, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.StockCvtermID = foreign.StockCvtermID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockCvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockCvtermID != foreign.StockCvtermID {
		t.Errorf("want: %v, got %v", foreign.StockCvtermID, check.StockCvtermID)
	}

	slice := StockCvtermpropSlice{&local}
	if err = local.L.LoadStockCvterm(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockCvterm == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockCvterm = nil
	if err = local.L.LoadStockCvterm(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockCvterm == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockCvtermpropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockCvtermprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockCvtermpropDBTypes, false, strmangle.SetComplement(stockCvtermpropPrimaryKeyColumns, stockCvtermpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeStockCvtermprop != &a {
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
func testStockCvtermpropToOneSetOpStockCvtermUsingStockCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockCvtermprop
	var b, c StockCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockCvtermpropDBTypes, false, strmangle.SetComplement(stockCvtermpropPrimaryKeyColumns, stockCvtermpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockCvtermDBTypes, false, strmangle.SetComplement(stockCvtermPrimaryKeyColumns, stockCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockCvtermDBTypes, false, strmangle.SetComplement(stockCvtermPrimaryKeyColumns, stockCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockCvterm{&b, &c} {
		err = a.SetStockCvterm(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockCvterm != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.StockCvtermprop != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StockCvtermID != x.StockCvtermID {
			t.Error("foreign key was wrong value", a.StockCvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.StockCvtermID))
		reflect.Indirect(reflect.ValueOf(&a.StockCvtermID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockCvtermID != x.StockCvtermID {
			t.Error("foreign key was wrong value", a.StockCvtermID, x.StockCvtermID)
		}
	}
}
func testStockCvtermpropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvtermprop := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockCvtermprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockCvtermpropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvtermprop := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockCvtermpropSlice{stockCvtermprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockCvtermpropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvtermprop := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockCvtermprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockCvtermpropDBTypes = map[string]string{"Rank": "integer", "StockCvtermID": "integer", "StockCvtermpropID": "integer", "TypeID": "integer", "Value": "text"}
	_                      = bytes.MinRead
)

func testStockCvtermpropsUpdate(t *testing.T) {
	t.Parallel()

	if len(stockCvtermpropColumns) == len(stockCvtermpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockCvtermprop := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	if err = stockCvtermprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockCvtermpropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockCvtermpropColumns) == len(stockCvtermpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockCvtermprop := &StockCvtermprop{}
	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockCvtermprop, stockCvtermpropDBTypes, true, stockCvtermpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockCvtermpropColumns, stockCvtermpropPrimaryKeyColumns) {
		fields = stockCvtermpropColumns
	} else {
		fields = strmangle.SetComplement(
			stockCvtermpropColumns,
			stockCvtermpropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockCvtermprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockCvtermpropSlice{stockCvtermprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockCvtermpropsUpsert(t *testing.T) {
	t.Parallel()

	if len(stockCvtermpropColumns) == len(stockCvtermpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockCvtermprop := StockCvtermprop{}
	if err = randomize.Struct(seed, &stockCvtermprop, stockCvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockCvtermprop: %s", err)
	}

	count, err := StockCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockCvtermprop, stockCvtermpropDBTypes, false, stockCvtermpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}

	if err = stockCvtermprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockCvtermprop: %s", err)
	}

	count, err = StockCvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
