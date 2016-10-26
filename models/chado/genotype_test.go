package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testGenotypes(t *testing.T) {
	t.Parallel()

	query := Genotypes(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testGenotypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotype := &Genotype{}
	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = genotype.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Genotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGenotypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotype := &Genotype{}
	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Genotypes(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Genotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGenotypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotype := &Genotype{}
	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotype.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := GenotypeSlice{genotype}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Genotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testGenotypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotype := &Genotype{}
	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotype.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := GenotypeExists(tx, genotype.GenotypeID)
	if err != nil {
		t.Errorf("Unable to check if Genotype exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GenotypeExistsG to return true, but got false.")
	}
}
func testGenotypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotype := &Genotype{}
	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotype.Insert(tx); err != nil {
		t.Error(err)
	}

	genotypeFound, err := FindGenotype(tx, genotype.GenotypeID)
	if err != nil {
		t.Error(err)
	}

	if genotypeFound == nil {
		t.Error("want a record, got nil")
	}
}
func testGenotypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotype := &Genotype{}
	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Genotypes(tx).Bind(genotype); err != nil {
		t.Error(err)
	}
}

func testGenotypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotype := &Genotype{}
	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Genotypes(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGenotypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotypeOne := &Genotype{}
	genotypeTwo := &Genotype{}
	if err = randomize.Struct(seed, genotypeOne, genotypeDBTypes, false, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}
	if err = randomize.Struct(seed, genotypeTwo, genotypeDBTypes, false, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = genotypeTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Genotypes(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGenotypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	genotypeOne := &Genotype{}
	genotypeTwo := &Genotype{}
	if err = randomize.Struct(seed, genotypeOne, genotypeDBTypes, false, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}
	if err = randomize.Struct(seed, genotypeTwo, genotypeDBTypes, false, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = genotypeTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Genotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func genotypeBeforeInsertHook(e boil.Executor, o *Genotype) error {
	*o = Genotype{}
	return nil
}

func genotypeAfterInsertHook(e boil.Executor, o *Genotype) error {
	*o = Genotype{}
	return nil
}

func genotypeAfterSelectHook(e boil.Executor, o *Genotype) error {
	*o = Genotype{}
	return nil
}

func genotypeBeforeUpdateHook(e boil.Executor, o *Genotype) error {
	*o = Genotype{}
	return nil
}

func genotypeAfterUpdateHook(e boil.Executor, o *Genotype) error {
	*o = Genotype{}
	return nil
}

func genotypeBeforeDeleteHook(e boil.Executor, o *Genotype) error {
	*o = Genotype{}
	return nil
}

func genotypeAfterDeleteHook(e boil.Executor, o *Genotype) error {
	*o = Genotype{}
	return nil
}

func genotypeBeforeUpsertHook(e boil.Executor, o *Genotype) error {
	*o = Genotype{}
	return nil
}

func genotypeAfterUpsertHook(e boil.Executor, o *Genotype) error {
	*o = Genotype{}
	return nil
}

func testGenotypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Genotype{}
	o := &Genotype{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, genotypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Genotype object: %s", err)
	}

	AddGenotypeHook(boil.BeforeInsertHook, genotypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	genotypeBeforeInsertHooks = []GenotypeHook{}

	AddGenotypeHook(boil.AfterInsertHook, genotypeAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	genotypeAfterInsertHooks = []GenotypeHook{}

	AddGenotypeHook(boil.AfterSelectHook, genotypeAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	genotypeAfterSelectHooks = []GenotypeHook{}

	AddGenotypeHook(boil.BeforeUpdateHook, genotypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	genotypeBeforeUpdateHooks = []GenotypeHook{}

	AddGenotypeHook(boil.AfterUpdateHook, genotypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	genotypeAfterUpdateHooks = []GenotypeHook{}

	AddGenotypeHook(boil.BeforeDeleteHook, genotypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	genotypeBeforeDeleteHooks = []GenotypeHook{}

	AddGenotypeHook(boil.AfterDeleteHook, genotypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	genotypeAfterDeleteHooks = []GenotypeHook{}

	AddGenotypeHook(boil.BeforeUpsertHook, genotypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	genotypeBeforeUpsertHooks = []GenotypeHook{}

	AddGenotypeHook(boil.AfterUpsertHook, genotypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	genotypeAfterUpsertHooks = []GenotypeHook{}
}
func testGenotypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotype := &Genotype{}
	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotype.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Genotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGenotypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotype := &Genotype{}
	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotype.Insert(tx, genotypeColumns...); err != nil {
		t.Error(err)
	}

	count, err := Genotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGenotypeOneToOneFeatureGenotypeUsingFeatureGenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureGenotype
	var local Genotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.GenotypeID = local.GenotypeID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureGenotype(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.GenotypeID != foreign.GenotypeID {
		t.Errorf("want: %v, got %v", foreign.GenotypeID, check.GenotypeID)
	}

	slice := GenotypeSlice{&local}
	if err = local.L.LoadFeatureGenotype(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureGenotype == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureGenotype = nil
	if err = local.L.LoadFeatureGenotype(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureGenotype == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testGenotypeOneToOneGenotypepropUsingGenotypeprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Genotypeprop
	var local Genotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, genotypepropDBTypes, true, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.GenotypeID = local.GenotypeID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Genotypeprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.GenotypeID != foreign.GenotypeID {
		t.Errorf("want: %v, got %v", foreign.GenotypeID, check.GenotypeID)
	}

	slice := GenotypeSlice{&local}
	if err = local.L.LoadGenotypeprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Genotypeprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Genotypeprop = nil
	if err = local.L.LoadGenotypeprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Genotypeprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testGenotypeOneToOnePhendescUsingPhendesc(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Phendesc
	var local Genotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.GenotypeID = local.GenotypeID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Phendesc(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.GenotypeID != foreign.GenotypeID {
		t.Errorf("want: %v, got %v", foreign.GenotypeID, check.GenotypeID)
	}

	slice := GenotypeSlice{&local}
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

func testGenotypeOneToOnePhenstatementUsingPhenstatement(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Phenstatement
	var local Genotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.GenotypeID = local.GenotypeID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Phenstatement(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.GenotypeID != foreign.GenotypeID {
		t.Errorf("want: %v, got %v", foreign.GenotypeID, check.GenotypeID)
	}

	slice := GenotypeSlice{&local}
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

func testGenotypeOneToOneStockGenotypeUsingStockGenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockGenotype
	var local Genotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockGenotypeDBTypes, true, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.GenotypeID = local.GenotypeID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockGenotype(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.GenotypeID != foreign.GenotypeID {
		t.Errorf("want: %v, got %v", foreign.GenotypeID, check.GenotypeID)
	}

	slice := GenotypeSlice{&local}
	if err = local.L.LoadStockGenotype(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockGenotype == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockGenotype = nil
	if err = local.L.LoadStockGenotype(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockGenotype == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testGenotypeOneToOnePhenotypeComparisonUsingGenotype1PhenotypeComparison(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign PhenotypeComparison
	var local Genotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.Genotype1ID = local.GenotypeID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Genotype1PhenotypeComparison(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.Genotype1ID != foreign.Genotype1ID {
		t.Errorf("want: %v, got %v", foreign.Genotype1ID, check.Genotype1ID)
	}

	slice := GenotypeSlice{&local}
	if err = local.L.LoadGenotype1PhenotypeComparison(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Genotype1PhenotypeComparison == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Genotype1PhenotypeComparison = nil
	if err = local.L.LoadGenotype1PhenotypeComparison(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Genotype1PhenotypeComparison == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testGenotypeOneToOnePhenotypeComparisonUsingGenotype2PhenotypeComparison(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign PhenotypeComparison
	var local Genotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.Genotype2ID = local.GenotypeID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Genotype2PhenotypeComparison(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.Genotype2ID != foreign.Genotype2ID {
		t.Errorf("want: %v, got %v", foreign.Genotype2ID, check.Genotype2ID)
	}

	slice := GenotypeSlice{&local}
	if err = local.L.LoadGenotype2PhenotypeComparison(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Genotype2PhenotypeComparison == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Genotype2PhenotypeComparison = nil
	if err = local.L.LoadGenotype2PhenotypeComparison(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Genotype2PhenotypeComparison == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testGenotypeOneToOneSetOpFeatureGenotypeUsingFeatureGenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Genotype
	var b, c FeatureGenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, genotypeDBTypes, false, strmangle.SetComplement(genotypePrimaryKeyColumns, genotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureGenotypeDBTypes, false, strmangle.SetComplement(featureGenotypePrimaryKeyColumns, featureGenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureGenotypeDBTypes, false, strmangle.SetComplement(featureGenotypePrimaryKeyColumns, featureGenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureGenotype{&b, &c} {
		err = a.SetFeatureGenotype(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureGenotype != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Genotype != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.GenotypeID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.GenotypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.GenotypeID))
		reflect.Indirect(reflect.ValueOf(&x.GenotypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GenotypeID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.GenotypeID, x.GenotypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testGenotypeOneToOneSetOpGenotypepropUsingGenotypeprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Genotype
	var b, c Genotypeprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, genotypeDBTypes, false, strmangle.SetComplement(genotypePrimaryKeyColumns, genotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, genotypepropDBTypes, false, strmangle.SetComplement(genotypepropPrimaryKeyColumns, genotypepropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, genotypepropDBTypes, false, strmangle.SetComplement(genotypepropPrimaryKeyColumns, genotypepropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Genotypeprop{&b, &c} {
		err = a.SetGenotypeprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Genotypeprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Genotype != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.GenotypeID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.GenotypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.GenotypeID))
		reflect.Indirect(reflect.ValueOf(&x.GenotypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GenotypeID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.GenotypeID, x.GenotypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testGenotypeOneToOneSetOpPhendescUsingPhendesc(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Genotype
	var b, c Phendesc

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, genotypeDBTypes, false, strmangle.SetComplement(genotypePrimaryKeyColumns, genotypeColumnsWithoutDefault)...); err != nil {
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
		if x.R.Genotype != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.GenotypeID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.GenotypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.GenotypeID))
		reflect.Indirect(reflect.ValueOf(&x.GenotypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GenotypeID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.GenotypeID, x.GenotypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testGenotypeOneToOneSetOpPhenstatementUsingPhenstatement(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Genotype
	var b, c Phenstatement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, genotypeDBTypes, false, strmangle.SetComplement(genotypePrimaryKeyColumns, genotypeColumnsWithoutDefault)...); err != nil {
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
		if x.R.Genotype != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.GenotypeID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.GenotypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.GenotypeID))
		reflect.Indirect(reflect.ValueOf(&x.GenotypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GenotypeID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.GenotypeID, x.GenotypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testGenotypeOneToOneSetOpStockGenotypeUsingStockGenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Genotype
	var b, c StockGenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, genotypeDBTypes, false, strmangle.SetComplement(genotypePrimaryKeyColumns, genotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockGenotypeDBTypes, false, strmangle.SetComplement(stockGenotypePrimaryKeyColumns, stockGenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockGenotypeDBTypes, false, strmangle.SetComplement(stockGenotypePrimaryKeyColumns, stockGenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockGenotype{&b, &c} {
		err = a.SetStockGenotype(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockGenotype != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Genotype != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.GenotypeID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.GenotypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.GenotypeID))
		reflect.Indirect(reflect.ValueOf(&x.GenotypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GenotypeID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.GenotypeID, x.GenotypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testGenotypeOneToOneSetOpPhenotypeComparisonUsingGenotype1PhenotypeComparison(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Genotype
	var b, c PhenotypeComparison

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, genotypeDBTypes, false, strmangle.SetComplement(genotypePrimaryKeyColumns, genotypeColumnsWithoutDefault)...); err != nil {
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
		err = a.SetGenotype1PhenotypeComparison(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Genotype1PhenotypeComparison != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Genotype1 != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.GenotypeID != x.Genotype1ID {
			t.Error("foreign key was wrong value", a.GenotypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.Genotype1ID))
		reflect.Indirect(reflect.ValueOf(&x.Genotype1ID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GenotypeID != x.Genotype1ID {
			t.Error("foreign key was wrong value", a.GenotypeID, x.Genotype1ID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testGenotypeOneToOneSetOpPhenotypeComparisonUsingGenotype2PhenotypeComparison(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Genotype
	var b, c PhenotypeComparison

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, genotypeDBTypes, false, strmangle.SetComplement(genotypePrimaryKeyColumns, genotypeColumnsWithoutDefault)...); err != nil {
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
		err = a.SetGenotype2PhenotypeComparison(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Genotype2PhenotypeComparison != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Genotype2 != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.GenotypeID != x.Genotype2ID {
			t.Error("foreign key was wrong value", a.GenotypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.Genotype2ID))
		reflect.Indirect(reflect.ValueOf(&x.Genotype2ID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GenotypeID != x.Genotype2ID {
			t.Error("foreign key was wrong value", a.GenotypeID, x.Genotype2ID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testGenotypeToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Genotype
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
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

	slice := GenotypeSlice{&local}
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

func testGenotypeToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Genotype
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, genotypeDBTypes, false, strmangle.SetComplement(genotypePrimaryKeyColumns, genotypeColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeGenotypes[0] != &a {
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
func testGenotypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotype := &Genotype{}
	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = genotype.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testGenotypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotype := &Genotype{}
	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotype.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := GenotypeSlice{genotype}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testGenotypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotype := &Genotype{}
	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotype.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Genotypes(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	genotypeDBTypes = map[string]string{"Description": "character varying", "GenotypeID": "integer", "Name": "text", "TypeID": "integer", "Uniquename": "text"}
	_               = bytes.MinRead
)

func testGenotypesUpdate(t *testing.T) {
	t.Parallel()

	if len(genotypeColumns) == len(genotypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	genotype := &Genotype{}
	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotype.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Genotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	if err = genotype.Update(tx); err != nil {
		t.Error(err)
	}
}

func testGenotypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(genotypeColumns) == len(genotypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	genotype := &Genotype{}
	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotype.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Genotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, genotype, genotypeDBTypes, true, genotypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(genotypeColumns, genotypePrimaryKeyColumns) {
		fields = genotypeColumns
	} else {
		fields = strmangle.SetComplement(
			genotypeColumns,
			genotypePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(genotype))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := GenotypeSlice{genotype}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testGenotypesUpsert(t *testing.T) {
	t.Parallel()

	if len(genotypeColumns) == len(genotypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	genotype := Genotype{}
	if err = randomize.Struct(seed, &genotype, genotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotype.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Genotype: %s", err)
	}

	count, err := Genotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &genotype, genotypeDBTypes, false, genotypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	if err = genotype.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Genotype: %s", err)
	}

	count, err = Genotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
