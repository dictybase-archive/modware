package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeatures(t *testing.T) {
	t.Parallel()

	query := Features(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeaturesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	feature := &Feature{}
	if err = randomize.Struct(seed, feature, featureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = feature.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = feature.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Features(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	feature := &Feature{}
	if err = randomize.Struct(seed, feature, featureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = feature.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Features(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Features(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	feature := &Feature{}
	if err = randomize.Struct(seed, feature, featureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = feature.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureSlice{feature}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Features(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeaturesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	feature := &Feature{}
	if err = randomize.Struct(seed, feature, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = feature.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeatureExists(tx, feature.FeatureID)
	if err != nil {
		t.Errorf("Unable to check if Feature exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeatureExistsG to return true, but got false.")
	}
}
func testFeaturesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	feature := &Feature{}
	if err = randomize.Struct(seed, feature, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = feature.Insert(tx); err != nil {
		t.Error(err)
	}

	featureFound, err := FindFeature(tx, feature.FeatureID)
	if err != nil {
		t.Error(err)
	}

	if featureFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeaturesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	feature := &Feature{}
	if err = randomize.Struct(seed, feature, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = feature.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Features(tx).Bind(feature); err != nil {
		t.Error(err)
	}
}

func testFeaturesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	feature := &Feature{}
	if err = randomize.Struct(seed, feature, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = feature.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Features(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeaturesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureOne := &Feature{}
	featureTwo := &Feature{}
	if err = randomize.Struct(seed, featureOne, featureDBTypes, false, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}
	if err = randomize.Struct(seed, featureTwo, featureDBTypes, false, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Features(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeaturesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featureOne := &Feature{}
	featureTwo := &Feature{}
	if err = randomize.Struct(seed, featureOne, featureDBTypes, false, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}
	if err = randomize.Struct(seed, featureTwo, featureDBTypes, false, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Features(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featureBeforeInsertHook(e boil.Executor, o *Feature) error {
	*o = Feature{}
	return nil
}

func featureAfterInsertHook(e boil.Executor, o *Feature) error {
	*o = Feature{}
	return nil
}

func featureAfterSelectHook(e boil.Executor, o *Feature) error {
	*o = Feature{}
	return nil
}

func featureBeforeUpdateHook(e boil.Executor, o *Feature) error {
	*o = Feature{}
	return nil
}

func featureAfterUpdateHook(e boil.Executor, o *Feature) error {
	*o = Feature{}
	return nil
}

func featureBeforeDeleteHook(e boil.Executor, o *Feature) error {
	*o = Feature{}
	return nil
}

func featureAfterDeleteHook(e boil.Executor, o *Feature) error {
	*o = Feature{}
	return nil
}

func featureBeforeUpsertHook(e boil.Executor, o *Feature) error {
	*o = Feature{}
	return nil
}

func featureAfterUpsertHook(e boil.Executor, o *Feature) error {
	*o = Feature{}
	return nil
}

func testFeaturesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Feature{}
	o := &Feature{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featureDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Feature object: %s", err)
	}

	AddFeatureHook(boil.BeforeInsertHook, featureBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featureBeforeInsertHooks = []FeatureHook{}

	AddFeatureHook(boil.AfterInsertHook, featureAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featureAfterInsertHooks = []FeatureHook{}

	AddFeatureHook(boil.AfterSelectHook, featureAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featureAfterSelectHooks = []FeatureHook{}

	AddFeatureHook(boil.BeforeUpdateHook, featureBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featureBeforeUpdateHooks = []FeatureHook{}

	AddFeatureHook(boil.AfterUpdateHook, featureAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featureAfterUpdateHooks = []FeatureHook{}

	AddFeatureHook(boil.BeforeDeleteHook, featureBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featureBeforeDeleteHooks = []FeatureHook{}

	AddFeatureHook(boil.AfterDeleteHook, featureAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featureAfterDeleteHooks = []FeatureHook{}

	AddFeatureHook(boil.BeforeUpsertHook, featureBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featureBeforeUpsertHooks = []FeatureHook{}

	AddFeatureHook(boil.AfterUpsertHook, featureAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featureAfterUpsertHooks = []FeatureHook{}
}
func testFeaturesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	feature := &Feature{}
	if err = randomize.Struct(seed, feature, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = feature.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Features(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeaturesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	feature := &Feature{}
	if err = randomize.Struct(seed, feature, featureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = feature.Insert(tx, featureColumns...); err != nil {
		t.Error(err)
	}

	count, err := Features(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureOneToOneFeatureDbxrefUsingFeatureDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureDbxref
	var local Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureDbxrefDBTypes, true, featureDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureDbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeatureID = local.FeatureID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureDbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureID != foreign.FeatureID {
		t.Errorf("want: %v, got %v", foreign.FeatureID, check.FeatureID)
	}

	slice := FeatureSlice{&local}
	if err = local.L.LoadFeatureDbxref(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureDbxref == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureDbxref = nil
	if err = local.L.LoadFeatureDbxref(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureDbxref == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureOneToOneFeaturePhenotypeUsingFeaturePhenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeaturePhenotype
	var local Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featurePhenotypeDBTypes, true, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeatureID = local.FeatureID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeaturePhenotype(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureID != foreign.FeatureID {
		t.Errorf("want: %v, got %v", foreign.FeatureID, check.FeatureID)
	}

	slice := FeatureSlice{&local}
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

func testFeatureOneToOneAnalysisfeatureUsingAnalysisfeature(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Analysisfeature
	var local Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, analysisfeatureDBTypes, true, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeatureID = local.FeatureID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Analysisfeature(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureID != foreign.FeatureID {
		t.Errorf("want: %v, got %v", foreign.FeatureID, check.FeatureID)
	}

	slice := FeatureSlice{&local}
	if err = local.L.LoadAnalysisfeature(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Analysisfeature == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Analysisfeature = nil
	if err = local.L.LoadAnalysisfeature(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Analysisfeature == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureOneToOneFeaturepropUsingFeatureprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Featureprop
	var local Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featurepropDBTypes, true, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeatureID = local.FeatureID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Featureprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureID != foreign.FeatureID {
		t.Errorf("want: %v, got %v", foreign.FeatureID, check.FeatureID)
	}

	slice := FeatureSlice{&local}
	if err = local.L.LoadFeatureprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Featureprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Featureprop = nil
	if err = local.L.LoadFeatureprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Featureprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureOneToOneFeatureCvtermUsingFeatureCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureCvterm
	var local Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeatureID = local.FeatureID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureCvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureID != foreign.FeatureID {
		t.Errorf("want: %v, got %v", foreign.FeatureID, check.FeatureID)
	}

	slice := FeatureSlice{&local}
	if err = local.L.LoadFeatureCvterm(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureCvterm == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureCvterm = nil
	if err = local.L.LoadFeatureCvterm(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureCvterm == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureOneToOneFeatureGenotypeUsingChromosomeFeatureGenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureGenotype
	var local Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	foreign.ChromosomeID.Valid = true

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.ChromosomeID.Int = local.FeatureID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ChromosomeFeatureGenotype(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ChromosomeID.Int != foreign.ChromosomeID.Int {
		t.Errorf("want: %v, got %v", foreign.ChromosomeID.Int, check.ChromosomeID.Int)
	}

	slice := FeatureSlice{&local}
	if err = local.L.LoadChromosomeFeatureGenotype(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.ChromosomeFeatureGenotype == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ChromosomeFeatureGenotype = nil
	if err = local.L.LoadChromosomeFeatureGenotype(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.ChromosomeFeatureGenotype == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureOneToOneFeatureGenotypeUsingFeatureGenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureGenotype
	var local Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeatureID = local.FeatureID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureGenotype(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureID != foreign.FeatureID {
		t.Errorf("want: %v, got %v", foreign.FeatureID, check.FeatureID)
	}

	slice := FeatureSlice{&local}
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

func testFeatureOneToOneFeaturePubUsingFeaturePub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeaturePub
	var local Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featurePubDBTypes, true, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeatureID = local.FeatureID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeaturePub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureID != foreign.FeatureID {
		t.Errorf("want: %v, got %v", foreign.FeatureID, check.FeatureID)
	}

	slice := FeatureSlice{&local}
	if err = local.L.LoadFeaturePub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturePub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeaturePub = nil
	if err = local.L.LoadFeaturePub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturePub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureOneToOneFeatureSynonymUsingFeatureSynonym(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureSynonym
	var local Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureSynonymDBTypes, true, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeatureID = local.FeatureID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureSynonym(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureID != foreign.FeatureID {
		t.Errorf("want: %v, got %v", foreign.FeatureID, check.FeatureID)
	}

	slice := FeatureSlice{&local}
	if err = local.L.LoadFeatureSynonym(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureSynonym == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureSynonym = nil
	if err = local.L.LoadFeatureSynonym(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureSynonym == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureOneToOneFeatureRelationshipUsingObjectFeatureRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureRelationship
	var local Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.ObjectID = local.FeatureID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ObjectFeatureRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ObjectID != foreign.ObjectID {
		t.Errorf("want: %v, got %v", foreign.ObjectID, check.ObjectID)
	}

	slice := FeatureSlice{&local}
	if err = local.L.LoadObjectFeatureRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.ObjectFeatureRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ObjectFeatureRelationship = nil
	if err = local.L.LoadObjectFeatureRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.ObjectFeatureRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureOneToOneFeatureRelationshipUsingSubjectFeatureRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureRelationship
	var local Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.SubjectID = local.FeatureID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.SubjectFeatureRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.SubjectID != foreign.SubjectID {
		t.Errorf("want: %v, got %v", foreign.SubjectID, check.SubjectID)
	}

	slice := FeatureSlice{&local}
	if err = local.L.LoadSubjectFeatureRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.SubjectFeatureRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SubjectFeatureRelationship = nil
	if err = local.L.LoadSubjectFeatureRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.SubjectFeatureRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureOneToOneFeaturelocUsingFeatureloc(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Featureloc
	var local Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featurelocDBTypes, true, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.FeatureID = local.FeatureID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Featureloc(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeatureID != foreign.FeatureID {
		t.Errorf("want: %v, got %v", foreign.FeatureID, check.FeatureID)
	}

	slice := FeatureSlice{&local}
	if err = local.L.LoadFeatureloc(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Featureloc == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Featureloc = nil
	if err = local.L.LoadFeatureloc(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Featureloc == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureOneToOneSetOpFeatureDbxrefUsingFeatureDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c FeatureDbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureDbxrefDBTypes, false, strmangle.SetComplement(featureDbxrefPrimaryKeyColumns, featureDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureDbxrefDBTypes, false, strmangle.SetComplement(featureDbxrefPrimaryKeyColumns, featureDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureDbxref{&b, &c} {
		err = a.SetFeatureDbxref(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureDbxref != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Feature != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeatureID))
		reflect.Indirect(reflect.ValueOf(&x.FeatureID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID, x.FeatureID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testFeatureOneToOneSetOpFeaturePhenotypeUsingFeaturePhenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c FeaturePhenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
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
		if x.R.Feature != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeatureID))
		reflect.Indirect(reflect.ValueOf(&x.FeatureID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID, x.FeatureID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testFeatureOneToOneSetOpAnalysisfeatureUsingAnalysisfeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c Analysisfeature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, analysisfeatureDBTypes, false, strmangle.SetComplement(analysisfeaturePrimaryKeyColumns, analysisfeatureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, analysisfeatureDBTypes, false, strmangle.SetComplement(analysisfeaturePrimaryKeyColumns, analysisfeatureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Analysisfeature{&b, &c} {
		err = a.SetAnalysisfeature(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Analysisfeature != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Feature != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeatureID))
		reflect.Indirect(reflect.ValueOf(&x.FeatureID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID, x.FeatureID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testFeatureOneToOneSetOpFeaturepropUsingFeatureprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c Featureprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featurepropDBTypes, false, strmangle.SetComplement(featurepropPrimaryKeyColumns, featurepropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featurepropDBTypes, false, strmangle.SetComplement(featurepropPrimaryKeyColumns, featurepropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Featureprop{&b, &c} {
		err = a.SetFeatureprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Featureprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Feature != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeatureID))
		reflect.Indirect(reflect.ValueOf(&x.FeatureID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID, x.FeatureID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testFeatureOneToOneSetOpFeatureCvtermUsingFeatureCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c FeatureCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureCvtermDBTypes, false, strmangle.SetComplement(featureCvtermPrimaryKeyColumns, featureCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureCvtermDBTypes, false, strmangle.SetComplement(featureCvtermPrimaryKeyColumns, featureCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureCvterm{&b, &c} {
		err = a.SetFeatureCvterm(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureCvterm != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Feature != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeatureID))
		reflect.Indirect(reflect.ValueOf(&x.FeatureID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID, x.FeatureID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testFeatureOneToOneSetOpFeatureGenotypeUsingChromosomeFeatureGenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c FeatureGenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
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
		err = a.SetChromosomeFeatureGenotype(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ChromosomeFeatureGenotype != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Chromosome != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureID != x.ChromosomeID.Int {
			t.Error("foreign key was wrong value", a.FeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.ChromosomeID.Int))
		reflect.Indirect(reflect.ValueOf(&x.ChromosomeID.Int)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureID != x.ChromosomeID.Int {
			t.Error("foreign key was wrong value", a.FeatureID, x.ChromosomeID.Int)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testFeatureOneToOneRemoveOpFeatureGenotypeUsingChromosomeFeatureGenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b FeatureGenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureGenotypeDBTypes, false, strmangle.SetComplement(featureGenotypePrimaryKeyColumns, featureGenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetChromosomeFeatureGenotype(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveChromosomeFeatureGenotype(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.ChromosomeFeatureGenotype(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.ChromosomeFeatureGenotype != nil {
		t.Error("R struct entry should be nil")
	}

	if b.ChromosomeID.Valid {
		t.Error("foreign key column should be nil")
	}

	if b.R.Chromosome != nil {
		t.Error("failed to remove a from b's relationships")
	}
}

func testFeatureOneToOneSetOpFeatureGenotypeUsingFeatureGenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c FeatureGenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
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
		if x.R.Feature != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeatureID))
		reflect.Indirect(reflect.ValueOf(&x.FeatureID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID, x.FeatureID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testFeatureOneToOneSetOpFeaturePubUsingFeaturePub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c FeaturePub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featurePubDBTypes, false, strmangle.SetComplement(featurePubPrimaryKeyColumns, featurePubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featurePubDBTypes, false, strmangle.SetComplement(featurePubPrimaryKeyColumns, featurePubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeaturePub{&b, &c} {
		err = a.SetFeaturePub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeaturePub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Feature != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeatureID))
		reflect.Indirect(reflect.ValueOf(&x.FeatureID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID, x.FeatureID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testFeatureOneToOneSetOpFeatureSynonymUsingFeatureSynonym(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c FeatureSynonym

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureSynonymDBTypes, false, strmangle.SetComplement(featureSynonymPrimaryKeyColumns, featureSynonymColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureSynonymDBTypes, false, strmangle.SetComplement(featureSynonymPrimaryKeyColumns, featureSynonymColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureSynonym{&b, &c} {
		err = a.SetFeatureSynonym(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureSynonym != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Feature != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeatureID))
		reflect.Indirect(reflect.ValueOf(&x.FeatureID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID, x.FeatureID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testFeatureOneToOneSetOpFeatureRelationshipUsingObjectFeatureRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c FeatureRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureRelationshipDBTypes, false, strmangle.SetComplement(featureRelationshipPrimaryKeyColumns, featureRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureRelationshipDBTypes, false, strmangle.SetComplement(featureRelationshipPrimaryKeyColumns, featureRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureRelationship{&b, &c} {
		err = a.SetObjectFeatureRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ObjectFeatureRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Object != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureID != x.ObjectID {
			t.Error("foreign key was wrong value", a.FeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.ObjectID))
		reflect.Indirect(reflect.ValueOf(&x.ObjectID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureID != x.ObjectID {
			t.Error("foreign key was wrong value", a.FeatureID, x.ObjectID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testFeatureOneToOneSetOpFeatureRelationshipUsingSubjectFeatureRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c FeatureRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureRelationshipDBTypes, false, strmangle.SetComplement(featureRelationshipPrimaryKeyColumns, featureRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureRelationshipDBTypes, false, strmangle.SetComplement(featureRelationshipPrimaryKeyColumns, featureRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureRelationship{&b, &c} {
		err = a.SetSubjectFeatureRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SubjectFeatureRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Subject != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureID != x.SubjectID {
			t.Error("foreign key was wrong value", a.FeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.SubjectID))
		reflect.Indirect(reflect.ValueOf(&x.SubjectID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureID != x.SubjectID {
			t.Error("foreign key was wrong value", a.FeatureID, x.SubjectID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testFeatureOneToOneSetOpFeaturelocUsingFeatureloc(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c Featureloc

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featurelocDBTypes, false, strmangle.SetComplement(featurelocPrimaryKeyColumns, featurelocColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featurelocDBTypes, false, strmangle.SetComplement(featurelocPrimaryKeyColumns, featurelocColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Featureloc{&b, &c} {
		err = a.SetFeatureloc(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Featureloc != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Feature != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.FeatureID))
		reflect.Indirect(reflect.ValueOf(&x.FeatureID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeatureID != x.FeatureID {
			t.Error("foreign key was wrong value", a.FeatureID, x.FeatureID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testFeatureToManySrcfeatureFeaturelocs(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c Featureloc

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, featurelocDBTypes, false, featurelocColumnsWithDefault...)
	randomize.Struct(seed, &c, featurelocDBTypes, false, featurelocColumnsWithDefault...)
	b.SrcfeatureID.Valid = true
	c.SrcfeatureID.Valid = true
	b.SrcfeatureID.Int = a.FeatureID
	c.SrcfeatureID.Int = a.FeatureID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	featureloc, err := a.SrcfeatureFeaturelocs(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range featureloc {
		if v.SrcfeatureID.Int == b.SrcfeatureID.Int {
			bFound = true
		}
		if v.SrcfeatureID.Int == c.SrcfeatureID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := FeatureSlice{&a}
	if err = a.L.LoadSrcfeatureFeaturelocs(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SrcfeatureFeaturelocs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.SrcfeatureFeaturelocs = nil
	if err = a.L.LoadSrcfeatureFeaturelocs(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SrcfeatureFeaturelocs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", featureloc)
	}
}

func testFeatureToManyAddOpSrcfeatureFeaturelocs(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c, d, e Featureloc

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Featureloc{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, featurelocDBTypes, false, strmangle.SetComplement(featurelocPrimaryKeyColumns, featurelocColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Featureloc{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSrcfeatureFeaturelocs(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.FeatureID != first.SrcfeatureID.Int {
			t.Error("foreign key was wrong value", a.FeatureID, first.SrcfeatureID.Int)
		}
		if a.FeatureID != second.SrcfeatureID.Int {
			t.Error("foreign key was wrong value", a.FeatureID, second.SrcfeatureID.Int)
		}

		if first.R.Srcfeature != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Srcfeature != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.SrcfeatureFeaturelocs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.SrcfeatureFeaturelocs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.SrcfeatureFeaturelocs(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testFeatureToManySetOpSrcfeatureFeaturelocs(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c, d, e Featureloc

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Featureloc{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, featurelocDBTypes, false, strmangle.SetComplement(featurelocPrimaryKeyColumns, featurelocColumnsWithoutDefault)...); err != nil {
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

	err = a.SetSrcfeatureFeaturelocs(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.SrcfeatureFeaturelocs(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetSrcfeatureFeaturelocs(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.SrcfeatureFeaturelocs(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.SrcfeatureID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.SrcfeatureID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.FeatureID != d.SrcfeatureID.Int {
		t.Error("foreign key was wrong value", a.FeatureID, d.SrcfeatureID.Int)
	}
	if a.FeatureID != e.SrcfeatureID.Int {
		t.Error("foreign key was wrong value", a.FeatureID, e.SrcfeatureID.Int)
	}

	if b.R.Srcfeature != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Srcfeature != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Srcfeature != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Srcfeature != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.SrcfeatureFeaturelocs[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.SrcfeatureFeaturelocs[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testFeatureToManyRemoveOpSrcfeatureFeaturelocs(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c, d, e Featureloc

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Featureloc{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, featurelocDBTypes, false, strmangle.SetComplement(featurelocPrimaryKeyColumns, featurelocColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddSrcfeatureFeaturelocs(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.SrcfeatureFeaturelocs(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveSrcfeatureFeaturelocs(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.SrcfeatureFeaturelocs(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.SrcfeatureID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.SrcfeatureID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Srcfeature != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Srcfeature != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Srcfeature != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Srcfeature != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.SrcfeatureFeaturelocs) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.SrcfeatureFeaturelocs[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.SrcfeatureFeaturelocs[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testFeatureToOneOrganismUsingOrganism(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Feature
	var foreign Organism

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
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

	slice := FeatureSlice{&local}
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

func testFeatureToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Feature
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
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

	slice := FeatureSlice{&local}
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

func testFeatureToOneDbxrefUsingDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Feature
	var foreign Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	local.DbxrefID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.DbxrefID.Int = foreign.DbxrefID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Dbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.DbxrefID != foreign.DbxrefID {
		t.Errorf("want: %v, got %v", foreign.DbxrefID, check.DbxrefID)
	}

	slice := FeatureSlice{&local}
	if err = local.L.LoadDbxref(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Dbxref == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Dbxref = nil
	if err = local.L.LoadDbxref(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Dbxref == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureToOneSetOpOrganismUsingOrganism(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c Organism

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
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

		if x.R.Feature != &a {
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
func testFeatureToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeFeature != &a {
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
func testFeatureToOneSetOpDbxrefUsingDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b, c Dbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Dbxref{&b, &c} {
		err = a.SetDbxref(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Dbxref != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Features[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DbxrefID.Int != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DbxrefID.Int))
		reflect.Indirect(reflect.ValueOf(&a.DbxrefID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DbxrefID.Int != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID.Int, x.DbxrefID)
		}
	}
}

func testFeatureToOneRemoveOpDbxrefUsingDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Feature
	var b Dbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dbxrefDBTypes, false, strmangle.SetComplement(dbxrefPrimaryKeyColumns, dbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetDbxref(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveDbxref(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Dbxref(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Dbxref != nil {
		t.Error("R struct entry should be nil")
	}

	if a.DbxrefID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Features) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testFeaturesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	feature := &Feature{}
	if err = randomize.Struct(seed, feature, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = feature.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = feature.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	feature := &Feature{}
	if err = randomize.Struct(seed, feature, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = feature.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureSlice{feature}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeaturesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	feature := &Feature{}
	if err = randomize.Struct(seed, feature, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = feature.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Features(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featureDBTypes = map[string]string{"DbxrefID": "integer", "FeatureID": "integer", "IsAnalysis": "boolean", "IsObsolete": "boolean", "Md5checksum": "character", "Name": "character varying", "OrganismID": "integer", "Residues": "text", "Seqlen": "integer", "Timeaccessioned": "timestamp without time zone", "Timelastmodified": "timestamp without time zone", "TypeID": "integer", "Uniquename": "text"}
	_              = bytes.MinRead
)

func testFeaturesUpdate(t *testing.T) {
	t.Parallel()

	if len(featureColumns) == len(featurePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	feature := &Feature{}
	if err = randomize.Struct(seed, feature, featureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = feature.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Features(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, feature, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err = feature.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featureColumns) == len(featurePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	feature := &Feature{}
	if err = randomize.Struct(seed, feature, featureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = feature.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Features(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, feature, featureDBTypes, true, featurePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featureColumns, featurePrimaryKeyColumns) {
		fields = featureColumns
	} else {
		fields = strmangle.SetComplement(
			featureColumns,
			featurePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(feature))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeatureSlice{feature}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeaturesUpsert(t *testing.T) {
	t.Parallel()

	if len(featureColumns) == len(featurePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	feature := Feature{}
	if err = randomize.Struct(seed, &feature, featureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = feature.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Feature: %s", err)
	}

	count, err := Features(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &feature, featureDBTypes, false, featurePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}

	if err = feature.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Feature: %s", err)
	}

	count, err = Features(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
