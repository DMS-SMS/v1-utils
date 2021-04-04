package golang

const (
	// CreateNewStudent
	StudentIDDuplicate = -101
	ParentUUIDNoExist = -102
	StudentNumberDuplicate = -103
	StudentPhoneNumberDuplicate = -104

	// CreateNewTeacher
	TeacherIDDuplicate = -201
	TeacherPhoneNumberDuplicate = -202
	TeacherClassDuplicate = -203

	// CreateNewParent
	ParentIDDuplicate = -301
	ParentPhoneNumberDuplicate = -302

	// Login(Student, Teacher, Parent)
	StudentIDNoExist = -401
	IncorrectStudentPWForLogin = -402
	TeacherIDNoExist = -411
	IncorrectTeacherPWForLogin = -412
	NotCertifiedTeacherAccount = -413
	TeacherAccountMismatch = -414
	ParentIDNoExist = -421
	IncorrectParentPWForLogin = -422
	AdminIDNoExist = -431
	IncorrectAdminPWForLogin = -432

	// Get(Student, Teacher, Parent)UUIDsWithInform
	StudentWithThatInformNoExist = -501
	TeacherWithThatInformNoExist = -511
	ParentWithThatInformNoExist = -521

	// GetStudentInformsWithUUIDs
	StudentUUIDsContainNoExistUUID = -601

	// Change(Student, Teacher, Parent)PW
	IncorrectStudentPWForChange = -701
	IncorrectTeacherPWForChange = -801
	IncorrectParentPWForChange = -901
)
