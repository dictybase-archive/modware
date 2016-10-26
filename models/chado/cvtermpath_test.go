package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testCvtermpaths(t *testing.T) {
	t.Parallel()

	query := Cvtermpaths(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testCvtermpathsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermpath := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpath.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cvtermpath.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermpaths(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCvtermpathsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermpath := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpath.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Cvtermpaths(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Cvtermpaths(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCvtermpathsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermpath := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpath.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CvtermpathSlice{cvtermpath}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermpaths(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testCvtermpathsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermpath := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpath.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := CvtermpathExists(tx, cvtermpath.CvtermpathID)
	if err != nil {
		t.Errorf("Unable to check if Cvtermpath exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CvtermpathExistsG to return true, but got false.")
	}
}
func testCvtermpathsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermpath := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpath.Insert(tx); err != nil {
		t.Error(err)
	}

	cvtermpathFound, err := FindCvtermpath(tx, cvtermpath.CvtermpathID)
	if err != nil {
		t.Error(err)
	}

	if cvtermpathFound == nil {
		t.Error("want a record, got nil")
	}
}
func testCvtermpathsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermpath := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpath.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Cvtermpaths(tx).Bind(cvtermpath); err != nil {
		t.Error(err)
	}
}

func testCvtermpathsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermpath := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpath.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Cvtermpaths(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCvtermpathsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermpathOne := &Cvtermpath{}
	cvtermpathTwo := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpathOne, cvtermpathDBTypes, false, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}
	if err = randomize.Struct(seed, cvtermpathTwo, cvtermpathDBTypes, false, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpathOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvtermpathTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Cvtermpaths(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCvtermpathsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	cvtermpathOne := &Cvtermpath{}
	cvtermpathTwo := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpathOne, cvtermpathDBTypes, false, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}
	if err = randomize.Struct(seed, cvtermpathTwo, cvtermpathDBTypes, false, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpathOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvtermpathTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermpaths(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func cvtermpathBeforeInsertHook(e boil.Executor, o *Cvtermpath) error {
	*o = Cvtermpath{}
	return nil
}

func cvtermpathAfterInsertHook(e boil.Executor, o *Cvtermpath) error {
	*o = Cvtermpath{}
	return nil
}

func cvtermpathAfterSelectHook(e boil.Executor, o *Cvtermpath) error {
	*o = Cvtermpath{}
	return nil
}

func cvtermpathBeforeUpdateHook(e boil.Executor, o *Cvtermpath) error {
	*o = Cvtermpath{}
	return nil
}

func cvtermpathAfterUpdateHook(e boil.Executor, o *Cvtermpath) error {
	*o = Cvtermpath{}
	return nil
}

func cvtermpathBeforeDeleteHook(e boil.Executor, o *Cvtermpath) error {
	*o = Cvtermpath{}
	return nil
}

func cvtermpathAfterDeleteHook(e boil.Executor, o *Cvtermpath) error {
	*o = Cvtermpath{}
	return nil
}

func cvtermpathBeforeUpsertHook(e boil.Executor, o *Cvtermpath) error {
	*o = Cvtermpath{}
	return nil
}

func cvtermpathAfterUpsertHook(e boil.Executor, o *Cvtermpath) error {
	*o = Cvtermpath{}
	return nil
}

func testCvtermpathsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Cvtermpath{}
	o := &Cvtermpath{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, cvtermpathDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Cvtermpath object: %s", err)
	}

	AddCvtermpathHook(boil.BeforeInsertHook, cvtermpathBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	cvtermpathBeforeInsertHooks = []CvtermpathHook{}

	AddCvtermpathHook(boil.AfterInsertHook, cvtermpathAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	cvtermpathAfterInsertHooks = []CvtermpathHook{}

	AddCvtermpathHook(boil.AfterSelectHook, cvtermpathAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	cvtermpathAfterSelectHooks = []CvtermpathHook{}

	AddCvtermpathHook(boil.BeforeUpdateHook, cvtermpathBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	cvtermpathBeforeUpdateHooks = []CvtermpathHook{}

	AddCvtermpathHook(boil.AfterUpdateHook, cvtermpathAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	cvtermpathAfterUpdateHooks = []CvtermpathHook{}

	AddCvtermpathHook(boil.BeforeDeleteHook, cvtermpathBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	cvtermpathBeforeDeleteHooks = []CvtermpathHook{}

	AddCvtermpathHook(boil.AfterDeleteHook, cvtermpathAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	cvtermpathAfterDeleteHooks = []CvtermpathHook{}

	AddCvtermpathHook(boil.BeforeUpsertHook, cvtermpathBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	cvtermpathBeforeUpsertHooks = []CvtermpathHook{}

	AddCvtermpathHook(boil.AfterUpsertHook, cvtermpathAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	cvtermpathAfterUpsertHooks = []CvtermpathHook{}
}
func testCvtermpathsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermpath := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpath.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermpaths(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCvtermpathsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermpath := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpath.Insert(tx, cvtermpathColumns...); err != nil {
		t.Error(err)
	}

	count, err := Cvtermpaths(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCvtermpathToOneCVUsingCV(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Cvtermpath
	var foreign CV

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
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

	slice := CvtermpathSlice{&local}
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

func testCvtermpathToOneCvtermUsingObject(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Cvtermpath
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
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

	slice := CvtermpathSlice{&local}
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

func testCvtermpathToOneCvtermUsingSubject(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Cvtermpath
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
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

	slice := CvtermpathSlice{&local}
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

func testCvtermpathToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Cvtermpath
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
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

	slice := CvtermpathSlice{&local}
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

func testCvtermpathToOneSetOpCVUsingCV(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvtermpath
	var b, c CV

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermpathDBTypes, false, strmangle.SetComplement(cvtermpathPrimaryKeyColumns, cvtermpathColumnsWithoutDefault)...); err != nil {
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

		if x.R.Cvtermpaths[0] != &a {
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
func testCvtermpathToOneSetOpCvtermUsingObject(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvtermpath
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermpathDBTypes, false, strmangle.SetComplement(cvtermpathPrimaryKeyColumns, cvtermpathColumnsWithoutDefault)...); err != nil {
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

		if x.R.ObjectCvtermpath != &a {
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
func testCvtermpathToOneSetOpCvtermUsingSubject(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvtermpath
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermpathDBTypes, false, strmangle.SetComplement(cvtermpathPrimaryKeyColumns, cvtermpathColumnsWithoutDefault)...); err != nil {
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

		if x.R.SubjectCvtermpath != &a {
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
func testCvtermpathToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvtermpath
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermpathDBTypes, false, strmangle.SetComplement(cvtermpathPrimaryKeyColumns, cvtermpathColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeCvtermpath != &a {
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

func testCvtermpathToOneRemoveOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvtermpath
	var b Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermpathDBTypes, false, strmangle.SetComplement(cvtermpathPrimaryKeyColumns, cvtermpathColumnsWithoutDefault)...); err != nil {
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

	if b.R.TypeCvtermpath != nil {
		t.Error("failed to remove a from b's relationships")
	}

}

func testCvtermpathsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermpath := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpath.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cvtermpath.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCvtermpathsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermpath := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpath.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CvtermpathSlice{cvtermpath}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testCvtermpathsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermpath := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpath.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Cvtermpaths(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	cvtermpathDBTypes = map[string]string{"CVID": "integer", "CvtermpathID": "integer", "ObjectID": "integer", "Pathdistance": "integer", "SubjectID": "integer", "TypeID": "integer"}
	_                 = bytes.MinRead
)

func testCvtermpathsUpdate(t *testing.T) {
	t.Parallel()

	if len(cvtermpathColumns) == len(cvtermpathPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cvtermpath := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpath.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermpaths(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	if err = cvtermpath.Update(tx); err != nil {
		t.Error(err)
	}
}

func testCvtermpathsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(cvtermpathColumns) == len(cvtermpathPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cvtermpath := &Cvtermpath{}
	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpath.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvtermpaths(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cvtermpath, cvtermpathDBTypes, true, cvtermpathPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(cvtermpathColumns, cvtermpathPrimaryKeyColumns) {
		fields = cvtermpathColumns
	} else {
		fields = strmangle.SetComplement(
			cvtermpathColumns,
			cvtermpathPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(cvtermpath))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := CvtermpathSlice{cvtermpath}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testCvtermpathsUpsert(t *testing.T) {
	t.Parallel()

	if len(cvtermpathColumns) == len(cvtermpathPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	cvtermpath := Cvtermpath{}
	if err = randomize.Struct(seed, &cvtermpath, cvtermpathDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermpath.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Cvtermpath: %s", err)
	}

	count, err := Cvtermpaths(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &cvtermpath, cvtermpathDBTypes, false, cvtermpathPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}

	if err = cvtermpath.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Cvtermpath: %s", err)
	}

	count, err = Cvtermpaths(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
