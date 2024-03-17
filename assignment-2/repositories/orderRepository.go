package repositories

import (
	"assignment-2/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}


func (r *OrderRepository) Create(order *models.Orders) error {
	return r.DB.Create(order).Error
}

func (r *OrderRepository) GetAll() ([]models.Orders, error) {
	var orders []models.Orders
	if err := r.DB.Preload("Items").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) GetById(id uint) (*models.Orders, error) {
	var order models.Orders

	if err := r.DB.Preload("Items").First(&order, id).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) UpdateByID(id uint, updatedOrder models.Orders) error {

	if err := r.DB.Preload("Items").First(&models.Orders{}, id).Error; err != nil {
		return err
	}

    if err := r.DB.Model(&models.Orders{}).Where("id = ?", id).Updates(&updatedOrder).Error; err != nil {
		return err
	}

	for _, item := range updatedOrder.Items {
		var existingItem models.Items

		r.DB.Where("order_id = ? AND id = ?", id, item.ID).First(&existingItem)

		if existingItem.ID != 0 {
			r.DB.Model(&existingItem).Updates(&item)
			continue
		}

		item.OrderID = uint(models.Orders{}.ID)
		r.DB.Create(&item)
	}

	if err := r.DB.Preload("Items").First(&models.Orders{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) Delete(id uint) error {
	if err := r.DB.First(&models.Orders{}, id).Error; err != nil {
		return err
	}

	if err := r.DB.Select("Items").Delete(&models.Orders{}).Error; err != nil {
		return err
	}

	return nil
}

