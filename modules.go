package msf

//
// func (msf *Metasploit) GetModule(moduleType string, SessionToken string) []string {
// 	modules := msf.PackAndSend([]string{"module." + moduleType, SessionToken})
// 	return []string{modules["modules"]}
// }
//
// func (msf *Metasploit) GetExploits() []string {
// 	return msf.GetExploitsById(0)
// }
//
// func (msf *Metasploit) GetExploitsById(index int) []string {
// 	session := msf.SessionTokens[index]
// 	//if session == nil {
// 	//	return nil, nil
// 	//}
// 	return msf.GetExploitsBySession(session)
// }
//
// func (msf *Metasploit) GetExploitsBySession(SessionToken string) []string {
// 	return msf.GetModule("exploit", SessionToken)
// }
//
// func (msf *Metasploit) GetAuxiliary() []string {
// 	return msf.GetAuxiliaryById(0)
// }
//
// func (msf *Metasploit) GetAuxiliaryById(index int) []string {
// 	session := msf.SessionTokens[index]
// 	//if session == nil {
// 	//	return nil, nil
// 	//}
// 	return msf.GetAuxiliaryBySession(session)
// }
//
// func (msf *Metasploit) GetAuxiliaryBySession(SessionToken string) []string {
// 	return msf.GetModule("auxiliary", SessionToken)
// }

// POST
// Encoders
// Nops

// [ "module.info", "<token>", "ModuleType", "ModuleName" ]
// [ "module.options", "<token>", "ModuleType", "ModuleName" ]
// [ "module.compatible_payloads", "<token>", "ModuleName" ]
