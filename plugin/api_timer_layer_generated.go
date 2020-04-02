// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

// Code generated by "make pluginapi"
// DO NOT EDIT

package plugin

import (
	"io"
	"net/http"
	timePkg "time"

	"github.com/mattermost/mattermost-server/v5/einterfaces"
	"github.com/mattermost/mattermost-server/v5/model"
)

type apiTimerLayer struct {
	pluginID string
	apiImpl  API
	metrics  einterfaces.MetricsInterface
}

func (api *apiTimerLayer) recordTime(startTime timePkg.Time, name string, success bool) {
	if api.metrics != nil {
		elapsedTime := float64(timePkg.Since(startTime)) / float64(timePkg.Second)
		api.metrics.ObservePluginApiDuration(api.pluginID, name, success, elapsedTime)
	}
}

func (api *apiTimerLayer) LoadPluginConfiguration(dest interface{}) error {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.LoadPluginConfiguration(dest)
	api.recordTime(startTime, "LoadPluginConfiguration", true)
	return _returnsA
}

func (api *apiTimerLayer) RegisterCommand(command *model.Command) error {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.RegisterCommand(command)
	api.recordTime(startTime, "RegisterCommand", true)
	return _returnsA
}

func (api *apiTimerLayer) UnregisterCommand(teamId, trigger string) error {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.UnregisterCommand(teamId, trigger)
	api.recordTime(startTime, "UnregisterCommand", true)
	return _returnsA
}

func (api *apiTimerLayer) GetSession(sessionId string) (*model.Session, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetSession(sessionId)
	api.recordTime(startTime, "GetSession", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetConfig() *model.Config {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.GetConfig()
	api.recordTime(startTime, "GetConfig", true)
	return _returnsA
}

func (api *apiTimerLayer) GetUnsanitizedConfig() *model.Config {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.GetUnsanitizedConfig()
	api.recordTime(startTime, "GetUnsanitizedConfig", true)
	return _returnsA
}

func (api *apiTimerLayer) SaveConfig(config *model.Config) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.SaveConfig(config)
	api.recordTime(startTime, "SaveConfig", true)
	return _returnsA
}

func (api *apiTimerLayer) GetPluginConfig() map[string]interface{} {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.GetPluginConfig()
	api.recordTime(startTime, "GetPluginConfig", true)
	return _returnsA
}

func (api *apiTimerLayer) SavePluginConfig(config map[string]interface{}) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.SavePluginConfig(config)
	api.recordTime(startTime, "SavePluginConfig", true)
	return _returnsA
}

