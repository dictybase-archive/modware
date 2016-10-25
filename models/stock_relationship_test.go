package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockRelationships(t *testing.T) {
	t.Parallel()

	query := StockRelationships(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockRelationshipsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationship := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockRelationship.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockRelationshipsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationship := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockRelationships(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := StockRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockRelationshipsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationship := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockRelationshipSlice{stockRelationship}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockRelationshipsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationship := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockRelationshipExists(tx, stockRelationship.StockRelationshipID)
	if err != nil {
		t.Errorf("Unable to check if StockRelationship exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockRelationshipExistsG to return true, but got false.")
	}
}
func testStockRelationshipsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationship := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	stockRelationshipFound, err := FindStockRelationship(tx, stockRelationship.StockRelationshipID)
	if err != nil {
		t.Error(err)
	}

	if stockRelationshipFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockRelationshipsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationship := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockRelationships(tx).Bind(stockRelationship); err != nil {
		t.Error(err)
	}
}

func testStockRelationshipsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationship := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := StockRelationships(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockRelationshipsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipOne := &StockRelationship{}
	stockRelationshipTwo := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationshipOne, stockRelationshipDBTypes, false, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}
	if err = randomize.Struct(seed, stockRelationshipTwo, stockRelationshipDBTypes, false, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockRelationshipTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockRelationships(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockRelationshipsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockRelationshipOne := &StockRelationship{}
	stockRelationshipTwo := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationshipOne, stockRelationshipDBTypes, false, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}
	if err = randomize.Struct(seed, stockRelationshipTwo, stockRelationshipDBTypes, false, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockRelationshipTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockRelationshipBeforeInsertHook(e boil.Executor, o *StockRelationship) error {
	*o = StockRelationship{}
	return nil
}

func stockRelationshipAfterInsertHook(e boil.Executor, o *StockRelationship) error {
	*o = StockRelationship{}
	return nil
}

func stockRelationshipAfterSelectHook(e boil.Executor, o *StockRelationship) error {
	*o = StockRelationship{}
	return nil
}

func stockRelationshipBeforeUpdateHook(e boil.Executor, o *StockRelationship) error {
	*o = StockRelationship{}
	return nil
}

func stockRelationshipAfterUpdateHook(e boil.Executor, o *StockRelationship) error {
	*o = StockRelationship{}
	return nil
}

func stockRelationshipBeforeDeleteHook(e boil.Executor, o *StockRelationship) error {
	*o = StockRelationship{}
	return nil
}

func stockRelationshipAfterDeleteHook(e boil.Executor, o *StockRelationship) error {
	*o = StockRelationship{}
	return nil
}

func stockRelationshipBeforeUpsertHook(e boil.Executor, o *StockRelationship) error {
	*o = StockRelationship{}
	return nil
}

func stockRelationshipAfterUpsertHook(e boil.Executor, o *StockRelationship) error {
	*o = StockRelationship{}
	return nil
}

func testStockRelationshipsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &StockRelationship{}
	o := &StockRelationship{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockRelationshipDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StockRelationship object: %s", err)
	}

	AddStockRelationshipHook(boil.BeforeInsertHook, stockRelationshipBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockRelationshipBeforeInsertHooks = []StockRelationshipHook{}

	AddStockRelationshipHook(boil.AfterInsertHook, stockRelationshipAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockRelationshipAfterInsertHooks = []StockRelationshipHook{}

	AddStockRelationshipHook(boil.AfterSelectHook, stockRelationshipAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockRelationshipAfterSelectHooks = []StockRelationshipHook{}

	AddStockRelationshipHook(boil.BeforeUpdateHook, stockRelationshipBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockRelationshipBeforeUpdateHooks = []StockRelationshipHook{}

	AddStockRelationshipHook(boil.AfterUpdateHook, stockRelationshipAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockRelationshipAfterUpdateHooks = []StockRelationshipHook{}

	AddStockRelationshipHook(boil.BeforeDeleteHook, stockRelationshipBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockRelationshipBeforeDeleteHooks = []StockRelationshipHook{}

	AddStockRelationshipHook(boil.AfterDeleteHook, stockRelationshipAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockRelationshipAfterDeleteHooks = []StockRelationshipHook{}

	AddStockRelationshipHook(boil.BeforeUpsertHook, stockRelationshipBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockRelationshipBeforeUpsertHooks = []StockRelationshipHook{}

	AddStockRelationshipHook(boil.AfterUpsertHook, stockRelationshipAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockRelationshipAfterUpsertHooks = []StockRelationshipHook{}
}
func testStockRelationshipsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationship := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockRelationshipsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationship := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationship.Insert(tx, stockRelationshipColumns...); err != nil {
		t.Error(err)
	}

	count, err := StockRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockRelationshipOneToOneStockRelationshipPubUsingStockRelationshipPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockRelationshipPub
	var local StockRelationship

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockRelationshipPubDBTypes, true, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.StockRelationshipID = local.StockRelationshipID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockRelationshipPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockRelationshipID != foreign.StockRelationshipID {
		t.Errorf("want: %v, got %v", foreign.StockRelationshipID, check.StockRelationshipID)
	}

	slice := StockRelationshipSlice{&local}
	if err = local.L.LoadStockRelationshipPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockRelationshipPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockRelationshipPub = nil
	if err = local.L.LoadStockRelationshipPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockRelationshipPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockRelationshipOneToOneSetOpStockRelationshipPubUsingStockRelationshipPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockRelationship
	var b, c StockRelationshipPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockRelationshipDBTypes, false, strmangle.SetComplement(stockRelationshipPrimaryKeyColumns, stockRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockRelationshipPubDBTypes, false, strmangle.SetComplement(stockRelationshipPubPrimaryKeyColumns, stockRelationshipPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockRelationshipPubDBTypes, false, strmangle.SetComplement(stockRelationshipPubPrimaryKeyColumns, stockRelationshipPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockRelationshipPub{&b, &c} {
		err = a.SetStockRelationshipPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockRelationshipPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.StockRelationship != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.StockRelationshipID != x.StockRelationshipID {
			t.Error("foreign key was wrong value", a.StockRelationshipID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.StockRelationshipID))
		reflect.Indirect(reflect.ValueOf(&x.StockRelationshipID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockRelationshipID != x.StockRelationshipID {
			t.Error("foreign key was wrong value", a.StockRelationshipID, x.StockRelationshipID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testStockRelationshipToManyStockRelationshipCvterms(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockRelationship
	var b, c StockRelationshipCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, stockRelationshipCvtermDBTypes, false, stockRelationshipCvtermColumnsWithDefault...)
	randomize.Struct(seed, &c, stockRelationshipCvtermDBTypes, false, stockRelationshipCvtermColumnsWithDefault...)

	b.StockRelationshipID = a.StockRelationshipID
	c.StockRelationshipID = a.StockRelationshipID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	stockRelationshipCvterm, err := a.StockRelationshipCvterms(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range stockRelationshipCvterm {
		if v.StockRelationshipID == b.StockRelationshipID {
			bFound = true
		}
		if v.StockRelationshipID == c.StockRelationshipID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := StockRelationshipSlice{&a}
	if err = a.L.LoadStockRelationshipCvterms(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.StockRelationshipCvterms); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.StockRelationshipCvterms = nil
	if err = a.L.LoadStockRelationshipCvterms(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.StockRelationshipCvterms); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", stockRelationshipCvterm)
	}
}

func testStockRelationshipToManyAddOpStockRelationshipCvterms(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockRelationship
	var b, c, d, e StockRelationshipCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockRelationshipDBTypes, false, strmangle.SetComplement(stockRelationshipPrimaryKeyColumns, stockRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*StockRelationshipCvterm{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, stockRelationshipCvtermDBTypes, false, strmangle.SetComplement(stockRelationshipCvtermPrimaryKeyColumns, stockRelationshipCvtermColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*StockRelationshipCvterm{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddStockRelationshipCvterms(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.StockRelationshipID != first.StockRelationshipID {
			t.Error("foreign key was wrong value", a.StockRelationshipID, first.StockRelationshipID)
		}
		if a.StockRelationshipID != second.StockRelationshipID {
			t.Error("foreign key was wrong value", a.StockRelationshipID, second.StockRelationshipID)
		}

		if first.R.StockRelationship != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.StockRelationship != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.StockRelationshipCvterms[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.StockRelationshipCvterms[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.StockRelationshipCvterms(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testStockRelationshipToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockRelationship
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
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

	slice := StockRelationshipSlice{&local}
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

func testStockRelationshipToOneStockUsingObject(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockRelationship
	var foreign Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ObjectID = foreign.StockID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Object(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockID != foreign.StockID {
		t.Errorf("want: %v, got %v", foreign.StockID, check.StockID)
	}

	slice := StockRelationshipSlice{&local}
	if err = local.L.LoadObject(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Object == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Object = nil
	if err = local.L.LoadObject(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Object == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockRelationshipToOneStockUsingSubject(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockRelationship
	var foreign Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.SubjectID = foreign.StockID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Subject(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockID != foreign.StockID {
		t.Errorf("want: %v, got %v", foreign.StockID, check.StockID)
	}

	slice := StockRelationshipSlice{&local}
	if err = local.L.LoadSubject(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Subject == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Subject = nil
	if err = local.L.LoadSubject(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Subject == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockRelationshipToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockRelationship
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockRelationshipDBTypes, false, strmangle.SetComplement(stockRelationshipPrimaryKeyColumns, stockRelationshipColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeStockRelationship != &a {
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
func testStockRelationshipToOneSetOpStockUsingObject(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockRelationship
	var b, c Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockRelationshipDBTypes, false, strmangle.SetComplement(stockRelationshipPrimaryKeyColumns, stockRelationshipColumnsWithoutDefault)...); err != nil {
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
		err = a.SetObject(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Object != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ObjectStockRelationship != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ObjectID != x.StockID {
			t.Error("foreign key was wrong value", a.ObjectID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ObjectID))
		reflect.Indirect(reflect.ValueOf(&a.ObjectID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ObjectID != x.StockID {
			t.Error("foreign key was wrong value", a.ObjectID, x.StockID)
		}
	}
}
func testStockRelationshipToOneSetOpStockUsingSubject(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockRelationship
	var b, c Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockRelationshipDBTypes, false, strmangle.SetComplement(stockRelationshipPrimaryKeyColumns, stockRelationshipColumnsWithoutDefault)...); err != nil {
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
		err = a.SetSubject(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Subject != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SubjectStockRelationship != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SubjectID != x.StockID {
			t.Error("foreign key was wrong value", a.SubjectID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SubjectID))
		reflect.Indirect(reflect.ValueOf(&a.SubjectID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SubjectID != x.StockID {
			t.Error("foreign key was wrong value", a.SubjectID, x.StockID)
		}
	}
}
func testStockRelationshipsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationship := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockRelationship.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockRelationshipsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationship := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockRelationshipSlice{stockRelationship}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockRelationshipsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationship := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockRelationships(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockRelationshipDBTypes = map[string]string{"ObjectID": "integer", "Rank": "integer", "StockRelationshipID": "integer", "SubjectID": "integer", "TypeID": "integer", "Value": "text"}
	_                        = bytes.MinRead
)

func testStockRelationshipsUpdate(t *testing.T) {
	t.Parallel()

	if len(stockRelationshipColumns) == len(stockRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockRelationship := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	if err = stockRelationship.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockRelationshipsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockRelationshipColumns) == len(stockRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockRelationship := &StockRelationship{}
	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockRelationship, stockRelationshipDBTypes, true, stockRelationshipPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockRelationshipColumns, stockRelationshipPrimaryKeyColumns) {
		fields = stockRelationshipColumns
	} else {
		fields = strmangle.SetComplement(
			stockRelationshipColumns,
			stockRelationshipPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockRelationship))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockRelationshipSlice{stockRelationship}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockRelationshipsUpsert(t *testing.T) {
	t.Parallel()

	if len(stockRelationshipColumns) == len(stockRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockRelationship := StockRelationship{}
	if err = randomize.Struct(seed, &stockRelationship, stockRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationship.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockRelationship: %s", err)
	}

	count, err := StockRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockRelationship, stockRelationshipDBTypes, false, stockRelationshipPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}

	if err = stockRelationship.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockRelationship: %s", err)
	}

	count, err = StockRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
