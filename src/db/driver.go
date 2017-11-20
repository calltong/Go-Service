package db

func InitSession(path string) error {
  // init method to start db
	if service.baseSession == nil {
    //service.URL = "mongodb://localhost:27017"
		service.URL = path
    return service.New()
  } else {
		return nil
	}
}
