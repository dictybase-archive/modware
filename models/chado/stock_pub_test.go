package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockPubs(t *testing.T) {
	t.Parallel()

	query := StockPubs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockPubsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockPub := &StockPub{}
	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockPub.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := StockPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockPubsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockPub := &StockPub{}
	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockPubs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := StockPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockPubsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockPub := &StockPub{}
	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockPubSlice{stockPub}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := StockPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockPubsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockPub := &StockPub{}
	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPub.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockPubExists(tx, stockPub.StockPubID)
	if err != nil {
		t.Errorf("Unable to check if StockPub exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockPubExistsG to return true, but got false.")
	}
}
func testStockPubsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockPub := &StockPub{}
	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPub.Insert(tx); err != nil {
		t.Error(err)
	}

	stockPubFound, err := FindStockPub(tx, stockPub.StockPubID)
	if err != nil {
		t.Error(err)
	}

	if stockPubFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockPubsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockPub := &StockPub{}
	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockPubs(tx).Bind(stockPub); err != nil {
		t.Error(err)
	}
}

func testStockPubsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockPub := &StockPub{}
	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := StockPubs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockPubsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockPubOne := &StockPub{}
	stockPubTwo := &StockPub{}
	if err = randomize.Struct(seed, stockPubOne, stockPubDBTypes, false, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}
	if err = randomize.Struct(seed, stockPubTwo, stockPubDBTypes, false, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockPubsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockPubOne := &StockPub{}
	stockPubTwo := &StockPub{}
	if err = randomize.Struct(seed, stockPubOne, stockPubDBTypes, false, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}
	if err = randomize.Struct(seed, stockPubTwo, stockPubDBTypes, false, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockPubBeforeInsertHook(e boil.Executor, o *StockPub) error {
	*o = StockPub{}
	return nil
}

func stockPubAfterInsertHook(e boil.Executor, o *StockPub) error {
	*o = StockPub{}
	return nil
}

func stockPubAfterSelectHook(e boil.Executor, o *StockPub) error {
	*o = StockPub{}
	return nil
}

func stockPubBeforeUpdateHook(e boil.Executor, o *StockPub) error {
	*o = StockPub{}
	return nil
}

func stockPubAfterUpdateHook(e boil.Executor, o *StockPub) error {
	*o = StockPub{}
	return nil
}

func stockPubBeforeDeleteHook(e boil.Executor, o *StockPub) error {
	*o = StockPub{}
	return nil
}

func stockPubAfterDeleteHook(e boil.Executor, o *StockPub) error {
	*o = StockPub{}
	return nil
}

func stockPubBeforeUpsertHook(e boil.Executor, o *StockPub) error {
	*o = StockPub{}
	return nil
}

func stockPubAfterUpsertHook(e boil.Executor, o *StockPub) error {
	*o = StockPub{}
	return nil
}

func testStockPubsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &StockPub{}
	o := &StockPub{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockPubDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StockPub object: %s", err)
	}

	AddStockPubHook(boil.BeforeInsertHook, stockPubBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockPubBeforeInsertHooks = []StockPubHook{}

	AddStockPubHook(boil.AfterInsertHook, stockPubAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockPubAfterInsertHooks = []StockPubHook{}

	AddStockPubHook(boil.AfterSelectHook, stockPubAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockPubAfterSelectHooks = []StockPubHook{}

	AddStockPubHook(boil.BeforeUpdateHook, stockPubBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockPubBeforeUpdateHooks = []StockPubHook{}

	AddStockPubHook(boil.AfterUpdateHook, stockPubAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockPubAfterUpdateHooks = []StockPubHook{}

	AddStockPubHook(boil.BeforeDeleteHook, stockPubBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockPubBeforeDeleteHooks = []StockPubHook{}

	AddStockPubHook(boil.AfterDeleteHook, stockPubAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockPubAfterDeleteHooks = []StockPubHook{}

	AddStockPubHook(boil.BeforeUpsertHook, stockPubBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockPubBeforeUpsertHooks = []StockPubHook{}

	AddStockPubHook(boil.AfterUpsertHook, stockPubAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockPubAfterUpsertHooks = []StockPubHook{}
}
func testStockPubsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockPub := &StockPub{}
	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockPubsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockPub := &StockPub{}
	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPub.Insert(tx, stockPubColumns...); err != nil {
		t.Error(err)
	}

	count, err := StockPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockPubToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockPub
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockPubDBTypes, true, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
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

	slice := StockPubSlice{&local}
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

func testStockPubToOneStockUsingStock(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockPub
	var foreign Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockPubDBTypes, true, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
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

	slice := StockPubSlice{&local}
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

func testStockPubToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockPub
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockPubDBTypes, false, strmangle.SetComplement(stockPubPrimaryKeyColumns, stockPubColumnsWithoutDefault)...); err != nil {
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

		if x.R.StockPub != &a {
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
func testStockPubToOneSetOpStockUsingStock(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockPub
	var b, c Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockPubDBTypes, false, strmangle.SetComplement(stockPubPrimaryKeyColumns, stockPubColumnsWithoutDefault)...); err != nil {
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

		if x.R.StockPub != &a {
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
func testStockPubsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockPub := &StockPub{}
	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockPub.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockPubsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockPub := &StockPub{}
	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockPubSlice{stockPub}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockPubsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockPub := &StockPub{}
	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockPubDBTypes = map[string]string{"PubID": "integer", "StockID": "integer", "StockPubID": "integer"}
	_               = bytes.MinRead
)

func testStockPubsUpdate(t *testing.T) {
	t.Parallel()

	if len(stockPubColumns) == len(stockPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockPub := &StockPub{}
	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	if err = stockPub.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockPubsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockPubColumns) == len(stockPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockPub := &StockPub{}
	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockPub, stockPubDBTypes, true, stockPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockPubColumns, stockPubPrimaryKeyColumns) {
		fields = stockPubColumns
	} else {
		fields = strmangle.SetComplement(
			stockPubColumns,
			stockPubPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockPub))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockPubSlice{stockPub}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockPubsUpsert(t *testing.T) {
	t.Parallel()

	if len(stockPubColumns) == len(stockPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockPub := StockPub{}
	if err = randomize.Struct(seed, &stockPub, stockPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockPub.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockPub: %s", err)
	}

	count, err := StockPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockPub, stockPubDBTypes, false, stockPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}

	if err = stockPub.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockPub: %s", err)
	}

	count, err = StockPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
