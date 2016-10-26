package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testOrganisms(t *testing.T) {
	t.Parallel()

	query := Organisms(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testOrganismsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organism := &Organism{}
	if err = randomize.Struct(seed, organism, organismDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organism.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = organism.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Organisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganismsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organism := &Organism{}
	if err = randomize.Struct(seed, organism, organismDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organism.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Organisms(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Organisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganismsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organism := &Organism{}
	if err = randomize.Struct(seed, organism, organismDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organism.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OrganismSlice{organism}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Organisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testOrganismsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organism := &Organism{}
	if err = randomize.Struct(seed, organism, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organism.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := OrganismExists(tx, organism.OrganismID)
	if err != nil {
		t.Errorf("Unable to check if Organism exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrganismExistsG to return true, but got false.")
	}
}
func testOrganismsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organism := &Organism{}
	if err = randomize.Struct(seed, organism, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organism.Insert(tx); err != nil {
		t.Error(err)
	}

	organismFound, err := FindOrganism(tx, organism.OrganismID)
	if err != nil {
		t.Error(err)
	}

	if organismFound == nil {
		t.Error("want a record, got nil")
	}
}
func testOrganismsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organism := &Organism{}
	if err = randomize.Struct(seed, organism, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organism.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Organisms(tx).Bind(organism); err != nil {
		t.Error(err)
	}
}

func testOrganismsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organism := &Organism{}
	if err = randomize.Struct(seed, organism, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organism.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Organisms(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrganismsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organismOne := &Organism{}
	organismTwo := &Organism{}
	if err = randomize.Struct(seed, organismOne, organismDBTypes, false, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}
	if err = randomize.Struct(seed, organismTwo, organismDBTypes, false, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = organismTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Organisms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrganismsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	organismOne := &Organism{}
	organismTwo := &Organism{}
	if err = randomize.Struct(seed, organismOne, organismDBTypes, false, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}
	if err = randomize.Struct(seed, organismTwo, organismDBTypes, false, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organismOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = organismTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func organismBeforeInsertHook(e boil.Executor, o *Organism) error {
	*o = Organism{}
	return nil
}

func organismAfterInsertHook(e boil.Executor, o *Organism) error {
	*o = Organism{}
	return nil
}

func organismAfterSelectHook(e boil.Executor, o *Organism) error {
	*o = Organism{}
	return nil
}

func organismBeforeUpdateHook(e boil.Executor, o *Organism) error {
	*o = Organism{}
	return nil
}

func organismAfterUpdateHook(e boil.Executor, o *Organism) error {
	*o = Organism{}
	return nil
}

func organismBeforeDeleteHook(e boil.Executor, o *Organism) error {
	*o = Organism{}
	return nil
}

func organismAfterDeleteHook(e boil.Executor, o *Organism) error {
	*o = Organism{}
	return nil
}

func organismBeforeUpsertHook(e boil.Executor, o *Organism) error {
	*o = Organism{}
	return nil
}

func organismAfterUpsertHook(e boil.Executor, o *Organism) error {
	*o = Organism{}
	return nil
}

func testOrganismsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Organism{}
	o := &Organism{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, organismDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Organism object: %s", err)
	}

	AddOrganismHook(boil.BeforeInsertHook, organismBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	organismBeforeInsertHooks = []OrganismHook{}

	AddOrganismHook(boil.AfterInsertHook, organismAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	organismAfterInsertHooks = []OrganismHook{}

	AddOrganismHook(boil.AfterSelectHook, organismAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	organismAfterSelectHooks = []OrganismHook{}

	AddOrganismHook(boil.BeforeUpdateHook, organismBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	organismBeforeUpdateHooks = []OrganismHook{}

	AddOrganismHook(boil.AfterUpdateHook, organismAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	organismAfterUpdateHooks = []OrganismHook{}

	AddOrganismHook(boil.BeforeDeleteHook, organismBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	organismBeforeDeleteHooks = []OrganismHook{}

	AddOrganismHook(boil.AfterDeleteHook, organismAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	organismAfterDeleteHooks = []OrganismHook{}

	AddOrganismHook(boil.BeforeUpsertHook, organismBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	organismBeforeUpsertHooks = []OrganismHook{}

	AddOrganismHook(boil.AfterUpsertHook, organismAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	organismAfterUpsertHooks = []OrganismHook{}
}
func testOrganismsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organism := &Organism{}
	if err = randomize.Struct(seed, organism, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organism.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganismsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organism := &Organism{}
	if err = randomize.Struct(seed, organism, organismDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organism.Insert(tx, organismColumns...); err != nil {
		t.Error(err)
	}

	count, err := Organisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganismOneToOneStockUsingStock(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Stock
	var local Organism

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	foreign.OrganismID.Valid = true

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.OrganismID.Int = local.OrganismID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Stock(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.OrganismID.Int != foreign.OrganismID.Int {
		t.Errorf("want: %v, got %v", foreign.OrganismID.Int, check.OrganismID.Int)
	}

	slice := OrganismSlice{&local}
	if err = local.L.LoadStock(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Stock == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Stock = nil
	if err = local.L.LoadStock(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Stock == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOrganismOneToOneFeatureUsingFeature(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Feature
	var local Organism

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, featureDBTypes, true, featureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Feature struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.OrganismID = local.OrganismID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Feature(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.OrganismID != foreign.OrganismID {
		t.Errorf("want: %v, got %v", foreign.OrganismID, check.OrganismID)
	}

	slice := OrganismSlice{&local}
	if err = local.L.LoadFeature(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Feature == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Feature = nil
	if err = local.L.LoadFeature(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Feature == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOrganismOneToOneJbrowseOrganismUsingJbrowseOrganism(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign JbrowseOrganism
	var local Organism

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, jbrowseOrganismDBTypes, true, jbrowseOrganismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JbrowseOrganism struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.OrganismID = local.OrganismID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.JbrowseOrganism(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.OrganismID != foreign.OrganismID {
		t.Errorf("want: %v, got %v", foreign.OrganismID, check.OrganismID)
	}

	slice := OrganismSlice{&local}
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

func testOrganismOneToOneOrganismDbxrefUsingOrganismDbxref(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign OrganismDbxref
	var local Organism

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, organismDbxrefDBTypes, true, organismDbxrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganismDbxref struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.OrganismID = local.OrganismID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.OrganismDbxref(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.OrganismID != foreign.OrganismID {
		t.Errorf("want: %v, got %v", foreign.OrganismID, check.OrganismID)
	}

	slice := OrganismSlice{&local}
	if err = local.L.LoadOrganismDbxref(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.OrganismDbxref == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.OrganismDbxref = nil
	if err = local.L.LoadOrganismDbxref(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.OrganismDbxref == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOrganismOneToOneOrganismpropUsingOrganismprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Organismprop
	var local Organism

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, organismpropDBTypes, true, organismpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organismprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.OrganismID = local.OrganismID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Organismprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.OrganismID != foreign.OrganismID {
		t.Errorf("want: %v, got %v", foreign.OrganismID, check.OrganismID)
	}

	slice := OrganismSlice{&local}
	if err = local.L.LoadOrganismprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Organismprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Organismprop = nil
	if err = local.L.LoadOrganismprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Organismprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOrganismOneToOneSetOpStockUsingStock(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organism
	var b, c Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organismDBTypes, false, strmangle.SetComplement(organismPrimaryKeyColumns, organismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Stock{&b, &c} {
		err = a.SetStock(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Stock != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Organism != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.OrganismID != x.OrganismID.Int {
			t.Error("foreign key was wrong value", a.OrganismID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.OrganismID.Int))
		reflect.Indirect(reflect.ValueOf(&x.OrganismID.Int)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrganismID != x.OrganismID.Int {
			t.Error("foreign key was wrong value", a.OrganismID, x.OrganismID.Int)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testOrganismOneToOneRemoveOpStockUsingStock(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organism
	var b Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organismDBTypes, false, strmangle.SetComplement(organismPrimaryKeyColumns, organismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetStock(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveStock(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Stock(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Stock != nil {
		t.Error("R struct entry should be nil")
	}

	if b.OrganismID.Valid {
		t.Error("foreign key column should be nil")
	}

	if b.R.Organism != nil {
		t.Error("failed to remove a from b's relationships")
	}
}

func testOrganismOneToOneSetOpFeatureUsingFeature(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organism
	var b, c Feature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organismDBTypes, false, strmangle.SetComplement(organismPrimaryKeyColumns, organismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, featureDBTypes, false, strmangle.SetComplement(featurePrimaryKeyColumns, featureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Feature{&b, &c} {
		err = a.SetFeature(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Feature != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Organism != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.OrganismID != x.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.OrganismID))
		reflect.Indirect(reflect.ValueOf(&x.OrganismID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrganismID != x.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID, x.OrganismID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testOrganismOneToOneSetOpJbrowseOrganismUsingJbrowseOrganism(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organism
	var b, c JbrowseOrganism

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organismDBTypes, false, strmangle.SetComplement(organismPrimaryKeyColumns, organismColumnsWithoutDefault)...); err != nil {
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
		if x.R.Organism != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.OrganismID != x.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.OrganismID))
		reflect.Indirect(reflect.ValueOf(&x.OrganismID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrganismID != x.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID, x.OrganismID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testOrganismOneToOneSetOpOrganismDbxrefUsingOrganismDbxref(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organism
	var b, c OrganismDbxref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organismDBTypes, false, strmangle.SetComplement(organismPrimaryKeyColumns, organismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organismDbxrefDBTypes, false, strmangle.SetComplement(organismDbxrefPrimaryKeyColumns, organismDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, organismDbxrefDBTypes, false, strmangle.SetComplement(organismDbxrefPrimaryKeyColumns, organismDbxrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*OrganismDbxref{&b, &c} {
		err = a.SetOrganismDbxref(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.OrganismDbxref != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Organism != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.OrganismID != x.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.OrganismID))
		reflect.Indirect(reflect.ValueOf(&x.OrganismID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrganismID != x.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID, x.OrganismID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testOrganismOneToOneSetOpOrganismpropUsingOrganismprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organism
	var b, c Organismprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organismDBTypes, false, strmangle.SetComplement(organismPrimaryKeyColumns, organismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organismpropDBTypes, false, strmangle.SetComplement(organismpropPrimaryKeyColumns, organismpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, organismpropDBTypes, false, strmangle.SetComplement(organismpropPrimaryKeyColumns, organismpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Organismprop{&b, &c} {
		err = a.SetOrganismprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Organismprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Organism != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.OrganismID != x.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.OrganismID))
		reflect.Indirect(reflect.ValueOf(&x.OrganismID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrganismID != x.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID, x.OrganismID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testOrganismToManyPhenotypeComparisons(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organism
	var b, c PhenotypeComparison

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, phenotypeComparisonDBTypes, false, phenotypeComparisonColumnsWithDefault...)
	randomize.Struct(seed, &c, phenotypeComparisonDBTypes, false, phenotypeComparisonColumnsWithDefault...)

	b.OrganismID = a.OrganismID
	c.OrganismID = a.OrganismID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	phenotypeComparison, err := a.PhenotypeComparisons(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range phenotypeComparison {
		if v.OrganismID == b.OrganismID {
			bFound = true
		}
		if v.OrganismID == c.OrganismID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OrganismSlice{&a}
	if err = a.L.LoadPhenotypeComparisons(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PhenotypeComparisons); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PhenotypeComparisons = nil
	if err = a.L.LoadPhenotypeComparisons(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PhenotypeComparisons); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", phenotypeComparison)
	}
}

func testOrganismToManyAddOpPhenotypeComparisons(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organism
	var b, c, d, e PhenotypeComparison

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organismDBTypes, false, strmangle.SetComplement(organismPrimaryKeyColumns, organismColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PhenotypeComparison{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, phenotypeComparisonDBTypes, false, strmangle.SetComplement(phenotypeComparisonPrimaryKeyColumns, phenotypeComparisonColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*PhenotypeComparison{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPhenotypeComparisons(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.OrganismID != first.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID, first.OrganismID)
		}
		if a.OrganismID != second.OrganismID {
			t.Error("foreign key was wrong value", a.OrganismID, second.OrganismID)
		}

		if first.R.Organism != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Organism != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PhenotypeComparisons[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PhenotypeComparisons[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PhenotypeComparisons(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testOrganismsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organism := &Organism{}
	if err = randomize.Struct(seed, organism, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organism.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = organism.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testOrganismsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organism := &Organism{}
	if err = randomize.Struct(seed, organism, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organism.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OrganismSlice{organism}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testOrganismsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organism := &Organism{}
	if err = randomize.Struct(seed, organism, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organism.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Organisms(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	organismDBTypes = map[string]string{"Abbreviation": "character varying", "Comment": "text", "CommonName": "character varying", "Genus": "character varying", "OrganismID": "integer", "Species": "character varying"}
	_               = bytes.MinRead
)

func testOrganismsUpdate(t *testing.T) {
	t.Parallel()

	if len(organismColumns) == len(organismPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	organism := &Organism{}
	if err = randomize.Struct(seed, organism, organismDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organism.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, organism, organismDBTypes, true, organismColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	if err = organism.Update(tx); err != nil {
		t.Error(err)
	}
}

func testOrganismsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(organismColumns) == len(organismPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	organism := &Organism{}
	if err = randomize.Struct(seed, organism, organismDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organism.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organisms(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, organism, organismDBTypes, true, organismPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(organismColumns, organismPrimaryKeyColumns) {
		fields = organismColumns
	} else {
		fields = strmangle.SetComplement(
			organismColumns,
			organismPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(organism))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := OrganismSlice{organism}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testOrganismsUpsert(t *testing.T) {
	t.Parallel()

	if len(organismColumns) == len(organismPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	organism := Organism{}
	if err = randomize.Struct(seed, &organism, organismDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organism.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Organism: %s", err)
	}

	count, err := Organisms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &organism, organismDBTypes, false, organismPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Organism struct: %s", err)
	}

	if err = organism.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Organism: %s", err)
	}

	count, err = Organisms(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
