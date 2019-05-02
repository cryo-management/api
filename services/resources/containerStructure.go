package resources

import (
	"fmt"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
	"github.com/cryo-management/api/services"
)

// CreateContainerStructure persists the request body creating a new object in the database
func CreateContainerStructure(r *http.Request) *services.Response {
	containerStructure := models.ContainerStructure{}

	return services.Create(r, &containerStructure, "CreateContainerStructure", models.TableCoreSchPagCntStructures)
}

// LoadAllContainerStructures return all instances from the object
func LoadAllContainerStructures(r *http.Request) *services.Response {
	containerStructures := []models.ContainerStructure{}
	containerID := chi.URLParam(r, "container_id")
	containerIDColumn := fmt.Sprintf("%s.container_id", models.TableCoreSchPagCntStructures)
	condition := builder.Equal(containerIDColumn, containerID)

	return services.Load(r, &containerStructures, "LoadAllContainerStructures", models.TableCoreSchPagCntStructures, condition)
}

// LoadContainerStructure return only one object from the database
func LoadContainerStructure(r *http.Request) *services.Response {
	containerStructure := models.ContainerStructure{}
	containerStructureID := chi.URLParam(r, "container_structure_id")
	containerStructureIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPagCntStructures)
	condition := builder.Equal(containerStructureIDColumn, containerStructureID)

	return services.Load(r, &containerStructure, "LoadContainerStructure", models.TableCoreSchPagCntStructures, condition)
}

// UpdateContainerStructure updates object data in the database
func UpdateContainerStructure(r *http.Request) *services.Response {
	containerStructureID := chi.URLParam(r, "container_structure_id")
	containerStructureIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPagCntStructures)
	condition := builder.Equal(containerStructureIDColumn, containerStructureID)
	containerStructure := models.ContainerStructure{
		ID: containerStructureID,
	}

	return services.Update(r, &containerStructure, "UpdateContainerStructure", models.TableCoreSchPagCntStructures, condition)
}

// DeleteContainerStructure deletes object from the database
func DeleteContainerStructure(r *http.Request) *services.Response {
	containerStructureID := chi.URLParam(r, "container_structure_id")
	containerStructureIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPagCntStructures)
	condition := builder.Equal(containerStructureIDColumn, containerStructureID)

	return services.Remove(r, "DeleteContainerStructure", models.TableCoreSchPagCntStructures, condition)
}
