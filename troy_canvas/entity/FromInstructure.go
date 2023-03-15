package entity

import "time"

/*
We use the example payload from https://canvas.instructure.com/doc/api/index.html
Student_Info_Instructure = Users
Course_Info_Instructure = Course
Course_Announcement_Instructure = Announcement for each course
Quiz_Instructure = Quizzes
*/
type Student_Info_Instructure struct {
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	CreatedAt       string      `json:"created_at"`
	SortableName    string      `json:"sortable_name"`
	ShortName       string      `json:"short_name"`
	AvatarURL       string      `json:"avatar_url"`
	Locale          interface{} `json:"locale"`
	EffectiveLocale string      `json:"effective_locale"`
	Permissions     struct {
		CanUpdateName           bool `json:"can_update_name"`
		CanUpdateAvatar         bool `json:"can_update_avatar"`
		LimitParentAppWebAccess bool `json:"limit_parent_app_web_access"`
	} `json:"permissions"`
}
type Course_Info_Instructure struct {
	ID                          int         `json:"id"`
	Name                        string      `json:"name"`
	AccountID                   int         `json:"account_id"`
	UUID                        string      `json:"uuid"`
	StartAt                     time.Time   `json:"start_at"`
	GradingStandardID           interface{} `json:"grading_standard_id"`
	IsPublic                    bool        `json:"is_public"`
	CreatedAt                   time.Time   `json:"created_at"`
	CourseCode                  string      `json:"course_code"`
	DefaultView                 string      `json:"default_view"`
	RootAccountID               int         `json:"root_account_id"`
	EnrollmentTermID            int         `json:"enrollment_term_id"`
	License                     string      `json:"license"`
	GradePassbackSetting        interface{} `json:"grade_passback_setting"`
	EndAt                       time.Time   `json:"end_at"`
	PublicSyllabus              bool        `json:"public_syllabus"`
	PublicSyllabusToAuth        bool        `json:"public_syllabus_to_auth"`
	StorageQuotaMb              int         `json:"storage_quota_mb"`
	IsPublicToAuthUsers         bool        `json:"is_public_to_auth_users"`
	HomeroomCourse              bool        `json:"homeroom_course"`
	CourseColor                 interface{} `json:"course_color"`
	FriendlyName                interface{} `json:"friendly_name"`
	ApplyAssignmentGroupWeights bool        `json:"apply_assignment_group_weights"`
	Calendar                    struct {
		Ics string `json:"ics"`
	} `json:"calendar"`
	TimeZone    string `json:"time_zone"`
	Blueprint   bool   `json:"blueprint"`
	Template    bool   `json:"template"`
	Enrollments []struct {
		Type                           string `json:"type"`
		Role                           string `json:"role"`
		RoleID                         int    `json:"role_id"`
		UserID                         int    `json:"user_id"`
		EnrollmentState                string `json:"enrollment_state"`
		LimitPrivilegesToCourseSection bool   `json:"limit_privileges_to_course_section"`
	} `json:"enrollments"`
	HideFinalGrades                  bool   `json:"hide_final_grades"`
	WorkflowState                    string `json:"workflow_state"`
	RestrictEnrollmentsToCourseDates bool   `json:"restrict_enrollments_to_course_dates"`
}
type Course_Announement_Instructure struct {
	ID                      int         `json:"id"`
	Title                   string      `json:"title"`
	LastReplyAt             time.Time   `json:"last_reply_at"`
	CreatedAt               time.Time   `json:"created_at"`
	DelayedPostAt           interface{} `json:"delayed_post_at"`
	PostedAt                time.Time   `json:"posted_at"`
	AssignmentID            interface{} `json:"assignment_id"`
	RootTopicID             interface{} `json:"root_topic_id"`
	Position                int         `json:"position"`
	PodcastHasStudentPosts  bool        `json:"podcast_has_student_posts"`
	DiscussionType          string      `json:"discussion_type"`
	LockAt                  interface{} `json:"lock_at"`
	AllowRating             bool        `json:"allow_rating"`
	OnlyGradersCanRate      bool        `json:"only_graders_can_rate"`
	SortByRating            bool        `json:"sort_by_rating"`
	IsSectionSpecific       bool        `json:"is_section_specific"`
	AnonymousState          interface{} `json:"anonymous_state"`
	UserName                string      `json:"user_name"`
	DiscussionSubentryCount int         `json:"discussion_subentry_count"`
	Permissions             struct {
		Attach bool `json:"attach"`
		Update bool `json:"update"`
		Reply  bool `json:"reply"`
		Delete bool `json:"delete"`
	} `json:"permissions"`
	RequireInitialPost interface{}   `json:"require_initial_post"`
	UserCanSeePosts    bool          `json:"user_can_see_posts"`
	PodcastURL         interface{}   `json:"podcast_url"`
	ReadState          string        `json:"read_state"`
	UnreadCount        int           `json:"unread_count"`
	Subscribed         bool          `json:"subscribed"`
	Attachments        []interface{} `json:"attachments"`
	Published          bool          `json:"published"`
	CanUnpublish       bool          `json:"can_unpublish"`
	Locked             bool          `json:"locked"`
	CanLock            bool          `json:"can_lock"`
	CommentsDisabled   bool          `json:"comments_disabled"`
	Author             struct {
		ID             int         `json:"id"`
		AnonymousID    string      `json:"anonymous_id"`
		DisplayName    string      `json:"display_name"`
		AvatarImageURL string      `json:"avatar_image_url"`
		HTMLURL        string      `json:"html_url"`
		Pronouns       interface{} `json:"pronouns"`
	} `json:"author"`
	HTMLURL            string        `json:"html_url"`
	URL                string        `json:"url"`
	Pinned             bool          `json:"pinned"`
	GroupCategoryID    interface{}   `json:"group_category_id"`
	CanGroup           bool          `json:"can_group"`
	TopicChildren      []interface{} `json:"topic_children"`
	GroupTopicChildren []interface{} `json:"group_topic_children"`
	ContextCode        string        `json:"context_code"`
	LockedForUser      bool          `json:"locked_for_user"`
	LockInfo           struct {
		CanView     bool   `json:"can_view"`
		AssetString string `json:"asset_string"`
	} `json:"lock_info"`
	LockExplanation  string      `json:"lock_explanation"`
	Message          string      `json:"message"`
	SubscriptionHold string      `json:"subscription_hold"`
	TodoDate         interface{} `json:"todo_date"`
	IsAnnouncement   bool        `json:"is_announcement"`
}
type Quiz_Instructure struct {
	ID                      int         `json:"id"`
	Title                   string      `json:"title"`
	HTMLURL                 string      `json:"html_url"`
	MobileURL               string      `json:"mobile_url"`
	Description             string      `json:"description"`
	QuizType                string      `json:"quiz_type"`
	TimeLimit               interface{} `json:"time_limit"`
	TimerAutosubmitDisabled bool        `json:"timer_autosubmit_disabled"`
	ShuffleAnswers          bool        `json:"shuffle_answers"`
	ShowCorrectAnswers      bool        `json:"show_correct_answers"`
	ScoringPolicy           string      `json:"scoring_policy"`
	AllowedAttempts         int         `json:"allowed_attempts"`
	OneQuestionAtATime      bool        `json:"one_question_at_a_time"`
	QuestionCount           int         `json:"question_count"`
	PointsPossible          float64     `json:"points_possible"`
	CantGoBack              bool        `json:"cant_go_back"`
	IPFilter                interface{} `json:"ip_filter"`
	DueAt                   interface{} `json:"due_at"`
	LockAt                  time.Time   `json:"lock_at"`
	UnlockAt                interface{} `json:"unlock_at"`
	Published               bool        `json:"published"`
	LockedForUser           bool        `json:"locked_for_user"`
	LockInfo                struct {
		LockAt      time.Time `json:"lock_at"`
		CanView     bool      `json:"can_view"`
		AssetString string    `json:"asset_string"`
	} `json:"lock_info,omitempty"`
	LockExplanation      string      `json:"lock_explanation,omitempty"`
	HideResults          interface{} `json:"hide_results"`
	ShowCorrectAnswersAt time.Time   `json:"show_correct_answers_at"`
	HideCorrectAnswersAt interface{} `json:"hide_correct_answers_at"`
	AllDates             []struct {
		DueAt    interface{} `json:"due_at"`
		UnlockAt interface{} `json:"unlock_at"`
		LockAt   time.Time   `json:"lock_at"`
		Base     bool        `json:"base"`
	} `json:"all_dates"`
	CanUpdate                        bool        `json:"can_update"`
	RequireLockdownBrowser           bool        `json:"require_lockdown_browser"`
	RequireLockdownBrowserForResults bool        `json:"require_lockdown_browser_for_results"`
	RequireLockdownBrowserMonitor    bool        `json:"require_lockdown_browser_monitor"`
	LockdownBrowserMonitorData       interface{} `json:"lockdown_browser_monitor_data"`
	Permissions                      struct {
		Manage           bool `json:"manage"`
		Read             bool `json:"read"`
		Create           bool `json:"create"`
		Update           bool `json:"update"`
		Submit           bool `json:"submit"`
		Preview          bool `json:"preview"`
		Delete           bool `json:"delete"`
		ReadStatistics   bool `json:"read_statistics"`
		Grade            bool `json:"grade"`
		ReviewGrades     bool `json:"review_grades"`
		ViewAnswerAudits bool `json:"view_answer_audits"`
	} `json:"permissions"`
	QuizReportsURL                string      `json:"quiz_reports_url"`
	QuizStatisticsURL             string      `json:"quiz_statistics_url"`
	ImportantDates                bool        `json:"important_dates"`
	QuizSubmissionVersionsHTMLURL string      `json:"quiz_submission_versions_html_url"`
	AssignmentID                  int         `json:"assignment_id"`
	OneTimeResults                bool        `json:"one_time_results"`
	AssignmentGroupID             int         `json:"assignment_group_id"`
	ShowCorrectAnswersLastAttempt bool        `json:"show_correct_answers_last_attempt"`
	VersionNumber                 int         `json:"version_number"`
	HasAccessCode                 bool        `json:"has_access_code"`
	PostToSis                     bool        `json:"post_to_sis"`
	MigrationID                   interface{} `json:"migration_id"`
	InPacedCourse                 bool        `json:"in_paced_course"`
}
