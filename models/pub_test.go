package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testPubs(t *testing.T) {
	t.Parallel()

	query := Pubs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testPubsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pub := &Pub{}
	if err = randomize.Struct(seed, pub, pubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = pub.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPubsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pub := &Pub{}
	if err = randomize.Struct(seed, pub, pubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Pubs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Pubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPubsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pub := &Pub{}
	if err = randomize.Struct(seed, pub, pubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PubSlice{pub}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testPubsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pub := &Pub{}
	if err = randomize.Struct(seed, pub, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pub.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := PubExists(tx, pub.PubID)
	if err != nil {
		t.Errorf("Unable to check if Pub exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PubExistsG to return true, but got false.")
	}
}
func testPubsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pub := &Pub{}
	if err = randomize.Struct(seed, pub, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pub.Insert(tx); err != nil {
		t.Error(err)
	}

	pubFound, err := FindPub(tx, pub.PubID)
	if err != nil {
		t.Error(err)
	}

	if pubFound == nil {
		t.Error("want a record, got nil")
	}
}
func testPubsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pub := &Pub{}
	if err = randomize.Struct(seed, pub, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Pubs(tx).Bind(pub); err != nil {
		t.Error(err)
	}
}

func testPubsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pub := &Pub{}
	if err = randomize.Struct(seed, pub, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pub.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Pubs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPubsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pubOne := &Pub{}
	pubTwo := &Pub{}
	if err = randomize.Struct(seed, pubOne, pubDBTypes, false, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}
	if err = randomize.Struct(seed, pubTwo, pubDBTypes, false, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = pubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Pubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPubsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	pubOne := &Pub{}
	pubTwo := &Pub{}
	if err = randomize.Struct(seed, pubOne, pubDBTypes, false, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}
	if err = randomize.Struct(seed, pubTwo, pubDBTypes, false, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = pubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func pubBeforeInsertHook(e boil.Executor, o *Pub) error {
	*o = Pub{}
	return nil
}

func pubAfterInsertHook(e boil.Executor, o *Pub) error {
	*o = Pub{}
	return nil
}

func pubAfterSelectHook(e boil.Executor, o *Pub) error {
	*o = Pub{}
	return nil
}

func pubBeforeUpdateHook(e boil.Executor, o *Pub) error {
	*o = Pub{}
	return nil
}

func pubAfterUpdateHook(e boil.Executor, o *Pub) error {
	*o = Pub{}
	return nil
}

func pubBeforeDeleteHook(e boil.Executor, o *Pub) error {
	*o = Pub{}
	return nil
}

func pubAfterDeleteHook(e boil.Executor, o *Pub) error {
	*o = Pub{}
	return nil
}

func pubBeforeUpsertHook(e boil.Executor, o *Pub) error {
	*o = Pub{}
	return nil
}

func pubAfterUpsertHook(e boil.Executor, o *Pub) error {
	*o = Pub{}
	return nil
}

func testPubsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Pub{}
	o := &Pub{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, pubDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Pub object: %s", err)
	}

	AddPubHook(boil.BeforeInsertHook, pubBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	pubBeforeInsertHooks = []PubHook{}

	AddPubHook(boil.AfterInsertHook, pubAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	pubAfterInsertHooks = []PubHook{}

	AddPubHook(boil.AfterSelectHook, pubAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	pubAfterSelectHooks = []PubHook{}

	AddPubHook(boil.BeforeUpdateHook, pubBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	pubBeforeUpdateHooks = []PubHook{}

	AddPubHook(boil.AfterUpdateHook, pubAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	pubAfterUpdateHooks = []PubHook{}

	AddPubHook(boil.BeforeDeleteHook, pubBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	pubBeforeDeleteHooks = []PubHook{}

	AddPubHook(boil.AfterDeleteHook, pubAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	pubAfterDeleteHooks = []PubHook{}

	AddPubHook(boil.BeforeUpsertHook, pubBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	pubBeforeUpsertHooks = []PubHook{}

	AddPubHook(boil.AfterUpsertHook, pubAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	pubAfterUpsertHooks = []PubHook{}
}
func testPubsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pub := &Pub{}
	if err = randomize.Struct(seed, pub, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPubsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pub := &Pub{}
	if err = randomize.Struct(seed, pub, pubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pub.Insert(tx, pubColumns...); err != nil {
		t.Error(err)
	}

	count, err := Pubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPubOneToOneStockPubUsingStockPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockPub
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockPubDBTypes, true, stockPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadStockPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockPub = nil
	if err = local.L.LoadStockPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOneFeaturelocPubUsingFeaturelocPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeaturelocPub
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featurelocPubDBTypes, true, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeaturelocPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadFeaturelocPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturelocPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeaturelocPub = nil
	if err = local.L.LoadFeaturelocPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturelocPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOnePubDbxrefUsingPubDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign PubDbxref
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, pubDbxrefDBTypes, true, pubDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubDbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.PubDbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadPubDbxref(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.PubDbxref == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.PubDbxref = nil
	if err = local.L.LoadPubDbxref(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.PubDbxref == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOnePubRelationshipUsingObjectPubRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign PubRelationship
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, pubRelationshipDBTypes, true, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.ObjectID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ObjectPubRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ObjectID != foreign.ObjectID {
		t.Errorf("want: %v, got %v", foreign.ObjectID, check.ObjectID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadObjectPubRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.ObjectPubRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ObjectPubRelationship = nil
	if err = local.L.LoadObjectPubRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.ObjectPubRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOnePubRelationshipUsingSubjectPubRelationship(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign PubRelationship
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, pubRelationshipDBTypes, true, pubRelationshipColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PubRelationship struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.SubjectID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.SubjectPubRelationship(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.SubjectID != foreign.SubjectID {
		t.Errorf("want: %v, got %v", foreign.SubjectID, check.SubjectID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadSubjectPubRelationship(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.SubjectPubRelationship == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SubjectPubRelationship = nil
	if err = local.L.LoadSubjectPubRelationship(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.SubjectPubRelationship == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOnePubauthorUsingPubauthor(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Pubauthor
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, pubauthorDBTypes, true, pubauthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubauthor struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Pubauthor(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadPubauthor(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Pubauthor == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Pubauthor = nil
	if err = local.L.LoadPubauthor(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Pubauthor == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOnePubpropUsingPubprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Pubprop
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, pubpropDBTypes, true, pubpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pubprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Pubprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadPubprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Pubprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Pubprop = nil
	if err = local.L.LoadPubprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Pubprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOneFeatureCvtermPubUsingFeatureCvtermPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureCvtermPub
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureCvtermPubDBTypes, true, featureCvtermPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvtermPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureCvtermPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadFeatureCvtermPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureCvtermPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureCvtermPub = nil
	if err = local.L.LoadFeatureCvtermPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureCvtermPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOneStockCvtermUsingStockCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockCvterm
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockCvtermDBTypes, true, stockCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockCvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadStockCvterm(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockCvterm == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockCvterm = nil
	if err = local.L.LoadStockCvterm(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockCvterm == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOneStockRelationshipPubUsingStockRelationshipPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockRelationshipPub
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockRelationshipPubDBTypes, true, stockRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockRelationshipPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockRelationshipPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadStockRelationshipPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockRelationshipPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockRelationshipPub = nil
	if err = local.L.LoadStockRelationshipPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockRelationshipPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOneFeatureCvtermUsingFeatureCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureCvterm
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureCvtermDBTypes, true, featureCvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureCvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureCvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadFeatureCvterm(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureCvterm == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureCvterm = nil
	if err = local.L.LoadFeatureCvterm(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureCvterm == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOneFeaturePubUsingFeaturePub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeaturePub
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featurePubDBTypes, true, featurePubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturePub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeaturePub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadFeaturePub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturePub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeaturePub = nil
	if err = local.L.LoadFeaturePub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturePub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOneFeatureSynonymUsingFeatureSynonym(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureSynonym
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureSynonymDBTypes, true, featureSynonymColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureSynonym struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureSynonym(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadFeatureSynonym(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureSynonym == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureSynonym = nil
	if err = local.L.LoadFeatureSynonym(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureSynonym == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOneFeatureRelationshippropPubUsingFeatureRelationshippropPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureRelationshippropPub
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureRelationshippropPubDBTypes, true, featureRelationshippropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshippropPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureRelationshippropPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadFeatureRelationshippropPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureRelationshippropPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureRelationshippropPub = nil
	if err = local.L.LoadFeatureRelationshippropPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureRelationshippropPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOneFeaturepropPubUsingFeaturepropPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeaturepropPub
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featurepropPubDBTypes, true, featurepropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturepropPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeaturepropPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadFeaturepropPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturepropPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeaturepropPub = nil
	if err = local.L.LoadFeaturepropPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeaturepropPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOneFeatureRelationshipPubUsingFeatureRelationshipPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign FeatureRelationshipPub
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureRelationshipPubDBTypes, true, featureRelationshipPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureRelationshipPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.FeatureRelationshipPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadFeatureRelationshipPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureRelationshipPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FeatureRelationshipPub = nil
	if err = local.L.LoadFeatureRelationshipPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.FeatureRelationshipPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOnePhendescUsingPhendesc(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Phendesc
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phendescDBTypes, true, phendescColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phendesc struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Phendesc(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadPhendesc(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Phendesc == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Phendesc = nil
	if err = local.L.LoadPhendesc(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Phendesc == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOnePhenstatementUsingPhenstatement(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Phenstatement
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenstatementDBTypes, true, phenstatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Phenstatement struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Phenstatement(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadPhenstatement(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Phenstatement == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Phenstatement = nil
	if err = local.L.LoadPhenstatement(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Phenstatement == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOneStockpropPubUsingStockpropPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign StockpropPub
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockpropPubDBTypes, true, stockpropPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockpropPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StockpropPub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadStockpropPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.StockpropPub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StockpropPub = nil
	if err = local.L.LoadStockpropPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StockpropPub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOnePhenotypeComparisonUsingPhenotypeComparison(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign PhenotypeComparison
	var local Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, phenotypeComparisonDBTypes, true, phenotypeComparisonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PhenotypeComparison struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.PubID = local.PubID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.PhenotypeComparison(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := PubSlice{&local}
	if err = local.L.LoadPhenotypeComparison(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.PhenotypeComparison == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.PhenotypeComparison = nil
	if err = local.L.LoadPhenotypeComparison(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.PhenotypeComparison == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPubOneToOneSetOpStockPubUsingStockPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c StockPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockPubDBTypes, false, strmangle.SetComplement(stockPubPrimaryKeyColumns, stockPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockPubDBTypes, false, strmangle.SetComplement(stockPubPrimaryKeyColumns, stockPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockPub{&b, &c} {
		err = a.SetStockPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpFeaturelocPubUsingFeaturelocPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c FeaturelocPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featurelocPubDBTypes, false, strmangle.SetComplement(featurelocPubPrimaryKeyColumns, featurelocPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featurelocPubDBTypes, false, strmangle.SetComplement(featurelocPubPrimaryKeyColumns, featurelocPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeaturelocPub{&b, &c} {
		err = a.SetFeaturelocPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeaturelocPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpPubDbxrefUsingPubDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c PubDbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pubDbxrefDBTypes, false, strmangle.SetComplement(pubDbxrefPrimaryKeyColumns, pubDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pubDbxrefDBTypes, false, strmangle.SetComplement(pubDbxrefPrimaryKeyColumns, pubDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*PubDbxref{&b, &c} {
		err = a.SetPubDbxref(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.PubDbxref != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpPubRelationshipUsingObjectPubRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c PubRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pubRelationshipDBTypes, false, strmangle.SetComplement(pubRelationshipPrimaryKeyColumns, pubRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pubRelationshipDBTypes, false, strmangle.SetComplement(pubRelationshipPrimaryKeyColumns, pubRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*PubRelationship{&b, &c} {
		err = a.SetObjectPubRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ObjectPubRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Object != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.ObjectID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.ObjectID))
		reflect.Indirect(reflect.ValueOf(&x.ObjectID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.ObjectID {
			t.Error("foreign key was wrong value", a.PubID, x.ObjectID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpPubRelationshipUsingSubjectPubRelationship(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c PubRelationship

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pubRelationshipDBTypes, false, strmangle.SetComplement(pubRelationshipPrimaryKeyColumns, pubRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pubRelationshipDBTypes, false, strmangle.SetComplement(pubRelationshipPrimaryKeyColumns, pubRelationshipColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*PubRelationship{&b, &c} {
		err = a.SetSubjectPubRelationship(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SubjectPubRelationship != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Subject != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.SubjectID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.SubjectID))
		reflect.Indirect(reflect.ValueOf(&x.SubjectID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.SubjectID {
			t.Error("foreign key was wrong value", a.PubID, x.SubjectID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpPubauthorUsingPubauthor(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c Pubauthor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pubauthorDBTypes, false, strmangle.SetComplement(pubauthorPrimaryKeyColumns, pubauthorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pubauthorDBTypes, false, strmangle.SetComplement(pubauthorPrimaryKeyColumns, pubauthorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Pubauthor{&b, &c} {
		err = a.SetPubauthor(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Pubauthor != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpPubpropUsingPubprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c Pubprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pubpropDBTypes, false, strmangle.SetComplement(pubpropPrimaryKeyColumns, pubpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pubpropDBTypes, false, strmangle.SetComplement(pubpropPrimaryKeyColumns, pubpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Pubprop{&b, &c} {
		err = a.SetPubprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Pubprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpFeatureCvtermPubUsingFeatureCvtermPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c FeatureCvtermPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureCvtermPubDBTypes, false, strmangle.SetComplement(featureCvtermPubPrimaryKeyColumns, featureCvtermPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureCvtermPubDBTypes, false, strmangle.SetComplement(featureCvtermPubPrimaryKeyColumns, featureCvtermPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureCvtermPub{&b, &c} {
		err = a.SetFeatureCvtermPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureCvtermPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpStockCvtermUsingStockCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c StockCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockCvtermDBTypes, false, strmangle.SetComplement(stockCvtermPrimaryKeyColumns, stockCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockCvtermDBTypes, false, strmangle.SetComplement(stockCvtermPrimaryKeyColumns, stockCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockCvterm{&b, &c} {
		err = a.SetStockCvterm(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockCvterm != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpStockRelationshipPubUsingStockRelationshipPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c StockRelationshipPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockRelationshipPubDBTypes, false, strmangle.SetComplement(stockRelationshipPubPrimaryKeyColumns, stockRelationshipPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockRelationshipPubDBTypes, false, strmangle.SetComplement(stockRelationshipPubPrimaryKeyColumns, stockRelationshipPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockRelationshipPub{&b, &c} {
		err = a.SetStockRelationshipPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockRelationshipPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpFeatureCvtermUsingFeatureCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c FeatureCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureCvtermDBTypes, false, strmangle.SetComplement(featureCvtermPrimaryKeyColumns, featureCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureCvtermDBTypes, false, strmangle.SetComplement(featureCvtermPrimaryKeyColumns, featureCvtermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureCvterm{&b, &c} {
		err = a.SetFeatureCvterm(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureCvterm != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpFeaturePubUsingFeaturePub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c FeaturePub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featurePubDBTypes, false, strmangle.SetComplement(featurePubPrimaryKeyColumns, featurePubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featurePubDBTypes, false, strmangle.SetComplement(featurePubPrimaryKeyColumns, featurePubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeaturePub{&b, &c} {
		err = a.SetFeaturePub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeaturePub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpFeatureSynonymUsingFeatureSynonym(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c FeatureSynonym

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureSynonymDBTypes, false, strmangle.SetComplement(featureSynonymPrimaryKeyColumns, featureSynonymColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureSynonymDBTypes, false, strmangle.SetComplement(featureSynonymPrimaryKeyColumns, featureSynonymColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureSynonym{&b, &c} {
		err = a.SetFeatureSynonym(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureSynonym != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpFeatureRelationshippropPubUsingFeatureRelationshippropPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c FeatureRelationshippropPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureRelationshippropPubDBTypes, false, strmangle.SetComplement(featureRelationshippropPubPrimaryKeyColumns, featureRelationshippropPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureRelationshippropPubDBTypes, false, strmangle.SetComplement(featureRelationshippropPubPrimaryKeyColumns, featureRelationshippropPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureRelationshippropPub{&b, &c} {
		err = a.SetFeatureRelationshippropPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureRelationshippropPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpFeaturepropPubUsingFeaturepropPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c FeaturepropPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featurepropPubDBTypes, false, strmangle.SetComplement(featurepropPubPrimaryKeyColumns, featurepropPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featurepropPubDBTypes, false, strmangle.SetComplement(featurepropPubPrimaryKeyColumns, featurepropPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeaturepropPub{&b, &c} {
		err = a.SetFeaturepropPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeaturepropPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpFeatureRelationshipPubUsingFeatureRelationshipPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c FeatureRelationshipPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureRelationshipPubDBTypes, false, strmangle.SetComplement(featureRelationshipPubPrimaryKeyColumns, featureRelationshipPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureRelationshipPubDBTypes, false, strmangle.SetComplement(featureRelationshipPubPrimaryKeyColumns, featureRelationshipPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FeatureRelationshipPub{&b, &c} {
		err = a.SetFeatureRelationshipPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FeatureRelationshipPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpPhendescUsingPhendesc(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c Phendesc

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, phendescDBTypes, false, strmangle.SetComplement(phendescPrimaryKeyColumns, phendescColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, phendescDBTypes, false, strmangle.SetComplement(phendescPrimaryKeyColumns, phendescColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Phendesc{&b, &c} {
		err = a.SetPhendesc(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Phendesc != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpPhenstatementUsingPhenstatement(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c Phenstatement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, phenstatementDBTypes, false, strmangle.SetComplement(phenstatementPrimaryKeyColumns, phenstatementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, phenstatementDBTypes, false, strmangle.SetComplement(phenstatementPrimaryKeyColumns, phenstatementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Phenstatement{&b, &c} {
		err = a.SetPhenstatement(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Phenstatement != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpStockpropPubUsingStockpropPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c StockpropPub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockpropPubDBTypes, false, strmangle.SetComplement(stockpropPubPrimaryKeyColumns, stockpropPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockpropPubDBTypes, false, strmangle.SetComplement(stockpropPubPrimaryKeyColumns, stockpropPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StockpropPub{&b, &c} {
		err = a.SetStockpropPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StockpropPub != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubOneToOneSetOpPhenotypeComparisonUsingPhenotypeComparison(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c PhenotypeComparison

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*PhenotypeComparison{&b, &c} {
		err = a.SetPhenotypeComparison(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.PhenotypeComparison != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Pub != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.PubID))
		reflect.Indirect(reflect.ValueOf(&x.PubID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testPubToManyStockRelationshipCvterms(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c StockRelationshipCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, stockRelationshipCvtermDBTypes, false, stockRelationshipCvtermColumnsWithDefault...)
	randomize.Struct(seed, &c, stockRelationshipCvtermDBTypes, false, stockRelationshipCvtermColumnsWithDefault...)
	b.PubID.Valid = true
	c.PubID.Valid = true
	b.PubID.Int = a.PubID
	c.PubID.Int = a.PubID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	stockRelationshipCvterm, err := a.StockRelationshipCvterms(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range stockRelationshipCvterm {
		if v.PubID.Int == b.PubID.Int {
			bFound = true
		}
		if v.PubID.Int == c.PubID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PubSlice{&a}
	if err = a.L.LoadStockRelationshipCvterms(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.StockRelationshipCvterms); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.StockRelationshipCvterms = nil
	if err = a.L.LoadStockRelationshipCvterms(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.StockRelationshipCvterms); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", stockRelationshipCvterm)
	}
}

func testPubToManyPhenotypeComparisonCvterms(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c PhenotypeComparisonCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, phenotypeComparisonCvtermDBTypes, false, phenotypeComparisonCvtermColumnsWithDefault...)
	randomize.Struct(seed, &c, phenotypeComparisonCvtermDBTypes, false, phenotypeComparisonCvtermColumnsWithDefault...)

	b.PubID = a.PubID
	c.PubID = a.PubID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	phenotypeComparisonCvterm, err := a.PhenotypeComparisonCvterms(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range phenotypeComparisonCvterm {
		if v.PubID == b.PubID {
			bFound = true
		}
		if v.PubID == c.PubID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PubSlice{&a}
	if err = a.L.LoadPhenotypeComparisonCvterms(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PhenotypeComparisonCvterms); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PhenotypeComparisonCvterms = nil
	if err = a.L.LoadPhenotypeComparisonCvterms(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PhenotypeComparisonCvterms); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", phenotypeComparisonCvterm)
	}
}

func testPubToManyAddOpStockRelationshipCvterms(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c, d, e StockRelationshipCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*StockRelationshipCvterm{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, stockRelationshipCvtermDBTypes, false, strmangle.SetComplement(stockRelationshipCvtermPrimaryKeyColumns, stockRelationshipCvtermColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*StockRelationshipCvterm{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddStockRelationshipCvterms(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.PubID != first.PubID.Int {
			t.Error("foreign key was wrong value", a.PubID, first.PubID.Int)
		}
		if a.PubID != second.PubID.Int {
			t.Error("foreign key was wrong value", a.PubID, second.PubID.Int)
		}

		if first.R.Pub != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Pub != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.StockRelationshipCvterms[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.StockRelationshipCvterms[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.StockRelationshipCvterms(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testPubToManySetOpStockRelationshipCvterms(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c, d, e StockRelationshipCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*StockRelationshipCvterm{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, stockRelationshipCvtermDBTypes, false, strmangle.SetComplement(stockRelationshipCvtermPrimaryKeyColumns, stockRelationshipCvtermColumnsWithoutDefault)...); err != nil {
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

	err = a.SetStockRelationshipCvterms(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.StockRelationshipCvterms(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetStockRelationshipCvterms(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.StockRelationshipCvterms(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.PubID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.PubID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.PubID != d.PubID.Int {
		t.Error("foreign key was wrong value", a.PubID, d.PubID.Int)
	}
	if a.PubID != e.PubID.Int {
		t.Error("foreign key was wrong value", a.PubID, e.PubID.Int)
	}

	if b.R.Pub != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Pub != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Pub != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Pub != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.StockRelationshipCvterms[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.StockRelationshipCvterms[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testPubToManyRemoveOpStockRelationshipCvterms(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c, d, e StockRelationshipCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*StockRelationshipCvterm{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, stockRelationshipCvtermDBTypes, false, strmangle.SetComplement(stockRelationshipCvtermPrimaryKeyColumns, stockRelationshipCvtermColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddStockRelationshipCvterms(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.StockRelationshipCvterms(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveStockRelationshipCvterms(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.StockRelationshipCvterms(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.PubID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.PubID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Pub != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Pub != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Pub != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Pub != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.StockRelationshipCvterms) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.StockRelationshipCvterms[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.StockRelationshipCvterms[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testPubToManyAddOpPhenotypeComparisonCvterms(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c, d, e PhenotypeComparisonCvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PhenotypeComparisonCvterm{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeComparisonCvtermDBTypes, false, strmangle.SetComplement(phenotypeComparisonCvtermPrimaryKeyColumns, phenotypeComparisonCvtermColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*PhenotypeComparisonCvterm{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPhenotypeComparisonCvterms(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.PubID != first.PubID {
			t.Error("foreign key was wrong value", a.PubID, first.PubID)
		}
		if a.PubID != second.PubID {
			t.Error("foreign key was wrong value", a.PubID, second.PubID)
		}

		if first.R.Pub != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Pub != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PhenotypeComparisonCvterms[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PhenotypeComparisonCvterms[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PhenotypeComparisonCvterms(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testPubToOneCvtermUsingType(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Pub
	var foreign Cvterm

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
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

	slice := PubSlice{&local}
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

func testPubToOneSetOpCvtermUsingType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Pub
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pubDBTypes, false, strmangle.SetComplement(pubPrimaryKeyColumns, pubColumnsWithoutDefault)...); err != nil {
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

		if x.R.TypePubs[0] != &a {
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
func testPubsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pub := &Pub{}
	if err = randomize.Struct(seed, pub, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = pub.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPubsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pub := &Pub{}
	if err = randomize.Struct(seed, pub, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PubSlice{pub}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testPubsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pub := &Pub{}
	if err = randomize.Struct(seed, pub, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Pubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	pubDBTypes = map[string]string{"IsObsolete": "boolean", "Issue": "character varying", "Miniref": "character varying", "Pages": "character varying", "PubID": "integer", "Publisher": "character varying", "Pubplace": "character varying", "Pyear": "character varying", "SeriesName": "character varying", "Title": "text", "TypeID": "integer", "Uniquename": "text", "Volume": "character varying", "Volumetitle": "text"}
	_          = bytes.MinRead
)

func testPubsUpdate(t *testing.T) {
	t.Parallel()

	if len(pubColumns) == len(pubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	pub := &Pub{}
	if err = randomize.Struct(seed, pub, pubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, pub, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err = pub.Update(tx); err != nil {
		t.Error(err)
	}
}

func testPubsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(pubColumns) == len(pubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	pub := &Pub{}
	if err = randomize.Struct(seed, pub, pubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Pubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, pub, pubDBTypes, true, pubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(pubColumns, pubPrimaryKeyColumns) {
		fields = pubColumns
	} else {
		fields = strmangle.SetComplement(
			pubColumns,
			pubPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(pub))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := PubSlice{pub}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testPubsUpsert(t *testing.T) {
	t.Parallel()

	if len(pubColumns) == len(pubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	pub := Pub{}
	if err = randomize.Struct(seed, &pub, pubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = pub.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Pub: %s", err)
	}

	count, err := Pubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &pub, pubDBTypes, false, pubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err = pub.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Pub: %s", err)
	}

	count, err = Pubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
