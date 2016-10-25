package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockRelationshipCvterms(t *testing.T) {
	t.Parallel()

	query := StockRelationshipCvterms(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockRelationshipCvtermsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipCvterm := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockRelationshipCvterm.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockRelationshipCvtermsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipCvterm := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockRelationshipCvterms(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockRelationshipCvtermsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipCvterm := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockRelationshipCvtermSlice{stockRelationshipCvterm}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockRelationshipCvtermsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipCvterm := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockRelationshipCvtermExists(tx, stockRelationshipCvterm.StockRelationshipCvtermID)
	if err != nil {
		t.Errorf("Unable to check if StockRelationshipCvterm exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockRelationshipCvtermExistsG to return true, but got false.")
	}
}
func testStockRelationshipCvtermsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipCvterm := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	stockRelationshipCvtermFound, err := FindStockRelationshipCvterm(tx, stockRelationshipCvterm.StockRelationshipCvtermID)
	if err != nil {
		t.Error(err)
	}

	if stockRelationshipCvtermFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockRelationshipCvtermsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipCvterm := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockRelationshipCvterms(tx).Bind(stockRelationshipCvterm); err != nil {
		t.Error(err)
	}
}

func testStockRelationshipCvtermsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipCvterm := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := StockRelationshipCvterms(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockRelationshipCvtermsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipCvtermOne := &StockRelationshipCvterm{}
	stockRelationshipCvtermTwo := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvtermOne, stockRelationshipCvtermDBTypes, false, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}
	if err = randomize.Struct(seed, stockRelationshipCvtermTwo, stockRelationshipCvtermDBTypes, false, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvtermOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockRelationshipCvtermTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockRelationshipCvterms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockRelationshipCvtermsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockRelationshipCvtermOne := &StockRelationshipCvterm{}
	stockRelationshipCvtermTwo := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvtermOne, stockRelationshipCvtermDBTypes, false, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}
	if err = randomize.Struct(seed, stockRelationshipCvtermTwo, stockRelationshipCvtermDBTypes, false, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvtermOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockRelationshipCvtermTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockRelationshipCvtermBeforeInsertHook(e boil.Executor, o *StockRelationshipCvterm) error {
	*o = StockRelationshipCvterm{}
	return nil
}

func stockRelationshipCvtermAfterInsertHook(e boil.Executor, o *StockRelationshipCvterm) error {
	*o = StockRelationshipCvterm{}
	return nil
}

func stockRelationshipCvtermAfterSelectHook(e boil.Executor, o *StockRelationshipCvterm) error {
	*o = StockRelationshipCvterm{}
	return nil
}

func stockRelationshipCvtermBeforeUpdateHook(e boil.Executor, o *StockRelationshipCvterm) error {
	*o = StockRelationshipCvterm{}
	return nil
}

func stockRelationshipCvtermAfterUpdateHook(e boil.Executor, o *StockRelationshipCvterm) error {
	*o = StockRelationshipCvterm{}
	return nil
}

func stockRelationshipCvtermBeforeDeleteHook(e boil.Executor, o *StockRelationshipCvterm) error {
	*o = StockRelationshipCvterm{}
	return nil
}

func stockRelationshipCvtermAfterDeleteHook(e boil.Executor, o *StockRelationshipCvterm) error {
	*o = StockRelationshipCvterm{}
	return nil
}

func stockRelationshipCvtermBeforeUpsertHook(e boil.Executor, o *StockRelationshipCvterm) error {
	*o = StockRelationshipCvterm{}
	return nil
}

func stockRelationshipCvtermAfterUpsertHook(e boil.Executor, o *StockRelationshipCvterm) error {
	*o = StockRelationshipCvterm{}
	return nil
}

func testStockRelationshipCvtermsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &StockRelationshipCvterm{}
	o := &StockRelationshipCvterm{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockRelationshipCvtermDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm object: %s", err)
	}

	AddStockRelationshipCvtermHook(boil.BeforeInsertHook, stockRelationshipCvtermBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockRelationshipCvtermBeforeInsertHooks = []StockRelationshipCvtermHook{}

	AddStockRelationshipCvtermHook(boil.AfterInsertHook, stockRelationshipCvtermAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockRelationshipCvtermAfterInsertHooks = []StockRelationshipCvtermHook{}

	AddStockRelationshipCvtermHook(boil.AfterSelectHook, stockRelationshipCvtermAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockRelationshipCvtermAfterSelectHooks = []StockRelationshipCvtermHook{}

	AddStockRelationshipCvtermHook(boil.BeforeUpdateHook, stockRelationshipCvtermBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockRelationshipCvtermBeforeUpdateHooks = []StockRelationshipCvtermHook{}

	AddStockRelationshipCvtermHook(boil.AfterUpdateHook, stockRelationshipCvtermAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockRelationshipCvtermAfterUpdateHooks = []StockRelationshipCvtermHook{}

	AddStockRelationshipCvtermHook(boil.BeforeDeleteHook, stockRelationshipCvtermBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockRelationshipCvtermBeforeDeleteHooks = []StockRelationshipCvtermHook{}

	AddStockRelationshipCvtermHook(boil.AfterDeleteHook, stockRelationshipCvtermAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockRelationshipCvtermAfterDeleteHooks = []StockRelationshipCvtermHook{}

	AddStockRelationshipCvtermHook(boil.BeforeUpsertHook, stockRelationshipCvtermBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockRelationshipCvtermBeforeUpsertHooks = []StockRelationshipCvtermHook{}

	AddStockRelationshipCvtermHook(boil.AfterUpsertHook, stockRelationshipCvtermAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockRelationshipCvtermAfterUpsertHooks = []StockRelationshipCvtermHook{}
}
func testStockRelationshipCvtermsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipCvterm := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockRelationshipCvtermsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipCvterm := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvterm.Insert(tx, stockRelationshipCvtermColumns...); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockRelationshipCvtermToOneCvtermUsingCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockRelationshipCvterm
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockRelationshipCvtermDBTypes, true, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
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

	slice := StockRelationshipCvtermSlice{&local}
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

func testStockRelationshipCvtermToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockRelationshipCvterm
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockRelationshipCvtermDBTypes, true, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	local.PubID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.PubID.Int = foreign.PubID
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

	slice := StockRelationshipCvtermSlice{&local}
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

func testStockRelationshipCvtermToOneStockRelationshipUsingStockRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockRelationshipCvterm
	var foreign StockRelationship

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockRelationshipCvtermDBTypes, true, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
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

	slice := StockRelationshipCvtermSlice{&local}
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

func testStockRelationshipCvtermToOneSetOpCvtermUsingCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockRelationshipCvterm
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockRelationshipCvtermDBTypes, false, strmangle.SetComplement(stockRelationshipCvtermPrimaryKeyColumns, stockRelationshipCvtermColumnsWithoutDefault)...); err != nil {
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

		if x.R.StockRelationshipCvterms[0] != &a {
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
func testStockRelationshipCvtermToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockRelationshipCvterm
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockRelationshipCvtermDBTypes, false, strmangle.SetComplement(stockRelationshipCvtermPrimaryKeyColumns, stockRelationshipCvtermColumnsWithoutDefault)...); err != nil {
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

		if x.R.StockRelationshipCvterms[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PubID.Int != x.PubID {
			t.Error("foreign key was wrong value", a.PubID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PubID.Int))
		reflect.Indirect(reflect.ValueOf(&a.PubID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID.Int != x.PubID {
			t.Error("foreign key was wrong value", a.PubID.Int, x.PubID)
		}
	}
}

func testStockRelationshipCvtermToOneRemoveOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockRelationshipCvterm
	var b Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockRelationshipCvtermDBTypes, false, strmangle.SetComplement(stockRelationshipCvtermPrimaryKeyColumns, stockRelationshipCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetPub(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemovePub(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Pub(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Pub != nil {
		t.Error("R struct entry should be nil")
	}

	if a.PubID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.StockRelationshipCvterms) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testStockRelationshipCvtermToOneSetOpStockRelationshipUsingStockRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockRelationshipCvterm
	var b, c StockRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockRelationshipCvtermDBTypes, false, strmangle.SetComplement(stockRelationshipCvtermPrimaryKeyColumns, stockRelationshipCvtermColumnsWithoutDefault)...); err != nil {
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

		if x.R.StockRelationshipCvterms[0] != &a {
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
func testStockRelationshipCvtermsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipCvterm := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockRelationshipCvterm.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockRelationshipCvtermsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipCvterm := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockRelationshipCvtermSlice{stockRelationshipCvterm}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockRelationshipCvtermsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockRelationshipCvterm := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockRelationshipCvterms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockRelationshipCvtermDBTypes = map[string]string{"CvtermID": "integer", "PubID": "integer", "StockRelationshipCvtermID": "integer", "StockRelationshipID": "integer"}
	_                              = bytes.MinRead
)

func testStockRelationshipCvtermsUpdate(t *testing.T) {
	t.Parallel()

	if len(stockRelationshipCvtermColumns) == len(stockRelationshipCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockRelationshipCvterm := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true, stockRelationshipCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	if err = stockRelationshipCvterm.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockRelationshipCvtermsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockRelationshipCvtermColumns) == len(stockRelationshipCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockRelationshipCvterm := &StockRelationshipCvterm{}
	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockRelationshipCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true, stockRelationshipCvtermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockRelationshipCvtermColumns, stockRelationshipCvtermPrimaryKeyColumns) {
		fields = stockRelationshipCvtermColumns
	} else {
		fields = strmangle.SetComplement(
			stockRelationshipCvtermColumns,
			stockRelationshipCvtermPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockRelationshipCvterm))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockRelationshipCvtermSlice{stockRelationshipCvterm}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockRelationshipCvtermsUpsert(t *testing.T) {
	t.Parallel()

	if len(stockRelationshipCvtermColumns) == len(stockRelationshipCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockRelationshipCvterm := StockRelationshipCvterm{}
	if err = randomize.Struct(seed, &stockRelationshipCvterm, stockRelationshipCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockRelationshipCvterm.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockRelationshipCvterm: %s", err)
	}

	count, err := StockRelationshipCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockRelationshipCvterm, stockRelationshipCvtermDBTypes, false, stockRelationshipCvtermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipCvterm struct: %s", err)
	}

	if err = stockRelationshipCvterm.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockRelationshipCvterm: %s", err)
	}

	count, err = StockRelationshipCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
