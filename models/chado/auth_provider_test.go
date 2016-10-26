package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testAuthProviders(t *testing.T) {
	t.Parallel()

	query := AuthProviders(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAuthProvidersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authProvider := &AuthProvider{}
	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authProvider.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthProvidersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authProvider := &AuthProvider{}
	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthProviders(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := AuthProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthProvidersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authProvider := &AuthProvider{}
	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthProviderSlice{authProvider}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAuthProvidersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authProvider := &AuthProvider{}
	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true, authProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AuthProviderExists(tx, authProvider.AuthProviderID)
	if err != nil {
		t.Errorf("Unable to check if AuthProvider exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthProviderExistsG to return true, but got false.")
	}
}
func testAuthProvidersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authProvider := &AuthProvider{}
	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true, authProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	authProviderFound, err := FindAuthProvider(tx, authProvider.AuthProviderID)
	if err != nil {
		t.Error(err)
	}

	if authProviderFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAuthProvidersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authProvider := &AuthProvider{}
	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true, authProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthProviders(tx).Bind(authProvider); err != nil {
		t.Error(err)
	}
}

func testAuthProvidersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authProvider := &AuthProvider{}
	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true, authProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := AuthProviders(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthProvidersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authProviderOne := &AuthProvider{}
	authProviderTwo := &AuthProvider{}
	if err = randomize.Struct(seed, authProviderOne, authProviderDBTypes, false, authProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}
	if err = randomize.Struct(seed, authProviderTwo, authProviderDBTypes, false, authProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProviderOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authProviderTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthProviders(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthProvidersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authProviderOne := &AuthProvider{}
	authProviderTwo := &AuthProvider{}
	if err = randomize.Struct(seed, authProviderOne, authProviderDBTypes, false, authProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}
	if err = randomize.Struct(seed, authProviderTwo, authProviderDBTypes, false, authProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProviderOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authProviderTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func authProviderBeforeInsertHook(e boil.Executor, o *AuthProvider) error {
	*o = AuthProvider{}
	return nil
}

func authProviderAfterInsertHook(e boil.Executor, o *AuthProvider) error {
	*o = AuthProvider{}
	return nil
}

func authProviderAfterSelectHook(e boil.Executor, o *AuthProvider) error {
	*o = AuthProvider{}
	return nil
}

func authProviderBeforeUpdateHook(e boil.Executor, o *AuthProvider) error {
	*o = AuthProvider{}
	return nil
}

func authProviderAfterUpdateHook(e boil.Executor, o *AuthProvider) error {
	*o = AuthProvider{}
	return nil
}

func authProviderBeforeDeleteHook(e boil.Executor, o *AuthProvider) error {
	*o = AuthProvider{}
	return nil
}

func authProviderAfterDeleteHook(e boil.Executor, o *AuthProvider) error {
	*o = AuthProvider{}
	return nil
}

func authProviderBeforeUpsertHook(e boil.Executor, o *AuthProvider) error {
	*o = AuthProvider{}
	return nil
}

func authProviderAfterUpsertHook(e boil.Executor, o *AuthProvider) error {
	*o = AuthProvider{}
	return nil
}

func testAuthProvidersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &AuthProvider{}
	o := &AuthProvider{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, authProviderDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AuthProvider object: %s", err)
	}

	AddAuthProviderHook(boil.BeforeInsertHook, authProviderBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	authProviderBeforeInsertHooks = []AuthProviderHook{}

	AddAuthProviderHook(boil.AfterInsertHook, authProviderAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	authProviderAfterInsertHooks = []AuthProviderHook{}

	AddAuthProviderHook(boil.AfterSelectHook, authProviderAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	authProviderAfterSelectHooks = []AuthProviderHook{}

	AddAuthProviderHook(boil.BeforeUpdateHook, authProviderBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	authProviderBeforeUpdateHooks = []AuthProviderHook{}

	AddAuthProviderHook(boil.AfterUpdateHook, authProviderAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	authProviderAfterUpdateHooks = []AuthProviderHook{}

	AddAuthProviderHook(boil.BeforeDeleteHook, authProviderBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	authProviderBeforeDeleteHooks = []AuthProviderHook{}

	AddAuthProviderHook(boil.AfterDeleteHook, authProviderAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	authProviderAfterDeleteHooks = []AuthProviderHook{}

	AddAuthProviderHook(boil.BeforeUpsertHook, authProviderBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	authProviderBeforeUpsertHooks = []AuthProviderHook{}

	AddAuthProviderHook(boil.AfterUpsertHook, authProviderAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	authProviderAfterUpsertHooks = []AuthProviderHook{}
}
func testAuthProvidersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authProvider := &AuthProvider{}
	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true, authProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthProvidersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authProvider := &AuthProvider{}
	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProvider.Insert(tx, authProviderColumns...); err != nil {
		t.Error(err)
	}

	count, err := AuthProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthProviderToManyAuthUserProviders(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthProvider
	var b, c AuthUserProvider

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authProviderDBTypes, true, authProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, authUserProviderDBTypes, false, authUserProviderColumnsWithDefault...)
	randomize.Struct(seed, &c, authUserProviderDBTypes, false, authUserProviderColumnsWithDefault...)

	b.AuthProviderID = a.AuthProviderID
	c.AuthProviderID = a.AuthProviderID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	authUserProvider, err := a.AuthUserProviders(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range authUserProvider {
		if v.AuthProviderID == b.AuthProviderID {
			bFound = true
		}
		if v.AuthProviderID == c.AuthProviderID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AuthProviderSlice{&a}
	if err = a.L.LoadAuthUserProviders(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AuthUserProviders); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AuthUserProviders = nil
	if err = a.L.LoadAuthUserProviders(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AuthUserProviders); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", authUserProvider)
	}
}

func testAuthProviderToManyAddOpAuthUserProviders(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthProvider
	var b, c, d, e AuthUserProvider

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authProviderDBTypes, false, strmangle.SetComplement(authProviderPrimaryKeyColumns, authProviderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AuthUserProvider{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, authUserProviderDBTypes, false, strmangle.SetComplement(authUserProviderPrimaryKeyColumns, authUserProviderColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*AuthUserProvider{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAuthUserProviders(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.AuthProviderID != first.AuthProviderID {
			t.Error("foreign key was wrong value", a.AuthProviderID, first.AuthProviderID)
		}
		if a.AuthProviderID != second.AuthProviderID {
			t.Error("foreign key was wrong value", a.AuthProviderID, second.AuthProviderID)
		}

		if first.R.AuthProvider != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.AuthProvider != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AuthUserProviders[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AuthUserProviders[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AuthUserProviders(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testAuthProvidersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authProvider := &AuthProvider{}
	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true, authProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authProvider.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAuthProvidersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authProvider := &AuthProvider{}
	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true, authProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthProviderSlice{authProvider}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAuthProvidersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authProvider := &AuthProvider{}
	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true, authProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthProviders(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authProviderDBTypes = map[string]string{"AuthProviderID": "integer", "Name": "text"}
	_                   = bytes.MinRead
)

func testAuthProvidersUpdate(t *testing.T) {
	t.Parallel()

	if len(authProviderColumns) == len(authProviderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authProvider := &AuthProvider{}
	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true, authProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	if err = authProvider.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAuthProvidersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authProviderColumns) == len(authProviderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authProvider := &AuthProvider{}
	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authProvider, authProviderDBTypes, true, authProviderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authProviderColumns, authProviderPrimaryKeyColumns) {
		fields = authProviderColumns
	} else {
		fields = strmangle.SetComplement(
			authProviderColumns,
			authProviderPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(authProvider))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AuthProviderSlice{authProvider}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAuthProvidersUpsert(t *testing.T) {
	t.Parallel()

	if len(authProviderColumns) == len(authProviderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	authProvider := AuthProvider{}
	if err = randomize.Struct(seed, &authProvider, authProviderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authProvider.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthProvider: %s", err)
	}

	count, err := AuthProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &authProvider, authProviderDBTypes, false, authProviderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	if err = authProvider.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthProvider: %s", err)
	}

	count, err = AuthProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