func (api *apiTimerLayer) GetBundlePath() (string, error) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetBundlePath()
	api.recordTime(startTime, "GetBundlePath", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetLicense() *model.License {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.GetLicense()
	api.recordTime(startTime, "GetLicense", true)
	return _returnsA
}

func (api *apiTimerLayer) GetServerVersion() string {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.GetServerVersion()
	api.recordTime(startTime, "GetServerVersion", true)
	return _returnsA
}

func (api *apiTimerLayer) GetSystemInstallDate() (int64, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetSystemInstallDate()
	api.recordTime(startTime, "GetSystemInstallDate", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetDiagnosticId() string {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.GetDiagnosticId()
	api.recordTime(startTime, "GetDiagnosticId", true)
	return _returnsA
}

func (api *apiTimerLayer) CreateUser(user *model.User) (*model.User, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.CreateUser(user)
	api.recordTime(startTime, "CreateUser", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) DeleteUser(userId string) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.DeleteUser(userId)
	api.recordTime(startTime, "DeleteUser", true)
	return _returnsA
}

func (api *apiTimerLayer) GetUsers(options *model.UserGetOptions) ([]*model.User, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetUsers(options)
	api.recordTime(startTime, "GetUsers", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetUser(userId string) (*model.User, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetUser(userId)
	api.recordTime(startTime, "GetUser", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetUserByEmail(email string) (*model.User, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetUserByEmail(email)
	api.recordTime(startTime, "GetUserByEmail", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetUserByUsername(name string) (*model.User, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetUserByUsername(name)
	api.recordTime(startTime, "GetUserByUsername", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetUsersByUsernames(usernames []string) ([]*model.User, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetUsersByUsernames(usernames)
	api.recordTime(startTime, "GetUsersByUsernames", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetUsersInTeam(teamId string, page int, perPage int) ([]*model.User, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetUsersInTeam(teamId, page, perPage)
	api.recordTime(startTime, "GetUsersInTeam", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetTeamIcon(teamId string) ([]byte, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetTeamIcon(teamId)
	api.recordTime(startTime, "GetTeamIcon", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) SetTeamIcon(teamId string, data []byte) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.SetTeamIcon(teamId, data)
	api.recordTime(startTime, "SetTeamIcon", true)
	return _returnsA
}

func (api *apiTimerLayer) RemoveTeamIcon(teamId string) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.RemoveTeamIcon(teamId)
	api.recordTime(startTime, "RemoveTeamIcon", true)
	return _returnsA
}

func (api *apiTimerLayer) UpdateUser(user *model.User) (*model.User, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.UpdateUser(user)
	api.recordTime(startTime, "UpdateUser", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetUserStatus(userId string) (*model.Status, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetUserStatus(userId)
	api.recordTime(startTime, "GetUserStatus", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetUserStatusesByIds(userIds []string) ([]*model.Status, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetUserStatusesByIds(userIds)
	api.recordTime(startTime, "GetUserStatusesByIds", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) UpdateUserStatus(userId, status string) (*model.Status, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.UpdateUserStatus(userId, status)
	api.recordTime(startTime, "UpdateUserStatus", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) UpdateUserActive(userId string, active bool) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.UpdateUserActive(userId, active)
	api.recordTime(startTime, "UpdateUserActive", true)
	return _returnsA
}

func (api *apiTimerLayer) GetUsersInChannel(channelId, sortBy string, page, perPage int) ([]*model.User, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetUsersInChannel(channelId, sortBy, page, perPage)
	api.recordTime(startTime, "GetUsersInChannel", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetLDAPUserAttributes(userId string, attributes []string) (map[string]string, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetLDAPUserAttributes(userId, attributes)
	api.recordTime(startTime, "GetLDAPUserAttributes", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) CreateTeam(team *model.Team) (*model.Team, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.CreateTeam(team)
	api.recordTime(startTime, "CreateTeam", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) DeleteTeam(teamId string) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.DeleteTeam(teamId)
	api.recordTime(startTime, "DeleteTeam", true)
	return _returnsA
}

func (api *apiTimerLayer) GetTeams() ([]*model.Team, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetTeams()
	api.recordTime(startTime, "GetTeams", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetTeam(teamId string) (*model.Team, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetTeam(teamId)
	api.recordTime(startTime, "GetTeam", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetTeamByName(name string) (*model.Team, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetTeamByName(name)
	api.recordTime(startTime, "GetTeamByName", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetTeamsUnreadForUser(userId string) ([]*model.TeamUnread, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetTeamsUnreadForUser(userId)
	api.recordTime(startTime, "GetTeamsUnreadForUser", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) UpdateTeam(team *model.Team) (*model.Team, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.UpdateTeam(team)
	api.recordTime(startTime, "UpdateTeam", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) SearchTeams(term string) ([]*model.Team, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.SearchTeams(term)
	api.recordTime(startTime, "SearchTeams", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetTeamsForUser(userId string) ([]*model.Team, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetTeamsForUser(userId)
	api.recordTime(startTime, "GetTeamsForUser", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) CreateTeamMember(teamId, userId string) (*model.TeamMember, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.CreateTeamMember(teamId, userId)
	api.recordTime(startTime, "CreateTeamMember", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) CreateTeamMembers(teamId string, userIds []string, requestorId string) ([]*model.TeamMember, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.CreateTeamMembers(teamId, userIds, requestorId)
	api.recordTime(startTime, "CreateTeamMembers", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) CreateTeamMembersGracefully(teamId string, userIds []string, requestorId string) ([]*model.TeamMemberWithError, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.CreateTeamMembersGracefully(teamId, userIds, requestorId)
	api.recordTime(startTime, "CreateTeamMembersGracefully", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) DeleteTeamMember(teamId, userId, requestorId string) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.DeleteTeamMember(teamId, userId, requestorId)
	api.recordTime(startTime, "DeleteTeamMember", true)
	return _returnsA
}

func (api *apiTimerLayer) GetTeamMembers(teamId string, page, perPage int) ([]*model.TeamMember, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetTeamMembers(teamId, page, perPage)
	api.recordTime(startTime, "GetTeamMembers", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetTeamMember(teamId, userId string) (*model.TeamMember, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetTeamMember(teamId, userId)
	api.recordTime(startTime, "GetTeamMember", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetTeamMembersForUser(userId string, page int, perPage int) ([]*model.TeamMember, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetTeamMembersForUser(userId, page, perPage)
	api.recordTime(startTime, "GetTeamMembersForUser", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) UpdateTeamMemberRoles(teamId, userId, newRoles string) (*model.TeamMember, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.UpdateTeamMemberRoles(teamId, userId, newRoles)
	api.recordTime(startTime, "UpdateTeamMemberRoles", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) CreateChannel(channel *model.Channel) (*model.Channel, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.CreateChannel(channel)
	api.recordTime(startTime, "CreateChannel", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) DeleteChannel(channelId string) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.DeleteChannel(channelId)
	api.recordTime(startTime, "DeleteChannel", true)
	return _returnsA
}

func (api *apiTimerLayer) GetPublicChannelsForTeam(teamId string, page, perPage int) ([]*model.Channel, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetPublicChannelsForTeam(teamId, page, perPage)
	api.recordTime(startTime, "GetPublicChannelsForTeam", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetChannel(channelId string) (*model.Channel, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetChannel(channelId)
	api.recordTime(startTime, "GetChannel", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetChannelByName(teamId, name string, includeDeleted bool) (*model.Channel, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetChannelByName(teamId, name, includeDeleted)
	api.recordTime(startTime, "GetChannelByName", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetChannelByNameForTeamName(teamName, channelName string, includeDeleted bool) (*model.Channel, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetChannelByNameForTeamName(teamName, channelName, includeDeleted)
	api.recordTime(startTime, "GetChannelByNameForTeamName", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetChannelsForTeamForUser(teamId, userId string, includeDeleted bool) ([]*model.Channel, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetChannelsForTeamForUser(teamId, userId, includeDeleted)
	api.recordTime(startTime, "GetChannelsForTeamForUser", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetChannelStats(channelId string) (*model.ChannelStats, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetChannelStats(channelId)
	api.recordTime(startTime, "GetChannelStats", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetDirectChannel(userId1, userId2 string) (*model.Channel, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetDirectChannel(userId1, userId2)
	api.recordTime(startTime, "GetDirectChannel", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetGroupChannel(userIds []string) (*model.Channel, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetGroupChannel(userIds)
	api.recordTime(startTime, "GetGroupChannel", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) UpdateChannel(channel *model.Channel) (*model.Channel, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.UpdateChannel(channel)
	api.recordTime(startTime, "UpdateChannel", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) SearchChannels(teamId string, term string) ([]*model.Channel, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.SearchChannels(teamId, term)
	api.recordTime(startTime, "SearchChannels", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) SearchUsers(search *model.UserSearch) ([]*model.User, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.SearchUsers(search)
	api.recordTime(startTime, "SearchUsers", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) SearchPostsInTeam(teamId string, paramsList []*model.SearchParams) ([]*model.Post, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.SearchPostsInTeam(teamId, paramsList)
	api.recordTime(startTime, "SearchPostsInTeam", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) AddChannelMember(channelId, userId string) (*model.ChannelMember, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.AddChannelMember(channelId, userId)
	api.recordTime(startTime, "AddChannelMember", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) AddUserToChannel(channelId, userId, asUserId string) (*model.ChannelMember, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.AddUserToChannel(channelId, userId, asUserId)
	api.recordTime(startTime, "AddUserToChannel", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetChannelMember(channelId, userId string) (*model.ChannelMember, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetChannelMember(channelId, userId)
	api.recordTime(startTime, "GetChannelMember", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetChannelMembers(channelId string, page, perPage int) (*model.ChannelMembers, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetChannelMembers(channelId, page, perPage)
	api.recordTime(startTime, "GetChannelMembers", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetChannelMembersByIds(channelId string, userIds []string) (*model.ChannelMembers, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetChannelMembersByIds(channelId, userIds)
	api.recordTime(startTime, "GetChannelMembersByIds", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetChannelMembersForUser(teamId, userId string, page, perPage int) ([]*model.ChannelMember, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetChannelMembersForUser(teamId, userId, page, perPage)
	api.recordTime(startTime, "GetChannelMembersForUser", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) UpdateChannelMemberRoles(channelId, userId, newRoles string) (*model.ChannelMember, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.UpdateChannelMemberRoles(channelId, userId, newRoles)
	api.recordTime(startTime, "UpdateChannelMemberRoles", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) UpdateChannelMemberNotifications(channelId, userId string, notifications map[string]string) (*model.ChannelMember, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.UpdateChannelMemberNotifications(channelId, userId, notifications)
	api.recordTime(startTime, "UpdateChannelMemberNotifications", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetGroup(groupId string) (*model.Group, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetGroup(groupId)
	api.recordTime(startTime, "GetGroup", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetGroupByName(name string) (*model.Group, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetGroupByName(name)
	api.recordTime(startTime, "GetGroupByName", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetGroupsForUser(userId string) ([]*model.Group, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetGroupsForUser(userId)
	api.recordTime(startTime, "GetGroupsForUser", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) DeleteChannelMember(channelId, userId string) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.DeleteChannelMember(channelId, userId)
	api.recordTime(startTime, "DeleteChannelMember", true)
	return _returnsA
}

func (api *apiTimerLayer) CreatePost(post *model.Post) (*model.Post, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.CreatePost(post)
	api.recordTime(startTime, "CreatePost", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) AddReaction(reaction *model.Reaction) (*model.Reaction, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.AddReaction(reaction)
	api.recordTime(startTime, "AddReaction", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) RemoveReaction(reaction *model.Reaction) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.RemoveReaction(reaction)
	api.recordTime(startTime, "RemoveReaction", true)
	return _returnsA
}

func (api *apiTimerLayer) GetReactions(postId string) ([]*model.Reaction, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetReactions(postId)
	api.recordTime(startTime, "GetReactions", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) SendEphemeralPost(userId string, post *model.Post) *model.Post {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.SendEphemeralPost(userId, post)
	api.recordTime(startTime, "SendEphemeralPost", true)
	return _returnsA
}

func (api *apiTimerLayer) UpdateEphemeralPost(userId string, post *model.Post) *model.Post {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.UpdateEphemeralPost(userId, post)
	api.recordTime(startTime, "UpdateEphemeralPost", true)
	return _returnsA
}

func (api *apiTimerLayer) DeleteEphemeralPost(userId, postId string) {
	startTime := timePkg.Now()
	api.apiImpl.DeleteEphemeralPost(userId, postId)
	api.recordTime(startTime, "DeleteEphemeralPost", true)
}

func (api *apiTimerLayer) DeletePost(postId string) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.DeletePost(postId)
	api.recordTime(startTime, "DeletePost", true)
	return _returnsA
}

func (api *apiTimerLayer) GetPostThread(postId string) (*model.PostList, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetPostThread(postId)
	api.recordTime(startTime, "GetPostThread", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetPost(postId string) (*model.Post, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetPost(postId)
	api.recordTime(startTime, "GetPost", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetPostsSince(channelId string, time int64) (*model.PostList, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetPostsSince(channelId, time)
	api.recordTime(startTime, "GetPostsSince", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetPostsAfter(channelId, postId string, page, perPage int) (*model.PostList, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetPostsAfter(channelId, postId, page, perPage)
	api.recordTime(startTime, "GetPostsAfter", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetPostsBefore(channelId, postId string, page, perPage int) (*model.PostList, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetPostsBefore(channelId, postId, page, perPage)
	api.recordTime(startTime, "GetPostsBefore", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetPostsForChannel(channelId string, page, perPage int) (*model.PostList, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetPostsForChannel(channelId, page, perPage)
	api.recordTime(startTime, "GetPostsForChannel", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetTeamStats(teamId string) (*model.TeamStats, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetTeamStats(teamId)
	api.recordTime(startTime, "GetTeamStats", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) UpdatePost(post *model.Post) (*model.Post, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.UpdatePost(post)
	api.recordTime(startTime, "UpdatePost", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetProfileImage(userId string) ([]byte, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetProfileImage(userId)
	api.recordTime(startTime, "GetProfileImage", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) SetProfileImage(userId string, data []byte) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.SetProfileImage(userId, data)
	api.recordTime(startTime, "SetProfileImage", true)
	return _returnsA
}

func (api *apiTimerLayer) GetEmojiList(sortBy string, page, perPage int) ([]*model.Emoji, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetEmojiList(sortBy, page, perPage)
	api.recordTime(startTime, "GetEmojiList", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetEmojiByName(name string) (*model.Emoji, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetEmojiByName(name)
	api.recordTime(startTime, "GetEmojiByName", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetEmoji(emojiId string) (*model.Emoji, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetEmoji(emojiId)
	api.recordTime(startTime, "GetEmoji", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) CopyFileInfos(userId string, fileIds []string) ([]string, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.CopyFileInfos(userId, fileIds)
	api.recordTime(startTime, "CopyFileInfos", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetFileInfo(fileId string) (*model.FileInfo, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetFileInfo(fileId)
	api.recordTime(startTime, "GetFileInfo", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetFileInfos(page, perPage int, opt *model.GetFileInfosOptions) ([]*model.FileInfo, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetFileInfos(page, perPage, opt)
	api.recordTime(startTime, "GetFileInfos", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetFile(fileId string) ([]byte, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetFile(fileId)
	api.recordTime(startTime, "GetFile", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetFileLink(fileId string) (string, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetFileLink(fileId)
	api.recordTime(startTime, "GetFileLink", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) ReadFile(path string) ([]byte, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.ReadFile(path)
	api.recordTime(startTime, "ReadFile", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetEmojiImage(emojiId string) ([]byte, string, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB, _returnsC := api.apiImpl.GetEmojiImage(emojiId)
	api.recordTime(startTime, "GetEmojiImage", true)
	return _returnsA, _returnsB, _returnsC
}

func (api *apiTimerLayer) UploadFile(data []byte, channelId string, filename string) (*model.FileInfo, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.UploadFile(data, channelId, filename)
	api.recordTime(startTime, "UploadFile", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) OpenInteractiveDialog(dialog model.OpenDialogRequest) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.OpenInteractiveDialog(dialog)
	api.recordTime(startTime, "OpenInteractiveDialog", true)
	return _returnsA
}

func (api *apiTimerLayer) GetPlugins() ([]*model.Manifest, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetPlugins()
	api.recordTime(startTime, "GetPlugins", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) EnablePlugin(id string) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.EnablePlugin(id)
	api.recordTime(startTime, "EnablePlugin", true)
	return _returnsA
}

func (api *apiTimerLayer) DisablePlugin(id string) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.DisablePlugin(id)
	api.recordTime(startTime, "DisablePlugin", true)
	return _returnsA
}

func (api *apiTimerLayer) RemovePlugin(id string) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.RemovePlugin(id)
	api.recordTime(startTime, "RemovePlugin", true)
	return _returnsA
}

func (api *apiTimerLayer) GetPluginStatus(id string) (*model.PluginStatus, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetPluginStatus(id)
	api.recordTime(startTime, "GetPluginStatus", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) InstallPlugin(file io.Reader, replace bool) (*model.Manifest, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.InstallPlugin(file, replace)
	api.recordTime(startTime, "InstallPlugin", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) KVSet(key string, value []byte) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.KVSet(key, value)
	api.recordTime(startTime, "KVSet", true)
	return _returnsA
}

func (api *apiTimerLayer) KVCompareAndSet(key string, oldValue, newValue []byte) (bool, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.KVCompareAndSet(key, oldValue, newValue)
	api.recordTime(startTime, "KVCompareAndSet", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) KVCompareAndDelete(key string, oldValue []byte) (bool, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.KVCompareAndDelete(key, oldValue)
	api.recordTime(startTime, "KVCompareAndDelete", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) KVSetWithOptions(key string, value []byte, options model.PluginKVSetOptions) (bool, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.KVSetWithOptions(key, value, options)
	api.recordTime(startTime, "KVSetWithOptions", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) KVSetWithExpiry(key string, value []byte, expireInSeconds int64) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.KVSetWithExpiry(key, value, expireInSeconds)
	api.recordTime(startTime, "KVSetWithExpiry", true)
	return _returnsA
}

func (api *apiTimerLayer) KVGet(key string) ([]byte, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.KVGet(key)
	api.recordTime(startTime, "KVGet", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) KVDelete(key string) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.KVDelete(key)
	api.recordTime(startTime, "KVDelete", true)
	return _returnsA
}

func (api *apiTimerLayer) KVDeleteAll() *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.KVDeleteAll()
	api.recordTime(startTime, "KVDeleteAll", true)
	return _returnsA
}

func (api *apiTimerLayer) KVList(page, perPage int) ([]string, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.KVList(page, perPage)
	api.recordTime(startTime, "KVList", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) PublishWebSocketEvent(event string, payload map[string]interface{}, broadcast *model.WebsocketBroadcast) {
	startTime := timePkg.Now()
	api.apiImpl.PublishWebSocketEvent(event, payload, broadcast)
	api.recordTime(startTime, "PublishWebSocketEvent", true)
}

func (api *apiTimerLayer) HasPermissionTo(userId string, permission *model.Permission) bool {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.HasPermissionTo(userId, permission)
	api.recordTime(startTime, "HasPermissionTo", true)
	return _returnsA
}

func (api *apiTimerLayer) HasPermissionToTeam(userId, teamId string, permission *model.Permission) bool {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.HasPermissionToTeam(userId, teamId, permission)
	api.recordTime(startTime, "HasPermissionToTeam", true)
	return _returnsA
}

func (api *apiTimerLayer) HasPermissionToChannel(userId, channelId string, permission *model.Permission) bool {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.HasPermissionToChannel(userId, channelId, permission)
	api.recordTime(startTime, "HasPermissionToChannel", true)
	return _returnsA
}

func (api *apiTimerLayer) LogDebug(msg string, keyValuePairs ...interface{}) {
	startTime := timePkg.Now()
	api.apiImpl.LogDebug(msg, keyValuePairs...)
	api.recordTime(startTime, "LogDebug", true)
}

func (api *apiTimerLayer) LogInfo(msg string, keyValuePairs ...interface{}) {
	startTime := timePkg.Now()
	api.apiImpl.LogInfo(msg, keyValuePairs...)
	api.recordTime(startTime, "LogInfo", true)
}

func (api *apiTimerLayer) LogError(msg string, keyValuePairs ...interface{}) {
	startTime := timePkg.Now()
	api.apiImpl.LogError(msg, keyValuePairs...)
	api.recordTime(startTime, "LogError", true)
}

func (api *apiTimerLayer) LogWarn(msg string, keyValuePairs ...interface{}) {
	startTime := timePkg.Now()
	api.apiImpl.LogWarn(msg, keyValuePairs...)
	api.recordTime(startTime, "LogWarn", true)
}

func (api *apiTimerLayer) SendMail(to, subject, htmlBody string) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.SendMail(to, subject, htmlBody)
	api.recordTime(startTime, "SendMail", true)
	return _returnsA
}

func (api *apiTimerLayer) CreateBot(bot *model.Bot) (*model.Bot, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.CreateBot(bot)
	api.recordTime(startTime, "CreateBot", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) PatchBot(botUserId string, botPatch *model.BotPatch) (*model.Bot, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.PatchBot(botUserId, botPatch)
	api.recordTime(startTime, "PatchBot", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetBot(botUserId string, includeDeleted bool) (*model.Bot, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetBot(botUserId, includeDeleted)
	api.recordTime(startTime, "GetBot", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) GetBots(options *model.BotGetOptions) ([]*model.Bot, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetBots(options)
	api.recordTime(startTime, "GetBots", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) UpdateBotActive(botUserId string, active bool) (*model.Bot, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.UpdateBotActive(botUserId, active)
	api.recordTime(startTime, "UpdateBotActive", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) PermanentDeleteBot(botUserId string) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.PermanentDeleteBot(botUserId)
	api.recordTime(startTime, "PermanentDeleteBot", true)
	return _returnsA
}

func (api *apiTimerLayer) GetBotIconImage(botUserId string) ([]byte, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := api.apiImpl.GetBotIconImage(botUserId)
	api.recordTime(startTime, "GetBotIconImage", true)
	return _returnsA, _returnsB
}

func (api *apiTimerLayer) SetBotIconImage(botUserId string, data []byte) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.SetBotIconImage(botUserId, data)
	api.recordTime(startTime, "SetBotIconImage", true)
	return _returnsA
}

func (api *apiTimerLayer) DeleteBotIconImage(botUserId string) *model.AppError {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.DeleteBotIconImage(botUserId)
	api.recordTime(startTime, "DeleteBotIconImage", true)
	return _returnsA
}

func (api *apiTimerLayer) PluginHTTP(request *http.Request) *http.Response {
	startTime := timePkg.Now()
	_returnsA := api.apiImpl.PluginHTTP(request)
	api.recordTime(startTime, "PluginHTTP", true)
	return _returnsA
}
