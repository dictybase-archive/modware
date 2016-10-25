package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStocks(t *testing.T) {
	t.Parallel()

	query := Stocks(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStocksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stock := &Stock{}
	if err = randomize.Struct(seed, stock, stockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stock.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stock.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Stocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStocksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stock := &Stock{}
	if err = randomize.Struct(seed, stock, stockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stock.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Stocks(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Stocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStocksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stock := &Stock{}
	if err = randomize.Struct(seed, stock, stockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stock.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockSlice{stock}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Stocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStocksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stock := &Stock{}
	if err = randomize.Struct(seed, stock, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stock.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockExists(tx, stock.StockID)
	if err != nil {
		t.Errorf("Unable to check if Stock exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockExistsG to return true, but got false.")
	}
}
func testStocksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stock := &Stock{}
	if err = randomize.Struct(seed, stock, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stock.Insert(tx); err != nil {
		t.Error(err)
	}

	stockFound, err := FindStock(tx, stock.StockID)
	if err != nil {
		t.Error(err)
	}

	if stockFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStocksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stock := &Stock{}
	if err = randomize.Struct(seed, stock, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stock.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Stocks(tx).Bind(stock); err != nil {
		t.Error(err)
	}
}

func testStocksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stock := &Stock{}
	if err = randomize.Struct(seed, stock, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stock.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Stocks(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStocksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOne := &Stock{}
	stockTwo := &Stock{}
	if err = randomize.Struct(seed, stockOne, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}
	if err = randomize.Struct(seed, stockTwo, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Stocks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStocksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockOne := &Stock{}
	stockTwo := &Stock{}
	if err = randomize.Struct(seed, stockOne, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}
	if err = randomize.Struct(seed, stockTwo, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockBeforeInsertHook(e boil.Executor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockAfterInsertHook(e boil.Executor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockAfterSelectHook(e boil.Executor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockBeforeUpdateHook(e boil.Executor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockAfterUpdateHook(e boil.Executor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockBeforeDeleteHook(e boil.Executor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockAfterDeleteHook(e boil.Executor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockBeforeUpsertHook(e boil.Executor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockAfterUpsertHook(e boil.Executor, o *Stock) error {
	*o = Stock{}
	return nil
}

func testStocksHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Stock{}
	o := &Stock{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Stock object: %s", err)
	}

	AddStockHook(boil.BeforeInsertHook, stockBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockBeforeInsertHooks = []StockHook{}

	AddStockHook(boil.AfterInsertHook, stockAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockAfterInsertHooks = []StockHook{}

	AddStockHook(boil.AfterSelectHook, stockAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockAfterSelectHooks = []StockHook{}

	AddStockHook(boil.BeforeUpdateHook, stockBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockBeforeUpdateHooks = []StockHook{}

	AddStockHook(boil.AfterUpdateHook, stockAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockAfterUpdateHooks = []StockHook{}

	AddStockHook(boil.BeforeDeleteHook, stockBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockBeforeDeleteHooks = []StockHook{}

	AddStockHook(boil.AfterDeleteHook, stockAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockAfterDeleteHooks = []StockHook{}

	AddStockHook(boil.BeforeUpsertHook, stockBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockBeforeUpsertHooks = []StockHook{}

	AddStockHook(boil.AfterUpsertHook, stockAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockAfterUpsertHooks = []StockHook{}
}
func testStocksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stock := &Stock{}
	if err = randomize.Struct(seed, stock, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stock.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStocksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stock := &Stock{}
	if err = randomize.Struct(seed, stock, stockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stock.Insert(tx, stockColumns...); err != nil {
		t.Error(err)
	}

	count, err := Stocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockOneToOneStockPubUsingStockPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockPub
	var local Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockPubDBTypes, true, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.StockID = local.StockID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockID != foreign.StockID {
		t.Errorf("want: %v, got %v", foreign.StockID, check.StockID)
	}

	slice := StockSlice{&local}
	if err = local.L.LoadStockPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockPub = nil
	if err = local.L.LoadStockPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockOneToOneStockItemOrderUsingItemStockItemOrder(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockItemOrder
	var local Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockItemOrderDBTypes, true, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.ItemID = local.StockID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ItemStockItemOrder(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ItemID != foreign.ItemID {
		t.Errorf("want: %v, got %v", foreign.ItemID, check.ItemID)
	}

	slice := StockSlice{&local}
	if err = local.L.LoadItemStockItemOrder(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.ItemStockItemOrder == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ItemStockItemOrder = nil
	if err = local.L.LoadItemStockItemOrder(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.ItemStockItemOrder == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockOneToOneStockcollectionStockUsingStockcollectionStock(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockcollectionStock
	var local Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockcollectionStockDBTypes, true, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.StockID = local.StockID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockcollectionStock(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockID != foreign.StockID {
		t.Errorf("want: %v, got %v", foreign.StockID, check.StockID)
	}

	slice := StockSlice{&local}
	if err = local.L.LoadStockcollectionStock(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockcollectionStock == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockcollectionStock = nil
	if err = local.L.LoadStockcollectionStock(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockcollectionStock == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockOneToOneStockCvtermUsingStockCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockCvterm
	var local Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.StockID = local.StockID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockCvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockID != foreign.StockID {
		t.Errorf("want: %v, got %v", foreign.StockID, check.StockID)
	}

	slice := StockSlice{&local}
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

func testStockOneToOneStockRelationshipUsingObjectStockRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockRelationship
	var local Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.ObjectID = local.StockID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ObjectStockRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ObjectID != foreign.ObjectID {
		t.Errorf("want: %v, got %v", foreign.ObjectID, check.ObjectID)
	}

	slice := StockSlice{&local}
	if err = local.L.LoadObjectStockRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.ObjectStockRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ObjectStockRelationship = nil
	if err = local.L.LoadObjectStockRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.ObjectStockRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockOneToOneStockRelationshipUsingSubjectStockRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockRelationship
	var local Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.SubjectID = local.StockID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.SubjectStockRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.SubjectID != foreign.SubjectID {
		t.Errorf("want: %v, got %v", foreign.SubjectID, check.SubjectID)
	}

	slice := StockSlice{&local}
	if err = local.L.LoadSubjectStockRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.SubjectStockRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SubjectStockRelationship = nil
	if err = local.L.LoadSubjectStockRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.SubjectStockRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockOneToOneStockDbxrefUsingStockDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockDbxref
	var local Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockDbxrefDBTypes, true, stockDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.StockID = local.StockID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockDbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockID != foreign.StockID {
		t.Errorf("want: %v, got %v", foreign.StockID, check.StockID)
	}

	slice := StockSlice{&local}
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

func testStockOneToOneStockGenotypeUsingStockGenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockGenotype
	var local Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockGenotypeDBTypes, true, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.StockID = local.StockID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockGenotype(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockID != foreign.StockID {
		t.Errorf("want: %v, got %v", foreign.StockID, check.StockID)
	}

	slice := StockSlice{&local}
	if err = local.L.LoadStockGenotype(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockGenotype == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockGenotype = nil
	if err = local.L.LoadStockGenotype(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockGenotype == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockOneToOneStockpropUsingStockprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Stockprop
	var local Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockpropDBTypes, true, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.StockID = local.StockID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Stockprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockID != foreign.StockID {
		t.Errorf("want: %v, got %v", foreign.StockID, check.StockID)
	}

	slice := StockSlice{&local}
	if err = local.L.LoadStockprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Stockprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Stockprop = nil
	if err = local.L.LoadStockprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Stockprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockOneToOneSetOpStockPubUsingStockPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stock
	var b, c StockPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockPubDBTypes, false, strmangle.SetComplement(stockPubPrimaryKeyColumns, stockPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockPubDBTypes, false, strmangle.SetComplement(stockPubPrimaryKeyColumns, stockPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockPub{&b, &c} {
		err = a.SetStockPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Stock != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.StockID))
		reflect.Indirect(reflect.ValueOf(&x.StockID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID, x.StockID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testStockOneToOneSetOpStockItemOrderUsingItemStockItemOrder(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stock
	var b, c StockItemOrder

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockItemOrderDBTypes, false, strmangle.SetComplement(stockItemOrderPrimaryKeyColumns, stockItemOrderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockItemOrderDBTypes, false, strmangle.SetComplement(stockItemOrderPrimaryKeyColumns, stockItemOrderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockItemOrder{&b, &c} {
		err = a.SetItemStockItemOrder(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ItemStockItemOrder != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Item != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockID != x.ItemID {
			t.Error("foreign key was wrong value", a.StockID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.ItemID))
		reflect.Indirect(reflect.ValueOf(&x.ItemID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockID != x.ItemID {
			t.Error("foreign key was wrong value", a.StockID, x.ItemID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testStockOneToOneSetOpStockcollectionStockUsingStockcollectionStock(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stock
	var b, c StockcollectionStock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockcollectionStockDBTypes, false, strmangle.SetComplement(stockcollectionStockPrimaryKeyColumns, stockcollectionStockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockcollectionStockDBTypes, false, strmangle.SetComplement(stockcollectionStockPrimaryKeyColumns, stockcollectionStockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockcollectionStock{&b, &c} {
		err = a.SetStockcollectionStock(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockcollectionStock != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Stock != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.StockID))
		reflect.Indirect(reflect.ValueOf(&x.StockID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID, x.StockID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testStockOneToOneSetOpStockCvtermUsingStockCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stock
	var b, c StockCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
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
		if x.R.Stock != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.StockID))
		reflect.Indirect(reflect.ValueOf(&x.StockID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID, x.StockID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testStockOneToOneSetOpStockRelationshipUsingObjectStockRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stock
	var b, c StockRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockRelationshipDBTypes, false, strmangle.SetComplement(stockRelationshipPrimaryKeyColumns, stockRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockRelationshipDBTypes, false, strmangle.SetComplement(stockRelationshipPrimaryKeyColumns, stockRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockRelationship{&b, &c} {
		err = a.SetObjectStockRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ObjectStockRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Object != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockID != x.ObjectID {
			t.Error("foreign key was wrong value", a.StockID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.ObjectID))
		reflect.Indirect(reflect.ValueOf(&x.ObjectID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockID != x.ObjectID {
			t.Error("foreign key was wrong value", a.StockID, x.ObjectID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testStockOneToOneSetOpStockRelationshipUsingSubjectStockRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stock
	var b, c StockRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockRelationshipDBTypes, false, strmangle.SetComplement(stockRelationshipPrimaryKeyColumns, stockRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockRelationshipDBTypes, false, strmangle.SetComplement(stockRelationshipPrimaryKeyColumns, stockRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockRelationship{&b, &c} {
		err = a.SetSubjectStockRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SubjectStockRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Subject != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockID != x.SubjectID {
			t.Error("foreign key was wrong value", a.StockID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.SubjectID))
		reflect.Indirect(reflect.ValueOf(&x.SubjectID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockID != x.SubjectID {
			t.Error("foreign key was wrong value", a.StockID, x.SubjectID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testStockOneToOneSetOpStockDbxrefUsingStockDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stock
	var b, c StockDbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
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
		if x.R.Stock != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.StockID))
		reflect.Indirect(reflect.ValueOf(&x.StockID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID, x.StockID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testStockOneToOneSetOpStockGenotypeUsingStockGenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stock
	var b, c StockGenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockGenotypeDBTypes, false, strmangle.SetComplement(stockGenotypePrimaryKeyColumns, stockGenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockGenotypeDBTypes, false, strmangle.SetComplement(stockGenotypePrimaryKeyColumns, stockGenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockGenotype{&b, &c} {
		err = a.SetStockGenotype(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockGenotype != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Stock != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.StockID))
		reflect.Indirect(reflect.ValueOf(&x.StockID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID, x.StockID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testStockOneToOneSetOpStockpropUsingStockprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stock
	var b, c Stockprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockpropDBTypes, false, strmangle.SetComplement(stockpropPrimaryKeyColumns, stockpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockpropDBTypes, false, strmangle.SetComplement(stockpropPrimaryKeyColumns, stockpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Stockprop{&b, &c} {
		err = a.SetStockprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Stockprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Stock != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.StockID))
		reflect.Indirect(reflect.ValueOf(&x.StockID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID, x.StockID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testStockToOneOrganismUsingOrganism(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Stock
	var foreign Organism

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	local.OrganismID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.OrganismID.Int = foreign.OrganismID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Organism(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.OrganismID != foreign.OrganismID {
		t.Errorf("want: %v, got %v", foreign.OrganismID, check.OrganismID)
	}

	slice := StockSlice{&local}
	if err = local.L.LoadOrganism(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Organism == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Organism = nil
	if err = local.L.LoadOrganism(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Organism == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Stock
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
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

	slice := StockSlice{&local}
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

func testStockToOneDbxrefUsingDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Stock
	var foreign Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	local.DbxrefID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.DbxrefID.Int = foreign.DbxrefID
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

	slice := StockSlice{&local}
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

func testStockToOneSetOpOrganismUsingOrganism(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stock
	var b, c Organism

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organismDBTypes, false, strmangle.SetComplement(organismPrimaryKeyColumns, organismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, organismDBTypes, false, strmangle.SetComplement(organismPrimaryKeyColumns, organismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Organism{&b, &c} {
		err = a.SetOrganism(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Organism != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Stock != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OrganismID.Int != x.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OrganismID.Int))
		reflect.Indirect(reflect.ValueOf(&a.OrganismID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrganismID.Int != x.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID.Int, x.OrganismID)
		}
	}
}

func testStockToOneRemoveOpOrganismUsingOrganism(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stock
	var b Organism

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organismDBTypes, false, strmangle.SetComplement(organismPrimaryKeyColumns, organismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetOrganism(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveOrganism(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Organism(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Organism != nil {
		t.Error("R struct entry should be nil")
	}

	if a.OrganismID.Valid {
		t.Error("foreign key value should be nil")
	}

	if b.R.Stock != nil {
		t.Error("failed to remove a from b's relationships")
	}

}

func testStockToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stock
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeStock != &a {
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
func testStockToOneSetOpDbxrefUsingDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stock
	var b, c Dbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
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

		if x.R.Stocks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DbxrefID.Int != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DbxrefID.Int))
		reflect.Indirect(reflect.ValueOf(&a.DbxrefID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DbxrefID.Int != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID.Int, x.DbxrefID)
		}
	}
}

func testStockToOneRemoveOpDbxrefUsingDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stock
	var b Dbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetDbxref(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveDbxref(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Dbxref(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Dbxref != nil {
		t.Error("R struct entry should be nil")
	}

	if a.DbxrefID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Stocks) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testStocksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stock := &Stock{}
	if err = randomize.Struct(seed, stock, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stock.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stock.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStocksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stock := &Stock{}
	if err = randomize.Struct(seed, stock, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stock.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockSlice{stock}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStocksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stock := &Stock{}
	if err = randomize.Struct(seed, stock, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stock.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Stocks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockDBTypes = map[string]string{"DbxrefID": "integer", "Description": "text", "IsObsolete": "boolean", "Name": "character varying", "OrganismID": "integer", "StockID": "integer", "TypeID": "integer", "Uniquename": "text"}
	_            = bytes.MinRead
)

func testStocksUpdate(t *testing.T) {
	t.Parallel()

	if len(stockColumns) == len(stockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stock := &Stock{}
	if err = randomize.Struct(seed, stock, stockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stock.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stock, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err = stock.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStocksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockColumns) == len(stockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stock := &Stock{}
	if err = randomize.Struct(seed, stock, stockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stock.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stock, stockDBTypes, true, stockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockColumns, stockPrimaryKeyColumns) {
		fields = stockColumns
	} else {
		fields = strmangle.SetComplement(
			stockColumns,
			stockPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stock))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockSlice{stock}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStocksUpsert(t *testing.T) {
	t.Parallel()

	if len(stockColumns) == len(stockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stock := Stock{}
	if err = randomize.Struct(seed, &stock, stockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stock.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Stock: %s", err)
	}

	count, err := Stocks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stock, stockDBTypes, false, stockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err = stock.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Stock: %s", err)
	}

	count, err = Stocks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
