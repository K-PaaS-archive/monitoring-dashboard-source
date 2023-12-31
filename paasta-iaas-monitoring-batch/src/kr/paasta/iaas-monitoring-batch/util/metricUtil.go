package util

import (
	cb "iaas-monitoring-batch/model/base"
)

const (
	SERVICE_NAME string = "serviceName"
	DATA_NAME    string = "data"
)


type errorMessage struct {
	cb.ErrMessage
}

func GetError() *errorMessage {
	return &errorMessage{}
}

func (e errorMessage) DbCheckError(err error) cb.ErrMessage {
	if err != nil {
		errMessage := cb.ErrMessage{
			"Message": err.Error(),
		}
		return errMessage
	} else {
		return nil
	}
}


