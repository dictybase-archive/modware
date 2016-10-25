package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testAuthUserProviders(t *testing.T) {
	t.Parallel()

	query := AuthUserProviders(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAuthUserProvidersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserProvider := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authUserProvider.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthUserProvidersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserProvider := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthUserProviders(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := AuthUserProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthUserProvidersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserProvider := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthUserProviderSlice{authUserProvider}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAuthUserProvidersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserProvider := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true, authUserProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AuthUserProviderExists(tx, authUserProvider.AuthUserProviderID)
	if err != nil {
		t.Errorf("Unable to check if AuthUserProvider exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthUserProviderExistsG to return true, but got false.")
	}
}
func testAuthUserProvidersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserProvider := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true, authUserProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	authUserProviderFound, err := FindAuthUserProvider(tx, authUserProvider.AuthUserProviderID)
	if err != nil {
		t.Error(err)
	}

	if authUserProviderFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAuthUserProvidersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserProvider := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true, authUserProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthUserProviders(tx).Bind(authUserProvider); err != nil {
		t.Error(err)
	}
}

func testAuthUserProvidersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserProvider := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true, authUserProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := AuthUserProviders(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthUserProvidersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserProviderOne := &AuthUserProvider{}
	authUserProviderTwo := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProviderOne, authUserProviderDBTypes, false, authUserProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}
	if err = randomize.Struct(seed, authUserProviderTwo, authUserProviderDBTypes, false, authUserProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProviderOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authUserProviderTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthUserProviders(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthUserProvidersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authUserProviderOne := &AuthUserProvider{}
	authUserProviderTwo := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProviderOne, authUserProviderDBTypes, false, authUserProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}
	if err = randomize.Struct(seed, authUserProviderTwo, authUserProviderDBTypes, false, authUserProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProviderOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authUserProviderTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func authUserProviderBeforeInsertHook(e boil.Executor, o *AuthUserProvider) error {
	*o = AuthUserProvider{}
	return nil
}

func authUserProviderAfterInsertHook(e boil.Executor, o *AuthUserProvider) error {
	*o = AuthUserProvider{}
	return nil
}

func authUserProviderAfterSelectHook(e boil.Executor, o *AuthUserProvider) error {
	*o = AuthUserProvider{}
	return nil
}

func authUserProviderBeforeUpdateHook(e boil.Executor, o *AuthUserProvider) error {
	*o = AuthUserProvider{}
	return nil
}

func authUserProviderAfterUpdateHook(e boil.Executor, o *AuthUserProvider) error {
	*o = AuthUserProvider{}
	return nil
}

func authUserProviderBeforeDeleteHook(e boil.Executor, o *AuthUserProvider) error {
	*o = AuthUserProvider{}
	return nil
}

func authUserProviderAfterDeleteHook(e boil.Executor, o *AuthUserProvider) error {
	*o = AuthUserProvider{}
	return nil
}

func authUserProviderBeforeUpsertHook(e boil.Executor, o *AuthUserProvider) error {
	*o = AuthUserProvider{}
	return nil
}

func authUserProviderAfterUpsertHook(e boil.Executor, o *AuthUserProvider) error {
	*o = AuthUserProvider{}
	return nil
}

func testAuthUserProvidersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &AuthUserProvider{}
	o := &AuthUserProvider{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, authUserProviderDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider object: %s", err)
	}

	AddAuthUserProviderHook(boil.BeforeInsertHook, authUserProviderBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	authUserProviderBeforeInsertHooks = []AuthUserProviderHook{}

	AddAuthUserProviderHook(boil.AfterInsertHook, authUserProviderAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	authUserProviderAfterInsertHooks = []AuthUserProviderHook{}

	AddAuthUserProviderHook(boil.AfterSelectHook, authUserProviderAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	authUserProviderAfterSelectHooks = []AuthUserProviderHook{}

	AddAuthUserProviderHook(boil.BeforeUpdateHook, authUserProviderBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	authUserProviderBeforeUpdateHooks = []AuthUserProviderHook{}

	AddAuthUserProviderHook(boil.AfterUpdateHook, authUserProviderAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	authUserProviderAfterUpdateHooks = []AuthUserProviderHook{}

	AddAuthUserProviderHook(boil.BeforeDeleteHook, authUserProviderBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	authUserProviderBeforeDeleteHooks = []AuthUserProviderHook{}

	AddAuthUserProviderHook(boil.AfterDeleteHook, authUserProviderAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	authUserProviderAfterDeleteHooks = []AuthUserProviderHook{}

	AddAuthUserProviderHook(boil.BeforeUpsertHook, authUserProviderBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	authUserProviderBeforeUpsertHooks = []AuthUserProviderHook{}

	AddAuthUserProviderHook(boil.AfterUpsertHook, authUserProviderAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	authUserProviderAfterUpsertHooks = []AuthUserProviderHook{}
}
func testAuthUserProvidersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserProvider := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true, authUserProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthUserProvidersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserProvider := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProvider.Insert(tx, authUserProviderColumns...); err != nil {
		t.Error(err)
	}

	count, err := AuthUserProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthUserProviderToOneAuthProviderUsingAuthProvider(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local AuthUserProvider
	var foreign AuthProvider

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, authUserProviderDBTypes, true, authUserProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, authProviderDBTypes, true, authProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthProvider struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AuthProviderID = foreign.AuthProviderID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.AuthProvider(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.AuthProviderID != foreign.AuthProviderID {
		t.Errorf("want: %v, got %v", foreign.AuthProviderID, check.AuthProviderID)
	}

	slice := AuthUserProviderSlice{&local}
	if err = local.L.LoadAuthProvider(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.AuthProvider == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.AuthProvider = nil
	if err = local.L.LoadAuthProvider(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.AuthProvider == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAuthUserProviderToOneSetOpAuthProviderUsingAuthProvider(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthUserProvider
	var b, c AuthProvider

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserProviderDBTypes, false, strmangle.SetComplement(authUserProviderPrimaryKeyColumns, authUserProviderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, authProviderDBTypes, false, strmangle.SetComplement(authProviderPrimaryKeyColumns, authProviderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authProviderDBTypes, false, strmangle.SetComplement(authProviderPrimaryKeyColumns, authProviderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*AuthProvider{&b, &c} {
		err = a.SetAuthProvider(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.AuthProvider != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AuthUserProviders[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AuthProviderID != x.AuthProviderID {
			t.Error("foreign key was wrong value", a.AuthProviderID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AuthProviderID))
		reflect.Indirect(reflect.ValueOf(&a.AuthProviderID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AuthProviderID != x.AuthProviderID {
			t.Error("foreign key was wrong value", a.AuthProviderID, x.AuthProviderID)
		}
	}
}
func testAuthUserProvidersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserProvider := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true, authUserProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authUserProvider.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAuthUserProvidersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserProvider := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true, authUserProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthUserProviderSlice{authUserProvider}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAuthUserProvidersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserProvider := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true, authUserProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthUserProviders(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authUserProviderDBTypes = map[string]string{"AuthProviderID": "integer", "AuthUserProviderID": "integer", "CreatedAt": "timestamp with time zone", "Email": "USER-DEFINED", "Name": "text", "UpdatedAt": "timestamp with time zone"}
	_                       = bytes.MinRead
)

func testAuthUserProvidersUpdate(t *testing.T) {
	t.Parallel()

	if len(authUserProviderColumns) == len(authUserProviderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authUserProvider := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true, authUserProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	if err = authUserProvider.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAuthUserProvidersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authUserProviderColumns) == len(authUserProviderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authUserProvider := &AuthUserProvider{}
	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProvider.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authUserProvider, authUserProviderDBTypes, true, authUserProviderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authUserProviderColumns, authUserProviderPrimaryKeyColumns) {
		fields = authUserProviderColumns
	} else {
		fields = strmangle.SetComplement(
			authUserProviderColumns,
			authUserProviderPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(authUserProvider))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AuthUserProviderSlice{authUserProvider}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAuthUserProvidersUpsert(t *testing.T) {
	t.Parallel()

	if len(authUserProviderColumns) == len(authUserProviderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	authUserProvider := AuthUserProvider{}
	if err = randomize.Struct(seed, &authUserProvider, authUserProviderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserProvider.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthUserProvider: %s", err)
	}

	count, err := AuthUserProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &authUserProvider, authUserProviderDBTypes, false, authUserProviderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthUserProvider struct: %s", err)
	}

	if err = authUserProvider.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthUserProvider: %s", err)
	}

	count, err = AuthUserProviders(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
