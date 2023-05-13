package repositories

import (
	"api/src/models"
	"database/sql"
)

// Publications representa um repositório de publications
type Publications struct {
	db *sql.DB
}

// PublicationRepository cria repositório de publications
func PublicationRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

func (repository Publications) Create(publication models.Publication) (uint64, error) {

	statement, erro := repository.db.Prepare(
		"INSERT into PUBLICATIONS (title, content, author_id) VALUES (?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if erro != nil {
		return 0, erro
	}

	lastInsertId, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertId), nil
}

func (repository Publications) ListPublicationsThatUserFollows(userId uint64) ([]models.Publication, error) {
	rows, erro := repository.db.Query(
		`
		SELECT p.id, p.title, p.content, p.author_id, u.nick, p.createdAt 
		FROM publications p
			LEFT JOIN users u
			ON p.author_id = u.ID
			INNER JOIN followers f
			ON u.ID = f.user_id
		WHERE f.follower_id = ?
		ORDER BY p.createdAt desc;
		`,
		userId,
	)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var publications []models.Publication
	for rows.Next() {

		var tempPublication models.Publication
		if erro = rows.Scan(
			&tempPublication.ID,
			&tempPublication.Title,
			&tempPublication.Content,
			&tempPublication.AuthorID,
			&tempPublication.AuthorNick,
			&tempPublication.CreatedAt,
		); erro != nil {
			return nil, erro
		}
		publications = append(publications, tempPublication)

	}
	return publications, nil
}

func (repository Publications) ListMyPublications(userId uint64) ([]models.Publication, error) {

	rows, erro := repository.db.Query(
		`
		SELECT p.id, p.title, p.content, p.author_id, u.nick, p.createdAt 
		FROM publications p
			INNER JOIN users u
			ON p.author_id = u.ID
		WHERE u.ID = ?
		ORDER BY p.createdAt desc;
		`,
		userId,
	)

	if erro != nil {
		return nil, erro
	}

	defer rows.Close()

	var publications []models.Publication
	for rows.Next() {

		var tempPublication models.Publication
		if erro = rows.Scan(
			&tempPublication.ID,
			&tempPublication.Title,
			&tempPublication.Content,
			&tempPublication.AuthorID,
			&tempPublication.AuthorNick,
			&tempPublication.CreatedAt,
		); erro != nil {
			return nil, erro
		}
		publications = append(publications, tempPublication)

	}
	return publications, nil
}

func (repository Publications) Retrieve(publicationId uint64) (models.Publication, error) {

	var tempPublication models.Publication
	row, erro := repository.db.Query(
		`
		SELECT p.id, p.title, p.content, p.author_id, u.nick, p.createdAt 
		FROM publications p
		LEFT JOIN users u
		ON p.author_id = u.ID
		WHERE p.ID = ?
		`,
		publicationId,
	)
	if erro != nil {
		return tempPublication, erro
	}

	defer row.Close()

	for row.Next() {
		if erro = row.Scan(
			&tempPublication.ID,
			&tempPublication.Title,
			&tempPublication.Content,
			&tempPublication.AuthorID,
			&tempPublication.AuthorNick,
			&tempPublication.CreatedAt,
		); erro != nil {
			return tempPublication, erro
		}
	}

	return tempPublication, nil
}

func (repository Publications) GetAuthorIdFromPublication(publicationId uint64) (uint64, error) {

	var authorID uint64
	row, erro := repository.db.Query(
		`
		SELECT author_id
		FROM publications
		WHERE ID = ?
		`,
		publicationId,
	)
	if erro != nil {
		return 0, erro
	}

	defer row.Close()

	for row.Next() {
		if erro = row.Scan(
			&authorID,
		); erro != nil {
			return 0, erro
		}
	}

	return authorID, nil
}

func (repository Publications) Update(publicationId uint64, data models.Publication) error {

	statement, erro := repository.db.Prepare(
		"UPDATE publications SET title = ?, content = ? WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(data.Title, data.Content, publicationId)

	if erro != nil {
		return erro
	}
	return nil
}

func (repository Publications) Delete(publicationId uint64) (int64, error) {

	statement, erro := repository.db.Prepare(
		"DELETE FROM publications WHERE id = ?;",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(publicationId)
	if erro != nil {
		return 0, erro
	}

	rows_deleted, erro := result.RowsAffected()
	if erro != nil {
		return 0, erro
	}

	return rows_deleted, nil
}

func (repository Publications) PublicationHaveSomeLike(publicationId uint64) (bool, error) {

	var haveSomeLike bool = false
	row, erro := repository.db.Query(
		`
		SELECT likes
		FROM publications
		WHERE ID = ?
		`,
		publicationId,
	)
	if erro != nil {
		return haveSomeLike, erro
	}
	defer row.Close()

	var likes uint64
	for row.Next() {
		if erro = row.Scan(&likes); erro != nil {
			return haveSomeLike, erro
		}
	}

	if likes > 0 {
		haveSomeLike = true
	}
	return haveSomeLike, nil
}

func (repository Publications) LikePublication(publicationId uint64) error {
	statement, erro := repository.db.Prepare(
		"UPDATE publications SET likes = likes + 1 WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicationId)
	if erro != nil {
		return erro
	}
	return nil
}

func (repository Publications) RemoveLikePublication(publicationId uint64) error {
	statement, erro := repository.db.Prepare(
		"UPDATE publications SET likes = likes - 1 WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicationId)
	if erro != nil {
		return erro
	}
	return nil
}
