// Code generated by cmd/api-generator. DO NOT EDIT.

package api

import (
	"github.com/Mrs4s/go-cqhttp/coolq"
	"github.com/Mrs4s/go-cqhttp/global"
)

func (c *Caller) call(action string, p Getter) global.MSG {
	switch action {
	default:
		return coolq.Failed(404, "API_NOT_FOUND", "API不存在")
	case ".get_word_slices":
		p0 := p.Get("content").String()
		return c.bot.CQGetWordSlices(p0)
	case ".handle_quick_operation":
		p0 := p.Get("context")
		p1 := p.Get("operation")
		return c.bot.CQHandleQuickOperation(p0, p1)
	case "_get_model_show":
		p0 := p.Get("model").String()
		return c.bot.CQGetModelShow(p0)
	case "_get_vip_info":
		p0 := p.Get("user_id").Int()
		return c.bot.CQGetVipInfo(p0)
	case "_send_group_notice":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("content").String()
		p2 := p.Get("image").String()
		return c.bot.CQSetGroupMemo(p0, p1, p2)
	case "_set_model_show":
		p0 := p.Get("model").String()
		p1 := p.Get("model_show").String()
		return c.bot.CQSetModelShow(p0, p1)
	case "can_send_image":
		return c.bot.CQCanSendImage()
	case "can_send_record":
		return c.bot.CQCanSendRecord()
	case "check_url_safely":
		p0 := p.Get("url").String()
		return c.bot.CQCheckURLSafely(p0)
	case "create_group_file_folder":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("parent_id").String()
		p2 := p.Get("name").String()
		return c.bot.CQGroupFileCreateFolder(p0, p1, p2)
	case "delete_essence_msg":
		p0 := int32(p.Get("message_id").Int())
		return c.bot.CQDeleteEssenceMessage(p0)
	case "delete_friend":
		p0 := p.Get("[user_id,id].0").Int()
		return c.bot.CQDeleteFriend(p0)
	case "delete_group_file":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("file_id").String()
		p2 := int32(p.Get("bus_id").Int())
		return c.bot.CQGroupFileDeleteFile(p0, p1, p2)
	case "delete_group_folder":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("folder_id").String()
		return c.bot.CQGroupFileDeleteFolder(p0, p1)
	case "delete_msg":
		p0 := int32(p.Get("message_id").Int())
		return c.bot.CQDeleteMessage(p0)
	case "delete_unidirectional_friend":
		p0 := p.Get("user_id").Int()
		return c.bot.CQDeleteUnidirectionalFriend(p0)
	case "download_file":
		p0 := p.Get("url").String()
		p1 := p.Get("headers")
		p2 := int(p.Get("thread_count").Int())
		return c.bot.CQDownloadFile(p0, p1, p2)
	case "get_essence_msg_list":
		p0 := p.Get("group_id").Int()
		return c.bot.CQGetEssenceMessageList(p0)
	case "get_forward_msg":
		p0 := p.Get("[message_id,id].0").String()
		return c.bot.CQGetForwardMessage(p0)
	case "get_friend_list":
		return c.bot.CQGetFriendList()
	case "get_group_at_all_remain":
		p0 := p.Get("group_id").Int()
		return c.bot.CQGetAtAllRemain(p0)
	case "get_group_file_system_info":
		p0 := p.Get("group_id").Int()
		return c.bot.CQGetGroupFileSystemInfo(p0)
	case "get_group_file_url":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("file_id").String()
		p2 := int32(p.Get("bus_id").Int())
		return c.bot.CQGetGroupFileURL(p0, p1, p2)
	case "get_group_files_by_folder":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("folder_id").String()
		return c.bot.CQGetGroupFilesByFolderID(p0, p1)
	case "get_group_honor_info":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("type").String()
		return c.bot.CQGetGroupHonorInfo(p0, p1)
	case "get_group_info":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("no_cache").Bool()
		return c.bot.CQGetGroupInfo(p0, p1)
	case "get_group_list":
		p0 := p.Get("no_cache").Bool()
		return c.bot.CQGetGroupList(p0)
	case "get_group_member_info":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("user_id").Int()
		p2 := p.Get("no_cache").Bool()
		return c.bot.CQGetGroupMemberInfo(p0, p1, p2)
	case "get_group_member_list":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("no_cache").Bool()
		return c.bot.CQGetGroupMemberList(p0, p1)
	case "get_group_msg_history":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("message_seq").Int()
		return c.bot.CQGetGroupMessageHistory(p0, p1)
	case "get_group_root_files":
		p0 := p.Get("group_id").Int()
		return c.bot.CQGetGroupRootFiles(p0)
	case "get_group_system_msg":
		return c.bot.CQGetGroupSystemMessages()
	case "get_image":
		p0 := p.Get("file").String()
		return c.bot.CQGetImage(p0)
	case "get_login_info":
		return c.bot.CQGetLoginInfo()
	case "get_msg":
		p0 := int32(p.Get("message_id").Int())
		return c.bot.CQGetMessage(p0)
	case "get_online_clients":
		p0 := p.Get("no_cache").Bool()
		return c.bot.CQGetOnlineClients(p0)
	case "get_status":
		return c.bot.CQGetStatus()
	case "get_stranger_info":
		p0 := p.Get("user_id").Int()
		return c.bot.CQGetStrangerInfo(p0)
	case "get_unidirectional_friend_list":
		return c.bot.CQGetUnidirectionalFriendList()
	case "get_version_info":
		return c.bot.CQGetVersionInfo()
	case "mark_msg_as_read":
		p0 := int32(p.Get("message_id").Int())
		return c.bot.CQMarkMessageAsRead(p0)
	case "ocr_image", ".ocr_image":
		p0 := p.Get("image").String()
		return c.bot.CQOcrImage(p0)
	case "qidian_get_account_info":
		return c.bot.CQGetQiDianAccountInfo()
	case "reload_event_filter":
		p0 := p.Get("file").String()
		return c.bot.CQReloadEventFilter(p0)
	case "send_group_forward_msg":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("messages")
		return c.bot.CQSendGroupForwardMessage(p0, p1)
	case "send_group_msg":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("message")
		p2 := p.Get("auto_escape").Bool()
		return c.bot.CQSendGroupMessage(p0, p1, p2)
	case "send_msg":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("user_id").Int()
		p2 := p.Get("message")
		p3 := p.Get("message_type").String()
		p4 := p.Get("auto_escape").Bool()
		return c.bot.CQSendMessage(p0, p1, p2, p3, p4)
	case "send_private_msg":
		p0 := p.Get("user_id").Int()
		p1 := p.Get("group_id").Int()
		p2 := p.Get("message")
		p3 := p.Get("auto_escape").Bool()
		return c.bot.CQSendPrivateMessage(p0, p1, p2, p3)
	case "set_essence_msg":
		p0 := int32(p.Get("message_id").Int())
		return c.bot.CQSetEssenceMessage(p0)
	case "set_friend_add_request":
		p0 := p.Get("flag").String()
		p1 := true
		if pt := p.Get("approve"); pt.Exists() {
			p1 = pt.Bool()
		}
		return c.bot.CQProcessFriendRequest(p0, p1)
	case "set_group_add_request":
		p0 := p.Get("flag").String()
		p1 := p.Get("[sub_type,type].0").String()
		p2 := p.Get("reason").String()
		p3 := true
		if pt := p.Get("approve"); pt.Exists() {
			p3 = pt.Bool()
		}
		return c.bot.CQProcessGroupRequest(p0, p1, p2, p3)
	case "set_group_admin":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("user_id").Int()
		p2 := true
		if pt := p.Get("enable"); pt.Exists() {
			p2 = pt.Bool()
		}
		return c.bot.CQSetGroupAdmin(p0, p1, p2)
	case "set_group_anonymous_ban":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("[anonymous_flag,anonymous.flag].0").String()
		p2 := int32(p.Get("duration").Int())
		return c.bot.CQSetGroupAnonymousBan(p0, p1, p2)
	case "set_group_ban":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("user_id").Int()
		p2 := uint32(1800)
		if pt := p.Get("duration"); pt.Exists() {
			p2 = uint32(pt.Int())
		}
		return c.bot.CQSetGroupBan(p0, p1, p2)
	case "set_group_card":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("user_id").Int()
		p2 := p.Get("card").String()
		return c.bot.CQSetGroupCard(p0, p1, p2)
	case "set_group_kick":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("user_id").Int()
		p2 := p.Get("message").String()
		p3 := p.Get("reject_add_request").Bool()
		return c.bot.CQSetGroupKick(p0, p1, p2, p3)
	case "set_group_leave":
		p0 := p.Get("group_id").Int()
		return c.bot.CQSetGroupLeave(p0)
	case "set_group_name":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("group_name").String()
		return c.bot.CQSetGroupName(p0, p1)
	case "set_group_portrait":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("file").String()
		p2 := p.Get("cache").String()
		return c.bot.CQSetGroupPortrait(p0, p1, p2)
	case "set_group_special_title":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("user_id").Int()
		p2 := p.Get("title").String()
		return c.bot.CQSetGroupSpecialTitle(p0, p1, p2)
	case "set_group_whole_ban":
		p0 := p.Get("group_id").Int()
		p1 := true
		if pt := p.Get("enable"); pt.Exists() {
			p1 = pt.Bool()
		}
		return c.bot.CQSetGroupWholeBan(p0, p1)
	case "upload_group_file":
		p0 := p.Get("group_id").Int()
		p1 := p.Get("file").String()
		p2 := p.Get("name").String()
		p3 := p.Get("folder").String()
		return c.bot.CQUploadGroupFile(p0, p1, p2, p3)
	}
}
