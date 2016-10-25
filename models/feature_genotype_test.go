package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeatureGenotypes(t *testing.T) {
	t.Parallel()

	query := FeatureGenotypes(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeatureGenotypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureGenotype := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureGenotype.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureGenotypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureGenotype := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureGenotypes(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeatureGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureGenotypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureGenotype := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureGenotypeSlice{featureGenotype}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeatureGenotypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureGenotype := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeatureGenotypeExists(tx, featureGenotype.FeatureGenotypeID)
	if err != nil {
		t.Errorf("Unable to check if FeatureGenotype exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeatureGenotypeExistsG to return true, but got false.")
	}
}
func testFeatureGenotypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureGenotype := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	featureGenotypeFound, err := FindFeatureGenotype(tx, featureGenotype.FeatureGenotypeID)
	if err != nil {
		t.Error(err)
	}

	if featureGenotypeFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeatureGenotypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureGenotype := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureGenotypes(tx).Bind(featureGenotype); err != nil {
		t.Error(err)
	}
}

func testFeatureGenotypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureGenotype := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeatureGenotypes(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeatureGenotypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureGenotypeOne := &FeatureGenotype{}
	featureGenotypeTwo := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotypeOne, featureGenotypeDBTypes, false, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}
	if err = randomize.Struct(seed, featureGenotypeTwo, featureGenotypeDBTypes, false, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotypeOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureGenotypeTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureGenotypes(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeatureGenotypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featureGenotypeOne := &FeatureGenotype{}
	featureGenotypeTwo := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotypeOne, featureGenotypeDBTypes, false, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}
	if err = randomize.Struct(seed, featureGenotypeTwo, featureGenotypeDBTypes, false, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotypeOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureGenotypeTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featureGenotypeBeforeInsertHook(e boil.Executor, o *FeatureGenotype) error {
	*o = FeatureGenotype{}
	return nil
}

func featureGenotypeAfterInsertHook(e boil.Executor, o *FeatureGenotype) error {
	*o = FeatureGenotype{}
	return nil
}

func featureGenotypeAfterSelectHook(e boil.Executor, o *FeatureGenotype) error {
	*o = FeatureGenotype{}
	return nil
}

func featureGenotypeBeforeUpdateHook(e boil.Executor, o *FeatureGenotype) error {
	*o = FeatureGenotype{}
	return nil
}

func featureGenotypeAfterUpdateHook(e boil.Executor, o *FeatureGenotype) error {
	*o = FeatureGenotype{}
	return nil
}

func featureGenotypeBeforeDeleteHook(e boil.Executor, o *FeatureGenotype) error {
	*o = FeatureGenotype{}
	return nil
}

func featureGenotypeAfterDeleteHook(e boil.Executor, o *FeatureGenotype) error {
	*o = FeatureGenotype{}
	return nil
}

func featureGenotypeBeforeUpsertHook(e boil.Executor, o *FeatureGenotype) error {
	*o = FeatureGenotype{}
	return nil
}

func featureGenotypeAfterUpsertHook(e boil.Executor, o *FeatureGenotype) error {
	*o = FeatureGenotype{}
	return nil
}

func testFeatureGenotypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeatureGenotype{}
	o := &FeatureGenotype{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featureGenotypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype object: %s", err)
	}

	AddFeatureGenotypeHook(boil.BeforeInsertHook, featureGenotypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featureGenotypeBeforeInsertHooks = []FeatureGenotypeHook{}

	AddFeatureGenotypeHook(boil.AfterInsertHook, featureGenotypeAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featureGenotypeAfterInsertHooks = []FeatureGenotypeHook{}

	AddFeatureGenotypeHook(boil.AfterSelectHook, featureGenotypeAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featureGenotypeAfterSelectHooks = []FeatureGenotypeHook{}

	AddFeatureGenotypeHook(boil.BeforeUpdateHook, featureGenotypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featureGenotypeBeforeUpdateHooks = []FeatureGenotypeHook{}

	AddFeatureGenotypeHook(boil.AfterUpdateHook, featureGenotypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featureGenotypeAfterUpdateHooks = []FeatureGenotypeHook{}

	AddFeatureGenotypeHook(boil.BeforeDeleteHook, featureGenotypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featureGenotypeBeforeDeleteHooks = []FeatureGenotypeHook{}

	AddFeatureGenotypeHook(boil.AfterDeleteHook, featureGenotypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featureGenotypeAfterDeleteHooks = []FeatureGenotypeHook{}

	AddFeatureGenotypeHook(boil.BeforeUpsertHook, featureGenotypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featureGenotypeBeforeUpsertHooks = []FeatureGenotypeHook{}

	AddFeatureGenotypeHook(boil.AfterUpsertHook, featureGenotypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featureGenotypeAfterUpsertHooks = []FeatureGenotypeHook{}
}
func testFeatureGenotypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureGenotype := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureGenotypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureGenotype := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotype.Insert(tx, featureGenotypeColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeatureGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureGenotypeToOneGenotypeUsingGenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureGenotype
	var foreign Genotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
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

	slice := FeatureGenotypeSlice{&local}
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

func testFeatureGenotypeToOneCvtermUsingCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureGenotype
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
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

	slice := FeatureGenotypeSlice{&local}
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

func testFeatureGenotypeToOneFeatureUsingChromosome(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureGenotype
	var foreign Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	local.ChromosomeID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ChromosomeID.Int = foreign.FeatureID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Chromosome(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureID != foreign.FeatureID {
		t.Errorf("want: %v, got %v", foreign.FeatureID, check.FeatureID)
	}

	slice := FeatureGenotypeSlice{&local}
	if err = local.L.LoadChromosome(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Chromosome == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Chromosome = nil
	if err = local.L.LoadChromosome(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Chromosome == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureGenotypeToOneFeatureUsingFeature(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureGenotype
	var foreign Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.FeatureID = foreign.FeatureID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Feature(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureID != foreign.FeatureID {
		t.Errorf("want: %v, got %v", foreign.FeatureID, check.FeatureID)
	}

	slice := FeatureGenotypeSlice{&local}
	if err = local.L.LoadFeature(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Feature == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Feature = nil
	if err = local.L.LoadFeature(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Feature == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureGenotypeToOneSetOpGenotypeUsingGenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureGenotype
	var b, c Genotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureGenotypeDBTypes, false, strmangle.SetComplement(featureGenotypePrimaryKeyColumns, featureGenotypeColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeatureGenotype != &a {
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
func testFeatureGenotypeToOneSetOpCvtermUsingCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureGenotype
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureGenotypeDBTypes, false, strmangle.SetComplement(featureGenotypePrimaryKeyColumns, featureGenotypeColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeatureGenotype != &a {
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
func testFeatureGenotypeToOneSetOpFeatureUsingChromosome(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureGenotype
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureGenotypeDBTypes, false, strmangle.SetComplement(featureGenotypePrimaryKeyColumns, featureGenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Feature{&b, &c} {
		err = a.SetChromosome(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Chromosome != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ChromosomeFeatureGenotype != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ChromosomeID.Int != x.FeatureID {
			t.Error("foreign key was wrong value", a.ChromosomeID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ChromosomeID.Int))
		reflect.Indirect(reflect.ValueOf(&a.ChromosomeID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ChromosomeID.Int != x.FeatureID {
			t.Error("foreign key was wrong value", a.ChromosomeID.Int, x.FeatureID)
		}
	}
}

func testFeatureGenotypeToOneRemoveOpFeatureUsingChromosome(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureGenotype
	var b Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureGenotypeDBTypes, false, strmangle.SetComplement(featureGenotypePrimaryKeyColumns, featureGenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetChromosome(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveChromosome(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Chromosome(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Chromosome != nil {
		t.Error("R struct entry should be nil")
	}

	if a.ChromosomeID.Valid {
		t.Error("foreign key value should be nil")
	}

	if b.R.ChromosomeFeatureGenotype != nil {
		t.Error("failed to remove a from b's relationships")
	}

}

func testFeatureGenotypeToOneSetOpFeatureUsingFeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureGenotype
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureGenotypeDBTypes, false, strmangle.SetComplement(featureGenotypePrimaryKeyColumns, featureGenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Feature{&b, &c} {
		err = a.SetFeature(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Feature != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FeatureGenotype != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FeatureID))
		reflect.Indirect(reflect.ValueOf(&a.FeatureID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID, x.FeatureID)
		}
	}
}
func testFeatureGenotypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureGenotype := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureGenotype.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureGenotypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureGenotype := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureGenotypeSlice{featureGenotype}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeatureGenotypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureGenotype := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureGenotypes(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featureGenotypeDBTypes = map[string]string{"Cgroup": "integer", "ChromosomeID": "integer", "CvtermID": "integer", "FeatureGenotypeID": "integer", "FeatureID": "integer", "GenotypeID": "integer", "Rank": "integer"}
	_                      = bytes.MinRead
)

func testFeatureGenotypesUpdate(t *testing.T) {
	t.Parallel()

	if len(featureGenotypeColumns) == len(featureGenotypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureGenotype := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	if err = featureGenotype.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureGenotypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featureGenotypeColumns) == len(featureGenotypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureGenotype := &FeatureGenotype{}
	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureGenotype, featureGenotypeDBTypes, true, featureGenotypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featureGenotypeColumns, featureGenotypePrimaryKeyColumns) {
		fields = featureGenotypeColumns
	} else {
		fields = strmangle.SetComplement(
			featureGenotypeColumns,
			featureGenotypePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featureGenotype))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeatureGenotypeSlice{featureGenotype}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeatureGenotypesUpsert(t *testing.T) {
	t.Parallel()

	if len(featureGenotypeColumns) == len(featureGenotypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featureGenotype := FeatureGenotype{}
	if err = randomize.Struct(seed, &featureGenotype, featureGenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureGenotype.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureGenotype: %s", err)
	}

	count, err := FeatureGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featureGenotype, featureGenotypeDBTypes, false, featureGenotypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}

	if err = featureGenotype.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureGenotype: %s", err)
	}

	count, err = FeatureGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
