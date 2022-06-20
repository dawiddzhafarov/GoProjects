package blogposts_test

import (
	"errors"
	blogposts "github.com/dawiddzhafarov/GoProjects/GoBasics/reading_files"
	"io/fs"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Oh no, I always fail")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("always failing execution", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})

		if err == nil {
			t.Errorf("Error should occur, but didn't")
		}
	})
	t.Run("num of posts test", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello hello world.md": {Data: []byte("hi")},
			"hello-world2.md":      {Data: []byte("hola")},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, but wanted %d posts", len(posts), len(fs))
		}
	})
}