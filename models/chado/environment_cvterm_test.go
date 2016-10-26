package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testEnvironmentCvterms(t *testing.T) {
	t.Parallel()

	query := EnvironmentCvterms(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testEnvironmentCvtermsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environmentCvterm := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = environmentCvterm.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := EnvironmentCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEnvironmentCvtermsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environmentCvterm := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = EnvironmentCvterms(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := EnvironmentCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEnvironmentCvtermsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environmentCvterm := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := EnvironmentCvtermSlice{environmentCvterm}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := EnvironmentCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testEnvironmentCvtermsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environmentCvterm := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := EnvironmentCvtermExists(tx, environmentCvterm.EnvironmentCvtermID)
	if err != nil {
		t.Errorf("Unable to check if EnvironmentCvterm exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EnvironmentCvtermExistsG to return true, but got false.")
	}
}
func testEnvironmentCvtermsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environmentCvterm := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	environmentCvtermFound, err := FindEnvironmentCvterm(tx, environmentCvterm.EnvironmentCvtermID)
	if err != nil {
		t.Error(err)
	}

	if environmentCvtermFound == nil {
		t.Error("want a record, got nil")
	}
}
func testEnvironmentCvtermsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environmentCvterm := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = EnvironmentCvterms(tx).Bind(environmentCvterm); err != nil {
		t.Error(err)
	}
}

func testEnvironmentCvtermsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environmentCvterm := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := EnvironmentCvterms(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEnvironmentCvtermsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environmentCvtermOne := &EnvironmentCvterm{}
	environmentCvtermTwo := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvtermOne, environmentCvtermDBTypes, false, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}
	if err = randomize.Struct(seed, environmentCvtermTwo, environmentCvtermDBTypes, false, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvtermOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = environmentCvtermTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := EnvironmentCvterms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEnvironmentCvtermsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	environmentCvtermOne := &EnvironmentCvterm{}
	environmentCvtermTwo := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvtermOne, environmentCvtermDBTypes, false, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}
	if err = randomize.Struct(seed, environmentCvtermTwo, environmentCvtermDBTypes, false, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvtermOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = environmentCvtermTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := EnvironmentCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func environmentCvtermBeforeInsertHook(e boil.Executor, o *EnvironmentCvterm) error {
	*o = EnvironmentCvterm{}
	return nil
}

func environmentCvtermAfterInsertHook(e boil.Executor, o *EnvironmentCvterm) error {
	*o = EnvironmentCvterm{}
	return nil
}

func environmentCvtermAfterSelectHook(e boil.Executor, o *EnvironmentCvterm) error {
	*o = EnvironmentCvterm{}
	return nil
}

func environmentCvtermBeforeUpdateHook(e boil.Executor, o *EnvironmentCvterm) error {
	*o = EnvironmentCvterm{}
	return nil
}

func environmentCvtermAfterUpdateHook(e boil.Executor, o *EnvironmentCvterm) error {
	*o = EnvironmentCvterm{}
	return nil
}

func environmentCvtermBeforeDeleteHook(e boil.Executor, o *EnvironmentCvterm) error {
	*o = EnvironmentCvterm{}
	return nil
}

func environmentCvtermAfterDeleteHook(e boil.Executor, o *EnvironmentCvterm) error {
	*o = EnvironmentCvterm{}
	return nil
}

func environmentCvtermBeforeUpsertHook(e boil.Executor, o *EnvironmentCvterm) error {
	*o = EnvironmentCvterm{}
	return nil
}

func environmentCvtermAfterUpsertHook(e boil.Executor, o *EnvironmentCvterm) error {
	*o = EnvironmentCvterm{}
	return nil
}

func testEnvironmentCvtermsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &EnvironmentCvterm{}
	o := &EnvironmentCvterm{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, environmentCvtermDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm object: %s", err)
	}

	AddEnvironmentCvtermHook(boil.BeforeInsertHook, environmentCvtermBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	environmentCvtermBeforeInsertHooks = []EnvironmentCvtermHook{}

	AddEnvironmentCvtermHook(boil.AfterInsertHook, environmentCvtermAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	environmentCvtermAfterInsertHooks = []EnvironmentCvtermHook{}

	AddEnvironmentCvtermHook(boil.AfterSelectHook, environmentCvtermAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	environmentCvtermAfterSelectHooks = []EnvironmentCvtermHook{}

	AddEnvironmentCvtermHook(boil.BeforeUpdateHook, environmentCvtermBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	environmentCvtermBeforeUpdateHooks = []EnvironmentCvtermHook{}

	AddEnvironmentCvtermHook(boil.AfterUpdateHook, environmentCvtermAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	environmentCvtermAfterUpdateHooks = []EnvironmentCvtermHook{}

	AddEnvironmentCvtermHook(boil.BeforeDeleteHook, environmentCvtermBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	environmentCvtermBeforeDeleteHooks = []EnvironmentCvtermHook{}

	AddEnvironmentCvtermHook(boil.AfterDeleteHook, environmentCvtermAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	environmentCvtermAfterDeleteHooks = []EnvironmentCvtermHook{}

	AddEnvironmentCvtermHook(boil.BeforeUpsertHook, environmentCvtermBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	environmentCvtermBeforeUpsertHooks = []EnvironmentCvtermHook{}

	AddEnvironmentCvtermHook(boil.AfterUpsertHook, environmentCvtermAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	environmentCvtermAfterUpsertHooks = []EnvironmentCvtermHook{}
}
func testEnvironmentCvtermsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environmentCvterm := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := EnvironmentCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEnvironmentCvtermsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environmentCvterm := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvterm.Insert(tx, environmentCvtermColumns...); err != nil {
		t.Error(err)
	}

	count, err := EnvironmentCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEnvironmentCvtermToOneEnvironmentUsingEnvironment(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local EnvironmentCvterm
	var foreign Environment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, environmentCvtermDBTypes, true, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.EnvironmentID = foreign.EnvironmentID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Environment(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.EnvironmentID != foreign.EnvironmentID {
		t.Errorf("want: %v, got %v", foreign.EnvironmentID, check.EnvironmentID)
	}

	slice := EnvironmentCvtermSlice{&local}
	if err = local.L.LoadEnvironment(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Environment == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Environment = nil
	if err = local.L.LoadEnvironment(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Environment == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testEnvironmentCvtermToOneCvtermUsingCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local EnvironmentCvterm
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, environmentCvtermDBTypes, true, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
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

	slice := EnvironmentCvtermSlice{&local}
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

func testEnvironmentCvtermToOneSetOpEnvironmentUsingEnvironment(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a EnvironmentCvterm
	var b, c Environment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, environmentCvtermDBTypes, false, strmangle.SetComplement(environmentCvtermPrimaryKeyColumns, environmentCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Environment{&b, &c} {
		err = a.SetEnvironment(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Environment != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.EnvironmentCvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.EnvironmentID != x.EnvironmentID {
			t.Error("foreign key was wrong value", a.EnvironmentID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.EnvironmentID))
		reflect.Indirect(reflect.ValueOf(&a.EnvironmentID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.EnvironmentID != x.EnvironmentID {
			t.Error("foreign key was wrong value", a.EnvironmentID, x.EnvironmentID)
		}
	}
}
func testEnvironmentCvtermToOneSetOpCvtermUsingCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a EnvironmentCvterm
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, environmentCvtermDBTypes, false, strmangle.SetComplement(environmentCvtermPrimaryKeyColumns, environmentCvtermColumnsWithoutDefault)...); err != nil {
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

		if x.R.EnvironmentCvterm != &a {
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
func testEnvironmentCvtermsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environmentCvterm := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = environmentCvterm.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testEnvironmentCvtermsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environmentCvterm := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := EnvironmentCvtermSlice{environmentCvterm}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testEnvironmentCvtermsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environmentCvterm := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := EnvironmentCvterms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	environmentCvtermDBTypes = map[string]string{"CvtermID": "integer", "EnvironmentCvtermID": "integer", "EnvironmentID": "integer"}
	_                        = bytes.MinRead
)

func testEnvironmentCvtermsUpdate(t *testing.T) {
	t.Parallel()

	if len(environmentCvtermColumns) == len(environmentCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	environmentCvterm := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := EnvironmentCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	if err = environmentCvterm.Update(tx); err != nil {
		t.Error(err)
	}
}

func testEnvironmentCvtermsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(environmentCvtermColumns) == len(environmentCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	environmentCvterm := &EnvironmentCvterm{}
	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := EnvironmentCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, environmentCvterm, environmentCvtermDBTypes, true, environmentCvtermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(environmentCvtermColumns, environmentCvtermPrimaryKeyColumns) {
		fields = environmentCvtermColumns
	} else {
		fields = strmangle.SetComplement(
			environmentCvtermColumns,
			environmentCvtermPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(environmentCvterm))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := EnvironmentCvtermSlice{environmentCvterm}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testEnvironmentCvtermsUpsert(t *testing.T) {
	t.Parallel()

	if len(environmentCvtermColumns) == len(environmentCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	environmentCvterm := EnvironmentCvterm{}
	if err = randomize.Struct(seed, &environmentCvterm, environmentCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentCvterm.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert EnvironmentCvterm: %s", err)
	}

	count, err := EnvironmentCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &environmentCvterm, environmentCvtermDBTypes, false, environmentCvtermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}

	if err = environmentCvterm.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert EnvironmentCvterm: %s", err)
	}

	count, err = EnvironmentCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
