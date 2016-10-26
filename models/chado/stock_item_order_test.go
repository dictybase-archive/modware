package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockItemOrders(t *testing.T) {
	t.Parallel()

	query := StockItemOrders(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockItemOrdersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockItemOrder := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockItemOrder.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := StockItemOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockItemOrdersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockItemOrder := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockItemOrders(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := StockItemOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockItemOrdersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockItemOrder := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockItemOrderSlice{stockItemOrder}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := StockItemOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockItemOrdersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockItemOrder := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockItemOrderExists(tx, stockItemOrder.StockItemOrderID)
	if err != nil {
		t.Errorf("Unable to check if StockItemOrder exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockItemOrderExistsG to return true, but got false.")
	}
}
func testStockItemOrdersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockItemOrder := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	stockItemOrderFound, err := FindStockItemOrder(tx, stockItemOrder.StockItemOrderID)
	if err != nil {
		t.Error(err)
	}

	if stockItemOrderFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockItemOrdersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockItemOrder := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockItemOrders(tx).Bind(stockItemOrder); err != nil {
		t.Error(err)
	}
}

func testStockItemOrdersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockItemOrder := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := StockItemOrders(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockItemOrdersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockItemOrderOne := &StockItemOrder{}
	stockItemOrderTwo := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrderOne, stockItemOrderDBTypes, false, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}
	if err = randomize.Struct(seed, stockItemOrderTwo, stockItemOrderDBTypes, false, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrderOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockItemOrderTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockItemOrders(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockItemOrdersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockItemOrderOne := &StockItemOrder{}
	stockItemOrderTwo := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrderOne, stockItemOrderDBTypes, false, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}
	if err = randomize.Struct(seed, stockItemOrderTwo, stockItemOrderDBTypes, false, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrderOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockItemOrderTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockItemOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockItemOrderBeforeInsertHook(e boil.Executor, o *StockItemOrder) error {
	*o = StockItemOrder{}
	return nil
}

func stockItemOrderAfterInsertHook(e boil.Executor, o *StockItemOrder) error {
	*o = StockItemOrder{}
	return nil
}

func stockItemOrderAfterSelectHook(e boil.Executor, o *StockItemOrder) error {
	*o = StockItemOrder{}
	return nil
}

func stockItemOrderBeforeUpdateHook(e boil.Executor, o *StockItemOrder) error {
	*o = StockItemOrder{}
	return nil
}

func stockItemOrderAfterUpdateHook(e boil.Executor, o *StockItemOrder) error {
	*o = StockItemOrder{}
	return nil
}

func stockItemOrderBeforeDeleteHook(e boil.Executor, o *StockItemOrder) error {
	*o = StockItemOrder{}
	return nil
}

func stockItemOrderAfterDeleteHook(e boil.Executor, o *StockItemOrder) error {
	*o = StockItemOrder{}
	return nil
}

func stockItemOrderBeforeUpsertHook(e boil.Executor, o *StockItemOrder) error {
	*o = StockItemOrder{}
	return nil
}

func stockItemOrderAfterUpsertHook(e boil.Executor, o *StockItemOrder) error {
	*o = StockItemOrder{}
	return nil
}

func testStockItemOrdersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &StockItemOrder{}
	o := &StockItemOrder{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockItemOrderDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StockItemOrder object: %s", err)
	}

	AddStockItemOrderHook(boil.BeforeInsertHook, stockItemOrderBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockItemOrderBeforeInsertHooks = []StockItemOrderHook{}

	AddStockItemOrderHook(boil.AfterInsertHook, stockItemOrderAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockItemOrderAfterInsertHooks = []StockItemOrderHook{}

	AddStockItemOrderHook(boil.AfterSelectHook, stockItemOrderAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockItemOrderAfterSelectHooks = []StockItemOrderHook{}

	AddStockItemOrderHook(boil.BeforeUpdateHook, stockItemOrderBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockItemOrderBeforeUpdateHooks = []StockItemOrderHook{}

	AddStockItemOrderHook(boil.AfterUpdateHook, stockItemOrderAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockItemOrderAfterUpdateHooks = []StockItemOrderHook{}

	AddStockItemOrderHook(boil.BeforeDeleteHook, stockItemOrderBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockItemOrderBeforeDeleteHooks = []StockItemOrderHook{}

	AddStockItemOrderHook(boil.AfterDeleteHook, stockItemOrderAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockItemOrderAfterDeleteHooks = []StockItemOrderHook{}

	AddStockItemOrderHook(boil.BeforeUpsertHook, stockItemOrderBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockItemOrderBeforeUpsertHooks = []StockItemOrderHook{}

	AddStockItemOrderHook(boil.AfterUpsertHook, stockItemOrderAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockItemOrderAfterUpsertHooks = []StockItemOrderHook{}
}
func testStockItemOrdersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockItemOrder := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockItemOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockItemOrdersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockItemOrder := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrder.Insert(tx, stockItemOrderColumns...); err != nil {
		t.Error(err)
	}

	count, err := StockItemOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockItemOrderToOneStockUsingItem(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockItemOrder
	var foreign Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockItemOrderDBTypes, true, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ItemID = foreign.StockID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Item(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockID != foreign.StockID {
		t.Errorf("want: %v, got %v", foreign.StockID, check.StockID)
	}

	slice := StockItemOrderSlice{&local}
	if err = local.L.LoadItem(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Item == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Item = nil
	if err = local.L.LoadItem(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Item == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockItemOrderToOneStockOrderUsingOrder(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockItemOrder
	var foreign StockOrder

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockItemOrderDBTypes, true, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, stockOrderDBTypes, true, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.OrderID = foreign.StockOrderID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Order(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockOrderID != foreign.StockOrderID {
		t.Errorf("want: %v, got %v", foreign.StockOrderID, check.StockOrderID)
	}

	slice := StockItemOrderSlice{&local}
	if err = local.L.LoadOrder(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Order == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Order = nil
	if err = local.L.LoadOrder(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Order == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockItemOrderToOneSetOpStockUsingItem(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockItemOrder
	var b, c Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockItemOrderDBTypes, false, strmangle.SetComplement(stockItemOrderPrimaryKeyColumns, stockItemOrderColumnsWithoutDefault)...); err != nil {
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
		err = a.SetItem(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Item != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ItemStockItemOrder != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ItemID != x.StockID {
			t.Error("foreign key was wrong value", a.ItemID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ItemID))
		reflect.Indirect(reflect.ValueOf(&a.ItemID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ItemID != x.StockID {
			t.Error("foreign key was wrong value", a.ItemID, x.StockID)
		}
	}
}
func testStockItemOrderToOneSetOpStockOrderUsingOrder(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockItemOrder
	var b, c StockOrder

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockItemOrderDBTypes, false, strmangle.SetComplement(stockItemOrderPrimaryKeyColumns, stockItemOrderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockOrderDBTypes, false, strmangle.SetComplement(stockOrderPrimaryKeyColumns, stockOrderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockOrderDBTypes, false, strmangle.SetComplement(stockOrderPrimaryKeyColumns, stockOrderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockOrder{&b, &c} {
		err = a.SetOrder(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Order != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OrderStockItemOrder != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OrderID != x.StockOrderID {
			t.Error("foreign key was wrong value", a.OrderID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OrderID))
		reflect.Indirect(reflect.ValueOf(&a.OrderID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrderID != x.StockOrderID {
			t.Error("foreign key was wrong value", a.OrderID, x.StockOrderID)
		}
	}
}
func testStockItemOrdersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockItemOrder := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockItemOrder.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockItemOrdersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockItemOrder := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockItemOrderSlice{stockItemOrder}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockItemOrdersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockItemOrder := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockItemOrders(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockItemOrderDBTypes = map[string]string{"ItemID": "integer", "OrderID": "integer", "Quantity": "integer", "StockItemOrderID": "integer"}
	_                     = bytes.MinRead
)

func testStockItemOrdersUpdate(t *testing.T) {
	t.Parallel()

	if len(stockItemOrderColumns) == len(stockItemOrderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockItemOrder := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockItemOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	if err = stockItemOrder.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockItemOrdersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockItemOrderColumns) == len(stockItemOrderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockItemOrder := &StockItemOrder{}
	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockItemOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockItemOrder, stockItemOrderDBTypes, true, stockItemOrderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockItemOrderColumns, stockItemOrderPrimaryKeyColumns) {
		fields = stockItemOrderColumns
	} else {
		fields = strmangle.SetComplement(
			stockItemOrderColumns,
			stockItemOrderPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockItemOrder))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockItemOrderSlice{stockItemOrder}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockItemOrdersUpsert(t *testing.T) {
	t.Parallel()

	if len(stockItemOrderColumns) == len(stockItemOrderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockItemOrder := StockItemOrder{}
	if err = randomize.Struct(seed, &stockItemOrder, stockItemOrderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockItemOrder.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockItemOrder: %s", err)
	}

	count, err := StockItemOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockItemOrder, stockItemOrderDBTypes, false, stockItemOrderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}

	if err = stockItemOrder.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockItemOrder: %s", err)
	}

	count, err = StockItemOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
