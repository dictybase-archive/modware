package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockRelationshipPubs(t *testing.T) {
	t.Parallel()

	query := StockRelationshipPubs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockRelationshipPubsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipPub := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockRelationshipPub.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockRelationshipPubsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipPub := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockRelationshipPubs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockRelationshipPubsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipPub := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockRelationshipPubSlice{stockRelationshipPub}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockRelationshipPubsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipPub := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockRelationshipPubExists(tx, stockRelationshipPub.StockRelationshipPubID)
	if err != nil {
		t.Errorf("Unable to check if StockRelationshipPub exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockRelationshipPubExistsG to return true, but got false.")
	}
}
func testStockRelationshipPubsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipPub := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	stockRelationshipPubFound, err := FindStockRelationshipPub(tx, stockRelationshipPub.StockRelationshipPubID)
	if err != nil {
		t.Error(err)
	}

	if stockRelationshipPubFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockRelationshipPubsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipPub := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockRelationshipPubs(tx).Bind(stockRelationshipPub); err != nil {
		t.Error(err)
	}
}

func testStockRelationshipPubsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipPub := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := StockRelationshipPubs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockRelationshipPubsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipPubOne := &StockRelationshipPub{}
	stockRelationshipPubTwo := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPubOne, stockRelationshipPubDBTypes, false, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}
	if err = randomize.Struct(seed, stockRelationshipPubTwo, stockRelationshipPubDBTypes, false, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockRelationshipPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockRelationshipPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockRelationshipPubsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockRelationshipPubOne := &StockRelationshipPub{}
	stockRelationshipPubTwo := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPubOne, stockRelationshipPubDBTypes, false, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}
	if err = randomize.Struct(seed, stockRelationshipPubTwo, stockRelationshipPubDBTypes, false, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockRelationshipPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockRelationshipPubBeforeInsertHook(e boil.Executor, o *StockRelationshipPub) error {
	*o = StockRelationshipPub{}
	return nil
}

func stockRelationshipPubAfterInsertHook(e boil.Executor, o *StockRelationshipPub) error {
	*o = StockRelationshipPub{}
	return nil
}

func stockRelationshipPubAfterSelectHook(e boil.Executor, o *StockRelationshipPub) error {
	*o = StockRelationshipPub{}
	return nil
}

func stockRelationshipPubBeforeUpdateHook(e boil.Executor, o *StockRelationshipPub) error {
	*o = StockRelationshipPub{}
	return nil
}

func stockRelationshipPubAfterUpdateHook(e boil.Executor, o *StockRelationshipPub) error {
	*o = StockRelationshipPub{}
	return nil
}

func stockRelationshipPubBeforeDeleteHook(e boil.Executor, o *StockRelationshipPub) error {
	*o = StockRelationshipPub{}
	return nil
}

func stockRelationshipPubAfterDeleteHook(e boil.Executor, o *StockRelationshipPub) error {
	*o = StockRelationshipPub{}
	return nil
}

func stockRelationshipPubBeforeUpsertHook(e boil.Executor, o *StockRelationshipPub) error {
	*o = StockRelationshipPub{}
	return nil
}

func stockRelationshipPubAfterUpsertHook(e boil.Executor, o *StockRelationshipPub) error {
	*o = StockRelationshipPub{}
	return nil
}

func testStockRelationshipPubsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &StockRelationshipPub{}
	o := &StockRelationshipPub{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockRelationshipPubDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub object: %s", err)
	}

	AddStockRelationshipPubHook(boil.BeforeInsertHook, stockRelationshipPubBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockRelationshipPubBeforeInsertHooks = []StockRelationshipPubHook{}

	AddStockRelationshipPubHook(boil.AfterInsertHook, stockRelationshipPubAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockRelationshipPubAfterInsertHooks = []StockRelationshipPubHook{}

	AddStockRelationshipPubHook(boil.AfterSelectHook, stockRelationshipPubAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockRelationshipPubAfterSelectHooks = []StockRelationshipPubHook{}

	AddStockRelationshipPubHook(boil.BeforeUpdateHook, stockRelationshipPubBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockRelationshipPubBeforeUpdateHooks = []StockRelationshipPubHook{}

	AddStockRelationshipPubHook(boil.AfterUpdateHook, stockRelationshipPubAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockRelationshipPubAfterUpdateHooks = []StockRelationshipPubHook{}

	AddStockRelationshipPubHook(boil.BeforeDeleteHook, stockRelationshipPubBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockRelationshipPubBeforeDeleteHooks = []StockRelationshipPubHook{}

	AddStockRelationshipPubHook(boil.AfterDeleteHook, stockRelationshipPubAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockRelationshipPubAfterDeleteHooks = []StockRelationshipPubHook{}

	AddStockRelationshipPubHook(boil.BeforeUpsertHook, stockRelationshipPubBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockRelationshipPubBeforeUpsertHooks = []StockRelationshipPubHook{}

	AddStockRelationshipPubHook(boil.AfterUpsertHook, stockRelationshipPubAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockRelationshipPubAfterUpsertHooks = []StockRelationshipPubHook{}
}
func testStockRelationshipPubsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipPub := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockRelationshipPubsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipPub := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPub.Insert(tx, stockRelationshipPubColumns...); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockRelationshipPubToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockRelationshipPub
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockRelationshipPubDBTypes, true, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
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

	slice := StockRelationshipPubSlice{&local}
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

func testStockRelationshipPubToOneStockRelationshipUsingStockRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockRelationshipPub
	var foreign StockRelationship

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockRelationshipPubDBTypes, true, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.StockRelationshipID = foreign.StockRelationshipID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockRelationshipID != foreign.StockRelationshipID {
		t.Errorf("want: %v, got %v", foreign.StockRelationshipID, check.StockRelationshipID)
	}

	slice := StockRelationshipPubSlice{&local}
	if err = local.L.LoadStockRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockRelationship = nil
	if err = local.L.LoadStockRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockRelationshipPubToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockRelationshipPub
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockRelationshipPubDBTypes, false, strmangle.SetComplement(stockRelationshipPubPrimaryKeyColumns, stockRelationshipPubColumnsWithoutDefault)...); err != nil {
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

		if x.R.StockRelationshipPub != &a {
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
func testStockRelationshipPubToOneSetOpStockRelationshipUsingStockRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockRelationshipPub
	var b, c StockRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockRelationshipPubDBTypes, false, strmangle.SetComplement(stockRelationshipPubPrimaryKeyColumns, stockRelationshipPubColumnsWithoutDefault)...); err != nil {
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
		err = a.SetStockRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockRelationship != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.StockRelationshipPub != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StockRelationshipID != x.StockRelationshipID {
			t.Error("foreign key was wrong value", a.StockRelationshipID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.StockRelationshipID))
		reflect.Indirect(reflect.ValueOf(&a.StockRelationshipID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockRelationshipID != x.StockRelationshipID {
			t.Error("foreign key was wrong value", a.StockRelationshipID, x.StockRelationshipID)
		}
	}
}
func testStockRelationshipPubsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipPub := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockRelationshipPub.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockRelationshipPubsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipPub := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockRelationshipPubSlice{stockRelationshipPub}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockRelationshipPubsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipPub := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockRelationshipPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockRelationshipPubDBTypes = map[string]string{"PubID": "integer", "StockRelationshipID": "integer", "StockRelationshipPubID": "integer"}
	_                           = bytes.MinRead
)

func testStockRelationshipPubsUpdate(t *testing.T) {
	t.Parallel()

	if len(stockRelationshipPubColumns) == len(stockRelationshipPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockRelationshipPub := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	if err = stockRelationshipPub.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockRelationshipPubsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockRelationshipPubColumns) == len(stockRelationshipPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockRelationshipPub := &StockRelationshipPub{}
	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockRelationshipPub, stockRelationshipPubDBTypes, true, stockRelationshipPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockRelationshipPubColumns, stockRelationshipPubPrimaryKeyColumns) {
		fields = stockRelationshipPubColumns
	} else {
		fields = strmangle.SetComplement(
			stockRelationshipPubColumns,
			stockRelationshipPubPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockRelationshipPub))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockRelationshipPubSlice{stockRelationshipPub}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockRelationshipPubsUpsert(t *testing.T) {
	t.Parallel()

	if len(stockRelationshipPubColumns) == len(stockRelationshipPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockRelationshipPub := StockRelationshipPub{}
	if err = randomize.Struct(seed, &stockRelationshipPub, stockRelationshipPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipPub.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockRelationshipPub: %s", err)
	}

	count, err := StockRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockRelationshipPub, stockRelationshipPubDBTypes, false, stockRelationshipPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}

	if err = stockRelationshipPub.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockRelationshipPub: %s", err)
	}

	count, err = StockRelationshipPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
