package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testOrganismDbxrefs(t *testing.T) {
	t.Parallel()

	query := OrganismDbxrefs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testOrganismDbxrefsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismDbxref := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = organismDbxref.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := OrganismDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganismDbxrefsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismDbxref := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OrganismDbxrefs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := OrganismDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganismDbxrefsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismDbxref := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OrganismDbxrefSlice{organismDbxref}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := OrganismDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testOrganismDbxrefsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismDbxref := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := OrganismDbxrefExists(tx, organismDbxref.OrganismDbxrefID)
	if err != nil {
		t.Errorf("Unable to check if OrganismDbxref exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrganismDbxrefExistsG to return true, but got false.")
	}
}
func testOrganismDbxrefsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismDbxref := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	organismDbxrefFound, err := FindOrganismDbxref(tx, organismDbxref.OrganismDbxrefID)
	if err != nil {
		t.Error(err)
	}

	if organismDbxrefFound == nil {
		t.Error("want a record, got nil")
	}
}
func testOrganismDbxrefsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismDbxref := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OrganismDbxrefs(tx).Bind(organismDbxref); err != nil {
		t.Error(err)
	}
}

func testOrganismDbxrefsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismDbxref := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := OrganismDbxrefs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrganismDbxrefsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismDbxrefOne := &OrganismDbxref{}
	organismDbxrefTwo := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxrefOne, organismDbxrefDBTypes, false, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}
	if err = randomize.Struct(seed, organismDbxrefTwo, organismDbxrefDBTypes, false, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxrefOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = organismDbxrefTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := OrganismDbxrefs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrganismDbxrefsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	organismDbxrefOne := &OrganismDbxref{}
	organismDbxrefTwo := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxrefOne, organismDbxrefDBTypes, false, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}
	if err = randomize.Struct(seed, organismDbxrefTwo, organismDbxrefDBTypes, false, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxrefOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = organismDbxrefTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OrganismDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func organismDbxrefBeforeInsertHook(e boil.Executor, o *OrganismDbxref) error {
	*o = OrganismDbxref{}
	return nil
}

func organismDbxrefAfterInsertHook(e boil.Executor, o *OrganismDbxref) error {
	*o = OrganismDbxref{}
	return nil
}

func organismDbxrefAfterSelectHook(e boil.Executor, o *OrganismDbxref) error {
	*o = OrganismDbxref{}
	return nil
}

func organismDbxrefBeforeUpdateHook(e boil.Executor, o *OrganismDbxref) error {
	*o = OrganismDbxref{}
	return nil
}

func organismDbxrefAfterUpdateHook(e boil.Executor, o *OrganismDbxref) error {
	*o = OrganismDbxref{}
	return nil
}

func organismDbxrefBeforeDeleteHook(e boil.Executor, o *OrganismDbxref) error {
	*o = OrganismDbxref{}
	return nil
}

func organismDbxrefAfterDeleteHook(e boil.Executor, o *OrganismDbxref) error {
	*o = OrganismDbxref{}
	return nil
}

func organismDbxrefBeforeUpsertHook(e boil.Executor, o *OrganismDbxref) error {
	*o = OrganismDbxref{}
	return nil
}

func organismDbxrefAfterUpsertHook(e boil.Executor, o *OrganismDbxref) error {
	*o = OrganismDbxref{}
	return nil
}

func testOrganismDbxrefsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &OrganismDbxref{}
	o := &OrganismDbxref{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, organismDbxrefDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref object: %s", err)
	}

	AddOrganismDbxrefHook(boil.BeforeInsertHook, organismDbxrefBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	organismDbxrefBeforeInsertHooks = []OrganismDbxrefHook{}

	AddOrganismDbxrefHook(boil.AfterInsertHook, organismDbxrefAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	organismDbxrefAfterInsertHooks = []OrganismDbxrefHook{}

	AddOrganismDbxrefHook(boil.AfterSelectHook, organismDbxrefAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	organismDbxrefAfterSelectHooks = []OrganismDbxrefHook{}

	AddOrganismDbxrefHook(boil.BeforeUpdateHook, organismDbxrefBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	organismDbxrefBeforeUpdateHooks = []OrganismDbxrefHook{}

	AddOrganismDbxrefHook(boil.AfterUpdateHook, organismDbxrefAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	organismDbxrefAfterUpdateHooks = []OrganismDbxrefHook{}

	AddOrganismDbxrefHook(boil.BeforeDeleteHook, organismDbxrefBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	organismDbxrefBeforeDeleteHooks = []OrganismDbxrefHook{}

	AddOrganismDbxrefHook(boil.AfterDeleteHook, organismDbxrefAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	organismDbxrefAfterDeleteHooks = []OrganismDbxrefHook{}

	AddOrganismDbxrefHook(boil.BeforeUpsertHook, organismDbxrefBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	organismDbxrefBeforeUpsertHooks = []OrganismDbxrefHook{}

	AddOrganismDbxrefHook(boil.AfterUpsertHook, organismDbxrefAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	organismDbxrefAfterUpsertHooks = []OrganismDbxrefHook{}
}
func testOrganismDbxrefsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismDbxref := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OrganismDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganismDbxrefsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismDbxref := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxref.Insert(tx, organismDbxrefColumns...); err != nil {
		t.Error(err)
	}

	count, err := OrganismDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganismDbxrefToOneOrganismUsingOrganism(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local OrganismDbxref
	var foreign Organism

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, organismDbxrefDBTypes, true, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
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

	slice := OrganismDbxrefSlice{&local}
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

func testOrganismDbxrefToOneDbxrefUsingDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local OrganismDbxref
	var foreign Dbxref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, organismDbxrefDBTypes, true, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
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

	slice := OrganismDbxrefSlice{&local}
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

func testOrganismDbxrefToOneSetOpOrganismUsingOrganism(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OrganismDbxref
	var b, c Organism

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organismDbxrefDBTypes, false, strmangle.SetComplement(organismDbxrefPrimaryKeyColumns, organismDbxrefColumnsWithoutDefault)...); err != nil {
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

		if x.R.OrganismDbxref != &a {
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
func testOrganismDbxrefToOneSetOpDbxrefUsingDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OrganismDbxref
	var b, c Dbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organismDbxrefDBTypes, false, strmangle.SetComplement(organismDbxrefPrimaryKeyColumns, organismDbxrefColumnsWithoutDefault)...); err != nil {
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

		if x.R.OrganismDbxref != &a {
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
func testOrganismDbxrefsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismDbxref := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = organismDbxref.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testOrganismDbxrefsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismDbxref := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OrganismDbxrefSlice{organismDbxref}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testOrganismDbxrefsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismDbxref := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := OrganismDbxrefs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	organismDbxrefDBTypes = map[string]string{"DbxrefID": "integer", "OrganismDbxrefID": "integer", "OrganismID": "integer"}
	_                     = bytes.MinRead
)

func testOrganismDbxrefsUpdate(t *testing.T) {
	t.Parallel()

	if len(organismDbxrefColumns) == len(organismDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	organismDbxref := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OrganismDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	if err = organismDbxref.Update(tx); err != nil {
		t.Error(err)
	}
}

func testOrganismDbxrefsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(organismDbxrefColumns) == len(organismDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	organismDbxref := &OrganismDbxref{}
	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxref.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OrganismDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, organismDbxref, organismDbxrefDBTypes, true, organismDbxrefPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(organismDbxrefColumns, organismDbxrefPrimaryKeyColumns) {
		fields = organismDbxrefColumns
	} else {
		fields = strmangle.SetComplement(
			organismDbxrefColumns,
			organismDbxrefPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(organismDbxref))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := OrganismDbxrefSlice{organismDbxref}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testOrganismDbxrefsUpsert(t *testing.T) {
	t.Parallel()

	if len(organismDbxrefColumns) == len(organismDbxrefPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	organismDbxref := OrganismDbxref{}
	if err = randomize.Struct(seed, &organismDbxref, organismDbxrefDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismDbxref.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert OrganismDbxref: %s", err)
	}

	count, err := OrganismDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &organismDbxref, organismDbxrefDBTypes, false, organismDbxrefPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}

	if err = organismDbxref.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert OrganismDbxref: %s", err)
	}

	count, err = OrganismDbxrefs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
