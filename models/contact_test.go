package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testContacts(t *testing.T) {
	t.Parallel()

	query := Contacts(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testContactsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contact := &Contact{}
	if err = randomize.Struct(seed, contact, contactDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contact.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = contact.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Contacts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContactsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contact := &Contact{}
	if err = randomize.Struct(seed, contact, contactDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contact.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Contacts(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Contacts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContactsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contact := &Contact{}
	if err = randomize.Struct(seed, contact, contactDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contact.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ContactSlice{contact}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Contacts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testContactsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contact := &Contact{}
	if err = randomize.Struct(seed, contact, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contact.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := ContactExists(tx, contact.ContactID)
	if err != nil {
		t.Errorf("Unable to check if Contact exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ContactExistsG to return true, but got false.")
	}
}
func testContactsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contact := &Contact{}
	if err = randomize.Struct(seed, contact, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contact.Insert(tx); err != nil {
		t.Error(err)
	}

	contactFound, err := FindContact(tx, contact.ContactID)
	if err != nil {
		t.Error(err)
	}

	if contactFound == nil {
		t.Error("want a record, got nil")
	}
}
func testContactsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contact := &Contact{}
	if err = randomize.Struct(seed, contact, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contact.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Contacts(tx).Bind(contact); err != nil {
		t.Error(err)
	}
}

func testContactsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contact := &Contact{}
	if err = randomize.Struct(seed, contact, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contact.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Contacts(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testContactsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contactOne := &Contact{}
	contactTwo := &Contact{}
	if err = randomize.Struct(seed, contactOne, contactDBTypes, false, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}
	if err = randomize.Struct(seed, contactTwo, contactDBTypes, false, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = contactTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Contacts(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testContactsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	contactOne := &Contact{}
	contactTwo := &Contact{}
	if err = randomize.Struct(seed, contactOne, contactDBTypes, false, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}
	if err = randomize.Struct(seed, contactTwo, contactDBTypes, false, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contactOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = contactTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Contacts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func contactBeforeInsertHook(e boil.Executor, o *Contact) error {
	*o = Contact{}
	return nil
}

func contactAfterInsertHook(e boil.Executor, o *Contact) error {
	*o = Contact{}
	return nil
}

func contactAfterSelectHook(e boil.Executor, o *Contact) error {
	*o = Contact{}
	return nil
}

func contactBeforeUpdateHook(e boil.Executor, o *Contact) error {
	*o = Contact{}
	return nil
}

func contactAfterUpdateHook(e boil.Executor, o *Contact) error {
	*o = Contact{}
	return nil
}

func contactBeforeDeleteHook(e boil.Executor, o *Contact) error {
	*o = Contact{}
	return nil
}

func contactAfterDeleteHook(e boil.Executor, o *Contact) error {
	*o = Contact{}
	return nil
}

func contactBeforeUpsertHook(e boil.Executor, o *Contact) error {
	*o = Contact{}
	return nil
}

func contactAfterUpsertHook(e boil.Executor, o *Contact) error {
	*o = Contact{}
	return nil
}

func testContactsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Contact{}
	o := &Contact{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, contactDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Contact object: %s", err)
	}

	AddContactHook(boil.BeforeInsertHook, contactBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	contactBeforeInsertHooks = []ContactHook{}

	AddContactHook(boil.AfterInsertHook, contactAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	contactAfterInsertHooks = []ContactHook{}

	AddContactHook(boil.AfterSelectHook, contactAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	contactAfterSelectHooks = []ContactHook{}

	AddContactHook(boil.BeforeUpdateHook, contactBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	contactBeforeUpdateHooks = []ContactHook{}

	AddContactHook(boil.AfterUpdateHook, contactAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	contactAfterUpdateHooks = []ContactHook{}

	AddContactHook(boil.BeforeDeleteHook, contactBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	contactBeforeDeleteHooks = []ContactHook{}

	AddContactHook(boil.AfterDeleteHook, contactAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	contactAfterDeleteHooks = []ContactHook{}

	AddContactHook(boil.BeforeUpsertHook, contactBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	contactBeforeUpsertHooks = []ContactHook{}

	AddContactHook(boil.AfterUpsertHook, contactAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	contactAfterUpsertHooks = []ContactHook{}
}
func testContactsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contact := &Contact{}
	if err = randomize.Struct(seed, contact, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contact.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Contacts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContactsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contact := &Contact{}
	if err = randomize.Struct(seed, contact, contactDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contact.Insert(tx, contactColumns...); err != nil {
		t.Error(err)
	}

	count, err := Contacts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContactOneToOneContactRelationshipUsingObjectContactRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign ContactRelationship
	var local Contact

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, contactRelationshipDBTypes, true, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.ObjectID = local.ContactID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ObjectContactRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ObjectID != foreign.ObjectID {
		t.Errorf("want: %v, got %v", foreign.ObjectID, check.ObjectID)
	}

	slice := ContactSlice{&local}
	if err = local.L.LoadObjectContactRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.ObjectContactRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ObjectContactRelationship = nil
	if err = local.L.LoadObjectContactRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.ObjectContactRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testContactOneToOneContactRelationshipUsingSubjectContactRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign ContactRelationship
	var local Contact

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, contactRelationshipDBTypes, true, contactRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.SubjectID = local.ContactID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.SubjectContactRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.SubjectID != foreign.SubjectID {
		t.Errorf("want: %v, got %v", foreign.SubjectID, check.SubjectID)
	}

	slice := ContactSlice{&local}
	if err = local.L.LoadSubjectContactRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.SubjectContactRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SubjectContactRelationship = nil
	if err = local.L.LoadSubjectContactRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.SubjectContactRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testContactOneToOneSetOpContactRelationshipUsingObjectContactRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Contact
	var b, c ContactRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, contactRelationshipDBTypes, false, strmangle.SetComplement(contactRelationshipPrimaryKeyColumns, contactRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, contactRelationshipDBTypes, false, strmangle.SetComplement(contactRelationshipPrimaryKeyColumns, contactRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ContactRelationship{&b, &c} {
		err = a.SetObjectContactRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ObjectContactRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Object != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.ContactID != x.ObjectID {
			t.Error("foreign key was wrong value", a.ContactID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.ObjectID))
		reflect.Indirect(reflect.ValueOf(&x.ObjectID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ContactID != x.ObjectID {
			t.Error("foreign key was wrong value", a.ContactID, x.ObjectID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testContactOneToOneSetOpContactRelationshipUsingSubjectContactRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Contact
	var b, c ContactRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, contactRelationshipDBTypes, false, strmangle.SetComplement(contactRelationshipPrimaryKeyColumns, contactRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, contactRelationshipDBTypes, false, strmangle.SetComplement(contactRelationshipPrimaryKeyColumns, contactRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ContactRelationship{&b, &c} {
		err = a.SetSubjectContactRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SubjectContactRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Subject != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.ContactID != x.SubjectID {
			t.Error("foreign key was wrong value", a.ContactID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.SubjectID))
		reflect.Indirect(reflect.ValueOf(&x.SubjectID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ContactID != x.SubjectID {
			t.Error("foreign key was wrong value", a.ContactID, x.SubjectID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testContactToManyStockcollections(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Contact
	var b, c Stockcollection

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, stockcollectionDBTypes, false, stockcollectionColumnsWithDefault...)
	randomize.Struct(seed, &c, stockcollectionDBTypes, false, stockcollectionColumnsWithDefault...)
	b.ContactID.Valid = true
	c.ContactID.Valid = true
	b.ContactID.Int = a.ContactID
	c.ContactID.Int = a.ContactID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	stockcollection, err := a.Stockcollections(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range stockcollection {
		if v.ContactID.Int == b.ContactID.Int {
			bFound = true
		}
		if v.ContactID.Int == c.ContactID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ContactSlice{&a}
	if err = a.L.LoadStockcollections(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Stockcollections); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Stockcollections = nil
	if err = a.L.LoadStockcollections(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Stockcollections); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", stockcollection)
	}
}

func testContactToManyAddOpStockcollections(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Contact
	var b, c, d, e Stockcollection

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Stockcollection{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, stockcollectionDBTypes, false, strmangle.SetComplement(stockcollectionPrimaryKeyColumns, stockcollectionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Stockcollection{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddStockcollections(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ContactID != first.ContactID.Int {
			t.Error("foreign key was wrong value", a.ContactID, first.ContactID.Int)
		}
		if a.ContactID != second.ContactID.Int {
			t.Error("foreign key was wrong value", a.ContactID, second.ContactID.Int)
		}

		if first.R.Contact != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Contact != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Stockcollections[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Stockcollections[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Stockcollections(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testContactToManySetOpStockcollections(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Contact
	var b, c, d, e Stockcollection

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Stockcollection{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, stockcollectionDBTypes, false, strmangle.SetComplement(stockcollectionPrimaryKeyColumns, stockcollectionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetStockcollections(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Stockcollections(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetStockcollections(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Stockcollections(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.ContactID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.ContactID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.ContactID != d.ContactID.Int {
		t.Error("foreign key was wrong value", a.ContactID, d.ContactID.Int)
	}
	if a.ContactID != e.ContactID.Int {
		t.Error("foreign key was wrong value", a.ContactID, e.ContactID.Int)
	}

	if b.R.Contact != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Contact != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Contact != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Contact != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Stockcollections[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Stockcollections[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testContactToManyRemoveOpStockcollections(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Contact
	var b, c, d, e Stockcollection

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Stockcollection{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, stockcollectionDBTypes, false, strmangle.SetComplement(stockcollectionPrimaryKeyColumns, stockcollectionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddStockcollections(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Stockcollections(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveStockcollections(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Stockcollections(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.ContactID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.ContactID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Contact != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Contact != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Contact != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Contact != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Stockcollections) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Stockcollections[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Stockcollections[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testContactToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Contact
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
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

	slice := ContactSlice{&local}
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

func testContactToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Contact
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypeContacts[0] != &a {
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

func testContactToOneRemoveOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Contact
	var b Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contactDBTypes, false, strmangle.SetComplement(contactPrimaryKeyColumns, contactColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.TypeContacts) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testContactsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contact := &Contact{}
	if err = randomize.Struct(seed, contact, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contact.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = contact.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testContactsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contact := &Contact{}
	if err = randomize.Struct(seed, contact, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contact.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ContactSlice{contact}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testContactsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contact := &Contact{}
	if err = randomize.Struct(seed, contact, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contact.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Contacts(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	contactDBTypes = map[string]string{"ContactID": "integer", "Description": "character varying", "Name": "character varying", "TypeID": "integer"}
	_              = bytes.MinRead
)

func testContactsUpdate(t *testing.T) {
	t.Parallel()

	if len(contactColumns) == len(contactPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	contact := &Contact{}
	if err = randomize.Struct(seed, contact, contactDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contact.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Contacts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, contact, contactDBTypes, true, contactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	if err = contact.Update(tx); err != nil {
		t.Error(err)
	}
}

func testContactsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(contactColumns) == len(contactPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	contact := &Contact{}
	if err = randomize.Struct(seed, contact, contactDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contact.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Contacts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, contact, contactDBTypes, true, contactPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(contactColumns, contactPrimaryKeyColumns) {
		fields = contactColumns
	} else {
		fields = strmangle.SetComplement(
			contactColumns,
			contactPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(contact))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := ContactSlice{contact}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testContactsUpsert(t *testing.T) {
	t.Parallel()

	if len(contactColumns) == len(contactPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	contact := Contact{}
	if err = randomize.Struct(seed, &contact, contactDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contact.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Contact: %s", err)
	}

	count, err := Contacts(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &contact, contactDBTypes, false, contactPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Contact struct: %s", err)
	}

	if err = contact.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Contact: %s", err)
	}

	count, err = Contacts(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
