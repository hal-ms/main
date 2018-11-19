package repo

import "testing"

func TestAny(t *testing.T) {
	Job.Create(JobCategory.All()[3].Name, "test!")
}
