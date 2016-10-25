package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testPubRelationships(t *testing.T) {
	t.Parallel()

	query := PubRelationships(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testPubRelationshipsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubRelationship := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = pubRelationship.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := PubRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPubRelationshipsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubRelationship := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = PubRelationships(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := PubRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPubRelationshipsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubRelationship := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PubRelationshipSlice{pubRelationship}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := PubRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testPubRelationshipsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubRelationship := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := PubRelationshipExists(tx, pubRelationship.PubRelationshipID)
	if err != nil {
		t.Errorf("Unable to check if PubRelationship exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PubRelationshipExistsG to return true, but got false.")
	}
}
func testPubRelationshipsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubRelationship := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	pubRelationshipFound, err := FindPubRelationship(tx, pubRelationship.PubRelationshipID)
	if err != nil {
		t.Error(err)
	}

	if pubRelationshipFound == nil {
		t.Error("want a record, got nil")
	}
}
func testPubRelationshipsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubRelationship := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = PubRelationships(tx).Bind(pubRelationship); err != nil {
		t.Error(err)
	}
}

func testPubRelationshipsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubRelationship := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := PubRelationships(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPubRelationshipsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubRelationshipOne := &PubRelationship{}
	pubRelationshipTwo := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationshipOne, pubRelationshipDBTypes, false, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}
	if err = randomize.Struct(seed, pubRelationshipTwo, pubRelationshipDBTypes, false, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationshipOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = pubRelationshipTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := PubRelationships(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPubRelationshipsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	pubRelationshipOne := &PubRelationship{}
	pubRelationshipTwo := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationshipOne, pubRelationshipDBTypes, false, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}
	if err = randomize.Struct(seed, pubRelationshipTwo, pubRelationshipDBTypes, false, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationshipOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = pubRelationshipTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PubRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func pubRelationshipBeforeInsertHook(e boil.Executor, o *PubRelationship) error {
	*o = PubRelationship{}
	return nil
}

func pubRelationshipAfterInsertHook(e boil.Executor, o *PubRelationship) error {
	*o = PubRelationship{}
	return nil
}

func pubRelationshipAfterSelectHook(e boil.Executor, o *PubRelationship) error {
	*o = PubRelationship{}
	return nil
}

func pubRelationshipBeforeUpdateHook(e boil.Executor, o *PubRelationship) error {
	*o = PubRelationship{}
	return nil
}

func pubRelationshipAfterUpdateHook(e boil.Executor, o *PubRelationship) error {
	*o = PubRelationship{}
	return nil
}

func pubRelationshipBeforeDeleteHook(e boil.Executor, o *PubRelationship) error {
	*o = PubRelationship{}
	return nil
}

func pubRelationshipAfterDeleteHook(e boil.Executor, o *PubRelationship) error {
	*o = PubRelationship{}
	return nil
}

func pubRelationshipBeforeUpsertHook(e boil.Executor, o *PubRelationship) error {
	*o = PubRelationship{}
	return nil
}

func pubRelationshipAfterUpsertHook(e boil.Executor, o *PubRelationship) error {
	*o = PubRelationship{}
	return nil
}

func testPubRelationshipsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &PubRelationship{}
	o := &PubRelationship{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, pubRelationshipDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PubRelationship object: %s", err)
	}

	AddPubRelationshipHook(boil.BeforeInsertHook, pubRelationshipBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	pubRelationshipBeforeInsertHooks = []PubRelationshipHook{}

	AddPubRelationshipHook(boil.AfterInsertHook, pubRelationshipAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	pubRelationshipAfterInsertHooks = []PubRelationshipHook{}

	AddPubRelationshipHook(boil.AfterSelectHook, pubRelationshipAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	pubRelationshipAfterSelectHooks = []PubRelationshipHook{}

	AddPubRelationshipHook(boil.BeforeUpdateHook, pubRelationshipBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	pubRelationshipBeforeUpdateHooks = []PubRelationshipHook{}

	AddPubRelationshipHook(boil.AfterUpdateHook, pubRelationshipAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	pubRelationshipAfterUpdateHooks = []PubRelationshipHook{}

	AddPubRelationshipHook(boil.BeforeDeleteHook, pubRelationshipBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	pubRelationshipBeforeDeleteHooks = []PubRelationshipHook{}

	AddPubRelationshipHook(boil.AfterDeleteHook, pubRelationshipAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	pubRelationshipAfterDeleteHooks = []PubRelationshipHook{}

	AddPubRelationshipHook(boil.BeforeUpsertHook, pubRelationshipBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	pubRelationshipBeforeUpsertHooks = []PubRelationshipHook{}

	AddPubRelationshipHook(boil.AfterUpsertHook, pubRelationshipAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	pubRelationshipAfterUpsertHooks = []PubRelationshipHook{}
}
func testPubRelationshipsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubRelationship := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PubRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPubRelationshipsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubRelationship := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationship.Insert(tx, pubRelationshipColumns...); err != nil {
		t.Error(err)
	}

	count, err := PubRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPubRelationshipToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PubRelationship
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, pubRelationshipDBTypes, true, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
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

	slice := PubRelationshipSlice{&local}
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

func testPubRelationshipToOnePubUsingObject(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PubRelationship
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, pubRelationshipDBTypes, true, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ObjectID = foreign.PubID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Object(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubRelationshipSlice{&local}
	if err = local.L.LoadObject(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Object == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Object = nil
	if err = local.L.LoadObject(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Object == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubRelationshipToOnePubUsingSubject(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local PubRelationship
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, pubRelationshipDBTypes, true, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.SubjectID = foreign.PubID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Subject(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubRelationshipSlice{&local}
	if err = local.L.LoadSubject(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Subject == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Subject = nil
	if err = local.L.LoadSubject(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Subject == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubRelationshipToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PubRelationship
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubRelationshipDBTypes, false, strmangle.SetComplement(pubRelationshipPrimaryKeyColumns, pubRelationshipColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypePubRelationship != &a {
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
func testPubRelationshipToOneSetOpPubUsingObject(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PubRelationship
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubRelationshipDBTypes, false, strmangle.SetComplement(pubRelationshipPrimaryKeyColumns, pubRelationshipColumnsWithoutDefault)...); err != nil {
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
		err = a.SetObject(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Object != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ObjectPubRelationship != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ObjectID != x.PubID {
			t.Error("foreign key was wrong value", a.ObjectID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ObjectID))
		reflect.Indirect(reflect.ValueOf(&a.ObjectID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ObjectID != x.PubID {
			t.Error("foreign key was wrong value", a.ObjectID, x.PubID)
		}
	}
}
func testPubRelationshipToOneSetOpPubUsingSubject(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a PubRelationship
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubRelationshipDBTypes, false, strmangle.SetComplement(pubRelationshipPrimaryKeyColumns, pubRelationshipColumnsWithoutDefault)...); err != nil {
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
		err = a.SetSubject(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Subject != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SubjectPubRelationship != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SubjectID != x.PubID {
			t.Error("foreign key was wrong value", a.SubjectID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SubjectID))
		reflect.Indirect(reflect.ValueOf(&a.SubjectID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SubjectID != x.PubID {
			t.Error("foreign key was wrong value", a.SubjectID, x.PubID)
		}
	}
}
func testPubRelationshipsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubRelationship := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = pubRelationship.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPubRelationshipsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubRelationship := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PubRelationshipSlice{pubRelationship}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testPubRelationshipsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubRelationship := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := PubRelationships(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	pubRelationshipDBTypes = map[string]string{"ObjectID": "integer", "PubRelationshipID": "integer", "SubjectID": "integer", "TypeID": "integer"}
	_                      = bytes.MinRead
)

func testPubRelationshipsUpdate(t *testing.T) {
	t.Parallel()

	if len(pubRelationshipColumns) == len(pubRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	pubRelationship := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PubRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	if err = pubRelationship.Update(tx); err != nil {
		t.Error(err)
	}
}

func testPubRelationshipsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(pubRelationshipColumns) == len(pubRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	pubRelationship := &PubRelationship{}
	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := PubRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, pubRelationship, pubRelationshipDBTypes, true, pubRelationshipPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(pubRelationshipColumns, pubRelationshipPrimaryKeyColumns) {
		fields = pubRelationshipColumns
	} else {
		fields = strmangle.SetComplement(
			pubRelationshipColumns,
			pubRelationshipPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(pubRelationship))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := PubRelationshipSlice{pubRelationship}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testPubRelationshipsUpsert(t *testing.T) {
	t.Parallel()

	if len(pubRelationshipColumns) == len(pubRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	pubRelationship := PubRelationship{}
	if err = randomize.Struct(seed, &pubRelationship, pubRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubRelationship.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert PubRelationship: %s", err)
	}

	count, err := PubRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &pubRelationship, pubRelationshipDBTypes, false, pubRelationshipPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}

	if err = pubRelationship.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert PubRelationship: %s", err)
	}

	count, err = PubRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
