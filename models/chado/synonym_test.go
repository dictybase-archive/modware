package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testSynonyms(t *testing.T) {
	t.Parallel()

	query := Synonyms(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testSynonymsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	synonym := &Synonym{}
	if err = randomize.Struct(seed, synonym, synonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonym.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = synonym.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Synonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSynonymsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	synonym := &Synonym{}
	if err = randomize.Struct(seed, synonym, synonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonym.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Synonyms(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Synonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSynonymsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	synonym := &Synonym{}
	if err = randomize.Struct(seed, synonym, synonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonym.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SynonymSlice{synonym}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Synonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testSynonymsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	synonym := &Synonym{}
	if err = randomize.Struct(seed, synonym, synonymDBTypes, true, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonym.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := SynonymExists(tx, synonym.SynonymID)
	if err != nil {
		t.Errorf("Unable to check if Synonym exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SynonymExistsG to return true, but got false.")
	}
}
func testSynonymsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	synonym := &Synonym{}
	if err = randomize.Struct(seed, synonym, synonymDBTypes, true, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonym.Insert(tx); err != nil {
		t.Error(err)
	}

	synonymFound, err := FindSynonym(tx, synonym.SynonymID)
	if err != nil {
		t.Error(err)
	}

	if synonymFound == nil {
		t.Error("want a record, got nil")
	}
}
func testSynonymsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	synonym := &Synonym{}
	if err = randomize.Struct(seed, synonym, synonymDBTypes, true, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonym.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Synonyms(tx).Bind(synonym); err != nil {
		t.Error(err)
	}
}

func testSynonymsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	synonym := &Synonym{}
	if err = randomize.Struct(seed, synonym, synonymDBTypes, true, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonym.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Synonyms(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSynonymsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	synonymOne := &Synonym{}
	synonymTwo := &Synonym{}
	if err = randomize.Struct(seed, synonymOne, synonymDBTypes, false, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}
	if err = randomize.Struct(seed, synonymTwo, synonymDBTypes, false, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonymOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = synonymTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Synonyms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSynonymsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	synonymOne := &Synonym{}
	synonymTwo := &Synonym{}
	if err = randomize.Struct(seed, synonymOne, synonymDBTypes, false, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}
	if err = randomize.Struct(seed, synonymTwo, synonymDBTypes, false, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonymOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = synonymTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Synonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func synonymBeforeInsertHook(e boil.Executor, o *Synonym) error {
	*o = Synonym{}
	return nil
}

func synonymAfterInsertHook(e boil.Executor, o *Synonym) error {
	*o = Synonym{}
	return nil
}

func synonymAfterSelectHook(e boil.Executor, o *Synonym) error {
	*o = Synonym{}
	return nil
}

func synonymBeforeUpdateHook(e boil.Executor, o *Synonym) error {
	*o = Synonym{}
	return nil
}

func synonymAfterUpdateHook(e boil.Executor, o *Synonym) error {
	*o = Synonym{}
	return nil
}

func synonymBeforeDeleteHook(e boil.Executor, o *Synonym) error {
	*o = Synonym{}
	return nil
}

func synonymAfterDeleteHook(e boil.Executor, o *Synonym) error {
	*o = Synonym{}
	return nil
}

func synonymBeforeUpsertHook(e boil.Executor, o *Synonym) error {
	*o = Synonym{}
	return nil
}

func synonymAfterUpsertHook(e boil.Executor, o *Synonym) error {
	*o = Synonym{}
	return nil
}

func testSynonymsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Synonym{}
	o := &Synonym{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, synonymDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Synonym object: %s", err)
	}

	AddSynonymHook(boil.BeforeInsertHook, synonymBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	synonymBeforeInsertHooks = []SynonymHook{}

	AddSynonymHook(boil.AfterInsertHook, synonymAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	synonymAfterInsertHooks = []SynonymHook{}

	AddSynonymHook(boil.AfterSelectHook, synonymAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	synonymAfterSelectHooks = []SynonymHook{}

	AddSynonymHook(boil.BeforeUpdateHook, synonymBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	synonymBeforeUpdateHooks = []SynonymHook{}

	AddSynonymHook(boil.AfterUpdateHook, synonymAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	synonymAfterUpdateHooks = []SynonymHook{}

	AddSynonymHook(boil.BeforeDeleteHook, synonymBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	synonymBeforeDeleteHooks = []SynonymHook{}

	AddSynonymHook(boil.AfterDeleteHook, synonymAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	synonymAfterDeleteHooks = []SynonymHook{}

	AddSynonymHook(boil.BeforeUpsertHook, synonymBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	synonymBeforeUpsertHooks = []SynonymHook{}

	AddSynonymHook(boil.AfterUpsertHook, synonymAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	synonymAfterUpsertHooks = []SynonymHook{}
}
func testSynonymsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	synonym := &Synonym{}
	if err = randomize.Struct(seed, synonym, synonymDBTypes, true, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonym.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Synonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSynonymsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	synonym := &Synonym{}
	if err = randomize.Struct(seed, synonym, synonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonym.Insert(tx, synonymColumns...); err != nil {
		t.Error(err)
	}

	count, err := Synonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSynonymOneToOneFeatureSynonymUsingFeatureSynonym(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureSynonym
	var local Synonym

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureSynonymDBTypes, true, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, synonymDBTypes, true, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.SynonymID = local.SynonymID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureSynonym(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.SynonymID != foreign.SynonymID {
		t.Errorf("want: %v, got %v", foreign.SynonymID, check.SynonymID)
	}

	slice := SynonymSlice{&local}
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

func testSynonymOneToOneSetOpFeatureSynonymUsingFeatureSynonym(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Synonym
	var b, c FeatureSynonym

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, synonymDBTypes, false, strmangle.SetComplement(synonymPrimaryKeyColumns, synonymColumnsWithoutDefault)...); err != nil {
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
		if x.R.Synonym != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.SynonymID != x.SynonymID {
			t.Error("foreign key was wrong value", a.SynonymID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.SynonymID))
		reflect.Indirect(reflect.ValueOf(&x.SynonymID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SynonymID != x.SynonymID {
			t.Error("foreign key was wrong value", a.SynonymID, x.SynonymID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testSynonymToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Synonym
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, synonymDBTypes, true, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
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

	slice := SynonymSlice{&local}
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

func testSynonymToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Synonym
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, synonymDBTypes, false, strmangle.SetComplement(synonymPrimaryKeyColumns, synonymColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeSynonym != &a {
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
func testSynonymsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	synonym := &Synonym{}
	if err = randomize.Struct(seed, synonym, synonymDBTypes, true, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonym.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = synonym.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testSynonymsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	synonym := &Synonym{}
	if err = randomize.Struct(seed, synonym, synonymDBTypes, true, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonym.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SynonymSlice{synonym}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testSynonymsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	synonym := &Synonym{}
	if err = randomize.Struct(seed, synonym, synonymDBTypes, true, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonym.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Synonyms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	synonymDBTypes = map[string]string{"Name": "character varying", "SynonymID": "integer", "SynonymSGML": "character varying", "TypeID": "integer"}
	_              = bytes.MinRead
)

func testSynonymsUpdate(t *testing.T) {
	t.Parallel()

	if len(synonymColumns) == len(synonymPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	synonym := &Synonym{}
	if err = randomize.Struct(seed, synonym, synonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonym.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Synonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, synonym, synonymDBTypes, true, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	if err = synonym.Update(tx); err != nil {
		t.Error(err)
	}
}

func testSynonymsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(synonymColumns) == len(synonymPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	synonym := &Synonym{}
	if err = randomize.Struct(seed, synonym, synonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonym.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Synonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, synonym, synonymDBTypes, true, synonymPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(synonymColumns, synonymPrimaryKeyColumns) {
		fields = synonymColumns
	} else {
		fields = strmangle.SetComplement(
			synonymColumns,
			synonymPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(synonym))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := SynonymSlice{synonym}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testSynonymsUpsert(t *testing.T) {
	t.Parallel()

	if len(synonymColumns) == len(synonymPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	synonym := Synonym{}
	if err = randomize.Struct(seed, &synonym, synonymDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = synonym.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Synonym: %s", err)
	}

	count, err := Synonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &synonym, synonymDBTypes, false, synonymPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}

	if err = synonym.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Synonym: %s", err)
	}

	count, err = Synonyms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
