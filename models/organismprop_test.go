package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testOrganismprops(t *testing.T) {
	t.Parallel()

	query := Organismprops(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testOrganismpropsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismprop := &Organismprop{}
	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = organismprop.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Organismprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganismpropsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismprop := &Organismprop{}
	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Organismprops(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Organismprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganismpropsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismprop := &Organismprop{}
	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OrganismpropSlice{organismprop}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Organismprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testOrganismpropsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismprop := &Organismprop{}
	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismprop.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := OrganismpropExists(tx, organismprop.OrganismpropID)
	if err != nil {
		t.Errorf("Unable to check if Organismprop exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrganismpropExistsG to return true, but got false.")
	}
}
func testOrganismpropsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismprop := &Organismprop{}
	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismprop.Insert(tx); err != nil {
		t.Error(err)
	}

	organismpropFound, err := FindOrganismprop(tx, organismprop.OrganismpropID)
	if err != nil {
		t.Error(err)
	}

	if organismpropFound == nil {
		t.Error("want a record, got nil")
	}
}
func testOrganismpropsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismprop := &Organismprop{}
	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Organismprops(tx).Bind(organismprop); err != nil {
		t.Error(err)
	}
}

func testOrganismpropsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismprop := &Organismprop{}
	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Organismprops(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrganismpropsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismpropOne := &Organismprop{}
	organismpropTwo := &Organismprop{}
	if err = randomize.Struct(seed, organismpropOne, organismpropDBTypes, false, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}
	if err = randomize.Struct(seed, organismpropTwo, organismpropDBTypes, false, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = organismpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Organismprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrganismpropsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	organismpropOne := &Organismprop{}
	organismpropTwo := &Organismprop{}
	if err = randomize.Struct(seed, organismpropOne, organismpropDBTypes, false, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}
	if err = randomize.Struct(seed, organismpropTwo, organismpropDBTypes, false, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismpropOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = organismpropTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organismprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func organismpropBeforeInsertHook(e boil.Executor, o *Organismprop) error {
	*o = Organismprop{}
	return nil
}

func organismpropAfterInsertHook(e boil.Executor, o *Organismprop) error {
	*o = Organismprop{}
	return nil
}

func organismpropAfterSelectHook(e boil.Executor, o *Organismprop) error {
	*o = Organismprop{}
	return nil
}

func organismpropBeforeUpdateHook(e boil.Executor, o *Organismprop) error {
	*o = Organismprop{}
	return nil
}

func organismpropAfterUpdateHook(e boil.Executor, o *Organismprop) error {
	*o = Organismprop{}
	return nil
}

func organismpropBeforeDeleteHook(e boil.Executor, o *Organismprop) error {
	*o = Organismprop{}
	return nil
}

func organismpropAfterDeleteHook(e boil.Executor, o *Organismprop) error {
	*o = Organismprop{}
	return nil
}

func organismpropBeforeUpsertHook(e boil.Executor, o *Organismprop) error {
	*o = Organismprop{}
	return nil
}

func organismpropAfterUpsertHook(e boil.Executor, o *Organismprop) error {
	*o = Organismprop{}
	return nil
}

func testOrganismpropsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Organismprop{}
	o := &Organismprop{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, organismpropDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Organismprop object: %s", err)
	}

	AddOrganismpropHook(boil.BeforeInsertHook, organismpropBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	organismpropBeforeInsertHooks = []OrganismpropHook{}

	AddOrganismpropHook(boil.AfterInsertHook, organismpropAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	organismpropAfterInsertHooks = []OrganismpropHook{}

	AddOrganismpropHook(boil.AfterSelectHook, organismpropAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	organismpropAfterSelectHooks = []OrganismpropHook{}

	AddOrganismpropHook(boil.BeforeUpdateHook, organismpropBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	organismpropBeforeUpdateHooks = []OrganismpropHook{}

	AddOrganismpropHook(boil.AfterUpdateHook, organismpropAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	organismpropAfterUpdateHooks = []OrganismpropHook{}

	AddOrganismpropHook(boil.BeforeDeleteHook, organismpropBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	organismpropBeforeDeleteHooks = []OrganismpropHook{}

	AddOrganismpropHook(boil.AfterDeleteHook, organismpropAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	organismpropAfterDeleteHooks = []OrganismpropHook{}

	AddOrganismpropHook(boil.BeforeUpsertHook, organismpropBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	organismpropBeforeUpsertHooks = []OrganismpropHook{}

	AddOrganismpropHook(boil.AfterUpsertHook, organismpropAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	organismpropAfterUpsertHooks = []OrganismpropHook{}
}
func testOrganismpropsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismprop := &Organismprop{}
	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organismprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganismpropsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismprop := &Organismprop{}
	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismprop.Insert(tx, organismpropColumns...); err != nil {
		t.Error(err)
	}

	count, err := Organismprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganismpropToOneOrganismUsingOrganism(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Organismprop
	var foreign Organism

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, organismpropDBTypes, true, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.OrganismID = foreign.OrganismID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Organism(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.OrganismID != foreign.OrganismID {
		t.Errorf("want: %v, got %v", foreign.OrganismID, check.OrganismID)
	}

	slice := OrganismpropSlice{&local}
	if err = local.L.LoadOrganism(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Organism == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Organism = nil
	if err = local.L.LoadOrganism(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Organism == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOrganismpropToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Organismprop
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, organismpropDBTypes, true, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
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

	slice := OrganismpropSlice{&local}
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

func testOrganismpropToOneSetOpOrganismUsingOrganism(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organismprop
	var b, c Organism

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organismpropDBTypes, false, strmangle.SetComplement(organismpropPrimaryKeyColumns, organismpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organismDBTypes, false, strmangle.SetComplement(organismPrimaryKeyColumns, organismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, organismDBTypes, false, strmangle.SetComplement(organismPrimaryKeyColumns, organismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Organism{&b, &c} {
		err = a.SetOrganism(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Organism != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Organismprop != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OrganismID != x.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OrganismID))
		reflect.Indirect(reflect.ValueOf(&a.OrganismID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrganismID != x.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID, x.OrganismID)
		}
	}
}
func testOrganismpropToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organismprop
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organismpropDBTypes, false, strmangle.SetComplement(organismpropPrimaryKeyColumns, organismpropColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeOrganismprop != &a {
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
func testOrganismpropsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismprop := &Organismprop{}
	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismprop.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = organismprop.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testOrganismpropsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismprop := &Organismprop{}
	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OrganismpropSlice{organismprop}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testOrganismpropsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismprop := &Organismprop{}
	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismprop.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Organismprops(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	organismpropDBTypes = map[string]string{"OrganismID": "integer", "OrganismpropID": "integer", "Rank": "integer", "TypeID": "integer", "Value": "text"}
	_                   = bytes.MinRead
)

func testOrganismpropsUpdate(t *testing.T) {
	t.Parallel()

	if len(organismpropColumns) == len(organismpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	organismprop := &Organismprop{}
	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organismprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	if err = organismprop.Update(tx); err != nil {
		t.Error(err)
	}
}

func testOrganismpropsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(organismpropColumns) == len(organismpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	organismprop := &Organismprop{}
	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismprop.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organismprops(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, organismprop, organismpropDBTypes, true, organismpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(organismpropColumns, organismpropPrimaryKeyColumns) {
		fields = organismpropColumns
	} else {
		fields = strmangle.SetComplement(
			organismpropColumns,
			organismpropPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(organismprop))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := OrganismpropSlice{organismprop}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testOrganismpropsUpsert(t *testing.T) {
	t.Parallel()

	if len(organismpropColumns) == len(organismpropPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	organismprop := Organismprop{}
	if err = randomize.Struct(seed, &organismprop, organismpropDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismprop.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Organismprop: %s", err)
	}

	count, err := Organismprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &organismprop, organismpropDBTypes, false, organismpropPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}

	if err = organismprop.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Organismprop: %s", err)
	}

	count, err = Organismprops(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
