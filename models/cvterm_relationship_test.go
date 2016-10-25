package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testCvtermRelationships(t *testing.T) {
	t.Parallel()

	query := CvtermRelationships(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testCvtermRelationshipsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermRelationship := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cvtermRelationship.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := CvtermRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCvtermRelationshipsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermRelationship := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = CvtermRelationships(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := CvtermRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCvtermRelationshipsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermRelationship := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CvtermRelationshipSlice{cvtermRelationship}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := CvtermRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testCvtermRelationshipsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermRelationship := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := CvtermRelationshipExists(tx, cvtermRelationship.CvtermRelationshipID)
	if err != nil {
		t.Errorf("Unable to check if CvtermRelationship exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CvtermRelationshipExistsG to return true, but got false.")
	}
}
func testCvtermRelationshipsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermRelationship := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	cvtermRelationshipFound, err := FindCvtermRelationship(tx, cvtermRelationship.CvtermRelationshipID)
	if err != nil {
		t.Error(err)
	}

	if cvtermRelationshipFound == nil {
		t.Error("want a record, got nil")
	}
}
func testCvtermRelationshipsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermRelationship := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = CvtermRelationships(tx).Bind(cvtermRelationship); err != nil {
		t.Error(err)
	}
}

func testCvtermRelationshipsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermRelationship := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := CvtermRelationships(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCvtermRelationshipsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermRelationshipOne := &CvtermRelationship{}
	cvtermRelationshipTwo := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationshipOne, cvtermRelationshipDBTypes, false, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}
	if err = randomize.Struct(seed, cvtermRelationshipTwo, cvtermRelationshipDBTypes, false, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationshipOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvtermRelationshipTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := CvtermRelationships(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCvtermRelationshipsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	cvtermRelationshipOne := &CvtermRelationship{}
	cvtermRelationshipTwo := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationshipOne, cvtermRelationshipDBTypes, false, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}
	if err = randomize.Struct(seed, cvtermRelationshipTwo, cvtermRelationshipDBTypes, false, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationshipOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvtermRelationshipTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CvtermRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func cvtermRelationshipBeforeInsertHook(e boil.Executor, o *CvtermRelationship) error {
	*o = CvtermRelationship{}
	return nil
}

func cvtermRelationshipAfterInsertHook(e boil.Executor, o *CvtermRelationship) error {
	*o = CvtermRelationship{}
	return nil
}

func cvtermRelationshipAfterSelectHook(e boil.Executor, o *CvtermRelationship) error {
	*o = CvtermRelationship{}
	return nil
}

func cvtermRelationshipBeforeUpdateHook(e boil.Executor, o *CvtermRelationship) error {
	*o = CvtermRelationship{}
	return nil
}

func cvtermRelationshipAfterUpdateHook(e boil.Executor, o *CvtermRelationship) error {
	*o = CvtermRelationship{}
	return nil
}

func cvtermRelationshipBeforeDeleteHook(e boil.Executor, o *CvtermRelationship) error {
	*o = CvtermRelationship{}
	return nil
}

func cvtermRelationshipAfterDeleteHook(e boil.Executor, o *CvtermRelationship) error {
	*o = CvtermRelationship{}
	return nil
}

func cvtermRelationshipBeforeUpsertHook(e boil.Executor, o *CvtermRelationship) error {
	*o = CvtermRelationship{}
	return nil
}

func cvtermRelationshipAfterUpsertHook(e boil.Executor, o *CvtermRelationship) error {
	*o = CvtermRelationship{}
	return nil
}

func testCvtermRelationshipsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &CvtermRelationship{}
	o := &CvtermRelationship{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, cvtermRelationshipDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship object: %s", err)
	}

	AddCvtermRelationshipHook(boil.BeforeInsertHook, cvtermRelationshipBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	cvtermRelationshipBeforeInsertHooks = []CvtermRelationshipHook{}

	AddCvtermRelationshipHook(boil.AfterInsertHook, cvtermRelationshipAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	cvtermRelationshipAfterInsertHooks = []CvtermRelationshipHook{}

	AddCvtermRelationshipHook(boil.AfterSelectHook, cvtermRelationshipAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	cvtermRelationshipAfterSelectHooks = []CvtermRelationshipHook{}

	AddCvtermRelationshipHook(boil.BeforeUpdateHook, cvtermRelationshipBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	cvtermRelationshipBeforeUpdateHooks = []CvtermRelationshipHook{}

	AddCvtermRelationshipHook(boil.AfterUpdateHook, cvtermRelationshipAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	cvtermRelationshipAfterUpdateHooks = []CvtermRelationshipHook{}

	AddCvtermRelationshipHook(boil.BeforeDeleteHook, cvtermRelationshipBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	cvtermRelationshipBeforeDeleteHooks = []CvtermRelationshipHook{}

	AddCvtermRelationshipHook(boil.AfterDeleteHook, cvtermRelationshipAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	cvtermRelationshipAfterDeleteHooks = []CvtermRelationshipHook{}

	AddCvtermRelationshipHook(boil.BeforeUpsertHook, cvtermRelationshipBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	cvtermRelationshipBeforeUpsertHooks = []CvtermRelationshipHook{}

	AddCvtermRelationshipHook(boil.AfterUpsertHook, cvtermRelationshipAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	cvtermRelationshipAfterUpsertHooks = []CvtermRelationshipHook{}
}
func testCvtermRelationshipsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermRelationship := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CvtermRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCvtermRelationshipsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermRelationship := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationship.Insert(tx, cvtermRelationshipColumns...); err != nil {
		t.Error(err)
	}

	count, err := CvtermRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCvtermRelationshipToOneCvtermUsingObject(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local CvtermRelationship
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvtermRelationshipDBTypes, true, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ObjectID = foreign.CvtermID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Object(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := CvtermRelationshipSlice{&local}
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

func testCvtermRelationshipToOneCvtermUsingSubject(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local CvtermRelationship
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvtermRelationshipDBTypes, true, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.SubjectID = foreign.CvtermID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Subject(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := CvtermRelationshipSlice{&local}
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

func testCvtermRelationshipToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local CvtermRelationship
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvtermRelationshipDBTypes, true, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
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

	slice := CvtermRelationshipSlice{&local}
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

func testCvtermRelationshipToOneSetOpCvtermUsingObject(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a CvtermRelationship
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermRelationshipDBTypes, false, strmangle.SetComplement(cvtermRelationshipPrimaryKeyColumns, cvtermRelationshipColumnsWithoutDefault)...); err != nil {
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
		err = a.SetObject(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Object != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ObjectCvtermRelationship != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ObjectID != x.CvtermID {
			t.Error("foreign key was wrong value", a.ObjectID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ObjectID))
		reflect.Indirect(reflect.ValueOf(&a.ObjectID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ObjectID != x.CvtermID {
			t.Error("foreign key was wrong value", a.ObjectID, x.CvtermID)
		}
	}
}
func testCvtermRelationshipToOneSetOpCvtermUsingSubject(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a CvtermRelationship
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermRelationshipDBTypes, false, strmangle.SetComplement(cvtermRelationshipPrimaryKeyColumns, cvtermRelationshipColumnsWithoutDefault)...); err != nil {
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
		err = a.SetSubject(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Subject != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SubjectCvtermRelationship != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SubjectID != x.CvtermID {
			t.Error("foreign key was wrong value", a.SubjectID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SubjectID))
		reflect.Indirect(reflect.ValueOf(&a.SubjectID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SubjectID != x.CvtermID {
			t.Error("foreign key was wrong value", a.SubjectID, x.CvtermID)
		}
	}
}
func testCvtermRelationshipToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a CvtermRelationship
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermRelationshipDBTypes, false, strmangle.SetComplement(cvtermRelationshipPrimaryKeyColumns, cvtermRelationshipColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeCvtermRelationship != &a {
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
func testCvtermRelationshipsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermRelationship := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cvtermRelationship.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCvtermRelationshipsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermRelationship := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CvtermRelationshipSlice{cvtermRelationship}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testCvtermRelationshipsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermRelationship := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := CvtermRelationships(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	cvtermRelationshipDBTypes = map[string]string{"CvtermRelationshipID": "integer", "ObjectID": "integer", "SubjectID": "integer", "TypeID": "integer"}
	_                         = bytes.MinRead
)

func testCvtermRelationshipsUpdate(t *testing.T) {
	t.Parallel()

	if len(cvtermRelationshipColumns) == len(cvtermRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cvtermRelationship := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CvtermRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	if err = cvtermRelationship.Update(tx); err != nil {
		t.Error(err)
	}
}

func testCvtermRelationshipsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(cvtermRelationshipColumns) == len(cvtermRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cvtermRelationship := &CvtermRelationship{}
	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CvtermRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cvtermRelationship, cvtermRelationshipDBTypes, true, cvtermRelationshipPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(cvtermRelationshipColumns, cvtermRelationshipPrimaryKeyColumns) {
		fields = cvtermRelationshipColumns
	} else {
		fields = strmangle.SetComplement(
			cvtermRelationshipColumns,
			cvtermRelationshipPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(cvtermRelationship))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := CvtermRelationshipSlice{cvtermRelationship}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testCvtermRelationshipsUpsert(t *testing.T) {
	t.Parallel()

	if len(cvtermRelationshipColumns) == len(cvtermRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	cvtermRelationship := CvtermRelationship{}
	if err = randomize.Struct(seed, &cvtermRelationship, cvtermRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermRelationship.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert CvtermRelationship: %s", err)
	}

	count, err := CvtermRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &cvtermRelationship, cvtermRelationshipDBTypes, false, cvtermRelationshipPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}

	if err = cvtermRelationship.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert CvtermRelationship: %s", err)
	}

	count, err = CvtermRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
