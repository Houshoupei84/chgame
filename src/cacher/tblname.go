package cacher

import "strconv"

func accountId(name string) string {
	return "accs."+name+":uid"
}

func users(uid int) string {
	return "users."+strconv.Itoa(uid)
}

func allUsers() string {
	return "users.*"
}

func servers(id int) string {
	return "servers."+strconv.Itoa(id)
}

func serversPattern() string {
	return "servers.*"
}

func ciduserid(uid uint32) string {
	return "cids."+strconv.Itoa(int(uid))
}

func notices(index int) string {
	return "notices."+strconv.Itoa(int(index))
}

func noticesPattern() string {
	return "notices.*"
}

func useritems(userid uint32, id uint32) string {
	return "items."+strconv.Itoa(int(userid))+"."+strconv.Itoa(int(id))
}

func alluseritems(userid uint32) string {
	return "items."+strconv.Itoa(int(userid))+".*"
}

func allItems() string {
	return "items.*"
}

func recordId() string {
	return "record.id"
}

func recordHead(id int) string {
	return "record.rh."+strconv.Itoa(id)
}

func recordContent(id int) string {
	return "record.rc."+strconv.Itoa(id)
}

func userAllRecord(userId int) string {
	return "record.u."+strconv.Itoa(userId)+".*"
}

func userRecord(userId, recordId int) string {
	return "record.u."+strconv.Itoa(userId)+"."+strconv.Itoa(recordId)
}


/*
func recordHead(userId, id int) string {
	return "record.ch."+strconv.Itoa(userId)+"."+strconv.Itoa(id)
}

func recordContent(userId, id int) string {
	return "record.cc."+strconv.Itoa(userId)+"."+strconv.Itoa(id)
}

func allUserRecordHead(userId int) string {
	return "record.ch."+strconv.Itoa(userId) + ".*"
}

func allRecords(userId int) string {
	return "record.c*."+strconv.Itoa(userId)+".*"
}
*/
