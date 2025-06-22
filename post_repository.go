package main

import (
	"database/sql"
	"encoding/json"
)

type PostRepository struct {
	Db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{Db: db}
}

type scanner interface {
	Scan(dest ...any) error
}

func scanRowIntoPost(scanner scanner, post *Post) error {
	var jsonTags []byte

	err := scanner.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &jsonTags, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonTags, &post.Tags)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) CreatePost(postDto PostDto) (*Post, error) {
	jsonTags, err := json.Marshal(postDto.Tags)
	if err != nil {
		return nil, err
	}

	var (
		title    = postDto.Title
		content  = postDto.Content
		category = postDto.Category
	)

	res, err := r.Db.Exec(`
    INSERT INTO posts (title, content, category, tags)
    VALUES (?, ?, ?, ?)`,
		title, content, category, jsonTags)
	if err != nil {
		return nil, err
	}

	postID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	row := r.Db.QueryRow(`
	SELECT id, title, content, category, tags, created_at, updated_at 
	FROM posts WHERE id = ?`, postID)

	post := &Post{}
	err = scanRowIntoPost(row, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostRepository) UpdatePost(id int, postDto PostDto) (*Post, error) {
	var post = &Post{}
	jsonTags, err := json.Marshal(postDto.Tags)
	if err != nil {
		return nil, err
	}

	post.Title = postDto.Title
	post.Content = postDto.Content
	post.Category = postDto.Category

	res, err := r.Db.Exec(`
    UPDATE posts SET title = ?, content = ?, category = ?, tags = ?
    WHERE id = ?`,
		post.Title, post.Content, post.Category, jsonTags, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return post, nil
	}

	row := r.Db.QueryRow(`
	SELECT id, title, content, category, tags, created_at, updated_at 
	FROM posts WHERE id = ?`, id)

	err = scanRowIntoPost(row, post)
	if err != nil {
		return nil, err
	}

	return post, nil

}

func (r *PostRepository) DeletePost(id int) (bool, error) {
	res, err := r.Db.Exec(`DELETE FROM posts WHERE id = ?`, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (r *PostRepository) GetPosts(term string) ([]Post, error) {
	posts := []Post{}

	rows, err := r.Db.Query(`
	SELECT id, title, content, category, tags, created_at, updated_at 
	FROM posts
	WHERE title like CONCAT('%', ?, '%')
	OR content like CONCAT('%', ?, '%')
	OR category like CONCAT('%', ?, '%')
	OR json_search(LOWER(tags), 'one', LOWER(CONCAT('%', ?, '%'))) IS NOT NULL
	`, term, term, term, term)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		post := &Post{}
		err := scanRowIntoPost(rows, post)
		if err != nil {
			return nil, err
		}

		posts = append(posts, *post)
	}

	return posts, nil
}

func (r *PostRepository) GetPostById(id int) (*Post, error) {
	row := r.Db.QueryRow(`
	SELECT id, title, content, category, tags, created_at, updated_at 
	FROM posts 
	WHERE id = ?`, id)

	post := &Post{}
	err := scanRowIntoPost(row, post)

	if err == sql.ErrNoRows {
		return post, nil
	}

	if err != nil {
		return nil, err
	}

	return post, nil
}
