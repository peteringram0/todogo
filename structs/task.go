package structs

// Task is a struct containing Task data
type Task struct {
	ID   int    `json:"id"`
	UID  int    `json:"uid"`
	Name string `json:"name"`
}

// TaskCollection is collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}
