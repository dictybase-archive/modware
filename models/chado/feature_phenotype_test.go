package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeaturePhenotypes(t *testing.T) {
	t.Parallel()

	query := FeaturePhenotypes(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeaturePhenotypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePhenotype := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featurePhenotype.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePhenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturePhenotypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePhenotype := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeaturePhenotypes(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeaturePhenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturePhenotypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePhenotype := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeaturePhenotypeSlice{featurePhenotype}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePhenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeaturePhenotypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePhenotype := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeaturePhenotypeExists(tx, featurePhenotype.FeaturePhenotypeID)
	if err != nil {
		t.Errorf("Unable to check if FeaturePhenotype exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeaturePhenotypeExistsG to return true, but got false.")
	}
}
func testFeaturePhenotypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePhenotype := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	featurePhenotypeFound, err := FindFeaturePhenotype(tx, featurePhenotype.FeaturePhenotypeID)
	if err != nil {
		t.Error(err)
	}

	if featurePhenotypeFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeaturePhenotypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePhenotype := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeaturePhenotypes(tx).Bind(featurePhenotype); err != nil {
		t.Error(err)
	}
}

func testFeaturePhenotypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePhenotype := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeaturePhenotypes(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeaturePhenotypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePhenotypeOne := &FeaturePhenotype{}
	featurePhenotypeTwo := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotypeOne, featurePhenotypeDBTypes, false, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}
	if err = randomize.Struct(seed, featurePhenotypeTwo, featurePhenotypeDBTypes, false, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotypeOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featurePhenotypeTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeaturePhenotypes(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeaturePhenotypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featurePhenotypeOne := &FeaturePhenotype{}
	featurePhenotypeTwo := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotypeOne, featurePhenotypeDBTypes, false, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}
	if err = randomize.Struct(seed, featurePhenotypeTwo, featurePhenotypeDBTypes, false, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotypeOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featurePhenotypeTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePhenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featurePhenotypeBeforeInsertHook(e boil.Executor, o *FeaturePhenotype) error {
	*o = FeaturePhenotype{}
	return nil
}

func featurePhenotypeAfterInsertHook(e boil.Executor, o *FeaturePhenotype) error {
	*o = FeaturePhenotype{}
	return nil
}

func featurePhenotypeAfterSelectHook(e boil.Executor, o *FeaturePhenotype) error {
	*o = FeaturePhenotype{}
	return nil
}

func featurePhenotypeBeforeUpdateHook(e boil.Executor, o *FeaturePhenotype) error {
	*o = FeaturePhenotype{}
	return nil
}

func featurePhenotypeAfterUpdateHook(e boil.Executor, o *FeaturePhenotype) error {
	*o = FeaturePhenotype{}
	return nil
}

func featurePhenotypeBeforeDeleteHook(e boil.Executor, o *FeaturePhenotype) error {
	*o = FeaturePhenotype{}
	return nil
}

func featurePhenotypeAfterDeleteHook(e boil.Executor, o *FeaturePhenotype) error {
	*o = FeaturePhenotype{}
	return nil
}

func featurePhenotypeBeforeUpsertHook(e boil.Executor, o *FeaturePhenotype) error {
	*o = FeaturePhenotype{}
	return nil
}

func featurePhenotypeAfterUpsertHook(e boil.Executor, o *FeaturePhenotype) error {
	*o = FeaturePhenotype{}
	return nil
}

func testFeaturePhenotypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeaturePhenotype{}
	o := &FeaturePhenotype{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featurePhenotypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype object: %s", err)
	}

	AddFeaturePhenotypeHook(boil.BeforeInsertHook, featurePhenotypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featurePhenotypeBeforeInsertHooks = []FeaturePhenotypeHook{}

	AddFeaturePhenotypeHook(boil.AfterInsertHook, featurePhenotypeAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featurePhenotypeAfterInsertHooks = []FeaturePhenotypeHook{}

	AddFeaturePhenotypeHook(boil.AfterSelectHook, featurePhenotypeAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featurePhenotypeAfterSelectHooks = []FeaturePhenotypeHook{}

	AddFeaturePhenotypeHook(boil.BeforeUpdateHook, featurePhenotypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featurePhenotypeBeforeUpdateHooks = []FeaturePhenotypeHook{}

	AddFeaturePhenotypeHook(boil.AfterUpdateHook, featurePhenotypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featurePhenotypeAfterUpdateHooks = []FeaturePhenotypeHook{}

	AddFeaturePhenotypeHook(boil.BeforeDeleteHook, featurePhenotypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featurePhenotypeBeforeDeleteHooks = []FeaturePhenotypeHook{}

	AddFeaturePhenotypeHook(boil.AfterDeleteHook, featurePhenotypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featurePhenotypeAfterDeleteHooks = []FeaturePhenotypeHook{}

	AddFeaturePhenotypeHook(boil.BeforeUpsertHook, featurePhenotypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featurePhenotypeBeforeUpsertHooks = []FeaturePhenotypeHook{}

	AddFeaturePhenotypeHook(boil.AfterUpsertHook, featurePhenotypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featurePhenotypeAfterUpsertHooks = []FeaturePhenotypeHook{}
}
func testFeaturePhenotypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePhenotype := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePhenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeaturePhenotypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePhenotype := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotype.Insert(tx, featurePhenotypeColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeaturePhenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeaturePhenotypeToOnePhenotypeUsingPhenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeaturePhenotype
	var foreign Phenotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featurePhenotypeDBTypes, true, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
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

	slice := FeaturePhenotypeSlice{&local}
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

func testFeaturePhenotypeToOneFeatureUsingFeature(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeaturePhenotype
	var foreign Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featurePhenotypeDBTypes, true, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
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

	slice := FeaturePhenotypeSlice{&local}
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

func testFeaturePhenotypeToOneSetOpPhenotypeUsingPhenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeaturePhenotype
	var b, c Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurePhenotypeDBTypes, false, strmangle.SetComplement(featurePhenotypePrimaryKeyColumns, featurePhenotypeColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeaturePhenotype != &a {
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
func testFeaturePhenotypeToOneSetOpFeatureUsingFeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeaturePhenotype
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurePhenotypeDBTypes, false, strmangle.SetComplement(featurePhenotypePrimaryKeyColumns, featurePhenotypeColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeaturePhenotype != &a {
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
func testFeaturePhenotypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePhenotype := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featurePhenotype.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturePhenotypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePhenotype := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeaturePhenotypeSlice{featurePhenotype}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeaturePhenotypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurePhenotype := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeaturePhenotypes(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featurePhenotypeDBTypes = map[string]string{"FeatureID": "integer", "FeaturePhenotypeID": "integer", "PhenotypeID": "integer"}
	_                       = bytes.MinRead
)

func testFeaturePhenotypesUpdate(t *testing.T) {
	t.Parallel()

	if len(featurePhenotypeColumns) == len(featurePhenotypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featurePhenotype := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePhenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true, featurePhenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	if err = featurePhenotype.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturePhenotypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featurePhenotypeColumns) == len(featurePhenotypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featurePhenotype := &FeaturePhenotype{}
	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturePhenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featurePhenotype, featurePhenotypeDBTypes, true, featurePhenotypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featurePhenotypeColumns, featurePhenotypePrimaryKeyColumns) {
		fields = featurePhenotypeColumns
	} else {
		fields = strmangle.SetComplement(
			featurePhenotypeColumns,
			featurePhenotypePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featurePhenotype))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeaturePhenotypeSlice{featurePhenotype}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeaturePhenotypesUpsert(t *testing.T) {
	t.Parallel()

	if len(featurePhenotypeColumns) == len(featurePhenotypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featurePhenotype := FeaturePhenotype{}
	if err = randomize.Struct(seed, &featurePhenotype, featurePhenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurePhenotype.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeaturePhenotype: %s", err)
	}

	count, err := FeaturePhenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featurePhenotype, featurePhenotypeDBTypes, false, featurePhenotypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeaturePhenotype struct: %s", err)
	}

	if err = featurePhenotype.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeaturePhenotype: %s", err)
	}

	count, err = FeaturePhenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
