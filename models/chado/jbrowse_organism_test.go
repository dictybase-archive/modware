package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testJbrowseOrganisms(t *testing.T) {
	t.Parallel()

	query := JbrowseOrganisms(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testJbrowseOrganismsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseOrganism := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganism.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = jbrowseOrganism.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := JbrowseOrganisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJbrowseOrganismsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseOrganism := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganism.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = JbrowseOrganisms(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := JbrowseOrganisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJbrowseOrganismsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseOrganism := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganism.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := JbrowseOrganismSlice{jbrowseOrganism}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := JbrowseOrganisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testJbrowseOrganismsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseOrganism := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganism.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := JbrowseOrganismExists(tx, jbrowseOrganism.JbrowseOrganismID)
	if err != nil {
		t.Errorf("Unable to check if JbrowseOrganism exists: %s", err)
	}
	if !e {
		t.Errorf("Expected JbrowseOrganismExistsG to return true, but got false.")
	}
}
func testJbrowseOrganismsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseOrganism := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganism.Insert(tx); err != nil {
		t.Error(err)
	}

	jbrowseOrganismFound, err := FindJbrowseOrganism(tx, jbrowseOrganism.JbrowseOrganismID)
	if err != nil {
		t.Error(err)
	}

	if jbrowseOrganismFound == nil {
		t.Error("want a record, got nil")
	}
}
func testJbrowseOrganismsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseOrganism := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganism.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = JbrowseOrganisms(tx).Bind(jbrowseOrganism); err != nil {
		t.Error(err)
	}
}

func testJbrowseOrganismsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseOrganism := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganism.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := JbrowseOrganisms(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testJbrowseOrganismsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseOrganismOne := &JbrowseOrganism{}
	jbrowseOrganismTwo := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganismOne, jbrowseOrganismDBTypes, false, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}
	if err = randomize.Struct(seed, jbrowseOrganismTwo, jbrowseOrganismDBTypes, false, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganismOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = jbrowseOrganismTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := JbrowseOrganisms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testJbrowseOrganismsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	jbrowseOrganismOne := &JbrowseOrganism{}
	jbrowseOrganismTwo := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganismOne, jbrowseOrganismDBTypes, false, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}
	if err = randomize.Struct(seed, jbrowseOrganismTwo, jbrowseOrganismDBTypes, false, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganismOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = jbrowseOrganismTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := JbrowseOrganisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func jbrowseOrganismBeforeInsertHook(e boil.Executor, o *JbrowseOrganism) error {
	*o = JbrowseOrganism{}
	return nil
}

func jbrowseOrganismAfterInsertHook(e boil.Executor, o *JbrowseOrganism) error {
	*o = JbrowseOrganism{}
	return nil
}

func jbrowseOrganismAfterSelectHook(e boil.Executor, o *JbrowseOrganism) error {
	*o = JbrowseOrganism{}
	return nil
}

func jbrowseOrganismBeforeUpdateHook(e boil.Executor, o *JbrowseOrganism) error {
	*o = JbrowseOrganism{}
	return nil
}

func jbrowseOrganismAfterUpdateHook(e boil.Executor, o *JbrowseOrganism) error {
	*o = JbrowseOrganism{}
	return nil
}

func jbrowseOrganismBeforeDeleteHook(e boil.Executor, o *JbrowseOrganism) error {
	*o = JbrowseOrganism{}
	return nil
}

func jbrowseOrganismAfterDeleteHook(e boil.Executor, o *JbrowseOrganism) error {
	*o = JbrowseOrganism{}
	return nil
}

func jbrowseOrganismBeforeUpsertHook(e boil.Executor, o *JbrowseOrganism) error {
	*o = JbrowseOrganism{}
	return nil
}

func jbrowseOrganismAfterUpsertHook(e boil.Executor, o *JbrowseOrganism) error {
	*o = JbrowseOrganism{}
	return nil
}

func testJbrowseOrganismsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &JbrowseOrganism{}
	o := &JbrowseOrganism{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, jbrowseOrganismDBTypes, false); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism object: %s", err)
	}

	AddJbrowseOrganismHook(boil.BeforeInsertHook, jbrowseOrganismBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	jbrowseOrganismBeforeInsertHooks = []JbrowseOrganismHook{}

	AddJbrowseOrganismHook(boil.AfterInsertHook, jbrowseOrganismAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	jbrowseOrganismAfterInsertHooks = []JbrowseOrganismHook{}

	AddJbrowseOrganismHook(boil.AfterSelectHook, jbrowseOrganismAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	jbrowseOrganismAfterSelectHooks = []JbrowseOrganismHook{}

	AddJbrowseOrganismHook(boil.BeforeUpdateHook, jbrowseOrganismBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	jbrowseOrganismBeforeUpdateHooks = []JbrowseOrganismHook{}

	AddJbrowseOrganismHook(boil.AfterUpdateHook, jbrowseOrganismAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	jbrowseOrganismAfterUpdateHooks = []JbrowseOrganismHook{}

	AddJbrowseOrganismHook(boil.BeforeDeleteHook, jbrowseOrganismBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	jbrowseOrganismBeforeDeleteHooks = []JbrowseOrganismHook{}

	AddJbrowseOrganismHook(boil.AfterDeleteHook, jbrowseOrganismAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	jbrowseOrganismAfterDeleteHooks = []JbrowseOrganismHook{}

	AddJbrowseOrganismHook(boil.BeforeUpsertHook, jbrowseOrganismBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	jbrowseOrganismBeforeUpsertHooks = []JbrowseOrganismHook{}

	AddJbrowseOrganismHook(boil.AfterUpsertHook, jbrowseOrganismAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	jbrowseOrganismAfterUpsertHooks = []JbrowseOrganismHook{}
}
func testJbrowseOrganismsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseOrganism := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganism.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := JbrowseOrganisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJbrowseOrganismsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseOrganism := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganism.Insert(tx, jbrowseOrganismColumns...); err != nil {
		t.Error(err)
	}

	count, err := JbrowseOrganisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJbrowseOrganismOneToOneJbrowseTrackUsingJbrowseTrack(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign JbrowseTrack
	var local JbrowseOrganism

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, jbrowseTrackDBTypes, true, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, jbrowseOrganismDBTypes, true, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.JbrowseOrganismID = local.JbrowseOrganismID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.JbrowseTrack(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.JbrowseOrganismID != foreign.JbrowseOrganismID {
		t.Errorf("want: %v, got %v", foreign.JbrowseOrganismID, check.JbrowseOrganismID)
	}

	slice := JbrowseOrganismSlice{&local}
	if err = local.L.LoadJbrowseTrack(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.JbrowseTrack == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.JbrowseTrack = nil
	if err = local.L.LoadJbrowseTrack(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.JbrowseTrack == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testJbrowseOrganismOneToOneSetOpJbrowseTrackUsingJbrowseTrack(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a JbrowseOrganism
	var b, c JbrowseTrack

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jbrowseOrganismDBTypes, false, strmangle.SetComplement(jbrowseOrganismPrimaryKeyColumns, jbrowseOrganismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, jbrowseTrackDBTypes, false, strmangle.SetComplement(jbrowseTrackPrimaryKeyColumns, jbrowseTrackColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, jbrowseTrackDBTypes, false, strmangle.SetComplement(jbrowseTrackPrimaryKeyColumns, jbrowseTrackColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*JbrowseTrack{&b, &c} {
		err = a.SetJbrowseTrack(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.JbrowseTrack != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.JbrowseOrganism != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.JbrowseOrganismID != x.JbrowseOrganismID {
			t.Error("foreign key was wrong value", a.JbrowseOrganismID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.JbrowseOrganismID))
		reflect.Indirect(reflect.ValueOf(&x.JbrowseOrganismID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.JbrowseOrganismID != x.JbrowseOrganismID {
			t.Error("foreign key was wrong value", a.JbrowseOrganismID, x.JbrowseOrganismID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testJbrowseOrganismToOneOrganismUsingOrganism(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local JbrowseOrganism
	var foreign Organism

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, jbrowseOrganismDBTypes, true, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
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

	slice := JbrowseOrganismSlice{&local}
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

func testJbrowseOrganismToOneJbrowseUsingJbrowse(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local JbrowseOrganism
	var foreign Jbrowse

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, jbrowseOrganismDBTypes, true, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, jbrowseDBTypes, true, jbrowseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jbrowse struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.JbrowseID = foreign.JbrowseID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Jbrowse(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.JbrowseID != foreign.JbrowseID {
		t.Errorf("want: %v, got %v", foreign.JbrowseID, check.JbrowseID)
	}

	slice := JbrowseOrganismSlice{&local}
	if err = local.L.LoadJbrowse(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Jbrowse == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Jbrowse = nil
	if err = local.L.LoadJbrowse(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Jbrowse == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testJbrowseOrganismToOneSetOpOrganismUsingOrganism(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a JbrowseOrganism
	var b, c Organism

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jbrowseOrganismDBTypes, false, strmangle.SetComplement(jbrowseOrganismPrimaryKeyColumns, jbrowseOrganismColumnsWithoutDefault)...); err != nil {
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

		if x.R.JbrowseOrganism != &a {
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
func testJbrowseOrganismToOneSetOpJbrowseUsingJbrowse(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a JbrowseOrganism
	var b, c Jbrowse

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jbrowseOrganismDBTypes, false, strmangle.SetComplement(jbrowseOrganismPrimaryKeyColumns, jbrowseOrganismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, jbrowseDBTypes, false, strmangle.SetComplement(jbrowsePrimaryKeyColumns, jbrowseColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, jbrowseDBTypes, false, strmangle.SetComplement(jbrowsePrimaryKeyColumns, jbrowseColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Jbrowse{&b, &c} {
		err = a.SetJbrowse(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Jbrowse != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.JbrowseOrganism != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.JbrowseID != x.JbrowseID {
			t.Error("foreign key was wrong value", a.JbrowseID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.JbrowseID))
		reflect.Indirect(reflect.ValueOf(&a.JbrowseID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.JbrowseID != x.JbrowseID {
			t.Error("foreign key was wrong value", a.JbrowseID, x.JbrowseID)
		}
	}
}
func testJbrowseOrganismsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseOrganism := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganism.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = jbrowseOrganism.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testJbrowseOrganismsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseOrganism := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganism.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := JbrowseOrganismSlice{jbrowseOrganism}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testJbrowseOrganismsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseOrganism := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganism.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := JbrowseOrganisms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	jbrowseOrganismDBTypes = map[string]string{"Dataset": "character varying", "JbrowseID": "integer", "JbrowseOrganismID": "integer", "OrganismID": "integer"}
	_                      = bytes.MinRead
)

func testJbrowseOrganismsUpdate(t *testing.T) {
	t.Parallel()

	if len(jbrowseOrganismColumns) == len(jbrowseOrganismPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	jbrowseOrganism := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganism.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := JbrowseOrganisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	if err = jbrowseOrganism.Update(tx); err != nil {
		t.Error(err)
	}
}

func testJbrowseOrganismsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(jbrowseOrganismColumns) == len(jbrowseOrganismPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	jbrowseOrganism := &JbrowseOrganism{}
	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganism.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := JbrowseOrganisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, jbrowseOrganism, jbrowseOrganismDBTypes, true, jbrowseOrganismPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(jbrowseOrganismColumns, jbrowseOrganismPrimaryKeyColumns) {
		fields = jbrowseOrganismColumns
	} else {
		fields = strmangle.SetComplement(
			jbrowseOrganismColumns,
			jbrowseOrganismPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(jbrowseOrganism))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := JbrowseOrganismSlice{jbrowseOrganism}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testJbrowseOrganismsUpsert(t *testing.T) {
	t.Parallel()

	if len(jbrowseOrganismColumns) == len(jbrowseOrganismPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	jbrowseOrganism := JbrowseOrganism{}
	if err = randomize.Struct(seed, &jbrowseOrganism, jbrowseOrganismDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseOrganism.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert JbrowseOrganism: %s", err)
	}

	count, err := JbrowseOrganisms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &jbrowseOrganism, jbrowseOrganismDBTypes, false, jbrowseOrganismPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	if err = jbrowseOrganism.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert JbrowseOrganism: %s", err)
	}

	count, err = JbrowseOrganisms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
