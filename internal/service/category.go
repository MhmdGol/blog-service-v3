package service

type Category struct {
	Name string
}

func (a *categoryApp) CreateCategory(name string) error {
	return a.catRepo.CreateCategory(name)
}

func (a *categoryApp) AllCategories() ([]Category, error) {
	categoryRows, err := a.catRepo.AllCategories()
	if err != nil {
		return nil, err
	}

	categories := []Category{}
	for _, categoryRow := range categoryRows {
		categories = append(categories, Category{
			Name: categoryRow.Name,
		})
	}

	return categories, nil
}

func (a *categoryApp) UpdateCategory(categoryID int, name string) error {
	return a.catRepo.UpdateCategory(categoryID, name)
}

func (a *categoryApp) DeleteCategory(categoryID int) error {
	return a.catRepo.DeleteCategory(categoryID)
}
