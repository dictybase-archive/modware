package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testPhenstatements(t *testing.T) {
	t.Parallel()

	query := Phenstatements(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testPhenstatementsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenstatement := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatement.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = phenstatement.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenstatements(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhenstatementsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenstatement := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatement.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Phenstatements(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Phenstatements(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhenstatementsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenstatement := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatement.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PhenstatementSlice{phenstatement}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenstatements(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testPhenstatementsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenstatement := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatement.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := PhenstatementExists(tx, phenstatement.PhenstatementID)
	if err != nil {
		t.Errorf("Unable to check if Phenstatement exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PhenstatementExistsG to return true, but got false.")
	}
}
func testPhenstatementsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenstatement := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatement.Insert(tx); err != nil {
		t.Error(err)
	}

	phenstatementFound, err := FindPhenstatement(tx, phenstatement.PhenstatementID)
	if err != nil {
		t.Error(err)
	}

	if phenstatementFound == nil {
		t.Error("want a record, got nil")
	}
}
func testPhenstatementsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenstatement := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatement.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Phenstatements(tx).Bind(phenstatement); err != nil {
		t.Error(err)
	}
}

func testPhenstatementsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenstatement := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatement.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Phenstatements(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPhenstatementsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenstatementOne := &Phenstatement{}
	phenstatementTwo := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatementOne, phenstatementDBTypes, false, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}
	if err = randomize.Struct(seed, phenstatementTwo, phenstatementDBTypes, false, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatementOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = phenstatementTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Phenstatements(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPhenstatementsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	phenstatementOne := &Phenstatement{}
	phenstatementTwo := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatementOne, phenstatementDBTypes, false, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}
	if err = randomize.Struct(seed, phenstatementTwo, phenstatementDBTypes, false, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatementOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = phenstatementTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenstatements(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func phenstatementBeforeInsertHook(e boil.Executor, o *Phenstatement) error {
	*o = Phenstatement{}
	return nil
}

func phenstatementAfterInsertHook(e boil.Executor, o *Phenstatement) error {
	*o = Phenstatement{}
	return nil
}

func phenstatementAfterSelectHook(e boil.Executor, o *Phenstatement) error {
	*o = Phenstatement{}
	return nil
}

func phenstatementBeforeUpdateHook(e boil.Executor, o *Phenstatement) error {
	*o = Phenstatement{}
	return nil
}

func phenstatementAfterUpdateHook(e boil.Executor, o *Phenstatement) error {
	*o = Phenstatement{}
	return nil
}

func phenstatementBeforeDeleteHook(e boil.Executor, o *Phenstatement) error {
	*o = Phenstatement{}
	return nil
}

func phenstatementAfterDeleteHook(e boil.Executor, o *Phenstatement) error {
	*o = Phenstatement{}
	return nil
}

func phenstatementBeforeUpsertHook(e boil.Executor, o *Phenstatement) error {
	*o = Phenstatement{}
	return nil
}

func phenstatementAfterUpsertHook(e boil.Executor, o *Phenstatement) error {
	*o = Phenstatement{}
	return nil
}

func testPhenstatementsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Phenstatement{}
	o := &Phenstatement{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, phenstatementDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Phenstatement object: %s", err)
	}

	AddPhenstatementHook(boil.BeforeInsertHook, phenstatementBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	phenstatementBeforeInsertHooks = []PhenstatementHook{}

	AddPhenstatementHook(boil.AfterInsertHook, phenstatementAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	phenstatementAfterInsertHooks = []PhenstatementHook{}

	AddPhenstatementHook(boil.AfterSelectHook, phenstatementAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	phenstatementAfterSelectHooks = []PhenstatementHook{}

	AddPhenstatementHook(boil.BeforeUpdateHook, phenstatementBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	phenstatementBeforeUpdateHooks = []PhenstatementHook{}

	AddPhenstatementHook(boil.AfterUpdateHook, phenstatementAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	phenstatementAfterUpdateHooks = []PhenstatementHook{}

	AddPhenstatementHook(boil.BeforeDeleteHook, phenstatementBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	phenstatementBeforeDeleteHooks = []PhenstatementHook{}

	AddPhenstatementHook(boil.AfterDeleteHook, phenstatementAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	phenstatementAfterDeleteHooks = []PhenstatementHook{}

	AddPhenstatementHook(boil.BeforeUpsertHook, phenstatementBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	phenstatementBeforeUpsertHooks = []PhenstatementHook{}

	AddPhenstatementHook(boil.AfterUpsertHook, phenstatementAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	phenstatementAfterUpsertHooks = []PhenstatementHook{}
}
func testPhenstatementsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenstatement := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatement.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenstatements(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhenstatementsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenstatement := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatement.Insert(tx, phenstatementColumns...); err != nil {
		t.Error(err)
	}

	count, err := Phenstatements(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhenstatementToOneEnvironmentUsingEnvironment(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Phenstatement
	var foreign Environment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
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

	slice := PhenstatementSlice{&local}
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

func testPhenstatementToOneGenotypeUsingGenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Phenstatement
	var foreign Genotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.GenotypeID = foreign.GenotypeID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Genotype(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.GenotypeID != foreign.GenotypeID {
		t.Errorf("want: %v, got %v", foreign.GenotypeID, check.GenotypeID)
	}

	slice := PhenstatementSlice{&local}
	if err = local.L.LoadGenotype(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Genotype == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Genotype = nil
	if err = local.L.LoadGenotype(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Genotype == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenstatementToOnePhenotypeUsingPhenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Phenstatement
	var foreign Phenotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.PhenotypeID = foreign.PhenotypeID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Phenotype(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PhenotypeID != foreign.PhenotypeID {
		t.Errorf("want: %v, got %v", foreign.PhenotypeID, check.PhenotypeID)
	}

	slice := PhenstatementSlice{&local}
	if err = local.L.LoadPhenotype(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Phenotype == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Phenotype = nil
	if err = local.L.LoadPhenotype(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Phenotype == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenstatementToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Phenstatement
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
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

	slice := PhenstatementSlice{&local}
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

func testPhenstatementToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Phenstatement
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
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

	slice := PhenstatementSlice{&local}
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

func testPhenstatementToOneSetOpEnvironmentUsingEnvironment(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenstatement
	var b, c Environment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenstatementDBTypes, false, strmangle.SetComplement(phenstatementPrimaryKeyColumns, phenstatementColumnsWithoutDefault)...); err != nil {
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

		if x.R.Phenstatement != &a {
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
func testPhenstatementToOneSetOpGenotypeUsingGenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenstatement
	var b, c Genotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenstatementDBTypes, false, strmangle.SetComplement(phenstatementPrimaryKeyColumns, phenstatementColumnsWithoutDefault)...); err != nil {
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
		err = a.SetGenotype(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Genotype != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Phenstatement != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.GenotypeID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.GenotypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.GenotypeID))
		reflect.Indirect(reflect.ValueOf(&a.GenotypeID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GenotypeID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.GenotypeID, x.GenotypeID)
		}
	}
}
func testPhenstatementToOneSetOpPhenotypeUsingPhenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenstatement
	var b, c Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenstatementDBTypes, false, strmangle.SetComplement(phenstatementPrimaryKeyColumns, phenstatementColumnsWithoutDefault)...); err != nil {
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
		err = a.SetPhenotype(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Phenotype != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Phenstatement != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PhenotypeID != x.PhenotypeID {
			t.Error("foreign key was wrong value", a.PhenotypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PhenotypeID))
		reflect.Indirect(reflect.ValueOf(&a.PhenotypeID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PhenotypeID != x.PhenotypeID {
			t.Error("foreign key was wrong value", a.PhenotypeID, x.PhenotypeID)
		}
	}
}
func testPhenstatementToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenstatement
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenstatementDBTypes, false, strmangle.SetComplement(phenstatementPrimaryKeyColumns, phenstatementColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypePhenstatement != &a {
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
func testPhenstatementToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenstatement
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenstatementDBTypes, false, strmangle.SetComplement(phenstatementPrimaryKeyColumns, phenstatementColumnsWithoutDefault)...); err != nil {
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

		if x.R.Phenstatement != &a {
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
func testPhenstatementsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenstatement := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatement.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = phenstatement.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPhenstatementsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenstatement := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatement.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PhenstatementSlice{phenstatement}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testPhenstatementsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenstatement := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatement.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Phenstatements(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	phenstatementDBTypes = map[string]string{"EnvironmentID": "integer", "GenotypeID": "integer", "PhenotypeID": "integer", "PhenstatementID": "integer", "PubID": "integer", "TypeID": "integer"}
	_                    = bytes.MinRead
)

func testPhenstatementsUpdate(t *testing.T) {
	t.Parallel()

	if len(phenstatementColumns) == len(phenstatementPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	phenstatement := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatement.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenstatements(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	if err = phenstatement.Update(tx); err != nil {
		t.Error(err)
	}
}

func testPhenstatementsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(phenstatementColumns) == len(phenstatementPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	phenstatement := &Phenstatement{}
	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatement.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenstatements(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, phenstatement, phenstatementDBTypes, true, phenstatementPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(phenstatementColumns, phenstatementPrimaryKeyColumns) {
		fields = phenstatementColumns
	} else {
		fields = strmangle.SetComplement(
			phenstatementColumns,
			phenstatementPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(phenstatement))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := PhenstatementSlice{phenstatement}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testPhenstatementsUpsert(t *testing.T) {
	t.Parallel()

	if len(phenstatementColumns) == len(phenstatementPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	phenstatement := Phenstatement{}
	if err = randomize.Struct(seed, &phenstatement, phenstatementDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenstatement.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Phenstatement: %s", err)
	}

	count, err := Phenstatements(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &phenstatement, phenstatementDBTypes, false, phenstatementPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}

	if err = phenstatement.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Phenstatement: %s", err)
	}

	count, err = Phenstatements(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
