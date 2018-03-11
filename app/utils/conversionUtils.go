package utils

import (
	"fmt"

	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"github.com/vahdet/go-auth-service/models"
	pb "github.com/vahdet/go-user-store-redis/proto"
)

func ConvertProtoToModel(proto *pb.User) (*models.User, error) {

	createdTimestamp, err := ptypes.Timestamp(proto.Created)
	lastChangedTimestamp, err := ptypes.Timestamp(proto.LastChanged)
	if err != nil {
		log.WithFields(log.Fields{
			"createdTime": proto.Created,
			"lastChanged": proto.LastChanged,
		}).Warn(fmt.Sprintf("Time conversion failed: '%#v'", err))
		return nil, err
	}

	return &models.User{
		Id:          proto.Id,
		Name:        proto.Name,
		Email:       proto.Email,
		Location:    proto.Location,
		Language:    proto.Language,
		Created:     createdTimestamp,
		LastChanged: lastChangedTimestamp,
	}, nil
}
