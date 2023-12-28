package types

type WordCountRequest struct {
	Content string
}

type WordCountReply struct {
	Counts map[string]int
}
