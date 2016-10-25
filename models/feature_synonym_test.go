package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeatureSynonyms(t *testing.T) {
	t.Parallel()

	query := FeatureSynonyms(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeatureSynonymsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureSynonym := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureSynonym.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureSynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureSynonymsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureSynonym := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureSynonyms(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeatureSynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureSynonymsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureSynonym := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureSynonymSlice{featureSynonym}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureSynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeatureSynonymsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureSynonym := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeatureSynonymExists(tx, featureSynonym.FeatureSynonymID)
	if err != nil {
		t.Errorf("Unable to check if FeatureSynonym exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeatureSynonymExistsG to return true, but got false.")
	}
}
func testFeatureSynonymsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureSynonym := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	featureSynonymFound, err := FindFeatureSynonym(tx, featureSynonym.FeatureSynonymID)
	if err != nil {
		t.Error(err)
	}

	if featureSynonymFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeatureSynonymsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureSynonym := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeatureSynonyms(tx).Bind(featureSynonym); err != nil {
		t.Error(err)
	}
}

func testFeatureSynonymsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureSynonym := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeatureSynonyms(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeatureSynonymsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureSynonymOne := &FeatureSynonym{}
	featureSynonymTwo := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonymOne, featureSynonymDBTypes, false, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}
	if err = randomize.Struct(seed, featureSynonymTwo, featureSynonymDBTypes, false, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonymOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureSynonymTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureSynonyms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeatureSynonymsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featureSynonymOne := &FeatureSynonym{}
	featureSynonymTwo := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonymOne, featureSynonymDBTypes, false, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}
	if err = randomize.Struct(seed, featureSynonymTwo, featureSynonymDBTypes, false, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonymOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featureSynonymTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureSynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featureSynonymBeforeInsertHook(e boil.Executor, o *FeatureSynonym) error {
	*o = FeatureSynonym{}
	return nil
}

func featureSynonymAfterInsertHook(e boil.Executor, o *FeatureSynonym) error {
	*o = FeatureSynonym{}
	return nil
}

func featureSynonymAfterSelectHook(e boil.Executor, o *FeatureSynonym) error {
	*o = FeatureSynonym{}
	return nil
}

func featureSynonymBeforeUpdateHook(e boil.Executor, o *FeatureSynonym) error {
	*o = FeatureSynonym{}
	return nil
}

func featureSynonymAfterUpdateHook(e boil.Executor, o *FeatureSynonym) error {
	*o = FeatureSynonym{}
	return nil
}

func featureSynonymBeforeDeleteHook(e boil.Executor, o *FeatureSynonym) error {
	*o = FeatureSynonym{}
	return nil
}

func featureSynonymAfterDeleteHook(e boil.Executor, o *FeatureSynonym) error {
	*o = FeatureSynonym{}
	return nil
}

func featureSynonymBeforeUpsertHook(e boil.Executor, o *FeatureSynonym) error {
	*o = FeatureSynonym{}
	return nil
}

func featureSynonymAfterUpsertHook(e boil.Executor, o *FeatureSynonym) error {
	*o = FeatureSynonym{}
	return nil
}

func testFeatureSynonymsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeatureSynonym{}
	o := &FeatureSynonym{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featureSynonymDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym object: %s", err)
	}

	AddFeatureSynonymHook(boil.BeforeInsertHook, featureSynonymBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featureSynonymBeforeInsertHooks = []FeatureSynonymHook{}

	AddFeatureSynonymHook(boil.AfterInsertHook, featureSynonymAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featureSynonymAfterInsertHooks = []FeatureSynonymHook{}

	AddFeatureSynonymHook(boil.AfterSelectHook, featureSynonymAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featureSynonymAfterSelectHooks = []FeatureSynonymHook{}

	AddFeatureSynonymHook(boil.BeforeUpdateHook, featureSynonymBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featureSynonymBeforeUpdateHooks = []FeatureSynonymHook{}

	AddFeatureSynonymHook(boil.AfterUpdateHook, featureSynonymAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featureSynonymAfterUpdateHooks = []FeatureSynonymHook{}

	AddFeatureSynonymHook(boil.BeforeDeleteHook, featureSynonymBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featureSynonymBeforeDeleteHooks = []FeatureSynonymHook{}

	AddFeatureSynonymHook(boil.AfterDeleteHook, featureSynonymAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featureSynonymAfterDeleteHooks = []FeatureSynonymHook{}

	AddFeatureSynonymHook(boil.BeforeUpsertHook, featureSynonymBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featureSynonymBeforeUpsertHooks = []FeatureSynonymHook{}

	AddFeatureSynonymHook(boil.AfterUpsertHook, featureSynonymAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featureSynonymAfterUpsertHooks = []FeatureSynonymHook{}
}
func testFeatureSynonymsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureSynonym := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureSynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureSynonymsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureSynonym := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonym.Insert(tx, featureSynonymColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeatureSynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureSynonymToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureSynonym
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureSynonymDBTypes, true, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
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

	slice := FeatureSynonymSlice{&local}
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

func testFeatureSynonymToOneSynonymUsingSynonym(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureSynonym
	var foreign Synonym

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureSynonymDBTypes, true, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, synonymDBTypes, true, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.SynonymID = foreign.SynonymID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Synonym(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.SynonymID != foreign.SynonymID {
		t.Errorf("want: %v, got %v", foreign.SynonymID, check.SynonymID)
	}

	slice := FeatureSynonymSlice{&local}
	if err = local.L.LoadSynonym(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Synonym == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Synonym = nil
	if err = local.L.LoadSynonym(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Synonym == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeatureSynonymToOneFeatureUsingFeature(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeatureSynonym
	var foreign Feature

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featureSynonymDBTypes, true, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
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

	slice := FeatureSynonymSlice{&local}
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

func testFeatureSynonymToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureSynonym
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureSynonymDBTypes, false, strmangle.SetComplement(featureSynonymPrimaryKeyColumns, featureSynonymColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeatureSynonym != &a {
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
func testFeatureSynonymToOneSetOpSynonymUsingSynonym(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureSynonym
	var b, c Synonym

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureSynonymDBTypes, false, strmangle.SetComplement(featureSynonymPrimaryKeyColumns, featureSynonymColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, synonymDBTypes, false, strmangle.SetComplement(synonymPrimaryKeyColumns, synonymColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, synonymDBTypes, false, strmangle.SetComplement(synonymPrimaryKeyColumns, synonymColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Synonym{&b, &c} {
		err = a.SetSynonym(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Synonym != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FeatureSynonym != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SynonymID != x.SynonymID {
			t.Error("foreign key was wrong value", a.SynonymID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SynonymID))
		reflect.Indirect(reflect.ValueOf(&a.SynonymID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SynonymID != x.SynonymID {
			t.Error("foreign key was wrong value", a.SynonymID, x.SynonymID)
		}
	}
}
func testFeatureSynonymToOneSetOpFeatureUsingFeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeatureSynonym
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featureSynonymDBTypes, false, strmangle.SetComplement(featureSynonymPrimaryKeyColumns, featureSynonymColumnsWithoutDefault)...); err != nil {
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

		if x.R.FeatureSynonym != &a {
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
func testFeatureSynonymsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureSynonym := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featureSynonym.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureSynonymsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureSynonym := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeatureSynonymSlice{featureSynonym}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeatureSynonymsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureSynonym := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeatureSynonyms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featureSynonymDBTypes = map[string]string{"FeatureID": "integer", "FeatureSynonymID": "integer", "IsCurrent": "boolean", "IsInternal": "boolean", "PubID": "integer", "SynonymID": "integer"}
	_                     = bytes.MinRead
)

func testFeatureSynonymsUpdate(t *testing.T) {
	t.Parallel()

	if len(featureSynonymColumns) == len(featureSynonymPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureSynonym := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureSynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	if err = featureSynonym.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeatureSynonymsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featureSynonymColumns) == len(featureSynonymPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featureSynonym := &FeatureSynonym{}
	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonym.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeatureSynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featureSynonym, featureSynonymDBTypes, true, featureSynonymPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featureSynonymColumns, featureSynonymPrimaryKeyColumns) {
		fields = featureSynonymColumns
	} else {
		fields = strmangle.SetComplement(
			featureSynonymColumns,
			featureSynonymPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featureSynonym))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeatureSynonymSlice{featureSynonym}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeatureSynonymsUpsert(t *testing.T) {
	t.Parallel()

	if len(featureSynonymColumns) == len(featureSynonymPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featureSynonym := FeatureSynonym{}
	if err = randomize.Struct(seed, &featureSynonym, featureSynonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featureSynonym.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureSynonym: %s", err)
	}

	count, err := FeatureSynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featureSynonym, featureSynonymDBTypes, false, featureSynonymPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}

	if err = featureSynonym.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeatureSynonym: %s", err)
	}

	count, err = FeatureSynonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
