package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/avtara/carehub/internal/models"
	"github.com/avtara/carehub/internal/service"
	"github.com/avtara/carehub/internal/service/repository/postgres/queries"
	"github.com/jmoiron/sqlx"
)

type categoryRepository struct {
	conn *sqlx.DB
}

func NewCategoryRepository(
	conn *sqlx.DB,
) service.CategoryRepository {
	return &categoryRepository{
		conn: conn,
	}
}

func (c *categoryRepository) GetAllCategories(ctx context.Context, limit int) (result []models.Category, err error) {
	rows, err := c.conn.QueryContext(ctx, queries.GetAllCategory, limit)
	if err != nil {
		if err != sql.ErrNoRows {

			return
		}
	}
	defer rows.Close()

	for rows.Next() {
		var tmp models.Category
		err = rows.Scan(&tmp.ID, &tmp.Name)
		if err != nil {
			err = fmt.Errorf("[Repository][GetAllCategories] failed scan GetAllCategory: %s", err.Error())
			return
		}
		result = append(result, tmp)
	}

	return
}

func (c *categoryRepository) GetCategoryByID(ctx context.Context, ID int64) (result models.Category, err error) {
	rows, err := c.conn.QueryContext(ctx, queries.GetCategoryByID, ID)
	if err != nil {
		err = fmt.Errorf("[Repository][GetCategoryByID] failed getting category: %s", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tmp models.Category
		err = rows.Scan(&tmp.ID, &tmp.Name)
		if err != nil {
			err = fmt.Errorf("[Repository][GetCategoryByID] failed scan Category: %s", err.Error())
			return
		}
		result = tmp
	}

	if result.ID == 0 {
		err = models.ErrorCategoryNotFound
		return
	}

	return
}

func (c *categoryRepository) GetExtraFieldByCategoryID(ctx context.Context, categoryID int64) (result []models.ExtraFieldCategory, err error) {
	rows, err := c.conn.QueryContext(ctx, queries.GetExtraFieldByCategoryID, categoryID)
	if err != nil {
		if err != sql.ErrNoRows {
			err = fmt.Errorf("[Repository][GetExtraFieldByCategoryID] failed getting extra field: %s", err.Error())
			return
		}
	}
	defer rows.Close()

	for rows.Next() {
		var tmp models.ExtraFieldCategory
		err = rows.Scan(&tmp.ID, &tmp.CategoryID, &tmp.FieldType, &tmp.FieldLabel, &tmp.FieldOptions)
		if err != nil {
			err = fmt.Errorf("[Repository][GetExtraFieldByCategoryID] failed scan ExtraFieldCategory: %s", err.Error())
			return
		}
		result = append(result, tmp)
	}

	return
}
