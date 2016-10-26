package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testJbrowses(t *testing.T) {
	t.Parallel()

	query := Jbrowses(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testJbrowsesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowse := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowse.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = jbrowse.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Jbrowses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJbrowsesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowse := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowse.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Jbrowses(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Jbrowses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJbrowsesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowse := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowse.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := JbrowseSlice{jbrowse}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Jbrowses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testJbrowsesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowse := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true, jbrowseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowse.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := JbrowseExists(tx, jbrowse.JbrowseID)
	if err != nil {
		t.Errorf("Unable to check if Jbrowse exists: %s", err)
	}
	if !e {
		t.Errorf("Expected JbrowseExistsG to return true, but got false.")
	}
}
func testJbrowsesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowse := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true, jbrowseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowse.Insert(tx); err != nil {
		t.Error(err)
	}

	jbrowseFound, err := FindJbrowse(tx, jbrowse.JbrowseID)
	if err != nil {
		t.Error(err)
	}

	if jbrowseFound == nil {
		t.Error("want a record, got nil")
	}
}
func testJbrowsesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowse := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true, jbrowseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowse.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Jbrowses(tx).Bind(jbrowse); err != nil {
		t.Error(err)
	}
}

func testJbrowsesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowse := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true, jbrowseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowse.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Jbrowses(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testJbrowsesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseOne := &Jbrowse{}
	jbrowseTwo := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowseOne, jbrowseDBTypes, false, jbrowseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}
	if err = randomize.Struct(seed, jbrowseTwo, jbrowseDBTypes, false, jbrowseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = jbrowseTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Jbrowses(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testJbrowsesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	jbrowseOne := &Jbrowse{}
	jbrowseTwo := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowseOne, jbrowseDBTypes, false, jbrowseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}
	if err = randomize.Struct(seed, jbrowseTwo, jbrowseDBTypes, false, jbrowseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = jbrowseTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Jbrowses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func jbrowseBeforeInsertHook(e boil.Executor, o *Jbrowse) error {
	*o = Jbrowse{}
	return nil
}

func jbrowseAfterInsertHook(e boil.Executor, o *Jbrowse) error {
	*o = Jbrowse{}
	return nil
}

func jbrowseAfterSelectHook(e boil.Executor, o *Jbrowse) error {
	*o = Jbrowse{}
	return nil
}

func jbrowseBeforeUpdateHook(e boil.Executor, o *Jbrowse) error {
	*o = Jbrowse{}
	return nil
}

func jbrowseAfterUpdateHook(e boil.Executor, o *Jbrowse) error {
	*o = Jbrowse{}
	return nil
}

func jbrowseBeforeDeleteHook(e boil.Executor, o *Jbrowse) error {
	*o = Jbrowse{}
	return nil
}

func jbrowseAfterDeleteHook(e boil.Executor, o *Jbrowse) error {
	*o = Jbrowse{}
	return nil
}

func jbrowseBeforeUpsertHook(e boil.Executor, o *Jbrowse) error {
	*o = Jbrowse{}
	return nil
}

func jbrowseAfterUpsertHook(e boil.Executor, o *Jbrowse) error {
	*o = Jbrowse{}
	return nil
}

func testJbrowsesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Jbrowse{}
	o := &Jbrowse{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, jbrowseDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Jbrowse object: %s", err)
	}

	AddJbrowseHook(boil.BeforeInsertHook, jbrowseBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	jbrowseBeforeInsertHooks = []JbrowseHook{}

	AddJbrowseHook(boil.AfterInsertHook, jbrowseAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	jbrowseAfterInsertHooks = []JbrowseHook{}

	AddJbrowseHook(boil.AfterSelectHook, jbrowseAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	jbrowseAfterSelectHooks = []JbrowseHook{}

	AddJbrowseHook(boil.BeforeUpdateHook, jbrowseBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	jbrowseBeforeUpdateHooks = []JbrowseHook{}

	AddJbrowseHook(boil.AfterUpdateHook, jbrowseAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	jbrowseAfterUpdateHooks = []JbrowseHook{}

	AddJbrowseHook(boil.BeforeDeleteHook, jbrowseBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	jbrowseBeforeDeleteHooks = []JbrowseHook{}

	AddJbrowseHook(boil.AfterDeleteHook, jbrowseAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	jbrowseAfterDeleteHooks = []JbrowseHook{}

	AddJbrowseHook(boil.BeforeUpsertHook, jbrowseBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	jbrowseBeforeUpsertHooks = []JbrowseHook{}

	AddJbrowseHook(boil.AfterUpsertHook, jbrowseAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	jbrowseAfterUpsertHooks = []JbrowseHook{}
}
func testJbrowsesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowse := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true, jbrowseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowse.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Jbrowses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJbrowsesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowse := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowse.Insert(tx, jbrowseColumns...); err != nil {
		t.Error(err)
	}

	count, err := Jbrowses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJbrowseOneToOneJbrowseOrganismUsingJbrowseOrganism(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign JbrowseOrganism
	var local Jbrowse

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, jbrowseOrganismDBTypes, true, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, jbrowseDBTypes, true, jbrowseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.JbrowseID = local.JbrowseID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.JbrowseOrganism(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.JbrowseID != foreign.JbrowseID {
		t.Errorf("want: %v, got %v", foreign.JbrowseID, check.JbrowseID)
	}

	slice := JbrowseSlice{&local}
	if err = local.L.LoadJbrowseOrganism(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.JbrowseOrganism == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.JbrowseOrganism = nil
	if err = local.L.LoadJbrowseOrganism(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.JbrowseOrganism == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testJbrowseOneToOneSetOpJbrowseOrganismUsingJbrowseOrganism(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Jbrowse
	var b, c JbrowseOrganism

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jbrowseDBTypes, false, strmangle.SetComplement(jbrowsePrimaryKeyColumns, jbrowseColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, jbrowseOrganismDBTypes, false, strmangle.SetComplement(jbrowseOrganismPrimaryKeyColumns, jbrowseOrganismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, jbrowseOrganismDBTypes, false, strmangle.SetComplement(jbrowseOrganismPrimaryKeyColumns, jbrowseOrganismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*JbrowseOrganism{&b, &c} {
		err = a.SetJbrowseOrganism(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.JbrowseOrganism != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Jbrowse != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.JbrowseID != x.JbrowseID {
			t.Error("foreign key was wrong value", a.JbrowseID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.JbrowseID))
		reflect.Indirect(reflect.ValueOf(&x.JbrowseID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.JbrowseID != x.JbrowseID {
			t.Error("foreign key was wrong value", a.JbrowseID, x.JbrowseID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testJbrowsesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowse := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true, jbrowseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowse.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = jbrowse.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testJbrowsesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowse := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true, jbrowseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowse.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := JbrowseSlice{jbrowse}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testJbrowsesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowse := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true, jbrowseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowse.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Jbrowses(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	jbrowseDBTypes = map[string]string{"Configuration": "jsonb", "JbrowseID": "integer", "Name": "character varying"}
	_              = bytes.MinRead
)

func testJbrowsesUpdate(t *testing.T) {
	t.Parallel()

	if len(jbrowseColumns) == len(jbrowsePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	jbrowse := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowse.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Jbrowses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true, jbrowseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	if err = jbrowse.Update(tx); err != nil {
		t.Error(err)
	}
}

func testJbrowsesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(jbrowseColumns) == len(jbrowsePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	jbrowse := &Jbrowse{}
	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowse.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Jbrowses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, jbrowse, jbrowseDBTypes, true, jbrowsePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(jbrowseColumns, jbrowsePrimaryKeyColumns) {
		fields = jbrowseColumns
	} else {
		fields = strmangle.SetComplement(
			jbrowseColumns,
			jbrowsePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(jbrowse))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := JbrowseSlice{jbrowse}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testJbrowsesUpsert(t *testing.T) {
	t.Parallel()

	if len(jbrowseColumns) == len(jbrowsePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	jbrowse := Jbrowse{}
	if err = randomize.Struct(seed, &jbrowse, jbrowseDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowse.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Jbrowse: %s", err)
	}

	count, err := Jbrowses(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &jbrowse, jbrowseDBTypes, false, jbrowsePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	if err = jbrowse.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Jbrowse: %s", err)
	}

	count, err = Jbrowses(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
