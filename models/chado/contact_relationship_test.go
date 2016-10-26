package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testContactRelationships(t *testing.T) {
	t.Parallel()

	query := ContactRelationships(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testContactRelationshipsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contactRelationship := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = contactRelationship.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := ContactRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContactRelationshipsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contactRelationship := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ContactRelationships(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := ContactRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContactRelationshipsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contactRelationship := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ContactRelationshipSlice{contactRelationship}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := ContactRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testContactRelationshipsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contactRelationship := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := ContactRelationshipExists(tx, contactRelationship.ContactRelationshipID)
	if err != nil {
		t.Errorf("Unable to check if ContactRelationship exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ContactRelationshipExistsG to return true, but got false.")
	}
}
func testContactRelationshipsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contactRelationship := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	contactRelationshipFound, err := FindContactRelationship(tx, contactRelationship.ContactRelationshipID)
	if err != nil {
		t.Error(err)
	}

	if contactRelationshipFound == nil {
		t.Error("want a record, got nil")
	}
}
func testContactRelationshipsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contactRelationship := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ContactRelationships(tx).Bind(contactRelationship); err != nil {
		t.Error(err)
	}
}

func testContactRelationshipsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contactRelationship := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := ContactRelationships(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testContactRelationshipsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contactRelationshipOne := &ContactRelationship{}
	contactRelationshipTwo := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationshipOne, contactRelationshipDBTypes, false, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}
	if err = randomize.Struct(seed, contactRelationshipTwo, contactRelationshipDBTypes, false, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationshipOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = contactRelationshipTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := ContactRelationships(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testContactRelationshipsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	contactRelationshipOne := &ContactRelationship{}
	contactRelationshipTwo := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationshipOne, contactRelationshipDBTypes, false, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}
	if err = randomize.Struct(seed, contactRelationshipTwo, contactRelationshipDBTypes, false, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationshipOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = contactRelationshipTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ContactRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func contactRelationshipBeforeInsertHook(e boil.Executor, o *ContactRelationship) error {
	*o = ContactRelationship{}
	return nil
}

func contactRelationshipAfterInsertHook(e boil.Executor, o *ContactRelationship) error {
	*o = ContactRelationship{}
	return nil
}

func contactRelationshipAfterSelectHook(e boil.Executor, o *ContactRelationship) error {
	*o = ContactRelationship{}
	return nil
}

func contactRelationshipBeforeUpdateHook(e boil.Executor, o *ContactRelationship) error {
	*o = ContactRelationship{}
	return nil
}

func contactRelationshipAfterUpdateHook(e boil.Executor, o *ContactRelationship) error {
	*o = ContactRelationship{}
	return nil
}

func contactRelationshipBeforeDeleteHook(e boil.Executor, o *ContactRelationship) error {
	*o = ContactRelationship{}
	return nil
}

func contactRelationshipAfterDeleteHook(e boil.Executor, o *ContactRelationship) error {
	*o = ContactRelationship{}
	return nil
}

func contactRelationshipBeforeUpsertHook(e boil.Executor, o *ContactRelationship) error {
	*o = ContactRelationship{}
	return nil
}

func contactRelationshipAfterUpsertHook(e boil.Executor, o *ContactRelationship) error {
	*o = ContactRelationship{}
	return nil
}

func testContactRelationshipsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &ContactRelationship{}
	o := &ContactRelationship{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, contactRelationshipDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ContactRelationship object: %s", err)
	}

	AddContactRelationshipHook(boil.BeforeInsertHook, contactRelationshipBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	contactRelationshipBeforeInsertHooks = []ContactRelationshipHook{}

	AddContactRelationshipHook(boil.AfterInsertHook, contactRelationshipAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	contactRelationshipAfterInsertHooks = []ContactRelationshipHook{}

	AddContactRelationshipHook(boil.AfterSelectHook, contactRelationshipAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	contactRelationshipAfterSelectHooks = []ContactRelationshipHook{}

	AddContactRelationshipHook(boil.BeforeUpdateHook, contactRelationshipBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	contactRelationshipBeforeUpdateHooks = []ContactRelationshipHook{}

	AddContactRelationshipHook(boil.AfterUpdateHook, contactRelationshipAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	contactRelationshipAfterUpdateHooks = []ContactRelationshipHook{}

	AddContactRelationshipHook(boil.BeforeDeleteHook, contactRelationshipBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	contactRelationshipBeforeDeleteHooks = []ContactRelationshipHook{}

	AddContactRelationshipHook(boil.AfterDeleteHook, contactRelationshipAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	contactRelationshipAfterDeleteHooks = []ContactRelationshipHook{}

	AddContactRelationshipHook(boil.BeforeUpsertHook, contactRelationshipBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	contactRelationshipBeforeUpsertHooks = []ContactRelationshipHook{}

	AddContactRelationshipHook(boil.AfterUpsertHook, contactRelationshipAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	contactRelationshipAfterUpsertHooks = []ContactRelationshipHook{}
}
func testContactRelationshipsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contactRelationship := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ContactRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContactRelationshipsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contactRelationship := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationship.Insert(tx, contactRelationshipColumns...); err != nil {
		t.Error(err)
	}

	count, err := ContactRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContactRelationshipToOneContactUsingObject(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local ContactRelationship
	var foreign Contact

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, contactRelationshipDBTypes, true, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ObjectID = foreign.ContactID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Object(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ContactID != foreign.ContactID {
		t.Errorf("want: %v, got %v", foreign.ContactID, check.ContactID)
	}

	slice := ContactRelationshipSlice{&local}
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

func testContactRelationshipToOneContactUsingSubject(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local ContactRelationship
	var foreign Contact

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, contactRelationshipDBTypes, true, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.SubjectID = foreign.ContactID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Subject(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ContactID != foreign.ContactID {
		t.Errorf("want: %v, got %v", foreign.ContactID, check.ContactID)
	}

	slice := ContactRelationshipSlice{&local}
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

func testContactRelationshipToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local ContactRelationship
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, contactRelationshipDBTypes, true, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
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

	slice := ContactRelationshipSlice{&local}
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

func testContactRelationshipToOneSetOpContactUsingObject(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a ContactRelationship
	var b, c Contact

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contactRelationshipDBTypes, false, strmangle.SetComplement(contactRelationshipPrimaryKeyColumns, contactRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Contact{&b, &c} {
		err = a.SetObject(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Object != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ObjectContactRelationship != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ObjectID != x.ContactID {
			t.Error("foreign key was wrong value", a.ObjectID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ObjectID))
		reflect.Indirect(reflect.ValueOf(&a.ObjectID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ObjectID != x.ContactID {
			t.Error("foreign key was wrong value", a.ObjectID, x.ContactID)
		}
	}
}
func testContactRelationshipToOneSetOpContactUsingSubject(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a ContactRelationship
	var b, c Contact

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contactRelationshipDBTypes, false, strmangle.SetComplement(contactRelationshipPrimaryKeyColumns, contactRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Contact{&b, &c} {
		err = a.SetSubject(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Subject != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SubjectContactRelationship != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SubjectID != x.ContactID {
			t.Error("foreign key was wrong value", a.SubjectID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SubjectID))
		reflect.Indirect(reflect.ValueOf(&a.SubjectID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SubjectID != x.ContactID {
			t.Error("foreign key was wrong value", a.SubjectID, x.ContactID)
		}
	}
}
func testContactRelationshipToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a ContactRelationship
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contactRelationshipDBTypes, false, strmangle.SetComplement(contactRelationshipPrimaryKeyColumns, contactRelationshipColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeContactRelationship != &a {
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
func testContactRelationshipsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contactRelationship := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = contactRelationship.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testContactRelationshipsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contactRelationship := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ContactRelationshipSlice{contactRelationship}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testContactRelationshipsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contactRelationship := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := ContactRelationships(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	contactRelationshipDBTypes = map[string]string{"ContactRelationshipID": "integer", "ObjectID": "integer", "SubjectID": "integer", "TypeID": "integer"}
	_                          = bytes.MinRead
)

func testContactRelationshipsUpdate(t *testing.T) {
	t.Parallel()

	if len(contactRelationshipColumns) == len(contactRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	contactRelationship := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ContactRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	if err = contactRelationship.Update(tx); err != nil {
		t.Error(err)
	}
}

func testContactRelationshipsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(contactRelationshipColumns) == len(contactRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	contactRelationship := &ContactRelationship{}
	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationship.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ContactRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, contactRelationship, contactRelationshipDBTypes, true, contactRelationshipPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(contactRelationshipColumns, contactRelationshipPrimaryKeyColumns) {
		fields = contactRelationshipColumns
	} else {
		fields = strmangle.SetComplement(
			contactRelationshipColumns,
			contactRelationshipPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(contactRelationship))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := ContactRelationshipSlice{contactRelationship}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testContactRelationshipsUpsert(t *testing.T) {
	t.Parallel()

	if len(contactRelationshipColumns) == len(contactRelationshipPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	contactRelationship := ContactRelationship{}
	if err = randomize.Struct(seed, &contactRelationship, contactRelationshipDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactRelationship.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert ContactRelationship: %s", err)
	}

	count, err := ContactRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &contactRelationship, contactRelationshipDBTypes, false, contactRelationshipPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}

	if err = contactRelationship.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert ContactRelationship: %s", err)
	}

	count, err = ContactRelationships(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
