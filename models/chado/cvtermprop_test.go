package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testCvtermprops(t *testing.T) {
	t.Parallel()

	query := Cvtermprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testCvtermpropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermprop := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cvtermprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCvtermpropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermprop := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Cvtermprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Cvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCvtermpropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermprop := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CvtermpropSlice{cvtermprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testCvtermpropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermprop := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := CvtermpropExists(tx, cvtermprop.CvtermpropID)
	if err != nil {
		t.Errorf("Unable to check if Cvtermprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CvtermpropExistsG to return true, but got false.")
	}
}
func testCvtermpropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermprop := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	cvtermpropFound, err := FindCvtermprop(tx, cvtermprop.CvtermpropID)
	if err != nil {
		t.Error(err)
	}

	if cvtermpropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testCvtermpropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermprop := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Cvtermprops(tx).Bind(cvtermprop); err != nil {
		t.Error(err)
	}
}

func testCvtermpropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermprop := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Cvtermprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCvtermpropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermpropOne := &Cvtermprop{}
	cvtermpropTwo := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermpropOne, cvtermpropDBTypes, false, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}
	if err = randomize.Struct(seed, cvtermpropTwo, cvtermpropDBTypes, false, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvtermpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Cvtermprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCvtermpropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	cvtermpropOne := &Cvtermprop{}
	cvtermpropTwo := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermpropOne, cvtermpropDBTypes, false, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}
	if err = randomize.Struct(seed, cvtermpropTwo, cvtermpropDBTypes, false, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvtermpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func cvtermpropBeforeInsertHook(e boil.Executor, o *Cvtermprop) error {
	*o = Cvtermprop{}
	return nil
}

func cvtermpropAfterInsertHook(e boil.Executor, o *Cvtermprop) error {
	*o = Cvtermprop{}
	return nil
}

func cvtermpropAfterSelectHook(e boil.Executor, o *Cvtermprop) error {
	*o = Cvtermprop{}
	return nil
}

func cvtermpropBeforeUpdateHook(e boil.Executor, o *Cvtermprop) error {
	*o = Cvtermprop{}
	return nil
}

func cvtermpropAfterUpdateHook(e boil.Executor, o *Cvtermprop) error {
	*o = Cvtermprop{}
	return nil
}

func cvtermpropBeforeDeleteHook(e boil.Executor, o *Cvtermprop) error {
	*o = Cvtermprop{}
	return nil
}

func cvtermpropAfterDeleteHook(e boil.Executor, o *Cvtermprop) error {
	*o = Cvtermprop{}
	return nil
}

func cvtermpropBeforeUpsertHook(e boil.Executor, o *Cvtermprop) error {
	*o = Cvtermprop{}
	return nil
}

func cvtermpropAfterUpsertHook(e boil.Executor, o *Cvtermprop) error {
	*o = Cvtermprop{}
	return nil
}

func testCvtermpropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Cvtermprop{}
	o := &Cvtermprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, cvtermpropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Cvtermprop object: %s", err)
	}

	AddCvtermpropHook(boil.BeforeInsertHook, cvtermpropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	cvtermpropBeforeInsertHooks = []CvtermpropHook{}

	AddCvtermpropHook(boil.AfterInsertHook, cvtermpropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	cvtermpropAfterInsertHooks = []CvtermpropHook{}

	AddCvtermpropHook(boil.AfterSelectHook, cvtermpropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	cvtermpropAfterSelectHooks = []CvtermpropHook{}

	AddCvtermpropHook(boil.BeforeUpdateHook, cvtermpropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	cvtermpropBeforeUpdateHooks = []CvtermpropHook{}

	AddCvtermpropHook(boil.AfterUpdateHook, cvtermpropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	cvtermpropAfterUpdateHooks = []CvtermpropHook{}

	AddCvtermpropHook(boil.BeforeDeleteHook, cvtermpropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	cvtermpropBeforeDeleteHooks = []CvtermpropHook{}

	AddCvtermpropHook(boil.AfterDeleteHook, cvtermpropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	cvtermpropAfterDeleteHooks = []CvtermpropHook{}

	AddCvtermpropHook(boil.BeforeUpsertHook, cvtermpropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	cvtermpropBeforeUpsertHooks = []CvtermpropHook{}

	AddCvtermpropHook(boil.AfterUpsertHook, cvtermpropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	cvtermpropAfterUpsertHooks = []CvtermpropHook{}
}
func testCvtermpropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermprop := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCvtermpropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermprop := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermprop.Insert(tx, cvtermpropColumns...); err != nil {
		t.Error(err)
	}

	count, err := Cvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCvtermpropToOneCvtermUsingCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Cvtermprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvtermpropDBTypes, true, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
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

	slice := CvtermpropSlice{&local}
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

func testCvtermpropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Cvtermprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvtermpropDBTypes, true, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
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

	slice := CvtermpropSlice{&local}
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

func testCvtermpropToOneSetOpCvtermUsingCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvtermprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermpropDBTypes, false, strmangle.SetComplement(cvtermpropPrimaryKeyColumns, cvtermpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.Cvtermprop != &a {
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
func testCvtermpropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvtermprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermpropDBTypes, false, strmangle.SetComplement(cvtermpropPrimaryKeyColumns, cvtermpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeCvtermprop != &a {
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
func testCvtermpropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermprop := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cvtermprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCvtermpropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermprop := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CvtermpropSlice{cvtermprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testCvtermpropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermprop := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Cvtermprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	cvtermpropDBTypes = map[string]string{"CvtermID": "integer", "CvtermpropID": "integer", "Rank": "integer", "TypeID": "integer", "Value": "text"}
	_                 = bytes.MinRead
)

func testCvtermpropsUpdate(t *testing.T) {
	t.Parallel()

	if len(cvtermpropColumns) == len(cvtermpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cvtermprop := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	if err = cvtermprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testCvtermpropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(cvtermpropColumns) == len(cvtermpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cvtermprop := &Cvtermprop{}
	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cvtermprop, cvtermpropDBTypes, true, cvtermpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(cvtermpropColumns, cvtermpropPrimaryKeyColumns) {
		fields = cvtermpropColumns
	} else {
		fields = strmangle.SetComplement(
			cvtermpropColumns,
			cvtermpropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(cvtermprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := CvtermpropSlice{cvtermprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testCvtermpropsUpsert(t *testing.T) {
	t.Parallel()

	if len(cvtermpropColumns) == len(cvtermpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	cvtermprop := Cvtermprop{}
	if err = randomize.Struct(seed, &cvtermprop, cvtermpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Cvtermprop: %s", err)
	}

	count, err := Cvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &cvtermprop, cvtermpropDBTypes, false, cvtermpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}

	if err = cvtermprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Cvtermprop: %s", err)
	}

	count, err = Cvtermprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
