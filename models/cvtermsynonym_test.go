package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testCvtermsynonyms(t *testing.T) {
	t.Parallel()

	query := Cvtermsynonyms(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testCvtermsynonymsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermsynonym := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cvtermsynonym.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermsynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCvtermsynonymsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermsynonym := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Cvtermsynonyms(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Cvtermsynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCvtermsynonymsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermsynonym := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CvtermsynonymSlice{cvtermsynonym}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermsynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testCvtermsynonymsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermsynonym := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := CvtermsynonymExists(tx, cvtermsynonym.CvtermsynonymID)
	if err != nil {
		t.Errorf("Unable to check if Cvtermsynonym exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CvtermsynonymExistsG to return true, but got false.")
	}
}
func testCvtermsynonymsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermsynonym := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	cvtermsynonymFound, err := FindCvtermsynonym(tx, cvtermsynonym.CvtermsynonymID)
	if err != nil {
		t.Error(err)
	}

	if cvtermsynonymFound == nil {
		t.Error("want a record, got nil")
	}
}
func testCvtermsynonymsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermsynonym := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Cvtermsynonyms(tx).Bind(cvtermsynonym); err != nil {
		t.Error(err)
	}
}

func testCvtermsynonymsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermsynonym := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Cvtermsynonyms(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCvtermsynonymsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermsynonymOne := &Cvtermsynonym{}
	cvtermsynonymTwo := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonymOne, cvtermsynonymDBTypes, false, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}
	if err = randomize.Struct(seed, cvtermsynonymTwo, cvtermsynonymDBTypes, false, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonymOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvtermsynonymTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Cvtermsynonyms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCvtermsynonymsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	cvtermsynonymOne := &Cvtermsynonym{}
	cvtermsynonymTwo := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonymOne, cvtermsynonymDBTypes, false, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}
	if err = randomize.Struct(seed, cvtermsynonymTwo, cvtermsynonymDBTypes, false, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonymOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvtermsynonymTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermsynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func cvtermsynonymBeforeInsertHook(e boil.Executor, o *Cvtermsynonym) error {
	*o = Cvtermsynonym{}
	return nil
}

func cvtermsynonymAfterInsertHook(e boil.Executor, o *Cvtermsynonym) error {
	*o = Cvtermsynonym{}
	return nil
}

func cvtermsynonymAfterSelectHook(e boil.Executor, o *Cvtermsynonym) error {
	*o = Cvtermsynonym{}
	return nil
}

func cvtermsynonymBeforeUpdateHook(e boil.Executor, o *Cvtermsynonym) error {
	*o = Cvtermsynonym{}
	return nil
}

func cvtermsynonymAfterUpdateHook(e boil.Executor, o *Cvtermsynonym) error {
	*o = Cvtermsynonym{}
	return nil
}

func cvtermsynonymBeforeDeleteHook(e boil.Executor, o *Cvtermsynonym) error {
	*o = Cvtermsynonym{}
	return nil
}

func cvtermsynonymAfterDeleteHook(e boil.Executor, o *Cvtermsynonym) error {
	*o = Cvtermsynonym{}
	return nil
}

func cvtermsynonymBeforeUpsertHook(e boil.Executor, o *Cvtermsynonym) error {
	*o = Cvtermsynonym{}
	return nil
}

func cvtermsynonymAfterUpsertHook(e boil.Executor, o *Cvtermsynonym) error {
	*o = Cvtermsynonym{}
	return nil
}

func testCvtermsynonymsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Cvtermsynonym{}
	o := &Cvtermsynonym{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, cvtermsynonymDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym object: %s", err)
	}

	AddCvtermsynonymHook(boil.BeforeInsertHook, cvtermsynonymBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	cvtermsynonymBeforeInsertHooks = []CvtermsynonymHook{}

	AddCvtermsynonymHook(boil.AfterInsertHook, cvtermsynonymAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	cvtermsynonymAfterInsertHooks = []CvtermsynonymHook{}

	AddCvtermsynonymHook(boil.AfterSelectHook, cvtermsynonymAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	cvtermsynonymAfterSelectHooks = []CvtermsynonymHook{}

	AddCvtermsynonymHook(boil.BeforeUpdateHook, cvtermsynonymBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	cvtermsynonymBeforeUpdateHooks = []CvtermsynonymHook{}

	AddCvtermsynonymHook(boil.AfterUpdateHook, cvtermsynonymAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	cvtermsynonymAfterUpdateHooks = []CvtermsynonymHook{}

	AddCvtermsynonymHook(boil.BeforeDeleteHook, cvtermsynonymBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	cvtermsynonymBeforeDeleteHooks = []CvtermsynonymHook{}

	AddCvtermsynonymHook(boil.AfterDeleteHook, cvtermsynonymAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	cvtermsynonymAfterDeleteHooks = []CvtermsynonymHook{}

	AddCvtermsynonymHook(boil.BeforeUpsertHook, cvtermsynonymBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	cvtermsynonymBeforeUpsertHooks = []CvtermsynonymHook{}

	AddCvtermsynonymHook(boil.AfterUpsertHook, cvtermsynonymAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	cvtermsynonymAfterUpsertHooks = []CvtermsynonymHook{}
}
func testCvtermsynonymsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermsynonym := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermsynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCvtermsynonymsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermsynonym := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonym.Insert(tx, cvtermsynonymColumns...); err != nil {
		t.Error(err)
	}

	count, err := Cvtermsynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCvtermsynonymToOneCvtermUsingCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Cvtermsynonym
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvtermsynonymDBTypes, true, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
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

	slice := CvtermsynonymSlice{&local}
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

func testCvtermsynonymToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Cvtermsynonym
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvtermsynonymDBTypes, true, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	local.TypeID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.TypeID.Int = foreign.CvtermID
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

	slice := CvtermsynonymSlice{&local}
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

func testCvtermsynonymToOneSetOpCvtermUsingCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvtermsynonym
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermsynonymDBTypes, false, strmangle.SetComplement(cvtermsynonymPrimaryKeyColumns, cvtermsynonymColumnsWithoutDefault)...); err != nil {
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

		if x.R.Cvtermsynonym != &a {
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
func testCvtermsynonymToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvtermsynonym
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermsynonymDBTypes, false, strmangle.SetComplement(cvtermsynonymPrimaryKeyColumns, cvtermsynonymColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeCvtermsynonyms[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TypeID.Int != x.CvtermID {
			t.Error("foreign key was wrong value", a.TypeID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TypeID.Int))
		reflect.Indirect(reflect.ValueOf(&a.TypeID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TypeID.Int != x.CvtermID {
			t.Error("foreign key was wrong value", a.TypeID.Int, x.CvtermID)
		}
	}
}

func testCvtermsynonymToOneRemoveOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvtermsynonym
	var b Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermsynonymDBTypes, false, strmangle.SetComplement(cvtermsynonymPrimaryKeyColumns, cvtermsynonymColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetType(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveType(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Type(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Type != nil {
		t.Error("R struct entry should be nil")
	}

	if a.TypeID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.TypeCvtermsynonyms) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testCvtermsynonymsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermsynonym := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cvtermsynonym.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCvtermsynonymsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermsynonym := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CvtermsynonymSlice{cvtermsynonym}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testCvtermsynonymsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermsynonym := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Cvtermsynonyms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	cvtermsynonymDBTypes = map[string]string{"CvtermID": "integer", "CvtermsynonymID": "integer", "Synonym": "character varying", "TypeID": "integer"}
	_                    = bytes.MinRead
)

func testCvtermsynonymsUpdate(t *testing.T) {
	t.Parallel()

	if len(cvtermsynonymColumns) == len(cvtermsynonymPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cvtermsynonym := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermsynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	if err = cvtermsynonym.Update(tx); err != nil {
		t.Error(err)
	}
}

func testCvtermsynonymsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(cvtermsynonymColumns) == len(cvtermsynonymPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cvtermsynonym := &Cvtermsynonym{}
	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermsynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cvtermsynonym, cvtermsynonymDBTypes, true, cvtermsynonymPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(cvtermsynonymColumns, cvtermsynonymPrimaryKeyColumns) {
		fields = cvtermsynonymColumns
	} else {
		fields = strmangle.SetComplement(
			cvtermsynonymColumns,
			cvtermsynonymPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(cvtermsynonym))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := CvtermsynonymSlice{cvtermsynonym}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testCvtermsynonymsUpsert(t *testing.T) {
	t.Parallel()

	if len(cvtermsynonymColumns) == len(cvtermsynonymPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	cvtermsynonym := Cvtermsynonym{}
	if err = randomize.Struct(seed, &cvtermsynonym, cvtermsynonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermsynonym.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Cvtermsynonym: %s", err)
	}

	count, err := Cvtermsynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &cvtermsynonym, cvtermsynonymDBTypes, false, cvtermsynonymPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}

	if err = cvtermsynonym.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Cvtermsynonym: %s", err)
	}

	count, err = Cvtermsynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
