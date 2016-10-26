package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testEnvironments(t *testing.T) {
	t.Parallel()

	query := Environments(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testEnvironmentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environment := &Environment{}
	if err = randomize.Struct(seed, environment, environmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environment.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = environment.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Environments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEnvironmentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environment := &Environment{}
	if err = randomize.Struct(seed, environment, environmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environment.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Environments(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Environments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEnvironmentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environment := &Environment{}
	if err = randomize.Struct(seed, environment, environmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environment.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := EnvironmentSlice{environment}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Environments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testEnvironmentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environment := &Environment{}
	if err = randomize.Struct(seed, environment, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environment.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := EnvironmentExists(tx, environment.EnvironmentID)
	if err != nil {
		t.Errorf("Unable to check if Environment exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EnvironmentExistsG to return true, but got false.")
	}
}
func testEnvironmentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environment := &Environment{}
	if err = randomize.Struct(seed, environment, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environment.Insert(tx); err != nil {
		t.Error(err)
	}

	environmentFound, err := FindEnvironment(tx, environment.EnvironmentID)
	if err != nil {
		t.Error(err)
	}

	if environmentFound == nil {
		t.Error("want a record, got nil")
	}
}
func testEnvironmentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environment := &Environment{}
	if err = randomize.Struct(seed, environment, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environment.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Environments(tx).Bind(environment); err != nil {
		t.Error(err)
	}
}

func testEnvironmentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environment := &Environment{}
	if err = randomize.Struct(seed, environment, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environment.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Environments(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEnvironmentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environmentOne := &Environment{}
	environmentTwo := &Environment{}
	if err = randomize.Struct(seed, environmentOne, environmentDBTypes, false, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}
	if err = randomize.Struct(seed, environmentTwo, environmentDBTypes, false, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = environmentTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Environments(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEnvironmentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	environmentOne := &Environment{}
	environmentTwo := &Environment{}
	if err = randomize.Struct(seed, environmentOne, environmentDBTypes, false, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}
	if err = randomize.Struct(seed, environmentTwo, environmentDBTypes, false, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environmentOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = environmentTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Environments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func environmentBeforeInsertHook(e boil.Executor, o *Environment) error {
	*o = Environment{}
	return nil
}

func environmentAfterInsertHook(e boil.Executor, o *Environment) error {
	*o = Environment{}
	return nil
}

func environmentAfterSelectHook(e boil.Executor, o *Environment) error {
	*o = Environment{}
	return nil
}

func environmentBeforeUpdateHook(e boil.Executor, o *Environment) error {
	*o = Environment{}
	return nil
}

func environmentAfterUpdateHook(e boil.Executor, o *Environment) error {
	*o = Environment{}
	return nil
}

func environmentBeforeDeleteHook(e boil.Executor, o *Environment) error {
	*o = Environment{}
	return nil
}

func environmentAfterDeleteHook(e boil.Executor, o *Environment) error {
	*o = Environment{}
	return nil
}

func environmentBeforeUpsertHook(e boil.Executor, o *Environment) error {
	*o = Environment{}
	return nil
}

func environmentAfterUpsertHook(e boil.Executor, o *Environment) error {
	*o = Environment{}
	return nil
}

func testEnvironmentsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Environment{}
	o := &Environment{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, environmentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Environment object: %s", err)
	}

	AddEnvironmentHook(boil.BeforeInsertHook, environmentBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	environmentBeforeInsertHooks = []EnvironmentHook{}

	AddEnvironmentHook(boil.AfterInsertHook, environmentAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	environmentAfterInsertHooks = []EnvironmentHook{}

	AddEnvironmentHook(boil.AfterSelectHook, environmentAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	environmentAfterSelectHooks = []EnvironmentHook{}

	AddEnvironmentHook(boil.BeforeUpdateHook, environmentBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	environmentBeforeUpdateHooks = []EnvironmentHook{}

	AddEnvironmentHook(boil.AfterUpdateHook, environmentAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	environmentAfterUpdateHooks = []EnvironmentHook{}

	AddEnvironmentHook(boil.BeforeDeleteHook, environmentBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	environmentBeforeDeleteHooks = []EnvironmentHook{}

	AddEnvironmentHook(boil.AfterDeleteHook, environmentAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	environmentAfterDeleteHooks = []EnvironmentHook{}

	AddEnvironmentHook(boil.BeforeUpsertHook, environmentBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	environmentBeforeUpsertHooks = []EnvironmentHook{}

	AddEnvironmentHook(boil.AfterUpsertHook, environmentAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	environmentAfterUpsertHooks = []EnvironmentHook{}
}
func testEnvironmentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environment := &Environment{}
	if err = randomize.Struct(seed, environment, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environment.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Environments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEnvironmentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environment := &Environment{}
	if err = randomize.Struct(seed, environment, environmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environment.Insert(tx, environmentColumns...); err != nil {
		t.Error(err)
	}

	count, err := Environments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEnvironmentOneToOneEnvironmentCvtermUsingEnvironmentCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign EnvironmentCvterm
	var local Environment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, environmentCvtermDBTypes, true, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.EnvironmentID = local.EnvironmentID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.EnvironmentCvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.EnvironmentID != foreign.EnvironmentID {
		t.Errorf("want: %v, got %v", foreign.EnvironmentID, check.EnvironmentID)
	}

	slice := EnvironmentSlice{&local}
	if err = local.L.LoadEnvironmentCvterm(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.EnvironmentCvterm == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.EnvironmentCvterm = nil
	if err = local.L.LoadEnvironmentCvterm(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.EnvironmentCvterm == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testEnvironmentOneToOnePhendescUsingPhendesc(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Phendesc
	var local Environment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.EnvironmentID = local.EnvironmentID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Phendesc(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.EnvironmentID != foreign.EnvironmentID {
		t.Errorf("want: %v, got %v", foreign.EnvironmentID, check.EnvironmentID)
	}

	slice := EnvironmentSlice{&local}
	if err = local.L.LoadPhendesc(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Phendesc == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Phendesc = nil
	if err = local.L.LoadPhendesc(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Phendesc == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testEnvironmentOneToOnePhenstatementUsingPhenstatement(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Phenstatement
	var local Environment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.EnvironmentID = local.EnvironmentID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Phenstatement(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.EnvironmentID != foreign.EnvironmentID {
		t.Errorf("want: %v, got %v", foreign.EnvironmentID, check.EnvironmentID)
	}

	slice := EnvironmentSlice{&local}
	if err = local.L.LoadPhenstatement(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Phenstatement == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Phenstatement = nil
	if err = local.L.LoadPhenstatement(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Phenstatement == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testEnvironmentOneToOnePhenotypeComparisonUsingEnvironment1PhenotypeComparison(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign PhenotypeComparison
	var local Environment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.Environment1ID = local.EnvironmentID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Environment1PhenotypeComparison(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.Environment1ID != foreign.Environment1ID {
		t.Errorf("want: %v, got %v", foreign.Environment1ID, check.Environment1ID)
	}

	slice := EnvironmentSlice{&local}
	if err = local.L.LoadEnvironment1PhenotypeComparison(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Environment1PhenotypeComparison == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Environment1PhenotypeComparison = nil
	if err = local.L.LoadEnvironment1PhenotypeComparison(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Environment1PhenotypeComparison == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testEnvironmentOneToOnePhenotypeComparisonUsingEnvironment2PhenotypeComparison(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign PhenotypeComparison
	var local Environment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.Environment2ID = local.EnvironmentID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Environment2PhenotypeComparison(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.Environment2ID != foreign.Environment2ID {
		t.Errorf("want: %v, got %v", foreign.Environment2ID, check.Environment2ID)
	}

	slice := EnvironmentSlice{&local}
	if err = local.L.LoadEnvironment2PhenotypeComparison(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Environment2PhenotypeComparison == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Environment2PhenotypeComparison = nil
	if err = local.L.LoadEnvironment2PhenotypeComparison(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Environment2PhenotypeComparison == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testEnvironmentOneToOneSetOpEnvironmentCvtermUsingEnvironmentCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Environment
	var b, c EnvironmentCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, environmentCvtermDBTypes, false, strmangle.SetComplement(environmentCvtermPrimaryKeyColumns, environmentCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, environmentCvtermDBTypes, false, strmangle.SetComplement(environmentCvtermPrimaryKeyColumns, environmentCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*EnvironmentCvterm{&b, &c} {
		err = a.SetEnvironmentCvterm(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.EnvironmentCvterm != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Environment != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.EnvironmentID != x.EnvironmentID {
			t.Error("foreign key was wrong value", a.EnvironmentID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.EnvironmentID))
		reflect.Indirect(reflect.ValueOf(&x.EnvironmentID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.EnvironmentID != x.EnvironmentID {
			t.Error("foreign key was wrong value", a.EnvironmentID, x.EnvironmentID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testEnvironmentOneToOneSetOpPhendescUsingPhendesc(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Environment
	var b, c Phendesc

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, phendescDBTypes, false, strmangle.SetComplement(phendescPrimaryKeyColumns, phendescColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, phendescDBTypes, false, strmangle.SetComplement(phendescPrimaryKeyColumns, phendescColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Phendesc{&b, &c} {
		err = a.SetPhendesc(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Phendesc != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Environment != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.EnvironmentID != x.EnvironmentID {
			t.Error("foreign key was wrong value", a.EnvironmentID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.EnvironmentID))
		reflect.Indirect(reflect.ValueOf(&x.EnvironmentID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.EnvironmentID != x.EnvironmentID {
			t.Error("foreign key was wrong value", a.EnvironmentID, x.EnvironmentID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testEnvironmentOneToOneSetOpPhenstatementUsingPhenstatement(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Environment
	var b, c Phenstatement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, phenstatementDBTypes, false, strmangle.SetComplement(phenstatementPrimaryKeyColumns, phenstatementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, phenstatementDBTypes, false, strmangle.SetComplement(phenstatementPrimaryKeyColumns, phenstatementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Phenstatement{&b, &c} {
		err = a.SetPhenstatement(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Phenstatement != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Environment != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.EnvironmentID != x.EnvironmentID {
			t.Error("foreign key was wrong value", a.EnvironmentID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.EnvironmentID))
		reflect.Indirect(reflect.ValueOf(&x.EnvironmentID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.EnvironmentID != x.EnvironmentID {
			t.Error("foreign key was wrong value", a.EnvironmentID, x.EnvironmentID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testEnvironmentOneToOneSetOpPhenotypeComparisonUsingEnvironment1PhenotypeComparison(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Environment
	var b, c PhenotypeComparison

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
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
		err = a.SetEnvironment1PhenotypeComparison(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Environment1PhenotypeComparison != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Environment1 != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.EnvironmentID != x.Environment1ID {
			t.Error("foreign key was wrong value", a.EnvironmentID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.Environment1ID))
		reflect.Indirect(reflect.ValueOf(&x.Environment1ID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.EnvironmentID != x.Environment1ID {
			t.Error("foreign key was wrong value", a.EnvironmentID, x.Environment1ID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testEnvironmentOneToOneSetOpPhenotypeComparisonUsingEnvironment2PhenotypeComparison(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Environment
	var b, c PhenotypeComparison

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
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
		err = a.SetEnvironment2PhenotypeComparison(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Environment2PhenotypeComparison != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Environment2 != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.EnvironmentID != x.Environment2ID {
			t.Error("foreign key was wrong value", a.EnvironmentID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.Environment2ID))
		reflect.Indirect(reflect.ValueOf(&x.Environment2ID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.EnvironmentID != x.Environment2ID {
			t.Error("foreign key was wrong value", a.EnvironmentID, x.Environment2ID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testEnvironmentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environment := &Environment{}
	if err = randomize.Struct(seed, environment, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environment.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = environment.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testEnvironmentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environment := &Environment{}
	if err = randomize.Struct(seed, environment, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environment.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := EnvironmentSlice{environment}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testEnvironmentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	environment := &Environment{}
	if err = randomize.Struct(seed, environment, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environment.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Environments(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	environmentDBTypes = map[string]string{"Description": "text", "EnvironmentID": "integer", "Uniquename": "text"}
	_                  = bytes.MinRead
)

func testEnvironmentsUpdate(t *testing.T) {
	t.Parallel()

	if len(environmentColumns) == len(environmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	environment := &Environment{}
	if err = randomize.Struct(seed, environment, environmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environment.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Environments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, environment, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	if err = environment.Update(tx); err != nil {
		t.Error(err)
	}
}

func testEnvironmentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(environmentColumns) == len(environmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	environment := &Environment{}
	if err = randomize.Struct(seed, environment, environmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environment.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Environments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, environment, environmentDBTypes, true, environmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(environmentColumns, environmentPrimaryKeyColumns) {
		fields = environmentColumns
	} else {
		fields = strmangle.SetComplement(
			environmentColumns,
			environmentPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(environment))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := EnvironmentSlice{environment}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testEnvironmentsUpsert(t *testing.T) {
	t.Parallel()

	if len(environmentColumns) == len(environmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	environment := Environment{}
	if err = randomize.Struct(seed, &environment, environmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = environment.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Environment: %s", err)
	}

	count, err := Environments(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &environment, environmentDBTypes, false, environmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	if err = environment.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Environment: %s", err)
	}

	count, err = Environments(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
