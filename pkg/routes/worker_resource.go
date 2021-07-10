package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kameshsampath/hybrid-cloud-backend/pkg/utils"
	"log"
	"net/http"
	"strings"
)

var (
	workerId string
)

func init() {
	workerId = "worker-" + uuid.New().String()[:4]
}

// Process godoc
// @Summary process the request message
// @Description process the request message from the front end and apply the needed transformations
// @Tags backend
// @Accept json
// @Produce json
// @Param message body routes.Message true "Message to process"
// @Success 200 {object} routes.Response
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
//@Router /process [post]
func (e *Endpoints) Process(c *gin.Context) {
	var message Message
	if err := c.BindJSON(&message); err != nil {
		utils.NewError(c, http.StatusInternalServerError, err)
		return
	} else {
		log.Printf("Processing message %v", message)
		if message.Request.Text == "" {
			utils.NewError(c, http.StatusBadRequest, err)
			return
		}
		response := message.process()
		log.Printf("Processed message %v", response)
		c.JSON(http.StatusOK, response)
	}
}

func (m *Message) process() Response {
	request := m.Request
	text := request.Text
	if request.Reverse {
		text = utils.Reverse(text)
	}
	if request.Uppercase {
		text = strings.ToUpper(text)
	}

	return Response{
		RequestId: m.RequestId,
		WorkerId:  workerId,
		Text:      text,
		CloudId:   utils.WhichCloud(),
	}
}
