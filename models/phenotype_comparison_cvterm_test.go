package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testPhenotypeComparisonCvterms(t *testing.T) {
	t.Parallel()

	query := PhenotypeComparisonCvterms(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testPhenotypeComparisonCvtermsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonCvterm := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = phenotypeComparisonCvterm.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisonCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhenotypeComparisonCvtermsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonCvterm := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = PhenotypeComparisonCvterms(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisonCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhenotypeComparisonCvtermsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonCvterm := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PhenotypeComparisonCvtermSlice{phenotypeComparisonCvterm}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisonCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testPhenotypeComparisonCvtermsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonCvterm := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := PhenotypeComparisonCvtermExists(tx, phenotypeComparisonCvterm.PhenotypeComparisonCvtermID)
	if err != nil {
		t.Errorf("Unable to check if PhenotypeComparisonCvterm exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PhenotypeComparisonCvtermExistsG to return true, but got false.")
	}
}
func testPhenotypeComparisonCvtermsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonCvterm := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	phenotypeComparisonCvtermFound, err := FindPhenotypeComparisonCvterm(tx, phenotypeComparisonCvterm.PhenotypeComparisonCvtermID)
	if err != nil {
		t.Error(err)
	}

	if phenotypeComparisonCvtermFound == nil {
		t.Error("want a record, got nil")
	}
}
func testPhenotypeComparisonCvtermsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonCvterm := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = PhenotypeComparisonCvterms(tx).Bind(phenotypeComparisonCvterm); err != nil {
		t.Error(err)
	}
}

func testPhenotypeComparisonCvtermsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonCvterm := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := PhenotypeComparisonCvterms(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPhenotypeComparisonCvtermsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonCvtermOne := &PhenotypeComparisonCvterm{}
	phenotypeComparisonCvtermTwo := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvtermOne, phenotypeComparisonCvtermDBTypes, false, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}
	if err = randomize.Struct(seed, phenotypeComparisonCvtermTwo, phenotypeComparisonCvtermDBTypes, false, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvtermOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = phenotypeComparisonCvtermTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := PhenotypeComparisonCvterms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPhenotypeComparisonCvtermsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	phenotypeComparisonCvtermOne := &PhenotypeComparisonCvterm{}
	phenotypeComparisonCvtermTwo := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvtermOne, phenotypeComparisonCvtermDBTypes, false, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}
	if err = randomize.Struct(seed, phenotypeComparisonCvtermTwo, phenotypeComparisonCvtermDBTypes, false, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvtermOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = phenotypeComparisonCvtermTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisonCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func phenotypeComparisonCvtermBeforeInsertHook(e boil.Executor, o *PhenotypeComparisonCvterm) error {
	*o = PhenotypeComparisonCvterm{}
	return nil
}

func phenotypeComparisonCvtermAfterInsertHook(e boil.Executor, o *PhenotypeComparisonCvterm) error {
	*o = PhenotypeComparisonCvterm{}
	return nil
}

func phenotypeComparisonCvtermAfterSelectHook(e boil.Executor, o *PhenotypeComparisonCvterm) error {
	*o = PhenotypeComparisonCvterm{}
	return nil
}

func phenotypeComparisonCvtermBeforeUpdateHook(e boil.Executor, o *PhenotypeComparisonCvterm) error {
	*o = PhenotypeComparisonCvterm{}
	return nil
}

func phenotypeComparisonCvtermAfterUpdateHook(e boil.Executor, o *PhenotypeComparisonCvterm) error {
	*o = PhenotypeComparisonCvterm{}
	return nil
}

func phenotypeComparisonCvtermBeforeDeleteHook(e boil.Executor, o *PhenotypeComparisonCvterm) error {
	*o = PhenotypeComparisonCvterm{}
	return nil
}

func phenotypeComparisonCvtermAfterDeleteHook(e boil.Executor, o *PhenotypeComparisonCvterm) error {
	*o = PhenotypeComparisonCvterm{}
	return nil
}

func phenotypeComparisonCvtermBeforeUpsertHook(e boil.Executor, o *PhenotypeComparisonCvterm) error {
	*o = PhenotypeComparisonCvterm{}
	return nil
}

func phenotypeComparisonCvtermAfterUpsertHook(e boil.Executor, o *PhenotypeComparisonCvterm) error {
	*o = PhenotypeComparisonCvterm{}
	return nil
}

func testPhenotypeComparisonCvtermsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &PhenotypeComparisonCvterm{}
	o := &PhenotypeComparisonCvterm{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, phenotypeComparisonCvtermDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm object: %s", err)
	}

	AddPhenotypeComparisonCvtermHook(boil.BeforeInsertHook, phenotypeComparisonCvtermBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonCvtermBeforeInsertHooks = []PhenotypeComparisonCvtermHook{}

	AddPhenotypeComparisonCvtermHook(boil.AfterInsertHook, phenotypeComparisonCvtermAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonCvtermAfterInsertHooks = []PhenotypeComparisonCvtermHook{}

	AddPhenotypeComparisonCvtermHook(boil.AfterSelectHook, phenotypeComparisonCvtermAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonCvtermAfterSelectHooks = []PhenotypeComparisonCvtermHook{}

	AddPhenotypeComparisonCvtermHook(boil.BeforeUpdateHook, phenotypeComparisonCvtermBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonCvtermBeforeUpdateHooks = []PhenotypeComparisonCvtermHook{}

	AddPhenotypeComparisonCvtermHook(boil.AfterUpdateHook, phenotypeComparisonCvtermAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonCvtermAfterUpdateHooks = []PhenotypeComparisonCvtermHook{}

	AddPhenotypeComparisonCvtermHook(boil.BeforeDeleteHook, phenotypeComparisonCvtermBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonCvtermBeforeDeleteHooks = []PhenotypeComparisonCvtermHook{}

	AddPhenotypeComparisonCvtermHook(boil.AfterDeleteHook, phenotypeComparisonCvtermAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonCvtermAfterDeleteHooks = []PhenotypeComparisonCvtermHook{}

	AddPhenotypeComparisonCvtermHook(boil.BeforeUpsertHook, phenotypeComparisonCvtermBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonCvtermBeforeUpsertHooks = []PhenotypeComparisonCvtermHook{}

	AddPhenotypeComparisonCvtermHook(boil.AfterUpsertHook, phenotypeComparisonCvtermAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonCvtermAfterUpsertHooks = []PhenotypeComparisonCvtermHook{}
}
func testPhenotypeComparisonCvtermsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonCvterm := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisonCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhenotypeComparisonCvtermsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonCvterm := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvterm.Insert(tx, phenotypeComparisonCvtermColumns...); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisonCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhenotypeComparisonCvtermToOnePhenotypeComparisonUsingPhenotypeComparison(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PhenotypeComparisonCvterm
	var foreign PhenotypeComparison

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeComparisonCvtermDBTypes, true, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.PhenotypeComparisonID = foreign.PhenotypeComparisonID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.PhenotypeComparison(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PhenotypeComparisonID != foreign.PhenotypeComparisonID {
		t.Errorf("want: %v, got %v", foreign.PhenotypeComparisonID, check.PhenotypeComparisonID)
	}

	slice := PhenotypeComparisonCvtermSlice{&local}
	if err = local.L.LoadPhenotypeComparison(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.PhenotypeComparison == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.PhenotypeComparison = nil
	if err = local.L.LoadPhenotypeComparison(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.PhenotypeComparison == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeComparisonCvtermToOneCvtermUsingCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PhenotypeComparisonCvterm
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeComparisonCvtermDBTypes, true, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
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

	slice := PhenotypeComparisonCvtermSlice{&local}
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

func testPhenotypeComparisonCvtermToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PhenotypeComparisonCvterm
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeComparisonCvtermDBTypes, true, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
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

	slice := PhenotypeComparisonCvtermSlice{&local}
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

func testPhenotypeComparisonCvtermToOneSetOpPhenotypeComparisonUsingPhenotypeComparison(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PhenotypeComparisonCvterm
	var b, c PhenotypeComparison

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeComparisonCvtermDBTypes, false, strmangle.SetComplement(phenotypeComparisonCvtermPrimaryKeyColumns, phenotypeComparisonCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*PhenotypeComparison{&b, &c} {
		err = a.SetPhenotypeComparison(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.PhenotypeComparison != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PhenotypeComparisonCvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PhenotypeComparisonID != x.PhenotypeComparisonID {
			t.Error("foreign key was wrong value", a.PhenotypeComparisonID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PhenotypeComparisonID))
		reflect.Indirect(reflect.ValueOf(&a.PhenotypeComparisonID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PhenotypeComparisonID != x.PhenotypeComparisonID {
			t.Error("foreign key was wrong value", a.PhenotypeComparisonID, x.PhenotypeComparisonID)
		}
	}
}
func testPhenotypeComparisonCvtermToOneSetOpCvtermUsingCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PhenotypeComparisonCvterm
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeComparisonCvtermDBTypes, false, strmangle.SetComplement(phenotypeComparisonCvtermPrimaryKeyColumns, phenotypeComparisonCvtermColumnsWithoutDefault)...); err != nil {
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

		if x.R.PhenotypeComparisonCvterm != &a {
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
func testPhenotypeComparisonCvtermToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PhenotypeComparisonCvterm
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeComparisonCvtermDBTypes, false, strmangle.SetComplement(phenotypeComparisonCvtermPrimaryKeyColumns, phenotypeComparisonCvtermColumnsWithoutDefault)...); err != nil {
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

		if x.R.PhenotypeComparisonCvterms[0] != &a {
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
func testPhenotypeComparisonCvtermsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonCvterm := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = phenotypeComparisonCvterm.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPhenotypeComparisonCvtermsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonCvterm := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PhenotypeComparisonCvtermSlice{phenotypeComparisonCvterm}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testPhenotypeComparisonCvtermsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonCvterm := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := PhenotypeComparisonCvterms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	phenotypeComparisonCvtermDBTypes = map[string]string{"CvtermID": "integer", "PhenotypeComparisonCvtermID": "integer", "PhenotypeComparisonID": "integer", "PubID": "integer", "Rank": "integer"}
	_                                = bytes.MinRead
)

func testPhenotypeComparisonCvtermsUpdate(t *testing.T) {
	t.Parallel()

	if len(phenotypeComparisonCvtermColumns) == len(phenotypeComparisonCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonCvterm := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisonCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	if err = phenotypeComparisonCvterm.Update(tx); err != nil {
		t.Error(err)
	}
}

func testPhenotypeComparisonCvtermsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(phenotypeComparisonCvtermColumns) == len(phenotypeComparisonCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonCvterm := &PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisonCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true, phenotypeComparisonCvtermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(phenotypeComparisonCvtermColumns, phenotypeComparisonCvtermPrimaryKeyColumns) {
		fields = phenotypeComparisonCvtermColumns
	} else {
		fields = strmangle.SetComplement(
			phenotypeComparisonCvtermColumns,
			phenotypeComparisonCvtermPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(phenotypeComparisonCvterm))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := PhenotypeComparisonCvtermSlice{phenotypeComparisonCvterm}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testPhenotypeComparisonCvtermsUpsert(t *testing.T) {
	t.Parallel()

	if len(phenotypeComparisonCvtermColumns) == len(phenotypeComparisonCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	phenotypeComparisonCvterm := PhenotypeComparisonCvterm{}
	if err = randomize.Struct(seed, &phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonCvterm.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert PhenotypeComparisonCvterm: %s", err)
	}

	count, err := PhenotypeComparisonCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &phenotypeComparisonCvterm, phenotypeComparisonCvtermDBTypes, false, phenotypeComparisonCvtermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}

	if err = phenotypeComparisonCvterm.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert PhenotypeComparisonCvterm: %s", err)
	}

	count, err = PhenotypeComparisonCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
