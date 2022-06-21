package blogposts_test

import (
	"errors"
	blogposts "github.com/dawiddzhafarov/GoProjects/GoBasics/reading_files"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

const (
	firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
	secondBody = `Title: Post 2
Description: Description 2
Tags: tdd2, go2
---
L
O
L`
)

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Oh no, I always fail")
}

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}
	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Errorf("error occured but shouldnt")
	}
	got := posts[0]
	want := blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	}
	assertPost(t, got, want)
}

func TestWrongFormat(t *testing.T) {
	t.Run("wrong format file with correct ones", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":   {Data: []byte(firstBody)},
			"hello world2.txt": {Data: []byte(firstBody)},
			"hello world3.md":  {Data: []byte(secondBody)},
		}
		posts, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Errorf("Shouldnt get an error but did")
		}
		got := posts[0]
		want := blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		}
		if len(posts) != 2 {
			t.Errorf("should get 2 posts, but got %d", len(posts))
		}
		assertPost(t, got, want)
	})

	t.Run("only incorrect files format", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world2.txt": {Data: []byte(firstBody)},
		}
		posts, err := blogposts.NewPostsFromFS(fs)
		if err == nil {
			t.Errorf("Shouldnt get an error but did")
		}

		if len(posts) != 0 {
			t.Errorf("Shouldnt create any posts, but created %d", (len(posts)))
		}
	})

}

func TestReadFileError(t *testing.T) {
	t.Run("always failing execution", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})

		if err == nil {
			t.Errorf("Error should occur, but didn't")
		}
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, but wanted %+v", got, want)
	}
}
