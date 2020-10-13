package golang

const (
	// CreateNewClub
	MemberUUIDsNotIncludeLeaderUUID = -1501
	MemberUUIDsIncludeNoExistUUID = -1502
	ClubLeaderAlreadyExist = -1503
	ClubNameDuplicate = -1504
	ClubLocationDuplicate = -1505
	ClubMemberDuplicate = -1506

	// GetRecruitmentUUIDWithClubUUID
	ThereIsNoCurrentRecruitment = -1601
	
	// GetClubUUIDWithLeaderUUID
	ThereIsNoClubWithThatLeaderUUID = -1611
)
