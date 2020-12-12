package utils

func DepartmentExist(department string) int {
	for k, v := range [...]string{"Marketing", "Design", "Development"} {
		if department == v {
			return k
		}
	}
	return -1
}