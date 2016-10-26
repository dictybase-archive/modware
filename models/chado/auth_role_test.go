package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testAuthRoles(t *testing.T) {
	t.Parallel()

	query := AuthRoles(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAuthRolesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRole := &AuthRole{}
	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authRole.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthRolesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRole := &AuthRole{}
	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthRoles(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := AuthRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthRolesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRole := &AuthRole{}
	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthRoleSlice{authRole}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAuthRolesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRole := &AuthRole{}
	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRole.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AuthRoleExists(tx, authRole.AuthRoleID)
	if err != nil {
		t.Errorf("Unable to check if AuthRole exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthRoleExistsG to return true, but got false.")
	}
}
func testAuthRolesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRole := &AuthRole{}
	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRole.Insert(tx); err != nil {
		t.Error(err)
	}

	authRoleFound, err := FindAuthRole(tx, authRole.AuthRoleID)
	if err != nil {
		t.Error(err)
	}

	if authRoleFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAuthRolesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRole := &AuthRole{}
	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthRoles(tx).Bind(authRole); err != nil {
		t.Error(err)
	}
}

func testAuthRolesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRole := &AuthRole{}
	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := AuthRoles(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthRolesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRoleOne := &AuthRole{}
	authRoleTwo := &AuthRole{}
	if err = randomize.Struct(seed, authRoleOne, authRoleDBTypes, false, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}
	if err = randomize.Struct(seed, authRoleTwo, authRoleDBTypes, false, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRoleOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authRoleTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthRoles(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthRolesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authRoleOne := &AuthRole{}
	authRoleTwo := &AuthRole{}
	if err = randomize.Struct(seed, authRoleOne, authRoleDBTypes, false, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}
	if err = randomize.Struct(seed, authRoleTwo, authRoleDBTypes, false, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRoleOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authRoleTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func authRoleBeforeInsertHook(e boil.Executor, o *AuthRole) error {
	*o = AuthRole{}
	return nil
}

func authRoleAfterInsertHook(e boil.Executor, o *AuthRole) error {
	*o = AuthRole{}
	return nil
}

func authRoleAfterSelectHook(e boil.Executor, o *AuthRole) error {
	*o = AuthRole{}
	return nil
}

func authRoleBeforeUpdateHook(e boil.Executor, o *AuthRole) error {
	*o = AuthRole{}
	return nil
}

func authRoleAfterUpdateHook(e boil.Executor, o *AuthRole) error {
	*o = AuthRole{}
	return nil
}

func authRoleBeforeDeleteHook(e boil.Executor, o *AuthRole) error {
	*o = AuthRole{}
	return nil
}

func authRoleAfterDeleteHook(e boil.Executor, o *AuthRole) error {
	*o = AuthRole{}
	return nil
}

func authRoleBeforeUpsertHook(e boil.Executor, o *AuthRole) error {
	*o = AuthRole{}
	return nil
}

func authRoleAfterUpsertHook(e boil.Executor, o *AuthRole) error {
	*o = AuthRole{}
	return nil
}

func testAuthRolesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &AuthRole{}
	o := &AuthRole{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, authRoleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AuthRole object: %s", err)
	}

	AddAuthRoleHook(boil.BeforeInsertHook, authRoleBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	authRoleBeforeInsertHooks = []AuthRoleHook{}

	AddAuthRoleHook(boil.AfterInsertHook, authRoleAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	authRoleAfterInsertHooks = []AuthRoleHook{}

	AddAuthRoleHook(boil.AfterSelectHook, authRoleAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	authRoleAfterSelectHooks = []AuthRoleHook{}

	AddAuthRoleHook(boil.BeforeUpdateHook, authRoleBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	authRoleBeforeUpdateHooks = []AuthRoleHook{}

	AddAuthRoleHook(boil.AfterUpdateHook, authRoleAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	authRoleAfterUpdateHooks = []AuthRoleHook{}

	AddAuthRoleHook(boil.BeforeDeleteHook, authRoleBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	authRoleBeforeDeleteHooks = []AuthRoleHook{}

	AddAuthRoleHook(boil.AfterDeleteHook, authRoleAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	authRoleAfterDeleteHooks = []AuthRoleHook{}

	AddAuthRoleHook(boil.BeforeUpsertHook, authRoleBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	authRoleBeforeUpsertHooks = []AuthRoleHook{}

	AddAuthRoleHook(boil.AfterUpsertHook, authRoleAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	authRoleAfterUpsertHooks = []AuthRoleHook{}
}
func testAuthRolesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRole := &AuthRole{}
	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthRolesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRole := &AuthRole{}
	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRole.Insert(tx, authRoleColumns...); err != nil {
		t.Error(err)
	}

	count, err := AuthRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthRoleOneToOneAuthRolePermissionUsingAuthRolePermission(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign AuthRolePermission
	var local AuthRole

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, authRolePermissionDBTypes, true, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, authRoleDBTypes, true, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreign.AuthRoleID = local.AuthRoleID
	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.AuthRolePermission(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.AuthRoleID != foreign.AuthRoleID {
		t.Errorf("want: %v, got %v", foreign.AuthRoleID, check.AuthRoleID)
	}

	slice := AuthRoleSlice{&local}
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

func testAuthRoleOneToOneSetOpAuthRolePermissionUsingAuthRolePermission(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthRole
	var b, c AuthRolePermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authRoleDBTypes, false, strmangle.SetComplement(authRolePrimaryKeyColumns, authRoleColumnsWithoutDefault)...); err != nil {
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
		if x.R.AuthRole != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.AuthRoleID != x.AuthRoleID {
			t.Error("foreign key was wrong value", a.AuthRoleID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.AuthRoleID))
		reflect.Indirect(reflect.ValueOf(&x.AuthRoleID)).Set(zero)

		if err = x.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AuthRoleID != x.AuthRoleID {
			t.Error("foreign key was wrong value", a.AuthRoleID, x.AuthRoleID)
		}

		if err = x.Delete(tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testAuthRoleToManyAuthUserRoles(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthRole
	var b, c AuthUserRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authRoleDBTypes, true, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, authUserRoleDBTypes, false, authUserRoleColumnsWithDefault...)
	randomize.Struct(seed, &c, authUserRoleDBTypes, false, authUserRoleColumnsWithDefault...)

	b.AuthRoleID = a.AuthRoleID
	c.AuthRoleID = a.AuthRoleID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	authUserRole, err := a.AuthUserRoles(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range authUserRole {
		if v.AuthRoleID == b.AuthRoleID {
			bFound = true
		}
		if v.AuthRoleID == c.AuthRoleID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AuthRoleSlice{&a}
	if err = a.L.LoadAuthUserRoles(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AuthUserRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AuthUserRoles = nil
	if err = a.L.LoadAuthUserRoles(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AuthUserRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", authUserRole)
	}
}

func testAuthRoleToManyAddOpAuthUserRoles(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthRole
	var b, c, d, e AuthUserRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authRoleDBTypes, false, strmangle.SetComplement(authRolePrimaryKeyColumns, authRoleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AuthUserRole{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, authUserRoleDBTypes, false, strmangle.SetComplement(authUserRolePrimaryKeyColumns, authUserRoleColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*AuthUserRole{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAuthUserRoles(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.AuthRoleID != first.AuthRoleID {
			t.Error("foreign key was wrong value", a.AuthRoleID, first.AuthRoleID)
		}
		if a.AuthRoleID != second.AuthRoleID {
			t.Error("foreign key was wrong value", a.AuthRoleID, second.AuthRoleID)
		}

		if first.R.AuthRole != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.AuthRole != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AuthUserRoles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AuthUserRoles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AuthUserRoles(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testAuthRolesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRole := &AuthRole{}
	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authRole.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAuthRolesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRole := &AuthRole{}
	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthRoleSlice{authRole}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAuthRolesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRole := &AuthRole{}
	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthRoles(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authRoleDBTypes = map[string]string{"AuthRoleID": "integer", "CreatedAt": "timestamp with time zone", "Description": "text", "Role": "character varying", "UpdatedAt": "timestamp with time zone"}
	_               = bytes.MinRead
)

func testAuthRolesUpdate(t *testing.T) {
	t.Parallel()

	if len(authRoleColumns) == len(authRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authRole := &AuthRole{}
	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	if err = authRole.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAuthRolesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authRoleColumns) == len(authRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authRole := &AuthRole{}
	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authRole, authRoleDBTypes, true, authRolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authRoleColumns, authRolePrimaryKeyColumns) {
		fields = authRoleColumns
	} else {
		fields = strmangle.SetComplement(
			authRoleColumns,
			authRolePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(authRole))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AuthRoleSlice{authRole}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAuthRolesUpsert(t *testing.T) {
	t.Parallel()

	if len(authRoleColumns) == len(authRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	authRole := AuthRole{}
	if err = randomize.Struct(seed, &authRole, authRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRole.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthRole: %s", err)
	}

	count, err := AuthRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &authRole, authRoleDBTypes, false, authRolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	if err = authRole.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthRole: %s", err)
	}

	count, err = AuthRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
