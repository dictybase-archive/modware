package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testPhenotypeCvterms(t *testing.T) {
	t.Parallel()

	query := PhenotypeCvterms(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testPhenotypeCvtermsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeCvterm := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = phenotypeCvterm.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhenotypeCvtermsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeCvterm := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = PhenotypeCvterms(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhenotypeCvtermsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeCvterm := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PhenotypeCvtermSlice{phenotypeCvterm}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testPhenotypeCvtermsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeCvterm := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := PhenotypeCvtermExists(tx, phenotypeCvterm.PhenotypeCvtermID)
	if err != nil {
		t.Errorf("Unable to check if PhenotypeCvterm exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PhenotypeCvtermExistsG to return true, but got false.")
	}
}
func testPhenotypeCvtermsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeCvterm := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	phenotypeCvtermFound, err := FindPhenotypeCvterm(tx, phenotypeCvterm.PhenotypeCvtermID)
	if err != nil {
		t.Error(err)
	}

	if phenotypeCvtermFound == nil {
		t.Error("want a record, got nil")
	}
}
func testPhenotypeCvtermsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeCvterm := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = PhenotypeCvterms(tx).Bind(phenotypeCvterm); err != nil {
		t.Error(err)
	}
}

func testPhenotypeCvtermsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeCvterm := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := PhenotypeCvterms(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPhenotypeCvtermsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeCvtermOne := &PhenotypeCvterm{}
	phenotypeCvtermTwo := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvtermOne, phenotypeCvtermDBTypes, false, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}
	if err = randomize.Struct(seed, phenotypeCvtermTwo, phenotypeCvtermDBTypes, false, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvtermOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = phenotypeCvtermTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := PhenotypeCvterms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPhenotypeCvtermsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	phenotypeCvtermOne := &PhenotypeCvterm{}
	phenotypeCvtermTwo := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvtermOne, phenotypeCvtermDBTypes, false, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}
	if err = randomize.Struct(seed, phenotypeCvtermTwo, phenotypeCvtermDBTypes, false, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvtermOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = phenotypeCvtermTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func phenotypeCvtermBeforeInsertHook(e boil.Executor, o *PhenotypeCvterm) error {
	*o = PhenotypeCvterm{}
	return nil
}

func phenotypeCvtermAfterInsertHook(e boil.Executor, o *PhenotypeCvterm) error {
	*o = PhenotypeCvterm{}
	return nil
}

func phenotypeCvtermAfterSelectHook(e boil.Executor, o *PhenotypeCvterm) error {
	*o = PhenotypeCvterm{}
	return nil
}

func phenotypeCvtermBeforeUpdateHook(e boil.Executor, o *PhenotypeCvterm) error {
	*o = PhenotypeCvterm{}
	return nil
}

func phenotypeCvtermAfterUpdateHook(e boil.Executor, o *PhenotypeCvterm) error {
	*o = PhenotypeCvterm{}
	return nil
}

func phenotypeCvtermBeforeDeleteHook(e boil.Executor, o *PhenotypeCvterm) error {
	*o = PhenotypeCvterm{}
	return nil
}

func phenotypeCvtermAfterDeleteHook(e boil.Executor, o *PhenotypeCvterm) error {
	*o = PhenotypeCvterm{}
	return nil
}

func phenotypeCvtermBeforeUpsertHook(e boil.Executor, o *PhenotypeCvterm) error {
	*o = PhenotypeCvterm{}
	return nil
}

func phenotypeCvtermAfterUpsertHook(e boil.Executor, o *PhenotypeCvterm) error {
	*o = PhenotypeCvterm{}
	return nil
}

func testPhenotypeCvtermsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &PhenotypeCvterm{}
	o := &PhenotypeCvterm{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, phenotypeCvtermDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm object: %s", err)
	}

	AddPhenotypeCvtermHook(boil.BeforeInsertHook, phenotypeCvtermBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	phenotypeCvtermBeforeInsertHooks = []PhenotypeCvtermHook{}

	AddPhenotypeCvtermHook(boil.AfterInsertHook, phenotypeCvtermAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	phenotypeCvtermAfterInsertHooks = []PhenotypeCvtermHook{}

	AddPhenotypeCvtermHook(boil.AfterSelectHook, phenotypeCvtermAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	phenotypeCvtermAfterSelectHooks = []PhenotypeCvtermHook{}

	AddPhenotypeCvtermHook(boil.BeforeUpdateHook, phenotypeCvtermBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	phenotypeCvtermBeforeUpdateHooks = []PhenotypeCvtermHook{}

	AddPhenotypeCvtermHook(boil.AfterUpdateHook, phenotypeCvtermAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	phenotypeCvtermAfterUpdateHooks = []PhenotypeCvtermHook{}

	AddPhenotypeCvtermHook(boil.BeforeDeleteHook, phenotypeCvtermBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	phenotypeCvtermBeforeDeleteHooks = []PhenotypeCvtermHook{}

	AddPhenotypeCvtermHook(boil.AfterDeleteHook, phenotypeCvtermAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	phenotypeCvtermAfterDeleteHooks = []PhenotypeCvtermHook{}

	AddPhenotypeCvtermHook(boil.BeforeUpsertHook, phenotypeCvtermBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	phenotypeCvtermBeforeUpsertHooks = []PhenotypeCvtermHook{}

	AddPhenotypeCvtermHook(boil.AfterUpsertHook, phenotypeCvtermAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	phenotypeCvtermAfterUpsertHooks = []PhenotypeCvtermHook{}
}
func testPhenotypeCvtermsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeCvterm := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhenotypeCvtermsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeCvterm := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvterm.Insert(tx, phenotypeCvtermColumns...); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhenotypeCvtermToOnePhenotypeUsingPhenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PhenotypeCvterm
	var foreign Phenotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeCvtermDBTypes, true, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
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

	slice := PhenotypeCvtermSlice{&local}
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

func testPhenotypeCvtermToOneCvtermUsingCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PhenotypeCvterm
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypeCvtermDBTypes, true, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
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

	slice := PhenotypeCvtermSlice{&local}
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

func testPhenotypeCvtermToOneSetOpPhenotypeUsingPhenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PhenotypeCvterm
	var b, c Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeCvtermDBTypes, false, strmangle.SetComplement(phenotypeCvtermPrimaryKeyColumns, phenotypeCvtermColumnsWithoutDefault)...); err != nil {
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

		if x.R.PhenotypeCvterm != &a {
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
func testPhenotypeCvtermToOneSetOpCvtermUsingCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PhenotypeCvterm
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypeCvtermDBTypes, false, strmangle.SetComplement(phenotypeCvtermPrimaryKeyColumns, phenotypeCvtermColumnsWithoutDefault)...); err != nil {
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

		if x.R.PhenotypeCvterm != &a {
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
func testPhenotypeCvtermsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeCvterm := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = phenotypeCvterm.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPhenotypeCvtermsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeCvterm := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PhenotypeCvtermSlice{phenotypeCvterm}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testPhenotypeCvtermsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeCvterm := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := PhenotypeCvterms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	phenotypeCvtermDBTypes = map[string]string{"CvtermID": "integer", "PhenotypeCvtermID": "integer", "PhenotypeID": "integer", "Rank": "integer"}
	_                      = bytes.MinRead
)

func testPhenotypeCvtermsUpdate(t *testing.T) {
	t.Parallel()

	if len(phenotypeCvtermColumns) == len(phenotypeCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	phenotypeCvterm := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true, phenotypeCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	if err = phenotypeCvterm.Update(tx); err != nil {
		t.Error(err)
	}
}

func testPhenotypeCvtermsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(phenotypeCvtermColumns) == len(phenotypeCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	phenotypeCvterm := &PhenotypeCvterm{}
	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvterm.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PhenotypeCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, phenotypeCvterm, phenotypeCvtermDBTypes, true, phenotypeCvtermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(phenotypeCvtermColumns, phenotypeCvtermPrimaryKeyColumns) {
		fields = phenotypeCvtermColumns
	} else {
		fields = strmangle.SetComplement(
			phenotypeCvtermColumns,
			phenotypeCvtermPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(phenotypeCvterm))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := PhenotypeCvtermSlice{phenotypeCvterm}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testPhenotypeCvtermsUpsert(t *testing.T) {
	t.Parallel()

	if len(phenotypeCvtermColumns) == len(phenotypeCvtermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	phenotypeCvterm := PhenotypeCvterm{}
	if err = randomize.Struct(seed, &phenotypeCvterm, phenotypeCvtermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeCvterm.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert PhenotypeCvterm: %s", err)
	}

	count, err := PhenotypeCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &phenotypeCvterm, phenotypeCvtermDBTypes, false, phenotypeCvtermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PhenotypeCvterm struct: %s", err)
	}

	if err = phenotypeCvterm.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert PhenotypeCvterm: %s", err)
	}

	count, err = PhenotypeCvterms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
