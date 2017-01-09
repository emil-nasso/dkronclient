package main

func main() {
	client := NewClient()
	tags := map[string]string{
		"role": "bar:10",
	}
	job := DkronJob{Name: "mass-merry-christmas", Schedule: "@midnight", Command: "echo \"merry christmas\"", Tags: tags}
	client.CreateJob(job)
}
