package gce

import (
	compute "google.golang.org/api/compute/v1"
)

type ImageService struct {
	GCE     *GCEClient
	Payload *compute.Image
}

func NewImageService(project string, image *compute.Image) (*ImageService, error) {
	is, err := New(project)
	if err != nil {
		return nil, err
	}
	return &ImageService{
		GCE:     is,
		Payload: image,
	}, nil
}

// Create an image.
func (is *ImageService) Create() error {
	op, err := is.GCE.service.Images.Insert(is.GCE.projectID, is.Payload).Do()
	if err != nil {
		return err
	}
	if err = is.GCE.waitForGlobalOp(op); err != nil {
		return err
	}
	return nil
}

// Get an Image
func (is *ImageService) Get() (*compute.Image, error) {
	image, err := is.GCE.service.Images.Get(is.GCE.projectID, is.Payload.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return image, nil
}

// Delete an image
func (is *ImageService) Delete() error {
	op, err := is.GCE.service.Images.Delete(is.GCE.projectID, is.Payload.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil
		}
		return err
	}
	if err = is.GCE.waitForGlobalOp(op); err != nil {
		return err
	}
	return nil
}

// Update an image
// currently do not support updating an image
func (is *ImageService) Update() error {
	return nil
}
