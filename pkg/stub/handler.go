package stub

import (
	"context"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/paulczar/gcp-operator/pkg/apis/cloud/v1alpha1"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.Instance:
		//err := sdk.Create(newInstance(o))
		if event.Deleted {
			return deleteInstance(o)
		}
		err := newInstance(o)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("failed to create instance : %v", err)
			return err
		}
	}
	return nil
}
