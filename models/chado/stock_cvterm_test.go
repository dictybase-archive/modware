package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockCvterms(t *testing.T) {
	t.Parallel()

	query := StockCvterms(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockCvtermsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvterm := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockCvterm.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := StockCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockCvtermsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvterm := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockCvterms(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := StockCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockCvtermsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvterm := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockCvtermSlice{stockCvterm}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := StockCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockCvtermsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvterm := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockCvtermExists(tx, stockCvterm.StockCvtermID)
	if err != nil {
		t.Errorf("Unable to check if StockCvterm exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockCvtermExistsG to return true, but got false.")
	}
}
func testStockCvtermsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvterm := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	stockCvtermFound, err := FindStockCvterm(tx, stockCvterm.StockCvtermID)
	if err != nil {
		t.Error(err)
	}

	if stockCvtermFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockCvtermsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvterm := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockCvterms(tx).Bind(stockCvterm); err != nil {
		t.Error(err)
	}
}

func testStockCvtermsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvterm := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := StockCvterms(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockCvtermsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvtermOne := &StockCvterm{}
	stockCvtermTwo := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvtermOne, stockCvtermDBTypes, false, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}
	if err = randomize.Struct(seed, stockCvtermTwo, stockCvtermDBTypes, false, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockCvtermTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockCvterms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockCvtermsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockCvtermOne := &StockCvterm{}
	stockCvtermTwo := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvtermOne, stockCvtermDBTypes, false, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}
	if err = randomize.Struct(seed, stockCvtermTwo, stockCvtermDBTypes, false, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvtermOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockCvtermTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockCvtermBeforeInsertHook(e boil.Executor, o *StockCvterm) error {
	*o = StockCvterm{}
	return nil
}

func stockCvtermAfterInsertHook(e boil.Executor, o *StockCvterm) error {
	*o = StockCvterm{}
	return nil
}

func stockCvtermAfterSelectHook(e boil.Executor, o *StockCvterm) error {
	*o = StockCvterm{}
	return nil
}

func stockCvtermBeforeUpdateHook(e boil.Executor, o *StockCvterm) error {
	*o = StockCvterm{}
	return nil
}

func stockCvtermAfterUpdateHook(e boil.Executor, o *StockCvterm) error {
	*o = StockCvterm{}
	return nil
}

func stockCvtermBeforeDeleteHook(e boil.Executor, o *StockCvterm) error {
	*o = StockCvterm{}
	return nil
}

func stockCvtermAfterDeleteHook(e boil.Executor, o *StockCvterm) error {
	*o = StockCvterm{}
	return nil
}

func stockCvtermBeforeUpsertHook(e boil.Executor, o *StockCvterm) error {
	*o = StockCvterm{}
	return nil
}

func stockCvtermAfterUpsertHook(e boil.Executor, o *StockCvterm) error {
	*o = StockCvterm{}
	return nil
}

func testStockCvtermsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &StockCvterm{}
	o := &StockCvterm{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockCvtermDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StockCvterm object: %s", err)
	}

	AddStockCvtermHook(boil.BeforeInsertHook, stockCvtermBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockCvtermBeforeInsertHooks = []StockCvtermHook{}

	AddStockCvtermHook(boil.AfterInsertHook, stockCvtermAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockCvtermAfterInsertHooks = []StockCvtermHook{}

	AddStockCvtermHook(boil.AfterSelectHook, stockCvtermAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockCvtermAfterSelectHooks = []StockCvtermHook{}

	AddStockCvtermHook(boil.BeforeUpdateHook, stockCvtermBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockCvtermBeforeUpdateHooks = []StockCvtermHook{}

	AddStockCvtermHook(boil.AfterUpdateHook, stockCvtermAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockCvtermAfterUpdateHooks = []StockCvtermHook{}

	AddStockCvtermHook(boil.BeforeDeleteHook, stockCvtermBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockCvtermBeforeDeleteHooks = []StockCvtermHook{}

	AddStockCvtermHook(boil.AfterDeleteHook, stockCvtermAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockCvtermAfterDeleteHooks = []StockCvtermHook{}

	AddStockCvtermHook(boil.BeforeUpsertHook, stockCvtermBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockCvtermBeforeUpsertHooks = []StockCvtermHook{}

	AddStockCvtermHook(boil.AfterUpsertHook, stockCvtermAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockCvtermAfterUpsertHooks = []StockCvtermHook{}
}
func testStockCvtermsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvterm := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockCvtermsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvterm := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvterm.Insert(tx, stockCvtermColumns...); err != nil {
		t.Error(err)
	}

	count, err := StockCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockCvtermOneToOneStockCvtermpropUsingStockCvtermprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockCvtermprop
	var local StockCvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockCvtermpropDBTypes, true, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.StockCvtermID = local.StockCvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockCvtermprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockCvtermID != foreign.StockCvtermID {
		t.Errorf("want: %v, got %v", foreign.StockCvtermID, check.StockCvtermID)
	}

	slice := StockCvtermSlice{&local}
	if err = local.L.LoadStockCvtermprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockCvtermprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockCvtermprop = nil
	if err = local.L.LoadStockCvtermprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockCvtermprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockCvtermOneToOneSetOpStockCvtermpropUsingStockCvtermprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockCvterm
	var b, c StockCvtermprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockCvtermDBTypes, false, strmangle.SetComplement(stockCvtermPrimaryKeyColumns, stockCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockCvtermpropDBTypes, false, strmangle.SetComplement(stockCvtermpropPrimaryKeyColumns, stockCvtermpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockCvtermpropDBTypes, false, strmangle.SetComplement(stockCvtermpropPrimaryKeyColumns, stockCvtermpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockCvtermprop{&b, &c} {
		err = a.SetStockCvtermprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockCvtermprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.StockCvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockCvtermID != x.StockCvtermID {
			t.Error("foreign key was wrong value", a.StockCvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.StockCvtermID))
		reflect.Indirect(reflect.ValueOf(&x.StockCvtermID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockCvtermID != x.StockCvtermID {
			t.Error("foreign key was wrong value", a.StockCvtermID, x.StockCvtermID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testStockCvtermToOneCvtermUsingCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockCvterm
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
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

	slice := StockCvtermSlice{&local}
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

func testStockCvtermToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockCvterm
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
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

	slice := StockCvtermSlice{&local}
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

func testStockCvtermToOneStockUsingStock(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockCvterm
	var foreign Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
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

	slice := StockCvtermSlice{&local}
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

func testStockCvtermToOneSetOpCvtermUsingCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockCvterm
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockCvtermDBTypes, false, strmangle.SetComplement(stockCvtermPrimaryKeyColumns, stockCvtermColumnsWithoutDefault)...); err != nil {
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

		if x.R.StockCvterm != &a {
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
func testStockCvtermToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockCvterm
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockCvtermDBTypes, false, strmangle.SetComplement(stockCvtermPrimaryKeyColumns, stockCvtermColumnsWithoutDefault)...); err != nil {
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

		if x.R.StockCvterm != &a {
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
func testStockCvtermToOneSetOpStockUsingStock(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockCvterm
	var b, c Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockCvtermDBTypes, false, strmangle.SetComplement(stockCvtermPrimaryKeyColumns, stockCvtermColumnsWithoutDefault)...); err != nil {
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

		if x.R.StockCvterm != &a {
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
func testStockCvtermsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvterm := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockCvterm.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockCvtermsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvterm := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockCvtermSlice{stockCvterm}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockCvtermsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockCvterm := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockCvterms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockCvtermDBTypes = map[string]string{"CvtermID": "integer", "IsNot": "boolean", "PubID": "integer", "Rank": "integer", "StockCvtermID": "integer", "StockID": "integer"}
	_                  = bytes.MinRead
)

func testStockCvtermsUpdate(t *testing.T) {
	t.Parallel()

	if len(stockCvtermColumns) == len(stockCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockCvterm := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	if err = stockCvterm.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockCvtermsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockCvtermColumns) == len(stockCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockCvterm := &StockCvterm{}
	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockCvterm, stockCvtermDBTypes, true, stockCvtermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockCvtermColumns, stockCvtermPrimaryKeyColumns) {
		fields = stockCvtermColumns
	} else {
		fields = strmangle.SetComplement(
			stockCvtermColumns,
			stockCvtermPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockCvterm))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockCvtermSlice{stockCvterm}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockCvtermsUpsert(t *testing.T) {
	t.Parallel()

	if len(stockCvtermColumns) == len(stockCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockCvterm := StockCvterm{}
	if err = randomize.Struct(seed, &stockCvterm, stockCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockCvterm.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockCvterm: %s", err)
	}

	count, err := StockCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockCvterm, stockCvtermDBTypes, false, stockCvtermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}

	if err = stockCvterm.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockCvterm: %s", err)
	}

	count, err = StockCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
