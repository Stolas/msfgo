# msfgo, Metasploit RPC Go client

As Metasploit RPC relies heavenly on msfpack. You'll need ``github.com/ugorji/go/codec`` thus first run:

```
go get github.com/ugorji/go/codec
```

## Metasploit RPC Modules
- [ ] Authentication
    - [X] Login
    - [X] Logout
    - [ ] Token Add
    - [ ] Token List
    - [ ] Token Remote
- [ ] Core
    - [ ] Add Module Path
    - [ ] Module Stats
    - [ ] Reload Modules
    - [ ] Save
    - [ ] Setg
    - [ ] Unsetg
    - [ ] Thread List
    - [ ] Thread Kill
    - [X] Version
    - [ ] Stop
- [ ] Console
    - [ ] Create
    - [ ] Destroy
    - [ ] List
    - [ ] Write
    - [ ] Read
    - [ ] Session Detach
    - [ ] Session Kill
    - [ ] Tabs
- [ ] Jobs
    - [ ] List
    - [ ] Info
    - [ ] Stop
- [ ] Modules
    - [ ] Exploits
    - [ ] Auxiliary
    - [ ] Post
    - [ ] Payloads
    - [ ] Encoders
    - [ ] Nops
    - [ ] Info
    - [ ] Options
    - [ ] Compatible Payloads
    - [ ] Target Compatible Payloads
    - [ ] Compatible Sessions
    - [ ] Encode
    - [ ] Execute
- [ ] Plugins
    - [ ] Load
    - [ ] Unload
    - [ ] Loaded
- [ ] Sessions
    - [ ] List
    - [ ] Stop
    - [ ] Shell Read
    - [ ] Shell Write
    - [ ] Meterpreter Write
    - [ ] Meterpreter Read
    - [ ] Meterpreter Run Single
    - [ ] Meterpreter Script
    - [ ] Meterpreter Session Detach
    - [ ] Meterpreter Session Kill
    - [ ] Meterpreter Tabs
    - [ ] Compatible Modules
    - [ ] Shell Upgrade
    - [ ] Ring Clear
    - [ ] Ring Last
    - [ ] Ring Put
- [ ] Pro
    - [ ] About
    - [ ] Workspaces
    - [ ] Projects
    - [ ] Workspace Add
    - [ ] Project Add
    - [ ] Workspace Delete
    - [ ] Project Delete
    - [ ] Users
- [ ] Pro License
    - [ ] Register
    - [ ] Activate
    - [ ] Activate Offline
    - [ ] License
    - [ ] Revert License
- [ ] Pro Updates
    - [ ] Update Available
    - [ ] Update Install
    - [ ] Update Install Offline
    - [ ] Update Update Status
    - [ ] Update Update Stop
    - [ ] Restart Service
- [ ] Pro Task
    - [ ] Task List
    - [ ] Task Status
    - [ ] Task Stop
    - [ ] Task Log
    - [ ] Task Delete Log
- [ ] Pro Feature
    - [ ] Start Discover
    - [ ] Start Import
    - [ ] Start Import Creds
    - [ ] Start Nexpose
    - [ ] Start Bruteforce
    - [ ] Start Exploit
    - [ ] Start Cleanup
    - [ ] Start Collect
    - [ ] Start Report
    - [ ] Report List
    - [ ] List Report
    - [ ] List Report Types
    - [ ] Report Download
    - [ ] Report Artifact Download
    - [ ] Start Export
    - [ ] Export List
    - [ ] Export Download
- [ ] Pro Import
    - [ ] Import Data
    - [ ] Import File
    - [ ] Validate Import File
- [ ] Pro Loot
    - [ ] Load Download
    - [ ] Load List
- [ ] Pro Module
    - [ ] Module Search
    - [ ] Module Validate
    - [ ] Modules
