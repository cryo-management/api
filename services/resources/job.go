package resources

import (
	"fmt"
	"net/http"
	"time"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/andreluzz/go-sql-builder/db"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
	"github.com/cryo-management/api/services"
)

// CreateJob persists the request body creating a new object in the database
func CreateJob(r *http.Request) *services.Response {
	job := models.Job{}

	return services.Create(r, &job, "CreateJob", models.TableCoreJobs)
}

// LoadAllJobs return all instances from the object
func LoadAllJobs(r *http.Request) *services.Response {
	jobs := []models.Job{}

	return services.Load(r, &jobs, "LoadAllJobs", models.TableCoreJobs, nil)
}

// LoadJob return only one object from the database
func LoadJob(r *http.Request) *services.Response {
	job := models.Job{}
	jobID := chi.URLParam(r, "job_id")
	jobIDColumn := fmt.Sprintf("%s.id", models.TableCoreJobs)
	condition := builder.Equal(jobIDColumn, jobID)

	return services.Load(r, &job, "LoadJob", models.TableCoreJobs, condition)
}

// UpdateJob updates object data in the database
func UpdateJob(r *http.Request) *services.Response {
	jobID := chi.URLParam(r, "job_id")
	jobIDColumn := fmt.Sprintf("%s.id", models.TableCoreJobs)
	condition := builder.Equal(jobIDColumn, jobID)
	job := models.Job{
		ID: jobID,
	}

	return services.Update(r, &job, "UpdateJob", models.TableCoreJobs, condition)
}

// DeleteJob deletes object from the database
func DeleteJob(r *http.Request) *services.Response {
	jobID := chi.URLParam(r, "job_id")
	jobIDColumn := fmt.Sprintf("%s.id", models.TableCoreJobs)
	condition := builder.Equal(jobIDColumn, jobID)

	return services.Remove(r, "DeleteJob", models.TableCoreJobs, condition)
}

// CreateJobTask persists the request body creating a new object in the database
func CreateJobTask(r *http.Request) *services.Response {
	jobTask := models.JobTask{}

	return services.Create(r, &jobTask, "CreateJobTask", models.TableCoreJobTasks)
}

// LoadAllJobTasks return all instances from the object
func LoadAllJobTasks(r *http.Request) *services.Response {
	jobTasks := []models.JobTask{}
	jobTaskID := chi.URLParam(r, "job_id")
	jobTaskIDColumn := fmt.Sprintf("%s.job_id", models.TableCoreJobTasks)
	condition := builder.Equal(jobTaskIDColumn, jobTaskID)

	return services.Load(r, &jobTasks, "LoadAllJobTasks", models.TableCoreJobTasks, condition)
}

// LoadJobTask return only one object from the database
func LoadJobTask(r *http.Request) *services.Response {
	jobTask := models.JobTask{}
	jobTaskID := chi.URLParam(r, "job_task_id")
	jobTaskIDColumn := fmt.Sprintf("%s.id", models.TableCoreJobTasks)
	condition := builder.Equal(jobTaskIDColumn, jobTaskID)

	return services.Load(r, &jobTask, "LoadJobTask", models.TableCoreJobTasks, condition)
}

// UpdateJobTask updates object data in the database
func UpdateJobTask(r *http.Request) *services.Response {
	jobTaskID := chi.URLParam(r, "job_task_id")
	jobTaskIDColumn := fmt.Sprintf("%s.id", models.TableCoreJobTasks)
	condition := builder.Equal(jobTaskIDColumn, jobTaskID)
	jobTask := models.JobTask{
		ID: jobTaskID,
	}

	return services.Update(r, &jobTask, "UpdateJobTask", models.TableCoreJobTasks, condition)
}

// DeleteJobTask deletes object from the database
func DeleteJobTask(r *http.Request) *services.Response {
	jobTaskID := chi.URLParam(r, "job_task_id")
	jobTaskIDColumn := fmt.Sprintf("%s.id", models.TableCoreJobTasks)
	condition := builder.Equal(jobTaskIDColumn, jobTaskID)

	return services.Remove(r, "DeleteJobTask", models.TableCoreJobTasks, condition)
}

// LoadAllJobFollowersAvaible return all instances from the object
func LoadAllJobFollowersAvaible(r *http.Request) *services.Response {
	viewFollowersAvailable := []models.ViewFollowerAvailable{}
	activeColumn := fmt.Sprintf("%s.active", models.ViewCoreUsersAndGroups)
	languageCode := r.Header.Get("Content-Language")
	languageCodeColumn := fmt.Sprintf("%s.language_code", models.ViewCoreUsersAndGroups)
	condition := builder.And(
		builder.Equal(activeColumn, true),
		builder.Or(
			builder.Equal(languageCodeColumn, languageCode),
			builder.Equal(languageCodeColumn, nil),
		),
	)

	return services.Load(r, &viewFollowersAvailable, "LoadAllJobFollowersAvaible", models.ViewCoreUsersAndGroups, condition)
}

// InsertFollowerInJob persists the request creating a new object in the database
func InsertFollowerInJob(r *http.Request) *services.Response {
	response := services.NewResponse()

	jobID := chi.URLParam(r, "job_id")
	followerID := chi.URLParam(r, "follower_id")
	followerType := chi.URLParam(r, "follower_type")

	userID := r.Header.Get("userID")
	now := time.Now()

	statemant := builder.Insert(
		models.TableCoreJobsFollowers,
		"job_id",
		"follower_id",
		"follower_type",
		"created_by",
		"created_at",
		"updated_by",
		"updated_at",
	).Values(
		jobID,
		followerID,
		followerType,
		userID,
		now,
		userID,
		now,
	)

	err := db.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorInsertingRecord, "InsertFollowerInJob", err.Error()))

		return response
	}

	return response
}

// LoadAllFollowersByJob return all instances from the object
func LoadAllFollowersByJob(r *http.Request) *services.Response {
	jobFollowers := []models.JobFollowers{}
	jobID := chi.URLParam(r, "job_id")
	jobIDColumn := fmt.Sprintf("%s.job_id", models.ViewCoreJobFollowers)
	languageCode := r.Header.Get("Content-Language")
	languageCodeColumn := fmt.Sprintf("%s.language_code", models.ViewCoreJobFollowers)
	followerTypeColumn := fmt.Sprintf("%s.follower_type", models.ViewCoreJobFollowers)
	condition := builder.And(
		builder.Equal(jobIDColumn, jobID),
		builder.Or(
			builder.Equal(followerTypeColumn, "user"),
			builder.And(
				builder.Equal(followerTypeColumn, "group"),
				builder.Equal(languageCodeColumn, languageCode),
			),
		),
	)

	return services.Load(r, &jobFollowers, "LoadAllFollowersByJob", models.ViewCoreJobFollowers, condition)
}

// RemoveFollowerFromJob deletes object from the database
func RemoveFollowerFromJob(r *http.Request) *services.Response {
	response := services.NewResponse()

	jobID := chi.URLParam(r, "job_id")
	followerID := chi.URLParam(r, "follower_id")

	statemant := builder.Delete(models.TableCoreJobsFollowers).Where(
		builder.And(
			builder.Equal("job_id", jobID),
			builder.Equal("follower_id", followerID),
		),
	)

	err := db.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorDeletingData, "RemoveFollowerFromJob", err.Error()))

		return response
	}

	return response
}
