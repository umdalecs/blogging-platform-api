package main

import (
	"context"
	"database/sql"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ctx = context.Background()

type PostRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(postDto PostDto, post *Post) error {
	var id int
	err := r.db.QueryRow(ctx, `
    INSERT INTO posts (title, content, category, tags)
    VALUES ($1, $2, $3, $4) RETURNING id`,
		postDto.Title, postDto.Content, postDto.Category, postDto.Tags).Scan(&id)
	if err != nil {
		return err
	}

	err = pgxscan.Get(ctx, r.db, post, `
	SELECT 1 id, title, content, category, tags, created_at, updated_at 
	FROM posts WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) UpdatePost(id int, postDto PostDto, post *Post) error {
	commandTag, err := r.db.Exec(ctx, `
    UPDATE posts SET title = $1, content = $2, category = $3, tags = $4
    WHERE id = $5`,
		postDto.Title, postDto.Content, postDto.Category, postDto.Tags, id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return nil
	}

	err = pgxscan.Get(ctx, r.db, post, `
	SELECT 1 id, title, content, category, tags, created_at, updated_at 
	FROM posts WHERE id = $1`, id)

	if err != nil {
		return err
	}

	return nil

}

func (r *PostRepository) DeletePost(id int) (bool, error) {
	commandTag, err := r.db.Exec(ctx, `DELETE FROM posts WHERE id = $1`, id)
	if err != nil {
		return false, err
	}

	if commandTag.RowsAffected() == 0 {
		return false, nil
	}

	return true, nil
}

func (r *PostRepository) GetPosts(term string, posts *[]Post) error {
	err := pgxscan.Select(ctx, r.db, posts, `
	SELECT id, title, content, category, tags, created_at, updated_at 
    FROM posts
    WHERE title ILIKE '%' || $1 || '%'
       OR content ILIKE '%' || $1 || '%'
       OR category ILIKE '%' || $1 || '%'
       OR EXISTS (
					SELECT 1 FROM unnest(tags) tag
					WHERE tag ILIKE '%' || $1 || '%'
       )
			ORDER BY id ASC
	`, term)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) GetPostById(id int, post *Post) error {
	err := pgxscan.Get(ctx, r.db, post, `
	SELECT id, title, content, category, tags, created_at, updated_at 
	FROM posts 
	WHERE id = $1`, id)

	if err == sql.ErrNoRows {
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}
