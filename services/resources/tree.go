package resources

import (
	"fmt"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
	"github.com/cryo-management/api/services"
)

// CreateTree persists the request body creating a new object in the database
func CreateTree(r *http.Request) *services.Response {
	tree := models.Tree{}

	return services.Create(r, &tree, "CreateTree", models.TableCoreTrees)
}

// LoadAllTrees return all instances from the object
func LoadAllTrees(r *http.Request) *services.Response {
	trees := []models.Tree{}

	return services.Load(r, &trees, "LoadAllTrees", models.TableCoreTrees, nil)
}

// LoadTree return only one object from the database
func LoadTree(r *http.Request) *services.Response {
	tree := models.Tree{}
	treeID := chi.URLParam(r, "tree_id")
	treeIDColumn := fmt.Sprintf("%s.id", models.TableCoreTrees)
	condition := builder.Equal(treeIDColumn, treeID)

	return services.Load(r, &tree, "LoadTree", models.TableCoreTrees, condition)
}

// UpdateTree updates object data in the database
func UpdateTree(r *http.Request) *services.Response {
	treeID := chi.URLParam(r, "tree_id")
	treeIDColumn := fmt.Sprintf("%s.id", models.TableCoreTrees)
	condition := builder.Equal(treeIDColumn, treeID)
	tree := models.Tree{
		ID: treeID,
	}

	return services.Update(r, &tree, "UpdateTree", models.TableCoreTrees, condition)
}

// DeleteTree deletes object from the database
func DeleteTree(r *http.Request) *services.Response {
	treeID := chi.URLParam(r, "tree_id")
	treeIDColumn := fmt.Sprintf("%s.id", models.TableCoreTrees)
	condition := builder.Equal(treeIDColumn, treeID)

	return services.Remove(r, "DeleteTree", models.TableCoreTrees, condition)
}

// CreateTreeLevel persists the request body creating a new object in the database
func CreateTreeLevel(r *http.Request) *services.Response {
	treeLevel := models.TreeLevel{}

	return services.Create(r, &treeLevel, "CreateTreeLevel", models.TableCoreTreLevels)
}

// LoadAllTreeLevels return all instances from the object
func LoadAllTreeLevels(r *http.Request) *services.Response {
	treeLevels := []models.TreeLevel{}
	treeID := chi.URLParam(r, "tree_id")
	treeIDColumn := fmt.Sprintf("%s.tree_id", models.TableCoreTreLevels)
	condition := builder.Equal(treeIDColumn, treeID)

	return services.Load(r, &treeLevels, "LoadAllTreeLevels", models.TableCoreTreLevels, condition)
}

// LoadTreeLevel return only one object from the database
func LoadTreeLevel(r *http.Request) *services.Response {
	treeLevel := models.TreeLevel{}
	treeLevelID := chi.URLParam(r, "tree_level_id")
	treeLevelIDColumn := fmt.Sprintf("%s.id", models.TableCoreTreLevels)
	condition := builder.Equal(treeLevelIDColumn, treeLevelID)

	return services.Load(r, &treeLevel, "LoadTreeLevel", models.TableCoreTreLevels, condition)
}

// UpdateTreeLevel updates object data in the database
func UpdateTreeLevel(r *http.Request) *services.Response {
	treeLevelID := chi.URLParam(r, "tree_level_id")
	treeLevelIDColumn := fmt.Sprintf("%s.id", models.TableCoreTreLevels)
	condition := builder.Equal(treeLevelIDColumn, treeLevelID)
	treeLevel := models.TreeLevel{
		ID: treeLevelID,
	}

	return services.Update(r, &treeLevel, "UpdateTreeLevel", models.TableCoreTreLevels, condition)
}

// DeleteTreeLevel deletes object from the database
func DeleteTreeLevel(r *http.Request) *services.Response {
	treeLevelID := chi.URLParam(r, "tree_level_id")
	treeLevelIDColumn := fmt.Sprintf("%s.id", models.TableCoreTreLevels)
	condition := builder.Equal(treeLevelIDColumn, treeLevelID)

	return services.Remove(r, "DeleteTreeLevel", models.TableCoreTreLevels, condition)
}

// CreateTreeUnit persists the request body creating a new object in the database
func CreateTreeUnit(r *http.Request) *services.Response {
	treeUnit := models.TreeUnit{}

	return services.Create(r, &treeUnit, "CreateTreeUnit", models.TableCoreTreUnits)
}

// LoadAllTreeUnits return all instances from the object
func LoadAllTreeUnits(r *http.Request) *services.Response {
	treeUnits := []models.TreeUnit{}
	treeID := chi.URLParam(r, "tree_id")
	treeIDColumn := fmt.Sprintf("%s.tree_id", models.TableCoreTreUnits)
	condition := builder.Equal(treeIDColumn, treeID)

	return services.Load(r, &treeUnits, "LoadAllTreeUnits", models.TableCoreTreUnits, condition)
}

// LoadTreeUnit return only one object from the database
func LoadTreeUnit(r *http.Request) *services.Response {
	treeUnit := models.TreeUnit{}
	treeUnitID := chi.URLParam(r, "tree_unit_id")
	treeUnitIDColumn := fmt.Sprintf("%s.id", models.TableCoreTreUnits)
	condition := builder.Equal(treeUnitIDColumn, treeUnitID)

	return services.Load(r, &treeUnit, "LoadTreeUnit", models.TableCoreTreUnits, condition)
}

// UpdateTreeUnit updates object data in the database
func UpdateTreeUnit(r *http.Request) *services.Response {
	treeUnitID := chi.URLParam(r, "tree_unit_id")
	treeUnitIDColumn := fmt.Sprintf("%s.id", models.TableCoreTreUnits)
	condition := builder.Equal(treeUnitIDColumn, treeUnitID)
	treeUnit := models.TreeUnit{
		ID: treeUnitID,
	}

	return services.Update(r, &treeUnit, "UpdateTreeUnit", models.TableCoreTreUnits, condition)
}

// DeleteTreeUnit deletes object from the database
func DeleteTreeUnit(r *http.Request) *services.Response {
	treeUnitID := chi.URLParam(r, "tree_unit_id")
	treeUnitIDColumn := fmt.Sprintf("%s.id", models.TableCoreTreUnits)
	condition := builder.Equal(treeUnitIDColumn, treeUnitID)

	return services.Remove(r, "DeleteTreeUnit", models.TableCoreTreUnits, condition)
}
