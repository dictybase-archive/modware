package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockcollectionStocks(t *testing.T) {
	t.Parallel()

	query := StockcollectionStocks(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockcollectionStocksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionStock := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStock.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockcollectionStock.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := StockcollectionStocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockcollectionStocksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionStock := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStock.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockcollectionStocks(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := StockcollectionStocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockcollectionStocksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionStock := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStock.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockcollectionStockSlice{stockcollectionStock}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := StockcollectionStocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockcollectionStocksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionStock := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStock.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockcollectionStockExists(tx, stockcollectionStock.StockcollectionStockID)
	if err != nil {
		t.Errorf("Unable to check if StockcollectionStock exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockcollectionStockExistsG to return true, but got false.")
	}
}
func testStockcollectionStocksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionStock := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStock.Insert(tx); err != nil {
		t.Error(err)
	}

	stockcollectionStockFound, err := FindStockcollectionStock(tx, stockcollectionStock.StockcollectionStockID)
	if err != nil {
		t.Error(err)
	}

	if stockcollectionStockFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockcollectionStocksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionStock := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStock.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockcollectionStocks(tx).Bind(stockcollectionStock); err != nil {
		t.Error(err)
	}
}

func testStockcollectionStocksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionStock := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStock.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := StockcollectionStocks(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockcollectionStocksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionStockOne := &StockcollectionStock{}
	stockcollectionStockTwo := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStockOne, stockcollectionStockDBTypes, false, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}
	if err = randomize.Struct(seed, stockcollectionStockTwo, stockcollectionStockDBTypes, false, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStockOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockcollectionStockTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockcollectionStocks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockcollectionStocksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockcollectionStockOne := &StockcollectionStock{}
	stockcollectionStockTwo := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStockOne, stockcollectionStockDBTypes, false, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}
	if err = randomize.Struct(seed, stockcollectionStockTwo, stockcollectionStockDBTypes, false, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStockOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockcollectionStockTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockcollectionStocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockcollectionStockBeforeInsertHook(e boil.Executor, o *StockcollectionStock) error {
	*o = StockcollectionStock{}
	return nil
}

func stockcollectionStockAfterInsertHook(e boil.Executor, o *StockcollectionStock) error {
	*o = StockcollectionStock{}
	return nil
}

func stockcollectionStockAfterSelectHook(e boil.Executor, o *StockcollectionStock) error {
	*o = StockcollectionStock{}
	return nil
}

func stockcollectionStockBeforeUpdateHook(e boil.Executor, o *StockcollectionStock) error {
	*o = StockcollectionStock{}
	return nil
}

func stockcollectionStockAfterUpdateHook(e boil.Executor, o *StockcollectionStock) error {
	*o = StockcollectionStock{}
	return nil
}

func stockcollectionStockBeforeDeleteHook(e boil.Executor, o *StockcollectionStock) error {
	*o = StockcollectionStock{}
	return nil
}

func stockcollectionStockAfterDeleteHook(e boil.Executor, o *StockcollectionStock) error {
	*o = StockcollectionStock{}
	return nil
}

func stockcollectionStockBeforeUpsertHook(e boil.Executor, o *StockcollectionStock) error {
	*o = StockcollectionStock{}
	return nil
}

func stockcollectionStockAfterUpsertHook(e boil.Executor, o *StockcollectionStock) error {
	*o = StockcollectionStock{}
	return nil
}

func testStockcollectionStocksHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &StockcollectionStock{}
	o := &StockcollectionStock{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockcollectionStockDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock object: %s", err)
	}

	AddStockcollectionStockHook(boil.BeforeInsertHook, stockcollectionStockBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockcollectionStockBeforeInsertHooks = []StockcollectionStockHook{}

	AddStockcollectionStockHook(boil.AfterInsertHook, stockcollectionStockAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockcollectionStockAfterInsertHooks = []StockcollectionStockHook{}

	AddStockcollectionStockHook(boil.AfterSelectHook, stockcollectionStockAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockcollectionStockAfterSelectHooks = []StockcollectionStockHook{}

	AddStockcollectionStockHook(boil.BeforeUpdateHook, stockcollectionStockBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockcollectionStockBeforeUpdateHooks = []StockcollectionStockHook{}

	AddStockcollectionStockHook(boil.AfterUpdateHook, stockcollectionStockAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockcollectionStockAfterUpdateHooks = []StockcollectionStockHook{}

	AddStockcollectionStockHook(boil.BeforeDeleteHook, stockcollectionStockBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockcollectionStockBeforeDeleteHooks = []StockcollectionStockHook{}

	AddStockcollectionStockHook(boil.AfterDeleteHook, stockcollectionStockAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockcollectionStockAfterDeleteHooks = []StockcollectionStockHook{}

	AddStockcollectionStockHook(boil.BeforeUpsertHook, stockcollectionStockBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockcollectionStockBeforeUpsertHooks = []StockcollectionStockHook{}

	AddStockcollectionStockHook(boil.AfterUpsertHook, stockcollectionStockAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockcollectionStockAfterUpsertHooks = []StockcollectionStockHook{}
}
func testStockcollectionStocksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionStock := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStock.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockcollectionStocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockcollectionStocksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionStock := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStock.Insert(tx, stockcollectionStockColumns...); err != nil {
		t.Error(err)
	}

	count, err := StockcollectionStocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockcollectionStockToOneStockUsingStock(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockcollectionStock
	var foreign Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockcollectionStockDBTypes, true, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
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

	slice := StockcollectionStockSlice{&local}
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

func testStockcollectionStockToOneStockcollectionUsingStockcollection(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockcollectionStock
	var foreign Stockcollection

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockcollectionStockDBTypes, true, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, stockcollectionDBTypes, true, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.StockcollectionID = foreign.StockcollectionID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Stockcollection(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockcollectionID != foreign.StockcollectionID {
		t.Errorf("want: %v, got %v", foreign.StockcollectionID, check.StockcollectionID)
	}

	slice := StockcollectionStockSlice{&local}
	if err = local.L.LoadStockcollection(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Stockcollection == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Stockcollection = nil
	if err = local.L.LoadStockcollection(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Stockcollection == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockcollectionStockToOneSetOpStockUsingStock(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockcollectionStock
	var b, c Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockcollectionStockDBTypes, false, strmangle.SetComplement(stockcollectionStockPrimaryKeyColumns, stockcollectionStockColumnsWithoutDefault)...); err != nil {
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

		if x.R.StockcollectionStock != &a {
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
func testStockcollectionStockToOneSetOpStockcollectionUsingStockcollection(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockcollectionStock
	var b, c Stockcollection

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockcollectionStockDBTypes, false, strmangle.SetComplement(stockcollectionStockPrimaryKeyColumns, stockcollectionStockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockcollectionDBTypes, false, strmangle.SetComplement(stockcollectionPrimaryKeyColumns, stockcollectionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockcollectionDBTypes, false, strmangle.SetComplement(stockcollectionPrimaryKeyColumns, stockcollectionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Stockcollection{&b, &c} {
		err = a.SetStockcollection(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Stockcollection != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.StockcollectionStock != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StockcollectionID != x.StockcollectionID {
			t.Error("foreign key was wrong value", a.StockcollectionID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.StockcollectionID))
		reflect.Indirect(reflect.ValueOf(&a.StockcollectionID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockcollectionID != x.StockcollectionID {
			t.Error("foreign key was wrong value", a.StockcollectionID, x.StockcollectionID)
		}
	}
}
func testStockcollectionStocksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionStock := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStock.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockcollectionStock.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockcollectionStocksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionStock := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStock.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockcollectionStockSlice{stockcollectionStock}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockcollectionStocksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionStock := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStock.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockcollectionStocks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockcollectionStockDBTypes = map[string]string{"StockID": "integer", "StockcollectionID": "integer", "StockcollectionStockID": "integer"}
	_                           = bytes.MinRead
)

func testStockcollectionStocksUpdate(t *testing.T) {
	t.Parallel()

	if len(stockcollectionStockColumns) == len(stockcollectionStockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockcollectionStock := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStock.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockcollectionStocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	if err = stockcollectionStock.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockcollectionStocksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockcollectionStockColumns) == len(stockcollectionStockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockcollectionStock := &StockcollectionStock{}
	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStock.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockcollectionStocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockcollectionStock, stockcollectionStockDBTypes, true, stockcollectionStockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockcollectionStockColumns, stockcollectionStockPrimaryKeyColumns) {
		fields = stockcollectionStockColumns
	} else {
		fields = strmangle.SetComplement(
			stockcollectionStockColumns,
			stockcollectionStockPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockcollectionStock))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockcollectionStockSlice{stockcollectionStock}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockcollectionStocksUpsert(t *testing.T) {
	t.Parallel()

	if len(stockcollectionStockColumns) == len(stockcollectionStockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockcollectionStock := StockcollectionStock{}
	if err = randomize.Struct(seed, &stockcollectionStock, stockcollectionStockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionStock.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockcollectionStock: %s", err)
	}

	count, err := StockcollectionStocks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockcollectionStock, stockcollectionStockDBTypes, false, stockcollectionStockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}

	if err = stockcollectionStock.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockcollectionStock: %s", err)
	}

	count, err = StockcollectionStocks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
