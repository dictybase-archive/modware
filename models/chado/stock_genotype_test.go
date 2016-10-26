package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStockGenotypes(t *testing.T) {
	t.Parallel()

	query := StockGenotypes(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStockGenotypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockGenotype := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockGenotype.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := StockGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockGenotypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockGenotype := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockGenotypes(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := StockGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockGenotypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockGenotype := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockGenotypeSlice{stockGenotype}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := StockGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStockGenotypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockGenotype := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StockGenotypeExists(tx, stockGenotype.StockGenotypeID)
	if err != nil {
		t.Errorf("Unable to check if StockGenotype exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockGenotypeExistsG to return true, but got false.")
	}
}
func testStockGenotypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockGenotype := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	stockGenotypeFound, err := FindStockGenotype(tx, stockGenotype.StockGenotypeID)
	if err != nil {
		t.Error(err)
	}

	if stockGenotypeFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStockGenotypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockGenotype := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = StockGenotypes(tx).Bind(stockGenotype); err != nil {
		t.Error(err)
	}
}

func testStockGenotypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockGenotype := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := StockGenotypes(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockGenotypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockGenotypeOne := &StockGenotype{}
	stockGenotypeTwo := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotypeOne, stockGenotypeDBTypes, false, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}
	if err = randomize.Struct(seed, stockGenotypeTwo, stockGenotypeDBTypes, false, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotypeOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockGenotypeTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockGenotypes(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockGenotypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockGenotypeOne := &StockGenotype{}
	stockGenotypeTwo := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotypeOne, stockGenotypeDBTypes, false, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}
	if err = randomize.Struct(seed, stockGenotypeTwo, stockGenotypeDBTypes, false, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotypeOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = stockGenotypeTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func stockGenotypeBeforeInsertHook(e boil.Executor, o *StockGenotype) error {
	*o = StockGenotype{}
	return nil
}

func stockGenotypeAfterInsertHook(e boil.Executor, o *StockGenotype) error {
	*o = StockGenotype{}
	return nil
}

func stockGenotypeAfterSelectHook(e boil.Executor, o *StockGenotype) error {
	*o = StockGenotype{}
	return nil
}

func stockGenotypeBeforeUpdateHook(e boil.Executor, o *StockGenotype) error {
	*o = StockGenotype{}
	return nil
}

func stockGenotypeAfterUpdateHook(e boil.Executor, o *StockGenotype) error {
	*o = StockGenotype{}
	return nil
}

func stockGenotypeBeforeDeleteHook(e boil.Executor, o *StockGenotype) error {
	*o = StockGenotype{}
	return nil
}

func stockGenotypeAfterDeleteHook(e boil.Executor, o *StockGenotype) error {
	*o = StockGenotype{}
	return nil
}

func stockGenotypeBeforeUpsertHook(e boil.Executor, o *StockGenotype) error {
	*o = StockGenotype{}
	return nil
}

func stockGenotypeAfterUpsertHook(e boil.Executor, o *StockGenotype) error {
	*o = StockGenotype{}
	return nil
}

func testStockGenotypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &StockGenotype{}
	o := &StockGenotype{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockGenotypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StockGenotype object: %s", err)
	}

	AddStockGenotypeHook(boil.BeforeInsertHook, stockGenotypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockGenotypeBeforeInsertHooks = []StockGenotypeHook{}

	AddStockGenotypeHook(boil.AfterInsertHook, stockGenotypeAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockGenotypeAfterInsertHooks = []StockGenotypeHook{}

	AddStockGenotypeHook(boil.AfterSelectHook, stockGenotypeAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockGenotypeAfterSelectHooks = []StockGenotypeHook{}

	AddStockGenotypeHook(boil.BeforeUpdateHook, stockGenotypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockGenotypeBeforeUpdateHooks = []StockGenotypeHook{}

	AddStockGenotypeHook(boil.AfterUpdateHook, stockGenotypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockGenotypeAfterUpdateHooks = []StockGenotypeHook{}

	AddStockGenotypeHook(boil.BeforeDeleteHook, stockGenotypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockGenotypeBeforeDeleteHooks = []StockGenotypeHook{}

	AddStockGenotypeHook(boil.AfterDeleteHook, stockGenotypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockGenotypeAfterDeleteHooks = []StockGenotypeHook{}

	AddStockGenotypeHook(boil.BeforeUpsertHook, stockGenotypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockGenotypeBeforeUpsertHooks = []StockGenotypeHook{}

	AddStockGenotypeHook(boil.AfterUpsertHook, stockGenotypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockGenotypeAfterUpsertHooks = []StockGenotypeHook{}
}
func testStockGenotypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockGenotype := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockGenotypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockGenotype := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotype.Insert(tx, stockGenotypeColumns...); err != nil {
		t.Error(err)
	}

	count, err := StockGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockGenotypeToOneGenotypeUsingGenotype(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockGenotype
	var foreign Genotype

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockGenotypeDBTypes, true, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, genotypeDBTypes, true, genotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genotype struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.GenotypeID = foreign.GenotypeID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Genotype(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.GenotypeID != foreign.GenotypeID {
		t.Errorf("want: %v, got %v", foreign.GenotypeID, check.GenotypeID)
	}

	slice := StockGenotypeSlice{&local}
	if err = local.L.LoadGenotype(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Genotype == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Genotype = nil
	if err = local.L.LoadGenotype(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Genotype == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockGenotypeToOneStockUsingStock(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local StockGenotype
	var foreign Stock

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockGenotypeDBTypes, true, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.StockID = foreign.StockID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Stock(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StockID != foreign.StockID {
		t.Errorf("want: %v, got %v", foreign.StockID, check.StockID)
	}

	slice := StockGenotypeSlice{&local}
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

func testStockGenotypeToOneSetOpGenotypeUsingGenotype(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockGenotype
	var b, c Genotype

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockGenotypeDBTypes, false, strmangle.SetComplement(stockGenotypePrimaryKeyColumns, stockGenotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, genotypeDBTypes, false, strmangle.SetComplement(genotypePrimaryKeyColumns, genotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, genotypeDBTypes, false, strmangle.SetComplement(genotypePrimaryKeyColumns, genotypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Genotype{&b, &c} {
		err = a.SetGenotype(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Genotype != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.StockGenotype != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.GenotypeID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.GenotypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.GenotypeID))
		reflect.Indirect(reflect.ValueOf(&a.GenotypeID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GenotypeID != x.GenotypeID {
			t.Error("foreign key was wrong value", a.GenotypeID, x.GenotypeID)
		}
	}
}
func testStockGenotypeToOneSetOpStockUsingStock(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a StockGenotype
	var b, c Stock

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockGenotypeDBTypes, false, strmangle.SetComplement(stockGenotypePrimaryKeyColumns, stockGenotypeColumnsWithoutDefault)...); err != nil {
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

		if x.R.StockGenotype != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.StockID))
		reflect.Indirect(reflect.ValueOf(&a.StockID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StockID != x.StockID {
			t.Error("foreign key was wrong value", a.StockID, x.StockID)
		}
	}
}
func testStockGenotypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockGenotype := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = stockGenotype.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStockGenotypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockGenotype := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StockGenotypeSlice{stockGenotype}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStockGenotypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockGenotype := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := StockGenotypes(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockGenotypeDBTypes = map[string]string{"GenotypeID": "integer", "StockGenotypeID": "integer", "StockID": "integer"}
	_                    = bytes.MinRead
)

func testStockGenotypesUpdate(t *testing.T) {
	t.Parallel()

	if len(stockGenotypeColumns) == len(stockGenotypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockGenotype := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true, stockGenotypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	if err = stockGenotype.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStockGenotypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockGenotypeColumns) == len(stockGenotypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	stockGenotype := &StockGenotype{}
	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotype.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := StockGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, stockGenotype, stockGenotypeDBTypes, true, stockGenotypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockGenotypeColumns, stockGenotypePrimaryKeyColumns) {
		fields = stockGenotypeColumns
	} else {
		fields = strmangle.SetComplement(
			stockGenotypeColumns,
			stockGenotypePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(stockGenotype))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StockGenotypeSlice{stockGenotype}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStockGenotypesUpsert(t *testing.T) {
	t.Parallel()

	if len(stockGenotypeColumns) == len(stockGenotypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	stockGenotype := StockGenotype{}
	if err = randomize.Struct(seed, &stockGenotype, stockGenotypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = stockGenotype.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockGenotype: %s", err)
	}

	count, err := StockGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &stockGenotype, stockGenotypeDBTypes, false, stockGenotypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockGenotype struct: %s", err)
	}

	if err = stockGenotype.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert StockGenotype: %s", err)
	}

	count, err = StockGenotypes(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
