package structs

// Task is a struct containing Task data
type Task struct {
	ID   int64  `json:"id"`
	UID  int64  `json:"uid"`
	Name string `json:"name"`
}

// TaskCollection is collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}
