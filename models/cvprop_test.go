package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testCvprops(t *testing.T) {
	t.Parallel()

	query := Cvprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testCvpropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvprop := &Cvprop{}
	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cvprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCvpropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvprop := &Cvprop{}
	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Cvprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Cvprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCvpropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvprop := &Cvprop{}
	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CvpropSlice{cvprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testCvpropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvprop := &Cvprop{}
	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := CvpropExists(tx, cvprop.CvpropID)
	if err != nil {
		t.Errorf("Unable to check if Cvprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CvpropExistsG to return true, but got false.")
	}
}
func testCvpropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvprop := &Cvprop{}
	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvprop.Insert(tx); err != nil {
		t.Error(err)
	}

	cvpropFound, err := FindCvprop(tx, cvprop.CvpropID)
	if err != nil {
		t.Error(err)
	}

	if cvpropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testCvpropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvprop := &Cvprop{}
	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Cvprops(tx).Bind(cvprop); err != nil {
		t.Error(err)
	}
}

func testCvpropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvprop := &Cvprop{}
	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Cvprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCvpropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvpropOne := &Cvprop{}
	cvpropTwo := &Cvprop{}
	if err = randomize.Struct(seed, cvpropOne, cvpropDBTypes, false, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}
	if err = randomize.Struct(seed, cvpropTwo, cvpropDBTypes, false, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Cvprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCvpropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	cvpropOne := &Cvprop{}
	cvpropTwo := &Cvprop{}
	if err = randomize.Struct(seed, cvpropOne, cvpropDBTypes, false, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}
	if err = randomize.Struct(seed, cvpropTwo, cvpropDBTypes, false, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func cvpropBeforeInsertHook(e boil.Executor, o *Cvprop) error {
	*o = Cvprop{}
	return nil
}

func cvpropAfterInsertHook(e boil.Executor, o *Cvprop) error {
	*o = Cvprop{}
	return nil
}

func cvpropAfterSelectHook(e boil.Executor, o *Cvprop) error {
	*o = Cvprop{}
	return nil
}

func cvpropBeforeUpdateHook(e boil.Executor, o *Cvprop) error {
	*o = Cvprop{}
	return nil
}

func cvpropAfterUpdateHook(e boil.Executor, o *Cvprop) error {
	*o = Cvprop{}
	return nil
}

func cvpropBeforeDeleteHook(e boil.Executor, o *Cvprop) error {
	*o = Cvprop{}
	return nil
}

func cvpropAfterDeleteHook(e boil.Executor, o *Cvprop) error {
	*o = Cvprop{}
	return nil
}

func cvpropBeforeUpsertHook(e boil.Executor, o *Cvprop) error {
	*o = Cvprop{}
	return nil
}

func cvpropAfterUpsertHook(e boil.Executor, o *Cvprop) error {
	*o = Cvprop{}
	return nil
}

func testCvpropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Cvprop{}
	o := &Cvprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, cvpropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Cvprop object: %s", err)
	}

	AddCvpropHook(boil.BeforeInsertHook, cvpropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	cvpropBeforeInsertHooks = []CvpropHook{}

	AddCvpropHook(boil.AfterInsertHook, cvpropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	cvpropAfterInsertHooks = []CvpropHook{}

	AddCvpropHook(boil.AfterSelectHook, cvpropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	cvpropAfterSelectHooks = []CvpropHook{}

	AddCvpropHook(boil.BeforeUpdateHook, cvpropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	cvpropBeforeUpdateHooks = []CvpropHook{}

	AddCvpropHook(boil.AfterUpdateHook, cvpropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	cvpropAfterUpdateHooks = []CvpropHook{}

	AddCvpropHook(boil.BeforeDeleteHook, cvpropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	cvpropBeforeDeleteHooks = []CvpropHook{}

	AddCvpropHook(boil.AfterDeleteHook, cvpropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	cvpropAfterDeleteHooks = []CvpropHook{}

	AddCvpropHook(boil.BeforeUpsertHook, cvpropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	cvpropBeforeUpsertHooks = []CvpropHook{}

	AddCvpropHook(boil.AfterUpsertHook, cvpropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	cvpropAfterUpsertHooks = []CvpropHook{}
}
func testCvpropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvprop := &Cvprop{}
	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCvpropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvprop := &Cvprop{}
	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvprop.Insert(tx, cvpropColumns...); err != nil {
		t.Error(err)
	}

	count, err := Cvprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCvpropToOneCVUsingCV(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Cvprop
	var foreign CV

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvpropDBTypes, true, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cvDBTypes, true, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.CVID = foreign.CVID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.CV(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CVID != foreign.CVID {
		t.Errorf("want: %v, got %v", foreign.CVID, check.CVID)
	}

	slice := CvpropSlice{&local}
	if err = local.L.LoadCV(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.CV == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.CV = nil
	if err = local.L.LoadCV(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.CV == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvpropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Cvprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvpropDBTypes, true, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
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

	slice := CvpropSlice{&local}
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

func testCvpropToOneSetOpCVUsingCV(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvprop
	var b, c CV

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvpropDBTypes, false, strmangle.SetComplement(cvpropPrimaryKeyColumns, cvpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvDBTypes, false, strmangle.SetComplement(cvPrimaryKeyColumns, cvColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvDBTypes, false, strmangle.SetComplement(cvPrimaryKeyColumns, cvColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*CV{&b, &c} {
		err = a.SetCV(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.CV != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Cvprop != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CVID != x.CVID {
			t.Error("foreign key was wrong value", a.CVID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CVID))
		reflect.Indirect(reflect.ValueOf(&a.CVID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CVID != x.CVID {
			t.Error("foreign key was wrong value", a.CVID, x.CVID)
		}
	}
}
func testCvpropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvpropDBTypes, false, strmangle.SetComplement(cvpropPrimaryKeyColumns, cvpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeCvprop != &a {
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
func testCvpropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvprop := &Cvprop{}
	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cvprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCvpropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvprop := &Cvprop{}
	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CvpropSlice{cvprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testCvpropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvprop := &Cvprop{}
	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Cvprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	cvpropDBTypes = map[string]string{"CVID": "integer", "CvpropID": "integer", "Rank": "integer", "TypeID": "integer", "Value": "text"}
	_             = bytes.MinRead
)

func testCvpropsUpdate(t *testing.T) {
	t.Parallel()

	if len(cvpropColumns) == len(cvpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cvprop := &Cvprop{}
	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	if err = cvprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testCvpropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(cvpropColumns) == len(cvpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cvprop := &Cvprop{}
	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cvprop, cvpropDBTypes, true, cvpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(cvpropColumns, cvpropPrimaryKeyColumns) {
		fields = cvpropColumns
	} else {
		fields = strmangle.SetComplement(
			cvpropColumns,
			cvpropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(cvprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := CvpropSlice{cvprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testCvpropsUpsert(t *testing.T) {
	t.Parallel()

	if len(cvpropColumns) == len(cvpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	cvprop := Cvprop{}
	if err = randomize.Struct(seed, &cvprop, cvpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Cvprop: %s", err)
	}

	count, err := Cvprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &cvprop, cvpropDBTypes, false, cvpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}

	if err = cvprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Cvprop: %s", err)
	}

	count, err = Cvprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
