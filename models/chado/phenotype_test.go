package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testPhenotypes(t *testing.T) {
	t.Parallel()

	query := Phenotypes(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testPhenotypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotype := &Phenotype{}
	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = phenotype.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhenotypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotype := &Phenotype{}
	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Phenotypes(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Phenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhenotypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotype := &Phenotype{}
	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PhenotypeSlice{phenotype}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testPhenotypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotype := &Phenotype{}
	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := PhenotypeExists(tx, phenotype.PhenotypeID)
	if err != nil {
		t.Errorf("Unable to check if Phenotype exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PhenotypeExistsG to return true, but got false.")
	}
}
func testPhenotypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotype := &Phenotype{}
	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	phenotypeFound, err := FindPhenotype(tx, phenotype.PhenotypeID)
	if err != nil {
		t.Error(err)
	}

	if phenotypeFound == nil {
		t.Error("want a record, got nil")
	}
}
func testPhenotypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotype := &Phenotype{}
	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Phenotypes(tx).Bind(phenotype); err != nil {
		t.Error(err)
	}
}

func testPhenotypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotype := &Phenotype{}
	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Phenotypes(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPhenotypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeOne := &Phenotype{}
	phenotypeTwo := &Phenotype{}
	if err = randomize.Struct(seed, phenotypeOne, phenotypeDBTypes, false, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}
	if err = randomize.Struct(seed, phenotypeTwo, phenotypeDBTypes, false, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = phenotypeTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Phenotypes(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPhenotypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	phenotypeOne := &Phenotype{}
	phenotypeTwo := &Phenotype{}
	if err = randomize.Struct(seed, phenotypeOne, phenotypeDBTypes, false, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}
	if err = randomize.Struct(seed, phenotypeTwo, phenotypeDBTypes, false, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = phenotypeTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func phenotypeBeforeInsertHook(e boil.Executor, o *Phenotype) error {
	*o = Phenotype{}
	return nil
}

func phenotypeAfterInsertHook(e boil.Executor, o *Phenotype) error {
	*o = Phenotype{}
	return nil
}

func phenotypeAfterSelectHook(e boil.Executor, o *Phenotype) error {
	*o = Phenotype{}
	return nil
}

func phenotypeBeforeUpdateHook(e boil.Executor, o *Phenotype) error {
	*o = Phenotype{}
	return nil
}

func phenotypeAfterUpdateHook(e boil.Executor, o *Phenotype) error {
	*o = Phenotype{}
	return nil
}

func phenotypeBeforeDeleteHook(e boil.Executor, o *Phenotype) error {
	*o = Phenotype{}
	return nil
}

func phenotypeAfterDeleteHook(e boil.Executor, o *Phenotype) error {
	*o = Phenotype{}
	return nil
}

func phenotypeBeforeUpsertHook(e boil.Executor, o *Phenotype) error {
	*o = Phenotype{}
	return nil
}

func phenotypeAfterUpsertHook(e boil.Executor, o *Phenotype) error {
	*o = Phenotype{}
	return nil
}

func testPhenotypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Phenotype{}
	o := &Phenotype{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, phenotypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Phenotype object: %s", err)
	}

	AddPhenotypeHook(boil.BeforeInsertHook, phenotypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	phenotypeBeforeInsertHooks = []PhenotypeHook{}

	AddPhenotypeHook(boil.AfterInsertHook, phenotypeAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	phenotypeAfterInsertHooks = []PhenotypeHook{}

	AddPhenotypeHook(boil.AfterSelectHook, phenotypeAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	phenotypeAfterSelectHooks = []PhenotypeHook{}

	AddPhenotypeHook(boil.BeforeUpdateHook, phenotypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	phenotypeBeforeUpdateHooks = []PhenotypeHook{}

	AddPhenotypeHook(boil.AfterUpdateHook, phenotypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	phenotypeAfterUpdateHooks = []PhenotypeHook{}

	AddPhenotypeHook(boil.BeforeDeleteHook, phenotypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	phenotypeBeforeDeleteHooks = []PhenotypeHook{}

	AddPhenotypeHook(boil.AfterDeleteHook, phenotypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	phenotypeAfterDeleteHooks = []PhenotypeHook{}

	AddPhenotypeHook(boil.BeforeUpsertHook, phenotypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	phenotypeBeforeUpsertHooks = []PhenotypeHook{}

	AddPhenotypeHook(boil.AfterUpsertHook, phenotypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	phenotypeAfterUpsertHooks = []PhenotypeHook{}
}
func testPhenotypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotype := &Phenotype{}
	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhenotypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotype := &Phenotype{}
	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotype.Insert(tx, phenotypeColumns...); err != nil {
		t.Error(err)
	}

	count, err := Phenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhenotypeOneToOnePhenotypepropUsingPhenotypeprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Phenotypeprop
	var local Phenotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenotypepropDBTypes, true, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PhenotypeID = local.PhenotypeID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Phenotypeprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PhenotypeID != foreign.PhenotypeID {
		t.Errorf("want: %v, got %v", foreign.PhenotypeID, check.PhenotypeID)
	}

	slice := PhenotypeSlice{&local}
	if err = local.L.LoadPhenotypeprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Phenotypeprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Phenotypeprop = nil
	if err = local.L.LoadPhenotypeprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Phenotypeprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeOneToOneFeaturePhenotypeUsingFeaturePhenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeaturePhenotype
	var local Phenotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featurePhenotypeDBTypes, true, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PhenotypeID = local.PhenotypeID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeaturePhenotype(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PhenotypeID != foreign.PhenotypeID {
		t.Errorf("want: %v, got %v", foreign.PhenotypeID, check.PhenotypeID)
	}

	slice := PhenotypeSlice{&local}
	if err = local.L.LoadFeaturePhenotype(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturePhenotype == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeaturePhenotype = nil
	if err = local.L.LoadFeaturePhenotype(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturePhenotype == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeOneToOnePhenstatementUsingPhenstatement(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Phenstatement
	var local Phenotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PhenotypeID = local.PhenotypeID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Phenstatement(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PhenotypeID != foreign.PhenotypeID {
		t.Errorf("want: %v, got %v", foreign.PhenotypeID, check.PhenotypeID)
	}

	slice := PhenotypeSlice{&local}
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

func testPhenotypeOneToOnePhenotypeCvtermUsingPhenotypeCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign PhenotypeCvterm
	var local Phenotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenotypeCvtermDBTypes, true, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PhenotypeID = local.PhenotypeID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.PhenotypeCvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PhenotypeID != foreign.PhenotypeID {
		t.Errorf("want: %v, got %v", foreign.PhenotypeID, check.PhenotypeID)
	}

	slice := PhenotypeSlice{&local}
	if err = local.L.LoadPhenotypeCvterm(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.PhenotypeCvterm == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.PhenotypeCvterm = nil
	if err = local.L.LoadPhenotypeCvterm(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.PhenotypeCvterm == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeOneToOnePhenotypeComparisonUsingPhenotype1PhenotypeComparison(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign PhenotypeComparison
	var local Phenotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.Phenotype1ID = local.PhenotypeID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Phenotype1PhenotypeComparison(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.Phenotype1ID != foreign.Phenotype1ID {
		t.Errorf("want: %v, got %v", foreign.Phenotype1ID, check.Phenotype1ID)
	}

	slice := PhenotypeSlice{&local}
	if err = local.L.LoadPhenotype1PhenotypeComparison(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Phenotype1PhenotypeComparison == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Phenotype1PhenotypeComparison = nil
	if err = local.L.LoadPhenotype1PhenotypeComparison(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Phenotype1PhenotypeComparison == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeOneToOneSetOpPhenotypepropUsingPhenotypeprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b, c Phenotypeprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, phenotypepropDBTypes, false, strmangle.SetComplement(phenotypepropPrimaryKeyColumns, phenotypepropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, phenotypepropDBTypes, false, strmangle.SetComplement(phenotypepropPrimaryKeyColumns, phenotypepropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Phenotypeprop{&b, &c} {
		err = a.SetPhenotypeprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Phenotypeprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Phenotype != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PhenotypeID != x.PhenotypeID {
			t.Error("foreign key was wrong value", a.PhenotypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PhenotypeID))
		reflect.Indirect(reflect.ValueOf(&x.PhenotypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PhenotypeID != x.PhenotypeID {
			t.Error("foreign key was wrong value", a.PhenotypeID, x.PhenotypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPhenotypeOneToOneSetOpFeaturePhenotypeUsingFeaturePhenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b, c FeaturePhenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featurePhenotypeDBTypes, false, strmangle.SetComplement(featurePhenotypePrimaryKeyColumns, featurePhenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featurePhenotypeDBTypes, false, strmangle.SetComplement(featurePhenotypePrimaryKeyColumns, featurePhenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeaturePhenotype{&b, &c} {
		err = a.SetFeaturePhenotype(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeaturePhenotype != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Phenotype != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PhenotypeID != x.PhenotypeID {
			t.Error("foreign key was wrong value", a.PhenotypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PhenotypeID))
		reflect.Indirect(reflect.ValueOf(&x.PhenotypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PhenotypeID != x.PhenotypeID {
			t.Error("foreign key was wrong value", a.PhenotypeID, x.PhenotypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPhenotypeOneToOneSetOpPhenstatementUsingPhenstatement(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b, c Phenstatement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
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
		if x.R.Phenotype != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PhenotypeID != x.PhenotypeID {
			t.Error("foreign key was wrong value", a.PhenotypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PhenotypeID))
		reflect.Indirect(reflect.ValueOf(&x.PhenotypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PhenotypeID != x.PhenotypeID {
			t.Error("foreign key was wrong value", a.PhenotypeID, x.PhenotypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPhenotypeOneToOneSetOpPhenotypeCvtermUsingPhenotypeCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b, c PhenotypeCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, phenotypeCvtermDBTypes, false, strmangle.SetComplement(phenotypeCvtermPrimaryKeyColumns, phenotypeCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, phenotypeCvtermDBTypes, false, strmangle.SetComplement(phenotypeCvtermPrimaryKeyColumns, phenotypeCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*PhenotypeCvterm{&b, &c} {
		err = a.SetPhenotypeCvterm(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.PhenotypeCvterm != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Phenotype != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PhenotypeID != x.PhenotypeID {
			t.Error("foreign key was wrong value", a.PhenotypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PhenotypeID))
		reflect.Indirect(reflect.ValueOf(&x.PhenotypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PhenotypeID != x.PhenotypeID {
			t.Error("foreign key was wrong value", a.PhenotypeID, x.PhenotypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPhenotypeOneToOneSetOpPhenotypeComparisonUsingPhenotype1PhenotypeComparison(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b, c PhenotypeComparison

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
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
		err = a.SetPhenotype1PhenotypeComparison(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Phenotype1PhenotypeComparison != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Phenotype1 != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PhenotypeID != x.Phenotype1ID {
			t.Error("foreign key was wrong value", a.PhenotypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.Phenotype1ID))
		reflect.Indirect(reflect.ValueOf(&x.Phenotype1ID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PhenotypeID != x.Phenotype1ID {
			t.Error("foreign key was wrong value", a.PhenotypeID, x.Phenotype1ID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPhenotypeToManyPhenotype2PhenotypeComparisons(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b, c PhenotypeComparison

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, phenotypeComparisonDBTypes, false, phenotypeComparisonColumnsWithDefault...)
	randomize.Struct(seed, &c, phenotypeComparisonDBTypes, false, phenotypeComparisonColumnsWithDefault...)
	b.Phenotype2ID.Valid = true
	c.Phenotype2ID.Valid = true
	b.Phenotype2ID.Int = a.PhenotypeID
	c.Phenotype2ID.Int = a.PhenotypeID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	phenotypeComparison, err := a.Phenotype2PhenotypeComparisons(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range phenotypeComparison {
		if v.Phenotype2ID.Int == b.Phenotype2ID.Int {
			bFound = true
		}
		if v.Phenotype2ID.Int == c.Phenotype2ID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PhenotypeSlice{&a}
	if err = a.L.LoadPhenotype2PhenotypeComparisons(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Phenotype2PhenotypeComparisons); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Phenotype2PhenotypeComparisons = nil
	if err = a.L.LoadPhenotype2PhenotypeComparisons(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Phenotype2PhenotypeComparisons); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", phenotypeComparison)
	}
}

func testPhenotypeToManyAddOpPhenotype2PhenotypeComparisons(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b, c, d, e PhenotypeComparison

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PhenotypeComparison{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*PhenotypeComparison{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPhenotype2PhenotypeComparisons(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.PhenotypeID != first.Phenotype2ID.Int {
			t.Error("foreign key was wrong value", a.PhenotypeID, first.Phenotype2ID.Int)
		}
		if a.PhenotypeID != second.Phenotype2ID.Int {
			t.Error("foreign key was wrong value", a.PhenotypeID, second.Phenotype2ID.Int)
		}

		if first.R.Phenotype2 != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Phenotype2 != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Phenotype2PhenotypeComparisons[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Phenotype2PhenotypeComparisons[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Phenotype2PhenotypeComparisons(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testPhenotypeToManySetOpPhenotype2PhenotypeComparisons(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b, c, d, e PhenotypeComparison

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PhenotypeComparison{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetPhenotype2PhenotypeComparisons(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Phenotype2PhenotypeComparisons(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetPhenotype2PhenotypeComparisons(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Phenotype2PhenotypeComparisons(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.Phenotype2ID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.Phenotype2ID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.PhenotypeID != d.Phenotype2ID.Int {
		t.Error("foreign key was wrong value", a.PhenotypeID, d.Phenotype2ID.Int)
	}
	if a.PhenotypeID != e.Phenotype2ID.Int {
		t.Error("foreign key was wrong value", a.PhenotypeID, e.Phenotype2ID.Int)
	}

	if b.R.Phenotype2 != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Phenotype2 != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Phenotype2 != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Phenotype2 != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Phenotype2PhenotypeComparisons[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Phenotype2PhenotypeComparisons[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testPhenotypeToManyRemoveOpPhenotype2PhenotypeComparisons(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b, c, d, e PhenotypeComparison

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PhenotypeComparison{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddPhenotype2PhenotypeComparisons(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Phenotype2PhenotypeComparisons(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemovePhenotype2PhenotypeComparisons(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Phenotype2PhenotypeComparisons(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.Phenotype2ID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.Phenotype2ID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Phenotype2 != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Phenotype2 != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Phenotype2 != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Phenotype2 != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Phenotype2PhenotypeComparisons) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Phenotype2PhenotypeComparisons[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Phenotype2PhenotypeComparisons[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testPhenotypeToOneCvtermUsingAssay(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Phenotype
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	local.AssayID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AssayID.Int = foreign.CvtermID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Assay(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := PhenotypeSlice{&local}
	if err = local.L.LoadAssay(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Assay == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Assay = nil
	if err = local.L.LoadAssay(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Assay == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeToOneCvtermUsingAttr(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Phenotype
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	local.AttrID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AttrID.Int = foreign.CvtermID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Attr(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := PhenotypeSlice{&local}
	if err = local.L.LoadAttr(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Attr == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Attr = nil
	if err = local.L.LoadAttr(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Attr == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeToOneCvtermUsingCvalue(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Phenotype
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	local.CvalueID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.CvalueID.Int = foreign.CvtermID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Cvalue(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := PhenotypeSlice{&local}
	if err = local.L.LoadCvalue(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Cvalue == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Cvalue = nil
	if err = local.L.LoadCvalue(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Cvalue == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeToOneCvtermUsingObservable(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Phenotype
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	local.ObservableID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ObservableID.Int = foreign.CvtermID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Observable(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := PhenotypeSlice{&local}
	if err = local.L.LoadObservable(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Observable == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Observable = nil
	if err = local.L.LoadObservable(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Observable == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhenotypeToOneSetOpCvtermUsingAssay(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
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
		err = a.SetAssay(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Assay != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AssayPhenotypes[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AssayID.Int != x.CvtermID {
			t.Error("foreign key was wrong value", a.AssayID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AssayID.Int))
		reflect.Indirect(reflect.ValueOf(&a.AssayID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AssayID.Int != x.CvtermID {
			t.Error("foreign key was wrong value", a.AssayID.Int, x.CvtermID)
		}
	}
}

func testPhenotypeToOneRemoveOpCvtermUsingAssay(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetAssay(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveAssay(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Assay(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Assay != nil {
		t.Error("R struct entry should be nil")
	}

	if a.AssayID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.AssayPhenotypes) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPhenotypeToOneSetOpCvtermUsingAttr(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
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
		err = a.SetAttr(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Attr != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AttrPhenotypes[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AttrID.Int != x.CvtermID {
			t.Error("foreign key was wrong value", a.AttrID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AttrID.Int))
		reflect.Indirect(reflect.ValueOf(&a.AttrID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AttrID.Int != x.CvtermID {
			t.Error("foreign key was wrong value", a.AttrID.Int, x.CvtermID)
		}
	}
}

func testPhenotypeToOneRemoveOpCvtermUsingAttr(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetAttr(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveAttr(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Attr(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Attr != nil {
		t.Error("R struct entry should be nil")
	}

	if a.AttrID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.AttrPhenotypes) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPhenotypeToOneSetOpCvtermUsingCvalue(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
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
		err = a.SetCvalue(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Cvalue != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.CvaluePhenotypes[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CvalueID.Int != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvalueID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CvalueID.Int))
		reflect.Indirect(reflect.ValueOf(&a.CvalueID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvalueID.Int != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvalueID.Int, x.CvtermID)
		}
	}
}

func testPhenotypeToOneRemoveOpCvtermUsingCvalue(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetCvalue(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveCvalue(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Cvalue(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Cvalue != nil {
		t.Error("R struct entry should be nil")
	}

	if a.CvalueID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.CvaluePhenotypes) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPhenotypeToOneSetOpCvtermUsingObservable(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
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
		err = a.SetObservable(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Observable != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ObservablePhenotypes[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ObservableID.Int != x.CvtermID {
			t.Error("foreign key was wrong value", a.ObservableID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ObservableID.Int))
		reflect.Indirect(reflect.ValueOf(&a.ObservableID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ObservableID.Int != x.CvtermID {
			t.Error("foreign key was wrong value", a.ObservableID.Int, x.CvtermID)
		}
	}
}

func testPhenotypeToOneRemoveOpCvtermUsingObservable(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotype
	var b Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetObservable(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveObservable(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Observable(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Observable != nil {
		t.Error("R struct entry should be nil")
	}

	if a.ObservableID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.ObservablePhenotypes) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPhenotypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotype := &Phenotype{}
	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = phenotype.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPhenotypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotype := &Phenotype{}
	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PhenotypeSlice{phenotype}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testPhenotypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotype := &Phenotype{}
	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Phenotypes(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	phenotypeDBTypes = map[string]string{"AssayID": "integer", "AttrID": "integer", "CvalueID": "integer", "Name": "text", "ObservableID": "integer", "PhenotypeID": "integer", "Uniquename": "text", "Value": "text"}
	_                = bytes.MinRead
)

func testPhenotypesUpdate(t *testing.T) {
	t.Parallel()

	if len(phenotypeColumns) == len(phenotypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	phenotype := &Phenotype{}
	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true, phenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	if err = phenotype.Update(tx); err != nil {
		t.Error(err)
	}
}

func testPhenotypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(phenotypeColumns) == len(phenotypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	phenotype := &Phenotype{}
	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, phenotype, phenotypeDBTypes, true, phenotypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(phenotypeColumns, phenotypePrimaryKeyColumns) {
		fields = phenotypeColumns
	} else {
		fields = strmangle.SetComplement(
			phenotypeColumns,
			phenotypePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(phenotype))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := PhenotypeSlice{phenotype}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testPhenotypesUpsert(t *testing.T) {
	t.Parallel()

	if len(phenotypeColumns) == len(phenotypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	phenotype := Phenotype{}
	if err = randomize.Struct(seed, &phenotype, phenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotype.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Phenotype: %s", err)
	}

	count, err := Phenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &phenotype, phenotypeDBTypes, false, phenotypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Phenotype struct: %s", err)
	}

	if err = phenotype.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Phenotype: %s", err)
	}

	count, err = Phenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
