package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testAuthPermissions(t *testing.T) {
	t.Parallel()

	query := AuthPermissions(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAuthPermissionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authPermission := &AuthPermission{}
	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermission.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authPermission.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthPermissionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authPermission := &AuthPermission{}
	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermission.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthPermissions(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthPermissionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authPermission := &AuthPermission{}
	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermission.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthPermissionSlice{authPermission}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAuthPermissionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authPermission := &AuthPermission{}
	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermission.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AuthPermissionExists(tx, authPermission.AuthPermissionID)
	if err != nil {
		t.Errorf("Unable to check if AuthPermission exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthPermissionExistsG to return true, but got false.")
	}
}
func testAuthPermissionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authPermission := &AuthPermission{}
	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermission.Insert(tx); err != nil {
		t.Error(err)
	}

	authPermissionFound, err := FindAuthPermission(tx, authPermission.AuthPermissionID)
	if err != nil {
		t.Error(err)
	}

	if authPermissionFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAuthPermissionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authPermission := &AuthPermission{}
	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermission.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthPermissions(tx).Bind(authPermission); err != nil {
		t.Error(err)
	}
}

func testAuthPermissionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authPermission := &AuthPermission{}
	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermission.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := AuthPermissions(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthPermissionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authPermissionOne := &AuthPermission{}
	authPermissionTwo := &AuthPermission{}
	if err = randomize.Struct(seed, authPermissionOne, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}
	if err = randomize.Struct(seed, authPermissionTwo, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermissionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authPermissionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthPermissions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthPermissionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authPermissionOne := &AuthPermission{}
	authPermissionTwo := &AuthPermission{}
	if err = randomize.Struct(seed, authPermissionOne, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}
	if err = randomize.Struct(seed, authPermissionTwo, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermissionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authPermissionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func authPermissionBeforeInsertHook(e boil.Executor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionAfterInsertHook(e boil.Executor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionAfterSelectHook(e boil.Executor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionBeforeUpdateHook(e boil.Executor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionAfterUpdateHook(e boil.Executor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionBeforeDeleteHook(e boil.Executor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionAfterDeleteHook(e boil.Executor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionBeforeUpsertHook(e boil.Executor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionAfterUpsertHook(e boil.Executor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func testAuthPermissionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &AuthPermission{}
	o := &AuthPermission{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, authPermissionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AuthPermission object: %s", err)
	}

	AddAuthPermissionHook(boil.BeforeInsertHook, authPermissionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	authPermissionBeforeInsertHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.AfterInsertHook, authPermissionAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	authPermissionAfterInsertHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.AfterSelectHook, authPermissionAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	authPermissionAfterSelectHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.BeforeUpdateHook, authPermissionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	authPermissionBeforeUpdateHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.AfterUpdateHook, authPermissionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	authPermissionAfterUpdateHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.BeforeDeleteHook, authPermissionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	authPermissionBeforeDeleteHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.AfterDeleteHook, authPermissionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	authPermissionAfterDeleteHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.BeforeUpsertHook, authPermissionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	authPermissionBeforeUpsertHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.AfterUpsertHook, authPermissionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	authPermissionAfterUpsertHooks = []AuthPermissionHook{}
}
func testAuthPermissionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authPermission := &AuthPermission{}
	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermission.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthPermissionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authPermission := &AuthPermission{}
	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermission.Insert(tx, authPermissionColumns...); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthPermissionOneToOneAuthRolePermissionUsingAuthRolePermission(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign AuthRolePermission
	var local AuthPermission

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, authRolePermissionDBTypes, true, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.AuthPermissionID = local.AuthPermissionID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.AuthRolePermission(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.AuthPermissionID != foreign.AuthPermissionID {
		t.Errorf("want: %v, got %v", foreign.AuthPermissionID, check.AuthPermissionID)
	}

	slice := AuthPermissionSlice{&local}
	if err = local.L.LoadAuthRolePermission(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.AuthRolePermission == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.AuthRolePermission = nil
	if err = local.L.LoadAuthRolePermission(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.AuthRolePermission == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAuthPermissionOneToOneSetOpAuthRolePermissionUsingAuthRolePermission(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthPermission
	var b, c AuthRolePermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authPermissionDBTypes, false, strmangle.SetComplement(authPermissionPrimaryKeyColumns, authPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, authRolePermissionDBTypes, false, strmangle.SetComplement(authRolePermissionPrimaryKeyColumns, authRolePermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authRolePermissionDBTypes, false, strmangle.SetComplement(authRolePermissionPrimaryKeyColumns, authRolePermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*AuthRolePermission{&b, &c} {
		err = a.SetAuthRolePermission(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.AuthRolePermission != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.AuthPermission != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.AuthPermissionID != x.AuthPermissionID {
			t.Error("foreign key was wrong value", a.AuthPermissionID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.AuthPermissionID))
		reflect.Indirect(reflect.ValueOf(&x.AuthPermissionID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AuthPermissionID != x.AuthPermissionID {
			t.Error("foreign key was wrong value", a.AuthPermissionID, x.AuthPermissionID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testAuthPermissionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authPermission := &AuthPermission{}
	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermission.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authPermission.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAuthPermissionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authPermission := &AuthPermission{}
	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermission.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthPermissionSlice{authPermission}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAuthPermissionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authPermission := &AuthPermission{}
	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermission.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthPermissions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authPermissionDBTypes = map[string]string{"AuthPermissionID": "integer", "CreatedAt": "timestamp with time zone", "Description": "text", "Permission": "character varying", "UpdatedAt": "timestamp with time zone"}
	_                     = bytes.MinRead
)

func testAuthPermissionsUpdate(t *testing.T) {
	t.Parallel()

	if len(authPermissionColumns) == len(authPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authPermission := &AuthPermission{}
	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermission.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	if err = authPermission.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAuthPermissionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authPermissionColumns) == len(authPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authPermission := &AuthPermission{}
	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermission.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authPermission, authPermissionDBTypes, true, authPermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authPermissionColumns, authPermissionPrimaryKeyColumns) {
		fields = authPermissionColumns
	} else {
		fields = strmangle.SetComplement(
			authPermissionColumns,
			authPermissionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(authPermission))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AuthPermissionSlice{authPermission}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAuthPermissionsUpsert(t *testing.T) {
	t.Parallel()

	if len(authPermissionColumns) == len(authPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	authPermission := AuthPermission{}
	if err = randomize.Struct(seed, &authPermission, authPermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authPermission.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthPermission: %s", err)
	}

	count, err := AuthPermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &authPermission, authPermissionDBTypes, false, authPermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	if err = authPermission.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthPermission: %s", err)
	}

	count, err = AuthPermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
