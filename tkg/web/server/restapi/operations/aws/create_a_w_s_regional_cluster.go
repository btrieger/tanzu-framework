// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateAWSRegionalClusterHandlerFunc turns a function with the right signature into a create a w s regional cluster handler
type CreateAWSRegionalClusterHandlerFunc func(CreateAWSRegionalClusterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAWSRegionalClusterHandlerFunc) Handle(params CreateAWSRegionalClusterParams) middleware.Responder {
	return fn(params)
}

// CreateAWSRegionalClusterHandler interface for that can handle valid create a w s regional cluster params
type CreateAWSRegionalClusterHandler interface {
	Handle(CreateAWSRegionalClusterParams) middleware.Responder
}

// NewCreateAWSRegionalCluster creates a new http.Handler for the create a w s regional cluster operation
func NewCreateAWSRegionalCluster(ctx *middleware.Context, handler CreateAWSRegionalClusterHandler) *CreateAWSRegionalCluster {
	return &CreateAWSRegionalCluster{Context: ctx, Handler: handler}
}

/*
CreateAWSRegionalCluster swagger:route POST /api/providers/aws/create aws createAWSRegionalCluster

Create AWS regional cluster
*/
type CreateAWSRegionalCluster struct {
	Context *middleware.Context
	Handler CreateAWSRegionalClusterHandler
}

func (o *CreateAWSRegionalCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAWSRegionalClusterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
