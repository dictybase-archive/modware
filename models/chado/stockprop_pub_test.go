package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockpropPubs(t *testing.T) {
	t.Parallel()

	query := StockpropPubs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockpropPubsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockpropPub := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockpropPub.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := StockpropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockpropPubsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockpropPub := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockpropPubs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := StockpropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockpropPubsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockpropPub := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockpropPubSlice{stockpropPub}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := StockpropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockpropPubsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockpropPub := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockpropPubExists(tx, stockpropPub.StockpropPubID)
	if err != nil {
		t.Errorf("Unable to check if StockpropPub exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockpropPubExistsG to return true, but got false.")
	}
}
func testStockpropPubsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockpropPub := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	stockpropPubFound, err := FindStockpropPub(tx, stockpropPub.StockpropPubID)
	if err != nil {
		t.Error(err)
	}

	if stockpropPubFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockpropPubsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockpropPub := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockpropPubs(tx).Bind(stockpropPub); err != nil {
		t.Error(err)
	}
}

func testStockpropPubsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockpropPub := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := StockpropPubs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockpropPubsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockpropPubOne := &StockpropPub{}
	stockpropPubTwo := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPubOne, stockpropPubDBTypes, false, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}
	if err = randomize.Struct(seed, stockpropPubTwo, stockpropPubDBTypes, false, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockpropPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockpropPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockpropPubsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockpropPubOne := &StockpropPub{}
	stockpropPubTwo := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPubOne, stockpropPubDBTypes, false, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}
	if err = randomize.Struct(seed, stockpropPubTwo, stockpropPubDBTypes, false, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockpropPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockpropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockpropPubBeforeInsertHook(e boil.Executor, o *StockpropPub) error {
	*o = StockpropPub{}
	return nil
}

func stockpropPubAfterInsertHook(e boil.Executor, o *StockpropPub) error {
	*o = StockpropPub{}
	return nil
}

func stockpropPubAfterSelectHook(e boil.Executor, o *StockpropPub) error {
	*o = StockpropPub{}
	return nil
}

func stockpropPubBeforeUpdateHook(e boil.Executor, o *StockpropPub) error {
	*o = StockpropPub{}
	return nil
}

func stockpropPubAfterUpdateHook(e boil.Executor, o *StockpropPub) error {
	*o = StockpropPub{}
	return nil
}

func stockpropPubBeforeDeleteHook(e boil.Executor, o *StockpropPub) error {
	*o = StockpropPub{}
	return nil
}

func stockpropPubAfterDeleteHook(e boil.Executor, o *StockpropPub) error {
	*o = StockpropPub{}
	return nil
}

func stockpropPubBeforeUpsertHook(e boil.Executor, o *StockpropPub) error {
	*o = StockpropPub{}
	return nil
}

func stockpropPubAfterUpsertHook(e boil.Executor, o *StockpropPub) error {
	*o = StockpropPub{}
	return nil
}

func testStockpropPubsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &StockpropPub{}
	o := &StockpropPub{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockpropPubDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StockpropPub object: %s", err)
	}

	AddStockpropPubHook(boil.BeforeInsertHook, stockpropPubBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockpropPubBeforeInsertHooks = []StockpropPubHook{}

	AddStockpropPubHook(boil.AfterInsertHook, stockpropPubAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockpropPubAfterInsertHooks = []StockpropPubHook{}

	AddStockpropPubHook(boil.AfterSelectHook, stockpropPubAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockpropPubAfterSelectHooks = []StockpropPubHook{}

	AddStockpropPubHook(boil.BeforeUpdateHook, stockpropPubBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockpropPubBeforeUpdateHooks = []StockpropPubHook{}

	AddStockpropPubHook(boil.AfterUpdateHook, stockpropPubAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockpropPubAfterUpdateHooks = []StockpropPubHook{}

	AddStockpropPubHook(boil.BeforeDeleteHook, stockpropPubBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockpropPubBeforeDeleteHooks = []StockpropPubHook{}

	AddStockpropPubHook(boil.AfterDeleteHook, stockpropPubAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockpropPubAfterDeleteHooks = []StockpropPubHook{}

	AddStockpropPubHook(boil.BeforeUpsertHook, stockpropPubBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockpropPubBeforeUpsertHooks = []StockpropPubHook{}

	AddStockpropPubHook(boil.AfterUpsertHook, stockpropPubAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockpropPubAfterUpsertHooks = []StockpropPubHook{}
}
func testStockpropPubsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockpropPub := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockpropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockpropPubsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockpropPub := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPub.Insert(tx, stockpropPubColumns...); err != nil {
		t.Error(err)
	}

	count, err := StockpropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockpropPubToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockpropPub
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockpropPubDBTypes, true, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
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

	slice := StockpropPubSlice{&local}
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

func testStockpropPubToOneStockpropUsingStockprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockpropPub
	var foreign Stockprop

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockpropPubDBTypes, true, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, stockpropDBTypes, true, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.StockpropID = foreign.StockpropID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Stockprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockpropID != foreign.StockpropID {
		t.Errorf("want: %v, got %v", foreign.StockpropID, check.StockpropID)
	}

	slice := StockpropPubSlice{&local}
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

func testStockpropPubToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockpropPub
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockpropPubDBTypes, false, strmangle.SetComplement(stockpropPubPrimaryKeyColumns, stockpropPubColumnsWithoutDefault)...); err != nil {
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

		if x.R.StockpropPub != &a {
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
func testStockpropPubToOneSetOpStockpropUsingStockprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockpropPub
	var b, c Stockprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockpropPubDBTypes, false, strmangle.SetComplement(stockpropPubPrimaryKeyColumns, stockpropPubColumnsWithoutDefault)...); err != nil {
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

		if x.R.StockpropPub != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StockpropID != x.StockpropID {
			t.Error("foreign key was wrong value", a.StockpropID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.StockpropID))
		reflect.Indirect(reflect.ValueOf(&a.StockpropID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockpropID != x.StockpropID {
			t.Error("foreign key was wrong value", a.StockpropID, x.StockpropID)
		}
	}
}
func testStockpropPubsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockpropPub := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockpropPub.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockpropPubsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockpropPub := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockpropPubSlice{stockpropPub}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockpropPubsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockpropPub := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockpropPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockpropPubDBTypes = map[string]string{"PubID": "integer", "StockpropID": "integer", "StockpropPubID": "integer"}
	_                   = bytes.MinRead
)

func testStockpropPubsUpdate(t *testing.T) {
	t.Parallel()

	if len(stockpropPubColumns) == len(stockpropPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockpropPub := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockpropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	if err = stockpropPub.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockpropPubsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockpropPubColumns) == len(stockpropPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockpropPub := &StockpropPub{}
	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockpropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockpropPub, stockpropPubDBTypes, true, stockpropPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockpropPubColumns, stockpropPubPrimaryKeyColumns) {
		fields = stockpropPubColumns
	} else {
		fields = strmangle.SetComplement(
			stockpropPubColumns,
			stockpropPubPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockpropPub))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockpropPubSlice{stockpropPub}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockpropPubsUpsert(t *testing.T) {
	t.Parallel()

	if len(stockpropPubColumns) == len(stockpropPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockpropPub := StockpropPub{}
	if err = randomize.Struct(seed, &stockpropPub, stockpropPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockpropPub.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockpropPub: %s", err)
	}

	count, err := StockpropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockpropPub, stockpropPubDBTypes, false, stockpropPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}

	if err = stockpropPub.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockpropPub: %s", err)
	}

	count, err = StockpropPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
