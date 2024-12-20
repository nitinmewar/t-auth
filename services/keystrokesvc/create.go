package keystrokesvc

import (
	"net/http"
	"tauth/entities"
	"tauth/models"

	"github.com/gin-gonic/gin"
)

func (h *KeyStrokeImpl) CreateKeyStroke(ctx *gin.Context, req RequestBody) (models.BaseResponse, entities.Users, error) {
	var res models.BaseResponse
	var keystroke entities.KeystrokeProfile
	var user entities.Users
	var err error

	// default response
	res.Message = "something went wrong"
	res.StatusCode = http.StatusInternalServerError

	// map the data
	keystroke.UserPID = req.Data.UserPID
	keystroke.SampleText = req.Data.SampleText
	keystroke.TextLength = len(req.Data.SampleText)
	keystroke.DwellTimes = req.Data.Metrics.RawMetrics.DwellTimes
	keystroke.FlightTimes = req.Data.Metrics.RawMetrics.FlightTimes
	keystroke.AverageDwellTime = average(req.Data.Metrics.RawMetrics.DwellTimes)
	keystroke.AverageFlightTime = average(req.Data.Metrics.RawMetrics.FlightTimes)
	keystroke.WordsPerMinute = float64(req.Data.Metrics.WPM)
	keystroke.DeviceInfo = req.Data.DeviceInfo.Browser
	keystroke.CreatedFrom = ctx.ClientIP()

	// create a keystroke profile
	keystroke, err = h.keyStrokeGorm.CreateKeyStroke(ctx, keystroke)
	if err != nil {
		return res, user, err
	}

	// update user
	user, err = h.userGorm.UpdateKeystrokeMetrics(ctx, req.Data.UserPID)
	if err != nil {
		return res, user, err
	}

	// success response
	res.Success = true
	res.Message = "metrics saved succesfully"
	res.StatusCode = http.StatusOK

	return res, user, err
}

func average(items []float64) float64 {
	var sum float64

	if len(items) == 0 {
		return 0
	}

	for _, item := range items {
		sum += item
	}

	return sum / float64(len(items))
}
