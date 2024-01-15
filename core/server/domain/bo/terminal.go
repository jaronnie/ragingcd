package bo

type ExecTerminalBo struct {
	ResourceType string `form:"resource_type"`
	ResourceID   string `form:"resource_id"`
	WorkingDir   string `form:"working_dir"`
	Rows         uint16 `form:"rows"`
	Cols         uint16 `form:"cols"`
}
