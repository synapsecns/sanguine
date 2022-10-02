package collection

import (
	"fmt"
	"github.com/barkimedes/go-deepcopy"
	"github.com/rbretecher/go-postman-collection"
)

// DeepCopyItems makes a deep copy of an item struct
func DeepCopyItems(og *postman.Items) (*postman.Items, error) {
	behavior := deepcopy.MustAnything(og.ProtocolProfileBehavior)
	variables := deepcopy.MustAnything(og.Variables)
	events := deepcopy.MustAnything(og.Events)
	responses := deepcopy.MustAnything(og.Responses)
	auth := deepcopy.MustAnything(og.Auth)

	request, err := copyRequest(og.Request)
	if err != nil {
		return nil, fmt.Errorf("could not copy request: %w", err)
	}

	newItems := &postman.Items{
		Name:                    og.Name,
		Description:             og.Description,
		Variables:               variables.([]*postman.Variable),
		Events:                  events.([]*postman.Event),
		ProtocolProfileBehavior: behavior,
		ID:                      og.ID,
		Request:                 request,
		Responses:               responses.([]*postman.Response),
		Auth:                    auth.(*postman.Auth),
		Items:                   make([]*postman.Items, len(og.Items)),
	}

	for i, item := range og.Items {
		newItems.Items[i], err = DeepCopyItems(item)
		if err != nil {
			return nil, err
		}
	}

	return newItems, nil
}

func copyRequest(req *postman.Request) (res *postman.Request, err error) {
	if req == nil {
		return res, nil
	}

	requestJSON, err := req.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("could not marshall request: %w", err)
	}

	res = &postman.Request{}

	err = res.UnmarshalJSON(requestJSON)
	return res, nil
}
