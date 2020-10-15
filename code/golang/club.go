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

	// Common code in club leader service
	ForbiddenNotStudentOrAdminUUID = -1711
	ForbiddenNotClubLeader = -1712
	NotFoundClubNoExist = -1721

	// AddClubMember
	ClubMemberAlreadyExist = -1701
	NotFoundStudentNoExist = -1722

	// DeleteClubMember
	NotFoundClubMemberNoExist = -1723

	// ChangeClubLeader
	AlreadyClubLeader = -1801
	ClubLeaderDuplicateForChange = -1802

	// DeleteClubWithUUID
	RecruitmentInProgressExist = -1901
)
