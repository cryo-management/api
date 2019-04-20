package admin

import (
	controller "github.com/cryo-management/api/controllers/admin"
	"github.com/go-chi/chi"
)

// StructureRoutes creates the api methods
func StructureRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/structures
	r.Route("/", func(r chi.Router) {
		r.Post("/{structure_id}/views", controller.PostView)
		r.Get("/{structure_id}/views", controller.GetAllViews)
		r.Get("/{structure_id}/views/{view_id}", controller.GetView)
		r.Patch("/{structure_id}/views/{view_id}", controller.UpdateView)
		r.Delete("/{structure_id}/views/{view_id}", controller.DeleteView)
		r.Post("/{structure_id}/views/{view_id}/sections", controller.PostSection)
		r.Get("/{structure_id}/views/{view_id}/sections", controller.GetAllSections)
		r.Get("/{structure_id}/views/{view_id}/sections/{section_id}", controller.GetSection)
		r.Patch("/{structure_id}/views/{view_id}/sections/{section_id}", controller.UpdateSection)
		r.Delete("/{structure_id}/views/{view_id}/sections/{section_id}", controller.DeleteSection)
		r.Post("/{structure_id}/views/{view_id}/sections/{section_id}/section_structures", controller.PostSectionStructure)
		r.Get("/{structure_id}/views/{view_id}/sections/{section_id}/section_structures", controller.GetAllSectionStructures)
		r.Get("/{structure_id}/views/{view_id}/sections/{section_id}/section_structures/{section_structure_id}", controller.GetSectionStructure)
		r.Patch("/{structure_id}/views/{view_id}/sections/{section_id}/section_structures/{section_structure_id}", controller.UpdateSectionStructure)
		r.Delete("/{structure_id}/views/{view_id}/sections/{section_id}/section_structures/{section_structure_id}", controller.DeleteSectionStructure)
	})

	return r
}
