package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testUserRelationships(t *testing.T) {
	t.Parallel()

	query := UserRelationships(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testUserRelationshipsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userRelationship := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = userRelationship.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := UserRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserRelationshipsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userRelationship := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = UserRelationships(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := UserRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserRelationshipsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userRelationship := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UserRelationshipSlice{userRelationship}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := UserRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testUserRelationshipsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userRelationship := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := UserRelationshipExists(tx, userRelationship.UserRelationshipID)
	if err != nil {
		t.Errorf("Unable to check if UserRelationship exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserRelationshipExistsG to return true, but got false.")
	}
}
func testUserRelationshipsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userRelationship := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	userRelationshipFound, err := FindUserRelationship(tx, userRelationship.UserRelationshipID)
	if err != nil {
		t.Error(err)
	}

	if userRelationshipFound == nil {
		t.Error("want a record, got nil")
	}
}
func testUserRelationshipsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userRelationship := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = UserRelationships(tx).Bind(userRelationship); err != nil {
		t.Error(err)
	}
}

func testUserRelationshipsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userRelationship := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := UserRelationships(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserRelationshipsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userRelationshipOne := &UserRelationship{}
	userRelationshipTwo := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationshipOne, userRelationshipDBTypes, false, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}
	if err = randomize.Struct(seed, userRelationshipTwo, userRelationshipDBTypes, false, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationshipOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = userRelationshipTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := UserRelationships(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserRelationshipsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userRelationshipOne := &UserRelationship{}
	userRelationshipTwo := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationshipOne, userRelationshipDBTypes, false, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}
	if err = randomize.Struct(seed, userRelationshipTwo, userRelationshipDBTypes, false, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationshipOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = userRelationshipTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UserRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func userRelationshipBeforeInsertHook(e boil.Executor, o *UserRelationship) error {
	*o = UserRelationship{}
	return nil
}

func userRelationshipAfterInsertHook(e boil.Executor, o *UserRelationship) error {
	*o = UserRelationship{}
	return nil
}

func userRelationshipAfterSelectHook(e boil.Executor, o *UserRelationship) error {
	*o = UserRelationship{}
	return nil
}

func userRelationshipBeforeUpdateHook(e boil.Executor, o *UserRelationship) error {
	*o = UserRelationship{}
	return nil
}

func userRelationshipAfterUpdateHook(e boil.Executor, o *UserRelationship) error {
	*o = UserRelationship{}
	return nil
}

func userRelationshipBeforeDeleteHook(e boil.Executor, o *UserRelationship) error {
	*o = UserRelationship{}
	return nil
}

func userRelationshipAfterDeleteHook(e boil.Executor, o *UserRelationship) error {
	*o = UserRelationship{}
	return nil
}

func userRelationshipBeforeUpsertHook(e boil.Executor, o *UserRelationship) error {
	*o = UserRelationship{}
	return nil
}

func userRelationshipAfterUpsertHook(e boil.Executor, o *UserRelationship) error {
	*o = UserRelationship{}
	return nil
}

func testUserRelationshipsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &UserRelationship{}
	o := &UserRelationship{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userRelationshipDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserRelationship object: %s", err)
	}

	AddUserRelationshipHook(boil.BeforeInsertHook, userRelationshipBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userRelationshipBeforeInsertHooks = []UserRelationshipHook{}

	AddUserRelationshipHook(boil.AfterInsertHook, userRelationshipAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userRelationshipAfterInsertHooks = []UserRelationshipHook{}

	AddUserRelationshipHook(boil.AfterSelectHook, userRelationshipAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userRelationshipAfterSelectHooks = []UserRelationshipHook{}

	AddUserRelationshipHook(boil.BeforeUpdateHook, userRelationshipBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userRelationshipBeforeUpdateHooks = []UserRelationshipHook{}

	AddUserRelationshipHook(boil.AfterUpdateHook, userRelationshipAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userRelationshipAfterUpdateHooks = []UserRelationshipHook{}

	AddUserRelationshipHook(boil.BeforeDeleteHook, userRelationshipBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userRelationshipBeforeDeleteHooks = []UserRelationshipHook{}

	AddUserRelationshipHook(boil.AfterDeleteHook, userRelationshipAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userRelationshipAfterDeleteHooks = []UserRelationshipHook{}

	AddUserRelationshipHook(boil.BeforeUpsertHook, userRelationshipBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userRelationshipBeforeUpsertHooks = []UserRelationshipHook{}

	AddUserRelationshipHook(boil.AfterUpsertHook, userRelationshipAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userRelationshipAfterUpsertHooks = []UserRelationshipHook{}
}
func testUserRelationshipsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userRelationship := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UserRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserRelationshipsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userRelationship := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationship.Insert(tx, userRelationshipColumns...); err != nil {
		t.Error(err)
	}

	count, err := UserRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserRelationshipToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local UserRelationship
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userRelationshipDBTypes, true, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
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

	slice := UserRelationshipSlice{&local}
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

func testUserRelationshipToOneCvtermUsingObject(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local UserRelationship
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userRelationshipDBTypes, true, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ObjectID = foreign.CvtermID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Object(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := UserRelationshipSlice{&local}
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

func testUserRelationshipToOneCvtermUsingSubject(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local UserRelationship
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userRelationshipDBTypes, true, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.SubjectID = foreign.CvtermID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Subject(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CvtermID != foreign.CvtermID {
		t.Errorf("want: %v, got %v", foreign.CvtermID, check.CvtermID)
	}

	slice := UserRelationshipSlice{&local}
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

func testUserRelationshipToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a UserRelationship
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userRelationshipDBTypes, false, strmangle.SetComplement(userRelationshipPrimaryKeyColumns, userRelationshipColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeUserRelationship != &a {
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
func testUserRelationshipToOneSetOpCvtermUsingObject(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a UserRelationship
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userRelationshipDBTypes, false, strmangle.SetComplement(userRelationshipPrimaryKeyColumns, userRelationshipColumnsWithoutDefault)...); err != nil {
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
		err = a.SetObject(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Object != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ObjectUserRelationship != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ObjectID != x.CvtermID {
			t.Error("foreign key was wrong value", a.ObjectID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ObjectID))
		reflect.Indirect(reflect.ValueOf(&a.ObjectID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ObjectID != x.CvtermID {
			t.Error("foreign key was wrong value", a.ObjectID, x.CvtermID)
		}
	}
}
func testUserRelationshipToOneSetOpCvtermUsingSubject(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a UserRelationship
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userRelationshipDBTypes, false, strmangle.SetComplement(userRelationshipPrimaryKeyColumns, userRelationshipColumnsWithoutDefault)...); err != nil {
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
		err = a.SetSubject(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Subject != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SubjectUserRelationship != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SubjectID != x.CvtermID {
			t.Error("foreign key was wrong value", a.SubjectID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SubjectID))
		reflect.Indirect(reflect.ValueOf(&a.SubjectID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SubjectID != x.CvtermID {
			t.Error("foreign key was wrong value", a.SubjectID, x.CvtermID)
		}
	}
}
func testUserRelationshipsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userRelationship := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = userRelationship.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testUserRelationshipsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userRelationship := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UserRelationshipSlice{userRelationship}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testUserRelationshipsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userRelationship := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := UserRelationships(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userRelationshipDBTypes = map[string]string{"IsActive": "boolean", "ObjectID": "integer", "SubjectID": "integer", "TypeID": "integer", "UserRelationshipID": "integer"}
	_                       = bytes.MinRead
)

func testUserRelationshipsUpdate(t *testing.T) {
	t.Parallel()

	if len(userRelationshipColumns) == len(userRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	userRelationship := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UserRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true, userRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	if err = userRelationship.Update(tx); err != nil {
		t.Error(err)
	}
}

func testUserRelationshipsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userRelationshipColumns) == len(userRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	userRelationship := &UserRelationship{}
	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UserRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, userRelationship, userRelationshipDBTypes, true, userRelationshipPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userRelationshipColumns, userRelationshipPrimaryKeyColumns) {
		fields = userRelationshipColumns
	} else {
		fields = strmangle.SetComplement(
			userRelationshipColumns,
			userRelationshipPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(userRelationship))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := UserRelationshipSlice{userRelationship}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testUserRelationshipsUpsert(t *testing.T) {
	t.Parallel()

	if len(userRelationshipColumns) == len(userRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	userRelationship := UserRelationship{}
	if err = randomize.Struct(seed, &userRelationship, userRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userRelationship.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert UserRelationship: %s", err)
	}

	count, err := UserRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &userRelationship, userRelationshipDBTypes, false, userRelationshipPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserRelationship struct: %s", err)
	}

	if err = userRelationship.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert UserRelationship: %s", err)
	}

	count, err = UserRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
