; Inno Setup Script

#define MyAppName "joinme.click-launcher"
#define MyAppVersion "0.0.2.0"
#define MyAppPublisher "cetteup"
#define MyAppPublisherURL "https://github.com/cetteup"
#define MyAppURL "https://github.com/cetteup/joinme.click-launcher"
#define MyAppExeName "joinme.click-launcher.exe"

[Setup]
AppId={{BAFF781C-B425-4D62-8ACD-299A40A2B752}
AppName={#MyAppName}
VersionInfoVersion={#MyAppVersion}
AppVersion={#MyAppVersion}
AppVerName={#MyAppName} v{#MyAppVersion}
AppPublisher={#MyAppPublisher}
AppPublisherURL={#MyAppPublisherURL}
AppSupportURL={#MyAppURL}
AppUpdatesURL={#MyAppURL}
DefaultDirName={autopf}\joinme.click-launcher
DisableProgramGroupPage=yes
OutputDir=.
OutputBaseFilename=joinme.click-launcher-setup
SetupIconFile=resource\icon.ico
UninstallDisplayIcon={app}\icon.ico
Compression=lzma2/max
SolidCompression=yes
WizardStyle=classic
PrivilegesRequired=lowest

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Files]
Source: "joinme.click-launcher.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "resource\icon.ico"; DestDir: "{app}"; Flags: ignoreversion

[Run]
Filename: "{app}\joinme.click-launcher.exe"; Parameters: "-quiet"

[UninstallRun]
Filename: "{app}\joinme.click-launcher.exe"; Parameters: "-deregister -quiet"; RunOnceId: "DeregisterHandlers"