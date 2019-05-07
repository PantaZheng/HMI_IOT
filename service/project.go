package service

import "github.com/pantazheng/bci/models"

func ProjectCreate(project *models.ProjectJson) (projectJson models.ProjectJson, err error) {
	return models.ProjectCreate(project)
}

func ProjectFindByID(id uint) (recordProjectJson models.ProjectJson, err error) {
	project := new(models.Project)
	project.ID = id
	return models.ProjectFind(project)
}

func ProjectsFindByLeaderID(id uint) (projects []models.ProjectBriefJson, err error) {
	leader := new(models.User)
	leader.ID = id
	return models.ProjectsFindByLeader(leader)
}

func ProjectsFindByParticipantID(id uint) (projects []models.ProjectBriefJson, err error) {
	participant := new(models.User)
	participant.ID = id
	return models.ProjectsFindByParticipant(participant)
}

func ProjectUpdate(projectJson *models.ProjectJson) (recordProjectJson models.ProjectJson, err error) {
	return models.ProjectUpdate(projectJson)
}

func ProjectDeleteByID(id uint) (recordProjectJson models.ProjectJson, err error) {
	project := new(models.Project)
	project.ID = id
	return models.ProjectDelete(project)
}
