package authsvc

import (
	"net/http"
	"tauth/database"
	"tauth/dbops/gorm/keystrokes"
	"tauth/entities"
	"tauth/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/* ---------------------------------- LOGIN --------------------------------- */
func (h *AuthSvcImpl) LoginUser(ctx *gin.Context, req LoginObject) (models.BaseResponse, entities.Users, error) {
	var res models.BaseResponse
	var user entities.Users
	var err error

	// default response
	res.Message = "something went wrong"
	res.StatusCode = http.StatusInternalServerError

	// get user
	user, err = h.userGorm.GetUserDetailsEmail(ctx, req.Email)
	if err != nil {
		return res, user, err
	}

	db, _ := database.Connection()
	keystrokeGorm := keystrokes.Gorm(db)

	savedProfile, err := keystrokeGorm.GetKeyStrokeByUserPID(ctx, user.PID.String)
	if err != nil {
		return res, user, err
	}

	if req.Password != "" {
		// check password
		err = h.checkPassword(ctx, req)
		if err != nil {
			res.Success = false
			res.Message = "incorrect password"
			res.StatusCode = http.StatusUnauthorized
			return res, user, nil
		}

		// success response
		res.Success = true
		res.Message = "login succcesfull"
		res.StatusCode = http.StatusOK

		return res, user, nil
	} else {
		var keystroke entities.KeystrokeProfile
		// map the data
		keystroke.UserPID = req.TypingDNA.UserPID
		keystroke.SampleText = req.TypingDNA.SampleText
		keystroke.TextLength = len(req.TypingDNA.SampleText)
		keystroke.DwellTimes = req.TypingDNA.Metrics.RawMetrics.DwellTimes
		keystroke.FlightTimes = req.TypingDNA.Metrics.RawMetrics.FlightTimes
		keystroke.AverageDwellTime = average(req.TypingDNA.Metrics.RawMetrics.DwellTimes)
		keystroke.AverageFlightTime = average(req.TypingDNA.Metrics.RawMetrics.FlightTimes)
		keystroke.WordsPerMinute = float64(req.TypingDNA.Metrics.WPM)
		keystroke.DeviceInfo = req.TypingDNA.DeviceInfo.Browser
		keystroke.CreatedFrom = ctx.ClientIP()

		match, _, err := compareKeystrokeProfiles(savedProfile, keystroke)
		if err != nil {
			return res, user, err
		}
		if match {
			// success response
			res.Success = true
			res.Message = "login succcesfull"
			res.StatusCode = http.StatusOK

			return res, user, nil
		} else {
			// success response
			res.Success = true
			res.Message = "login failed, pattern does not match"
			res.StatusCode = http.StatusOK

			return res, user, nil
		}

	}

	return res, user, err
}

func (h *AuthSvcImpl) checkPassword(c *gin.Context, req LoginObject) error {
	res, err := h.userGorm.GetUserDetailsEmail(c, req.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(req.Password))
	if err != nil {
		return err
	}

	return nil
}
