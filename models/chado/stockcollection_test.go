package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockcollections(t *testing.T) {
	t.Parallel()

	query := Stockcollections(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockcollectionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollection := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollection.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockcollection.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockcollections(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockcollectionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollection := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollection.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Stockcollections(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Stockcollections(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockcollectionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollection := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollection.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockcollectionSlice{stockcollection}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockcollections(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockcollectionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollection := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollection.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockcollectionExists(tx, stockcollection.StockcollectionID)
	if err != nil {
		t.Errorf("Unable to check if Stockcollection exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockcollectionExistsG to return true, but got false.")
	}
}
func testStockcollectionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollection := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollection.Insert(tx); err != nil {
		t.Error(err)
	}

	stockcollectionFound, err := FindStockcollection(tx, stockcollection.StockcollectionID)
	if err != nil {
		t.Error(err)
	}

	if stockcollectionFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockcollectionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollection := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollection.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Stockcollections(tx).Bind(stockcollection); err != nil {
		t.Error(err)
	}
}

func testStockcollectionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollection := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollection.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Stockcollections(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockcollectionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionOne := &Stockcollection{}
	stockcollectionTwo := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollectionOne, stockcollectionDBTypes, false, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}
	if err = randomize.Struct(seed, stockcollectionTwo, stockcollectionDBTypes, false, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockcollectionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Stockcollections(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockcollectionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockcollectionOne := &Stockcollection{}
	stockcollectionTwo := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollectionOne, stockcollectionDBTypes, false, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}
	if err = randomize.Struct(seed, stockcollectionTwo, stockcollectionDBTypes, false, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockcollectionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockcollections(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockcollectionBeforeInsertHook(e boil.Executor, o *Stockcollection) error {
	*o = Stockcollection{}
	return nil
}

func stockcollectionAfterInsertHook(e boil.Executor, o *Stockcollection) error {
	*o = Stockcollection{}
	return nil
}

func stockcollectionAfterSelectHook(e boil.Executor, o *Stockcollection) error {
	*o = Stockcollection{}
	return nil
}

func stockcollectionBeforeUpdateHook(e boil.Executor, o *Stockcollection) error {
	*o = Stockcollection{}
	return nil
}

func stockcollectionAfterUpdateHook(e boil.Executor, o *Stockcollection) error {
	*o = Stockcollection{}
	return nil
}

func stockcollectionBeforeDeleteHook(e boil.Executor, o *Stockcollection) error {
	*o = Stockcollection{}
	return nil
}

func stockcollectionAfterDeleteHook(e boil.Executor, o *Stockcollection) error {
	*o = Stockcollection{}
	return nil
}

func stockcollectionBeforeUpsertHook(e boil.Executor, o *Stockcollection) error {
	*o = Stockcollection{}
	return nil
}

func stockcollectionAfterUpsertHook(e boil.Executor, o *Stockcollection) error {
	*o = Stockcollection{}
	return nil
}

func testStockcollectionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Stockcollection{}
	o := &Stockcollection{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockcollectionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Stockcollection object: %s", err)
	}

	AddStockcollectionHook(boil.BeforeInsertHook, stockcollectionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockcollectionBeforeInsertHooks = []StockcollectionHook{}

	AddStockcollectionHook(boil.AfterInsertHook, stockcollectionAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockcollectionAfterInsertHooks = []StockcollectionHook{}

	AddStockcollectionHook(boil.AfterSelectHook, stockcollectionAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockcollectionAfterSelectHooks = []StockcollectionHook{}

	AddStockcollectionHook(boil.BeforeUpdateHook, stockcollectionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockcollectionBeforeUpdateHooks = []StockcollectionHook{}

	AddStockcollectionHook(boil.AfterUpdateHook, stockcollectionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockcollectionAfterUpdateHooks = []StockcollectionHook{}

	AddStockcollectionHook(boil.BeforeDeleteHook, stockcollectionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockcollectionBeforeDeleteHooks = []StockcollectionHook{}

	AddStockcollectionHook(boil.AfterDeleteHook, stockcollectionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockcollectionAfterDeleteHooks = []StockcollectionHook{}

	AddStockcollectionHook(boil.BeforeUpsertHook, stockcollectionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockcollectionBeforeUpsertHooks = []StockcollectionHook{}

	AddStockcollectionHook(boil.AfterUpsertHook, stockcollectionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockcollectionAfterUpsertHooks = []StockcollectionHook{}
}
func testStockcollectionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollection := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollection.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockcollections(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockcollectionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollection := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollection.Insert(tx, stockcollectionColumns...); err != nil {
		t.Error(err)
	}

	count, err := Stockcollections(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockcollectionOneToOneStockcollectionStockUsingStockcollectionStock(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockcollectionStock
	var local Stockcollection

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockcollectionStockDBTypes, true, stockcollectionStockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockcollectionStock struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockcollectionDBTypes, true, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.StockcollectionID = local.StockcollectionID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockcollectionStock(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockcollectionID != foreign.StockcollectionID {
		t.Errorf("want: %v, got %v", foreign.StockcollectionID, check.StockcollectionID)
	}

	slice := StockcollectionSlice{&local}
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

func testStockcollectionOneToOneStockcollectionpropUsingStockcollectionprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Stockcollectionprop
	var local Stockcollection

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockcollectionpropDBTypes, true, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockcollectionDBTypes, true, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.StockcollectionID = local.StockcollectionID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Stockcollectionprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockcollectionID != foreign.StockcollectionID {
		t.Errorf("want: %v, got %v", foreign.StockcollectionID, check.StockcollectionID)
	}

	slice := StockcollectionSlice{&local}
	if err = local.L.LoadStockcollectionprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Stockcollectionprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Stockcollectionprop = nil
	if err = local.L.LoadStockcollectionprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Stockcollectionprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockcollectionOneToOneSetOpStockcollectionStockUsingStockcollectionStock(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stockcollection
	var b, c StockcollectionStock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockcollectionDBTypes, false, strmangle.SetComplement(stockcollectionPrimaryKeyColumns, stockcollectionColumnsWithoutDefault)...); err != nil {
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
		if x.R.Stockcollection != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockcollectionID != x.StockcollectionID {
			t.Error("foreign key was wrong value", a.StockcollectionID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.StockcollectionID))
		reflect.Indirect(reflect.ValueOf(&x.StockcollectionID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockcollectionID != x.StockcollectionID {
			t.Error("foreign key was wrong value", a.StockcollectionID, x.StockcollectionID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testStockcollectionOneToOneSetOpStockcollectionpropUsingStockcollectionprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stockcollection
	var b, c Stockcollectionprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockcollectionDBTypes, false, strmangle.SetComplement(stockcollectionPrimaryKeyColumns, stockcollectionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockcollectionpropDBTypes, false, strmangle.SetComplement(stockcollectionpropPrimaryKeyColumns, stockcollectionpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockcollectionpropDBTypes, false, strmangle.SetComplement(stockcollectionpropPrimaryKeyColumns, stockcollectionpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Stockcollectionprop{&b, &c} {
		err = a.SetStockcollectionprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Stockcollectionprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Stockcollection != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockcollectionID != x.StockcollectionID {
			t.Error("foreign key was wrong value", a.StockcollectionID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.StockcollectionID))
		reflect.Indirect(reflect.ValueOf(&x.StockcollectionID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockcollectionID != x.StockcollectionID {
			t.Error("foreign key was wrong value", a.StockcollectionID, x.StockcollectionID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testStockcollectionToOneContactUsingContact(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Stockcollection
	var foreign Contact

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockcollectionDBTypes, true, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	local.ContactID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ContactID.Int = foreign.ContactID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Contact(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ContactID != foreign.ContactID {
		t.Errorf("want: %v, got %v", foreign.ContactID, check.ContactID)
	}

	slice := StockcollectionSlice{&local}
	if err = local.L.LoadContact(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Contact == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Contact = nil
	if err = local.L.LoadContact(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Contact == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockcollectionToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Stockcollection
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockcollectionDBTypes, true, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
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

	slice := StockcollectionSlice{&local}
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

func testStockcollectionToOneSetOpContactUsingContact(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stockcollection
	var b, c Contact

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockcollectionDBTypes, false, strmangle.SetComplement(stockcollectionPrimaryKeyColumns, stockcollectionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Contact{&b, &c} {
		err = a.SetContact(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Contact != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Stockcollections[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ContactID.Int != x.ContactID {
			t.Error("foreign key was wrong value", a.ContactID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ContactID.Int))
		reflect.Indirect(reflect.ValueOf(&a.ContactID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ContactID.Int != x.ContactID {
			t.Error("foreign key was wrong value", a.ContactID.Int, x.ContactID)
		}
	}
}

func testStockcollectionToOneRemoveOpContactUsingContact(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stockcollection
	var b Contact

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockcollectionDBTypes, false, strmangle.SetComplement(stockcollectionPrimaryKeyColumns, stockcollectionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetContact(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveContact(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Contact(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Contact != nil {
		t.Error("R struct entry should be nil")
	}

	if a.ContactID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Stockcollections) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testStockcollectionToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stockcollection
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockcollectionDBTypes, false, strmangle.SetComplement(stockcollectionPrimaryKeyColumns, stockcollectionColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeStockcollection != &a {
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
func testStockcollectionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollection := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollection.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockcollection.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockcollectionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollection := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollection.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockcollectionSlice{stockcollection}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockcollectionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollection := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollection.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Stockcollections(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockcollectionDBTypes = map[string]string{"ContactID": "integer", "Name": "character varying", "StockcollectionID": "integer", "TypeID": "integer", "Uniquename": "text"}
	_                      = bytes.MinRead
)

func testStockcollectionsUpdate(t *testing.T) {
	t.Parallel()

	if len(stockcollectionColumns) == len(stockcollectionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockcollection := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollection.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockcollections(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	if err = stockcollection.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockcollectionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockcollectionColumns) == len(stockcollectionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockcollection := &Stockcollection{}
	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollection.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockcollections(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockcollection, stockcollectionDBTypes, true, stockcollectionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockcollectionColumns, stockcollectionPrimaryKeyColumns) {
		fields = stockcollectionColumns
	} else {
		fields = strmangle.SetComplement(
			stockcollectionColumns,
			stockcollectionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockcollection))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockcollectionSlice{stockcollection}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockcollectionsUpsert(t *testing.T) {
	t.Parallel()

	if len(stockcollectionColumns) == len(stockcollectionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockcollection := Stockcollection{}
	if err = randomize.Struct(seed, &stockcollection, stockcollectionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollection.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Stockcollection: %s", err)
	}

	count, err := Stockcollections(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockcollection, stockcollectionDBTypes, false, stockcollectionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}

	if err = stockcollection.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Stockcollection: %s", err)
	}

	count, err = Stockcollections(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
