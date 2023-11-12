package books

type BookEvent interface {
	GetBookEventKind() int
}
