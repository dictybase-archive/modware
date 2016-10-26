package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testAnalysisfeatures(t *testing.T) {
	t.Parallel()

	query := Analysisfeatures(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAnalysisfeaturesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeature := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeature.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = analysisfeature.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnalysisfeaturesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeature := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeature.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Analysisfeatures(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnalysisfeaturesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeature := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeature.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AnalysisfeatureSlice{analysisfeature}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAnalysisfeaturesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeature := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeature.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AnalysisfeatureExists(tx, analysisfeature.AnalysisfeatureID)
	if err != nil {
		t.Errorf("Unable to check if Analysisfeature exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AnalysisfeatureExistsG to return true, but got false.")
	}
}
func testAnalysisfeaturesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeature := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeature.Insert(tx); err != nil {
		t.Error(err)
	}

	analysisfeatureFound, err := FindAnalysisfeature(tx, analysisfeature.AnalysisfeatureID)
	if err != nil {
		t.Error(err)
	}

	if analysisfeatureFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAnalysisfeaturesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeature := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeature.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Analysisfeatures(tx).Bind(analysisfeature); err != nil {
		t.Error(err)
	}
}

func testAnalysisfeaturesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeature := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeature.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Analysisfeatures(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAnalysisfeaturesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeatureOne := &Analysisfeature{}
	analysisfeatureTwo := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeatureOne, analysisfeatureDBTypes, false, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}
	if err = randomize.Struct(seed, analysisfeatureTwo, analysisfeatureDBTypes, false, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = analysisfeatureTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Analysisfeatures(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAnalysisfeaturesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	analysisfeatureOne := &Analysisfeature{}
	analysisfeatureTwo := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeatureOne, analysisfeatureDBTypes, false, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}
	if err = randomize.Struct(seed, analysisfeatureTwo, analysisfeatureDBTypes, false, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeatureOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = analysisfeatureTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func analysisfeatureBeforeInsertHook(e boil.Executor, o *Analysisfeature) error {
	*o = Analysisfeature{}
	return nil
}

func analysisfeatureAfterInsertHook(e boil.Executor, o *Analysisfeature) error {
	*o = Analysisfeature{}
	return nil
}

func analysisfeatureAfterSelectHook(e boil.Executor, o *Analysisfeature) error {
	*o = Analysisfeature{}
	return nil
}

func analysisfeatureBeforeUpdateHook(e boil.Executor, o *Analysisfeature) error {
	*o = Analysisfeature{}
	return nil
}

func analysisfeatureAfterUpdateHook(e boil.Executor, o *Analysisfeature) error {
	*o = Analysisfeature{}
	return nil
}

func analysisfeatureBeforeDeleteHook(e boil.Executor, o *Analysisfeature) error {
	*o = Analysisfeature{}
	return nil
}

func analysisfeatureAfterDeleteHook(e boil.Executor, o *Analysisfeature) error {
	*o = Analysisfeature{}
	return nil
}

func analysisfeatureBeforeUpsertHook(e boil.Executor, o *Analysisfeature) error {
	*o = Analysisfeature{}
	return nil
}

func analysisfeatureAfterUpsertHook(e boil.Executor, o *Analysisfeature) error {
	*o = Analysisfeature{}
	return nil
}

func testAnalysisfeaturesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Analysisfeature{}
	o := &Analysisfeature{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, analysisfeatureDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Analysisfeature object: %s", err)
	}

	AddAnalysisfeatureHook(boil.BeforeInsertHook, analysisfeatureBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	analysisfeatureBeforeInsertHooks = []AnalysisfeatureHook{}

	AddAnalysisfeatureHook(boil.AfterInsertHook, analysisfeatureAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	analysisfeatureAfterInsertHooks = []AnalysisfeatureHook{}

	AddAnalysisfeatureHook(boil.AfterSelectHook, analysisfeatureAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	analysisfeatureAfterSelectHooks = []AnalysisfeatureHook{}

	AddAnalysisfeatureHook(boil.BeforeUpdateHook, analysisfeatureBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	analysisfeatureBeforeUpdateHooks = []AnalysisfeatureHook{}

	AddAnalysisfeatureHook(boil.AfterUpdateHook, analysisfeatureAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	analysisfeatureAfterUpdateHooks = []AnalysisfeatureHook{}

	AddAnalysisfeatureHook(boil.BeforeDeleteHook, analysisfeatureBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	analysisfeatureBeforeDeleteHooks = []AnalysisfeatureHook{}

	AddAnalysisfeatureHook(boil.AfterDeleteHook, analysisfeatureAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	analysisfeatureAfterDeleteHooks = []AnalysisfeatureHook{}

	AddAnalysisfeatureHook(boil.BeforeUpsertHook, analysisfeatureBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	analysisfeatureBeforeUpsertHooks = []AnalysisfeatureHook{}

	AddAnalysisfeatureHook(boil.AfterUpsertHook, analysisfeatureAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	analysisfeatureAfterUpsertHooks = []AnalysisfeatureHook{}
}
func testAnalysisfeaturesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeature := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeature.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAnalysisfeaturesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeature := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeature.Insert(tx, analysisfeatureColumns...); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAnalysisfeatureOneToOneAnalysisfeaturepropUsingAnalysisfeatureprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Analysisfeatureprop
	var local Analysisfeature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, analysisfeaturepropDBTypes, true, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, analysisfeatureDBTypes, true, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.AnalysisfeatureID = local.AnalysisfeatureID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Analysisfeatureprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.AnalysisfeatureID != foreign.AnalysisfeatureID {
		t.Errorf("want: %v, got %v", foreign.AnalysisfeatureID, check.AnalysisfeatureID)
	}

	slice := AnalysisfeatureSlice{&local}
	if err = local.L.LoadAnalysisfeatureprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Analysisfeatureprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Analysisfeatureprop = nil
	if err = local.L.LoadAnalysisfeatureprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Analysisfeatureprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAnalysisfeatureOneToOneSetOpAnalysisfeaturepropUsingAnalysisfeatureprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Analysisfeature
	var b, c Analysisfeatureprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, analysisfeatureDBTypes, false, strmangle.SetComplement(analysisfeaturePrimaryKeyColumns, analysisfeatureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, analysisfeaturepropDBTypes, false, strmangle.SetComplement(analysisfeaturepropPrimaryKeyColumns, analysisfeaturepropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, analysisfeaturepropDBTypes, false, strmangle.SetComplement(analysisfeaturepropPrimaryKeyColumns, analysisfeaturepropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Analysisfeatureprop{&b, &c} {
		err = a.SetAnalysisfeatureprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Analysisfeatureprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Analysisfeature != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.AnalysisfeatureID != x.AnalysisfeatureID {
			t.Error("foreign key was wrong value", a.AnalysisfeatureID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.AnalysisfeatureID))
		reflect.Indirect(reflect.ValueOf(&x.AnalysisfeatureID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AnalysisfeatureID != x.AnalysisfeatureID {
			t.Error("foreign key was wrong value", a.AnalysisfeatureID, x.AnalysisfeatureID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testAnalysisfeatureToOneAnalysiUsingAnalysi(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Analysisfeature
	var foreign Analysi

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, analysisfeatureDBTypes, true, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, analysiDBTypes, true, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AnalysisID = foreign.AnalysisID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Analysi(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.AnalysisID != foreign.AnalysisID {
		t.Errorf("want: %v, got %v", foreign.AnalysisID, check.AnalysisID)
	}

	slice := AnalysisfeatureSlice{&local}
	if err = local.L.LoadAnalysi(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Analysi == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Analysi = nil
	if err = local.L.LoadAnalysi(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Analysi == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAnalysisfeatureToOneFeatureUsingFeature(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Analysisfeature
	var foreign Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, analysisfeatureDBTypes, true, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
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

	slice := AnalysisfeatureSlice{&local}
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

func testAnalysisfeatureToOneSetOpAnalysiUsingAnalysi(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Analysisfeature
	var b, c Analysi

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, analysisfeatureDBTypes, false, strmangle.SetComplement(analysisfeaturePrimaryKeyColumns, analysisfeatureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, analysiDBTypes, false, strmangle.SetComplement(analysiPrimaryKeyColumns, analysiColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, analysiDBTypes, false, strmangle.SetComplement(analysiPrimaryKeyColumns, analysiColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Analysi{&b, &c} {
		err = a.SetAnalysi(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Analysi != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Analysisfeature != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AnalysisID != x.AnalysisID {
			t.Error("foreign key was wrong value", a.AnalysisID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AnalysisID))
		reflect.Indirect(reflect.ValueOf(&a.AnalysisID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AnalysisID != x.AnalysisID {
			t.Error("foreign key was wrong value", a.AnalysisID, x.AnalysisID)
		}
	}
}
func testAnalysisfeatureToOneSetOpFeatureUsingFeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Analysisfeature
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, analysisfeatureDBTypes, false, strmangle.SetComplement(analysisfeaturePrimaryKeyColumns, analysisfeatureColumnsWithoutDefault)...); err != nil {
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

		if x.R.Analysisfeature != &a {
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
func testAnalysisfeaturesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeature := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeature.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = analysisfeature.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAnalysisfeaturesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeature := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeature.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AnalysisfeatureSlice{analysisfeature}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAnalysisfeaturesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysisfeature := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeature.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Analysisfeatures(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	analysisfeatureDBTypes = map[string]string{"AnalysisID": "integer", "AnalysisfeatureID": "integer", "FeatureID": "integer", "Identity": "double precision", "Normscore": "double precision", "Rawscore": "double precision", "Significance": "double precision"}
	_                      = bytes.MinRead
)

func testAnalysisfeaturesUpdate(t *testing.T) {
	t.Parallel()

	if len(analysisfeatureColumns) == len(analysisfeaturePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	analysisfeature := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeature.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	if err = analysisfeature.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAnalysisfeaturesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(analysisfeatureColumns) == len(analysisfeaturePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	analysisfeature := &Analysisfeature{}
	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeature.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analysisfeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, analysisfeature, analysisfeatureDBTypes, true, analysisfeaturePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(analysisfeatureColumns, analysisfeaturePrimaryKeyColumns) {
		fields = analysisfeatureColumns
	} else {
		fields = strmangle.SetComplement(
			analysisfeatureColumns,
			analysisfeaturePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(analysisfeature))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AnalysisfeatureSlice{analysisfeature}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAnalysisfeaturesUpsert(t *testing.T) {
	t.Parallel()

	if len(analysisfeatureColumns) == len(analysisfeaturePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	analysisfeature := Analysisfeature{}
	if err = randomize.Struct(seed, &analysisfeature, analysisfeatureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysisfeature.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Analysisfeature: %s", err)
	}

	count, err := Analysisfeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &analysisfeature, analysisfeatureDBTypes, false, analysisfeaturePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}

	if err = analysisfeature.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Analysisfeature: %s", err)
	}

	count, err = Analysisfeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
