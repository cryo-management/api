package services

import (
	"fmt"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
)

// CreateTree persists the request body creating a new object in the database
func CreateTree(r *http.Request) *Response {
	tree := models.Tree{}

	return create(r, &tree, "CreateTree", models.TableCoreTrees)
}

// LoadAllTrees return all instances from the object
func LoadAllTrees(r *http.Request) *Response {
	trees := []models.Tree{}

	return load(r, &trees, "LoadAllTrees", models.TableCoreTrees, nil)
}

// LoadTree return only one object from the database
func LoadTree(r *http.Request) *Response {
	tree := models.Tree{}
	treeID := chi.URLParam(r, "tree_id")
	treeIDColumn := fmt.Sprintf("%s.id", models.TableCoreTrees)
	condition := builder.Equal(treeIDColumn, treeID)

	return load(r, &tree, "LoadTree", models.TableCoreTrees, condition)
}

// UpdateTree updates object data in the database
func UpdateTree(r *http.Request) *Response {
	treeID := chi.URLParam(r, "tree_id")
	treeIDColumn := fmt.Sprintf("%s.id", models.TableCoreTrees)
	condition := builder.Equal(treeIDColumn, treeID)
	tree := models.Tree{
		ID: treeID,
	}

	return update(r, &tree, "UpdateTree", models.TableCoreTrees, condition)
}

// DeleteTree deletes object from the database
func DeleteTree(r *http.Request) *Response {
	treeID := chi.URLParam(r, "tree_id")
	treeIDColumn := fmt.Sprintf("%s.id", models.TableCoreTrees)
	condition := builder.Equal(treeIDColumn, treeID)

	return remove(r, "DeleteTree", models.TableCoreTrees, condition)
}

// CreateTreeLevel persists the request body creating a new object in the database
func CreateTreeLevel(r *http.Request) *Response {
	treeLevel := models.TreeLevel{}

	return create(r, &treeLevel, "CreateTreeLevel", models.TableCoreTreLevels)
}

// LoadAllTreeLevels return all instances from the object
func LoadAllTreeLevels(r *http.Request) *Response {
	treeLevels := []models.TreeLevel{}
	treeID := chi.URLParam(r, "tree_id")
	treeIDColumn := fmt.Sprintf("%s.tree_id", models.TableCoreTreLevels)
	condition := builder.Equal(treeIDColumn, treeID)

	return load(r, &treeLevels, "LoadAllTreeLevels", models.TableCoreTreLevels, condition)
}

// LoadTreeLevel return only one object from the database
func LoadTreeLevel(r *http.Request) *Response {
	treeLevel := models.TreeLevel{}
	treeLevelID := chi.URLParam(r, "tree_level_id")
	treeLevelIDColumn := fmt.Sprintf("%s.id", models.TableCoreTreLevels)
	condition := builder.Equal(treeLevelIDColumn, treeLevelID)

	return load(r, &treeLevel, "LoadTreeLevel", models.TableCoreTreLevels, condition)
}

// UpdateTreeLevel updates object data in the database
func UpdateTreeLevel(r *http.Request) *Response {
	treeLevelID := chi.URLParam(r, "tree_level_id")
	treeLevelIDColumn := fmt.Sprintf("%s.id", models.TableCoreTreLevels)
	condition := builder.Equal(treeLevelIDColumn, treeLevelID)
	treeLevel := models.TreeLevel{
		ID: treeLevelID,
	}

	return update(r, &treeLevel, "UpdateTreeLevel", models.TableCoreTreLevels, condition)
}

// DeleteTreeLevel deletes object from the database
func DeleteTreeLevel(r *http.Request) *Response {
	treeLevelID := chi.URLParam(r, "tree_level_id")
	treeLevelIDColumn := fmt.Sprintf("%s.id", models.TableCoreTreLevels)
	condition := builder.Equal(treeLevelIDColumn, treeLevelID)

	return remove(r, "DeleteTreeLevel", models.TableCoreTreLevels, condition)
}

// CreateTreeUnit persists the request body creating a new object in the database
func CreateTreeUnit(r *http.Request) *Response {
	treeUnit := models.TreeUnit{}

	return create(r, &treeUnit, "CreateTreeUnit", models.TableCoreTreUnits)
}

// LoadAllTreeUnits return all instances from the object
func LoadAllTreeUnits(r *http.Request) *Response {
	treeUnits := []models.TreeUnit{}
	treeID := chi.URLParam(r, "tree_id")
	treeIDColumn := fmt.Sprintf("%s.tree_id", models.TableCoreTreUnits)
	condition := builder.Equal(treeIDColumn, treeID)

	return load(r, &treeUnits, "LoadAllTreeUnits", models.TableCoreTreUnits, condition)
}

// LoadTreeUnit return only one object from the database
func LoadTreeUnit(r *http.Request) *Response {
	treeUnit := models.TreeUnit{}
	treeUnitID := chi.URLParam(r, "tree_unit_id")
	treeUnitIDColumn := fmt.Sprintf("%s.id", models.TableCoreTreUnits)
	condition := builder.Equal(treeUnitIDColumn, treeUnitID)

	return load(r, &treeUnit, "LoadTreeUnit", models.TableCoreTreUnits, condition)
}

// UpdateTreeUnit updates object data in the database
func UpdateTreeUnit(r *http.Request) *Response {
	treeUnitID := chi.URLParam(r, "tree_unit_id")
	treeUnitIDColumn := fmt.Sprintf("%s.id", models.TableCoreTreUnits)
	condition := builder.Equal(treeUnitIDColumn, treeUnitID)
	treeUnit := models.TreeUnit{
		ID: treeUnitID,
	}

	return update(r, &treeUnit, "UpdateTreeUnit", models.TableCoreTreUnits, condition)
}

// DeleteTreeUnit deletes object from the database
func DeleteTreeUnit(r *http.Request) *Response {
	treeUnitID := chi.URLParam(r, "tree_unit_id")
	treeUnitIDColumn := fmt.Sprintf("%s.id", models.TableCoreTreUnits)
	condition := builder.Equal(treeUnitIDColumn, treeUnitID)

	return remove(r, "DeleteTreeUnit", models.TableCoreTreUnits, condition)
}
