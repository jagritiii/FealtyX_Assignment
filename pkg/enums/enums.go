package enums

type Message string

var (
	Faileddecode      Message = "failed to decode the object"
	Validationerror   Message = " validation failed"
	Successful        Message = "Succesfully updated "
	Servererror       Message = "Internal Server Error"
	Paramserror       Message = "Params send in wrong way"
	InsertSucceeded   Message = "Successfully inserted the student record"
	InsertFailed      Message = "Failed to insert the student record"
	DeleteSucceeded   Message = "Successfully deleted the student record"
	DeleteFailed      Message = "Failed to delete the student record"
	UpdateSucceeded   Message = "Successfully updated the student record"
	UpdateFailed      Message = "Failed to update the student record"
	FailedAllStudents Message = "Failed to retrieve all students"
	NoStudentById     Message = "No student record was found for this id"
	InvalidStudentId  Message = "Invalid student ID"
	BadRequest        Message = "Bad Request"
	NotFound          Message = "Record not found"
)
