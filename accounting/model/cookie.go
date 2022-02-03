package model

//const SelectedProjectIdCookie string = "SelectedProjectId"
//
//func GetSelectedProjectId(request *http.Request) int64 {
//	selectedProjectId, _ := strconv.ParseInt(GetSelectedProjectIdString(request), 10, 64)
//	return selectedProjectId
//}
//
//func GetSelectedProjectIdString(request *http.Request) string {
//	cookie, err := request.Cookie(SelectedProjectIdCookie)
//	if err != nil {
//		return "0"
//	}
//	return cookie.Value
//}
//
//func SetSelectedProjectId(writer http.ResponseWriter, projectId string) {
//	expiration := time.Now().Add(365 * 24 * time.Hour)
//	http.SetCookie(writer, &http.Cookie{
//		Name:    SelectedProjectIdCookie,
//		Value:   projectId,
//		Path:    "/",
//		Expires: expiration,
//	})
//}
