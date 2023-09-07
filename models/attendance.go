package models

import (
	"time"
)

type Attendance struct {
	UniqueID              string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	LoginMode             string     `json:"loginMode,omitempty" bson:"loginMode,omitempty"`
	Status                string     `json:"status,omitempty" bson:"status,omitempty"`
	PunchIn               *time.Time `json:"punchIn" bson:"punchIn,omitempty"`
	PunchOut              *time.Time `json:"punchOut" bson:"punchOut,omitempty"`
	EmployeeId            string     `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OrganisationId        string     `json:"organisationId" bson:"organisationId,omitempty"`
	WorkingHours          float64    `json:"workingHours,omitempty" bson:"workingHours,omitempty"`
	OverTime              float64    `json:"overtime,omitempty" bson:"overtime,omitempty"`
	TotalOverTimehoursstr string     `json:"totalOverTimehoursstr,omitempty" bson:"totalOverTimehoursstr,omitempty"`
	TotalWorkingMins      float64    `json:"totalWorkingMins,omitempty" bson:"totalWorkingMins,omitempty"`
	TotalWorkingHoursStr  string     `json:"totalworkingHoursStr,omitempty" bson:"totalworkingHoursStr,omitempty"`
	TotalBreakHoursStr    string     `json:"totalbreakHoursStr,omitempty" bson:"totalbreakHoursStr,omitempty"`
	TotalBreakMins        float64    `json:"totalbreakMins,omitempty" bson:"totalbreakMins,omitempty"`
	Date                  *time.Time `json:"date,omitempty" bson:"date,omitempty"`
	LeaveType             string     `json:"leaveType,omitempty" bson:"leaveType,omitempty"`
	LineManager           string     `json:"lineManager,omitempty" bson:"lineManager,omitempty"`
	ImageURL              string     `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
	WorkScheduleId        string     `json:"workscheduleId,omitempty" bson:"workscheduleId,omitempty"`
	PayRoll               string     `json:"payRoll" bson:"payRoll,omitempty"`
	CaseLOP               string     `json:"caseLOP" bson:"caseLOP,omitempty"`
	Notes                 string     `json:"notes,omitempty" bson:"notes,omitempty"`
	Created               *Created   `json:"createdOn" bson:"createdOn,omitempty"`
}

type ClockoutAttendance struct {
	EmployeeId     string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OrganisationId string `json:"organisationId" bson:"organisationId,omitempty"`
	Notes          string `json:"notes,omitempty" bson:"notes,omitempty"`
}

// type AttendanceEmployeeLeaveRequest struct {
// 	UniqueID       string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
// 	State          string     `json:"state,omitempty" bson:"state,omitempty"`
// 	Status         string     `json:"status,omitempty" bson:"status,omitempty"`
// 	EmployeeId     string     `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
// 	OrganisationId string     `json:"organisationId" bson:"organisationId,omitempty"`
// 	Notes          string     `json:"notes,omitempty" bson:"notes,omitempty"`
// 	Date           *time.Time `json:"date,omitempty" bson:"date,omitempty"`
// 	LeaveType      string     `json:"leaveType,omitempty" bson:"leaveType,omitempty"`
// 	LineManager    string     `json:"lineManager,omitempty" bson:"lineManager,omitempty"`
// 	Created        *Created   `json:"createdOn" bson:"createdOn,omitempty"`
// }

type EmployeeAttendanceTodayStatus struct {
	//Attendance         `bson:",inline"`
	RecentPunchinTime  *time.Time `json:"recentpunchinTime,omitempty" bson:"recentpunchinTime,omitempty"`
	RecentPunchoutTime *time.Time `json:"recentpunchoutTime,omitempty" bson:"recentpunchoutTime,omitempty"`
	FirstPunchinTime   *time.Time `json:"firstpunchinTime,omitempty" bson:"firstpunchinTime,omitempty"`
	LastpunchoutTime   *time.Time `json:"lastpunchoutTime,omitempty" bson:"lastpunchoutTime,omitempty"`
	TotalWorkingHours  float64    `json:"totalworkingHours,omitempty" bson:"totalworkingHours,omitempty"`
	TotalBreakHours    float64    `json:"totalbreakHours,omitempty" bson:"totalbreakHours,omitempty"`
	OverTime           float64    `json:"overTime,omitempty" bson:"overTime,omitempty"`
}

