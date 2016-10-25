package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testCvterms(t *testing.T) {
	t.Parallel()

	query := Cvterms(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testCvtermsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvterm := &Cvterm{}
	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cvterm.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCvtermsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvterm := &Cvterm{}
	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Cvterms(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Cvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCvtermsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvterm := &Cvterm{}
	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CvtermSlice{cvterm}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testCvtermsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvterm := &Cvterm{}
	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := CvtermExists(tx, cvterm.CvtermID)
	if err != nil {
		t.Errorf("Unable to check if Cvterm exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CvtermExistsG to return true, but got false.")
	}
}
func testCvtermsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvterm := &Cvterm{}
	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	cvtermFound, err := FindCvterm(tx, cvterm.CvtermID)
	if err != nil {
		t.Error(err)
	}

	if cvtermFound == nil {
		t.Error("want a record, got nil")
	}
}
func testCvtermsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvterm := &Cvterm{}
	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Cvterms(tx).Bind(cvterm); err != nil {
		t.Error(err)
	}
}

func testCvtermsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvterm := &Cvterm{}
	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Cvterms(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCvtermsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvtermOne := &Cvterm{}
	cvtermTwo := &Cvterm{}
	if err = randomize.Struct(seed, cvtermOne, cvtermDBTypes, false, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}
	if err = randomize.Struct(seed, cvtermTwo, cvtermDBTypes, false, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvtermTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Cvterms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCvtermsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	cvtermOne := &Cvterm{}
	cvtermTwo := &Cvterm{}
	if err = randomize.Struct(seed, cvtermOne, cvtermDBTypes, false, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}
	if err = randomize.Struct(seed, cvtermTwo, cvtermDBTypes, false, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvtermOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvtermTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func cvtermBeforeInsertHook(e boil.Executor, o *Cvterm) error {
	*o = Cvterm{}
	return nil
}

func cvtermAfterInsertHook(e boil.Executor, o *Cvterm) error {
	*o = Cvterm{}
	return nil
}

func cvtermAfterSelectHook(e boil.Executor, o *Cvterm) error {
	*o = Cvterm{}
	return nil
}

func cvtermBeforeUpdateHook(e boil.Executor, o *Cvterm) error {
	*o = Cvterm{}
	return nil
}

func cvtermAfterUpdateHook(e boil.Executor, o *Cvterm) error {
	*o = Cvterm{}
	return nil
}

func cvtermBeforeDeleteHook(e boil.Executor, o *Cvterm) error {
	*o = Cvterm{}
	return nil
}

func cvtermAfterDeleteHook(e boil.Executor, o *Cvterm) error {
	*o = Cvterm{}
	return nil
}

func cvtermBeforeUpsertHook(e boil.Executor, o *Cvterm) error {
	*o = Cvterm{}
	return nil
}

func cvtermAfterUpsertHook(e boil.Executor, o *Cvterm) error {
	*o = Cvterm{}
	return nil
}

func testCvtermsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Cvterm{}
	o := &Cvterm{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, cvtermDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Cvterm object: %s", err)
	}

	AddCvtermHook(boil.BeforeInsertHook, cvtermBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	cvtermBeforeInsertHooks = []CvtermHook{}

	AddCvtermHook(boil.AfterInsertHook, cvtermAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	cvtermAfterInsertHooks = []CvtermHook{}

	AddCvtermHook(boil.AfterSelectHook, cvtermAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	cvtermAfterSelectHooks = []CvtermHook{}

	AddCvtermHook(boil.BeforeUpdateHook, cvtermBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	cvtermBeforeUpdateHooks = []CvtermHook{}

	AddCvtermHook(boil.AfterUpdateHook, cvtermAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	cvtermAfterUpdateHooks = []CvtermHook{}

	AddCvtermHook(boil.BeforeDeleteHook, cvtermBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	cvtermBeforeDeleteHooks = []CvtermHook{}

	AddCvtermHook(boil.AfterDeleteHook, cvtermAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	cvtermAfterDeleteHooks = []CvtermHook{}

	AddCvtermHook(boil.BeforeUpsertHook, cvtermBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	cvtermBeforeUpsertHooks = []CvtermHook{}

	AddCvtermHook(boil.AfterUpsertHook, cvtermAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	cvtermAfterUpsertHooks = []CvtermHook{}
}
func testCvtermsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvterm := &Cvterm{}
	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCvtermsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvterm := &Cvterm{}
	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvterm.Insert(tx, cvtermColumns...); err != nil {
		t.Error(err)
	}

	count, err := Cvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCvtermOneToOneSynonymUsingTypeSynonym(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Synonym
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, synonymDBTypes, true, synonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Synonym struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeSynonym(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeSynonym(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeSynonym == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeSynonym = nil
	if err = local.L.LoadTypeSynonym(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeSynonym == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOnePhenotypepropUsingTypePhenotypeprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Phenotypeprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenotypepropDBTypes, true, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypePhenotypeprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypePhenotypeprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypePhenotypeprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypePhenotypeprop = nil
	if err = local.L.LoadTypePhenotypeprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypePhenotypeprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneStockUsingTypeStock(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Stock
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeStock(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeStock(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeStock == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeStock = nil
	if err = local.L.LoadTypeStock(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeStock == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOnePubRelationshipUsingTypePubRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign PubRelationship
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, pubRelationshipDBTypes, true, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypePubRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypePubRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypePubRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypePubRelationship = nil
	if err = local.L.LoadTypePubRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypePubRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOnePubpropUsingTypePubprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Pubprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, pubpropDBTypes, true, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypePubprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypePubprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypePubprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypePubprop = nil
	if err = local.L.LoadTypePubprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypePubprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneCvtermsynonymUsingCvtermsynonym(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Cvtermsynonym
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, cvtermsynonymDBTypes, true, cvtermsynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermsynonym struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.CvtermID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Cvtermsynonym(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadCvtermsynonym(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Cvtermsynonym == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Cvtermsynonym = nil
	if err = local.L.LoadCvtermsynonym(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Cvtermsynonym == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneDbxrefpropUsingTypeDbxrefprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Dbxrefprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, dbxrefpropDBTypes, true, dbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxrefprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeDbxrefprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeDbxrefprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeDbxrefprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeDbxrefprop = nil
	if err = local.L.LoadTypeDbxrefprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeDbxrefprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneFeatureUsingTypeFeature(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Feature
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeFeature(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeFeature(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeFeature == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeFeature = nil
	if err = local.L.LoadTypeFeature(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeFeature == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneStockDbxrefpropUsingTypeStockDbxrefprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockDbxrefprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockDbxrefpropDBTypes, true, stockDbxrefpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockDbxrefprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeStockDbxrefprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeStockDbxrefprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeStockDbxrefprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeStockDbxrefprop = nil
	if err = local.L.LoadTypeStockDbxrefprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeStockDbxrefprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneEnvironmentCvtermUsingEnvironmentCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign EnvironmentCvterm
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, environmentCvtermDBTypes, true, environmentCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EnvironmentCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.CvtermID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.EnvironmentCvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadEnvironmentCvterm(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.EnvironmentCvterm == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.EnvironmentCvterm = nil
	if err = local.L.LoadEnvironmentCvterm(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.EnvironmentCvterm == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneStockCvtermUsingStockCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockCvterm
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.CvtermID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockCvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadStockCvterm(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockCvterm == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockCvterm = nil
	if err = local.L.LoadStockCvterm(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockCvterm == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneStockCvtermpropUsingTypeStockCvtermprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockCvtermprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockCvtermpropDBTypes, true, stockCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvtermprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeStockCvtermprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeStockCvtermprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeStockCvtermprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeStockCvtermprop = nil
	if err = local.L.LoadTypeStockCvtermprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeStockCvtermprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneStockcollectionUsingTypeStockcollection(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Stockcollection
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockcollectionDBTypes, true, stockcollectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollection struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeStockcollection(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeStockcollection(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeStockcollection == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeStockcollection = nil
	if err = local.L.LoadTypeStockcollection(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeStockcollection == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneAnalysisfeaturepropUsingTypeAnalysisfeatureprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Analysisfeatureprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, analysisfeaturepropDBTypes, true, analysisfeaturepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisfeatureprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeAnalysisfeatureprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeAnalysisfeatureprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeAnalysisfeatureprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeAnalysisfeatureprop = nil
	if err = local.L.LoadTypeAnalysisfeatureprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeAnalysisfeatureprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneCvtermDbxrefUsingCvtermDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign CvtermDbxref
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, cvtermDbxrefDBTypes, true, cvtermDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermDbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.CvtermID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.CvtermDbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadCvtermDbxref(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.CvtermDbxref == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.CvtermDbxref = nil
	if err = local.L.LoadCvtermDbxref(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.CvtermDbxref == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneFeaturepropUsingTypeFeatureprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Featureprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featurepropDBTypes, true, featurepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeFeatureprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeFeatureprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeFeatureprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeFeatureprop = nil
	if err = local.L.LoadTypeFeatureprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeFeatureprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneAnalysispropUsingTypeAnalysisprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Analysisprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, analysispropDBTypes, true, analysispropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Analysisprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeAnalysisprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeAnalysisprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeAnalysisprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeAnalysisprop = nil
	if err = local.L.LoadTypeAnalysisprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeAnalysisprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneFeatureCvtermUsingFeatureCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureCvterm
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.CvtermID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureCvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := CvtermSlice{&local}
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

func testCvtermOneToOneChadopropUsingTypeChadoprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Chadoprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, chadopropDBTypes, true, chadopropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chadoprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeChadoprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeChadoprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeChadoprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeChadoprop = nil
	if err = local.L.LoadTypeChadoprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeChadoprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneFeatureCvtermpropUsingTypeFeatureCvtermprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureCvtermprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureCvtermpropDBTypes, true, featureCvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeFeatureCvtermprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeFeatureCvtermprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeFeatureCvtermprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeFeatureCvtermprop = nil
	if err = local.L.LoadTypeFeatureCvtermprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeFeatureCvtermprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneFeatureGenotypeUsingFeatureGenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureGenotype
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureGenotypeDBTypes, true, featureGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureGenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.CvtermID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureGenotype(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := CvtermSlice{&local}
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

func testCvtermOneToOneContactRelationshipUsingTypeContactRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign ContactRelationship
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, contactRelationshipDBTypes, true, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeContactRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeContactRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeContactRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeContactRelationship = nil
	if err = local.L.LoadTypeContactRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeContactRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneFeaturePubpropUsingTypeFeaturePubprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeaturePubprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featurePubpropDBTypes, true, featurePubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePubprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeFeaturePubprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeFeaturePubprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeFeaturePubprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeFeaturePubprop = nil
	if err = local.L.LoadTypeFeaturePubprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeFeaturePubprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneCvpropUsingTypeCvprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Cvprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, cvpropDBTypes, true, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeCvprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeCvprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeCvprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeCvprop = nil
	if err = local.L.LoadTypeCvprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeCvprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneStockRelationshipUsingTypeStockRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockRelationship
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockRelationshipDBTypes, true, stockRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeStockRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeStockRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeStockRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeStockRelationship = nil
	if err = local.L.LoadTypeStockRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeStockRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneFeatureRelationshippropUsingTypeFeatureRelationshipprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureRelationshipprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureRelationshippropDBTypes, true, featureRelationshippropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeFeatureRelationshipprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeFeatureRelationshipprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeFeatureRelationshipprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeFeatureRelationshipprop = nil
	if err = local.L.LoadTypeFeatureRelationshipprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeFeatureRelationshipprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneFeatureRelationshipUsingTypeFeatureRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureRelationship
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureRelationshipDBTypes, true, featureRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeFeatureRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeFeatureRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeFeatureRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeFeatureRelationship = nil
	if err = local.L.LoadTypeFeatureRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeFeatureRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneGenotypepropUsingTypeGenotypeprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Genotypeprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, genotypepropDBTypes, true, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeGenotypeprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeGenotypeprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeGenotypeprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeGenotypeprop = nil
	if err = local.L.LoadTypeGenotypeprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeGenotypeprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneOrganismpropUsingTypeOrganismprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Organismprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, organismpropDBTypes, true, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeOrganismprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeOrganismprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeOrganismprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeOrganismprop = nil
	if err = local.L.LoadTypeOrganismprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeOrganismprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOnePhendescUsingTypePhendesc(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Phendesc
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypePhendesc(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypePhendesc(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypePhendesc == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypePhendesc = nil
	if err = local.L.LoadTypePhendesc(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypePhendesc == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOnePhenstatementUsingTypePhenstatement(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Phenstatement
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypePhenstatement(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypePhenstatement(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypePhenstatement == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypePhenstatement = nil
	if err = local.L.LoadTypePhenstatement(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypePhenstatement == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneStockcollectionpropUsingTypeStockcollectionprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Stockcollectionprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockcollectionpropDBTypes, true, stockcollectionpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockcollectionprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeStockcollectionprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeStockcollectionprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeStockcollectionprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeStockcollectionprop = nil
	if err = local.L.LoadTypeStockcollectionprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeStockcollectionprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneStockpropUsingTypeStockprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Stockprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockpropDBTypes, true, stockpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stockprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeStockprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeStockprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeStockprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeStockprop = nil
	if err = local.L.LoadTypeStockprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeStockprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneCvtermRelationshipUsingObjectCvtermRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign CvtermRelationship
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, cvtermRelationshipDBTypes, true, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.ObjectID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ObjectCvtermRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ObjectID != foreign.ObjectID {
		t.Errorf("want: %v, got %v", foreign.ObjectID, check.ObjectID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadObjectCvtermRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.ObjectCvtermRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ObjectCvtermRelationship = nil
	if err = local.L.LoadObjectCvtermRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.ObjectCvtermRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneCvtermRelationshipUsingSubjectCvtermRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign CvtermRelationship
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, cvtermRelationshipDBTypes, true, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.SubjectID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.SubjectCvtermRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.SubjectID != foreign.SubjectID {
		t.Errorf("want: %v, got %v", foreign.SubjectID, check.SubjectID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadSubjectCvtermRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.SubjectCvtermRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SubjectCvtermRelationship = nil
	if err = local.L.LoadSubjectCvtermRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.SubjectCvtermRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneCvtermRelationshipUsingTypeCvtermRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign CvtermRelationship
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, cvtermRelationshipDBTypes, true, cvtermRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CvtermRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeCvtermRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeCvtermRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeCvtermRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeCvtermRelationship = nil
	if err = local.L.LoadTypeCvtermRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeCvtermRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOnePhenotypeCvtermUsingPhenotypeCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign PhenotypeCvterm
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenotypeCvtermDBTypes, true, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.CvtermID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.PhenotypeCvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := CvtermSlice{&local}
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

func testCvtermOneToOnePhenotypeComparisonCvtermUsingPhenotypeComparisonCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign PhenotypeComparisonCvterm
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenotypeComparisonCvtermDBTypes, true, phenotypeComparisonCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparisonCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.CvtermID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.PhenotypeComparisonCvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadPhenotypeComparisonCvterm(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.PhenotypeComparisonCvterm == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.PhenotypeComparisonCvterm = nil
	if err = local.L.LoadPhenotypeComparisonCvterm(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.PhenotypeComparisonCvterm == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneCvtermpathUsingObjectCvtermpath(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Cvtermpath
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.ObjectID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ObjectCvtermpath(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ObjectID != foreign.ObjectID {
		t.Errorf("want: %v, got %v", foreign.ObjectID, check.ObjectID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadObjectCvtermpath(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.ObjectCvtermpath == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ObjectCvtermpath = nil
	if err = local.L.LoadObjectCvtermpath(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.ObjectCvtermpath == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneCvtermpathUsingSubjectCvtermpath(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Cvtermpath
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.SubjectID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.SubjectCvtermpath(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.SubjectID != foreign.SubjectID {
		t.Errorf("want: %v, got %v", foreign.SubjectID, check.SubjectID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadSubjectCvtermpath(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.SubjectCvtermpath == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SubjectCvtermpath = nil
	if err = local.L.LoadSubjectCvtermpath(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.SubjectCvtermpath == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneCvtermpathUsingTypeCvtermpath(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Cvtermpath
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, cvtermpathDBTypes, true, cvtermpathColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermpath struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	foreign.TypeID.Valid = true

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID.Int = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeCvtermpath(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID.Int != foreign.TypeID.Int {
		t.Errorf("want: %v, got %v", foreign.TypeID.Int, check.TypeID.Int)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeCvtermpath(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeCvtermpath == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeCvtermpath = nil
	if err = local.L.LoadTypeCvtermpath(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeCvtermpath == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneUserRelationshipUsingTypeUserRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign UserRelationship
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, userRelationshipDBTypes, true, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeUserRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeUserRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeUserRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeUserRelationship = nil
	if err = local.L.LoadTypeUserRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeUserRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneUserRelationshipUsingObjectUserRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign UserRelationship
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, userRelationshipDBTypes, true, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.ObjectID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ObjectUserRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ObjectID != foreign.ObjectID {
		t.Errorf("want: %v, got %v", foreign.ObjectID, check.ObjectID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadObjectUserRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.ObjectUserRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ObjectUserRelationship = nil
	if err = local.L.LoadObjectUserRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.ObjectUserRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneUserRelationshipUsingSubjectUserRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign UserRelationship
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, userRelationshipDBTypes, true, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.SubjectID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.SubjectUserRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.SubjectID != foreign.SubjectID {
		t.Errorf("want: %v, got %v", foreign.SubjectID, check.SubjectID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadSubjectUserRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.SubjectUserRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SubjectUserRelationship = nil
	if err = local.L.LoadSubjectUserRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.SubjectUserRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneCvtermpropUsingCvtermprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Cvtermprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, cvtermpropDBTypes, true, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.CvtermID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Cvtermprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadCvtermprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Cvtermprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Cvtermprop = nil
	if err = local.L.LoadCvtermprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Cvtermprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneCvtermpropUsingTypeCvtermprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Cvtermprop
	var local Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, cvtermpropDBTypes, true, cvtermpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvtermprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.TypeID = local.CvtermID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TypeCvtermprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.TypeID != foreign.TypeID {
		t.Errorf("want: %v, got %v", foreign.TypeID, check.TypeID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadTypeCvtermprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeCvtermprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TypeCvtermprop = nil
	if err = local.L.LoadTypeCvtermprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.TypeCvtermprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermOneToOneSetOpSynonymUsingTypeSynonym(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Synonym

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
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
		err = a.SetTypeSynonym(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeSynonym != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpPhenotypepropUsingTypePhenotypeprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Phenotypeprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
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
		err = a.SetTypePhenotypeprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypePhenotypeprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpStockUsingTypeStock(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Stock{&b, &c} {
		err = a.SetTypeStock(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeStock != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpPubRelationshipUsingTypePubRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c PubRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pubRelationshipDBTypes, false, strmangle.SetComplement(pubRelationshipPrimaryKeyColumns, pubRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pubRelationshipDBTypes, false, strmangle.SetComplement(pubRelationshipPrimaryKeyColumns, pubRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*PubRelationship{&b, &c} {
		err = a.SetTypePubRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypePubRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpPubpropUsingTypePubprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Pubprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pubpropDBTypes, false, strmangle.SetComplement(pubpropPrimaryKeyColumns, pubpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pubpropDBTypes, false, strmangle.SetComplement(pubpropPrimaryKeyColumns, pubpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Pubprop{&b, &c} {
		err = a.SetTypePubprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypePubprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpCvtermsynonymUsingCvtermsynonym(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Cvtermsynonym

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermsynonymDBTypes, false, strmangle.SetComplement(cvtermsynonymPrimaryKeyColumns, cvtermsynonymColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvtermsynonymDBTypes, false, strmangle.SetComplement(cvtermsynonymPrimaryKeyColumns, cvtermsynonymColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Cvtermsynonym{&b, &c} {
		err = a.SetCvtermsynonym(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Cvtermsynonym != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Cvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.CvtermID))
		reflect.Indirect(reflect.ValueOf(&x.CvtermID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID, x.CvtermID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpDbxrefpropUsingTypeDbxrefprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Dbxrefprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dbxrefpropDBTypes, false, strmangle.SetComplement(dbxrefpropPrimaryKeyColumns, dbxrefpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dbxrefpropDBTypes, false, strmangle.SetComplement(dbxrefpropPrimaryKeyColumns, dbxrefpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Dbxrefprop{&b, &c} {
		err = a.SetTypeDbxrefprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeDbxrefprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpFeatureUsingTypeFeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
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
		err = a.SetTypeFeature(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeFeature != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpStockDbxrefpropUsingTypeStockDbxrefprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c StockDbxrefprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockDbxrefpropDBTypes, false, strmangle.SetComplement(stockDbxrefpropPrimaryKeyColumns, stockDbxrefpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockDbxrefpropDBTypes, false, strmangle.SetComplement(stockDbxrefpropPrimaryKeyColumns, stockDbxrefpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockDbxrefprop{&b, &c} {
		err = a.SetTypeStockDbxrefprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeStockDbxrefprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpEnvironmentCvtermUsingEnvironmentCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c EnvironmentCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, environmentCvtermDBTypes, false, strmangle.SetComplement(environmentCvtermPrimaryKeyColumns, environmentCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, environmentCvtermDBTypes, false, strmangle.SetComplement(environmentCvtermPrimaryKeyColumns, environmentCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*EnvironmentCvterm{&b, &c} {
		err = a.SetEnvironmentCvterm(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.EnvironmentCvterm != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Cvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.CvtermID))
		reflect.Indirect(reflect.ValueOf(&x.CvtermID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID, x.CvtermID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpStockCvtermUsingStockCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c StockCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockCvtermDBTypes, false, strmangle.SetComplement(stockCvtermPrimaryKeyColumns, stockCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockCvtermDBTypes, false, strmangle.SetComplement(stockCvtermPrimaryKeyColumns, stockCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockCvterm{&b, &c} {
		err = a.SetStockCvterm(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockCvterm != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Cvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.CvtermID))
		reflect.Indirect(reflect.ValueOf(&x.CvtermID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID, x.CvtermID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpStockCvtermpropUsingTypeStockCvtermprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c StockCvtermprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockCvtermpropDBTypes, false, strmangle.SetComplement(stockCvtermpropPrimaryKeyColumns, stockCvtermpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockCvtermpropDBTypes, false, strmangle.SetComplement(stockCvtermpropPrimaryKeyColumns, stockCvtermpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockCvtermprop{&b, &c} {
		err = a.SetTypeStockCvtermprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeStockCvtermprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpStockcollectionUsingTypeStockcollection(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Stockcollection

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockcollectionDBTypes, false, strmangle.SetComplement(stockcollectionPrimaryKeyColumns, stockcollectionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockcollectionDBTypes, false, strmangle.SetComplement(stockcollectionPrimaryKeyColumns, stockcollectionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Stockcollection{&b, &c} {
		err = a.SetTypeStockcollection(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeStockcollection != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpAnalysisfeaturepropUsingTypeAnalysisfeatureprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Analysisfeatureprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
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
		err = a.SetTypeAnalysisfeatureprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeAnalysisfeatureprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpCvtermDbxrefUsingCvtermDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c CvtermDbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermDbxrefDBTypes, false, strmangle.SetComplement(cvtermDbxrefPrimaryKeyColumns, cvtermDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvtermDbxrefDBTypes, false, strmangle.SetComplement(cvtermDbxrefPrimaryKeyColumns, cvtermDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*CvtermDbxref{&b, &c} {
		err = a.SetCvtermDbxref(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.CvtermDbxref != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Cvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.CvtermID))
		reflect.Indirect(reflect.ValueOf(&x.CvtermID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID, x.CvtermID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpFeaturepropUsingTypeFeatureprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Featureprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
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
		err = a.SetTypeFeatureprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeFeatureprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpAnalysispropUsingTypeAnalysisprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Analysisprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
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
		err = a.SetTypeAnalysisprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeAnalysisprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpFeatureCvtermUsingFeatureCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c FeatureCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
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
		if x.R.Cvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.CvtermID))
		reflect.Indirect(reflect.ValueOf(&x.CvtermID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID, x.CvtermID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpChadopropUsingTypeChadoprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Chadoprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, chadopropDBTypes, false, strmangle.SetComplement(chadopropPrimaryKeyColumns, chadopropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, chadopropDBTypes, false, strmangle.SetComplement(chadopropPrimaryKeyColumns, chadopropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Chadoprop{&b, &c} {
		err = a.SetTypeChadoprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeChadoprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpFeatureCvtermpropUsingTypeFeatureCvtermprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c FeatureCvtermprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureCvtermpropDBTypes, false, strmangle.SetComplement(featureCvtermpropPrimaryKeyColumns, featureCvtermpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureCvtermpropDBTypes, false, strmangle.SetComplement(featureCvtermpropPrimaryKeyColumns, featureCvtermpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureCvtermprop{&b, &c} {
		err = a.SetTypeFeatureCvtermprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeFeatureCvtermprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpFeatureGenotypeUsingFeatureGenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c FeatureGenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
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
		if x.R.Cvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.CvtermID))
		reflect.Indirect(reflect.ValueOf(&x.CvtermID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID, x.CvtermID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpContactRelationshipUsingTypeContactRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c ContactRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, contactRelationshipDBTypes, false, strmangle.SetComplement(contactRelationshipPrimaryKeyColumns, contactRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, contactRelationshipDBTypes, false, strmangle.SetComplement(contactRelationshipPrimaryKeyColumns, contactRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ContactRelationship{&b, &c} {
		err = a.SetTypeContactRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeContactRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpFeaturePubpropUsingTypeFeaturePubprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c FeaturePubprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featurePubpropDBTypes, false, strmangle.SetComplement(featurePubpropPrimaryKeyColumns, featurePubpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featurePubpropDBTypes, false, strmangle.SetComplement(featurePubpropPrimaryKeyColumns, featurePubpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeaturePubprop{&b, &c} {
		err = a.SetTypeFeaturePubprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeFeaturePubprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpCvpropUsingTypeCvprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Cvprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvpropDBTypes, false, strmangle.SetComplement(cvpropPrimaryKeyColumns, cvpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvpropDBTypes, false, strmangle.SetComplement(cvpropPrimaryKeyColumns, cvpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Cvprop{&b, &c} {
		err = a.SetTypeCvprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeCvprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpStockRelationshipUsingTypeStockRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c StockRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockRelationshipDBTypes, false, strmangle.SetComplement(stockRelationshipPrimaryKeyColumns, stockRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockRelationshipDBTypes, false, strmangle.SetComplement(stockRelationshipPrimaryKeyColumns, stockRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockRelationship{&b, &c} {
		err = a.SetTypeStockRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeStockRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpFeatureRelationshippropUsingTypeFeatureRelationshipprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c FeatureRelationshipprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureRelationshippropDBTypes, false, strmangle.SetComplement(featureRelationshippropPrimaryKeyColumns, featureRelationshippropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureRelationshippropDBTypes, false, strmangle.SetComplement(featureRelationshippropPrimaryKeyColumns, featureRelationshippropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureRelationshipprop{&b, &c} {
		err = a.SetTypeFeatureRelationshipprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeFeatureRelationshipprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpFeatureRelationshipUsingTypeFeatureRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c FeatureRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
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
		err = a.SetTypeFeatureRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeFeatureRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpGenotypepropUsingTypeGenotypeprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Genotypeprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
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
		err = a.SetTypeGenotypeprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeGenotypeprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpOrganismpropUsingTypeOrganismprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Organismprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organismpropDBTypes, false, strmangle.SetComplement(organismpropPrimaryKeyColumns, organismpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, organismpropDBTypes, false, strmangle.SetComplement(organismpropPrimaryKeyColumns, organismpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Organismprop{&b, &c} {
		err = a.SetTypeOrganismprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeOrganismprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpPhendescUsingTypePhendesc(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Phendesc

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
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
		err = a.SetTypePhendesc(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypePhendesc != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpPhenstatementUsingTypePhenstatement(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Phenstatement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
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
		err = a.SetTypePhenstatement(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypePhenstatement != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpStockcollectionpropUsingTypeStockcollectionprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Stockcollectionprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockcollectionpropDBTypes, false, strmangle.SetComplement(stockcollectionpropPrimaryKeyColumns, stockcollectionpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockcollectionpropDBTypes, false, strmangle.SetComplement(stockcollectionpropPrimaryKeyColumns, stockcollectionpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Stockcollectionprop{&b, &c} {
		err = a.SetTypeStockcollectionprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeStockcollectionprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpStockpropUsingTypeStockprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Stockprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockpropDBTypes, false, strmangle.SetComplement(stockpropPrimaryKeyColumns, stockpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockpropDBTypes, false, strmangle.SetComplement(stockpropPrimaryKeyColumns, stockpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Stockprop{&b, &c} {
		err = a.SetTypeStockprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeStockprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpCvtermRelationshipUsingObjectCvtermRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c CvtermRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermRelationshipDBTypes, false, strmangle.SetComplement(cvtermRelationshipPrimaryKeyColumns, cvtermRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvtermRelationshipDBTypes, false, strmangle.SetComplement(cvtermRelationshipPrimaryKeyColumns, cvtermRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*CvtermRelationship{&b, &c} {
		err = a.SetObjectCvtermRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ObjectCvtermRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Object != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.ObjectID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.ObjectID))
		reflect.Indirect(reflect.ValueOf(&x.ObjectID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.ObjectID {
			t.Error("foreign key was wrong value", a.CvtermID, x.ObjectID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpCvtermRelationshipUsingSubjectCvtermRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c CvtermRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermRelationshipDBTypes, false, strmangle.SetComplement(cvtermRelationshipPrimaryKeyColumns, cvtermRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvtermRelationshipDBTypes, false, strmangle.SetComplement(cvtermRelationshipPrimaryKeyColumns, cvtermRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*CvtermRelationship{&b, &c} {
		err = a.SetSubjectCvtermRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SubjectCvtermRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Subject != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.SubjectID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.SubjectID))
		reflect.Indirect(reflect.ValueOf(&x.SubjectID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.SubjectID {
			t.Error("foreign key was wrong value", a.CvtermID, x.SubjectID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpCvtermRelationshipUsingTypeCvtermRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c CvtermRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermRelationshipDBTypes, false, strmangle.SetComplement(cvtermRelationshipPrimaryKeyColumns, cvtermRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvtermRelationshipDBTypes, false, strmangle.SetComplement(cvtermRelationshipPrimaryKeyColumns, cvtermRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*CvtermRelationship{&b, &c} {
		err = a.SetTypeCvtermRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeCvtermRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpPhenotypeCvtermUsingPhenotypeCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c PhenotypeCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
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
		if x.R.Cvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.CvtermID))
		reflect.Indirect(reflect.ValueOf(&x.CvtermID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID, x.CvtermID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpPhenotypeComparisonCvtermUsingPhenotypeComparisonCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c PhenotypeComparisonCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, phenotypeComparisonCvtermDBTypes, false, strmangle.SetComplement(phenotypeComparisonCvtermPrimaryKeyColumns, phenotypeComparisonCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, phenotypeComparisonCvtermDBTypes, false, strmangle.SetComplement(phenotypeComparisonCvtermPrimaryKeyColumns, phenotypeComparisonCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*PhenotypeComparisonCvterm{&b, &c} {
		err = a.SetPhenotypeComparisonCvterm(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.PhenotypeComparisonCvterm != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Cvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.CvtermID))
		reflect.Indirect(reflect.ValueOf(&x.CvtermID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID, x.CvtermID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpCvtermpathUsingObjectCvtermpath(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Cvtermpath

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermpathDBTypes, false, strmangle.SetComplement(cvtermpathPrimaryKeyColumns, cvtermpathColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvtermpathDBTypes, false, strmangle.SetComplement(cvtermpathPrimaryKeyColumns, cvtermpathColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Cvtermpath{&b, &c} {
		err = a.SetObjectCvtermpath(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ObjectCvtermpath != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Object != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.ObjectID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.ObjectID))
		reflect.Indirect(reflect.ValueOf(&x.ObjectID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.ObjectID {
			t.Error("foreign key was wrong value", a.CvtermID, x.ObjectID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpCvtermpathUsingSubjectCvtermpath(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Cvtermpath

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermpathDBTypes, false, strmangle.SetComplement(cvtermpathPrimaryKeyColumns, cvtermpathColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvtermpathDBTypes, false, strmangle.SetComplement(cvtermpathPrimaryKeyColumns, cvtermpathColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Cvtermpath{&b, &c} {
		err = a.SetSubjectCvtermpath(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SubjectCvtermpath != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Subject != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.SubjectID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.SubjectID))
		reflect.Indirect(reflect.ValueOf(&x.SubjectID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.SubjectID {
			t.Error("foreign key was wrong value", a.CvtermID, x.SubjectID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpCvtermpathUsingTypeCvtermpath(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Cvtermpath

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermpathDBTypes, false, strmangle.SetComplement(cvtermpathPrimaryKeyColumns, cvtermpathColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvtermpathDBTypes, false, strmangle.SetComplement(cvtermpathPrimaryKeyColumns, cvtermpathColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Cvtermpath{&b, &c} {
		err = a.SetTypeCvtermpath(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeCvtermpath != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID.Int {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID.Int))
		reflect.Indirect(reflect.ValueOf(&x.TypeID.Int)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID.Int {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID.Int)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testCvtermOneToOneRemoveOpCvtermpathUsingTypeCvtermpath(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b Cvtermpath

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermpathDBTypes, false, strmangle.SetComplement(cvtermpathPrimaryKeyColumns, cvtermpathColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetTypeCvtermpath(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveTypeCvtermpath(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.TypeCvtermpath(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.TypeCvtermpath != nil {
		t.Error("R struct entry should be nil")
	}

	if b.TypeID.Valid {
		t.Error("foreign key column should be nil")
	}

	if b.R.Type != nil {
		t.Error("failed to remove a from b's relationships")
	}
}

func testCvtermOneToOneSetOpUserRelationshipUsingTypeUserRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c UserRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userRelationshipDBTypes, false, strmangle.SetComplement(userRelationshipPrimaryKeyColumns, userRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userRelationshipDBTypes, false, strmangle.SetComplement(userRelationshipPrimaryKeyColumns, userRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*UserRelationship{&b, &c} {
		err = a.SetTypeUserRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeUserRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpUserRelationshipUsingObjectUserRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c UserRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userRelationshipDBTypes, false, strmangle.SetComplement(userRelationshipPrimaryKeyColumns, userRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userRelationshipDBTypes, false, strmangle.SetComplement(userRelationshipPrimaryKeyColumns, userRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*UserRelationship{&b, &c} {
		err = a.SetObjectUserRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ObjectUserRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Object != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.ObjectID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.ObjectID))
		reflect.Indirect(reflect.ValueOf(&x.ObjectID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.ObjectID {
			t.Error("foreign key was wrong value", a.CvtermID, x.ObjectID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpUserRelationshipUsingSubjectUserRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c UserRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userRelationshipDBTypes, false, strmangle.SetComplement(userRelationshipPrimaryKeyColumns, userRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userRelationshipDBTypes, false, strmangle.SetComplement(userRelationshipPrimaryKeyColumns, userRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*UserRelationship{&b, &c} {
		err = a.SetSubjectUserRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SubjectUserRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Subject != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.SubjectID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.SubjectID))
		reflect.Indirect(reflect.ValueOf(&x.SubjectID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.SubjectID {
			t.Error("foreign key was wrong value", a.CvtermID, x.SubjectID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpCvtermpropUsingCvtermprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Cvtermprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermpropDBTypes, false, strmangle.SetComplement(cvtermpropPrimaryKeyColumns, cvtermpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvtermpropDBTypes, false, strmangle.SetComplement(cvtermpropPrimaryKeyColumns, cvtermpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Cvtermprop{&b, &c} {
		err = a.SetCvtermprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Cvtermprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Cvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.CvtermID))
		reflect.Indirect(reflect.ValueOf(&x.CvtermID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID, x.CvtermID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermOneToOneSetOpCvtermpropUsingTypeCvtermprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Cvtermprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermpropDBTypes, false, strmangle.SetComplement(cvtermpropPrimaryKeyColumns, cvtermpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvtermpropDBTypes, false, strmangle.SetComplement(cvtermpropPrimaryKeyColumns, cvtermpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Cvtermprop{&b, &c} {
		err = a.SetTypeCvtermprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TypeCvtermprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Type != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.TypeID))
		reflect.Indirect(reflect.ValueOf(&x.TypeID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CvtermID != x.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, x.TypeID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCvtermToManyAssayPhenotypes(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, phenotypeDBTypes, false, phenotypeColumnsWithDefault...)
	randomize.Struct(seed, &c, phenotypeDBTypes, false, phenotypeColumnsWithDefault...)
	b.AssayID.Valid = true
	c.AssayID.Valid = true
	b.AssayID.Int = a.CvtermID
	c.AssayID.Int = a.CvtermID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	phenotype, err := a.AssayPhenotypes(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range phenotype {
		if v.AssayID.Int == b.AssayID.Int {
			bFound = true
		}
		if v.AssayID.Int == c.AssayID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CvtermSlice{&a}
	if err = a.L.LoadAssayPhenotypes(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AssayPhenotypes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AssayPhenotypes = nil
	if err = a.L.LoadAssayPhenotypes(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AssayPhenotypes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", phenotype)
	}
}

func testCvtermToManyAttrPhenotypes(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, phenotypeDBTypes, false, phenotypeColumnsWithDefault...)
	randomize.Struct(seed, &c, phenotypeDBTypes, false, phenotypeColumnsWithDefault...)
	b.AttrID.Valid = true
	c.AttrID.Valid = true
	b.AttrID.Int = a.CvtermID
	c.AttrID.Int = a.CvtermID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	phenotype, err := a.AttrPhenotypes(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range phenotype {
		if v.AttrID.Int == b.AttrID.Int {
			bFound = true
		}
		if v.AttrID.Int == c.AttrID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CvtermSlice{&a}
	if err = a.L.LoadAttrPhenotypes(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AttrPhenotypes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AttrPhenotypes = nil
	if err = a.L.LoadAttrPhenotypes(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AttrPhenotypes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", phenotype)
	}
}

func testCvtermToManyCvaluePhenotypes(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, phenotypeDBTypes, false, phenotypeColumnsWithDefault...)
	randomize.Struct(seed, &c, phenotypeDBTypes, false, phenotypeColumnsWithDefault...)
	b.CvalueID.Valid = true
	c.CvalueID.Valid = true
	b.CvalueID.Int = a.CvtermID
	c.CvalueID.Int = a.CvtermID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	phenotype, err := a.CvaluePhenotypes(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range phenotype {
		if v.CvalueID.Int == b.CvalueID.Int {
			bFound = true
		}
		if v.CvalueID.Int == c.CvalueID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CvtermSlice{&a}
	if err = a.L.LoadCvaluePhenotypes(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CvaluePhenotypes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.CvaluePhenotypes = nil
	if err = a.L.LoadCvaluePhenotypes(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CvaluePhenotypes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", phenotype)
	}
}

func testCvtermToManyObservablePhenotypes(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, phenotypeDBTypes, false, phenotypeColumnsWithDefault...)
	randomize.Struct(seed, &c, phenotypeDBTypes, false, phenotypeColumnsWithDefault...)
	b.ObservableID.Valid = true
	c.ObservableID.Valid = true
	b.ObservableID.Int = a.CvtermID
	c.ObservableID.Int = a.CvtermID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	phenotype, err := a.ObservablePhenotypes(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range phenotype {
		if v.ObservableID.Int == b.ObservableID.Int {
			bFound = true
		}
		if v.ObservableID.Int == c.ObservableID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CvtermSlice{&a}
	if err = a.L.LoadObservablePhenotypes(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ObservablePhenotypes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ObservablePhenotypes = nil
	if err = a.L.LoadObservablePhenotypes(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ObservablePhenotypes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", phenotype)
	}
}

func testCvtermToManyTypeCvtermsynonyms(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Cvtermsynonym

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, cvtermsynonymDBTypes, false, cvtermsynonymColumnsWithDefault...)
	randomize.Struct(seed, &c, cvtermsynonymDBTypes, false, cvtermsynonymColumnsWithDefault...)
	b.TypeID.Valid = true
	c.TypeID.Valid = true
	b.TypeID.Int = a.CvtermID
	c.TypeID.Int = a.CvtermID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	cvtermsynonym, err := a.TypeCvtermsynonyms(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range cvtermsynonym {
		if v.TypeID.Int == b.TypeID.Int {
			bFound = true
		}
		if v.TypeID.Int == c.TypeID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CvtermSlice{&a}
	if err = a.L.LoadTypeCvtermsynonyms(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TypeCvtermsynonyms); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TypeCvtermsynonyms = nil
	if err = a.L.LoadTypeCvtermsynonyms(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TypeCvtermsynonyms); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", cvtermsynonym)
	}
}

func testCvtermToManyTypeJbrowseTracks(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c JbrowseTrack

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, jbrowseTrackDBTypes, false, jbrowseTrackColumnsWithDefault...)
	randomize.Struct(seed, &c, jbrowseTrackDBTypes, false, jbrowseTrackColumnsWithDefault...)
	b.TypeID.Valid = true
	c.TypeID.Valid = true
	b.TypeID.Int = a.CvtermID
	c.TypeID.Int = a.CvtermID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	jbrowseTrack, err := a.TypeJbrowseTracks(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range jbrowseTrack {
		if v.TypeID.Int == b.TypeID.Int {
			bFound = true
		}
		if v.TypeID.Int == c.TypeID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CvtermSlice{&a}
	if err = a.L.LoadTypeJbrowseTracks(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TypeJbrowseTracks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TypeJbrowseTracks = nil
	if err = a.L.LoadTypeJbrowseTracks(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TypeJbrowseTracks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", jbrowseTrack)
	}
}

func testCvtermToManyTypePubs(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, pubDBTypes, false, pubColumnsWithDefault...)
	randomize.Struct(seed, &c, pubDBTypes, false, pubColumnsWithDefault...)

	b.TypeID = a.CvtermID
	c.TypeID = a.CvtermID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	pub, err := a.TypePubs(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range pub {
		if v.TypeID == b.TypeID {
			bFound = true
		}
		if v.TypeID == c.TypeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CvtermSlice{&a}
	if err = a.L.LoadTypePubs(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TypePubs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TypePubs = nil
	if err = a.L.LoadTypePubs(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TypePubs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", pub)
	}
}

func testCvtermToManyStockRelationshipCvterms(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c StockRelationshipCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, stockRelationshipCvtermDBTypes, false, stockRelationshipCvtermColumnsWithDefault...)
	randomize.Struct(seed, &c, stockRelationshipCvtermDBTypes, false, stockRelationshipCvtermColumnsWithDefault...)

	b.CvtermID = a.CvtermID
	c.CvtermID = a.CvtermID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	stockRelationshipCvterm, err := a.StockRelationshipCvterms(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range stockRelationshipCvterm {
		if v.CvtermID == b.CvtermID {
			bFound = true
		}
		if v.CvtermID == c.CvtermID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CvtermSlice{&a}
	if err = a.L.LoadStockRelationshipCvterms(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.StockRelationshipCvterms); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.StockRelationshipCvterms = nil
	if err = a.L.LoadStockRelationshipCvterms(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.StockRelationshipCvterms); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", stockRelationshipCvterm)
	}
}

func testCvtermToManyTypeContacts(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Contact

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, contactDBTypes, false, contactColumnsWithDefault...)
	randomize.Struct(seed, &c, contactDBTypes, false, contactColumnsWithDefault...)
	b.TypeID.Valid = true
	c.TypeID.Valid = true
	b.TypeID.Int = a.CvtermID
	c.TypeID.Int = a.CvtermID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	contact, err := a.TypeContacts(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range contact {
		if v.TypeID.Int == b.TypeID.Int {
			bFound = true
		}
		if v.TypeID.Int == c.TypeID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CvtermSlice{&a}
	if err = a.L.LoadTypeContacts(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TypeContacts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TypeContacts = nil
	if err = a.L.LoadTypeContacts(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TypeContacts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", contact)
	}
}

func testCvtermToManyTypeGenotypes(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Genotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, genotypeDBTypes, false, genotypeColumnsWithDefault...)
	randomize.Struct(seed, &c, genotypeDBTypes, false, genotypeColumnsWithDefault...)

	b.TypeID = a.CvtermID
	c.TypeID = a.CvtermID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	genotype, err := a.TypeGenotypes(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range genotype {
		if v.TypeID == b.TypeID {
			bFound = true
		}
		if v.TypeID == c.TypeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CvtermSlice{&a}
	if err = a.L.LoadTypeGenotypes(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TypeGenotypes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TypeGenotypes = nil
	if err = a.L.LoadTypeGenotypes(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TypeGenotypes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", genotype)
	}
}

func testCvtermToManyAddOpAssayPhenotypes(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Phenotype{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Phenotype{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAssayPhenotypes(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.CvtermID != first.AssayID.Int {
			t.Error("foreign key was wrong value", a.CvtermID, first.AssayID.Int)
		}
		if a.CvtermID != second.AssayID.Int {
			t.Error("foreign key was wrong value", a.CvtermID, second.AssayID.Int)
		}

		if first.R.Assay != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Assay != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AssayPhenotypes[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AssayPhenotypes[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AssayPhenotypes(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCvtermToManySetOpAssayPhenotypes(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Phenotype{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
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

	err = a.SetAssayPhenotypes(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.AssayPhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetAssayPhenotypes(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.AssayPhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.AssayID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.AssayID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.CvtermID != d.AssayID.Int {
		t.Error("foreign key was wrong value", a.CvtermID, d.AssayID.Int)
	}
	if a.CvtermID != e.AssayID.Int {
		t.Error("foreign key was wrong value", a.CvtermID, e.AssayID.Int)
	}

	if b.R.Assay != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Assay != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Assay != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Assay != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.AssayPhenotypes[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.AssayPhenotypes[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testCvtermToManyRemoveOpAssayPhenotypes(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Phenotype{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddAssayPhenotypes(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.AssayPhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveAssayPhenotypes(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.AssayPhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.AssayID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.AssayID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Assay != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Assay != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Assay != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Assay != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.AssayPhenotypes) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.AssayPhenotypes[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.AssayPhenotypes[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testCvtermToManyAddOpAttrPhenotypes(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Phenotype{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Phenotype{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAttrPhenotypes(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.CvtermID != first.AttrID.Int {
			t.Error("foreign key was wrong value", a.CvtermID, first.AttrID.Int)
		}
		if a.CvtermID != second.AttrID.Int {
			t.Error("foreign key was wrong value", a.CvtermID, second.AttrID.Int)
		}

		if first.R.Attr != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Attr != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AttrPhenotypes[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AttrPhenotypes[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AttrPhenotypes(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCvtermToManySetOpAttrPhenotypes(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Phenotype{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
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

	err = a.SetAttrPhenotypes(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.AttrPhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetAttrPhenotypes(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.AttrPhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.AttrID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.AttrID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.CvtermID != d.AttrID.Int {
		t.Error("foreign key was wrong value", a.CvtermID, d.AttrID.Int)
	}
	if a.CvtermID != e.AttrID.Int {
		t.Error("foreign key was wrong value", a.CvtermID, e.AttrID.Int)
	}

	if b.R.Attr != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Attr != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Attr != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Attr != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.AttrPhenotypes[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.AttrPhenotypes[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testCvtermToManyRemoveOpAttrPhenotypes(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Phenotype{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddAttrPhenotypes(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.AttrPhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveAttrPhenotypes(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.AttrPhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.AttrID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.AttrID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Attr != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Attr != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Attr != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Attr != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.AttrPhenotypes) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.AttrPhenotypes[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.AttrPhenotypes[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testCvtermToManyAddOpCvaluePhenotypes(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Phenotype{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Phenotype{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCvaluePhenotypes(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.CvtermID != first.CvalueID.Int {
			t.Error("foreign key was wrong value", a.CvtermID, first.CvalueID.Int)
		}
		if a.CvtermID != second.CvalueID.Int {
			t.Error("foreign key was wrong value", a.CvtermID, second.CvalueID.Int)
		}

		if first.R.Cvalue != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Cvalue != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.CvaluePhenotypes[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.CvaluePhenotypes[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.CvaluePhenotypes(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCvtermToManySetOpCvaluePhenotypes(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Phenotype{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
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

	err = a.SetCvaluePhenotypes(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.CvaluePhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetCvaluePhenotypes(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.CvaluePhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.CvalueID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.CvalueID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.CvtermID != d.CvalueID.Int {
		t.Error("foreign key was wrong value", a.CvtermID, d.CvalueID.Int)
	}
	if a.CvtermID != e.CvalueID.Int {
		t.Error("foreign key was wrong value", a.CvtermID, e.CvalueID.Int)
	}

	if b.R.Cvalue != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Cvalue != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Cvalue != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Cvalue != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.CvaluePhenotypes[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.CvaluePhenotypes[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testCvtermToManyRemoveOpCvaluePhenotypes(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Phenotype{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddCvaluePhenotypes(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.CvaluePhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveCvaluePhenotypes(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.CvaluePhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.CvalueID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.CvalueID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Cvalue != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Cvalue != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Cvalue != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Cvalue != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.CvaluePhenotypes) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.CvaluePhenotypes[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.CvaluePhenotypes[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testCvtermToManyAddOpObservablePhenotypes(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Phenotype{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Phenotype{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddObservablePhenotypes(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.CvtermID != first.ObservableID.Int {
			t.Error("foreign key was wrong value", a.CvtermID, first.ObservableID.Int)
		}
		if a.CvtermID != second.ObservableID.Int {
			t.Error("foreign key was wrong value", a.CvtermID, second.ObservableID.Int)
		}

		if first.R.Observable != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Observable != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ObservablePhenotypes[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ObservablePhenotypes[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ObservablePhenotypes(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCvtermToManySetOpObservablePhenotypes(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Phenotype{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
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

	err = a.SetObservablePhenotypes(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ObservablePhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetObservablePhenotypes(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ObservablePhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.ObservableID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.ObservableID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.CvtermID != d.ObservableID.Int {
		t.Error("foreign key was wrong value", a.CvtermID, d.ObservableID.Int)
	}
	if a.CvtermID != e.ObservableID.Int {
		t.Error("foreign key was wrong value", a.CvtermID, e.ObservableID.Int)
	}

	if b.R.Observable != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Observable != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Observable != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Observable != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.ObservablePhenotypes[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.ObservablePhenotypes[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testCvtermToManyRemoveOpObservablePhenotypes(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Phenotype{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeDBTypes, false, strmangle.SetComplement(phenotypePrimaryKeyColumns, phenotypeColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddObservablePhenotypes(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ObservablePhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveObservablePhenotypes(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ObservablePhenotypes(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.ObservableID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.ObservableID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Observable != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Observable != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Observable != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Observable != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.ObservablePhenotypes) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.ObservablePhenotypes[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.ObservablePhenotypes[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testCvtermToManyAddOpTypeCvtermsynonyms(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Cvtermsynonym

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Cvtermsynonym{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, cvtermsynonymDBTypes, false, strmangle.SetComplement(cvtermsynonymPrimaryKeyColumns, cvtermsynonymColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Cvtermsynonym{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTypeCvtermsynonyms(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.CvtermID != first.TypeID.Int {
			t.Error("foreign key was wrong value", a.CvtermID, first.TypeID.Int)
		}
		if a.CvtermID != second.TypeID.Int {
			t.Error("foreign key was wrong value", a.CvtermID, second.TypeID.Int)
		}

		if first.R.Type != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Type != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.TypeCvtermsynonyms[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TypeCvtermsynonyms[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TypeCvtermsynonyms(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCvtermToManySetOpTypeCvtermsynonyms(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Cvtermsynonym

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Cvtermsynonym{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, cvtermsynonymDBTypes, false, strmangle.SetComplement(cvtermsynonymPrimaryKeyColumns, cvtermsynonymColumnsWithoutDefault)...); err != nil {
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

	err = a.SetTypeCvtermsynonyms(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.TypeCvtermsynonyms(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetTypeCvtermsynonyms(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.TypeCvtermsynonyms(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.TypeID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.TypeID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.CvtermID != d.TypeID.Int {
		t.Error("foreign key was wrong value", a.CvtermID, d.TypeID.Int)
	}
	if a.CvtermID != e.TypeID.Int {
		t.Error("foreign key was wrong value", a.CvtermID, e.TypeID.Int)
	}

	if b.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Type != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Type != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.TypeCvtermsynonyms[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.TypeCvtermsynonyms[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testCvtermToManyRemoveOpTypeCvtermsynonyms(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Cvtermsynonym

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Cvtermsynonym{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, cvtermsynonymDBTypes, false, strmangle.SetComplement(cvtermsynonymPrimaryKeyColumns, cvtermsynonymColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddTypeCvtermsynonyms(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.TypeCvtermsynonyms(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveTypeCvtermsynonyms(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.TypeCvtermsynonyms(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.TypeID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.TypeID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Type != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Type != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.TypeCvtermsynonyms) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.TypeCvtermsynonyms[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.TypeCvtermsynonyms[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testCvtermToManyAddOpTypeJbrowseTracks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e JbrowseTrack

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*JbrowseTrack{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, jbrowseTrackDBTypes, false, strmangle.SetComplement(jbrowseTrackPrimaryKeyColumns, jbrowseTrackColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*JbrowseTrack{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTypeJbrowseTracks(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.CvtermID != first.TypeID.Int {
			t.Error("foreign key was wrong value", a.CvtermID, first.TypeID.Int)
		}
		if a.CvtermID != second.TypeID.Int {
			t.Error("foreign key was wrong value", a.CvtermID, second.TypeID.Int)
		}

		if first.R.Type != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Type != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.TypeJbrowseTracks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TypeJbrowseTracks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TypeJbrowseTracks(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCvtermToManySetOpTypeJbrowseTracks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e JbrowseTrack

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*JbrowseTrack{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, jbrowseTrackDBTypes, false, strmangle.SetComplement(jbrowseTrackPrimaryKeyColumns, jbrowseTrackColumnsWithoutDefault)...); err != nil {
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

	err = a.SetTypeJbrowseTracks(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.TypeJbrowseTracks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetTypeJbrowseTracks(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.TypeJbrowseTracks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.TypeID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.TypeID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.CvtermID != d.TypeID.Int {
		t.Error("foreign key was wrong value", a.CvtermID, d.TypeID.Int)
	}
	if a.CvtermID != e.TypeID.Int {
		t.Error("foreign key was wrong value", a.CvtermID, e.TypeID.Int)
	}

	if b.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Type != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Type != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.TypeJbrowseTracks[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.TypeJbrowseTracks[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testCvtermToManyRemoveOpTypeJbrowseTracks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e JbrowseTrack

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*JbrowseTrack{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, jbrowseTrackDBTypes, false, strmangle.SetComplement(jbrowseTrackPrimaryKeyColumns, jbrowseTrackColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddTypeJbrowseTracks(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.TypeJbrowseTracks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveTypeJbrowseTracks(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.TypeJbrowseTracks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.TypeID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.TypeID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Type != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Type != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.TypeJbrowseTracks) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.TypeJbrowseTracks[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.TypeJbrowseTracks[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testCvtermToManyAddOpTypePubs(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Pub{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Pub{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTypePubs(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.CvtermID != first.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, first.TypeID)
		}
		if a.CvtermID != second.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, second.TypeID)
		}

		if first.R.Type != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Type != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.TypePubs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TypePubs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TypePubs(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testCvtermToManyAddOpStockRelationshipCvterms(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e StockRelationshipCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*StockRelationshipCvterm{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, stockRelationshipCvtermDBTypes, false, strmangle.SetComplement(stockRelationshipCvtermPrimaryKeyColumns, stockRelationshipCvtermColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*StockRelationshipCvterm{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddStockRelationshipCvterms(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.CvtermID != first.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID, first.CvtermID)
		}
		if a.CvtermID != second.CvtermID {
			t.Error("foreign key was wrong value", a.CvtermID, second.CvtermID)
		}

		if first.R.Cvterm != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Cvterm != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.StockRelationshipCvterms[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.StockRelationshipCvterms[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.StockRelationshipCvterms(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testCvtermToManyAddOpTypeContacts(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Contact

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Contact{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Contact{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTypeContacts(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.CvtermID != first.TypeID.Int {
			t.Error("foreign key was wrong value", a.CvtermID, first.TypeID.Int)
		}
		if a.CvtermID != second.TypeID.Int {
			t.Error("foreign key was wrong value", a.CvtermID, second.TypeID.Int)
		}

		if first.R.Type != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Type != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.TypeContacts[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TypeContacts[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TypeContacts(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCvtermToManySetOpTypeContacts(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Contact

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Contact{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
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

	err = a.SetTypeContacts(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.TypeContacts(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetTypeContacts(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.TypeContacts(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.TypeID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.TypeID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.CvtermID != d.TypeID.Int {
		t.Error("foreign key was wrong value", a.CvtermID, d.TypeID.Int)
	}
	if a.CvtermID != e.TypeID.Int {
		t.Error("foreign key was wrong value", a.CvtermID, e.TypeID.Int)
	}

	if b.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Type != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Type != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.TypeContacts[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.TypeContacts[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testCvtermToManyRemoveOpTypeContacts(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Contact

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Contact{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddTypeContacts(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.TypeContacts(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveTypeContacts(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.TypeContacts(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.TypeID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.TypeID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Type != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Type != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.TypeContacts) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.TypeContacts[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.TypeContacts[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testCvtermToManyAddOpTypeGenotypes(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c, d, e Genotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Genotype{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, genotypeDBTypes, false, strmangle.SetComplement(genotypePrimaryKeyColumns, genotypeColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Genotype{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTypeGenotypes(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.CvtermID != first.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, first.TypeID)
		}
		if a.CvtermID != second.TypeID {
			t.Error("foreign key was wrong value", a.CvtermID, second.TypeID)
		}

		if first.R.Type != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Type != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.TypeGenotypes[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TypeGenotypes[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TypeGenotypes(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testCvtermToOneCVUsingCV(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Cvterm
	var foreign CV

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cvDBTypes, true, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.CVID = foreign.CVID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.CV(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CVID != foreign.CVID {
		t.Errorf("want: %v, got %v", foreign.CVID, check.CVID)
	}

	slice := CvtermSlice{&local}
	if err = local.L.LoadCV(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.CV == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.CV = nil
	if err = local.L.LoadCV(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.CV == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCvtermToOneDbxrefUsingDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Cvterm
	var foreign Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dbxrefDBTypes, true, dbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dbxref struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.DbxrefID = foreign.DbxrefID
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

	slice := CvtermSlice{&local}
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

func testCvtermToOneSetOpCVUsingCV(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c CV

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvDBTypes, false, strmangle.SetComplement(cvPrimaryKeyColumns, cvColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvDBTypes, false, strmangle.SetComplement(cvPrimaryKeyColumns, cvColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*CV{&b, &c} {
		err = a.SetCV(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.CV != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Cvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CVID != x.CVID {
			t.Error("foreign key was wrong value", a.CVID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CVID))
		reflect.Indirect(reflect.ValueOf(&a.CVID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CVID != x.CVID {
			t.Error("foreign key was wrong value", a.CVID, x.CVID)
		}
	}
}
func testCvtermToOneSetOpDbxrefUsingDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Cvterm
	var b, c Dbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
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

		if x.R.Cvterm != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DbxrefID))
		reflect.Indirect(reflect.ValueOf(&a.DbxrefID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DbxrefID != x.DbxrefID {
			t.Error("foreign key was wrong value", a.DbxrefID, x.DbxrefID)
		}
	}
}
func testCvtermsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvterm := &Cvterm{}
	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cvterm.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCvtermsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvterm := &Cvterm{}
	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CvtermSlice{cvterm}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testCvtermsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvterm := &Cvterm{}
	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Cvterms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	cvtermDBTypes = map[string]string{"CVID": "integer", "CvtermID": "integer", "DbxrefID": "integer", "Definition": "text", "IsObsolete": "integer", "IsRelationshiptype": "integer", "Name": "character varying"}
	_             = bytes.MinRead
)

func testCvtermsUpdate(t *testing.T) {
	t.Parallel()

	if len(cvtermColumns) == len(cvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cvterm := &Cvterm{}
	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err = cvterm.Update(tx); err != nil {
		t.Error(err)
	}
}

func testCvtermsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(cvtermColumns) == len(cvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cvterm := &Cvterm{}
	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Cvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cvterm, cvtermDBTypes, true, cvtermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(cvtermColumns, cvtermPrimaryKeyColumns) {
		fields = cvtermColumns
	} else {
		fields = strmangle.SetComplement(
			cvtermColumns,
			cvtermPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(cvterm))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := CvtermSlice{cvterm}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testCvtermsUpsert(t *testing.T) {
	t.Parallel()

	if len(cvtermColumns) == len(cvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	cvterm := Cvterm{}
	if err = randomize.Struct(seed, &cvterm, cvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvterm.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Cvterm: %s", err)
	}

	count, err := Cvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &cvterm, cvtermDBTypes, false, cvtermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err = cvterm.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Cvterm: %s", err)
	}

	count, err = Cvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
