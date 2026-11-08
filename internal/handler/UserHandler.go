package handler


import(
	"errors"
	"github.com/napollytop/manga/internal/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

)


type AuthHandler struct{
	DB *gorm.DB
}


func NewAuthHandler(db *gorm.DB) *AuthHandler{
	return &AuthHandler{DB: db}
}

type RegisterRequest struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}


type LoginRequest struct{
	Email string `json:"email"`
	Password string `json:"password"`
}


// register handler
func (h *AuthHandler) Register(c *fiber.Ctx) error{
	var req RegisterRequest
	if err := c.BodyParser(req); err != nil{
		return  c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "format json tidak valid",
		})
	}

	// validasi input
	if req.Username == "" || req.Email == ""|| req.Password == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "kredensial tidak valid",
		})
	}
// hash password
hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "gagal melakukan hashing terhadap password",
		})
	}


	user := models.User{
		Username: req.Username,
		Email: req.Email,
		Password: string(hashedPassword),

	}

	if err := h.DB.Create(&user).Error; err != nil{
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "username atau email sudah terdaftar",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "berhasil daftar",
		"user": fiber.Map{
			"id": user.ID,
			"username": user.Username,
			"email": user.Email,
		},
	})

}

// login handler

func (h *AuthHandler) Login(c *fiber.Ctx) error{
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "format json tidak valid",
		})

	}

	var user models.User
	// cari user berdasarkan email
	err := h.DB.Where("email = ? ", req.Email).First(&user).Error
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":"email atau password sala",
			})

		}
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "username atau password salah",
			})
		}
		
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Login Berhasil",
			"user": fiber.Map{
				"id": user.ID,
				"username": user.Username,
				"email": user.Email,
			},
		})
	
}




