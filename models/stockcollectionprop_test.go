package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockcollectionprops(t *testing.T) {
	t.Parallel()

	query := Stockcollectionprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockcollectionpropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionprop := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockcollectionprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockcollectionprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockcollectionpropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionprop := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Stockcollectionprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Stockcollectionprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockcollectionpropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionprop := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockcollectionpropSlice{stockcollectionprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockcollectionprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockcollectionpropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionprop := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockcollectionpropExists(tx, stockcollectionprop.StockcollectionpropID)
	if err != nil {
		t.Errorf("Unable to check if Stockcollectionprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockcollectionpropExistsG to return true, but got false.")
	}
}
func testStockcollectionpropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionprop := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionprop.Insert(tx); err != nil {
		t.Error(err)
	}

	stockcollectionpropFound, err := FindStockcollectionprop(tx, stockcollectionprop.StockcollectionpropID)
	if err != nil {
		t.Error(err)
	}

	if stockcollectionpropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockcollectionpropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionprop := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Stockcollectionprops(tx).Bind(stockcollectionprop); err != nil {
		t.Error(err)
	}
}

func testStockcollectionpropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionprop := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Stockcollectionprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockcollectionpropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionpropOne := &Stockcollectionprop{}
	stockcollectionpropTwo := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionpropOne, stockcollectionpropDBTypes, false, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}
	if err = randomize.Struct(seed, stockcollectionpropTwo, stockcollectionpropDBTypes, false, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockcollectionpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Stockcollectionprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockcollectionpropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockcollectionpropOne := &Stockcollectionprop{}
	stockcollectionpropTwo := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionpropOne, stockcollectionpropDBTypes, false, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}
	if err = randomize.Struct(seed, stockcollectionpropTwo, stockcollectionpropDBTypes, false, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockcollectionpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockcollectionprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockcollectionpropBeforeInsertHook(e boil.Executor, o *Stockcollectionprop) error {
	*o = Stockcollectionprop{}
	return nil
}

func stockcollectionpropAfterInsertHook(e boil.Executor, o *Stockcollectionprop) error {
	*o = Stockcollectionprop{}
	return nil
}

func stockcollectionpropAfterSelectHook(e boil.Executor, o *Stockcollectionprop) error {
	*o = Stockcollectionprop{}
	return nil
}

func stockcollectionpropBeforeUpdateHook(e boil.Executor, o *Stockcollectionprop) error {
	*o = Stockcollectionprop{}
	return nil
}

func stockcollectionpropAfterUpdateHook(e boil.Executor, o *Stockcollectionprop) error {
	*o = Stockcollectionprop{}
	return nil
}

func stockcollectionpropBeforeDeleteHook(e boil.Executor, o *Stockcollectionprop) error {
	*o = Stockcollectionprop{}
	return nil
}

func stockcollectionpropAfterDeleteHook(e boil.Executor, o *Stockcollectionprop) error {
	*o = Stockcollectionprop{}
	return nil
}

func stockcollectionpropBeforeUpsertHook(e boil.Executor, o *Stockcollectionprop) error {
	*o = Stockcollectionprop{}
	return nil
}

func stockcollectionpropAfterUpsertHook(e boil.Executor, o *Stockcollectionprop) error {
	*o = Stockcollectionprop{}
	return nil
}

func testStockcollectionpropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Stockcollectionprop{}
	o := &Stockcollectionprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockcollectionpropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop object: %s", err)
	}

	AddStockcollectionpropHook(boil.BeforeInsertHook, stockcollectionpropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockcollectionpropBeforeInsertHooks = []StockcollectionpropHook{}

	AddStockcollectionpropHook(boil.AfterInsertHook, stockcollectionpropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockcollectionpropAfterInsertHooks = []StockcollectionpropHook{}

	AddStockcollectionpropHook(boil.AfterSelectHook, stockcollectionpropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockcollectionpropAfterSelectHooks = []StockcollectionpropHook{}

	AddStockcollectionpropHook(boil.BeforeUpdateHook, stockcollectionpropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockcollectionpropBeforeUpdateHooks = []StockcollectionpropHook{}

	AddStockcollectionpropHook(boil.AfterUpdateHook, stockcollectionpropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockcollectionpropAfterUpdateHooks = []StockcollectionpropHook{}

	AddStockcollectionpropHook(boil.BeforeDeleteHook, stockcollectionpropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockcollectionpropBeforeDeleteHooks = []StockcollectionpropHook{}

	AddStockcollectionpropHook(boil.AfterDeleteHook, stockcollectionpropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockcollectionpropAfterDeleteHooks = []StockcollectionpropHook{}

	AddStockcollectionpropHook(boil.BeforeUpsertHook, stockcollectionpropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockcollectionpropBeforeUpsertHooks = []StockcollectionpropHook{}

	AddStockcollectionpropHook(boil.AfterUpsertHook, stockcollectionpropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockcollectionpropAfterUpsertHooks = []StockcollectionpropHook{}
}
func testStockcollectionpropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionprop := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockcollectionprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockcollectionpropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionprop := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionprop.Insert(tx, stockcollectionpropColumns...); err != nil {
		t.Error(err)
	}

	count, err := Stockcollectionprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockcollectionpropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Stockcollectionprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockcollectionpropDBTypes, true, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
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

	slice := StockcollectionpropSlice{&local}
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

func testStockcollectionpropToOneStockcollectionUsingStockcollection(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Stockcollectionprop
	var foreign Stockcollection

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockcollectionpropDBTypes, true, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
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

	slice := StockcollectionpropSlice{&local}
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

func testStockcollectionpropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stockcollectionprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockcollectionpropDBTypes, false, strmangle.SetComplement(stockcollectionpropPrimaryKeyColumns, stockcollectionpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeStockcollectionprop != &a {
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
func testStockcollectionpropToOneSetOpStockcollectionUsingStockcollection(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Stockcollectionprop
	var b, c Stockcollection

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockcollectionpropDBTypes, false, strmangle.SetComplement(stockcollectionpropPrimaryKeyColumns, stockcollectionpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.Stockcollectionprop != &a {
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
func testStockcollectionpropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionprop := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockcollectionprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockcollectionpropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionprop := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockcollectionpropSlice{stockcollectionprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockcollectionpropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockcollectionprop := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Stockcollectionprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockcollectionpropDBTypes = map[string]string{"Rank": "integer", "StockcollectionID": "integer", "StockcollectionpropID": "integer", "TypeID": "integer", "Value": "text"}
	_                          = bytes.MinRead
)

func testStockcollectionpropsUpdate(t *testing.T) {
	t.Parallel()

	if len(stockcollectionpropColumns) == len(stockcollectionpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockcollectionprop := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockcollectionprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	if err = stockcollectionprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockcollectionpropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockcollectionpropColumns) == len(stockcollectionpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockcollectionprop := &Stockcollectionprop{}
	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stockcollectionprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockcollectionprop, stockcollectionpropDBTypes, true, stockcollectionpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockcollectionpropColumns, stockcollectionpropPrimaryKeyColumns) {
		fields = stockcollectionpropColumns
	} else {
		fields = strmangle.SetComplement(
			stockcollectionpropColumns,
			stockcollectionpropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockcollectionprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockcollectionpropSlice{stockcollectionprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockcollectionpropsUpsert(t *testing.T) {
	t.Parallel()

	if len(stockcollectionpropColumns) == len(stockcollectionpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockcollectionprop := Stockcollectionprop{}
	if err = randomize.Struct(seed, &stockcollectionprop, stockcollectionpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockcollectionprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Stockcollectionprop: %s", err)
	}

	count, err := Stockcollectionprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockcollectionprop, stockcollectionpropDBTypes, false, stockcollectionpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}

	if err = stockcollectionprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Stockcollectionprop: %s", err)
	}

	count, err = Stockcollectionprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
