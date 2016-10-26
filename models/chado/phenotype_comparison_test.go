package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testPhenotypeComparisons(t *testing.T) {
	t.Parallel()

	query := PhenotypeComparisons(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testPhenotypeComparisonsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparison := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparison.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = phenotypeComparison.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisons(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhenotypeComparisonsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparison := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparison.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = PhenotypeComparisons(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisons(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhenotypeComparisonsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparison := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparison.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PhenotypeComparisonSlice{phenotypeComparison}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisons(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testPhenotypeComparisonsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparison := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparison.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := PhenotypeComparisonExists(tx, phenotypeComparison.PhenotypeComparisonID)
	if err != nil {
		t.Errorf("Unable to check if PhenotypeComparison exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PhenotypeComparisonExistsG to return true, but got false.")
	}
}
func testPhenotypeComparisonsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparison := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparison.Insert(tx); err != nil {
		t.Error(err)
	}

	phenotypeComparisonFound, err := FindPhenotypeComparison(tx, phenotypeComparison.PhenotypeComparisonID)
	if err != nil {
		t.Error(err)
	}

	if phenotypeComparisonFound == nil {
		t.Error("want a record, got nil")
	}
}
func testPhenotypeComparisonsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparison := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparison.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = PhenotypeComparisons(tx).Bind(phenotypeComparison); err != nil {
		t.Error(err)
	}
}

func testPhenotypeComparisonsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparison := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparison.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := PhenotypeComparisons(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPhenotypeComparisonsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparisonOne := &PhenotypeComparison{}
	phenotypeComparisonTwo := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparisonOne, phenotypeComparisonDBTypes, false, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}
	if err = randomize.Struct(seed, phenotypeComparisonTwo, phenotypeComparisonDBTypes, false, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = phenotypeComparisonTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := PhenotypeComparisons(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPhenotypeComparisonsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	phenotypeComparisonOne := &PhenotypeComparison{}
	phenotypeComparisonTwo := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparisonOne, phenotypeComparisonDBTypes, false, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}
	if err = randomize.Struct(seed, phenotypeComparisonTwo, phenotypeComparisonDBTypes, false, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparisonOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = phenotypeComparisonTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisons(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func phenotypeComparisonBeforeInsertHook(e boil.Executor, o *PhenotypeComparison) error {
	*o = PhenotypeComparison{}
	return nil
}

func phenotypeComparisonAfterInsertHook(e boil.Executor, o *PhenotypeComparison) error {
	*o = PhenotypeComparison{}
	return nil
}

func phenotypeComparisonAfterSelectHook(e boil.Executor, o *PhenotypeComparison) error {
	*o = PhenotypeComparison{}
	return nil
}

func phenotypeComparisonBeforeUpdateHook(e boil.Executor, o *PhenotypeComparison) error {
	*o = PhenotypeComparison{}
	return nil
}

func phenotypeComparisonAfterUpdateHook(e boil.Executor, o *PhenotypeComparison) error {
	*o = PhenotypeComparison{}
	return nil
}

func phenotypeComparisonBeforeDeleteHook(e boil.Executor, o *PhenotypeComparison) error {
	*o = PhenotypeComparison{}
	return nil
}

func phenotypeComparisonAfterDeleteHook(e boil.Executor, o *PhenotypeComparison) error {
	*o = PhenotypeComparison{}
	return nil
}

func phenotypeComparisonBeforeUpsertHook(e boil.Executor, o *PhenotypeComparison) error {
	*o = PhenotypeComparison{}
	return nil
}

func phenotypeComparisonAfterUpsertHook(e boil.Executor, o *PhenotypeComparison) error {
	*o = PhenotypeComparison{}
	return nil
}

func testPhenotypeComparisonsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &PhenotypeComparison{}
	o := &PhenotypeComparison{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, phenotypeComparisonDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison object: %s", err)
	}

	AddPhenotypeComparisonHook(boil.BeforeInsertHook, phenotypeComparisonBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonBeforeInsertHooks = []PhenotypeComparisonHook{}

	AddPhenotypeComparisonHook(boil.AfterInsertHook, phenotypeComparisonAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonAfterInsertHooks = []PhenotypeComparisonHook{}

	AddPhenotypeComparisonHook(boil.AfterSelectHook, phenotypeComparisonAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonAfterSelectHooks = []PhenotypeComparisonHook{}

	AddPhenotypeComparisonHook(boil.BeforeUpdateHook, phenotypeComparisonBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonBeforeUpdateHooks = []PhenotypeComparisonHook{}

	AddPhenotypeComparisonHook(boil.AfterUpdateHook, phenotypeComparisonAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonAfterUpdateHooks = []PhenotypeComparisonHook{}

	AddPhenotypeComparisonHook(boil.BeforeDeleteHook, phenotypeComparisonBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonBeforeDeleteHooks = []PhenotypeComparisonHook{}

	AddPhenotypeComparisonHook(boil.AfterDeleteHook, phenotypeComparisonAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonAfterDeleteHooks = []PhenotypeComparisonHook{}

	AddPhenotypeComparisonHook(boil.BeforeUpsertHook, phenotypeComparisonBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonBeforeUpsertHooks = []PhenotypeComparisonHook{}

	AddPhenotypeComparisonHook(boil.AfterUpsertHook, phenotypeComparisonAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	phenotypeComparisonAfterUpsertHooks = []PhenotypeComparisonHook{}
}
func testPhenotypeComparisonsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparison := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparison.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisons(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhenotypeComparisonsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparison := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparison.Insert(tx, phenotypeComparisonColumns...); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisons(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhenotypeComparisonOneToOnePhenotypeComparisonCvtermUsingPhenotypeComparisonCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign PhenotypeComparisonCvterm
	var local PhenotypeComparison

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenotypeComparisonCvtermDBTypes, true, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PhenotypeComparisonID = local.PhenotypeComparisonID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.PhenotypeComparisonCvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PhenotypeComparisonID != foreign.PhenotypeComparisonID {
		t.Errorf("want: %v, got %v", foreign.PhenotypeComparisonID, check.PhenotypeComparisonID)
	}

	slice := PhenotypeComparisonSlice{&local}
	if err = local.L.LoadPhenotypeComparisonCvterm(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.PhenotypeComparisonCvterm == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.PhenotypeComparisonCvterm = nil
	if err = local.L.LoadPhenotypeComparisonCvterm(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.PhenotypeComparisonCvterm == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeComparisonOneToOneSetOpPhenotypeComparisonCvtermUsingPhenotypeComparisonCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PhenotypeComparison
	var b, c PhenotypeComparisonCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, phenotypeComparisonCvtermDBTypes, false, strmangle.SetComplement(phenotypeComparisonCvtermPrimaryKeyColumns, phenotypeComparisonCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, phenotypeComparisonCvtermDBTypes, false, strmangle.SetComplement(phenotypeComparisonCvtermPrimaryKeyColumns, phenotypeComparisonCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*PhenotypeComparisonCvterm{&b, &c} {
		err = a.SetPhenotypeComparisonCvterm(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.PhenotypeComparisonCvterm != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.PhenotypeComparison != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PhenotypeComparisonID != x.PhenotypeComparisonID {
			t.Error("foreign key was wrong value", a.PhenotypeComparisonID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PhenotypeComparisonID))
		reflect.Indirect(reflect.ValueOf(&x.PhenotypeComparisonID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PhenotypeComparisonID != x.PhenotypeComparisonID {
			t.Error("foreign key was wrong value", a.PhenotypeComparisonID, x.PhenotypeComparisonID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testPhenotypeComparisonToOneEnvironmentUsingEnvironment1(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PhenotypeComparison
	var foreign Environment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Environment1ID = foreign.EnvironmentID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Environment1(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.EnvironmentID != foreign.EnvironmentID {
		t.Errorf("want: %v, got %v", foreign.EnvironmentID, check.EnvironmentID)
	}

	slice := PhenotypeComparisonSlice{&local}
	if err = local.L.LoadEnvironment1(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Environment1 == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Environment1 = nil
	if err = local.L.LoadEnvironment1(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Environment1 == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeComparisonToOneEnvironmentUsingEnvironment2(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PhenotypeComparison
	var foreign Environment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Environment2ID = foreign.EnvironmentID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Environment2(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.EnvironmentID != foreign.EnvironmentID {
		t.Errorf("want: %v, got %v", foreign.EnvironmentID, check.EnvironmentID)
	}

	slice := PhenotypeComparisonSlice{&local}
	if err = local.L.LoadEnvironment2(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Environment2 == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Environment2 = nil
	if err = local.L.LoadEnvironment2(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Environment2 == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeComparisonToOneGenotypeUsingGenotype1(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PhenotypeComparison
	var foreign Genotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Genotype1ID = foreign.GenotypeID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Genotype1(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.GenotypeID != foreign.GenotypeID {
		t.Errorf("want: %v, got %v", foreign.GenotypeID, check.GenotypeID)
	}

	slice := PhenotypeComparisonSlice{&local}
	if err = local.L.LoadGenotype1(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Genotype1 == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Genotype1 = nil
	if err = local.L.LoadGenotype1(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Genotype1 == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeComparisonToOneGenotypeUsingGenotype2(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PhenotypeComparison
	var foreign Genotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Genotype2ID = foreign.GenotypeID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Genotype2(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.GenotypeID != foreign.GenotypeID {
		t.Errorf("want: %v, got %v", foreign.GenotypeID, check.GenotypeID)
	}

	slice := PhenotypeComparisonSlice{&local}
	if err = local.L.LoadGenotype2(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Genotype2 == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Genotype2 = nil
	if err = local.L.LoadGenotype2(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Genotype2 == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeComparisonToOneOrganismUsingOrganism(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PhenotypeComparison
	var foreign Organism

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.OrganismID = foreign.OrganismID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Organism(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.OrganismID != foreign.OrganismID {
		t.Errorf("want: %v, got %v", foreign.OrganismID, check.OrganismID)
	}

	slice := PhenotypeComparisonSlice{&local}
	if err = local.L.LoadOrganism(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Organism == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Organism = nil
	if err = local.L.LoadOrganism(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Organism == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeComparisonToOnePhenotypeUsingPhenotype1(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PhenotypeComparison
	var foreign Phenotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Phenotype1ID = foreign.PhenotypeID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Phenotype1(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PhenotypeID != foreign.PhenotypeID {
		t.Errorf("want: %v, got %v", foreign.PhenotypeID, check.PhenotypeID)
	}

	slice := PhenotypeComparisonSlice{&local}
	if err = local.L.LoadPhenotype1(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Phenotype1 == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Phenotype1 = nil
	if err = local.L.LoadPhenotype1(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Phenotype1 == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeComparisonToOnePhenotypeUsingPhenotype2(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PhenotypeComparison
	var foreign Phenotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	local.Phenotype2ID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Phenotype2ID.Int = foreign.PhenotypeID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Phenotype2(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PhenotypeID != foreign.PhenotypeID {
		t.Errorf("want: %v, got %v", foreign.PhenotypeID, check.PhenotypeID)
	}

	slice := PhenotypeComparisonSlice{&local}
	if err = local.L.LoadPhenotype2(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Phenotype2 == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Phenotype2 = nil
	if err = local.L.LoadPhenotype2(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Phenotype2 == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeComparisonToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PhenotypeComparison
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
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

	slice := PhenotypeComparisonSlice{&local}
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

func testPhenotypeComparisonToOneSetOpEnvironmentUsingEnvironment1(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PhenotypeComparison
	var b, c Environment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
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
		err = a.SetEnvironment1(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Environment1 != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Environment1PhenotypeComparison != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Environment1ID != x.EnvironmentID {
			t.Error("foreign key was wrong value", a.Environment1ID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Environment1ID))
		reflect.Indirect(reflect.ValueOf(&a.Environment1ID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.Environment1ID != x.EnvironmentID {
			t.Error("foreign key was wrong value", a.Environment1ID, x.EnvironmentID)
		}
	}
}
func testPhenotypeComparisonToOneSetOpEnvironmentUsingEnvironment2(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PhenotypeComparison
	var b, c Environment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
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
		err = a.SetEnvironment2(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Environment2 != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Environment2PhenotypeComparison != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Environment2ID != x.EnvironmentID {
			t.Error("foreign key was wrong value", a.Environment2ID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Environment2ID))
		reflect.Indirect(reflect.ValueOf(&a.Environment2ID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.Environment2ID != x.EnvironmentID {
			t.Error("foreign key was wrong value", a.Environment2ID, x.EnvironmentID)
		}
	}
}
func testPhenotypeComparisonToOneSetOpGenotypeUsingGenotype1(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PhenotypeComparison
	var b, c Genotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, genotypeDBTypes, false, strmangle.SetComplement(genotypePrimaryKeyColumns, genotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, genotypeDBTypes, false, strmangle.SetComplement(genotypePrimaryKeyColumns, genotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Genotype{&b, &c} {
		err = a.SetGenotype1(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Genotype1 != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Genotype1PhenotypeComparison != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Genotype1ID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.Genotype1ID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Genotype1ID))
		reflect.Indirect(reflect.ValueOf(&a.Genotype1ID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.Genotype1ID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.Genotype1ID, x.GenotypeID)
		}
	}
}
func testPhenotypeComparisonToOneSetOpGenotypeUsingGenotype2(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PhenotypeComparison
	var b, c Genotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, genotypeDBTypes, false, strmangle.SetComplement(genotypePrimaryKeyColumns, genotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, genotypeDBTypes, false, strmangle.SetComplement(genotypePrimaryKeyColumns, genotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Genotype{&b, &c} {
		err = a.SetGenotype2(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Genotype2 != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Genotype2PhenotypeComparison != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Genotype2ID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.Genotype2ID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Genotype2ID))
		reflect.Indirect(reflect.ValueOf(&a.Genotype2ID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.Genotype2ID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.Genotype2ID, x.GenotypeID)
		}
	}
}
func testPhenotypeComparisonToOneSetOpOrganismUsingOrganism(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PhenotypeComparison
	var b, c Organism

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organismDBTypes, false, strmangle.SetComplement(organismPrimaryKeyColumns, organismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, organismDBTypes, false, strmangle.SetComplement(organismPrimaryKeyColumns, organismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Organism{&b, &c} {
		err = a.SetOrganism(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Organism != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PhenotypeComparisons[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OrganismID != x.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OrganismID))
		reflect.Indirect(reflect.ValueOf(&a.OrganismID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrganismID != x.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID, x.OrganismID)
		}
	}
}
func testPhenotypeComparisonToOneSetOpPhenotypeUsingPhenotype1(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PhenotypeComparison
	var b, c Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Phenotype{&b, &c} {
		err = a.SetPhenotype1(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Phenotype1 != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Phenotype1PhenotypeComparison != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Phenotype1ID != x.PhenotypeID {
			t.Error("foreign key was wrong value", a.Phenotype1ID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Phenotype1ID))
		reflect.Indirect(reflect.ValueOf(&a.Phenotype1ID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.Phenotype1ID != x.PhenotypeID {
			t.Error("foreign key was wrong value", a.Phenotype1ID, x.PhenotypeID)
		}
	}
}
func testPhenotypeComparisonToOneSetOpPhenotypeUsingPhenotype2(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PhenotypeComparison
	var b, c Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Phenotype{&b, &c} {
		err = a.SetPhenotype2(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Phenotype2 != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Phenotype2PhenotypeComparisons[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Phenotype2ID.Int != x.PhenotypeID {
			t.Error("foreign key was wrong value", a.Phenotype2ID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Phenotype2ID.Int))
		reflect.Indirect(reflect.ValueOf(&a.Phenotype2ID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.Phenotype2ID.Int != x.PhenotypeID {
			t.Error("foreign key was wrong value", a.Phenotype2ID.Int, x.PhenotypeID)
		}
	}
}

func testPhenotypeComparisonToOneRemoveOpPhenotypeUsingPhenotype2(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PhenotypeComparison
	var b Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetPhenotype2(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemovePhenotype2(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Phenotype2(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Phenotype2 != nil {
		t.Error("R struct entry should be nil")
	}

	if a.Phenotype2ID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Phenotype2PhenotypeComparisons) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPhenotypeComparisonToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PhenotypeComparison
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
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

		if x.R.PhenotypeComparison != &a {
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
func testPhenotypeComparisonsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparison := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparison.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = phenotypeComparison.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPhenotypeComparisonsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparison := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparison.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PhenotypeComparisonSlice{phenotypeComparison}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testPhenotypeComparisonsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeComparison := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparison.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := PhenotypeComparisons(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	phenotypeComparisonDBTypes = map[string]string{"Environment1ID": "integer", "Environment2ID": "integer", "Genotype1ID": "integer", "Genotype2ID": "integer", "OrganismID": "integer", "Phenotype1ID": "integer", "Phenotype2ID": "integer", "PhenotypeComparisonID": "integer", "PubID": "integer"}
	_                          = bytes.MinRead
)

func testPhenotypeComparisonsUpdate(t *testing.T) {
	t.Parallel()

	if len(phenotypeComparisonColumns) == len(phenotypeComparisonPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	phenotypeComparison := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparison.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisons(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	if err = phenotypeComparison.Update(tx); err != nil {
		t.Error(err)
	}
}

func testPhenotypeComparisonsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(phenotypeComparisonColumns) == len(phenotypeComparisonPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	phenotypeComparison := &PhenotypeComparison{}
	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparison.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeComparisons(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, phenotypeComparison, phenotypeComparisonDBTypes, true, phenotypeComparisonPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(phenotypeComparisonColumns, phenotypeComparisonPrimaryKeyColumns) {
		fields = phenotypeComparisonColumns
	} else {
		fields = strmangle.SetComplement(
			phenotypeComparisonColumns,
			phenotypeComparisonPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(phenotypeComparison))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := PhenotypeComparisonSlice{phenotypeComparison}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testPhenotypeComparisonsUpsert(t *testing.T) {
	t.Parallel()

	if len(phenotypeComparisonColumns) == len(phenotypeComparisonPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	phenotypeComparison := PhenotypeComparison{}
	if err = randomize.Struct(seed, &phenotypeComparison, phenotypeComparisonDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeComparison.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert PhenotypeComparison: %s", err)
	}

	count, err := PhenotypeComparisons(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &phenotypeComparison, phenotypeComparisonDBTypes, false, phenotypeComparisonPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}

	if err = phenotypeComparison.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert PhenotypeComparison: %s", err)
	}

	count, err = PhenotypeComparisons(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
