package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockOrders(t *testing.T) {
	t.Parallel()

	query := StockOrders(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockOrdersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOrder := &StockOrder{}
	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockOrder.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := StockOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockOrdersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOrder := &StockOrder{}
	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockOrders(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := StockOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockOrdersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOrder := &StockOrder{}
	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockOrderSlice{stockOrder}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := StockOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockOrdersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOrder := &StockOrder{}
	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockOrderExists(tx, stockOrder.StockOrderID)
	if err != nil {
		t.Errorf("Unable to check if StockOrder exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockOrderExistsG to return true, but got false.")
	}
}
func testStockOrdersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOrder := &StockOrder{}
	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	stockOrderFound, err := FindStockOrder(tx, stockOrder.StockOrderID)
	if err != nil {
		t.Error(err)
	}

	if stockOrderFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockOrdersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOrder := &StockOrder{}
	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockOrders(tx).Bind(stockOrder); err != nil {
		t.Error(err)
	}
}

func testStockOrdersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOrder := &StockOrder{}
	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := StockOrders(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockOrdersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOrderOne := &StockOrder{}
	stockOrderTwo := &StockOrder{}
	if err = randomize.Struct(seed, stockOrderOne, stockOrderDBTypes, false, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}
	if err = randomize.Struct(seed, stockOrderTwo, stockOrderDBTypes, false, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrderOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockOrderTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockOrders(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockOrdersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockOrderOne := &StockOrder{}
	stockOrderTwo := &StockOrder{}
	if err = randomize.Struct(seed, stockOrderOne, stockOrderDBTypes, false, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}
	if err = randomize.Struct(seed, stockOrderTwo, stockOrderDBTypes, false, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrderOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockOrderTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockOrderBeforeInsertHook(e boil.Executor, o *StockOrder) error {
	*o = StockOrder{}
	return nil
}

func stockOrderAfterInsertHook(e boil.Executor, o *StockOrder) error {
	*o = StockOrder{}
	return nil
}

func stockOrderAfterSelectHook(e boil.Executor, o *StockOrder) error {
	*o = StockOrder{}
	return nil
}

func stockOrderBeforeUpdateHook(e boil.Executor, o *StockOrder) error {
	*o = StockOrder{}
	return nil
}

func stockOrderAfterUpdateHook(e boil.Executor, o *StockOrder) error {
	*o = StockOrder{}
	return nil
}

func stockOrderBeforeDeleteHook(e boil.Executor, o *StockOrder) error {
	*o = StockOrder{}
	return nil
}

func stockOrderAfterDeleteHook(e boil.Executor, o *StockOrder) error {
	*o = StockOrder{}
	return nil
}

func stockOrderBeforeUpsertHook(e boil.Executor, o *StockOrder) error {
	*o = StockOrder{}
	return nil
}

func stockOrderAfterUpsertHook(e boil.Executor, o *StockOrder) error {
	*o = StockOrder{}
	return nil
}

func testStockOrdersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &StockOrder{}
	o := &StockOrder{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockOrderDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StockOrder object: %s", err)
	}

	AddStockOrderHook(boil.BeforeInsertHook, stockOrderBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockOrderBeforeInsertHooks = []StockOrderHook{}

	AddStockOrderHook(boil.AfterInsertHook, stockOrderAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockOrderAfterInsertHooks = []StockOrderHook{}

	AddStockOrderHook(boil.AfterSelectHook, stockOrderAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockOrderAfterSelectHooks = []StockOrderHook{}

	AddStockOrderHook(boil.BeforeUpdateHook, stockOrderBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockOrderBeforeUpdateHooks = []StockOrderHook{}

	AddStockOrderHook(boil.AfterUpdateHook, stockOrderAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockOrderAfterUpdateHooks = []StockOrderHook{}

	AddStockOrderHook(boil.BeforeDeleteHook, stockOrderBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockOrderBeforeDeleteHooks = []StockOrderHook{}

	AddStockOrderHook(boil.AfterDeleteHook, stockOrderAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockOrderAfterDeleteHooks = []StockOrderHook{}

	AddStockOrderHook(boil.BeforeUpsertHook, stockOrderBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockOrderBeforeUpsertHooks = []StockOrderHook{}

	AddStockOrderHook(boil.AfterUpsertHook, stockOrderAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockOrderAfterUpsertHooks = []StockOrderHook{}
}
func testStockOrdersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOrder := &StockOrder{}
	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockOrdersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOrder := &StockOrder{}
	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrder.Insert(tx, stockOrderColumns...); err != nil {
		t.Error(err)
	}

	count, err := StockOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockOrderOneToOneStockItemOrderUsingOrderStockItemOrder(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockItemOrder
	var local StockOrder

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockItemOrderDBTypes, true, stockItemOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItemOrder struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockOrderDBTypes, true, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.OrderID = local.StockOrderID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.OrderStockItemOrder(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.OrderID != foreign.OrderID {
		t.Errorf("want: %v, got %v", foreign.OrderID, check.OrderID)
	}

	slice := StockOrderSlice{&local}
	if err = local.L.LoadOrderStockItemOrder(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.OrderStockItemOrder == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.OrderStockItemOrder = nil
	if err = local.L.LoadOrderStockItemOrder(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.OrderStockItemOrder == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockOrderOneToOneSetOpStockItemOrderUsingOrderStockItemOrder(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockOrder
	var b, c StockItemOrder

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockOrderDBTypes, false, strmangle.SetComplement(stockOrderPrimaryKeyColumns, stockOrderColumnsWithoutDefault)...); err != nil {
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
		err = a.SetOrderStockItemOrder(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.OrderStockItemOrder != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Order != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockOrderID != x.OrderID {
			t.Error("foreign key was wrong value", a.StockOrderID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.OrderID))
		reflect.Indirect(reflect.ValueOf(&x.OrderID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockOrderID != x.OrderID {
			t.Error("foreign key was wrong value", a.StockOrderID, x.OrderID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testStockOrderToOneAuthUserUsingUser(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockOrder
	var foreign AuthUser

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockOrderDBTypes, true, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.AuthUserID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.User(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.AuthUserID != foreign.AuthUserID {
		t.Errorf("want: %v, got %v", foreign.AuthUserID, check.AuthUserID)
	}

	slice := StockOrderSlice{&local}
	if err = local.L.LoadUser(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockOrderToOneSetOpAuthUserUsingUser(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockOrder
	var b, c AuthUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockOrderDBTypes, false, strmangle.SetComplement(stockOrderPrimaryKeyColumns, stockOrderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, authUserDBTypes, false, strmangle.SetComplement(authUserPrimaryKeyColumns, authUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authUserDBTypes, false, strmangle.SetComplement(authUserPrimaryKeyColumns, authUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*AuthUser{&b, &c} {
		err = a.SetUser(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserStockOrders[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.AuthUserID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.AuthUserID {
			t.Error("foreign key was wrong value", a.UserID, x.AuthUserID)
		}
	}
}
func testStockOrdersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOrder := &StockOrder{}
	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockOrder.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockOrdersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOrder := &StockOrder{}
	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockOrderSlice{stockOrder}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockOrdersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOrder := &StockOrder{}
	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockOrders(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockOrderDBTypes = map[string]string{"Comments": "text", "Courier": "USER-DEFINED", "CourierAccount": "integer", "CreatedAt": "timestamp with time zone", "Payment": "USER-DEFINED", "PurchaseOrderID": "integer", "Status": "character varying", "StockOrderID": "integer", "UpdatedAt": "timestamp with time zone", "UserID": "integer"}
	_                 = bytes.MinRead
)

func testStockOrdersUpdate(t *testing.T) {
	t.Parallel()

	if len(stockOrderColumns) == len(stockOrderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockOrder := &StockOrder{}
	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true, stockOrderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	if err = stockOrder.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockOrdersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockOrderColumns) == len(stockOrderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockOrder := &StockOrder{}
	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrder.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockOrder, stockOrderDBTypes, true, stockOrderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockOrderColumns, stockOrderPrimaryKeyColumns) {
		fields = stockOrderColumns
	} else {
		fields = strmangle.SetComplement(
			stockOrderColumns,
			stockOrderPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockOrder))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockOrderSlice{stockOrder}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockOrdersUpsert(t *testing.T) {
	t.Parallel()

	if len(stockOrderColumns) == len(stockOrderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockOrder := StockOrder{}
	if err = randomize.Struct(seed, &stockOrder, stockOrderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockOrder.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockOrder: %s", err)
	}

	count, err := StockOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockOrder, stockOrderDBTypes, false, stockOrderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockOrder struct: %s", err)
	}

	if err = stockOrder.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockOrder: %s", err)
	}

	count, err = StockOrders(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
