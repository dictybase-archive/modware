package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testPhendescs(t *testing.T) {
	t.Parallel()

	query := Phendescs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testPhendescsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phendesc := &Phendesc{}
	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendesc.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = phendesc.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Phendescs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhendescsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phendesc := &Phendesc{}
	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendesc.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Phendescs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Phendescs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhendescsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phendesc := &Phendesc{}
	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendesc.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PhendescSlice{phendesc}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Phendescs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testPhendescsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phendesc := &Phendesc{}
	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendesc.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := PhendescExists(tx, phendesc.PhendescID)
	if err != nil {
		t.Errorf("Unable to check if Phendesc exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PhendescExistsG to return true, but got false.")
	}
}
func testPhendescsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phendesc := &Phendesc{}
	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendesc.Insert(tx); err != nil {
		t.Error(err)
	}

	phendescFound, err := FindPhendesc(tx, phendesc.PhendescID)
	if err != nil {
		t.Error(err)
	}

	if phendescFound == nil {
		t.Error("want a record, got nil")
	}
}
func testPhendescsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phendesc := &Phendesc{}
	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendesc.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Phendescs(tx).Bind(phendesc); err != nil {
		t.Error(err)
	}
}

func testPhendescsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phendesc := &Phendesc{}
	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendesc.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Phendescs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPhendescsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phendescOne := &Phendesc{}
	phendescTwo := &Phendesc{}
	if err = randomize.Struct(seed, phendescOne, phendescDBTypes, false, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}
	if err = randomize.Struct(seed, phendescTwo, phendescDBTypes, false, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendescOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = phendescTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Phendescs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPhendescsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	phendescOne := &Phendesc{}
	phendescTwo := &Phendesc{}
	if err = randomize.Struct(seed, phendescOne, phendescDBTypes, false, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}
	if err = randomize.Struct(seed, phendescTwo, phendescDBTypes, false, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendescOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = phendescTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phendescs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func phendescBeforeInsertHook(e boil.Executor, o *Phendesc) error {
	*o = Phendesc{}
	return nil
}

func phendescAfterInsertHook(e boil.Executor, o *Phendesc) error {
	*o = Phendesc{}
	return nil
}

func phendescAfterSelectHook(e boil.Executor, o *Phendesc) error {
	*o = Phendesc{}
	return nil
}

func phendescBeforeUpdateHook(e boil.Executor, o *Phendesc) error {
	*o = Phendesc{}
	return nil
}

func phendescAfterUpdateHook(e boil.Executor, o *Phendesc) error {
	*o = Phendesc{}
	return nil
}

func phendescBeforeDeleteHook(e boil.Executor, o *Phendesc) error {
	*o = Phendesc{}
	return nil
}

func phendescAfterDeleteHook(e boil.Executor, o *Phendesc) error {
	*o = Phendesc{}
	return nil
}

func phendescBeforeUpsertHook(e boil.Executor, o *Phendesc) error {
	*o = Phendesc{}
	return nil
}

func phendescAfterUpsertHook(e boil.Executor, o *Phendesc) error {
	*o = Phendesc{}
	return nil
}

func testPhendescsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Phendesc{}
	o := &Phendesc{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, phendescDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Phendesc object: %s", err)
	}

	AddPhendescHook(boil.BeforeInsertHook, phendescBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	phendescBeforeInsertHooks = []PhendescHook{}

	AddPhendescHook(boil.AfterInsertHook, phendescAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	phendescAfterInsertHooks = []PhendescHook{}

	AddPhendescHook(boil.AfterSelectHook, phendescAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	phendescAfterSelectHooks = []PhendescHook{}

	AddPhendescHook(boil.BeforeUpdateHook, phendescBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	phendescBeforeUpdateHooks = []PhendescHook{}

	AddPhendescHook(boil.AfterUpdateHook, phendescAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	phendescAfterUpdateHooks = []PhendescHook{}

	AddPhendescHook(boil.BeforeDeleteHook, phendescBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	phendescBeforeDeleteHooks = []PhendescHook{}

	AddPhendescHook(boil.AfterDeleteHook, phendescAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	phendescAfterDeleteHooks = []PhendescHook{}

	AddPhendescHook(boil.BeforeUpsertHook, phendescBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	phendescBeforeUpsertHooks = []PhendescHook{}

	AddPhendescHook(boil.AfterUpsertHook, phendescAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	phendescAfterUpsertHooks = []PhendescHook{}
}
func testPhendescsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phendesc := &Phendesc{}
	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendesc.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phendescs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhendescsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phendesc := &Phendesc{}
	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendesc.Insert(tx, phendescColumns...); err != nil {
		t.Error(err)
	}

	count, err := Phendescs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhendescToOneEnvironmentUsingEnvironment(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Phendesc
	var foreign Environment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, environmentDBTypes, true, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.EnvironmentID = foreign.EnvironmentID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Environment(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.EnvironmentID != foreign.EnvironmentID {
		t.Errorf("want: %v, got %v", foreign.EnvironmentID, check.EnvironmentID)
	}

	slice := PhendescSlice{&local}
	if err = local.L.LoadEnvironment(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Environment == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Environment = nil
	if err = local.L.LoadEnvironment(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Environment == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPhendescToOneGenotypeUsingGenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Phendesc
	var foreign Genotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
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

	slice := PhendescSlice{&local}
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

func testPhendescToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Phendesc
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
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

	slice := PhendescSlice{&local}
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

func testPhendescToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Phendesc
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
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

	slice := PhendescSlice{&local}
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

func testPhendescToOneSetOpEnvironmentUsingEnvironment(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phendesc
	var b, c Environment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phendescDBTypes, false, strmangle.SetComplement(phendescPrimaryKeyColumns, phendescColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Environment{&b, &c} {
		err = a.SetEnvironment(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Environment != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Phendesc != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.EnvironmentID != x.EnvironmentID {
			t.Error("foreign key was wrong value", a.EnvironmentID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.EnvironmentID))
		reflect.Indirect(reflect.ValueOf(&a.EnvironmentID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.EnvironmentID != x.EnvironmentID {
			t.Error("foreign key was wrong value", a.EnvironmentID, x.EnvironmentID)
		}
	}
}
func testPhendescToOneSetOpGenotypeUsingGenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phendesc
	var b, c Genotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phendescDBTypes, false, strmangle.SetComplement(phendescPrimaryKeyColumns, phendescColumnsWithoutDefault)...); err != nil {
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

		if x.R.Phendesc != &a {
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
func testPhendescToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phendesc
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phendescDBTypes, false, strmangle.SetComplement(phendescPrimaryKeyColumns, phendescColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypePhendesc != &a {
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
func testPhendescToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Phendesc
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, phendescDBTypes, false, strmangle.SetComplement(phendescPrimaryKeyColumns, phendescColumnsWithoutDefault)...); err != nil {
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

		if x.R.Phendesc != &a {
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
func testPhendescsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phendesc := &Phendesc{}
	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendesc.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = phendesc.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPhendescsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phendesc := &Phendesc{}
	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendesc.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PhendescSlice{phendesc}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testPhendescsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	phendesc := &Phendesc{}
	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendesc.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Phendescs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	phendescDBTypes = map[string]string{"Description": "text", "EnvironmentID": "integer", "GenotypeID": "integer", "PhendescID": "integer", "PubID": "integer", "TypeID": "integer"}
	_               = bytes.MinRead
)

func testPhendescsUpdate(t *testing.T) {
	t.Parallel()

	if len(phendescColumns) == len(phendescPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	phendesc := &Phendesc{}
	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendesc.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phendescs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	if err = phendesc.Update(tx); err != nil {
		t.Error(err)
	}
}

func testPhendescsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(phendescColumns) == len(phendescPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	phendesc := &Phendesc{}
	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendesc.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Phendescs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, phendesc, phendescDBTypes, true, phendescPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(phendescColumns, phendescPrimaryKeyColumns) {
		fields = phendescColumns
	} else {
		fields = strmangle.SetComplement(
			phendescColumns,
			phendescPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(phendesc))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := PhendescSlice{phendesc}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testPhendescsUpsert(t *testing.T) {
	t.Parallel()

	if len(phendescColumns) == len(phendescPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	phendesc := Phendesc{}
	if err = randomize.Struct(seed, &phendesc, phendescDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = phendesc.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Phendesc: %s", err)
	}

	count, err := Phendescs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &phendesc, phendescDBTypes, false, phendescPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}

	if err = phendesc.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Phendesc: %s", err)
	}

	count, err = Phendescs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
