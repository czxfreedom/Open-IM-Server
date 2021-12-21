package main

import (
	apiAuth "Open_IM/internal/api/auth"
	apiChat "Open_IM/internal/api/chat"
	"Open_IM/internal/api/friend"
	"Open_IM/internal/api/group"
	"Open_IM/internal/api/manage"
	apiThird "Open_IM/internal/api/third"
	"Open_IM/internal/api/user"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/utils"
	"flag"
	"github.com/gin-gonic/gin"
	"strconv"
	//"syscall"
)

func main() {

	//logFile, err := os.OpenFile("./fatal.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	//	if err != nil {

	//	return
	//	}
	//syscall.Dup2(int(logFile.Fd()), int(os.Stderr.Fd()))

	//log.Info("", "", "api server running...")
	r := gin.Default()
	r.Use(utils.CorsHandler())
	// user routing group, which handles user registration and login services
	userRouterGroup := r.Group("/user")
	{
		//更新用户信息
		userRouterGroup.POST("/update_user_info", user.UpdateUserInfo)
		//管理员调用获取用户详细信息接口可以获取多个用户注册的详细信息。
		userRouterGroup.POST("/get_user_info", user.GetUserInfo)
	}
	//friend routing group
	friendRouterGroup := r.Group("/friend")
	{
		friendRouterGroup.POST("/get_friends_info", friend.GetFriendsInfo)
		//
		friendRouterGroup.POST("/add_friend", friend.AddFriend)
		//
		friendRouterGroup.POST("/get_friend_apply_list", friend.GetFriendApplyList)
		//
		friendRouterGroup.POST("/get_self_apply_list", friend.GetSelfApplyList)
		//
		friendRouterGroup.POST("/get_friend_list", friend.GetFriendList)
		//APP管理员包uid添加到ownerUid的黑名单中
		friendRouterGroup.POST("/add_blacklist", friend.AddBlacklist)
		//
		friendRouterGroup.POST("/get_blacklist", friend.GetBlacklist)
		//
		friendRouterGroup.POST("/remove_blacklist", friend.RemoveBlacklist)
		//
		friendRouterGroup.POST("/delete_friend", friend.DeleteFriend)
		//
		friendRouterGroup.POST("/add_friend_response", friend.AddFriendResponse)
		//
		friendRouterGroup.POST("/set_friend_comment", friend.SetFriendComment)
		//
		friendRouterGroup.POST("/is_friend", friend.IsFriend)
		//APP管理员使用户A和其他人成为好友
		friendRouterGroup.POST("/import_friend", friend.ImportFriend)
	}
	//group related routing group
	groupRouterGroup := r.Group("/group")
	{
		//创建群组
		groupRouterGroup.POST("/create_group", group.CreateGroup)
		//设置群信息
		groupRouterGroup.POST("/set_group_info", group.SetGroupInfo)
		//加入群组
		groupRouterGroup.POST("join_group", group.JoinGroup)
		//退出群组
		groupRouterGroup.POST("/quit_group", group.QuitGroup)

		groupRouterGroup.POST("/group_application_response", group.ApplicationGroupResponse)

		groupRouterGroup.POST("/transfer_group", group.TransferGroupOwner)

		groupRouterGroup.POST("/get_group_applicationList", group.GetGroupApplicationList)

		groupRouterGroup.POST("/get_groups_info", group.GetGroupsInfo)
		//APP管理员把用户从群里直接踢出
		groupRouterGroup.POST("/kick_group", group.KickGroupMember)

		groupRouterGroup.POST("/get_group_member_list", group.GetGroupMemberList)

		groupRouterGroup.POST("/get_group_all_member_list", group.GetGroupAllMember)

		groupRouterGroup.POST("/get_group_members_info", group.GetGroupMembersInfo)
		//APP管理员邀请用户直接进群
		groupRouterGroup.POST("/invite_user_to_group", group.InviteUserToGroup)
		//
		groupRouterGroup.POST("/get_joined_group_list", group.GetJoinedGroupList)
	}
	//certificate
	authRouterGroup := r.Group("/auth")
	{ //注册新用户
		authRouterGroup.POST("/user_register", apiAuth.UserRegister)
		//换取 IMToken
		authRouterGroup.POST("/user_token", apiAuth.UserToken)
	}
	//Third service
	thirdGroup := r.Group("/third")
	{
		//腾讯云认证服务
		thirdGroup.POST("/tencent_cloud_storage_credential", apiThird.TencentCloudStorageCredential)
	}
	//Message
	chatGroup := r.Group("/chat")
	{
		chatGroup.POST("/newest_seq", apiChat.UserGetSeq)
		chatGroup.POST("/pull_msg", apiChat.UserPullMsg)
		chatGroup.POST("/send_msg", apiChat.UserSendMsg)
		chatGroup.POST("/pull_msg_by_seq", apiChat.UserPullMsgBySeqList)
	}
	//Manager
	managementGroup := r.Group("/manager")
	{
		//删除用户
		managementGroup.POST("/delete_user", manage.DeleteUser)
		//管理员通过后台接口发送单聊群聊消息，可以以管理员身份发消息，也可以以其他用户的身份发消息，通过sendID区分。
		managementGroup.POST("/send_msg", manage.ManagementSendMsg)
		//管理员调用获取IM已经注册的所有用户的UID接口。
		managementGroup.POST("/get_all_users_uid", manage.GetAllUsersUid)
	}
	log.NewPrivateLog("api")
	ginPort := flag.Int("port", 10000, "get ginServerPort from cmd,default 10000 as port")
	flag.Parse()
	r.Run(utils.ServerIP + ":" + strconv.Itoa(*ginPort))
}
