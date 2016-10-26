package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testPhenotypeprops(t *testing.T) {
	t.Parallel()

	query := Phenotypeprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testPhenotypepropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeprop := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = phenotypeprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhenotypepropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeprop := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Phenotypeprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Phenotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhenotypepropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeprop := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PhenotypepropSlice{phenotypeprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testPhenotypepropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeprop := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := PhenotypepropExists(tx, phenotypeprop.PhenotypepropID)
	if err != nil {
		t.Errorf("Unable to check if Phenotypeprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PhenotypepropExistsG to return true, but got false.")
	}
}
func testPhenotypepropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeprop := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	phenotypepropFound, err := FindPhenotypeprop(tx, phenotypeprop.PhenotypepropID)
	if err != nil {
		t.Error(err)
	}

	if phenotypepropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testPhenotypepropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeprop := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Phenotypeprops(tx).Bind(phenotypeprop); err != nil {
		t.Error(err)
	}
}

func testPhenotypepropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeprop := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Phenotypeprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPhenotypepropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypepropOne := &Phenotypeprop{}
	phenotypepropTwo := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypepropOne, phenotypepropDBTypes, false, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}
	if err = randomize.Struct(seed, phenotypepropTwo, phenotypepropDBTypes, false, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypepropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = phenotypepropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Phenotypeprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPhenotypepropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	phenotypepropOne := &Phenotypeprop{}
	phenotypepropTwo := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypepropOne, phenotypepropDBTypes, false, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}
	if err = randomize.Struct(seed, phenotypepropTwo, phenotypepropDBTypes, false, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypepropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = phenotypepropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func phenotypepropBeforeInsertHook(e boil.Executor, o *Phenotypeprop) error {
	*o = Phenotypeprop{}
	return nil
}

func phenotypepropAfterInsertHook(e boil.Executor, o *Phenotypeprop) error {
	*o = Phenotypeprop{}
	return nil
}

func phenotypepropAfterSelectHook(e boil.Executor, o *Phenotypeprop) error {
	*o = Phenotypeprop{}
	return nil
}

func phenotypepropBeforeUpdateHook(e boil.Executor, o *Phenotypeprop) error {
	*o = Phenotypeprop{}
	return nil
}

func phenotypepropAfterUpdateHook(e boil.Executor, o *Phenotypeprop) error {
	*o = Phenotypeprop{}
	return nil
}

func phenotypepropBeforeDeleteHook(e boil.Executor, o *Phenotypeprop) error {
	*o = Phenotypeprop{}
	return nil
}

func phenotypepropAfterDeleteHook(e boil.Executor, o *Phenotypeprop) error {
	*o = Phenotypeprop{}
	return nil
}

func phenotypepropBeforeUpsertHook(e boil.Executor, o *Phenotypeprop) error {
	*o = Phenotypeprop{}
	return nil
}

func phenotypepropAfterUpsertHook(e boil.Executor, o *Phenotypeprop) error {
	*o = Phenotypeprop{}
	return nil
}

func testPhenotypepropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Phenotypeprop{}
	o := &Phenotypeprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, phenotypepropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop object: %s", err)
	}

	AddPhenotypepropHook(boil.BeforeInsertHook, phenotypepropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	phenotypepropBeforeInsertHooks = []PhenotypepropHook{}

	AddPhenotypepropHook(boil.AfterInsertHook, phenotypepropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	phenotypepropAfterInsertHooks = []PhenotypepropHook{}

	AddPhenotypepropHook(boil.AfterSelectHook, phenotypepropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	phenotypepropAfterSelectHooks = []PhenotypepropHook{}

	AddPhenotypepropHook(boil.BeforeUpdateHook, phenotypepropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	phenotypepropBeforeUpdateHooks = []PhenotypepropHook{}

	AddPhenotypepropHook(boil.AfterUpdateHook, phenotypepropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	phenotypepropAfterUpdateHooks = []PhenotypepropHook{}

	AddPhenotypepropHook(boil.BeforeDeleteHook, phenotypepropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	phenotypepropBeforeDeleteHooks = []PhenotypepropHook{}

	AddPhenotypepropHook(boil.AfterDeleteHook, phenotypepropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	phenotypepropAfterDeleteHooks = []PhenotypepropHook{}

	AddPhenotypepropHook(boil.BeforeUpsertHook, phenotypepropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	phenotypepropBeforeUpsertHooks = []PhenotypepropHook{}

	AddPhenotypepropHook(boil.AfterUpsertHook, phenotypepropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	phenotypepropAfterUpsertHooks = []PhenotypepropHook{}
}
func testPhenotypepropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeprop := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhenotypepropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeprop := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeprop.Insert(tx, phenotypepropColumns...); err != nil {
		t.Error(err)
	}

	count, err := Phenotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhenotypepropToOnePhenotypeUsingPhenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Phenotypeprop
	var foreign Phenotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypepropDBTypes, true, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
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

	slice := PhenotypepropSlice{&local}
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

func testPhenotypepropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Phenotypeprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phenotypepropDBTypes, true, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
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

	slice := PhenotypepropSlice{&local}
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

func testPhenotypepropToOneSetOpPhenotypeUsingPhenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotypeprop
	var b, c Phenotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypepropDBTypes, false, strmangle.SetComplement(phenotypepropPrimaryKeyColumns, phenotypepropColumnsWithoutDefault)...); err != nil {
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

		if x.R.Phenotypeprop != &a {
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
func testPhenotypepropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phenotypeprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phenotypepropDBTypes, false, strmangle.SetComplement(phenotypepropPrimaryKeyColumns, phenotypepropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypePhenotypeprop != &a {
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
func testPhenotypepropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeprop := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = phenotypeprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPhenotypepropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeprop := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PhenotypepropSlice{phenotypeprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testPhenotypepropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phenotypeprop := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Phenotypeprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	phenotypepropDBTypes = map[string]string{"PhenotypeID": "integer", "PhenotypepropID": "integer", "Rank": "integer", "TypeID": "integer", "Value": "text"}
	_                    = bytes.MinRead
)

func testPhenotypepropsUpdate(t *testing.T) {
	t.Parallel()

	if len(phenotypepropColumns) == len(phenotypepropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	phenotypeprop := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true, phenotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	if err = phenotypeprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testPhenotypepropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(phenotypepropColumns) == len(phenotypepropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	phenotypeprop := &Phenotypeprop{}
	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phenotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, phenotypeprop, phenotypepropDBTypes, true, phenotypepropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(phenotypepropColumns, phenotypepropPrimaryKeyColumns) {
		fields = phenotypepropColumns
	} else {
		fields = strmangle.SetComplement(
			phenotypepropColumns,
			phenotypepropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(phenotypeprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := PhenotypepropSlice{phenotypeprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testPhenotypepropsUpsert(t *testing.T) {
	t.Parallel()

	if len(phenotypepropColumns) == len(phenotypepropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	phenotypeprop := Phenotypeprop{}
	if err = randomize.Struct(seed, &phenotypeprop, phenotypepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phenotypeprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Phenotypeprop: %s", err)
	}

	count, err := Phenotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &phenotypeprop, phenotypepropDBTypes, false, phenotypepropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Phenotypeprop struct: %s", err)
	}

	if err = phenotypeprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Phenotypeprop: %s", err)
	}

	count, err = Phenotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
