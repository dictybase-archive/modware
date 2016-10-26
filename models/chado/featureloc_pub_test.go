package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testFeaturelocPubs(t *testing.T) {
	t.Parallel()

	query := FeaturelocPubs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFeaturelocPubsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurelocPub := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featurelocPub.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturelocPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturelocPubsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurelocPub := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeaturelocPubs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := FeaturelocPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeaturelocPubsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurelocPub := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeaturelocPubSlice{featurelocPub}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturelocPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFeaturelocPubsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurelocPub := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPub.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FeaturelocPubExists(tx, featurelocPub.FeaturelocPubID)
	if err != nil {
		t.Errorf("Unable to check if FeaturelocPub exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeaturelocPubExistsG to return true, but got false.")
	}
}
func testFeaturelocPubsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurelocPub := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPub.Insert(tx); err != nil {
		t.Error(err)
	}

	featurelocPubFound, err := FindFeaturelocPub(tx, featurelocPub.FeaturelocPubID)
	if err != nil {
		t.Error(err)
	}

	if featurelocPubFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFeaturelocPubsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurelocPub := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = FeaturelocPubs(tx).Bind(featurelocPub); err != nil {
		t.Error(err)
	}
}

func testFeaturelocPubsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurelocPub := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := FeaturelocPubs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeaturelocPubsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurelocPubOne := &FeaturelocPub{}
	featurelocPubTwo := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPubOne, featurelocPubDBTypes, false, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}
	if err = randomize.Struct(seed, featurelocPubTwo, featurelocPubDBTypes, false, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featurelocPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeaturelocPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeaturelocPubsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featurelocPubOne := &FeaturelocPub{}
	featurelocPubTwo := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPubOne, featurelocPubDBTypes, false, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}
	if err = randomize.Struct(seed, featurelocPubTwo, featurelocPubDBTypes, false, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPubOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = featurelocPubTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturelocPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func featurelocPubBeforeInsertHook(e boil.Executor, o *FeaturelocPub) error {
	*o = FeaturelocPub{}
	return nil
}

func featurelocPubAfterInsertHook(e boil.Executor, o *FeaturelocPub) error {
	*o = FeaturelocPub{}
	return nil
}

func featurelocPubAfterSelectHook(e boil.Executor, o *FeaturelocPub) error {
	*o = FeaturelocPub{}
	return nil
}

func featurelocPubBeforeUpdateHook(e boil.Executor, o *FeaturelocPub) error {
	*o = FeaturelocPub{}
	return nil
}

func featurelocPubAfterUpdateHook(e boil.Executor, o *FeaturelocPub) error {
	*o = FeaturelocPub{}
	return nil
}

func featurelocPubBeforeDeleteHook(e boil.Executor, o *FeaturelocPub) error {
	*o = FeaturelocPub{}
	return nil
}

func featurelocPubAfterDeleteHook(e boil.Executor, o *FeaturelocPub) error {
	*o = FeaturelocPub{}
	return nil
}

func featurelocPubBeforeUpsertHook(e boil.Executor, o *FeaturelocPub) error {
	*o = FeaturelocPub{}
	return nil
}

func featurelocPubAfterUpsertHook(e boil.Executor, o *FeaturelocPub) error {
	*o = FeaturelocPub{}
	return nil
}

func testFeaturelocPubsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FeaturelocPub{}
	o := &FeaturelocPub{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featurelocPubDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub object: %s", err)
	}

	AddFeaturelocPubHook(boil.BeforeInsertHook, featurelocPubBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featurelocPubBeforeInsertHooks = []FeaturelocPubHook{}

	AddFeaturelocPubHook(boil.AfterInsertHook, featurelocPubAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featurelocPubAfterInsertHooks = []FeaturelocPubHook{}

	AddFeaturelocPubHook(boil.AfterSelectHook, featurelocPubAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featurelocPubAfterSelectHooks = []FeaturelocPubHook{}

	AddFeaturelocPubHook(boil.BeforeUpdateHook, featurelocPubBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featurelocPubBeforeUpdateHooks = []FeaturelocPubHook{}

	AddFeaturelocPubHook(boil.AfterUpdateHook, featurelocPubAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featurelocPubAfterUpdateHooks = []FeaturelocPubHook{}

	AddFeaturelocPubHook(boil.BeforeDeleteHook, featurelocPubBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featurelocPubBeforeDeleteHooks = []FeaturelocPubHook{}

	AddFeaturelocPubHook(boil.AfterDeleteHook, featurelocPubAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featurelocPubAfterDeleteHooks = []FeaturelocPubHook{}

	AddFeaturelocPubHook(boil.BeforeUpsertHook, featurelocPubBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featurelocPubBeforeUpsertHooks = []FeaturelocPubHook{}

	AddFeaturelocPubHook(boil.AfterUpsertHook, featurelocPubAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featurelocPubAfterUpsertHooks = []FeaturelocPubHook{}
}
func testFeaturelocPubsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurelocPub := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturelocPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeaturelocPubsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurelocPub := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPub.Insert(tx, featurelocPubColumns...); err != nil {
		t.Error(err)
	}

	count, err := FeaturelocPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeaturelocPubToOnePubUsingPub(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeaturelocPub
	var foreign Pub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featurelocPubDBTypes, true, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, pubDBTypes, true, pubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pub struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.PubID = foreign.PubID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Pub(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PubID != foreign.PubID {
		t.Errorf("want: %v, got %v", foreign.PubID, check.PubID)
	}

	slice := FeaturelocPubSlice{&local}
	if err = local.L.LoadPub(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Pub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Pub = nil
	if err = local.L.LoadPub(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Pub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeaturelocPubToOneFeaturelocUsingFeatureloc(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local FeaturelocPub
	var foreign Featureloc

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, featurelocPubDBTypes, true, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, featurelocDBTypes, true, featurelocColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Featureloc struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.FeaturelocID = foreign.FeaturelocID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Featureloc(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.FeaturelocID != foreign.FeaturelocID {
		t.Errorf("want: %v, got %v", foreign.FeaturelocID, check.FeaturelocID)
	}

	slice := FeaturelocPubSlice{&local}
	if err = local.L.LoadFeatureloc(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Featureloc == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Featureloc = nil
	if err = local.L.LoadFeatureloc(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Featureloc == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFeaturelocPubToOneSetOpPubUsingPub(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeaturelocPub
	var b, c Pub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurelocPubDBTypes, false, strmangle.SetComplement(featurelocPubPrimaryKeyColumns, featurelocPubColumnsWithoutDefault)...); err != nil {
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
		err = a.SetPub(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Pub != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FeaturelocPub != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PubID))
		reflect.Indirect(reflect.ValueOf(&a.PubID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PubID != x.PubID {
			t.Error("foreign key was wrong value", a.PubID, x.PubID)
		}
	}
}
func testFeaturelocPubToOneSetOpFeaturelocUsingFeatureloc(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a FeaturelocPub
	var b, c Featureloc

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, featurelocPubDBTypes, false, strmangle.SetComplement(featurelocPubPrimaryKeyColumns, featurelocPubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featurelocDBTypes, false, strmangle.SetComplement(featurelocPrimaryKeyColumns, featurelocColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featurelocDBTypes, false, strmangle.SetComplement(featurelocPrimaryKeyColumns, featurelocColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Featureloc{&b, &c} {
		err = a.SetFeatureloc(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Featureloc != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FeaturelocPub != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FeaturelocID != x.FeaturelocID {
			t.Error("foreign key was wrong value", a.FeaturelocID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FeaturelocID))
		reflect.Indirect(reflect.ValueOf(&a.FeaturelocID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FeaturelocID != x.FeaturelocID {
			t.Error("foreign key was wrong value", a.FeaturelocID, x.FeaturelocID)
		}
	}
}
func testFeaturelocPubsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurelocPub := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPub.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = featurelocPub.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturelocPubsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurelocPub := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FeaturelocPubSlice{featurelocPub}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFeaturelocPubsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featurelocPub := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPub.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := FeaturelocPubs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featurelocPubDBTypes = map[string]string{"FeaturelocID": "integer", "FeaturelocPubID": "integer", "PubID": "integer"}
	_                    = bytes.MinRead
)

func testFeaturelocPubsUpdate(t *testing.T) {
	t.Parallel()

	if len(featurelocPubColumns) == len(featurelocPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featurelocPub := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturelocPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true, featurelocPubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	if err = featurelocPub.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFeaturelocPubsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featurelocPubColumns) == len(featurelocPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	featurelocPub := &FeaturelocPub{}
	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPub.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := FeaturelocPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, featurelocPub, featurelocPubDBTypes, true, featurelocPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featurelocPubColumns, featurelocPubPrimaryKeyColumns) {
		fields = featurelocPubColumns
	} else {
		fields = strmangle.SetComplement(
			featurelocPubColumns,
			featurelocPubPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(featurelocPub))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FeaturelocPubSlice{featurelocPub}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFeaturelocPubsUpsert(t *testing.T) {
	t.Parallel()

	if len(featurelocPubColumns) == len(featurelocPubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	featurelocPub := FeaturelocPub{}
	if err = randomize.Struct(seed, &featurelocPub, featurelocPubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = featurelocPub.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeaturelocPub: %s", err)
	}

	count, err := FeaturelocPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &featurelocPub, featurelocPubDBTypes, false, featurelocPubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeaturelocPub struct: %s", err)
	}

	if err = featurelocPub.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert FeaturelocPub: %s", err)
	}

	count, err = FeaturelocPubs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
