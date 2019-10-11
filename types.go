package togotime

type MeEndpoint struct {
	Since int64 `json:"since"`
	Data  Data  `json:"data"`
}

type Data struct {
	ID                     int64       `json:"id"`
	APIToken               string      `json:"api_token"`
	DefaultWid             int64       `json:"default_wid"`
	Email                  string      `json:"email"`
	Fullname               string      `json:"fullname"`
	JqueryTimeofdayFormat  string      `json:"jquery_timeofday_format"`
	JqueryDateFormat       string      `json:"jquery_date_format"`
	TimeofdayFormat        string      `json:"timeofday_format"`
	DateFormat             string      `json:"date_format"`
	StoreStartAndStopTime  bool        `json:"store_start_and_stop_time"`
	BeginningOfWeek        int64       `json:"beginning_of_week"`
	Language               string      `json:"language"`
	ImageURL               string      `json:"image_url"`
	SidebarPiechart        bool        `json:"sidebar_piechart"`
	At                     string      `json:"at"`
	CreatedAt              string      `json:"created_at"`
	Retention              int64       `json:"retention"`
	RecordTimeline         bool        `json:"record_timeline"`
	RenderTimeline         bool        `json:"render_timeline"`
	TimelineEnabled        bool        `json:"timeline_enabled"`
	TimelineExperiment     bool        `json:"timeline_experiment"`
	ShouldUpgrade          bool        `json:"should_upgrade"`
	AchievementsEnabled    bool        `json:"achievements_enabled"`
	Timezone               string      `json:"timezone"`
	OpenidEnabled          bool        `json:"openid_enabled"`
	SendProductEmails      bool        `json:"send_product_emails"`
	SendWeeklyReport       bool        `json:"send_weekly_report"`
	SendTimerNotifications bool        `json:"send_timer_notifications"`
	LastBlogEntry          string      `json:"last_blog_entry"`
	Workspaces             []Workspace `json:"workspaces"`
	DurationFormat         string      `json:"duration_format"`
}

type Workspace struct {
	ID                          int64   `json:"id"`
	Name                        string  `json:"name"`
	Profile                     int64   `json:"profile"`
	Premium                     bool    `json:"premium"`
	Admin                       bool    `json:"admin"`
	DefaultHourlyRate           int64   `json:"default_hourly_rate"`
	DefaultCurrency             string  `json:"default_currency"`
	OnlyAdminsMayCreateProjects bool    `json:"only_admins_may_create_projects"`
	OnlyAdminsSeeBillableRates  bool    `json:"only_admins_see_billable_rates"`
	OnlyAdminsSeeTeamDashboard  bool    `json:"only_admins_see_team_dashboard"`
	ProjectsBillableByDefault   bool    `json:"projects_billable_by_default"`
	Rounding                    int64   `json:"rounding"`
	RoundingMinutes             int64   `json:"rounding_minutes"`
	APIToken                    *string `json:"api_token,omitempty"`
	At                          string  `json:"at"`
	IcalEnabled                 bool    `json:"ical_enabled"`
	Projects                    []Project
	Clients                     []Client
}

type Project struct {
	ID            int64  `json:"id"`
	Wid           int64  `json:"wid"`
	Cid           int64  `json:"cid"`
	Name          string `json:"name"`
	Billable      bool   `json:"billable"`
	IsPrivate     bool   `json:"is_private"`
	Active        bool   `json:"active"`
	Template      bool   `json:"template"`
	At            string `json:"at"`
	CreatedAt     string `json:"created_at"`
	Color         string `json:"color"`
	AutoEstimates bool   `json:"auto_estimates"`
	HexColor      string `json:"hex_color"`
}

type Client struct {
	ID    int64  `json:"id"`
	Wid   int64  `json:"wid"`
	Name  string `json:"name"`
	At    string `json:"at"`
	Notes string `json:"notes"`
	Hrate int64  `json:"hrate"`
	Cur   string `json:"cur"`
}

type TimeEntry struct {
	ID          int64    `json:"id"`
	Wid         int64    `json:"wid"`
	PID         *int64   `json:"pid,omitempty"`
	Billable    bool     `json:"billable"`
	Start       string   `json:"start"`
	Stop        string   `json:"stop"`
	Duration    int64    `json:"duration"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	At          string   `json:"at"`
}
