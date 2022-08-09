package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gmo_2022_summer/pkg/model"
	"gmo_2022_summer/pkg/view"
	"golang.org/x/crypto/bcrypt"

	//"log"
	"math"
	"time"
)

func hashPW(pw string) string {
	hpw, _ := bcrypt.GenerateFromPassword([]byte(pw), 4)
	return string(hpw)
}

func checkPW(userid string, pw string) bool {
	hash := model.GetUser(userid).Password
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw)); err != nil {
		return false
	} else {
		return true
	}
}

func Register(c *gin.Context) {
	type request struct {
		UserID    string `json:"UserID" binding:"required"`
		Name      string `json:"Name" binding:"required"`
		Birthdate int    `json:"Birthdate" binding:"required"`
		Sex       int    `json:"Sex" binding:"required"`
		Height    int    `json:"Height" binding:"required"`
		Weight    int    `json:"Weight" binding:"required"`
		Password  string `json:"Password" binding:"required"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		view.BadRequest(
			c,
			"Body not complete",
		)
		return
	}

	var u model.User
	copier.Copy(&u, &req)

	if !CheckUserIDDuplication(u.UserID) {
		view.BadRequest(
			c,
			"UserID is already taken",
		)
		return

	}

	u.Password = hashPW(u.Password)

	// 年齢計算
	bDtimeTime := time.Unix(int64(u.Birthdate), 0)
	age := RoundTime((time.Now()).Sub(bDtimeTime).Seconds() / 31207680)

	// objective計算
	if true {
		if u.Sex == 1 {
			if age <= 7 {
				u.Objective = 1550
			} else if age <= 9 {
				u.Objective = 1850
			} else if age <= 11 {
				u.Objective = 2250
			} else if age <= 14 {
				u.Objective = 2600
			} else if age <= 17 {
				u.Objective = 2800
			} else if age <= 29 {
				u.Objective = 2650
			} else if age <= 49 {
				u.Objective = 2700
			} else if age <= 64 {
				u.Objective = 2600
			} else if age <= 74 {
				u.Objective = 2400
			} else {
				u.Objective = 2100
			}
		} else {
			if age <= 7 {
				u.Objective = 1450
			} else if age <= 9 {
				u.Objective = 1700
			} else if age <= 11 {
				u.Objective = 2100
			} else if age <= 14 {
				u.Objective = 2400
			} else if age <= 17 {
				u.Objective = 2300
			} else if age <= 29 {
				u.Objective = 2000
			} else if age <= 49 {
				u.Objective = 2050
			} else if age <= 64 {
				u.Objective = 1950
			} else if age <= 74 {
				u.Objective = 1850
			} else {
				u.Objective = 1650
			}
		}
	}

	if err := model.CreateUser(u); err != nil {
		view.BadRequest(c, "SQL error has occured")
		return
	}
	view.StatusOK(c, "OK", u)
}

func Login(c *gin.Context) {
	type request struct {
		UserID   string `json:"UserID" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"Detail": 1})
		return
	}
	if !checkPW(req.UserID, req.Password) {
		c.JSON(403, gin.H{
			"Message": "Password is wrong",
		})
		return
	}
	c.JSON(200, gin.H{"Detail": true})
}

func CheckUserIDDuplication(userid string) bool {
	u := model.GetUser(userid)
	if u.UserID == "" {
		return true
	} else {
		return false
	}
}

func RoundTime(input float64) int {
	var result float64

	if input < 0 {
		result = math.Ceil(input - 0.5)
	} else {
		result = math.Floor(input + 0.5)
	}

	// only interested in integer, ignore fractional
	i, _ := math.Modf(result)

	return int(i)
}

func UpdateUser(c *gin.Context) {
	type request struct {
		UserID    string `json:"UserID" binding:"required"`
		Name      string `json:"Name" binding:"required"`
		Birthdate int    `json:"Birthdate" binding:"required"`
		Sex       int    `json:"Sex" binding:"required"`
		Height    int    `json:"Height" binding:"required"`
		Weight    int    `json:"Weight" binding:"required"`
		Objective int    `json:"Objective" binding:"required"`
		Password  string `json:"Password" json:"-"`
		NPassword string `json:"NPassword" json:"-"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, nil)
		return
	}

	if (req.Password == "") || (req.NPassword == "") {
		req.NPassword = ""
	} else {
		if !checkPW(req.UserID, req.Password) {
			c.JSON(403, nil)
			return
		}
	}

	req.Password = req.NPassword

	var newu model.User

	if err := copier.Copy(&newu, &req); err != nil {
		c.JSON(400, nil)
		return
	}

	if err := model.UpdateUser(newu); err != nil {
		c.JSON(400, gin.H{"Message": "CRUD error"})
	}
	c.JSON(200, map[string]any{"Detail": newu})
}

func GetUser(c *gin.Context) {
	type request struct {
		UserID string `json:"UserID" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, nil)
		return
	}
	c.JSON(200, gin.H{"Details": model.GetUser(req.UserID)})
}
