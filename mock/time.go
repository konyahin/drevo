package mock

type FakeTime struct{}

func (_ FakeTime) CurrentDate() string {
	return "2025-06-20"
}
