package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testGenotypeprops(t *testing.T) {
	t.Parallel()

	query := Genotypeprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testGenotypepropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotypeprop := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = genotypeprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Genotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGenotypepropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotypeprop := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Genotypeprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Genotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGenotypepropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotypeprop := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := GenotypepropSlice{genotypeprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Genotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testGenotypepropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotypeprop := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := GenotypepropExists(tx, genotypeprop.GenotypepropID)
	if err != nil {
		t.Errorf("Unable to check if Genotypeprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GenotypepropExistsG to return true, but got false.")
	}
}
func testGenotypepropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotypeprop := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	genotypepropFound, err := FindGenotypeprop(tx, genotypeprop.GenotypepropID)
	if err != nil {
		t.Error(err)
	}

	if genotypepropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testGenotypepropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotypeprop := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Genotypeprops(tx).Bind(genotypeprop); err != nil {
		t.Error(err)
	}
}

func testGenotypepropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotypeprop := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Genotypeprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGenotypepropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotypepropOne := &Genotypeprop{}
	genotypepropTwo := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypepropOne, genotypepropDBTypes, false, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}
	if err = randomize.Struct(seed, genotypepropTwo, genotypepropDBTypes, false, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypepropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = genotypepropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Genotypeprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGenotypepropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	genotypepropOne := &Genotypeprop{}
	genotypepropTwo := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypepropOne, genotypepropDBTypes, false, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}
	if err = randomize.Struct(seed, genotypepropTwo, genotypepropDBTypes, false, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypepropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = genotypepropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Genotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func genotypepropBeforeInsertHook(e boil.Executor, o *Genotypeprop) error {
	*o = Genotypeprop{}
	return nil
}

func genotypepropAfterInsertHook(e boil.Executor, o *Genotypeprop) error {
	*o = Genotypeprop{}
	return nil
}

func genotypepropAfterSelectHook(e boil.Executor, o *Genotypeprop) error {
	*o = Genotypeprop{}
	return nil
}

func genotypepropBeforeUpdateHook(e boil.Executor, o *Genotypeprop) error {
	*o = Genotypeprop{}
	return nil
}

func genotypepropAfterUpdateHook(e boil.Executor, o *Genotypeprop) error {
	*o = Genotypeprop{}
	return nil
}

func genotypepropBeforeDeleteHook(e boil.Executor, o *Genotypeprop) error {
	*o = Genotypeprop{}
	return nil
}

func genotypepropAfterDeleteHook(e boil.Executor, o *Genotypeprop) error {
	*o = Genotypeprop{}
	return nil
}

func genotypepropBeforeUpsertHook(e boil.Executor, o *Genotypeprop) error {
	*o = Genotypeprop{}
	return nil
}

func genotypepropAfterUpsertHook(e boil.Executor, o *Genotypeprop) error {
	*o = Genotypeprop{}
	return nil
}

func testGenotypepropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Genotypeprop{}
	o := &Genotypeprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, genotypepropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Genotypeprop object: %s", err)
	}

	AddGenotypepropHook(boil.BeforeInsertHook, genotypepropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	genotypepropBeforeInsertHooks = []GenotypepropHook{}

	AddGenotypepropHook(boil.AfterInsertHook, genotypepropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	genotypepropAfterInsertHooks = []GenotypepropHook{}

	AddGenotypepropHook(boil.AfterSelectHook, genotypepropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	genotypepropAfterSelectHooks = []GenotypepropHook{}

	AddGenotypepropHook(boil.BeforeUpdateHook, genotypepropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	genotypepropBeforeUpdateHooks = []GenotypepropHook{}

	AddGenotypepropHook(boil.AfterUpdateHook, genotypepropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	genotypepropAfterUpdateHooks = []GenotypepropHook{}

	AddGenotypepropHook(boil.BeforeDeleteHook, genotypepropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	genotypepropBeforeDeleteHooks = []GenotypepropHook{}

	AddGenotypepropHook(boil.AfterDeleteHook, genotypepropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	genotypepropAfterDeleteHooks = []GenotypepropHook{}

	AddGenotypepropHook(boil.BeforeUpsertHook, genotypepropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	genotypepropBeforeUpsertHooks = []GenotypepropHook{}

	AddGenotypepropHook(boil.AfterUpsertHook, genotypepropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	genotypepropAfterUpsertHooks = []GenotypepropHook{}
}
func testGenotypepropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotypeprop := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Genotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGenotypepropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotypeprop := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeprop.Insert(tx, genotypepropColumns...); err != nil {
		t.Error(err)
	}

	count, err := Genotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGenotypepropToOneGenotypeUsingGenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Genotypeprop
	var foreign Genotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, genotypepropDBTypes, true, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
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

	slice := GenotypepropSlice{&local}
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

func testGenotypepropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Genotypeprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, genotypepropDBTypes, true, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
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

	slice := GenotypepropSlice{&local}
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

func testGenotypepropToOneSetOpGenotypeUsingGenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Genotypeprop
	var b, c Genotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, genotypepropDBTypes, false, strmangle.SetComplement(genotypepropPrimaryKeyColumns, genotypepropColumnsWithoutDefault)...); err != nil {
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

		if x.R.Genotypeprop != &a {
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
func testGenotypepropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Genotypeprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, genotypepropDBTypes, false, strmangle.SetComplement(genotypepropPrimaryKeyColumns, genotypepropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeGenotypeprop != &a {
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
func testGenotypepropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotypeprop := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = genotypeprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testGenotypepropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotypeprop := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := GenotypepropSlice{genotypeprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testGenotypepropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genotypeprop := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Genotypeprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	genotypepropDBTypes = map[string]string{"GenotypeID": "integer", "GenotypepropID": "integer", "Rank": "integer", "TypeID": "integer", "Value": "text"}
	_                   = bytes.MinRead
)

func testGenotypepropsUpdate(t *testing.T) {
	t.Parallel()

	if len(genotypepropColumns) == len(genotypepropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	genotypeprop := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Genotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true, genotypepropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	if err = genotypeprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testGenotypepropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(genotypepropColumns) == len(genotypepropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	genotypeprop := &Genotypeprop{}
	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Genotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, genotypeprop, genotypepropDBTypes, true, genotypepropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(genotypepropColumns, genotypepropPrimaryKeyColumns) {
		fields = genotypepropColumns
	} else {
		fields = strmangle.SetComplement(
			genotypepropColumns,
			genotypepropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(genotypeprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := GenotypepropSlice{genotypeprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testGenotypepropsUpsert(t *testing.T) {
	t.Parallel()

	if len(genotypepropColumns) == len(genotypepropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	genotypeprop := Genotypeprop{}
	if err = randomize.Struct(seed, &genotypeprop, genotypepropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = genotypeprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Genotypeprop: %s", err)
	}

	count, err := Genotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &genotypeprop, genotypepropDBTypes, false, genotypepropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Genotypeprop struct: %s", err)
	}

	if err = genotypeprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Genotypeprop: %s", err)
	}

	count, err = Genotypeprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
