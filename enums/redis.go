package enums

import "fmt"

const (
	// GuildRedisKey - the key for retrieving information about the guild
	GuildRedisKey string = "cache:guilds:%s"
	// GuildMemberRedisKey - for retrieving information on a member of a guild
	GuildMemberRedisKey string = "cache:guilds:%s:members:%s"
	// GuildRoleRedisKey - for getting info on a role in a guild
	GuildRoleRedisKey string = "cache:guilds:%s:roles:%s"
	// GuildTextChannelRedisKey - for getting info on a text channel in a guild
	GuildTextChannelRedisKey string = "cache:guilds:%s:text_channels:%s"
	// GuildVoiceChannelRedisKey - for getting info on a voice channel in a guild
	GuildVoiceChannelRedisKey string = "cache:guilds:%s:voice_channels:%s"
)

// GetGuildRedisKey - formats the redis key for a guild
func GetGuildRedisKey(gID string) string {
	return fmt.Sprintf(GuildRedisKey, gID)
}

// GetMemberRedisKey - formats the redis key for a guild member
func GetMemberRedisKey(gID string, mID string) string {
	return fmt.Sprintf(GuildMemberRedisKey, gID, mID)
}

// GetRoleRedisKey - formats the redis key for a guild role
func GetRoleRedisKey(gID string, rID string) string {
	return fmt.Sprintf(GuildRoleRedisKey, gID, rID)
}

// GetTextChannelRedisKey - formats the redis key for a guild text channel
func GetTextChannelRedisKey(gID string, cID string) string {
	return fmt.Sprintf(GuildTextChannelRedisKey, gID, cID)
}

// GetVoiceChannelRedisKey - formats the redis key for a guild voice channel
func GetVoiceChannelRedisKey(gID string, cID string) string {
	return fmt.Sprintf(GuildVoiceChannelRedisKey, gID, cID)
}