type RefAttendance struct {
	Attendance `bson:",inline"`
	Ref        struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type TodayTimeoff struct {
	EmployeeId *User `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
}
type FilterAttendance struct {
	Status     []string   `json:"status,omitempty" bson:"status,omitempty"`
	EmployeeId []string   `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Employee   string     `json:"employee,omitempty" bson:"employee,omitempty"`
	StartDate  *time.Time `json:"startDate,omitempty"  bson:"startDate,omitempty"`
	SortBy     string     `json:"sortBy"`
	SortOrder  int        `json:"sortOrder"`
	DateRange  struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
}
type DayWiseAttendanceReportFilter struct {
	EmployeeId string     `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Manager    string     `json:"manager" bson:"manager,omitempty"`
	StartDate  *time.Time `json:"startDate,omitempty"  bson:"startDate,omitempty"`
	Status     string     `json:"status,omitempty" bson:"status,omitempty"`
	SearchBox  struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchBox" bson:"searchBox"`
}
type EmployeeAttendanceApprove struct {
	EmployeeId string   `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	UniqueID   []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
}
type DayWiseAttendanceReport struct {
	WorkSchedule        float64 `json:"workSchedule" bson:"workSchedule,omitempty"`
	EmployeeName        string  `json:"employeeName" bson:"employeeName"`
	PaidTime            float64 `json:"paidTime" bson:"paidTime,omitempty"`
	TotalWorkingMins    float64 `json:"totalWorkingMins" bson:"totalWorkingMins,omitempty"`
	Deficit             float64 `json:"deficit" bson:"deficit,omitempty"`
	TotalOvertime       float64 `json:"totalOvertime" bson:"totalOvertime,omitempty"`
	DeficitStr          string  `json:"deficitStr" bson:"deficitStr,omitempty"`
	TotalWorkingMinsStr string  `json:"totalWorkingMinsStr" bson:"totalWorkingMinsStr,omitempty"`
	TotalOvertimeStr    string  `json:"totalOvertimeStr" bson:"totalOvertimeStr,omitempty"`
	Days                []struct {
		Date                *time.Time      `json:"date" bson:"date,omitempty"`
		DayOfWeek           string          `json:"dayOfWeek" bson:"dayOfWeek,omitempty"`
		PunchIn             *time.Time      `json:"punchIn" bson:"punchIn,omitempty"`
		PunchOut            *time.Time      `json:"punchOut" bson:"punchOut,omitempty"`
		PaidTime            float64         `json:"paidTime" bson:"paidTime,omitempty"`
		Status              string          `json:"status,omitempty" bson:"status,omitempty"`
		TotalWorkingMins    float64         `json:"totalWorkingMins,omitempty" bson:"totalWorkingMins,omitempty"`
		TotalBreakMins      float64         `json:"totalbreakMins,omitempty" bson:"totalbreakMins,omitempty"`
		OverTime            float64         `json:"overtime" bson:"overtime,omitempty"`
		Deficit             float64         `json:"deficit" bson:"deficit,omitempty"`
		DeficitStr          string          `json:"deficitStr" bson:"deficitStr,omitempty"`
		TotalWorkingMinsStr string          `json:"totalWorkingMinsStr" bson:"totalWorkingMinsStr,omitempty"`
		TotalOvertimeStr    string          `json:"totalOvertimeStr" bson:"totalOvertimeStr,omitempty"`
		AttendanceLog       []AttendanceLog `json:"attendanceLog" bson:"attendanceLog,omitempty"`
	} `json:"days" bson:"days"`
}
type EmployeeDayWiseAttendanceReport struct {
	User `bson:",inline"`
	Days []struct {
		Date       *time.Time `json:"date" bson:"date,omitempty"`
		Attendance Attendance `json:"attendance" bson:"attendance,omitempty"`
	} `json:"days,omitempty" bson:"days"`
	// Ref struct {
	// 	Employee `bson:",inline"`
	// } `json:"ref,omitempty" bson:"ref,omitempty"`
}
type AttendanceEmployeeStatistics struct {
	//	AttendanceLog          AttendanceLog `json:"attendanceLog,omitempty" bson:"attendanceLog,omitempty"`
	//	TodayWorkingHours      float64 `json:"totalworkingHours" bson:"totalworkingHours,omitempty"`
	TotalWorkingMins         float64 `json:"totalWorkingMins" bson:"totalWorkingMins,omitempty"`
	TodayTotalWorkingHours   float64 `json:"todaytotalworkingHours" bson:"todaytotalworkingHours,omitempty"`
	WeekWorkingHours         float64 `json:"weekworkingHours" bson:"weekworkingHours,omitempty"`
	WeekTotalWorkingHours    float64 `json:"weektotalworkingHours" bson:"weektotalworkingHours,omitempty"`
	MonthWorkingHours        float64 `json:"monthworkingHours" bson:"monthworkingHours,omitempty"`
	MonthTotalWorkingHours   float64 `json:"monthtotalworkingHours" bson:"monthtotalworkingHours,omitempty"`
	RemainingWorkingHours    float64 `json:"remainingWorkingHours" bson:"remainingWorkingHours,omitempty"`
	TotalOvertime            float64 `json:"totalOvertime" bson:"totalOvertime,omitempty"`
	RemainingWorkingHoursStr string  `json:"remainingWorkingHoursStr" bson:"remainingWorkingHoursStr,omitempty"`
	TotalOvertimeStr         string  `json:"totalOvertimeStr" bson:"totalOvertimeStr,omitempty"`
	MonthWorkingHoursStr     string  `json:"monthworkingHoursStr" bson:"monthworkingHoursStr,omitempty"`
	WeekWorkingHoursStr      string  `json:"weekworkingHoursStr" bson:"weekworkingHoursStr,omitempty"`
	TodayTotalWorkingMinsStr string  `json:"todaytotalWorkingMinsStr" bson:"todaytotalWorkingMinsStr,omitempty"`
}
type TodayEmployessLeave struct {
	Planned        float64 `json:"planned" bson:"planned,omitempty"`
	UnPlanned      float64 `json:"unPlanned" bson:"unPlanned,omitempty"`
	PendingTimeOff float64 `json:"pendingTimeOff" bson:"pendingTimeOff,omitempty"`
	TodayPresent   float64 `json:"todayPresent" bson:"todayPresent,omitempty"`
	TotalEmployee  float64 `json:"totalEmployee" bson:"totalEmployee,omitempty"`
}
type PayRollEmployessLeave struct {
	EmployeeId string  `json:"employeeId" bson:"employeeId,omitempty"`
	LossOfPay  float64 `json:"lossOfPay" bson:"lossOfPay,omitempty"`
	Paid       float64 `json:"paid" bson:"paid,omitempty"`
	PartialPay float64 `json:"partialPay" bson:"partialPay,omitempty"`
}
type UpdatePaidTime struct {
	EmployeeId     string     `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OrganisationId string     `json:"organisationId" bson:"organisationId,omitempty"`
	Date           *time.Time `json:"date,omitempty" bson:"date,omitempty"`
	WorkingHours   float64    `json:"workingHours,omitempty" bson:"workingHours,omitempty"`
}
type UpdateOverTimeTime struct {
	EmployeeId     string     `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OrganisationId string     `json:"organisationId" bson:"organisationId,omitempty"`
	Date           *time.Time `json:"date,omitempty" bson:"date,omitempty"`
	OverTime       float64    `json:"overtime,omitempty" bson:"overtime,omitempty"`
}
