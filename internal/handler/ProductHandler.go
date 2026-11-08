package handler


import(
	"github.com/napollytop/manga/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"errors"
)

type ProductHandler struct{
	DB *gorm.DB
}

func NewProductHandler(db *gorm.DB) *ProductHandler{
	return &ProductHandler{DB: db}
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error{
	var products []models.Product
	if err := h.DB.Find(&products).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "gagal mengambil produk",
		})
	}
	return c.JSON(products)
}

func (h *ProductHandler) PostProduct(c *fiber.Ctx) error{
	var products models.Product

	if err := c.BodyParser(&products); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "format JSON tidak valied",
		})
	}
	if products.Name == "" || products.Price <= 0{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "nama dan harga produk wajib di sertakan",
		})
	}

	if err := h.DB.Create(&products).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "gagal menyimpan produk",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"ok": "produk berhasil di tambahkan",
	})
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error{
    id, err := c.ParamsInt("id")
	
	if err != nil || id <= 0{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id Produk tidak valid",
		})
	}

	var product models.Product
	if err := h.DB.First(&product, id).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return c.Status(fiber.StatusNotFound).JSON((fiber.Map{
				"error": "Produk Tidak Ditemukan",
			}))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":"gagal mencari produk",
		})
	}

	if err := h.DB.Delete(&product).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "gagal menghapus produk",
		})
	}

	return c.JSON(fiber.Map{
		"ok": "produk berhasil di hapus",
	})

}