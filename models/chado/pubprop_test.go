package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testPubprops(t *testing.T) {
	t.Parallel()

	query := Pubprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testPubpropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubprop := &Pubprop{}
	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = pubprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPubpropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubprop := &Pubprop{}
	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Pubprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Pubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPubpropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubprop := &Pubprop{}
	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PubpropSlice{pubprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testPubpropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubprop := &Pubprop{}
	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := PubpropExists(tx, pubprop.PubpropID)
	if err != nil {
		t.Errorf("Unable to check if Pubprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PubpropExistsG to return true, but got false.")
	}
}
func testPubpropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubprop := &Pubprop{}
	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	pubpropFound, err := FindPubprop(tx, pubprop.PubpropID)
	if err != nil {
		t.Error(err)
	}

	if pubpropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testPubpropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubprop := &Pubprop{}
	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Pubprops(tx).Bind(pubprop); err != nil {
		t.Error(err)
	}
}

func testPubpropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubprop := &Pubprop{}
	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Pubprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPubpropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubpropOne := &Pubprop{}
	pubpropTwo := &Pubprop{}
	if err = randomize.Struct(seed, pubpropOne, pubpropDBTypes, false, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}
	if err = randomize.Struct(seed, pubpropTwo, pubpropDBTypes, false, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = pubpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Pubprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPubpropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	pubpropOne := &Pubprop{}
	pubpropTwo := &Pubprop{}
	if err = randomize.Struct(seed, pubpropOne, pubpropDBTypes, false, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}
	if err = randomize.Struct(seed, pubpropTwo, pubpropDBTypes, false, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = pubpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func pubpropBeforeInsertHook(e boil.Executor, o *Pubprop) error {
	*o = Pubprop{}
	return nil
}

func pubpropAfterInsertHook(e boil.Executor, o *Pubprop) error {
	*o = Pubprop{}
	return nil
}

func pubpropAfterSelectHook(e boil.Executor, o *Pubprop) error {
	*o = Pubprop{}
	return nil
}

func pubpropBeforeUpdateHook(e boil.Executor, o *Pubprop) error {
	*o = Pubprop{}
	return nil
}

func pubpropAfterUpdateHook(e boil.Executor, o *Pubprop) error {
	*o = Pubprop{}
	return nil
}

func pubpropBeforeDeleteHook(e boil.Executor, o *Pubprop) error {
	*o = Pubprop{}
	return nil
}

func pubpropAfterDeleteHook(e boil.Executor, o *Pubprop) error {
	*o = Pubprop{}
	return nil
}

func pubpropBeforeUpsertHook(e boil.Executor, o *Pubprop) error {
	*o = Pubprop{}
	return nil
}

func pubpropAfterUpsertHook(e boil.Executor, o *Pubprop) error {
	*o = Pubprop{}
	return nil
}

func testPubpropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Pubprop{}
	o := &Pubprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, pubpropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Pubprop object: %s", err)
	}

	AddPubpropHook(boil.BeforeInsertHook, pubpropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	pubpropBeforeInsertHooks = []PubpropHook{}

	AddPubpropHook(boil.AfterInsertHook, pubpropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	pubpropAfterInsertHooks = []PubpropHook{}

	AddPubpropHook(boil.AfterSelectHook, pubpropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	pubpropAfterSelectHooks = []PubpropHook{}

	AddPubpropHook(boil.BeforeUpdateHook, pubpropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	pubpropBeforeUpdateHooks = []PubpropHook{}

	AddPubpropHook(boil.AfterUpdateHook, pubpropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	pubpropAfterUpdateHooks = []PubpropHook{}

	AddPubpropHook(boil.BeforeDeleteHook, pubpropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	pubpropBeforeDeleteHooks = []PubpropHook{}

	AddPubpropHook(boil.AfterDeleteHook, pubpropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	pubpropAfterDeleteHooks = []PubpropHook{}

	AddPubpropHook(boil.BeforeUpsertHook, pubpropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	pubpropBeforeUpsertHooks = []PubpropHook{}

	AddPubpropHook(boil.AfterUpsertHook, pubpropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	pubpropAfterUpsertHooks = []PubpropHook{}
}
func testPubpropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubprop := &Pubprop{}
	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPubpropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubprop := &Pubprop{}
	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubprop.Insert(tx, pubpropColumns...); err != nil {
		t.Error(err)
	}

	count, err := Pubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPubpropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Pubprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, pubpropDBTypes, true, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
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

	slice := PubpropSlice{&local}
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

func testPubpropToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Pubprop
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, pubpropDBTypes, true, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
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

	slice := PubpropSlice{&local}
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

func testPubpropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pubprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubpropDBTypes, false, strmangle.SetComplement(pubpropPrimaryKeyColumns, pubpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypePubprop != &a {
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
func testPubpropToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pubprop
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubpropDBTypes, false, strmangle.SetComplement(pubpropPrimaryKeyColumns, pubpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.Pubprop != &a {
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
func testPubpropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubprop := &Pubprop{}
	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = pubprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPubpropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubprop := &Pubprop{}
	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PubpropSlice{pubprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testPubpropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubprop := &Pubprop{}
	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Pubprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	pubpropDBTypes = map[string]string{"PubID": "integer", "PubpropID": "integer", "Rank": "integer", "TypeID": "integer", "Value": "text"}
	_              = bytes.MinRead
)

func testPubpropsUpdate(t *testing.T) {
	t.Parallel()

	if len(pubpropColumns) == len(pubpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	pubprop := &Pubprop{}
	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	if err = pubprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testPubpropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(pubpropColumns) == len(pubpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	pubprop := &Pubprop{}
	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, pubprop, pubpropDBTypes, true, pubpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(pubpropColumns, pubpropPrimaryKeyColumns) {
		fields = pubpropColumns
	} else {
		fields = strmangle.SetComplement(
			pubpropColumns,
			pubpropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(pubprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := PubpropSlice{pubprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testPubpropsUpsert(t *testing.T) {
	t.Parallel()

	if len(pubpropColumns) == len(pubpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	pubprop := Pubprop{}
	if err = randomize.Struct(seed, &pubprop, pubpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Pubprop: %s", err)
	}

	count, err := Pubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &pubprop, pubpropDBTypes, false, pubpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}

	if err = pubprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Pubprop: %s", err)
	}

	count, err = Pubprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
