package util

import S "ConfBackend/services"

func PadUrlLinkToPcdFile(fileName string) string {
	pref := S.S.Conf.Pcd.ServerPCDFileUrlPrefix
	// if pref ends with /
	if pref[len(pref)-1] == '/' {
		return pref + fileName
	} else {
		return pref + "/" + fileName
	}
}
