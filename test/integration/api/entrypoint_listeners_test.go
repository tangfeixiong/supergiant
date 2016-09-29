package api

import (
	"testing"

	"github.com/supergiant/supergiant/pkg/core"
	"github.com/supergiant/supergiant/pkg/core/fake"
	"github.com/supergiant/supergiant/pkg/model"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEntrypointListenersList(t *testing.T) {
	srv := newTestServer()
	go srv.Start()
	defer srv.Stop()

	Convey("EntrypointListeners List works correctly", t, func() {

		table := []struct {
			// Input
			parentCloudAccount *model.CloudAccount
			parentKube         *model.Kube
			parentEntrypoint   *model.Entrypoint
			existingModels     []*model.EntrypointListener
			// Expectations
			err *model.Error
		}{
			// A successful example
			{
				parentCloudAccount: &model.CloudAccount{
					Name:        "test",
					Provider:    "aws",
					Credentials: map[string]string{"test": "test"},
				},
				parentKube: &model.Kube{
					CloudAccountName: "test",
					Name:             "test",
					MasterNodeSize:   "t2.micro",
					NodeSizes:        []string{"t2.micro"},
					AWSConfig: &model.AWSKubeConfig{
						Region:           "us-east-1",
						AvailabilityZone: "us-east-1a",
					},
				},
				parentEntrypoint: &model.Entrypoint{
					KubeName: "test",
					Name:     "test",
				},
				existingModels: []*model.EntrypointListener{
					{
						EntrypointName: "test",
						Name:           "port-test",
						EntrypointPort: 80,
						NodePort:       30303,
					},
				},
				err: nil,
			},
		}

		for _, item := range table {

			wipeAndInitialize(srv.Core)

			srv.Core.AWSProvider = func(_ map[string]string) core.Provider {
				return new(fake.Provider)
			}

			requestor := createAdmin(srv.Core)
			sg := srv.Core.NewAPIClient("token", requestor.APIToken)

			srv.Core.CloudAccounts.Create(item.parentCloudAccount)
			srv.Core.Kubes.Create(item.parentKube)
			srv.Core.Entrypoints.Create(item.parentEntrypoint)

			for _, existingModel := range item.existingModels {
				srv.Core.EntrypointListeners.Create(existingModel)
			}

			list := new(model.EntrypointListenerList)
			err := sg.EntrypointListeners.List(list)

			if item.err == nil {
				So(err, ShouldBeNil)
			} else {
				So(err, ShouldResemble, item.err)
			}

			So(len(list.Items), ShouldEqual, len(item.existingModels))
		}
	})
}

//------------------------------------------------------------------------------

func TestEntrypointListenersCreate(t *testing.T) {
	srv := newTestServer()
	go srv.Start()
	defer srv.Stop()

	Convey("EntrypointListeners Create works correctly", t, func() {

		table := []struct {
			// Input
			parentCloudAccount *model.CloudAccount
			parentKube         *model.Kube
			parentEntrypoint   *model.Entrypoint
			model              *model.EntrypointListener
			// Mocks
			mockCreateEntrypointListenerError error
			// Expectations
			err *model.Error
		}{
			// A successful example
			{
				parentCloudAccount: &model.CloudAccount{
					Name:        "test",
					Provider:    "aws",
					Credentials: map[string]string{"test": "test"},
				},
				parentKube: &model.Kube{
					CloudAccountName: "test",
					Name:             "test",
					MasterNodeSize:   "t2.micro",
					NodeSizes:        []string{"t2.micro"},
					AWSConfig: &model.AWSKubeConfig{
						Region:           "us-east-1",
						AvailabilityZone: "us-east-1a",
					},
				},
				parentEntrypoint: &model.Entrypoint{
					KubeName: "test",
					Name:     "test",
				},
				model: &model.EntrypointListener{
					EntrypointName: "test",
					Name:           "port-test",
					EntrypointPort: 80,
					NodePort:       30303,
				},
				mockCreateEntrypointListenerError: nil,
				err: nil,
			},

			// Correctly sets defaults

			// No Entrypoint

			// Invalid Entrypoint

			// Invalid provider
			// {
			// 	model: &model.EntrypointListener{
			// 		Name:        "test",
			// 		Provider:    "nocloud",
			// 		Credentials: map[string]string{"access_key": "blah", "secret_key": "bleh"},
			// 	},
			// 	mockCreateEntrypointListenerError: nil,
			// 	err: &model.Error{Status: 422, Message: "Validation failed: Provider: regular expression mismatch"},
			// },

			// On Provider CreateEntrypointListener error

		}

		for _, item := range table {

			wipeAndInitialize(srv.Core)

			srv.Core.AWSProvider = func(_ map[string]string) core.Provider {
				return &fake.Provider{
					CreateEntrypointListenerFn: func(m *model.EntrypointListener) error {
						return item.mockCreateEntrypointListenerError
					},
				}
			}

			requestor := createAdmin(srv.Core)
			sg := srv.Core.NewAPIClient("token", requestor.APIToken)

			srv.Core.CloudAccounts.Create(item.parentCloudAccount)
			srv.Core.Kubes.Create(item.parentKube)
			srv.Core.Entrypoints.Create(item.parentEntrypoint)

			err := sg.EntrypointListeners.Create(item.model)

			if item.err == nil {
				So(err, ShouldBeNil)
			} else {
				So(err, ShouldResemble, item.err)
			}
		}
	})
}

// //------------------------------------------------------------------------------
//
// func TestCloudAccountsGet(t *testing.T) {
// 	srv := newTestServer()
// 	go srv.Start()
// 	defer srv.Stop()
//
// 	Convey("CloudAccounts Get works correctly", t, func() {
//
// 		table := []struct {
// 			// Input
// 			existingModel *model.CloudAccount
// 			// Expectations
// 			err *model.Error
// 		}{
// 			// A successful example
// 			{
// 				existingModel: &model.CloudAccount{
// 					Name:        "test",
// 					Provider:    "aws",
// 					Credentials: map[string]string{"access_key": "blah", "secret_key": "bleh"},
// 				},
// 				err: nil,
// 			},
// 		}
//
// 		for _, item := range table {
//
// 			wipeAndInitialize(srv.Core)
//
// 			// For ValidateAccount on Create
// 			srv.Core.AWSProvider = func(_ map[string]string) core.Provider {
// 				return new(fake.Provider)
// 			}
//
// 			requestor := createAdmin(srv.Core)
// 			sg := srv.Core.NewAPIClient("token", requestor.APIToken)
//
// 			srv.Core.CloudAccounts.Create(item.existingModel)
//
// 			err := sg.CloudAccounts.Get(item.existingModel.ID, item.existingModel)
//
// 			if item.err == nil {
// 				So(err, ShouldBeNil)
// 			} else {
// 				So(err, ShouldResemble, item.err)
// 			}
// 		}
// 	})
// }
//
// //------------------------------------------------------------------------------
//
// func TestCloudAccountsUpdate(t *testing.T) {
// 	srv := newTestServer()
// 	go srv.Start()
// 	defer srv.Stop()
//
// 	Convey("CloudAccounts Update works correctly", t, func() {
//
// 		table := []struct {
// 			// Input
// 			existingModel *model.CloudAccount
// 			modelUpdate   *model.CloudAccount
// 			// Expectations
// 			err *model.Error
// 		}{
// 			// Can't update Name
// 			{
// 				existingModel: &model.CloudAccount{
// 					Name:        "test",
// 					Provider:    "aws",
// 					Credentials: map[string]string{"access_key": "blah", "secret_key": "bleh"},
// 				},
// 				modelUpdate: &model.CloudAccount{
// 					Name: "new-name",
// 				},
// 				err: &model.Error{Status: 422, Message: "Name cannot be changed"},
// 			},
//
// 			// Can't update Provider
// 			{
// 				existingModel: &model.CloudAccount{
// 					Name:        "test",
// 					Provider:    "aws",
// 					Credentials: map[string]string{"access_key": "blah", "secret_key": "bleh"},
// 				},
// 				modelUpdate: &model.CloudAccount{
// 					Provider: "do",
// 				},
// 				err: &model.Error{Status: 422, Message: "Provider cannot be changed"},
// 			},
//
// 			// Can't update Credentials
// 			{
// 				existingModel: &model.CloudAccount{
// 					Name:        "test",
// 					Provider:    "aws",
// 					Credentials: map[string]string{"access_key": "blah", "secret_key": "bleh"},
// 				},
// 				modelUpdate: &model.CloudAccount{
// 					Credentials: map[string]string{"new": "credz"},
// 				},
// 				err: &model.Error{Status: 422, Message: "Credentials cannot be changed"},
// 			},
// 		}
//
// 		for _, item := range table {
//
// 			wipeAndInitialize(srv.Core)
//
// 			// For Create
// 			srv.Core.AWSProvider = func(_ map[string]string) core.Provider {
// 				return new(fake.Provider)
// 			}
//
// 			requestor := createAdmin(srv.Core)
// 			sg := srv.Core.NewAPIClient("token", requestor.APIToken)
//
// 			srv.Core.CloudAccounts.Create(item.existingModel)
//
// 			err := sg.CloudAccounts.Update(item.existingModel.ID, item.modelUpdate)
//
// 			if item.err == nil {
// 				So(err, ShouldBeNil)
// 			} else {
// 				So(err, ShouldResemble, item.err)
// 			}
// 		}
// 	})
// }
//
// //------------------------------------------------------------------------------
//
// func TestCloudAccountsDelete(t *testing.T) {
// 	srv := newTestServer()
// 	go srv.Start()
// 	defer srv.Stop()
//
// 	Convey("CloudAccounts Delete works correctly", t, func() {
//
// 		table := []struct {
// 			// Input
// 			existingModel    *model.CloudAccount
// 			hasExistingKubes []*model.Kube
// 			// Expectations
// 			err *model.Error
// 		}{
// 			// Successful example
// 			{
// 				existingModel: &model.CloudAccount{
// 					Name:        "test",
// 					Provider:    "aws",
// 					Credentials: map[string]string{"access_key": "blah", "secret_key": "bleh"},
// 				},
// 				hasExistingKubes: nil,
// 				err:              nil,
// 			},
//
// 			// Can't delete if there's Kubes (cuz then we couldn't tear down)
// 			{
// 				existingModel: &model.CloudAccount{
// 					Name:        "test",
// 					Provider:    "aws",
// 					Credentials: map[string]string{"access_key": "blah", "secret_key": "bleh"},
// 				},
// 				hasExistingKubes: []*model.Kube{
// 					{
// 						CloudAccountName: "test",
// 						Name:             "testkube",
// 						MasterNodeSize:   "t2.micro",
// 						NodeSizes:        []string{"t2.micro"},
// 						AWSConfig: &model.AWSKubeConfig{
// 							Region:           "us-east-1",
// 							AvailabilityZone: "us-east-1a",
// 						},
// 					},
// 				},
// 				err: &model.Error{Status: 422, Message: "Validation failed: Cannot delete CloudAccount that has active Kubes"},
// 			},
// 		}
//
// 		for _, item := range table {
//
// 			wipeAndInitialize(srv.Core)
//
// 			// For Create
// 			srv.Core.AWSProvider = func(_ map[string]string) core.Provider {
// 				return new(fake.Provider)
// 			}
//
// 			requestor := createAdmin(srv.Core)
// 			sg := srv.Core.NewAPIClient("token", requestor.APIToken)
//
// 			srv.Core.CloudAccounts.Create(item.existingModel)
//
// 			for _, existingKube := range item.hasExistingKubes {
// 				srv.Core.Kubes.Create(existingKube)
// 			}
//
// 			err := sg.CloudAccounts.Delete(item.existingModel.ID, item.existingModel)
//
// 			if item.err == nil {
// 				So(err, ShouldBeNil)
// 			} else {
// 				So(err, ShouldResemble, item.err)
// 			}
// 		}
// 	})
// }
