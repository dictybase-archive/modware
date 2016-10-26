package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testCVS(t *testing.T) {
	t.Parallel()

	query := CVS(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testCVSDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cv := &CV{}
	if err = randomize.Struct(seed, cv, cvDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cv.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cv.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := CVS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCVSQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cv := &CV{}
	if err = randomize.Struct(seed, cv, cvDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cv.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = CVS(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := CVS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCVSSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cv := &CV{}
	if err = randomize.Struct(seed, cv, cvDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cv.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CVSlice{cv}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := CVS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testCVSExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cv := &CV{}
	if err = randomize.Struct(seed, cv, cvDBTypes, true, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cv.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := CVExists(tx, cv.CVID)
	if err != nil {
		t.Errorf("Unable to check if CV exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CVExistsG to return true, but got false.")
	}
}
func testCVSFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cv := &CV{}
	if err = randomize.Struct(seed, cv, cvDBTypes, true, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cv.Insert(tx); err != nil {
		t.Error(err)
	}

	cvFound, err := FindCV(tx, cv.CVID)
	if err != nil {
		t.Error(err)
	}

	if cvFound == nil {
		t.Error("want a record, got nil")
	}
}
func testCVSBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cv := &CV{}
	if err = randomize.Struct(seed, cv, cvDBTypes, true, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cv.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = CVS(tx).Bind(cv); err != nil {
		t.Error(err)
	}
}

func testCVSOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cv := &CV{}
	if err = randomize.Struct(seed, cv, cvDBTypes, true, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cv.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := CVS(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCVSAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cvOne := &CV{}
	cvTwo := &CV{}
	if err = randomize.Struct(seed, cvOne, cvDBTypes, false, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}
	if err = randomize.Struct(seed, cvTwo, cvDBTypes, false, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := CVS(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCVSCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	cvOne := &CV{}
	cvTwo := &CV{}
	if err = randomize.Struct(seed, cvOne, cvDBTypes, false, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}
	if err = randomize.Struct(seed, cvTwo, cvDBTypes, false, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cvOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = cvTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CVS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func cvBeforeInsertHook(e boil.Executor, o *CV) error {
	*o = CV{}
	return nil
}

func cvAfterInsertHook(e boil.Executor, o *CV) error {
	*o = CV{}
	return nil
}

func cvAfterSelectHook(e boil.Executor, o *CV) error {
	*o = CV{}
	return nil
}

func cvBeforeUpdateHook(e boil.Executor, o *CV) error {
	*o = CV{}
	return nil
}

func cvAfterUpdateHook(e boil.Executor, o *CV) error {
	*o = CV{}
	return nil
}

func cvBeforeDeleteHook(e boil.Executor, o *CV) error {
	*o = CV{}
	return nil
}

func cvAfterDeleteHook(e boil.Executor, o *CV) error {
	*o = CV{}
	return nil
}

func cvBeforeUpsertHook(e boil.Executor, o *CV) error {
	*o = CV{}
	return nil
}

func cvAfterUpsertHook(e boil.Executor, o *CV) error {
	*o = CV{}
	return nil
}

func testCVSHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &CV{}
	o := &CV{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, cvDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CV object: %s", err)
	}

	AddCVHook(boil.BeforeInsertHook, cvBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	cvBeforeInsertHooks = []CVHook{}

	AddCVHook(boil.AfterInsertHook, cvAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	cvAfterInsertHooks = []CVHook{}

	AddCVHook(boil.AfterSelectHook, cvAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	cvAfterSelectHooks = []CVHook{}

	AddCVHook(boil.BeforeUpdateHook, cvBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	cvBeforeUpdateHooks = []CVHook{}

	AddCVHook(boil.AfterUpdateHook, cvAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	cvAfterUpdateHooks = []CVHook{}

	AddCVHook(boil.BeforeDeleteHook, cvBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	cvBeforeDeleteHooks = []CVHook{}

	AddCVHook(boil.AfterDeleteHook, cvAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	cvAfterDeleteHooks = []CVHook{}

	AddCVHook(boil.BeforeUpsertHook, cvBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	cvBeforeUpsertHooks = []CVHook{}

	AddCVHook(boil.AfterUpsertHook, cvAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	cvAfterUpsertHooks = []CVHook{}
}
func testCVSInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cv := &CV{}
	if err = randomize.Struct(seed, cv, cvDBTypes, true, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cv.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CVS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCVSInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cv := &CV{}
	if err = randomize.Struct(seed, cv, cvDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cv.Insert(tx, cvColumns...); err != nil {
		t.Error(err)
	}

	count, err := CVS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCVOneToOneCvtermUsingCvterm(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Cvterm
	var local CV

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, cvtermDBTypes, true, cvtermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvterm struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvDBTypes, true, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.CVID = local.CVID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Cvterm(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CVID != foreign.CVID {
		t.Errorf("want: %v, got %v", foreign.CVID, check.CVID)
	}

	slice := CVSlice{&local}
	if err = local.L.LoadCvterm(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Cvterm == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Cvterm = nil
	if err = local.L.LoadCvterm(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Cvterm == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCVOneToOneCvpropUsingCvprop(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign Cvprop
	var local CV

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, cvpropDBTypes, true, cvpropColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cvprop struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, cvDBTypes, true, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.CVID = local.CVID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Cvprop(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CVID != foreign.CVID {
		t.Errorf("want: %v, got %v", foreign.CVID, check.CVID)
	}

	slice := CVSlice{&local}
	if err = local.L.LoadCvprop(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Cvprop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Cvprop = nil
	if err = local.L.LoadCvprop(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Cvprop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCVOneToOneSetOpCvtermUsingCvterm(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a CV
	var b, c Cvterm

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvDBTypes, false, strmangle.SetComplement(cvPrimaryKeyColumns, cvColumnsWithoutDefault)...); err != nil {
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
		err = a.SetCvterm(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Cvterm != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.CV != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CVID != x.CVID {
			t.Error("foreign key was wrong value", a.CVID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.CVID))
		reflect.Indirect(reflect.ValueOf(&x.CVID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CVID != x.CVID {
			t.Error("foreign key was wrong value", a.CVID, x.CVID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCVOneToOneSetOpCvpropUsingCvprop(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a CV
	var b, c Cvprop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvDBTypes, false, strmangle.SetComplement(cvPrimaryKeyColumns, cvColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cvpropDBTypes, false, strmangle.SetComplement(cvpropPrimaryKeyColumns, cvpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cvpropDBTypes, false, strmangle.SetComplement(cvpropPrimaryKeyColumns, cvpropColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Cvprop{&b, &c} {
		err = a.SetCvprop(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Cvprop != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.CV != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.CVID != x.CVID {
			t.Error("foreign key was wrong value", a.CVID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.CVID))
		reflect.Indirect(reflect.ValueOf(&x.CVID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CVID != x.CVID {
			t.Error("foreign key was wrong value", a.CVID, x.CVID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCVToManyCvtermpaths(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a CV
	var b, c Cvtermpath

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvDBTypes, true, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, cvtermpathDBTypes, false, cvtermpathColumnsWithDefault...)
	randomize.Struct(seed, &c, cvtermpathDBTypes, false, cvtermpathColumnsWithDefault...)

	b.CVID = a.CVID
	c.CVID = a.CVID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	cvtermpath, err := a.Cvtermpaths(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range cvtermpath {
		if v.CVID == b.CVID {
			bFound = true
		}
		if v.CVID == c.CVID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CVSlice{&a}
	if err = a.L.LoadCvtermpaths(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Cvtermpaths); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Cvtermpaths = nil
	if err = a.L.LoadCvtermpaths(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Cvtermpaths); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", cvtermpath)
	}
}

func testCVToManyAddOpCvtermpaths(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a CV
	var b, c, d, e Cvtermpath

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cvDBTypes, false, strmangle.SetComplement(cvPrimaryKeyColumns, cvColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Cvtermpath{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, cvtermpathDBTypes, false, strmangle.SetComplement(cvtermpathPrimaryKeyColumns, cvtermpathColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Cvtermpath{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCvtermpaths(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.CVID != first.CVID {
			t.Error("foreign key was wrong value", a.CVID, first.CVID)
		}
		if a.CVID != second.CVID {
			t.Error("foreign key was wrong value", a.CVID, second.CVID)
		}

		if first.R.CV != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.CV != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Cvtermpaths[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Cvtermpaths[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Cvtermpaths(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCVSReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cv := &CV{}
	if err = randomize.Struct(seed, cv, cvDBTypes, true, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cv.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = cv.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCVSReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cv := &CV{}
	if err = randomize.Struct(seed, cv, cvDBTypes, true, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cv.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CVSlice{cv}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testCVSSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cv := &CV{}
	if err = randomize.Struct(seed, cv, cvDBTypes, true, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cv.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := CVS(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	cvDBTypes = map[string]string{"CVID": "integer", "Definition": "text", "Name": "character varying"}
	_         = bytes.MinRead
)

func testCVSUpdate(t *testing.T) {
	t.Parallel()

	if len(cvColumns) == len(cvPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cv := &CV{}
	if err = randomize.Struct(seed, cv, cvDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cv.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CVS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cv, cvDBTypes, true, cvColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	if err = cv.Update(tx); err != nil {
		t.Error(err)
	}
}

func testCVSSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(cvColumns) == len(cvPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	cv := &CV{}
	if err = randomize.Struct(seed, cv, cvDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cv.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CVS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, cv, cvDBTypes, true, cvPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(cvColumns, cvPrimaryKeyColumns) {
		fields = cvColumns
	} else {
		fields = strmangle.SetComplement(
			cvColumns,
			cvPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(cv))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := CVSlice{cv}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testCVSUpsert(t *testing.T) {
	t.Parallel()

	if len(cvColumns) == len(cvPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	cv := CV{}
	if err = randomize.Struct(seed, &cv, cvDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = cv.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert CV: %s", err)
	}

	count, err := CVS(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &cv, cvDBTypes, false, cvPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CV struct: %s", err)
	}

	if err = cv.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert CV: %s", err)
	}

	count, err = CVS(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
