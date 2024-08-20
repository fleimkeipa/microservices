package interfaces

type MessageInterfaces interface {
	Send(string, string) error
}
