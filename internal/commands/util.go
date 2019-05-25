package commands

import (
	"github.com/bwmarrin/discordgo"
)

type (
	// CmdFunc is the form of functions used for commands
	CmdFunc func(Context)

	// Command contains a function and description for commands
	Command struct {
		fnc  CmdFunc
		desc string
	}

	// CmdMap is a map of commands
	CmdMap map[string]Command

	// CommandHandler contains a map of commands
	CommandHandler struct {
		cmds CmdMap
	}
)

// Context provides context for the Command
type Context struct {
	Session      *discordgo.Session
	Guild        *discordgo.Guild
	VoiceChannel *discordgo.Channel
	TextChannel  *discordgo.Channel
	User         *discordgo.User
	Message      *discordgo.MessageCreate
	Args         []string
	cmdHandler   *CommandHandler
	Prefix       string
}

// NewContext creates a new Context that can be used by commands
func NewContext(session *discordgo.Session, guild *discordgo.Guild, channel *discordgo.Channel, usr *discordgo.User, msg *discordgo.MessageCreate, cmdHandler *CommandHandler, prefix string) *Context {
	ctx := new(Context)
	ctx.Session = session
	ctx.Guild = guild
	ctx.TextChannel = channel
	ctx.User = usr
	ctx.Message = msg
	ctx.cmdHandler = cmdHandler
	ctx.Prefix = prefix
	return ctx
}

// NewCommandHandler creates a new CommandHandler
func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make(CmdMap)}
}

// GetCmds returns a map of commands
func (handler CommandHandler) GetCmds() CmdMap {
	return handler.cmds
}

// Get returns the command with the same name as the argument, and a bool telling if the command was found
func (handler CommandHandler) Get(name string) (*CmdFunc, bool) {
	cmd, found := handler.cmds[name]
	return &cmd.fnc, found
}

// Register registers new commands
func (handler CommandHandler) Register(name string, desc string, cmd CmdFunc) {
	newCmd := new(Command)
	newCmd.fnc = cmd
	newCmd.desc = desc

	handler.cmds[name] = *newCmd
}

// MemberHasPermission checks if a member has the given permission
// for example, If you would like to check if user has the administrator
// permission you would use
// --- MemberHasPermission(s, guildID, userID, discordgo.PermissionAdministrator)
// If you want to check for multiple permissions you would use the bitwise OR
// operator to pack more bits in. (e.g): PermissionAdministrator|PermissionAddReactions
// =================================================================================
//     s          :  discordgo session
//     guildID    :  guildID of the member you wish to check the roles of
//     userID     :  userID of the member you wish to retrieve
//     permission :  the permission you wish to check for
func MemberHasPermission(s *discordgo.Session, guildID string, userID string, permission int) (bool, error) {
	member, err := s.State.Member(guildID, userID)
	if err != nil {
		if member, err = s.GuildMember(guildID, userID); err != nil {
			return false, err
		}
	}

	// Iterate through the role IDs stored in member.Roles
	// to check permissions
	for _, roleID := range member.Roles {
		role, err := s.State.Role(guildID, roleID)
		if err != nil {
			return false, err
		}
		if role.Permissions&permission != 0 {
			return true, nil
		}
	}

	return false, nil
}
