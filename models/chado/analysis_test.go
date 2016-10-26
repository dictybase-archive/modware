package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testAnalyses(t *testing.T) {
	t.Parallel()

	query := Analyses(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAnalysesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysi := &Analysi{}
	if err = randomize.Struct(seed, analysi, analysiDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysi.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = analysi.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Analyses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnalysesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysi := &Analysi{}
	if err = randomize.Struct(seed, analysi, analysiDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysi.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Analyses(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Analyses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnalysesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysi := &Analysi{}
	if err = randomize.Struct(seed, analysi, analysiDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysi.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AnalysiSlice{analysi}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Analyses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAnalysesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysi := &Analysi{}
	if err = randomize.Struct(seed, analysi, analysiDBTypes, true, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysi.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AnalysiExists(tx, analysi.AnalysisID)
	if err != nil {
		t.Errorf("Unable to check if Analysi exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AnalysiExistsG to return true, but got false.")
	}
}
func testAnalysesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysi := &Analysi{}
	if err = randomize.Struct(seed, analysi, analysiDBTypes, true, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysi.Insert(tx); err != nil {
		t.Error(err)
	}

	analysiFound, err := FindAnalysi(tx, analysi.AnalysisID)
	if err != nil {
		t.Error(err)
	}

	if analysiFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAnalysesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysi := &Analysi{}
	if err = randomize.Struct(seed, analysi, analysiDBTypes, true, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysi.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Analyses(tx).Bind(analysi); err != nil {
		t.Error(err)
	}
}

func testAnalysesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysi := &Analysi{}
	if err = randomize.Struct(seed, analysi, analysiDBTypes, true, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysi.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Analyses(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAnalysesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysiOne := &Analysi{}
	analysiTwo := &Analysi{}
	if err = randomize.Struct(seed, analysiOne, analysiDBTypes, false, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}
	if err = randomize.Struct(seed, analysiTwo, analysiDBTypes, false, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysiOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = analysiTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Analyses(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAnalysesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	analysiOne := &Analysi{}
	analysiTwo := &Analysi{}
	if err = randomize.Struct(seed, analysiOne, analysiDBTypes, false, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}
	if err = randomize.Struct(seed, analysiTwo, analysiDBTypes, false, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysiOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = analysiTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analyses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func analysiBeforeInsertHook(e boil.Executor, o *Analysi) error {
	*o = Analysi{}
	return nil
}

func analysiAfterInsertHook(e boil.Executor, o *Analysi) error {
	*o = Analysi{}
	return nil
}

func analysiAfterSelectHook(e boil.Executor, o *Analysi) error {
	*o = Analysi{}
	return nil
}

func analysiBeforeUpdateHook(e boil.Executor, o *Analysi) error {
	*o = Analysi{}
	return nil
}

func analysiAfterUpdateHook(e boil.Executor, o *Analysi) error {
	*o = Analysi{}
	return nil
}

func analysiBeforeDeleteHook(e boil.Executor, o *Analysi) error {
	*o = Analysi{}
	return nil
}

func analysiAfterDeleteHook(e boil.Executor, o *Analysi) error {
	*o = Analysi{}
	return nil
}

func analysiBeforeUpsertHook(e boil.Executor, o *Analysi) error {
	*o = Analysi{}
	return nil
}

func analysiAfterUpsertHook(e boil.Executor, o *Analysi) error {
	*o = Analysi{}
	return nil
}

func testAnalysesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Analysi{}
	o := &Analysi{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, analysiDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Analysi object: %s", err)
	}

	AddAnalysiHook(boil.BeforeInsertHook, analysiBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	analysiBeforeInsertHooks = []AnalysiHook{}

	AddAnalysiHook(boil.AfterInsertHook, analysiAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	analysiAfterInsertHooks = []AnalysiHook{}

	AddAnalysiHook(boil.AfterSelectHook, analysiAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	analysiAfterSelectHooks = []AnalysiHook{}

	AddAnalysiHook(boil.BeforeUpdateHook, analysiBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	analysiBeforeUpdateHooks = []AnalysiHook{}

	AddAnalysiHook(boil.AfterUpdateHook, analysiAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	analysiAfterUpdateHooks = []AnalysiHook{}

	AddAnalysiHook(boil.BeforeDeleteHook, analysiBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	analysiBeforeDeleteHooks = []AnalysiHook{}

	AddAnalysiHook(boil.AfterDeleteHook, analysiAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	analysiAfterDeleteHooks = []AnalysiHook{}

	AddAnalysiHook(boil.BeforeUpsertHook, analysiBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	analysiBeforeUpsertHooks = []AnalysiHook{}

	AddAnalysiHook(boil.AfterUpsertHook, analysiAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	analysiAfterUpsertHooks = []AnalysiHook{}
}
func testAnalysesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysi := &Analysi{}
	if err = randomize.Struct(seed, analysi, analysiDBTypes, true, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysi.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analyses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAnalysesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysi := &Analysi{}
	if err = randomize.Struct(seed, analysi, analysiDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysi.Insert(tx, analysiColumns...); err != nil {
		t.Error(err)
	}

	count, err := Analyses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAnalysiOneToOneAnalysisfeatureUsingAnalysisfeature(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Analysisfeature
	var local Analysi

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, analysisfeatureDBTypes, true, analysisfeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeature struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, analysiDBTypes, true, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.AnalysisID = local.AnalysisID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Analysisfeature(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.AnalysisID != foreign.AnalysisID {
		t.Errorf("want: %v, got %v", foreign.AnalysisID, check.AnalysisID)
	}

	slice := AnalysiSlice{&local}
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

func testAnalysiOneToOneAnalysispropUsingAnalysisprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Analysisprop
	var local Analysi

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, analysispropDBTypes, true, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, analysiDBTypes, true, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.AnalysisID = local.AnalysisID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Analysisprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.AnalysisID != foreign.AnalysisID {
		t.Errorf("want: %v, got %v", foreign.AnalysisID, check.AnalysisID)
	}

	slice := AnalysiSlice{&local}
	if err = local.L.LoadAnalysisprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Analysisprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Analysisprop = nil
	if err = local.L.LoadAnalysisprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Analysisprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAnalysiOneToOneSetOpAnalysisfeatureUsingAnalysisfeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Analysi
	var b, c Analysisfeature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, analysiDBTypes, false, strmangle.SetComplement(analysiPrimaryKeyColumns, analysiColumnsWithoutDefault)...); err != nil {
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
		if x.R.Analysi != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.AnalysisID != x.AnalysisID {
			t.Error("foreign key was wrong value", a.AnalysisID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.AnalysisID))
		reflect.Indirect(reflect.ValueOf(&x.AnalysisID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AnalysisID != x.AnalysisID {
			t.Error("foreign key was wrong value", a.AnalysisID, x.AnalysisID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testAnalysiOneToOneSetOpAnalysispropUsingAnalysisprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Analysi
	var b, c Analysisprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, analysiDBTypes, false, strmangle.SetComplement(analysiPrimaryKeyColumns, analysiColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, analysispropDBTypes, false, strmangle.SetComplement(analysispropPrimaryKeyColumns, analysispropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, analysispropDBTypes, false, strmangle.SetComplement(analysispropPrimaryKeyColumns, analysispropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Analysisprop{&b, &c} {
		err = a.SetAnalysisprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Analysisprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Analysi != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.AnalysisID != x.AnalysisID {
			t.Error("foreign key was wrong value", a.AnalysisID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.AnalysisID))
		reflect.Indirect(reflect.ValueOf(&x.AnalysisID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AnalysisID != x.AnalysisID {
			t.Error("foreign key was wrong value", a.AnalysisID, x.AnalysisID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testAnalysesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysi := &Analysi{}
	if err = randomize.Struct(seed, analysi, analysiDBTypes, true, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysi.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = analysi.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAnalysesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysi := &Analysi{}
	if err = randomize.Struct(seed, analysi, analysiDBTypes, true, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysi.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AnalysiSlice{analysi}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAnalysesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	analysi := &Analysi{}
	if err = randomize.Struct(seed, analysi, analysiDBTypes, true, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysi.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Analyses(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	analysiDBTypes = map[string]string{"Algorithm": "character varying", "AnalysisID": "integer", "Description": "text", "Name": "character varying", "Program": "character varying", "Programversion": "character varying", "Sourcename": "character varying", "Sourceuri": "text", "Sourceversion": "character varying", "Timeexecuted": "timestamp without time zone"}
	_              = bytes.MinRead
)

func testAnalysesUpdate(t *testing.T) {
	t.Parallel()

	if len(analysiColumns) == len(analysiPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	analysi := &Analysi{}
	if err = randomize.Struct(seed, analysi, analysiDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysi.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analyses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, analysi, analysiDBTypes, true, analysiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	if err = analysi.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAnalysesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(analysiColumns) == len(analysiPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	analysi := &Analysi{}
	if err = randomize.Struct(seed, analysi, analysiDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysi.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Analyses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, analysi, analysiDBTypes, true, analysiPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(analysiColumns, analysiPrimaryKeyColumns) {
		fields = analysiColumns
	} else {
		fields = strmangle.SetComplement(
			analysiColumns,
			analysiPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(analysi))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AnalysiSlice{analysi}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAnalysesUpsert(t *testing.T) {
	t.Parallel()

	if len(analysiColumns) == len(analysiPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	analysi := Analysi{}
	if err = randomize.Struct(seed, &analysi, analysiDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = analysi.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Analysi: %s", err)
	}

	count, err := Analyses(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &analysi, analysiDBTypes, false, analysiPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Analysi struct: %s", err)
	}

	if err = analysi.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Analysi: %s", err)
	}

	count, err = Analyses(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
