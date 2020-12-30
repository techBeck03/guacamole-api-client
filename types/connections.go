package types

// GuacConnection base guacamole connection info
type GuacConnection struct {
	Name              string                   `json:"name"`
	Identifier        string                   `json:"identifier,omitempty"`
	ParentIdentifier  string                   `json:"parentIdentifier"`
	Protocol          string                   `json:"protocol"`
	Attributes        GuacConnectionAttributes `json:"attributes"`
	Properties        GuacConnectionParameters `json:"parameters"`
	ActiveConnections int                      `json:"activeConnections,omitempty"`
}

// GuacConnectionAttributes guacd attributes
type GuacConnectionAttributes struct {
	GuacdEncryption       string `json:"guacd-encryption"`
	FailoverOnly          string `json:"failover-only"`
	Weight                string `json:"weight"`
	MaxConnections        string `json:"max-connections"`
	GuacdHostname         string `json:"guacd-hostname,omitempty"`
	GuacdPort             string `json:"guacd-port"`
	MaxConnectionsPerUser string `json:"max-connections-per-user"`
}

// GuacConnectionParameters defines guacamole connection parameters
type GuacConnectionParameters struct {
	/*** Network ***/
	// All
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
	// SSH
	PublicHostKey string `json:"host-key"`
	// Kubernetes
	UseSSL string `json:"use-ssl"`
	CACert string `json:"ca-cert"`

	/*** Authentication ***/
	// SSH
	Username   string `json:"username"`
	Password   string `json:"password"`
	PrivateKey string `json:"private-key"`
	Passphrase string `json:"passphrase"`
	// RDP
	Domain                string `json:"domain"`
	Security              string `json:"security"`
	DisableAuthentication string `json:"disable-auth"`
	IgnoreCert            string `json:"ignore-cert"`
	// Telnet
	UsernameRegex     string `json:"username-regex"`
	PasswordRegex     string `json:"password-regex"`
	LoginSuccessRegex string `json:"login-success-regex"`
	LoginFailureRegex string `json:"login-failure-regex"`
	// Kubernetes
	ClientCert string `json:"client-cert"`
	ClientKey  string `json:"client-key"`

	/*** Display ***/
	// SSH
	ColorScheme string `json:"color-scheme"`
	FontName    string `json:"font-name"`
	FontSize    string `json:"font-size"`
	Scrollback  string `json:"scrollback"`
	ReadOnly    string `json:"read-only"`
	// RDP
	Width        string `json:"width"`
	Height       string `json:"height"`
	DPI          string `json:"dpi"`
	ColorDepth   string `json:"color-depth"`
	ResizeMethod string `json:"resize-method"`
	// VNC
	SwapRedBlue string `json:"swap-red-blue"`
	Cursor      string `json:"cursor"`

	/*** Clipboard ***/
	DisableCopy  string `json:"disable-copy"`
	DisablePaste string `json:"disable-paste"`
	// VNC
	ClipboardEncoding string `json:"clipboard-encoding"`

	/*** Session Environment, Basic Settings ***/
	ExecuteCommand          string `json:"command"`
	Locale                  string `json:"locale"`
	Timezone                string `json:"timezone"`
	ServerKeepaliveInterval string `json:"server-alive-interval"`
	InitialProgram          string `json:"initial-program"`
	ClientName              string `json:client-name`
	KeyboardLayout          string `json:server-layout`
	AdministratorConsole    string `json:"console"`

	/*** Terminal Behavior ***/
	Backspace    string `json:"backspace"`
	TerminalType string `json:"terminal-type"`

	/*** Typescript (Text Session Recording) ***/
	TypescriptPath       string `json:"typescript-path"`
	TypescriptName       string `json:"typescript-name"`
	CreateTypescriptPath string `json:"create-typescript-path"`

	/*** Screen Recording ***/
	RecordingPath          string `json:"recording-path"`
	RecordingName          string `json:"recording-name"`
	RecordingExcludeOutput string `json:"recording-exclude-output"`
	RecordingExcludeMouse  string `json:"recording-exclude-mouse"`
	RecordingIncludeKeys   string `json:"recording-include-keys"`
	CreateRecordingPath    string `json:"create-recording-path"`

	/*** SFTP ***/
	EnableSFTP              string `json:"enable-sftp"`
	SFTPRootDirectory       string `json:"sftp-root-directory"`
	SFTPDisableFileDownload string `json:"sftp-disable-download"`
	SFTPDisableFileUpload   string `json:"sftp-disable-upload"`
	SFTPHostname            string `json:"sftp-hostname"`
	SFTPPort                string `json:"sftp-port"`
	SFTPHostKey             string `json:"sftp-host-key"`
	SFTPUsername            string `json:"sftp-username"`
	SFTPPassword            string `json:"sftp-password"`
	SFTPPrivateKey          string `json:"sftp-private-key"`
	SFTPPassphrase          string `json:"sftp-passphrase"`
	SFTPUploadDirectory     string `json:"sftp-directory"`
	SFTPKeepAliveInterval   string `json:"sftp-server-alive-interval"`

	/*** Wake-on-LAN ***/
	WOLSendPacket       string `json:"wol-send-packet"`
	WOLMacAddress       string `json:"wol-mac-addr"`
	WOLBroadcastAddress string `json:"wol-broadcast-addr"`
	WOLBootWaitTime     string `json:"wol-wait-time"`

	/*** RDP - Remote Desktop Gateway ***/
	GatewayHostname string `json:"gateway-hostname"`
	GatewayPort     string `json:"gateway-port"`
	GatewayUsername string `json:"gateway-username"`
	GatewayPassword string `json:"gateway-password"`
	GatewayDomain   string `json:"gateway-domain"`

	/*** RDP - Device Redirection ***/
	ConsoleAudio        string `json:"console-audio"`
	DisableAudio        string `json:"disable-audio"`
	EnableAudioInput    string `json:"enable-audio-input"`
	EnablePrinting      string `json:"enable-printing"`
	PrinterName         string `json:"printer-name"`
	EnableDrive         string `json:"enable-drive"`
	DriveName           string `json:"drive-name"`
	DisableFileDownload string `json:"disable-download"`
	DisableFileUpload   string `json:"disable-upload"`
	DrivePath           string `json:"drive-path"`
	CreateDrivePath     string `json:"create-drive-path"`
	StaticChannels      string `json:"static-channels"`

	/*** RDP - Performance ***/
	EnableWallpaper          string `json:"enable-wallpaper"`
	EnableTheming            string `json:"enable-theming"`
	EnableFontSmoothing      string `json:"enable-font-smoothing"`
	EnableFullWindowDrag     string `json:"enable-full-window-drag"`
	EnableDesktopComposition string `json:"enable-desktop-composition"`
	EnableMenuAnimations     string `json:"enable-menu-animations"`
	DisableBitmapCaching     string `json:"disable-bitmap-caching"`
	DisableOffscreenCaching  string `json:"disable-offscreen-caching"`
	DisableGlyphCaching      string `json:"disable-glyph-caching"`

	/*** RDP - RemoteApp ***/
	RemoteApp                 string `json:"remote-app"`
	RemoteAppWorkingDirectory string `json:"remote-app-dir"`
	RemoteAppParameters       string `json:"remote-app-args"`

	/*** RDP - Preconnection PDU/Hyper-V ***/
	PreconnectionID   string `json:"preconnection-id"`
	PreconnectionBLOB string `json:"preconnection-blob"`

	/*** RDP - Load Balancing ***/
	LoadBalanceInfo string `json:"laod-balance-info"`

	/*** VNC Repeater ***/
	DestinationHost string `json:"dest-host"`
	DestinationPort string `json:"dest-port"`

	/*** Audio ***/
	// VNC
	AudioServerName string `json:"audio-servername"`
	EnableAudio     string `json:"enable-audio"`

	/*** Container ***/
	Container string `json:"container"`
	Namespace string `json:"namespace"`
	Pod       string `json:"pod"`
}
