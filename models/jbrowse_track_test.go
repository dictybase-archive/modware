package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testJbrowseTracks(t *testing.T) {
	t.Parallel()

	query := JbrowseTracks(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testJbrowseTracksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseTrack := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrack.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = jbrowseTrack.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := JbrowseTracks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJbrowseTracksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseTrack := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrack.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = JbrowseTracks(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := JbrowseTracks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJbrowseTracksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseTrack := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrack.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := JbrowseTrackSlice{jbrowseTrack}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := JbrowseTracks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testJbrowseTracksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseTrack := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrack.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := JbrowseTrackExists(tx, jbrowseTrack.JbrowseTrackID)
	if err != nil {
		t.Errorf("Unable to check if JbrowseTrack exists: %s", err)
	}
	if !e {
		t.Errorf("Expected JbrowseTrackExistsG to return true, but got false.")
	}
}
func testJbrowseTracksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseTrack := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrack.Insert(tx); err != nil {
		t.Error(err)
	}

	jbrowseTrackFound, err := FindJbrowseTrack(tx, jbrowseTrack.JbrowseTrackID)
	if err != nil {
		t.Error(err)
	}

	if jbrowseTrackFound == nil {
		t.Error("want a record, got nil")
	}
}
func testJbrowseTracksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseTrack := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrack.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = JbrowseTracks(tx).Bind(jbrowseTrack); err != nil {
		t.Error(err)
	}
}

func testJbrowseTracksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseTrack := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrack.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := JbrowseTracks(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testJbrowseTracksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseTrackOne := &JbrowseTrack{}
	jbrowseTrackTwo := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrackOne, jbrowseTrackDBTypes, false, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}
	if err = randomize.Struct(seed, jbrowseTrackTwo, jbrowseTrackDBTypes, false, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrackOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = jbrowseTrackTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := JbrowseTracks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testJbrowseTracksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	jbrowseTrackOne := &JbrowseTrack{}
	jbrowseTrackTwo := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrackOne, jbrowseTrackDBTypes, false, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}
	if err = randomize.Struct(seed, jbrowseTrackTwo, jbrowseTrackDBTypes, false, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrackOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = jbrowseTrackTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := JbrowseTracks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func jbrowseTrackBeforeInsertHook(e boil.Executor, o *JbrowseTrack) error {
	*o = JbrowseTrack{}
	return nil
}

func jbrowseTrackAfterInsertHook(e boil.Executor, o *JbrowseTrack) error {
	*o = JbrowseTrack{}
	return nil
}

func jbrowseTrackAfterSelectHook(e boil.Executor, o *JbrowseTrack) error {
	*o = JbrowseTrack{}
	return nil
}

func jbrowseTrackBeforeUpdateHook(e boil.Executor, o *JbrowseTrack) error {
	*o = JbrowseTrack{}
	return nil
}

func jbrowseTrackAfterUpdateHook(e boil.Executor, o *JbrowseTrack) error {
	*o = JbrowseTrack{}
	return nil
}

func jbrowseTrackBeforeDeleteHook(e boil.Executor, o *JbrowseTrack) error {
	*o = JbrowseTrack{}
	return nil
}

func jbrowseTrackAfterDeleteHook(e boil.Executor, o *JbrowseTrack) error {
	*o = JbrowseTrack{}
	return nil
}

func jbrowseTrackBeforeUpsertHook(e boil.Executor, o *JbrowseTrack) error {
	*o = JbrowseTrack{}
	return nil
}

func jbrowseTrackAfterUpsertHook(e boil.Executor, o *JbrowseTrack) error {
	*o = JbrowseTrack{}
	return nil
}

func testJbrowseTracksHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &JbrowseTrack{}
	o := &JbrowseTrack{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, jbrowseTrackDBTypes, false); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack object: %s", err)
	}

	AddJbrowseTrackHook(boil.BeforeInsertHook, jbrowseTrackBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	jbrowseTrackBeforeInsertHooks = []JbrowseTrackHook{}

	AddJbrowseTrackHook(boil.AfterInsertHook, jbrowseTrackAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	jbrowseTrackAfterInsertHooks = []JbrowseTrackHook{}

	AddJbrowseTrackHook(boil.AfterSelectHook, jbrowseTrackAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	jbrowseTrackAfterSelectHooks = []JbrowseTrackHook{}

	AddJbrowseTrackHook(boil.BeforeUpdateHook, jbrowseTrackBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	jbrowseTrackBeforeUpdateHooks = []JbrowseTrackHook{}

	AddJbrowseTrackHook(boil.AfterUpdateHook, jbrowseTrackAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	jbrowseTrackAfterUpdateHooks = []JbrowseTrackHook{}

	AddJbrowseTrackHook(boil.BeforeDeleteHook, jbrowseTrackBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	jbrowseTrackBeforeDeleteHooks = []JbrowseTrackHook{}

	AddJbrowseTrackHook(boil.AfterDeleteHook, jbrowseTrackAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	jbrowseTrackAfterDeleteHooks = []JbrowseTrackHook{}

	AddJbrowseTrackHook(boil.BeforeUpsertHook, jbrowseTrackBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	jbrowseTrackBeforeUpsertHooks = []JbrowseTrackHook{}

	AddJbrowseTrackHook(boil.AfterUpsertHook, jbrowseTrackAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	jbrowseTrackAfterUpsertHooks = []JbrowseTrackHook{}
}
func testJbrowseTracksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseTrack := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrack.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := JbrowseTracks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJbrowseTracksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseTrack := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrack.Insert(tx, jbrowseTrackColumns...); err != nil {
		t.Error(err)
	}

	count, err := JbrowseTracks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJbrowseTrackToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local JbrowseTrack
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, jbrowseTrackDBTypes, true, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	local.TypeID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.TypeID.Int = foreign.CvtermID
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

	slice := JbrowseTrackSlice{&local}
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

func testJbrowseTrackToOneJbrowseOrganismUsingJbrowseOrganism(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local JbrowseTrack
	var foreign JbrowseOrganism

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, jbrowseTrackDBTypes, true, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, jbrowseOrganismDBTypes, true, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.JbrowseOrganismID = foreign.JbrowseOrganismID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.JbrowseOrganism(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.JbrowseOrganismID != foreign.JbrowseOrganismID {
		t.Errorf("want: %v, got %v", foreign.JbrowseOrganismID, check.JbrowseOrganismID)
	}

	slice := JbrowseTrackSlice{&local}
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

func testJbrowseTrackToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a JbrowseTrack
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jbrowseTrackDBTypes, false, strmangle.SetComplement(jbrowseTrackPrimaryKeyColumns, jbrowseTrackColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeJbrowseTracks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TypeID.Int != x.CvtermID {
			t.Error("foreign key was wrong value", a.TypeID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TypeID.Int))
		reflect.Indirect(reflect.ValueOf(&a.TypeID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TypeID.Int != x.CvtermID {
			t.Error("foreign key was wrong value", a.TypeID.Int, x.CvtermID)
		}
	}
}

func testJbrowseTrackToOneRemoveOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a JbrowseTrack
	var b Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jbrowseTrackDBTypes, false, strmangle.SetComplement(jbrowseTrackPrimaryKeyColumns, jbrowseTrackColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvtermDBTypes, false, strmangle.SetComplement(cvtermPrimaryKeyColumns, cvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetType(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveType(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Type(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Type != nil {
		t.Error("R struct entry should be nil")
	}

	if a.TypeID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.TypeJbrowseTracks) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testJbrowseTrackToOneSetOpJbrowseOrganismUsingJbrowseOrganism(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a JbrowseTrack
	var b, c JbrowseOrganism

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jbrowseTrackDBTypes, false, strmangle.SetComplement(jbrowseTrackPrimaryKeyColumns, jbrowseTrackColumnsWithoutDefault)...); err != nil {
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

		if x.R.JbrowseTrack != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.JbrowseOrganismID != x.JbrowseOrganismID {
			t.Error("foreign key was wrong value", a.JbrowseOrganismID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.JbrowseOrganismID))
		reflect.Indirect(reflect.ValueOf(&a.JbrowseOrganismID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.JbrowseOrganismID != x.JbrowseOrganismID {
			t.Error("foreign key was wrong value", a.JbrowseOrganismID, x.JbrowseOrganismID)
		}
	}
}
func testJbrowseTracksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseTrack := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrack.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = jbrowseTrack.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testJbrowseTracksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseTrack := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrack.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := JbrowseTrackSlice{jbrowseTrack}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testJbrowseTracksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jbrowseTrack := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrack.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := JbrowseTracks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	jbrowseTrackDBTypes = map[string]string{"Configuration": "jsonb", "JbrowseOrganismID": "integer", "JbrowseTrackID": "integer", "TypeID": "integer"}
	_                   = bytes.MinRead
)

func testJbrowseTracksUpdate(t *testing.T) {
	t.Parallel()

	if len(jbrowseTrackColumns) == len(jbrowseTrackPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	jbrowseTrack := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrack.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := JbrowseTracks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true, jbrowseTrackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	if err = jbrowseTrack.Update(tx); err != nil {
		t.Error(err)
	}
}

func testJbrowseTracksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(jbrowseTrackColumns) == len(jbrowseTrackPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	jbrowseTrack := &JbrowseTrack{}
	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrack.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := JbrowseTracks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, jbrowseTrack, jbrowseTrackDBTypes, true, jbrowseTrackPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(jbrowseTrackColumns, jbrowseTrackPrimaryKeyColumns) {
		fields = jbrowseTrackColumns
	} else {
		fields = strmangle.SetComplement(
			jbrowseTrackColumns,
			jbrowseTrackPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(jbrowseTrack))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := JbrowseTrackSlice{jbrowseTrack}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testJbrowseTracksUpsert(t *testing.T) {
	t.Parallel()

	if len(jbrowseTrackColumns) == len(jbrowseTrackPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	jbrowseTrack := JbrowseTrack{}
	if err = randomize.Struct(seed, &jbrowseTrack, jbrowseTrackDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jbrowseTrack.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert JbrowseTrack: %s", err)
	}

	count, err := JbrowseTracks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &jbrowseTrack, jbrowseTrackDBTypes, false, jbrowseTrackPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize JbrowseTrack struct: %s", err)
	}

	if err = jbrowseTrack.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert JbrowseTrack: %s", err)
	}

	count, err = JbrowseTracks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
